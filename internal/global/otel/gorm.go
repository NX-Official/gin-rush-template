package otel

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"
	"github.com/uptrace/opentelemetry-go-extra/otelsql"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	semconv "go.opentelemetry.io/otel/semconv/v1.25.0"
	"go.opentelemetry.io/otel/trace"
	"gorm.io/gorm"
	"io"
)

type gormRegister interface {
	Register(name string, fn func(*gorm.DB)) error
}

type gormPlugin struct {
}

func GetGormPlugin() gorm.Plugin {
	return &gormPlugin{}
}

func (p *gormPlugin) Name() string {
	return "otelgorm"
}

func (p *gormPlugin) Initialize(db *gorm.DB) error {
	if db, ok := db.ConnPool.(*sql.DB); ok {
		otelsql.ReportDBStatsMetrics(db)
	}
	cb := db.Callback()
	hooks := []struct {
		callback gormRegister
		hook     func(tx *gorm.DB)
		name     string
	}{
		{cb.Create().Before("gorm:create"), p.before("gorm.Create"), "before:create"},
		{cb.Create().After("gorm:create"), p.after(), "after:create"},

		{cb.Query().Before("gorm:query"), p.before("gorm.Query"), "before:select"},
		{cb.Query().After("gorm:query"), p.after(), "after:select"},

		{cb.Delete().Before("gorm:delete"), p.before("gorm.Delete"), "before:delete"},
		{cb.Delete().After("gorm:delete"), p.after(), "after:delete"},

		{cb.Update().Before("gorm:update"), p.before("gorm.Update"), "before:update"},
		{cb.Update().After("gorm:update"), p.after(), "after:update"},

		{cb.Row().Before("gorm:row"), p.before("gorm.Row"), "before:row"},
		{cb.Row().After("gorm:row"), p.after(), "after:row"},

		{cb.Raw().Before("gorm:raw"), p.before("gorm.Raw"), "before:raw"},
		{cb.Raw().After("gorm:raw"), p.after(), "after:raw"},
	}

	var firstErr error

	for _, h := range hooks {
		if err := h.callback.Register("otel:"+h.name, h.hook); err != nil && firstErr == nil {
			firstErr = fmt.Errorf("callback register %s failed: %w", h.name, err)
		}
	}

	return firstErr
}

type parentCtxKey struct{}

func (p *gormPlugin) before(spanName string) func(tx *gorm.DB) {
	return func(tx *gorm.DB) {
		ctx := tx.Statement.Context
		ctx = context.WithValue(ctx, parentCtxKey{}, ctx)
		ctx, _ = trace.
			SpanFromContext(ctx).
			TracerProvider().
			Tracer("gorm").
			Start(ctx, spanName, trace.WithSpanKind(trace.SpanKindClient))
		tx.Statement.Context = ctx
	}
}

func (p *gormPlugin) after() func(tx *gorm.DB) {
	return func(tx *gorm.DB) {
		span := trace.SpanFromContext(tx.Statement.Context)
		if !span.IsRecording() {
			return
		}
		defer span.End()

		span.SetAttributes(
			dbSystem(tx),
			semconv.DBStatement(tx.Dialector.Explain(tx.Statement.SQL.String(), tx.Statement.Vars...)),
			semconv.DBSQLTable(tx.Statement.Table),
			attribute.Key("db.rows_affected").Int64(tx.Statement.RowsAffected),
		)
		switch {
		case tx.Error == nil,
			tx.Error == io.EOF,
			errors.Is(tx.Error, gorm.ErrRecordNotFound),
			errors.Is(tx.Error, driver.ErrSkip),
			errors.Is(tx.Error, sql.ErrNoRows):
			span.SetStatus(codes.Ok, "")
		default:
			span.RecordError(tx.Error)
			span.SetStatus(codes.Error, tx.Error.Error())
		}

		switch parentCtx := tx.Statement.Context.Value(parentCtxKey{}).(type) {
		case context.Context:
			tx.Statement.Context = parentCtx
		}
	}
}

func dbSystem(tx *gorm.DB) attribute.KeyValue {
	switch tx.Dialector.Name() {
	case "mysql":
		return semconv.DBSystemMySQL
	case "mssql":
		return semconv.DBSystemMSSQL
	case "postgres", "postgresql":
		return semconv.DBSystemPostgreSQL
	case "sqlite":
		return semconv.DBSystemSqlite
	case "sqlserver":
		return semconv.DBSystemKey.String("sqlserver")
	case "clickhouse":
		return semconv.DBSystemKey.String("clickhouse")
	default:
		return attribute.KeyValue{}
	}
}

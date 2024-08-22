package otel

import (
	"context"
	"fmt"
	"gin-rush-template/config"
	"gin-rush-template/tools"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	oteltrace "go.opentelemetry.io/otel/sdk/trace"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
)

// List of supported exporters
// https://opentelemetry.io/docs/instrumentation/go/exporters/

// OTLP Exporter
func newOTLPExporter(ctx context.Context) (oteltrace.SpanExporter, error) {
	// Change default HTTPS -> HTTP
	insecureOpt := otlptracehttp.WithInsecure()

	// Update default OTLP reciver endpoint
	endPoint := fmt.Sprintf("%s:%s", config.Get().OTel.AgentHost, config.Get().OTel.AgentPort)
	endpointOpt := otlptracehttp.WithEndpoint(endPoint)
	return otlptracehttp.New(ctx, insecureOpt, endpointOpt)
}

func InitTracer() {
	r, err := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(config.Get().OTel.ServiceName),
		),
	)
	tools.PanicOnErr(err)

	provider := sdktrace.NewTracerProvider(sdktrace.WithResource(r))
	otel.SetTracerProvider(provider)
	exp, err := newOTLPExporter(context.Background())
	tools.PanicOnErr(err)

	bsp := sdktrace.NewBatchSpanProcessor(exp)
	provider.RegisterSpanProcessor(bsp)
}

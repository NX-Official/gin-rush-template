package otel

import (
	"context"
	"gin-rush-template/config"
	"gin-rush-template/tools"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/prometheus"
	metr "go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
)

func newPrometheusExporter() (*prometheus.Exporter, error) {
	uOps := prometheus.WithoutUnits()
	namespace := prometheus.WithNamespace("gin_rush_template")
	return prometheus.New(uOps, namespace)
}

func InitPrometheus() {
	r, err := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(config.Get().OTel.ServiceName),
		),
	)
	tools.PanicOnErr(err)

	exporter, err := newPrometheusExporter()
	tools.PanicOnErr(err)

	provider := metric.NewMeterProvider(
		metric.WithResource(r),
		metric.WithReader(exporter),
	)

	otel.SetMeterProvider(provider)
}

func CustomMetrics() {
	ctx := context.Background()
	meter := otel.Meter(config.Get().OTel.ServiceName)

	opt := metr.WithAttributes(
		attribute.Key("ping").String("pong"),
		attribute.Key("version").String("v4"),
	)

	// create a counter 计数器
	counter, err := meter.Float64Counter("foo", metr.WithDescription("a simple counter"))
	if err != nil {
		panic(err)
	}
	counter.Add(ctx, 5, opt)

	// Create a gauge 测量仪表
	gauge, err := meter.Float64ObservableGauge("bar", metr.WithDescription("a fun little gauge"))
	if err != nil {
		panic(err)
	}
	_, err = meter.RegisterCallback(func(_ context.Context, o metr.Observer) error {
		n := 42.0
		o.ObserveFloat64(gauge, n, opt)
		return nil
	}, gauge)
	if err != nil {
		panic(err)
	}

	//  create a histogram 直方图
	histogram, err := meter.Float64Histogram(
		"baz",
		metr.WithDescription("a histogram with custom buckets and rename"),
		metr.WithExplicitBucketBoundaries(64, 128, 256, 512, 1024, 2048, 4096),
	)
	if err != nil {
		panic(err)
	}
	histogram.Record(ctx, 136, opt)
	histogram.Record(ctx, 64, opt)
	histogram.Record(ctx, 701, opt)
	histogram.Record(ctx, 830, opt)

}

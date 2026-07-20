package observops

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"go.opentelemetry.io/contrib/bridges/otelslog"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploggrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploghttp"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/exporters/stdout/stdoutlog"
	"go.opentelemetry.io/otel/exporters/stdout/stdoutmetric"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	otellog "go.opentelemetry.io/otel/log/global"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/propagation"
	sdklog "go.opentelemetry.io/otel/sdk/log"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
	"go.opentelemetry.io/otel/trace"
)

// Telemetry holds initialized OpenTelemetry handles for direct use with the standard OTel
// API. Setup wires every enabled signal and their exporters; callers then instrument with
// vanilla OTel (trace.Tracer, metric.Meter, slog) and retain full ecosystem interop.
type Telemetry struct {
	// Tracer is a ready-to-use tracer named after the service.
	Tracer trace.Tracer
	// Meter is a ready-to-use meter named after the service.
	Meter metric.Meter
	// Logger is an slog.Logger fanning out to the console and, when logs are enabled, to
	// the OpenTelemetry log pipeline. It is also installed as slog.Default().
	Logger *slog.Logger
	// MetricsHandler serves Prometheus metrics when WithPrometheus is set, else nil.
	MetricsHandler http.Handler

	// Providers are exposed for advanced use (custom instruments, registration).
	TracerProvider *sdktrace.TracerProvider
	MeterProvider  *sdkmetric.MeterProvider
	LoggerProvider *sdklog.LoggerProvider

	shutdowns []func(context.Context) error
}

// Setup bootstraps OpenTelemetry from options and returns native handles. Metrics, traces,
// and logs are enabled by default; disable any with WithMetrics/WithTraces/WithLogs(false).
// Push export to an OTLP collector is active when WithEndpoint is set; WithPrometheus adds a
// pull endpoint; WithStdout mirrors to stdout for debugging.
func Setup(ctx context.Context, opts ...ClientOption) (*Telemetry, error) {
	cfg := ApplyOptions(append([]ClientOption{setupDefaults()}, opts...)...)

	name := cfg.ServiceName
	if name == "" {
		name = "service"
	}

	res, err := buildSetupResource(cfg)
	if err != nil {
		return nil, err
	}

	t := &Telemetry{}
	otlp := cfg.Endpoint != ""

	if err := t.setupMetrics(ctx, cfg, res, otlp); err != nil {
		return nil, t.abort(ctx, err)
	}
	if err := t.setupTraces(ctx, cfg, res, otlp); err != nil {
		return nil, t.abort(ctx, err)
	}
	if err := t.setupLogs(ctx, cfg, res, otlp, name); err != nil {
		return nil, t.abort(ctx, err)
	}

	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	))

	t.Tracer = otel.Tracer(name)
	t.Meter = otel.Meter(name)
	return t, nil
}

// setupDefaults seeds the signal toggles that default to enabled.
func setupDefaults() ClientOption {
	return func(c *Config) {
		c.EnableMetrics = true
		c.EnableTraces = true
		c.EnableLogs = true
		c.TraceSampleRatio = 1.0
	}
}

func buildSetupResource(cfg *Config) (*resource.Resource, error) {
	attrs := []attribute.KeyValue{}
	if cfg.ServiceName != "" {
		attrs = append(attrs, semconv.ServiceNameKey.String(cfg.ServiceName))
	}
	if cfg.ServiceVersion != "" {
		attrs = append(attrs, semconv.ServiceVersionKey.String(cfg.ServiceVersion))
	}
	if cfg.Resource != nil {
		if cfg.Resource.ServiceNamespace != "" {
			attrs = append(attrs, semconv.ServiceNamespaceKey.String(cfg.Resource.ServiceNamespace))
		}
		if cfg.Resource.DeploymentEnv != "" {
			attrs = append(attrs, attribute.String("deployment.environment", cfg.Resource.DeploymentEnv))
		}
		for k, v := range cfg.Resource.Attributes {
			attrs = append(attrs, attribute.String(k, v))
		}
	}

	// Merge with the default (process/host/runtime) resource so telemetry is
	// self-describing. Our attributes are schemaless so the merge adopts the default
	// resource's schema URL rather than conflicting with it.
	return resource.Merge(
		resource.Default(),
		resource.NewSchemaless(attrs...),
	)
}

func (t *Telemetry) setupMetrics(ctx context.Context, cfg *Config, res *resource.Resource, otlp bool) error {
	if !cfg.EnableMetrics {
		return nil
	}

	readers := []sdkmetric.Option{sdkmetric.WithResource(res)}

	if cfg.EnablePrometheus {
		exp, err := prometheus.New()
		if err != nil {
			return fmt.Errorf("prometheus exporter: %w", err)
		}
		readers = append(readers, sdkmetric.WithReader(exp))
		t.MetricsHandler = promHandler()
	}

	if otlp {
		exp, err := newOTLPMetricExporter(ctx, cfg)
		if err != nil {
			return fmt.Errorf("otlp metric exporter: %w", err)
		}
		readers = append(readers, sdkmetric.WithReader(sdkmetric.NewPeriodicReader(exp)))
	}

	if cfg.EnableStdout {
		exp, err := stdoutmetric.New()
		if err != nil {
			return fmt.Errorf("stdout metric exporter: %w", err)
		}
		readers = append(readers, sdkmetric.WithReader(sdkmetric.NewPeriodicReader(exp)))
	}

	t.MeterProvider = sdkmetric.NewMeterProvider(readers...)
	otel.SetMeterProvider(t.MeterProvider)
	t.addShutdown(t.MeterProvider.Shutdown)
	return nil
}

func (t *Telemetry) setupTraces(ctx context.Context, cfg *Config, res *resource.Resource, otlp bool) error {
	if !cfg.EnableTraces {
		return nil
	}

	ratio := cfg.TraceSampleRatio
	if ratio <= 0 {
		ratio = 1.0
	}

	opts := []sdktrace.TracerProviderOption{
		sdktrace.WithResource(res),
		sdktrace.WithSampler(sdktrace.ParentBased(sdktrace.TraceIDRatioBased(ratio))),
	}

	if otlp {
		exp, err := newOTLPTraceExporter(ctx, cfg)
		if err != nil {
			return fmt.Errorf("otlp trace exporter: %w", err)
		}
		opts = append(opts, sdktrace.WithBatcher(exp))
	}

	if cfg.EnableStdout {
		exp, err := stdouttrace.New()
		if err != nil {
			return fmt.Errorf("stdout trace exporter: %w", err)
		}
		opts = append(opts, sdktrace.WithBatcher(exp))
	}

	t.TracerProvider = sdktrace.NewTracerProvider(opts...)
	otel.SetTracerProvider(t.TracerProvider)
	t.addShutdown(t.TracerProvider.Shutdown)
	return nil
}

func (t *Telemetry) setupLogs(ctx context.Context, cfg *Config, res *resource.Resource, otlp bool, name string) error {
	// The console handler is always present; the OTel bridge is added when a log exporter
	// exists. The result is installed both on Telemetry.Logger and as slog.Default().
	handlers := []slog.Handler{consoleSlogHandler(cfg.Debug)}

	if cfg.EnableLogs && (otlp || cfg.EnableStdout) {
		var processors []sdklog.LoggerProviderOption
		processors = append(processors, sdklog.WithResource(res))

		if otlp {
			exp, err := newOTLPLogExporter(ctx, cfg)
			if err != nil {
				return fmt.Errorf("otlp log exporter: %w", err)
			}
			processors = append(processors, sdklog.WithProcessor(sdklog.NewBatchProcessor(exp)))
		}
		if cfg.EnableStdout {
			exp, err := stdoutlog.New()
			if err != nil {
				return fmt.Errorf("stdout log exporter: %w", err)
			}
			processors = append(processors, sdklog.WithProcessor(sdklog.NewBatchProcessor(exp)))
		}

		t.LoggerProvider = sdklog.NewLoggerProvider(processors...)
		otellog.SetLoggerProvider(t.LoggerProvider)
		t.addShutdown(t.LoggerProvider.Shutdown)
		handlers = append(handlers, otelslog.NewHandler(name, otelslog.WithLoggerProvider(t.LoggerProvider)))
	}

	logger := slog.New(newSetupFanout(handlers...))
	t.Logger = logger
	slog.SetDefault(logger)
	return nil
}

// --- exporter constructors (gRPC or HTTP) ---

func newOTLPMetricExporter(ctx context.Context, cfg *Config) (sdkmetric.Exporter, error) {
	if cfg.OverHTTP {
		opts := []otlpmetrichttp.Option{otlpmetrichttp.WithEndpoint(cfg.Endpoint)}
		if cfg.Insecure {
			opts = append(opts, otlpmetrichttp.WithInsecure())
		}
		if len(cfg.Headers) > 0 {
			opts = append(opts, otlpmetrichttp.WithHeaders(cfg.Headers))
		}
		return otlpmetrichttp.New(ctx, opts...)
	}
	opts := []otlpmetricgrpc.Option{otlpmetricgrpc.WithEndpoint(cfg.Endpoint)}
	if cfg.Insecure {
		opts = append(opts, otlpmetricgrpc.WithInsecure())
	}
	if len(cfg.Headers) > 0 {
		opts = append(opts, otlpmetricgrpc.WithHeaders(cfg.Headers))
	}
	return otlpmetricgrpc.New(ctx, opts...)
}

func newOTLPTraceExporter(ctx context.Context, cfg *Config) (sdktrace.SpanExporter, error) {
	if cfg.OverHTTP {
		opts := []otlptracehttp.Option{otlptracehttp.WithEndpoint(cfg.Endpoint)}
		if cfg.Insecure {
			opts = append(opts, otlptracehttp.WithInsecure())
		}
		if len(cfg.Headers) > 0 {
			opts = append(opts, otlptracehttp.WithHeaders(cfg.Headers))
		}
		return otlptracehttp.New(ctx, opts...)
	}
	opts := []otlptracegrpc.Option{otlptracegrpc.WithEndpoint(cfg.Endpoint)}
	if cfg.Insecure {
		opts = append(opts, otlptracegrpc.WithInsecure())
	}
	if len(cfg.Headers) > 0 {
		opts = append(opts, otlptracegrpc.WithHeaders(cfg.Headers))
	}
	return otlptracegrpc.New(ctx, opts...)
}

func newOTLPLogExporter(ctx context.Context, cfg *Config) (sdklog.Exporter, error) {
	if cfg.OverHTTP {
		opts := []otlploghttp.Option{otlploghttp.WithEndpoint(cfg.Endpoint)}
		if cfg.Insecure {
			opts = append(opts, otlploghttp.WithInsecure())
		}
		if len(cfg.Headers) > 0 {
			opts = append(opts, otlploghttp.WithHeaders(cfg.Headers))
		}
		return otlploghttp.New(ctx, opts...)
	}
	opts := []otlploggrpc.Option{otlploggrpc.WithEndpoint(cfg.Endpoint)}
	if cfg.Insecure {
		opts = append(opts, otlploggrpc.WithInsecure())
	}
	if len(cfg.Headers) > 0 {
		opts = append(opts, otlploggrpc.WithHeaders(cfg.Headers))
	}
	return otlploggrpc.New(ctx, opts...)
}

func consoleSlogHandler(debug bool) slog.Handler {
	level := slog.LevelInfo
	if debug {
		level = slog.LevelDebug
	}
	return slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: level})
}

func (t *Telemetry) addShutdown(fn func(context.Context) error) {
	t.shutdowns = append(t.shutdowns, fn)
}

// abort shuts down whatever was already initialized when a later step fails, and returns the
// original error.
func (t *Telemetry) abort(ctx context.Context, err error) error {
	_ = t.Shutdown(ctx)
	return err
}

// Shutdown flushes and stops every initialized signal provider in reverse order. It returns
// the first error but always attempts them all.
func (t *Telemetry) Shutdown(ctx context.Context) error {
	var firstErr error
	for i := len(t.shutdowns) - 1; i >= 0; i-- {
		if err := t.shutdowns[i](ctx); err != nil && firstErr == nil {
			firstErr = err
		}
	}
	return firstErr
}

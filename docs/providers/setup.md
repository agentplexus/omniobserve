# Setup (recommended)

`observops.Setup` is the recommended way to add observability to a service. One call wires
every enabled signal — metrics, traces, and logs — and returns **native OpenTelemetry
handles**. You instrument with the standard OTel API and keep full ecosystem interop
(otelhttp, gRPC interceptors, database instrumentation); `observops` owns only the wiring.

## Why Setup instead of a driver

The driver registry (`Open("otlp")`, `Open("datadog")`, …) wraps the OpenTelemetry API behind
its own interfaces. That is a lossy subset and cuts you off from the OTel contrib ecosystem.
`Setup` abstracts the *wiring* instead of the *API*, so you get one-line configuration **and**
the full power of OpenTelemetry.

## Quick start

```go
import "github.com/plexusone/omniobserve/observops"

tel, err := observops.Setup(ctx,
    observops.WithServiceName("my-service"),
    observops.WithServiceVersion("1.2.3"),
    observops.WithPrometheus(),               // pull: expose tel.MetricsHandler
    observops.WithEndpoint("localhost:4317"), // push: OTLP gRPC collector
    observops.WithInsecure(),
)
if err != nil {
    log.Fatal(err)
}
defer tel.Shutdown(context.Background())
```

`Setup` returns a `*Telemetry`:

| Field | Type | Purpose |
|-------|------|---------|
| `Tracer` | `trace.Tracer` | Native OTel tracer, named after the service |
| `Meter` | `metric.Meter` | Native OTel meter (int64/float64, sync/observable) |
| `Logger` | `*slog.Logger` | Fans out to the console and, when logs are enabled, to OTel; also installed as `slog.Default()` |
| `MetricsHandler` | `http.Handler` | Prometheus scrape handler (nil unless `WithPrometheus`) |
| `TracerProvider` / `MeterProvider` / `LoggerProvider` | SDK providers | Escape hatch for advanced use |

## Using the handles

```go
// Metrics — full native API, including observable instruments.
reqs, _ := tel.Meter.Int64Counter("requests.total")
reqs.Add(ctx, 1)

_, _ = tel.Meter.Int64ObservableGauge("queue.depth",
    metric.WithInt64Callback(func(_ context.Context, o metric.Int64Observer) error {
        o.Observe(currentDepth())
        return nil
    }))

// Tracing.
ctx, span := tel.Tracer.Start(ctx, "handle-request")
defer span.End()

// Logging (exported via OTel and echoed to the console).
tel.Logger.InfoContext(ctx, "request processed", "user_id", "123")

// Serve metrics and instrument HTTP.
mux.Handle("/metrics", tel.MetricsHandler)
handler = tel.Middleware("api")(handler)
client := &http.Client{Transport: tel.Transport(nil)}
```

## Signals and exporters

Metrics, traces, and logs are **enabled by default**. Toggle them, and compose export targets:

| Option | Effect |
|--------|--------|
| `WithMetrics(false)` / `WithTraces(false)` / `WithLogs(false)` | Disable a signal |
| `WithPrometheus()` | Add a Prometheus pull exporter; exposes `MetricsHandler` |
| `WithEndpoint("host:4317")` | Push all enabled signals to an OTLP collector |
| `WithOTLPOverHTTP()` | Use OTLP over HTTP (port 4318) instead of gRPC |
| `WithInsecure()` | Plaintext OTLP (local collectors) |
| `WithStdout()` | Mirror telemetry to stdout for debugging |
| `WithTraceSampleRatio(0.1)` | Head sampling ratio for traces |
| `WithHeaders(...)` | Auth headers for the collector |

Prometheus pull and OTLP push can be used together. Logs are exported through the real
OpenTelemetry log SDK (via the `otelslog` bridge), not as span events.

## Shutdown

`tel.Shutdown(ctx)` flushes and stops every provider in reverse order of initialization.
Always defer it so buffered telemetry is exported on exit.

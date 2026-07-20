# observops

General service observability (metrics, traces, logs) on OpenTelemetry, with two entry
points.

## Setup (recommended)

One call wires every signal and returns **native OpenTelemetry handles** — you instrument
with the standard OTel API and keep full ecosystem interop. `observops` owns only the wiring:
exporters, resource detection, propagation, batch processors, the slog bridge, and shutdown.

```go
tel, err := observops.Setup(ctx,
    observops.WithServiceName("my-service"),
    observops.WithPrometheus(),               // pull: tel.MetricsHandler
    observops.WithEndpoint("localhost:4317"), // push: OTLP gRPC
    observops.WithInsecure(),
)
if err != nil {
    log.Fatal(err)
}
defer tel.Shutdown(context.Background())

reqs, _ := tel.Meter.Int64Counter("requests.total")
reqs.Add(ctx, 1)

ctx, span := tel.Tracer.Start(ctx, "handle-request")
defer span.End()

tel.Logger.InfoContext(ctx, "request processed", "user_id", "123")

mux.Handle("/metrics", tel.MetricsHandler)
handler = tel.Middleware("api")(handler)
```

Metrics, traces, and logs are on by default; disable with `WithMetrics/WithTraces/WithLogs`.
Export targets compose: Prometheus pull (`WithPrometheus`), OTLP push over gRPC or HTTP
(`WithEndpoint`, `WithOTLPOverHTTP`), and stdout (`WithStdout`).

See [`docs/providers/setup.md`](../docs/providers/setup.md) for the full reference.

## Driver registry (vendor-neutral abstraction)

`Open(name)` returns a provider whose own interfaces wrap OpenTelemetry, for switching
vendors through one call. It exposes an OTel subset; prefer `Setup` unless you specifically
want the vendor switch.

```go
import _ "github.com/plexusone/omniobserve/observops/otlp"

provider, err := observops.Open("otlp",
    observops.WithEndpoint("localhost:4317"),
    observops.WithServiceName("my-service"),
)
defer provider.Shutdown(context.Background())
```

Registered drivers: `otlp`, `datadog`, `newrelic`, `dynatrace`.

package observops_test

import (
	"context"
	"log"
	"net/http"

	"github.com/plexusone/omniobserve/observops"
)

// ExampleSetup shows the recommended bootstrap: one call wires metrics, traces, and logs,
// and returns native OpenTelemetry handles plus a Prometheus handler and HTTP middleware.
func ExampleSetup() {
	ctx := context.Background()

	tel, err := observops.Setup(ctx,
		observops.WithServiceName("my-service"),
		observops.WithServiceVersion("1.2.3"),
		observops.WithPrometheus(),               // pull endpoint (tel.MetricsHandler)
		observops.WithEndpoint("localhost:4317"), // OTLP push
		observops.WithInsecure(),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = tel.Shutdown(ctx) }()

	// Metrics with the native OTel meter.
	requests, _ := tel.Meter.Int64Counter("requests.total")
	requests.Add(ctx, 1)

	// Tracing with the native OTel tracer.
	ctx, span := tel.Tracer.Start(ctx, "handle-request")
	defer span.End()

	// Structured logging exported via OpenTelemetry and echoed to the console.
	tel.Logger.InfoContext(ctx, "request processed", "user_id", "123")

	// Expose Prometheus metrics and instrument an HTTP handler.
	mux := http.NewServeMux()
	mux.Handle("/metrics", tel.MetricsHandler)
	mux.Handle("/api/", tel.Middleware("api")(http.NotFoundHandler()))
}

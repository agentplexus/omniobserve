package observops_test

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/plexusone/omniobserve/observops"
)

func TestSetupReturnsUsableHandles(t *testing.T) {
	ctx := context.Background()
	tel, err := observops.Setup(ctx, observops.WithServiceName("test-svc"))
	if err != nil {
		t.Fatalf("Setup: %v", err)
	}
	defer func() {
		if err := tel.Shutdown(ctx); err != nil {
			t.Errorf("Shutdown: %v", err)
		}
	}()

	if tel.Tracer == nil || tel.Meter == nil || tel.Logger == nil {
		t.Fatal("expected non-nil Tracer, Meter, Logger")
	}

	// The native handles must be usable without panicking, with no exporter configured.
	_, span := tel.Tracer.Start(ctx, "unit")
	span.End()

	counter, err := tel.Meter.Int64Counter("test.counter")
	if err != nil {
		t.Fatalf("Int64Counter: %v", err)
	}
	counter.Add(ctx, 1)

	tel.Logger.Info("hello from setup")
}

func TestSetupPrometheusHandlerServesMetrics(t *testing.T) {
	ctx := context.Background()
	tel, err := observops.Setup(ctx,
		observops.WithServiceName("prom-svc"),
		observops.WithPrometheus(),
		observops.WithTraces(false),
		observops.WithLogs(false),
	)
	if err != nil {
		t.Fatalf("Setup: %v", err)
	}
	defer func() { _ = tel.Shutdown(ctx) }()

	if tel.MetricsHandler == nil {
		t.Fatal("expected a Prometheus MetricsHandler")
	}

	// Record a metric, then scrape and confirm it is exported in Prometheus text format.
	counter, err := tel.Meter.Int64Counter("orders.total")
	if err != nil {
		t.Fatalf("counter: %v", err)
	}
	counter.Add(ctx, 3)

	srv := httptest.NewServer(tel.MetricsHandler)
	defer srv.Close()

	resp, err := http.Get(srv.URL)
	if err != nil {
		t.Fatalf("scrape: %v", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("read: %v", err)
	}
	if !strings.Contains(string(body), "orders_total") {
		t.Errorf("scrape output missing orders_total counter; got:\n%s", body)
	}
}

func TestSetupNoPrometheusHandlerWhenDisabled(t *testing.T) {
	ctx := context.Background()
	tel, err := observops.Setup(ctx, observops.WithServiceName("no-prom"))
	if err != nil {
		t.Fatalf("Setup: %v", err)
	}
	defer func() { _ = tel.Shutdown(ctx) }()

	if tel.MetricsHandler != nil {
		t.Error("expected nil MetricsHandler without WithPrometheus")
	}
}

func TestSetupDisableSignals(t *testing.T) {
	ctx := context.Background()
	tel, err := observops.Setup(ctx,
		observops.WithServiceName("minimal"),
		observops.WithMetrics(false),
		observops.WithTraces(false),
		observops.WithLogs(false),
	)
	if err != nil {
		t.Fatalf("Setup: %v", err)
	}
	defer func() { _ = tel.Shutdown(ctx) }()

	// Handles must still be safe to call even when the signals are disabled.
	_, span := tel.Tracer.Start(ctx, "noop")
	span.End()
	tel.Logger.Info("still safe")

	if tel.MeterProvider != nil {
		t.Error("expected no MeterProvider when metrics disabled")
	}
	if tel.TracerProvider != nil {
		t.Error("expected no TracerProvider when traces disabled")
	}
}

func TestSetupHTTPMiddlewareAndTransport(t *testing.T) {
	ctx := context.Background()
	tel, err := observops.Setup(ctx, observops.WithServiceName("http-svc"))
	if err != nil {
		t.Fatalf("Setup: %v", err)
	}
	defer func() { _ = tel.Shutdown(ctx) }()

	handler := tel.Middleware("test-server")(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	}))
	srv := httptest.NewServer(handler)
	defer srv.Close()

	client := &http.Client{Transport: tel.Transport(nil)}
	resp, err := client.Get(srv.URL)
	if err != nil {
		t.Fatalf("instrumented request: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusNoContent {
		t.Errorf("status = %d, want %d", resp.StatusCode, http.StatusNoContent)
	}
}

func TestSetupShutdownIsIdempotentEnough(t *testing.T) {
	ctx := context.Background()
	tel, err := observops.Setup(ctx, observops.WithServiceName("svc"), observops.WithPrometheus())
	if err != nil {
		t.Fatalf("Setup: %v", err)
	}
	if err := tel.Shutdown(ctx); err != nil {
		t.Fatalf("first shutdown: %v", err)
	}
	// A second shutdown must not panic (providers may return an error, which is tolerated).
	_ = tel.Shutdown(ctx)
}

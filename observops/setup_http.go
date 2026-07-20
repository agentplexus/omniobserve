package observops

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

// promHandler returns the Prometheus scrape handler for the default gatherer, which the
// OTel Prometheus exporter registers into.
func promHandler() http.Handler {
	return promhttp.Handler()
}

// Middleware wraps an http.Handler with OpenTelemetry server instrumentation: a span per
// request (with context extracted from inbound headers) plus HTTP server metrics. Use it on
// the serving side.
func (t *Telemetry) Middleware(operation string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return otelhttp.NewHandler(next, operation)
	}
}

// Transport wraps an http.RoundTripper with OpenTelemetry client instrumentation: a span per
// outbound request plus trace-context propagation to the upstream. Pass nil to wrap
// http.DefaultTransport.
func (t *Telemetry) Transport(base http.RoundTripper) http.RoundTripper {
	if base == nil {
		base = http.DefaultTransport
	}
	return otelhttp.NewTransport(base)
}

// setupFanout dispatches each log record to every underlying handler, so logs reach both the
// console and the OpenTelemetry log pipeline.
type setupFanout struct {
	handlers []slog.Handler
}

func newSetupFanout(handlers ...slog.Handler) *setupFanout {
	return &setupFanout{handlers: handlers}
}

func (f *setupFanout) Enabled(ctx context.Context, level slog.Level) bool {
	for _, h := range f.handlers {
		if h.Enabled(ctx, level) {
			return true
		}
	}
	return false
}

func (f *setupFanout) Handle(ctx context.Context, r slog.Record) error {
	var firstErr error
	for _, h := range f.handlers {
		if !h.Enabled(ctx, r.Level) {
			continue
		}
		if err := h.Handle(ctx, r.Clone()); err != nil && firstErr == nil {
			firstErr = err
		}
	}
	return firstErr
}

func (f *setupFanout) WithAttrs(attrs []slog.Attr) slog.Handler {
	next := make([]slog.Handler, len(f.handlers))
	for i, h := range f.handlers {
		next[i] = h.WithAttrs(attrs)
	}
	return &setupFanout{handlers: next}
}

func (f *setupFanout) WithGroup(name string) slog.Handler {
	next := make([]slog.Handler, len(f.handlers))
	for i, h := range f.handlers {
		next[i] = h.WithGroup(name)
	}
	return &setupFanout{handlers: next}
}

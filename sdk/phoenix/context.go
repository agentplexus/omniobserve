package phoenix

import "context"

// Context keys for storing trace and span data.
type contextKey int

const (
	traceContextKey contextKey = iota
	spanContextKey
	clientContextKey
)

// ContextWithTrace returns a new context with the trace attached.
func ContextWithTrace(ctx context.Context, trace *Trace) context.Context {
	return context.WithValue(ctx, traceContextKey, trace)
}

// TraceFromContext returns the trace from the context, or nil if none.
func TraceFromContext(ctx context.Context) *Trace {
	if trace, ok := ctx.Value(traceContextKey).(*Trace); ok {
		return trace
	}
	return nil
}

// ContextWithSpan returns a new context with the span attached.
func ContextWithSpan(ctx context.Context, span *Span) context.Context {
	return context.WithValue(ctx, spanContextKey, span)
}

// SpanFromContext returns the span from the context, or nil if none.
func SpanFromContext(ctx context.Context) *Span {
	if span, ok := ctx.Value(spanContextKey).(*Span); ok {
		return span
	}
	return nil
}

// ContextWithClient returns a new context with the client attached.
func ContextWithClient(ctx context.Context, client *Client) context.Context {
	return context.WithValue(ctx, clientContextKey, client)
}

// ClientFromContext returns the client from the context, or nil if none.
func ClientFromContext(ctx context.Context) *Client {
	if client, ok := ctx.Value(clientContextKey).(*Client); ok {
		return client
	}
	return nil
}

// StartSpan starts a span from the current context.
func StartSpan(ctx context.Context, name string, opts ...SpanOption) (context.Context, *Span, error) {
	// Try parent span first
	if parentSpan := SpanFromContext(ctx); parentSpan != nil {
		return parentSpan.Span(ctx, name, opts...)
	}

	// Try trace
	if trace := TraceFromContext(ctx); trace != nil {
		return trace.Span(ctx, name, opts...)
	}

	return ctx, nil, ErrNoActiveTrace
}

// StartLLMSpan starts an LLM span from the current context.
func StartLLMSpan(ctx context.Context, name string, opts ...SpanOption) (context.Context, *Span, error) {
	opts = append([]SpanOption{WithSpanKind(SpanKindLLM)}, opts...)
	return StartSpan(ctx, name, opts...)
}

// EndSpan ends the current span in context.
func EndSpan(ctx context.Context, opts ...SpanOption) error {
	span := SpanFromContext(ctx)
	if span == nil {
		return ErrNoActiveSpan
	}
	return span.End(ctx, opts...)
}

// EndTrace ends the current trace in context.
func EndTrace(ctx context.Context, opts ...TraceOption) error {
	trace := TraceFromContext(ctx)
	if trace == nil {
		return ErrNoActiveTrace
	}
	return trace.End(ctx, opts...)
}

// CurrentTraceID returns the current trace ID from context.
func CurrentTraceID(ctx context.Context) string {
	if trace := TraceFromContext(ctx); trace != nil {
		return trace.ID()
	}
	return ""
}

// CurrentSpanID returns the current span ID from context.
func CurrentSpanID(ctx context.Context) string {
	if span := SpanFromContext(ctx); span != nil {
		return span.ID()
	}
	return ""
}

package fluxllm

import (
	"context"

	"github.com/grokify/observai/llmops"
)

// contextKey is a private type used for storing spans in context.
type contextKey struct{}

// contextWithSpan returns a new context with the span attached.
func contextWithSpan(ctx context.Context, span llmops.Span) context.Context {
	return context.WithValue(ctx, contextKey{}, span)
}

// spanFromContext retrieves the span from the context.
// Returns nil if no span is attached.
func spanFromContext(ctx context.Context) llmops.Span {
	span, _ := ctx.Value(contextKey{}).(llmops.Span)
	return span
}

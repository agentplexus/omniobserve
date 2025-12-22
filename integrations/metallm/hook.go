// Package metallm provides an ObservabilityHook implementation for MetaLLM
// that integrates with ObservAI's llmops providers (Opik, Langfuse, Phoenix).
package metallm

import (
	"context"

	"github.com/grokify/metallm"
	"github.com/grokify/metallm/provider"

	"github.com/grokify/metaobserve/llmops"
)

// Hook implements metallm.ObservabilityHook using an llmops.Provider.
// It automatically creates spans for each LLM call with model, provider,
// input/output, and token usage information.
type Hook struct {
	provider llmops.Provider
}

// NewHook creates a new MetaLLM observability hook.
// The provider should be initialized before passing to this function.
func NewHook(provider llmops.Provider) *Hook {
	return &Hook{provider: provider}
}

// Ensure Hook implements the interface at compile time
var _ metallm.ObservabilityHook = (*Hook)(nil)

// BeforeRequest is called before each LLM call.
// It starts a new span and returns a context with the span attached.
func (h *Hook) BeforeRequest(ctx context.Context, info metallm.LLMCallInfo, req *provider.ChatCompletionRequest) context.Context {
	// Start a span for this LLM call
	ctx, span, err := h.provider.StartSpan(ctx, "llm-completion",
		llmops.WithSpanType(llmops.SpanTypeLLM),
		llmops.WithModel(req.Model),
		llmops.WithProvider(info.ProviderName),
		llmops.WithSpanInput(req.Messages),
	)
	if err != nil {
		// Don't fail the request if observability fails
		return ctx
	}

	// Store span in context for AfterResponse
	return contextWithSpan(ctx, span)
}

// AfterResponse is called after each LLM call completes.
// It records the response output, token usage, and ends the span.
func (h *Hook) AfterResponse(ctx context.Context, info metallm.LLMCallInfo, req *provider.ChatCompletionRequest, resp *provider.ChatCompletionResponse, err error) {
	span := spanFromContext(ctx)
	if span == nil {
		return
	}

	if err != nil {
		_ = span.End(llmops.WithEndError(err))
		return
	}

	if resp != nil {
		// Set output
		if len(resp.Choices) > 0 {
			_ = span.SetOutput(resp.Choices[0].Message.Content)
		}

		// Set token usage (Usage is a struct, not a pointer)
		_ = span.SetUsage(llmops.TokenUsage{
			PromptTokens:     resp.Usage.PromptTokens,
			CompletionTokens: resp.Usage.CompletionTokens,
			TotalTokens:      resp.Usage.TotalTokens,
		})
	}

	_ = span.End()
}

// WrapStream wraps a stream for observability.
// The wrapped stream will buffer content and record it when the stream ends.
func (h *Hook) WrapStream(ctx context.Context, info metallm.LLMCallInfo, req *provider.ChatCompletionRequest, stream provider.ChatCompletionStream) provider.ChatCompletionStream {
	span := spanFromContext(ctx)
	if span == nil {
		return stream
	}
	return &observedStream{
		stream: stream,
		span:   span,
		info:   info,
	}
}

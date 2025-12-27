package phoenix

import (
	"context"
	"time"

	"github.com/agentplexus/omniobserve/llmops"
	sdk "github.com/agentplexus/omniobserve/sdk/phoenix"
)

// traceAdapter adapts sdk.Trace to llmops.Trace.
type traceAdapter struct {
	trace *sdk.Trace
}

func (t *traceAdapter) ID() string {
	return t.trace.ID()
}

func (t *traceAdapter) Name() string {
	return t.trace.Name()
}

func (t *traceAdapter) StartSpan(ctx context.Context, name string, opts ...llmops.SpanOption) (context.Context, llmops.Span, error) {
	cfg := llmops.ApplySpanOptions(opts...)
	sdkOpts := mapSpanOptions(cfg)

	newCtx, span, err := t.trace.Span(ctx, name, sdkOpts...)
	if err != nil {
		return ctx, nil, err
	}
	return newCtx, &spanAdapter{span: span}, nil
}

func (t *traceAdapter) SetInput(input any) error {
	return t.trace.Update(context.Background(), sdk.WithInput(input))
}

func (t *traceAdapter) SetOutput(output any) error {
	return t.trace.Update(context.Background(), sdk.WithOutput(output))
}

func (t *traceAdapter) SetMetadata(metadata map[string]any) error {
	return t.trace.Update(context.Background(), sdk.WithMetadata(metadata))
}

func (t *traceAdapter) AddTag(tag string) error {
	// Phoenix uses metadata for tags
	return t.trace.Update(context.Background(), sdk.WithMetadata(map[string]any{"tags": []string{tag}}))
}

func (t *traceAdapter) AddFeedbackScore(ctx context.Context, name string, score float64, opts ...llmops.FeedbackOption) error {
	return t.trace.AddAnnotation(ctx, name, score)
}

func (t *traceAdapter) End(opts ...llmops.EndOption) error {
	cfg := &llmops.EndOptions{}
	for _, opt := range opts {
		opt(cfg)
	}

	sdkOpts := []sdk.TraceOption{}
	if cfg.Output != nil {
		sdkOpts = append(sdkOpts, sdk.WithOutput(cfg.Output))
	}
	if cfg.Metadata != nil {
		sdkOpts = append(sdkOpts, sdk.WithMetadata(cfg.Metadata))
	}

	return t.trace.End(context.Background(), sdkOpts...)
}

func (t *traceAdapter) EndTime() *time.Time {
	return t.trace.EndTime()
}

func (t *traceAdapter) Duration() time.Duration {
	startTime := t.trace.StartTime()
	if endTime := t.trace.EndTime(); endTime != nil {
		return endTime.Sub(startTime)
	}
	return time.Since(startTime)
}

// spanAdapter adapts sdk.Span to llmops.Span.
type spanAdapter struct {
	span *sdk.Span
}

func (s *spanAdapter) ID() string {
	return s.span.ID()
}

func (s *spanAdapter) TraceID() string {
	return s.span.TraceID()
}

func (s *spanAdapter) ParentSpanID() string {
	return s.span.ParentSpanID()
}

func (s *spanAdapter) Name() string {
	return s.span.Name()
}

func (s *spanAdapter) Type() llmops.SpanType {
	switch s.span.Kind() {
	case sdk.SpanKindLLM:
		return llmops.SpanTypeLLM
	case sdk.SpanKindTool:
		return llmops.SpanTypeTool
	case sdk.SpanKindRetriever:
		return llmops.SpanTypeRetrieval
	case sdk.SpanKindAgent:
		return llmops.SpanTypeAgent
	case sdk.SpanKindChain:
		return llmops.SpanTypeChain
	case sdk.SpanKindGuardrail:
		return llmops.SpanTypeGuardrail
	default:
		return llmops.SpanTypeGeneral
	}
}

func (s *spanAdapter) StartSpan(ctx context.Context, name string, opts ...llmops.SpanOption) (context.Context, llmops.Span, error) {
	cfg := llmops.ApplySpanOptions(opts...)
	sdkOpts := mapSpanOptions(cfg)

	newCtx, span, err := s.span.Span(ctx, name, sdkOpts...)
	if err != nil {
		return ctx, nil, err
	}
	return newCtx, &spanAdapter{span: span}, nil
}

func (s *spanAdapter) SetInput(input any) error {
	return s.span.SetInput(input)
}

func (s *spanAdapter) SetOutput(output any) error {
	return s.span.SetOutput(output)
}

func (s *spanAdapter) SetMetadata(metadata map[string]any) error {
	return s.span.Update(context.Background(), sdk.WithSpanMetadata(metadata))
}

func (s *spanAdapter) SetModel(model string) error {
	return s.span.SetModel(model)
}

func (s *spanAdapter) SetProvider(provider string) error {
	return s.span.Update(context.Background(), sdk.WithProvider(provider))
}

func (s *spanAdapter) SetUsage(usage llmops.TokenUsage) error {
	return s.span.SetUsage(usage.PromptTokens, usage.CompletionTokens, usage.TotalTokens)
}

func (s *spanAdapter) AddTag(tag string) error {
	return s.span.Update(context.Background(), sdk.WithSpanMetadata(map[string]any{"tags": []string{tag}}))
}

func (s *spanAdapter) AddFeedbackScore(ctx context.Context, name string, score float64, opts ...llmops.FeedbackOption) error {
	return s.span.AddAnnotation(ctx, name, score)
}

func (s *spanAdapter) End(opts ...llmops.EndOption) error {
	cfg := &llmops.EndOptions{}
	for _, opt := range opts {
		opt(cfg)
	}

	sdkOpts := []sdk.SpanOption{}
	if cfg.Output != nil {
		sdkOpts = append(sdkOpts, sdk.WithSpanOutput(cfg.Output))
	}
	if cfg.Metadata != nil {
		sdkOpts = append(sdkOpts, sdk.WithSpanMetadata(cfg.Metadata))
	}

	return s.span.End(context.Background(), sdkOpts...)
}

func (s *spanAdapter) EndTime() *time.Time {
	return s.span.EndTime()
}

func (s *spanAdapter) Duration() time.Duration {
	startTime := s.span.StartTime()
	if endTime := s.span.EndTime(); endTime != nil {
		return endTime.Sub(startTime)
	}
	return time.Since(startTime)
}

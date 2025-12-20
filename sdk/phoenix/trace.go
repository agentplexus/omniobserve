package phoenix

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// Trace represents an execution trace.
type Trace struct {
	client    *Client
	id        string
	name      string
	startTime time.Time
	endTime   *time.Time
	input     any
	output    any
	metadata  map[string]any
	sessionID string
	userID    string
	disabled  bool
	ended     bool
}

// ID returns the trace ID.
func (t *Trace) ID() string {
	return t.id
}

// Name returns the trace name.
func (t *Trace) Name() string {
	return t.name
}

// StartTime returns the start time.
func (t *Trace) StartTime() time.Time {
	return t.startTime
}

// EndTime returns the end time.
func (t *Trace) EndTime() *time.Time {
	return t.endTime
}

// End ends the trace.
func (t *Trace) End(ctx context.Context, opts ...TraceOption) error {
	if t.disabled || t.ended {
		return nil
	}
	t.ended = true

	cfg := &traceConfig{}
	for _, opt := range opts {
		opt(cfg)
	}

	endTime := time.Now()
	t.endTime = &endTime

	if cfg.output != nil {
		t.output = cfg.output
	}
	if cfg.metadata != nil {
		for k, v := range cfg.metadata {
			if t.metadata == nil {
				t.metadata = make(map[string]any)
			}
			t.metadata[k] = v
		}
	}

	// In a full implementation, this would send the trace to Phoenix
	return nil
}

// Update updates the trace.
func (t *Trace) Update(ctx context.Context, opts ...TraceOption) error {
	if t.disabled {
		return nil
	}

	cfg := &traceConfig{}
	for _, opt := range opts {
		opt(cfg)
	}

	if cfg.output != nil {
		t.output = cfg.output
	}
	if cfg.metadata != nil {
		for k, v := range cfg.metadata {
			if t.metadata == nil {
				t.metadata = make(map[string]any)
			}
			t.metadata[k] = v
		}
	}

	return nil
}

// Span creates a child span.
func (t *Trace) Span(ctx context.Context, name string, opts ...SpanOption) (context.Context, *Span, error) {
	if t.disabled {
		return ctx, &Span{disabled: true}, nil
	}

	cfg := &spanConfig{
		spanKind: SpanKindUnknown,
	}
	for _, opt := range opts {
		opt(cfg)
	}

	startTime := time.Now()
	if cfg.startTime != nil {
		startTime = *cfg.startTime
	}

	span := &Span{
		client:    t.client,
		id:        uuid.New().String(),
		traceID:   t.id,
		name:      name,
		kind:      cfg.spanKind,
		startTime: startTime,
		input:     cfg.input,
		metadata:  cfg.metadata,
		model:     cfg.model,
		provider:  cfg.provider,
		usage:     cfg.usage,
	}

	t.client.mu.Lock()
	t.client.spans = append(t.client.spans, span)
	t.client.mu.Unlock()

	newCtx := ContextWithSpan(ctx, span)
	return newCtx, span, nil
}

// LLMSpan creates an LLM span.
func (t *Trace) LLMSpan(ctx context.Context, name string, opts ...SpanOption) (context.Context, *Span, error) {
	opts = append([]SpanOption{WithSpanKind(SpanKindLLM)}, opts...)
	return t.Span(ctx, name, opts...)
}

// ChainSpan creates a chain span.
func (t *Trace) ChainSpan(ctx context.Context, name string, opts ...SpanOption) (context.Context, *Span, error) {
	opts = append([]SpanOption{WithSpanKind(SpanKindChain)}, opts...)
	return t.Span(ctx, name, opts...)
}

// ToolSpan creates a tool span.
func (t *Trace) ToolSpan(ctx context.Context, name string, opts ...SpanOption) (context.Context, *Span, error) {
	opts = append([]SpanOption{WithSpanKind(SpanKindTool)}, opts...)
	return t.Span(ctx, name, opts...)
}

// RetrieverSpan creates a retriever span.
func (t *Trace) RetrieverSpan(ctx context.Context, name string, opts ...SpanOption) (context.Context, *Span, error) {
	opts = append([]SpanOption{WithSpanKind(SpanKindRetriever)}, opts...)
	return t.Span(ctx, name, opts...)
}

// AddAnnotation adds an annotation to the trace.
func (t *Trace) AddAnnotation(ctx context.Context, name string, score float64, opts ...AnnotationOption) error {
	if t.disabled {
		return nil
	}
	return t.client.AddAnnotation(ctx, t.id, name, score, opts...)
}

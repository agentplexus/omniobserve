package phoenix

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// Span represents a span within a trace.
type Span struct {
	client       *Client
	id           string
	traceID      string
	parentSpanID string
	name         string
	kind         SpanKind
	startTime    time.Time
	endTime      *time.Time
	input        any
	output       any
	metadata     map[string]any
	model        string
	provider     string
	usage        *Usage
	events       []SpanEvent
	disabled     bool
	ended        bool
}

// ID returns the span ID.
func (s *Span) ID() string {
	return s.id
}

// TraceID returns the parent trace ID.
func (s *Span) TraceID() string {
	return s.traceID
}

// ParentSpanID returns the parent span ID.
func (s *Span) ParentSpanID() string {
	return s.parentSpanID
}

// Name returns the span name.
func (s *Span) Name() string {
	return s.name
}

// Kind returns the span kind.
func (s *Span) Kind() SpanKind {
	return s.kind
}

// StartTime returns the start time.
func (s *Span) StartTime() time.Time {
	return s.startTime
}

// EndTime returns the end time.
func (s *Span) EndTime() *time.Time {
	return s.endTime
}

// End ends the span.
func (s *Span) End(ctx context.Context, opts ...SpanOption) error {
	if s.disabled || s.ended {
		return nil
	}
	s.ended = true

	cfg := &spanConfig{}
	for _, opt := range opts {
		opt(cfg)
	}

	endTime := time.Now()
	s.endTime = &endTime

	if cfg.output != nil {
		s.output = cfg.output
	}
	if cfg.usage != nil {
		s.usage = cfg.usage
	}
	if cfg.metadata != nil {
		for k, v := range cfg.metadata {
			if s.metadata == nil {
				s.metadata = make(map[string]any)
			}
			s.metadata[k] = v
		}
	}

	return nil
}

// Update updates the span.
func (s *Span) Update(ctx context.Context, opts ...SpanOption) error {
	if s.disabled {
		return nil
	}

	cfg := &spanConfig{}
	for _, opt := range opts {
		opt(cfg)
	}

	if cfg.input != nil {
		s.input = cfg.input
	}
	if cfg.output != nil {
		s.output = cfg.output
	}
	if cfg.usage != nil {
		s.usage = cfg.usage
	}
	if cfg.model != "" {
		s.model = cfg.model
	}
	if cfg.provider != "" {
		s.provider = cfg.provider
	}
	if cfg.metadata != nil {
		for k, v := range cfg.metadata {
			if s.metadata == nil {
				s.metadata = make(map[string]any)
			}
			s.metadata[k] = v
		}
	}

	return nil
}

// Span creates a child span.
func (s *Span) Span(ctx context.Context, name string, opts ...SpanOption) (context.Context, *Span, error) {
	if s.disabled {
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

	child := &Span{
		client:       s.client,
		id:           uuid.New().String(),
		traceID:      s.traceID,
		parentSpanID: s.id,
		name:         name,
		kind:         cfg.spanKind,
		startTime:    startTime,
		input:        cfg.input,
		metadata:     cfg.metadata,
		model:        cfg.model,
		provider:     cfg.provider,
		usage:        cfg.usage,
	}

	s.client.mu.Lock()
	s.client.spans = append(s.client.spans, child)
	s.client.mu.Unlock()

	newCtx := ContextWithSpan(ctx, child)
	return newCtx, child, nil
}

// LLMSpan creates an LLM child span.
func (s *Span) LLMSpan(ctx context.Context, name string, opts ...SpanOption) (context.Context, *Span, error) {
	opts = append([]SpanOption{WithSpanKind(SpanKindLLM)}, opts...)
	return s.Span(ctx, name, opts...)
}

// ToolSpan creates a tool child span.
func (s *Span) ToolSpan(ctx context.Context, name string, opts ...SpanOption) (context.Context, *Span, error) {
	opts = append([]SpanOption{WithSpanKind(SpanKindTool)}, opts...)
	return s.Span(ctx, name, opts...)
}

// RetrieverSpan creates a retriever child span.
func (s *Span) RetrieverSpan(ctx context.Context, name string, opts ...SpanOption) (context.Context, *Span, error) {
	opts = append([]SpanOption{WithSpanKind(SpanKindRetriever)}, opts...)
	return s.Span(ctx, name, opts...)
}

// AddEvent adds an event to the span.
func (s *Span) AddEvent(name string, attrs map[string]any) {
	if s.disabled {
		return
	}
	s.events = append(s.events, SpanEvent{
		Name:       name,
		Timestamp:  time.Now(),
		Attributes: attrs,
	})
}

// SetInput sets the span input.
func (s *Span) SetInput(input any) error {
	return s.Update(context.Background(), WithSpanInput(input))
}

// SetOutput sets the span output.
func (s *Span) SetOutput(output any) error {
	return s.Update(context.Background(), WithSpanOutput(output))
}

// SetModel sets the model name.
func (s *Span) SetModel(model string) error {
	return s.Update(context.Background(), WithModel(model))
}

// SetUsage sets token usage.
func (s *Span) SetUsage(prompt, completion, total int) error {
	return s.Update(context.Background(), WithUsage(prompt, completion, total))
}

// AddAnnotation adds an annotation to the span.
func (s *Span) AddAnnotation(ctx context.Context, name string, score float64, opts ...AnnotationOption) error {
	if s.disabled {
		return nil
	}
	return s.client.AddAnnotation(ctx, s.id, name, score, opts...)
}

// SetDocuments sets retrieved documents (for retriever spans).
func (s *Span) SetDocuments(docs []Document) error {
	if s.disabled {
		return nil
	}
	if s.metadata == nil {
		s.metadata = make(map[string]any)
	}
	s.metadata["retriever.documents"] = docs
	return nil
}

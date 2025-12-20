package langfuse

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
	startTime    time.Time
	endTime      *time.Time
	input        any
	output       any
	metadata     map[string]any
	level        string
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
	if cfg.metadata != nil {
		for k, v := range cfg.metadata {
			if s.metadata == nil {
				s.metadata = make(map[string]any)
			}
			s.metadata[k] = v
		}
	}

	s.client.enqueue(Event{
		ID:        uuid.New().String(),
		Type:      EventTypeSpanUpdate,
		Timestamp: time.Now(),
		Body: SpanBody{
			ID:                  s.id,
			TraceID:             s.traceID,
			ParentObservationID: s.parentSpanID,
			Name:                s.name,
			StartTime:           s.startTime,
			EndTime:             s.endTime,
			Input:               s.input,
			Output:              s.output,
			Metadata:            s.metadata,
			Level:               s.level,
		},
	})

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
	if cfg.metadata != nil {
		for k, v := range cfg.metadata {
			if s.metadata == nil {
				s.metadata = make(map[string]any)
			}
			s.metadata[k] = v
		}
	}
	if cfg.level != "" {
		s.level = cfg.level
	}

	s.client.enqueue(Event{
		ID:        uuid.New().String(),
		Type:      EventTypeSpanUpdate,
		Timestamp: time.Now(),
		Body: SpanBody{
			ID:                  s.id,
			TraceID:             s.traceID,
			ParentObservationID: s.parentSpanID,
			Name:                s.name,
			StartTime:           s.startTime,
			EndTime:             s.endTime,
			Input:               s.input,
			Output:              s.output,
			Metadata:            s.metadata,
			Level:               s.level,
		},
	})

	return nil
}

// Span creates a child span.
func (s *Span) Span(ctx context.Context, name string, opts ...SpanOption) (context.Context, *Span, error) {
	if s.disabled {
		return ctx, &Span{disabled: true}, nil
	}

	cfg := &spanConfig{}
	for _, opt := range opts {
		opt(cfg)
	}

	child := &Span{
		client:       s.client,
		id:           uuid.New().String(),
		traceID:      s.traceID,
		parentSpanID: s.id,
		name:         name,
		startTime:    time.Now(),
		input:        cfg.input,
		metadata:     cfg.metadata,
		level:        cfg.level,
	}

	s.client.enqueue(Event{
		ID:        uuid.New().String(),
		Type:      EventTypeSpanCreate,
		Timestamp: time.Now(),
		Body: SpanBody{
			ID:                  child.id,
			TraceID:             s.traceID,
			ParentObservationID: s.id,
			Name:                name,
			StartTime:           child.startTime,
			Input:               cfg.input,
			Metadata:            cfg.metadata,
			Level:               cfg.level,
			Version:             cfg.version,
		},
	})

	newCtx := ContextWithSpan(ctx, child)
	return newCtx, child, nil
}

// Generation creates an LLM generation within this span.
func (s *Span) Generation(ctx context.Context, name string, opts ...GenerationOption) (context.Context, *Generation, error) {
	if s.disabled {
		return ctx, &Generation{disabled: true}, nil
	}

	cfg := &generationConfig{}
	for _, opt := range opts {
		opt(cfg)
	}

	gen := &Generation{
		client:          s.client,
		id:              uuid.New().String(),
		traceID:         s.traceID,
		parentSpanID:    s.id,
		name:            name,
		startTime:       time.Now(),
		model:           cfg.model,
		modelParameters: cfg.modelParameters,
		input:           cfg.input,
		metadata:        cfg.metadata,
		promptName:      cfg.promptName,
		promptVersion:   cfg.promptVersion,
	}

	s.client.enqueue(Event{
		ID:        uuid.New().String(),
		Type:      EventTypeGenerationCreate,
		Timestamp: time.Now(),
		Body: GenerationBody{
			ID:                  gen.id,
			TraceID:             s.traceID,
			ParentObservationID: s.id,
			Name:                name,
			StartTime:           gen.startTime,
			Model:               cfg.model,
			ModelParameters:     cfg.modelParameters,
			Input:               cfg.input,
			Metadata:            cfg.metadata,
			PromptName:          cfg.promptName,
			PromptVersion:       cfg.promptVersion,
			Level:               cfg.level,
		},
	})

	newCtx := ContextWithGeneration(ctx, gen)
	return newCtx, gen, nil
}

// Score adds a score to the span.
func (s *Span) Score(ctx context.Context, name string, value float64, opts ...ScoreOption) error {
	if s.disabled {
		return nil
	}

	cfg := &scoreConfig{}
	for _, opt := range opts {
		opt(cfg)
	}

	s.client.enqueue(Event{
		ID:        uuid.New().String(),
		Type:      EventTypeScoreCreate,
		Timestamp: time.Now(),
		Body: ScoreBody{
			ID:            uuid.New().String(),
			TraceID:       s.traceID,
			ObservationID: s.id,
			Name:          name,
			Value:         value,
			Comment:       cfg.comment,
			Source:        cfg.source,
			DataType:      "NUMERIC",
		},
	})

	return nil
}

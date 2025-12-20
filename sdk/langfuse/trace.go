package langfuse

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
	tags      []string
	userId    string
	sessionId string
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

	// Send trace update
	t.client.enqueue(Event{
		ID:        uuid.New().String(),
		Type:      EventTypeTraceCreate, // Langfuse uses same event for update
		Timestamp: time.Now(),
		Body: TraceBody{
			ID:        t.id,
			Name:      t.name,
			Timestamp: t.startTime,
			Metadata:  t.metadata,
			Tags:      t.tags,
			UserID:    t.userId,
			SessionID: t.sessionId,
			Input:     t.input,
			Output:    t.output,
		},
	})

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
	if len(cfg.tags) > 0 {
		t.tags = append(t.tags, cfg.tags...)
	}

	t.client.enqueue(Event{
		ID:        uuid.New().String(),
		Type:      EventTypeTraceCreate,
		Timestamp: time.Now(),
		Body: TraceBody{
			ID:        t.id,
			Name:      t.name,
			Timestamp: t.startTime,
			Metadata:  t.metadata,
			Tags:      t.tags,
			UserID:    t.userId,
			SessionID: t.sessionId,
			Input:     t.input,
			Output:    t.output,
		},
	})

	return nil
}

// Span creates a child span.
func (t *Trace) Span(ctx context.Context, name string, opts ...SpanOption) (context.Context, *Span, error) {
	if t.disabled {
		return ctx, &Span{disabled: true}, nil
	}

	cfg := &spanConfig{}
	for _, opt := range opts {
		opt(cfg)
	}

	span := &Span{
		client:    t.client,
		id:        uuid.New().String(),
		traceID:   t.id,
		name:      name,
		startTime: time.Now(),
		input:     cfg.input,
		metadata:  cfg.metadata,
		level:     cfg.level,
	}

	t.client.enqueue(Event{
		ID:        uuid.New().String(),
		Type:      EventTypeSpanCreate,
		Timestamp: time.Now(),
		Body: SpanBody{
			ID:        span.id,
			TraceID:   t.id,
			Name:      name,
			StartTime: span.startTime,
			Input:     cfg.input,
			Metadata:  cfg.metadata,
			Level:     cfg.level,
			Version:   cfg.version,
		},
	})

	newCtx := ContextWithSpan(ctx, span)
	return newCtx, span, nil
}

// Generation creates an LLM generation span.
func (t *Trace) Generation(ctx context.Context, name string, opts ...GenerationOption) (context.Context, *Generation, error) {
	if t.disabled {
		return ctx, &Generation{disabled: true}, nil
	}

	cfg := &generationConfig{}
	for _, opt := range opts {
		opt(cfg)
	}

	gen := &Generation{
		client:          t.client,
		id:              uuid.New().String(),
		traceID:         t.id,
		name:            name,
		startTime:       time.Now(),
		model:           cfg.model,
		modelParameters: cfg.modelParameters,
		input:           cfg.input,
		metadata:        cfg.metadata,
		promptName:      cfg.promptName,
		promptVersion:   cfg.promptVersion,
	}

	t.client.enqueue(Event{
		ID:        uuid.New().String(),
		Type:      EventTypeGenerationCreate,
		Timestamp: time.Now(),
		Body: GenerationBody{
			ID:              gen.id,
			TraceID:         t.id,
			Name:            name,
			StartTime:       gen.startTime,
			Model:           cfg.model,
			ModelParameters: cfg.modelParameters,
			Input:           cfg.input,
			Metadata:        cfg.metadata,
			PromptName:      cfg.promptName,
			PromptVersion:   cfg.promptVersion,
			Level:           cfg.level,
		},
	})

	newCtx := ContextWithGeneration(ctx, gen)
	return newCtx, gen, nil
}

// Score adds a score to the trace.
func (t *Trace) Score(ctx context.Context, name string, value float64, opts ...ScoreOption) error {
	if t.disabled {
		return nil
	}

	cfg := &scoreConfig{}
	for _, opt := range opts {
		opt(cfg)
	}

	t.client.enqueue(Event{
		ID:        uuid.New().String(),
		Type:      EventTypeScoreCreate,
		Timestamp: time.Now(),
		Body: ScoreBody{
			ID:       uuid.New().String(),
			TraceID:  t.id,
			Name:     name,
			Value:    value,
			Comment:  cfg.comment,
			Source:   cfg.source,
			DataType: "NUMERIC",
		},
	})

	return nil
}

// scoreConfig holds score configuration.
type scoreConfig struct {
	comment   string
	source    string
	dataType  string
}

// ScoreOption configures score creation.
type ScoreOption func(*scoreConfig)

// WithScoreComment sets the score comment.
func WithScoreComment(comment string) ScoreOption {
	return func(c *scoreConfig) {
		c.comment = comment
	}
}

// WithScoreSource sets the score source.
func WithScoreSource(source string) ScoreOption {
	return func(c *scoreConfig) {
		c.source = source
	}
}

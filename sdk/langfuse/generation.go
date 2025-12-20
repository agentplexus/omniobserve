package langfuse

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// Generation represents an LLM generation (call) within a trace.
type Generation struct {
	client          *Client
	id              string
	traceID         string
	parentSpanID    string
	name            string
	startTime       time.Time
	endTime         *time.Time
	completionStart *time.Time
	model           string
	modelParameters map[string]any
	input           any
	output          any
	metadata        map[string]any
	usage           *Usage
	promptName      string
	promptVersion   int
	level           string
	disabled        bool
	ended           bool
}

// ID returns the generation ID.
func (g *Generation) ID() string {
	return g.id
}

// TraceID returns the parent trace ID.
func (g *Generation) TraceID() string {
	return g.traceID
}

// ParentSpanID returns the parent span ID.
func (g *Generation) ParentSpanID() string {
	return g.parentSpanID
}

// Name returns the generation name.
func (g *Generation) Name() string {
	return g.name
}

// Model returns the model name.
func (g *Generation) Model() string {
	return g.model
}

// StartTime returns the start time.
func (g *Generation) StartTime() time.Time {
	return g.startTime
}

// EndTime returns the end time.
func (g *Generation) EndTime() *time.Time {
	return g.endTime
}

// End ends the generation.
func (g *Generation) End(ctx context.Context, opts ...GenerationOption) error {
	if g.disabled || g.ended {
		return nil
	}
	g.ended = true

	cfg := &generationConfig{}
	for _, opt := range opts {
		opt(cfg)
	}

	endTime := time.Now()
	g.endTime = &endTime

	if cfg.output != nil {
		g.output = cfg.output
	}
	if cfg.usage != nil {
		g.usage = cfg.usage
	}
	if cfg.metadata != nil {
		for k, v := range cfg.metadata {
			if g.metadata == nil {
				g.metadata = make(map[string]any)
			}
			g.metadata[k] = v
		}
	}
	if cfg.model != "" {
		g.model = cfg.model
	}

	g.client.enqueue(Event{
		ID:        uuid.New().String(),
		Type:      EventTypeGenerationUpdate,
		Timestamp: time.Now(),
		Body: GenerationBody{
			ID:                  g.id,
			TraceID:             g.traceID,
			ParentObservationID: g.parentSpanID,
			Name:                g.name,
			StartTime:           g.startTime,
			EndTime:             g.endTime,
			CompletionStartTime: g.completionStart,
			Model:               g.model,
			ModelParameters:     g.modelParameters,
			Input:               g.input,
			Output:              g.output,
			Metadata:            g.metadata,
			Usage:               g.usage,
			PromptName:          g.promptName,
			PromptVersion:       g.promptVersion,
			Level:               g.level,
		},
	})

	return nil
}

// Update updates the generation.
func (g *Generation) Update(ctx context.Context, opts ...GenerationOption) error {
	if g.disabled {
		return nil
	}

	cfg := &generationConfig{}
	for _, opt := range opts {
		opt(cfg)
	}

	if cfg.input != nil {
		g.input = cfg.input
	}
	if cfg.output != nil {
		g.output = cfg.output
	}
	if cfg.usage != nil {
		g.usage = cfg.usage
	}
	if cfg.metadata != nil {
		for k, v := range cfg.metadata {
			if g.metadata == nil {
				g.metadata = make(map[string]any)
			}
			g.metadata[k] = v
		}
	}
	if cfg.model != "" {
		g.model = cfg.model
	}

	g.client.enqueue(Event{
		ID:        uuid.New().String(),
		Type:      EventTypeGenerationUpdate,
		Timestamp: time.Now(),
		Body: GenerationBody{
			ID:                  g.id,
			TraceID:             g.traceID,
			ParentObservationID: g.parentSpanID,
			Name:                g.name,
			StartTime:           g.startTime,
			EndTime:             g.endTime,
			CompletionStartTime: g.completionStart,
			Model:               g.model,
			ModelParameters:     g.modelParameters,
			Input:               g.input,
			Output:              g.output,
			Metadata:            g.metadata,
			Usage:               g.usage,
			PromptName:          g.promptName,
			PromptVersion:       g.promptVersion,
			Level:               g.level,
		},
	})

	return nil
}

// SetOutput sets the generation output.
func (g *Generation) SetOutput(output any) error {
	return g.Update(context.Background(), WithGenerationOutput(output))
}

// SetUsage sets token usage.
func (g *Generation) SetUsage(promptTokens, completionTokens, totalTokens int) error {
	return g.Update(context.Background(), WithUsage(promptTokens, completionTokens, totalTokens))
}

// MarkCompletionStart marks when the first token was received (for streaming).
func (g *Generation) MarkCompletionStart() {
	if g.completionStart == nil {
		now := time.Now()
		g.completionStart = &now
	}
}

// Score adds a score to the generation.
func (g *Generation) Score(ctx context.Context, name string, value float64, opts ...ScoreOption) error {
	if g.disabled {
		return nil
	}

	cfg := &scoreConfig{}
	for _, opt := range opts {
		opt(cfg)
	}

	g.client.enqueue(Event{
		ID:        uuid.New().String(),
		Type:      EventTypeScoreCreate,
		Timestamp: time.Now(),
		Body: ScoreBody{
			ID:            uuid.New().String(),
			TraceID:       g.traceID,
			ObservationID: g.id,
			Name:          name,
			Value:         value,
			Comment:       cfg.comment,
			Source:        cfg.source,
			DataType:      "NUMERIC",
		},
	})

	return nil
}

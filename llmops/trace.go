package llmops

import (
	"context"
	"time"
)

// Trace represents an execution trace for an LLM operation or workflow.
type Trace interface {
	// ID returns the unique identifier for this trace.
	ID() string

	// Name returns the trace name.
	Name() string

	// StartSpan creates a child span within this trace.
	StartSpan(ctx context.Context, name string, opts ...SpanOption) (context.Context, Span, error)

	// SetInput sets the trace input data.
	SetInput(input any) error

	// SetOutput sets the trace output data.
	SetOutput(output any) error

	// SetMetadata sets additional metadata on the trace.
	SetMetadata(metadata map[string]any) error

	// AddTag adds a tag to the trace.
	AddTag(tag string) error

	// AddFeedbackScore adds a feedback score to this trace.
	AddFeedbackScore(ctx context.Context, name string, score float64, opts ...FeedbackOption) error

	// End completes the trace with optional final output and metadata.
	End(opts ...EndOption) error

	// EndTime returns when the trace ended, if it has ended.
	EndTime() *time.Time

	// Duration returns the trace duration, or time since start if not ended.
	Duration() time.Duration
}

// Span represents a unit of work within a trace, such as an LLM call.
type Span interface {
	// ID returns the unique identifier for this span.
	ID() string

	// TraceID returns the parent trace ID.
	TraceID() string

	// ParentSpanID returns the parent span ID, if this is a nested span.
	ParentSpanID() string

	// Name returns the span name.
	Name() string

	// Type returns the span type (general, llm, tool, guardrail).
	Type() SpanType

	// StartSpan creates a child span within this span.
	StartSpan(ctx context.Context, name string, opts ...SpanOption) (context.Context, Span, error)

	// SetInput sets the span input data.
	SetInput(input any) error

	// SetOutput sets the span output data.
	SetOutput(output any) error

	// SetMetadata sets additional metadata on the span.
	SetMetadata(metadata map[string]any) error

	// SetModel sets the LLM model name (e.g., "gpt-4", "claude-3-opus").
	SetModel(model string) error

	// SetProvider sets the LLM provider name (e.g., "openai", "anthropic").
	SetProvider(provider string) error

	// SetUsage sets token usage information.
	SetUsage(usage TokenUsage) error

	// AddTag adds a tag to the span.
	AddTag(tag string) error

	// AddFeedbackScore adds a feedback score to this span.
	AddFeedbackScore(ctx context.Context, name string, score float64, opts ...FeedbackOption) error

	// End completes the span with optional final output and metadata.
	End(opts ...EndOption) error

	// EndTime returns when the span ended, if it has ended.
	EndTime() *time.Time

	// Duration returns the span duration, or time since start if not ended.
	Duration() time.Duration
}

// SpanType categorizes spans by their function.
type SpanType string

const (
	SpanTypeGeneral   SpanType = "general"
	SpanTypeLLM       SpanType = "llm"
	SpanTypeTool      SpanType = "tool"
	SpanTypeGuardrail SpanType = "guardrail"
	SpanTypeRetrieval SpanType = "retrieval"
	SpanTypeAgent     SpanType = "agent"
	SpanTypeChain     SpanType = "chain"
)

// TokenUsage represents token consumption for an LLM call.
type TokenUsage struct {
	PromptTokens     int `json:"prompt_tokens,omitempty"`
	CompletionTokens int `json:"completion_tokens,omitempty"`
	TotalTokens      int `json:"total_tokens,omitempty"`

	// Cost tracking (optional, provider-dependent)
	PromptCost     float64 `json:"prompt_cost,omitempty"`
	CompletionCost float64 `json:"completion_cost,omitempty"`
	TotalCost      float64 `json:"total_cost,omitempty"`
	Currency       string  `json:"currency,omitempty"` // e.g., "USD"
}

// FeedbackScore represents a score given to a trace or span.
type FeedbackScore struct {
	Name     string  `json:"name"`
	Score    float64 `json:"score"`
	Reason   string  `json:"reason,omitempty"`
	Category string  `json:"category,omitempty"`
	Source   string  `json:"source,omitempty"` // e.g., "user", "llm", "heuristic"
}

// TraceInfo provides read-only information about a trace.
type TraceInfo struct {
	ID        string         `json:"id"`
	Name      string         `json:"name"`
	ProjectID string         `json:"project_id,omitempty"`
	StartTime time.Time      `json:"start_time"`
	EndTime   *time.Time     `json:"end_time,omitempty"`
	Input     any            `json:"input,omitempty"`
	Output    any            `json:"output,omitempty"`
	Metadata  map[string]any `json:"metadata,omitempty"`
	Tags      []string       `json:"tags,omitempty"`
	Feedback  []FeedbackScore `json:"feedback,omitempty"`
}

// SpanInfo provides read-only information about a span.
type SpanInfo struct {
	ID           string         `json:"id"`
	TraceID      string         `json:"trace_id"`
	ParentSpanID string         `json:"parent_span_id,omitempty"`
	Name         string         `json:"name"`
	Type         SpanType       `json:"type"`
	StartTime    time.Time      `json:"start_time"`
	EndTime      *time.Time     `json:"end_time,omitempty"`
	Input        any            `json:"input,omitempty"`
	Output       any            `json:"output,omitempty"`
	Metadata     map[string]any `json:"metadata,omitempty"`
	Model        string         `json:"model,omitempty"`
	Provider     string         `json:"provider,omitempty"`
	Usage        *TokenUsage    `json:"usage,omitempty"`
	Tags         []string       `json:"tags,omitempty"`
	Feedback     []FeedbackScore `json:"feedback,omitempty"`
}

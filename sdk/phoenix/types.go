package phoenix

import "time"

// SpanKind represents the type of span.
type SpanKind string

const (
	SpanKindLLM       SpanKind = "LLM"
	SpanKindChain     SpanKind = "CHAIN"
	SpanKindTool      SpanKind = "TOOL"
	SpanKindAgent     SpanKind = "AGENT"
	SpanKindRetriever SpanKind = "RETRIEVER"
	SpanKindEmbedding SpanKind = "EMBEDDING"
	SpanKindReranker  SpanKind = "RERANKER"
	SpanKindGuardrail SpanKind = "GUARDRAIL"
	SpanKindEvaluator SpanKind = "EVALUATOR"
	SpanKindUnknown   SpanKind = "UNKNOWN"
)

// Usage represents token usage for LLM calls.
type Usage struct {
	PromptTokens     int `json:"prompt_tokens,omitempty"`
	CompletionTokens int `json:"completion_tokens,omitempty"`
	TotalTokens      int `json:"total_tokens,omitempty"`
}

// SpanAttributes holds semantic attributes for spans.
// Phoenix uses OpenTelemetry semantic conventions for LLM observability.
type SpanAttributes struct {
	// LLM attributes
	LLMModelName       string `json:"llm.model_name,omitempty"`
	LLMProvider        string `json:"llm.provider,omitempty"`
	LLMTokenCountPrompt     int    `json:"llm.token_count.prompt,omitempty"`
	LLMTokenCountCompletion int    `json:"llm.token_count.completion,omitempty"`
	LLMTokenCountTotal      int    `json:"llm.token_count.total,omitempty"`

	// Input/Output
	InputValue  any `json:"input.value,omitempty"`
	OutputValue any `json:"output.value,omitempty"`

	// Retrieval attributes
	RetrieverQueryEmbedding []float64 `json:"retriever.query.embedding,omitempty"`
	RetrieverDocumentCount  int       `json:"retriever.document_count,omitempty"`

	// Tool attributes
	ToolName        string `json:"tool.name,omitempty"`
	ToolDescription string `json:"tool.description,omitempty"`

	// Embedding attributes
	EmbeddingModelName string    `json:"embedding.model_name,omitempty"`
	EmbeddingVector    []float64 `json:"embedding.vector,omitempty"`

	// Session attributes
	SessionID string `json:"session.id,omitempty"`
	UserID    string `json:"user.id,omitempty"`

	// Custom metadata
	Metadata map[string]any `json:"metadata,omitempty"`
}

// TraceInfo provides read-only information about a trace.
type TraceInfo struct {
	ID        string         `json:"id"`
	Name      string         `json:"name"`
	ProjectID string         `json:"projectId,omitempty"`
	StartTime time.Time      `json:"startTime"`
	EndTime   *time.Time     `json:"endTime,omitempty"`
	Latency   float64        `json:"latencyMs,omitempty"`
	Status    string         `json:"status,omitempty"`
	Metadata  map[string]any `json:"metadata,omitempty"`
	SpanCount int            `json:"numSpans,omitempty"`
}

// SpanInfo provides read-only information about a span.
type SpanInfo struct {
	ID           string         `json:"id"`
	TraceID      string         `json:"traceId"`
	ParentID     string         `json:"parentId,omitempty"`
	Name         string         `json:"name"`
	Kind         SpanKind       `json:"spanKind"`
	StartTime    time.Time      `json:"startTime"`
	EndTime      *time.Time     `json:"endTime,omitempty"`
	Latency      float64        `json:"latencyMs,omitempty"`
	Status       string         `json:"status,omitempty"`
	Attributes   SpanAttributes `json:"attributes,omitempty"`
	Events       []SpanEvent    `json:"events,omitempty"`
	StatusMessage string        `json:"statusMessage,omitempty"`
}

// SpanEvent represents an event within a span.
type SpanEvent struct {
	Name       string         `json:"name"`
	Timestamp  time.Time      `json:"timestamp"`
	Attributes map[string]any `json:"attributes,omitempty"`
}

// Annotation represents a human or automated annotation on a span.
type Annotation struct {
	ID          string    `json:"id"`
	SpanID      string    `json:"spanId"`
	Name        string    `json:"name"`
	AnnotatorKind string  `json:"annotatorKind"` // HUMAN, LLM
	Score       *float64  `json:"score,omitempty"`
	Label       string    `json:"label,omitempty"`
	Explanation string    `json:"explanation,omitempty"`
	Metadata    map[string]any `json:"metadata,omitempty"`
	CreatedAt   time.Time `json:"createdAt"`
}

// Evaluation represents an evaluation result.
type Evaluation struct {
	ID          string         `json:"id"`
	Name        string         `json:"name"`
	TraceID     string         `json:"traceId,omitempty"`
	SpanID      string         `json:"spanId,omitempty"`
	Score       float64        `json:"score"`
	Label       string         `json:"label,omitempty"`
	Explanation string         `json:"explanation,omitempty"`
	Metadata    map[string]any `json:"metadata,omitempty"`
	CreatedAt   time.Time      `json:"createdAt"`
}

// Document represents a retrieved document.
type Document struct {
	ID       string         `json:"id,omitempty"`
	Content  string         `json:"content"`
	Score    float64        `json:"score,omitempty"`
	Metadata map[string]any `json:"metadata,omitempty"`
}

// Message represents a chat message.
type Message struct {
	Role       string         `json:"role"`
	Content    string         `json:"content"`
	Name       string         `json:"name,omitempty"`
	ToolCalls  []ToolCall     `json:"toolCalls,omitempty"`
	ToolCallID string         `json:"toolCallId,omitempty"`
}

// ToolCall represents a tool call.
type ToolCall struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Arguments string `json:"arguments"`
}

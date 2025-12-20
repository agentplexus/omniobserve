package langfuse

import "time"

// Event types for batch ingestion.
const (
	EventTypeTraceCreate      = "trace-create"
	EventTypeSpanCreate       = "span-create"
	EventTypeSpanUpdate       = "span-update"
	EventTypeGenerationCreate = "generation-create"
	EventTypeGenerationUpdate = "generation-update"
	EventTypeScoreCreate      = "score-create"
	EventTypeEventCreate      = "event-create"
)

// BatchIngestionRequest is the request body for batch ingestion.
type BatchIngestionRequest struct {
	Batch    []Event        `json:"batch"`
	Metadata map[string]any `json:"metadata,omitempty"`
}

// Event represents a single event in the batch.
type Event struct {
	ID        string    `json:"id"`
	Type      string    `json:"type"`
	Timestamp time.Time `json:"timestamp"`
	Body      any       `json:"body"`
}

// TraceBody represents the body of a trace event.
type TraceBody struct {
	ID        string         `json:"id"`
	Name      string         `json:"name,omitempty"`
	Timestamp time.Time      `json:"timestamp"`
	Metadata  map[string]any `json:"metadata,omitempty"`
	Tags      []string       `json:"tags,omitempty"`
	UserID    string         `json:"userId,omitempty"`
	SessionID string         `json:"sessionId,omitempty"`
	Input     any            `json:"input,omitempty"`
	Output    any            `json:"output,omitempty"`
	Public    bool           `json:"public,omitempty"`
}

// SpanBody represents the body of a span event.
type SpanBody struct {
	ID                  string         `json:"id"`
	TraceID             string         `json:"traceId"`
	ParentObservationID string         `json:"parentObservationId,omitempty"`
	Name                string         `json:"name,omitempty"`
	StartTime           time.Time      `json:"startTime"`
	EndTime             *time.Time     `json:"endTime,omitempty"`
	Metadata            map[string]any `json:"metadata,omitempty"`
	Input               any            `json:"input,omitempty"`
	Output              any            `json:"output,omitempty"`
	Level               string         `json:"level,omitempty"` // DEBUG, DEFAULT, WARNING, ERROR
	Version             string         `json:"version,omitempty"`
	StatusMessage       string         `json:"statusMessage,omitempty"`
}

// GenerationBody represents the body of a generation event (LLM call).
type GenerationBody struct {
	ID                  string         `json:"id"`
	TraceID             string         `json:"traceId"`
	ParentObservationID string         `json:"parentObservationId,omitempty"`
	Name                string         `json:"name,omitempty"`
	StartTime           time.Time      `json:"startTime"`
	EndTime             *time.Time     `json:"endTime,omitempty"`
	CompletionStartTime *time.Time     `json:"completionStartTime,omitempty"`
	Model               string         `json:"model,omitempty"`
	ModelParameters     map[string]any `json:"modelParameters,omitempty"`
	Input               any            `json:"input,omitempty"`
	Output              any            `json:"output,omitempty"`
	Metadata            map[string]any `json:"metadata,omitempty"`
	Usage               *Usage         `json:"usage,omitempty"`
	Level               string         `json:"level,omitempty"`
	PromptName          string         `json:"promptName,omitempty"`
	PromptVersion       int            `json:"promptVersion,omitempty"`
	StatusMessage       string         `json:"statusMessage,omitempty"`
}

// Usage represents token usage for a generation.
type Usage struct {
	PromptTokens     int     `json:"promptTokens,omitempty"`
	CompletionTokens int     `json:"completionTokens,omitempty"`
	TotalTokens      int     `json:"totalTokens,omitempty"`
	Input            int     `json:"input,omitempty"`  // Alternative field
	Output           int     `json:"output,omitempty"` // Alternative field
	Total            int     `json:"total,omitempty"`  // Alternative field
	Unit             string  `json:"unit,omitempty"`   // TOKENS, CHARACTERS, etc.
	InputCost        float64 `json:"inputCost,omitempty"`
	OutputCost       float64 `json:"outputCost,omitempty"`
	TotalCost        float64 `json:"totalCost,omitempty"`
}

// ScoreBody represents the body of a score event.
type ScoreBody struct {
	ID            string  `json:"id"`
	TraceID       string  `json:"traceId"`
	ObservationID string  `json:"observationId,omitempty"`
	Name          string  `json:"name"`
	Value         float64 `json:"value,omitempty"`
	StringValue   string  `json:"stringValue,omitempty"`
	DataType      string  `json:"dataType,omitempty"` // NUMERIC, CATEGORICAL, BOOLEAN
	Comment       string  `json:"comment,omitempty"`
	Source        string  `json:"source,omitempty"` // API, ANNOTATION, EVAL
}

// Observation represents an observation from the API.
type Observation struct {
	ID            string         `json:"id"`
	TraceID       string         `json:"traceId"`
	Type          string         `json:"type"` // SPAN, GENERATION, EVENT
	Name          string         `json:"name"`
	StartTime     time.Time      `json:"startTime"`
	EndTime       *time.Time     `json:"endTime,omitempty"`
	Model         string         `json:"model,omitempty"`
	Input         any            `json:"input,omitempty"`
	Output        any            `json:"output,omitempty"`
	Metadata      map[string]any `json:"metadata,omitempty"`
	Usage         *Usage         `json:"usage,omitempty"`
	Level         string         `json:"level,omitempty"`
	PromptName    string         `json:"promptName,omitempty"`
	PromptVersion int            `json:"promptVersion,omitempty"`
}

// TraceInfo represents trace information from the API.
type TraceInfo struct {
	ID           string         `json:"id"`
	Name         string         `json:"name"`
	Timestamp    time.Time      `json:"timestamp"`
	Metadata     map[string]any `json:"metadata,omitempty"`
	Tags         []string       `json:"tags,omitempty"`
	UserID       string         `json:"userId,omitempty"`
	SessionID    string         `json:"sessionId,omitempty"`
	Input        any            `json:"input,omitempty"`
	Output       any            `json:"output,omitempty"`
	Public       bool           `json:"public,omitempty"`
	Observations []Observation  `json:"observations,omitempty"`
	Scores       []Score        `json:"scores,omitempty"`
}

// Score represents a score from the API.
type Score struct {
	ID            string    `json:"id"`
	TraceID       string    `json:"traceId"`
	ObservationID string    `json:"observationId,omitempty"`
	Name          string    `json:"name"`
	Value         float64   `json:"value,omitempty"`
	StringValue   string    `json:"stringValue,omitempty"`
	DataType      string    `json:"dataType,omitempty"`
	Comment       string    `json:"comment,omitempty"`
	Source        string    `json:"source,omitempty"`
	Timestamp     time.Time `json:"timestamp,omitempty"`
}

// Dataset represents a dataset.
type Dataset struct {
	ID          string         `json:"id"`
	Name        string         `json:"name"`
	Description string         `json:"description,omitempty"`
	Metadata    map[string]any `json:"metadata,omitempty"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
}

// DatasetItem represents an item in a dataset.
type DatasetItem struct {
	ID                  string         `json:"id"`
	DatasetID           string         `json:"datasetId"`
	Input               any            `json:"input,omitempty"`
	ExpectedOutput      any            `json:"expectedOutput,omitempty"`
	Metadata            map[string]any `json:"metadata,omitempty"`
	SourceTraceID       string         `json:"sourceTraceId,omitempty"`
	SourceObservationID string         `json:"sourceObservationId,omitempty"`
	Status              string         `json:"status,omitempty"` // ACTIVE, ARCHIVED
	CreatedAt           time.Time      `json:"createdAt"`
	UpdatedAt           time.Time      `json:"updatedAt"`
}

// DatasetRun represents a dataset run (experiment).
type DatasetRun struct {
	ID        string         `json:"id"`
	Name      string         `json:"name"`
	DatasetID string         `json:"datasetId"`
	Metadata  map[string]any `json:"metadata,omitempty"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
}

// DatasetRunItem represents an item in a dataset run.
type DatasetRunItem struct {
	ID            string    `json:"id"`
	DatasetRunID  string    `json:"datasetRunId"`
	DatasetItemID string    `json:"datasetItemId"`
	TraceID       string    `json:"traceId"`
	ObservationID string    `json:"observationId,omitempty"`
	CreatedAt     time.Time `json:"createdAt"`
}

// Prompt represents a prompt template.
type Prompt struct {
	ID        string         `json:"id"`
	Name      string         `json:"name"`
	Version   int            `json:"version"`
	Template  string         `json:"prompt"` // The actual template content
	Config    map[string]any `json:"config,omitempty"`
	Labels    []string       `json:"labels,omitempty"`
	Tags      []string       `json:"tags,omitempty"`
	IsActive  bool           `json:"isActive,omitempty"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
}

// Project represents a project.
type Project struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// PaginatedResponse is a generic paginated response.
type PaginatedResponse[T any] struct {
	Data []T            `json:"data"`
	Meta PaginationMeta `json:"meta"`
}

// PaginationMeta contains pagination metadata.
type PaginationMeta struct {
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	TotalItems int `json:"totalItems"`
	TotalPages int `json:"totalPages"`
}

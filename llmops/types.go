package llmops

import "time"

// EvalInput represents input for evaluation.
type EvalInput struct {
	Input    any            `json:"input,omitempty"`    // User query/prompt
	Output   any            `json:"output,omitempty"`   // LLM response
	Expected any            `json:"expected,omitempty"` // Ground truth/expected output
	Context  []string       `json:"context,omitempty"`  // Retrieved context (for RAG)
	Metadata map[string]any `json:"metadata,omitempty"` // Additional metadata

	// Optional references
	TraceID string `json:"trace_id,omitempty"`
	SpanID  string `json:"span_id,omitempty"`
}

// EvalResult contains evaluation results.
type EvalResult struct {
	Scores   []MetricScore  `json:"scores"`
	Metadata map[string]any `json:"metadata,omitempty"`
	Duration time.Duration  `json:"duration,omitempty"`
}

// MetricScore represents a single metric evaluation result.
type MetricScore struct {
	Name     string  `json:"name"`
	Score    float64 `json:"score"`
	Reason   string  `json:"reason,omitempty"`
	Metadata any     `json:"metadata,omitempty"`
	Error    string  `json:"error,omitempty"`
}

// Metric defines an evaluation metric.
type Metric interface {
	// Name returns the metric name.
	Name() string

	// Evaluate computes the metric score for the given input.
	Evaluate(input EvalInput) (MetricScore, error)
}

// FeedbackScoreOpts configures feedback score creation.
type FeedbackScoreOpts struct {
	TraceID  string  `json:"trace_id,omitempty"`
	SpanID   string  `json:"span_id,omitempty"`
	Name     string  `json:"name"`
	Score    float64 `json:"score"`
	Reason   string  `json:"reason,omitempty"`
	Category string  `json:"category,omitempty"`
	Source   string  `json:"source,omitempty"`
}

// Prompt represents a prompt template.
type Prompt struct {
	ID          string         `json:"id"`
	Name        string         `json:"name"`
	Template    string         `json:"template"`
	Description string         `json:"description,omitempty"`
	Version     string         `json:"version,omitempty"`
	Tags        []string       `json:"tags,omitempty"`
	Metadata    map[string]any `json:"metadata,omitempty"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
}

// Render renders the prompt template with the given variables.
func (p *Prompt) Render(vars map[string]any) (string, error) {
	// Simple implementation - providers may override with more sophisticated templating
	result := p.Template
	for key, value := range vars {
		// Simple placeholder replacement: {{key}} -> value
		placeholder := "{{" + key + "}}"
		if str, ok := value.(string); ok {
			result = replaceAll(result, placeholder, str)
		}
	}
	return result, nil
}

// replaceAll is a simple string replacement helper.
func replaceAll(s, old, new string) string {
	// Simple implementation without importing strings to avoid dependency
	result := ""
	for {
		i := indexOf(s, old)
		if i == -1 {
			return result + s
		}
		result += s[:i] + new
		s = s[i+len(old):]
	}
}

func indexOf(s, substr string) int {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return i
		}
	}
	return -1
}

// Dataset represents a test dataset for evaluation.
type Dataset struct {
	ID          string         `json:"id"`
	Name        string         `json:"name"`
	Description string         `json:"description,omitempty"`
	Tags        []string       `json:"tags,omitempty"`
	Metadata    map[string]any `json:"metadata,omitempty"`
	ItemCount   int            `json:"item_count,omitempty"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
}

// DatasetItem represents a single item in a dataset.
type DatasetItem struct {
	ID       string         `json:"id,omitempty"`
	Input    any            `json:"input,omitempty"`
	Expected any            `json:"expected,omitempty"`
	Metadata map[string]any `json:"metadata,omitempty"`
	Tags     []string       `json:"tags,omitempty"`

	// Optional references for items created from traces
	TraceID string `json:"trace_id,omitempty"`
	SpanID  string `json:"span_id,omitempty"`
}

// Experiment represents an evaluation experiment.
type Experiment struct {
	ID          string         `json:"id"`
	Name        string         `json:"name"`
	DatasetID   string         `json:"dataset_id,omitempty"`
	DatasetName string         `json:"dataset_name,omitempty"`
	Status      string         `json:"status,omitempty"` // running, completed, cancelled
	Metadata    map[string]any `json:"metadata,omitempty"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
}

// ExperimentItem represents a single evaluation result in an experiment.
type ExperimentItem struct {
	ID            string         `json:"id,omitempty"`
	ExperimentID  string         `json:"experiment_id"`
	DatasetItemID string         `json:"dataset_item_id,omitempty"`
	Input         any            `json:"input,omitempty"`
	Output        any            `json:"output,omitempty"`
	Expected      any            `json:"expected,omitempty"`
	Scores        []MetricScore  `json:"scores,omitempty"`
	Metadata      map[string]any `json:"metadata,omitempty"`
	TraceID       string         `json:"trace_id,omitempty"`
}

// Project represents a project or workspace.
type Project struct {
	ID          string         `json:"id"`
	Name        string         `json:"name"`
	Description string         `json:"description,omitempty"`
	Metadata    map[string]any `json:"metadata,omitempty"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
}

// StreamChunk represents a chunk from streaming LLM output.
type StreamChunk struct {
	Content      string         `json:"content,omitempty"`
	TokenCount   int            `json:"token_count,omitempty"`
	Index        int            `json:"index,omitempty"`
	FinishReason string         `json:"finish_reason,omitempty"`
	Metadata     map[string]any `json:"metadata,omitempty"`
}

// StreamAccumulator accumulates streaming chunks.
type StreamAccumulator struct {
	Chunks       []StreamChunk `json:"chunks"`
	TotalContent string        `json:"total_content"`
	TotalTokens  int           `json:"total_tokens"`
	ChunkCount   int           `json:"chunk_count"`
	StartTime    time.Time     `json:"start_time"`
	FirstChunkAt *time.Time    `json:"first_chunk_at,omitempty"`
	FinishReason string        `json:"finish_reason,omitempty"`
}

// NewStreamAccumulator creates a new stream accumulator.
func NewStreamAccumulator() *StreamAccumulator {
	return &StreamAccumulator{
		Chunks:    make([]StreamChunk, 0),
		StartTime: time.Now(),
	}
}

// AddChunk adds a chunk to the accumulator.
func (a *StreamAccumulator) AddChunk(chunk StreamChunk) {
	if a.FirstChunkAt == nil {
		now := time.Now()
		a.FirstChunkAt = &now
	}
	a.Chunks = append(a.Chunks, chunk)
	a.TotalContent += chunk.Content
	a.TotalTokens += chunk.TokenCount
	a.ChunkCount++
	if chunk.FinishReason != "" {
		a.FinishReason = chunk.FinishReason
	}
}

// TimeToFirstChunk returns the duration until the first chunk was received.
func (a *StreamAccumulator) TimeToFirstChunk() time.Duration {
	if a.FirstChunkAt == nil {
		return 0
	}
	return a.FirstChunkAt.Sub(a.StartTime)
}

// TotalDuration returns the total streaming duration.
func (a *StreamAccumulator) TotalDuration() time.Duration {
	if len(a.Chunks) == 0 {
		return 0
	}
	return time.Since(a.StartTime)
}

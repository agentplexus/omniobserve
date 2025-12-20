// Package llmops provides a unified interface for LLM observability platforms.
// It abstracts common functionality across providers like Comet Opik, Arize Phoenix,
// Langfuse, Lunary, and others.
package llmops

import (
	"context"
	"io"
)

// Provider is the main interface that all LLM observability backends implement.
// It composes specialized interfaces for different capabilities.
type Provider interface {
	Tracer
	Evaluator
	PromptManager
	DatasetManager
	ProjectManager
	io.Closer

	// Name returns the provider name (e.g., "opik", "langfuse", "phoenix")
	Name() string
}

// Tracer handles trace and span operations for LLM call tracking.
type Tracer interface {
	// StartTrace begins a new trace, optionally attaching it to the context.
	StartTrace(ctx context.Context, name string, opts ...TraceOption) (context.Context, Trace, error)

	// StartSpan begins a new span within the current trace context.
	// If no trace exists in context, implementations may create one or return an error.
	StartSpan(ctx context.Context, name string, opts ...SpanOption) (context.Context, Span, error)

	// TraceFromContext retrieves the current trace from context, if any.
	TraceFromContext(ctx context.Context) (Trace, bool)

	// SpanFromContext retrieves the current span from context, if any.
	SpanFromContext(ctx context.Context) (Span, bool)
}

// Evaluator handles evaluation and scoring of LLM outputs.
type Evaluator interface {
	// Evaluate runs evaluation metrics on the given input.
	Evaluate(ctx context.Context, input EvalInput, metrics ...Metric) (*EvalResult, error)

	// AddFeedbackScore adds a feedback score to a trace or span.
	AddFeedbackScore(ctx context.Context, opts FeedbackScoreOpts) error
}

// PromptManager handles prompt template management and versioning.
type PromptManager interface {
	// CreatePrompt creates a new prompt template.
	CreatePrompt(ctx context.Context, name string, template string, opts ...PromptOption) (*Prompt, error)

	// GetPrompt retrieves a prompt by name, optionally at a specific version.
	GetPrompt(ctx context.Context, name string, version ...string) (*Prompt, error)

	// ListPrompts lists available prompts.
	ListPrompts(ctx context.Context, opts ...ListOption) ([]*Prompt, error)
}

// DatasetManager handles test dataset management.
type DatasetManager interface {
	// CreateDataset creates a new dataset for evaluation.
	CreateDataset(ctx context.Context, name string, opts ...DatasetOption) (*Dataset, error)

	// GetDataset retrieves a dataset by name.
	GetDataset(ctx context.Context, name string) (*Dataset, error)

	// AddDatasetItems adds items to a dataset.
	AddDatasetItems(ctx context.Context, datasetName string, items []DatasetItem) error

	// ListDatasets lists available datasets.
	ListDatasets(ctx context.Context, opts ...ListOption) ([]*Dataset, error)
}

// ProjectManager handles project and workspace management.
type ProjectManager interface {
	// CreateProject creates a new project.
	CreateProject(ctx context.Context, name string, opts ...ProjectOption) (*Project, error)

	// GetProject retrieves a project by name.
	GetProject(ctx context.Context, name string) (*Project, error)

	// ListProjects lists available projects.
	ListProjects(ctx context.Context, opts ...ListOption) ([]*Project, error)

	// SetProject sets the current project for subsequent operations.
	SetProject(ctx context.Context, name string) error
}

// CapabilityChecker allows querying provider capabilities.
type CapabilityChecker interface {
	// HasCapability checks if the provider supports a given capability.
	HasCapability(cap Capability) bool

	// Capabilities returns all supported capabilities.
	Capabilities() []Capability
}

// Capability represents a specific feature a provider may support.
type Capability string

const (
	CapabilityTracing      Capability = "tracing"
	CapabilityEvaluation   Capability = "evaluation"
	CapabilityPrompts      Capability = "prompts"
	CapabilityDatasets     Capability = "datasets"
	CapabilityExperiments  Capability = "experiments"
	CapabilityStreaming    Capability = "streaming"
	CapabilityDistributed  Capability = "distributed_tracing"
	CapabilityCostTracking Capability = "cost_tracking"
	CapabilityOTel         Capability = "opentelemetry"
)

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
	AnnotationManager
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

	// GetPrompt retrieves a prompt by name, optionally at a specific version or tag.
	// The version parameter can be:
	//   - Empty/omitted: returns the latest version
	//   - A tag name (e.g., "production", "staging"): returns the version with that tag
	//   - A version ID: returns that specific version (if provider supports)
	//
	// Tag-based versioning allows deployment patterns like:
	//   prompt, _ := provider.GetPrompt(ctx, "my-prompt", "production")
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

	// GetDatasetByID retrieves a dataset by ID.
	GetDatasetByID(ctx context.Context, id string) (*Dataset, error)

	// AddDatasetItems adds items to a dataset.
	AddDatasetItems(ctx context.Context, datasetName string, items []DatasetItem) error

	// ListDatasets lists available datasets.
	ListDatasets(ctx context.Context, opts ...ListOption) ([]*Dataset, error)

	// DeleteDataset deletes a dataset by ID.
	DeleteDataset(ctx context.Context, datasetID string) error
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

// AnnotationManager handles annotations on spans and traces.
type AnnotationManager interface {
	// CreateAnnotation creates an annotation on a span or trace.
	// Either SpanID or TraceID must be set in the annotation.
	CreateAnnotation(ctx context.Context, annotation Annotation) error

	// ListAnnotations lists annotations for spans or traces.
	// Provide either spanIDs or traceIDs (not both).
	ListAnnotations(ctx context.Context, opts ListAnnotationsOptions) ([]*Annotation, error)
}

// ListAnnotationsOptions configures annotation listing.
type ListAnnotationsOptions struct {
	SpanIDs  []string // List annotations for these span IDs
	TraceIDs []string // List annotations for these trace IDs
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
	CapabilityAnnotations  Capability = "annotations"
	CapabilityStreaming    Capability = "streaming"
	CapabilityDistributed  Capability = "distributed_tracing"
	CapabilityCostTracking Capability = "cost_tracking"
	CapabilityOTel         Capability = "opentelemetry"
)

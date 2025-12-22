// Package metaobserve provides a unified interface for LLM and ML observability platforms.
//
// This library abstracts common functionality across providers like:
//   - Comet Opik
//   - Arize Phoenix
//   - Langfuse
//   - Lunary
//   - MLflow
//   - Weights & Biases
//
// # Quick Start
//
// Import the provider you want to use:
//
//	import (
//		"github.com/grokify/observai/llmops"
//		_ "github.com/grokify/observai/llmops/opik"    // Register Opik
//		// or
//		_ "github.com/grokify/observai/llmops/langfuse" // Register Langfuse
//		// or
//		_ "github.com/grokify/observai/llmops/phoenix"  // Register Phoenix
//	)
//
// Then open a provider:
//
//	provider, err := llmops.Open("opik", llmops.WithAPIKey("..."))
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer provider.Close()
//
// Start tracing:
//
//	ctx, trace, _ := provider.StartTrace(ctx, "my-workflow")
//	defer trace.End()
//
//	ctx, span, _ := provider.StartSpan(ctx, "llm-call",
//		llmops.WithSpanType(llmops.SpanTypeLLM),
//		llmops.WithModel("gpt-4"),
//	)
//	defer span.End()
//
// # Architecture
//
// The library is organized into two main packages:
//
//   - llmops: LLM observability (traces, spans, evaluations, prompts)
//   - mlops: ML operations (experiments, model registry, artifacts)
//
// Each package defines interfaces that providers implement. Provider-specific
// implementations are in subpackages (e.g., llmops/opik, llmops/langfuse).
//
// # Providers
//
// Available LLM observability providers:
//
//   - opik: Comet Opik (open-source, self-hosted)
//   - langfuse: Langfuse (open-source, cloud & self-hosted)
//   - phoenix: Arize Phoenix (open-source, uses OpenTelemetry)
//
// # Features
//
// Common features across providers:
//
//   - Trace/span creation and context propagation
//   - Input/output capture
//   - Token usage and cost tracking
//   - Feedback scores and evaluations
//   - Dataset management
//   - Prompt versioning (provider-dependent)
//
// # SDK Access
//
// For provider-specific features, you can use the underlying SDKs directly:
//
//	import "github.com/grokify/metaobserve/sdk/langfuse"
//	import "github.com/grokify/metaobserve/sdk/phoenix"
//
// Or use the existing Opik SDK:
//
//	import "github.com/grokify/go-comet-ml-opik"
package observai

import (
	"github.com/grokify/metaobserve/llmops"
	"github.com/grokify/metaobserve/mlops"
)

// Version is the library version.
const Version = "0.1.0"

// Re-export commonly used types for convenience.
type (
	// Provider is an alias for llmops.Provider.
	Provider = llmops.Provider

	// Trace is an alias for llmops.Trace.
	Trace = llmops.Trace

	// Span is an alias for llmops.Span.
	Span = llmops.Span

	// TokenUsage is an alias for llmops.TokenUsage.
	TokenUsage = llmops.TokenUsage

	// SpanType is an alias for llmops.SpanType.
	SpanType = llmops.SpanType

	// EvalInput is an alias for llmops.EvalInput.
	EvalInput = llmops.EvalInput

	// EvalResult is an alias for llmops.EvalResult.
	EvalResult = llmops.EvalResult

	// MLProvider is an alias for mlops.Provider.
	MLProvider = mlops.Provider

	// Experiment is an alias for mlops.Experiment.
	Experiment = mlops.Experiment

	// Run is an alias for mlops.Run.
	Run = mlops.Run

	// Model is an alias for mlops.Model.
	Model = mlops.Model
)

// Span type constants.
const (
	SpanTypeGeneral   = llmops.SpanTypeGeneral
	SpanTypeLLM       = llmops.SpanTypeLLM
	SpanTypeTool      = llmops.SpanTypeTool
	SpanTypeRetrieval = llmops.SpanTypeRetrieval
	SpanTypeAgent     = llmops.SpanTypeAgent
	SpanTypeChain     = llmops.SpanTypeChain
	SpanTypeGuardrail = llmops.SpanTypeGuardrail
)

// OpenLLMOps opens an LLM observability provider.
// This is a convenience function that wraps llmops.Open.
func OpenLLMOps(name string, opts ...llmops.ClientOption) (llmops.Provider, error) {
	return llmops.Open(name, opts...)
}

// Providers returns the names of registered LLM providers.
func Providers() []string {
	return llmops.Providers()
}

// Re-export option functions for convenience.
var (
	// Client options
	WithAPIKey      = llmops.WithAPIKey
	WithEndpoint    = llmops.WithEndpoint
	WithWorkspace   = llmops.WithWorkspace
	WithProjectName = llmops.WithProjectName
	WithHTTPClient  = llmops.WithHTTPClient
	WithTimeout     = llmops.WithTimeout
	WithDisabled    = llmops.WithDisabled
	WithDebug       = llmops.WithDebug

	// Trace options
	WithTraceProject  = llmops.WithTraceProject
	WithTraceInput    = llmops.WithTraceInput
	WithTraceOutput   = llmops.WithTraceOutput
	WithTraceMetadata = llmops.WithTraceMetadata
	WithTraceTags     = llmops.WithTraceTags
	WithThreadID      = llmops.WithThreadID

	// Span options
	WithSpanType     = llmops.WithSpanType
	WithSpanInput    = llmops.WithSpanInput
	WithSpanOutput   = llmops.WithSpanOutput
	WithSpanMetadata = llmops.WithSpanMetadata
	WithSpanTags     = llmops.WithSpanTags
	WithModel        = llmops.WithModel
	WithProvider     = llmops.WithProvider
	WithTokenUsage   = llmops.WithTokenUsage

	// End options
	WithEndOutput   = llmops.WithEndOutput
	WithEndMetadata = llmops.WithEndMetadata
	WithEndError    = llmops.WithEndError
)

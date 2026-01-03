# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.5.0] - 2026-01-03

### Added

- `AnnotationManager` interface for span/trace annotations
  - `CreateAnnotation` for adding annotations to spans or traces
  - `ListAnnotations` for querying annotations by span or trace IDs
- `DatasetManager` new methods
  - `GetDatasetByID` for retrieving datasets by ID
  - `DeleteDataset` for removing datasets
- Prompt model/provider options
  - `WithPromptModel` and `WithPromptProvider` options
  - `ModelName` and `ModelProvider` fields on `Prompt` type
- OmniLLM hook auto-creates traces when none exists in context
- Trace context helpers (`contextWithTrace`, `traceFromContext`)
- `llmops/metrics` package with evaluation metrics
  - LLM-based: `HallucinationMetric`, `RelevanceMetric`, `QACorrectnessMetric`, `ToxicityMetric`
  - Code-based: `ExactMatchMetric`, `RegexMetric`, `ContainsMetric`
- `examples/evaluation` - Example demonstrating metrics usage

### Changed

- Provider adapters moved to standalone SDKs
  - Opik: `github.com/agentplexus/go-opik/llmops`
  - Phoenix: `github.com/agentplexus/go-phoenix/llmops`
  - Langfuse remains in omniobserve

### Removed

- `llmops/opik` adapter (moved to go-opik)
- `llmops/phoenix` adapter (moved to go-phoenix)
- `sdk/phoenix` package (use go-phoenix directly)

### Migration

Update imports:

```go
// Before
import _ "github.com/agentplexus/omniobserve/llmops/opik"
import _ "github.com/agentplexus/omniobserve/llmops/phoenix"

// After
import _ "github.com/agentplexus/go-opik/llmops"
import _ "github.com/agentplexus/go-phoenix/llmops"
```

## [0.4.0] - Previous Release

See git history for earlier changes.

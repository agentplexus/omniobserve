# ObservAI

[![Build Status][build-status-svg]][build-status-url]
[![Lint Status][lint-status-svg]][lint-status-url]
[![Go Report Card][goreport-svg]][goreport-url]
[![Docs][docs-godoc-svg]][docs-godoc-url]
[![License][license-svg]][license-url]

A unified Go library for LLM and ML observability. ObservAI provides a vendor-agnostic abstraction layer that enables you to instrument your AI applications once and seamlessly switch between different observability backends without code changes.

## Features

- 🔗 **Unified Interface**: Single API for tracing, evaluation, prompts, and datasets across all providers
- 🔄 **Provider Agnostic**: Switch between Opik, Langfuse, and Phoenix without changing your code
- 🔍 **Full Tracing**: Trace LLM calls with spans, token usage, and cost tracking
- 📊 **Evaluation Support**: Run metrics and add feedback scores to traces
- 📦 **Dataset Management**: Create and manage evaluation datasets
- 📝 **Prompt Versioning**: Store and version prompt templates (provider-dependent)
- 🔀 **Context Propagation**: Automatic trace/span context propagation via `context.Context`
- ⚙️ **Functional Options**: Clean, extensible configuration using the options pattern

## Installation

```bash
go get github.com/grokify/observai
```

## Quick Start

```go
package main

import (
    "context"
    "log"

    "github.com/grokify/observai/llmops"
    _ "github.com/grokify/observai/llmops/opik"  // Register Opik provider
)

func main() {
    // Open a provider
    provider, err := llmops.Open("opik",
        llmops.WithAPIKey("your-api-key"),
        llmops.WithProjectName("my-project"),
    )
    if err != nil {
        log.Fatal(err)
    }
    defer provider.Close()

    ctx := context.Background()

    // Start a trace
    ctx, trace, err := provider.StartTrace(ctx, "chat-workflow",
        llmops.WithTraceInput(map[string]any{"query": "Hello, world!"}),
    )
    if err != nil {
        log.Fatal(err)
    }
    defer trace.End()

    // Start a span for the LLM call
    ctx, span, err := provider.StartSpan(ctx, "gpt-4-completion",
        llmops.WithSpanType(llmops.SpanTypeLLM),
        llmops.WithModel("gpt-4"),
        llmops.WithProvider("openai"),
    )
    if err != nil {
        log.Fatal(err)
    }

    // Record the LLM interaction
    span.SetInput(map[string]any{
        "messages": []map[string]string{
            {"role": "user", "content": "Hello!"},
        },
    })

    // ... call your LLM here ...

    span.SetOutput(map[string]any{
        "response": "Hello! How can I help you today?",
    })
    span.SetUsage(llmops.TokenUsage{
        PromptTokens:     10,
        CompletionTokens: 8,
        TotalTokens:      18,
    })

    span.End()
    trace.SetOutput(map[string]any{"response": "Hello! How can I help you today?"})
}
```

## Supported Providers

| Provider | Package | Description |
|----------|---------|-------------|
| **Opik** | `llmops/opik` | Comet Opik - Open-source, full-featured |
| **Langfuse** | `llmops/langfuse` | Cloud & self-hosted, batch ingestion |
| **Phoenix** | `llmops/phoenix` | Arize Phoenix - OpenTelemetry-based |

### Provider Capabilities

| Feature | Opik | Langfuse | Phoenix |
|---------|:----:|:--------:|:-------:|
| Tracing | :white_check_mark: | :white_check_mark: | :white_check_mark: |
| Evaluation | :white_check_mark: | :white_check_mark: | :white_check_mark: |
| Prompts | :white_check_mark: | Partial | :x: |
| Datasets | :white_check_mark: | :white_check_mark: | Partial |
| Experiments | :white_check_mark: | :white_check_mark: | Partial |
| Streaming | :white_check_mark: | :white_check_mark: | Planned |
| Distributed Tracing | :white_check_mark: | :x: | :white_check_mark: |
| Cost Tracking | :white_check_mark: | :white_check_mark: | :x: |
| OpenTelemetry | :x: | :x: | :white_check_mark: |

## Architecture

```
observai/
├── observai.go          # Main package with re-exports
├── llmops/              # LLM observability interfaces
│   ├── llmops.go        # Core interfaces (Provider, Tracer, Evaluator, etc.)
│   ├── trace.go         # Trace and Span interfaces
│   ├── types.go         # Data types (EvalInput, Dataset, Prompt, etc.)
│   ├── options.go       # Functional options
│   ├── provider.go      # Provider registration system
│   ├── errors.go        # Error definitions
│   ├── opik/            # Opik provider adapter
│   ├── langfuse/        # Langfuse provider adapter
│   └── phoenix/         # Phoenix provider adapter
├── integrations/        # Integrations with LLM libraries
│   └── fluxllm/         # FluxLLM observability hook (separate module)
├── mlops/               # ML operations interfaces (experiments, model registry)
└── sdk/                 # Provider-specific SDKs
    ├── langfuse/        # Langfuse Go SDK
    └── phoenix/         # Phoenix Go SDK
```

## Core Interfaces

### Provider

The main interface that all observability backends implement:

```go
type Provider interface {
    Tracer           // Trace/span operations
    Evaluator        // Evaluation and feedback
    PromptManager    // Prompt template management
    DatasetManager   // Test dataset management
    ProjectManager   // Project/workspace management
    io.Closer

    Name() string
}
```

### Trace and Span

```go
type Trace interface {
    ID() string
    Name() string
    StartSpan(ctx context.Context, name string, opts ...SpanOption) (context.Context, Span, error)
    SetInput(input any)
    SetOutput(output any)
    SetMetadata(metadata map[string]any)
    AddTag(key, value string)
    AddFeedbackScore(ctx context.Context, name string, score float64, opts ...FeedbackOption) error
    End(opts ...EndOption)
}

type Span interface {
    ID() string
    TraceID() string
    Name() string
    Type() SpanType
    StartSpan(ctx context.Context, name string, opts ...SpanOption) (context.Context, Span, error)
    SetInput(input any)
    SetOutput(output any)
    SetModel(model string)
    SetProvider(provider string)
    SetUsage(usage TokenUsage)
    End(opts ...EndOption)
}
```

### Span Types

```go
const (
    SpanTypeGeneral   SpanType = "general"
    SpanTypeLLM       SpanType = "llm"
    SpanTypeTool      SpanType = "tool"
    SpanTypeRetrieval SpanType = "retrieval"
    SpanTypeAgent     SpanType = "agent"
    SpanTypeChain     SpanType = "chain"
    SpanTypeGuardrail SpanType = "guardrail"
)
```

## Usage Examples

### Using Different Providers

```go
// Opik
import _ "github.com/grokify/observai/llmops/opik"
provider, _ := llmops.Open("opik", llmops.WithAPIKey("..."))

// Langfuse
import _ "github.com/grokify/observai/llmops/langfuse"
provider, _ := llmops.Open("langfuse",
    llmops.WithAPIKey("sk-lf-..."),
    llmops.WithEndpoint("https://cloud.langfuse.com"),
)

// Phoenix
import _ "github.com/grokify/observai/llmops/phoenix"
provider, _ := llmops.Open("phoenix",
    llmops.WithEndpoint("http://localhost:6006"),
)
```

### Nested Spans

```go
ctx, trace, _ := provider.StartTrace(ctx, "rag-pipeline")
defer trace.End()

// Retrieval span
ctx, retrievalSpan, _ := provider.StartSpan(ctx, "vector-search",
    llmops.WithSpanType(llmops.SpanTypeRetrieval),
)
// ... perform retrieval ...
retrievalSpan.SetOutput(documents)
retrievalSpan.End()

// LLM span
ctx, llmSpan, _ := provider.StartSpan(ctx, "generate-response",
    llmops.WithSpanType(llmops.SpanTypeLLM),
    llmops.WithModel("gpt-4"),
)
// ... call LLM ...
llmSpan.SetUsage(llmops.TokenUsage{
    PromptTokens:     150,
    CompletionTokens: 50,
    TotalTokens:      200,
})
llmSpan.End()
```

### Adding Feedback Scores

```go
// Add a score to a span
span.AddFeedbackScore(ctx, "relevance", 0.95,
    llmops.WithFeedbackReason("Response directly addressed the query"),
    llmops.WithFeedbackCategory("quality"),
)

// Add a score to a trace
trace.AddFeedbackScore(ctx, "user_satisfaction", 0.8)
```

### Working with Datasets

```go
// Create a dataset
dataset, _ := provider.CreateDataset(ctx, "test-cases",
    llmops.WithDatasetDescription("Test cases for RAG evaluation"),
)

// Add items
provider.AddDatasetItems(ctx, "test-cases", []llmops.DatasetItem{
    {
        Input:    map[string]any{"query": "What is Go?"},
        Expected: map[string]any{"answer": "Go is a programming language..."},
    },
})
```

### Working with Prompts (Opik)

```go
// Create a versioned prompt
prompt, _ := provider.CreatePrompt(ctx, "chat-template",
    `You are a helpful assistant. User: {{.query}}`,
    llmops.WithPromptDescription("Main chat template"),
)

// Get a prompt
prompt, _ := provider.GetPrompt(ctx, "chat-template")

// Render with variables
rendered := prompt.Render(map[string]any{"query": "Hello!"})
```

## Configuration Options

### Client Options

```go
llmops.WithAPIKey("...")           // API key for authentication
llmops.WithEndpoint("...")         // Custom endpoint URL
llmops.WithWorkspace("...")        // Workspace/organization name
llmops.WithProjectName("...")      // Default project name
llmops.WithHTTPClient(client)      // Custom HTTP client
llmops.WithTimeout(30 * time.Second)
llmops.WithDisabled(true)          // Disable tracing (no-op mode)
llmops.WithDebug(true)             // Enable debug logging
```

### Trace Options

```go
llmops.WithTraceProject("...")
llmops.WithTraceInput(input)
llmops.WithTraceOutput(output)
llmops.WithTraceMetadata(map[string]any{...})
llmops.WithTraceTags(map[string]string{...})
llmops.WithThreadID("...")
```

### Span Options

```go
llmops.WithSpanType(llmops.SpanTypeLLM)
llmops.WithSpanInput(input)
llmops.WithSpanOutput(output)
llmops.WithSpanMetadata(map[string]any{...})
llmops.WithModel("gpt-4")
llmops.WithProvider("openai")
llmops.WithTokenUsage(usage)
llmops.WithParentSpan(parentSpan)
```

## Error Handling

The library provides typed errors for common conditions:

```go
if errors.Is(err, llmops.ErrMissingAPIKey) {
    // Handle missing API key
}

if llmops.IsNotFound(err) {
    // Handle not found
}

if llmops.IsRateLimited(err) {
    // Handle rate limiting
}
```

## Direct SDK Access

For provider-specific features, you can use the underlying SDKs directly:

```go
import "github.com/grokify/observai/sdk/langfuse"
import "github.com/grokify/observai/sdk/phoenix"
```

## FluxLLM Integration

ObservAI provides an integration with [FluxLLM](https://github.com/grokify/fluxllm), a multi-LLM abstraction layer. This allows you to automatically instrument all LLM calls made through FluxLLM with any ObservAI provider.

```bash
go get github.com/grokify/observai/integrations/fluxllm
```

```go
package main

import (
    "github.com/grokify/fluxllm"
    fluxllmhook "github.com/grokify/observai/integrations/fluxllm"
    "github.com/grokify/observai/llmops"
    _ "github.com/grokify/observai/llmops/opik"
)

func main() {
    // Initialize an ObservAI provider
    provider, _ := llmops.Open("opik",
        llmops.WithAPIKey("your-api-key"),
        llmops.WithProjectName("my-project"),
    )
    defer provider.Close()

    // Create the observability hook
    hook := fluxllmhook.NewHook(provider)

    // Attach to your FluxLLM client
    client := fluxllm.NewClient(
        fluxllm.WithObservabilityHook(hook),
    )

    // All LLM calls through this client are now automatically traced
}
```

The hook automatically captures:
- Model and provider information
- Input messages and output responses
- Token usage (prompt, completion, total)
- Streaming responses
- Errors

## Requirements

- Go 1.24.5 or later

## Contributing

Contributions are welcome! Please feel free to submit issues and pull requests.

## License

See [LICENSE](LICENSE) for details.

 [build-status-svg]: https://github.com/grokify/observai/actions/workflows/ci.yaml/badge.svg?branch=main
 [build-status-url]: https://github.com/grokify/observai/actions/workflows/ci.yaml
 [lint-status-svg]: https://github.com/grokify/observai/actions/workflows/lint.yaml/badge.svg?branch=main
 [lint-status-url]: https://github.com/grokify/observai/actions/workflows/lint.yaml
 [goreport-svg]: https://goreportcard.com/badge/github.com/grokify/observai
 [goreport-url]: https://goreportcard.com/report/github.com/grokify/observai
 [docs-godoc-svg]: https://pkg.go.dev/badge/github.com/grokify/observai
 [docs-godoc-url]: https://pkg.go.dev/github.com/grokify/observai
 [license-svg]: https://img.shields.io/badge/license-MIT-blue.svg
 [license-url]: https://github.com/grokify/observai/blob/master/LICENSE
 [used-by-svg]: https://sourcegraph.com/github.com/grokify/observai/-/badge.svg
 [used-by-url]: https://sourcegraph.com/github.com/grokify/observai?badge

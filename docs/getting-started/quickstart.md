# Quick Start

This page covers two systems: **observops** for app observability (metrics, traces, logs)
and **llmops** for LLM/ML tracing. Start with whichever you need.

## App Observability (observops.Setup)

`observops.Setup` is the recommended entry point: one call wires metrics, traces, and logs
and returns native OpenTelemetry handles, so you instrument with the standard OTel API.

```go
package main

import (
    "context"
    "log"
    "net/http"

    "github.com/plexusone/omniobserve/observops"
)

func main() {
    ctx := context.Background()

    tel, err := observops.Setup(ctx,
        observops.WithServiceName("my-service"),
        observops.WithServiceVersion("1.2.3"),
        observops.WithPrometheus(),               // pull endpoint: tel.MetricsHandler
        observops.WithEndpoint("localhost:4317"), // OTLP push (gRPC)
        observops.WithInsecure(),
    )
    if err != nil {
        log.Fatal(err)
    }
    defer tel.Shutdown(ctx)

    // Metrics — native OTel meter (int64/float64, sync/observable).
    requests, _ := tel.Meter.Int64Counter("requests.total")
    requests.Add(ctx, 1)

    // Tracing — native OTel tracer.
    ctx, span := tel.Tracer.Start(ctx, "handle-request")
    defer span.End()

    // Logging — exported via OpenTelemetry and echoed to the console.
    tel.Logger.InfoContext(ctx, "request processed", "user_id", "123")

    // Serve Prometheus metrics and instrument an HTTP handler.
    mux := http.NewServeMux()
    mux.Handle("/metrics", tel.MetricsHandler)
    handler := tel.Middleware("api")(mux)

    log.Fatal(http.ListenAndServe(":8080", handler))
}
```

Metrics, traces, and logs are enabled by default; disable any with `WithMetrics`,
`WithTraces`, or `WithLogs`. Prometheus pull and OTLP push compose; add `WithStdout` to mirror
telemetry to stdout while debugging. See [Setup (recommended)](../providers/setup.md) for the
full option reference, and the [driver providers](../providers/otlp.md) for the vendor-neutral
`Open` API.

## LLM Tracing

### Basic Tracing

```go
package main

import (
    "context"
    "log"

    "github.com/plexusone/omniobserve/llmops"
    _ "github.com/agentplexus/go-opik/llmops"  // Register Opik provider
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

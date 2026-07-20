# OmniObserve

**Unified Go library for observability**

OmniObserve provides vendor-agnostic abstraction layers for observability, enabling you to instrument your applications once and seamlessly switch between different backends without code changes.

## Two Provider Systems

| Package | Purpose | Providers |
|---------|---------|-----------|
| **llmops** | LLM/ML observability | Opik, Langfuse, Phoenix, slog |
| **observops** | App observability (metrics, traces, logs) | OTLP, Datadog, New Relic, Dynatrace |

## Features

### LLM Observability (llmops)

- **Unified Interface**: Single API for tracing, evaluation, prompts, and datasets across all providers
- **Provider Agnostic**: Switch between Opik, Langfuse, Phoenix, and slog without changing your code
- **Full Tracing**: Trace LLM calls with spans, token usage, and cost tracking
- **Evaluation Support**: Run metrics and add feedback scores to traces
- **Dataset Management**: Create and manage evaluation datasets
- **Prompt Versioning**: Store and version prompt templates (provider-dependent)

### App Observability (observops)

- **One-Call Setup**: `observops.Setup` wires metrics, traces, and logs and returns native OpenTelemetry handles — full ecosystem interop, no wrapper tax
- **Composable Exporters**: Prometheus pull and OTLP push (gRPC or HTTP), plus stdout for debugging
- **Full Telemetry**: Metrics, distributed traces, and real OpenTelemetry logs
- **HTTP Instrumentation**: One-line server middleware and client transport
- **Vendor-Agnostic Drivers**: Single interface for OTLP, Datadog, New Relic, and Dynatrace

### Common

- **Context Propagation**: Automatic trace/span context propagation via `context.Context`
- **Functional Options**: Clean, extensible configuration using the options pattern

## Quick Examples

### LLM Observability (llmops)

```go
import (
    "github.com/plexusone/omniobserve/llmops"
    _ "github.com/agentplexus/go-opik/llmops"  // Register Opik provider
)

provider, _ := llmops.Open("opik",
    llmops.WithAPIKey("your-api-key"),
    llmops.WithProjectName("my-project"),
)
defer provider.Close()

ctx, trace, _ := provider.StartTrace(ctx, "chat-workflow")
defer trace.End()

_, span, _ := provider.StartSpan(ctx, "gpt-4-completion",
    llmops.WithSpanType(llmops.SpanTypeLLM),
    llmops.WithModel("gpt-4"),
)
span.SetUsage(llmops.TokenUsage{TotalTokens: 18})
span.End()
```

### App Observability (observops)

```go
import "github.com/plexusone/omniobserve/observops"

tel, _ := observops.Setup(ctx,
    observops.WithServiceName("my-service"),
    observops.WithPrometheus(),               // pull: tel.MetricsHandler
    observops.WithEndpoint("localhost:4317"), // OTLP push
    observops.WithInsecure(),
)
defer tel.Shutdown(ctx)

// Native OpenTelemetry handles:
reqs, _ := tel.Meter.Int64Counter("requests.total")
reqs.Add(ctx, 1)

ctx, span := tel.Tracer.Start(ctx, "handle-request")
defer span.End()

tel.Logger.InfoContext(ctx, "request processed", "user_id", "123")

// Prometheus endpoint + HTTP instrumentation:
mux.Handle("/metrics", tel.MetricsHandler)
handler = tel.Middleware("api")(handler)
```

See [Setup (recommended)](providers/setup.md) for the full reference, or the
[driver providers](providers/otlp.md) for the vendor-neutral `Open` API.

## Supported Providers

### LLM Providers (llmops)

| Provider | Package | Description |
|----------|---------|-------------|
| [Opik](providers/opik.md) | `go-opik/llmops` | Comet Opik - Open-source, full-featured |
| [Langfuse](providers/langfuse.md) | `omniobserve/llmops/langfuse` | Cloud & self-hosted, batch ingestion |
| [Phoenix](providers/phoenix.md) | `go-phoenix/llmops` | Arize Phoenix - OpenTelemetry-based |
| [slog](providers/slog.md) | `omniobserve/llmops/slog` | Local structured logging for development |

### App Observability Providers (observops)

| Provider | Package | Description |
|----------|---------|-------------|
| [OTLP](providers/otlp.md) | `omniobserve/observops/otlp` | OpenTelemetry Protocol - vendor-agnostic |
| [Datadog](providers/datadog.md) | `omniobserve/observops/datadog` | Datadog APM via OTLP |
| [New Relic](providers/newrelic.md) | `omniobserve/observops/newrelic` | New Relic via OTLP |
| [Dynatrace](providers/dynatrace.md) | `omniobserve/observops/dynatrace` | Dynatrace via OTLP |

## Next Steps

- [Installation](getting-started/installation.md) - Get OmniObserve set up
- [Quick Start](getting-started/quickstart.md) - Your first trace
- [Providers](providers/index.md) - Configure specific providers
- [OmniLLM Integration](integrations/omnillm.md) - Auto-instrument LLM calls

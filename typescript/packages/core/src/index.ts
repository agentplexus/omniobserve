/**
 * @omniobserve/core
 *
 * TypeScript interfaces and semantic conventions for OmniObserve observability.
 * This package mirrors the Go implementation in omniobserve.
 *
 * @example Basic usage
 * ```typescript
 * import type { Trace, Span, Provider } from '@omniobserve/core';
 * import { semconv } from '@omniobserve/core';
 *
 * // Use semantic convention constants
 * console.log(semconv.agent.TaskID); // 'gen_ai.agent.task.id'
 * console.log(semconv.journey.SessionID); // 'session.id'
 * ```
 *
 * @example Implementing a provider
 * ```typescript
 * import type { Provider, Trace, Span, TraceOptions, SpanOptions } from '@omniobserve/core';
 *
 * class MyProvider implements Provider {
 *   name = 'my-provider';
 *
 *   async startTrace(name: string, options?: TraceOptions): Promise<Trace> {
 *     // Implementation
 *   }
 *
 *   async startSpan(trace: Trace, name: string, options?: SpanOptions): Promise<Span> {
 *     // Implementation
 *   }
 *
 *   // ... other methods
 * }
 * ```
 *
 * @packageDocumentation
 */

// =============================================================================
// Core Types
// =============================================================================

export type {
  // Span types
  SpanType,
  TokenUsage,
  Trace,
  Span,
  FeedbackScore,

  // Evaluation types
  EvalInput,
  MetricScore,
  EvalMetric,

  // Dataset types
  Dataset,
  DatasetItem,

  // Prompt types
  Prompt,

  // Status types
  Status,
  HandoffStatus,
  HandoffType,
  Severity,
  EventCategory,

  // AgentOps types
  Workflow,
  Task,
  Handoff,
  ToolInvocation,
  AgentEvent,

  // Options types
  TraceOptions,
  SpanOptions,

  // Provider interfaces
  Tracer,
  Evaluator,
  DatasetManager,
  PromptManager,
  Provider,
} from './types';

// =============================================================================
// Semantic Conventions
// =============================================================================

export * as semconv from './semconv';

// Direct exports for tree-shaking
export * as agentSemconv from './semconv/agent';
export * as journeySemconv from './semconv/journey';

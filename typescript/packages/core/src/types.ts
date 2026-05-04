/**
 * Core types for OmniObserve TypeScript SDK
 *
 * These interfaces mirror the Go implementation in omniobserve/llmops
 * and omniobserve/agentops packages.
 */

// =============================================================================
// Trace & Span Types (from llmops)
// =============================================================================

/**
 * Span type categorizes the type of operation being traced
 */
export type SpanType =
  | 'general'
  | 'llm'
  | 'tool'
  | 'retrieval'
  | 'agent'
  | 'chain'
  | 'guardrail';

/**
 * Token usage information for LLM calls
 */
export interface TokenUsage {
  promptTokens?: number;
  completionTokens?: number;
  totalTokens?: number;
}

/**
 * Trace represents a top-level execution session
 */
export interface Trace {
  id: string;
  name: string;
  traceId: string;
  input?: unknown;
  output?: unknown;
  metadata?: Record<string, unknown>;
  tags?: string[];
  startTime: Date;
  endTime?: Date;
}

/**
 * Span represents a unit of work within a trace
 */
export interface Span {
  id: string;
  traceId: string;
  parentSpanId?: string;
  name: string;
  type: SpanType;
  input?: unknown;
  output?: unknown;
  model?: string;
  provider?: string;
  usage?: TokenUsage;
  metadata?: Record<string, unknown>;
  startTime: Date;
  endTime?: Date;
  durationMs?: number;
}

/**
 * Feedback score for evaluation
 */
export interface FeedbackScore {
  traceId?: string;
  spanId?: string;
  name: string;
  score: number; // 0-1 range
  reason?: string;
  category?: string;
  source?: 'human' | 'llm' | 'code';
}

// =============================================================================
// Evaluation Types (from llmops/metrics)
// =============================================================================

/**
 * Input for evaluation metrics
 */
export interface EvalInput {
  input?: unknown;
  output?: unknown;
  expected?: unknown;
  context?: string[];
  metadata?: Record<string, unknown>;
}

/**
 * Result from an evaluation metric
 */
export interface MetricScore {
  name: string;
  score: number;
  reason?: string;
  metadata?: Record<string, unknown>;
}

/**
 * Evaluation metric interface
 */
export interface EvalMetric {
  name: string;
  evaluate(input: EvalInput): Promise<MetricScore>;
}

// =============================================================================
// Dataset Types (from llmops)
// =============================================================================

/**
 * Dataset for evaluation testing
 */
export interface Dataset {
  id: string;
  name: string;
  description?: string;
  metadata?: Record<string, unknown>;
  createdAt: Date;
  updatedAt: Date;
}

/**
 * Item within a dataset
 */
export interface DatasetItem {
  id: string;
  datasetId: string;
  input: unknown;
  expected?: unknown;
  metadata?: Record<string, unknown>;
}

// =============================================================================
// Prompt Types (from llmops)
// =============================================================================

/**
 * Prompt template for LLM interactions
 */
export interface Prompt {
  id: string;
  name: string;
  template: string;
  version?: string;
  tags?: string[];
  model?: string;
  provider?: string;
  metadata?: Record<string, unknown>;
}

// =============================================================================
// Workflow Types (from agentops)
// =============================================================================

/**
 * Status for workflows, tasks, and handoffs
 */
export type Status = 'pending' | 'running' | 'completed' | 'failed' | 'cancelled';

/**
 * Handoff status values
 */
export type HandoffStatus = Status | 'accepted' | 'rejected';

/**
 * Handoff type values
 */
export type HandoffType = 'request' | 'response' | 'broadcast' | 'delegate';

/**
 * Event severity levels
 */
export type Severity = 'debug' | 'info' | 'warn' | 'error';

/**
 * Event category values
 */
export type EventCategory = 'agent' | 'workflow' | 'tool' | 'domain' | 'system';

/**
 * Workflow represents an end-to-end execution session
 */
export interface Workflow {
  id: string;
  name: string;
  status: Status;
  traceId?: string;
  parentWorkflowId?: string;
  initiator?: string;
  input?: Record<string, unknown>;
  output?: Record<string, unknown>;
  metadata?: Record<string, unknown>;
  taskCount: number;
  completedTaskCount: number;
  failedTaskCount: number;
  totalCostUSD: number;
  totalTokens: number;
  durationMs?: number;
  errorMessage?: string;
  startedAt: Date;
  endedAt?: Date;
}

/**
 * Task represents an agent task within a workflow
 */
export interface Task {
  id: string;
  workflowId?: string;
  agentId: string;
  agentType?: string;
  taskType: string;
  name: string;
  status: Status;
  traceId?: string;
  spanId?: string;
  parentSpanId?: string;
  input?: Record<string, unknown>;
  output?: Record<string, unknown>;
  metadata?: Record<string, unknown>;
  llmCallCount: number;
  toolCallCount: number;
  retryCount: number;
  tokensPrompt: number;
  tokensCompletion: number;
  tokensTotal: number;
  costUSD: number;
  durationMs?: number;
  errorType?: string;
  errorMessage?: string;
  startedAt: Date;
  endedAt?: Date;
}

/**
 * Handoff represents agent-to-agent communication
 */
export interface Handoff {
  id: string;
  workflowId?: string;
  fromAgentId: string;
  fromAgentType?: string;
  toAgentId: string;
  toAgentType?: string;
  handoffType: HandoffType;
  status: HandoffStatus;
  traceId?: string;
  fromTaskId?: string;
  toTaskId?: string;
  payload?: Record<string, unknown>;
  metadata?: Record<string, unknown>;
  payloadSizeBytes: number;
  latencyMs?: number;
  errorMessage?: string;
  initiatedAt: Date;
  acceptedAt?: Date;
  completedAt?: Date;
}

/**
 * ToolInvocation represents a tool/function call
 */
export interface ToolInvocation {
  id: string;
  taskId?: string;
  agentId: string;
  toolName: string;
  toolType?: string;
  status: Status;
  traceId?: string;
  spanId?: string;
  input?: Record<string, unknown>;
  output?: Record<string, unknown>;
  metadata?: Record<string, unknown>;
  httpMethod?: string;
  httpUrl?: string;
  httpStatusCode?: number;
  durationMs?: number;
  requestSizeBytes: number;
  responseSizeBytes: number;
  retryCount: number;
  errorType?: string;
  errorMessage?: string;
  startedAt: Date;
  endedAt?: Date;
}

/**
 * Event represents a generic event
 */
export interface AgentEvent {
  id: string;
  eventType: string;
  eventCategory: EventCategory;
  workflowId?: string;
  taskId?: string;
  agentId?: string;
  traceId?: string;
  spanId?: string;
  severity: Severity;
  data?: Record<string, unknown>;
  metadata?: Record<string, unknown>;
  tags?: string[];
  source?: string;
  version?: string;
  timestamp: Date;
}

// =============================================================================
// Provider Interfaces (from llmops)
// =============================================================================

/**
 * Options for starting a trace
 */
export interface TraceOptions {
  input?: unknown;
  metadata?: Record<string, unknown>;
  tags?: string[];
}

/**
 * Options for starting a span
 */
export interface SpanOptions {
  type?: SpanType;
  input?: unknown;
  model?: string;
  provider?: string;
  metadata?: Record<string, unknown>;
}

/**
 * Tracer interface for creating traces and spans
 */
export interface Tracer {
  /**
   * Start a new trace
   */
  startTrace(name: string, options?: TraceOptions): Promise<Trace>;

  /**
   * Start a span within a trace
   */
  startSpan(trace: Trace, name: string, options?: SpanOptions): Promise<Span>;

  /**
   * Start a child span
   */
  startChildSpan(parentSpan: Span, name: string, options?: SpanOptions): Promise<Span>;

  /**
   * End a trace
   */
  endTrace(trace: Trace, output?: unknown): Promise<void>;

  /**
   * End a span
   */
  endSpan(span: Span, output?: unknown, usage?: TokenUsage): Promise<void>;
}

/**
 * Evaluator interface for feedback and scoring
 */
export interface Evaluator {
  /**
   * Add feedback score to a trace or span
   */
  addFeedbackScore(score: FeedbackScore): Promise<void>;

  /**
   * Run an evaluation metric
   */
  evaluate(metric: EvalMetric, input: EvalInput): Promise<MetricScore>;
}

/**
 * DatasetManager interface for managing test datasets
 */
export interface DatasetManager {
  /**
   * Create a new dataset
   */
  createDataset(
    name: string,
    options?: { description?: string; metadata?: Record<string, unknown> }
  ): Promise<Dataset>;

  /**
   * Get a dataset by ID
   */
  getDataset(id: string): Promise<Dataset | null>;

  /**
   * List all datasets
   */
  listDatasets(): Promise<Dataset[]>;

  /**
   * Add an item to a dataset
   */
  addDatasetItem(datasetId: string, item: Omit<DatasetItem, 'id' | 'datasetId'>): Promise<DatasetItem>;

  /**
   * Get items from a dataset
   */
  getDatasetItems(datasetId: string): Promise<DatasetItem[]>;
}

/**
 * PromptManager interface for managing prompts
 */
export interface PromptManager {
  /**
   * Create a new prompt
   */
  createPrompt(
    name: string,
    template: string,
    options?: { version?: string; tags?: string[]; model?: string; metadata?: Record<string, unknown> }
  ): Promise<Prompt>;

  /**
   * Get a prompt by name (optionally by version/tag)
   */
  getPrompt(name: string, options?: { version?: string; tag?: string }): Promise<Prompt | null>;

  /**
   * Render a prompt with variables
   */
  renderPrompt(prompt: Prompt, variables: Record<string, string>): string;
}

/**
 * Full provider interface combining all capabilities
 */
export interface Provider extends Tracer, Evaluator {
  /**
   * Provider name
   */
  readonly name: string;

  /**
   * Close the provider and release resources
   */
  close(): Promise<void>;
}

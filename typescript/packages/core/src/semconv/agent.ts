/**
 * OpenTelemetry Semantic Conventions for Agentic AI
 *
 * These attribute names extend OpenTelemetry's gen_ai.* namespace with
 * agent-specific concepts for multi-agent system observability.
 *
 * Mirrors: omniobserve/semconv/agent/attributes.go
 */

// =============================================================================
// Agent Attributes (gen_ai.agent.*)
// =============================================================================

/** Unique identifier of the agent instance */
export const AgentID = 'gen_ai.agent.id';

/** Human-readable name of the agent */
export const AgentName = 'gen_ai.agent.name';

/** Categorizes the agent's role or function */
export const AgentType = 'gen_ai.agent.type';

/** Version of the agent implementation */
export const AgentVersion = 'gen_ai.agent.version';

// =============================================================================
// Workflow Attributes (gen_ai.agent.workflow.*)
// =============================================================================

/** Unique identifier of the workflow/session */
export const WorkflowID = 'gen_ai.agent.workflow.id';

/** Name or type of the workflow */
export const WorkflowName = 'gen_ai.agent.workflow.name';

/** Current status of the workflow */
export const WorkflowStatus = 'gen_ai.agent.workflow.status';

/** Parent workflow ID for nested workflows */
export const WorkflowParentID = 'gen_ai.agent.workflow.parent_id';

/** What initiated the workflow */
export const WorkflowInitiator = 'gen_ai.agent.workflow.initiator';

/** Total number of tasks in the workflow */
export const WorkflowTaskCount = 'gen_ai.agent.workflow.task.count';

/** Number of completed tasks */
export const WorkflowTaskCompletedCount = 'gen_ai.agent.workflow.task.completed_count';

/** Number of failed tasks */
export const WorkflowTaskFailedCount = 'gen_ai.agent.workflow.task.failed_count';

/** Total duration in milliseconds */
export const WorkflowDuration = 'gen_ai.agent.workflow.duration';

// =============================================================================
// Task Attributes (gen_ai.agent.task.*)
// =============================================================================

/** Unique identifier of the task */
export const TaskID = 'gen_ai.agent.task.id';

/** Human-readable name of the task */
export const TaskName = 'gen_ai.agent.task.name';

/** Categorizes the type of task */
export const TaskType = 'gen_ai.agent.task.type';

/** Current status of the task */
export const TaskStatus = 'gen_ai.agent.task.status';

/** Parent task ID for nested tasks */
export const TaskParentID = 'gen_ai.agent.task.parent_id';

/** Number of retry attempts */
export const TaskRetryCount = 'gen_ai.agent.task.retry_count';

/** Task duration in milliseconds */
export const TaskDuration = 'gen_ai.agent.task.duration';

/** Error type if the task failed */
export const TaskErrorType = 'gen_ai.agent.task.error.type';

/** Error message if the task failed */
export const TaskErrorMessage = 'gen_ai.agent.task.error.message';

/** Number of LLM calls made during the task */
export const TaskLLMCallCount = 'gen_ai.agent.task.llm.call_count';

/** Number of tool calls made during the task */
export const TaskToolCallCount = 'gen_ai.agent.task.tool_call.count';

// =============================================================================
// Handoff Attributes (gen_ai.agent.handoff.*)
// =============================================================================

/** Unique identifier of the handoff */
export const HandoffID = 'gen_ai.agent.handoff.id';

/** Type of handoff */
export const HandoffType = 'gen_ai.agent.handoff.type';

/** Current status of the handoff */
export const HandoffStatus = 'gen_ai.agent.handoff.status';

/** Agent initiating the handoff */
export const HandoffFromAgentID = 'gen_ai.agent.handoff.from.agent.id';

/** Type of the source agent */
export const HandoffFromAgentType = 'gen_ai.agent.handoff.from.agent.type';

/** Agent receiving the handoff */
export const HandoffToAgentID = 'gen_ai.agent.handoff.to.agent.id';

/** Type of the target agent */
export const HandoffToAgentType = 'gen_ai.agent.handoff.to.agent.type';

/** Task ID in the source agent */
export const HandoffFromTaskID = 'gen_ai.agent.handoff.from.task.id';

/** Task ID in the target agent */
export const HandoffToTaskID = 'gen_ai.agent.handoff.to.task.id';

/** Size of the handoff payload in bytes */
export const HandoffPayloadSize = 'gen_ai.agent.handoff.payload.size';

/** Time from initiation to acceptance in ms */
export const HandoffLatency = 'gen_ai.agent.handoff.latency';

/** Error message if handoff failed */
export const HandoffErrorMessage = 'gen_ai.agent.handoff.error.message';

// =============================================================================
// Tool Call Attributes (gen_ai.agent.tool_call.*)
// =============================================================================

/** Unique identifier of the tool invocation */
export const ToolCallID = 'gen_ai.agent.tool_call.id';

/** Name of the tool being invoked */
export const ToolCallName = 'gen_ai.agent.tool_call.name';

/** Type of tool */
export const ToolCallType = 'gen_ai.agent.tool_call.type';

/** Status of the tool invocation */
export const ToolCallStatus = 'gen_ai.agent.tool_call.status';

/** Duration of the tool invocation in ms */
export const ToolCallDuration = 'gen_ai.agent.tool_call.duration';

/** Size of the request payload in bytes */
export const ToolCallRequestSize = 'gen_ai.agent.tool_call.request.size';

/** Size of the response payload in bytes */
export const ToolCallResponseSize = 'gen_ai.agent.tool_call.response.size';

/** Number of retry attempts */
export const ToolCallRetryCount = 'gen_ai.agent.tool_call.retry_count';

/** Error type if the tool call failed */
export const ToolCallErrorType = 'gen_ai.agent.tool_call.error.type';

/** Error message if the tool call failed */
export const ToolCallErrorMessage = 'gen_ai.agent.tool_call.error.message';

/** HTTP method used (for HTTP-based tools) */
export const ToolCallHTTPMethod = 'gen_ai.agent.tool_call.http.method';

/** URL called (for HTTP-based tools) */
export const ToolCallHTTPURL = 'gen_ai.agent.tool_call.http.url';

/** HTTP status code returned */
export const ToolCallHTTPStatusCode = 'gen_ai.agent.tool_call.http.status_code';

// =============================================================================
// Event Attributes (gen_ai.agent.event.*)
// =============================================================================

/** Unique identifier of the event */
export const EventID = 'gen_ai.agent.event.id';

/** Name/type of the event */
export const EventName = 'gen_ai.agent.event.name';

/** Event category */
export const EventCategory = 'gen_ai.agent.event.category';

/** Source of the event */
export const EventSource = 'gen_ai.agent.event.source';

/** Severity level of the event */
export const EventSeverity = 'gen_ai.agent.event.severity';

// =============================================================================
// GenAI Attributes (gen_ai.*) - Reused from OpenTelemetry
// =============================================================================

/** GenAI system identifier */
export const GenAISystem = 'gen_ai.system';

/** Model used for the request */
export const GenAIRequestModel = 'gen_ai.request.model';

/** Number of input/prompt tokens */
export const GenAIUsageInputTokens = 'gen_ai.usage.input_tokens';

/** Number of output/completion tokens */
export const GenAIUsageOutputTokens = 'gen_ai.usage.output_tokens';

/** Total tokens (input + output) */
export const GenAIUsageTotalTokens = 'gen_ai.usage.total_tokens';

/** Cost in USD */
export const GenAIUsageCost = 'gen_ai.usage.cost';

// =============================================================================
// Status Values
// =============================================================================

export const StatusPending = 'pending';
export const StatusRunning = 'running';
export const StatusCompleted = 'completed';
export const StatusFailed = 'failed';
export const StatusCancelled = 'cancelled';
export const StatusAccepted = 'accepted';
export const StatusRejected = 'rejected';

// =============================================================================
// Handoff Type Values
// =============================================================================

export const HandoffTypeRequest = 'request';
export const HandoffTypeResponse = 'response';
export const HandoffTypeBroadcast = 'broadcast';
export const HandoffTypeDelegate = 'delegate';

// =============================================================================
// Error Type Values
// =============================================================================

export const ErrorTypeTimeout = 'timeout';
export const ErrorTypeRateLimit = 'rate_limit';
export const ErrorTypeValidation = 'validation';
export const ErrorTypeInternal = 'internal';
export const ErrorTypeNetwork = 'network';
export const ErrorTypeAuth = 'auth';

// =============================================================================
// Event Category Values
// =============================================================================

export const EventCategoryAgent = 'agent';
export const EventCategoryWorkflow = 'workflow';
export const EventCategoryTool = 'tool';
export const EventCategoryDomain = 'domain';
export const EventCategorySystem = 'system';

// =============================================================================
// Severity Values
// =============================================================================

export const SeverityDebug = 'debug';
export const SeverityInfo = 'info';
export const SeverityWarn = 'warn';
export const SeverityError = 'error';

// =============================================================================
// Common Event Names
// =============================================================================

export const EventNameTaskStarted = 'gen_ai.agent.task.started';
export const EventNameTaskCompleted = 'gen_ai.agent.task.completed';
export const EventNameTaskFailed = 'gen_ai.agent.task.failed';
export const EventNameHandoffInitiated = 'gen_ai.agent.handoff.initiated';
export const EventNameHandoffCompleted = 'gen_ai.agent.handoff.completed';
export const EventNameToolCallInvoked = 'gen_ai.agent.tool_call.invoked';
export const EventNameToolCallCompleted = 'gen_ai.agent.tool_call.completed';
export const EventNameWorkflowStarted = 'gen_ai.agent.workflow.started';
export const EventNameWorkflowCompleted = 'gen_ai.agent.workflow.completed';
export const EventNameRetryAttempted = 'gen_ai.agent.retry.attempted';

package agent

// OpenTelemetry Semantic Conventions for Agentic AI
//
// These attribute names extend OpenTelemetry's gen_ai.* namespace with
// agent-specific concepts for multi-agent system observability.

// =============================================================================
// Agent Attributes (gen_ai.agent.*)
// Aligned with OpenTelemetry GenAI Agent Spans conventions
// =============================================================================

const (
	// AgentID is the unique identifier of the agent instance.
	// Type: string
	// Example: "synthesis-agent-1", "verification-agent-2"
	AgentID = "gen_ai.agent.id"

	// AgentName is the human-readable name of the agent.
	// Type: string
	// Example: "Synthesis Agent", "Research Agent"
	AgentName = "gen_ai.agent.name"

	// AgentType categorizes the agent's role or function.
	// Type: string
	// Example: "synthesis", "verification", "research", "orchestration"
	AgentType = "gen_ai.agent.type"

	// AgentVersion is the version of the agent implementation.
	// Type: string
	// Example: "1.0.0", "2.1.3"
	AgentVersion = "gen_ai.agent.version"
)

// =============================================================================
// Workflow Attributes (gen_ai.agent.workflow.*)
// =============================================================================

const (
	// WorkflowID is the unique identifier of the workflow/session.
	// Type: string
	// Example: "wf-550e8400-e29b-41d4-a716-446655440000"
	WorkflowID = "gen_ai.agent.workflow.id"

	// WorkflowName is the name or type of the workflow.
	// Type: string
	// Example: "statistics-extraction", "document-analysis"
	WorkflowName = "gen_ai.agent.workflow.name"

	// WorkflowStatus is the current status of the workflow.
	// Type: string
	// Enum: "pending", "running", "completed", "failed", "cancelled"
	WorkflowStatus = "gen_ai.agent.workflow.status"

	// WorkflowParentID is the parent workflow ID for nested workflows.
	// Type: string
	WorkflowParentID = "gen_ai.agent.workflow.parent_id"

	// WorkflowInitiator identifies what initiated the workflow.
	// Type: string
	// Example: "user:123", "api_key:abc", "system:scheduler"
	WorkflowInitiator = "gen_ai.agent.workflow.initiator"

	// WorkflowTaskCount is the total number of tasks in the workflow.
	// Type: int
	WorkflowTaskCount = "gen_ai.agent.workflow.task.count"

	// WorkflowTaskCompletedCount is the number of completed tasks.
	// Type: int
	WorkflowTaskCompletedCount = "gen_ai.agent.workflow.task.completed_count"

	// WorkflowTaskFailedCount is the number of failed tasks.
	// Type: int
	WorkflowTaskFailedCount = "gen_ai.agent.workflow.task.failed_count"

	// WorkflowDuration is the total duration in milliseconds.
	// Type: int64
	WorkflowDuration = "gen_ai.agent.workflow.duration"
)

// =============================================================================
// Task Attributes (gen_ai.agent.task.*)
// =============================================================================

const (
	// TaskID is the unique identifier of the task.
	// Type: string
	TaskID = "gen_ai.agent.task.id"

	// TaskName is the human-readable name of the task.
	// Type: string
	// Example: "extract_gdp_statistics", "verify_sources"
	TaskName = "gen_ai.agent.task.name"

	// TaskType categorizes the type of task.
	// Type: string
	// Example: "extraction", "verification", "synthesis", "research"
	TaskType = "gen_ai.agent.task.type"

	// TaskStatus is the current status of the task.
	// Type: string
	// Enum: "pending", "running", "completed", "failed", "cancelled"
	TaskStatus = "gen_ai.agent.task.status"

	// TaskParentID links to the parent task for nested tasks.
	// Type: string
	TaskParentID = "gen_ai.agent.task.parent_id"

	// TaskRetryCount is the number of retry attempts.
	// Type: int
	TaskRetryCount = "gen_ai.agent.task.retry_count"

	// TaskDuration is the task duration in milliseconds.
	// Type: int64
	TaskDuration = "gen_ai.agent.task.duration"

	// TaskErrorType categorizes the error if the task failed.
	// Type: string
	// Example: "timeout", "rate_limit", "validation", "internal"
	TaskErrorType = "gen_ai.agent.task.error.type"

	// TaskErrorMessage is the error message if the task failed.
	// Type: string
	TaskErrorMessage = "gen_ai.agent.task.error.message"

	// TaskLLMCallCount is the number of LLM calls made during the task.
	// Type: int
	TaskLLMCallCount = "gen_ai.agent.task.llm.call_count"

	// TaskToolCallCount is the number of tool calls made during the task.
	// Type: int
	TaskToolCallCount = "gen_ai.agent.task.tool_call.count"
)

// =============================================================================
// Handoff Attributes (gen_ai.agent.handoff.*)
// =============================================================================

const (
	// HandoffID is the unique identifier of the handoff.
	// Type: string
	HandoffID = "gen_ai.agent.handoff.id"

	// HandoffType categorizes the type of handoff.
	// Type: string
	// Enum: "request", "response", "broadcast", "delegate"
	HandoffType = "gen_ai.agent.handoff.type"

	// HandoffStatus is the current status of the handoff.
	// Type: string
	// Enum: "pending", "accepted", "rejected", "completed", "failed"
	HandoffStatus = "gen_ai.agent.handoff.status"

	// HandoffFromAgentID is the agent initiating the handoff.
	// Type: string
	HandoffFromAgentID = "gen_ai.agent.handoff.from.agent.id"

	// HandoffFromAgentType is the type of the source agent.
	// Type: string
	HandoffFromAgentType = "gen_ai.agent.handoff.from.agent.type"

	// HandoffToAgentID is the agent receiving the handoff.
	// Type: string
	HandoffToAgentID = "gen_ai.agent.handoff.to.agent.id"

	// HandoffToAgentType is the type of the target agent.
	// Type: string
	HandoffToAgentType = "gen_ai.agent.handoff.to.agent.type"

	// HandoffFromTaskID is the task ID in the source agent.
	// Type: string
	HandoffFromTaskID = "gen_ai.agent.handoff.from.task.id"

	// HandoffToTaskID is the task ID in the target agent.
	// Type: string
	HandoffToTaskID = "gen_ai.agent.handoff.to.task.id"

	// HandoffPayloadSize is the size of the handoff payload in bytes.
	// Type: int
	HandoffPayloadSize = "gen_ai.agent.handoff.payload.size"

	// HandoffLatency is the time from initiation to acceptance in ms.
	// Type: int64
	HandoffLatency = "gen_ai.agent.handoff.latency"

	// HandoffErrorMessage is the error message if handoff failed.
	// Type: string
	HandoffErrorMessage = "gen_ai.agent.handoff.error.message"
)

// =============================================================================
// Tool Call Attributes (gen_ai.agent.tool_call.*)
// Distinct from OTel's gen_ai.tool.* which describes tool definitions
// =============================================================================

const (
	// ToolCallID is the unique identifier of the tool invocation.
	// Type: string
	ToolCallID = "gen_ai.agent.tool_call.id"

	// ToolCallName is the name of the tool being invoked.
	// Type: string
	// Example: "web_search", "database_query", "api_call"
	ToolCallName = "gen_ai.agent.tool_call.name"

	// ToolCallType categorizes the type of tool.
	// Type: string
	// Example: "search", "database", "api", "file", "compute"
	ToolCallType = "gen_ai.agent.tool_call.type"

	// ToolCallStatus is the status of the tool invocation.
	// Type: string
	// Enum: "running", "completed", "failed"
	ToolCallStatus = "gen_ai.agent.tool_call.status"

	// ToolCallDuration is the duration of the tool invocation in ms.
	// Type: int64
	ToolCallDuration = "gen_ai.agent.tool_call.duration"

	// ToolCallRequestSize is the size of the request payload in bytes.
	// Type: int
	ToolCallRequestSize = "gen_ai.agent.tool_call.request.size"

	// ToolCallResponseSize is the size of the response payload in bytes.
	// Type: int
	ToolCallResponseSize = "gen_ai.agent.tool_call.response.size"

	// ToolCallRetryCount is the number of retry attempts.
	// Type: int
	ToolCallRetryCount = "gen_ai.agent.tool_call.retry_count"

	// ToolCallErrorType categorizes the error if the tool call failed.
	// Type: string
	ToolCallErrorType = "gen_ai.agent.tool_call.error.type"

	// ToolCallErrorMessage is the error message if the tool call failed.
	// Type: string
	ToolCallErrorMessage = "gen_ai.agent.tool_call.error.message"

	// ToolCallHTTPMethod is the HTTP method used (for HTTP-based tools).
	// Type: string
	ToolCallHTTPMethod = "gen_ai.agent.tool_call.http.method"

	// ToolCallHTTPURL is the URL called (for HTTP-based tools).
	// Type: string
	ToolCallHTTPURL = "gen_ai.agent.tool_call.http.url"

	// ToolCallHTTPStatusCode is the HTTP status code returned.
	// Type: int
	ToolCallHTTPStatusCode = "gen_ai.agent.tool_call.http.status_code"
)

// =============================================================================
// Event Attributes (gen_ai.agent.event.*)
// =============================================================================

const (
	// EventID is the unique identifier of the event.
	// Type: string
	EventID = "gen_ai.agent.event.id"

	// EventName is the name/type of the event.
	// Type: string
	// Example: "agent.task.started", "agent.handoff.completed"
	EventName = "gen_ai.agent.event.name"

	// EventCategory categorizes the event.
	// Type: string
	// Example: "agent", "workflow", "tool", "domain", "system"
	EventCategory = "gen_ai.agent.event.category"

	// EventSource identifies the source of the event.
	// Type: string
	// Example: "synthesis-agent", "orchestrator"
	EventSource = "gen_ai.agent.event.source"

	// EventSeverity is the severity level of the event.
	// Type: string
	// Enum: "debug", "info", "warn", "error"
	EventSeverity = "gen_ai.agent.event.severity"
)

// =============================================================================
// GenAI Attributes (gen_ai.*) - Reused from OpenTelemetry
// =============================================================================

const (
	// GenAISystem identifies the GenAI system (reused from OTel).
	// Type: string
	// Example: "openai", "anthropic", "google_ai"
	GenAISystem = "gen_ai.system"

	// GenAIRequestModel is the model used for the request.
	// Type: string
	// Example: "gpt-4", "claude-3-opus", "gemini-pro"
	GenAIRequestModel = "gen_ai.request.model"

	// GenAIUsageInputTokens is the number of input/prompt tokens.
	// Type: int
	GenAIUsageInputTokens = "gen_ai.usage.input_tokens" //nolint:gosec // Not a credential, semantic convention attribute name

	// GenAIUsageOutputTokens is the number of output/completion tokens.
	// Type: int
	GenAIUsageOutputTokens = "gen_ai.usage.output_tokens" //nolint:gosec // Not a credential, semantic convention attribute name

	// GenAIUsageTotalTokens is the total tokens (input + output).
	// Type: int
	GenAIUsageTotalTokens = "gen_ai.usage.total_tokens" //nolint:gosec // Not a credential, semantic convention attribute name

	// GenAIUsageCost is the cost in USD.
	// Type: float64
	GenAIUsageCost = "gen_ai.usage.cost"
)

// =============================================================================
// Status Values
// =============================================================================

// Status values for workflow, task, and handoff status attributes.
const (
	StatusPending   = "pending"
	StatusRunning   = "running"
	StatusCompleted = "completed"
	StatusFailed    = "failed"
	StatusCancelled = "cancelled"
	StatusAccepted  = "accepted"
	StatusRejected  = "rejected"
)

// =============================================================================
// Handoff Type Values
// =============================================================================

const (
	HandoffTypeRequest   = "request"
	HandoffTypeResponse  = "response"
	HandoffTypeBroadcast = "broadcast"
	HandoffTypeDelegate  = "delegate"
)

// =============================================================================
// Error Type Values
// =============================================================================

const (
	ErrorTypeTimeout    = "timeout"
	ErrorTypeRateLimit  = "rate_limit"
	ErrorTypeValidation = "validation"
	ErrorTypeInternal   = "internal"
	ErrorTypeNetwork    = "network"
	ErrorTypeAuth       = "auth"
)

// =============================================================================
// Event Category Values
// =============================================================================

const (
	EventCategoryAgent    = "agent"
	EventCategoryWorkflow = "workflow"
	EventCategoryTool     = "tool"
	EventCategoryDomain   = "domain"
	EventCategorySystem   = "system"
)

// =============================================================================
// Severity Values
// =============================================================================

const (
	SeverityDebug = "debug"
	SeverityInfo  = "info"
	SeverityWarn  = "warn"
	SeverityError = "error"
)

// =============================================================================
// Common Event Names
// =============================================================================

const (
	EventNameTaskStarted       = "gen_ai.agent.task.started"
	EventNameTaskCompleted     = "gen_ai.agent.task.completed"
	EventNameTaskFailed        = "gen_ai.agent.task.failed"
	EventNameHandoffInitiated  = "gen_ai.agent.handoff.initiated"
	EventNameHandoffCompleted  = "gen_ai.agent.handoff.completed"
	EventNameToolCallInvoked   = "gen_ai.agent.tool_call.invoked"
	EventNameToolCallCompleted = "gen_ai.agent.tool_call.completed"
	EventNameWorkflowStarted   = "gen_ai.agent.workflow.started"
	EventNameWorkflowCompleted = "gen_ai.agent.workflow.completed"
	EventNameRetryAttempted    = "gen_ai.agent.retry.attempted"
)

package agentops

import "time"

// =============================================================================
// Client Options
// =============================================================================

// ClientOption configures a store client.
type ClientOption func(*ClientConfig)

// ClientConfig holds client configuration.
type ClientConfig struct {
	DSN          string
	MaxOpenConns int
	MaxIdleConns int
	ConnMaxLife  time.Duration
	Debug        bool
	AutoMigrate  bool
}

// WithDSN sets the database connection string.
func WithDSN(dsn string) ClientOption {
	return func(c *ClientConfig) {
		c.DSN = dsn
	}
}

// WithMaxOpenConns sets the maximum number of open connections.
func WithMaxOpenConns(n int) ClientOption {
	return func(c *ClientConfig) {
		c.MaxOpenConns = n
	}
}

// WithMaxIdleConns sets the maximum number of idle connections.
func WithMaxIdleConns(n int) ClientOption {
	return func(c *ClientConfig) {
		c.MaxIdleConns = n
	}
}

// WithConnMaxLifetime sets the maximum connection lifetime.
func WithConnMaxLifetime(d time.Duration) ClientOption {
	return func(c *ClientConfig) {
		c.ConnMaxLife = d
	}
}

// WithDebug enables debug logging.
func WithDebug() ClientOption {
	return func(c *ClientConfig) {
		c.Debug = true
	}
}

// WithAutoMigrate enables automatic schema migration.
func WithAutoMigrate() ClientOption {
	return func(c *ClientConfig) {
		c.AutoMigrate = true
	}
}

// ApplyClientOptions applies options to a config.
func ApplyClientOptions(opts ...ClientOption) *ClientConfig {
	cfg := &ClientConfig{
		MaxOpenConns: 25,
		MaxIdleConns: 5,
		ConnMaxLife:  5 * time.Minute,
	}
	for _, opt := range opts {
		opt(cfg)
	}
	return cfg
}

// =============================================================================
// Workflow Options
// =============================================================================

// WorkflowOption configures workflow creation.
type WorkflowOption func(*WorkflowConfig)

// WorkflowConfig holds workflow creation configuration.
type WorkflowConfig struct {
	TraceID          string
	ParentWorkflowID string
	Initiator        string
	Input            map[string]any
	Metadata         map[string]any
}

// ApplyWorkflowOptions applies options to a config.
func ApplyWorkflowOptions(opts ...WorkflowOption) *WorkflowConfig {
	cfg := &WorkflowConfig{}
	for _, opt := range opts {
		opt(cfg)
	}
	return cfg
}

// WithWorkflowTraceID sets the trace ID.
func WithWorkflowTraceID(traceID string) WorkflowOption {
	return func(c *WorkflowConfig) {
		c.TraceID = traceID
	}
}

// WithParentWorkflow sets the parent workflow ID.
func WithParentWorkflow(id string) WorkflowOption {
	return func(c *WorkflowConfig) {
		c.ParentWorkflowID = id
	}
}

// WithParentWorkflowID is an alias for WithParentWorkflow.
func WithParentWorkflowID(id string) WorkflowOption {
	return WithParentWorkflow(id)
}

// WithWorkflowInitiator sets who initiated the workflow.
func WithWorkflowInitiator(initiator string) WorkflowOption {
	return func(c *WorkflowConfig) {
		c.Initiator = initiator
	}
}

// WithWorkflowInput sets the workflow input.
func WithWorkflowInput(input map[string]any) WorkflowOption {
	return func(c *WorkflowConfig) {
		c.Input = input
	}
}

// WithWorkflowMetadata sets workflow metadata.
func WithWorkflowMetadata(metadata map[string]any) WorkflowOption {
	return func(c *WorkflowConfig) {
		c.Metadata = metadata
	}
}

// WorkflowUpdateOption configures workflow updates.
type WorkflowUpdateOption func(*WorkflowUpdateConfig)

// WorkflowUpdateConfig holds workflow update configuration.
type WorkflowUpdateConfig struct {
	Output    map[string]any
	Metadata  map[string]any
	AddCost   float64
	AddTokens int
	Duration  int64
}

// ApplyWorkflowUpdateOptions applies options to a config.
func ApplyWorkflowUpdateOptions(opts ...WorkflowUpdateOption) *WorkflowUpdateConfig {
	cfg := &WorkflowUpdateConfig{}
	for _, opt := range opts {
		opt(cfg)
	}
	return cfg
}

// WithWorkflowOutput sets the workflow output.
func WithWorkflowOutput(output map[string]any) WorkflowUpdateOption {
	return func(c *WorkflowUpdateConfig) {
		c.Output = output
	}
}

// WithWorkflowAddCost adds to the total cost.
func WithWorkflowAddCost(cost float64) WorkflowUpdateOption {
	return func(c *WorkflowUpdateConfig) {
		c.AddCost = cost
	}
}

// WithWorkflowAddTokens adds to the total tokens.
func WithWorkflowAddTokens(tokens int) WorkflowUpdateOption {
	return func(c *WorkflowUpdateConfig) {
		c.AddTokens = tokens
	}
}

// WithWorkflowUpdateDuration sets the workflow duration.
func WithWorkflowUpdateDuration(duration int64) WorkflowUpdateOption {
	return func(c *WorkflowUpdateConfig) {
		c.Duration = duration
	}
}

// WorkflowCompleteOption configures workflow completion.
type WorkflowCompleteOption func(*WorkflowCompleteConfig)

// WorkflowCompleteConfig holds workflow completion configuration.
type WorkflowCompleteConfig struct {
	Output   map[string]any
	Metadata map[string]any
}

// ApplyWorkflowCompleteOptions applies options to a config.
func ApplyWorkflowCompleteOptions(opts ...WorkflowCompleteOption) *WorkflowCompleteConfig {
	cfg := &WorkflowCompleteConfig{}
	for _, opt := range opts {
		opt(cfg)
	}
	return cfg
}

// WithWorkflowCompleteOutput sets the output on completion.
func WithWorkflowCompleteOutput(output map[string]any) WorkflowCompleteOption {
	return func(c *WorkflowCompleteConfig) {
		c.Output = output
	}
}

// =============================================================================
// Task Options
// =============================================================================

// TaskOption configures task creation.
type TaskOption func(*TaskConfig)

// TaskConfig holds task creation configuration.
type TaskConfig struct {
	AgentType    string
	TaskType     string
	TraceID      string
	SpanID       string
	ParentSpanID string
	Input        map[string]any
	Metadata     map[string]any
}

// ApplyTaskOptions applies options to a config.
func ApplyTaskOptions(opts ...TaskOption) *TaskConfig {
	cfg := &TaskConfig{}
	for _, opt := range opts {
		opt(cfg)
	}
	return cfg
}

// WithAgentType sets the agent type.
func WithAgentType(agentType string) TaskOption {
	return func(c *TaskConfig) {
		c.AgentType = agentType
	}
}

// WithTaskType sets the task type.
func WithTaskType(taskType string) TaskOption {
	return func(c *TaskConfig) {
		c.TaskType = taskType
	}
}

// WithTaskTraceID sets the trace ID.
func WithTaskTraceID(traceID string) TaskOption {
	return func(c *TaskConfig) {
		c.TraceID = traceID
	}
}

// WithTaskSpanID sets the span ID.
func WithTaskSpanID(spanID string) TaskOption {
	return func(c *TaskConfig) {
		c.SpanID = spanID
	}
}

// WithTaskParentSpanID sets the parent span ID.
func WithTaskParentSpanID(parentSpanID string) TaskOption {
	return func(c *TaskConfig) {
		c.ParentSpanID = parentSpanID
	}
}

// WithTaskInput sets the task input.
func WithTaskInput(input map[string]any) TaskOption {
	return func(c *TaskConfig) {
		c.Input = input
	}
}

// WithTaskMetadata sets task metadata.
func WithTaskMetadata(metadata map[string]any) TaskOption {
	return func(c *TaskConfig) {
		c.Metadata = metadata
	}
}

// TaskUpdateOption configures task updates.
type TaskUpdateOption func(*TaskUpdateConfig)

// TaskUpdateConfig holds task update configuration.
type TaskUpdateConfig struct {
	AddLLMCalls  int
	AddToolCalls int
	AddRetries   int
	AddTokens    TokenUsage
	AddCost      float64
	Metadata     map[string]any
}

// ApplyTaskUpdateOptions applies options to a config.
func ApplyTaskUpdateOptions(opts ...TaskUpdateOption) *TaskUpdateConfig {
	cfg := &TaskUpdateConfig{}
	for _, opt := range opts {
		opt(cfg)
	}
	return cfg
}

// TokenUsage holds token counts.
type TokenUsage struct {
	Prompt     int
	Completion int
}

// WithTaskAddLLMCall increments LLM call count.
func WithTaskAddLLMCall() TaskUpdateOption {
	return func(c *TaskUpdateConfig) {
		c.AddLLMCalls++
	}
}

// WithTaskAddToolCall increments tool call count.
func WithTaskAddToolCall() TaskUpdateOption {
	return func(c *TaskUpdateConfig) {
		c.AddToolCalls++
	}
}

// WithTaskAddRetry increments retry count.
func WithTaskAddRetry() TaskUpdateOption {
	return func(c *TaskUpdateConfig) {
		c.AddRetries++
	}
}

// WithTaskAddTokens adds token usage.
func WithTaskAddTokens(prompt, completion int) TaskUpdateOption {
	return func(c *TaskUpdateConfig) {
		c.AddTokens.Prompt += prompt
		c.AddTokens.Completion += completion
	}
}

// WithTaskAddCost adds to the cost.
func WithTaskAddCost(cost float64) TaskUpdateOption {
	return func(c *TaskUpdateConfig) {
		c.AddCost += cost
	}
}

// TaskCompleteOption configures task completion.
type TaskCompleteOption func(*TaskCompleteConfig)

// TaskCompleteConfig holds task completion configuration.
type TaskCompleteConfig struct {
	Output   map[string]any
	Metadata map[string]any
	Duration int64
}

// ApplyTaskCompleteOptions applies options to a config.
func ApplyTaskCompleteOptions(opts ...TaskCompleteOption) *TaskCompleteConfig {
	cfg := &TaskCompleteConfig{}
	for _, opt := range opts {
		opt(cfg)
	}
	return cfg
}

// WithTaskOutput sets the task output.
func WithTaskOutput(output map[string]any) TaskCompleteOption {
	return func(c *TaskCompleteConfig) {
		c.Output = output
	}
}

// WithTaskCompleteMetadata sets metadata on completion.
func WithTaskCompleteMetadata(metadata map[string]any) TaskCompleteOption {
	return func(c *TaskCompleteConfig) {
		c.Metadata = metadata
	}
}

// WithTaskCompleteDuration sets the task duration on completion.
func WithTaskCompleteDuration(duration int64) TaskCompleteOption {
	return func(c *TaskCompleteConfig) {
		c.Duration = duration
	}
}

// TaskFailOption configures task failure.
type TaskFailOption func(*TaskFailConfig)

// TaskFailConfig holds task failure configuration.
type TaskFailConfig struct {
	ErrorType string
	Duration  int64
}

// ApplyTaskFailOptions applies options to a config.
func ApplyTaskFailOptions(opts ...TaskFailOption) *TaskFailConfig {
	cfg := &TaskFailConfig{}
	for _, opt := range opts {
		opt(cfg)
	}
	return cfg
}

// WithTaskErrorType sets the error type.
func WithTaskErrorType(errorType string) TaskFailOption {
	return func(c *TaskFailConfig) {
		c.ErrorType = errorType
	}
}

// WithTaskFailDuration sets the task duration on failure.
func WithTaskFailDuration(duration int64) TaskFailOption {
	return func(c *TaskFailConfig) {
		c.Duration = duration
	}
}

// =============================================================================
// Handoff Options
// =============================================================================

// HandoffOption configures handoff creation.
type HandoffOption func(*HandoffConfig)

// HandoffConfig holds handoff creation configuration.
type HandoffConfig struct {
	WorkflowID    string
	FromAgentType string
	ToAgentType   string
	HandoffType   string
	TraceID       string
	FromTaskID    string
	ToTaskID      string
	Payload       map[string]any
	Metadata      map[string]any
	PayloadSize   int
}

// ApplyHandoffOptions applies options to a config.
func ApplyHandoffOptions(opts ...HandoffOption) *HandoffConfig {
	cfg := &HandoffConfig{}
	for _, opt := range opts {
		opt(cfg)
	}
	return cfg
}

// WithHandoffWorkflow sets the workflow ID.
func WithHandoffWorkflow(workflowID string) HandoffOption {
	return func(c *HandoffConfig) {
		c.WorkflowID = workflowID
	}
}

// WithHandoffWorkflowID is an alias for WithHandoffWorkflow.
func WithHandoffWorkflowID(workflowID string) HandoffOption {
	return WithHandoffWorkflow(workflowID)
}

// WithFromTaskID sets the source task ID.
func WithFromTaskID(taskID string) HandoffOption {
	return func(c *HandoffConfig) {
		c.FromTaskID = taskID
	}
}

// WithHandoffPayloadSize sets the payload size in bytes.
func WithHandoffPayloadSize(size int) HandoffOption {
	return func(c *HandoffConfig) {
		c.PayloadSize = size
	}
}

// WithFromAgentType sets the source agent type.
func WithFromAgentType(agentType string) HandoffOption {
	return func(c *HandoffConfig) {
		c.FromAgentType = agentType
	}
}

// WithToAgentType sets the target agent type.
func WithToAgentType(agentType string) HandoffOption {
	return func(c *HandoffConfig) {
		c.ToAgentType = agentType
	}
}

// WithHandoffType sets the handoff type.
func WithHandoffType(handoffType string) HandoffOption {
	return func(c *HandoffConfig) {
		c.HandoffType = handoffType
	}
}

// WithHandoffTraceID sets the trace ID.
func WithHandoffTraceID(traceID string) HandoffOption {
	return func(c *HandoffConfig) {
		c.TraceID = traceID
	}
}

// WithHandoffPayload sets the payload.
func WithHandoffPayload(payload map[string]any) HandoffOption {
	return func(c *HandoffConfig) {
		c.Payload = payload
	}
}

// WithHandoffMetadata sets metadata.
func WithHandoffMetadata(metadata map[string]any) HandoffOption {
	return func(c *HandoffConfig) {
		c.Metadata = metadata
	}
}

// HandoffUpdateOption configures handoff updates.
type HandoffUpdateOption func(*HandoffUpdateConfig)

// HandoffUpdateConfig holds handoff update configuration.
type HandoffUpdateConfig struct {
	Status       string
	ToTaskID     string
	ErrorMessage string
	Latency      int64
}

// ApplyHandoffUpdateOptions applies options to a config.
func ApplyHandoffUpdateOptions(opts ...HandoffUpdateOption) *HandoffUpdateConfig {
	cfg := &HandoffUpdateConfig{}
	for _, opt := range opts {
		opt(cfg)
	}
	return cfg
}

// WithHandoffStatus sets the status.
func WithHandoffStatus(status string) HandoffUpdateOption {
	return func(c *HandoffUpdateConfig) {
		c.Status = status
	}
}

// WithHandoffUpdateStatus is an alias for WithHandoffStatus.
func WithHandoffUpdateStatus(status string) HandoffUpdateOption {
	return WithHandoffStatus(status)
}

// WithHandoffUpdateLatency sets the handoff latency.
func WithHandoffUpdateLatency(latency int64) HandoffUpdateOption {
	return func(c *HandoffUpdateConfig) {
		c.Latency = latency
	}
}

// WithHandoffUpdateError sets the error message.
func WithHandoffUpdateError(msg string) HandoffUpdateOption {
	return func(c *HandoffUpdateConfig) {
		c.ErrorMessage = msg
	}
}

// WithHandoffToTaskID sets the target task ID.
func WithHandoffToTaskID(taskID string) HandoffUpdateOption {
	return func(c *HandoffUpdateConfig) {
		c.ToTaskID = taskID
	}
}

// WithHandoffError sets the error message.
func WithHandoffError(msg string) HandoffUpdateOption {
	return func(c *HandoffUpdateConfig) {
		c.ErrorMessage = msg
	}
}

// =============================================================================
// Tool Invocation Options
// =============================================================================

// ToolInvocationOption configures tool invocation creation.
type ToolInvocationOption func(*ToolInvocationConfig)

// ToolInvocationConfig holds tool invocation creation configuration.
type ToolInvocationConfig struct {
	ToolType    string
	TraceID     string
	SpanID      string
	Input       map[string]any
	Metadata    map[string]any
	HTTPMethod  string
	HTTPURL     string
	RequestSize int
}

// ApplyToolInvocationOptions applies options to a config.
func ApplyToolInvocationOptions(opts ...ToolInvocationOption) *ToolInvocationConfig {
	cfg := &ToolInvocationConfig{}
	for _, opt := range opts {
		opt(cfg)
	}
	return cfg
}

// WithToolType sets the tool type.
func WithToolType(toolType string) ToolInvocationOption {
	return func(c *ToolInvocationConfig) {
		c.ToolType = toolType
	}
}

// WithToolTraceID sets the trace ID.
func WithToolTraceID(traceID string) ToolInvocationOption {
	return func(c *ToolInvocationConfig) {
		c.TraceID = traceID
	}
}

// WithToolSpanID sets the span ID.
func WithToolSpanID(spanID string) ToolInvocationOption {
	return func(c *ToolInvocationConfig) {
		c.SpanID = spanID
	}
}

// WithToolInput sets the input.
func WithToolInput(input map[string]any) ToolInvocationOption {
	return func(c *ToolInvocationConfig) {
		c.Input = input
	}
}

// WithToolMetadata sets metadata.
func WithToolMetadata(metadata map[string]any) ToolInvocationOption {
	return func(c *ToolInvocationConfig) {
		c.Metadata = metadata
	}
}

// WithToolHTTP sets HTTP details.
func WithToolHTTP(method, url string) ToolInvocationOption {
	return func(c *ToolInvocationConfig) {
		c.HTTPMethod = method
		c.HTTPURL = url
	}
}

// WithToolHTTPMethod sets the HTTP method.
func WithToolHTTPMethod(method string) ToolInvocationOption {
	return func(c *ToolInvocationConfig) {
		c.HTTPMethod = method
	}
}

// WithToolHTTPURL sets the HTTP URL.
func WithToolHTTPURL(url string) ToolInvocationOption {
	return func(c *ToolInvocationConfig) {
		c.HTTPURL = url
	}
}

// WithToolRequestSize sets the request size in bytes.
func WithToolRequestSize(size int) ToolInvocationOption {
	return func(c *ToolInvocationConfig) {
		c.RequestSize = size
	}
}

// ToolInvocationUpdateOption configures tool invocation updates.
type ToolInvocationUpdateOption func(*ToolInvocationUpdateConfig)

// ToolInvocationUpdateConfig holds tool invocation update configuration.
type ToolInvocationUpdateConfig struct {
	RetryCount   int
	Status       string
	Duration     int64
	ErrorType    string
	ErrorMessage string
}

// ApplyToolInvocationUpdateOptions applies options to a config.
func ApplyToolInvocationUpdateOptions(opts ...ToolInvocationUpdateOption) *ToolInvocationUpdateConfig {
	cfg := &ToolInvocationUpdateConfig{}
	for _, opt := range opts {
		opt(cfg)
	}
	return cfg
}

// WithToolRetry increments retry count.
func WithToolRetry() ToolInvocationUpdateOption {
	return func(c *ToolInvocationUpdateConfig) {
		c.RetryCount++
	}
}

// WithToolUpdateStatus sets the tool invocation status.
func WithToolUpdateStatus(status string) ToolInvocationUpdateOption {
	return func(c *ToolInvocationUpdateConfig) {
		c.Status = status
	}
}

// WithToolUpdateDuration sets the tool invocation duration.
func WithToolUpdateDuration(duration int64) ToolInvocationUpdateOption {
	return func(c *ToolInvocationUpdateConfig) {
		c.Duration = duration
	}
}

// WithToolUpdateError sets the error type and message.
func WithToolUpdateError(errorType, errorMessage string) ToolInvocationUpdateOption {
	return func(c *ToolInvocationUpdateConfig) {
		c.ErrorType = errorType
		c.ErrorMessage = errorMessage
	}
}

// ToolInvocationCompleteOption configures tool invocation completion.
type ToolInvocationCompleteOption func(*ToolInvocationCompleteConfig)

// ToolInvocationCompleteConfig holds tool invocation completion configuration.
type ToolInvocationCompleteConfig struct {
	Output            map[string]any
	HTTPStatusCode    int
	ResponseSizeBytes int
	Duration          int64
}

// ApplyToolInvocationCompleteOptions applies options to a config.
func ApplyToolInvocationCompleteOptions(opts ...ToolInvocationCompleteOption) *ToolInvocationCompleteConfig {
	cfg := &ToolInvocationCompleteConfig{}
	for _, opt := range opts {
		opt(cfg)
	}
	return cfg
}

// WithToolOutput sets the output.
func WithToolOutput(output map[string]any) ToolInvocationCompleteOption {
	return func(c *ToolInvocationCompleteConfig) {
		c.Output = output
	}
}

// WithToolHTTPStatus sets the HTTP status code.
func WithToolHTTPStatus(code int) ToolInvocationCompleteOption {
	return func(c *ToolInvocationCompleteConfig) {
		c.HTTPStatusCode = code
	}
}

// WithToolResponseSize sets the response size.
func WithToolResponseSize(size int) ToolInvocationCompleteOption {
	return func(c *ToolInvocationCompleteConfig) {
		c.ResponseSizeBytes = size
	}
}

// WithToolCompleteResponseSize is an alias for WithToolResponseSize.
func WithToolCompleteResponseSize(size int) ToolInvocationCompleteOption {
	return WithToolResponseSize(size)
}

// WithToolCompleteDuration sets the duration on completion.
func WithToolCompleteDuration(duration int64) ToolInvocationCompleteOption {
	return func(c *ToolInvocationCompleteConfig) {
		c.Duration = duration
	}
}

// =============================================================================
// Event Options
// =============================================================================

// EventOption configures event creation.
type EventOption func(*EventConfig)

// EventConfig holds event creation configuration.
type EventConfig struct {
	Category   string
	WorkflowID string
	TaskID     string
	AgentID    string
	TraceID    string
	SpanID     string
	Severity   string
	Data       map[string]any
	Metadata   map[string]any
	Tags       []string
	Source     string
}

// ApplyEventOptions applies options to a config.
func ApplyEventOptions(opts ...EventOption) *EventConfig {
	cfg := &EventConfig{}
	for _, opt := range opts {
		opt(cfg)
	}
	return cfg
}

// WithEventCategory sets the event category.
func WithEventCategory(category string) EventOption {
	return func(c *EventConfig) {
		c.Category = category
	}
}

// WithEventWorkflow sets the workflow ID.
func WithEventWorkflow(workflowID string) EventOption {
	return func(c *EventConfig) {
		c.WorkflowID = workflowID
	}
}

// WithEventTask sets the task ID.
func WithEventTask(taskID string) EventOption {
	return func(c *EventConfig) {
		c.TaskID = taskID
	}
}

// WithEventAgent sets the agent ID.
func WithEventAgent(agentID string) EventOption {
	return func(c *EventConfig) {
		c.AgentID = agentID
	}
}

// WithEventTraceID sets the trace ID.
func WithEventTraceID(traceID string) EventOption {
	return func(c *EventConfig) {
		c.TraceID = traceID
	}
}

// WithEventSpanID sets the span ID.
func WithEventSpanID(spanID string) EventOption {
	return func(c *EventConfig) {
		c.SpanID = spanID
	}
}

// WithEventSeverity sets the severity.
func WithEventSeverity(severity string) EventOption {
	return func(c *EventConfig) {
		c.Severity = severity
	}
}

// WithEventData sets the event data.
func WithEventData(data map[string]any) EventOption {
	return func(c *EventConfig) {
		c.Data = data
	}
}

// WithEventMetadata sets event metadata.
func WithEventMetadata(metadata map[string]any) EventOption {
	return func(c *EventConfig) {
		c.Metadata = metadata
	}
}

// WithEventTags sets event tags.
func WithEventTags(tags ...string) EventOption {
	return func(c *EventConfig) {
		c.Tags = tags
	}
}

// WithEventSource sets the event source.
func WithEventSource(source string) EventOption {
	return func(c *EventConfig) {
		c.Source = source
	}
}

// =============================================================================
// List Options
// =============================================================================

// ListOption configures list queries.
type ListOption func(*ListConfig)

// ListConfig holds list query configuration.
type ListConfig struct {
	Limit      int
	Offset     int
	WorkflowID string
	TaskID     string
	AgentID    string
	Status     string
	EventType  string
	StartTime  *time.Time
	EndTime    *time.Time
	OrderBy    string
	OrderDesc  bool
}

// ApplyListOptions applies options to a config.
func ApplyListOptions(opts ...ListOption) *ListConfig {
	cfg := &ListConfig{}
	for _, opt := range opts {
		opt(cfg)
	}
	return cfg
}

// WithLimit sets the result limit.
func WithLimit(limit int) ListOption {
	return func(c *ListConfig) {
		c.Limit = limit
	}
}

// WithOffset sets the pagination offset.
func WithOffset(offset int) ListOption {
	return func(c *ListConfig) {
		c.Offset = offset
	}
}

// WithFilterWorkflow filters by workflow ID.
func WithFilterWorkflow(workflowID string) ListOption {
	return func(c *ListConfig) {
		c.WorkflowID = workflowID
	}
}

// WithFilterTask filters by task ID.
func WithFilterTask(taskID string) ListOption {
	return func(c *ListConfig) {
		c.TaskID = taskID
	}
}

// WithFilterAgent filters by agent ID.
func WithFilterAgent(agentID string) ListOption {
	return func(c *ListConfig) {
		c.AgentID = agentID
	}
}

// WithFilterStatus filters by status.
func WithFilterStatus(status string) ListOption {
	return func(c *ListConfig) {
		c.Status = status
	}
}

// WithFilterEventType filters by event type.
func WithFilterEventType(eventType string) ListOption {
	return func(c *ListConfig) {
		c.EventType = eventType
	}
}

// WithFilterTimeRange filters by time range.
func WithFilterTimeRange(start, end time.Time) ListOption {
	return func(c *ListConfig) {
		c.StartTime = &start
		c.EndTime = &end
	}
}

// WithOrderBy sets the ordering field.
func WithOrderBy(field string, desc bool) ListOption {
	return func(c *ListConfig) {
		c.OrderBy = field
		c.OrderDesc = desc
	}
}

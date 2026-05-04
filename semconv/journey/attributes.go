package journey

// OpenTelemetry Semantic Conventions for User Journey Observability
//
// These attribute names extend OpenTelemetry's gen_ai.* namespace with
// user journey-specific concepts for product analytics and UX observability.

// =============================================================================
// Journey Attributes (gen_ai.journey.*)
// Describes user journey definitions and their steps
// =============================================================================

const (
	// JourneyID uniquely identifies a journey definition.
	// A journey represents a designed path users should follow to achieve a goal.
	// Type: string
	// Example: "checkout-flow", "onboarding-v2", "feature-discovery"
	JourneyID = "gen_ai.journey.id"

	// JourneyName is the human-readable name of the journey.
	// Type: string
	// Example: "Checkout Flow", "New User Onboarding"
	JourneyName = "gen_ai.journey.name"

	// JourneyVersion is the version of the journey definition.
	// Type: string
	// Example: "1.0.0", "2.1"
	JourneyVersion = "gen_ai.journey.version"

	// JourneyType categorizes the journey.
	// Type: string
	// Example: "conversion", "onboarding", "engagement", "retention"
	JourneyType = "gen_ai.journey.type"

	// JourneyDescription provides a detailed description of the journey goal.
	// Type: string
	JourneyDescription = "gen_ai.journey.description"
)

// =============================================================================
// Journey Step Attributes (gen_ai.journey.step.*)
// Describes individual steps within a journey
// =============================================================================

const (
	// StepID uniquely identifies a step within a journey.
	// Type: string
	// Example: "add-to-cart", "enter-shipping", "confirm-payment"
	StepID = "gen_ai.journey.step.id"

	// StepName is the human-readable name of the step.
	// Type: string
	// Example: "Add to Cart", "Enter Shipping Address"
	StepName = "gen_ai.journey.step.name"

	// StepType categorizes the type of step.
	// Type: string
	// Enum: "page", "action", "checkpoint", "milestone", "entry", "exit"
	StepType = "gen_ai.journey.step.type"

	// StepOrder is the expected sequence number within the journey.
	// Type: int
	// Example: 1, 2, 3
	StepOrder = "gen_ai.journey.step.order"

	// StepIsRequired indicates if this step must be completed.
	// Type: bool
	StepIsRequired = "gen_ai.journey.step.is_required"

	// StepIsOptional indicates if this step can be skipped.
	// Type: bool
	StepIsOptional = "gen_ai.journey.step.is_optional"

	// StepDuration is the time spent at this step in milliseconds.
	// Type: int64
	StepDuration = "gen_ai.journey.step.duration_ms"

	// StepEnteredAt is when the user entered this step.
	// Type: timestamp
	StepEnteredAt = "gen_ai.journey.step.entered_at"

	// StepExitedAt is when the user left this step.
	// Type: timestamp
	StepExitedAt = "gen_ai.journey.step.exited_at"
)

// =============================================================================
// Journey Conversion Attributes (gen_ai.journey.conversion.*)
// Tracks conversion status and outcomes
// =============================================================================

const (
	// ConversionStatus indicates the user's conversion status in the journey.
	// Type: string
	// Enum: "pending", "converted", "dropped", "bounced", "abandoned"
	ConversionStatus = "gen_ai.journey.conversion.status"

	// ConversionStepID is the step where conversion occurred.
	// Type: string
	ConversionStepID = "gen_ai.journey.conversion.step_id"

	// ConversionTimestamp is when conversion occurred.
	// Type: timestamp
	ConversionTimestamp = "gen_ai.journey.conversion.timestamp"

	// DropoffStepID is the last step completed before user dropped off.
	// Type: string
	DropoffStepID = "gen_ai.journey.dropoff.step_id"

	// DropoffReason categorizes why the user dropped off.
	// Type: string
	// Example: "timeout", "error", "navigation_away", "rage_quit", "unknown"
	DropoffReason = "gen_ai.journey.dropoff.reason"
)

// =============================================================================
// Session Attributes (session.*)
// Describes user sessions (collections of events within a time window)
// =============================================================================

const (
	// SessionID uniquely identifies a user session.
	// A session groups related events within a configurable timeout window.
	// Type: string
	// Example: "sess-550e8400-e29b-41d4-a716-446655440000"
	SessionID = "session.id"

	// SessionDuration is the total session duration in milliseconds.
	// Type: int64
	SessionDuration = "session.duration_ms"

	// SessionPageCount is the number of unique pages viewed in the session.
	// Type: int
	SessionPageCount = "session.page_count"

	// SessionEventCount is the total number of events in the session.
	// Type: int
	SessionEventCount = "session.event_count"

	// SessionEntryPage is the first page visited in the session.
	// Type: string
	// Example: "/landing", "/products/widget-123"
	SessionEntryPage = "session.entry_page"

	// SessionExitPage is the last page visited before session end.
	// Type: string
	SessionExitPage = "session.exit_page"

	// SessionReferrer is the external referrer URL that started the session.
	// Type: string
	// Example: "https://google.com/search?q=widgets"
	SessionReferrer = "session.referrer"

	// SessionStartedAt is when the session began.
	// Type: timestamp
	SessionStartedAt = "session.started_at"

	// SessionEndedAt is when the session ended.
	// Type: timestamp
	SessionEndedAt = "session.ended_at"

	// SessionIsActive indicates if the session is still active.
	// Type: bool
	SessionIsActive = "session.is_active"

	// SessionTimeout is the inactivity timeout in milliseconds.
	// Type: int64
	// Default: 1800000 (30 minutes)
	SessionTimeout = "session.timeout_ms"
)

// =============================================================================
// Page Attributes (page.*)
// Describes page navigation and context
// =============================================================================

const (
	// PagePath is the URL path of the page.
	// Type: string
	// Example: "/products/widget-123", "/checkout/shipping"
	PagePath = "page.path"

	// PageTitle is the document title of the page.
	// Type: string
	// Example: "Widget 123 - Our Store"
	PageTitle = "page.title"

	// PageURL is the full URL of the page.
	// Type: string
	PageURL = "page.url"

	// PageReferrer is the previous page URL (internal navigation).
	// Type: string
	PageReferrer = "page.referrer"

	// PageLoadTime is the page load time in milliseconds.
	// Type: int64
	PageLoadTime = "page.load_time_ms"

	// PageViewDuration is time spent on the page in milliseconds.
	// Type: int64
	PageViewDuration = "page.view_duration_ms"
)

// =============================================================================
// UI Interaction Attributes (ui.*)
// Describes user interface interactions
// =============================================================================

const (
	// UIComponentName is the name of the UI component.
	// Type: string
	// Example: "AddToCartButton", "SearchInput", "NavMenu"
	UIComponentName = "ui.component.name"

	// UIComponentPath is the full path to the component in the component tree.
	// Type: string
	// Example: "App/Layout/ProductPage/ProductDetails/AddToCartButton"
	UIComponentPath = "ui.component.path"

	// UIComponentType is the type of component.
	// Type: string
	// Example: "button", "input", "link", "form", "modal"
	UIComponentType = "ui.component.type"

	// UIAction is the user action performed.
	// Type: string
	// Enum: "click", "input", "scroll", "hover", "focus", "blur", "submit", "change"
	UIAction = "ui.action"

	// UIElement is the DOM element type or identifier.
	// Type: string
	// Example: "button", "input[type=text]", "#checkout-form"
	UIElement = "ui.element"

	// UIElementText is the visible text of the element (if applicable).
	// Type: string
	// Example: "Add to Cart", "Submit Order"
	UIElementText = "ui.element.text"

	// UIViewport describes the viewport dimensions.
	// Type: string
	// Format: "widthxheight"
	// Example: "1920x1080", "375x812"
	UIViewport = "ui.viewport"

	// UIScrollPosition is the scroll position as a percentage.
	// Type: float64
	// Range: 0.0 - 1.0
	UIScrollPosition = "ui.scroll.position"

	// UIScrollDirection is the scroll direction.
	// Type: string
	// Enum: "up", "down", "left", "right"
	UIScrollDirection = "ui.scroll.direction"
)

// =============================================================================
// UI State Attributes (ui.state.*)
// Tracks application state changes
// =============================================================================

const (
	// UIStateKey is the name of the state variable that changed.
	// Type: string
	// Example: "cart", "user.preferences", "filters.category"
	UIStateKey = "ui.state.key"

	// UIStateBefore is the JSON-serialized state before the change.
	// Type: string (JSON)
	UIStateBefore = "ui.state.before"

	// UIStateAfter is the JSON-serialized state after the change.
	// Type: string (JSON)
	UIStateAfter = "ui.state.after"

	// UIStateChangeType categorizes the type of state change.
	// Type: string
	// Enum: "set", "update", "delete", "reset", "merge"
	UIStateChangeType = "ui.state.change_type"

	// UIStateSource indicates what triggered the state change.
	// Type: string
	// Example: "user_action", "api_response", "websocket", "timer"
	UIStateSource = "ui.state.source"
)

// =============================================================================
// Snapshot Attributes (snapshot.*)
// Describes visual captures and recordings
// =============================================================================

const (
	// SnapshotID uniquely identifies the snapshot.
	// Type: string
	SnapshotID = "snapshot.id"

	// SnapshotURL is the URL where the snapshot is stored.
	// Type: string
	// Example: "https://storage.example.com/snapshots/abc123.png"
	SnapshotURL = "snapshot.url"

	// SnapshotType is the type of snapshot.
	// Type: string
	// Enum: "screenshot", "dom", "state", "video_frame"
	SnapshotType = "snapshot.type"

	// SnapshotViewport describes the captured viewport dimensions.
	// Type: string
	// Format: "widthxheight"
	SnapshotViewport = "snapshot.viewport"

	// SnapshotTimestamp is when the snapshot was captured.
	// Type: timestamp
	SnapshotTimestamp = "snapshot.timestamp"

	// SnapshotFileSize is the file size in bytes.
	// Type: int
	SnapshotFileSize = "snapshot.file_size"

	// SnapshotFormat is the image format.
	// Type: string
	// Enum: "png", "jpeg", "webp"
	SnapshotFormat = "snapshot.format"

	// SnapshotQuality is the compression quality (for lossy formats).
	// Type: float64
	// Range: 0.0 - 1.0
	SnapshotQuality = "snapshot.quality"
)

// =============================================================================
// Event Type Attributes (event.*)
// Standard event classification
// =============================================================================

const (
	// EventType is the primary event type.
	// Type: string
	// Enum: See EventType* constants below
	EventType = "event.type"

	// EventName is the specific event name within the type.
	// Type: string
	EventName = "event.name"

	// EventTimestamp is when the event occurred.
	// Type: timestamp
	EventTimestamp = "event.timestamp"

	// EventSequence is the event sequence number within the session.
	// Type: int64
	EventSequence = "event.sequence"
)

// =============================================================================
// User Attributes (user.*)
// User identification (privacy-conscious)
// =============================================================================

const (
	// UserID is the user's unique identifier.
	// Type: string
	// Note: May be hashed or pseudonymous for privacy
	UserID = "user.id"

	// UserAnonymousID is the anonymous user identifier (pre-login).
	// Type: string
	UserAnonymousID = "user.anonymous_id"

	// UserType categorizes the user.
	// Type: string
	// Example: "free", "paid", "enterprise", "admin"
	UserType = "user.type"

	// UserFirstSeen is when the user was first observed.
	// Type: timestamp
	UserFirstSeen = "user.first_seen"
)

// =============================================================================
// Project Attributes (project.*)
// Multi-tenant project context
// =============================================================================

const (
	// ProjectID identifies the project/application.
	// Type: string
	ProjectID = "project.id"

	// ProjectName is the project's display name.
	// Type: string
	ProjectName = "project.name"

	// ProjectEnvironment is the deployment environment.
	// Type: string
	// Enum: "development", "staging", "production"
	ProjectEnvironment = "project.environment"
)

// =============================================================================
// API Request Attributes (api.*)
// Tracks frontend-to-backend API calls
// =============================================================================

const (
	// APIMethod is the HTTP method.
	// Type: string
	// Example: "GET", "POST", "PUT", "DELETE"
	APIMethod = "api.method"

	// APIPath is the API endpoint path.
	// Type: string
	// Example: "/api/v1/cart", "/api/users/me"
	APIPath = "api.path"

	// APIStatusCode is the HTTP response status code.
	// Type: int
	APIStatusCode = "api.status_code"

	// APIDuration is the request duration in milliseconds.
	// Type: int64
	APIDuration = "api.duration_ms"

	// APIRequestSize is the request body size in bytes.
	// Type: int
	APIRequestSize = "api.request_size"

	// APIResponseSize is the response body size in bytes.
	// Type: int
	APIResponseSize = "api.response_size"
)

// =============================================================================
// Error Attributes (error.*)
// Error tracking and classification
// =============================================================================

const (
	// ErrorType is the error type or class name.
	// Type: string
	// Example: "TypeError", "NetworkError", "ValidationError"
	ErrorType = "error.type"

	// ErrorMessage is the error message.
	// Type: string
	ErrorMessage = "error.message"

	// ErrorStack is the error stack trace.
	// Type: string
	ErrorStack = "error.stack"

	// ErrorComponent is the component where the error occurred.
	// Type: string
	ErrorComponent = "error.component"

	// ErrorRecoverable indicates if the error was recoverable.
	// Type: bool
	ErrorRecoverable = "error.recoverable"
)

// =============================================================================
// Performance Attributes (performance.*)
// Web Vitals and performance metrics
// =============================================================================

const (
	// PerformanceLCP is Largest Contentful Paint in milliseconds.
	// Type: float64
	PerformanceLCP = "performance.lcp_ms"

	// PerformanceFID is First Input Delay in milliseconds.
	// Type: float64
	PerformanceFID = "performance.fid_ms"

	// PerformanceCLS is Cumulative Layout Shift score.
	// Type: float64
	PerformanceCLS = "performance.cls"

	// PerformanceTTFB is Time to First Byte in milliseconds.
	// Type: float64
	PerformanceTTFB = "performance.ttfb_ms"

	// PerformanceFCP is First Contentful Paint in milliseconds.
	// Type: float64
	PerformanceFCP = "performance.fcp_ms"

	// PerformanceINP is Interaction to Next Paint in milliseconds.
	// Type: float64
	PerformanceINP = "performance.inp_ms"
)

// =============================================================================
// Event Type Values
// Standard event type enum values
// =============================================================================

const (
	// EventTypePageView indicates a page navigation event.
	EventTypePageView = "page.view"

	// EventTypePageLeave indicates leaving a page.
	EventTypePageLeave = "page.leave"

	// EventTypeUIClick indicates a click interaction.
	EventTypeUIClick = "ui.click"

	// EventTypeUIInput indicates text input.
	EventTypeUIInput = "ui.input"

	// EventTypeUIScroll indicates scrolling.
	EventTypeUIScroll = "ui.scroll"

	// EventTypeUIHover indicates hover/mouse movement.
	EventTypeUIHover = "ui.hover"

	// EventTypeUIFocus indicates focus gained.
	EventTypeUIFocus = "ui.focus"

	// EventTypeUIBlur indicates focus lost.
	EventTypeUIBlur = "ui.blur"

	// EventTypeUISubmit indicates form submission.
	EventTypeUISubmit = "ui.submit"

	// EventTypeStateChange indicates application state changed.
	EventTypeStateChange = "state.change"

	// EventTypeAPIRequest indicates an outbound API request started.
	EventTypeAPIRequest = "api.request"

	// EventTypeAPIResponse indicates an API response received.
	EventTypeAPIResponse = "api.response"

	// EventTypeError indicates an error occurred.
	EventTypeError = "error"

	// EventTypePerformance indicates a performance metric.
	EventTypePerformance = "performance"

	// EventTypeJourneyStep indicates journey step progression.
	EventTypeJourneyStep = "journey.step"

	// EventTypeJourneyConversion indicates journey conversion.
	EventTypeJourneyConversion = "journey.conversion"

	// EventTypeSnapshotCaptured indicates a snapshot was taken.
	EventTypeSnapshotCaptured = "snapshot.captured"

	// EventTypeCustom indicates a custom event.
	EventTypeCustom = "custom"
)

// =============================================================================
// Step Type Values
// Journey step type enum values
// =============================================================================

const (
	// StepTypePage indicates a page visit step.
	StepTypePage = "page"

	// StepTypeAction indicates a user action step.
	StepTypeAction = "action"

	// StepTypeCheckpoint indicates a required checkpoint.
	StepTypeCheckpoint = "checkpoint"

	// StepTypeMilestone indicates an important milestone.
	StepTypeMilestone = "milestone"

	// StepTypeEntry indicates the journey entry point.
	StepTypeEntry = "entry"

	// StepTypeExit indicates the journey exit/completion point.
	StepTypeExit = "exit"
)

// =============================================================================
// Conversion Status Values
// Journey conversion status enum values
// =============================================================================

const (
	// ConversionStatusPending indicates user is still in journey.
	ConversionStatusPending = "pending"

	// ConversionStatusConverted indicates user completed the journey.
	ConversionStatusConverted = "converted"

	// ConversionStatusDropped indicates user left mid-journey.
	ConversionStatusDropped = "dropped"

	// ConversionStatusBounced indicates user left on entry.
	ConversionStatusBounced = "bounced"

	// ConversionStatusAbandoned indicates user was inactive and timed out.
	ConversionStatusAbandoned = "abandoned"
)

// =============================================================================
// UI Action Values
// User action enum values
// =============================================================================

const (
	UIActionClick  = "click"
	UIActionInput  = "input"
	UIActionScroll = "scroll"
	UIActionHover  = "hover"
	UIActionFocus  = "focus"
	UIActionBlur   = "blur"
	UIActionSubmit = "submit"
	UIActionChange = "change"
)

// =============================================================================
// Snapshot Type Values
// Snapshot type enum values
// =============================================================================

const (
	SnapshotTypeScreenshot = "screenshot"
	SnapshotTypeDOM        = "dom"
	SnapshotTypeState      = "state"
	SnapshotTypeVideoFrame = "video_frame"
)

// =============================================================================
// Dropoff Reason Values
// Journey dropoff reason enum values
// =============================================================================

const (
	DropoffReasonTimeout        = "timeout"
	DropoffReasonError          = "error"
	DropoffReasonNavigationAway = "navigation_away"
	DropoffReasonRageQuit       = "rage_quit" // Multiple rage clicks detected
	DropoffReasonUnknown        = "unknown"
)

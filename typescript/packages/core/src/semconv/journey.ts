/**
 * OpenTelemetry Semantic Conventions for User Journey Observability
 *
 * These attribute names extend OpenTelemetry's gen_ai.* namespace with
 * user journey-specific concepts for product analytics and UX observability.
 *
 * Mirrors: omniobserve/semconv/journey/attributes.go
 */

// =============================================================================
// Journey Attributes (gen_ai.journey.*)
// =============================================================================

/** Unique identifier of a journey definition */
export const JourneyID = 'gen_ai.journey.id';

/** Human-readable name of the journey */
export const JourneyName = 'gen_ai.journey.name';

/** Version of the journey definition */
export const JourneyVersion = 'gen_ai.journey.version';

/** Categorizes the journey */
export const JourneyType = 'gen_ai.journey.type';

/** Detailed description of the journey goal */
export const JourneyDescription = 'gen_ai.journey.description';

// =============================================================================
// Journey Step Attributes (gen_ai.journey.step.*)
// =============================================================================

/** Unique identifier of a step within a journey */
export const StepID = 'gen_ai.journey.step.id';

/** Human-readable name of the step */
export const StepName = 'gen_ai.journey.step.name';

/** Categorizes the type of step */
export const StepType = 'gen_ai.journey.step.type';

/** Expected sequence number within the journey */
export const StepOrder = 'gen_ai.journey.step.order';

/** Indicates if this step must be completed */
export const StepIsRequired = 'gen_ai.journey.step.is_required';

/** Indicates if this step can be skipped */
export const StepIsOptional = 'gen_ai.journey.step.is_optional';

/** Time spent at this step in milliseconds */
export const StepDuration = 'gen_ai.journey.step.duration_ms';

/** When the user entered this step */
export const StepEnteredAt = 'gen_ai.journey.step.entered_at';

/** When the user left this step */
export const StepExitedAt = 'gen_ai.journey.step.exited_at';

// =============================================================================
// Journey Conversion Attributes (gen_ai.journey.conversion.*)
// =============================================================================

/** User's conversion status in the journey */
export const ConversionStatus = 'gen_ai.journey.conversion.status';

/** Step where conversion occurred */
export const ConversionStepID = 'gen_ai.journey.conversion.step_id';

/** When conversion occurred */
export const ConversionTimestamp = 'gen_ai.journey.conversion.timestamp';

/** Last step completed before user dropped off */
export const DropoffStepID = 'gen_ai.journey.dropoff.step_id';

/** Categorizes why the user dropped off */
export const DropoffReason = 'gen_ai.journey.dropoff.reason';

// =============================================================================
// Session Attributes (session.*)
// =============================================================================

/** Unique identifier of a user session */
export const SessionID = 'session.id';

/** Total session duration in milliseconds */
export const SessionDuration = 'session.duration_ms';

/** Number of unique pages viewed in the session */
export const SessionPageCount = 'session.page_count';

/** Total number of events in the session */
export const SessionEventCount = 'session.event_count';

/** First page visited in the session */
export const SessionEntryPage = 'session.entry_page';

/** Last page visited before session end */
export const SessionExitPage = 'session.exit_page';

/** External referrer URL that started the session */
export const SessionReferrer = 'session.referrer';

/** When the session began */
export const SessionStartedAt = 'session.started_at';

/** When the session ended */
export const SessionEndedAt = 'session.ended_at';

/** Indicates if the session is still active */
export const SessionIsActive = 'session.is_active';

/** Inactivity timeout in milliseconds */
export const SessionTimeout = 'session.timeout_ms';

// =============================================================================
// Page Attributes (page.*)
// =============================================================================

/** URL path of the page */
export const PagePath = 'page.path';

/** Document title of the page */
export const PageTitle = 'page.title';

/** Full URL of the page */
export const PageURL = 'page.url';

/** Previous page URL (internal navigation) */
export const PageReferrer = 'page.referrer';

/** Page load time in milliseconds */
export const PageLoadTime = 'page.load_time_ms';

/** Time spent on the page in milliseconds */
export const PageViewDuration = 'page.view_duration_ms';

// =============================================================================
// UI Interaction Attributes (ui.*)
// =============================================================================

/** Name of the UI component */
export const UIComponentName = 'ui.component.name';

/** Full path to the component in the component tree */
export const UIComponentPath = 'ui.component.path';

/** Type of component */
export const UIComponentType = 'ui.component.type';

/** User action performed */
export const UIAction = 'ui.action';

/** DOM element type or identifier */
export const UIElement = 'ui.element';

/** Visible text of the element */
export const UIElementText = 'ui.element.text';

/** Viewport dimensions */
export const UIViewport = 'ui.viewport';

/** Scroll position as a percentage (0-1) */
export const UIScrollPosition = 'ui.scroll.position';

/** Scroll direction */
export const UIScrollDirection = 'ui.scroll.direction';

// =============================================================================
// UI State Attributes (ui.state.*)
// =============================================================================

/** Name of the state variable that changed */
export const UIStateKey = 'ui.state.key';

/** JSON-serialized state before the change */
export const UIStateBefore = 'ui.state.before';

/** JSON-serialized state after the change */
export const UIStateAfter = 'ui.state.after';

/** Type of state change */
export const UIStateChangeType = 'ui.state.change_type';

/** What triggered the state change */
export const UIStateSource = 'ui.state.source';

// =============================================================================
// Snapshot Attributes (snapshot.*)
// =============================================================================

/** Unique identifier of the snapshot */
export const SnapshotID = 'snapshot.id';

/** URL where the snapshot is stored */
export const SnapshotURL = 'snapshot.url';

/** Type of snapshot */
export const SnapshotType = 'snapshot.type';

/** Captured viewport dimensions */
export const SnapshotViewport = 'snapshot.viewport';

/** When the snapshot was captured */
export const SnapshotTimestamp = 'snapshot.timestamp';

/** File size in bytes */
export const SnapshotFileSize = 'snapshot.file_size';

/** Image format */
export const SnapshotFormat = 'snapshot.format';

/** Compression quality (0-1) */
export const SnapshotQuality = 'snapshot.quality';

// =============================================================================
// Event Attributes (event.*)
// =============================================================================

/** Primary event type */
export const EventType = 'event.type';

/** Specific event name within the type */
export const EventName = 'event.name';

/** When the event occurred */
export const EventTimestamp = 'event.timestamp';

/** Event sequence number within the session */
export const EventSequence = 'event.sequence';

// =============================================================================
// User Attributes (user.*)
// =============================================================================

/** User's unique identifier */
export const UserID = 'user.id';

/** Anonymous user identifier (pre-login) */
export const UserAnonymousID = 'user.anonymous_id';

/** User type/category */
export const UserType = 'user.type';

/** When the user was first observed */
export const UserFirstSeen = 'user.first_seen';

// =============================================================================
// Project Attributes (project.*)
// =============================================================================

/** Project/application identifier */
export const ProjectID = 'project.id';

/** Project's display name */
export const ProjectName = 'project.name';

/** Deployment environment */
export const ProjectEnvironment = 'project.environment';

// =============================================================================
// API Request Attributes (api.*)
// =============================================================================

/** HTTP method */
export const APIMethod = 'api.method';

/** API endpoint path */
export const APIPath = 'api.path';

/** HTTP response status code */
export const APIStatusCode = 'api.status_code';

/** Request duration in milliseconds */
export const APIDuration = 'api.duration_ms';

/** Request body size in bytes */
export const APIRequestSize = 'api.request_size';

/** Response body size in bytes */
export const APIResponseSize = 'api.response_size';

// =============================================================================
// Error Attributes (error.*)
// =============================================================================

/** Error type or class name */
export const ErrorType = 'error.type';

/** Error message */
export const ErrorMessage = 'error.message';

/** Error stack trace */
export const ErrorStack = 'error.stack';

/** Component where the error occurred */
export const ErrorComponent = 'error.component';

/** Indicates if the error was recoverable */
export const ErrorRecoverable = 'error.recoverable';

// =============================================================================
// Performance Attributes (performance.*)
// =============================================================================

/** Largest Contentful Paint in milliseconds */
export const PerformanceLCP = 'performance.lcp_ms';

/** First Input Delay in milliseconds */
export const PerformanceFID = 'performance.fid_ms';

/** Cumulative Layout Shift score */
export const PerformanceCLS = 'performance.cls';

/** Time to First Byte in milliseconds */
export const PerformanceTTFB = 'performance.ttfb_ms';

/** First Contentful Paint in milliseconds */
export const PerformanceFCP = 'performance.fcp_ms';

/** Interaction to Next Paint in milliseconds */
export const PerformanceINP = 'performance.inp_ms';

// =============================================================================
// Event Type Values
// =============================================================================

export const EventTypePageView = 'page.view';
export const EventTypePageLeave = 'page.leave';
export const EventTypeUIClick = 'ui.click';
export const EventTypeUIInput = 'ui.input';
export const EventTypeUIScroll = 'ui.scroll';
export const EventTypeUIHover = 'ui.hover';
export const EventTypeUIFocus = 'ui.focus';
export const EventTypeUIBlur = 'ui.blur';
export const EventTypeUISubmit = 'ui.submit';
export const EventTypeStateChange = 'state.change';
export const EventTypeAPIRequest = 'api.request';
export const EventTypeAPIResponse = 'api.response';
export const EventTypeError = 'error';
export const EventTypePerformance = 'performance';
export const EventTypeJourneyStep = 'journey.step';
export const EventTypeJourneyConversion = 'journey.conversion';
export const EventTypeSnapshotCaptured = 'snapshot.captured';
export const EventTypeCustom = 'custom';

// =============================================================================
// Step Type Values
// =============================================================================

export const StepTypePage = 'page';
export const StepTypeAction = 'action';
export const StepTypeCheckpoint = 'checkpoint';
export const StepTypeMilestone = 'milestone';
export const StepTypeEntry = 'entry';
export const StepTypeExit = 'exit';

// =============================================================================
// Conversion Status Values
// =============================================================================

export const ConversionStatusPending = 'pending';
export const ConversionStatusConverted = 'converted';
export const ConversionStatusDropped = 'dropped';
export const ConversionStatusBounced = 'bounced';
export const ConversionStatusAbandoned = 'abandoned';

// =============================================================================
// UI Action Values
// =============================================================================

export const UIActionClick = 'click';
export const UIActionInput = 'input';
export const UIActionScroll = 'scroll';
export const UIActionHover = 'hover';
export const UIActionFocus = 'focus';
export const UIActionBlur = 'blur';
export const UIActionSubmit = 'submit';
export const UIActionChange = 'change';

// =============================================================================
// Snapshot Type Values
// =============================================================================

export const SnapshotTypeScreenshot = 'screenshot';
export const SnapshotTypeDOM = 'dom';
export const SnapshotTypeState = 'state';
export const SnapshotTypeVideoFrame = 'video_frame';

// =============================================================================
// Dropoff Reason Values
// =============================================================================

export const DropoffReasonTimeout = 'timeout';
export const DropoffReasonError = 'error';
export const DropoffReasonNavigationAway = 'navigation_away';
export const DropoffReasonRageQuit = 'rage_quit';
export const DropoffReasonUnknown = 'unknown';

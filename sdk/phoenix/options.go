package phoenix

import (
	"net/http"
	"time"
)

// Option configures the client.
type Option func(*Client)

// WithEndpoint sets the Phoenix API endpoint.
func WithEndpoint(endpoint string) Option {
	return func(c *Client) {
		c.endpoint = endpoint
	}
}

// WithAPIKey sets the API key for authentication.
func WithAPIKey(key string) Option {
	return func(c *Client) {
		c.apiKey = key
	}
}

// WithHTTPClient sets a custom HTTP client.
func WithHTTPClient(client *http.Client) Option {
	return func(c *Client) {
		c.httpClient = client
	}
}

// WithProjectID sets the default project ID.
func WithProjectID(id string) Option {
	return func(c *Client) {
		c.projectID = id
	}
}

// WithDisabled disables tracing.
func WithDisabled(disabled bool) Option {
	return func(c *Client) {
		c.disabled = disabled
	}
}

// WithDebug enables debug mode.
func WithDebug(debug bool) Option {
	return func(c *Client) {
		c.debug = debug
	}
}

// WithHeaders sets additional HTTP headers.
func WithHeaders(headers map[string]string) Option {
	return func(c *Client) {
		for k, v := range headers {
			c.headers[k] = v
		}
	}
}

// traceConfig holds trace configuration.
type traceConfig struct {
	input     any
	output    any
	metadata  map[string]any
	sessionID string
	userID    string
}

// TraceOption configures trace creation.
type TraceOption func(*traceConfig)

// WithInput sets the trace input.
func WithInput(input any) TraceOption {
	return func(c *traceConfig) {
		c.input = input
	}
}

// WithOutput sets the trace output.
func WithOutput(output any) TraceOption {
	return func(c *traceConfig) {
		c.output = output
	}
}

// WithMetadata sets trace metadata.
func WithMetadata(metadata map[string]any) TraceOption {
	return func(c *traceConfig) {
		c.metadata = metadata
	}
}

// WithSessionID sets the session ID.
func WithSessionID(sessionID string) TraceOption {
	return func(c *traceConfig) {
		c.sessionID = sessionID
	}
}

// WithUserID sets the user ID.
func WithUserID(userID string) TraceOption {
	return func(c *traceConfig) {
		c.userID = userID
	}
}

// spanConfig holds span configuration.
type spanConfig struct {
	spanKind  SpanKind
	input     any
	output    any
	metadata  map[string]any
	model     string
	provider  string
	usage     *Usage
	startTime *time.Time
}

// SpanOption configures span creation.
type SpanOption func(*spanConfig)

// WithSpanKind sets the span kind.
func WithSpanKind(kind SpanKind) SpanOption {
	return func(c *spanConfig) {
		c.spanKind = kind
	}
}

// WithSpanInput sets the span input.
func WithSpanInput(input any) SpanOption {
	return func(c *spanConfig) {
		c.input = input
	}
}

// WithSpanOutput sets the span output.
func WithSpanOutput(output any) SpanOption {
	return func(c *spanConfig) {
		c.output = output
	}
}

// WithSpanMetadata sets span metadata.
func WithSpanMetadata(metadata map[string]any) SpanOption {
	return func(c *spanConfig) {
		c.metadata = metadata
	}
}

// WithModel sets the model name.
func WithModel(model string) SpanOption {
	return func(c *spanConfig) {
		c.model = model
	}
}

// WithProvider sets the provider name.
func WithProvider(provider string) SpanOption {
	return func(c *spanConfig) {
		c.provider = provider
	}
}

// WithUsage sets token usage.
func WithUsage(prompt, completion, total int) SpanOption {
	return func(c *spanConfig) {
		c.usage = &Usage{
			PromptTokens:     prompt,
			CompletionTokens: completion,
			TotalTokens:      total,
		}
	}
}

// WithStartTime sets a custom start time.
func WithStartTime(t time.Time) SpanOption {
	return func(c *spanConfig) {
		c.startTime = &t
	}
}

// ProjectOption configures project creation.
type ProjectOption func(*projectConfig)

type projectConfig struct {
	description string
	metadata    map[string]any
}

// WithProjectDescription sets the project description.
func WithProjectDescription(desc string) ProjectOption {
	return func(c *projectConfig) {
		c.description = desc
	}
}

// DatasetOption configures dataset creation.
type DatasetOption func(*datasetConfig)

type datasetConfig struct {
	description string
	metadata    map[string]any
}

// WithDatasetDescription sets the dataset description.
func WithDatasetDescription(desc string) DatasetOption {
	return func(c *datasetConfig) {
		c.description = desc
	}
}

// WithDatasetMetadata sets dataset metadata.
func WithDatasetMetadata(metadata map[string]any) DatasetOption {
	return func(c *datasetConfig) {
		c.metadata = metadata
	}
}

// AnnotationOption configures annotation creation.
type AnnotationOption func(*annotationConfig)

type annotationConfig struct {
	label       string
	explanation string
	metadata    map[string]any
}

// WithAnnotationLabel sets the annotation label.
func WithAnnotationLabel(label string) AnnotationOption {
	return func(c *annotationConfig) {
		c.label = label
	}
}

// WithAnnotationExplanation sets the annotation explanation.
func WithAnnotationExplanation(explanation string) AnnotationOption {
	return func(c *annotationConfig) {
		c.explanation = explanation
	}
}

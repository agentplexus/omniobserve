package langfuse

import (
	"net/http"
	"time"
)

// Option configures the client.
type Option func(*Client)

// WithPublicKey sets the Langfuse public key.
func WithPublicKey(key string) Option {
	return func(c *Client) {
		c.publicKey = key
	}
}

// WithSecretKey sets the Langfuse secret key.
func WithSecretKey(key string) Option {
	return func(c *Client) {
		c.secretKey = key
	}
}

// WithEndpoint sets the Langfuse API endpoint.
func WithEndpoint(endpoint string) Option {
	return func(c *Client) {
		c.endpoint = endpoint
	}
}

// WithHTTPClient sets a custom HTTP client.
func WithHTTPClient(client *http.Client) Option {
	return func(c *Client) {
		c.httpClient = client
	}
}

// WithTimeout sets the request timeout.
func WithTimeout(timeout time.Duration) Option {
	return func(c *Client) {
		c.timeout = timeout
		c.httpClient.Timeout = timeout
	}
}

// WithBatchSize sets the batch size for event ingestion.
func WithBatchSize(size int) Option {
	return func(c *Client) {
		c.batchSize = size
	}
}

// WithFlushPeriod sets the flush interval for batched events.
func WithFlushPeriod(period time.Duration) Option {
	return func(c *Client) {
		c.flushPeriod = period
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

// traceConfig holds trace configuration.
type traceConfig struct {
	input     any
	output    any
	metadata  map[string]any
	tags      []string
	userId    string
	sessionId string
	public    bool
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

// WithTags sets trace tags.
func WithTags(tags ...string) TraceOption {
	return func(c *traceConfig) {
		c.tags = tags
	}
}

// WithUserID sets the user ID.
func WithUserID(userId string) TraceOption {
	return func(c *traceConfig) {
		c.userId = userId
	}
}

// WithSessionID sets the session ID.
func WithSessionID(sessionId string) TraceOption {
	return func(c *traceConfig) {
		c.sessionId = sessionId
	}
}

// WithPublic sets whether the trace is public.
func WithPublic(public bool) TraceOption {
	return func(c *traceConfig) {
		c.public = public
	}
}

// spanConfig holds span configuration.
type spanConfig struct {
	input    any
	output   any
	metadata map[string]any
	level    string // DEBUG, DEFAULT, WARNING, ERROR
	version  string
}

// SpanOption configures span creation.
type SpanOption func(*spanConfig)

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

// WithLevel sets the log level.
func WithLevel(level string) SpanOption {
	return func(c *spanConfig) {
		c.level = level
	}
}

// WithVersion sets the version.
func WithVersion(version string) SpanOption {
	return func(c *spanConfig) {
		c.version = version
	}
}

// generationConfig holds generation configuration.
type generationConfig struct {
	spanConfig
	model           string
	modelParameters map[string]any
	promptName      string
	promptVersion   int
	usage           *Usage
	completionStart *time.Time
}

// GenerationOption configures generation creation.
type GenerationOption func(*generationConfig)

// WithModel sets the model name.
func WithModel(model string) GenerationOption {
	return func(c *generationConfig) {
		c.model = model
	}
}

// WithModelParameters sets model parameters.
func WithModelParameters(params map[string]any) GenerationOption {
	return func(c *generationConfig) {
		c.modelParameters = params
	}
}

// WithPromptName sets the prompt name.
func WithPromptName(name string) GenerationOption {
	return func(c *generationConfig) {
		c.promptName = name
	}
}

// WithPromptVersion sets the prompt version.
func WithPromptVersion(version int) GenerationOption {
	return func(c *generationConfig) {
		c.promptVersion = version
	}
}

// WithUsage sets token usage.
func WithUsage(promptTokens, completionTokens, totalTokens int) GenerationOption {
	return func(c *generationConfig) {
		c.usage = &Usage{
			PromptTokens:     promptTokens,
			CompletionTokens: completionTokens,
			TotalTokens:      totalTokens,
		}
	}
}

// WithGenerationInput sets generation input.
func WithGenerationInput(input any) GenerationOption {
	return func(c *generationConfig) {
		c.input = input
	}
}

// WithGenerationOutput sets generation output.
func WithGenerationOutput(output any) GenerationOption {
	return func(c *generationConfig) {
		c.output = output
	}
}

// WithGenerationMetadata sets generation metadata.
func WithGenerationMetadata(metadata map[string]any) GenerationOption {
	return func(c *generationConfig) {
		c.metadata = metadata
	}
}

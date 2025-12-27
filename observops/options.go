package observops

import "time"

// ClientOption configures a provider client.
type ClientOption func(*Config)

// WithServiceName sets the service name.
func WithServiceName(name string) ClientOption {
	return func(c *Config) {
		c.ServiceName = name
	}
}

// WithServiceVersion sets the service version.
func WithServiceVersion(version string) ClientOption {
	return func(c *Config) {
		c.ServiceVersion = version
	}
}

// WithEndpoint sets the backend endpoint.
func WithEndpoint(endpoint string) ClientOption {
	return func(c *Config) {
		c.Endpoint = endpoint
	}
}

// WithAPIKey sets the API key for authentication.
func WithAPIKey(apiKey string) ClientOption {
	return func(c *Config) {
		c.APIKey = apiKey
	}
}

// WithInsecure disables TLS for the connection.
func WithInsecure() ClientOption {
	return func(c *Config) {
		c.Insecure = true
	}
}

// WithHeaders sets additional headers to send with requests.
func WithHeaders(headers map[string]string) ClientOption {
	return func(c *Config) {
		c.Headers = headers
	}
}

// WithResource sets the resource describing the service.
func WithResource(resource *Resource) ClientOption {
	return func(c *Config) {
		c.Resource = resource
	}
}

// WithBatchTimeout sets the maximum time to wait before exporting.
func WithBatchTimeout(timeout time.Duration) ClientOption {
	return func(c *Config) {
		c.BatchTimeout = timeout
	}
}

// WithBatchSize sets the maximum number of items per batch.
func WithBatchSize(size int) ClientOption {
	return func(c *Config) {
		c.BatchSize = size
	}
}

// WithDisabled disables telemetry collection.
func WithDisabled() ClientOption {
	return func(c *Config) {
		c.Disabled = true
	}
}

// WithDebug enables debug logging.
func WithDebug() ClientOption {
	return func(c *Config) {
		c.Debug = true
	}
}

// ApplyOptions applies the given options to a config.
func ApplyOptions(opts ...ClientOption) *Config {
	cfg := &Config{
		BatchTimeout: 5 * time.Second,
		BatchSize:    512,
	}
	for _, opt := range opts {
		opt(cfg)
	}
	return cfg
}

// MetricOption configures a metric instrument.
type MetricOption func(*metricConfig)

type metricConfig struct {
	description string
	unit        string
}

// WithDescription sets the metric description.
func WithDescription(desc string) MetricOption {
	return func(c *metricConfig) {
		c.description = desc
	}
}

// WithUnit sets the metric unit.
func WithUnit(unit string) MetricOption {
	return func(c *metricConfig) {
		c.unit = unit
	}
}

// ApplyMetricOptions applies metric options and returns the config.
func ApplyMetricOptions(opts ...MetricOption) *metricConfig {
	cfg := &metricConfig{}
	for _, opt := range opts {
		opt(cfg)
	}
	return cfg
}

// GetDescription returns the description from metric options.
func GetDescription(opts ...MetricOption) string {
	cfg := ApplyMetricOptions(opts...)
	return cfg.description
}

// GetUnit returns the unit from metric options.
func GetUnit(opts ...MetricOption) string {
	cfg := ApplyMetricOptions(opts...)
	return cfg.unit
}

// RecordOption configures a metric recording.
type RecordOption func(*recordConfig)

type recordConfig struct {
	attributes []KeyValue
}

// WithAttributes sets attributes for a metric recording.
func WithAttributes(attrs ...KeyValue) RecordOption {
	return func(c *recordConfig) {
		c.attributes = append(c.attributes, attrs...)
	}
}

// ApplyRecordOptions applies record options and returns the config.
func ApplyRecordOptions(opts ...RecordOption) *recordConfig {
	cfg := &recordConfig{}
	for _, opt := range opts {
		opt(cfg)
	}
	return cfg
}

// GetAttributes returns the attributes from record options.
func GetAttributes(opts ...RecordOption) []KeyValue {
	cfg := ApplyRecordOptions(opts...)
	return cfg.attributes
}

// SpanOption configures span creation.
type SpanOption func(*spanConfig)

type spanConfig struct {
	kind       SpanKind
	attributes []KeyValue
	links      []SpanContext
}

// WithSpanKind sets the span kind.
func WithSpanKind(kind SpanKind) SpanOption {
	return func(c *spanConfig) {
		c.kind = kind
	}
}

// WithSpanAttributes sets initial span attributes.
func WithSpanAttributes(attrs ...KeyValue) SpanOption {
	return func(c *spanConfig) {
		c.attributes = append(c.attributes, attrs...)
	}
}

// WithSpanLinks sets span links.
func WithSpanLinks(links ...SpanContext) SpanOption {
	return func(c *spanConfig) {
		c.links = append(c.links, links...)
	}
}

// ApplySpanOptions applies span options and returns the config.
func ApplySpanOptions(opts ...SpanOption) *spanConfig {
	cfg := &spanConfig{}
	for _, opt := range opts {
		opt(cfg)
	}
	return cfg
}

// GetSpanKind returns the span kind from options.
func GetSpanKind(opts ...SpanOption) SpanKind {
	cfg := ApplySpanOptions(opts...)
	return cfg.kind
}

// GetSpanAttributes returns the attributes from span options.
func GetSpanAttributes(opts ...SpanOption) []KeyValue {
	cfg := ApplySpanOptions(opts...)
	return cfg.attributes
}

// GetSpanLinks returns the links from span options.
func GetSpanLinks(opts ...SpanOption) []SpanContext {
	cfg := ApplySpanOptions(opts...)
	return cfg.links
}

// SpanEndOption configures span end behavior.
type SpanEndOption func(*spanEndConfig)

type spanEndConfig struct {
	timestamp *time.Time
}

// WithEndTimestamp sets a custom end timestamp.
func WithEndTimestamp(t time.Time) SpanEndOption {
	return func(c *spanEndConfig) {
		c.timestamp = &t
	}
}

// ApplySpanEndOptions applies span end options and returns the config.
func ApplySpanEndOptions(opts ...SpanEndOption) *spanEndConfig {
	cfg := &spanEndConfig{}
	for _, opt := range opts {
		opt(cfg)
	}
	return cfg
}

// GetEndTimestamp returns the timestamp from span end options.
func GetEndTimestamp(opts ...SpanEndOption) *time.Time {
	cfg := ApplySpanEndOptions(opts...)
	return cfg.timestamp
}

// EventOption configures a span event.
type EventOption func(*eventConfig)

type eventConfig struct {
	attributes []KeyValue
	timestamp  *time.Time
}

// WithEventAttributes sets event attributes.
func WithEventAttributes(attrs ...KeyValue) EventOption {
	return func(c *eventConfig) {
		c.attributes = append(c.attributes, attrs...)
	}
}

// WithEventTimestamp sets a custom event timestamp.
func WithEventTimestamp(t time.Time) EventOption {
	return func(c *eventConfig) {
		c.timestamp = &t
	}
}

// ApplyEventOptions applies event options and returns the config.
func ApplyEventOptions(opts ...EventOption) *eventConfig {
	cfg := &eventConfig{}
	for _, opt := range opts {
		opt(cfg)
	}
	return cfg
}

// GetEventAttributes returns the attributes from event options.
func GetEventAttributes(opts ...EventOption) []KeyValue {
	cfg := ApplyEventOptions(opts...)
	return cfg.attributes
}

// GetEventTimestamp returns the timestamp from event options.
func GetEventTimestamp(opts ...EventOption) *time.Time {
	cfg := ApplyEventOptions(opts...)
	return cfg.timestamp
}

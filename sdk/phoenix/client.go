// Package phoenix provides a Go SDK for Arize Phoenix, an open-source LLM observability platform.
//
// Phoenix uses OpenTelemetry for trace collection. This SDK wraps OTEL with Phoenix-specific
// semantics and convenience methods.
//
// Usage:
//
//	client, err := phoenix.NewClient(
//		phoenix.WithEndpoint("http://localhost:6006"),
//		phoenix.WithAPIKey("..."), // Optional for cloud
//	)
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer client.Shutdown(context.Background())
//
//	ctx, trace, _ := client.StartTrace(ctx, "my-trace")
//	defer trace.End(ctx)
package phoenix

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/google/uuid"
)

// Version is the SDK version.
const Version = "0.1.0"

// Default endpoints.
const (
	DefaultEndpoint = "http://localhost:6006"
	CloudEndpoint   = "https://app.phoenix.arize.com"
)

// Client is the main Phoenix client.
type Client struct {
	endpoint   string
	apiKey     string
	httpClient *http.Client
	projectID  string
	headers    map[string]string

	// For non-OTEL mode (direct REST API)
	mu     sync.Mutex
	spans  []*Span
	traces []*Trace

	disabled bool
	debug    bool
}

// NewClient creates a new Phoenix client.
func NewClient(opts ...Option) (*Client, error) {
	c := &Client{
		endpoint:   DefaultEndpoint,
		httpClient: &http.Client{Timeout: 30 * time.Second},
		headers:    make(map[string]string),
		spans:      make([]*Span, 0),
		traces:     make([]*Trace, 0),
	}

	for _, opt := range opts {
		opt(c)
	}

	// Add auth header if API key provided
	if c.apiKey != "" {
		c.headers["Authorization"] = "Bearer " + c.apiKey
	}

	return c, nil
}

// Shutdown gracefully shuts down the client and flushes any pending data.
func (c *Client) Shutdown(ctx context.Context) error {
	if c.disabled {
		return nil
	}

	// Flush any pending spans
	return c.flush(ctx)
}

// Close is an alias for Shutdown.
func (c *Client) Close() error {
	return c.Shutdown(context.Background())
}

// flush sends any pending spans to Phoenix.
func (c *Client) flush(_ context.Context) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	// In a full implementation, this would batch send spans via REST or OTLP
	c.spans = make([]*Span, 0)
	c.traces = make([]*Trace, 0)
	return nil
}

// StartTrace starts a new trace.
func (c *Client) StartTrace(ctx context.Context, name string, opts ...TraceOption) (context.Context, *Trace, error) {
	if c.disabled {
		return ctx, &Trace{disabled: true}, nil
	}

	cfg := &traceConfig{}
	for _, opt := range opts {
		opt(cfg)
	}

	trace := &Trace{
		client:    c,
		id:        uuid.New().String(),
		name:      name,
		startTime: time.Now(),
		metadata:  cfg.metadata,
		sessionID: cfg.sessionID,
		userID:    cfg.userID,
		input:     cfg.input,
	}

	c.mu.Lock()
	c.traces = append(c.traces, trace)
	c.mu.Unlock()

	newCtx := ContextWithTrace(ctx, trace)
	newCtx = ContextWithClient(newCtx, c)
	return newCtx, trace, nil
}

// Project represents a Phoenix project.
type Project struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description,omitempty"`
	CreatedAt   time.Time `json:"createdAt"`
}

// GetProject retrieves a project by name.
func (c *Client) GetProject(ctx context.Context, name string) (*Project, error) {
	// This would call the Phoenix REST API
	return nil, fmt.Errorf("not implemented: GetProject")
}

// CreateProject creates a new project.
func (c *Client) CreateProject(ctx context.Context, name string, opts ...ProjectOption) (*Project, error) {
	return nil, fmt.Errorf("not implemented: CreateProject")
}

// ListProjects lists all projects.
func (c *Client) ListProjects(ctx context.Context, limit, offset int) ([]*Project, error) {
	return nil, fmt.Errorf("not implemented: ListProjects")
}

// Dataset represents a Phoenix dataset.
type Dataset struct {
	ID          string         `json:"id"`
	Name        string         `json:"name"`
	Description string         `json:"description,omitempty"`
	Metadata    map[string]any `json:"metadata,omitempty"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
}

// DatasetItem represents an item in a dataset.
type DatasetItem struct {
	ID       string         `json:"id"`
	Input    any            `json:"input,omitempty"`
	Output   any            `json:"output,omitempty"`
	Expected any            `json:"expected,omitempty"`
	Metadata map[string]any `json:"metadata,omitempty"`
}

// CreateDataset creates a new dataset.
func (c *Client) CreateDataset(ctx context.Context, name string, opts ...DatasetOption) (*Dataset, error) {
	return nil, fmt.Errorf("not implemented: CreateDataset")
}

// GetDataset retrieves a dataset by name.
func (c *Client) GetDataset(ctx context.Context, name string) (*Dataset, error) {
	return nil, fmt.Errorf("not implemented: GetDataset")
}

// ListDatasets lists all datasets.
func (c *Client) ListDatasets(ctx context.Context, limit, offset int) ([]*Dataset, error) {
	return nil, fmt.Errorf("not implemented: ListDatasets")
}

// AddDatasetItems adds items to a dataset.
func (c *Client) AddDatasetItems(ctx context.Context, datasetID string, items []DatasetItem) error {
	return fmt.Errorf("not implemented: AddDatasetItems")
}

// Experiment represents an evaluation experiment.
type Experiment struct {
	ID        string         `json:"id"`
	Name      string         `json:"name"`
	DatasetID string         `json:"datasetId"`
	Metadata  map[string]any `json:"metadata,omitempty"`
	CreatedAt time.Time      `json:"createdAt"`
}

// CreateExperiment creates a new experiment.
func (c *Client) CreateExperiment(ctx context.Context, datasetID string, name string) (*Experiment, error) {
	return nil, fmt.Errorf("not implemented: CreateExperiment")
}

// AddAnnotation adds an annotation/feedback to a span.
func (c *Client) AddAnnotation(ctx context.Context, spanID string, name string, score float64, opts ...AnnotationOption) error {
	return fmt.Errorf("not implemented: AddAnnotation")
}

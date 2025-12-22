// Package phoenix provides a Phoenix adapter for the llmops abstraction.
//
// Import this package to register the Phoenix provider:
//
//	import _ "github.com/grokify/metaobserve/llmops/phoenix"
//
// Then open it:
//
//	provider, err := llmops.Open("phoenix", llmops.WithEndpoint("http://localhost:6006"))
package phoenix

import (
	"context"
	"time"

	"github.com/grokify/metaobserve/llmops"
	sdk "github.com/grokify/metaobserve/sdk/phoenix"
)

const ProviderName = "phoenix"

func init() {
	llmops.Register(ProviderName, New)
	llmops.RegisterInfo(llmops.ProviderInfo{
		Name:        ProviderName,
		Description: "Arize Phoenix - Open-source LLM observability platform",
		Website:     "https://phoenix.arize.com",
		OpenSource:  true,
		SelfHosted:  true,
		Capabilities: []llmops.Capability{
			llmops.CapabilityTracing,
			llmops.CapabilityEvaluation,
			llmops.CapabilityDatasets,
			llmops.CapabilityExperiments,
			llmops.CapabilityOTel,
		},
	})
}

// Provider implements llmops.Provider for Phoenix.
type Provider struct {
	client *sdk.Client
}

// New creates a new Phoenix provider.
func New(opts ...llmops.ClientOption) (llmops.Provider, error) {
	cfg := llmops.ApplyClientOptions(opts...)

	sdkOpts := []sdk.Option{}
	if cfg.Endpoint != "" {
		sdkOpts = append(sdkOpts, sdk.WithEndpoint(cfg.Endpoint))
	}
	if cfg.APIKey != "" {
		sdkOpts = append(sdkOpts, sdk.WithAPIKey(cfg.APIKey))
	}
	if cfg.HTTPClient != nil {
		sdkOpts = append(sdkOpts, sdk.WithHTTPClient(cfg.HTTPClient))
	}
	if cfg.ProjectName != "" {
		sdkOpts = append(sdkOpts, sdk.WithProjectID(cfg.ProjectName))
	}
	if cfg.Disabled {
		sdkOpts = append(sdkOpts, sdk.WithDisabled(true))
	}
	if cfg.Debug {
		sdkOpts = append(sdkOpts, sdk.WithDebug(true))
	}

	client, err := sdk.NewClient(sdkOpts...)
	if err != nil {
		return nil, err
	}

	return &Provider{client: client}, nil
}

// Name returns the provider name.
func (p *Provider) Name() string {
	return ProviderName
}

// Close closes the provider.
func (p *Provider) Close() error {
	return p.client.Close()
}

// StartTrace starts a new trace.
func (p *Provider) StartTrace(ctx context.Context, name string, opts ...llmops.TraceOption) (context.Context, llmops.Trace, error) {
	cfg := llmops.ApplyTraceOptions(opts...)

	sdkOpts := []sdk.TraceOption{}
	if cfg.Input != nil {
		sdkOpts = append(sdkOpts, sdk.WithInput(cfg.Input))
	}
	if cfg.Metadata != nil {
		sdkOpts = append(sdkOpts, sdk.WithMetadata(cfg.Metadata))
	}
	if cfg.ThreadID != "" {
		sdkOpts = append(sdkOpts, sdk.WithSessionID(cfg.ThreadID))
	}

	newCtx, trace, err := p.client.StartTrace(ctx, name, sdkOpts...)
	if err != nil {
		return ctx, nil, err
	}

	return newCtx, &traceAdapter{trace: trace}, nil
}

// StartSpan starts a new span.
func (p *Provider) StartSpan(ctx context.Context, name string, opts ...llmops.SpanOption) (context.Context, llmops.Span, error) {
	cfg := llmops.ApplySpanOptions(opts...)

	sdkOpts := mapSpanOptions(cfg)
	newCtx, span, err := sdk.StartSpan(ctx, name, sdkOpts...)
	if err != nil {
		return ctx, nil, err
	}

	return newCtx, &spanAdapter{span: span}, nil
}

// TraceFromContext gets the current trace from context.
func (p *Provider) TraceFromContext(ctx context.Context) (llmops.Trace, bool) {
	trace := sdk.TraceFromContext(ctx)
	if trace == nil {
		return nil, false
	}
	return &traceAdapter{trace: trace}, true
}

// SpanFromContext gets the current span from context.
func (p *Provider) SpanFromContext(ctx context.Context) (llmops.Span, bool) {
	span := sdk.SpanFromContext(ctx)
	if span == nil {
		return nil, false
	}
	return &spanAdapter{span: span}, true
}

// Evaluate runs evaluation metrics.
func (p *Provider) Evaluate(ctx context.Context, input llmops.EvalInput, metrics ...llmops.Metric) (*llmops.EvalResult, error) {
	startTime := time.Now()

	scores := make([]llmops.MetricScore, 0, len(metrics))
	for _, metric := range metrics {
		score, err := metric.Evaluate(input)
		if err != nil {
			scores = append(scores, llmops.MetricScore{
				Name:  metric.Name(),
				Error: err.Error(),
			})
		} else {
			scores = append(scores, score)
		}
	}

	return &llmops.EvalResult{
		Scores:   scores,
		Duration: time.Since(startTime),
	}, nil
}

// AddFeedbackScore adds a feedback score.
func (p *Provider) AddFeedbackScore(ctx context.Context, opts llmops.FeedbackScoreOpts) error {
	if span := sdk.SpanFromContext(ctx); span != nil {
		return span.AddAnnotation(ctx, opts.Name, opts.Score)
	}
	if trace := sdk.TraceFromContext(ctx); trace != nil {
		return trace.AddAnnotation(ctx, opts.Name, opts.Score)
	}
	return llmops.ErrNoActiveTrace
}

// CreatePrompt is not directly supported in Phoenix.
func (p *Provider) CreatePrompt(ctx context.Context, name string, template string, opts ...llmops.PromptOption) (*llmops.Prompt, error) {
	return nil, llmops.WrapNotImplemented(ProviderName, "CreatePrompt")
}

// GetPrompt is not directly supported in Phoenix.
func (p *Provider) GetPrompt(ctx context.Context, name string, version ...string) (*llmops.Prompt, error) {
	return nil, llmops.WrapNotImplemented(ProviderName, "GetPrompt")
}

// ListPrompts is not directly supported in Phoenix.
func (p *Provider) ListPrompts(ctx context.Context, opts ...llmops.ListOption) ([]*llmops.Prompt, error) {
	return nil, llmops.WrapNotImplemented(ProviderName, "ListPrompts")
}

// CreateDataset creates a new dataset.
func (p *Provider) CreateDataset(ctx context.Context, name string, opts ...llmops.DatasetOption) (*llmops.Dataset, error) {
	cfg := &llmops.DatasetOptions{}
	for _, opt := range opts {
		opt(cfg)
	}

	sdkOpts := []sdk.DatasetOption{}
	if cfg.Description != "" {
		sdkOpts = append(sdkOpts, sdk.WithDatasetDescription(cfg.Description))
	}

	dataset, err := p.client.CreateDataset(ctx, name, sdkOpts...)
	if err != nil {
		return nil, err
	}

	return &llmops.Dataset{
		ID:          dataset.ID,
		Name:        dataset.Name,
		Description: dataset.Description,
		CreatedAt:   dataset.CreatedAt,
		UpdatedAt:   dataset.UpdatedAt,
	}, nil
}

// GetDataset gets a dataset by name.
func (p *Provider) GetDataset(ctx context.Context, name string) (*llmops.Dataset, error) {
	dataset, err := p.client.GetDataset(ctx, name)
	if err != nil {
		return nil, err
	}

	return &llmops.Dataset{
		ID:          dataset.ID,
		Name:        dataset.Name,
		Description: dataset.Description,
		CreatedAt:   dataset.CreatedAt,
		UpdatedAt:   dataset.UpdatedAt,
	}, nil
}

// AddDatasetItems adds items to a dataset.
func (p *Provider) AddDatasetItems(ctx context.Context, datasetName string, items []llmops.DatasetItem) error {
	dataset, err := p.client.GetDataset(ctx, datasetName)
	if err != nil {
		return err
	}

	sdkItems := make([]sdk.DatasetItem, len(items))
	for i, item := range items {
		sdkItems[i] = sdk.DatasetItem{
			Input:    item.Input,
			Expected: item.Expected,
		}
	}

	return p.client.AddDatasetItems(ctx, dataset.ID, sdkItems)
}

// ListDatasets lists datasets.
func (p *Provider) ListDatasets(ctx context.Context, opts ...llmops.ListOption) ([]*llmops.Dataset, error) {
	cfg := llmops.ApplyListOptions(opts...)

	datasets, err := p.client.ListDatasets(ctx, cfg.Limit, cfg.Offset)
	if err != nil {
		return nil, err
	}

	result := make([]*llmops.Dataset, len(datasets))
	for i, ds := range datasets {
		result[i] = &llmops.Dataset{
			ID:          ds.ID,
			Name:        ds.Name,
			Description: ds.Description,
			CreatedAt:   ds.CreatedAt,
			UpdatedAt:   ds.UpdatedAt,
		}
	}
	return result, nil
}

// CreateProject creates a new project.
func (p *Provider) CreateProject(ctx context.Context, name string, opts ...llmops.ProjectOption) (*llmops.Project, error) {
	cfg := &llmops.ProjectOptions{}
	for _, opt := range opts {
		opt(cfg)
	}

	project, err := p.client.CreateProject(ctx, name)
	if err != nil {
		return nil, err
	}

	return &llmops.Project{
		ID:   project.ID,
		Name: project.Name,
	}, nil
}

// GetProject gets a project by name.
func (p *Provider) GetProject(ctx context.Context, name string) (*llmops.Project, error) {
	project, err := p.client.GetProject(ctx, name)
	if err != nil {
		return nil, err
	}

	return &llmops.Project{
		ID:   project.ID,
		Name: project.Name,
	}, nil
}

// ListProjects lists projects.
func (p *Provider) ListProjects(ctx context.Context, opts ...llmops.ListOption) ([]*llmops.Project, error) {
	cfg := llmops.ApplyListOptions(opts...)

	projects, err := p.client.ListProjects(ctx, cfg.Limit, cfg.Offset)
	if err != nil {
		return nil, err
	}

	result := make([]*llmops.Project, len(projects))
	for i, proj := range projects {
		result[i] = &llmops.Project{
			ID:   proj.ID,
			Name: proj.Name,
		}
	}
	return result, nil
}

// SetProject sets the current project.
func (p *Provider) SetProject(ctx context.Context, name string) error {
	// Phoenix doesn't have a direct "set project" - it's done via endpoint/headers
	return nil
}

// mapSpanOptions converts llmops span options to SDK options.
func mapSpanOptions(cfg *llmops.SpanOptions) []sdk.SpanOption {
	sdkOpts := []sdk.SpanOption{}

	if cfg.Type != "" {
		sdkOpts = append(sdkOpts, sdk.WithSpanKind(mapSpanType(cfg.Type)))
	}
	if cfg.Input != nil {
		sdkOpts = append(sdkOpts, sdk.WithSpanInput(cfg.Input))
	}
	if cfg.Metadata != nil {
		sdkOpts = append(sdkOpts, sdk.WithSpanMetadata(cfg.Metadata))
	}
	if cfg.Model != "" {
		sdkOpts = append(sdkOpts, sdk.WithModel(cfg.Model))
	}
	if cfg.Provider != "" {
		sdkOpts = append(sdkOpts, sdk.WithProvider(cfg.Provider))
	}
	if cfg.Usage != nil {
		sdkOpts = append(sdkOpts, sdk.WithUsage(
			cfg.Usage.PromptTokens,
			cfg.Usage.CompletionTokens,
			cfg.Usage.TotalTokens,
		))
	}

	return sdkOpts
}

// mapSpanType maps llmops span type to Phoenix span kind.
func mapSpanType(t llmops.SpanType) sdk.SpanKind {
	switch t {
	case llmops.SpanTypeLLM:
		return sdk.SpanKindLLM
	case llmops.SpanTypeTool:
		return sdk.SpanKindTool
	case llmops.SpanTypeRetrieval:
		return sdk.SpanKindRetriever
	case llmops.SpanTypeAgent:
		return sdk.SpanKindAgent
	case llmops.SpanTypeChain:
		return sdk.SpanKindChain
	case llmops.SpanTypeGuardrail:
		return sdk.SpanKindGuardrail
	default:
		return sdk.SpanKindUnknown
	}
}

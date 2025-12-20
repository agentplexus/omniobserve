package langfuse

import (
	"context"
	"fmt"
	"net/url"
)

// CreateDataset creates a new dataset.
func (c *Client) CreateDataset(ctx context.Context, name string, opts ...DatasetOption) (*Dataset, error) {
	cfg := &datasetConfig{}
	for _, opt := range opts {
		opt(cfg)
	}

	req := map[string]any{
		"name": name,
	}
	if cfg.description != "" {
		req["description"] = cfg.description
	}
	if cfg.metadata != nil {
		req["metadata"] = cfg.metadata
	}

	var result Dataset
	err := c.doPost(ctx, "/api/public/v2/datasets", req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetDataset retrieves a dataset by name.
func (c *Client) GetDataset(ctx context.Context, name string) (*Dataset, error) {
	var result Dataset
	err := c.doGet(ctx, "/api/public/v2/datasets/"+url.PathEscape(name), &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ListDatasets lists datasets with pagination.
func (c *Client) ListDatasets(ctx context.Context, limit, page int) ([]Dataset, error) {
	path := fmt.Sprintf("/api/public/v2/datasets?limit=%d&page=%d", limit, page)

	var result struct {
		Data []Dataset `json:"data"`
	}
	err := c.doGet(ctx, path, &result)
	if err != nil {
		return nil, err
	}
	return result.Data, nil
}

// CreateDatasetItem creates an item in a dataset.
func (c *Client) CreateDatasetItem(ctx context.Context, datasetName string, item DatasetItem) (*DatasetItem, error) {
	req := map[string]any{
		"datasetName": datasetName,
	}
	if item.Input != nil {
		req["input"] = item.Input
	}
	if item.ExpectedOutput != nil {
		req["expectedOutput"] = item.ExpectedOutput
	}
	if item.Metadata != nil {
		req["metadata"] = item.Metadata
	}
	if item.SourceTraceID != "" {
		req["sourceTraceId"] = item.SourceTraceID
	}
	if item.SourceObservationID != "" {
		req["sourceObservationId"] = item.SourceObservationID
	}

	var result DatasetItem
	err := c.doPost(ctx, "/api/public/dataset-items", req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetDatasetItems retrieves items from a dataset.
func (c *Client) GetDatasetItems(ctx context.Context, datasetName string, limit, page int) ([]DatasetItem, error) {
	path := fmt.Sprintf("/api/public/dataset-items?datasetName=%s&limit=%d&page=%d",
		url.QueryEscape(datasetName), limit, page)

	var result struct {
		Data []DatasetItem `json:"data"`
	}
	err := c.doGet(ctx, path, &result)
	if err != nil {
		return nil, err
	}
	return result.Data, nil
}

// CreateDatasetRun creates a dataset run (experiment).
func (c *Client) CreateDatasetRun(ctx context.Context, datasetName, runName string, opts ...DatasetRunOption) (*DatasetRun, error) {
	cfg := &datasetRunConfig{}
	for _, opt := range opts {
		opt(cfg)
	}

	req := map[string]any{
		"datasetName": datasetName,
		"name":        runName,
	}
	if cfg.metadata != nil {
		req["metadata"] = cfg.metadata
	}

	var result DatasetRun
	err := c.doPost(ctx, "/api/public/dataset-run-items", req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// LinkTraceToDatasetItem links a trace to a dataset item for a run.
func (c *Client) LinkTraceToDatasetItem(ctx context.Context, datasetItemID, traceID string, runName string, observationID string) error {
	req := map[string]any{
		"datasetItemId": datasetItemID,
		"traceId":       traceID,
		"runName":       runName,
	}
	if observationID != "" {
		req["observationId"] = observationID
	}

	return c.doPost(ctx, "/api/public/dataset-run-items", req, nil)
}

// GetDatasetRuns retrieves runs for a dataset.
func (c *Client) GetDatasetRuns(ctx context.Context, datasetName string, limit, page int) ([]DatasetRun, error) {
	path := fmt.Sprintf("/api/public/datasets/%s/runs?limit=%d&page=%d",
		url.PathEscape(datasetName), limit, page)

	var result struct {
		Data []DatasetRun `json:"data"`
	}
	err := c.doGet(ctx, path, &result)
	if err != nil {
		return nil, err
	}
	return result.Data, nil
}

// datasetConfig holds dataset creation configuration.
type datasetConfig struct {
	description string
	metadata    map[string]any
}

// DatasetOption configures dataset creation.
type DatasetOption func(*datasetConfig)

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

// datasetRunConfig holds dataset run configuration.
type datasetRunConfig struct {
	metadata map[string]any
}

// DatasetRunOption configures dataset run creation.
type DatasetRunOption func(*datasetRunConfig)

// WithRunMetadata sets the run metadata.
func WithRunMetadata(metadata map[string]any) DatasetRunOption {
	return func(c *datasetRunConfig) {
		c.metadata = metadata
	}
}

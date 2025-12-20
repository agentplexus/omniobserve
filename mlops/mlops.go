// Package mlops provides interfaces for ML operations platforms.
// This package will expand to include experiment tracking, model registry,
// dataset versioning, and other MLOps capabilities.
//
// Planned providers:
//   - MLflow
//   - Weights & Biases
//   - DVC
//   - Neptune
package mlops

import (
	"context"
	"io"
	"time"
)

// Provider is the main interface for MLOps platforms.
type Provider interface {
	ExperimentTracker
	ModelRegistry
	ArtifactStore
	io.Closer

	// Name returns the provider name.
	Name() string
}

// ExperimentTracker handles ML experiment tracking.
type ExperimentTracker interface {
	// CreateExperiment creates a new experiment.
	CreateExperiment(ctx context.Context, name string, opts ...ExperimentOption) (*Experiment, error)

	// GetExperiment retrieves an experiment by name.
	GetExperiment(ctx context.Context, name string) (*Experiment, error)

	// ListExperiments lists experiments.
	ListExperiments(ctx context.Context, opts ...ListOption) ([]*Experiment, error)

	// StartRun starts a new run within an experiment.
	StartRun(ctx context.Context, experimentName string, opts ...RunOption) (*Run, error)

	// LogMetric logs a metric value.
	LogMetric(ctx context.Context, runID string, key string, value float64, step int) error

	// LogParam logs a parameter.
	LogParam(ctx context.Context, runID string, key string, value string) error

	// EndRun ends a run.
	EndRun(ctx context.Context, runID string, status RunStatus) error
}

// ModelRegistry handles model versioning and registry.
type ModelRegistry interface {
	// RegisterModel registers a model version.
	RegisterModel(ctx context.Context, name string, opts ...ModelOption) (*Model, error)

	// GetModel retrieves a model by name.
	GetModel(ctx context.Context, name string, version ...string) (*Model, error)

	// ListModels lists registered models.
	ListModels(ctx context.Context, opts ...ListOption) ([]*Model, error)

	// TransitionModelStage transitions a model version to a new stage.
	TransitionModelStage(ctx context.Context, name string, version string, stage ModelStage) error
}

// ArtifactStore handles artifact storage.
type ArtifactStore interface {
	// LogArtifact uploads an artifact.
	LogArtifact(ctx context.Context, runID string, localPath string, artifactPath string) error

	// DownloadArtifact downloads an artifact.
	DownloadArtifact(ctx context.Context, runID string, artifactPath string, destPath string) error

	// ListArtifacts lists artifacts for a run.
	ListArtifacts(ctx context.Context, runID string, path string) ([]*Artifact, error)
}

// Experiment represents an ML experiment.
type Experiment struct {
	ID          string         `json:"id"`
	Name        string         `json:"name"`
	Description string         `json:"description,omitempty"`
	Tags        map[string]string `json:"tags,omitempty"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
}

// Run represents a single run within an experiment.
type Run struct {
	ID             string            `json:"id"`
	ExperimentID   string            `json:"experiment_id"`
	Name           string            `json:"name,omitempty"`
	Status         RunStatus         `json:"status"`
	StartTime      time.Time         `json:"start_time"`
	EndTime        *time.Time        `json:"end_time,omitempty"`
	Params         map[string]string `json:"params,omitempty"`
	Metrics        map[string]float64 `json:"metrics,omitempty"`
	Tags           map[string]string `json:"tags,omitempty"`
	ArtifactURI    string            `json:"artifact_uri,omitempty"`
}

// RunStatus represents the status of a run.
type RunStatus string

const (
	RunStatusRunning   RunStatus = "RUNNING"
	RunStatusScheduled RunStatus = "SCHEDULED"
	RunStatusFinished  RunStatus = "FINISHED"
	RunStatusFailed    RunStatus = "FAILED"
	RunStatusKilled    RunStatus = "KILLED"
)

// Model represents a registered model.
type Model struct {
	ID          string         `json:"id"`
	Name        string         `json:"name"`
	Version     string         `json:"version"`
	Stage       ModelStage     `json:"stage"`
	Description string         `json:"description,omitempty"`
	RunID       string         `json:"run_id,omitempty"`
	Source      string         `json:"source,omitempty"`
	Tags        map[string]string `json:"tags,omitempty"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
}

// ModelStage represents the deployment stage of a model.
type ModelStage string

const (
	ModelStageNone        ModelStage = "None"
	ModelStageStaging     ModelStage = "Staging"
	ModelStageProduction  ModelStage = "Production"
	ModelStageArchived    ModelStage = "Archived"
)

// Artifact represents a stored artifact.
type Artifact struct {
	Path     string `json:"path"`
	IsDir    bool   `json:"is_dir"`
	FileSize int64  `json:"file_size,omitempty"`
}

// Options

// ExperimentOption configures experiment creation.
type ExperimentOption func(*experimentConfig)

type experimentConfig struct {
	description string
	tags        map[string]string
}

// WithExperimentDescription sets the experiment description.
func WithExperimentDescription(desc string) ExperimentOption {
	return func(c *experimentConfig) {
		c.description = desc
	}
}

// WithExperimentTags sets experiment tags.
func WithExperimentTags(tags map[string]string) ExperimentOption {
	return func(c *experimentConfig) {
		c.tags = tags
	}
}

// RunOption configures run creation.
type RunOption func(*runConfig)

type runConfig struct {
	name   string
	tags   map[string]string
	params map[string]string
}

// WithRunName sets the run name.
func WithRunName(name string) RunOption {
	return func(c *runConfig) {
		c.name = name
	}
}

// WithRunTags sets run tags.
func WithRunTags(tags map[string]string) RunOption {
	return func(c *runConfig) {
		c.tags = tags
	}
}

// WithRunParams sets initial run parameters.
func WithRunParams(params map[string]string) RunOption {
	return func(c *runConfig) {
		c.params = params
	}
}

// ModelOption configures model registration.
type ModelOption func(*modelConfig)

type modelConfig struct {
	description string
	runID       string
	source      string
	tags        map[string]string
}

// WithModelDescription sets the model description.
func WithModelDescription(desc string) ModelOption {
	return func(c *modelConfig) {
		c.description = desc
	}
}

// WithModelRunID sets the source run ID.
func WithModelRunID(runID string) ModelOption {
	return func(c *modelConfig) {
		c.runID = runID
	}
}

// WithModelSource sets the model source path.
func WithModelSource(source string) ModelOption {
	return func(c *modelConfig) {
		c.source = source
	}
}

// ListOption configures list operations.
type ListOption func(*listConfig)

type listConfig struct {
	limit   int
	offset  int
	orderBy string
	filter  map[string]any
}

// WithLimit sets the maximum results.
func WithLimit(limit int) ListOption {
	return func(c *listConfig) {
		c.limit = limit
	}
}

// WithOffset sets the pagination offset.
func WithOffset(offset int) ListOption {
	return func(c *listConfig) {
		c.offset = offset
	}
}

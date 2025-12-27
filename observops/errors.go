package observops

import (
	"errors"
	"fmt"
)

// Sentinel errors for common failure cases.
var (
	// ErrProviderDisabled indicates the provider is disabled.
	ErrProviderDisabled = errors.New("observops: provider is disabled")

	// ErrMissingServiceName indicates a service name is required but not provided.
	ErrMissingServiceName = errors.New("observops: service name is required")

	// ErrMissingEndpoint indicates an endpoint is required but not provided.
	ErrMissingEndpoint = errors.New("observops: endpoint is required")

	// ErrMissingAPIKey indicates an API key is required but not provided.
	ErrMissingAPIKey = errors.New("observops: API key is required")

	// ErrShutdown indicates the provider has been shut down.
	ErrShutdown = errors.New("observops: provider has been shut down")

	// ErrNotSupported indicates the operation is not supported by this provider.
	ErrNotSupported = errors.New("observops: operation not supported")
)

// ProviderError wraps errors from a specific provider.
type ProviderError struct {
	Provider string
	Op       string
	Err      error
}

func (e *ProviderError) Error() string {
	if e.Op != "" {
		return fmt.Sprintf("observops: %s: %s: %v", e.Provider, e.Op, e.Err)
	}
	return fmt.Sprintf("observops: %s: %v", e.Provider, e.Err)
}

func (e *ProviderError) Unwrap() error {
	return e.Err
}

// WrapError wraps an error with provider context.
func WrapError(provider, op string, err error) error {
	if err == nil {
		return nil
	}
	return &ProviderError{
		Provider: provider,
		Op:       op,
		Err:      err,
	}
}

// ExportError represents an error during telemetry export.
type ExportError struct {
	Signal  string // "metrics", "traces", or "logs"
	Count   int    // Number of items that failed to export
	Details string
	Err     error
}

func (e *ExportError) Error() string {
	if e.Count > 0 {
		return fmt.Sprintf("observops: failed to export %d %s: %s", e.Count, e.Signal, e.Details)
	}
	return fmt.Sprintf("observops: failed to export %s: %s", e.Signal, e.Details)
}

func (e *ExportError) Unwrap() error {
	return e.Err
}

// ConfigError represents a configuration error.
type ConfigError struct {
	Field   string
	Message string
}

func (e *ConfigError) Error() string {
	return fmt.Sprintf("observops: config error for %s: %s", e.Field, e.Message)
}

// WrapNotSupported wraps an error indicating a feature is not supported.
func WrapNotSupported(provider, feature string) error {
	return &ProviderError{
		Provider: provider,
		Op:       feature,
		Err:      ErrNotSupported,
	}
}

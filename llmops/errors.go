package llmops

import (
	"errors"
	"fmt"
)

// Sentinel errors for common conditions.
var (
	// Configuration errors
	ErrMissingAPIKey    = errors.New("llmops: missing API key")
	ErrMissingEndpoint  = errors.New("llmops: missing endpoint URL")
	ErrMissingWorkspace = errors.New("llmops: missing workspace")
	ErrMissingProject   = errors.New("llmops: missing project name")

	// State errors
	ErrTracingDisabled = errors.New("llmops: tracing is disabled")
	ErrNoActiveTrace   = errors.New("llmops: no active trace in context")
	ErrNoActiveSpan    = errors.New("llmops: no active span in context")
	ErrAlreadyEnded    = errors.New("llmops: trace or span already ended")

	// Not found errors
	ErrTraceNotFound      = errors.New("llmops: trace not found")
	ErrSpanNotFound       = errors.New("llmops: span not found")
	ErrDatasetNotFound    = errors.New("llmops: dataset not found")
	ErrPromptNotFound     = errors.New("llmops: prompt not found")
	ErrProjectNotFound    = errors.New("llmops: project not found")
	ErrExperimentNotFound = errors.New("llmops: experiment not found")

	// Validation errors
	ErrInvalidInput     = errors.New("llmops: invalid input")
	ErrInvalidSpanType  = errors.New("llmops: invalid span type")
	ErrInvalidMetric    = errors.New("llmops: invalid metric")

	// Provider errors
	ErrProviderNotFound   = errors.New("llmops: provider not found")
	ErrNotImplemented     = errors.New("llmops: not implemented")
	ErrCapabilityNotSupported = errors.New("llmops: capability not supported")
)

// APIError represents an error from a provider API.
type APIError struct {
	StatusCode int
	Message    string
	Details    map[string]any
	Provider   string
	Err        error
}

func (e *APIError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s API error (status %d): %s: %v", e.Provider, e.StatusCode, e.Message, e.Err)
	}
	return fmt.Sprintf("%s API error (status %d): %s", e.Provider, e.StatusCode, e.Message)
}

func (e *APIError) Unwrap() error {
	return e.Err
}

// NewAPIError creates a new API error.
func NewAPIError(provider string, statusCode int, message string, err error) *APIError {
	return &APIError{
		Provider:   provider,
		StatusCode: statusCode,
		Message:    message,
		Err:        err,
	}
}

// IsNotFound returns true if the error indicates a resource was not found.
func IsNotFound(err error) bool {
	if err == nil {
		return false
	}
	if errors.Is(err, ErrTraceNotFound) ||
		errors.Is(err, ErrSpanNotFound) ||
		errors.Is(err, ErrDatasetNotFound) ||
		errors.Is(err, ErrPromptNotFound) ||
		errors.Is(err, ErrProjectNotFound) ||
		errors.Is(err, ErrExperimentNotFound) {
		return true
	}
	var apiErr *APIError
	if errors.As(err, &apiErr) {
		return apiErr.StatusCode == 404
	}
	return false
}

// IsUnauthorized returns true if the error indicates an authentication failure.
func IsUnauthorized(err error) bool {
	var apiErr *APIError
	if errors.As(err, &apiErr) {
		return apiErr.StatusCode == 401
	}
	return false
}

// IsRateLimited returns true if the error indicates rate limiting.
func IsRateLimited(err error) bool {
	var apiErr *APIError
	if errors.As(err, &apiErr) {
		return apiErr.StatusCode == 429
	}
	return false
}

// IsDisabled returns true if tracing is disabled.
func IsDisabled(err error) bool {
	return errors.Is(err, ErrTracingDisabled)
}

// IsNotImplemented returns true if the operation is not implemented.
func IsNotImplemented(err error) bool {
	return errors.Is(err, ErrNotImplemented)
}

// WrapNotImplemented wraps an operation name in a not implemented error.
func WrapNotImplemented(provider, operation string) error {
	return fmt.Errorf("%w: %s does not support %s", ErrNotImplemented, provider, operation)
}

// WrapCapabilityNotSupported wraps a capability in a not supported error.
func WrapCapabilityNotSupported(provider string, cap Capability) error {
	return fmt.Errorf("%w: %s does not support %s", ErrCapabilityNotSupported, provider, cap)
}

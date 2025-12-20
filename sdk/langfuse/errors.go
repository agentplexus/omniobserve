package langfuse

import (
	"errors"
	"fmt"
)

// Sentinel errors.
var (
	ErrMissingPublicKey   = errors.New("langfuse: missing public key")
	ErrMissingSecretKey   = errors.New("langfuse: missing secret key")
	ErrNoActiveTrace      = errors.New("langfuse: no active trace in context")
	ErrNoActiveSpan       = errors.New("langfuse: no active span in context")
	ErrNoActiveGeneration = errors.New("langfuse: no active generation in context")
	ErrTraceNotFound      = errors.New("langfuse: trace not found")
	ErrDatasetNotFound    = errors.New("langfuse: dataset not found")
	ErrPromptNotFound     = errors.New("langfuse: prompt not found")
)

// APIError represents an error from the Langfuse API.
type APIError struct {
	StatusCode int
	Message    string
	Details    map[string]any
}

func (e *APIError) Error() string {
	return fmt.Sprintf("langfuse API error (status %d): %s", e.StatusCode, e.Message)
}

// IsNotFound returns true if the error is a not found error.
func IsNotFound(err error) bool {
	if err == nil {
		return false
	}
	if errors.Is(err, ErrTraceNotFound) ||
		errors.Is(err, ErrDatasetNotFound) ||
		errors.Is(err, ErrPromptNotFound) {
		return true
	}
	var apiErr *APIError
	if errors.As(err, &apiErr) {
		return apiErr.StatusCode == 404
	}
	return false
}

// IsUnauthorized returns true if the error is an authentication error.
func IsUnauthorized(err error) bool {
	var apiErr *APIError
	if errors.As(err, &apiErr) {
		return apiErr.StatusCode == 401
	}
	return false
}

// IsRateLimited returns true if the error is a rate limit error.
func IsRateLimited(err error) bool {
	var apiErr *APIError
	if errors.As(err, &apiErr) {
		return apiErr.StatusCode == 429
	}
	return false
}

package phoenix

import (
	"errors"
	"fmt"
)

// Sentinel errors.
var (
	ErrMissingEndpoint = errors.New("phoenix: missing endpoint")
	ErrMissingAPIKey   = errors.New("phoenix: missing API key (required for cloud)")
	ErrNoActiveTrace   = errors.New("phoenix: no active trace in context")
	ErrNoActiveSpan    = errors.New("phoenix: no active span in context")
	ErrTraceNotFound   = errors.New("phoenix: trace not found")
	ErrSpanNotFound    = errors.New("phoenix: span not found")
	ErrDatasetNotFound = errors.New("phoenix: dataset not found")
	ErrProjectNotFound = errors.New("phoenix: project not found")
)

// APIError represents an error from the Phoenix API.
type APIError struct {
	StatusCode int
	Message    string
	Details    map[string]any
}

func (e *APIError) Error() string {
	return fmt.Sprintf("phoenix API error (status %d): %s", e.StatusCode, e.Message)
}

// IsNotFound returns true if the error is a not found error.
func IsNotFound(err error) bool {
	if err == nil {
		return false
	}
	if errors.Is(err, ErrTraceNotFound) ||
		errors.Is(err, ErrSpanNotFound) ||
		errors.Is(err, ErrDatasetNotFound) ||
		errors.Is(err, ErrProjectNotFound) {
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

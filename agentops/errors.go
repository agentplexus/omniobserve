package agentops

import (
	"errors"
	"fmt"
)

// Sentinel errors for common failure cases.
var (
	// ErrNotFound indicates the requested entity was not found.
	ErrNotFound = errors.New("agentops: not found")

	// ErrMissingDSN indicates a DSN is required but not provided.
	ErrMissingDSN = errors.New("agentops: DSN is required")

	// ErrInvalidStatus indicates an invalid status transition.
	ErrInvalidStatus = errors.New("agentops: invalid status transition")

	// ErrAlreadyCompleted indicates the entity is already completed.
	ErrAlreadyCompleted = errors.New("agentops: already completed")

	// ErrStoreClosed indicates the store has been closed.
	ErrStoreClosed = errors.New("agentops: store is closed")

	// ErrConnectionFailed indicates a connection failure.
	ErrConnectionFailed = errors.New("agentops: connection failed")
)

// StoreError wraps errors from a specific store.
type StoreError struct {
	Store string
	Op    string
	Err   error
}

func (e *StoreError) Error() string {
	if e.Op != "" {
		return fmt.Sprintf("agentops: %s: %s: %v", e.Store, e.Op, e.Err)
	}
	return fmt.Sprintf("agentops: %s: %v", e.Store, e.Err)
}

func (e *StoreError) Unwrap() error {
	return e.Err
}

// WrapError wraps an error with store context.
func WrapError(store, op string, err error) error {
	if err == nil {
		return nil
	}
	return &StoreError{
		Store: store,
		Op:    op,
		Err:   err,
	}
}

// IsNotFound returns true if the error indicates a not found condition.
func IsNotFound(err error) bool {
	return errors.Is(err, ErrNotFound)
}

// IsAlreadyCompleted returns true if the error indicates already completed.
func IsAlreadyCompleted(err error) bool {
	return errors.Is(err, ErrAlreadyCompleted)
}

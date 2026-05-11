package sendry

import "fmt"

// Error is the base error type for all API errors returned by the Sendr API.
// It implements the error interface and includes the HTTP status code, error code,
// and human-readable message from the API response.
type Error struct {
	// StatusCode is the HTTP status code returned by the API.
	StatusCode int
	// Code is the machine-readable error code (e.g. "rate_limit_error").
	Code string
	// Message is the human-readable error description.
	Message string
}

// Error implements the error interface.
func (e *Error) Error() string {
	return fmt.Sprintf("sendr: %s (status=%d, code=%s)", e.Message, e.StatusCode, e.Code)
}

// AuthenticationError is returned when the API key is invalid or missing (HTTP 401).
type AuthenticationError struct {
	*Error
}

// ValidationError is returned when the request body or parameters fail validation (HTTP 422).
type ValidationError struct {
	*Error
	// Details contains field-level validation errors from the API.
	Details any
}

// RateLimitError is returned when the rate limit has been exceeded (HTTP 429).
type RateLimitError struct {
	*Error
	// RetryAfter is the number of seconds to wait before retrying, or 0 if not specified.
	RetryAfter int
}

// NotFoundError is returned when the requested resource does not exist (HTTP 404).
type NotFoundError struct {
	*Error
}

// NetworkError is returned when a network-level failure occurs (DNS, connection refused,
// timeout, etc.) and no HTTP response was received from the server.
type NetworkError struct {
	// Err is the underlying network error.
	Err error
}

// Error implements the error interface.
func (e *NetworkError) Error() string {
	return fmt.Sprintf("sendr: network error: %v", e.Err)
}

// Unwrap returns the underlying error so errors.Is and errors.As work correctly.
func (e *NetworkError) Unwrap() error {
	return e.Err
}

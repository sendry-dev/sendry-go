package sendry

import (
	"context"
	"time"
)

const (
	defaultMaxRetries  = 3
	defaultBaseBackoff = 1 * time.Second
	maxBackoff         = 30 * time.Second
)

// retryConfig holds the retry policy for a client.
type retryConfig struct {
	maxRetries int
}

// shouldRetry reports whether an error should trigger a retry attempt.
// Retries are performed for network errors and 5xx API errors.
// Client errors (4xx) are never retried.
func shouldRetry(err error) bool {
	if err == nil {
		return false
	}
	if _, ok := err.(*NetworkError); ok {
		return true
	}
	if apiErr, ok := err.(*APIError); ok {
		return apiErr.StatusCode >= 500
	}
	// Subtypes embed *APIError — check them explicitly.
	if authErr, ok := err.(*AuthenticationError); ok {
		return authErr.StatusCode >= 500
	}
	return false
}

// backoffDuration returns the exponential backoff duration for a given attempt
// (0-indexed). The sequence is 1s, 2s, 4s, 8s... capped at maxBackoff.
func backoffDuration(attempt int) time.Duration {
	d := defaultBaseBackoff * (1 << uint(attempt))
	if d > maxBackoff {
		d = maxBackoff
	}
	return d
}

// sleep pauses execution for d, respecting ctx cancellation.
func sleep(ctx context.Context, d time.Duration) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-time.After(d):
		return nil
	}
}

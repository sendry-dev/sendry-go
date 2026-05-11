package sendry

import (
	"testing"
)

func TestNewClientDefaults(t *testing.T) {
	c := NewClient("sn_live_test")
	if c.apiKey != "sn_live_test" {
		t.Fatalf("apiKey not set: got %q", c.apiKey)
	}
	if c.baseURL != defaultBaseURL {
		t.Fatalf("baseURL default mismatch: got %q", c.baseURL)
	}
	if c.Emails == nil || c.Domains == nil || c.Webhooks == nil {
		t.Fatal("resource accessors not wired")
	}
}

func TestWithBaseURL(t *testing.T) {
	c := NewClient("k", WithBaseURL("https://example.test"))
	if c.baseURL != "https://example.test" {
		t.Fatalf("WithBaseURL not applied: %q", c.baseURL)
	}
}

func TestAPIErrorError(t *testing.T) {
	e := &APIError{StatusCode: 422, Code: "validation_error", Message: "bad"}
	if got := e.Error(); got == "" {
		t.Fatal("Error() empty")
	}
}

func TestRateLimitErrorImplementsError(t *testing.T) {
	var err error = &RateLimitError{APIError: &APIError{StatusCode: 429, Code: "rate_limit", Message: "x"}, RetryAfter: 1}
	if err.Error() == "" {
		t.Fatal("RateLimitError.Error() empty")
	}
}

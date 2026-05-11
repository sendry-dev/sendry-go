package sendry

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const (
	defaultBaseURL = "https://api.sendry.online"
	defaultTimeout = 30 * time.Second
	sdkVersion     = "0.1.0"
	userAgent      = "sendry-go/" + sdkVersion
)

// Option is a functional option for configuring a Client.
type Option func(*Client)

// WithBaseURL overrides the API base URL. Useful for testing against a local server.
func WithBaseURL(baseURL string) Option {
	return func(c *Client) {
		c.baseURL = baseURL
	}
}

// WithHTTPClient replaces the default http.Client with a custom one.
func WithHTTPClient(httpClient *http.Client) Option {
	return func(c *Client) {
		c.httpClient = httpClient
	}
}

// WithTimeout sets the per-request timeout. Defaults to 30s.
func WithTimeout(d time.Duration) Option {
	return func(c *Client) {
		c.httpClient.Timeout = d
	}
}

// WithMaxRetries sets the maximum number of retry attempts for 5xx / network errors.
// Defaults to 3. Set to 0 to disable retries.
func WithMaxRetries(n int) Option {
	return func(c *Client) {
		c.retry.maxRetries = n
	}
}

// Client is the Sendry API client. Create one with NewClient and call methods on
// the resource fields (e.g. client.Emails.Send).
type Client struct {
	apiKey     string
	baseURL    string
	httpClient *http.Client
	retry      retryConfig

	// Resource accessors
	Emails       *EmailsResource
	Domains      *DomainsResource
	Templates    *TemplatesResource
	APIKeys      *APIKeysResource
	Webhooks     *WebhooksResource
	Analytics    *AnalyticsResource
	Suppression  *SuppressionResource
	Unsubscribes *UnsubscribesResource
	Contacts     *ContactsResource
	Audiences    *AudiencesResource
	Campaigns    *CampaignsResource
	Billing      *BillingResource
	Team         *TeamResource
}

// NewClient creates a new Sendry API client authenticated with the given API key.
// Pass functional options to customise the base URL, HTTP client, or timeout.
//
// Example:
//
//	client := sendry.NewClient("sn_live_abc123")
//
//	// With options
//	client := sendry.NewClient(
//	    "sn_live_abc123",
//	    sendry.WithTimeout(10*time.Second),
//	    sendry.WithBaseURL("https://api.sendry.online"),
//	)
func NewClient(apiKey string, opts ...Option) *Client {
	c := &Client{
		apiKey:  apiKey,
		baseURL: defaultBaseURL,
		httpClient: &http.Client{
			Timeout: defaultTimeout,
		},
		retry: retryConfig{
			maxRetries: defaultMaxRetries,
		},
	}

	for _, opt := range opts {
		opt(c)
	}

	// Wire up resource accessors
	c.Emails = &EmailsResource{client: c}
	c.Domains = &DomainsResource{client: c}
	c.Templates = &TemplatesResource{client: c}
	c.APIKeys = &APIKeysResource{client: c}
	c.Webhooks = &WebhooksResource{client: c}
	c.Analytics = &AnalyticsResource{client: c}
	c.Suppression = &SuppressionResource{client: c}
	c.Unsubscribes = &UnsubscribesResource{client: c}
	c.Contacts = &ContactsResource{client: c}
	c.Audiences = &AudiencesResource{client: c}
	c.Campaigns = &CampaignsResource{client: c}
	c.Billing = &BillingResource{client: c}
	c.Team = &TeamResource{client: c}

	return c
}

// apiErrorBody is the shape of the API's error response JSON.
type apiErrorBody struct {
	Error struct {
		Code    string `json:"code"`
		Message string `json:"message"`
		Details any    `json:"details"`
	} `json:"error"`
}

// do executes an HTTP request with retry logic and returns the decoded response body.
// A nil body pointer means the caller expects no response body (e.g. 204 No Content).
func (c *Client) do(ctx context.Context, method, path string, query url.Values, reqBody, respBody any) error {
	var (
		lastErr    error
		maxAttempts = c.retry.maxRetries + 1
	)

	for attempt := 0; attempt < maxAttempts; attempt++ {
		if attempt > 0 {
			retryDelay := backoffDuration(attempt - 1)
			// If the last error was a rate-limit with Retry-After, use that instead.
			if rl, ok := lastErr.(*RateLimitError); ok && rl.RetryAfter > 0 {
				retryDelay = time.Duration(rl.RetryAfter) * time.Second
			}
			if err := sleep(ctx, retryDelay); err != nil {
				return &NetworkError{Err: err}
			}
		}

		err := c.doOnce(ctx, method, path, query, reqBody, respBody)
		if err == nil {
			return nil
		}

		// Non-retryable errors: return immediately.
		if !shouldRetry(err) {
			return err
		}

		lastErr = err
	}

	return lastErr
}

// doOnce performs a single HTTP request without any retry logic.
func (c *Client) doOnce(ctx context.Context, method, path string, query url.Values, reqBody, respBody any) error {
	// Build URL
	rawURL := c.baseURL + path
	if len(query) > 0 {
		rawURL += "?" + query.Encode()
	}

	// Encode request body
	var bodyReader io.Reader
	if reqBody != nil {
		data, err := json.Marshal(reqBody)
		if err != nil {
			return fmt.Errorf("sendry: marshal request body: %w", err)
		}
		bodyReader = bytes.NewReader(data)
	}

	req, err := http.NewRequestWithContext(ctx, method, rawURL, bodyReader)
	if err != nil {
		return &NetworkError{Err: fmt.Errorf("build request: %w", err)}
	}

	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("User-Agent", userAgent)
	if reqBody != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return &NetworkError{Err: err}
	}
	defer resp.Body.Close()

	// 204 No Content — success, no body
	if resp.StatusCode == http.StatusNoContent {
		return nil
	}

	// Success
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		if respBody != nil {
			if err := json.NewDecoder(resp.Body).Decode(respBody); err != nil {
				return fmt.Errorf("sendry: decode response: %w", err)
			}
		}
		return nil
	}

	// Error response — parse body
	var errBody apiErrorBody
	_ = json.NewDecoder(resp.Body).Decode(&errBody)

	code := errBody.Error.Code
	if code == "" {
		code = "unknown_error"
	}
	msg := errBody.Error.Message
	if msg == "" {
		msg = resp.Status
	}

	base := &APIError{
		StatusCode: resp.StatusCode,
		Code:       code,
		Message:    msg,
	}

	switch resp.StatusCode {
	case http.StatusUnauthorized:
		return &AuthenticationError{APIError: base}
	case http.StatusNotFound:
		return &NotFoundError{APIError: base}
	case http.StatusUnprocessableEntity:
		return &ValidationError{APIError: base, Details: errBody.Error.Details}
	case http.StatusTooManyRequests:
		retryAfter := 0
		if raw := resp.Header.Get("Retry-After"); raw != "" {
			if n, err := strconv.Atoi(raw); err == nil {
				retryAfter = n
			}
		}
		return &RateLimitError{APIError: base, RetryAfter: retryAfter}
	}

	return base
}

// get performs a GET request and decodes the JSON response into out.
func (c *Client) get(ctx context.Context, path string, query url.Values, out any) error {
	return c.do(ctx, http.MethodGet, path, query, nil, out)
}

// post performs a POST request with a JSON body and decodes the JSON response into out.
func (c *Client) post(ctx context.Context, path string, body, out any) error {
	return c.do(ctx, http.MethodPost, path, nil, body, out)
}

// put performs a PUT request with a JSON body and decodes the JSON response into out.
func (c *Client) put(ctx context.Context, path string, body, out any) error {
	return c.do(ctx, http.MethodPut, path, nil, body, out)
}

// patch performs a PATCH request with a JSON body and decodes the JSON response into out.
func (c *Client) patch(ctx context.Context, path string, body, out any) error {
	return c.do(ctx, http.MethodPatch, path, nil, body, out)
}

// delete performs a DELETE request and decodes the JSON response into out.
func (c *Client) delete(ctx context.Context, path string, out any) error {
	return c.do(ctx, http.MethodDelete, path, nil, nil, out)
}

// buildQuery converts a map of string key/value pairs into a url.Values, omitting
// any entries whose value is the empty string.
func buildQuery(params map[string]string) url.Values {
	v := url.Values{}
	for k, val := range params {
		if val != "" {
			v.Set(k, val)
		}
	}
	return v
}

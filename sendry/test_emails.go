package sendry

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

// TestEmailsResource provides methods for retrieving test-mode emails.
type TestEmailsResource struct {
	client *Client
}

// List returns test emails captured in test mode with cursor pagination.
//
// Example:
//
//	page, err := client.TestEmails.List(ctx, nil)
func (r *TestEmailsResource) List(ctx context.Context, params *PaginationParams) (*PaginatedResponse[TestEmailSummary], error) {
	q := url.Values{}
	if params != nil {
		if params.Limit != nil {
			q.Set("limit", strconv.Itoa(*params.Limit))
		}
		if params.Cursor != nil {
			q.Set("cursor", *params.Cursor)
		}
	}
	var out PaginatedResponse[TestEmailSummary]
	if err := r.client.get(ctx, "/v1/test-emails", q, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Get returns a specific test email by ID including HTML and plain-text body.
//
// Example:
//
//	email, err := client.TestEmails.Get(ctx, "te_abc123")
func (r *TestEmailsResource) Get(ctx context.Context, id string) (*TestEmail, error) {
	var out TestEmail
	if err := r.client.get(ctx, fmt.Sprintf("/v1/test-emails/%s", url.PathEscape(id)), nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

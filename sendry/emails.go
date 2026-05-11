package sendry

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

// EmailsResource provides methods for sending and managing emails.
type EmailsResource struct {
	client *Client
}

// Send sends a single transactional email.
//
// Example:
//
//	resp, err := client.Emails.Send(ctx, sendr.SendEmailParams{
//	    From:    "hello@example.com",
//	    To:      "user@example.com",
//	    Subject: "Hello",
//	    HTML:    "<p>World</p>",
//	})
func (r *EmailsResource) Send(ctx context.Context, params SendEmailParams) (*SendEmailResponse, error) {
	var out SendEmailResponse
	if err := r.client.post(ctx, "/v1/emails", params, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Get retrieves a single email by its ID.
//
// Example:
//
//	email, err := client.Emails.Get(ctx, "em_abc123")
func (r *EmailsResource) Get(ctx context.Context, id string) (*Email, error) {
	var out Email
	if err := r.client.get(ctx, fmt.Sprintf("/v1/emails/%s", url.PathEscape(id)), nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// List returns a cursor-paginated list of emails.
//
// Example:
//
//	page, err := client.Emails.List(ctx, &sendr.ListEmailsParams{
//	    PaginationParams: sendr.PaginationParams{Limit: sendr.IntPtr(25)},
//	})
func (r *EmailsResource) List(ctx context.Context, params *ListEmailsParams) (*PaginatedResponse[Email], error) {
	q := url.Values{}
	if params != nil {
		if params.Limit != nil {
			q.Set("limit", strconv.Itoa(*params.Limit))
		}
		if params.Cursor != nil {
			q.Set("cursor", *params.Cursor)
		}
		if params.Status != "" {
			q.Set("status", params.Status)
		}
	}
	var out PaginatedResponse[Email]
	if err := r.client.get(ctx, "/v1/emails", q, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// SendBatch sends a batch of up to 100 emails in a single API call.
//
// Example:
//
//	result, err := client.Emails.SendBatch(ctx, sendr.SendBatchParams{
//	    From: "hello@example.com",
//	    Emails: []sendr.BatchEmailItem{
//	        {To: "a@example.com", Subject: "Hi A", HTML: "<p>A</p>"},
//	    },
//	})
func (r *EmailsResource) SendBatch(ctx context.Context, params SendBatchParams) (*BatchEmailResponse, error) {
	var out BatchEmailResponse
	if err := r.client.post(ctx, "/v1/emails/batch", params, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// SendMarketing sends a marketing email with built-in unsubscribe support.
//
// Example:
//
//	resp, err := client.Emails.SendMarketing(ctx, sendr.SendMarketingEmailParams{
//	    From:           "news@example.com",
//	    To:             "subscriber@example.com",
//	    Subject:        "Newsletter",
//	    HTML:           "<p>News!</p>",
//	    UnsubscribeURL: "https://example.com/unsub",
//	})
func (r *EmailsResource) SendMarketing(ctx context.Context, params SendMarketingEmailParams) (*SendEmailResponse, error) {
	var out SendEmailResponse
	if err := r.client.post(ctx, "/v1/emails/marketing", params, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Cancel cancels a queued email. Only emails that have not yet been sent can be cancelled.
//
// Example:
//
//	result, err := client.Emails.Cancel(ctx, "em_abc123")
func (r *EmailsResource) Cancel(ctx context.Context, id string) (*CancelEmailResponse, error) {
	var out CancelEmailResponse
	if err := r.client.post(ctx, fmt.Sprintf("/v1/emails/%s/cancel", url.PathEscape(id)), nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// IntPtr is a helper that returns a pointer to an int literal.
func IntPtr(n int) *int { return &n }

// StringPtr is a helper that returns a pointer to a string literal.
func StringPtr(s string) *string { return &s }

// BoolPtr is a helper that returns a pointer to a bool literal.
func BoolPtr(b bool) *bool { return &b }

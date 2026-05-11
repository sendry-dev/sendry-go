package sendry

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

// WebhooksResource provides methods for managing webhook endpoints.
type WebhooksResource struct {
	client *Client
}

// Create registers a new webhook endpoint.
//
// Example:
//
//	hook, err := client.Webhooks.Create(ctx, sendry.CreateWebhookParams{
//	    URL:    "https://example.com/webhook",
//	    Events: []string{"email.delivered", "email.bounced"},
//	})
//	// hook.Secret is the HMAC secret for verifying payloads.
func (r *WebhooksResource) Create(ctx context.Context, params CreateWebhookParams) (*Webhook, error) {
	var out Webhook
	if err := r.client.post(ctx, "/v1/webhooks", params, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// List returns all webhook endpoints.
//
// Example:
//
//	page, err := client.Webhooks.List(ctx, nil)
func (r *WebhooksResource) List(ctx context.Context, params *PaginationParams) (*PaginatedResponse[WebhookListItem], error) {
	q := url.Values{}
	if params != nil {
		if params.Limit != nil {
			q.Set("limit", strconv.Itoa(*params.Limit))
		}
		if params.Cursor != nil {
			q.Set("cursor", *params.Cursor)
		}
	}
	var out PaginatedResponse[WebhookListItem]
	if err := r.client.get(ctx, "/v1/webhooks", q, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Get retrieves a webhook by its ID. The response includes the signing secret.
//
// Example:
//
//	hook, err := client.Webhooks.Get(ctx, "wh_abc123")
func (r *WebhooksResource) Get(ctx context.Context, id string) (*Webhook, error) {
	var out Webhook
	if err := r.client.get(ctx, fmt.Sprintf("/v1/webhooks/%s", url.PathEscape(id)), nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Update modifies a webhook endpoint.
//
// Example:
//
//	updated, err := client.Webhooks.Update(ctx, "wh_abc123", sendry.UpdateWebhookParams{
//	    Active: sendry.BoolPtr(false),
//	})
func (r *WebhooksResource) Update(ctx context.Context, id string, params UpdateWebhookParams) (*Webhook, error) {
	var out Webhook
	if err := r.client.put(ctx, fmt.Sprintf("/v1/webhooks/%s", url.PathEscape(id)), params, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Remove deletes a webhook endpoint.
//
// Example:
//
//	_, err := client.Webhooks.Remove(ctx, "wh_abc123")
func (r *WebhooksResource) Remove(ctx context.Context, id string) (*DeleteResponse, error) {
	var out DeleteResponse
	if err := r.client.delete(ctx, fmt.Sprintf("/v1/webhooks/%s", url.PathEscape(id)), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

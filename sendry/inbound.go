package sendry

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

// InboundResource provides methods for retrieving inbound emails and managing
// the inbound webhook forwarding configuration.
type InboundResource struct {
	client *Client
}

// List returns received inbound emails with cursor-based pagination.
//
// Example:
//
//	page, err := client.Inbound.List(ctx, nil)
func (r *InboundResource) List(ctx context.Context, params *PaginationParams) (*PaginatedResponse[InboundEmail], error) {
	q := url.Values{}
	if params != nil {
		if params.Limit != nil {
			q.Set("limit", strconv.Itoa(*params.Limit))
		}
		if params.Cursor != nil {
			q.Set("cursor", *params.Cursor)
		}
	}
	var out PaginatedResponse[InboundEmail]
	if err := r.client.get(ctx, "/v1/inbound", q, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Get retrieves a specific inbound email by ID including body and attachments.
//
// Example:
//
//	email, err := client.Inbound.Get(ctx, "inb_abc123")
func (r *InboundResource) Get(ctx context.Context, id string) (*InboundEmail, error) {
	var out InboundEmail
	if err := r.client.get(ctx, fmt.Sprintf("/v1/inbound/%s", url.PathEscape(id)), nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetConfig returns the inbound webhook forwarding configuration.
//
// Example:
//
//	config, err := client.Inbound.GetConfig(ctx)
func (r *InboundResource) GetConfig(ctx context.Context) (*InboundConfig, error) {
	var out InboundConfig
	if err := r.client.get(ctx, "/v1/inbound/config", nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// UpdateConfig updates the inbound webhook forwarding configuration.
//
// Example:
//
//	url := "https://api.acme.com/webhooks/inbound"
//	secret := "my-hmac-secret"
//	config, err := client.Inbound.UpdateConfig(ctx, sendry.UpdateInboundConfigParams{
//	    URL: &url, Secret: &secret,
//	})
func (r *InboundResource) UpdateConfig(ctx context.Context, params UpdateInboundConfigParams) (*InboundConfig, error) {
	var out InboundConfig
	if err := r.client.put(ctx, "/v1/inbound/config", params, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

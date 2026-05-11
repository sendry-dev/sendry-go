package sendry

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

// APIKeysResource provides methods for managing API keys.
type APIKeysResource struct {
	client *Client
}

// Create creates a new API key. The full key value is only returned once in
// the response — store it securely.
//
// Example:
//
//	created, err := client.APIKeys.Create(ctx, sendry.CreateAPIKeyParams{
//	    Name:  "Production Key",
//	    Scope: sendry.APIKeyScopeSendingAccess,
//	})
//	// Store created.Key securely — it cannot be retrieved again.
func (r *APIKeysResource) Create(ctx context.Context, params CreateAPIKeyParams) (*APIKeyCreated, error) {
	var out APIKeyCreated
	if err := r.client.post(ctx, "/v1/api-keys", params, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// List returns all API keys. The key values are masked; only the prefix is shown.
//
// Example:
//
//	page, err := client.APIKeys.List(ctx, nil)
func (r *APIKeysResource) List(ctx context.Context, params *PaginationParams) (*PaginatedResponse[APIKey], error) {
	q := url.Values{}
	if params != nil {
		if params.Limit != nil {
			q.Set("limit", strconv.Itoa(*params.Limit))
		}
		if params.Cursor != nil {
			q.Set("cursor", *params.Cursor)
		}
	}
	var out PaginatedResponse[APIKey]
	if err := r.client.get(ctx, "/v1/api-keys", q, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Remove revokes (deletes) an API key.
//
// Example:
//
//	_, err := client.APIKeys.Remove(ctx, "key_abc123")
func (r *APIKeysResource) Remove(ctx context.Context, id string) (*DeleteResponse, error) {
	var out DeleteResponse
	if err := r.client.delete(ctx, fmt.Sprintf("/v1/api-keys/%s", url.PathEscape(id)), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

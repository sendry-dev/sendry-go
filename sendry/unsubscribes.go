package sendry

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

// UnsubscribesResource provides methods for managing the unsubscribe list.
type UnsubscribesResource struct {
	client *Client
}

// List returns unsubscribed email addresses with optional filters.
//
// Example:
//
//	page, err := client.Unsubscribes.List(ctx, &sendry.ListUnsubscribesParams{
//	    ListID: "newsletter",
//	})
func (r *UnsubscribesResource) List(ctx context.Context, params *ListUnsubscribesParams) (*PaginatedResponse[UnsubscribeEntry], error) {
	q := url.Values{}
	if params != nil {
		if params.Limit != nil {
			q.Set("limit", strconv.Itoa(*params.Limit))
		}
		if params.Cursor != nil {
			q.Set("cursor", *params.Cursor)
		}
		if params.Email != "" {
			q.Set("email", params.Email)
		}
		if params.ListID != "" {
			q.Set("list_id", params.ListID)
		}
	}
	var out PaginatedResponse[UnsubscribeEntry]
	if err := r.client.get(ctx, "/v1/unsubscribes", q, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Create adds a single email address to the unsubscribe list.
//
// Example:
//
//	entry, err := client.Unsubscribes.Create(ctx, sendry.CreateUnsubscribeParams{
//	    Email:  "user@example.com",
//	    ListID: "newsletter",
//	})
func (r *UnsubscribesResource) Create(ctx context.Context, params CreateUnsubscribeParams) (*UnsubscribeEntry, error) {
	var out UnsubscribeEntry
	if err := r.client.post(ctx, "/v1/unsubscribes", params, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// CreateBatch adds up to 1000 email addresses to the unsubscribe list in one request.
//
// Example:
//
//	result, err := client.Unsubscribes.CreateBatch(ctx, sendry.BatchUnsubscribeParams{
//	    Emails: []string{"a@example.com", "b@example.com"},
//	    ListID: "newsletter",
//	})
//	fmt.Println(result.Inserted)
func (r *UnsubscribesResource) CreateBatch(ctx context.Context, params BatchUnsubscribeParams) (*BatchUnsubscribeResponse, error) {
	var out BatchUnsubscribeResponse
	if err := r.client.post(ctx, "/v1/unsubscribes/batch", params, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Get retrieves a single unsubscribe record by its ID.
//
// Example:
//
//	entry, err := client.Unsubscribes.Get(ctx, "unsub_abc123")
func (r *UnsubscribesResource) Get(ctx context.Context, id string) (*UnsubscribeEntry, error) {
	var out UnsubscribeEntry
	if err := r.client.get(ctx, fmt.Sprintf("/v1/unsubscribes/%s", url.PathEscape(id)), nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Remove deletes an unsubscribe record.
//
// Example:
//
//	_, err := client.Unsubscribes.Remove(ctx, "unsub_abc123")
func (r *UnsubscribesResource) Remove(ctx context.Context, id string) (*DeleteResponse, error) {
	var out DeleteResponse
	if err := r.client.delete(ctx, fmt.Sprintf("/v1/unsubscribes/%s", url.PathEscape(id)), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

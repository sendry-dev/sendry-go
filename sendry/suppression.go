package sendry

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

// SuppressionResource provides methods for managing the email suppression list.
type SuppressionResource struct {
	client *Client
}

// List returns all suppressed email addresses.
//
// Example:
//
//	page, err := client.Suppression.List(ctx, nil)
func (r *SuppressionResource) List(ctx context.Context, params *PaginationParams) (*PaginatedResponse[SuppressionEntry], error) {
	q := url.Values{}
	if params != nil {
		if params.Limit != nil {
			q.Set("limit", strconv.Itoa(*params.Limit))
		}
		if params.Cursor != nil {
			q.Set("cursor", *params.Cursor)
		}
	}
	var out PaginatedResponse[SuppressionEntry]
	if err := r.client.get(ctx, "/v1/suppression", q, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Add adds an email address to the suppression list.
//
// Example:
//
//	entry, err := client.Suppression.Add(ctx, sendr.AddSuppressionParams{
//	    Email:  "bounced@example.com",
//	    Reason: sendr.SuppressionReasonHardBounce,
//	})
func (r *SuppressionResource) Add(ctx context.Context, params AddSuppressionParams) (*SuppressionEntry, error) {
	var out SuppressionEntry
	if err := r.client.post(ctx, "/v1/suppression", params, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Remove removes an email address from the suppression list.
//
// Example:
//
//	_, err := client.Suppression.Remove(ctx, "bounced@example.com")
func (r *SuppressionResource) Remove(ctx context.Context, email string) (*DeleteResponse, error) {
	var out DeleteResponse
	if err := r.client.delete(ctx, fmt.Sprintf("/v1/suppression/%s", url.PathEscape(email)), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

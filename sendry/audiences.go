package sendry

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

// AudiencesResource provides methods for managing audiences (contact lists).
type AudiencesResource struct {
	client *Client
}

// Create creates a new audience.
//
// Example:
//
//	audience, err := client.Audiences.Create(ctx, sendr.CreateAudienceParams{
//	    Name:        "Newsletter Subscribers",
//	    Description: "Weekly newsletter list",
//	})
func (r *AudiencesResource) Create(ctx context.Context, params CreateAudienceParams) (*Audience, error) {
	var out Audience
	if err := r.client.post(ctx, "/v1/audiences", params, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// List returns all audiences with cursor-based pagination.
//
// Example:
//
//	page, err := client.Audiences.List(ctx, nil)
func (r *AudiencesResource) List(ctx context.Context, params *PaginationParams) (*PaginatedResponse[Audience], error) {
	q := url.Values{}
	if params != nil {
		if params.Limit != nil {
			q.Set("limit", strconv.Itoa(*params.Limit))
		}
		if params.Cursor != nil {
			q.Set("cursor", *params.Cursor)
		}
	}
	var out PaginatedResponse[Audience]
	if err := r.client.get(ctx, "/v1/audiences", q, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Get retrieves an audience by its ID.
//
// Example:
//
//	audience, err := client.Audiences.Get(ctx, "aud_abc123")
//	fmt.Println(audience.MemberCount)
func (r *AudiencesResource) Get(ctx context.Context, id string) (*Audience, error) {
	var out Audience
	if err := r.client.get(ctx, fmt.Sprintf("/v1/audiences/%s", url.PathEscape(id)), nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Update updates an audience's name or description.
//
// Example:
//
//	updated, err := client.Audiences.Update(ctx, "aud_abc123", sendr.UpdateAudienceParams{
//	    Name: "VIP Subscribers",
//	})
func (r *AudiencesResource) Update(ctx context.Context, id string, params UpdateAudienceParams) (*Audience, error) {
	var out Audience
	if err := r.client.put(ctx, fmt.Sprintf("/v1/audiences/%s", url.PathEscape(id)), params, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Remove deletes an audience. Contacts themselves are not deleted.
//
// Example:
//
//	_, err := client.Audiences.Remove(ctx, "aud_abc123")
func (r *AudiencesResource) Remove(ctx context.Context, id string) (*DeleteResponse, error) {
	var out DeleteResponse
	if err := r.client.delete(ctx, fmt.Sprintf("/v1/audiences/%s", url.PathEscape(id)), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// AddContacts adds one or more contacts to an audience by their IDs (max 100 per call).
//
// Example:
//
//	result, err := client.Audiences.AddContacts(ctx, "aud_abc123", sendr.AddContactsToAudienceParams{
//	    ContactIDs: []string{"ct_1", "ct_2"},
//	})
//	fmt.Println(result.Added)
func (r *AudiencesResource) AddContacts(ctx context.Context, audienceID string, params AddContactsToAudienceParams) (*AddContactsToAudienceResult, error) {
	var out AddContactsToAudienceResult
	if err := r.client.post(ctx, fmt.Sprintf("/v1/audiences/%s/contacts", url.PathEscape(audienceID)), params, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// ListContacts returns all contacts in an audience.
//
// Example:
//
//	page, err := client.Audiences.ListContacts(ctx, "aud_abc123", nil)
func (r *AudiencesResource) ListContacts(ctx context.Context, audienceID string, params *PaginationParams) (*PaginatedResponse[Contact], error) {
	q := url.Values{}
	if params != nil {
		if params.Limit != nil {
			q.Set("limit", strconv.Itoa(*params.Limit))
		}
		if params.Cursor != nil {
			q.Set("cursor", *params.Cursor)
		}
	}
	var out PaginatedResponse[Contact]
	if err := r.client.get(ctx, fmt.Sprintf("/v1/audiences/%s/contacts", url.PathEscape(audienceID)), q, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// RemoveContact removes a contact from an audience.
//
// Example:
//
//	_, err := client.Audiences.RemoveContact(ctx, "aud_abc123", "ct_xyz456")
func (r *AudiencesResource) RemoveContact(ctx context.Context, audienceID, contactID string) (*DeleteResponse, error) {
	var out DeleteResponse
	path := fmt.Sprintf("/v1/audiences/%s/contacts/%s", url.PathEscape(audienceID), url.PathEscape(contactID))
	if err := r.client.delete(ctx, path, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

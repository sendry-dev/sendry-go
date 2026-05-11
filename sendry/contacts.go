package sendry

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

// ContactsResource provides methods for managing contacts.
type ContactsResource struct {
	client *Client
}

// Create creates a new contact.
//
// Example:
//
//	contact, err := client.Contacts.Create(ctx, sendr.CreateContactParams{
//	    Email:     "jane@example.com",
//	    FirstName: "Jane",
//	    LastName:  "Doe",
//	})
func (r *ContactsResource) Create(ctx context.Context, params CreateContactParams) (*Contact, error) {
	var out Contact
	if err := r.client.post(ctx, "/v1/contacts", params, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// List returns a cursor-paginated list of contacts with optional filters.
//
// Example:
//
//	page, err := client.Contacts.List(ctx, &sendr.ListContactsParams{
//	    AudienceID: "aud_abc123",
//	})
func (r *ContactsResource) List(ctx context.Context, params *ListContactsParams) (*PaginatedResponse[Contact], error) {
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
		if params.AudienceID != "" {
			q.Set("audience_id", params.AudienceID)
		}
	}
	var out PaginatedResponse[Contact]
	if err := r.client.get(ctx, "/v1/contacts", q, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Get retrieves a contact by its ID.
//
// Example:
//
//	contact, err := client.Contacts.Get(ctx, "ct_abc123")
func (r *ContactsResource) Get(ctx context.Context, id string) (*Contact, error) {
	var out Contact
	if err := r.client.get(ctx, fmt.Sprintf("/v1/contacts/%s", url.PathEscape(id)), nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Update updates a contact.
//
// Example:
//
//	updated, err := client.Contacts.Update(ctx, "ct_abc123", sendr.UpdateContactParams{
//	    Unsubscribed: sendr.BoolPtr(true),
//	})
func (r *ContactsResource) Update(ctx context.Context, id string, params UpdateContactParams) (*Contact, error) {
	var out Contact
	if err := r.client.put(ctx, fmt.Sprintf("/v1/contacts/%s", url.PathEscape(id)), params, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Remove deletes a contact.
//
// Example:
//
//	_, err := client.Contacts.Remove(ctx, "ct_abc123")
func (r *ContactsResource) Remove(ctx context.Context, id string) (*DeleteResponse, error) {
	var out DeleteResponse
	if err := r.client.delete(ctx, fmt.Sprintf("/v1/contacts/%s", url.PathEscape(id)), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// BulkImport imports up to 1000 contacts at once.
// Existing contacts matched by email are updated; new ones are created.
//
// Example:
//
//	result, err := client.Contacts.BulkImport(ctx, sendr.BulkImportContactsParams{
//	    Contacts: []sendr.BulkImportContactItem{
//	        {Email: "alice@example.com", FirstName: "Alice"},
//	        {Email: "bob@example.com", FirstName: "Bob"},
//	    },
//	    AudienceID: "aud_abc123",
//	})
//	fmt.Println(result.Created, result.Updated)
func (r *ContactsResource) BulkImport(ctx context.Context, params BulkImportContactsParams) (*BulkImportResult, error) {
	var out BulkImportResult
	if err := r.client.post(ctx, "/v1/contacts/import", params, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

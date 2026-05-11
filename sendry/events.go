package sendry

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

// EventsResource provides methods for ingesting and querying automation
// trigger events.
type EventsResource struct {
	client *Client
}

// Ingest ingests an event into Sendry. Events may trigger automations.
//
// Example:
//
//	event, err := client.Events.Ingest(ctx, sendry.IngestEventParams{
//	    Name:         "signup",
//	    ContactEmail: "jane@example.com",
//	    Payload:      map[string]any{"plan": "pro"},
//	})
func (r *EventsResource) Ingest(ctx context.Context, params IngestEventParams) (*IngestedEvent, error) {
	var out IngestedEvent
	if err := r.client.post(ctx, "/v1/events", params, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// List returns ingested events with cursor-based pagination.
//
// Example:
//
//	page, err := client.Events.List(ctx, &sendry.ListEventsParams{Name: "signup"})
func (r *EventsResource) List(ctx context.Context, params *ListEventsParams) (*PaginatedResponse[IngestedEvent], error) {
	q := url.Values{}
	if params != nil {
		if params.Limit != nil {
			q.Set("limit", strconv.Itoa(*params.Limit))
		}
		if params.Cursor != nil {
			q.Set("cursor", *params.Cursor)
		}
		if params.Name != "" {
			q.Set("name", params.Name)
		}
	}
	var out PaginatedResponse[IngestedEvent]
	if err := r.client.get(ctx, "/v1/events", q, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Get retrieves a single ingested event by ID.
//
// Example:
//
//	event, err := client.Events.Get(ctx, "evt_abc123")
func (r *EventsResource) Get(ctx context.Context, id string) (*IngestedEvent, error) {
	var out IngestedEvent
	if err := r.client.get(ctx, fmt.Sprintf("/v1/events/%s", url.PathEscape(id)), nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

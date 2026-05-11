package sendry

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

// CampaignsResource provides methods for creating and managing email campaigns.
type CampaignsResource struct {
	client *Client
}

// Create creates a new campaign in draft status.
//
// Example:
//
//	campaign, err := client.Campaigns.Create(ctx, sendry.CreateCampaignParams{
//	    Name:       "March Newsletter",
//	    Subject:    "What's new in March",
//	    From:       "Acme <hello@acme.com>",
//	    AudienceID: "aud_abc123",
//	    HTML:       "<h1>Hello!</h1>",
//	})
func (r *CampaignsResource) Create(ctx context.Context, params CreateCampaignParams) (*Campaign, error) {
	var out Campaign
	if err := r.client.post(ctx, "/v1/campaigns", params, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// List returns campaigns with cursor-based pagination and optional status filter.
//
// Example:
//
//	page, err := client.Campaigns.List(ctx, &sendry.ListCampaignsParams{Status: "draft"})
func (r *CampaignsResource) List(ctx context.Context, params *ListCampaignsParams) (*PaginatedResponse[CampaignListItem], error) {
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
	var out PaginatedResponse[CampaignListItem]
	if err := r.client.get(ctx, "/v1/campaigns", q, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Get retrieves a campaign by its ID, including delivery statistics.
//
// Example:
//
//	campaign, err := client.Campaigns.Get(ctx, "cp_abc123")
//	fmt.Println(campaign.Stats.DeliveredCount)
func (r *CampaignsResource) Get(ctx context.Context, id string) (*Campaign, error) {
	var out Campaign
	if err := r.client.get(ctx, fmt.Sprintf("/v1/campaigns/%s", url.PathEscape(id)), nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Update modifies a campaign. Only draft campaigns can be updated.
//
// Example:
//
//	updated, err := client.Campaigns.Update(ctx, "cp_abc123", sendry.UpdateCampaignParams{
//	    Subject: "Updated subject line",
//	})
func (r *CampaignsResource) Update(ctx context.Context, id string, params UpdateCampaignParams) (*Campaign, error) {
	var out Campaign
	if err := r.client.put(ctx, fmt.Sprintf("/v1/campaigns/%s", url.PathEscape(id)), params, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Remove deletes a campaign. Only draft or cancelled campaigns can be deleted.
//
// Example:
//
//	_, err := client.Campaigns.Remove(ctx, "cp_abc123")
func (r *CampaignsResource) Remove(ctx context.Context, id string) (*DeleteResponse, error) {
	var out DeleteResponse
	if err := r.client.delete(ctx, fmt.Sprintf("/v1/campaigns/%s", url.PathEscape(id)), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Schedule schedules a draft campaign to be sent at a specific time.
//
// Example:
//
//	result, err := client.Campaigns.Schedule(ctx, "cp_abc123", sendry.ScheduleCampaignParams{
//	    ScheduledAt: "2026-03-15T10:00:00Z",
//	})
func (r *CampaignsResource) Schedule(ctx context.Context, id string, params ScheduleCampaignParams) (*CampaignActionResponse, error) {
	var out CampaignActionResponse
	if err := r.client.post(ctx, fmt.Sprintf("/v1/campaigns/%s/schedule", url.PathEscape(id)), params, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Send immediately enqueues a draft or scheduled campaign for sending.
//
// Example:
//
//	result, err := client.Campaigns.Send(ctx, "cp_abc123")
func (r *CampaignsResource) Send(ctx context.Context, id string) (*CampaignActionResponse, error) {
	var out CampaignActionResponse
	if err := r.client.post(ctx, fmt.Sprintf("/v1/campaigns/%s/send", url.PathEscape(id)), nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Cancel cancels a scheduled or paused campaign.
//
// Example:
//
//	result, err := client.Campaigns.Cancel(ctx, "cp_abc123")
func (r *CampaignsResource) Cancel(ctx context.Context, id string) (*CampaignActionResponse, error) {
	var out CampaignActionResponse
	if err := r.client.post(ctx, fmt.Sprintf("/v1/campaigns/%s/cancel", url.PathEscape(id)), nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Pause pauses a campaign that is currently sending.
//
// Example:
//
//	result, err := client.Campaigns.Pause(ctx, "cp_abc123")
func (r *CampaignsResource) Pause(ctx context.Context, id string) (*CampaignActionResponse, error) {
	var out CampaignActionResponse
	if err := r.client.post(ctx, fmt.Sprintf("/v1/campaigns/%s/pause", url.PathEscape(id)), nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Resume resumes a paused campaign.
//
// Example:
//
//	result, err := client.Campaigns.Resume(ctx, "cp_abc123")
func (r *CampaignsResource) Resume(ctx context.Context, id string) (*CampaignActionResponse, error) {
	var out CampaignActionResponse
	if err := r.client.post(ctx, fmt.Sprintf("/v1/campaigns/%s/resume", url.PathEscape(id)), nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

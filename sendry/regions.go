package sendry

import (
	"context"
	"fmt"
	"net/url"
)

// RegionsResource provides methods for managing SES sending regions.
type RegionsResource struct {
	client *Client
}

// List returns all active SES regions available for sending.
//
// Example:
//
//	regions, err := client.Regions.List(ctx)
func (r *RegionsResource) List(ctx context.Context) ([]Region, error) {
	var out struct {
		Data []Region `json:"data"`
	}
	if err := r.client.get(ctx, "/v1/regions", nil, &out); err != nil {
		return nil, err
	}
	return out.Data, nil
}

// GetOrgSettings returns the organisation-level default SES sending region and
// data residency constraint.
//
// Example:
//
//	settings, err := client.Regions.GetOrgSettings(ctx)
func (r *RegionsResource) GetOrgSettings(ctx context.Context) (*OrgRegionSettings, error) {
	var out OrgRegionSettings
	if err := r.client.get(ctx, "/v1/organizations/me/region", nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// UpdateOrgSettings sets the organisation-level default SES sending region and
// optional data residency constraint.
//
// Example:
//
//	region := "eu-west-1"
//	settings, err := client.Regions.UpdateOrgSettings(ctx, sendry.UpdateOrgRegionParams{
//	    DefaultRegion: &region, DataResidency: "eu",
//	})
func (r *RegionsResource) UpdateOrgSettings(ctx context.Context, params UpdateOrgRegionParams) (*OrgRegionSettings, error) {
	var out OrgRegionSettings
	if err := r.client.patch(ctx, "/v1/organizations/me/region", params, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// SetDomainRegion overrides the SES sending region for a specific domain.
// Pass a nil Region to clear the override.
//
// Example:
//
//	region := "eu-west-1"
//	_, err := client.Regions.SetDomainRegion(ctx, "dom_abc123", sendry.UpdateDomainRegionParams{Region: &region})
func (r *RegionsResource) SetDomainRegion(ctx context.Context, domainID string, params UpdateDomainRegionParams) (map[string]any, error) {
	var out map[string]any
	if err := r.client.patch(ctx, fmt.Sprintf("/v1/domains/%s/region", url.PathEscape(domainID)), params, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// GetRegionAnalytics returns a breakdown of sent emails by SES region.
//
// Example:
//
//	analytics, err := client.Regions.GetRegionAnalytics(ctx, sendry.RegionAnalyticsParams{
//	    From: "2025-01-01", To: "2025-01-31",
//	})
func (r *RegionsResource) GetRegionAnalytics(ctx context.Context, params RegionAnalyticsParams) (*RegionAnalyticsResponse, error) {
	q := url.Values{}
	if params.From != "" {
		q.Set("from", params.From)
	}
	if params.To != "" {
		q.Set("to", params.To)
	}
	var out RegionAnalyticsResponse
	if err := r.client.get(ctx, "/v1/analytics/regions", q, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

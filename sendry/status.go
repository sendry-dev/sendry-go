package sendry

import (
	"context"
	"net/url"
	"strconv"
)

// StatusResource provides methods for querying Sendry's operational status.
type StatusResource struct {
	client *Client
}

// GetCurrent returns the current operational status of all Sendry components
// and any active incidents.
//
// Example:
//
//	status, err := client.Status.GetCurrent(ctx)
func (r *StatusResource) GetCurrent(ctx context.Context) (*SystemStatus, error) {
	var out SystemStatus
	if err := r.client.get(ctx, "/v1/status", nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetHistory returns a paginated list of resolved incidents.
//
// Example:
//
//	page, err := client.Status.GetHistory(ctx, nil)
func (r *StatusResource) GetHistory(ctx context.Context, params *PaginationParams) (*PaginatedResponse[Incident], error) {
	q := url.Values{}
	if params != nil {
		if params.Limit != nil {
			q.Set("limit", strconv.Itoa(*params.Limit))
		}
		if params.Cursor != nil {
			q.Set("cursor", *params.Cursor)
		}
	}
	var out PaginatedResponse[Incident]
	if err := r.client.get(ctx, "/v1/status/history", q, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetLatency returns hourly P50/P95/P99 latency rollups for a component.
//
// Example:
//
//	latency, err := client.Status.GetLatency(ctx, &sendry.GetLatencyParams{
//	    Component: "api-gateway", Hours: sendry.IntPtr(48),
//	})
func (r *StatusResource) GetLatency(ctx context.Context, params *GetLatencyParams) (*LatencyStats, error) {
	q := url.Values{}
	if params != nil {
		if params.Component != "" {
			q.Set("component", params.Component)
		}
		if params.Hours != nil {
			q.Set("hours", strconv.Itoa(*params.Hours))
		}
	}
	var out LatencyStats
	if err := r.client.get(ctx, "/v1/status/latency", q, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

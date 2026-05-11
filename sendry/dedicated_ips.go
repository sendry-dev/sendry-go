package sendry

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

// DedicatedIpsResource provides methods for managing dedicated sending IPs.
type DedicatedIpsResource struct {
	client *Client
}

// Provision provisions a new dedicated IP address for email sending.
// Available on Business and Enterprise plans.
//
// Example:
//
//	ip, err := client.DedicatedIps.Provision(ctx, &sendry.ProvisionDedicatedIpParams{
//	    Provider: "ses",
//	})
func (r *DedicatedIpsResource) Provision(ctx context.Context, params *ProvisionDedicatedIpParams) (*DedicatedIp, error) {
	body := params
	if body == nil {
		body = &ProvisionDedicatedIpParams{}
	}
	var out DedicatedIp
	if err := r.client.post(ctx, "/v1/ips", body, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// List returns all dedicated IPs for the organisation.
//
// Example:
//
//	page, err := client.DedicatedIps.List(ctx, nil)
func (r *DedicatedIpsResource) List(ctx context.Context, params *PaginationParams) (*PaginatedResponse[DedicatedIp], error) {
	q := url.Values{}
	if params != nil {
		if params.Limit != nil {
			q.Set("limit", strconv.Itoa(*params.Limit))
		}
		if params.Cursor != nil {
			q.Set("cursor", *params.Cursor)
		}
	}
	var out PaginatedResponse[DedicatedIp]
	if err := r.client.get(ctx, "/v1/ips", q, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Get retrieves details for a specific dedicated IP.
//
// Example:
//
//	ip, err := client.DedicatedIps.Get(ctx, "dip_abc123")
func (r *DedicatedIpsResource) Get(ctx context.Context, id string) (*DedicatedIp, error) {
	var out DedicatedIp
	if err := r.client.get(ctx, fmt.Sprintf("/v1/ips/%s", url.PathEscape(id)), nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Assign assigns a dedicated IP to a domain. The IP must be in warming or active status.
//
// Example:
//
//	assignment, err := client.DedicatedIps.Assign(ctx, "dip_abc123", sendry.AssignIpParams{
//	    DomainID: "dom_xyz456",
//	})
func (r *DedicatedIpsResource) Assign(ctx context.Context, ipID string, params AssignIpParams) (*IpAssignment, error) {
	var out IpAssignment
	if err := r.client.post(ctx, fmt.Sprintf("/v1/ips/%s/assign", url.PathEscape(ipID)), params, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// RemoveAssignment removes a dedicated IP assignment from a domain.
//
// Example:
//
//	_, err := client.DedicatedIps.RemoveAssignment(ctx, "dip_abc123", "asgn_xyz456")
func (r *DedicatedIpsResource) RemoveAssignment(ctx context.Context, ipID, assignmentID string) (*DeleteResponse, error) {
	var out DeleteResponse
	if err := r.client.delete(ctx, fmt.Sprintf("/v1/ips/%s/assign/%s", url.PathEscape(ipID), url.PathEscape(assignmentID)), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Release releases a dedicated IP address. This deletes the associated SES pool and all domain assignments.
//
// Example:
//
//	_, err := client.DedicatedIps.Release(ctx, "dip_abc123")
func (r *DedicatedIpsResource) Release(ctx context.Context, id string) (*DeleteResponse, error) {
	var out DeleteResponse
	if err := r.client.delete(ctx, fmt.Sprintf("/v1/ips/%s", url.PathEscape(id)), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

package sendry

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

// DomainsResource provides methods for managing sending domains.
type DomainsResource struct {
	client *Client
}

// Create adds a new domain to the organisation.
//
// Example:
//
//	domain, err := client.Domains.Create(ctx, sendry.CreateDomainParams{Name: "example.com"})
//	// domain.DnsRecords contains records to add to your DNS provider
func (r *DomainsResource) Create(ctx context.Context, params CreateDomainParams) (*Domain, error) {
	var out Domain
	if err := r.client.post(ctx, "/v1/domains", params, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// List returns all domains registered in the organisation.
//
// Example:
//
//	page, err := client.Domains.List(ctx, nil)
func (r *DomainsResource) List(ctx context.Context, params *PaginationParams) (*PaginatedResponse[Domain], error) {
	q := url.Values{}
	if params != nil {
		if params.Limit != nil {
			q.Set("limit", strconv.Itoa(*params.Limit))
		}
		if params.Cursor != nil {
			q.Set("cursor", *params.Cursor)
		}
	}
	var out PaginatedResponse[Domain]
	if err := r.client.get(ctx, "/v1/domains", q, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Get retrieves a domain by its ID.
//
// Example:
//
//	domain, err := client.Domains.Get(ctx, "dom_abc123")
func (r *DomainsResource) Get(ctx context.Context, id string) (*Domain, error) {
	var out Domain
	if err := r.client.get(ctx, fmt.Sprintf("/v1/domains/%s", url.PathEscape(id)), nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Verify triggers DNS verification for a domain.
//
// Example:
//
//	result, err := client.Domains.Verify(ctx, "dom_abc123")
//	if result.SpfVerified && result.DkimVerified {
//	    fmt.Println("Domain is fully verified!")
//	}
func (r *DomainsResource) Verify(ctx context.Context, id string) (*VerifyDomainResponse, error) {
	var out VerifyDomainResponse
	if err := r.client.post(ctx, fmt.Sprintf("/v1/domains/%s/verify", url.PathEscape(id)), nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Remove deletes a domain.
//
// Example:
//
//	result, err := client.Domains.Remove(ctx, "dom_abc123")
func (r *DomainsResource) Remove(ctx context.Context, id string) (*DeleteResponse, error) {
	var out DeleteResponse
	if err := r.client.delete(ctx, fmt.Sprintf("/v1/domains/%s", url.PathEscape(id)), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// ConfigureBimi configures BIMI (Brand Indicators for Message Identification) for a domain.
//
// Example:
//
//	bimi, err := client.Domains.ConfigureBimi(ctx, "dom_abc123", sendry.ConfigureBimiParams{
//	    LogoURL: "https://example.com/logo.svg",
//	})
func (r *DomainsResource) ConfigureBimi(ctx context.Context, domainID string, params ConfigureBimiParams) (*BimiConfig, error) {
	var out BimiConfig
	if err := r.client.post(ctx, fmt.Sprintf("/v1/domains/%s/bimi", url.PathEscape(domainID)), params, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetBimi retrieves the BIMI configuration for a domain.
//
// Example:
//
//	bimi, err := client.Domains.GetBimi(ctx, "dom_abc123")
func (r *DomainsResource) GetBimi(ctx context.Context, domainID string) (*BimiConfig, error) {
	var out BimiConfig
	if err := r.client.get(ctx, fmt.Sprintf("/v1/domains/%s/bimi", url.PathEscape(domainID)), nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// VerifyBimi triggers DNS verification for a domain's BIMI record.
//
// Example:
//
//	result, err := client.Domains.VerifyBimi(ctx, "dom_abc123")
func (r *DomainsResource) VerifyBimi(ctx context.Context, domainID string) (*VerifyBimiResponse, error) {
	var out VerifyBimiResponse
	if err := r.client.post(ctx, fmt.Sprintf("/v1/domains/%s/bimi/verify", url.PathEscape(domainID)), nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// RemoveBimi deletes the BIMI configuration for a domain.
//
// Example:
//
//	_, err := client.Domains.RemoveBimi(ctx, "dom_abc123")
func (r *DomainsResource) RemoveBimi(ctx context.Context, domainID string) (*DeleteResponse, error) {
	var out DeleteResponse
	if err := r.client.delete(ctx, fmt.Sprintf("/v1/domains/%s/bimi", url.PathEscape(domainID)), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

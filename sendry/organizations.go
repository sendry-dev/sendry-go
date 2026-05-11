package sendry

import "context"

// OrganizationsResource provides methods for managing the current organisation.
type OrganizationsResource struct {
	client *Client
}

// GetCurrent returns the current organisation's details.
//
// Example:
//
//	org, err := client.Organizations.GetCurrent(ctx)
func (r *OrganizationsResource) GetCurrent(ctx context.Context) (*Organization, error) {
	var out Organization
	if err := r.client.get(ctx, "/v1/organizations/me", nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Update updates the current organisation's name.
//
// Example:
//
//	org, err := client.Organizations.Update(ctx, sendry.UpdateOrganizationParams{Name: "Acme Corp"})
func (r *OrganizationsResource) Update(ctx context.Context, params UpdateOrganizationParams) (*Organization, error) {
	var out Organization
	if err := r.client.patch(ctx, "/v1/organizations/me", params, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetBranding returns the current organisation's branding settings.
//
// Example:
//
//	branding, err := client.Organizations.GetBranding(ctx)
func (r *OrganizationsResource) GetBranding(ctx context.Context) (*BrandingSettings, error) {
	var out BrandingSettings
	if err := r.client.get(ctx, "/v1/organizations/me/branding", nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// UpdateBranding updates the current organisation's branding settings.
//
// Example:
//
//	color := "#6366f1"
//	branding, err := client.Organizations.UpdateBranding(ctx, sendry.UpdateBrandingParams{
//	    BrandColor: &color,
//	})
func (r *OrganizationsResource) UpdateBranding(ctx context.Context, params UpdateBrandingParams) (*BrandingSettings, error) {
	var out BrandingSettings
	if err := r.client.patch(ctx, "/v1/organizations/me/branding", params, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

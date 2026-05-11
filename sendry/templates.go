package sendry

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

// TemplatesResource provides methods for managing and rendering email templates.
type TemplatesResource struct {
	client *Client
}

// Create creates a new email template.
//
// Example:
//
//	tmpl, err := client.Templates.Create(ctx, sendry.CreateTemplateParams{
//	    Name:    "Welcome Email",
//	    Subject: "Welcome, {{name}}!",
//	    HTML:    "<h1>Hello {{name}}</h1>",
//	})
func (r *TemplatesResource) Create(ctx context.Context, params CreateTemplateParams) (*Template, error) {
	var out Template
	if err := r.client.post(ctx, "/v1/templates", params, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// List returns all templates with cursor-based pagination.
//
// Example:
//
//	page, err := client.Templates.List(ctx, nil)
func (r *TemplatesResource) List(ctx context.Context, params *PaginationParams) (*PaginatedResponse[Template], error) {
	q := url.Values{}
	if params != nil {
		if params.Limit != nil {
			q.Set("limit", strconv.Itoa(*params.Limit))
		}
		if params.Cursor != nil {
			q.Set("cursor", *params.Cursor)
		}
	}
	var out PaginatedResponse[Template]
	if err := r.client.get(ctx, "/v1/templates", q, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Get retrieves a template by its ID.
//
// Example:
//
//	tmpl, err := client.Templates.Get(ctx, "tmpl_abc123")
func (r *TemplatesResource) Get(ctx context.Context, id string) (*Template, error) {
	var out Template
	if err := r.client.get(ctx, fmt.Sprintf("/v1/templates/%s", url.PathEscape(id)), nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Update updates a template.
//
// Example:
//
//	updated, err := client.Templates.Update(ctx, "tmpl_abc123", sendry.UpdateTemplateParams{
//	    Subject: "Updated Subject",
//	})
func (r *TemplatesResource) Update(ctx context.Context, id string, params UpdateTemplateParams) (*Template, error) {
	var out Template
	if err := r.client.put(ctx, fmt.Sprintf("/v1/templates/%s", url.PathEscape(id)), params, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Remove deletes a template.
//
// Example:
//
//	_, err := client.Templates.Remove(ctx, "tmpl_abc123")
func (r *TemplatesResource) Remove(ctx context.Context, id string) (*DeleteResponse, error) {
	var out DeleteResponse
	if err := r.client.delete(ctx, fmt.Sprintf("/v1/templates/%s", url.PathEscape(id)), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Render renders a saved template with the provided variable values.
//
// Example:
//
//	result, err := client.Templates.Render(ctx, "tmpl_abc123", &sendry.RenderTemplateParams{
//	    Variables: map[string]string{"name": "World"},
//	})
func (r *TemplatesResource) Render(ctx context.Context, id string, params *RenderTemplateParams) (*RenderTemplateResponse, error) {
	body := params
	if body == nil {
		body = &RenderTemplateParams{}
	}
	var out RenderTemplateResponse
	if err := r.client.post(ctx, fmt.Sprintf("/v1/templates/%s/render", url.PathEscape(id)), body, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// ListStarters returns all pre-built starter templates.
//
// Example:
//
//	starters, err := client.Templates.ListStarters(ctx)
func (r *TemplatesResource) ListStarters(ctx context.Context) ([]TemplateStarter, error) {
	var out struct {
		Data []TemplateStarter `json:"data"`
	}
	if err := r.client.get(ctx, "/v1/templates/starters", nil, &out); err != nil {
		return nil, err
	}
	return out.Data, nil
}

// ListVisualStarters returns summaries of all available visual (block-based) starter templates.
//
// Example:
//
//	starters, err := client.Templates.ListVisualStarters(ctx)
func (r *TemplatesResource) ListVisualStarters(ctx context.Context) ([]VisualStarterSummary, error) {
	var out struct {
		Data []VisualStarterSummary `json:"data"`
	}
	if err := r.client.get(ctx, "/v1/templates/visual-starters", nil, &out); err != nil {
		return nil, err
	}
	return out.Data, nil
}

// GetVisualStarter retrieves the full design JSON for a visual starter template.
//
// Example:
//
//	design, err := client.Templates.GetVisualStarter(ctx, "welcome-blocks")
func (r *TemplatesResource) GetVisualStarter(ctx context.Context, starterID string) (map[string]any, error) {
	var out map[string]any
	if err := r.client.get(ctx, fmt.Sprintf("/v1/templates/visual-starters/%s", url.PathEscape(starterID)), nil, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// CompileBlocks compiles a visual block design JSON to email-safe HTML.
//
// Example:
//
//	result, err := client.Templates.CompileBlocks(ctx, sendry.CompileBlocksParams{
//	    Design:    myBlockDesign,
//	    Variables: map[string]string{"name": "Alice"},
//	})
func (r *TemplatesResource) CompileBlocks(ctx context.Context, params CompileBlocksParams) (*RenderTemplateResponse, error) {
	var out RenderTemplateResponse
	if err := r.client.post(ctx, "/v1/templates/compile-blocks", params, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// RenderAdhoc renders arbitrary HTML with variable substitution without saving a template.
//
// Example:
//
//	result, err := client.Templates.RenderAdhoc(ctx, sendry.RenderAdhocParams{
//	    HTML:      "<h1>Hello {{name}}</h1>",
//	    Variables: map[string]string{"name": "Bob"},
//	})
func (r *TemplatesResource) RenderAdhoc(ctx context.Context, params RenderAdhocParams) (*RenderTemplateResponse, error) {
	var out RenderTemplateResponse
	if err := r.client.post(ctx, "/v1/templates/render", params, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

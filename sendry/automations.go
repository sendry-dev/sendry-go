package sendry

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

// AutomationsResource provides methods for managing automation workflows,
// their steps, and execution runs.
type AutomationsResource struct {
	client *Client
}

// List returns automations with cursor-based pagination and optional status filter.
//
// Example:
//
//	page, err := client.Automations.List(ctx, &sendry.ListAutomationsParams{
//	    Status: sendry.AutomationStatusActive,
//	})
func (r *AutomationsResource) List(ctx context.Context, params *ListAutomationsParams) (*PaginatedResponse[Automation], error) {
	q := url.Values{}
	if params != nil {
		if params.Limit != nil {
			q.Set("limit", strconv.Itoa(*params.Limit))
		}
		if params.Cursor != nil {
			q.Set("cursor", *params.Cursor)
		}
		if params.Status != "" {
			q.Set("status", string(params.Status))
		}
	}
	var out PaginatedResponse[Automation]
	if err := r.client.get(ctx, "/v1/automations", q, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Get retrieves a single automation by ID.
//
// Example:
//
//	automation, err := client.Automations.Get(ctx, "auto_abc123")
func (r *AutomationsResource) Get(ctx context.Context, id string) (*Automation, error) {
	var out Automation
	if err := r.client.get(ctx, fmt.Sprintf("/v1/automations/%s", url.PathEscape(id)), nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Create creates a new automation in draft status.
//
// Example:
//
//	automation, err := client.Automations.Create(ctx, sendry.CreateAutomationParams{
//	    Name:        "Welcome series",
//	    TriggerType: sendry.AutomationTriggerEvent,
//	})
func (r *AutomationsResource) Create(ctx context.Context, params CreateAutomationParams) (*Automation, error) {
	var out Automation
	if err := r.client.post(ctx, "/v1/automations", params, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Update patches an automation. Only provided fields are changed.
//
// Example:
//
//	updated, err := client.Automations.Update(ctx, "auto_abc123", sendry.UpdateAutomationParams{
//	    Name: "Renamed automation",
//	})
func (r *AutomationsResource) Update(ctx context.Context, id string, params UpdateAutomationParams) (*Automation, error) {
	var out Automation
	if err := r.client.patch(ctx, fmt.Sprintf("/v1/automations/%s", url.PathEscape(id)), params, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Delete deletes an automation.
//
// Example:
//
//	_, err := client.Automations.Delete(ctx, "auto_abc123")
func (r *AutomationsResource) Delete(ctx context.Context, id string) (*DeleteResponse, error) {
	var out DeleteResponse
	if err := r.client.delete(ctx, fmt.Sprintf("/v1/automations/%s", url.PathEscape(id)), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Activate activates a draft or paused automation.
//
// Example:
//
//	automation, err := client.Automations.Activate(ctx, "auto_abc123")
func (r *AutomationsResource) Activate(ctx context.Context, id string) (*Automation, error) {
	var out Automation
	if err := r.client.post(ctx, fmt.Sprintf("/v1/automations/%s/activate", url.PathEscape(id)), nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Pause pauses an active automation.
//
// Example:
//
//	automation, err := client.Automations.Pause(ctx, "auto_abc123")
func (r *AutomationsResource) Pause(ctx context.Context, id string) (*Automation, error) {
	var out Automation
	if err := r.client.post(ctx, fmt.Sprintf("/v1/automations/%s/pause", url.PathEscape(id)), nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Archive archives an automation.
//
// Example:
//
//	automation, err := client.Automations.Archive(ctx, "auto_abc123")
func (r *AutomationsResource) Archive(ctx context.Context, id string) (*Automation, error) {
	var out Automation
	if err := r.client.post(ctx, fmt.Sprintf("/v1/automations/%s/archive", url.PathEscape(id)), nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// ListSteps returns all steps belonging to an automation.
//
// Example:
//
//	steps, err := client.Automations.ListSteps(ctx, "auto_abc123")
func (r *AutomationsResource) ListSteps(ctx context.Context, automationID string) ([]AutomationStep, error) {
	var out struct {
		Data []AutomationStep `json:"data"`
	}
	if err := r.client.get(ctx, fmt.Sprintf("/v1/automations/%s/steps", url.PathEscape(automationID)), nil, &out); err != nil {
		return nil, err
	}
	return out.Data, nil
}

// AddStep adds a new step to an automation. Use one of the typed *StepConfig
// constructors (SendEmailStepConfig, WaitStepConfig, BranchStepConfig,
// ABSplitStepConfig) and call .ToConfig() to build the Config field.
//
// Example:
//
//	step, err := client.Automations.AddStep(ctx, "auto_abc123", sendry.AddAutomationStepParams{
//	    Config: sendry.WaitStepConfig{DurationSeconds: 86400}.ToConfig(),
//	})
func (r *AutomationsResource) AddStep(ctx context.Context, automationID string, params AddAutomationStepParams) (*AutomationStep, error) {
	var out AutomationStep
	if err := r.client.post(ctx, fmt.Sprintf("/v1/automations/%s/steps", url.PathEscape(automationID)), params, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// UpdateStep patches an existing automation step.
//
// Example:
//
//	step, err := client.Automations.UpdateStep(ctx, "auto_abc123", "step_xyz", sendry.UpdateAutomationStepParams{
//	    Position: sendry.IntPtr(2),
//	})
func (r *AutomationsResource) UpdateStep(ctx context.Context, automationID, stepID string, params UpdateAutomationStepParams) (*AutomationStep, error) {
	var out AutomationStep
	if err := r.client.patch(ctx, fmt.Sprintf("/v1/automations/%s/steps/%s", url.PathEscape(automationID), url.PathEscape(stepID)), params, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// DeleteStep deletes a step from an automation.
//
// Example:
//
//	_, err := client.Automations.DeleteStep(ctx, "auto_abc123", "step_xyz")
func (r *AutomationsResource) DeleteStep(ctx context.Context, automationID, stepID string) (*DeleteResponse, error) {
	var out DeleteResponse
	if err := r.client.delete(ctx, fmt.Sprintf("/v1/automations/%s/steps/%s", url.PathEscape(automationID), url.PathEscape(stepID)), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// ListRuns returns automation runs with cursor-based pagination and optional status filter.
//
// Example:
//
//	page, err := client.Automations.ListRuns(ctx, "auto_abc123", nil)
func (r *AutomationsResource) ListRuns(ctx context.Context, automationID string, params *ListAutomationRunsParams) (*PaginatedResponse[AutomationRun], error) {
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
	var out PaginatedResponse[AutomationRun]
	if err := r.client.get(ctx, fmt.Sprintf("/v1/automations/%s/runs", url.PathEscape(automationID)), q, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetRun retrieves a single automation run by ID.
//
// Example:
//
//	run, err := client.Automations.GetRun(ctx, "auto_abc123", "run_xyz")
func (r *AutomationsResource) GetRun(ctx context.Context, automationID, runID string) (*AutomationRun, error) {
	var out AutomationRun
	if err := r.client.get(ctx, fmt.Sprintf("/v1/automations/%s/runs/%s", url.PathEscape(automationID), url.PathEscape(runID)), nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// ListRunSteps returns all step-execution records for an automation run.
//
// Example:
//
//	steps, err := client.Automations.ListRunSteps(ctx, "auto_abc123", "run_xyz")
func (r *AutomationsResource) ListRunSteps(ctx context.Context, automationID, runID string) ([]AutomationRunStep, error) {
	var out struct {
		Data []AutomationRunStep `json:"data"`
	}
	if err := r.client.get(ctx, fmt.Sprintf("/v1/automations/%s/runs/%s/steps", url.PathEscape(automationID), url.PathEscape(runID)), nil, &out); err != nil {
		return nil, err
	}
	return out.Data, nil
}

// CancelRun cancels an in-progress automation run.
//
// Example:
//
//	run, err := client.Automations.CancelRun(ctx, "auto_abc123", "run_xyz")
func (r *AutomationsResource) CancelRun(ctx context.Context, automationID, runID string) (*AutomationRun, error) {
	var out AutomationRun
	if err := r.client.post(ctx, fmt.Sprintf("/v1/automations/%s/runs/%s/cancel", url.PathEscape(automationID), url.PathEscape(runID)), nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// CreateRun creates a new manual automation run for a contact.
//
// Example:
//
//	run, err := client.Automations.CreateRun(ctx, "auto_abc123", sendry.CreateAutomationRunParams{
//	    ContactEmail: "jane@example.com",
//	})
func (r *AutomationsResource) CreateRun(ctx context.Context, automationID string, params CreateAutomationRunParams) (*AutomationRun, error) {
	var out AutomationRun
	if err := r.client.post(ctx, fmt.Sprintf("/v1/automations/%s/runs", url.PathEscape(automationID)), params, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

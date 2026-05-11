package sendry

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

// DeliverabilityResource provides methods for reputation, blocklist, and
// deliverability reporting.
type DeliverabilityResource struct {
	client *Client
}

// GetReputation returns a reputation overview for the organisation or a specific domain.
//
// Example:
//
//	rep, err := client.Deliverability.GetReputation(ctx, &sendry.ReputationQuery{Days: sendry.IntPtr(30)})
func (r *DeliverabilityResource) GetReputation(ctx context.Context, params *ReputationQuery) (*ReputationResponse, error) {
	q := url.Values{}
	if params != nil {
		if params.DomainID != "" {
			q.Set("domain_id", params.DomainID)
		}
		if params.Days != nil {
			q.Set("days", strconv.Itoa(*params.Days))
		}
	}
	var out ReputationResponse
	if err := r.client.get(ctx, "/v1/deliverability/reputation", q, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetReputationHistory returns reputation score history for a specific domain.
//
// Example:
//
//	history, err := client.Deliverability.GetReputationHistory(ctx, "dom_abc123", &sendry.ReputationHistoryQuery{
//	    From: "2025-01-01", To: "2025-01-31",
//	})
func (r *DeliverabilityResource) GetReputationHistory(ctx context.Context, domainID string, params *ReputationHistoryQuery) (map[string]any, error) {
	q := url.Values{}
	if params != nil {
		if params.From != "" {
			q.Set("from", params.From)
		}
		if params.To != "" {
			q.Set("to", params.To)
		}
	}
	var out map[string]any
	if err := r.client.get(ctx, fmt.Sprintf("/v1/deliverability/reputation/%s/history", url.PathEscape(domainID)), q, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// GetBlocklist returns blocklist status for the organisation's domains and IPs.
//
// Example:
//
//	bl, err := client.Deliverability.GetBlocklist(ctx, nil)
func (r *DeliverabilityResource) GetBlocklist(ctx context.Context, params *BlocklistQuery) (*BlocklistResponse, error) {
	q := url.Values{}
	if params != nil {
		if params.Target != "" {
			q.Set("target", params.Target)
		}
		if params.TargetType != "" {
			q.Set("target_type", params.TargetType)
		}
	}
	var out BlocklistResponse
	if err := r.client.get(ctx, "/v1/deliverability/blocklist", q, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// RunBlocklistCheck runs an on-demand blocklist check for a specific domain or IP.
//
// Example:
//
//	_, err := client.Deliverability.RunBlocklistCheck(ctx, sendry.BlocklistCheckBody{
//	    Target: "example.com", TargetType: "domain",
//	})
func (r *DeliverabilityResource) RunBlocklistCheck(ctx context.Context, params BlocklistCheckBody) (map[string]any, error) {
	var out map[string]any
	if err := r.client.post(ctx, "/v1/deliverability/blocklist/check", params, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// DismissAlert dismisses a blocklist alert by ID.
//
// Example:
//
//	_, err := client.Deliverability.DismissAlert(ctx, "alert_abc123")
func (r *DeliverabilityResource) DismissAlert(ctx context.Context, alertID string) (map[string]any, error) {
	var out map[string]any
	body := map[string]string{"status": "dismissed"}
	if err := r.client.patch(ctx, fmt.Sprintf("/v1/deliverability/blocklist/alerts/%s", url.PathEscape(alertID)), body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// GetReport returns a comprehensive deliverability report for a date range.
//
// Example:
//
//	report, err := client.Deliverability.GetReport(ctx, &sendry.DeliverabilityReportQuery{Days: sendry.IntPtr(30)})
func (r *DeliverabilityResource) GetReport(ctx context.Context, params *DeliverabilityReportQuery) (*DeliverabilityReport, error) {
	q := url.Values{}
	if params != nil {
		if params.DomainID != "" {
			q.Set("domain_id", params.DomainID)
		}
		if params.Days != nil {
			q.Set("days", strconv.Itoa(*params.Days))
		}
	}
	var out DeliverabilityReport
	if err := r.client.get(ctx, "/v1/deliverability/report", q, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

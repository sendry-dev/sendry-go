package sendry

import (
	"context"
	"net/url"
	"strconv"
)

// AnalyticsResource provides methods for fetching email analytics and event logs.
type AnalyticsResource struct {
	client *Client
}

// Stats returns aggregated email analytics including a summary and timeseries breakdown.
//
// Example:
//
//	stats, err := client.Analytics.Stats(ctx, sendr.AnalyticsParams{
//	    From:        "2025-01-01",
//	    To:          "2025-01-31",
//	    Granularity: "day",
//	})
//	fmt.Println(stats.Summary.DeliveryRate)
func (r *AnalyticsResource) Stats(ctx context.Context, params AnalyticsParams) (*AnalyticsResponse, error) {
	q := buildQuery(map[string]string{
		"from":        params.From,
		"to":          params.To,
		"granularity": params.Granularity,
		"tag":         params.Tag,
		"domain":      params.Domain,
	})
	var out AnalyticsResponse
	if err := r.client.get(ctx, "/v1/analytics", q, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Logs returns a cursor-paginated list of email tracking events.
//
// Example:
//
//	logs, err := client.Analytics.Logs(ctx, &sendr.LogsParams{
//	    EmailID: "em_abc123",
//	    Type:    "delivered",
//	})
func (r *AnalyticsResource) Logs(ctx context.Context, params *LogsParams) (*PaginatedResponse[LogEvent], error) {
	q := url.Values{}
	if params != nil {
		if params.Limit != nil {
			q.Set("limit", strconv.Itoa(*params.Limit))
		}
		if params.Cursor != nil {
			q.Set("cursor", *params.Cursor)
		}
		if params.EmailID != "" {
			q.Set("email_id", params.EmailID)
		}
		if params.Type != "" {
			q.Set("type", params.Type)
		}
		if params.To != "" {
			q.Set("to", params.To)
		}
		if params.FromDate != "" {
			q.Set("from_date", params.FromDate)
		}
		if params.ToDate != "" {
			q.Set("to_date", params.ToDate)
		}
	}
	var out PaginatedResponse[LogEvent]
	if err := r.client.get(ctx, "/v1/logs", q, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetCohorts returns cohort analysis data showing engagement decay over time.
//
// Example:
//
//	cohorts, err := client.Analytics.GetCohorts(ctx, sendr.CohortParams{
//	    From:   "2025-01-01",
//	    To:     "2025-01-31",
//	    Metric: "open_rate",
//	})
func (r *AnalyticsResource) GetCohorts(ctx context.Context, params CohortParams) (*CohortResponse, error) {
	q := buildQuery(map[string]string{
		"from":        params.From,
		"to":          params.To,
		"granularity": params.Granularity,
		"metric":      params.Metric,
	})
	var out CohortResponse
	if err := r.client.get(ctx, "/v1/analytics/cohorts", q, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetBenchmarks returns cross-organisation benchmark comparisons.
//
// Example:
//
//	benchmarks, err := client.Analytics.GetBenchmarks(ctx, sendr.BenchmarkParams{
//	    From: "2025-01-01",
//	    To:   "2025-01-31",
//	})
func (r *AnalyticsResource) GetBenchmarks(ctx context.Context, params BenchmarkParams) (*BenchmarkResponse, error) {
	q := buildQuery(map[string]string{
		"from":        params.From,
		"to":          params.To,
		"granularity": params.Granularity,
	})
	var out BenchmarkResponse
	if err := r.client.get(ctx, "/v1/analytics/benchmarks", q, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// ToggleBenchmarkOptIn opts the organisation in or out of anonymous benchmark data sharing.
//
// Example:
//
//	err := client.Analytics.ToggleBenchmarkOptIn(ctx, true)
func (r *AnalyticsResource) ToggleBenchmarkOptIn(ctx context.Context, optIn bool) error {
	body := map[string]bool{"opt_in": optIn}
	return r.client.post(ctx, "/v1/analytics/benchmarks/opt-in", body, nil)
}

// GetBreakdowns returns analytics broken down by domain or template.
//
// Example:
//
//	result, err := client.Analytics.GetBreakdowns(ctx, sendr.BreakdownParams{
//	    From:    "2025-01-01",
//	    To:      "2025-01-31",
//	    GroupBy: "domain",
//	})
func (r *AnalyticsResource) GetBreakdowns(ctx context.Context, params BreakdownParams) (*BreakdownResponse, error) {
	q := buildQuery(map[string]string{
		"from":     params.From,
		"to":       params.To,
		"group_by": params.GroupBy,
	})
	if params.Limit != nil {
		q.Set("limit", strconv.Itoa(*params.Limit))
	}
	var out BreakdownResponse
	if err := r.client.get(ctx, "/v1/analytics/breakdowns", q, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetComparison compares current period metrics against the equivalent previous period.
//
// Example:
//
//	comparison, err := client.Analytics.GetComparison(ctx, sendr.AnalyticsParams{
//	    From: "2025-01-01",
//	    To:   "2025-01-31",
//	})
//	fmt.Println(comparison.Changes.OpenRateDelta)
func (r *AnalyticsResource) GetComparison(ctx context.Context, params AnalyticsParams) (*ComparisonResponse, error) {
	q := buildQuery(map[string]string{
		"from":        params.From,
		"to":          params.To,
		"granularity": params.Granularity,
		"tag":         params.Tag,
		"domain":      params.Domain,
	})
	var out ComparisonResponse
	if err := r.client.get(ctx, "/v1/analytics/comparison", q, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// ExportData exports analytics data as CSV or JSON.
// The returned string is raw CSV or JSON depending on the params.Format field.
//
// Example:
//
//	data, err := client.Analytics.ExportData(ctx, sendr.ExportParams{
//	    From:   "2025-01-01",
//	    To:     "2025-01-31",
//	    Format: "csv",
//	})
func (r *AnalyticsResource) ExportData(ctx context.Context, params ExportParams) (any, error) {
	q := buildQuery(map[string]string{
		"from":        params.From,
		"to":          params.To,
		"granularity": params.Granularity,
		"format":      params.Format,
		"domain":      params.Domain,
	})
	var out any
	if err := r.client.get(ctx, "/v1/analytics/export", q, &out); err != nil {
		return nil, err
	}
	return out, nil
}


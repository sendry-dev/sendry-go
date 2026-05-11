// Package sendr provides a Go client for the Sendr email API.
//
// See https://api.sendr.dev for the API reference.
package sendry

// ---------------------------------------------------------------------------
// Shared
// ---------------------------------------------------------------------------

// PaginatedResponse is a cursor-paginated list response returned by list endpoints.
type PaginatedResponse[T any] struct {
	Data       []T     `json:"data"`
	HasMore    bool    `json:"has_more"`
	NextCursor *string `json:"next_cursor"`
}

// PaginationParams holds common cursor-based pagination query parameters.
type PaginationParams struct {
	// Limit is the maximum number of items to return (e.g. 25).
	Limit *int `json:"limit,omitempty"`
	// Cursor is the opaque cursor from the previous response's next_cursor field.
	Cursor *string `json:"cursor,omitempty"`
}

// DeleteResponse is returned by DELETE endpoints.
type DeleteResponse struct {
	Deleted bool `json:"deleted"`
}

// ---------------------------------------------------------------------------
// Emails
// ---------------------------------------------------------------------------

// Tag is a name/value pair attached to an email for analytics and filtering.
type Tag struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// Attachment represents a file attachment to include in an email.
type Attachment struct {
	// Filename is the file name including extension (e.g. "report.pdf").
	Filename string `json:"filename"`
	// Content is the base64-encoded file content.
	Content string `json:"content"`
	// ContentType is the MIME type. Defaults to application/octet-stream.
	ContentType string `json:"content_type,omitempty"`
}

// SendEmailParams are the parameters for sending a single email.
type SendEmailParams struct {
	From        string            `json:"from"`
	To          interface{}       `json:"to"` // string or []string
	CC          interface{}       `json:"cc,omitempty"`
	BCC         interface{}       `json:"bcc,omitempty"`
	ReplyTo     string            `json:"reply_to,omitempty"`
	Subject     string            `json:"subject"`
	HTML        string            `json:"html,omitempty"`
	Text        string            `json:"text,omitempty"`
	Tags        []Tag             `json:"tags,omitempty"`
	Headers     map[string]string `json:"headers,omitempty"`
	ScheduledAt string            `json:"scheduled_at,omitempty"`
	TemplateID  string            `json:"template_id,omitempty"`
	Variables   map[string]any    `json:"variables,omitempty"`
	Attachments []Attachment      `json:"attachments,omitempty"`
	Tracking    *bool             `json:"tracking,omitempty"`
}

// SendEmailResponse is returned after successfully sending an email.
type SendEmailResponse struct {
	ID        string `json:"id"`
	From      string `json:"from"`
	To        []string `json:"to"`
	Subject   string `json:"subject"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
}

// EmailStatus represents the delivery status of an email.
type EmailStatus string

const (
	EmailStatusQueued    EmailStatus = "queued"
	EmailStatusSending   EmailStatus = "sending"
	EmailStatusSent      EmailStatus = "sent"
	EmailStatusDelivered EmailStatus = "delivered"
	EmailStatusBounced   EmailStatus = "bounced"
	EmailStatusComplained EmailStatus = "complained"
	EmailStatusFailed    EmailStatus = "failed"
	EmailStatusCancelled EmailStatus = "cancelled"
)

// Email represents a sent email with its full details and status.
type Email struct {
	ID          string                           `json:"id"`
	From        string                           `json:"from"`
	To          []string                         `json:"to"`
	Subject     string                           `json:"subject"`
	Status      EmailStatus                      `json:"status"`
	CreatedAt   string                           `json:"created_at"`
	SentAt      *string                          `json:"sent_at"`
	LastEvent   *string                          `json:"last_event"`
	Attachments []struct {
		Filename    string `json:"filename"`
		ContentType string `json:"content_type"`
	} `json:"attachments,omitempty"`
}

// ListEmailsParams are the query parameters for listing emails.
type ListEmailsParams struct {
	PaginationParams
	Status string `json:"status,omitempty"`
}

// BatchEmailItem is a single email within a batch send request.
type BatchEmailItem struct {
	From        interface{}    `json:"from,omitempty"`
	To          interface{}    `json:"to"`
	CC          interface{}    `json:"cc,omitempty"`
	BCC         interface{}    `json:"bcc,omitempty"`
	ReplyTo     string         `json:"reply_to,omitempty"`
	Subject     string         `json:"subject,omitempty"`
	HTML        string         `json:"html,omitempty"`
	Text        string         `json:"text,omitempty"`
	Tags        []Tag          `json:"tags,omitempty"`
	Variables   map[string]any `json:"variables,omitempty"`
	Attachments []Attachment   `json:"attachments,omitempty"`
}

// SendBatchParams are the parameters for sending a batch of emails.
type SendBatchParams struct {
	From       string           `json:"from,omitempty"`
	Subject    string           `json:"subject,omitempty"`
	TemplateID string           `json:"template_id,omitempty"`
	Emails     []BatchEmailItem `json:"emails"`
}

// BatchEmailResponse is returned after a batch send operation.
type BatchEmailResponse struct {
	Data []struct {
		ID     string `json:"id"`
		Status string `json:"status"`
	} `json:"data"`
}

// SendMarketingEmailParams are the parameters for sending a marketing email.
type SendMarketingEmailParams struct {
	From           string            `json:"from"`
	To             interface{}       `json:"to"`
	Subject        string            `json:"subject"`
	HTML           string            `json:"html,omitempty"`
	Text           string            `json:"text,omitempty"`
	ReplyTo        string            `json:"reply_to,omitempty"`
	Headers        map[string]string `json:"headers,omitempty"`
	Tags           []Tag             `json:"tags,omitempty"`
	UnsubscribeURL string            `json:"unsubscribe_url"`
	ListID         string            `json:"list_id,omitempty"`
	ScheduledAt    string            `json:"scheduled_at,omitempty"`
	TemplateID     string            `json:"template_id,omitempty"`
}

// CancelEmailResponse is returned after cancelling a queued email.
type CancelEmailResponse struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

// ---------------------------------------------------------------------------
// Domains
// ---------------------------------------------------------------------------

// DnsRecord is a DNS record that must be added to verify a domain.
type DnsRecord struct {
	Type     string  `json:"type"`
	Host     string  `json:"host"`
	Value    string  `json:"value"`
	Name     string  `json:"name"`
	Priority *int    `json:"priority,omitempty"`
	Verified bool    `json:"verified"`
}

// Domain represents a sending domain registered in Sendr.
type Domain struct {
	ID         string      `json:"id"`
	Name       string      `json:"name"`
	Status     string      `json:"status"` // "pending" | "verified" | "failed"
	DnsRecords []DnsRecord `json:"dns_records"`
	CreatedAt  string      `json:"created_at"`
}

// CreateDomainParams are the parameters for adding a domain.
type CreateDomainParams struct {
	Name string `json:"name"`
}

// VerifyDomainResponse is returned after triggering domain DNS verification.
type VerifyDomainResponse struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Status       string `json:"status"`
	SpfVerified  bool   `json:"spf_verified"`
	DkimVerified bool   `json:"dkim_verified"`
	DmarcVerified bool  `json:"dmarc_verified"`
}

// ConfigureBimiParams are the parameters for setting up BIMI on a domain.
type ConfigureBimiParams struct {
	// LogoURL is the SVG logo URL (must be publicly accessible).
	LogoURL string `json:"logo_url"`
	// VmcURL is the Verified Mark Certificate URL (optional).
	VmcURL string `json:"vmc_url,omitempty"`
	// Selector is the BIMI selector name. Defaults to "default".
	Selector string `json:"selector,omitempty"`
}

// BimiConfig represents the BIMI configuration for a domain.
type BimiConfig struct {
	ID         string  `json:"id"`
	DomainID   string  `json:"domain_id"`
	LogoURL    string  `json:"logo_url"`
	VmcURL     *string `json:"vmc_url"`
	Selector   string  `json:"selector"`
	Status     string  `json:"status"`
	DnsRecord  *string `json:"dns_record"`
	VerifiedAt *string `json:"verified_at"`
	CreatedAt  string  `json:"created_at"`
	UpdatedAt  string  `json:"updated_at"`
}

// VerifyBimiResponse is returned after triggering BIMI DNS verification.
type VerifyBimiResponse struct {
	Verified bool   `json:"verified"`
	Status   string `json:"status"`
}

// ---------------------------------------------------------------------------
// Templates
// ---------------------------------------------------------------------------

// TemplateVariable describes a variable slot within a template.
type TemplateVariable struct {
	Type     string `json:"type"`
	Required bool   `json:"required,omitempty"`
	Default  any    `json:"default,omitempty"`
}

// CreateTemplateParams are the parameters for creating a new template.
type CreateTemplateParams struct {
	Name      string                      `json:"name"`
	Subject   string                      `json:"subject,omitempty"`
	HTML      string                      `json:"html,omitempty"`
	Engine    string                      `json:"engine,omitempty"` // "html" | "react"
	Variables map[string]TemplateVariable `json:"variables,omitempty"`
}

// UpdateTemplateParams are the parameters for updating an existing template.
type UpdateTemplateParams struct {
	Name      string                      `json:"name,omitempty"`
	Subject   string                      `json:"subject,omitempty"`
	HTML      string                      `json:"html,omitempty"`
	Engine    string                      `json:"engine,omitempty"`
	Variables map[string]TemplateVariable `json:"variables,omitempty"`
}

// Template represents an email template stored in Sendr.
type Template struct {
	ID        string                      `json:"id"`
	Name      string                      `json:"name"`
	Subject   *string                     `json:"subject"`
	HTML      *string                     `json:"html"`
	Engine    string                      `json:"engine"`
	Variables map[string]TemplateVariable `json:"variables,omitempty"`
	CreatedAt string                      `json:"created_at"`
	UpdatedAt string                      `json:"updated_at"`
}

// RenderTemplateParams are the parameters for rendering a template.
type RenderTemplateParams struct {
	Variables map[string]string `json:"variables,omitempty"`
}

// RenderTemplateResponse is returned after rendering a template.
type RenderTemplateResponse struct {
	HTML string `json:"html"`
	Text string `json:"text"`
}

// TemplateStarter is a pre-built starter template.
type TemplateStarter struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Subject     string   `json:"subject"`
	HTML        string   `json:"html"`
	Variables   []string `json:"variables"`
	Engine      string   `json:"engine"`
}

// VisualStarterSummary is a summary of a visual (block-based) starter template.
type VisualStarterSummary struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Category    string `json:"category"`
}

// CompileBlocksParams are the parameters for compiling a visual block design.
type CompileBlocksParams struct {
	Design    map[string]any    `json:"design"`
	Variables map[string]string `json:"variables,omitempty"`
}

// RenderAdhocParams are the parameters for rendering arbitrary HTML without saving.
type RenderAdhocParams struct {
	HTML      string            `json:"html"`
	Engine    string            `json:"engine,omitempty"` // "html" | "react" | "visual"
	Variables map[string]string `json:"variables,omitempty"`
}

// ---------------------------------------------------------------------------
// API Keys
// ---------------------------------------------------------------------------

// APIKeyScope represents the permission level of an API key.
type APIKeyScope string

const (
	APIKeyScopeFullAccess    APIKeyScope = "full_access"
	APIKeyScopeSendingAccess APIKeyScope = "sending_access"
	APIKeyScopeReadOnly      APIKeyScope = "read_only"
)

// CreateAPIKeyParams are the parameters for creating a new API key.
type CreateAPIKeyParams struct {
	Name  string      `json:"name"`
	Scope APIKeyScope `json:"scope,omitempty"`
}

// APIKey represents an existing API key (key value is masked).
type APIKey struct {
	ID         string  `json:"id"`
	Name       string  `json:"name"`
	Scope      string  `json:"scope"`
	KeyPrefix  string  `json:"key_prefix"`
	LastUsedAt *string `json:"last_used_at"`
	CreatedAt  string  `json:"created_at"`
}

// APIKeyCreated is returned immediately after creating a new API key.
// The Key field contains the full key and is only shown once.
type APIKeyCreated struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Scope     string `json:"scope"`
	Key       string `json:"key"`
	CreatedAt string `json:"created_at"`
}

// ---------------------------------------------------------------------------
// Webhooks
// ---------------------------------------------------------------------------

// CreateWebhookParams are the parameters for creating a webhook endpoint.
type CreateWebhookParams struct {
	URL    string   `json:"url"`
	Events []string `json:"events"`
}

// UpdateWebhookParams are the parameters for updating a webhook.
type UpdateWebhookParams struct {
	URL    string   `json:"url,omitempty"`
	Events []string `json:"events,omitempty"`
	Active *bool    `json:"active,omitempty"`
}

// Webhook represents a webhook endpoint with its signing secret.
type Webhook struct {
	ID        string   `json:"id"`
	URL       string   `json:"url"`
	Events    []string `json:"events"`
	Secret    string   `json:"secret"`
	Active    bool     `json:"active"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
}

// WebhookListItem represents a webhook endpoint in a list response (no secret).
type WebhookListItem struct {
	ID        string   `json:"id"`
	URL       string   `json:"url"`
	Events    []string `json:"events"`
	Active    bool     `json:"active"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
}

// ---------------------------------------------------------------------------
// Analytics
// ---------------------------------------------------------------------------

// AnalyticsParams are the query parameters for fetching analytics.
type AnalyticsParams struct {
	From        string `json:"from"`
	To          string `json:"to"`
	Granularity string `json:"granularity,omitempty"` // "hour" | "day" | "week" | "month"
	Tag         string `json:"tag,omitempty"`
	Domain      string `json:"domain,omitempty"`
}

// AnalyticsBucket holds aggregated metrics for a single time bucket.
type AnalyticsBucket struct {
	Date       string `json:"date"`
	Sent       int    `json:"sent"`
	Delivered  int    `json:"delivered"`
	Opened     int    `json:"opened"`
	Clicked    int    `json:"clicked"`
	Bounced    int    `json:"bounced"`
	Complained int    `json:"complained"`
}

// AnalyticsSummary holds aggregate totals and rates across all time buckets.
type AnalyticsSummary struct {
	Sent          int     `json:"sent"`
	Delivered     int     `json:"delivered"`
	Opened        int     `json:"opened"`
	Clicked       int     `json:"clicked"`
	Bounced       int     `json:"bounced"`
	Complained    int     `json:"complained"`
	DeliveryRate  float64 `json:"delivery_rate"`
	OpenRate      float64 `json:"open_rate"`
	ClickRate     float64 `json:"click_rate"`
	BounceRate    float64 `json:"bounce_rate"`
	ComplaintRate float64 `json:"complaint_rate"`
}

// AnalyticsResponse is returned by the analytics stats endpoint.
type AnalyticsResponse struct {
	Summary    AnalyticsSummary  `json:"summary"`
	Timeseries []AnalyticsBucket `json:"timeseries"`
}

// LogsParams are the query parameters for fetching email event logs.
type LogsParams struct {
	PaginationParams
	EmailID  string `json:"email_id,omitempty"`
	Type     string `json:"type,omitempty"`
	To       string `json:"to,omitempty"`
	FromDate string `json:"from_date,omitempty"`
	ToDate   string `json:"to_date,omitempty"`
}

// LogEvent is a single email tracking event entry.
type LogEvent struct {
	ID        string `json:"id"`
	EmailID   string `json:"email_id"`
	Type      string `json:"type"`
	Recipient string `json:"recipient"`
	Metadata  any    `json:"metadata,omitempty"`
	CreatedAt string `json:"created_at"`
}

// CohortParams are the query parameters for cohort analysis.
type CohortParams struct {
	From        string `json:"from"`
	To          string `json:"to"`
	Granularity string `json:"granularity,omitempty"` // "day" | "week" | "month"
	Metric      string `json:"metric,omitempty"`      // "open_rate" | "click_rate" | "delivery_rate"
}

// CohortBucket holds cohort analysis data for a specific period offset.
type CohortBucket struct {
	CohortDate   string  `json:"cohort_date"`
	PeriodOffset int     `json:"period_offset"`
	TotalSent    int     `json:"total_sent"`
	MetricValue  float64 `json:"metric_value"`
}

// CohortResponse is returned by the analytics cohorts endpoint.
type CohortResponse struct {
	Cohorts     []CohortBucket `json:"cohorts"`
	Granularity string         `json:"granularity"`
	Metric      string         `json:"metric"`
}

// BenchmarkParams are the query parameters for benchmark comparisons.
type BenchmarkParams struct {
	From        string `json:"from"`
	To          string `json:"to"`
	Granularity string `json:"granularity,omitempty"` // "day" | "week" | "month"
}

// BenchmarkBucket holds benchmark comparison data for a single date bucket.
type BenchmarkBucket struct {
	Date               string  `json:"date"`
	YourDeliveryRate   float64 `json:"your_delivery_rate"`
	YourOpenRate       float64 `json:"your_open_rate"`
	YourClickRate      float64 `json:"your_click_rate"`
	AvgDeliveryRate    float64 `json:"avg_delivery_rate"`
	AvgOpenRate        float64 `json:"avg_open_rate"`
	AvgClickRate       float64 `json:"avg_click_rate"`
	P50DeliveryRate    float64 `json:"p50_delivery_rate"`
	P50OpenRate        float64 `json:"p50_open_rate"`
	P50ClickRate       float64 `json:"p50_click_rate"`
	P75DeliveryRate    float64 `json:"p75_delivery_rate"`
	P75OpenRate        float64 `json:"p75_open_rate"`
	P75ClickRate       float64 `json:"p75_click_rate"`
}

// BenchmarkResponse is returned by the analytics benchmarks endpoint.
type BenchmarkResponse struct {
	Data            []BenchmarkBucket `json:"data"`
	BenchmarkOptIn  bool              `json:"benchmark_opt_in"`
	OrgCount        int               `json:"org_count"`
}

// BreakdownParams are the query parameters for analytics breakdowns.
type BreakdownParams struct {
	From    string `json:"from"`
	To      string `json:"to"`
	GroupBy string `json:"group_by"` // "domain" | "template"
	Limit   *int   `json:"limit,omitempty"`
}

// BreakdownItem holds metrics for a single dimension in a breakdown report.
type BreakdownItem struct {
	ID           *string `json:"id"`
	Name         *string `json:"name"`
	Sent         int     `json:"sent"`
	Delivered    int     `json:"delivered"`
	Opened       int     `json:"opened"`
	Clicked      int     `json:"clicked"`
	Bounced      int     `json:"bounced"`
	Complained   int     `json:"complained"`
	DeliveryRate float64 `json:"delivery_rate"`
	OpenRate     float64 `json:"open_rate"`
	ClickRate    float64 `json:"click_rate"`
}

// BreakdownResponse is returned by the analytics breakdowns endpoint.
type BreakdownResponse struct {
	Data    []BreakdownItem `json:"data"`
	GroupBy string          `json:"group_by"`
}

// ComparisonPeriodStats holds metrics for a single comparison period.
type ComparisonPeriodStats struct {
	Sent          int     `json:"sent"`
	Delivered     int     `json:"delivered"`
	Opened        int     `json:"opened"`
	Clicked       int     `json:"clicked"`
	Bounced       int     `json:"bounced"`
	Complained    int     `json:"complained"`
	DeliveryRate  float64 `json:"delivery_rate"`
	OpenRate      float64 `json:"open_rate"`
	ClickRate     float64 `json:"click_rate"`
	BounceRate    float64 `json:"bounce_rate"`
	ComplaintRate float64 `json:"complaint_rate"`
}

// ComparisonChanges holds percentage deltas between periods.
type ComparisonChanges struct {
	SentPct             float64 `json:"sent_pct"`
	DeliveredPct        float64 `json:"delivered_pct"`
	OpenedPct           float64 `json:"opened_pct"`
	ClickedPct          float64 `json:"clicked_pct"`
	BouncedPct          float64 `json:"bounced_pct"`
	ComplainedPct       float64 `json:"complained_pct"`
	DeliveryRateDelta   float64 `json:"delivery_rate_delta"`
	OpenRateDelta       float64 `json:"open_rate_delta"`
	ClickRateDelta      float64 `json:"click_rate_delta"`
	BounceRateDelta     float64 `json:"bounce_rate_delta"`
	ComplaintRateDelta  float64 `json:"complaint_rate_delta"`
}

// ComparisonResponse is returned by the analytics comparison endpoint.
type ComparisonResponse struct {
	Current  ComparisonPeriodStats `json:"current"`
	Previous ComparisonPeriodStats `json:"previous"`
	Changes  ComparisonChanges     `json:"changes"`
}

// ExportParams are the query parameters for exporting analytics data.
type ExportParams struct {
	From        string `json:"from"`
	To          string `json:"to"`
	Granularity string `json:"granularity,omitempty"` // "hour" | "day" | "week" | "month"
	Format      string `json:"format,omitempty"`      // "csv" | "json"
	Domain      string `json:"domain,omitempty"`
}

// ---------------------------------------------------------------------------
// Suppression
// ---------------------------------------------------------------------------

// SuppressionReason is the reason an address was added to the suppression list.
type SuppressionReason string

const (
	SuppressionReasonHardBounce  SuppressionReason = "hard_bounce"
	SuppressionReasonComplaint   SuppressionReason = "complaint"
	SuppressionReasonUnsubscribe SuppressionReason = "unsubscribe"
	SuppressionReasonManual      SuppressionReason = "manual"
)

// AddSuppressionParams are the parameters for adding an address to the suppression list.
type AddSuppressionParams struct {
	Email  string            `json:"email"`
	Reason SuppressionReason `json:"reason,omitempty"`
}

// SuppressionEntry is a single suppressed email address.
type SuppressionEntry struct {
	Email     string `json:"email"`
	Reason    string `json:"reason"`
	CreatedAt string `json:"created_at"`
}

// ---------------------------------------------------------------------------
// Unsubscribes
// ---------------------------------------------------------------------------

// CreateUnsubscribeParams are the parameters for adding a single unsubscribe record.
type CreateUnsubscribeParams struct {
	Email  string `json:"email"`
	ListID string `json:"list_id,omitempty"`
	Reason string `json:"reason,omitempty"`
}

// BatchUnsubscribeParams are the parameters for bulk-adding unsubscribe records.
type BatchUnsubscribeParams struct {
	Emails []string `json:"emails"`
	ListID string   `json:"list_id,omitempty"`
	Reason string   `json:"reason,omitempty"`
}

// BatchUnsubscribeResponse is returned after a batch unsubscribe operation.
type BatchUnsubscribeResponse struct {
	Inserted int `json:"inserted"`
}

// UnsubscribeEntry is a single unsubscribe record.
type UnsubscribeEntry struct {
	ID        string  `json:"id"`
	Email     string  `json:"email"`
	ListID    *string `json:"list_id"`
	Reason    *string `json:"reason"`
	CreatedAt string  `json:"created_at"`
}

// ListUnsubscribesParams are the query parameters for listing unsubscribe records.
type ListUnsubscribesParams struct {
	PaginationParams
	Email  string `json:"email,omitempty"`
	ListID string `json:"list_id,omitempty"`
}

// ---------------------------------------------------------------------------
// Contacts
// ---------------------------------------------------------------------------

// CreateContactParams are the parameters for creating a contact.
type CreateContactParams struct {
	Email        string         `json:"email"`
	FirstName    string         `json:"first_name,omitempty"`
	LastName     string         `json:"last_name,omitempty"`
	Metadata     map[string]any `json:"metadata,omitempty"`
	Unsubscribed *bool          `json:"unsubscribed,omitempty"`
	AudienceID   string         `json:"audience_id,omitempty"`
}

// UpdateContactParams are the parameters for updating an existing contact.
type UpdateContactParams struct {
	Email        string         `json:"email,omitempty"`
	FirstName    *string        `json:"first_name"`
	LastName     *string        `json:"last_name"`
	Metadata     map[string]any `json:"metadata"`
	Unsubscribed *bool          `json:"unsubscribed,omitempty"`
}

// Contact represents a contact stored in Sendr.
type Contact struct {
	ID           string  `json:"id"`
	Email        string  `json:"email"`
	FirstName    *string `json:"first_name"`
	LastName     *string `json:"last_name"`
	Metadata     any     `json:"metadata"`
	Unsubscribed bool    `json:"unsubscribed"`
	CreatedAt    string  `json:"created_at"`
	UpdatedAt    string  `json:"updated_at"`
}

// ListContactsParams are the query parameters for listing contacts.
type ListContactsParams struct {
	PaginationParams
	Email      string `json:"email,omitempty"`
	AudienceID string `json:"audience_id,omitempty"`
}

// BulkImportContactItem is a single contact in a bulk import operation.
type BulkImportContactItem struct {
	Email        string         `json:"email"`
	FirstName    string         `json:"first_name,omitempty"`
	LastName     string         `json:"last_name,omitempty"`
	Metadata     map[string]any `json:"metadata,omitempty"`
	Unsubscribed *bool          `json:"unsubscribed,omitempty"`
}

// BulkImportContactsParams are the parameters for bulk-importing contacts.
type BulkImportContactsParams struct {
	Contacts   []BulkImportContactItem `json:"contacts"`
	AudienceID string                  `json:"audience_id,omitempty"`
}

// BulkImportResult is returned after a bulk contact import.
type BulkImportResult struct {
	Created int `json:"created"`
	Updated int `json:"updated"`
	Total   int `json:"total"`
}

// ---------------------------------------------------------------------------
// Audiences
// ---------------------------------------------------------------------------

// CreateAudienceParams are the parameters for creating an audience.
type CreateAudienceParams struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

// UpdateAudienceParams are the parameters for updating an audience.
type UpdateAudienceParams struct {
	Name        string  `json:"name,omitempty"`
	Description *string `json:"description"`
}

// Audience represents a contact list (audience) in Sendr.
type Audience struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description *string `json:"description"`
	MemberCount int    `json:"member_count"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

// AddContactsToAudienceParams are the parameters for adding contacts to an audience.
type AddContactsToAudienceParams struct {
	ContactIDs []string `json:"contact_ids"`
}

// AddContactsToAudienceResult is returned after adding contacts to an audience.
type AddContactsToAudienceResult struct {
	Added int `json:"added"`
}

// ---------------------------------------------------------------------------
// Campaigns
// ---------------------------------------------------------------------------

// CampaignStatus represents the status of a campaign.
type CampaignStatus string

const (
	CampaignStatusDraft     CampaignStatus = "draft"
	CampaignStatusScheduled CampaignStatus = "scheduled"
	CampaignStatusSending   CampaignStatus = "sending"
	CampaignStatusPaused    CampaignStatus = "paused"
	CampaignStatusSent      CampaignStatus = "sent"
	CampaignStatusCancelled CampaignStatus = "cancelled"
)

// CampaignStats holds delivery statistics for a campaign.
type CampaignStats struct {
	TotalRecipients   int `json:"total_recipients"`
	SentCount         int `json:"sent_count"`
	DeliveredCount    int `json:"delivered_count"`
	OpenedCount       int `json:"opened_count"`
	ClickedCount      int `json:"clicked_count"`
	BouncedCount      int `json:"bounced_count"`
	ComplainedCount   int `json:"complained_count"`
	UnsubscribedCount int `json:"unsubscribed_count"`
	FailedCount       int `json:"failed_count"`
}

// Campaign represents a full campaign with content and stats.
type Campaign struct {
	ID            string         `json:"id"`
	Name          string         `json:"name"`
	Subject       string         `json:"subject"`
	From          string         `json:"from"`
	ReplyTo       *string        `json:"reply_to"`
	PreviewText   *string        `json:"preview_text"`
	AudienceID    string         `json:"audience_id"`
	TemplateID    *string        `json:"template_id"`
	HTML          *string        `json:"html"`
	Text          *string        `json:"text"`
	Status        string         `json:"status"`
	ScheduledAt   *string        `json:"scheduled_at"`
	SendStartedAt *string        `json:"send_started_at"`
	SentAt        *string        `json:"sent_at"`
	Stats         CampaignStats  `json:"stats"`
	CreatedAt     string         `json:"created_at"`
	UpdatedAt     string         `json:"updated_at"`
}

// CampaignListItem is a campaign summary for list responses.
type CampaignListItem struct {
	ID          string        `json:"id"`
	Name        string        `json:"name"`
	Subject     string        `json:"subject"`
	From        string        `json:"from"`
	AudienceID  string        `json:"audience_id"`
	Status      string        `json:"status"`
	ScheduledAt *string       `json:"scheduled_at"`
	SentAt      *string       `json:"sent_at"`
	Stats       CampaignStats `json:"stats"`
	CreatedAt   string        `json:"created_at"`
	UpdatedAt   string        `json:"updated_at"`
}

// CreateCampaignParams are the parameters for creating a new campaign.
type CreateCampaignParams struct {
	Name        string `json:"name"`
	Subject     string `json:"subject"`
	From        string `json:"from"`
	ReplyTo     string `json:"reply_to,omitempty"`
	PreviewText string `json:"preview_text,omitempty"`
	AudienceID  string `json:"audience_id"`
	TemplateID  string `json:"template_id,omitempty"`
	HTML        string `json:"html,omitempty"`
	Text        string `json:"text,omitempty"`
}

// UpdateCampaignParams are the parameters for updating a campaign.
type UpdateCampaignParams struct {
	Name        string  `json:"name,omitempty"`
	Subject     string  `json:"subject,omitempty"`
	From        string  `json:"from,omitempty"`
	ReplyTo     *string `json:"reply_to"`
	PreviewText *string `json:"preview_text"`
	AudienceID  string  `json:"audience_id,omitempty"`
	TemplateID  *string `json:"template_id"`
	HTML        *string `json:"html"`
	Text        *string `json:"text"`
}

// ScheduleCampaignParams are the parameters for scheduling a campaign.
type ScheduleCampaignParams struct {
	// ScheduledAt is an ISO 8601 datetime to send the campaign.
	ScheduledAt string `json:"scheduled_at"`
}

// CampaignActionResponse is returned after campaign state-change actions.
type CampaignActionResponse struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

// ListCampaignsParams are the query parameters for listing campaigns.
type ListCampaignsParams struct {
	PaginationParams
	Status string `json:"status,omitempty"`
}

// ---------------------------------------------------------------------------
// Billing
// ---------------------------------------------------------------------------

// BillingPlan represents the organisation's current billing plan.
type BillingPlan struct {
	Plan            string `json:"plan"`
	HasSubscription bool   `json:"hasSubscription"`
	BillingPeriod   string `json:"billingPeriod"` // "monthly" | "annual"
}

// BillingUsage represents the organisation's current billing usage summary.
type BillingUsage struct {
	EmailsSentThisPeriod int      `json:"emails_sent_this_period"`
	PlanLimit            int      `json:"plan_limit"`
	OverageCount         int      `json:"overage_count"`
	OverageRate          *float64 `json:"overage_rate"`
	PeriodEnd            *string  `json:"period_end"`
}

// CreateCheckoutParams are the parameters for creating a Stripe Checkout session.
type CreateCheckoutParams struct {
	Plan          string `json:"plan"`
	BillingPeriod string `json:"billingPeriod,omitempty"` // "monthly" | "annual"
	SuccessURL    string `json:"successUrl,omitempty"`
	CancelURL     string `json:"cancelUrl,omitempty"`
}

// CreatePortalParams are the parameters for creating a Stripe Billing Portal session.
type CreatePortalParams struct {
	ReturnURL string `json:"returnUrl,omitempty"`
}

// CheckoutSession contains the URL to redirect the user to for Stripe Checkout.
type CheckoutSession struct {
	URL string `json:"url"`
}

// PortalSession contains the URL to redirect the user to for the Stripe Billing Portal.
type PortalSession struct {
	URL string `json:"url"`
}

// ---------------------------------------------------------------------------
// Team
// ---------------------------------------------------------------------------

// TeamMember represents a member of the organisation's team.
type TeamMember struct {
	ID       string  `json:"id"`
	OrgID    string  `json:"org_id"`
	UserID   *string `json:"user_id"`
	Email    string  `json:"email"`
	Name     *string `json:"name"`
	Role     string  `json:"role"`
	Status   string  `json:"status"`
	InvitedAt string  `json:"invited_at"`
	JoinedAt  *string `json:"joined_at"`
}

// TeamSeats describes the team seat usage for the organisation.
type TeamSeats struct {
	Used      int  `json:"used"`
	Limit     int  `json:"limit"`
	Unlimited bool `json:"unlimited"`
}

// ListTeamResponse is returned by the list team endpoint.
type ListTeamResponse struct {
	Data  []TeamMember `json:"data"`
	Seats TeamSeats    `json:"seats"`
	Plan  string       `json:"plan"`
}

// InviteTeamMemberParams are the parameters for inviting a new team member.
type InviteTeamMemberParams struct {
	Email string `json:"email"`
	Role  string `json:"role,omitempty"` // "admin" | "member"
}

// UpdateTeamMemberRoleParams are the parameters for updating a team member's role.
type UpdateTeamMemberRoleParams struct {
	Role string `json:"role"`
}

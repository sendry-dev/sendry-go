// Package sendry provides a Go client for the Sendry email API.
//
// See https://api.sendry.online for the API reference.
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

// Domain represents a sending domain registered in Sendry.
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

// Template represents an email template stored in Sendry.
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

// Contact represents a contact stored in Sendry.
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

// Audience represents a contact list (audience) in Sendry.
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

// ---------------------------------------------------------------------------
// Dedicated IPs
// ---------------------------------------------------------------------------

// DedicatedIpAssignment describes a domain-to-IP assignment summary embedded
// inside a DedicatedIp.
type DedicatedIpAssignment struct {
	ID         string `json:"id"`
	DomainID   string `json:"domain_id"`
	DomainName string `json:"domain_name"`
	CreatedAt  string `json:"created_at"`
}

// DedicatedIp represents a dedicated sending IP provisioned for the organisation.
type DedicatedIp struct {
	ID             string                  `json:"id"`
	IPAddress      string                  `json:"ip_address"`
	Provider       string                  `json:"provider"`
	Status         string                  `json:"status"` // "provisioning" | "warming" | "active" | "inactive"
	WarmupDay      int                     `json:"warmup_day"`
	WarmupProgress int                     `json:"warmup_progress"`
	PoolName       *string                 `json:"pool_name"`
	Assignments    []DedicatedIpAssignment `json:"assignments,omitempty"`
	CreatedAt      string                  `json:"created_at"`
}

// ProvisionDedicatedIpParams are the parameters for provisioning a new dedicated IP.
type ProvisionDedicatedIpParams struct {
	Provider string `json:"provider,omitempty"` // "ses" | "mailgun"
}

// AssignIpParams are the parameters for assigning a dedicated IP to a domain.
type AssignIpParams struct {
	DomainID string `json:"domain_id"`
}

// IpAssignment is returned after assigning a dedicated IP to a domain.
type IpAssignment struct {
	ID        string `json:"id"`
	IPID      string `json:"ip_id"`
	DomainID  string `json:"domain_id"`
	CreatedAt string `json:"created_at"`
}

// ---------------------------------------------------------------------------
// Deliverability
// ---------------------------------------------------------------------------

// ReputationQuery are the query parameters for the reputation overview endpoint.
type ReputationQuery struct {
	DomainID string `json:"domain_id,omitempty"`
	Days     *int   `json:"days,omitempty"`
}

// ReputationHistoryQuery are the query parameters for reputation history.
type ReputationHistoryQuery struct {
	From string `json:"from,omitempty"`
	To   string `json:"to,omitempty"`
}

// ReputationFactors holds the underlying factors that make up a reputation score.
type ReputationFactors struct {
	BounceRate     float64 `json:"bounceRate"`
	ComplaintRate  float64 `json:"complaintRate"`
	DeliveryRate   float64 `json:"deliveryRate"`
	EngagementRate float64 `json:"engagementRate"`
}

// ReputationCurrent describes the current reputation snapshot.
type ReputationCurrent struct {
	Score           int               `json:"score"`
	Rating          string            `json:"rating"`
	Factors         ReputationFactors `json:"factors"`
	Recommendations []string          `json:"recommendations"`
}

// ReputationSnapshot is a single historical reputation data point.
type ReputationSnapshot struct {
	Date            string  `json:"date"`
	TotalSent       int     `json:"total_sent"`
	TotalDelivered  int     `json:"total_delivered"`
	TotalBounced    int     `json:"total_bounced"`
	TotalComplained int     `json:"total_complained"`
	TotalOpened     int     `json:"total_opened"`
	TotalClicked    int     `json:"total_clicked"`
	DeliveryRate    float64 `json:"delivery_rate"`
	BounceRate      float64 `json:"bounce_rate"`
	ComplaintRate   float64 `json:"complaint_rate"`
	OpenRate        float64 `json:"open_rate"`
	ClickRate       float64 `json:"click_rate"`
	ReputationScore int     `json:"reputation_score"`
}

// ReputationDomain summarises reputation for a single domain.
type ReputationDomain struct {
	DomainID   string `json:"domain_id"`
	DomainName string `json:"domain_name"`
	Score      int    `json:"score"`
	Rating     string `json:"rating"`
}

// ReputationResponse is returned by the reputation overview endpoint.
type ReputationResponse struct {
	Current ReputationCurrent    `json:"current"`
	History []ReputationSnapshot `json:"history"`
	Domains []ReputationDomain   `json:"domains"`
}

// BlocklistQuery are the query parameters for fetching blocklist status.
type BlocklistQuery struct {
	Target     string `json:"target,omitempty"`
	TargetType string `json:"target_type,omitempty"` // "domain" | "ip"
}

// BlocklistCheckBody are the parameters for an on-demand blocklist check.
type BlocklistCheckBody struct {
	Target     string `json:"target"`
	TargetType string `json:"target_type"` // "domain" | "ip"
}

// BlocklistCheckItem is a single blocklist check result.
type BlocklistCheckItem struct {
	ID             string  `json:"id"`
	Target         string  `json:"target"`
	TargetType     string  `json:"target_type"`
	Provider       string  `json:"provider"`
	Listed         bool    `json:"listed"`
	ListingReason  *string `json:"listing_reason"`
	ResponseTimeMS *int    `json:"response_time_ms"`
	CheckedAt      string  `json:"checked_at"`
}

// BlocklistAlertItem is an active blocklist alert.
type BlocklistAlertItem struct {
	ID         string  `json:"id"`
	Target     string  `json:"target"`
	TargetType string  `json:"target_type"`
	Provider   string  `json:"provider"`
	Status     string  `json:"status"`
	ListedAt   string  `json:"listed_at"`
	ResolvedAt *string `json:"resolved_at"`
}

// BlocklistSummary summarises blocklist check totals.
type BlocklistSummary struct {
	TotalTargets int `json:"total_targets"`
	ListedCount  int `json:"listed_count"`
	CleanCount   int `json:"clean_count"`
}

// BlocklistResponse is returned by the blocklist status endpoint.
type BlocklistResponse struct {
	Checks  []BlocklistCheckItem `json:"checks"`
	Alerts  []BlocklistAlertItem `json:"alerts"`
	Summary BlocklistSummary     `json:"summary"`
}

// DeliverabilityReportQuery are the query parameters for the deliverability report.
type DeliverabilityReportQuery struct {
	DomainID string `json:"domain_id,omitempty"`
	Days     *int   `json:"days,omitempty"`
}

// DeliverabilityPeriod describes the time range covered by a deliverability report.
type DeliverabilityPeriod struct {
	From string `json:"from"`
	To   string `json:"to"`
	Days int    `json:"days"`
}

// DeliverabilityMetrics holds aggregate metrics for a deliverability report.
type DeliverabilityMetrics struct {
	TotalSent       int     `json:"total_sent"`
	TotalDelivered  int     `json:"total_delivered"`
	TotalBounced    int     `json:"total_bounced"`
	TotalComplained int     `json:"total_complained"`
	DeliveryRate    float64 `json:"delivery_rate"`
	BounceRate      float64 `json:"bounce_rate"`
	ComplaintRate   float64 `json:"complaint_rate"`
	OpenRate        float64 `json:"open_rate"`
	ClickRate       float64 `json:"click_rate"`
}

// DeliverabilityReputation holds reputation summary in a deliverability report.
type DeliverabilityReputation struct {
	Score  int    `json:"score"`
	Rating string `json:"rating"`
	Trend  string `json:"trend"`
}

// DeliverabilityBlocklistStatus holds blocklist summary in a deliverability report.
type DeliverabilityBlocklistStatus struct {
	TotalListsChecked int  `json:"total_lists_checked"`
	ActiveListings    int  `json:"active_listings"`
	Clean             bool `json:"clean"`
}

// DeliverabilityInboxPlacement holds inbox placement estimates.
type DeliverabilityInboxPlacement struct {
	InboxPct   float64 `json:"inbox_pct"`
	SpamPct    float64 `json:"spam_pct"`
	MissingPct float64 `json:"missing_pct"`
}

// DeliverabilityAuthentication holds authentication status for a domain.
type DeliverabilityAuthentication struct {
	SPF   bool `json:"spf"`
	DKIM  bool `json:"dkim"`
	DMARC bool `json:"dmarc"`
	BIMI  bool `json:"bimi"`
}

// DeliverabilityReport is returned by the deliverability report endpoint.
type DeliverabilityReport struct {
	Period                 DeliverabilityPeriod          `json:"period"`
	Metrics                DeliverabilityMetrics         `json:"metrics"`
	Reputation             DeliverabilityReputation      `json:"reputation"`
	BlocklistStatus        DeliverabilityBlocklistStatus `json:"blocklist_status"`
	InboxPlacementEstimate DeliverabilityInboxPlacement  `json:"inbox_placement_estimate"`
	Recommendations        []string                      `json:"recommendations"`
	Authentication         DeliverabilityAuthentication  `json:"authentication"`
}

// ---------------------------------------------------------------------------
// Inbound
// ---------------------------------------------------------------------------

// InboundEmailAttachment is an attachment on a received inbound email.
type InboundEmailAttachment struct {
	Filename    string `json:"filename"`
	ContentType string `json:"contentType"`
	Size        int    `json:"size"`
	ContentID   string `json:"contentId,omitempty"`
}

// InboundEmail represents a received inbound email.
type InboundEmail struct {
	ID               string                   `json:"id"`
	From             string                   `json:"from"`
	To               []string                 `json:"to"`
	CC               []string                 `json:"cc"`
	Subject          *string                  `json:"subject"`
	Text             *string                  `json:"text"`
	HTML             *string                  `json:"html"`
	Headers          map[string]string        `json:"headers"`
	Attachments      []InboundEmailAttachment `json:"attachments"`
	WebhookDelivered bool                     `json:"webhook_delivered"`
	CreatedAt        string                   `json:"created_at"`
}

// InboundConfig is the inbound webhook forwarding configuration.
type InboundConfig struct {
	URL    *string `json:"url"`
	Secret *string `json:"secret"`
}

// UpdateInboundConfigParams are the parameters for updating the inbound webhook config.
type UpdateInboundConfigParams struct {
	URL    *string `json:"url"`
	Secret *string `json:"secret,omitempty"`
}

// ---------------------------------------------------------------------------
// Notification Preferences
// ---------------------------------------------------------------------------

// NotificationPreferences represents the current user's notification preferences.
type NotificationPreferences struct {
	ID                string `json:"id"`
	BounceAlerts      bool   `json:"bounce_alerts"`
	ComplaintAlerts   bool   `json:"complaint_alerts"`
	DeliveryFailures  bool   `json:"delivery_failures"`
	DomainIssues      bool   `json:"domain_issues"`
	DailySummary      bool   `json:"daily_summary"`
	WeeklyDigest      bool   `json:"weekly_digest"`
	MonthlyReport     bool   `json:"monthly_report"`
	AllEvents         bool   `json:"all_events"`
	DeliveryEvents    bool   `json:"delivery_events"`
	EngagementEvents  bool   `json:"engagement_events"`
	ComplianceEvents  bool   `json:"compliance_events"`
	CreatedAt         string `json:"created_at"`
	UpdatedAt         string `json:"updated_at"`
}

// UpdateNotificationPreferencesParams are the parameters for updating notification preferences.
// Note: the API expects camelCase keys here, not snake_case.
type UpdateNotificationPreferencesParams struct {
	BounceAlerts     *bool `json:"bounceAlerts,omitempty"`
	ComplaintAlerts  *bool `json:"complaintAlerts,omitempty"`
	DeliveryFailures *bool `json:"deliveryFailures,omitempty"`
	DomainIssues     *bool `json:"domainIssues,omitempty"`
	DailySummary     *bool `json:"dailySummary,omitempty"`
	WeeklyDigest     *bool `json:"weeklyDigest,omitempty"`
	MonthlyReport    *bool `json:"monthlyReport,omitempty"`
	AllEvents        *bool `json:"allEvents,omitempty"`
	DeliveryEvents   *bool `json:"deliveryEvents,omitempty"`
	EngagementEvents *bool `json:"engagementEvents,omitempty"`
	ComplianceEvents *bool `json:"complianceEvents,omitempty"`
}

// ---------------------------------------------------------------------------
// Organizations
// ---------------------------------------------------------------------------

// Organization represents the current organisation.
type Organization struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Plan      string `json:"plan"`
	CreatedAt string `json:"createdAt"`
}

// UpdateOrganizationParams are the parameters for updating organisation details.
type UpdateOrganizationParams struct {
	Name string `json:"name"`
}

// BrandingSettings are the organisation's branding settings.
type BrandingSettings struct {
	BrandColor             string  `json:"brand_color"`
	BrandLogo              *string `json:"brand_logo"`
	UnsubscribeHeading     *string `json:"unsubscribe_heading"`
	UnsubscribeMessage     *string `json:"unsubscribe_message"`
	UnsubscribeRedirectURL *string `json:"unsubscribe_redirect_url"`
}

// UpdateBrandingParams are the parameters for updating branding settings.
// All fields are optional and nullable. Use pointer-to-string-pointer-style
// helpers from the caller if you need to explicitly clear a field; the API
// treats a missing key as "leave unchanged" and a JSON null as "clear".
type UpdateBrandingParams struct {
	BrandColor             *string `json:"brand_color,omitempty"`
	BrandLogo              *string `json:"brand_logo,omitempty"`
	UnsubscribeHeading     *string `json:"unsubscribe_heading,omitempty"`
	UnsubscribeMessage     *string `json:"unsubscribe_message,omitempty"`
	UnsubscribeRedirectURL *string `json:"unsubscribe_redirect_url,omitempty"`
}

// ---------------------------------------------------------------------------
// Regions
// ---------------------------------------------------------------------------

// Region describes a single SES region available for sending.
type Region struct {
	RegionCode  string `json:"region_code"`
	DisplayName string `json:"display_name"`
	IsDefault   bool   `json:"is_default"`
}

// OrgRegionSettings are the organisation-level region settings.
type OrgRegionSettings struct {
	DefaultRegion *string `json:"default_region"`
	DataResidency string  `json:"data_residency"`
}

// UpdateOrgRegionParams are the parameters for updating organisation region settings.
type UpdateOrgRegionParams struct {
	DefaultRegion *string `json:"default_region,omitempty"`
	DataResidency string  `json:"data_residency,omitempty"` // "none" | "eu" | "us" | "ap"
}

// UpdateDomainRegionParams are the parameters for setting a domain's region override.
type UpdateDomainRegionParams struct {
	Region *string `json:"region"`
}

// RegionAnalyticsParams are the query parameters for region analytics.
type RegionAnalyticsParams struct {
	From string `json:"from"`
	To   string `json:"to"`
}

// RegionAnalyticsItem is a single region's send breakdown.
type RegionAnalyticsItem struct {
	Region     string  `json:"region"`
	Count      int     `json:"count"`
	Percentage float64 `json:"percentage"`
}

// RegionAnalyticsResponse is returned by the region analytics endpoint.
type RegionAnalyticsResponse struct {
	Data []RegionAnalyticsItem `json:"data"`
}

// ---------------------------------------------------------------------------
// Status
// ---------------------------------------------------------------------------

// IncidentUpdate is a single update posted to an incident.
type IncidentUpdate struct {
	ID        string `json:"id"`
	Status    string `json:"status"`
	Message   string `json:"message"`
	CreatedAt string `json:"created_at"`
}

// AffectedComponent identifies a component affected by an incident.
type AffectedComponent struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

// Incident represents a status-page incident.
type Incident struct {
	ID                 string              `json:"id"`
	Title              string              `json:"title"`
	Status             string              `json:"status"`
	Impact             string              `json:"impact"`
	StartsAt           *string             `json:"starts_at"`
	EndsAt             *string             `json:"ends_at"`
	ResolvedAt         *string             `json:"resolved_at"`
	CreatedAt          string              `json:"created_at"`
	UpdatedAt          string              `json:"updated_at"`
	Updates            []IncidentUpdate    `json:"updates"`
	AffectedComponents []AffectedComponent `json:"affected_components"`
}

// StatusComponent describes a single component on the status page.
type StatusComponent struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	Group       *string `json:"group"`
	Slug        string  `json:"slug"`
	Status      string  `json:"status"`
	Uptime90d   float64 `json:"uptime_90d"`
	SLATarget   float64 `json:"sla_target"`
	SLAMet      bool    `json:"sla_met"`
}

// SLASummary summarises SLA compliance.
type SLASummary struct {
	Target         float64 `json:"target"`
	CurrentUptime  float64 `json:"current_uptime"`
	SLAMet         bool    `json:"sla_met"`
}

// SystemStatus is the current operational status of Sendry.
type SystemStatus struct {
	Status          string            `json:"status"`
	Components      []StatusComponent `json:"components"`
	ActiveIncidents []Incident        `json:"active_incidents"`
	SLASummary      SLASummary        `json:"sla_summary"`
}

// LatencyHourBucket is an hourly latency rollup.
type LatencyHourBucket struct {
	Hour         string   `json:"hour"`
	P50MS        *float64 `json:"p50_ms"`
	P95MS        *float64 `json:"p95_ms"`
	P99MS        *float64 `json:"p99_ms"`
	SampleCount  int      `json:"sample_count"`
	TargetMetPct float64  `json:"target_met_pct"`
}

// LatencyStats holds latency rollups for a component over a window.
type LatencyStats struct {
	Component    string              `json:"component"`
	TargetMS     float64             `json:"target_ms"`
	CurrentP50MS *float64            `json:"current_p50_ms"`
	TargetMet    bool                `json:"target_met"`
	Hourly       []LatencyHourBucket `json:"hourly"`
}

// GetLatencyParams are the query parameters for the latency endpoint.
type GetLatencyParams struct {
	Component string `json:"component,omitempty"`
	Hours     *int   `json:"hours,omitempty"`
}

// ---------------------------------------------------------------------------
// Test Emails
// ---------------------------------------------------------------------------

// TestEmailSummary is a summary of a test-mode email for list responses.
type TestEmailSummary struct {
	ID          string   `json:"id"`
	From        string   `json:"from"`
	To          []string `json:"to"`
	Subject     string   `json:"subject"`
	MessageType string   `json:"message_type"`
	CreatedAt   string   `json:"created_at"`
}

// TestEmail is the full details of a captured test-mode email.
type TestEmail struct {
	ID          string   `json:"id"`
	From        string   `json:"from"`
	To          []string `json:"to"`
	CC          []string `json:"cc"`
	Subject     string   `json:"subject"`
	HTML        *string  `json:"html"`
	Text        *string  `json:"text"`
	MessageType string   `json:"message_type"`
	CreatedAt   string   `json:"created_at"`
}

// ---------------------------------------------------------------------------
// Automations
// ---------------------------------------------------------------------------

// AutomationStatus represents the lifecycle state of an automation.
type AutomationStatus string

const (
	AutomationStatusDraft    AutomationStatus = "draft"
	AutomationStatusActive   AutomationStatus = "active"
	AutomationStatusPaused   AutomationStatus = "paused"
	AutomationStatusArchived AutomationStatus = "archived"
)

// AutomationTriggerType describes how an automation is triggered.
type AutomationTriggerType string

const (
	AutomationTriggerEvent                  AutomationTriggerType = "event"
	AutomationTriggerContactAddedToSegment  AutomationTriggerType = "contact_added_to_segment"
	AutomationTriggerSchedule               AutomationTriggerType = "schedule"
	AutomationTriggerManual                 AutomationTriggerType = "manual"
)

// AutomationReentryPolicy describes how a contact may re-enter an automation.
type AutomationReentryPolicy string

const (
	AutomationReentryOncePerContact AutomationReentryPolicy = "once_per_contact"
	AutomationReentryCooldown       AutomationReentryPolicy = "cooldown"
	AutomationReentryAlways         AutomationReentryPolicy = "always"
)

// Automation represents an automation workflow definition.
type Automation struct {
	ID                     string                 `json:"id"`
	Name                   string                 `json:"name"`
	Description            *string                `json:"description"`
	Status                 AutomationStatus       `json:"status"`
	TriggerType            AutomationTriggerType  `json:"trigger_type"`
	TriggerConfig          map[string]any         `json:"trigger_config"`
	EntrySegmentID         *string                `json:"entry_segment_id"`
	ReentryPolicy          string                 `json:"reentry_policy"`
	ReentryCooldownSeconds *int                   `json:"reentry_cooldown_seconds"`
	TotalRuns              int                    `json:"total_runs"`
	ActiveRuns             int                    `json:"active_runs"`
	CompletedRuns          int                    `json:"completed_runs"`
	FailedRuns             int                    `json:"failed_runs"`
	CreatedAt              string                 `json:"created_at"`
	UpdatedAt              string                 `json:"updated_at"`
}

// AutomationStepType identifies an automation step's kind.
type AutomationStepType string

const (
	AutomationStepSendEmail AutomationStepType = "send_email"
	AutomationStepWait      AutomationStepType = "wait"
	AutomationStepBranch    AutomationStepType = "branch"
	AutomationStepABSplit   AutomationStepType = "ab_split"
)

// AutomationStep represents a single step within an automation.
type AutomationStep struct {
	ID            string             `json:"id"`
	AutomationID  string             `json:"automation_id"`
	ParentStepID  *string            `json:"parent_step_id"`
	BranchLabel   *string            `json:"branch_label"`
	Position      int                `json:"position"`
	Type          AutomationStepType `json:"type"`
	Config        map[string]any     `json:"config"`
	CreatedAt     string             `json:"created_at"`
	UpdatedAt     string             `json:"updated_at"`
}

// AutomationRunStatus is the lifecycle status of a single run.
type AutomationRunStatus string

const (
	AutomationRunPending    AutomationRunStatus = "pending"
	AutomationRunInProgress AutomationRunStatus = "in_progress"
	AutomationRunCompleted  AutomationRunStatus = "completed"
	AutomationRunFailed     AutomationRunStatus = "failed"
	AutomationRunCancelled  AutomationRunStatus = "cancelled"
)

// AutomationRun represents a single contact's execution of an automation.
type AutomationRun struct {
	ID              string              `json:"id"`
	AutomationID    string              `json:"automation_id"`
	ContactID       *string             `json:"contact_id"`
	ContactEmail    string              `json:"contact_email"`
	TriggerEventID  *string             `json:"trigger_event_id"`
	Status          AutomationRunStatus `json:"status"`
	CurrentStepID   *string             `json:"current_step_id"`
	Context         map[string]any      `json:"context"`
	StartedAt       string              `json:"started_at"`
	CompletedAt     *string             `json:"completed_at"`
	FailedAt        *string             `json:"failed_at"`
	FailureReason   *string             `json:"failure_reason"`
	CreatedAt       string              `json:"created_at"`
	UpdatedAt       string              `json:"updated_at"`
}

// AutomationRunStepStatus is the lifecycle status of a single run step.
type AutomationRunStepStatus string

const (
	AutomationRunStepPending    AutomationRunStepStatus = "pending"
	AutomationRunStepInProgress AutomationRunStepStatus = "in_progress"
	AutomationRunStepCompleted  AutomationRunStepStatus = "completed"
	AutomationRunStepSkipped    AutomationRunStepStatus = "skipped"
	AutomationRunStepFailed     AutomationRunStepStatus = "failed"
)

// AutomationRunStep is a single step's execution record within a run.
type AutomationRunStep struct {
	ID           string                  `json:"id"`
	RunID        string                  `json:"run_id"`
	StepID       string                  `json:"step_id"`
	Status       AutomationRunStepStatus `json:"status"`
	EmailID      *string                 `json:"email_id"`
	BranchTaken  *string                 `json:"branch_taken"`
	ScheduledFor *string                 `json:"scheduled_for"`
	StartedAt    *string                 `json:"started_at"`
	CompletedAt  *string                 `json:"completed_at"`
	Error        *string                 `json:"error"`
	CreatedAt    string                  `json:"created_at"`
}

// ListAutomationsParams are the query parameters for listing automations.
type ListAutomationsParams struct {
	PaginationParams
	Status AutomationStatus `json:"status,omitempty"`
}

// CreateAutomationParams are the parameters for creating an automation.
type CreateAutomationParams struct {
	Name                   string                  `json:"name"`
	Description            string                  `json:"description,omitempty"`
	TriggerType            AutomationTriggerType   `json:"trigger_type"`
	TriggerConfig          map[string]any          `json:"trigger_config,omitempty"`
	EntrySegmentID         string                  `json:"entry_segment_id,omitempty"`
	ReentryPolicy          AutomationReentryPolicy `json:"reentry_policy,omitempty"`
	ReentryCooldownSeconds *int                    `json:"reentry_cooldown_seconds,omitempty"`
}

// UpdateAutomationParams are the parameters for patching an automation.
type UpdateAutomationParams struct {
	Name                   string                  `json:"name,omitempty"`
	Description            *string                 `json:"description,omitempty"`
	TriggerConfig          map[string]any          `json:"trigger_config,omitempty"`
	EntrySegmentID         *string                 `json:"entry_segment_id,omitempty"`
	ReentryPolicy          AutomationReentryPolicy `json:"reentry_policy,omitempty"`
	ReentryCooldownSeconds *int                    `json:"reentry_cooldown_seconds,omitempty"`
}

// AddAutomationStepParams are the parameters for adding a step to an automation.
type AddAutomationStepParams struct {
	ParentStepID *string        `json:"parent_step_id,omitempty"`
	BranchLabel  *string        `json:"branch_label,omitempty"`
	Position     *int           `json:"position,omitempty"`
	Config       map[string]any `json:"config"`
}

// UpdateAutomationStepParams are the parameters for patching an automation step.
type UpdateAutomationStepParams struct {
	ParentStepID *string        `json:"parent_step_id,omitempty"`
	BranchLabel  *string        `json:"branch_label,omitempty"`
	Position     *int           `json:"position,omitempty"`
	Config       map[string]any `json:"config,omitempty"`
}

// ListAutomationRunsParams are the query parameters for listing automation runs.
type ListAutomationRunsParams struct {
	PaginationParams
	Status string `json:"status,omitempty"`
}

// CreateAutomationRunParams are the parameters for creating a manual automation run.
type CreateAutomationRunParams struct {
	ContactID    string         `json:"contact_id,omitempty"`
	ContactEmail string         `json:"contact_email,omitempty"`
	Context      map[string]any `json:"context,omitempty"`
}

// SendEmailStepConfig builds a map[string]any config for a send_email step.
type SendEmailStepConfig struct {
	TemplateID  string            `json:"template_id,omitempty"`
	From        string            `json:"from"`
	ReplyTo     string            `json:"reply_to,omitempty"`
	Subject     string            `json:"subject,omitempty"`
	HTML        string            `json:"html,omitempty"`
	Text        string            `json:"text,omitempty"`
	MessageType string            `json:"message_type,omitempty"` // "transactional" | "marketing"
	TopicID     string            `json:"topic_id,omitempty"`
	Variables   map[string]string `json:"variables,omitempty"`
}

// ToConfig converts the typed step config to a generic map suitable for the
// automation step config field. The "type" discriminator is included.
func (s SendEmailStepConfig) ToConfig() map[string]any {
	m := map[string]any{"type": "send_email", "from": s.From}
	if s.TemplateID != "" {
		m["template_id"] = s.TemplateID
	}
	if s.ReplyTo != "" {
		m["reply_to"] = s.ReplyTo
	}
	if s.Subject != "" {
		m["subject"] = s.Subject
	}
	if s.HTML != "" {
		m["html"] = s.HTML
	}
	if s.Text != "" {
		m["text"] = s.Text
	}
	if s.MessageType != "" {
		m["message_type"] = s.MessageType
	}
	if s.TopicID != "" {
		m["topic_id"] = s.TopicID
	}
	if len(s.Variables) > 0 {
		m["variables"] = s.Variables
	}
	return m
}

// WaitStepConfig builds a map[string]any config for a wait step.
type WaitStepConfig struct {
	DurationSeconds int `json:"duration_seconds"`
}

// ToConfig converts the typed step config to a generic map.
func (s WaitStepConfig) ToConfig() map[string]any {
	return map[string]any{"type": "wait", "duration_seconds": s.DurationSeconds}
}

// BranchStepCondition is the condition payload for a branch step.
type BranchStepCondition struct {
	Op            string `json:"op"`
	Property      string `json:"property,omitempty"`
	Value         any    `json:"value,omitempty"`
	SegmentID     string `json:"segment_id,omitempty"`
	EventName     string `json:"event_name,omitempty"`
	WithinSeconds *int   `json:"within_seconds,omitempty"`
}

// BranchStepConfig builds a map[string]any config for a branch step.
type BranchStepConfig struct {
	Condition BranchStepCondition `json:"condition"`
}

// ToConfig converts the typed step config to a generic map.
func (s BranchStepConfig) ToConfig() map[string]any {
	cond := map[string]any{"op": s.Condition.Op}
	if s.Condition.Property != "" {
		cond["property"] = s.Condition.Property
	}
	if s.Condition.Value != nil {
		cond["value"] = s.Condition.Value
	}
	if s.Condition.SegmentID != "" {
		cond["segment_id"] = s.Condition.SegmentID
	}
	if s.Condition.EventName != "" {
		cond["event_name"] = s.Condition.EventName
	}
	if s.Condition.WithinSeconds != nil {
		cond["within_seconds"] = *s.Condition.WithinSeconds
	}
	return map[string]any{"type": "branch", "condition": cond}
}

// ABSplitStepConfig builds a map[string]any config for an A/B split step.
type ABSplitStepConfig struct {
	A    int    `json:"a"`
	B    int    `json:"b"`
	Seed string `json:"seed,omitempty"`
}

// ToConfig converts the typed step config to a generic map.
func (s ABSplitStepConfig) ToConfig() map[string]any {
	m := map[string]any{
		"type":  "ab_split",
		"split": map[string]any{"a": s.A, "b": s.B},
	}
	if s.Seed != "" {
		m["seed"] = s.Seed
	}
	return m
}

// ---------------------------------------------------------------------------
// Events
// ---------------------------------------------------------------------------

// IngestedEvent represents an event ingested into Sendry.
type IngestedEvent struct {
	ID            string         `json:"id"`
	ExternalID    *string        `json:"external_id"`
	Name          string         `json:"name"`
	ContactEmail  *string        `json:"contact_email"`
	ContactID     *string        `json:"contact_id"`
	Payload       map[string]any `json:"payload"`
	ReceivedAt    string         `json:"received_at"`
	ProcessedAt   *string        `json:"processed_at"`
	TriggeredRuns int            `json:"triggered_runs"`
	Deduped       bool           `json:"deduped,omitempty"`
}

// IngestEventParams are the parameters for ingesting an event.
type IngestEventParams struct {
	Name         string         `json:"name"`
	EventID      string         `json:"event_id,omitempty"`
	ContactEmail string         `json:"contact_email,omitempty"`
	ContactID    string         `json:"contact_id,omitempty"`
	Payload      map[string]any `json:"payload,omitempty"`
}

// ListEventsParams are the query parameters for listing events.
type ListEventsParams struct {
	PaginationParams
	Name string `json:"name,omitempty"`
}

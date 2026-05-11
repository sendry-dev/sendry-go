# sendry-go

Official Go SDK for the [Sendry](https://sendry.online) email API.

## Installation

```bash
go get github.com/sendry-dev/sendry-go
```

Requires Go 1.21 or later.

## Quick Start

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/sendry-dev/sendry-go/sendry"
)

func main() {
    client := sendry.NewClient("sn_live_abc123")

    ctx := context.Background()

    resp, err := client.Emails.Send(ctx, sendry.SendEmailParams{
        From:    "hello@example.com",
        To:      "user@example.com",
        Subject: "Hello from Sendry",
        HTML:    "<h1>Hello!</h1><p>This email was sent via the Sendry Go SDK.</p>",
    })
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Email sent:", resp.ID)
}
```

## Options

Customise the client with functional options:

```go
import (
    "net/http"
    "time"

    "github.com/sendry-dev/sendry-go/sendry"
)

client := sendry.NewClient(
    "sn_live_abc123",
    sendry.WithBaseURL("https://api.sendry.online"),     // override base URL (e.g. for testing)
    sendry.WithTimeout(10*time.Second),              // per-request timeout (default 30s)
    sendry.WithMaxRetries(3),                        // retries for 5xx / network errors
    sendry.WithHTTPClient(&http.Client{              // bring your own http.Client
        Timeout: 15 * time.Second,
    }),
)
```

## Emails

### Send a transactional email

```go
resp, err := client.Emails.Send(ctx, sendry.SendEmailParams{
    From:    "Acme <noreply@acme.com>",
    To:      "customer@example.com",
    Subject: "Your order is confirmed",
    HTML:    "<p>Order #1234 is confirmed.</p>",
    Text:    "Order #1234 is confirmed.",
    Tags: []sendry.Tag{
        {Name: "category", Value: "transactional"},
    },
})
```

### Send to multiple recipients

```go
resp, err := client.Emails.Send(ctx, sendry.SendEmailParams{
    From:    "hello@example.com",
    To:      []string{"alice@example.com", "bob@example.com"},
    Subject: "Team update",
    HTML:    "<p>Hello team!</p>",
})
```

### Send with attachments

```go
import "encoding/base64"

data, _ := os.ReadFile("report.pdf")

resp, err := client.Emails.Send(ctx, sendry.SendEmailParams{
    From:    "reports@example.com",
    To:      "manager@example.com",
    Subject: "Monthly Report",
    HTML:    "<p>Please find the report attached.</p>",
    Attachments: []sendry.Attachment{
        {
            Filename:    "report.pdf",
            Content:     base64.StdEncoding.EncodeToString(data),
            ContentType: "application/pdf",
        },
    },
})
```

### Schedule an email

```go
resp, err := client.Emails.Send(ctx, sendry.SendEmailParams{
    From:        "hello@example.com",
    To:          "user@example.com",
    Subject:     "Reminder",
    HTML:        "<p>Don't forget!</p>",
    ScheduledAt: "2026-04-01T09:00:00Z",
})
```

### Get email status

```go
email, err := client.Emails.Get(ctx, "em_abc123")
fmt.Println(email.Status) // "delivered"
```

### List emails

```go
page, err := client.Emails.List(ctx, &sendry.ListEmailsParams{
    PaginationParams: sendry.PaginationParams{Limit: sendry.IntPtr(25)},
    Status: "delivered",
})

for _, email := range page.Data {
    fmt.Println(email.ID, email.Status)
}

if page.HasMore {
    // fetch next page
    next, err := client.Emails.List(ctx, &sendry.ListEmailsParams{
        PaginationParams: sendry.PaginationParams{Cursor: page.NextCursor},
    })
    _ = next
    _ = err
}
```

### Send a batch of emails

```go
result, err := client.Emails.SendBatch(ctx, sendry.SendBatchParams{
    From: "hello@example.com",
    Emails: []sendry.BatchEmailItem{
        {To: "alice@example.com", Subject: "Hi Alice", HTML: "<p>Hi Alice!</p>"},
        {To: "bob@example.com",   Subject: "Hi Bob",   HTML: "<p>Hi Bob!</p>"},
    },
})
```

### Send a marketing email

```go
resp, err := client.Emails.SendMarketing(ctx, sendry.SendMarketingEmailParams{
    From:           "newsletter@example.com",
    To:             "subscriber@example.com",
    Subject:        "March Newsletter",
    HTML:           "<p>Here's what's new...</p>",
    UnsubscribeURL: "https://example.com/unsubscribe?token=abc",
})
```

### Cancel a queued email

```go
result, err := client.Emails.Cancel(ctx, "em_abc123")
fmt.Println(result.Status) // "cancelled"
```

## Domains

```go
// Add a domain
domain, err := client.Domains.Create(ctx, sendry.CreateDomainParams{Name: "mail.example.com"})
fmt.Println(domain.DnsRecords) // Add these to your DNS provider

// Verify DNS records
result, err := client.Domains.Verify(ctx, domain.ID)
if result.SpfVerified && result.DkimVerified {
    fmt.Println("Domain is fully verified!")
}

// List domains
page, err := client.Domains.List(ctx, nil)

// Get a domain
d, err := client.Domains.Get(ctx, "dom_abc123")

// Delete a domain
_, err = client.Domains.Remove(ctx, "dom_abc123")

// Configure BIMI
bimi, err := client.Domains.ConfigureBimi(ctx, "dom_abc123", sendry.ConfigureBimiParams{
    LogoURL: "https://example.com/logo.svg",
})
```

## Templates

```go
// Create a template
tmpl, err := client.Templates.Create(ctx, sendry.CreateTemplateParams{
    Name:    "Welcome Email",
    Subject: "Welcome, {{name}}!",
    HTML:    "<h1>Hello {{name}}</h1>",
})

// Render a template with variables
rendered, err := client.Templates.Render(ctx, tmpl.ID, &sendry.RenderTemplateParams{
    Variables: map[string]string{"name": "Alice"},
})
fmt.Println(rendered.HTML)

// List, get, update, delete
page, err := client.Templates.List(ctx, nil)
t, err := client.Templates.Get(ctx, "tmpl_abc123")
updated, err := client.Templates.Update(ctx, "tmpl_abc123", sendry.UpdateTemplateParams{Subject: "New Subject"})
_, err = client.Templates.Remove(ctx, "tmpl_abc123")

// Browse starter templates
starters, err := client.Templates.ListStarters(ctx)
```

## API Keys

```go
// Create an API key (key is only shown once)
created, err := client.APIKeys.Create(ctx, sendry.CreateAPIKeyParams{
    Name:  "Production Key",
    Scope: sendry.APIKeyScopeSendingAccess,
})
fmt.Println("Save this key:", created.Key)

// List API keys (values masked)
page, err := client.APIKeys.List(ctx, nil)

// Revoke a key
_, err = client.APIKeys.Remove(ctx, "key_abc123")
```

## Webhooks

```go
// Create a webhook
hook, err := client.Webhooks.Create(ctx, sendry.CreateWebhookParams{
    URL:    "https://example.com/webhooks/sendry",
    Events: []string{"email.delivered", "email.bounced", "email.complained"},
})
fmt.Println("Signing secret:", hook.Secret)

// Update a webhook
updated, err := client.Webhooks.Update(ctx, hook.ID, sendry.UpdateWebhookParams{
    Active: sendry.BoolPtr(false),
})

// List, get, delete
page, err := client.Webhooks.List(ctx, nil)
w, err := client.Webhooks.Get(ctx, "wh_abc123")
_, err = client.Webhooks.Remove(ctx, "wh_abc123")
```

### Verify webhook signatures

Sendry signs each webhook delivery with an HMAC-SHA256 signature. Always verify
the signature before processing the event:

```go
import "github.com/sendry-dev/sendry-go/sendry"

func handleWebhook(w http.ResponseWriter, r *http.Request) {
    body, err := io.ReadAll(r.Body)
    if err != nil {
        http.Error(w, "cannot read body", http.StatusBadRequest)
        return
    }

    sig := r.Header.Get("X-Sendry-Signature")
    if !sendry.VerifyWebhookSignature(string(body), sig, "whsec_your_signing_secret") {
        http.Error(w, "invalid signature", http.StatusUnauthorized)
        return
    }

    // Process the verified event...
    w.WriteHeader(http.StatusOK)
}
```

## Analytics

```go
// Aggregated stats
stats, err := client.Analytics.Stats(ctx, sendry.AnalyticsParams{
    From:        "2025-01-01",
    To:          "2025-01-31",
    Granularity: "day",
})
fmt.Printf("Delivery rate: %.1f%%\n", stats.Summary.DeliveryRate*100)

// Event logs
logs, err := client.Analytics.Logs(ctx, &sendry.LogsParams{
    EmailID: "em_abc123",
    Type:    "delivered",
})

// Breakdowns by domain or template
breakdowns, err := client.Analytics.GetBreakdowns(ctx, sendry.BreakdownParams{
    From:    "2025-01-01",
    To:      "2025-01-31",
    GroupBy: "domain",
})

// Period-over-period comparison
comparison, err := client.Analytics.GetComparison(ctx, sendry.AnalyticsParams{
    From: "2025-01-01",
    To:   "2025-01-31",
})
fmt.Printf("Open rate change: %+.1f pp\n", comparison.Changes.OpenRateDelta*100)

// Export as CSV
data, err := client.Analytics.ExportData(ctx, sendry.ExportParams{
    From:   "2025-01-01",
    To:     "2025-01-31",
    Format: "csv",
})
```

## Suppression

```go
// Add to suppression list
_, err := client.Suppression.Add(ctx, sendry.AddSuppressionParams{
    Email:  "bounced@example.com",
    Reason: sendry.SuppressionReasonHardBounce,
})

// List suppressed addresses
page, err := client.Suppression.List(ctx, nil)

// Remove from suppression list
_, err = client.Suppression.Remove(ctx, "bounced@example.com")
```

## Unsubscribes

```go
// Add a single unsubscribe
_, err := client.Unsubscribes.Create(ctx, sendry.CreateUnsubscribeParams{
    Email:  "user@example.com",
    ListID: "newsletter",
})

// Batch add
result, err := client.Unsubscribes.CreateBatch(ctx, sendry.BatchUnsubscribeParams{
    Emails: []string{"a@example.com", "b@example.com"},
    ListID: "newsletter",
})
fmt.Println(result.Inserted)

// List, get, remove
page, err := client.Unsubscribes.List(ctx, &sendry.ListUnsubscribesParams{ListID: "newsletter"})
entry, err := client.Unsubscribes.Get(ctx, "unsub_abc123")
_, err = client.Unsubscribes.Remove(ctx, "unsub_abc123")
```

## Contacts

```go
// Create a contact
contact, err := client.Contacts.Create(ctx, sendry.CreateContactParams{
    Email:      "jane@example.com",
    FirstName:  "Jane",
    LastName:   "Doe",
    AudienceID: "aud_abc123",
})

// List contacts
page, err := client.Contacts.List(ctx, &sendry.ListContactsParams{
    AudienceID: "aud_abc123",
})

// Update a contact
updated, err := client.Contacts.Update(ctx, contact.ID, sendry.UpdateContactParams{
    Unsubscribed: sendry.BoolPtr(true),
})

// Bulk import (upsert by email)
result, err := client.Contacts.BulkImport(ctx, sendry.BulkImportContactsParams{
    Contacts: []sendry.BulkImportContactItem{
        {Email: "alice@example.com", FirstName: "Alice"},
        {Email: "bob@example.com",   FirstName: "Bob"},
    },
    AudienceID: "aud_abc123",
})
fmt.Println(result.Created, "created,", result.Updated, "updated")
```

## Audiences

```go
// Create an audience
audience, err := client.Audiences.Create(ctx, sendry.CreateAudienceParams{
    Name:        "Newsletter Subscribers",
    Description: "Our weekly digest list",
})

// Add contacts to an audience
result, err := client.Audiences.AddContacts(ctx, audience.ID, sendry.AddContactsToAudienceParams{
    ContactIDs: []string{"ct_1", "ct_2", "ct_3"},
})
fmt.Println(result.Added)

// List contacts in an audience
page, err := client.Audiences.ListContacts(ctx, audience.ID, nil)

// Remove a contact from an audience
_, err = client.Audiences.RemoveContact(ctx, audience.ID, "ct_1")

// Update, delete audience
_, err = client.Audiences.Update(ctx, audience.ID, sendry.UpdateAudienceParams{Name: "VIP List"})
_, err = client.Audiences.Remove(ctx, audience.ID)
```

## Campaigns

```go
// Create a campaign in draft status
campaign, err := client.Campaigns.Create(ctx, sendry.CreateCampaignParams{
    Name:       "March Newsletter",
    Subject:    "What's new in March",
    From:       "Acme <hello@acme.com>",
    AudienceID: "aud_abc123",
    HTML:       "<h1>Hello subscribers!</h1>",
})

// Schedule it
_, err = client.Campaigns.Schedule(ctx, campaign.ID, sendry.ScheduleCampaignParams{
    ScheduledAt: "2026-03-15T10:00:00Z",
})

// Or send immediately
_, err = client.Campaigns.Send(ctx, campaign.ID)

// Pause, resume, cancel
_, err = client.Campaigns.Pause(ctx, campaign.ID)
_, err = client.Campaigns.Resume(ctx, campaign.ID)
_, err = client.Campaigns.Cancel(ctx, campaign.ID)

// List campaigns
page, err := client.Campaigns.List(ctx, &sendry.ListCampaignsParams{Status: "sent"})

// Get campaign with stats
c, err := client.Campaigns.Get(ctx, campaign.ID)
fmt.Printf("Delivered: %d / %d\n", c.Stats.DeliveredCount, c.Stats.TotalRecipients)
```

## Billing

```go
// Get current plan
plan, err := client.Billing.GetPlan(ctx)
fmt.Println(plan.Plan, plan.BillingPeriod) // "pro" "monthly"

// Get usage
usage, err := client.Billing.GetUsage(ctx)
fmt.Printf("%d / %d emails sent\n", usage.EmailsSentThisPeriod, usage.PlanLimit)

// Create checkout session (redirect user to URL)
session, err := client.Billing.CreateCheckout(ctx, sendry.CreateCheckoutParams{
    Plan:          "pro",
    BillingPeriod: "annual",
    SuccessURL:    "https://app.acme.com/billing?success=1",
    CancelURL:     "https://app.acme.com/billing",
})
fmt.Println("Redirect to:", session.URL)

// Create billing portal session
portal, err := client.Billing.CreatePortal(ctx, &sendry.CreatePortalParams{
    ReturnURL: "https://app.acme.com/billing",
})
fmt.Println("Redirect to:", portal.URL)
```

## Team

```go
// List team members
team, err := client.Team.List(ctx)
fmt.Printf("Using %d of %d seats\n", team.Seats.Used, team.Seats.Limit)

// Invite a team member
member, err := client.Team.Invite(ctx, sendry.InviteTeamMemberParams{
    Email: "colleague@acme.com",
    Role:  "member",
})

// Update role
updated, err := client.Team.UpdateRole(ctx, member.ID, sendry.UpdateTeamMemberRoleParams{
    Role: "admin",
})

// Remove a member
_, err = client.Team.Remove(ctx, member.ID)
```

## Error Handling

All methods return a typed error. Use `errors.As` to check for specific error types:

```go
import (
    "errors"
    "fmt"

    "github.com/sendry-dev/sendry-go/sendry"
)

resp, err := client.Emails.Send(ctx, params)
if err != nil {
    var authErr *sendry.AuthenticationError
    var validErr *sendry.ValidationError
    var rateLimitErr *sendry.RateLimitError
    var notFoundErr *sendry.NotFoundError
    var netErr *sendry.NetworkError
    var apiErr *sendry.Error

    switch {
    case errors.As(err, &authErr):
        fmt.Println("Invalid API key")
    case errors.As(err, &validErr):
        fmt.Println("Validation failed:", validErr.Message)
        fmt.Printf("Details: %v\n", validErr.Details)
    case errors.As(err, &rateLimitErr):
        fmt.Printf("Rate limited — retry after %ds\n", rateLimitErr.RetryAfter)
    case errors.As(err, &notFoundErr):
        fmt.Println("Resource not found")
    case errors.As(err, &netErr):
        fmt.Println("Network error:", netErr.Err)
    case errors.As(err, &apiErr):
        fmt.Printf("API error %d (%s): %s\n", apiErr.StatusCode, apiErr.Code, apiErr.Message)
    default:
        fmt.Println("Unexpected error:", err)
    }
    return
}
```

## Pagination

All list endpoints use cursor-based pagination. Iterate through all pages:

```go
var cursor *string

for {
    page, err := client.Emails.List(ctx, &sendry.ListEmailsParams{
        PaginationParams: sendry.PaginationParams{
            Limit:  sendry.IntPtr(50),
            Cursor: cursor,
        },
    })
    if err != nil {
        return err
    }

    for _, email := range page.Data {
        fmt.Println(email.ID, email.Status)
    }

    if !page.HasMore {
        break
    }
    cursor = page.NextCursor
}
```

## Pointer Helpers

The SDK provides helpers for creating pointers to literal values, which are
required for optional fields in request structs:

```go
sendry.IntPtr(25)       // *int
sendry.StringPtr("foo") // *string
sendry.BoolPtr(true)    // *bool
```

## License

MIT

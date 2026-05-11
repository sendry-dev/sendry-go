package sendry

import "context"

// BillingResource provides methods for managing billing plans and Stripe sessions.
type BillingResource struct {
	client *Client
}

// GetPlan returns the organisation's current billing plan and subscription status.
//
// Example:
//
//	plan, err := client.Billing.GetPlan(ctx)
//	fmt.Println(plan.Plan, plan.BillingPeriod) // "pro" "monthly"
func (r *BillingResource) GetPlan(ctx context.Context) (*BillingPlan, error) {
	var out BillingPlan
	if err := r.client.get(ctx, "/v1/billing/plan", nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetUsage returns the current billing usage summary for the organisation.
//
// Example:
//
//	usage, err := client.Billing.GetUsage(ctx)
//	fmt.Printf("%d / %d emails sent this period\n", usage.EmailsSentThisPeriod, usage.PlanLimit)
func (r *BillingResource) GetUsage(ctx context.Context) (*BillingUsage, error) {
	var out BillingUsage
	if err := r.client.get(ctx, "/v1/billing/usage", nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// CreateCheckout creates a Stripe Checkout session for upgrading to a paid plan.
// Redirect the user to the returned URL to complete payment.
//
// Example:
//
//	session, err := client.Billing.CreateCheckout(ctx, sendry.CreateCheckoutParams{
//	    Plan:          "pro",
//	    BillingPeriod: "annual",
//	    SuccessURL:    "https://app.acme.com/billing?success=1",
//	})
//	// Redirect user to session.URL
func (r *BillingResource) CreateCheckout(ctx context.Context, params CreateCheckoutParams) (*CheckoutSession, error) {
	var out CheckoutSession
	if err := r.client.post(ctx, "/v1/billing/checkout", params, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// CreatePortal creates a Stripe Billing Portal session for managing the subscription.
// Redirect the user to the returned URL to manage their plan.
//
// Example:
//
//	portal, err := client.Billing.CreatePortal(ctx, &sendry.CreatePortalParams{
//	    ReturnURL: "https://app.acme.com/billing",
//	})
//	// Redirect user to portal.URL
func (r *BillingResource) CreatePortal(ctx context.Context, params *CreatePortalParams) (*PortalSession, error) {
	body := params
	if body == nil {
		body = &CreatePortalParams{}
	}
	var out PortalSession
	if err := r.client.post(ctx, "/v1/billing/portal", body, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

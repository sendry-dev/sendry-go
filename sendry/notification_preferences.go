package sendry

import "context"

// NotificationPreferencesResource provides methods for managing notification preferences.
type NotificationPreferencesResource struct {
	client *Client
}

// Get returns the current user's notification preferences.
//
// Example:
//
//	prefs, err := client.NotificationPreferences.Get(ctx)
func (r *NotificationPreferencesResource) Get(ctx context.Context) (*NotificationPreferences, error) {
	var out NotificationPreferences
	if err := r.client.get(ctx, "/v1/notification-preferences", nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Update updates notification preferences. Only provided fields are changed.
//
// Example:
//
//	prefs, err := client.NotificationPreferences.Update(ctx, sendry.UpdateNotificationPreferencesParams{
//	    BounceAlerts: sendry.BoolPtr(true),
//	})
func (r *NotificationPreferencesResource) Update(ctx context.Context, params UpdateNotificationPreferencesParams) (*NotificationPreferences, error) {
	var out NotificationPreferences
	if err := r.client.put(ctx, "/v1/notification-preferences", params, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

package sendry

import (
	"context"
	"fmt"
	"net/url"
)

// TeamResource provides methods for managing team members.
type TeamResource struct {
	client *Client
}

// List returns all team members including pending invitations, plus seat usage.
//
// Example:
//
//	team, err := client.Team.List(ctx)
//	fmt.Printf("%d / %d seats used\n", team.Seats.Used, team.Seats.Limit)
func (r *TeamResource) List(ctx context.Context) (*ListTeamResponse, error) {
	var out ListTeamResponse
	if err := r.client.get(ctx, "/v1/team", nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Invite sends an email invitation to a new team member. Requires admin or owner role.
//
// Example:
//
//	member, err := client.Team.Invite(ctx, sendry.InviteTeamMemberParams{
//	    Email: "colleague@acme.com",
//	    Role:  "member",
//	})
func (r *TeamResource) Invite(ctx context.Context, params InviteTeamMemberParams) (*TeamMember, error) {
	var out TeamMember
	if err := r.client.post(ctx, "/v1/team/invite", params, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Remove removes a team member. Requires admin or owner role.
//
// Example:
//
//	_, err := client.Team.Remove(ctx, "mem_abc123")
func (r *TeamResource) Remove(ctx context.Context, id string) (*DeleteResponse, error) {
	var out DeleteResponse
	if err := r.client.delete(ctx, fmt.Sprintf("/v1/team/%s", url.PathEscape(id)), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// UpdateRole updates a team member's role. Requires admin or owner role.
//
// Example:
//
//	updated, err := client.Team.UpdateRole(ctx, "mem_abc123", sendry.UpdateTeamMemberRoleParams{
//	    Role: "admin",
//	})
func (r *TeamResource) UpdateRole(ctx context.Context, id string, params UpdateTeamMemberRoleParams) (*TeamMember, error) {
	var out TeamMember
	if err := r.client.patch(ctx, fmt.Sprintf("/v1/team/%s/role", url.PathEscape(id)), params, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

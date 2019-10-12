//
// DISCLAIMER
//
// Copyright 2019 ArangoDB Inc, Cologne, Germany
//
// Author Ewout Prangsma
//

package v1

import (
	context "context"

	common "github.com/arangodb-managed/apis/common/v1"
)

type (
	// OrganizationCallback is a callback for individual organizations.
	OrganizationCallback func(context.Context, *Organization) error

	// OrganizationMemberCallback is a callback for individual organization members.
	OrganizationMemberCallback func(context.Context, *Member) error

	// OrganizationInviteCallback is a callback for individual organization invites.
	OrganizationInviteCallback func(context.Context, *OrganizationInvite) error

	// QuotaCallback is a callback for individual quotas.
	QuotaCallback func(context.Context, *Quota) error

	// QuotaDescriptionCallback is a callback for individual quota descriptions.
	QuotaDescriptionCallback func(context.Context, *QuotaDescription) error

	// ProjectCallback is a callback for individual projects.
	ProjectCallback func(context.Context, *Project) error

	// EventCallback is a callback for individual events.
	EventCallback func(context.Context, *Event) error
)

// ForEachOrganization iterates over all organizations that the authenticated
// caller is a member of, invoking the given callback for each organization.
func ForEachOrganization(ctx context.Context, c ResourceManagerServiceClient,
	req *common.ListOptions, cb OrganizationCallback) error {
	opts := req.CloneOrDefault(50)
	for {
		list, err := c.ListOrganizations(ctx, opts)
		if err != nil {
			return err
		}
		for _, item := range list.GetItems() {
			if err := cb(ctx, item); err != nil {
				return err
			}
			if err := ctx.Err(); err != nil {
				return err
			}
		}
		if len(list.GetItems()) < int(opts.PageSize) {
			// We're done
			return nil
		}
		opts.Page++
	}
}

// ForEachOrganizationMember iterates over all members of the organization
// identified by the given context ID.
func ForEachOrganizationMember(ctx context.Context, c ResourceManagerServiceClient,
	req *common.ListOptions, cb OrganizationMemberCallback) error {
	opts := req.CloneOrDefault(50)
	for {
		list, err := c.ListOrganizationMembers(ctx, opts)
		if err != nil {
			return err
		}
		for _, item := range list.GetItems() {
			if err := cb(ctx, item); err != nil {
				return err
			}
			if err := ctx.Err(); err != nil {
				return err
			}
		}
		if len(list.GetItems()) < int(opts.PageSize) {
			// We're done
			return nil
		}
		opts.Page++
	}
}

// ForEachOrganizationInvite iterates over all invites of the organization
// identified by the given context ID.
func ForEachOrganizationInvite(ctx context.Context, c ResourceManagerServiceClient,
	req *common.ListOptions, cb OrganizationInviteCallback) error {
	opts := req.CloneOrDefault(50)
	for {
		list, err := c.ListOrganizationInvites(ctx, opts)
		if err != nil {
			return err
		}
		for _, item := range list.GetItems() {
			if err := cb(ctx, item); err != nil {
				return err
			}
			if err := ctx.Err(); err != nil {
				return err
			}
		}
		if len(list.GetItems()) < int(opts.PageSize) {
			// We're done
			return nil
		}
		opts.Page++
	}
}

// ForEachMyOrganizationInvite iterates over all my invites.
func ForEachMyOrganizationInvite(ctx context.Context, c ResourceManagerServiceClient,
	req *common.ListOptions, cb OrganizationInviteCallback) error {
	opts := req.CloneOrDefault(50)
	for {
		list, err := c.ListMyOrganizationInvites(ctx, opts)
		if err != nil {
			return err
		}
		for _, item := range list.GetItems() {
			if err := cb(ctx, item); err != nil {
				return err
			}
			if err := ctx.Err(); err != nil {
				return err
			}
		}
		if len(list.GetItems()) < int(opts.PageSize) {
			// We're done
			return nil
		}
		opts.Page++
	}
}

// ForEachProject iterates over all projects of the organization
// identified by the given context ID.
func ForEachProject(ctx context.Context, c ResourceManagerServiceClient,
	req *common.ListOptions, cb ProjectCallback) error {
	opts := req.CloneOrDefault(50)
	for {
		list, err := c.ListProjects(ctx, opts)
		if err != nil {
			return err
		}
		for _, item := range list.GetItems() {
			if err := cb(ctx, item); err != nil {
				return err
			}
			if err := ctx.Err(); err != nil {
				return err
			}
		}
		if len(list.GetItems()) < int(opts.PageSize) {
			// We're done
			return nil
		}
		opts.Page++
	}
}

// ForEachOrganizationQuota iterates over all organization related
// quotas specified by the given request.
func ForEachOrganizationQuota(ctx context.Context, c ResourceManagerServiceClient,
	req ListQuotasRequest, cb QuotaCallback) error {
	req.Options = req.GetOptions().CloneOrDefault(50)
	for {
		list, err := c.ListOrganizationQuotas(ctx, &req)
		if err != nil {
			return err
		}
		for _, item := range list.GetItems() {
			if err := cb(ctx, item); err != nil {
				return err
			}
			if err := ctx.Err(); err != nil {
				return err
			}
		}
		if len(list.GetItems()) < int(req.Options.PageSize) {
			// We're done
			return nil
		}
		req.Options.Page++
	}
}

// ForEachProjectQuota iterates over all project related quotas
// specified by the given request.
func ForEachProjectQuota(ctx context.Context, c ResourceManagerServiceClient,
	req ListQuotasRequest, cb QuotaCallback) error {
	req.Options = req.GetOptions().CloneOrDefault(50)
	for {
		list, err := c.ListProjectQuotas(ctx, &req)
		if err != nil {
			return err
		}
		for _, item := range list.GetItems() {
			if err := cb(ctx, item); err != nil {
				return err
			}
			if err := ctx.Err(); err != nil {
				return err
			}
		}
		if len(list.GetItems()) < int(req.Options.PageSize) {
			// We're done
			return nil
		}
		req.Options.Page++
	}
}

// ForEachEvent iterates over all events specified by the given request.
func ForEachEvent(ctx context.Context, c ResourceManagerServiceClient,
	req ListEventOptions, cb EventCallback) error {
	req.Options = req.GetOptions().CloneOrDefault(50)
	for {
		list, err := c.ListEvents(ctx, &req)
		if err != nil {
			return err
		}
		for _, item := range list.GetItems() {
			if err := cb(ctx, item); err != nil {
				return err
			}
			if err := ctx.Err(); err != nil {
				return err
			}
		}
		if len(list.GetItems()) < int(req.Options.PageSize) {
			// We're done
			return nil
		}
		req.Options.Page++
	}
}

// ForEachQuotaDescription iterates over all quota descriptions.
func ForEachQuotaDescription(ctx context.Context, c ResourceManagerServiceClient,
	req *common.ListOptions, cb QuotaDescriptionCallback) error {
	req = req.CloneOrDefault(50)
	for {
		list, err := c.ListQuotaDescriptions(ctx, req)
		if err != nil {
			return err
		}
		for _, item := range list.GetItems() {
			if err := cb(ctx, item); err != nil {
				return err
			}
			if err := ctx.Err(); err != nil {
				return err
			}
		}
		if len(list.GetItems()) < int(req.PageSize) {
			// We're done
			return nil
		}
		req.Page++
	}
}

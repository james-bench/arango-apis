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
	// GroupCallback is a callback for individual groups.
	GroupCallback func(context.Context, *Group) error

	// GroupMemberCallback is a callback for individual group member IDs.
	GroupMemberCallback func(context.Context, string) error

	// RoleCallback is a callback for individual roles.
	RoleCallback func(context.Context, *Role) error

	// APIKeyCallback is a callback for individual API key.
	APIKeyCallback func(context.Context, *APIKey) error
)

// ForEachGroup iterates over all groups in an organization
// identified by given context ID, invoking the given callback for
// each group.
func ForEachGroup(ctx context.Context, listFunc func(ctx context.Context, req *common.ListOptions) (*GroupList, error),
	opts *common.ListOptions, cb GroupCallback) error {
	opts = opts.CloneOrDefault()
	for {
		list, err := listFunc(ctx, opts)
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

// ForEachGroupMember iterates over all members in the group
// identified by given context ID, invoking the given callback for
// each group member ID.
func ForEachGroupMember(ctx context.Context, listFunc func(ctx context.Context, req *common.ListOptions) (*GroupMemberList, error),
	opts *common.ListOptions, cb GroupMemberCallback) error {
	opts = opts.CloneOrDefault()
	for {
		list, err := listFunc(ctx, opts)
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

// ForEachRole iterates over all roles in an organization
// identified by given context ID, invoking the given callback for
// each role.
func ForEachRole(ctx context.Context, listFunc func(ctx context.Context, req *common.ListOptions) (*RoleList, error),
	opts *common.ListOptions, cb RoleCallback) error {
	opts = opts.CloneOrDefault()
	for {
		list, err := listFunc(ctx, opts)
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

// ForEachAPIKey iterates over all API keys of the authenticated user,
// invoking the given callback for each key.
func ForEachAPIKey(ctx context.Context, listFunc func(ctx context.Context, req *common.ListOptions) (*APIKeyList, error),
	opts *common.ListOptions, cb APIKeyCallback) error {
	opts = opts.CloneOrDefault()
	for {
		list, err := listFunc(ctx, opts)
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

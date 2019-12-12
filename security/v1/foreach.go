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
	// IPWhitelistCallback is a callback for individual IP whitelist.
	IPWhitelistCallback func(context.Context, *IPWhitelist) error
)

// ForEachIPWhitelist iterates over all IP whitelists in a project
// identified by given context ID, invoking the given callback for
// each IP whitelist.
func ForEachIPWhitelist(ctx context.Context, listFunc func(ctx context.Context, req *common.ListOptions) (*IPWhitelistList, error),
	opts *common.ListOptions, cb IPWhitelistCallback) error {
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

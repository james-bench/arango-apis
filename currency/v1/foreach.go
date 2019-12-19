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
	// CurrencyCallback is a callback for individual currencies.
	CurrencyCallback func(context.Context, *Currency) error
)

// ForEachCurrency iterates over all currencies,
// invoking the given callback for each IP currencies.
func ForEachCurrency(ctx context.Context, listFunc func(ctx context.Context, req *common.ListOptions) (*CurrencyList, error),
	opts *common.ListOptions, cb CurrencyCallback) error {
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

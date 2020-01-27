//
// DISCLAIMER
//
// Copyright 2019 ArangoDB GmbH, Cologne, Germany
//
// Author Ewout Prangsma
//

package v1

import (
	context "context"
)

type (
	// ProviderCallback is a callback for individual (cloud) provider.
	ProviderCallback func(context.Context, *Provider) error

	// RegionCallback is a callback for individual (cloud) region.
	RegionCallback func(context.Context, *Region) error
)

// ForEachProvider iterates over all providers that match the given request,
// invoking the given callback for each item.
func ForEachProvider(ctx context.Context, listFunc func(ctx context.Context, req *ListProvidersRequest) (*ProviderList, error),
	req ListProvidersRequest, cb ProviderCallback) error {
	req.Options = req.Options.CloneOrDefault()
	for {
		list, err := listFunc(ctx, &req)
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

// ForEachRegion iterates over all regions that match the given request,
// invoking the given callback for each item.
func ForEachRegion(ctx context.Context, listFunc func(ctx context.Context, req *ListRegionsRequest) (*RegionList, error),
	req ListRegionsRequest, cb RegionCallback) error {
	req.Options = req.Options.CloneOrDefault()
	for {
		list, err := listFunc(ctx, &req)
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

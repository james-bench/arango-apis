//
// DISCLAIMER
//
// Copyright 2020 ArangoDB GmbH, Cologne, Germany
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Copyright holder is ArangoDB GmbH, Cologne, Germany
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

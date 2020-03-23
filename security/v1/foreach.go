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

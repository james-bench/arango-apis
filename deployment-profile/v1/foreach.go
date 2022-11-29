//
// DISCLAIMER
//
// Copyright 2022 ArangoDB GmbH, Cologne, Germany
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

package v1

import "context"

type (
	// DeploymentProfileCallback is a callback for individual deployment profiles.
	DeploymentProfileCallback func(context.Context, *DeploymentProfile) error
)

// ForEachDeploymentProfile iterates over all deployment profiles in the organization
// identified by the given context ID,
// invoking the given callback for each deployment profile.
func ForEachDeploymentProfile(ctx context.Context, listFunc func(ctx context.Context, req *ListDeploymentProfilesRequest) (*DeploymentProfileList, error),
	opts ListDeploymentProfilesRequest, cb DeploymentProfileCallback) error {
	opts.Options = opts.Options.CloneOrDefault()
	for {
		list, err := listFunc(ctx, &opts)
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
		if len(list.GetItems()) < int(opts.GetOptions().PageSize) {
			// We're done
			return nil
		}
		opts.GetOptions().Page++
	}
}

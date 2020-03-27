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
// Author Gergely Brautigam
//

package v1

import context "context"

type (
	// ExampleDatasetInstallationCallback is a callback for individual example dataset installations.
	ExampleDatasetInstallationCallback func(context.Context, *ExampleDatasetInstallation) error

	// ExampleDatasetCallback is a callback for individual example datasets.
	ExampleDatasetCallback func(context.Context, *ExampleDataset) error
)

// ForEachExampleDatasetInstallation iterates over all example dataset installations for a specific deployment,
// invoking the given callback for each of them.
func ForEachExampleDatasetInstallation(ctx context.Context, listFunc func(ctx context.Context, req *ListExampleDatasetInstallationsRequest) (*ExampleDatasetInstallationList, error),
	req ListExampleDatasetInstallationsRequest, cb ExampleDatasetInstallationCallback) error {
	req.Options = req.GetOptions().CloneOrDefault()
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

// ForEachExampleDataset iterates over all example datasets,
// invoking the given callback for each of them.
func ForEachExampleDataset(ctx context.Context, listFunc func(ctx context.Context, req *ListExampleDatasetsRequest) (*ExampleDatasetList, error),
	req ListExampleDatasetsRequest, cb ExampleDatasetCallback) error {
	req.Options = req.GetOptions().CloneOrDefault()
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

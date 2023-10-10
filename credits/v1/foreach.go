//
// DISCLAIMER
// Copyright 2023 ArangoDB GmbH, Cologne, Germany
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
	// CreditUsageCallback is a callback for each credit usage item.
	CreditUsageCallback func(context.Context, *CreditBundleUsage) error
	// CreditUsageReportCallback is a callback for each credit usage report.
	CreditUsageReportCallback func(context.Context, *CreditUsageReport) error
)

// ForEachCreditUsage iterates over all credit usages that match the given filter,
// and invoke the callback for each item.
func ForEachCreditUsage(
	ctx context.Context,
	listFunc func(ctx context.Context, req *ListCreditBundleUsageRequest) (*CreditBundleUsageList, error),
	req *ListCreditBundleUsageRequest,
	cb CreditUsageCallback,
) error {
	req.Options = req.Options.CloneOrDefault()
	for {
		list, err := listFunc(ctx, req)
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
		if len(list.GetItems()) < int(req.GetOptions().GetPageSize()) {
			return nil
		}
		req.Options.Page++
	}
}

// ForEachCreditUsageReport iterates over all credit usage reports that match the given filter,
// and invoke the callback for each item.
func ForEachCreditUsageReport(
	ctx context.Context,
	listFunc func(ctx context.Context, req *ListCreditUsageReportsRequest) (*CreditUsageReportList, error),
	req *ListCreditUsageReportsRequest,
	cb CreditUsageReportCallback,
) error {
	req.Options = req.Options.CloneOrDefault()
	for {
		list, err := listFunc(ctx, req)
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
		if len(list.GetItems()) < int(req.GetOptions().GetPageSize()) {
			return nil
		}
		req.Options.Page++
	}
}

//
// DISCLAIMER
//
// Copyright 2019 ArangoDB Inc, Cologne, Germany
//
// Author Robert Stam
//

package v1

import (
	context "context"
)

type (
	// UsageItemCallback is a callback for individual usage items.
	UsageItemCallback func(context.Context, *UsageItem) error
)

// ForEachUsageItem iterates over all usage items that match given given filter,
// invoking the given callback for each usage item.
func ForEachUsageItem(ctx context.Context, listFunc func(ctx context.Context, req *ListUsageItemsRequest) (*UsageItemList, error),
	req ListUsageItemsRequest, cb UsageItemCallback) error {
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

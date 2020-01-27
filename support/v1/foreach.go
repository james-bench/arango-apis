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
	// PlanCallback is a callback for individual support plan.
	PlanCallback func(context.Context, *Plan) error
)

// ForEachPlan iterates over all support plans that match given given filter,
// invoking the given callback for each plan.
func ForEachPlan(ctx context.Context, listFunc func(ctx context.Context, req *ListPlansRequest) (*PlanList, error),
	req ListPlansRequest, cb PlanCallback) error {
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

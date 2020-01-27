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

	common "github.com/arangodb-managed/apis/common/v1"
)

type (
	// DeploymentCallback is a callback for individual deployments.
	DeploymentCallback func(context.Context, *Deployment) error

	// VersionCallback is a callback for individual (ArangoDB) versions.
	VersionCallback func(context.Context, *Version) error
)

// ForEachDeployment iterates over all deployments in a project
// identified by given context ID, invoking the given callback for
// each deployment.
func ForEachDeployment(ctx context.Context, listFunc func(ctx context.Context, req *common.ListOptions) (*DeploymentList, error),
	opts *common.ListOptions, cb DeploymentCallback) error {
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

// ForEachVersion iterates over all (ArangoDB) versions that match the given filter
// invoking the given callback for each version.
func ForEachVersion(ctx context.Context, listFunc func(ctx context.Context, req *ListVersionsRequest) (*VersionList, error),
	req ListVersionsRequest, cb VersionCallback) error {
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

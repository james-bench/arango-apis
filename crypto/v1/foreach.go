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
	// CACertificateCallback is a callback for individual certificates.
	CACertificateCallback func(context.Context, *CACertificate) error
)

// ForEachCACertificate iterates over all CA certificates in the project
// identified by the given context ID,
// invoking the given callback for each certificate.
func ForEachCACertificate(ctx context.Context, listFunc func(ctx context.Context, req *common.ListOptions) (*CACertificateList, error),
	opts *common.ListOptions, cb CACertificateCallback) error {
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

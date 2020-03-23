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
	// OrganizationCallback is a callback for individual organizations.
	OrganizationCallback func(context.Context, *Organization) error

	// OrganizationMemberCallback is a callback for individual organization members.
	OrganizationMemberCallback func(context.Context, *Member) error

	// OrganizationInviteCallback is a callback for individual organization invites.
	OrganizationInviteCallback func(context.Context, *OrganizationInvite) error

	// QuotaCallback is a callback for individual quotas.
	QuotaCallback func(context.Context, *Quota) error

	// QuotaDescriptionCallback is a callback for individual quota descriptions.
	QuotaDescriptionCallback func(context.Context, *QuotaDescription) error

	// ProjectCallback is a callback for individual projects.
	ProjectCallback func(context.Context, *Project) error

	// EventCallback is a callback for individual events.
	EventCallback func(context.Context, *Event) error
)

// ForEachOrganization iterates over all organizations that the authenticated
// caller is a member of, invoking the given callback for each organization.
func ForEachOrganization(ctx context.Context, listFunc func(ctx context.Context, req *common.ListOptions) (*OrganizationList, error),
	req *common.ListOptions, cb OrganizationCallback) error {
	opts := req.CloneOrDefault()
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

// ForEachOrganizationMember iterates over all members of the organization
// identified by the given context ID.
func ForEachOrganizationMember(ctx context.Context, listFunc func(ctx context.Context, req *common.ListOptions) (*MemberList, error),
	req *common.ListOptions, cb OrganizationMemberCallback) error {
	opts := req.CloneOrDefault()
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

// ForEachOrganizationInvite iterates over all invites of the organization
// identified by the given context ID.
func ForEachOrganizationInvite(ctx context.Context, listFunc func(ctx context.Context, req *common.ListOptions) (*OrganizationInviteList, error),
	req *common.ListOptions, cb OrganizationInviteCallback) error {
	opts := req.CloneOrDefault()
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

// ForEachMyOrganizationInvite iterates over all my invites.
func ForEachMyOrganizationInvite(ctx context.Context, listFunc func(ctx context.Context, req *common.ListOptions) (*OrganizationInviteList, error),
	req *common.ListOptions, cb OrganizationInviteCallback) error {
	opts := req.CloneOrDefault()
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

// ForEachProject iterates over all projects of the organization
// identified by the given context ID.
func ForEachProject(ctx context.Context, listFunc func(ctx context.Context, req *common.ListOptions) (*ProjectList, error),
	req *common.ListOptions, cb ProjectCallback) error {
	opts := req.CloneOrDefault()
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

// ForEachOrganizationQuota iterates over all organization related
// quotas specified by the given request.
func ForEachOrganizationQuota(ctx context.Context, listFunc func(ctx context.Context, req *ListQuotasRequest) (*QuotaList, error),
	req ListQuotasRequest, cb QuotaCallback) error {
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

// ForEachProjectQuota iterates over all project related quotas
// specified by the given request.
func ForEachProjectQuota(ctx context.Context, listFunc func(ctx context.Context, req *ListQuotasRequest) (*QuotaList, error),
	req ListQuotasRequest, cb QuotaCallback) error {
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

// ForEachEvent iterates over all events specified by the given request.
func ForEachEvent(ctx context.Context, listFunc func(ctx context.Context, req *ListEventOptions) (*EventList, error),
	req ListEventOptions, cb EventCallback) error {
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

// ForEachQuotaDescription iterates over all quota descriptions.
func ForEachQuotaDescription(ctx context.Context, listFunc func(ctx context.Context, req *common.ListOptions) (*QuotaDescriptionList, error),
	req *common.ListOptions, cb QuotaDescriptionCallback) error {
	req = req.CloneOrDefault()
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
		if len(list.GetItems()) < int(req.PageSize) {
			// We're done
			return nil
		}
		req.Page++
	}
}

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
// Author Robert Stam
//

package v1

import (
	context "context"
)

type (
	// AuditLogCallback is a callback for individual audit log.
	AuditLogCallback func(context.Context, *AuditLog) error

	// AuditLogArchiveCallback is a callback for individual audit log archive.
	AuditLogArchiveCallback func(context.Context, *AuditLogArchive) error
)

// ForEachAuditLog iterates over all audit logs for a specific organization,
// invoking the given callback for each audit log.
func ForEachAuditLog(ctx context.Context, listFunc func(ctx context.Context, req *ListAuditLogsRequest) (*AuditLogList, error),
	req ListAuditLogsRequest, cb AuditLogCallback) error {
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

// ForEachAuditLogArchive iterates over all audit log archives for a specific audit log,
// invoking the given callback for each audit log archive.
func ForEachAuditLogArchive(ctx context.Context, listFunc func(ctx context.Context, req *ListAuditLogArchivesRequest) (*AuditLogArchiveList, error),
	req ListAuditLogArchivesRequest, cb AuditLogArchiveCallback) error {
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

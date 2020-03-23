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
)

type (
	// BackupPolicyCallback is a callback for individual backup policy.
	BackupPolicyCallback func(context.Context, *BackupPolicy) error

	// BackupCallback is a callback for individual backup.
	BackupCallback func(context.Context, *Backup) error
)

// ForEachBackupPolicy iterates over all backup policies for a specific deployment,
// invoking the given callback for each backup policy.
func ForEachBackupPolicy(ctx context.Context, listFunc func(ctx context.Context, req *ListBackupPoliciesRequest) (*BackupPolicyList, error),
	req ListBackupPoliciesRequest, cb BackupPolicyCallback) error {
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

// ForEachBackup iterates over all backups for a specific deployment,
// invoking the given callback for each backup.
func ForEachBackup(ctx context.Context, listFunc func(ctx context.Context, req *ListBackupsRequest) (*BackupList, error),
	req ListBackupsRequest, cb BackupCallback) error {
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

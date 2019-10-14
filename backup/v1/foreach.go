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
)

type (
	// BackupPolicyCallback is a callback for individual backup policy.
	BackupPolicyCallback func(context.Context, *BackupPolicy) error

	// BackupCallback is a callback for individual backup.
	BackupCallback func(context.Context, *Backup) error
)

// ForEachBackupPolicy iterates over all backup policies for a specific deployment,
// invoking the given callback for each backup policy.
func ForEachBackupPolicy(ctx context.Context, c BackupServiceClient,
	req ListBackupPoliciesRequest, cb BackupPolicyCallback) error {
	req.Options = req.GetOptions().CloneOrDefault()
	for {
		list, err := c.ListBackupPolicies(ctx, &req)
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
func ForEachBackup(ctx context.Context, c BackupServiceClient,
	req ListBackupsRequest, cb BackupCallback) error {
	req.Options = req.GetOptions().CloneOrDefault()
	for {
		list, err := c.ListBackups(ctx, &req)
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

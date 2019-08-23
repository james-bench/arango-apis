//
// DISCLAIMER
//
// Copyright 2019 ArangoDB Inc, Cologne, Germany
//
// Author Robert Stam
//

package v1

import (
	"net/url"
	"path"

	data "github.com/arangodb-managed/apis/data/v1"
)

const (
	// KindBackupPolicy is a constants for the kind of BackupPolicy resources.
	KindBackupPolicy = "BackupPolicy"
	// KindBackup is a constants for the kind of Backup resources.
	KindBackup = "Backup"
)

// BackupPolicyURL creates a resource URL for the BackupPolicy with given ID
// in given context.
func BackupPolicyURL(organizationID, projectID, deploymentID, backupPolicyID string) string {
	return path.Join(data.DeploymentURL(organizationID, projectID, deploymentID), KindBackupPolicy, url.PathEscape(backupPolicyID))
}

// BackupURL creates a resource URL for the Backup with given ID
// in given context.
func BackupURL(organizationID, projectID, deploymentID, backupID string) string {
	return path.Join(data.DeploymentURL(organizationID, projectID, deploymentID), KindBackup, url.PathEscape(backupID))
}

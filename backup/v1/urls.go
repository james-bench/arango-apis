//
// DISCLAIMER
//
// Copyright 2019 ArangoDB GmbH, Cologne, Germany
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
// in given context (as individual IDs).
func BackupPolicyURL(organizationID, projectID, deploymentID, backupPolicyID string) string {
	return BackupPolicyURL2(data.DeploymentURL(organizationID, projectID, deploymentID), backupPolicyID)
}

// BackupPolicyURL2 creates a resource URL for the BackupPolicy with given ID
// in given context (as base URL).
func BackupPolicyURL2(deploymentURL, backupPolicyID string) string {
	return path.Join(deploymentURL, KindBackupPolicy, url.PathEscape(backupPolicyID))
}

// BackupURL creates a resource URL for the Backup with given ID
// in given context (as individual IDs).
func BackupURL(organizationID, projectID, deploymentID, backupID string) string {
	return BackupURL2(data.DeploymentURL(organizationID, projectID, deploymentID), backupID)
}

// BackupURL2 creates a resource URL for the Backup with given ID
// in given context (as base URL).
func BackupURL2(deploymentURL, backupID string) string {
	return path.Join(deploymentURL, KindBackup, url.PathEscape(backupID))
}

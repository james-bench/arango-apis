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

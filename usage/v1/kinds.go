//
// DISCLAIMER
//
// Copyright 2020-2021 ArangoDB GmbH, Cologne, Germany
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
// Author Robert Stam
//

package v1

const (
	// UsageItems kind

	// UsageItemKindDeploymentSize indicates a UsageItem that contains cloud
	// resources for Deployment.
	// UsageItems from this kind will not be closed automatically, so open items can exists.
	UsageItemKindDeploymentSize = "DeploymentSize"
	// UsageItemKindNetworkTransferSize indicates a UsageItem that contains
	// the amount of network traffic caused by a deployment (or member of a deployment).
	// UsageItems from this kind will be closed automatically, so no open items can exists.
	// The timespan of this UsageItem is normally 24h (unless the deployment is deleted)
	UsageItemKindNetworkTransferSize = "NetworkTransferSize"
	// UsageItemKindBackupStorageSize indicates a UsageItem that contains
	// the amount of cloud storage used by backups of a deployment.
	// UsageItems from this kind will not be closed automatically, so open items can exists.
	UsageItemKindBackupStorageSize = "BackupStorageSize"
	// UsageItemKindAuditLogSize indicates a UsageItem that contains
	// the amount of resources used by audit log (of a deployment).
	// UsageItems from this kind will be closed automatically, so no open items can exists.
	// The timespan of this UsageItem is normally 24h (unless the auditlog or deployment is deleted)
	UsageItemKindAuditLogSize = "AuditLogSize"
	// UsageItemKindAuditLogStorageSize indicates a UsageItem that contains
	// the amount of cloud storage used by audit log (of a deployment).
	// UsageItems from this kind will not be closed automatically, so open items can exists.
	UsageItemKindAuditLogStorageSize = "AuditLogStorageSize"
	// UpdateItemKindNotebookSize indicates a UsageItem that contains
	// the amount of resources used by the notebooks (of a deployment).
	// UsageItems from this kind will not be closed automatically, so open items can exists.
	UsageItemKindNotebookSize = "NotebookSize"
)

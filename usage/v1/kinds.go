//
// DISCLAIMER
//
// Copyright 2019 ArangoDB GmbH, Cologne, Germany
//
// Author Ewout Prangsma
//

package v1

const (
	// UsageItems kind

	// UsageItemKindDeploymentSize indicates a UsageItem that contains cloud
	// resources for Deployment.
	UsageItemKindDeploymentSize = "DeploymentSize"
	// UsageItemKindNetworkTransferSize indicates a UsageItem that contains
	// the amount of network traffic caused by a deployment (or member of a deployment).
	UsageItemKindNetworkTransferSize = "NetworkTransferSize"
	// UsageItemKindBackupStorageSize indicates a UsageItem that contains
	// the amount of cloud storage used by backups of a deployment.
	UsageItemKindBackupStorageSize = "BackupStorageSize"
)

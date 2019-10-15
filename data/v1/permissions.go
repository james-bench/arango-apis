//
// DISCLAIMER
//
// Copyright 2019 ArangoDB Inc, Cologne, Germany
//
// Author Ewout Prangsma
//

package v1

const (
	// Deployment permissions

	// PermissionDeploymentList is needed for listing deployments in a project
	PermissionDeploymentList = "data.deployment.list"
	// PermissionDeploymentGet is needed for fetching an individual deployments in a project
	PermissionDeploymentGet = "data.deployment.get"
	// PermissionDeploymentCreate is needed for create a deployment
	PermissionDeploymentCreate = "data.deployment.create"
	// PermissionDeploymentUpdate is needed for updating a deployment
	PermissionDeploymentUpdate = "data.deployment.update"
	// PermissionDeploymentDelete is needed for deleting a deployment
	PermissionDeploymentDelete = "data.deployment.delete"
	// PermissionDeploymentRestoreBackup is needed for restoring a backup of a deployment
	PermissionDeploymentRestoreBackup = "data.deployment.restore-backup"
)

const (
	// DeploymentCredentials permissions

	// PermissionDeploymentCredentialsGet is needed for fetching credentials of an individual deployments in a project
	PermissionDeploymentCredentialsGet = "data.deploymentcredentials.get"
)

const (
	// ServerSpecLimits permissions

	// PermissionLimitsGet is needed for fetching server spec limits in a project
	PermissionLimitsGet = "data.limits.get"
)

const (
	// ServerSpecPresets permissions

	// PermissionPresetsList is needed for listing server spec presets in a project
	PermissionPresetsList = "data.presets.list"
)

const (
	// NodeSize permissions

	// PermissionPresetsList is needed for listing server spec presets in a project
	PermissionNodeSizeList = "data.nodesize.list"
)

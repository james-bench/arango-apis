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
)

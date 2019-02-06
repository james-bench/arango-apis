//
// DISCLAIMER
//
// Copyright 2019 ArangoDB Inc, Cologne, Germany
//
// Author Ewout Prangsma
//

package v1

const (
	// Organization permissions

	// PermissionOrganizationGet is needed for fetching an individual group in an organization
	PermissionOrganizationGet = "resourcemanager.organization.get"
	// PermissionOrganizationUpdate is needed for updating an organization
	PermissionOrganizationUpdate = "resourcemanager.organization.update"
	// PermissionOrganizationDelete is needed for deleting an organization
	PermissionOrganizationDelete = "resourcemanager.organization.delete"
)

const (
	// Project permissions

	// PermissionProjectList is needed for listing projects in an organization
	PermissionProjectList = "resourcemanager.project.list"
	// PermissionProjectGet is needed for fetching an individual project in an organization
	PermissionProjectGet = "resourcemanager.project.get"
	// PermissionProjectCreate is needed for create a project
	PermissionProjectCreate = "resourcemanager.project.create"
	// PermissionProjectUpdate is needed for updating a project
	PermissionProjectUpdate = "resourcemanager.project.update"
	// PermissionProjectDelete is needed for deleting a project
	PermissionProjectDelete = "resourcemanager.project.delete"
)

const (
	// Event permissions

	// PermissionEventList is needed for listing events in an organization
	PermissionEventList = "resourcemanager.event.list"
)

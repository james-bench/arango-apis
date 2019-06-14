//
// DISCLAIMER
//
// Copyright 2019 ArangoDB Inc, Cologne, Germany
//
// Author Ewout Prangsma
//

package v1

const (
	// IPWhitelist permissions

	// PermissionIPWhitelistList is needed for listing deployments in a project
	PermissionIPWhitelistList = "security.ipwhitelist.list"
	// PermissionIPWhitelistGet is needed for fetching an individual deployments in a project
	PermissionIPWhitelistGet = "security.ipwhitelist.get"
	// PermissionIPWhitelistCreate is needed for create a deployment
	PermissionIPWhitelistCreate = "security.ipwhitelist.create"
	// PermissionIPWhitelistUpdate is needed for updating a deployment
	PermissionIPWhitelistUpdate = "security.ipwhitelist.update"
	// PermissionIPWhitelistDelete is needed for deleting a deployment
	PermissionIPWhitelistDelete = "security.ipwhitelist.delete"
)

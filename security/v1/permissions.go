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

	// PermissionIPWhitelistList is needed for listing IP whitelists in a project
	PermissionIPWhitelistList = "security.ipwhitelist.list"
	// PermissionIPWhitelistGet is needed for fetching an individual IP whitelists in a project
	PermissionIPWhitelistGet = "security.ipwhitelist.get"
	// PermissionIPWhitelistCreate is needed for create an IP whitelist
	PermissionIPWhitelistCreate = "security.ipwhitelist.create"
	// PermissionIPWhitelistUpdate is needed for updating an IP whitelist
	PermissionIPWhitelistUpdate = "security.ipwhitelist.update"
	// PermissionIPWhitelistDelete is needed for deleting an IP whitelist
	PermissionIPWhitelistDelete = "security.ipwhitelist.delete"
)

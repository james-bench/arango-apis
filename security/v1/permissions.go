//
// DISCLAIMER
//
// Copyright 2019 ArangoDB GmbH, Cologne, Germany
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

const (
	// IAMProvider permissions

	// PermissionIAMProviderList is needed for listing IAM providers in a project
	PermissionIAMProviderList = "security.iamprovider.list"
	// PermissionIAMProviderGet is needed for fetching an individual IAM providers in a project
	PermissionIAMProviderGet = "security.iamprovider.get"
	// PermissionIAMProviderCreate is needed for create an IAM provider
	PermissionIAMProviderCreate = "security.iamprovider.create"
	// PermissionIAMProviderUpdate is needed for updating an IAM provider
	PermissionIAMProviderUpdate = "security.iamprovider.update"
	// PermissionIAMProviderDelete is needed for deleting an IAM provider
	PermissionIAMProviderDelete = "security.iamprovider.delete"
)

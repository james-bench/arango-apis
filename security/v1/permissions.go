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
// Author Ewout Prangsma
//

package v1

const (
	// IPAllowlist permissions

	// PermissionIPAllowlistList is needed for listing IP allowlists in a project
	PermissionIPAllowlistList = "security.ipallowlist.list"
	// PermissionIPAllowlistGet is needed for fetching an individual IP allowlists in a project
	PermissionIPAllowlistGet = "security.ipallowlist.get"
	// PermissionIPAllowlistCreate is needed for create an IP allowlist
	PermissionIPAllowlistCreate = "security.ipallowlist.create"
	// PermissionIPAllowlistUpdate is needed for updating an IP allowlist
	PermissionIPAllowlistUpdate = "security.ipallowlist.update"
	// PermissionIPAllowlistDelete is needed for deleting an IP allowlist
	PermissionIPAllowlistDelete = "security.ipallowlist.delete"
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
	// PermissionIAMProviderSetDefault is needed for marking an IAM provider as the default in a project
	PermissionIAMProviderSetDefault = "security.iamprovider.set-default"
)

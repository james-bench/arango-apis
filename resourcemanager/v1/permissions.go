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

const (
	// OrganizationInvite permissions

	// PermissionOrganizationInviteList is needed for fetching invites in an organization
	PermissionOrganizationInviteList = "resourcemanager.organization-invite.list"
	// PermissionOrganizationInviteGet is needed for fetching an individual invite in an organization
	PermissionOrganizationInviteGet = "resourcemanager.organization-invite.get"
	// PermissionOrganizationInviteCreate is needed for creating an organization invite
	PermissionOrganizationInviteCreate = "resourcemanager.organization-invite.create"
	// PermissionOrganizationInviteDelete is needed for deleting an organization invite
	PermissionOrganizationInviteDelete = "resourcemanager.organization-invite.delete"
)

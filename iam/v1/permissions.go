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
	// Group permissions

	// PermissionGroupList is needed for listing groups in an organization
	PermissionGroupList = "iam.group.list"
	// PermissionGroupGet is needed for fetching an individual group in an organization
	PermissionGroupGet = "iam.group.get"
	// PermissionGroupCreate is needed for create a group in an organization
	PermissionGroupCreate = "iam.group.create"
	// PermissionGroupUpdate is needed for updating a group in an organization
	PermissionGroupUpdate = "iam.group.update"
	// PermissionGroupDelete is needed for deleting a group in an organization
	PermissionGroupDelete = "iam.group.delete"
)

const (
	// Role permissions

	// PermissionRoleList is needed for listing roles in an organization
	PermissionRoleList = "iam.role.list"
	// PermissionRoleGet is needed for fetching an individual role in an organization
	PermissionRoleGet = "iam.role.get"
	// PermissionRoleCreate is needed for create a role in an organization
	PermissionRoleCreate = "iam.role.create"
	// PermissionRoleUpdate is needed for updating a role in an organization
	PermissionRoleUpdate = "iam.role.update"
	// PermissionRoleDelete is needed for deleting a role in an organization
	PermissionRoleDelete = "iam.role.delete"
)

const (
	// Policy permissions

	// PermissionPolicyGet is needed for fetching an individual policy on a resource
	PermissionPolicyGet = "iam.policy.get"
	// PermissionPolicyUpdate is needed for updating a policy on a resource
	PermissionPolicyUpdate = "iam.policy.update"
	// PermissionPolicyList is needed for listing role bindings
	PermissionPolicyList = "iam.policy.list"
)

const (
	// User permissions

	// PermissionUserGetPersonalData is needed for fetching a personal data such as
	// mobile phone number for a user that is not the currently authenticated user.
	PermissionUserGetPersonalData = "iam.user.get-personal-data"
	// PermissionUserUpdate is needed to update the data of a user (other
	// than the currently authenticated user).
	PermissionUserUpdate = "iam.user.update"
)

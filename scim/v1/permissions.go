//
// DISCLAIMER
//
// Copyright 2022 ArangoDB GmbH, Cologne, Germany
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

package v1

const (
	// User permissions

	// PermissionUserGet is needed for listing / getting user infomration
	PermissionUserGet = "scim.user.get"
	// PermissionUserAdd is needed for adding a user
	PermissionUserAdd = "scim.user.add"
	// PermissionUserUpdate is needed for updating a user
	PermissionUserUpdate = "scim.user.update"
	// PermissionUserDelete is needed for removing a user
	PermissionUserDelete = "scim.user.delete"
)

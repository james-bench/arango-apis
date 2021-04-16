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

import (
	"net/url"
	"path"

	rm "github.com/arangodb-managed/apis/resourcemanager/url"
)

const (
	// KindGroup is a constants for the kind of Group resources.
	KindGroup = "Group"
	// KindRole is a constants for the kind of Role resources.
	KindRole = "Role"
	// KindAPIKey is a constants for the kind of APIKey resources.
	KindAPIKey = "APIKey"
	// KindUser is a constants for the kind of User resources.
	KindUser = "User"
)

// GroupURL creates a resource URL for the Group with given ID in given context.
func GroupURL(organizationID, groupID string) string {
	return path.Join(rm.OrganizationURL(organizationID), KindGroup, url.PathEscape(groupID))
}

// RoleURL creates a resource URL for the Role with given ID in given context.
func RoleURL(organizationID, roleID string) string {
	return path.Join(rm.OrganizationURL(organizationID), KindRole, url.PathEscape(roleID))
}

// UserURL creates a resource URL for the User with given ID.
func UserURL(userID string) string {
	return "/" + path.Join(KindUser, url.PathEscape(userID))
}

// APIKeyURL creates a resource URL for the APIKey with given ID in given context.
func APIKeyURL(userID, apiKeyID string) string {
	return path.Join(UserURL(userID), KindAPIKey, url.PathEscape(apiKeyID))
}

//
// DISCLAIMER
//
// Copyright 2019 ArangoDB Inc, Cologne, Germany
//
// Author Ewout Prangsma
//

package v1

import (
	"net/url"
	"path"

	rm "github.com/arangodb-managed/apis/resourcemanager/v1"
)

const (
	// KindGroup is a constants for the kind of Group resources.
	KindGroup = "Group"
	// KindRole is a constants for the kind of Role resources.
	KindRole = "Role"
)

// GroupURL creates a resource URL for the Group with given ID in given context.
func GroupURL(organizationID, groupID string) string {
	return path.Join(rm.OrganizationURL(organizationID), KindGroup, url.PathEscape(groupID))
}

// RoleURL creates a resource URL for the Role with given ID in given context.
func RoleURL(organizationID, roleID string) string {
	return path.Join(rm.OrganizationURL(organizationID), KindRole, url.PathEscape(roleID))
}

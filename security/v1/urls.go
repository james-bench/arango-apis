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
	// KindIPWhitelist is a constants for the kind of IPWhitelist resources.
	KindIPWhitelist = "IPWhitelist"
)

// IPWhitelistURL creates a resource URL for the IPWhitelist with given ID
// in given context.
func IPWhitelistURL(organizationID, projectID, ipwhitelistID string) string {
	return path.Join(rm.ProjectURL(organizationID, projectID), KindIPWhitelist, url.PathEscape(ipwhitelistID))
}

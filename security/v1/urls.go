//
// DISCLAIMER
//
// Copyright 2019 ArangoDB GmbH, Cologne, Germany
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
	// KindIAMProvider is a constants for the kind of IAMProvider resources.
	KindIAMProvider = "IAMProvider"
)

// IPWhitelistURL creates a resource URL for the IPWhitelist with given ID
// in given context.
func IPWhitelistURL(organizationID, projectID, ipwhitelistID string) string {
	return path.Join(rm.ProjectURL(organizationID, projectID), KindIPWhitelist, url.PathEscape(ipwhitelistID))
}

// IAMProviderURL creates a resource URL for the IAMProvider with given ID
// in given context.
func IAMProviderURL(organizationID, projectID, iamproviderID string) string {
	return path.Join(rm.ProjectURL(organizationID, projectID), KindIAMProvider, url.PathEscape(iamproviderID))
}

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
	usageKind = "Usage"
)

// UsageURL creates a resource URL for the Usage with given ID
// in given context.
func UsageURL(organizationID, usageID string) string {
	return path.Join(rm.OrganizationURL(organizationID), usageKind, url.PathEscape(usageID))
}

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
	// KindUsageItem is a constants for the kind of UsageItem resources.
	KindUsageItem = "UsageItem"
)

// UsageItemURL creates a resource URL for the UsageItem with given ID
// in given context.
func UsageItemURL(organizationID, usageItemID string) string {
	return path.Join(rm.OrganizationURL(organizationID), KindUsageItem, url.PathEscape(usageItemID))
}

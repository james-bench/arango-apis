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
	// KindCACertificate is a constants for the kind of CACertificate resources.
	KindCACertificate = "CACertificate"
)

// CACertificateURL creates a resource URL for the CACertificate with given ID
// in given context.
func CACertificateURL(organizationID, projectID, cacertID string) string {
	return path.Join(rm.ProjectURL(organizationID, projectID), KindCACertificate, url.PathEscape(cacertID))
}

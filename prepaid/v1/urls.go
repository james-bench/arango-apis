package v1

import (
	"net/url"
	"path"

	rm "github.com/arangodb-managed/apis/resourcemanager/v1"
)

const (
	// KindPrepaidDeployment is a constant for the kind of PrepaidDeployment resources.
	KindPrepaidDeployment = "PrepaidDeployment"
)

// PrepaidDeploymentURL returns an url for given organization and prepaid deployment
func PrepaidDeploymentURL(organizationID, prepaidDeploymentID string) string {
	return path.Join("/", rm.OrganizationURL(organizationID), KindPrepaidDeployment, url.PathEscape(organizationID))
}

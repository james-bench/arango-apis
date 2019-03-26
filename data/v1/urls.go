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
	// KindDeployment is a constants for the kind of Deployment resources.
	KindDeployment = "Deployment"
)

// DeploymentURL creates a resource URL for the Deployment with given ID
// in given context.
func DeploymentURL(organizationID, projectID, deploymentID string) string {
	return path.Join(rm.ProjectURL(organizationID, projectID), KindDeployment, url.PathEscape(deploymentID))
}

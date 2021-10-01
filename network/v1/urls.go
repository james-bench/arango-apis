//
// DISCLAIMER
//
// Copyright 2021 ArangoDB GmbH, Cologne, Germany
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
// Author Robert Stam
//

package v1

import (
	"net/url"
	"path"

	data "github.com/arangodb-managed/apis/data/v1"
)

const (
	// KindPrivateEndpointService is a constants for the kind of PrivateEndpointService resources.
	KindPrivateEndpointService = "PrivateEndpointService"
)

// PrivateEndpointServiceURL creates a resource URL for the PrivateEndpointService with given ID
// in given context (as individual IDs).
func PrivateEndpointServiceURL(organizationID, projectID, deploymentID, privateEndpointServiceID string) string {
	return PrivateEndpointServiceURL2(data.DeploymentURL(organizationID, projectID, deploymentID), privateEndpointServiceID)
}

// PrivateEndpointServiceURL2 creates a resource URL for the PrivateEndpointService with given ID
// in given context (as base URL).
func PrivateEndpointServiceURL2(deploymentURL, privateEndpointServiceID string) string {
	return path.Join(deploymentURL, KindPrivateEndpointService, url.PathEscape(privateEndpointServiceID))
}

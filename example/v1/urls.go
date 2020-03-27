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
// Author Gergely Brautigam
//

package v1

import (
	"net/url"
	"path"

	data "github.com/arangodb-managed/apis/data/v1"

	rm "github.com/arangodb-managed/apis/resourcemanager/v1"
)

const (
	// KindExampleDataset is a constants for the kind of ExampleDataset resources.
	KindExampleDataset = "ExampleDataset"
	// KindExampleDatasetInstallation is a constants for the kind of ExampleDatasetInstallation resources.
	KindExampleDatasetInstallation = "ExampleDatasetInstallation"
)

// ExampleDatasetURL creates a resource URL for the ExampleDataset with given ID
// in given context (as individual IDs).
func ExampleDatasetURL(organizationID, exampleDatasetID string) string {
	return ExampleDatasetURL2(rm.OrganizationURL(organizationID), exampleDatasetID)
}

// ExampleDatasetURL2 creates a resource URL for the ExampleDataset with given ID
// in given context (as base URL).
func ExampleDatasetURL2(organizationURL, exampleDatasetID string) string {
	return path.Join(organizationURL, KindExampleDataset, url.PathEscape(exampleDatasetID))
}

// ExampleDatasetInstallationURL creates a resource URL for the ExampleDatasetInstallation with given ID
// in given context (as individual IDs).
func ExampleDatasetInstallationURL(organizationID, projectID, deploymentID, exampleInstallationID string) string {
	return ExampleDatasetInstallationURL2(data.DeploymentURL(organizationID, projectID, deploymentID), exampleInstallationID)
}

// ExampleDatasetInstallationURL2 creates a resource URL for the ExampleDatasetInstallation with given ID
// in given context (as base URL).
func ExampleDatasetInstallationURL2(deploymentURL, exampleInstallationID string) string {
	return path.Join(deploymentURL, KindExampleDatasetInstallation, url.PathEscape(exampleInstallationID))
}

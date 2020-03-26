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

const (
	// ExampleDataset permissions
	// PermissionExampleDatasetList is needed in order to be able to list example datasets.
	PermissionExampleDatasetList = "example.exampledataset.list"
	// PermissionExampleDatasetGet is needed to get a single example dataset.
	PermissionExampleDatasetGet = "example.exampledataset.get"
	// PermissionExampleDatasetCreate is needed to create an example dataset.
	PermissionExampleDatasetCreate = "example.exampledataset.create"
	// PermissionExampleDatasetDelete is needed to delete an example dataset.
	PermissionExampleDatasetDelete = "example.exampledataset.delete"
	// PermissionExampleDatasetUpdate is needed to update an example dataset.
	PermissionExampleDatasetUpdate = "example.exampledataset.update"
)

const (
	// ExampleDatasetInstallation permissions
	// PermissionExampleDatasetInstallationList is needed in order to be able to list example datasets.
	PermissionExampleDatasetInstallationList = "example.exampledatasetinstallation.list"
	// PermissionExampleDatasetInstallationGet is needed to get a single example dataset.
	PermissionExampleDatasetInstallationGet = "example.exampledatasetinstallation.get"
	// PermissionExampleDatasetInstallationCreate is needed to create an example dataset.
	PermissionExampleDatasetInstallationCreate = "example.exampledatasetinstallation.create"
	// PermissionExampleDatasetInstallationDelete is needed to delete an example dataset.
	PermissionExampleDatasetInstallationDelete = "example.exampledatasetinstallation.delete"
	// PermissionExampleDatasetInstallationUpdate is needed to update an example dataset.
	PermissionExampleDatasetInstallationUpdate = "example.exampledatasetinstallation.update"
)

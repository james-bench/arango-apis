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
	// Example installations event types

	// EventTypeExampleDatasetInstallationCreated is the type of event fired after an installations has been created
	// SubjectID contains the ID of the installation.
	EventTypeExampleDatasetInstallationCreated = "example.exampledatasetinstallation.created"
	// EventTypeExampleDatasetInstallationDeleted is the type of event fired after an installations has been deleted
	// SubjectID contains the ID of the installation.
	EventTypeExampleDatasetInstallationDeleted = "example.exampledatasetinstallation.deleted"
	// EventTypeExampleDatasetInstallationUpdated is the type of event fired after an installations has been updated
	// SubjectID contains the ID of the installation.
	EventTypeExampleDatasetInstallationUpdated = "example.exampledatasetinstallation.updated"
)

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

import common "github.com/arangodb-managed/apis/common/v1"

// Equals returns true when source & other have the same values
func (source *ExampleDataset) Equals(other *ExampleDataset) bool {
	return source.GetName() == other.GetName() &&
		source.GetDescription() == other.GetDescription() &&
		source.GetGuide() == other.GetGuide()
}

// Equals returns true when source & other have the same values
func (source *ExampleDatasetInstallation) Equals(other *ExampleDatasetInstallation) bool {
	return source.GetDeploymentId() == other.GetDeploymentId() &&
		source.GetExampledatasetId() == other.GetExampledatasetId() &&
		source.GetStatus().Equals(other.GetStatus())
}

// Equals returns true when source & other have the same values
func (source *ExampleDatasetInstallation_Status) Equals(other *ExampleDatasetInstallation_Status) bool {
	return source.GetState() == other.GetState() &&
		source.GetDatabaseName() == other.GetDatabaseName() &&
		source.GetIsAvailable() == other.GetIsAvailable() &&
		source.GetIsFailed() == other.GetIsFailed()
}

// CloneExampleDataset creates a deep copy of the given source
func (source *ExampleDataset) Clone() *ExampleDataset {
	if source == nil {
		return nil
	}
	clone := *source
	clone.CreatedAt = common.CloneTimestamp(source.GetCreatedAt())
	return &clone
}

// CloneExampleDatasetInstallation creates a deep copy of the given source
func (source *ExampleDatasetInstallation) Clone() *ExampleDatasetInstallation {
	if source == nil {
		return nil
	}
	clone := *source
	clone.CreatedAt = common.CloneTimestamp(source.GetCreatedAt())
	clone.DeletedAt = common.CloneTimestamp(source.GetDeletedAt())
	clone.Status = source.GetStatus().Clone()
	return &clone
}

// CloneExampleDatasetInstallationStatus creates a deep copy fo the given source
func (source *ExampleDatasetInstallation_Status) Clone() *ExampleDatasetInstallation_Status {
	if source == nil {
		return nil
	}
	clone := *source
	return &clone
}

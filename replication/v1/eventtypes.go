//
// DISCLAIMER
//
// Copyright 2022 ArangoDB GmbH, Cologne, Germany
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

package v1

const (
	// DeploymentReplication event types

	// EventTypeDeploymentReplicationCreated is the type of event fired after a DeploymentReplication has been created.
	// SubjectID contains the Deployment ID of the DeploymentReplication.
	EventTypeDeploymentReplicationCreated = "replication.deploymentreplication.created"
	// EventTypeDeploymentReplicationUpdated is the type of event fired after a DeploymentReplication has been updated.
	// SubjectID contains the Deployment ID of the DeploymentReplication.
	EventTypeDeploymentReplicationUpdated = "replication.deploymentreplication.updated"
	// EventTypeDeploymentReplicationDeleted is the type of event fired after a DeploymentReplication has been deleted.
	// SubjectID contains the Deployment ID of the DeploymentReplication.
	EventTypeDeploymentReplicationDeleted = "replication.deploymentreplication.deleted"
)

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
//

package v1

const (
	// Deployment permissions

	// PermissionDeploymentCloneFromBackup is needed to clone a deployment from an existing backup.
	PermissionDeploymentCloneFromBackup = "replication.deployment.clone-from-backup"
)

const (
	// PermissionGetDeploymentReplication is needed to get a DeploymentReplication for a given Deployment
	PermissionGetDeploymentReplication = "replication.deploymentreplication.get"

	// PermissionUpdateDeploymentReplication is needed to update / create a DeploymentReplication
	PermissionUpdateDeploymentReplication = "replication.deploymentreplication.update"
)

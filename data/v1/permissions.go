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
// Author Ewout Prangsma
//

package v1

const (
	// Deployment permissions

	// PermissionDeploymentList is needed for listing deployments in a project
	PermissionDeploymentList = "data.deployment.list"
	// PermissionDeploymentGet is needed for fetching an individual deployments in a project
	PermissionDeploymentGet = "data.deployment.get"
	// PermissionDeploymentCreate is needed for create a deployment
	PermissionDeploymentCreate = "data.deployment.create"
	// PermissionDeploymentUpdate is needed for updating a deployment
	PermissionDeploymentUpdate = "data.deployment.update"
	// PermissionDeploymentDelete is needed for deleting a deployment
	PermissionDeploymentDelete = "data.deployment.delete"
	// PermissionDeploymentRestoreBackup is needed for restoring a backup of a deployment
	PermissionDeploymentRestoreBackup = "data.deployment.restore-backup"
	// PermissionDeploymentResume is needed for resuming a paused deployment
	PermissionDeploymentResume = "data.deployment.resume"
	// PermissionDeploymentRotateServer is needed to rotate a server belonging to a deployment
	PermissionDeploymentRotateServer = "data.deployment.rotate-server"
	// PermissionDeploymentCreateTestDatabase is needed for creating a test database
	PermissionDeploymentCreateTestDatabase = "data.deployment.create-test-database"
	// PermissionDeploymentRebalanceShards is needed for rebalancing shards for a deployment
	PermissionDeploymentRebalanceShards = "data.deployment.rebalance-shards"
	// PermissionDeploymentUpdateScheduledRootPasswordRotation is needed to update scheduled root password rotation setting for a deployment
	PermissionDeploymentUpdateScheduledRootPasswordRotation = "data.deployment.update-scheduled-root-password-rotation"
)

const (
	// DeploymentCredentials permissions

	// PermissionDeploymentCredentialsGet is needed for fetching credentials of an individual deployments in a project
	PermissionDeploymentCredentialsGet = "data.deploymentcredentials.get"
)

const (
	// DeploymentFeatures permissions

	// PermissionDeploymentFeaturesGet is needed for fetching features that are available to deployments in a project
	PermissionDeploymentFeaturesGet = "data.deploymentfeatures.get"
)

const (
	// ServerSpecLimits permissions

	// PermissionLimitsGet is needed for fetching server spec limits in a project
	PermissionLimitsGet = "data.limits.get"
)

const (
	// ServerSpecPresets permissions

	// PermissionPresetsList is needed for listing server spec presets in a project
	PermissionPresetsList = "data.presets.list"
)

const (
	// NodeSize permissions

	// PermissionNodeSizeList is needed for listing server spec presets in a project
	PermissionNodeSizeList = "data.nodesize.list"
)

const (
	// DeploymentPrice permissions

	// PermissionDeploymentPriceCalculate is needed for calculating deployment prices
	PermissionDeploymentPriceCalculate = "data.deploymentprice.calculate"
)

const (
	// DeploymentModel permissions

	// PermissionDeploymentModelList is needed for listing deployment models in a project.
	PermissionDeploymentModelList = "data.deploymentmodel.list"
)

const (
	// CPUSize permissions

	// PermissionCPUSizeList is needed for listing CPU sizes in a project.
	PermissionCPUSizeList = "data.cpusize.list"
)

const (
	// DiskPerformance permissions

	// PermissionDiskPerformanceList is needed for listing disk performances for a specific deployment
	PermissionDiskPerformanceList = "data.diskperformance.list"
)

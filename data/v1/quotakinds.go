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
	// Deployment related quota kinds

	// QuotaKindDeploymentsPerProject limits the number of deployments in a single project.
	// This kinds of quota must be requested on project level.
	QuotaKindDeploymentsPerProject = "data.deployments-per-project"
	// QuotaKindDeploymentCoordinatorMemory limits the amount of memory (in GB) used
	// by a single coordinator in a deployment.
	// This kinds of quota must be requested on project level.
	QuotaKindDeploymentCoordinatorMemory = "data.deployment-coordinator-memory"
	// QuotaKindDeploymentCoordinators limits the number of coordinators per deployment.
	// This kinds of quota must be requested on project level.
	QuotaKindDeploymentCoordinators = "data.deployment-coordinators"
	// QuotaKindDeploymentDBServerMemory limits the amount of memory (in GB) used
	// by a single dbserver in a deployment.
	// This kinds of quota must be requested on project level.
	QuotaKindDeploymentDBServerMemory = "data.deployment-dbserver-memory"
	// QuotaKindDeploymentDBServerDisk limits the amount of disk (in GB) used
	// by a single dbserver in a deployment.
	// This kinds of quota must be requested on project level.
	QuotaKindDeploymentDBServerDisk = "data.deployment-dbserver-disk"
	// QuotaKindDeploymentDBServers limits the number of dbservers per deployment.
	// This kinds of quota must be requested on project level.
	QuotaKindDeploymentDBServers = "data.deployment-dbservers"
	// QuotaKindTotalDeploymentsMemoryPerProject limits the total amount of memory used
	// by all deployments in a project in GB.
	// This kinds of quota must be requested on project level.
	QuotaKindTotalDeploymentsMemoryPerProject = "data.total-deployments-memory-per-project"
	// QuotaKindTotalDeploymentsDiskPerProject limits the total amount of disk used
	// by all deployments in a project in GB.
	// This kinds of quota must be requested on project level.
	QuotaKindTotalDeploymentsDiskPerProject = "data.total-deployments-disk-per-project"
	// QuotaKindTotalBackupPoliciesPerDeployment limits the total number of back up policies
	// for a given deployment
	// This kind of quota must be requested on project level.
	QuotaKindTotalBackupPoliciesPerDeployment = "backup.total-backup-policy-per-deployment"
)

//
// DISCLAIMER
//
// Copyright 2019 ArangoDB GmbH, Cologne, Germany
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
)

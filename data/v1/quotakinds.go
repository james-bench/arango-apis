//
// DISCLAIMER
//
// Copyright 2019 ArangoDB Inc, Cologne, Germany
//
// Author Ewout Prangsma
//

package v1

const (
	// Deployment related quota kinds

	// QuotaKindDeploymentsPerProject limits the number of deployments in a single project.
	// This kinds of quota must be requested on project level.
	QuotaKindDeploymentsPerProject = "data.deployments-per-project"
	// QuotaKindTotalDeploymentsMemoryPerProject limits the total amount of memory used
	// by all deployments in a project in GB.
	// This kinds of quota must be requested on project level.
	QuotaKindTotalDeploymentsMemoryPerProject = "data.total-deployments-memory-per-project"
	// QuotaKindTotalDeploymentsDiskPerProject limits the total amount of disk used
	// by all deployments in a project in GB.
	// This kinds of quota must be requested on project level.
	QuotaKindTotalDeploymentsDiskPerProject = "data.total-deployments-disk-per-project"
)

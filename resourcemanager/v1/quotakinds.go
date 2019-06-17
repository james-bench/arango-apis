//
// DISCLAIMER
//
// Copyright 2019 ArangoDB Inc, Cologne, Germany
//
// Author Ewout Prangsma
//

package v1

const (
	// Quota kinds

	// QuotaKindProjectsPerOrganization limits the number of projects in a single organization.
	// This kinds of quota must be requested on organization level.
	QuotaKindProjectsPerOrganization = "resourcemanager.projects-per-organization"
	// QuotaKindTotalDeploymentsPerTierTemplate is a go-template for a quota kind
	// tha limits the total number of deployments (over the lifetime on the organization)
	// in a specific tier.
	// The template variable is the tier-id.
	// This kinds of quota must be requested on organization level.
	QuotaKindTotalDeploymentsPerTierTemplate = "resourcemanager.total-deployment-%s"
)

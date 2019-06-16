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
)

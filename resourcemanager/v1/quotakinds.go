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

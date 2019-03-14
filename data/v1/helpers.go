//
// DISCLAIMER
//
// Copyright 2019 ArangoDB Inc, Cologne, Germany
//
// Author Ewout Prangsma
//

package v1

import "strings"

// DeploymentStatusEqual returns true when the fields of a & b are equal.
func DeploymentStatusEqual(a, b *Deployment_Status) bool {
	return a.GetEndpoint() == b.GetEndpoint() &&
		a.GetDescription() == b.GetDescription() &&
		a.GetCreated() == b.GetCreated() &&
		a.GetReady() == b.GetReady() &&
		a.GetUpgrading() == b.GetUpgrading() &&
		strings.Join(a.GetServerVersions(), ",") == strings.Join(b.GetServerVersions(), ",")
}

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

// GetOrCreateCertificates returns the Certificates field, creating it if needed.
func (d *Deployment) GetOrCreateCertificates() *Deployment_CertificateSpec {
	if d.GetCertificates() == nil {
		d.Certificates = &Deployment_CertificateSpec{}
	}
	return d.GetCertificates()
}

// GetOrCreateExpiration returns the Expiration field, creating it if needed.
func (d *Deployment) GetOrCreateExpiration() *Deployment_Expiration {
	if d.GetExpiration() == nil {
		d.Expiration = &Deployment_Expiration{}
	}
	return d.GetExpiration()
}

// GetOrCreateServers returns the Servers field, creating it if needed.
func (d *Deployment) GetOrCreateServers() *Deployment_ServersSpec {
	if d.GetServers() == nil {
		d.Servers = &Deployment_ServersSpec{}
	}
	return d.GetServers()
}

// GetOrCreateStatus returns the Status field, creating it if needed.
func (d *Deployment) GetOrCreateStatus() *Deployment_Status {
	if d.GetStatus() == nil {
		d.Status = &Deployment_Status{}
	}
	return d.GetStatus()
}

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
		strings.Join(a.GetServerVersions(), ",") == strings.Join(b.GetServerVersions(), ",") &&
		DeploymentServerStatusListEqual(a.GetServers(), b.GetServers())
}

// DeploymentServerStatusListEqual returns true when the elements of a & b are equal.
func DeploymentServerStatusListEqual(a, b []*Deployment_ServerStatus) bool {
	if len(a) != len(b) {
		return false
	}
	for i, x := range a {
		if !DeploymentServerStatusEqual(x, b[i]) {
			return false
		}
	}
	return true
}

// DeploymentServerStatusEqual returns true when the fields of a & b are equal.
func DeploymentServerStatusEqual(a, b *Deployment_ServerStatus) bool {
	return a.GetId() == b.GetId() &&
		a.GetType() == b.GetType() &&
		a.GetDescription() == b.GetDescription() &&
		a.GetReady() == b.GetReady() &&
		a.GetMemberOfCluster() == b.GetMemberOfCluster() &&
		a.GetFailed() == b.GetFailed() &&
		DataVolumeInfoEqual(a.GetDataVolumeInfo(), b.GetDataVolumeInfo())
}

// DataVolumeInfoEqual returns true when the fields of a & b are equal.
func DataVolumeInfoEqual(a, b *DataVolumeInfo) bool {
	return a.GetAvailableBytes() == b.GetAvailableBytes() &&
		a.GetMeasuredAt() == b.GetMeasuredAt() &&
		a.GetTotalBytes() == b.GetTotalBytes() &&
		a.GetUsedBytes() == b.GetUsedBytes()
}

// GetOrCreateAuthentication returns the Authentication field, creating it if needed.
func (d *Deployment) GetOrCreateAuthentication() *Deployment_AuthenticationSpec {
	if d.GetAuthentication() == nil {
		d.Authentication = &Deployment_AuthenticationSpec{}
	}
	return d.GetAuthentication()
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

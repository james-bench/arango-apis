//
// DISCLAIMER
//
// Copyright 2019 ArangoDB Inc, Cologne, Germany
//
// Author Ewout Prangsma
//

package v1

import "strings"

// SpecEquals returns true when source & other have the same specification values
func (source *Deployment) SpecEquals(other *Deployment) bool {
	return source.GetVersion() == other.GetVersion() &&
		source.GetRegionId() == other.GetRegionId() &&
		source.GetIpwhitelistId() == other.GetIpwhitelistId() &&
		source.GetCertificates().Equals(other.GetCertificates()) &&
		source.GetServers().Equals(other.GetServers())
}

// Equals returns true when source & other have the same values
func (source *Deployment_CertificateSpec) Equals(other *Deployment_CertificateSpec) bool {
	return source.GetCaCertificateId() == other.GetCaCertificateId() &&
		strings.Join(source.GetAlternateDnsNames(), ",") == strings.Join(other.GetAlternateDnsNames(), ",")
}

// Equals returns true when source & other have the same values
func (source *Deployment_ServersSpec) Equals(other *Deployment_ServersSpec) bool {
	return source.GetCoordinators() == other.GetCoordinators() &&
		source.GetCoordinatorMemorySize() == other.GetCoordinatorMemorySize() &&
		strings.Join(source.GetCoordinatorArgs(), ",") == strings.Join(other.GetCoordinatorArgs(), ",") &&
		source.GetDbservers() == other.GetDbservers() &&
		source.GetDbserverDiskSize() == other.GetDbserverDiskSize() &&
		source.GetDbserverMemorySize() == other.GetDbserverMemorySize() &&
		strings.Join(source.GetDbserverArgs(), ",") == strings.Join(other.GetDbserverArgs(), ",")
}

// DeploymentStatusEqual returns true when the fields of a & b are equal.
func DeploymentStatusEqual(a, b *Deployment_Status, ignoreTimestamps bool) bool {
	return a.GetEndpoint() == b.GetEndpoint() &&
		a.GetDescription() == b.GetDescription() &&
		a.GetCreated() == b.GetCreated() &&
		a.GetReady() == b.GetReady() &&
		a.GetUpgrading() == b.GetUpgrading() &&
		strings.Join(a.GetServerVersions(), ",") == strings.Join(b.GetServerVersions(), ",") &&
		DeploymentServerStatusListEqual(a.GetServers(), b.GetServers(), ignoreTimestamps)
}

// DeploymentServerStatusListEqual returns true when the elements of a & b are equal.
func DeploymentServerStatusListEqual(a, b []*Deployment_ServerStatus, ignoreTimestamps bool) bool {
	if len(a) != len(b) {
		return false
	}
	for i, x := range a {
		if !DeploymentServerStatusEqual(x, b[i], ignoreTimestamps) {
			return false
		}
	}
	return true
}

// DeploymentServerStatusEqual returns true when the fields of a & b are equal.
func DeploymentServerStatusEqual(a, b *Deployment_ServerStatus, ignoreTimestamps bool) bool {
	return a.GetId() == b.GetId() &&
		a.GetType() == b.GetType() &&
		a.GetDescription() == b.GetDescription() &&
		a.GetReady() == b.GetReady() &&
		a.GetMemberOfCluster() == b.GetMemberOfCluster() &&
		a.GetFailed() == b.GetFailed() &&
		DataVolumeInfoEqual(a.GetDataVolumeInfo(), b.GetDataVolumeInfo(), ignoreTimestamps)
}

// DataVolumeInfoEqual returns true when the fields of a & b are equal.
func DataVolumeInfoEqual(a, b *DataVolumeInfo, ignoreTimestamps bool) bool {
	return a.GetAvailableBytes() == b.GetAvailableBytes() &&
		(ignoreTimestamps || (a.GetMeasuredAt().Equal(b.GetMeasuredAt()))) &&
		a.GetTotalBytes() == b.GetTotalBytes() &&
		a.GetUsedBytes() == b.GetUsedBytes()
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

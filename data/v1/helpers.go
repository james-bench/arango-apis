//
// DISCLAIMER
//
// Copyright 2020-2021 ArangoDB GmbH, Cologne, Germany
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

import (
	"strings"

	common "github.com/arangodb-managed/apis/common/v1"
	rm "github.com/arangodb-managed/apis/resourcemanager/v1"
)

// GetOrganizationID returns the organization ID of a deployment based on the URL.
func (depl *Deployment) GetOrganizationID() (string, error) {
	resURL, err := rm.ParseResourceURL(depl.GetUrl())
	if err != nil {
		return "", err
	}
	return resURL.OrganizationID(), nil
}

// Clone creates a deep clone of the given source
func (s *Deployment_Status) Clone() *Deployment_Status {
	if s == nil {
		return nil
	}
	clone := *s
	clone.ServerVersions = append([]string{}, s.GetServerVersions()...)
	clone.BootstrappedAt = common.CloneTimestamp(s.GetBootstrappedAt())
	clone.Servers = make([]*Deployment_ServerStatus, 0, len(s.GetServers()))
	for _, x := range s.GetServers() {
		clone.Servers = append(clone.Servers, x.Clone())
	}
	return &clone
}

// Clone creates a deep clone of the given source
func (s *Deployment_ServerStatus) Clone() *Deployment_ServerStatus {
	if s == nil {
		return nil
	}
	clone := *s
	clone.CreatedAt = common.CloneTimestamp(s.GetCreatedAt())
	clone.DataVolumeInfo = s.GetDataVolumeInfo().Clone()
	return &clone
}

// Clone creates a deep clone of the given source
func (s *DataVolumeInfo) Clone() *DataVolumeInfo {
	if s == nil {
		return nil
	}
	clone := *s
	clone.MeasuredAt = common.CloneTimestamp(s.GetMeasuredAt())
	return &clone
}

// SpecEquals returns true when source & other have the same specification values
func (source *Deployment) SpecEquals(other *Deployment) bool {
	return source.GetVersion() == other.GetVersion() &&
		source.GetRegionId() == other.GetRegionId() &&
		source.GetIpallowlistId() == other.GetIpallowlistId() &&
		source.GetCertificates().Equals(other.GetCertificates()) &&
		source.GetServers().Equals(other.GetServers()) &&
		source.GetBackupRestore().Equals(other.GetBackupRestore()) &&
		source.GetLocked() == other.GetLocked()
}

// Equals returns true when source & other have the same values
func (source *Deployment_BackupRestoreSpec) Equals(other *Deployment_BackupRestoreSpec) bool {
	return source.GetRevision() == other.GetRevision() &&
		source.GetLastUpdatedAt().Equal(other.GetLastUpdatedAt()) &&
		source.GetBackupId() == other.GetBackupId()
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

// Equals returns true when source and other have the same values
func (source *Deployment_ModelSpec) Equals(other *Deployment_ModelSpec) bool {
	return source.GetModel() == other.GetModel() &&
		source.GetNodeCount() == other.GetNodeCount() &&
		source.GetNodeDiskSize() == other.GetNodeDiskSize() &&
		source.GetNodeSizeId() == other.GetNodeSizeId()
}

// DeploymentStatusEqual returns true when the fields of a & b are equal.
func DeploymentStatusEqual(a, b *Deployment_Status, ignoreTimestamps, ignoreVolatile bool) bool {
	return a.GetEndpoint() == b.GetEndpoint() &&
		a.GetEndpointSelfSigned() == b.GetEndpointSelfSigned() &&
		a.GetDescription() == b.GetDescription() &&
		a.GetCreated() == b.GetCreated() &&
		a.GetReady() == b.GetReady() &&
		a.GetUpgrading() == b.GetUpgrading() &&
		a.GetBootstrapped() == b.GetBootstrapped() &&
		strings.Join(a.GetServerVersions(), ",") == strings.Join(b.GetServerVersions(), ",") &&
		DeploymentServerStatusListEqual(a.GetServers(), b.GetServers(), ignoreTimestamps, ignoreVolatile) &&
		a.GetBackupRestoreStatus().Equals(b.GetBackupRestoreStatus()) &&
		a.GetTotalBackupSizeBytes() == b.GetTotalBackupSizeBytes() &&
		a.GetBackupUploadInProgress() == b.GetBackupUploadInProgress()
}

// Equals returns true when source & other have the same values
func (source *Deployment_BackupRestoreStatus) Equals(other *Deployment_BackupRestoreStatus) bool {
	return source.GetRevision() == other.GetRevision() &&
		source.GetRestoring() == other.GetRestoring() &&
		source.GetStatus() == other.GetStatus() &&
		source.GetFailureReason() == other.GetFailureReason()
}

// DeploymentServerStatusListEqual returns true when the elements of a & b are equal.
func DeploymentServerStatusListEqual(a, b []*Deployment_ServerStatus, ignoreTimestamps, ignoreVolatile bool) bool {
	if len(a) != len(b) {
		return false
	}
	for i, x := range a {
		if !DeploymentServerStatusEqual(x, b[i], ignoreTimestamps, ignoreVolatile) {
			return false
		}
	}
	return true
}

// DeploymentServerStatusEqual returns true when the fields of a & b are equal.
func DeploymentServerStatusEqual(a, b *Deployment_ServerStatus, ignoreTimestamps, ignoreVolatile bool) bool {
	return a.GetId() == b.GetId() &&
		a.GetType() == b.GetType() &&
		a.GetDescription() == b.GetDescription() &&
		a.GetReady() == b.GetReady() &&
		a.GetMemberOfCluster() == b.GetMemberOfCluster() &&
		a.GetFailed() == b.GetFailed() &&
		a.GetVersion() == b.GetVersion() &&
		a.GetLastStartedAt().Equal(b.GetLastStartedAt()) &&
		a.GetRotationPending() == b.GetRotationPending() &&
		(ignoreVolatile || DataVolumeInfoEqual(a.GetDataVolumeInfo(), b.GetDataVolumeInfo(), ignoreTimestamps)) &&
		a.GetRecentRestarts() == b.GetRecentRestarts() &&
		(ignoreVolatile || a.GetLastMemoryUsage() == b.GetLastMemoryUsage()) &&
		(ignoreVolatile || a.GetLastCpuUsage() == b.GetLastCpuUsage()) &&
		a.GetIsLeader() == b.GetIsLeader()
}

// DataVolumeInfoEqual returns true when the fields of a & b are equal.
func DataVolumeInfoEqual(a, b *DataVolumeInfo, ignoreTimestamps bool) bool {
	return a.GetAvailableBytes() == b.GetAvailableBytes() &&
		a.GetTotalBytes() == b.GetTotalBytes() &&
		a.GetUsedBytes() == b.GetUsedBytes() &&
		a.GetAvailableInodes() == b.GetAvailableInodes() &&
		a.GetTotalInodes() == b.GetTotalInodes() &&
		a.GetUsedInodes() == b.GetUsedInodes() &&
		(ignoreTimestamps || (a.GetMeasuredAt().Equal(b.GetMeasuredAt())))

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

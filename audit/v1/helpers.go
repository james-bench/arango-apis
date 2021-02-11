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
// Author Gergely Brautigam
// Author Robert Stam
//

package v1

import (
	"sort"
	"strings"
)

// Equals returns true when source & other have the same values, the destination statuses are compared depending of the provided bool.
func (source *AuditLog) Equals(other *AuditLog, includeStatuses bool) bool {
	return source.GetId() == other.GetId() &&
		source.GetUrl() == other.GetUrl() &&
		source.GetName() == other.GetName() &&
		source.GetDescription() == other.GetDescription() &&
		source.GetCreatedAt().Equal(other.GetCreatedAt()) &&
		source.GetDeletedAt().Equal(other.GetDeletedAt()) &&
		source.GetCreatedById() == other.GetCreatedById() &&
		source.GetOrganizationId() == other.GetOrganizationId() &&
		source.GetCreatedById() == other.GetCreatedById() &&
		source.GetIsDefault() == other.GetIsDefault() &&
		auditLogDestinationList(source.GetDestinations()).Equals(other.GetDestinations(), includeStatuses)
}

type auditLogDestinationList []*AuditLog_Destination

// Equals returns true when source & other have the same values, the statuses are compared depending of the provided bool.
func (source auditLogDestinationList) Equals(other auditLogDestinationList, includeStatuses bool) bool {
	if len(source) != len(other) {
		return false
	}
	for i, x := range source {
		if !x.Equals(other[i], includeStatuses) {
			return false
		}
	}
	return true
}

// Equals returns true when source & other have the same values, the statuses are compared depending of the provided bool.
func (source *AuditLog_Destination) Equals(other *AuditLog_Destination, includeStatuses bool) bool {
	return source.GetType() == other.GetType() &&
		strings.Join(source.GetExcludedTopics(), ",") == strings.Join(other.GetExcludedTopics(), ",") &&
		source.GetHttpPost().Equals(other.GetHttpPost()) &&
		(!includeStatuses || auditLogDestinationStatusList(source.GetStatuses()).Equals(other.Statuses))
}

// Equals returns true when source & other have the same values.
func (source *AuditLog_HttpsPostSettings) Equals(other *AuditLog_HttpsPostSettings) bool {
	return source.GetClientCertificatePem() == other.GetClientCertificatePem() &&
		source.GetClientKeyPem() == other.GetClientKeyPem() &&
		source.GetTrustedServerCaPem() == other.GetTrustedServerCaPem() &&
		HeaderListEquals(source.GetHeaders(), other.GetHeaders())
}

type auditLogDestinationStatusList []*AuditLog_DestinationStatus

// Equals returns true when source & other have the same values
func (source auditLogDestinationStatusList) Equals(other auditLogDestinationStatusList) bool {
	if len(source) != len(other) {
		return false
	}
	for i, x := range source {
		if !x.Equals(other[i]) {
			return false
		}
	}
	return true
}

// Equals returns true when source & other have the same values.
func (source *AuditLog_DestinationStatus) Equals(other *AuditLog_DestinationStatus) bool {
	return source.GetDeploymentId() == other.GetDeploymentId() &&
		source.GetHasErrors() == other.GetHasErrors() &&
		source.GetErrorDetails() == other.GetErrorDetails() &&
		source.GetCountersSinceMidnight().Equals(other.GetCountersSinceMidnight()) &&
		source.GetCountersYesterday().Equals(other.GetCountersYesterday()) &&
		source.GetUpdatedAt().Equal(other.GetUpdatedAt())
}

// Equals returns true when source & other have the same values.
func (source *AuditLog_DestinationCounters) Equals(other *AuditLog_DestinationCounters) bool {
	return source.GetEvents() == other.GetEvents() &&
		source.GetEventsExcluded() == other.GetEventsExcluded() &&
		source.GetEventsUndeliverable() == other.GetEventsUndeliverable() &&
		source.GetBytesSucceeded() == other.GetBytesSucceeded() &&
		source.GetBytesFailed() == other.GetBytesFailed() &&
		source.GetHttpsPostsSucceeded() == other.GetHttpsPostsSucceeded() &&
		source.GetHttpsPostsFailed() == other.GetHttpsPostsFailed()
}

// Equals returns true when source & other have the same values.
func (source *AuditLog_Header) Equals(other *AuditLog_Header) bool {
	return source.GetKey() == other.GetKey() && source.GetValue() == other.GetValue()
}

// HeaderListEquals compares two arbitrary ordered header slices.
func HeaderListEquals(source []*AuditLog_Header, other []*AuditLog_Header) bool {
	if len(source) != len(other) {
		return false
	}
	a := make([]*AuditLog_Header, len(source))
	copy(a, source)
	b := make([]*AuditLog_Header, len(other))
	copy(b, other)
	// We don't have guarantee that the header slice will have the same order
	// so we must order the slice to compare.
	sort.SliceStable(a, func(i, j int) bool {
		return a[i].GetKey() < a[j].GetKey()
	})
	sort.SliceStable(b, func(i, j int) bool {
		return b[i].GetKey() < b[j].GetKey()
	})

	for i := range a {
		if !a[i].Equals(b[i]) {
			return false
		}
	}
	return true
}

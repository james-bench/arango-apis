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
// Author Gergely Brautigam
//

package v1

import (
	"sort"
	"strings"
)

// Equals returns true when source & other have the same values.
func (source *AuditLog_Destination) Equals(other *AuditLog_Destination) bool {
	return source.GetType() == other.GetType() &&
		source.GetHttpPost().Equals(other.GetHttpPost())
}

// Equals returns true when source & other have the same values.
func (source *AuditLog_HttpsPostSettings) Equals(other *AuditLog_HttpsPostSettings) bool {
	return source.GetClientCertificatePem() == other.GetClientCertificatePem() &&
		source.GetClientKeyPem() == other.GetClientKeyPem() &&
		strings.Join(source.GetExcludedTopics(), ",") == strings.Join(other.GetExcludedTopics(), ",") &&
		source.GetTrustedServerCaPem() == other.GetTrustedServerCaPem() &&
		HeaderListEquals(source.GetHeaders(), other.GetHeaders())
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

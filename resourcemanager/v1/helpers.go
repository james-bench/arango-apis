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

// Contains returns true if the given quota list contains a quota of given kind.
func (list *QuotaList) Contains(kind string) bool {
	for _, q := range list.GetItems() {
		if q.GetKind() == kind {
			return true
		}
	}
	return false
}

// Get returns the first quota with given kind from the given list or nil if no
// such quota exists.
func (list *QuotaList) Get(kind string) *Quota {
	for _, q := range list.GetItems() {
		if q.GetKind() == kind {
			return q
		}
	}
	return nil
}

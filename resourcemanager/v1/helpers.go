//
// DISCLAIMER
//
// Copyright 2019 ArangoDB GmbH, Cologne, Germany
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

//
// DISCLAIMER
//
// Copyright 2022 ArangoDB GmbH, Cologne, Germany
//

package v1

import "reflect"

// Equals returns true when source and other have the same values.
func (source *Status) Equals(other *Status) bool {
	return source.GetEndpoint() == other.GetEndpoint() &&
		source.GetPhase() == other.GetPhase() &&
		source.GetMessage() == other.GetMessage() &&
		source.GetLastUpdatedAt().Equal(other.GetLastUpdatedAt()) &&
		reflect.DeepEqual(source.GetUsage(), other.GetUsage())
}

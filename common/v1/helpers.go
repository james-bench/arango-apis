//
// DISCLAIMER
//
// Copyright 2019 ArangoDB Inc, Cologne, Germany
//
// Author Ewout Prangsma
//

package v1

import (
	types "github.com/gogo/protobuf/types"
)

// CloneTimestamp creates a deep clone of the given timestamp
func CloneTimestamp(s *types.Timestamp) *types.Timestamp {
	if s == nil {
		return nil
	}
	clone := *s
	return &clone
}

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

// CloneDuration creates a deep copy of the given duration
func CloneDuration(s *types.Duration) *types.Duration {
	if s == nil {
		return nil
	}
	clone := *s
	return &clone
}

// CloneOrDefault creates a clone of the given options if not nil,
// or creates an empty ListOptions.
// In either case, if current pageSize is zero, it is set to the given
// default page size.
func (opts *ListOptions) CloneOrDefault(defaultPageSize int32) *ListOptions {
	if opts == nil {
		opts = &ListOptions{}
	} else {
		clone := *opts
		opts = &clone
	}
	if opts.PageSize == 0 {
		opts.PageSize = defaultPageSize
	}
	return opts
}

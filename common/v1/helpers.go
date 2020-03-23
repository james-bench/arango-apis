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

import (
	types "github.com/gogo/protobuf/types"
)

const (
	// DefaultPageSize is the default number of items per List request.
	DefaultPageSize = 50
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
// default page size or DefaultPageSize is not default page size is given.
func (opts *ListOptions) CloneOrDefault(defaultPageSize ...int32) *ListOptions {
	if opts == nil {
		opts = &ListOptions{}
	} else {
		clone := *opts
		opts = &clone
	}
	if opts.PageSize == 0 {
		if len(defaultPageSize) > 0 {
			opts.PageSize = defaultPageSize[0]
		} else {
			opts.PageSize = DefaultPageSize
		}
	}
	return opts
}

//
// DISCLAIMER
//
// Copyright 2021 ArangoDB GmbH, Cologne, Germany
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
// Author Marcin Swiderski
//

package v1

import (
	"testing"
	"time"

	types "github.com/gogo/protobuf/types"
)

func TestIsActivePrepaidDeployment(t *testing.T) {
	now := time.Now()
	tests := []struct {
		name     string
		item     UsageItem
		expected bool
	}{
		{
			name:     "empty resource",
			item:     UsageItem{},
			expected: false,
		},
		{
			name: "empty prepaid deployment",
			item: UsageItem{
				Resource: &UsageItem_Resource{
					PrepaidDeploymentId: "",
				},
			},
			expected: false,
		},
		{
			name: "empty prepaid deployment",
			item: UsageItem{
				Resource: &UsageItem_Resource{
					PrepaidDeploymentId: "",
				},
			},
			expected: false,
		},
		{
			name: "empty dates",
			item: UsageItem{
				Resource: &UsageItem_Resource{
					PrepaidDeploymentId: "id",
				},
			},
			expected: false,
		},
		{
			//----------------------ppppppppppppppppp-
			//-------uuuuuuuuuuuuu--------------------
			name: "usage ends before prepaid deployment starts",
			item: UsageItem{
				Resource: &UsageItem_Resource{
					PrepaidDeploymentId:       "id",
					PrepaidDeploymentStartsAt: TimestampMust(now.Add(time.Hour)),
					PrepaidDeploymentEndsAt:   TimestampMust(now.Add(time.Hour * 24)),
				},
				StartsAt: TimestampMust(now.Add(-time.Hour * 24)),
				EndsAt:   TimestampMust(now),
			},
			expected: false,
		},
		{
			//-------ppppppppppppp--------------------
			//----------------------uuuuuuuuuuuuuuuuu-
			name: "usage starts after prepaid deployment ends",
			item: UsageItem{
				Resource: &UsageItem_Resource{
					PrepaidDeploymentId:       "id",
					PrepaidDeploymentStartsAt: TimestampMust(now.Add(-time.Hour * 24)),
					PrepaidDeploymentEndsAt:   TimestampMust(now),
				},
				StartsAt: TimestampMust(now.Add(time.Hour)),
				EndsAt:   TimestampMust(now.Add(time.Hour * 24)),
			},
			expected: false,
		},
		{
			//-------ppppppppppppp--------------------
			//--------------------uuuuuuuuuuuuuuuuuuu-
			name: "usage starts exactly at prepaid deployment ends",
			item: UsageItem{
				Resource: &UsageItem_Resource{
					PrepaidDeploymentId:       "id",
					PrepaidDeploymentStartsAt: TimestampMust(now.Add(-time.Hour * 24)),
					PrepaidDeploymentEndsAt:   TimestampMust(now),
				},
				StartsAt: TimestampMust(now),
				EndsAt:   TimestampMust(now.Add(time.Hour * 24)),
			},
			expected: false,
		},
		{
			//----------------------ppppppppppppppp---
			//-------uuuuuuuuuuuuuuu------------------
			name: "usage ends exactly at prepaid deployment start",
			item: UsageItem{
				Resource: &UsageItem_Resource{
					PrepaidDeploymentId:       "id",
					PrepaidDeploymentStartsAt: TimestampMust(now),
					PrepaidDeploymentEndsAt:   TimestampMust(now.Add(time.Hour * 24)),
				},
				StartsAt: TimestampMust(now.Add(-time.Hour * 48)),
				EndsAt:   TimestampMust(now),
			},
			expected: false,
		},
		{
			//-------ppppppppppppp--------------------
			//----------------uuuuuuuuuuuuuuuuu-------
			name: "usage starts before prepaid deployment ends",
			item: UsageItem{
				Resource: &UsageItem_Resource{
					PrepaidDeploymentId:       "id",
					PrepaidDeploymentStartsAt: TimestampMust(now.Add(-time.Hour * 24)),
					PrepaidDeploymentEndsAt:   TimestampMust(now.Add(time.Hour)),
				},
				StartsAt: TimestampMust(now.Add(-time.Hour)),
				EndsAt:   TimestampMust(now.Add(time.Hour * 24)),
			},
			expected: true,
		},
		{
			//-----ppppppppppppppppppppppp------------
			//--------uuuuuuuuuuuuuuuuu---------------
			name: "usage starts after prepaid deployment start & ends before prepaid deploment ends",
			item: UsageItem{
				Resource: &UsageItem_Resource{
					PrepaidDeploymentId:       "id",
					PrepaidDeploymentStartsAt: TimestampMust(now.Add(-time.Hour * 24)),
					PrepaidDeploymentEndsAt:   TimestampMust(now.Add(time.Hour * 24)),
				},
				StartsAt: TimestampMust(now.Add(-time.Hour)),
				EndsAt:   TimestampMust(now.Add(time.Hour * 4)),
			},
			expected: true,
		},
		{
			//----------ppppppppppp-------------------
			//--------uuuuuuuuuuuuuuuuu---------------
			name: "usage starts before prepaid deployment and ends after prepaid deployment",
			item: UsageItem{
				Resource: &UsageItem_Resource{
					PrepaidDeploymentId:       "id",
					PrepaidDeploymentStartsAt: TimestampMust(now.Add(-time.Hour * 24)),
					PrepaidDeploymentEndsAt:   TimestampMust(now.Add(time.Hour * 24)),
				},
				StartsAt: TimestampMust(now.Add(-time.Hour)),
				EndsAt:   TimestampMust(now.Add(time.Hour * 4)),
			},
			expected: true,
		},
		{
			//--------ppppppppppppppppp---------------
			//--------uuuuuuuuuuuuuuuuu---------------
			name: "usage starts before prepaid deployment and ends after prepaid deployment",
			item: UsageItem{
				Resource: &UsageItem_Resource{
					PrepaidDeploymentId:       "id",
					PrepaidDeploymentStartsAt: TimestampMust(now.Add(-time.Hour * 24)),
					PrepaidDeploymentEndsAt:   TimestampMust(now),
				},
				StartsAt: TimestampMust(now.Add(-time.Hour * 24)),
				EndsAt:   TimestampMust(now),
			},
			expected: true,
		},
	}

	for _, data := range tests {
		t.Run(data.name, func(t *testing.T) {
			actual := data.item.IsActivePrepaidDeployment()
			if actual != data.expected {
				t.Fail()
			}
		})
	}
}

func TimestampMust(t time.Time) *types.Timestamp {
	val, err := types.TimestampProto(t)
	if err != nil {
		panic(err)
	}
	return val
}

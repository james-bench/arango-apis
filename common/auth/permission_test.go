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

package auth

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPermission(t *testing.T) {
	tests := []struct {
		Input string
		Valid bool
		API   string
		Kind  string
		Verb  string
	}{
		{"iam.group.get", true, "iam", "group", "get"},
		{"foo.something.create", true, "foo", "something", "create"},
		{"foo.some-thing.create", true, "foo", "some-thing", "create"},
		{"foo .something.create", false, "", "", ""},
		{"foo.some_thing.create", false, "", "", ""},
		{"foo.something", false, "", "", ""},
		{"foo.something.create.more", false, "", "", ""},
	}
	for _, test := range tests {
		// Test parse
		api, kind, verb, err := ParsePermission(test.Input)
		if test.Valid {
			assert.NoError(t, err, test.Input)
			assert.Equal(t, test.API, api, "api")
			assert.Equal(t, test.Kind, kind, "kind")
			assert.Equal(t, test.Verb, verb, "verb")
		} else {
			assert.Error(t, err, test.Input)
		}

		// Test create
		if test.Valid {
			p := CreatePermission(test.API, test.Kind, test.Verb)
			assert.Equal(t, test.Input, p, test.Input)
		}
	}
}

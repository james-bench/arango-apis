//
// DISCLAIMER
//
// Copyright 2019 ArangoDB Inc, Cologne, Germany
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

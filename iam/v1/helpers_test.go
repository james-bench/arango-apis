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
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestGetAllEmails tests GetAllEmails
func TestGetAllEmails(t *testing.T) {
	var u *User
	assert.Len(t, u.GetAllEmails(), 0)

	u = &User{Email: "test@example.com"}
	assert.EqualValues(t, []string{"test@example.com"}, u.GetAllEmails())

	u = &User{Email: "e1", AdditionalEmails: []string{"a1", "z2", "c3", "t5", "m6"}}
	assert.EqualValues(t, []string{"a1", "c3", "e1", "m6", "t5", "z2"}, u.GetAllEmails())

	u = &User{Email: "", AdditionalEmails: []string{"a1", "z2", "c3", "t5", "m6"}}
	assert.EqualValues(t, []string{"a1", "c3", "m6", "t5", "z2"}, u.GetAllEmails())

	u = &User{Email: "z2", AdditionalEmails: []string{"a1", "z2", "c3", "t5", "m6"}}
	assert.EqualValues(t, []string{"a1", "c3", "m6", "t5", "z2"}, u.GetAllEmails())
}

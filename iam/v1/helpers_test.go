//
// DISCLAIMER
//
// Copyright 2019 ArangoDB GmbH, Cologne, Germany
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

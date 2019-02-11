//
// DISCLAIMER
//
// Copyright 2019 ArangoDB Inc, Cologne, Germany
//
// Author Ewout Prangsma
//

package v1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateMemberIDFromUser(t *testing.T) {
	assert.Equal(t, "user:123", CreateMemberIDFromUserID("123"))
	assert.Equal(t, "user:567", CreateMemberIDFromUser(&User{
		Id:   "567",
		Name: "foo",
	}))
}

func TestCreateMemberIDFromGroup(t *testing.T) {
	assert.Equal(t, "group:123", CreateMemberIDFromGroupID("123"))
	assert.Equal(t, "group:567", CreateMemberIDFromGroup(&Group{
		Id:   "567",
		Name: "foo",
	}))
}

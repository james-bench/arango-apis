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

func TestGroupURL(t *testing.T) {
	assert.Equal(t, "/Organization/123/Group/c1", GroupURL("123", "c1"))
	assert.Equal(t, "/Organization/123%2F567/Group/c2%2F3", GroupURL("123/567", "c2/3"))
	assert.Equal(t, "/Organization/123%2F567/Group/e%25f", GroupURL("123/567", "e%f"))
}

func TestRoleURL(t *testing.T) {
	assert.Equal(t, "/Organization/123/Role/c1", RoleURL("123", "c1"))
	assert.Equal(t, "/Organization/123%2F567/Role/c2%2F3", RoleURL("123/567", "c2/3"))
	assert.Equal(t, "/Organization/123%2F567/Role/e%25f", RoleURL("123/567", "e%f"))
}

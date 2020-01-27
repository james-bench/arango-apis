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

func TestIPWhitelistURL(t *testing.T) {
	assert.Equal(t, "/Organization/123/Project/p1/IPWhitelist/c1", IPWhitelistURL("123", "p1", "c1"))
	assert.Equal(t, "/Organization/123%2F567/Project/ab/IPWhitelist/c2%2F3", IPWhitelistURL("123/567", "ab", "c2/3"))
	assert.Equal(t, "/Organization/123%2F567/Project/a%25b/IPWhitelist/e%25f", IPWhitelistURL("123/567", "a%b", "e%f"))
}

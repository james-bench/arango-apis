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

func TestCACertificateURL(t *testing.T) {
	assert.Equal(t, "/Organization/123/Project/p1/CACertificate/c1", CACertificateURL("123", "p1", "c1"))
	assert.Equal(t, "/Organization/123%2F567/Project/ab/CACertificate/c2%2F3", CACertificateURL("123/567", "ab", "c2/3"))
	assert.Equal(t, "/Organization/123%2F567/Project/a%25b/CACertificate/e%25f", CACertificateURL("123/567", "a%b", "e%f"))
}

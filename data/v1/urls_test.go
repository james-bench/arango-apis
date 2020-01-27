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

func TestDeploymentURL(t *testing.T) {
	assert.Equal(t, "/Organization/123/Project/p1/Deployment/c1", DeploymentURL("123", "p1", "c1"))
	assert.Equal(t, "/Organization/123%2F567/Project/ab/Deployment/c2%2F3", DeploymentURL("123/567", "ab", "c2/3"))
	assert.Equal(t, "/Organization/123%2F567/Project/a%25b/Deployment/e%25f", DeploymentURL("123/567", "a%b", "e%f"))
}

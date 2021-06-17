package v1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrepaidDeploymentURL(t *testing.T) {
	assert.Equal(t, "/Organization/org_id/PrepaidDeployment/pd_id", PrepaidDeploymentURL("org_id", "pd_id"))
	assert.Equal(t, "/Organization/org%2Fid/PrepaidDeployment/pd%2Fid", PrepaidDeploymentURL("org/id", "pd/id"))
	assert.Equal(t, "/Organization/123/PrepaidDeployment/456", PrepaidDeploymentURL("123", "456"))
	assert.Equal(t, "/Organization/123/PrepaidDeployment/%3Fx=456", PrepaidDeploymentURL("123", "?x=456"))
}

//
// DISCLAIMER
//
// Copyright 2019 ArangoDB Inc, Cologne, Germany
//
// Author Robert Stam
//

package v1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBackupPolicyURL(t *testing.T) {
	assert.Equal(t, "/Organization/123/Project/p1/Deployment/c1/BackupPolicy/bp1", BackupPolicyURL("123", "p1", "c1", "bp1"))
	assert.Equal(t, "/Organization/123%2F567/Project/ab/Deployment/c2%2F3/BackupPolicy/bp%2F1", BackupPolicyURL("123/567", "ab", "c2/3", "bp/1"))
	assert.Equal(t, "/Organization/123%2F567/Project/a%25b/Deployment/e%25f/BackupPolicy/bp%251", BackupPolicyURL("123/567", "a%b", "e%f", "bp%1"))
}

func TestBackupURL(t *testing.T) {
	assert.Equal(t, "/Organization/123/Project/p1/Deployment/c1/Backup/b1", BackupURL("123", "p1", "c1", "b1"))
	assert.Equal(t, "/Organization/123%2F567/Project/ab/Deployment/c2%2F3/Backup/b%2F1", BackupURL("123/567", "ab", "c2/3", "b/1"))
	assert.Equal(t, "/Organization/123%2F567/Project/a%25b/Deployment/e%25f/Backup/b%251", BackupURL("123/567", "a%b", "e%f", "b%1"))
}

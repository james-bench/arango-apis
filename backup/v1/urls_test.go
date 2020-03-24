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

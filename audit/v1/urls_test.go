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

func TestAuditLogURL(t *testing.T) {
	assert.Equal(t, "/Organization/123/AuditLog/c1", AuditLogURL("123", "c1"))
	assert.Equal(t, "/Organization/123%2F567/AuditLog/c2%2F3", AuditLogURL("123/567", "c2/3"))
	assert.Equal(t, "/Organization/123%2F567/AuditLog/e%25f", AuditLogURL("123/567", "e%f"))
}

func TestAuditLogArchiveURL(t *testing.T) {
	assert.Equal(t, "/Organization/123/AuditLog/c1/AuditLogArchive/a1", AuditLogArchiveURL("123", "c1", "a1"))
	assert.Equal(t, "/Organization/123%2F567/AuditLog/c2%2F3/AuditLogArchive/d%2F5", AuditLogArchiveURL("123/567", "c2/3", "d/5"))
	assert.Equal(t, "/Organization/123%2F567/AuditLog/e%25f/AuditLogArchive/t%25T", AuditLogArchiveURL("123/567", "e%f", "t%T"))
}

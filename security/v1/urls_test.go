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

func TestIPAllowlistURL(t *testing.T) {
	assert.Equal(t, "/Organization/123/Project/p1/IPAllowlist/c1", IPAllowlistURL("123", "p1", "c1"))
	assert.Equal(t, "/Organization/123%2F567/Project/ab/IPAllowlist/c2%2F3", IPAllowlistURL("123/567", "ab", "c2/3"))
	assert.Equal(t, "/Organization/123%2F567/Project/a%25b/IPAllowlist/e%25f", IPAllowlistURL("123/567", "a%b", "e%f"))
}

func TestIAMProviderURL(t *testing.T) {
	assert.Equal(t, "/Organization/123/Project/p1/IAMProvider/c1", IAMProviderURL("123", "p1", "c1"))
	assert.Equal(t, "/Organization/123%2F567/Project/ab/IAMProvider/c2%2F3", IAMProviderURL("123/567", "ab", "c2/3"))
	assert.Equal(t, "/Organization/123%2F567/Project/a%25b/IAMProvider/e%25f", IAMProviderURL("123/567", "a%b", "e%f"))
}

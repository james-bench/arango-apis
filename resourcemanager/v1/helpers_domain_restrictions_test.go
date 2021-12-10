//
// DISCLAIMER
//
// Copyright 2021 ArangoDB GmbH, Cologne, Germany
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

package v1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsEmailAddressAllowed(t *testing.T) {
	// Empty restrictions
	var dr *DomainRestrictions
	assert.True(t, dr.IsEmailAddressAllowed("foo@bar"))
	// Even invalid address is allowed in this case
	assert.True(t, dr.IsEmailAddressAllowed("invalid"))

	// Explicit restrictions
	dr = &DomainRestrictions{
		AllowedDomains: []string{"foo.bar", "example.com"},
	}
	assert.True(t, dr.IsEmailAddressAllowed("foo@foo.bar"))
	assert.True(t, dr.IsEmailAddressAllowed("Foo <foo@foo.bar>"))
	assert.True(t, dr.IsEmailAddressAllowed("foo@example.com"))
	assert.True(t, dr.IsEmailAddressAllowed("Foo <foo@example.com>"))
	assert.True(t, dr.IsEmailAddressAllowed("Complex Foo <foo@example.com>"))
	// No subdomains are allowed
	assert.False(t, dr.IsEmailAddressAllowed("foo@sub.foo.bar"))
	assert.False(t, dr.IsEmailAddressAllowed("Foo <foo@sub.foo.bar>"))
	// No other domains are allowed
	assert.False(t, dr.IsEmailAddressAllowed("foo@other.com"))
	assert.False(t, dr.IsEmailAddressAllowed("Foo <foo@other.com>"))
	// No invalid email addresses are allowed
	assert.False(t, dr.IsEmailAddressAllowed("invalid@foo@foo.bar"))
	assert.False(t, dr.IsEmailAddressAllowed("foo.bar"))
}

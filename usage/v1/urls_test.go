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

func TestUsageURL(t *testing.T) {
	assert.Equal(t, "/Organization/123/UsageItem/u1", UsageItemURL("123", "u1"))
	assert.Equal(t, "/Organization/123%2F567/UsageItem/c2%2F3", UsageItemURL("123/567", "c2/3"))
	assert.Equal(t, "/Organization/123%2F567/UsageItem/a%25b", UsageItemURL("123/567", "a%b"))
}

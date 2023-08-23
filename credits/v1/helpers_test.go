//
// DISCLAIMER
//
// Copyright 2023 ArangoDB GmbH, Cologne, Germany
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
	"time"

	"github.com/gogo/protobuf/types"
	"github.com/stretchr/testify/assert"
)

func TestIsExpired(t *testing.T) {
	t.Run("expired", func(t *testing.T) {
		year := time.Now().Year() - 2
		validUntil, _ := types.TimestampProto(time.Date(year, 7, 16, 0, 0, 0, 0, time.UTC))
		b := &CreditBundle{
			ValidUntil: validUntil,
		}
		assert.True(t, b.IsExpired())
	})
	t.Run("not expired", func(t *testing.T) {
		year := time.Now().Year() + 2
		validUntil, _ := types.TimestampProto(time.Date(year, 7, 16, 0, 0, 0, 0, time.UTC))
		expiredBundle := &CreditBundle{
			ValidUntil: validUntil,
		}
		assert.False(t, expiredBundle.IsExpired())
	})
}

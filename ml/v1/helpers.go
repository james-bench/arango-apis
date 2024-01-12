//
// DISCLAIMER
//
// Copyright 2024 ArangoDB GmbH, Cologne, Germany
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

import "github.com/gogo/protobuf/types"

// IsExpired returns true if the MLServices has expired.
func (ml *MLServices) IsExpired() bool {
	// Check if we have exceeded the allowed quota?
	sts := ml.GetStatus()
	if hoursUsed, hoursAllowed := sts.GetHoursUsed(),
		sts.GetHoursAllowed(); hoursAllowed > 0 && hoursUsed > hoursAllowed {
		return true
	}
	// Check if we have exceeded to specified expiry deadline?
	now := types.TimestampNow()
	if expiresAt := sts.GetExpiresAt(); expiresAt != nil && now.Compare(expiresAt) > 0 {
		return true
	}
	return false
}

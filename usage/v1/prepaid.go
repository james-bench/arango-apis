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
// Author Marcin Swiderski
//

package v1

// IsActivePrepaidDeployment return true if usage item's range overlaps prepaid deployment range
func (u *UsageItem) IsActivePrepaidDeployment() bool {
	res := u.GetResource()
	if res == nil {
		return false
	}
	if res.GetPrepaidDeploymentId() == "" {
		return false
	}
	// there are only 2 cases when they do not overlap:
	// When usage item ends before prepaid deployment start
	if u.GetEndsAt().Compare(res.GetPrepaidDeploymentStartsAt()) <= 0 {
		return false
	}
	// When usage item starts after prepaid deployment end
	if u.GetStartsAt().Compare(res.GetPrepaidDeploymentEndsAt()) >= 0 {
		return false
	}

	return true
}

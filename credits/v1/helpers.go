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
	"github.com/gogo/protobuf/types"
)

// IsExpired returns true if a credit bundle is expired.
func (bundle *CreditBundle) IsExpired() bool {
	now := types.TimestampNow()
	return now.Compare(bundle.GetValidUntil()) > 0
}

// IsUsed returns true if a credit bundle has been used.
func (bundle *CreditBundle) IsUsed() bool {
	return bundle.GetCreditsPurchased() != bundle.GetCreditsRemaining()
}

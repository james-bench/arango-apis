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
	"net/mail"
	"strings"
)

// IsEmailAddressAllowed returns true if the given email address is allowed for
// the given email domain restrictions.
func (dr *DomainRestrictions) IsEmailAddressAllowed(emailAddress string) bool {
	allowedDomains := dr.GetAllowedDomains()
	if len(allowedDomains) == 0 {
		// All is allowed
		return true
	}
	addr, err := mail.ParseAddress(emailAddress)
	if err != nil {
		// Cannot parse email address, so do not allow
		return false
	}
	addrParts := strings.Split(addr.Address, "@")
	if len(addrParts) < 2 {
		// Not a valid email address
		return false
	}
	domain := strings.ToLower(addrParts[len(addrParts)-1])
	for _, x := range allowedDomains {
		if strings.ToLower(x) == domain {
			return true
		}
	}
	return false
}

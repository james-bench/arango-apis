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
	"net/url"
	"path"

	rm "github.com/arangodb-managed/apis/resourcemanager/v1"
)

const (
	// KindIPWhitelist is a constants for the kind of IPWhitelist resources.
	KindIPWhitelist = "IPWhitelist"
	// KindIAMProvider is a constants for the kind of IAMProvider resources.
	KindIAMProvider = "IAMProvider"
)

// IPWhitelistURL creates a resource URL for the IPWhitelist with given ID
// in given context.
func IPWhitelistURL(organizationID, projectID, ipwhitelistID string) string {
	return path.Join(rm.ProjectURL(organizationID, projectID), KindIPWhitelist, url.PathEscape(ipwhitelistID))
}

// IAMProviderURL creates a resource URL for the IAMProvider with given ID
// in given context.
func IAMProviderURL(organizationID, projectID, iamproviderID string) string {
	return path.Join(rm.ProjectURL(organizationID, projectID), KindIAMProvider, url.PathEscape(iamproviderID))
}

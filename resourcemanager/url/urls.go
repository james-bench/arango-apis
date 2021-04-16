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

package url

import (
	"net/url"
	"path"
)

const (
	// KindOrganization is a constants for the kind of Organization resources.
	KindOrganization = "Organization"
	// KindProject is a constants for the kind of Project resources.
	KindProject = "Project"
	// KindOrganizationInvite is a constants for the kind of OrganizationInvite resources.
	KindOrganizationInvite = "OrganizationInvite"
)

// OrganizationURL creates a resource URL for the organization with given ID.
func OrganizationURL(organizationID string) string {
	return path.Join("/", KindOrganization, url.PathEscape(organizationID))
}

// ProjectURL creates a resource URL for the project & organization with given IDs.
func ProjectURL(organizationID, projectID string) string {
	return path.Join(OrganizationURL(organizationID), KindProject, url.PathEscape(projectID))
}

// OrganizationInviteURL creates a resource URL for the invite & organization with given IDs.
func OrganizationInviteURL(organizationID, inviteID string) string {
	return path.Join(OrganizationURL(organizationID), KindOrganizationInvite, url.PathEscape(inviteID))
}

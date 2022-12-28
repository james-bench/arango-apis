//
// DISCLAIMER
//
// Copyright 2022 ArangoDB GmbH, Cologne, Germany
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

const (
	// EventTypeOrganizationMemberAdded is the type of event fired after a member has been added to an organization.
	// SubjectID contains the ID of the added member.
	EventTypeOrganizationMemberAdded = "scim.organization-member.added"
	// EventTypeOrganizationMemberRemoved is the type of event fired after a member has been removed from an organization.
	// SubjectID contains the ID of the removed member.
	EventTypeOrganizationMemberRemoved = "scim.organization-member.removed"
	// EventTypeOrganizationMemberUpdated is the type of event fired after a member has been updated in an organization
	// from owner to non-owner or from non-owner to owner.
	// SubjectID contains the ID of the updated member.
	EventTypeOrganizationMemberUpdated = "scim.organization-member.updated"
)

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

const (
	// Group event types

	// EventTypeGroupCreated is the type of event fired after a group has been created
	// SubjectID contains the ID of the group.
	EventTypeGroupCreated = "iam.group.created"
	// EventTypeGroupUpdated is the type of event fired after a group has been updated
	// SubjectID contains the ID of the group.
	EventTypeGroupUpdated = "iam.group.updated"
	// EventTypeGroupDeleted is the type of event fired after a group has been (marked for) deleted
	// SubjectID contains the ID of the group.
	EventTypeGroupDeleted = "iam.group.deleted"

	// EventTypeGroupMemberAdded is the type of event fired after a member has been added to a group.
	// SubjectID contains the ID of the added member.
	EventTypeGroupMemberAdded = "iam.group-member.added"
	// EventTypeGroupMemberRemoved is the type of event fired after a member has been removed from a group.
	// SubjectID contains the ID of the removed member.
	EventTypeGroupMemberRemoved = "iam.group-member.removed"
)

const (
	// Role event types

	// EventTypeRoleCreated is the type of event fired after a (custom) role has been created
	// SubjectID contains the ID of the role.
	EventTypeRoleCreated = "iam.role.created"
	// EventTypeRoleUpdated is the type of event fired after a (custom) role has been updated
	// SubjectID contains the ID of the role.
	EventTypeRoleUpdated = "iam.role.updated"
	// EventTypeRoleDeleted is the type of event fired after a (custom) role has been (marked for) deleted
	// SubjectID contains the ID of the role.
	EventTypeRoleDeleted = "iam.role.deleted"
)

const (
	// Policy event types

	// EventTypePolicyUpdated is the type of event fired after a policy has been updated
	// SubjectID contains the URL of the policy.
	EventTypePolicyUpdated = "iam.policy.updated"
)

const (
	// APIKey event types

	// EventTypeAPIKeyCreated is the type of event fired after an API key has been created.
	// SubjectID contains the ID of the api key.
	EventTypeAPIKeyCreated = "iam.apikey.created"
	// EventTypeAPIKeyRevoked is the type of event fired after an API key has been revoked.
	// SubjectID contains the ID of the api key.
	EventTypeAPIKeyRevoked = "iam.apikey.revoked"
	// EventTypeAPIKeyDeleted is the type of event fired after an API key has been deleted.
	// SubjectID contains the ID of the api key.
	EventTypeAPIKeyDeleted = "iam.apikey.deleted"
)

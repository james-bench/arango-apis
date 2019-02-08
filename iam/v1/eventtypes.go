//
// DISCLAIMER
//
// Copyright 2019 ArangoDB Inc, Cologne, Germany
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

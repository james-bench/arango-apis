//
// DISCLAIMER
//
// Copyright 2019 ArangoDB Inc, Cologne, Germany
//
// Author Ewout Prangsma
//

package v1

// CreateMemberIDFromUser creates a member ID from the given user.
func CreateMemberIDFromUser(user *User) string {
	return CreateMemberIDFromUserID(user.GetId())
}

// CreateMemberIDFromUserID creates a member ID from the given user ID.
func CreateMemberIDFromUserID(userID string) string {
	return "user:" + userID
}

// CreateMemberIDFromGroup creates a member ID from the given group.
func CreateMemberIDFromGroup(group *Group) string {
	return CreateMemberIDFromGroupID(group.GetId())
}

// CreateMemberIDFromGroupID creates a member ID from the given group ID.
func CreateMemberIDFromGroupID(groupID string) string {
	return "group:" + groupID
}

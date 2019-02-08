//
// DISCLAIMER
//
// Copyright 2019 ArangoDB Inc, Cologne, Germany
//
// Author Ewout Prangsma
//

package v1

const (
	// Organization event types

	// EventTypeOrganizationCreated is the type of event fired after an organization has been created
	// SubjectID contains the ID of the organization.
	EventTypeOrganizationCreated = "resourcemanager.organization.created"
	// EventTypeOrganizationUpdated is the type of event fired after an organization has been updated
	// SubjectID contains the ID of the organization.
	EventTypeOrganizationUpdated = "resourcemanager.organization.updated"
	// EventTypeOrganizationDeleted is the type of event fired after an organization has been (marked for) deleted
	// SubjectID contains the ID of the organization.
	EventTypeOrganizationDeleted = "resourcemanager.organization.deleted"

	// EventTypeOrganizationMemberAdded is the type of event fired after a member has been added to an organization.
	// SubjectID contains the ID of the added member.
	EventTypeOrganizationMemberAdded = "resourcemanager.organization-member.added"
	// EventTypeOrganizationMemberRemoved is the type of event fired after a member has been removed from an organization.
	// SubjectID contains the ID of the removed member.
	EventTypeOrganizationMemberRemoved = "resourcemanager.organization-member.removed"
)

const (
	// Project event types

	// EventTypeProjectCreated is the type of event fired after a project has been created.
	// SubjectID contains the ID of the project.
	EventTypeProjectCreated = "resourcemanager.project.created"
	// EventTypeProjectUpdated is the type of event fired after a project has been updated.
	// SubjectID contains the ID of the project.
	EventTypeProjectUpdated = "resourcemanager.project.updated"
	// EventTypeProjectDeleted is the type of event fired after a project has been (marked for) deleted.
	// SubjectID contains the ID of the project.
	EventTypeProjectDeleted = "resourcemanager.project.deleted"
)

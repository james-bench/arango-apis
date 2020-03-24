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
	// EventTypeOrganizationMemberUpdated is the type of event fired after a member has been updated in an organization
	// from owner to non-owner or from non-owner to owner.
	// SubjectID contains the ID of the updated member.
	EventTypeOrganizationMemberUpdated = "resourcemanager.organization-member.updated"
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

const (
	// OrganizationInvite event types

	// EventTypeOrganizationInviteCreated is the type of event fired after an organization invite has been created
	// SubjectID contains the ID of the organization invite.
	EventTypeOrganizationInviteCreated = "resourcemanager.organization-invite.created"
	// EventTypeOrganizationInviteDeleted is the type of event fired after an organization has been (marked for) deleted
	// SubjectID contains the ID of the organization.
	EventTypeOrganizationInviteDeleted = "resourcemanager.organization-invite.deleted"
	// EventTypeOrganizationInviteAccepted is the type of event fired after an invite has been accepted for an organization.
	// SubjectID contains the ID of the accepted invite.
	EventTypeOrganizationInviteAccepted = "resourcemanager.organization-invite.accepted"
	// EventTypeOrganizationInviteRejected is the type of event fired after an invite has been rejected for an organization.
	// SubjectID contains the ID of the rejected invite.
	EventTypeOrganizationInviteRejected = "resourcemanager.organization-invite.rejected"
)

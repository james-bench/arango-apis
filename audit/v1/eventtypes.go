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
	// AuditLog event types

	// EventTypeAuditLogCreated is the type of event fired after an AuditLog has been created
	// SubjectID contains the ID of the AuditLog.
	EventTypeAuditLogCreated = "audit.auditlog.created"
	// EventTypeAuditLogUpdated is the type of event fired after an AuditLog has been updated
	// SubjectID contains the ID of the AuditLog.
	EventTypeAuditLogUpdated = "audit.auditlog.updated"
	// EventTypeAuditLogDeleted is the type of event fired after an AuditLog has been (marked for) deleted
	// SubjectID contains the ID of the AuditLog.
	EventTypeAuditLogDeleted = "audit.auditlog.deleted"
	// EventTypeAuditLogAttached is the type of event fired after an AuditLog has been
	// attached to a Project.
	// SubjectID contains the ID of the AuditLog.
	// The Payload contains a "project_id" field containing the ID of the project.
	EventTypeAuditLogAttached = "audit.auditlog.attached"
	// EventTypeAuditLogDetached is the type of event fired after an AuditLog has been
	// detached from a Project.
	// SubjectID contains the ID of the AuditLog.
	// The Payload contains a "project_id" field containing the ID of the project.
	EventTypeAuditLogDetached = "audit.auditlog.detached"
)

const (
	// PayloadAuditLogAttachmentProjectID is the key of the payload field, of
	// an event (of type EventTypeAuditLogAttached or EventTypeAuditLogDetached),
	// that contains the project ID.
	PayloadAuditLogAttachmentProjectID = "project_id"
)

const (
	// AuditLogArchive event types

	// EventTypeAuditLogArchiveCreated is the type of event fired after an AuditLogArchive has been created
	// SubjectID contains the ID of the AuditLogArchive.
	EventTypeAuditLogArchiveCreated = "audit.auditlogarchive.created"
	// EventTypeAuditLogArchiveUpdated is the type of event fired after an AuditLogArchive has been updated
	// SubjectID contains the ID of the AuditLogArchive.
	EventTypeAuditLogArchiveUpdated = "audit.auditlogarchive.updated"
	// EventTypeAuditLogArchiveDeleted is the type of event fired after an AuditLogArchive has been (marked for) deleted
	// SubjectID contains the ID of the AuditLogArchive.
	EventTypeAuditLogArchiveDeleted = "audit.auditlogarchive.deleted"
)

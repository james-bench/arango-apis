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
	// AuditLog permissions

	// PermissionAuditLogList is needed to list AuditLogs.
	PermissionAuditLogList = "audit.auditlog.list"
	// PermissionAuditLogGet is needed to get an individual AuditLog.
	PermissionAuditLogGet = "audit.auditlog.get"
	// PermissionAuditLogCreate is needed to create a new AuditLog.
	PermissionAuditLogCreate = "audit.auditlog.create"
	// PermissionAuditLogUpdate is needed to update an existing AuditLog.
	PermissionAuditLogUpdate = "audit.auditlog.update"
	// PermissionAuditLogDelete is needed to delete an existing AuditLog.
	PermissionAuditLogDelete = "audit.auditlog.delete"
)

const (
	// AuditLogArchive permissions

	// PermissionAuditLogArchiveList is needed to list AuditLogArchives.
	PermissionAuditLogArchiveList = "audit.auditlogarchive.list"
	// PermissionAuditLogArchiveGet is needed to get an individual AuditLogArchive.
	PermissionAuditLogArchiveGet = "audit.auditlogarchive.get"
	// PermissionAuditLogArchiveDelete is needed to delete an existing AuditLogArchive.
	PermissionAuditLogArchiveDelete = "audit.auditlogarchive.delete"
)

const (
	// AuditLogEvent permissions

	// PermissionAuditLogEventDelete is needed to delete AuditLogEvents from an archive.
	PermissionAuditLogEventDelete = "audit.auditlogevent.delete"
)

const (
	// AuditLogEvents permissions

	// PermissionAuditLogEventsGet is needed to fetching AuditLogEvents.
	PermissionAuditLogEventsGet = "audit.auditlogevents.get"
)

const (
	// AuditLogAttachment permissions

	// PermissionAuditLogAttachmentGet is needed to get an an AuditLog attachment to a Project.
	PermissionAuditLogAttachmentGet = "audit.auditlogattachment.get"
	// PermissionAuditLogAttachmentCreate is needed to attach an AuditLog to a Project.
	PermissionAuditLogAttachmentCreate = "audit.auditlogattachment.create"
	// PermissionAuditLogAttachmentDelete is needed to detach an AuditLog from a Project.
	PermissionAuditLogAttachmentDelete = "audit.auditlogattachment.delete"
)

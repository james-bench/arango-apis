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
	// KindAuditLog is a constants for the kind of AuditLog resources.
	KindAuditLog = "AuditLog"
	// KindAuditLogArchive is a constants for the kind of AuditLogArchive resources.
	KindAuditLogArchive = "AuditLogArchive"
	// KindAuditLogAttachment is a constants for the kind of AuditLogAttachment resources.
	KindAuditLogAttachment = "AuditLogAttachment"
)

// AuditLogURL creates a resource URL for the AuditLog with given ID
// in given context.
func AuditLogURL(organizationID, auditLogID string) string {
	return path.Join(rm.OrganizationURL(organizationID), KindAuditLog, url.PathEscape(auditLogID))
}

// AuditLogArchiveURL creates a resource URL for the AuditLogArchive with given ID
// in given context.
func AuditLogArchiveURL(organizationID, auditLogID, auditLogArchiveID string) string {
	return path.Join(AuditLogURL(organizationID, auditLogID), KindAuditLogArchive, url.PathEscape(auditLogArchiveID))
}

// AuditLogAttachmentURL creates a resource URL for the AuditLogAttachment with given ID
// in given context.
func AuditLogAttachmentURL(organizationID, projectID, auditLogID string) string {
	return path.Join(rm.ProjectURL(organizationID, projectID), KindAuditLogAttachment, url.PathEscape(auditLogID))
}

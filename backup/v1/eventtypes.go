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
// Author Robert Stam
//

package v1

const (
	// BackupPolicy event types

	// EventTypeBackupPolicy Created is the type of event fired after a backup policy has been created
	// SubjectID contains the ID of the backup policy.
	EventTypeBackupPolicyCreated = "backup.backuppolicy.created"
	// EventTypeBackupPolicyUpdated is the type of event fired after a backup policy has been updated
	// SubjectID contains the ID of the backup policy.
	// Note that this type of event is also fired when the status of the backup policy has changed.
	EventTypeBackupPolicyUpdated = "backup.backuppolicy.updated"
	// EventTypeBackupPolicyDeleted is the type of event fired after a backup policy has been (marked for) deleted
	// SubjectID contains the ID of the backup policy.
	EventTypeBackupPolicyDeleted = "backup.backuppolicy.deleted"
)

const (
	// Backup event types

	// EventTypeBackupCreated is the type of event fired after a backup has been created
	// SubjectID contains the ID of the backup.
	EventTypeBackupCreated = "backup.backup.created"
	// EventTypeBackupUpdated is the type of event fired after a backup has been updated
	// SubjectID contains the ID of the backup.
	// Note that this type of event is also fired when the status of the backup has changed.
	EventTypeBackupUpdated = "backup.backup.updated"
	// EventTypeBackupDeleted is the type of event fired after a backup has been (marked for) deleted
	// SubjectID contains the ID of the backup.
	EventTypeBackupDeleted = "backup.backup.deleted"
)

//
// DISCLAIMER
//
// Copyright 2019 ArangoDB Inc, Cologne, Germany
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

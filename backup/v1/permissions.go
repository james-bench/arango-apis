//
// DISCLAIMER
//
// Copyright 2019 ArangoDB Inc, Cologne, Germany
//
// Author Robert Stam
//

package v1

const (
	// BackupPolicy permissions

	// PermissionBackupPolicyList is needed for listing backup policies in a deployment
	PermissionBackupPolicyList = "backup.backuppolicy.list"
	// PermissionBackupPolicyGet is needed for getting indivual backup policy in a deployment
	PermissionBackupPolicyGet = "backup.backuppolicy.get"
	// PermissionBackupPolicyCreate is needed for create a backup policy
	PermissionBackupPolicyCreate = "backup.backuppolicy.create"
	// PermissionBackupPolicyUpdate is needed for updating a backup policy
	PermissionBackupPolicyUpdate = "backup.backuppolicy.update"
	// PermissionBackupPolicyDelete is needed for deleting a backup policy
	PermissionBackupPolicyDelete = "backup.backuppolicy.delete"
)

const (
	// Backup permissions

	// PermissionBackupList is needed for listing backups in a deployment
	PermissionBackupList = "backup.backup.list"
	// PermissionBackupGet is needed for getting indivual backup in a deployment
	PermissionBackupGet = "backup.backup.get"
	// PermissionBackupCreate is needed for create a backup
	PermissionBackupCreate = "backup.backup.create"
	// PermissionBackupUpdate is needed for updating a backup (this includes upload and download)
	PermissionBackupUpdate = "backup.backup.update"
	// PermissionBackupRestore is needed for restoring a backup
	PermissionBackupRestore = "backup.backup.restore"
	// PermissionBackupDelete is needed for deleting a backup
	PermissionBackupDelete = "backup.backup.delete"
)

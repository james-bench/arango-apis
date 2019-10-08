//
// This file is AUTO-GENERATED by protoc-gen-ts.
// Do not modify it manually.
///
import api from '../../api'
import * as googleTypes from '../../googleTypes'
import { IDOptions as arangodb_cloud_common_v1_IDOptions } from '../../common/v1/common'
import { ListOptions as arangodb_cloud_common_v1_ListOptions } from '../../common/v1/common'
import { YesOrNo as arangodb_cloud_common_v1_YesOrNo } from '../../common/v1/common'
import { Deployment_ServersSpec as arangodb_cloud_data_v1_Deployment_ServersSpec } from '../../data/v1/data'

// File: backup/v1/backup.proto
// Package: arangodb.cloud.backup.v1

// Backup represents a single backup of a deployment.
export interface Backup {
  // System identifier of the backup.
  // This is a read-only value.
  // string
  id?: string;
  
  // URL of this resource
  // This is a read-only value.
  // string
  url?: string;
  
  // Name of the backup
  // string
  name?: string;
  
  // Description of the backup
  // string
  description?: string;
  
  // Identifier of the deployment that owns this backup.
  // After creation, this value cannot be changed.
  // string
  deployment_id?: string;
  
  // Identifier of the backup policy that triggered this backup
  // After creation, this value cannot be changed.
  // If this field is empty, this is a manual backup
  // string
  backup_policy_id?: string;
  
  // The creation timestamp of the backup (database object)
  // This is a read-only value.
  // googleTypes.Timestamp
  created_at?: googleTypes.Timestamp;
  
  // The deletion timestamp of the backup
  // This is a read-only value.
  // googleTypes.Timestamp
  deleted_at?: googleTypes.Timestamp;
  
  // Set when this backup is deleted.
  // This is a read-only value.
  // boolean
  is_deleted?: boolean;
  
  // The timestamp that this backup will be automatically removed
  // You cannot provide a value in the past,
  // If the field is not set, the backup will not be automatically removed.
  // googleTypes.Timestamp
  auto_deleted_at?: googleTypes.Timestamp;
  
  // Information about the deployment during backup
  // Backup_DeploymentInfo
  deployment_info?: Backup_DeploymentInfo;
  
  // Upload the backup, created by the backup policy, to an external source.
  // Setting or unsetting this fields after the backup has been created will upload/delete the backup from the external source.
  // Setting this field when status.available = false will result in an error
  // boolean
  upload?: boolean;
  
  // The timestamp of when the upload boolean has been updated.
  // This is a read-only value.
  // googleTypes.Timestamp
  upload_updated_at?: googleTypes.Timestamp;
  
  // Information about a backup download.
  // If this field is set the backup will be downloaded the deployment.
  // This is a read-only field. To set this field please use the DownloadBackup method.
  // Backup_DownloadSpec
  download?: Backup_DownloadSpec;
  
  // Status of the actual backup
  // Backup_Status
  status?: Backup_Status;
}

// Information about the deployment during backup
// All members of this field are read-only.
export interface Backup_DeploymentInfo {
  // ArangoDB version of the deployment during backup.
  // string
  version?: string;
  
  // Servers spec of the deployment during backup.
  // arangodb.cloud.data.v1.Deployment.ServersSpec
  servers?: arangodb_cloud_data_v1_Deployment_ServersSpec;
}

// Information about a backup download.
// All members of this message are read-only.
export interface Backup_DownloadSpec {
  // The revision of this DownloadSpec
  // number
  revision?: number;
  
  // The timestamp of when the last revision has been updated.
  // googleTypes.Timestamp
  last_updated_at?: googleTypes.Timestamp;
}

// The status of backup download
// All members of this message are read-only.
export interface Backup_DownloadStatus {
  // The revision of the used DownloadStatus
  // number
  revision?: number;
  
  // Set when the backup has been fully downloaded
  // boolean
  downloaded?: boolean;
  
  // The timestamp of when the backup has been fully downloaded.
  // googleTypes.Timestamp
  downloaded_at?: googleTypes.Timestamp;
}

// Status of the actual backup
// All members of this field are read-only.
export interface Backup_Status {
  // The creation timestamp of the backup
  // googleTypes.Timestamp
  created_at?: googleTypes.Timestamp;
  
  // ArangoDB version of the backup
  // string
  version?: string;
  
  // The state of the backup
  // Will be one of the following: "Pending|Unavailable|Scheduled|Download|DownloadError|Downloading|Create|Upload|Uploading|UploadError|Ready|Deleted|Failed"
  // string
  state?: string;
  
  // Set when the backup is failed
  // boolean
  is_failed?: boolean;
  
  // State message
  // string
  message?: string;
  
  // Progress of the backup (upload or download)
  // string
  progress?: string;
  
  // Size of the backup (in bytes)
  // number
  size_bytes?: number;
  
  // If set the backup is available on the cluster and can be restored
  // boolean
  available?: boolean;
  
  // Number of dbservers of the deployment during backup
  // number
  dbservers?: number;
  
  // Indicates that the backup is available in the external source only.
  // You should download the backup before you can restore it.
  // boolean
  upload_only?: boolean;
  
  // The status of backup upload (if applicable).
  // Backup_UploadStatus
  upload_status?: Backup_UploadStatus;
  
  // The status of backup download (if applicable).
  // This field will be set to empty if a new revision of the spec is available
  // Backup_DownloadStatus
  download_status?: Backup_DownloadStatus;
}

// The status of backup upload
// All members of this message are read-only.
export interface Backup_UploadStatus {
  // Set when the backup has been fully uploaded
  // boolean
  uploaded?: boolean;
  
  // The timestamp of when the backup has been fully uploaded
  // googleTypes.Timestamp
  uploaded_at?: googleTypes.Timestamp;
  
  // Size of the backup in the external source (in bytes)
  // number
  size_bytes?: number;
}

// List of backups.
export interface BackupList {
  // Backup
  items?: Backup[];
}

// BackupPolicy represents a single backup policy for a deployment.
export interface BackupPolicy {
  // System identifier of the backup policy.
  // This is a read-only value.
  // string
  id?: string;
  
  // URL of this resource
  // This is a read-only value.
  // string
  url?: string;
  
  // Name of the backup policy
  // string
  name?: string;
  
  // Description of the backup policy
  // string
  description?: string;
  
  // Identifier of the deployment that owns this backup policy.
  // After creation, this value cannot be changed.
  // string
  deployment_id?: string;
  
  // The creation timestamp of the backup policy
  // This is a read-only value.
  // googleTypes.Timestamp
  created_at?: googleTypes.Timestamp;
  
  // The deletion timestamp of the backup policy
  // This is a read-only value.
  // googleTypes.Timestamp
  deleted_at?: googleTypes.Timestamp;
  
  // Set when this backup policy is deleted.
  // This is a read-only value.
  // boolean
  is_deleted?: boolean;
  
  // Pause this backup policy.
  // If a backup policy is paused, the backup policy will not result in new backups.
  // The backup policy isn't deleted, unsetting this field will resume the creation of backups again.
  // boolean
  is_paused?: boolean;
  
  // The schedule for this backup policy
  // BackupPolicy_Schedule
  schedule?: BackupPolicy_Schedule;
  
  // Upload the backup, created by the backup policy, to an external source.
  // boolean
  upload?: boolean;
  
  // Backups created by this policy will be automatically deleted after the specified retention period
  // A value of 0 means that backup will never be deleted.
  // googleTypes.Duration
  retention_period?: googleTypes.Duration;
  
  // The owners of the organization can be notified by email
  // This field support the following values: "None|FailureOnly|Always"
  // string
  email_notification?: string;
  
  // Status of the backup policy
  // BackupPolicy_Status
  status?: BackupPolicy_Status;
}

// Note: Nested types inside nested types is not supported by the typescript generator
export interface BackupPolicy_DailySchedule {
  // If set, a backup will be created on Mondays.
  // boolean
  monday?: boolean;
  
  // If set, a backup will be created on Tuesdays.
  // boolean
  tuesday?: boolean;
  
  // If set, a backup will be created on Wednesdays.
  // boolean
  wednesday?: boolean;
  
  // If set, a backup will be created on Thursdays.
  // boolean
  thursday?: boolean;
  
  // If set, a backup will be created on Fridays.
  // boolean
  friday?: boolean;
  
  // If set, a backup will be created on Saturdays.
  // boolean
  saturday?: boolean;
  
  // If set, a backup will be created on Sundays.
  // boolean
  sunday?: boolean;
  
  // The (target) time of the schedule
  // TimeOfDay
  schedule_at?: TimeOfDay;
}

// Note: Nested types inside nested types is not supported by the typescript generator
export interface BackupPolicy_HourlySchedule {
  // Schedule should run with an interval of the specified hours (1-23)
  // number
  schedule_every_interval_hours?: number;
}

// Note: Nested types inside nested types is not supported by the typescript generator
export interface BackupPolicy_MonthlySchedule {
  // Run the backup on the specified day of the month (1-31)
  // Note: Specifying a number larger than some months have days will result in no backup for those months (e.g. 29 for February (unless leap year)).
  // number
  day_of_month?: number;
  
  // The (target) time of the schedule
  // TimeOfDay
  schedule_at?: TimeOfDay;
}
export interface BackupPolicy_Schedule {
  // Schedule type should be one of the following string: "Hourly|Daily|Monthly"
  // The schedule_hourly, schedule_daily or schedule_montly field should be set
  // Setting multiple fields, or inconsistent with this field result in an error during create/update
  // string
  schedule_type?: string;
  
  // Schedule applies to the selected day of the week
  // This is applicable for Hourly type only, ignored for Daily and Monthly
  // BackupPolicy_HourlySchedule
  hourly_schedule?: BackupPolicy_HourlySchedule;
  
  // Schedule applies to the selected day of the week
  // This is applicable for Daily type only, ignored for Hourly and Monthly
  // BackupPolicy_DailySchedule
  daily_schedule?: BackupPolicy_DailySchedule;
  
  // Schedule applies to the selected day of the month
  // This is applicable for Monthly type only, ignored for Hourly and Daily
  // BackupPolicy_MonthlySchedule
  monthly_schedule?: BackupPolicy_MonthlySchedule;
}

// Status of the backup policy
// All members of this field are read-only.
export interface BackupPolicy_Status {
  // The timestamp when the next backup - initiated by this backup policy - will be created
  // googleTypes.Timestamp
  next_backup?: googleTypes.Timestamp;
  
  // Message in case of failure, otherwise an empty string
  // string
  message?: string;
}

// List of backup policies.
export interface BackupPolicyList {
  // BackupPolicy
  items?: BackupPolicy[];
}

// Request arguments for ListBackupPolicies
export interface ListBackupPoliciesRequest {
  // Identifier of the deployment to request the backup policies for.
  // string
  deployment_id?: string;
  
  // If set, the result includes all backup policies, including those who set to deleted,
  // however are not removed from the system currently.
  // If not set, only backup policies not indicated as deleted are returned.
  // boolean
  include_deleted?: boolean;
  
  // Optional common list options, the context_id is ignored
  // arangodb.cloud.common.v1.ListOptions
  options?: arangodb_cloud_common_v1_ListOptions;
}

// Request arguments for ListBackups
export interface ListBackupsRequest {
  // Identifier of the deployment to request the backups for.
  // string
  deployment_id?: string;
  
  // Request backups that are created at or after this timestamp.
  // This is an optional field.
  // googleTypes.Timestamp
  from?: googleTypes.Timestamp;
  
  // Request backups that are created before this timestamp.
  // This is an optional field.
  // googleTypes.Timestamp
  to?: googleTypes.Timestamp;
  
  // Optional common list options, the context_id is ignored
  // arangodb.cloud.common.v1.ListOptions
  options?: arangodb_cloud_common_v1_ListOptions;
}

// TimeOfDay describes a specific moment on a day
export interface TimeOfDay {
  // Hours part of the time of day (0-23)
  // number
  hours?: number;
  
  // Minutes part of the time of day (0-59)
  // number
  minutes?: number;
  
  // The time-zone this time of day applies to (empty means UTC)
  // Names MUST be exactly as defined in RFC-822.
  // string
  time_zone?: string;
}

// BackupService is the API used to configure backup objects.
export class BackupService {
  // Checks if the backup feature is enabled and available for a specific deployment.
  // Required permissions:
  // - backup.feature.get on the deployment that is identified by the given ID.
  async IsBackupFeatureAvailable(req: arangodb_cloud_common_v1_IDOptions): Promise<arangodb_cloud_common_v1_YesOrNo> {
    const path = `/api/backup/v1/deployment/${encodeURIComponent(req.id || '')}/feature`;
    const url = path + api.queryString(req, [`id`]);
    return api.get(url, undefined);
  }
  
  // Fetch all backup policies for a specific deployment.
  // Required permissions:
  // - backup.backuppolicy.list on the deployment that owns the backup policies and is identified by the given ID.
  async ListBackupPolicies(req: ListBackupPoliciesRequest): Promise<BackupPolicyList> {
    const path = `/api/backup/v1/deployment/${encodeURIComponent(req.deployment_id || '')}/backuppolicies`;
    const url = path + api.queryString(req, [`deployment_id`]);
    return api.get(url, undefined);
  }
  
  // Fetch a backup policy identified by the given ID.
  // Required permissions:
  // - backup.backuppolicy.get on the backup policy identified by the given ID.
  async GetBackupPolicy(req: arangodb_cloud_common_v1_IDOptions): Promise<BackupPolicy> {
    const path = `/api/backup/v1/backuppolicies/${encodeURIComponent(req.id || '')}`;
    const url = path + api.queryString(req, [`id`]);
    return api.get(url, undefined);
  }
  
  // Create a new backup policy
  // Required permissions:
  // -  backup.backuppolicy.create on the deployment that owns the backup policy and is identified by the given ID.
  async CreateBackupPolicy(req: BackupPolicy): Promise<BackupPolicy> {
    const url = `/api/backup/v1/deployment/${encodeURIComponent(req.deployment_id || '')}/backuppolicies`;
    return api.post(url, req);
  }
  
  // Update a backup policy
  // Required permissions:
  // -  backup.backuppolicy.update on the backup policy identified by the given ID.
  async UpdateBackupPolicy(req: BackupPolicy): Promise<BackupPolicy> {
    const url = `/api/backup/v1/backuppolicies/${encodeURIComponent(req.id || '')}`;
    return api.patch(url, req);
  }
  
  // Delete a backup policy identified by the given ID.
  // Note that the backup policy are initially only marked for deletion, no backups will be deleted with this operation.
  // Once all their dependent backups are removed, the backup policy is removed.
  // Required permissions:
  // -  backup.backuppolicy.delete on the backup policy identified by the given ID.
  async DeleteBackupPolicy(req: arangodb_cloud_common_v1_IDOptions): Promise<void> {
    const path = `/api/backup/v1/backuppolicies/${encodeURIComponent(req.id || '')}`;
    const url = path + api.queryString(req, [`id`]);
    return api.delete(url, undefined);
  }
  
  // Fetch all backups for a specific deployment.
  // Required permissions:
  // - backup.backup.list on the deployment that owns the backup and is identified by the given ID.
  async ListBackups(req: ListBackupsRequest): Promise<BackupList> {
    const path = `/api/backup/v1/deployment/${encodeURIComponent(req.deployment_id || '')}/backups`;
    const url = path + api.queryString(req, [`deployment_id`]);
    return api.get(url, undefined);
  }
  
  // Fetch a backup identified by the given ID.
  // Required permissions:
  // - backup.backup.get on the backup identified by the given ID.
  async GetBackup(req: arangodb_cloud_common_v1_IDOptions): Promise<Backup> {
    const path = `/api/backup/v1/backup/${encodeURIComponent(req.id || '')}`;
    const url = path + api.queryString(req, [`id`]);
    return api.get(url, undefined);
  }
  
  // Create a new manual backup
  // Setting the backup_policy_id field in the backup is not allowed
  // Required permissions:
  // -  backup.backup.create on the deployment that owns the backup and is identified by the given ID.
  async CreateBackup(req: Backup): Promise<Backup> {
    const url = `/api/backup/v1/deployment/${encodeURIComponent(req.deployment_id || '')}/backup`;
    return api.post(url, req);
  }
  
  // Update a backup
  // Required permissions:
  // -  backup.backup.update on the backup identified by the given ID.
  async UpdateBackup(req: Backup): Promise<Backup> {
    const url = `/api/backup/v1/backup/${encodeURIComponent(req.id || '')}`;
    return api.patch(url, req);
  }
  
  // Download a backup identified by the given ID from remote storage to the volumes of the servers of the deployment
  // If this backup was already be downloaded, another download will be done.
  // If the backup is still available on the cluster there is no need to explicitly download the backup before restoring.
  // This function will return immediately.
  // To track status, please invoke GetBackup and check the .status field inside the returned backup object
  // Required permissions:
  // -  backup.backup.download on the backup identified by the given ID.
  async DownloadBackup(req: arangodb_cloud_common_v1_IDOptions): Promise<void> {
    const path = `/api/backup/v1/backup/${encodeURIComponent(req.id || '')}/download`;
    const url = path + api.queryString(req, [`id`]);
    return api.post(url, undefined);
  }
  
  // Restore (or recover) a backup identified by the given ID
  // This operation can only be executed on backups where status.available is set
  // This function will return immediately.
  // To track status, please invoke GetDeployment on the data API and check the
  // .status.restoring_backup and .status.restore_backup_status fields inside the returned deployment object
  // Required permissions (both are needed):
  // -  backup.backup.restore on the backup identified by the given ID.
  // -  data.deployment.restore-backup on the deployment that owns this backup
  async RestoreBackup(req: arangodb_cloud_common_v1_IDOptions): Promise<void> {
    const path = `/api/backup/v1/backup/${encodeURIComponent(req.id || '')}/restore`;
    const url = path + api.queryString(req, [`id`]);
    return api.post(url, undefined);
  }
  
  // Delete a backup identified by the given ID, after which removal of any remote storage of the backup is started.
  // Note that the backup are initially only marked for deletion.
  // Once all remote storage for the backup has been removed, the backup itself is removed.
  // Required permissions:
  // -  backup.backup.delete on the backup identified by the given ID.
  async DeleteBackup(req: arangodb_cloud_common_v1_IDOptions): Promise<void> {
    const path = `/api/backup/v1/backup/${encodeURIComponent(req.id || '')}`;
    const url = path + api.queryString(req, [`id`]);
    return api.delete(url, undefined);
  }
}

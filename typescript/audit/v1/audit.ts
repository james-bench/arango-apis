//
// This file is AUTO-GENERATED by protoc-gen-ts.
// Do not modify it manually.
///
import api from '../../api'
import * as googleTypes from '../../googleTypes'
import { Empty as arangodb_cloud_common_v1_Empty } from '../../common/v1/common'
import { IDOptions as arangodb_cloud_common_v1_IDOptions } from '../../common/v1/common'
import { ListOptions as arangodb_cloud_common_v1_ListOptions } from '../../common/v1/common'
import { Version as arangodb_cloud_common_v1_Version } from '../../common/v1/common'

// File: audit/v1/audit.proto
// Package: arangodb.cloud.audit.v1

// Request arguments for AttachProjectToAuditLog.
export interface AttachProjectToAuditLogRequest {
  // ID of project to attach the AuditLog to.
  // string
  project_id?: string;
  
  // ID of the AuditLog to attach.
  // string
  auditlog_id?: string;
}

// AuditLog holds a specification for filtering audit events, a specification for
// destinations that audit events should be sent to and it acts as a grouping
// of audit log archives.
export interface AuditLog {
  // The ID of this resource.
  // string
  id?: string;
  
  // URL of this resource
  // This is a read-only value.
  // string
  url?: string;
  
  // Name of the audit log
  // string
  name?: string;
  
  // Description of the audit log
  // string
  description?: string;
  
  // The creation timestamp of the resource
  // This is a read-only value.
  // googleTypes.Timestamp
  created_at?: googleTypes.Timestamp;
  
  // The deletion timestamp of the resource
  // This is a read-only value.
  // googleTypes.Timestamp
  deleted_at?: googleTypes.Timestamp;
  
  // Set when this resource is deleted.
  // This is a read-only value.
  // boolean
  is_deleted?: boolean;
  
  // Identifier of the user who created this resource.
  // This is a read-only value.
  // string
  created_by_id?: string;
  
  // Identifier of the organization that owns this audit log.
  // This is a read-only value.
  // string
  organization_id?: string;
  
  // If set, this AuditLog is the default for the organization.
  // boolean
  is_default?: boolean;
  
  // Audit event topic filter
  // AuditLog_Filter
  filter?: AuditLog_Filter;
  
  // Destinations that events of this AuditLog should be sent to.
  // Note that there can only be 1 destination of type "cloud".
  // AuditLog_Destination
  destinations?: AuditLog_Destination[];
}

// Specification of a destination for audit events.
export interface AuditLog_Destination {
  // Type of destination.
  // Possible values are: "cloud", "https-post"
  // string
  type?: string;
}

// Event filter specification
export interface AuditLog_Filter {
  // If non-empty, only audit events with these topics will pass
  // the filter.
  // string
  included_topics?: string[];
  
  // Audit events with these topics will not pass the filter.
  // If a topic is in both included_topics & excluded_topics,
  // it will not pass the filter.
  // string
  excluded_topics?: string[];
}

// AuditLogArchive collects files of audit events in a specific region,
// usually for a specific deployment.
export interface AuditLogArchive {
  // The ID of this resource.
  // string
  id?: string;
  
  // URL of this resource
  // This is a read-only value.
  // string
  url?: string;
  
  // The creation timestamp of the resource
  // This is a read-only value.
  // googleTypes.Timestamp
  created_at?: googleTypes.Timestamp;
  
  // The deletion timestamp of the resource
  // This is a read-only value.
  // googleTypes.Timestamp
  deleted_at?: googleTypes.Timestamp;
  
  // Set when this resource is deleted.
  // This is a read-only value.
  // boolean
  is_deleted?: boolean;
  
  // Identifier of the auditlog that owns this audit log archive.
  // This is a read-only value.
  // string
  auditlog_id?: string;
  
  // If set, this archive is collecting audit events for a deployment with this ID.
  // Note that the deployment may have already been deleted.
  // string
  deployment_id?: string;
}

// AuditLogArchiveChunk collects audit events in a specific time frame
// on an AuditLogArchive.
export interface AuditLogArchiveChunk {
  // The ID of this resource.
  // string
  id?: string;
  
  // URL of this resource
  // This is a read-only value.
  // string
  url?: string;
  
  // The creation timestamp of the resource
  // This is a read-only value.
  // googleTypes.Timestamp
  created_at?: googleTypes.Timestamp;
  
  // Identifier of the AuditLogArchive that owns this chunk.
  // This is a read-only value.
  // string
  auditlogarchive_id?: string;
  
  // The timestamp of the first event in the chunk.
  // This is a read-only value.
  // googleTypes.Timestamp
  first_event?: googleTypes.Timestamp;
  
  // The timestamp of the last event in the chunk.
  // This is a read-only value.
  // googleTypes.Timestamp
  last_event?: googleTypes.Timestamp;
}

// List of AuditLogArchiveChunk's.
export interface AuditLogArchiveChunkList {
  // AuditLogArchiveChunk
  items?: AuditLogArchiveChunk[];
}

// List of AuditLogArchive's.
export interface AuditLogArchiveList {
  // AuditLogArchive
  items?: AuditLogArchive[];
}

// Single audit log event
export interface AuditLogEvent {
  // When did the event happen
  // googleTypes.Timestamp
  timestamp?: googleTypes.Timestamp;
  
  // Topic of the event
  // string
  topic?: string;
  
  // ID of the project that the event happened in (if applicable)
  // string
  project_id?: string;
  
  // ID of the deployment that the event happened in (if applicable)
  // string
  deployment_id?: string;
  
  // ID of the server that the event happened in (if applicable)
  // string
  server_id?: string;
  
  // Instance ID of the server that the event happened in (if applicable)
  // string
  instance_id?: string;
  
  // Sequence number. Must be even increasing for (deployment_id, server_id, instance_id) pairs.
  // number
  sequence?: number;
  
  // ID of the user that caused the event
  // string
  user_id?: string;
  
  // Name of database the operation is in (if applicable)
  // string
  database?: string;
  
  // IP address of source of operation (if available)
  // string
  client_ip?: string;
  
  // Authentication details
  // string
  authentication?: string;
  
  // Free format text describing the event
  // string
  message?: string;
}

// List of AuditLogEvent's.
export interface AuditLogEventList {
  // AuditLogEvent
  items?: AuditLogEvent[];
}

// List of AuditLog's.
export interface AuditLogList {
  // AuditLog
  items?: AuditLog[];
}

// Request arguments for ListAuditLogArchiveChunks
export interface ListAuditLogArchiveChunksRequest {
  // Identifier of the audit log archive to request the chunks for.
  // string
  auditlogarchive_id?: string;
  
  // Request chunks that contain events created at or after this timestamp.
  // This is an optional field.
  // googleTypes.Timestamp
  from?: googleTypes.Timestamp;
  
  // Request chunks that contain events created before this timestamp.
  // This is an optional field.
  // googleTypes.Timestamp
  to?: googleTypes.Timestamp;
  
  // Optional common list options, the context_id is ignored
  // arangodb.cloud.common.v1.ListOptions
  options?: arangodb_cloud_common_v1_ListOptions;
}

// Request arguments for ListAuditLogArchives
export interface ListAuditLogArchivesRequest {
  // Identifier of the audit log to request the audit log archives for.
  // string
  auditlog_id?: string;
  
  // If set, the result includes all audit log archives, including those who set to deleted,
  // however are not removed from the system currently.
  // If not set, only audit log archives not indicated as deleted are returned.
  // boolean
  include_deleted?: boolean;
  
  // Optional common list options, the context_id is ignored
  // arangodb.cloud.common.v1.ListOptions
  options?: arangodb_cloud_common_v1_ListOptions;
}

// Request arguments for ListAuditLogEvents.
export interface ListAuditLogEventsRequest {
  // Identifier of the audit log to request events for.
  // string
  auditlog_id?: string;
  
  // If set, include only events from this AuditLogArchive.
  // string
  auditlogarchive_id?: string;
  
  // If set, include only events from this AuditLogArchiveChunk.
  // string
  auditlogarchivechunk_id?: string;
  
  // Request events created at or after this timestamp.
  // This is an optional field.
  // googleTypes.Timestamp
  from?: googleTypes.Timestamp;
  
  // Request events created before this timestamp.
  // This is an optional field.
  // googleTypes.Timestamp
  to?: googleTypes.Timestamp;
  
  // Optional common list options, the context_id is ignored
  // arangodb.cloud.common.v1.ListOptions
  options?: arangodb_cloud_common_v1_ListOptions;
}

// Request arguments for ListAuditLogs
export interface ListAuditLogsRequest {
  // Identifier of the organization to request the audit logs for.
  // string
  organization_id?: string;
  
  // If set, the result includes all audit logs, including those who set to deleted,
  // however are not removed from the system currently.
  // If not set, only audit logs not indicated as deleted are returned.
  // boolean
  include_deleted?: boolean;
  
  // Optional common list options, the context_id is ignored
  // arangodb.cloud.common.v1.ListOptions
  options?: arangodb_cloud_common_v1_ListOptions;
}

// AuditService is the API used to provide access to audit events.
export interface IAuditService {
  // Get the current API version of this service.
  // Required permissions:
  // - None
  GetAPIVersion: (req?: arangodb_cloud_common_v1_Empty) => Promise<arangodb_cloud_common_v1_Version>;
  
  // Fetch all audit logs in the organization identified by the given ID.
  // Required permissions:
  // - audit.auditlog.list on the organization identified by the given ID.
  ListAuditLogs: (req: ListAuditLogsRequest) => Promise<AuditLogList>;
  
  // Fetch a specific AuditLog identified by the given ID.
  // Required permissions:
  // - audit.auditlog.get on the audit log identified by the given ID.
  GetAuditLog: (req: arangodb_cloud_common_v1_IDOptions) => Promise<AuditLog>;
  
  // Create a new audit log.
  // Required permissions:
  // - audit.auditlog.create on the organization identified by the given ID.
  CreateAuditLog: (req: AuditLog) => Promise<AuditLog>;
  
  // Update an audit log.
  // Required permissions:
  // - audit.auditlog.update on the audit log identified by the given ID.
  UpdateAuditLog: (req: AuditLog) => Promise<AuditLog>;
  
  // Delete an audit log.
  // Note that audit logs are initially only marked for deleted.
  // Once all their resources are removed the audit log itself is deleted
  // and cannot be restored.
  // Required permissions:
  // - audit.auditlog.delete on the audit log.
  DeleteAuditLog: (req: arangodb_cloud_common_v1_IDOptions) => Promise<void>;
  
  // Fetch all audit log archives in the audit log identified by the given ID.
  // Required permissions:
  // - audit.auditlogarchive.list on the audit log identified by the given ID.
  ListAuditLogArchives: (req: ListAuditLogArchivesRequest) => Promise<AuditLogArchiveList>;
  
  // Fetch a specific AuditLogArchive identified by the given ID.
  // Required permissions:
  // - audit.auditlogarchive.get on the audit log archive identified by the given ID.
  GetAuditLogArchive: (req: arangodb_cloud_common_v1_IDOptions) => Promise<AuditLogArchive>;
  
  // Delete an audit log archive.
  // Note that audit log archives are initially only marked for deleted.
  // Once all their resources are removed the audit log archive itself is deleted
  // and cannot be restored.
  // Required permissions:
  // - audit.auditlogarchive.delete on the audit log archive.
  DeleteAuditLogArchive: (req: arangodb_cloud_common_v1_IDOptions) => Promise<void>;
  
  // Fetch all chunks of an audit log archive identified by the given ID.
  // Required permissions:
  // - audit.auditlogarchivechunk.list on the audit log archive identified by the given ID.
  ListAuditLogArchiveChunks: (req: ListAuditLogArchiveChunksRequest) => Promise<AuditLogArchiveChunkList>;
  
  // Fetch a specific AuditLogArchiveChunk identified by the given ID.
  // Required permissions:
  // - audit.auditlogarchivechunk.get on the audit log archive chunk identified by the given ID.
  GetAuditLogArchiveChunk: (req: arangodb_cloud_common_v1_IDOptions) => Promise<AuditLogArchiveChunk>;
  
  // Delete an audit log archive chunk.
  // Required permissions:
  // - audit.auditlogarchivechunk.delete on the audit log archive chunk.
  DeleteAuditLogArchiveChunk: (req: arangodb_cloud_common_v1_IDOptions) => Promise<void>;
  
  // Fetch all audit events that match the given filter.
  // Note that this method will return a precondition-failed error
  // if there is no destination of type "cloud" in the AuditLog.
  // Required permissions:
  // - audit.auditlogevent.list on the audit log identified by the given ID.
  ListAuditLogEvents: (req: ListAuditLogEventsRequest) => Promise<AuditLogEventList>;
  
  // Fetch the AuditLog that is attached to the project identified by the given ID.
  // If no AuditLog is attached to the project, a not-found error is returned.
  // Required permissions:
  // - audit.auditlogattachment.get on the project identified by the given ID.
  GetAuditLogAttachedToProject: (req: arangodb_cloud_common_v1_IDOptions) => Promise<AuditLog>;
  
  // Attach the AuditLog identified by given ID to the project identified with given ID.
  // This replaces any existing AuditLog attachment for the project.
  // Required permissions:
  // - audit.auditlogattachment.create on the project identified by the given ID.
  AttachProjectToAuditLog: (req: AttachProjectToAuditLogRequest) => Promise<void>;
  
  // Detach the current AuditLog from the project identified with given ID.
  // After a detachment, no audit events in the context of the project will be sent
  // to an AuditLog.
  // Required permissions:
  // - audit.auditlogattachment.delete on the project identified by the given ID.
  DetachProjectFromAuditLog: (req: arangodb_cloud_common_v1_IDOptions) => Promise<void>;
}

// AuditService is the API used to provide access to audit events.
export class AuditService implements IAuditService {
  // Get the current API version of this service.
  // Required permissions:
  // - None
  async GetAPIVersion(req?: arangodb_cloud_common_v1_Empty): Promise<arangodb_cloud_common_v1_Version> {
    const path = `/api/audit/v1/api-version`;
    const url = path + api.queryString(req, []);
    return api.get(url, undefined);
  }
  
  // Fetch all audit logs in the organization identified by the given ID.
  // Required permissions:
  // - audit.auditlog.list on the organization identified by the given ID.
  async ListAuditLogs(req: ListAuditLogsRequest): Promise<AuditLogList> {
    const path = `/api/audit/v1/organizations/${encodeURIComponent(req.organization_id || '')}/auditlogs`;
    const url = path + api.queryString(req, [`organization_id`]);
    return api.get(url, undefined);
  }
  
  // Fetch a specific AuditLog identified by the given ID.
  // Required permissions:
  // - audit.auditlog.get on the audit log identified by the given ID.
  async GetAuditLog(req: arangodb_cloud_common_v1_IDOptions): Promise<AuditLog> {
    const path = `/api/audit/v1/auditlogs/${encodeURIComponent(req.id || '')}`;
    const url = path + api.queryString(req, [`id`]);
    return api.get(url, undefined);
  }
  
  // Create a new audit log.
  // Required permissions:
  // - audit.auditlog.create on the organization identified by the given ID.
  async CreateAuditLog(req: AuditLog): Promise<AuditLog> {
    const url = `/api/audit/v1/organizations/${encodeURIComponent(req.organization_id || '')}/auditlogs`;
    return api.post(url, req);
  }
  
  // Update an audit log.
  // Required permissions:
  // - audit.auditlog.update on the audit log identified by the given ID.
  async UpdateAuditLog(req: AuditLog): Promise<AuditLog> {
    const url = `/api/audit/v1/auditlogs/${encodeURIComponent(req.id || '')}`;
    return api.patch(url, req);
  }
  
  // Delete an audit log.
  // Note that audit logs are initially only marked for deleted.
  // Once all their resources are removed the audit log itself is deleted
  // and cannot be restored.
  // Required permissions:
  // - audit.auditlog.delete on the audit log.
  async DeleteAuditLog(req: arangodb_cloud_common_v1_IDOptions): Promise<void> {
    const path = `/api/audit/v1/auditlogs/${encodeURIComponent(req.id || '')}`;
    const url = path + api.queryString(req, [`id`]);
    return api.delete(url, undefined);
  }
  
  // Fetch all audit log archives in the audit log identified by the given ID.
  // Required permissions:
  // - audit.auditlogarchive.list on the audit log identified by the given ID.
  async ListAuditLogArchives(req: ListAuditLogArchivesRequest): Promise<AuditLogArchiveList> {
    const path = `/api/audit/v1/auditlogs/${encodeURIComponent(req.auditlog_id || '')}/auditlogarchives`;
    const url = path + api.queryString(req, [`auditlog_id`]);
    return api.get(url, undefined);
  }
  
  // Fetch a specific AuditLogArchive identified by the given ID.
  // Required permissions:
  // - audit.auditlogarchive.get on the audit log archive identified by the given ID.
  async GetAuditLogArchive(req: arangodb_cloud_common_v1_IDOptions): Promise<AuditLogArchive> {
    const path = `/api/audit/v1/auditlogarchives/${encodeURIComponent(req.id || '')}`;
    const url = path + api.queryString(req, [`id`]);
    return api.get(url, undefined);
  }
  
  // Delete an audit log archive.
  // Note that audit log archives are initially only marked for deleted.
  // Once all their resources are removed the audit log archive itself is deleted
  // and cannot be restored.
  // Required permissions:
  // - audit.auditlogarchive.delete on the audit log archive.
  async DeleteAuditLogArchive(req: arangodb_cloud_common_v1_IDOptions): Promise<void> {
    const path = `/api/audit/v1/auditlogarchives/${encodeURIComponent(req.id || '')}`;
    const url = path + api.queryString(req, [`id`]);
    return api.delete(url, undefined);
  }
  
  // Fetch all chunks of an audit log archive identified by the given ID.
  // Required permissions:
  // - audit.auditlogarchivechunk.list on the audit log archive identified by the given ID.
  async ListAuditLogArchiveChunks(req: ListAuditLogArchiveChunksRequest): Promise<AuditLogArchiveChunkList> {
    const path = `/api/audit/v1/auditlogarchives/${encodeURIComponent(req.auditlogarchive_id || '')}/auditlogarchivechunks`;
    const url = path + api.queryString(req, [`auditlogarchive_id`]);
    return api.get(url, undefined);
  }
  
  // Fetch a specific AuditLogArchiveChunk identified by the given ID.
  // Required permissions:
  // - audit.auditlogarchivechunk.get on the audit log archive chunk identified by the given ID.
  async GetAuditLogArchiveChunk(req: arangodb_cloud_common_v1_IDOptions): Promise<AuditLogArchiveChunk> {
    const path = `/api/audit/v1/auditlogarchivechunks/${encodeURIComponent(req.id || '')}`;
    const url = path + api.queryString(req, [`id`]);
    return api.get(url, undefined);
  }
  
  // Delete an audit log archive chunk.
  // Required permissions:
  // - audit.auditlogarchivechunk.delete on the audit log archive chunk.
  async DeleteAuditLogArchiveChunk(req: arangodb_cloud_common_v1_IDOptions): Promise<void> {
    const path = `/api/audit/v1/auditlogarchivechunks/${encodeURIComponent(req.id || '')}`;
    const url = path + api.queryString(req, [`id`]);
    return api.delete(url, undefined);
  }
  
  // Fetch all audit events that match the given filter.
  // Note that this method will return a precondition-failed error
  // if there is no destination of type "cloud" in the AuditLog.
  // Required permissions:
  // - audit.auditlogevent.list on the audit log identified by the given ID.
  async ListAuditLogEvents(req: ListAuditLogEventsRequest): Promise<AuditLogEventList> {
    const path = `/api/audit/v1/auditlogs/${encodeURIComponent(req.auditlog_id || '')}/events`;
    const url = path + api.queryString(req, [`auditlog_id`]);
    return api.get(url, undefined);
  }
  
  // Fetch the AuditLog that is attached to the project identified by the given ID.
  // If no AuditLog is attached to the project, a not-found error is returned.
  // Required permissions:
  // - audit.auditlogattachment.get on the project identified by the given ID.
  async GetAuditLogAttachedToProject(req: arangodb_cloud_common_v1_IDOptions): Promise<AuditLog> {
    const path = `/api/audit/v1/projects/${encodeURIComponent(req.id || '')}/auditlog`;
    const url = path + api.queryString(req, [`id`]);
    return api.get(url, undefined);
  }
  
  // Attach the AuditLog identified by given ID to the project identified with given ID.
  // This replaces any existing AuditLog attachment for the project.
  // Required permissions:
  // - audit.auditlogattachment.create on the project identified by the given ID.
  async AttachProjectToAuditLog(req: AttachProjectToAuditLogRequest): Promise<void> {
    const url = `/api/audit/v1/projects/${encodeURIComponent(req.project_id || '')}/auditlogs/${encodeURIComponent(req.auditlog_id || '')}/attach`;
    return api.post(url, req);
  }
  
  // Detach the current AuditLog from the project identified with given ID.
  // After a detachment, no audit events in the context of the project will be sent
  // to an AuditLog.
  // Required permissions:
  // - audit.auditlogattachment.delete on the project identified by the given ID.
  async DetachProjectFromAuditLog(req: arangodb_cloud_common_v1_IDOptions): Promise<void> {
    const path = `/api/audit/v1/projects/${encodeURIComponent(req.id || '')}/auditlogs`;
    const url = path + api.queryString(req, [`id`]);
    return api.delete(url, undefined);
  }
}

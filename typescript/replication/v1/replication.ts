//
// This file is AUTO-GENERATED by protoc-gen-ts.
// Do not modify it manually.
///
import api from '../../api'
import * as googleTypes from '../../googleTypes'
import { Empty as arangodb_cloud_common_v1_Empty } from '../../common/v1/common'
import { IDOptions as arangodb_cloud_common_v1_IDOptions } from '../../common/v1/common'
import { Version as arangodb_cloud_common_v1_Version } from '../../common/v1/common'
import { Deployment as arangodb_cloud_data_v1_Deployment } from '../../data/v1/data'

// File: replication/v1/replication.proto
// Package: arangodb.cloud.replication.v1

// CloneDeploymentFromBackupRequest defines a request object for clone deployment call.
export interface CloneDeploymentFromBackupRequest {
  // The ID of the backup to clone a deployment from.
  // string
  backup_id?: string;
  
  // Target region.
  // This is an optional field
  // string
  region_id?: string;
  
  // This field must be set to the identifier of the current Terms&Conditions
  // when cloning a deployment.
  // If the tier of the organization does not require a non-empty Terms&Condition
  // identifier, this field may be left empty.
  // If this field is not set the terms and conditions of the source deployment will be used.
  // string
  accepted_terms_and_conditions_id?: string;
  
  // Target project identifier.
  // This is an optional field
  // string
  project_id?: string;
}

// DeploymentMigration defines a request for performing the migration of a deployment.
export interface DeploymentMigration {
  // Identifier of the source deployment that needs to be migrated.
  // This is a required field.
  // string
  source_deployment_id?: string;
  
  // Specification of the target deployment.
  // DeploymentMigration_DeploymentSpec
  target_deployment?: DeploymentMigration_DeploymentSpec;
  
  // Timestamp of when this migration was initiated.
  // This is a read-only field.
  // googleTypes.Timestamp
  created_at?: googleTypes.Timestamp;
  
  // Status of the migration.
  // This is a read-only field.
  // DeploymentMigration_Status
  status?: DeploymentMigration_Status;
}

// Specification of the target deployment.
export interface DeploymentMigration_DeploymentSpec {
  // Type of model being used
  // string
  model?: string;
  
  // Size of nodes being used
  // This field is ignored set in case the flexible model is used.
  // string
  node_size_id?: string;
  
  // Number of nodes being used
  // This field is ignored set in case the flexible model is used.
  // number
  node_count?: number;
  
  // Amount of disk space per node (in GB)
  // This field is ignored set in case the flexible model is used.
  // number
  node_disk_size?: number;
  
  // Identifier of the region in which the deployment is created.
  // After creation, this value cannot be changed.
  // string
  region_id?: string;
}

// Status of the DeploymentMigration.
export interface DeploymentMigration_Status {
  // The current phase of the migration.
  // This will contain only one of the following values:
  // - SourceBackupInProgress:                Creation of backup of source deployment is in progress
  // - TargetDeploymentCreationInProgress:    Creation of target deployment is in progress
  // - TargetDeploymentModelChangeInProgress: The model of the target deployment is being updated.
  // - Failed:                                Migration process has failed
  // - Complete:                              Migration process has completed
  // string
  phase?: string;
  
  // Additional information regarding the phase
  // string
  description?: string;
  
  // Timestamp of when the status was last updated.
  // googleTypes.Timestamp
  last_updated_at?: googleTypes.Timestamp;
  
  // ID of the backup at the source deployment.
  // This backup will be used to perform a restore at the target deployment.
  // string
  backup_id?: string;
  
  // ID of the target deployment.
  // string
  target_deployment_id?: string;
}

// DeploymentReplication defines a request object for creating or updating a deployment replication
export interface DeploymentReplication {
  // Identifier of the deployment for a given DeploymentReplication
  // string
  deployment_id?: string;
  
  // Start the replication process for the given deployment.
  // boolean
  started?: boolean;
  
  // A PEM encoded representation of the public key of the CA certificate used to verify sync master in source deployment.
  // string
  certificate_pem?: string;
  
  // A PEM encoded representation of the keyfile used for client authentication of the sync master (with the sync master in the source deployment).
  // A keyfile contains 1 or more certificates and a private key.
  // string
  tls_keyfile?: string;
  
  // Identifier of the user that initiated this deployment replication.
  // This is a read-only value.
  // string
  started_by_id?: string;
  
  // CancelationOptions describes what to do during cancellation process of the migration-agent.
  // DeploymentReplication_CancelationOptions
  cancelation_options?: DeploymentReplication_CancelationOptions;
  
  // Status of the DeploymentReplication.
  // DeploymentReplication_Status
  status?: DeploymentReplication_Status;
}

// CancelationOptions describes what to do during cancellation process of the migration-agent.
export interface DeploymentReplication_CancelationOptions {
  // If set, during cancellation process data consistency is not required (otherwise data consistency is required).
  // boolean
  data_consistency_not_required?: boolean;
  
  // If set, after cancellation the source deployment will be in read-only mode.
  // boolean
  make_source_deployment_read_only?: boolean;
}

// DeploymentReplicationStatus defines the status of a deployment replication.
// Note: All fields in this message block are read-only.
export interface DeploymentReplication_Status {
  // Where the deployment replication process is in its lifecycle at any given time.
  // Should only contain only one of the following values:
  // "Initialising"   - Replication has started, waiting for sync masters / workers.
  // "In-Progress"    - Replication has started and currently in progress.
  // "Error"          - Replication is in an errored state.
  // "Failed"         - Replication could not complete successfully.
  // "Stopping"       - Replication is being stopped.
  // "Completed"      - Replication is stopped and all resources cleaned up properly.
  // string
  phase?: string;
  
  // Supporting information about the deployment replication phase - such as error messages in case of failures.
  // This field will be in JSON format and can be built using the `AsJSON()` helper.
  // Use `FromJSON()` helper to parse this field.
  // See - replication/v1/message.go in this repository.
  // string
  message?: string;
  
  // Service (LoadBalancer) endpoint of the SyncMasters
  // This field has the format of a URL.
  // This is a readonly field.
  // string
  sync_endpoint?: string;
  
  // The timestamp of when the Phase of the Deployment Replication was last updated.
  // This is a readonly field.
  // googleTypes.Timestamp
  phase_updated_at?: googleTypes.Timestamp;
  
  // Service (LoadBalancer) endpoint of the Forwarder service which allows to start streaming connection.
  // This field has the format of a URL.
  // This is a readonly field.
  // string
  forwarder_endpoint?: string;
  
  // Progress of replication in percents (value from 0.0 to 1.0).
  // This is a readonly field.
  // number
  progress?: number;
}

// ReplicationService is the API used to replicate a deployment.
export interface IReplicationService {
  // Get the current API version of this service.
  // Required permissions:
  // - None
  GetAPIVersion: (req?: arangodb_cloud_common_v1_Empty) => Promise<arangodb_cloud_common_v1_Version>;
  
  // Takes a backup and creates a deployment from it. For all intents and purposes this new deployment
  // will be the same as the deployment at that exact moment when the backup was taken from it. This means that
  // the new deployment will be in the same project and use the same provider as the old deployment did. Optionally
  // a different region can be provided using the region id field on the request. Furthermore, the new deployment
  // will have the same server settings ( count, mode, replication factor ) as the old deployment did at the time
  // of taking the backup. After the new deployment successfully started, the backup will be used to restore the
  // data into the new deployment. The new deployment will have a different endpoint, and the password will also
  // be reset for it. All other user settings will remain the same.
  // The old deployment will not be touched.
  // Required permissions:
  // if project_id is specified
  // - backup.backup.get on the backup specified by backup_id in request
  // - replication.deployment.clone-from-backup on the project specified in request
  // if project_id is not specified
  // - replication.deployment.clone-from-backup on the backup specified by backup_id
  CloneDeploymentFromBackup: (req: CloneDeploymentFromBackupRequest) => Promise<arangodb_cloud_data_v1_Deployment>;
  
  // Get an existing DeploymentReplication using its deployment ID
  // Required permissions:
  // - replication.deploymentreplication.get
  GetDeploymentReplication: (req: arangodb_cloud_common_v1_IDOptions) => Promise<DeploymentReplication>;
  
  // Update an existing DeploymentReplication spec. If does not exist, this will create a new one.
  // This call expects the complete entity with the updated fields.
  // Required permissions:
  // - replication.deploymentreplication.update
  UpdateDeploymentReplication: (req: DeploymentReplication) => Promise<DeploymentReplication>;
  
  // Create a new deployment migration.
  // Note: currently migration is supported only for Deployments with 'free' model.
  // Required permissions:
  // - replication.deploymentmigration.create on the specified deployment ID
  CreateDeploymentMigration: (req: DeploymentMigration) => Promise<DeploymentMigration>;
  
  // Get info about the deployment migration for a deployment identified by the given ID.
  // Required permissions:
  // - replication.deploymentmigration.get on the specified deployment ID
  GetDeploymentMigration: (req: arangodb_cloud_common_v1_IDOptions) => Promise<DeploymentMigration>;
  
  // Delete an existing DeploymentMigration.
  // A DeploymentMigration may be deleted only if it is in COMPLETE or FAILED state.
  // Required permissions:
  // - replication.deploymentmigration.delete on the specified deployment ID.
  DeleteDeploymentMigration: (req: arangodb_cloud_common_v1_IDOptions) => Promise<void>;
}

// ReplicationService is the API used to replicate a deployment.
export class ReplicationService implements IReplicationService {
  // Get the current API version of this service.
  // Required permissions:
  // - None
  async GetAPIVersion(req?: arangodb_cloud_common_v1_Empty): Promise<arangodb_cloud_common_v1_Version> {
    const path = `/api/replication/v1/api-version`;
    const url = path + api.queryString(req, []);
    return api.get(url, undefined);
  }
  
  // Takes a backup and creates a deployment from it. For all intents and purposes this new deployment
  // will be the same as the deployment at that exact moment when the backup was taken from it. This means that
  // the new deployment will be in the same project and use the same provider as the old deployment did. Optionally
  // a different region can be provided using the region id field on the request. Furthermore, the new deployment
  // will have the same server settings ( count, mode, replication factor ) as the old deployment did at the time
  // of taking the backup. After the new deployment successfully started, the backup will be used to restore the
  // data into the new deployment. The new deployment will have a different endpoint, and the password will also
  // be reset for it. All other user settings will remain the same.
  // The old deployment will not be touched.
  // Required permissions:
  // if project_id is specified
  // - backup.backup.get on the backup specified by backup_id in request
  // - replication.deployment.clone-from-backup on the project specified in request
  // if project_id is not specified
  // - replication.deployment.clone-from-backup on the backup specified by backup_id
  async CloneDeploymentFromBackup(req: CloneDeploymentFromBackupRequest): Promise<arangodb_cloud_data_v1_Deployment> {
    const url = `/api/replication/v1/backup/${encodeURIComponent(req.backup_id || '')}/clone`;
    return api.post(url, req);
  }
  
  // Get an existing DeploymentReplication using its deployment ID
  // Required permissions:
  // - replication.deploymentreplication.get
  async GetDeploymentReplication(req: arangodb_cloud_common_v1_IDOptions): Promise<DeploymentReplication> {
    const path = `/api/replication/v1/deployment/${encodeURIComponent(req.id || '')}/replication`;
    const url = path + api.queryString(req, [`id`]);
    return api.get(url, undefined);
  }
  
  // Update an existing DeploymentReplication spec. If does not exist, this will create a new one.
  // This call expects the complete entity with the updated fields.
  // Required permissions:
  // - replication.deploymentreplication.update
  async UpdateDeploymentReplication(req: DeploymentReplication): Promise<DeploymentReplication> {
    const url = `/api/replication/v1/deployment/${encodeURIComponent(req.deployment_id || '')}/replication`;
    return api.put(url, req);
  }
  
  // Create a new deployment migration.
  // Note: currently migration is supported only for Deployments with 'free' model.
  // Required permissions:
  // - replication.deploymentmigration.create on the specified deployment ID
  async CreateDeploymentMigration(req: DeploymentMigration): Promise<DeploymentMigration> {
    const path = `/api/replication/v1/deploymentmigration`;
    const url = path + api.queryString(req, []);
    return api.post(url, undefined);
  }
  
  // Get info about the deployment migration for a deployment identified by the given ID.
  // Required permissions:
  // - replication.deploymentmigration.get on the specified deployment ID
  async GetDeploymentMigration(req: arangodb_cloud_common_v1_IDOptions): Promise<DeploymentMigration> {
    const path = `/api/replication/v1/deploymentmigration/${encodeURIComponent(req.id || '')}`;
    const url = path + api.queryString(req, [`id`]);
    return api.get(url, undefined);
  }
  
  // Delete an existing DeploymentMigration.
  // A DeploymentMigration may be deleted only if it is in COMPLETE or FAILED state.
  // Required permissions:
  // - replication.deploymentmigration.delete on the specified deployment ID.
  async DeleteDeploymentMigration(req: arangodb_cloud_common_v1_IDOptions): Promise<void> {
    const path = `/api/replication/v1/deploymentmigration/${encodeURIComponent(req.id || '')}`;
    const url = path + api.queryString(req, [`id`]);
    return api.delete(url, undefined);
  }
}

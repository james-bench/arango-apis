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

// DeploymentReplication defines a request object for creating or updating a deployment replication
export interface DeploymentReplication {
  // Identifier of the deployments for a given DeploymentReplication
  // string
  deployment_id?: string;
  
  // Start the replication process for a given deployment.
  // boolean
  started?: boolean;
  
  // A PEM encoded representation of the public key of the CA certificate used to verify sync master in source datacenter.
  // The value stored here is base64 encoded.
  // string
  certificate_pem?: string;
  
  // A PEM encoded representation of the keyfile used for client authentication of the sync master (with the sync master in the source datacenter).
  // A keyfile contains 1 or more certificates and a private key.
  // The value stored here is base64 encoded.
  // string
  tls_keyfile?: string;
  
  // List of master endpoints at source data center.
  // string
  master_endpoint?: string[];
  
  // Status of the DeploymentReplication.
  // DeploymentReplication_Status
  status?: DeploymentReplication_Status;
}

// DeploymentReplicationStatus defines the status of a deployment replication.
export interface DeploymentReplication_Status {
  // Where the deployment replication process is in its lifecycle at any given time.
  // Should only contain only one of the following values:
  // "Initialising"   - Replication is supported, waiting for sync masters / workers.
  // "Initialised"    - Sync masters / workers are ready, deployment is ready to start replication process.
  // "In-Progress"    - Replication has started and currently in progress.
  // "In-Sync"        - All data currently is in-sync.
  // "Error"          - Replication is in an errored state.
  // "Failed"         - Replication could not complete successfully.
  // string
  phase?: string;
  
  // Supporting information about the deployment replication phase - such as error messages in case of failures.
  // string
  message?: string;
  
  // Total number of shards that should be in-sync.
  // number
  total_shards?: number;
  
  // Number of shards currently in-sync.
  // number
  shards_in_sync?: number;
  
  // Service (LoadBalancer) endpoint of the SyncMasters
  // This field has the format of a URL.
  // This is a readonly field.
  // string
  sync_endpoint?: string;
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
}

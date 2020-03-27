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

// File: example/v1/example.proto
// Package: arangodb.cloud.example.v1

// ExampleDataset represents a single example dataaset.
export interface ExampleDataset {
  // System identifier of the example dataset.
  // This is a read-only value.
  // string
  id?: string;
  
  // URL of this resource.
  // This is a read-only value.
  // string
  url?: string;
  
  // Name of the example dataset
  // string
  name?: string;
  
  // Description of the example dataset
  // string
  description?: string;
  
  // Guide contains description of the example dataset including several example queries.
  // Content type of guide is markdown.
  // string
  guide?: string;
  
  // The creation timestamp of the example dataset (database object)
  // This is a read-only value.
  // googleTypes.Timestamp
  created_at?: googleTypes.Timestamp;
}

// ExampleDatasetInstallation represents an installation of a dataset initiated by the user.
export interface ExampleDatasetInstallation {
  // System identifier of the installation.
  // This is a read-only value.
  // string
  id?: string;
  
  // URL of this resource
  // This is a read-only value.
  // string
  url?: string;
  
  // Identifier of the deployment that owns this installation.
  // After creation, this value cannot be changed.
  // string
  deployment_id?: string;
  
  // Identifier of the example dataset which will be installed.
  // After creation, this value cannot be changed.
  // string
  exampledataset_id?: string;
  
  // The creation timestamp of the installation.
  // googleTypes.Timestamp
  created_at?: googleTypes.Timestamp;
  
  // Status of the actual installation.
  // ExampleDatasetInstallation_Status
  status?: ExampleDatasetInstallation_Status;
}

// Status of the installation.
// All members of this field are read-only.
export interface ExampleDatasetInstallation_Status {
  // Name of the database into which this example dataset was installed.
  // string
  database_name?: string;
  
  // The state of the installation.
  // Will be one of the following: "InProgress|Create|Ready|Failed"
  // string
  state?: string;
  
  // Set when the installation is failed.
  // boolean
  is_failed?: boolean;
  
  // Set when once the installation finished successfully and the dataset is available to be used.
  // boolean
  is_available?: boolean;
}

// List of example datasets.
export interface ExampleDatasetInstallationList {
  // ExampleDatasetInstallation
  items?: ExampleDatasetInstallation[];
}

// List of example datasets.
export interface ExampleDatasetList {
  // ExampleDataset
  items?: ExampleDataset[];
}

// Request arguments for ListExampleDatasetInstallations.
export interface ListExampleDatasetInstallationsRequest {
  // Identifier of the deployment to request the installations for.
  // string
  deployment_id?: string;
  
  // Optional common list options, the context_id is ignored.
  // arangodb.cloud.common.v1.ListOptions
  options?: arangodb_cloud_common_v1_ListOptions;
}

// ListExampleDatasetsRequest provides an extendable list request for listing datasets.
export interface ListExampleDatasetsRequest {
  // Optional common list options, the context_id is ignored.
  // arangodb.cloud.common.v1.ListOptions
  options?: arangodb_cloud_common_v1_ListOptions;
}

// ExampleDatasetService is the API used to query for example datasets..
export interface IExampleDatasetService {
  // Get the current API version of this service.
  // Required permissions:
  // - None
  GetAPIVersion: (req?: arangodb_cloud_common_v1_Empty) => Promise<arangodb_cloud_common_v1_Version>;
  
  // Fetch all example datasets.
  // Required permissions:
  // - example.exampledataset.list
  ListExampleDatasets: (req: ListExampleDatasetsRequest) => Promise<ExampleDatasetList>;
  
  // Fetch an example dataset identified by the given ID.
  // Required permissions:
  // - example.exampledataset.get on the example identified by the given ID.
  GetExampleDataset: (req: arangodb_cloud_common_v1_IDOptions) => Promise<ExampleDataset>;
  
  // Fetch all installations for a specific deployment.
  // Required permissions:
  // - example.exampledatasetinstallation.list on the deployment that owns the installation and is identified by the given ID.
  ListExampleDatasetInstallations: (req: ListExampleDatasetInstallationsRequest) => Promise<ExampleDatasetInstallationList>;
  
  // Fetch an installations identified by the given ID.
  // Required permissions:
  // - example.exampledatasetinstallation.get on the installation identified by the given ID.
  GetExampleDatasetInstallation: (req: arangodb_cloud_common_v1_IDOptions) => Promise<ExampleDatasetInstallation>;
  
  // Create a new example installation. This represents a request made by the user to create an installation
  // for a deployment given a specified example dataset.
  // Required permissions:
  // -  example.exampledatasetinstallation.create on the deployment that the installation is for and is identified by the given ID.
  CreateExampleDatasetInstallation: (req: ExampleDatasetInstallation) => Promise<ExampleDatasetInstallation>;
  
  // Update an installation.
  // Required permissions:
  // -  example.exampledatasetinstallation.update on the installation identified by the given ID.
  UpdateExampleDatasetInstallation: (req: ExampleDatasetInstallation) => Promise<ExampleDatasetInstallation>;
  
  // Delete an installation identified by the given ID.
  // Required permissions:
  // -  example.exampledatasetinstallation.delete on the installation identified by the given ID.
  DeleteExampleDatasetInstallation: (req: arangodb_cloud_common_v1_IDOptions) => Promise<void>;
}

// ExampleDatasetService is the API used to query for example datasets..
export class ExampleDatasetService implements IExampleDatasetService {
  // Get the current API version of this service.
  // Required permissions:
  // - None
  async GetAPIVersion(req?: arangodb_cloud_common_v1_Empty): Promise<arangodb_cloud_common_v1_Version> {
    const path = `/api/example/v1/api-version`;
    const url = path + api.queryString(req, []);
    return api.get(url, undefined);
  }
  
  // Fetch all example datasets.
  // Required permissions:
  // - example.exampledataset.list
  async ListExampleDatasets(req: ListExampleDatasetsRequest): Promise<ExampleDatasetList> {
    const path = `/api/example/v1/exampledataset`;
    const url = path + api.queryString(req, []);
    return api.get(url, undefined);
  }
  
  // Fetch an example dataset identified by the given ID.
  // Required permissions:
  // - example.exampledataset.get on the example identified by the given ID.
  async GetExampleDataset(req: arangodb_cloud_common_v1_IDOptions): Promise<ExampleDataset> {
    const path = `/api/example/v1/exampledataset/${encodeURIComponent(req.id || '')}`;
    const url = path + api.queryString(req, [`id`]);
    return api.get(url, undefined);
  }
  
  // Fetch all installations for a specific deployment.
  // Required permissions:
  // - example.exampledatasetinstallation.list on the deployment that owns the installation and is identified by the given ID.
  async ListExampleDatasetInstallations(req: ListExampleDatasetInstallationsRequest): Promise<ExampleDatasetInstallationList> {
    const path = `/api/example/v1/deployment/${encodeURIComponent(req.deployment_id || '')}/exampledatasetinstallation`;
    const url = path + api.queryString(req, [`deployment_id`]);
    return api.get(url, undefined);
  }
  
  // Fetch an installations identified by the given ID.
  // Required permissions:
  // - example.exampledatasetinstallation.get on the installation identified by the given ID.
  async GetExampleDatasetInstallation(req: arangodb_cloud_common_v1_IDOptions): Promise<ExampleDatasetInstallation> {
    const path = `/api/example/v1/exampledatasetinstallation/${encodeURIComponent(req.id || '')}`;
    const url = path + api.queryString(req, [`id`]);
    return api.get(url, undefined);
  }
  
  // Create a new example installation. This represents a request made by the user to create an installation
  // for a deployment given a specified example dataset.
  // Required permissions:
  // -  example.exampledatasetinstallation.create on the deployment that the installation is for and is identified by the given ID.
  async CreateExampleDatasetInstallation(req: ExampleDatasetInstallation): Promise<ExampleDatasetInstallation> {
    const url = `/api/example/v1/deployment/${encodeURIComponent(req.deployment_id || '')}/exampledatasetinstallation`;
    return api.post(url, req);
  }
  
  // Update an installation.
  // Required permissions:
  // -  example.exampledatasetinstallation.update on the installation identified by the given ID.
  async UpdateExampleDatasetInstallation(req: ExampleDatasetInstallation): Promise<ExampleDatasetInstallation> {
    const url = `/api/example/v1/exampledatasetinstallation/${encodeURIComponent(req.id || '')}`;
    return api.patch(url, req);
  }
  
  // Delete an installation identified by the given ID.
  // Required permissions:
  // -  example.exampledatasetinstallation.delete on the installation identified by the given ID.
  async DeleteExampleDatasetInstallation(req: arangodb_cloud_common_v1_IDOptions): Promise<void> {
    const path = `/api/example/v1/exampledatasetinstallation/${encodeURIComponent(req.id || '')}`;
    const url = path + api.queryString(req, [`id`]);
    return api.delete(url, undefined);
  }
}

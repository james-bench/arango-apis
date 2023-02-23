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

// File: support/v1/support.proto
// Package: arangodb.cloud.support.v1

// FaqGroup contains groups of faq entries
export interface FaqGroup {
  // ID of the FAQ Group
  // string
  id?: string;
  
  // Name of the FAQ Group
  // string
  name?: string;
}

// FaqGroupEntry contains entries for a group
export interface FaqGroupEntry {
  // The question of this entry
  // string
  question?: string;
  
  // The answer to the question in this entry
  // string
  answer?: string;
}

// List of faq group entries.
export interface FaqGroupEntryList {
  // FaqGroupEntry
  items?: FaqGroupEntry[];
}

// List of faq groups.
export interface FaqGroupList {
  // FaqGroup
  items?: FaqGroup[];
}

// Arguments for a ListPlans request
export interface ListPlansRequest {
  // Common list options
  // arangodb.cloud.common.v1.ListOptions
  options?: arangodb_cloud_common_v1_ListOptions;
  
  // If set, list plans as they are available for the organization identified by this ID.
  // string
  organization_id?: string;
  
  // If set, list plans as they are available for the deployment model identified by this field.
  // string
  model?: string;
}

// Plan represents a specific support plan such as Bronze, Silver or Gold.
export interface Plan {
  // System identifier of the plan.
  // string
  id?: string;
  
  // Name of the plan.
  // string
  name?: string;
  
  // If set, this plan is the default support plan.
  // boolean
  is_default?: boolean;
  
  // Human readable description of the plan
  // string
  description?: string;
  
  // If set, this plan is shown, but not selectable.
  // boolean
  is_unavailable?: boolean;
  
  // SLA times to first response for various situations.
  // When this plan is unavailable, this field is optional.
  // ResponseTimes
  first_response_times?: ResponseTimes;
  
  // Support hours applicable to the SLA.
  // Possible values are "business" and "24x7".
  // string
  support_hours?: string;
}

// List of plans.
export interface PlanList {
  // Plan
  items?: Plan[];
}

// Response for various categories on situations.
// All values are in minutes.
// A value of 0 means "best effort".
export interface ResponseTimes {
  // Response time for operation-impeding Error in a production environment.
  // number
  critical?: number;
  
  // Response time for operation-limiting error.
  // number
  high?: number;
  
  // Response time for minor error.
  // number
  normal?: number;
  
  // Response time for usage question.
  // number
  low?: number;
}

// SupportRequest contains information about the request
export interface SupportRequest {
  // ID of the request
  // This is a read-only field
  // string
  id?: string;
  
  // Name of the user who submitted the support request
  // This is a required field
  // string
  user_name?: string;
  
  // User ID is only provided if the user submitting the request has authenticated
  // This is an optional field
  // string
  user_id?: string;
  
  // Email address of the user if applicable submitting the request
  // string
  email_address?: string;
  
  // Optional identifier of the organization that is the subject of the support request
  // string
  organization_id?: string;
  
  // Optional identifier of the project that is the subject of the support request
  // string
  project_id?: string;
  
  // Optional identifier of the deployment that is the subject of the support request
  // string
  deployment_id?: string;
  
  // Desciption which will contain the details provided by the submitter
  // string
  description?: string;
  
  // Severity of the request. Can be one of the following: (low|normal|high|critical)
  // string
  severity?: string;
  
  // Optional title which will contain the summary provided by the submitter (max 100 characters)
  // If not provided a title will be extracted from the description.
  // string
  title?: string;
}

// SupportService is the API used to query for support.
export interface ISupportService {
  // Get the current API version of this service.
  // Required permissions:
  // - None
  GetAPIVersion: (req?: arangodb_cloud_common_v1_Empty) => Promise<arangodb_cloud_common_v1_Version>;
  
  // Fetch all support plans that are supported by the ArangoDB cloud.
  // Required permissions:
  // - None
  ListPlans: (req: ListPlansRequest) => Promise<PlanList>;
  
  // Fetch a support plan by its id.
  // Required permissions:
  // - None
  GetPlan: (req: arangodb_cloud_common_v1_IDOptions) => Promise<Plan>;
  
  // Fetch all FAQ groups.
  // Required permissions:
  // - None
  ListFaqGroups: (req: arangodb_cloud_common_v1_ListOptions) => Promise<FaqGroupList>;
  
  // Fetch all FAQ group entries of the FAQ group identified by the given context ID.
  // Required permissions:
  // - None
  ListFaqGroupEntries: (req: arangodb_cloud_common_v1_ListOptions) => Promise<FaqGroupEntryList>;
  
  // Submit a support request.
  // Required permissions:
  // - resourcemanager.organization.get (if organization id is set)
  // - resourcemanager.project.get (if project id is set)
  // - data.deployment.get (if deployment id is set)
  // - None (if no ids are set)
  CreateSupportRequest: (req: SupportRequest) => Promise<SupportRequest>;
}

// SupportService is the API used to query for support.
export class SupportService implements ISupportService {
  // Get the current API version of this service.
  // Required permissions:
  // - None
  async GetAPIVersion(req?: arangodb_cloud_common_v1_Empty): Promise<arangodb_cloud_common_v1_Version> {
    const path = `/api/support/v1/api-version`;
    const url = path + api.queryString(req, []);
    return api.get(url, undefined);
  }
  
  // Fetch all support plans that are supported by the ArangoDB cloud.
  // Required permissions:
  // - None
  async ListPlans(req: ListPlansRequest): Promise<PlanList> {
    const path = `/api/support/v1/plans`;
    const url = path + api.queryString(req, []);
    return api.get(url, undefined);
  }
  
  // Fetch a support plan by its id.
  // Required permissions:
  // - None
  async GetPlan(req: arangodb_cloud_common_v1_IDOptions): Promise<Plan> {
    const path = `/api/support/v1/plans/${encodeURIComponent(req.id || '')}`;
    const url = path + api.queryString(req, [`id`]);
    return api.get(url, undefined);
  }
  
  // Fetch all FAQ groups.
  // Required permissions:
  // - None
  async ListFaqGroups(req: arangodb_cloud_common_v1_ListOptions): Promise<FaqGroupList> {
    const path = `/api/support/v1/faqgroups`;
    const url = path + api.queryString(req, []);
    return api.get(url, undefined);
  }
  
  // Fetch all FAQ group entries of the FAQ group identified by the given context ID.
  // Required permissions:
  // - None
  async ListFaqGroupEntries(req: arangodb_cloud_common_v1_ListOptions): Promise<FaqGroupEntryList> {
    const path = `/api/support/v1/faqgroups/${encodeURIComponent(req.context_id || '')}/entries`;
    const url = path + api.queryString(req, [`context_id`]);
    return api.get(url, undefined);
  }
  
  // Submit a support request.
  // Required permissions:
  // - resourcemanager.organization.get (if organization id is set)
  // - resourcemanager.project.get (if project id is set)
  // - data.deployment.get (if deployment id is set)
  // - None (if no ids are set)
  async CreateSupportRequest(req: SupportRequest): Promise<SupportRequest> {
    const url = `/api/support/v1/supportrequests`;
    return api.post(url, req);
  }
}

//
// This file is AUTO-GENERATED by protoc-gen-ts.
// Do not modify it manually.
///
import api from '../../api'
import * as googleTypes from '../../googleTypes'
import { Budget as arangodb_cloud_common_v1_Budget } from '../../common/v1/common'
import { IDOptions as arangodb_cloud_common_v1_IDOptions } from '../../common/v1/common'
import { ListOptions as arangodb_cloud_common_v1_ListOptions } from '../../common/v1/common'

// File: resourcemanager/v1/resourcemanager.proto
// Package: arangodb.cloud.resourcemanager.v1
export interface DataProcessingAddendum {
  // Identifier of this version of the DPA
  // string
  id?: string;
  
  // Content of DPA in markdown format
  // string
  content?: string;
  
  // Creation date of this version of the DPA.
  // googleTypes.Timestamp
  created_at?: googleTypes.Timestamp;
}

// An Event represents something that happened to an organization
// in the ArangoDB Managed service.
export interface Event {
  // System identifier of the event.
  // This is a read-only value.
  // string
  id?: string;
  
  // URL of this resource
  // This is a read-only value and cannot be initialized.
  // string
  url?: string;
  
  // Identifier of the organization that owns this event.
  // This is a read-only value.
  // string
  organization_id?: string;
  
  // Identifier of the subject of this event.
  // This is a read-only value.
  // If the subject of this event is an organization,
  // this value is a duplicate of organization_id.
  // string
  subject_id?: string;
  
  // Type of the event.
  // string
  type?: string;
  
  // Payload of the event.
  // The fields used in the payload are specific
  // to the type of event.
  // Event_PayloadEntry
  payload?: Event_PayloadEntry[];
  
  // The creation timestamp of the event
  // googleTypes.Timestamp
  created_at?: googleTypes.Timestamp;
  
  // URL of the subject of this event.
  // This is a read-only value.
  // string
  subject_url?: string;
  
  // If set, this event is not persisted.
  // This is a read-only value.
  // boolean
  volatile?: boolean;
}
export interface Event_PayloadEntry {
  // string
  key?: string;
  
  // string
  value?: string;
}

// List of Events.
export interface EventList {
  // Event
  items?: Event[];
}

// Request arguments for IsMemberOfOrganization.
export interface IsMemberOfOrganizationRequest {
  // Identifier of the user
  // string
  user_id?: string;
  
  // Identifier of the organization
  // string
  organization_id?: string;
}

// Response for IsMemberOfOrganization.
export interface IsMemberOfOrganizationResponse {
  // Set if the requested user is a member of the requested organization.
  // boolean
  member?: boolean;
  
  // Set if the requested user is an owner of the requested organization.
  // boolean
  owner?: boolean;
}

// Options for ListEvents
export interface ListEventOptions {
  // Standard list options
  // arangodb.cloud.common.v1.ListOptions
  options?: arangodb_cloud_common_v1_ListOptions;
  
  // If set, filter on the subject_id of event
  // string
  subject_ids?: string[];
  
  // If set, filter on the type of event
  // string
  types?: string[];
  
  // If set, filter of events created after this timestamp
  // googleTypes.Timestamp
  created_after?: googleTypes.Timestamp;
  
  // If set, filter of events created before this timestamp
  // googleTypes.Timestamp
  created_before?: googleTypes.Timestamp;
}

// Request arguments for ListXyzQuotas
export interface ListQuotasRequest {
  // Common list options
  // arangodb.cloud.common.v1.ListOptions
  options?: arangodb_cloud_common_v1_ListOptions;
  
  // If set, limit the returned list of quota's to these kinds.
  // string
  kinds?: string[];
}

// Member of an organization.
// A member is always a user.
export interface Member {
  // Identifier of the user
  // string
  user_id?: string;
  
  // Set if this user is owner of the organization
  // boolean
  owner?: boolean;
}

// List of Members.
export interface MemberList {
  // Member
  items?: Member[];
}

// An Organization is represents a real world organization such as a company.
export interface Organization {
  // System identifier of the organization.
  // This is a read-only value.
  // string
  id?: string;
  
  // URL of this resource
  // This is a read-only value and cannot be initialized.
  // string
  url?: string;
  
  // Name of the organization
  // string
  name?: string;
  
  // Description of the organization
  // string
  description?: string;
  
  // Set when this organization is deleted.
  // This is a read-only value.
  // boolean
  is_deleted?: boolean;
  
  // The creation timestamp of the organization
  // googleTypes.Timestamp
  created_at?: googleTypes.Timestamp;
  
  // The deletion timestamp of the organization
  // googleTypes.Timestamp
  deleted_at?: googleTypes.Timestamp;
  
  // Tier used for this organization.
  // This is a read-only value and cannot be initialized.
  // Tier
  tier?: Tier;
  
  // Total number of deployments created in this organization throughout
  // its entire lifetime per tier-id.
  // map: tier-id -> count
  // This is a read-only value.
  // Organization_TotalDeploymentsEntry
  total_deployments?: Organization_TotalDeploymentsEntry[];
  
  // If set, all projects in this organization are allowed to use deployments using the flexible model.
  // boolean
  is_flexible_deployments_enabled?: boolean;
}
export interface Organization_TotalDeploymentsEntry {
  // string
  key?: string;
  
  // number
  value?: number;
}

// An OrganizationInvite represents an invite for a human to join an
// organization.
export interface OrganizationInvite {
  // System identifier of the invite.
  // This is a read-only value.
  // string
  id?: string;
  
  // URL of this resource
  // This is a read-only value and cannot be initialized.
  // string
  url?: string;
  
  // Identifier of the organization that the human is invited to join.
  // This is a read-only value.
  // string
  organization_id?: string;
  
  // Email address of the human who is invited.
  // string
  email?: string;
  
  // If set, the invitee accepted the invite.
  // This is a read-only value.
  // boolean
  accepted?: boolean;
  
  // If set, the invitee rejected the invite.
  // This is a read-only value.
  // boolean
  rejected?: boolean;
  
  // The creation timestamp of the invite
  // This is a read-only value.
  // googleTypes.Timestamp
  created_at?: googleTypes.Timestamp;
  
  // The acceptance timestamp of the invite
  // This is a read-only value.
  // googleTypes.Timestamp
  accepted_at?: googleTypes.Timestamp;
  
  // The rejection timestamp of the invite
  // This is a read-only value.
  // googleTypes.Timestamp
  rejected_at?: googleTypes.Timestamp;
  
  // Identifier of the user that accepted or rejected this invite.
  // This is a read-only value.
  // string
  user_id?: string;
  
  // Identifier of the user that created this invite.
  // string
  created_by_id?: string;
  
  // Identifier of the organization that the human is invited to join.
  // This is a read-only value.
  // string
  organization_name?: string;
  
  // Name of the user that created this invite.
  // This is a read-only value.
  // string
  created_by_name?: string;
}

// List of OrganizationInvites.
export interface OrganizationInviteList {
  // OrganizationInvite
  items?: OrganizationInvite[];
}

// List of organizations.
export interface OrganizationList {
  // Actual organizations
  // Organization
  items?: Organization[];
  
  // Budget for organizations (owned by the caller)
  // arangodb.cloud.common.v1.Budget
  budget?: arangodb_cloud_common_v1_Budget;
}

// Request arguments for Add/DeleteOrganizationMembers.
export interface OrganizationMembersRequest {
  // Identifier of the organization to add/remove a user from
  // string
  organization_id?: string;
  
  // Users to add/remove.
  // For every user, an owner flag is provided as well.
  // If you add an existing user, the owner flag or the add request
  // will overwrite the value of the existing owner flag.
  // MemberList
  members?: MemberList;
}

// A Project is represents a unit within an organization such as a department.
export interface Project {
  // System identifier of the project.
  // This is a read-only value.
  // It can be set when creating the project.
  // string
  id?: string;
  
  // URL of this resource
  // This is a read-only value and cannot be initialized.
  // string
  url?: string;
  
  // Name of the project
  // string
  name?: string;
  
  // Description of the project
  // string
  description?: string;
  
  // Identifier of the organization that owns this project.
  // This is a read-only value.
  // string
  organization_id?: string;
  
  // Set when this project is deleted
  // boolean
  is_deleted?: boolean;
  
  // The creation timestamp of the project
  // googleTypes.Timestamp
  created_at?: googleTypes.Timestamp;
  
  // The deletion timestamp of the project
  // googleTypes.Timestamp
  deleted_at?: googleTypes.Timestamp;
  
  // If set, this project is allowed to use deployments using the flexible model.
  // boolean
  is_flexible_deployments_enabled?: boolean;
}

// List of Projects.
export interface ProjectList {
  // Resulting projects
  // Project
  items?: Project[];
  
  // Budget for projects
  // arangodb.cloud.common.v1.Budget
  budget?: arangodb_cloud_common_v1_Budget;
}

// Quota limit
export interface Quota {
  // Kind of quota
  // string
  kind?: string;
  
  // Human readable description of the quota
  // string
  description?: string;
  
  // Current limit of the quota.
  // A value of 0 means unlimited.
  // number
  limit?: number;
}

// Description of a kind of quota's
export interface QuotaDescription {
  // Kind of the quota
  // string
  kind?: string;
  
  // Human readable description
  // string
  description?: string;
  
  // If set, this kind of quota is valid at organization level
  // boolean
  for_organizations?: boolean;
  
  // If set, this kind of quota is valid at project level
  // boolean
  for_projects?: boolean;
}

// List of QuotaDescription's
export interface QuotaDescriptionList {
  // QuotaDescription
  items?: QuotaDescription[];
}

// List of Quota's
export interface QuotaList {
  // Quota
  items?: Quota[];
}
export interface TermsAndConditions {
  // Identifier of this version of the terms & conditions
  // string
  id?: string;
  
  // Content of terms & conditions in markdown format
  // string
  content?: string;
  
  // Creation date of this version of the terms & conditions.
  // googleTypes.Timestamp
  created_at?: googleTypes.Timestamp;
}

// Tier of an organization.
export interface Tier {
  // Identifier of the tier.
  // This is a read-only value and cannot be initialized.
  // string
  id?: string;
  
  // Human readable name of the tier.
  // This is a read-only value and cannot be initialized.
  // string
  name?: string;
  
  // If set the tier has support plans.
  // This is a read-only value and cannot be initialized.
  // boolean
  has_support_plans?: boolean;
  
  // If set the tier has backup uploads.
  // This is a read-only value and cannot be initialized.
  // boolean
  has_backup_uploads?: boolean;
  
  // If set, the tier requires that new deployments accept the
  // current terms & conditions.
  // This is a read-only value and cannot be initialized.
  // boolean
  requires_terms_and_conditions?: boolean;
}

// ResourceManagerService is the API used to configure basic resource objects.
export class ResourceManagerService {
  // Fetch all organizations that the authenticated user is a member of.
  // Required permissions:
  // - None
  async ListOrganizations(req: arangodb_cloud_common_v1_ListOptions): Promise<OrganizationList> {
    const path = `/api/resourcemanager/v1/self/organizations`;
    const url = path + api.queryString(req, []);
    return api.get(url, undefined);
  }
  
  // Fetch an organization by its id.
  // The authenticated user must be a member of the organization.
  // Required permissions:
  // - None
  async GetOrganization(req: arangodb_cloud_common_v1_IDOptions): Promise<Organization> {
    const path = `/api/resourcemanager/v1/organizations/${encodeURIComponent(req.id || '')}`;
    const url = path + api.queryString(req, [`id`]);
    return api.get(url, undefined);
  }
  
  // Create a new organization
  // Required permissions:
  // - None
  async CreateOrganization(req: Organization): Promise<Organization> {
    const url = `/api/resourcemanager/v1/organizations`;
    return api.post(url, req);
  }
  
  // Update an organization
  // Required permissions:
  // - resourcemanager.organization.update on the organization
  async UpdateOrganization(req: Organization): Promise<Organization> {
    const url = `/api/resourcemanager/v1/organizations/${encodeURIComponent(req.id || '')}`;
    return api.patch(url, req);
  }
  
  // Delete an organization
  // Note that organization are never really removed.
  // Instead their is_deleted field is set to true.
  // Required permissions:
  // - resourcemanager.organization.delete on the organization
  async DeleteOrganization(req: arangodb_cloud_common_v1_IDOptions): Promise<void> {
    const path = `/api/resourcemanager/v1/organizations/${encodeURIComponent(req.id || '')}`;
    const url = path + api.queryString(req, [`id`]);
    return api.delete(url, undefined);
  }
  
  // Get a list of members of the organization identified by the given context ID.
  // Required permissions:
  // - resourcemanager.organization.get on the organization
  async ListOrganizationMembers(req: arangodb_cloud_common_v1_ListOptions): Promise<MemberList> {
    const path = `/api/resourcemanager/v1/organizations/${encodeURIComponent(req.context_id || '')}/members`;
    const url = path + api.queryString(req, [`context_id`]);
    return api.get(url, undefined);
  }
  
  // Add one or more members to an organization.
  // If there are members (in the request arguments) that are already member of the
  // organization an AlreadyExists error is returned.
  // Required permissions:
  // - resourcemanager.organization.update on the organization
  async AddOrganizationMembers(req: OrganizationMembersRequest): Promise<void> {
    const url = `/api/resourcemanager/v1/organizations/${encodeURIComponent(req.organization_id || '')}/members`;
    return api.post(url, req);
  }
  
  // Update the ownership flag of one or more members of an organization.
  // If there are members (in the request arguments) that are not yet member of
  // the organization, an InvalidArgument error is returned.
  // If the request would result in the last owner no longer being an owner,
  // an InvalidArgument error is returned.
  // Required permissions:
  // - resourcemanager.organization.update on the organization
  async UpdateOrganizationMembers(req: OrganizationMembersRequest): Promise<void> {
    const url = `/api/resourcemanager/v1/organizations/${encodeURIComponent(req.organization_id || '')}/members`;
    return api.patch(url, req);
  }
  
  // Remove one or more members from an organization.
  // If the request would result in the last owner being removed as member
  // of the organization, an InvalidArgument error is returned.
  // Required permissions:
  // - resourcemanager.organization.update on the organization
  async DeleteOrganizationMembers(req: OrganizationMembersRequest): Promise<void> {
    const url = `/api/resourcemanager/v1/organizations/${encodeURIComponent(req.organization_id || '')}/members`;
    return api.delete(url, req);
  }
  
  // Is the user identified by the given user ID a member
  // of the organization identified by the given organization ID.
  // Required permissions:
  // - resourcemanager.organization.get on the organization, unless the requested user is identical to the authenticated user.
  // Note that if the identified user or organization does not exist, no is returned.
  async IsMemberOfOrganization(req: IsMemberOfOrganizationRequest): Promise<IsMemberOfOrganizationResponse> {
    const path = `/api/resourcemanager/v1/organizations/${encodeURIComponent(req.organization_id || '')}/members/${encodeURIComponent(req.user_id || '')}`;
    const url = path + api.queryString(req, [`organization_id`, `user_id`]);
    return api.get(url, undefined);
  }
  
  // Get a list of quota values for the organization identified by the given context ID.
  // If a quota is not specified on organization level, a (potentially tier specific) default
  // value is returned.
  // Required permissions:
  // - resourcemanager.organization.get on the organization
  async ListOrganizationQuotas(req: ListQuotasRequest): Promise<QuotaList> {
    const path = `/api/resourcemanager/v1/organizations/${encodeURIComponent((req.options || {}).context_id || '')}/quotas`;
    const url = path + api.queryString(req, [`options.context_id`]);
    return api.get(url, undefined);
  }
  
  // Fetch all projects in the organization identified by the given context ID.
  // The authenticated user must be a member of the organization identifier by the given context ID.
  // Required permissions:
  // - resourcemanager.project.list on the organization identified by the given context ID
  async ListProjects(req: arangodb_cloud_common_v1_ListOptions): Promise<ProjectList> {
    const path = `/api/resourcemanager/v1/organizations/${encodeURIComponent(req.context_id || '')}/projects`;
    const url = path + api.queryString(req, [`context_id`]);
    return api.get(url, undefined);
  }
  
  // Fetch a project by its id.
  // The authenticated user must be a member of the organization that owns the project.
  // Required permissions:
  // - resourcemanager.project.get on the project identified by the given ID
  async GetProject(req: arangodb_cloud_common_v1_IDOptions): Promise<Project> {
    const path = `/api/resourcemanager/v1/projects/${encodeURIComponent(req.id || '')}`;
    const url = path + api.queryString(req, [`id`]);
    return api.get(url, undefined);
  }
  
  // Create a new project
  // The authenticated user must be a member of the organization that owns the project.
  // Required permissions:
  // - resourcemanager.project.create on the organization that owns the project
  async CreateProject(req: Project): Promise<Project> {
    const url = `/api/resourcemanager/v1/organizations/${encodeURIComponent(req.organization_id || '')}/projects`;
    return api.post(url, req);
  }
  
  // Update a project
  // The authenticated user must be a member of the organization that owns the project.
  // Required permissions:
  // - resourcemanager.project.update on the project
  async UpdateProject(req: Project): Promise<Project> {
    const url = `/api/resourcemanager/v1/projects/${encodeURIComponent(req.id || '')}`;
    return api.patch(url, req);
  }
  
  // Delete a project
  // Note that project are initially only marked for deleted.
  // Once all their resources are removed the project itself is deleted
  // and cannot be restored.
  // The authenticated user must be a member of the organization that owns the project.
  // Required permissions:
  // - resourcemanager.project.delete on the project
  async DeleteProject(req: arangodb_cloud_common_v1_IDOptions): Promise<void> {
    const path = `/api/resourcemanager/v1/projects/${encodeURIComponent(req.id || '')}`;
    const url = path + api.queryString(req, [`id`]);
    return api.delete(url, undefined);
  }
  
  // Get a list of quota values for the project identified by the given context ID.
  // If a quota is not specified on project level, a value from organization level
  // is returned.
  // Required permissions:
  // - resourcemanager.project.get on the project
  async ListProjectQuotas(req: ListQuotasRequest): Promise<QuotaList> {
    const path = `/api/resourcemanager/v1/projects/${encodeURIComponent((req.options || {}).context_id || '')}/quotas`;
    const url = path + api.queryString(req, [`options.context_id`]);
    return api.get(url, undefined);
  }
  
  // Fetch all events in the organization identified by the given context ID.
  // The authenticated user must be a member of the organization identifier by the given context ID.
  // Required permissions:
  // - resourcemanager.event.list on the organization identified by the given context ID
  async ListEvents(req: ListEventOptions): Promise<EventList> {
    const path = `/api/resourcemanager/v1/organizations/${encodeURIComponent((req.options || {}).context_id || '')}/events`;
    const url = path + api.queryString(req, [`options.context_id`]);
    return api.get(url, undefined);
  }
  
  // Fetch all organization invites in the organization identified by the given context ID.
  // The authenticated user must be a member of the organization identifier by the given context ID.
  // Required permissions:
  // - resourcemanager.organization-invite.list on the invite.
  async ListOrganizationInvites(req: arangodb_cloud_common_v1_ListOptions): Promise<OrganizationInviteList> {
    const path = `/api/resourcemanager/v1/organizations/${encodeURIComponent(req.context_id || '')}/organization-invites`;
    const url = path + api.queryString(req, [`context_id`]);
    return api.get(url, undefined);
  }
  
  // Fetch all organization invites for the email address of the authenticated user.
  // Required permissions:
  // - None
  async ListMyOrganizationInvites(req: arangodb_cloud_common_v1_ListOptions): Promise<OrganizationInviteList> {
    const path = `/api/resourcemanager/v1/self/organization-invites`;
    const url = path + api.queryString(req, []);
    return api.get(url, undefined);
  }
  
  // Fetch an organization invite by its id.
  // The authenticated user must be a member of the organization that the invite is for.
  // Required permissions:
  // - resourcemanager.organization-invite.get on the invite.
  async GetOrganizationInvite(req: arangodb_cloud_common_v1_IDOptions): Promise<OrganizationInvite> {
    const path = `/api/resourcemanager/v1/organization-invites/${encodeURIComponent(req.id || '')}`;
    const url = path + api.queryString(req, [`id`]);
    return api.get(url, undefined);
  }
  
  // Create a new organization invite.
  // The authenticated user must be a member of the organization that the invite is for.
  // Required permissions:
  // - resourcemanager.organization-invite.create on the organization that the invite is for.
  async CreateOrganizationInvite(req: OrganizationInvite): Promise<OrganizationInvite> {
    const url = `/api/resourcemanager/v1/organizations/${encodeURIComponent(req.organization_id || '')}/organization-invites`;
    return api.post(url, req);
  }
  
  // Delete an organization invite
  // The authenticated user must be a member of the organization that the invite is for.
  // Required permissions:
  // - resourcemanager.organization-invite.delete on the invite
  async DeleteOrganizationInvite(req: arangodb_cloud_common_v1_IDOptions): Promise<void> {
    const path = `/api/resourcemanager/v1/organization-invites/${encodeURIComponent(req.id || '')}`;
    const url = path + api.queryString(req, [`id`]);
    return api.delete(url, undefined);
  }
  
  // Accept an organization invite
  // The authenticated user's email address must match the email address specified in
  // the invite.
  // Required permissions:
  // - None
  async AcceptOrganizationInvite(req: arangodb_cloud_common_v1_IDOptions): Promise<void> {
    const path = `/api/resourcemanager/v1/organization-invites/${encodeURIComponent(req.id || '')}/accept`;
    const url = path + api.queryString(req, [`id`]);
    return api.post(url, undefined);
  }
  
  // Reject an organization invite
  // The authenticated user's email address must match the email address specified in
  // the invite.
  // Required permissions:
  // - None
  async RejectOrganizationInvite(req: arangodb_cloud_common_v1_IDOptions): Promise<void> {
    const path = `/api/resourcemanager/v1/organization-invites/${encodeURIComponent(req.id || '')}/reject`;
    const url = path + api.queryString(req, [`id`]);
    return api.post(url, undefined);
  }
  
  // Fetch descriptions for all quota kinds know by the platform.
  // Required permissions:
  // - None
  async ListQuotaDescriptions(req: arangodb_cloud_common_v1_ListOptions): Promise<QuotaDescriptionList> {
    const path = `/api/resourcemanager/v1/quotas/descriptions`;
    const url = path + api.queryString(req, []);
    return api.get(url, undefined);
  }
  
  // Fetch a specific version of the Terms & Conditions.
  // Required permissions:
  // - None
  async GetTermsAndConditions(req: arangodb_cloud_common_v1_IDOptions): Promise<TermsAndConditions> {
    const path = `/api/resourcemanager/v1/termsandconditions/${encodeURIComponent(req.id || '')}`;
    const url = path + api.queryString(req, [`id`]);
    return api.get(url, undefined);
  }
  
  // Fetch the current version of the Terms & Conditions for the organization
  // identified by the given (optional) ID.
  // Required permissions:
  // - None If ID is empty.
  // - resourcemanager.organization.get If ID is not empty.
  async GetCurrentTermsAndConditions(req: arangodb_cloud_common_v1_IDOptions): Promise<TermsAndConditions> {
    const path = `/api/resourcemanager/v1/current-termsandconditions`;
    const url = path + api.queryString(req, []);
    return api.get(url, undefined);
  }
  
  // Fetch a specific version of the Data Processing Addendum.
  // Required permissions:
  // - None
  async GetDataProcessingAddendum(req: arangodb_cloud_common_v1_IDOptions): Promise<DataProcessingAddendum> {
    const path = `/api/resourcemanager/v1/dpa/${encodeURIComponent(req.id || '')}`;
    const url = path + api.queryString(req, [`id`]);
    return api.get(url, undefined);
  }
  
  // Fetch the current version of the Data Processing Addendum for the organization
  // identified by the given (optional) ID.
  // Required permissions:
  // - None If ID is empty.
  // - resourcemanager.organization.get If ID is not empty.
  async GetCurrentDataProcessingAddendum(req: arangodb_cloud_common_v1_IDOptions): Promise<DataProcessingAddendum> {
    const path = `/api/resourcemanager/v1/current-dpa`;
    const url = path + api.queryString(req, []);
    return api.get(url, undefined);
  }
}

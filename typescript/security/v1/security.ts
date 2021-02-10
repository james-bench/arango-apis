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

// File: security/v1/security.proto
// Package: arangodb.cloud.security.v1

// IAMProvider provides configuration for a custom Identity & Access management provider
// for deployments.
export interface IAMProvider {
  // System identifier of the provider.
  // This is a read-only value.
  // string
  id?: string;
  
  // URL of the provider.
  // This is a read-only value.
  // string
  url?: string;
  
  // Name of the provider.
  // string
  name?: string;
  
  // Description of the provider.
  // string
  description?: string;
  
  // Identifier of the project that contains this provider.
  // string
  project_id?: string;
  
  // Type of provider
  // string
  type?: string;
  
  // The creation timestamp of this provider.
  // This is a read-only value.
  // googleTypes.Timestamp
  created_at?: googleTypes.Timestamp;
  
  // The deletion timestamp of the provider
  // This is a read-only value.
  // googleTypes.Timestamp
  deleted_at?: googleTypes.Timestamp;
  
  // Set when this provider is deleted.
  // This is a read-only value.
  // boolean
  is_deleted?: boolean;
  
  // Identifier of the user who created this provider.
  // This is a read-only value.
  // string
  created_by_id?: string;
  
  // Set when this provider is the default in its project.
  // This is a read-only value.
  // boolean
  is_default?: boolean;
  
  // If set, this IAM provider cannot be deleted.
  // To delete, first update the with locked set to false.
  // boolean
  locked?: boolean;
  
  // IAMProvider_LDAPSettings
  ldap_settings?: IAMProvider_LDAPSettings;
}

// LDAP provider specific settings
export interface IAMProvider_LDAPSettings {
  // Hostname or IP address of the server
  // string
  server?: string;
  
  // Port number of the server (defaults to 389)
  // number
  port?: number;
  
  // Base distinguished name under which the search takes place
  // string
  base_distinguished_name?: string;
  
  // distinguished name for a read-only LDAP user to which ArangoDB can bind to search the LDAP server
  // string
  bind_distinguished_name?: string;
  
  // Password name for a read-only LDAP user to which ArangoDB can bind to search the LDAP server.
  // This is a set-only field. During get/list requests, this field will be empty.
  // string
  bind_password?: string;
  
  // Refresh rate in seconds (defaults to 300)
  // number
  refresh_rate?: number;
  
  // PEM encoded version of the CA certificate used by the LDAP server.
  // string
  tls_ca_certificate_pem?: string;
  
  // If set, calls into the underlying LDAP library are serialized.
  // This option can be used to work around thread-unsafe LDAP library functionality.
  // boolean
  serialized?: boolean;
  
  // Timeout (in seconds) used when waiting to enter the LDAP library call serialization lock.
  // This is only meaningful when serialized has been set to true.
  // number
  serialize_timeout_sec?: number;
  
  // Number of retries to attempt a connection to the LDAP server.
  // Setting this to values greater than zero will make ArangoDB retry to contact the
  // LDAP server in case no connection can be made initially.
  // number
  retries?: number;
  
  // If set, the LDAP library will implicitly restart connections.
  // boolean
  restart?: boolean;
  
  // If set, the LDAP library will implicitly chase referrals.
  // boolean
  referrals?: boolean;
  
  // Timeout value (in seconds) for synchronous LDAP API calls (a value of 0 means default timeout).
  // number
  timeout_sec?: number;
  
  // Timeout value (in seconds) after which network operations following the initial
  // connection return in case of no activity (a value of 0 means default timeout).
  // number
  network_timeout_sec?: number;
  
  // If set, the LDAP library will connect asynchronously.
  // boolean
  async_connect?: boolean;
  
  // Prefix for simple authentication
  // string
  prefix?: string;
  
  // Suffix for simple authentication
  // string
  suffix?: string;
  
  // LDAP search scope with possible values "base" (just search the base distinguished name),
  // "sub" (recursive search under the base distinguished name) or
  // "one" (search the base’s immediate children) (default: "sub").
  // string
  search_scope?: string;
  
  // LDAP filter expression which limits the set of LDAP users being considered
  // (default: "objectClass=*"" which means all objects).
  // string
  search_filter?: string;
  
  // Specifies the attribute in the user objects which is used to match the ArangoDB user name (default: "uid").
  // string
  search_attribute?: string;
  
  // If set, this field specifies the name of the attribute used to fetch the roles of a user.
  // string
  roles_attribute_name?: string;
  
  // If set, then the string {USER} in the value of this field is replaced with the distinguished
  // name of the authenticated LDAP user and the resulting search expression is used to
  // match distinguished names of LDAP objects representing roles of that user.
  // string
  roles_search?: string;
  
  // Regular expression that is used to filter roles.
  // Only roles that match the regular expression are used.
  // string
  roles_include?: string;
  
  // Regular expression that is used to filter roles.
  // Only roles that do not match the regular expression are used.
  // string
  roles_exclude?: string;
  
  // A regular expression in the format of a replacement text (/re/text/).
  // This regular expression is applied to the role name found.
  // This is especially useful in the roles-search variant to extract the real role name out of the dn value.
  // string
  roles_transformation?: string;
  
  // Name of role associated with the superuser.
  // Any user belonging to this role gains superuser status.
  // This role is checked after applying the roles_transformation expression.
  // string
  super_user_role?: string;
}

// List of IAM providers.
export interface IAMProviderList {
  // IAMProvider
  items?: IAMProvider[];
}

// IPAllowlist represents a list of CIDR ranges from which a deployment is accessible.
export interface IPAllowlist {
  // System identifier of the allowlist.
  // This is a read-only value.
  // string
  id?: string;
  
  // URL of the allowlist.
  // This is a read-only value.
  // string
  url?: string;
  
  // Name of the allowlist.
  // string
  name?: string;
  
  // Description of the allowlist.
  // string
  description?: string;
  
  // Identifier of the project that contains this allowlist.
  // string
  project_id?: string;
  
  // List of CIDR ranges.
  // Values must follow format as defined in RFC 4632 and RFC 4291.
  // string
  cidr_ranges?: string[];
  
  // The creation timestamp of this allowlist.
  // This is a read-only value.
  // googleTypes.Timestamp
  created_at?: googleTypes.Timestamp;
  
  // The deletion timestamp of the allowlist
  // This is a read-only value.
  // googleTypes.Timestamp
  deleted_at?: googleTypes.Timestamp;
  
  // Set when this allowlist is deleted.
  // This is a read-only value.
  // boolean
  is_deleted?: boolean;
  
  // Identifier of the user who created this allowlist.
  // This is a read-only value.
  // string
  created_by_id?: string;
  
  // If set, this allow list cannot be deleted.
  // To delete, first update the with locked set to false.
  // boolean
  locked?: boolean;
  
  // The list of warnings which are related to the IP allow list.
  // string
  warnings?: string[];
}

// List of IP allowlists.
export interface IPAllowlistList {
  // IPAllowlist
  items?: IPAllowlist[];
}

// SecurityService is the API used to access security entities.
export interface ISecurityService {
  // Get the current API version of this service.
  // Required permissions:
  // - None
  GetAPIVersion: (req?: arangodb_cloud_common_v1_Empty) => Promise<arangodb_cloud_common_v1_Version>;
  
  // Fetch all IP allowlists that belong to the project identified by the given
  // context ID.
  // Required permissions:
  // - security.ipallowlist.list on the project identified by the given context ID.
  ListIPAllowlists: (req: arangodb_cloud_common_v1_ListOptions) => Promise<IPAllowlistList>;
  
  // Fetch an IP allowlist by its id.
  // Required permissions:
  // - security.ipallowlist.get on the IP allowlist
  GetIPAllowlist: (req: arangodb_cloud_common_v1_IDOptions) => Promise<IPAllowlist>;
  
  // Create a new IP allowlist
  // Required permissions:
  // - security.ipallowlist.create on the project that owns the IP allowlist.
  CreateIPAllowlist: (req: IPAllowlist) => Promise<IPAllowlist>;
  
  // Update an IP allowlist
  // Required permissions:
  // - security.ipallowlist.update on the IP allowlist
  UpdateIPAllowlist: (req: IPAllowlist) => Promise<IPAllowlist>;
  
  // Delete an IP allowlist.
  // Note that IP allowlists are initially only marked for deletion.
  // Once all their dependent deployments are removed, the allowlist is removed.
  // Required permissions:
  // - security.ipallowlist.delete on the IP allowlist
  DeleteIPAllowlist: (req: arangodb_cloud_common_v1_IDOptions) => Promise<void>;
  
  // Fetch all IAM providers that belong to the project identified by the given
  // context ID.
  // Required permissions:
  // - security.iamprovider.list on the project identified by the given context ID.
  ListIAMProviders: (req: arangodb_cloud_common_v1_ListOptions) => Promise<IAMProviderList>;
  
  // Fetch an IAM provider by its id.
  // Required permissions:
  // - security.iamprovider.get on the IAM provider
  GetIAMProvider: (req: arangodb_cloud_common_v1_IDOptions) => Promise<IAMProvider>;
  
  // Create a new IAM provider
  // Required permissions:
  // - security.iamprovider.create on the project that owns the IAM provider.
  CreateIAMProvider: (req: IAMProvider) => Promise<IAMProvider>;
  
  // Update an IAM provider
  // Required permissions:
  // - security.iamprovider.update on the IAM provider
  UpdateIAMProvider: (req: IAMProvider) => Promise<IAMProvider>;
  
  // Delete an IAM provider.
  // Note that IAM providers are initially only marked for deletion.
  // Once all their dependent deployments are removed, the provider is removed.
  // Required permissions:
  // - security.iamprovider.delete on the IP whitelist
  DeleteIAMProvider: (req: arangodb_cloud_common_v1_IDOptions) => Promise<void>;
  
  // Mark the given IAM provider as default for its containing project.
  // Required permissions:
  // - security.iamprovider.set-default on the project that owns the provider.
  SetDefaultIAMProvider: (req: IAMProvider) => Promise<void>;
}

// SecurityService is the API used to access security entities.
export class SecurityService implements ISecurityService {
  // Get the current API version of this service.
  // Required permissions:
  // - None
  async GetAPIVersion(req?: arangodb_cloud_common_v1_Empty): Promise<arangodb_cloud_common_v1_Version> {
    const path = `/api/security/v1/api-version`;
    const url = path + api.queryString(req, []);
    return api.get(url, undefined);
  }
  
  // Fetch all IP allowlists that belong to the project identified by the given
  // context ID.
  // Required permissions:
  // - security.ipallowlist.list on the project identified by the given context ID.
  async ListIPAllowlists(req: arangodb_cloud_common_v1_ListOptions): Promise<IPAllowlistList> {
    const path = `/api/security/v1/projects/${encodeURIComponent(req.context_id || '')}/ipallowlists`;
    const url = path + api.queryString(req, [`context_id`]);
    return api.get(url, undefined);
  }
  
  // Fetch an IP allowlist by its id.
  // Required permissions:
  // - security.ipallowlist.get on the IP allowlist
  async GetIPAllowlist(req: arangodb_cloud_common_v1_IDOptions): Promise<IPAllowlist> {
    const path = `/api/security/v1/ipallowlists/${encodeURIComponent(req.id || '')}`;
    const url = path + api.queryString(req, [`id`]);
    return api.get(url, undefined);
  }
  
  // Create a new IP allowlist
  // Required permissions:
  // - security.ipallowlist.create on the project that owns the IP allowlist.
  async CreateIPAllowlist(req: IPAllowlist): Promise<IPAllowlist> {
    const url = `/api/security/v1/project/${encodeURIComponent(req.project_id || '')}/ipallowlists`;
    return api.post(url, req);
  }
  
  // Update an IP allowlist
  // Required permissions:
  // - security.ipallowlist.update on the IP allowlist
  async UpdateIPAllowlist(req: IPAllowlist): Promise<IPAllowlist> {
    const url = `/api/security/v1/ipallowlists/${encodeURIComponent(req.id || '')}`;
    return api.patch(url, req);
  }
  
  // Delete an IP allowlist.
  // Note that IP allowlists are initially only marked for deletion.
  // Once all their dependent deployments are removed, the allowlist is removed.
  // Required permissions:
  // - security.ipallowlist.delete on the IP allowlist
  async DeleteIPAllowlist(req: arangodb_cloud_common_v1_IDOptions): Promise<void> {
    const path = `/api/security/v1/ipallowlists/${encodeURIComponent(req.id || '')}`;
    const url = path + api.queryString(req, [`id`]);
    return api.delete(url, undefined);
  }
  
  // Fetch all IAM providers that belong to the project identified by the given
  // context ID.
  // Required permissions:
  // - security.iamprovider.list on the project identified by the given context ID.
  async ListIAMProviders(req: arangodb_cloud_common_v1_ListOptions): Promise<IAMProviderList> {
    const path = `/api/security/v1/projects/${encodeURIComponent(req.context_id || '')}/iamproviders`;
    const url = path + api.queryString(req, [`context_id`]);
    return api.get(url, undefined);
  }
  
  // Fetch an IAM provider by its id.
  // Required permissions:
  // - security.iamprovider.get on the IAM provider
  async GetIAMProvider(req: arangodb_cloud_common_v1_IDOptions): Promise<IAMProvider> {
    const path = `/api/security/v1/iamproviders/${encodeURIComponent(req.id || '')}`;
    const url = path + api.queryString(req, [`id`]);
    return api.get(url, undefined);
  }
  
  // Create a new IAM provider
  // Required permissions:
  // - security.iamprovider.create on the project that owns the IAM provider.
  async CreateIAMProvider(req: IAMProvider): Promise<IAMProvider> {
    const url = `/api/security/v1/project/${encodeURIComponent(req.project_id || '')}/iamproviders`;
    return api.post(url, req);
  }
  
  // Update an IAM provider
  // Required permissions:
  // - security.iamprovider.update on the IAM provider
  async UpdateIAMProvider(req: IAMProvider): Promise<IAMProvider> {
    const url = `/api/security/v1/iamproviders/${encodeURIComponent(req.id || '')}`;
    return api.patch(url, req);
  }
  
  // Delete an IAM provider.
  // Note that IAM providers are initially only marked for deletion.
  // Once all their dependent deployments are removed, the provider is removed.
  // Required permissions:
  // - security.iamprovider.delete on the IP whitelist
  async DeleteIAMProvider(req: arangodb_cloud_common_v1_IDOptions): Promise<void> {
    const path = `/api/security/v1/iamproviders/${encodeURIComponent(req.id || '')}`;
    const url = path + api.queryString(req, [`id`]);
    return api.delete(url, undefined);
  }
  
  // Mark the given IAM provider as default for its containing project.
  // Required permissions:
  // - security.iamprovider.set-default on the project that owns the provider.
  async SetDefaultIAMProvider(req: IAMProvider): Promise<void> {
    const url = `/api/security/v1/projects/${encodeURIComponent(req.project_id || '')}/iamproviders/default`;
    return api.post(url, req);
  }
}

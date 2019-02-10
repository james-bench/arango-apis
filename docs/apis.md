# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [resourcemanager/v1/resourcemanager.proto](#resourcemanager/v1/resourcemanager.proto)
    - [Event](#arangodb.cloud.resourcemanager.v1.Event)
    - [Event.PayloadEntry](#arangodb.cloud.resourcemanager.v1.Event.PayloadEntry)
    - [EventList](#arangodb.cloud.resourcemanager.v1.EventList)
    - [IsMemberOfOrganizationRequest](#arangodb.cloud.resourcemanager.v1.IsMemberOfOrganizationRequest)
    - [IsMemberOfOrganizationResponse](#arangodb.cloud.resourcemanager.v1.IsMemberOfOrganizationResponse)
    - [ListEventOptions](#arangodb.cloud.resourcemanager.v1.ListEventOptions)
    - [Member](#arangodb.cloud.resourcemanager.v1.Member)
    - [MemberList](#arangodb.cloud.resourcemanager.v1.MemberList)
    - [Organization](#arangodb.cloud.resourcemanager.v1.Organization)
    - [OrganizationList](#arangodb.cloud.resourcemanager.v1.OrganizationList)
    - [OrganizationMembersRequest](#arangodb.cloud.resourcemanager.v1.OrganizationMembersRequest)
    - [Project](#arangodb.cloud.resourcemanager.v1.Project)
    - [ProjectList](#arangodb.cloud.resourcemanager.v1.ProjectList)
  
  
  
    - [ResourceManagerService](#arangodb.cloud.resourcemanager.v1.ResourceManagerService)
  

- [crypto/v1/crypto.proto](#crypto/v1/crypto.proto)
    - [CACertificate](#arangodb.cloud.crypto.v1.CACertificate)
    - [CACertificateList](#arangodb.cloud.crypto.v1.CACertificateList)
  
  
  
    - [CryptoService](#arangodb.cloud.crypto.v1.CryptoService)
  

- [platform/v1/platform.proto](#platform/v1/platform.proto)
    - [Provider](#arangodb.cloud.platform.v1.Provider)
    - [ProviderList](#arangodb.cloud.platform.v1.ProviderList)
    - [Region](#arangodb.cloud.platform.v1.Region)
    - [RegionList](#arangodb.cloud.platform.v1.RegionList)
  
  
  
    - [PlatformService](#arangodb.cloud.platform.v1.PlatformService)
  

- [common/v1/common.proto](#common/v1/common.proto)
    - [Empty](#arangodb.cloud.common.v1.Empty)
    - [IDOptions](#arangodb.cloud.common.v1.IDOptions)
    - [ListOptions](#arangodb.cloud.common.v1.ListOptions)
    - [URLOptions](#arangodb.cloud.common.v1.URLOptions)
    - [YesOrNo](#arangodb.cloud.common.v1.YesOrNo)
  
  
  
  

- [iam/v1/iam.proto](#iam/v1/iam.proto)
    - [Group](#arangodb.cloud.iam.v1.Group)
    - [GroupList](#arangodb.cloud.iam.v1.GroupList)
    - [GroupMemberList](#arangodb.cloud.iam.v1.GroupMemberList)
    - [GroupMembersRequest](#arangodb.cloud.iam.v1.GroupMembersRequest)
    - [HasPermissionsRequest](#arangodb.cloud.iam.v1.HasPermissionsRequest)
    - [IsMemberOfGroupRequest](#arangodb.cloud.iam.v1.IsMemberOfGroupRequest)
    - [PermissionList](#arangodb.cloud.iam.v1.PermissionList)
    - [Policy](#arangodb.cloud.iam.v1.Policy)
    - [Role](#arangodb.cloud.iam.v1.Role)
    - [RoleBinding](#arangodb.cloud.iam.v1.RoleBinding)
    - [RoleBindingsRequest](#arangodb.cloud.iam.v1.RoleBindingsRequest)
    - [RoleList](#arangodb.cloud.iam.v1.RoleList)
    - [User](#arangodb.cloud.iam.v1.User)
  
  
  
    - [IAMService](#arangodb.cloud.iam.v1.IAMService)
  

- [data/v1/data.proto](#data/v1/data.proto)
    - [Deployment](#arangodb.cloud.data.v1.Deployment)
    - [Deployment.CertificateSpec](#arangodb.cloud.data.v1.Deployment.CertificateSpec)
    - [Deployment.ServersSpec](#arangodb.cloud.data.v1.Deployment.ServersSpec)
    - [Deployment.Status](#arangodb.cloud.data.v1.Deployment.Status)
    - [DeploymentList](#arangodb.cloud.data.v1.DeploymentList)
    - [ServersSpecLimits](#arangodb.cloud.data.v1.ServersSpecLimits)
    - [ServersSpecLimits.Limits](#arangodb.cloud.data.v1.ServersSpecLimits.Limits)
    - [ServersSpecLimitsRequest](#arangodb.cloud.data.v1.ServersSpecLimitsRequest)
    - [Version](#arangodb.cloud.data.v1.Version)
    - [VersionList](#arangodb.cloud.data.v1.VersionList)
  
  
  
    - [DataService](#arangodb.cloud.data.v1.DataService)
  

- [Scalar Value Types](#scalar-value-types)



<a name="resourcemanager/v1/resourcemanager.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## resourcemanager/v1/resourcemanager.proto



<a name="arangodb.cloud.resourcemanager.v1.Event"></a>

### Event
An Event represents something that happened to an organization
in the ArangoDB Managed service.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | System identifier of the event. This is a read-only value. |
| url | [string](#string) |  | URL of this resource This is a read-only value and cannot be initialized. |
| organization_id | [string](#string) |  | Identifier of the organization that owns this event. This is a read-only value. |
| subject_id | [string](#string) |  | Identifier of the subject of this event. This is a read-only value. If the subject of this event is an organization, this value is a duplicate of organization_id. |
| type | [string](#string) |  | Type of the event. |
| payload | [Event.PayloadEntry](#arangodb.cloud.resourcemanager.v1.Event.PayloadEntry) | repeated | Payload of the event. The fields used in the payload are specific to the type of event. |
| created_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | The creation timestamp of the event |






<a name="arangodb.cloud.resourcemanager.v1.Event.PayloadEntry"></a>

### Event.PayloadEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="arangodb.cloud.resourcemanager.v1.EventList"></a>

### EventList
List of Events.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| items | [Event](#arangodb.cloud.resourcemanager.v1.Event) | repeated |  |






<a name="arangodb.cloud.resourcemanager.v1.IsMemberOfOrganizationRequest"></a>

### IsMemberOfOrganizationRequest
Request arguments for IsMemberOfOrganization.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [string](#string) |  | Identifier of the user |
| organization_id | [string](#string) |  | Identifier of the organization |






<a name="arangodb.cloud.resourcemanager.v1.IsMemberOfOrganizationResponse"></a>

### IsMemberOfOrganizationResponse
Response for IsMemberOfOrganization.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| member | [bool](#bool) |  | Set if the requested user is a member of the requested organization. |
| owner | [bool](#bool) |  | Set if the requested user is an owner of the requested organization. |






<a name="arangodb.cloud.resourcemanager.v1.ListEventOptions"></a>

### ListEventOptions
Options for ListEvents


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| options | [arangodb.cloud.common.v1.ListOptions](#arangodb.cloud.common.v1.ListOptions) |  | Standard list options |
| subject_ids | [string](#string) | repeated | If set, filter on the subject_id of event |
| types | [string](#string) | repeated | If set, filter on the type of event |
| created_after | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | If set, filter of events created after this timestamp |
| created_before | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | If set, filter of events created before this timestamp |






<a name="arangodb.cloud.resourcemanager.v1.Member"></a>

### Member
Member of an organization.
A member is always a user.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [string](#string) |  | Identifier of the user |
| owner | [bool](#bool) |  | Set if this user is owner of the organization |






<a name="arangodb.cloud.resourcemanager.v1.MemberList"></a>

### MemberList
List of Members.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| items | [Member](#arangodb.cloud.resourcemanager.v1.Member) | repeated |  |






<a name="arangodb.cloud.resourcemanager.v1.Organization"></a>

### Organization
An Organization is represents a real world organization such as a company.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | System identifier of the organization. This is a read-only value. |
| url | [string](#string) |  | URL of this resource This is a read-only value and cannot be initialized. |
| name | [string](#string) |  | Name of the organization |
| description | [string](#string) |  | Description of the organization |
| is_deleted | [bool](#bool) |  | Set when this organization is deleted. This is a read-only value. |
| created_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | The creation timestamp of the organization |
| deleted_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | The deletion timestamp of the organization |






<a name="arangodb.cloud.resourcemanager.v1.OrganizationList"></a>

### OrganizationList
List of organizations.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| items | [Organization](#arangodb.cloud.resourcemanager.v1.Organization) | repeated |  |






<a name="arangodb.cloud.resourcemanager.v1.OrganizationMembersRequest"></a>

### OrganizationMembersRequest
Request arguments for Add/DeleteOrganizationMembers.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| organization_id | [string](#string) |  | Identifier of the organization to add/remove a user from |
| members | [MemberList](#arangodb.cloud.resourcemanager.v1.MemberList) |  | Users to add/remove. For every user, an owner flag is provided as well. If you add an existing user, the owner flag or the add request will overwrite the value of the existing owner flag. |






<a name="arangodb.cloud.resourcemanager.v1.Project"></a>

### Project
A Project is represents a unit within an organization such as a department.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | System identifier of the project. This is a read-only value. It can be set when creating the project. |
| url | [string](#string) |  | URL of this resource This is a read-only value and cannot be initialized. |
| name | [string](#string) |  | Name of the project |
| description | [string](#string) |  | Description of the project |
| organization_id | [string](#string) |  | Identifier of the organization that owns this project. This is a read-only value. |
| is_deleted | [bool](#bool) |  | Set when this project is deleted |
| created_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | The creation timestamp of the project |
| deleted_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | The deletion timestamp of the project |






<a name="arangodb.cloud.resourcemanager.v1.ProjectList"></a>

### ProjectList
List of Projects.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| items | [Project](#arangodb.cloud.resourcemanager.v1.Project) | repeated |  |





 

 

 


<a name="arangodb.cloud.resourcemanager.v1.ResourceManagerService"></a>

### ResourceManagerService
ResourceManagerService is the API used to configure basic resource objects.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| ListOrganizations | [.arangodb.cloud.common.v1.ListOptions](#arangodb.cloud.common.v1.ListOptions) | [OrganizationList](#arangodb.cloud.resourcemanager.v1.OrganizationList) | Fetch all organizations that the authenticated user is a member of. Required permissions: - None |
| GetOrganization | [.arangodb.cloud.common.v1.IDOptions](#arangodb.cloud.common.v1.IDOptions) | [Organization](#arangodb.cloud.resourcemanager.v1.Organization) | Fetch an organization by its id. The authenticated user must be a member of the organization. Required permissions: - None |
| CreateOrganization | [Organization](#arangodb.cloud.resourcemanager.v1.Organization) | [Organization](#arangodb.cloud.resourcemanager.v1.Organization) | Create a new organization Required permissions: - None |
| UpdateOrganization | [Organization](#arangodb.cloud.resourcemanager.v1.Organization) | [Organization](#arangodb.cloud.resourcemanager.v1.Organization) | Update an organization Required permissions: - resourcemanager.organization.update on the organization |
| DeleteOrganization | [Organization](#arangodb.cloud.resourcemanager.v1.Organization) | [.arangodb.cloud.common.v1.Empty](#arangodb.cloud.common.v1.Empty) | Delete an organization Note that organization are never really removed. Instead their is_deleted field is set to true. Required permissions: - resourcemanager.organization.delete on the organization |
| ListOrganizationMembers | [.arangodb.cloud.common.v1.ListOptions](#arangodb.cloud.common.v1.ListOptions) | [MemberList](#arangodb.cloud.resourcemanager.v1.MemberList) | Get a list of members of the organization identified by the given context ID. Required permissions: - resourcemanager.organization.get on the organization |
| AddOrganizationMembers | [OrganizationMembersRequest](#arangodb.cloud.resourcemanager.v1.OrganizationMembersRequest) | [.arangodb.cloud.common.v1.Empty](#arangodb.cloud.common.v1.Empty) | Add one or more members to an organization. Required permissions: - resourcemanager.organization.update on the organization |
| DeleteOrganizationMembers | [OrganizationMembersRequest](#arangodb.cloud.resourcemanager.v1.OrganizationMembersRequest) | [.arangodb.cloud.common.v1.Empty](#arangodb.cloud.common.v1.Empty) | Remove one or more members from an organization. Required permissions: - resourcemanager.organization.update on the organization |
| IsMemberOfOrganization | [IsMemberOfOrganizationRequest](#arangodb.cloud.resourcemanager.v1.IsMemberOfOrganizationRequest) | [IsMemberOfOrganizationResponse](#arangodb.cloud.resourcemanager.v1.IsMemberOfOrganizationResponse) | Is the user identified by the given user ID a member of the organization identified by the given organization ID. Required permissions: - resourcemanager.organization.get on the organization, unless the requested user is identical to the authenticated user. Note that if the identified user or organization does not exist, no is returned. |
| ListProjects | [.arangodb.cloud.common.v1.ListOptions](#arangodb.cloud.common.v1.ListOptions) | [ProjectList](#arangodb.cloud.resourcemanager.v1.ProjectList) | Fetch all projects in the organization identified by the given context ID. The authenticated user must be a member of the organization identifier by the given context ID. Required permissions: - resourcemanager.project.list on the organization identified by the given context ID |
| GetProject | [.arangodb.cloud.common.v1.IDOptions](#arangodb.cloud.common.v1.IDOptions) | [Project](#arangodb.cloud.resourcemanager.v1.Project) | Fetch a project by its id. The authenticated user must be a member of the organization that owns the project. Required permissions: - resourcemanager.project.get on the project identified by the given ID |
| CreateProject | [Project](#arangodb.cloud.resourcemanager.v1.Project) | [Project](#arangodb.cloud.resourcemanager.v1.Project) | Create a new project The authenticated user must be a member of the organization that owns the project. Required permissions: - resourcemanager.project.create on the organization that owns the project |
| UpdateProject | [Project](#arangodb.cloud.resourcemanager.v1.Project) | [Project](#arangodb.cloud.resourcemanager.v1.Project) | Update a project The authenticated user must be a member of the organization that owns the project. Required permissions: - resourcemanager.project.update on the project |
| DeleteProject | [Project](#arangodb.cloud.resourcemanager.v1.Project) | [.arangodb.cloud.common.v1.Empty](#arangodb.cloud.common.v1.Empty) | Delete a project Note that project are initially only marked for deleted. Once all their resources are removed the project itself is deleted and cannot be restored. The authenticated user must be a member of the organization that owns the project. Required permissions: - resourcemanager.project.delete on the project |
| ListEvents | [ListEventOptions](#arangodb.cloud.resourcemanager.v1.ListEventOptions) | [EventList](#arangodb.cloud.resourcemanager.v1.EventList) | Fetch all events in the organization identified by the given context ID. The authenticated user must be a member of the organization identifier by the given context ID. Required permissions: - resourcemanager.event.list on the organization identified by the given context ID |

 



<a name="crypto/v1/crypto.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## crypto/v1/crypto.proto



<a name="arangodb.cloud.crypto.v1.CACertificate"></a>

### CACertificate
A CACertificate is represents a self-signed certificate authority used to sign
TLS certificates for deployments &amp; client authentication.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | System identifier of the CA certificate. This is a read-only value. |
| url | [string](#string) |  | URL of this resource This is a read-only value. |
| name | [string](#string) |  | Name of the CA certificate |
| description | [string](#string) |  | Description of the CA certificate |
| project_id | [string](#string) |  | Identifier of the project that owns this CA certificate. This value cannot be changed after creation. |
| lifetime | [google.protobuf.Duration](#google.protobuf.Duration) |  | Time from creation of the CA certificate to expiration. This value cannot be changed after creation. |
| created_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | The creation timestamp of the CA certificate This is a read-only value. |
| deleted_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | The deletion timestamp of the CA certificate This is a read-only value. |
| expires_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | The expiration timestamp of the CA certificate This is a read-only value. |
| certificate_pem | [string](#string) |  | A PEM encoded representation of the public key of the CA certificate. This is a read-only value. |
| is_deleted | [bool](#bool) |  | Set when this CA certificate is deleted. This is a read-only value. |






<a name="arangodb.cloud.crypto.v1.CACertificateList"></a>

### CACertificateList
List of CACertificates.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| items | [CACertificate](#arangodb.cloud.crypto.v1.CACertificate) | repeated |  |





 

 

 


<a name="arangodb.cloud.crypto.v1.CryptoService"></a>

### CryptoService
CryptoService is the API used to configure various crypto objects.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| ListCACertificates | [.arangodb.cloud.common.v1.ListOptions](#arangodb.cloud.common.v1.ListOptions) | [CACertificateList](#arangodb.cloud.crypto.v1.CACertificateList) | Fetch all CA certificates in the project identified by the given context ID. Required permissions: - crypto.cacertificate.list on the project identified by the given context ID |
| GetCACertificate | [.arangodb.cloud.common.v1.IDOptions](#arangodb.cloud.common.v1.IDOptions) | [CACertificate](#arangodb.cloud.crypto.v1.CACertificate) | Fetch a CA certificate by its id. Required permissions: - crypto.cacertificate.get on the CA certificate identified by the given ID |
| CreateCACertificate | [CACertificate](#arangodb.cloud.crypto.v1.CACertificate) | [CACertificate](#arangodb.cloud.crypto.v1.CACertificate) | Create a new CA certificate Required permissions: - crypto.cacertificate.create on the project that owns the CA certificate |
| UpdateCACertificate | [CACertificate](#arangodb.cloud.crypto.v1.CACertificate) | [CACertificate](#arangodb.cloud.crypto.v1.CACertificate) | Update a CA certificate Required permissions: - crypto.cacertificate.update on the CA certificate |
| DeleteCACertificate | [CACertificate](#arangodb.cloud.crypto.v1.CACertificate) | [.arangodb.cloud.common.v1.Empty](#arangodb.cloud.common.v1.Empty) | Delete a CA certificate Note that CA certificate are initially only marked for deleted. Once all the resources that depend on it are removed the CA certificate itself is deleted and cannot be restored. Required permissions: - crypto.cacertificate.delete on the CA certificate |

 



<a name="platform/v1/platform.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## platform/v1/platform.proto



<a name="arangodb.cloud.platform.v1.Provider"></a>

### Provider
Provider represents a specific cloud provider such as AWS or GCP.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | System identifier of the provider. |
| name | [string](#string) |  | Name of the provider |






<a name="arangodb.cloud.platform.v1.ProviderList"></a>

### ProviderList
List of providers.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| items | [Provider](#arangodb.cloud.platform.v1.Provider) | repeated |  |






<a name="arangodb.cloud.platform.v1.Region"></a>

### Region
Region represents a geographical region in which deployments are run.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | System identifier of the region. |
| provider_id | [string](#string) |  | Identifier of the provider that hosts this region. |
| location | [string](#string) |  | Location of the region |
| available | [bool](#bool) |  | Is this region available for creating new deployments? |






<a name="arangodb.cloud.platform.v1.RegionList"></a>

### RegionList
List of regions.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| items | [Region](#arangodb.cloud.platform.v1.Region) | repeated |  |





 

 

 


<a name="arangodb.cloud.platform.v1.PlatformService"></a>

### PlatformService
PlatformService is the API used to query for cloud provider &amp; regional info.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| ListProviders | [.arangodb.cloud.common.v1.ListOptions](#arangodb.cloud.common.v1.ListOptions) | [ProviderList](#arangodb.cloud.platform.v1.ProviderList) | Fetch all providers that are supported by the ArangoDB cloud. Required permissions: - None |
| GetProvider | [.arangodb.cloud.common.v1.IDOptions](#arangodb.cloud.common.v1.IDOptions) | [Provider](#arangodb.cloud.platform.v1.Provider) | Fetch a provider by its id. Required permissions: - None |
| ListRegions | [.arangodb.cloud.common.v1.ListOptions](#arangodb.cloud.common.v1.ListOptions) | [RegionList](#arangodb.cloud.platform.v1.RegionList) | Fetch all regions provided by the provided identified by the given context ID. Required permissions: - None |
| GetRegion | [.arangodb.cloud.common.v1.IDOptions](#arangodb.cloud.common.v1.IDOptions) | [Region](#arangodb.cloud.platform.v1.Region) | Fetch a region by its id. Required permissions: - None |

 



<a name="common/v1/common.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## common/v1/common.proto



<a name="arangodb.cloud.common.v1.Empty"></a>

### Empty
Empty message






<a name="arangodb.cloud.common.v1.IDOptions"></a>

### IDOptions
Options for a get-by-id request


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | System identifier of the object to fetch. |






<a name="arangodb.cloud.common.v1.ListOptions"></a>

### ListOptions
Options for a list request.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| page_size | [int32](#int32) |  | Maximum number of items to return. If not specified, all remaining items are returned. |
| page | [int64](#int64) |  | Page to start with (defaults to 0). |
| context_id | [string](#string) |  | Identifier of the resource in which the list request is made. |






<a name="arangodb.cloud.common.v1.URLOptions"></a>

### URLOptions
Options for a get-by-url request


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| url | [string](#string) |  | URL of the resource to fetch. |






<a name="arangodb.cloud.common.v1.YesOrNo"></a>

### YesOrNo
Response for single boolean.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| result | [bool](#bool) |  |  |





 

 

 

 



<a name="iam/v1/iam.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## iam/v1/iam.proto



<a name="arangodb.cloud.iam.v1.Group"></a>

### Group
Group of user accounts.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | System identifier of the group. This is a read-only value. |
| organization_id | [string](#string) |  | Identifier of the organization that owns this group. |
| name | [string](#string) |  | Name of the group |
| description | [string](#string) |  | Description of the group |
| created_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | The creation timestamp of the group |
| deleted_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | The deletion timestamp of the group |
| is_deleted | [bool](#bool) |  | Set when this organization is deleted. This is a read-only value. |
| url | [string](#string) |  | URL of this resource This is a read-only value and cannot be initialized. |
| is_virtual | [bool](#bool) |  | Set if this group is virtual and managed by the system. This is a read-only value. |






<a name="arangodb.cloud.iam.v1.GroupList"></a>

### GroupList
List of groups.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| items | [Group](#arangodb.cloud.iam.v1.Group) | repeated |  |






<a name="arangodb.cloud.iam.v1.GroupMemberList"></a>

### GroupMemberList
List of group members (user ID&#39;s)


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| items | [string](#string) | repeated | List of ID&#39;s of users that are member of the group. |






<a name="arangodb.cloud.iam.v1.GroupMembersRequest"></a>

### GroupMembersRequest
Request arguments for Add/DeleteGroupMembers.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| group_id | [string](#string) |  | ID of the group to add/remove members to/from. |
| user_ids | [string](#string) | repeated | ID&#39;s of users to add/remove to/from the group. |






<a name="arangodb.cloud.iam.v1.HasPermissionsRequest"></a>

### HasPermissionsRequest
Request arguments for HasPermissionsRequest.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| url | [string](#string) |  | URL of the resource to query permissions for. |
| permissions | [string](#string) | repeated | The list of permissions that are required. |






<a name="arangodb.cloud.iam.v1.IsMemberOfGroupRequest"></a>

### IsMemberOfGroupRequest
Request arguments for IsMemberOfGroup.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [string](#string) |  | Identifier of the user |
| group_id | [string](#string) |  | Identifier of the group |






<a name="arangodb.cloud.iam.v1.PermissionList"></a>

### PermissionList
List of permissions.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| items | [string](#string) | repeated |  |






<a name="arangodb.cloud.iam.v1.Policy"></a>

### Policy
Policy bindings members to roles for access to a resource.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| resource_url | [string](#string) |  | URL of the resource to which this policy applies. |
| bindings | [RoleBinding](#arangodb.cloud.iam.v1.RoleBinding) | repeated | Role bindings to apply to the resource. |






<a name="arangodb.cloud.iam.v1.Role"></a>

### Role
A role is a list of permissions.
Roles can be bound to resources for members.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | System identifier of the role. This is a read-only value. |
| organization_id | [string](#string) |  | Identifier of the organization that owns this role. This value is undefined for predefined roles. |
| name | [string](#string) |  | Name of the role |
| description | [string](#string) |  | Description of the role |
| permissions | [string](#string) | repeated | Permissions to grant when this role is bound. |
| is_predefined | [bool](#bool) |  | Set if this role is predefined. This is a read-only value. |
| created_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | The creation timestamp of the role |
| deleted_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | The deletion timestamp of the role |
| is_deleted | [bool](#bool) |  | Set when this organization is deleted. This is a read-only value. |
| url | [string](#string) |  | URL of this resource This is a read-only value and cannot be initialized. |






<a name="arangodb.cloud.iam.v1.RoleBinding"></a>

### RoleBinding
RoleBinding binds a Role to a member.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | System identifier of the role-binding. This is a read-only value. |
| member_id | [string](#string) |  | Identifier of the member to bind a role to. Member ID is formatted as: - user:&lt;user_id&gt; - group:&lt;group_id&gt; |
| role_id | [string](#string) |  | Identifier of the Role to grant to member |






<a name="arangodb.cloud.iam.v1.RoleBindingsRequest"></a>

### RoleBindingsRequest
Request arguments for Add/DeleteRoleBindings.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| resource_url | [string](#string) |  | URL of the resource to add/remove policy binding to/from. |
| bindings | [RoleBinding](#arangodb.cloud.iam.v1.RoleBinding) | repeated | Role bindings to add/remove to the policy. |






<a name="arangodb.cloud.iam.v1.RoleList"></a>

### RoleList
List of roles.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| items | [Role](#arangodb.cloud.iam.v1.Role) | repeated |  |






<a name="arangodb.cloud.iam.v1.User"></a>

### User
User represents an actual person.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | Identifier of the user. This is a read-only value. |
| email | [string](#string) |  | Email address of the user. |
| name | [string](#string) |  | Name of the user. This may be empty if not filled out by the user. |
| given_name | [string](#string) |  | Given name of the user. This may be empty if not filled out by the user. |
| family_name | [string](#string) |  | Family name of the user. This may be empty if not filled out by the user. |
| created_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | The creation timestamp of the user. |





 

 

 


<a name="arangodb.cloud.iam.v1.IAMService"></a>

### IAMService
IAMService is the API used to configure IAM objects.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetThisUser | [.arangodb.cloud.common.v1.Empty](#arangodb.cloud.common.v1.Empty) | [User](#arangodb.cloud.iam.v1.User) | Fetch all available information of the currently authenticated user. Required permissions: - None |
| GetUser | [.arangodb.cloud.common.v1.IDOptions](#arangodb.cloud.common.v1.IDOptions) | [User](#arangodb.cloud.iam.v1.User) | Fetch all available information of the user identified by the given ID. Required permissions: - resourcemanager.organization.get on one of the organizations that the request user and authenticated user are both a member of |
| ListGroups | [.arangodb.cloud.common.v1.ListOptions](#arangodb.cloud.common.v1.ListOptions) | [GroupList](#arangodb.cloud.iam.v1.GroupList) | Fetch all groups of the organization identified by the given context ID. Required permissions: - iam.group.list on organization identified by given context ID. |
| GetGroup | [.arangodb.cloud.common.v1.IDOptions](#arangodb.cloud.common.v1.IDOptions) | [Group](#arangodb.cloud.iam.v1.Group) | Fetch a group by its id. Required permissions: - iam.group.get on organization that owns the group |
| CreateGroup | [Group](#arangodb.cloud.iam.v1.Group) | [Group](#arangodb.cloud.iam.v1.Group) | Create a group Required permissions: - iam.group.create on organization that owns the group |
| UpdateGroup | [Group](#arangodb.cloud.iam.v1.Group) | [Group](#arangodb.cloud.iam.v1.Group) | Update a group Required permissions: - iam.group.update on organization that owns the group |
| DeleteGroup | [Group](#arangodb.cloud.iam.v1.Group) | [.arangodb.cloud.common.v1.Empty](#arangodb.cloud.common.v1.Empty) | Delete a group Required permissions: - iam.group.delete on organization that owns the group |
| ListGroupMembers | [.arangodb.cloud.common.v1.ListOptions](#arangodb.cloud.common.v1.ListOptions) | [GroupMemberList](#arangodb.cloud.iam.v1.GroupMemberList) | List of members of the group identified by the given context ID. Required permissions: - iam.group.get on organization that owns the group |
| AddGroupMembers | [GroupMembersRequest](#arangodb.cloud.iam.v1.GroupMembersRequest) | [Group](#arangodb.cloud.iam.v1.Group) | Add one or more members to the group identified by given ID. Required permissions: - iam.group.update on organization that owns the group |
| DeleteGroupMembers | [GroupMembersRequest](#arangodb.cloud.iam.v1.GroupMembersRequest) | [Group](#arangodb.cloud.iam.v1.Group) | Remove one or more members from the group identified by given ID. Required permissions: - iam.group.update on organization that owns the group |
| IsMemberOfGroup | [IsMemberOfGroupRequest](#arangodb.cloud.iam.v1.IsMemberOfGroupRequest) | [.arangodb.cloud.common.v1.YesOrNo](#arangodb.cloud.common.v1.YesOrNo) | Is the user identified by the given user ID a member of the group identified by the given group ID. Required permissions: - iam.group.get on organization that owns the group, unless the requested user is identical to the authenticated user. Note that if the identified group does not exist, no is returned. |
| ListRoles | [.arangodb.cloud.common.v1.ListOptions](#arangodb.cloud.common.v1.ListOptions) | [RoleList](#arangodb.cloud.iam.v1.RoleList) | Fetch all roles in the organization identified by the given context ID. Required permissions: - iam.role.list on organization identified by given context ID. |
| GetRole | [.arangodb.cloud.common.v1.IDOptions](#arangodb.cloud.common.v1.IDOptions) | [Role](#arangodb.cloud.iam.v1.Role) | Fetch a role by its id. Required permissions: - iam.role.get on organization that owns the role |
| CreateRole | [Role](#arangodb.cloud.iam.v1.Role) | [Role](#arangodb.cloud.iam.v1.Role) | Create a custom role Required permissions: - iam.role.create on organization that owns the role |
| UpdateRole | [Role](#arangodb.cloud.iam.v1.Role) | [Role](#arangodb.cloud.iam.v1.Role) | Update a custom role Required permissions: - iam.role.update on organization that owns the role |
| DeleteRole | [Role](#arangodb.cloud.iam.v1.Role) | [.arangodb.cloud.common.v1.Empty](#arangodb.cloud.common.v1.Empty) | Delete a custom role Required permissions: - iam.role.delete on organization that owns the role |
| GetPolicy | [.arangodb.cloud.common.v1.URLOptions](#arangodb.cloud.common.v1.URLOptions) | [Policy](#arangodb.cloud.iam.v1.Policy) | Get the policy for a resource identified by given URL. Required permissions: - iam.policy.get on resource identified by the url |
| AddRoleBindings | [RoleBindingsRequest](#arangodb.cloud.iam.v1.RoleBindingsRequest) | [Policy](#arangodb.cloud.iam.v1.Policy) | Add one or more RoleBindings to the policy of a resource identified by given URL. Required permissions: - iam.policy.update on resource identified by the url |
| DeleteRoleBindings | [RoleBindingsRequest](#arangodb.cloud.iam.v1.RoleBindingsRequest) | [Policy](#arangodb.cloud.iam.v1.Policy) | Remove one or more RoleBindings from the policy of a resource identified by given URL. Required permissions: - iam.policy.update on resource identified by the url |
| GetEffectivePermissions | [.arangodb.cloud.common.v1.URLOptions](#arangodb.cloud.common.v1.URLOptions) | [PermissionList](#arangodb.cloud.iam.v1.PermissionList) | Return the list of permissions that are available to the currently authenticated used for actions on the resource identified by the given URL. Required permissions: - None |
| HasPermissions | [HasPermissionsRequest](#arangodb.cloud.iam.v1.HasPermissionsRequest) | [.arangodb.cloud.common.v1.YesOrNo](#arangodb.cloud.common.v1.YesOrNo) | Does the authenticated user have all of the requested permissions for the resource identified by the given URL? Required permissions: - None |
| ListPermissions | [.arangodb.cloud.common.v1.Empty](#arangodb.cloud.common.v1.Empty) | [PermissionList](#arangodb.cloud.iam.v1.PermissionList) | List all known permissions. Required permissions: - None |

 



<a name="data/v1/data.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## data/v1/data.proto



<a name="arangodb.cloud.data.v1.Deployment"></a>

### Deployment
A Deployment is represents one deployment of an ArangoDB cluster.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | System identifier of the deployment. This is a read-only value. |
| url | [string](#string) |  | URL of this resource This is a read-only value. |
| name | [string](#string) |  | Name of the deployment |
| description | [string](#string) |  | Description of the deployment |
| project_id | [string](#string) |  | Identifier of the project that owns this deployment. After creation, this value cannot be changed. |
| region_id | [string](#string) |  | Identifier of the region in which the deployment is created. After creation, this value cannot be changed. |
| created_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | The creation timestamp of the deployment This is a read-only value. |
| deleted_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | The deletion timestamp of the deployment This is a read-only value. |
| is_deleted | [bool](#bool) |  | Set when this deployment is deleted. This is a read-only value. |
| version | [string](#string) |  | ArangoDB version to use for this deployment. See Version.version. If you change this value to a higher version, the deployment will be upgraded. If you change this value to a lower patch value, the deployment will be downgraded. Any attempt to change to a lower minor or major version is considered an invalid request. Any attempt to change to a version that is not in the list of available versions is considered an invalid request. |
| certificates | [Deployment.CertificateSpec](#arangodb.cloud.data.v1.Deployment.CertificateSpec) |  |  |
| servers | [Deployment.ServersSpec](#arangodb.cloud.data.v1.Deployment.ServersSpec) |  |  |
| status | [Deployment.Status](#arangodb.cloud.data.v1.Deployment.Status) |  |  |






<a name="arangodb.cloud.data.v1.Deployment.CertificateSpec"></a>

### Deployment.CertificateSpec



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ca_certificate_id | [string](#string) |  | Identifier of the CACertificate used to sign TLS certificates for the deployment. If this value is empty during creation of the deployment, a new CA certificate will be created for this deployment. If you change this value after the creation of the deployment a complete rotation of the deployment is required, which will result in some downtime. |
| alternate_dns_names | [string](#string) | repeated | Zero or more DNS names to include in the TLS certificate of the deployment. |






<a name="arangodb.cloud.data.v1.Deployment.ServersSpec"></a>

### Deployment.ServersSpec



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| coordinators | [int32](#int32) |  | Number of coordinators of the deployment |
| coordinator_memory_size | [int32](#int32) |  | Amount of memory (in GB) to allocate for coordinators. |
| dbservers | [int32](#int32) |  | Number of dbservers of the deployment |
| dbserver_memory_size | [int32](#int32) |  | Amount of memory (in GB) to allocate for dbservers. |
| dbserver_disk_size | [int32](#int32) |  | Amount of disk space (in GB) to allocate for dbservers. |






<a name="arangodb.cloud.data.v1.Deployment.Status"></a>

### Deployment.Status
Status of the deployment
All members of this field are read-only.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| endpoint | [string](#string) |  | Endpoint URL used to reach the deployment This value will be empty during the creation of the deployment. |
| description | [string](#string) |  | Human readable description of the status of the deployment. |
| created | [bool](#bool) |  | Set once the deployment has been created. |
| ready | [bool](#bool) |  | Set if the deployment is ready to be used. If the deployment has downtime (e.g. because of changing a CA certificate) this will go to false until the downtime is over. |
| upgrading | [bool](#bool) |  | Set if the deployment is being upgraded. |
| server_versions | [string](#string) | repeated | Versions of running servers |






<a name="arangodb.cloud.data.v1.DeploymentList"></a>

### DeploymentList
List of Deployments.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| items | [Deployment](#arangodb.cloud.data.v1.Deployment) | repeated |  |






<a name="arangodb.cloud.data.v1.ServersSpecLimits"></a>

### ServersSpecLimits
Limits of allowed values for fields of Deployment.ServersSpec.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| coordinators | [ServersSpecLimits.Limits](#arangodb.cloud.data.v1.ServersSpecLimits.Limits) |  | Limits for the number of coordinators of the deployment |
| coordinator_memory_size | [ServersSpecLimits.Limits](#arangodb.cloud.data.v1.ServersSpecLimits.Limits) |  | Possible values for the amount of memory (in GB) to allocate for coordinators. |
| dbservers | [ServersSpecLimits.Limits](#arangodb.cloud.data.v1.ServersSpecLimits.Limits) |  | Limits for the number of dbservers of the deployment |
| dbserver_memory_size | [ServersSpecLimits.Limits](#arangodb.cloud.data.v1.ServersSpecLimits.Limits) |  | Possible values for the amount of memory (in GB) to allocate for dbservers. |
| dbserver_disk_size | [ServersSpecLimits.Limits](#arangodb.cloud.data.v1.ServersSpecLimits.Limits) |  | Amount of disk space (in GB) to allocate for dbservers. |






<a name="arangodb.cloud.data.v1.ServersSpecLimits.Limits"></a>

### ServersSpecLimits.Limits



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| min | [int32](#int32) |  | Minimum value |
| max | [int32](#int32) |  | Maximum value |
| allowed_values | [int32](#int32) | repeated | Set of allowed values. If this field is non-empty, only one of these values is allowed. |






<a name="arangodb.cloud.data.v1.ServersSpecLimitsRequest"></a>

### ServersSpecLimitsRequest
Request arguments for GetServersSpecLimits


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| project_id | [string](#string) |  | Identifier of project that will own a deployment. |
| region_id | [string](#string) |  | Identifier of a region in which a deployment will be created. |






<a name="arangodb.cloud.data.v1.Version"></a>

### Version
Version of an ArangoDB release


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| version | [string](#string) |  | Version in the format of major.minor.patch |






<a name="arangodb.cloud.data.v1.VersionList"></a>

### VersionList
List of Versions.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| items | [Version](#arangodb.cloud.data.v1.Version) | repeated |  |





 

 

 


<a name="arangodb.cloud.data.v1.DataService"></a>

### DataService
DataService is the API used to configure data objects.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| ListDeployments | [.arangodb.cloud.common.v1.ListOptions](#arangodb.cloud.common.v1.ListOptions) | [DeploymentList](#arangodb.cloud.data.v1.DeploymentList) | Fetch all deployments in the project identified by the given context ID. Required permissions: - data.deployment.list on the project identified by the given context ID |
| GetDeployment | [.arangodb.cloud.common.v1.IDOptions](#arangodb.cloud.common.v1.IDOptions) | [Deployment](#arangodb.cloud.data.v1.Deployment) | Fetch a deployment by its id. Required permissions: - data.deployment.get on the deployment identified by the given ID |
| CreateDeployment | [Deployment](#arangodb.cloud.data.v1.Deployment) | [Deployment](#arangodb.cloud.data.v1.Deployment) | Create a new deployment Required permissions: - data.deployment.create on the project that owns the deployment |
| UpdateDeployment | [Deployment](#arangodb.cloud.data.v1.Deployment) | [Deployment](#arangodb.cloud.data.v1.Deployment) | Update a deployment Required permissions: - data.deployment.update on the deployment |
| DeleteDeployment | [Deployment](#arangodb.cloud.data.v1.Deployment) | [.arangodb.cloud.common.v1.Empty](#arangodb.cloud.common.v1.Empty) | Delete a deployment Note that deployments are initially only marked for deletion. Once all their resources are removed the deployment itself is removed. Required permissions: - data.deployment.delete on the deployment |
| ListVersions | [.arangodb.cloud.common.v1.ListOptions](#arangodb.cloud.common.v1.ListOptions) | [VersionList](#arangodb.cloud.data.v1.VersionList) | Fetch all ArangoDB versions that are available for deployments. Required permissions: - None |
| GetDefaultVersion | [.arangodb.cloud.common.v1.Empty](#arangodb.cloud.common.v1.Empty) | [Version](#arangodb.cloud.data.v1.Version) | Fetch the default ArangoDB version for new deployment. Required permissions: - None |
| GetServersSpecLimits | [ServersSpecLimitsRequest](#arangodb.cloud.data.v1.ServersSpecLimitsRequest) | [ServersSpecLimits](#arangodb.cloud.data.v1.ServersSpecLimits) | Fetch the limits for server specifications for deployments owned by the given projected, created in the given region. Required permissions: - data.limits.get on the requested project |

 



## Scalar Value Types

| .proto Type | Notes | C++ Type | Java Type | Python Type |
| ----------- | ----- | -------- | --------- | ----------- |
| <a name="double" /> double |  | double | double | float |
| <a name="float" /> float |  | float | float | float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long |
| <a name="bool" /> bool |  | bool | boolean | boolean |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str |


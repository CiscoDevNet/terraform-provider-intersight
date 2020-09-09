# HyperflexNodeRelationship

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountMoid** | Pointer to **string** | The Account ID for this managed object. | [optional] [readonly] 
**ClassId** | **string** | The concrete type of this complex type. Its value must be the same as the &#39;objectType&#39; property. The OpenAPI document references this property as a discriminator value. | [readonly] 
**CreateTime** | Pointer to [**time.Time**](time.Time.md) | The time when this managed object was created. | [optional] [readonly] 
**DomainGroupMoid** | Pointer to **string** | The DomainGroup ID for this managed object. | [optional] [readonly] 
**ModTime** | Pointer to [**time.Time**](time.Time.md) | The time when this managed object was last modified. | [optional] [readonly] 
**Moid** | Pointer to **string** | The unique identifier of this Managed Object instance. | [optional] 
**ObjectType** | **string** | The fully-qualified type of this managed object, i.e. the class name. This property is optional. The ObjectType is implied from the URL path. If specified, the value of objectType must match the class name specified in the URL path. | [readonly] 
**Owners** | Pointer to **[]string** |  | [optional] 
**SharedScope** | Pointer to **string** | Intersight provides pre-built workflows, tasks and policies to end users through global catalogs. Objects that are made available through global catalogs are said to have a &#39;shared&#39; ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. | [optional] [readonly] 
**Tags** | Pointer to [**[]MoTag**](mo.Tag.md) |  | [optional] 
**VersionContext** | Pointer to [**MoVersionContext**](mo.VersionContext.md) |  | [optional] 
**Ancestors** | Pointer to [**[]MoBaseMoRelationship**](mo.BaseMo.Relationship.md) | An array of relationships to moBaseMo resources. | [optional] [readonly] 
**Parent** | Pointer to [**MoBaseMoRelationship**](mo.BaseMo.Relationship.md) |  | [optional] 
**PermissionResources** | Pointer to [**[]MoBaseMoRelationship**](mo.BaseMo.Relationship.md) | An array of relationships to moBaseMo resources. | [optional] [readonly] 
**DisplayNames** | Pointer to [**map[string][]string**](array.md) | A set of display names for the MO resource. These names are calculated based on other properties of the MO and potentially properties of Ancestor MOs. Displaynames are intended as a way to provide a normalized user appropriate name for an MO, especially for MOs which do not have a &#39;Name&#39; property, which is the case for much of the inventory discovered from managed targets. There are a limited number of keys, currently &#39;short&#39; and &#39;hierarchical&#39;. The value is an array and clients should use the first element of the array. | [optional] [readonly] 
**BuildNumber** | Pointer to **string** |  | [optional] [readonly] 
**DisplayVersion** | Pointer to **string** |  | [optional] [readonly] 
**HostName** | Pointer to **string** |  | [optional] [readonly] 
**Hypervisor** | Pointer to **string** |  | [optional] [readonly] 
**Identity** | Pointer to [**HyperflexHxUuIdDt**](hyperflex.HxUuIdDt.md) |  | [optional] 
**Ip** | Pointer to [**HyperflexHxNetworkAddressDt**](hyperflex.HxNetworkAddressDt.md) |  | [optional] 
**Lockdown** | Pointer to **bool** |  | [optional] [readonly] 
**ModelNumber** | Pointer to **string** |  | [optional] [readonly] 
**Role** | Pointer to **string** |  | [optional] [readonly] [default to "UNKNOWN"]
**SerialNumber** | Pointer to **string** |  | [optional] [readonly] 
**Status** | Pointer to **string** |  | [optional] [readonly] [default to "UNKNOWN"]
**Version** | Pointer to **string** |  | [optional] [readonly] 
**Cluster** | Pointer to [**HyperflexClusterRelationship**](hyperflex.Cluster.Relationship.md) |  | [optional] 
**ClusterMember** | Pointer to [**AssetClusterMemberRelationship**](asset.ClusterMember.Relationship.md) |  | [optional] 
**PhysicalServer** | Pointer to [**ComputePhysicalRelationship**](compute.Physical.Relationship.md) |  | [optional] 

## Methods

### NewHyperflexNodeRelationship

`func NewHyperflexNodeRelationship(classId string, objectType string, ) *HyperflexNodeRelationship`

NewHyperflexNodeRelationship instantiates a new HyperflexNodeRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexNodeRelationshipWithDefaults

`func NewHyperflexNodeRelationshipWithDefaults() *HyperflexNodeRelationship`

NewHyperflexNodeRelationshipWithDefaults instantiates a new HyperflexNodeRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountMoid

`func (o *HyperflexNodeRelationship) GetAccountMoid() string`

GetAccountMoid returns the AccountMoid field if non-nil, zero value otherwise.

### GetAccountMoidOk

`func (o *HyperflexNodeRelationship) GetAccountMoidOk() (*string, bool)`

GetAccountMoidOk returns a tuple with the AccountMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMoid

`func (o *HyperflexNodeRelationship) SetAccountMoid(v string)`

SetAccountMoid sets AccountMoid field to given value.

### HasAccountMoid

`func (o *HyperflexNodeRelationship) HasAccountMoid() bool`

HasAccountMoid returns a boolean if a field has been set.

### GetClassId

`func (o *HyperflexNodeRelationship) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexNodeRelationship) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexNodeRelationship) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetCreateTime

`func (o *HyperflexNodeRelationship) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *HyperflexNodeRelationship) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *HyperflexNodeRelationship) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *HyperflexNodeRelationship) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDomainGroupMoid

`func (o *HyperflexNodeRelationship) GetDomainGroupMoid() string`

GetDomainGroupMoid returns the DomainGroupMoid field if non-nil, zero value otherwise.

### GetDomainGroupMoidOk

`func (o *HyperflexNodeRelationship) GetDomainGroupMoidOk() (*string, bool)`

GetDomainGroupMoidOk returns a tuple with the DomainGroupMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGroupMoid

`func (o *HyperflexNodeRelationship) SetDomainGroupMoid(v string)`

SetDomainGroupMoid sets DomainGroupMoid field to given value.

### HasDomainGroupMoid

`func (o *HyperflexNodeRelationship) HasDomainGroupMoid() bool`

HasDomainGroupMoid returns a boolean if a field has been set.

### GetModTime

`func (o *HyperflexNodeRelationship) GetModTime() time.Time`

GetModTime returns the ModTime field if non-nil, zero value otherwise.

### GetModTimeOk

`func (o *HyperflexNodeRelationship) GetModTimeOk() (*time.Time, bool)`

GetModTimeOk returns a tuple with the ModTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModTime

`func (o *HyperflexNodeRelationship) SetModTime(v time.Time)`

SetModTime sets ModTime field to given value.

### HasModTime

`func (o *HyperflexNodeRelationship) HasModTime() bool`

HasModTime returns a boolean if a field has been set.

### GetMoid

`func (o *HyperflexNodeRelationship) GetMoid() string`

GetMoid returns the Moid field if non-nil, zero value otherwise.

### GetMoidOk

`func (o *HyperflexNodeRelationship) GetMoidOk() (*string, bool)`

GetMoidOk returns a tuple with the Moid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoid

`func (o *HyperflexNodeRelationship) SetMoid(v string)`

SetMoid sets Moid field to given value.

### HasMoid

`func (o *HyperflexNodeRelationship) HasMoid() bool`

HasMoid returns a boolean if a field has been set.

### GetObjectType

`func (o *HyperflexNodeRelationship) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexNodeRelationship) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexNodeRelationship) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOwners

`func (o *HyperflexNodeRelationship) GetOwners() []string`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *HyperflexNodeRelationship) GetOwnersOk() (*[]string, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *HyperflexNodeRelationship) SetOwners(v []string)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *HyperflexNodeRelationship) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetSharedScope

`func (o *HyperflexNodeRelationship) GetSharedScope() string`

GetSharedScope returns the SharedScope field if non-nil, zero value otherwise.

### GetSharedScopeOk

`func (o *HyperflexNodeRelationship) GetSharedScopeOk() (*string, bool)`

GetSharedScopeOk returns a tuple with the SharedScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedScope

`func (o *HyperflexNodeRelationship) SetSharedScope(v string)`

SetSharedScope sets SharedScope field to given value.

### HasSharedScope

`func (o *HyperflexNodeRelationship) HasSharedScope() bool`

HasSharedScope returns a boolean if a field has been set.

### GetTags

`func (o *HyperflexNodeRelationship) GetTags() []MoTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *HyperflexNodeRelationship) GetTagsOk() (*[]MoTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *HyperflexNodeRelationship) SetTags(v []MoTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *HyperflexNodeRelationship) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersionContext

`func (o *HyperflexNodeRelationship) GetVersionContext() MoVersionContext`

GetVersionContext returns the VersionContext field if non-nil, zero value otherwise.

### GetVersionContextOk

`func (o *HyperflexNodeRelationship) GetVersionContextOk() (*MoVersionContext, bool)`

GetVersionContextOk returns a tuple with the VersionContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionContext

`func (o *HyperflexNodeRelationship) SetVersionContext(v MoVersionContext)`

SetVersionContext sets VersionContext field to given value.

### HasVersionContext

`func (o *HyperflexNodeRelationship) HasVersionContext() bool`

HasVersionContext returns a boolean if a field has been set.

### GetAncestors

`func (o *HyperflexNodeRelationship) GetAncestors() []MoBaseMoRelationship`

GetAncestors returns the Ancestors field if non-nil, zero value otherwise.

### GetAncestorsOk

`func (o *HyperflexNodeRelationship) GetAncestorsOk() (*[]MoBaseMoRelationship, bool)`

GetAncestorsOk returns a tuple with the Ancestors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestors

`func (o *HyperflexNodeRelationship) SetAncestors(v []MoBaseMoRelationship)`

SetAncestors sets Ancestors field to given value.

### HasAncestors

`func (o *HyperflexNodeRelationship) HasAncestors() bool`

HasAncestors returns a boolean if a field has been set.

### SetAncestorsNil

`func (o *HyperflexNodeRelationship) SetAncestorsNil(b bool)`

 SetAncestorsNil sets the value for Ancestors to be an explicit nil

### UnsetAncestors
`func (o *HyperflexNodeRelationship) UnsetAncestors()`

UnsetAncestors ensures that no value is present for Ancestors, not even an explicit nil
### GetParent

`func (o *HyperflexNodeRelationship) GetParent() MoBaseMoRelationship`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *HyperflexNodeRelationship) GetParentOk() (*MoBaseMoRelationship, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *HyperflexNodeRelationship) SetParent(v MoBaseMoRelationship)`

SetParent sets Parent field to given value.

### HasParent

`func (o *HyperflexNodeRelationship) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPermissionResources

`func (o *HyperflexNodeRelationship) GetPermissionResources() []MoBaseMoRelationship`

GetPermissionResources returns the PermissionResources field if non-nil, zero value otherwise.

### GetPermissionResourcesOk

`func (o *HyperflexNodeRelationship) GetPermissionResourcesOk() (*[]MoBaseMoRelationship, bool)`

GetPermissionResourcesOk returns a tuple with the PermissionResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionResources

`func (o *HyperflexNodeRelationship) SetPermissionResources(v []MoBaseMoRelationship)`

SetPermissionResources sets PermissionResources field to given value.

### HasPermissionResources

`func (o *HyperflexNodeRelationship) HasPermissionResources() bool`

HasPermissionResources returns a boolean if a field has been set.

### SetPermissionResourcesNil

`func (o *HyperflexNodeRelationship) SetPermissionResourcesNil(b bool)`

 SetPermissionResourcesNil sets the value for PermissionResources to be an explicit nil

### UnsetPermissionResources
`func (o *HyperflexNodeRelationship) UnsetPermissionResources()`

UnsetPermissionResources ensures that no value is present for PermissionResources, not even an explicit nil
### GetDisplayNames

`func (o *HyperflexNodeRelationship) GetDisplayNames() map[string][]string`

GetDisplayNames returns the DisplayNames field if non-nil, zero value otherwise.

### GetDisplayNamesOk

`func (o *HyperflexNodeRelationship) GetDisplayNamesOk() (*map[string][]string, bool)`

GetDisplayNamesOk returns a tuple with the DisplayNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNames

`func (o *HyperflexNodeRelationship) SetDisplayNames(v map[string][]string)`

SetDisplayNames sets DisplayNames field to given value.

### HasDisplayNames

`func (o *HyperflexNodeRelationship) HasDisplayNames() bool`

HasDisplayNames returns a boolean if a field has been set.

### SetDisplayNamesNil

`func (o *HyperflexNodeRelationship) SetDisplayNamesNil(b bool)`

 SetDisplayNamesNil sets the value for DisplayNames to be an explicit nil

### UnsetDisplayNames
`func (o *HyperflexNodeRelationship) UnsetDisplayNames()`

UnsetDisplayNames ensures that no value is present for DisplayNames, not even an explicit nil
### GetBuildNumber

`func (o *HyperflexNodeRelationship) GetBuildNumber() string`

GetBuildNumber returns the BuildNumber field if non-nil, zero value otherwise.

### GetBuildNumberOk

`func (o *HyperflexNodeRelationship) GetBuildNumberOk() (*string, bool)`

GetBuildNumberOk returns a tuple with the BuildNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildNumber

`func (o *HyperflexNodeRelationship) SetBuildNumber(v string)`

SetBuildNumber sets BuildNumber field to given value.

### HasBuildNumber

`func (o *HyperflexNodeRelationship) HasBuildNumber() bool`

HasBuildNumber returns a boolean if a field has been set.

### GetDisplayVersion

`func (o *HyperflexNodeRelationship) GetDisplayVersion() string`

GetDisplayVersion returns the DisplayVersion field if non-nil, zero value otherwise.

### GetDisplayVersionOk

`func (o *HyperflexNodeRelationship) GetDisplayVersionOk() (*string, bool)`

GetDisplayVersionOk returns a tuple with the DisplayVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayVersion

`func (o *HyperflexNodeRelationship) SetDisplayVersion(v string)`

SetDisplayVersion sets DisplayVersion field to given value.

### HasDisplayVersion

`func (o *HyperflexNodeRelationship) HasDisplayVersion() bool`

HasDisplayVersion returns a boolean if a field has been set.

### GetHostName

`func (o *HyperflexNodeRelationship) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *HyperflexNodeRelationship) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *HyperflexNodeRelationship) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *HyperflexNodeRelationship) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetHypervisor

`func (o *HyperflexNodeRelationship) GetHypervisor() string`

GetHypervisor returns the Hypervisor field if non-nil, zero value otherwise.

### GetHypervisorOk

`func (o *HyperflexNodeRelationship) GetHypervisorOk() (*string, bool)`

GetHypervisorOk returns a tuple with the Hypervisor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisor

`func (o *HyperflexNodeRelationship) SetHypervisor(v string)`

SetHypervisor sets Hypervisor field to given value.

### HasHypervisor

`func (o *HyperflexNodeRelationship) HasHypervisor() bool`

HasHypervisor returns a boolean if a field has been set.

### GetIdentity

`func (o *HyperflexNodeRelationship) GetIdentity() HyperflexHxUuIdDt`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *HyperflexNodeRelationship) GetIdentityOk() (*HyperflexHxUuIdDt, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *HyperflexNodeRelationship) SetIdentity(v HyperflexHxUuIdDt)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *HyperflexNodeRelationship) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetIp

`func (o *HyperflexNodeRelationship) GetIp() HyperflexHxNetworkAddressDt`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *HyperflexNodeRelationship) GetIpOk() (*HyperflexHxNetworkAddressDt, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *HyperflexNodeRelationship) SetIp(v HyperflexHxNetworkAddressDt)`

SetIp sets Ip field to given value.

### HasIp

`func (o *HyperflexNodeRelationship) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetLockdown

`func (o *HyperflexNodeRelationship) GetLockdown() bool`

GetLockdown returns the Lockdown field if non-nil, zero value otherwise.

### GetLockdownOk

`func (o *HyperflexNodeRelationship) GetLockdownOk() (*bool, bool)`

GetLockdownOk returns a tuple with the Lockdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockdown

`func (o *HyperflexNodeRelationship) SetLockdown(v bool)`

SetLockdown sets Lockdown field to given value.

### HasLockdown

`func (o *HyperflexNodeRelationship) HasLockdown() bool`

HasLockdown returns a boolean if a field has been set.

### GetModelNumber

`func (o *HyperflexNodeRelationship) GetModelNumber() string`

GetModelNumber returns the ModelNumber field if non-nil, zero value otherwise.

### GetModelNumberOk

`func (o *HyperflexNodeRelationship) GetModelNumberOk() (*string, bool)`

GetModelNumberOk returns a tuple with the ModelNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelNumber

`func (o *HyperflexNodeRelationship) SetModelNumber(v string)`

SetModelNumber sets ModelNumber field to given value.

### HasModelNumber

`func (o *HyperflexNodeRelationship) HasModelNumber() bool`

HasModelNumber returns a boolean if a field has been set.

### GetRole

`func (o *HyperflexNodeRelationship) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *HyperflexNodeRelationship) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *HyperflexNodeRelationship) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *HyperflexNodeRelationship) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetSerialNumber

`func (o *HyperflexNodeRelationship) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *HyperflexNodeRelationship) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *HyperflexNodeRelationship) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *HyperflexNodeRelationship) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetStatus

`func (o *HyperflexNodeRelationship) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HyperflexNodeRelationship) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HyperflexNodeRelationship) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HyperflexNodeRelationship) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVersion

`func (o *HyperflexNodeRelationship) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *HyperflexNodeRelationship) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *HyperflexNodeRelationship) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *HyperflexNodeRelationship) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetCluster

`func (o *HyperflexNodeRelationship) GetCluster() HyperflexClusterRelationship`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *HyperflexNodeRelationship) GetClusterOk() (*HyperflexClusterRelationship, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *HyperflexNodeRelationship) SetCluster(v HyperflexClusterRelationship)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *HyperflexNodeRelationship) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetClusterMember

`func (o *HyperflexNodeRelationship) GetClusterMember() AssetClusterMemberRelationship`

GetClusterMember returns the ClusterMember field if non-nil, zero value otherwise.

### GetClusterMemberOk

`func (o *HyperflexNodeRelationship) GetClusterMemberOk() (*AssetClusterMemberRelationship, bool)`

GetClusterMemberOk returns a tuple with the ClusterMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterMember

`func (o *HyperflexNodeRelationship) SetClusterMember(v AssetClusterMemberRelationship)`

SetClusterMember sets ClusterMember field to given value.

### HasClusterMember

`func (o *HyperflexNodeRelationship) HasClusterMember() bool`

HasClusterMember returns a boolean if a field has been set.

### GetPhysicalServer

`func (o *HyperflexNodeRelationship) GetPhysicalServer() ComputePhysicalRelationship`

GetPhysicalServer returns the PhysicalServer field if non-nil, zero value otherwise.

### GetPhysicalServerOk

`func (o *HyperflexNodeRelationship) GetPhysicalServerOk() (*ComputePhysicalRelationship, bool)`

GetPhysicalServerOk returns a tuple with the PhysicalServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalServer

`func (o *HyperflexNodeRelationship) SetPhysicalServer(v ComputePhysicalRelationship)`

SetPhysicalServer sets PhysicalServer field to given value.

### HasPhysicalServer

`func (o *HyperflexNodeRelationship) HasPhysicalServer() bool`

HasPhysicalServer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



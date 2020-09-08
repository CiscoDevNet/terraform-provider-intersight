# MemoryPersistentMemoryRegionRelationship

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
**DeviceMoId** | Pointer to **string** | The database identifier of the registered device of an object. | [optional] [readonly] 
**Dn** | Pointer to **string** | The Distinguished Name unambiguously identifies an object in the system. | [optional] [readonly] 
**Rn** | Pointer to **string** | The Relative Name uniquely identifies an object within a given context. | [optional] [readonly] 
**FreeCapacity** | Pointer to **string** | Free capacity in GiB of the Persistent Memory Region. | [optional] [readonly] 
**HealthState** | Pointer to **string** | Health state of the Persistent Memory Region. | [optional] [readonly] 
**InterleavedSetId** | Pointer to **string** | ID of the Interleaved Set formed for this Persistent Memory Region. | [optional] [readonly] 
**LocaterIds** | Pointer to **string** | Set of locator IDs that are included in the Persistent Memory Region. | [optional] [readonly] 
**PersistentMemoryType** | Pointer to **string** | Persistent Memory type of the Persistent Memory Region. | [optional] [readonly] 
**RegionId** | Pointer to **string** | ID of the Persistent Memory Region. | [optional] [readonly] 
**SocketId** | Pointer to **string** | Socket ID of the Persistent Memory Region. | [optional] [readonly] 
**SocketMemoryId** | Pointer to **string** | Socket Memory ID of the Persistent Memory Region. | [optional] [readonly] 
**TotalCapacity** | Pointer to **string** | Total capacity in GiB of the Persistent Memory Region. | [optional] [readonly] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](inventory.DeviceInfo.Relationship.md) |  | [optional] 
**MemoryPersistentMemoryConfiguration** | Pointer to [**MemoryPersistentMemoryConfigurationRelationship**](memory.PersistentMemoryConfiguration.Relationship.md) |  | [optional] 
**PersistentMemoryNamespaces** | Pointer to [**[]MemoryPersistentMemoryNamespaceRelationship**](memory.PersistentMemoryNamespace.Relationship.md) | An array of relationships to memoryPersistentMemoryNamespace resources. | [optional] [readonly] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewMemoryPersistentMemoryRegionRelationship

`func NewMemoryPersistentMemoryRegionRelationship(classId string, objectType string, ) *MemoryPersistentMemoryRegionRelationship`

NewMemoryPersistentMemoryRegionRelationship instantiates a new MemoryPersistentMemoryRegionRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemoryPersistentMemoryRegionRelationshipWithDefaults

`func NewMemoryPersistentMemoryRegionRelationshipWithDefaults() *MemoryPersistentMemoryRegionRelationship`

NewMemoryPersistentMemoryRegionRelationshipWithDefaults instantiates a new MemoryPersistentMemoryRegionRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountMoid

`func (o *MemoryPersistentMemoryRegionRelationship) GetAccountMoid() string`

GetAccountMoid returns the AccountMoid field if non-nil, zero value otherwise.

### GetAccountMoidOk

`func (o *MemoryPersistentMemoryRegionRelationship) GetAccountMoidOk() (*string, bool)`

GetAccountMoidOk returns a tuple with the AccountMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMoid

`func (o *MemoryPersistentMemoryRegionRelationship) SetAccountMoid(v string)`

SetAccountMoid sets AccountMoid field to given value.

### HasAccountMoid

`func (o *MemoryPersistentMemoryRegionRelationship) HasAccountMoid() bool`

HasAccountMoid returns a boolean if a field has been set.

### GetClassId

`func (o *MemoryPersistentMemoryRegionRelationship) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *MemoryPersistentMemoryRegionRelationship) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *MemoryPersistentMemoryRegionRelationship) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetCreateTime

`func (o *MemoryPersistentMemoryRegionRelationship) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *MemoryPersistentMemoryRegionRelationship) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *MemoryPersistentMemoryRegionRelationship) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *MemoryPersistentMemoryRegionRelationship) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDomainGroupMoid

`func (o *MemoryPersistentMemoryRegionRelationship) GetDomainGroupMoid() string`

GetDomainGroupMoid returns the DomainGroupMoid field if non-nil, zero value otherwise.

### GetDomainGroupMoidOk

`func (o *MemoryPersistentMemoryRegionRelationship) GetDomainGroupMoidOk() (*string, bool)`

GetDomainGroupMoidOk returns a tuple with the DomainGroupMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGroupMoid

`func (o *MemoryPersistentMemoryRegionRelationship) SetDomainGroupMoid(v string)`

SetDomainGroupMoid sets DomainGroupMoid field to given value.

### HasDomainGroupMoid

`func (o *MemoryPersistentMemoryRegionRelationship) HasDomainGroupMoid() bool`

HasDomainGroupMoid returns a boolean if a field has been set.

### GetModTime

`func (o *MemoryPersistentMemoryRegionRelationship) GetModTime() time.Time`

GetModTime returns the ModTime field if non-nil, zero value otherwise.

### GetModTimeOk

`func (o *MemoryPersistentMemoryRegionRelationship) GetModTimeOk() (*time.Time, bool)`

GetModTimeOk returns a tuple with the ModTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModTime

`func (o *MemoryPersistentMemoryRegionRelationship) SetModTime(v time.Time)`

SetModTime sets ModTime field to given value.

### HasModTime

`func (o *MemoryPersistentMemoryRegionRelationship) HasModTime() bool`

HasModTime returns a boolean if a field has been set.

### GetMoid

`func (o *MemoryPersistentMemoryRegionRelationship) GetMoid() string`

GetMoid returns the Moid field if non-nil, zero value otherwise.

### GetMoidOk

`func (o *MemoryPersistentMemoryRegionRelationship) GetMoidOk() (*string, bool)`

GetMoidOk returns a tuple with the Moid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoid

`func (o *MemoryPersistentMemoryRegionRelationship) SetMoid(v string)`

SetMoid sets Moid field to given value.

### HasMoid

`func (o *MemoryPersistentMemoryRegionRelationship) HasMoid() bool`

HasMoid returns a boolean if a field has been set.

### GetObjectType

`func (o *MemoryPersistentMemoryRegionRelationship) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *MemoryPersistentMemoryRegionRelationship) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *MemoryPersistentMemoryRegionRelationship) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOwners

`func (o *MemoryPersistentMemoryRegionRelationship) GetOwners() []string`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *MemoryPersistentMemoryRegionRelationship) GetOwnersOk() (*[]string, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *MemoryPersistentMemoryRegionRelationship) SetOwners(v []string)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *MemoryPersistentMemoryRegionRelationship) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetSharedScope

`func (o *MemoryPersistentMemoryRegionRelationship) GetSharedScope() string`

GetSharedScope returns the SharedScope field if non-nil, zero value otherwise.

### GetSharedScopeOk

`func (o *MemoryPersistentMemoryRegionRelationship) GetSharedScopeOk() (*string, bool)`

GetSharedScopeOk returns a tuple with the SharedScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedScope

`func (o *MemoryPersistentMemoryRegionRelationship) SetSharedScope(v string)`

SetSharedScope sets SharedScope field to given value.

### HasSharedScope

`func (o *MemoryPersistentMemoryRegionRelationship) HasSharedScope() bool`

HasSharedScope returns a boolean if a field has been set.

### GetTags

`func (o *MemoryPersistentMemoryRegionRelationship) GetTags() []MoTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *MemoryPersistentMemoryRegionRelationship) GetTagsOk() (*[]MoTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *MemoryPersistentMemoryRegionRelationship) SetTags(v []MoTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *MemoryPersistentMemoryRegionRelationship) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersionContext

`func (o *MemoryPersistentMemoryRegionRelationship) GetVersionContext() MoVersionContext`

GetVersionContext returns the VersionContext field if non-nil, zero value otherwise.

### GetVersionContextOk

`func (o *MemoryPersistentMemoryRegionRelationship) GetVersionContextOk() (*MoVersionContext, bool)`

GetVersionContextOk returns a tuple with the VersionContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionContext

`func (o *MemoryPersistentMemoryRegionRelationship) SetVersionContext(v MoVersionContext)`

SetVersionContext sets VersionContext field to given value.

### HasVersionContext

`func (o *MemoryPersistentMemoryRegionRelationship) HasVersionContext() bool`

HasVersionContext returns a boolean if a field has been set.

### GetAncestors

`func (o *MemoryPersistentMemoryRegionRelationship) GetAncestors() []MoBaseMoRelationship`

GetAncestors returns the Ancestors field if non-nil, zero value otherwise.

### GetAncestorsOk

`func (o *MemoryPersistentMemoryRegionRelationship) GetAncestorsOk() (*[]MoBaseMoRelationship, bool)`

GetAncestorsOk returns a tuple with the Ancestors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestors

`func (o *MemoryPersistentMemoryRegionRelationship) SetAncestors(v []MoBaseMoRelationship)`

SetAncestors sets Ancestors field to given value.

### HasAncestors

`func (o *MemoryPersistentMemoryRegionRelationship) HasAncestors() bool`

HasAncestors returns a boolean if a field has been set.

### SetAncestorsNil

`func (o *MemoryPersistentMemoryRegionRelationship) SetAncestorsNil(b bool)`

 SetAncestorsNil sets the value for Ancestors to be an explicit nil

### UnsetAncestors
`func (o *MemoryPersistentMemoryRegionRelationship) UnsetAncestors()`

UnsetAncestors ensures that no value is present for Ancestors, not even an explicit nil
### GetParent

`func (o *MemoryPersistentMemoryRegionRelationship) GetParent() MoBaseMoRelationship`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *MemoryPersistentMemoryRegionRelationship) GetParentOk() (*MoBaseMoRelationship, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *MemoryPersistentMemoryRegionRelationship) SetParent(v MoBaseMoRelationship)`

SetParent sets Parent field to given value.

### HasParent

`func (o *MemoryPersistentMemoryRegionRelationship) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPermissionResources

`func (o *MemoryPersistentMemoryRegionRelationship) GetPermissionResources() []MoBaseMoRelationship`

GetPermissionResources returns the PermissionResources field if non-nil, zero value otherwise.

### GetPermissionResourcesOk

`func (o *MemoryPersistentMemoryRegionRelationship) GetPermissionResourcesOk() (*[]MoBaseMoRelationship, bool)`

GetPermissionResourcesOk returns a tuple with the PermissionResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionResources

`func (o *MemoryPersistentMemoryRegionRelationship) SetPermissionResources(v []MoBaseMoRelationship)`

SetPermissionResources sets PermissionResources field to given value.

### HasPermissionResources

`func (o *MemoryPersistentMemoryRegionRelationship) HasPermissionResources() bool`

HasPermissionResources returns a boolean if a field has been set.

### SetPermissionResourcesNil

`func (o *MemoryPersistentMemoryRegionRelationship) SetPermissionResourcesNil(b bool)`

 SetPermissionResourcesNil sets the value for PermissionResources to be an explicit nil

### UnsetPermissionResources
`func (o *MemoryPersistentMemoryRegionRelationship) UnsetPermissionResources()`

UnsetPermissionResources ensures that no value is present for PermissionResources, not even an explicit nil
### GetDisplayNames

`func (o *MemoryPersistentMemoryRegionRelationship) GetDisplayNames() map[string][]string`

GetDisplayNames returns the DisplayNames field if non-nil, zero value otherwise.

### GetDisplayNamesOk

`func (o *MemoryPersistentMemoryRegionRelationship) GetDisplayNamesOk() (*map[string][]string, bool)`

GetDisplayNamesOk returns a tuple with the DisplayNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNames

`func (o *MemoryPersistentMemoryRegionRelationship) SetDisplayNames(v map[string][]string)`

SetDisplayNames sets DisplayNames field to given value.

### HasDisplayNames

`func (o *MemoryPersistentMemoryRegionRelationship) HasDisplayNames() bool`

HasDisplayNames returns a boolean if a field has been set.

### SetDisplayNamesNil

`func (o *MemoryPersistentMemoryRegionRelationship) SetDisplayNamesNil(b bool)`

 SetDisplayNamesNil sets the value for DisplayNames to be an explicit nil

### UnsetDisplayNames
`func (o *MemoryPersistentMemoryRegionRelationship) UnsetDisplayNames()`

UnsetDisplayNames ensures that no value is present for DisplayNames, not even an explicit nil
### GetDeviceMoId

`func (o *MemoryPersistentMemoryRegionRelationship) GetDeviceMoId() string`

GetDeviceMoId returns the DeviceMoId field if non-nil, zero value otherwise.

### GetDeviceMoIdOk

`func (o *MemoryPersistentMemoryRegionRelationship) GetDeviceMoIdOk() (*string, bool)`

GetDeviceMoIdOk returns a tuple with the DeviceMoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMoId

`func (o *MemoryPersistentMemoryRegionRelationship) SetDeviceMoId(v string)`

SetDeviceMoId sets DeviceMoId field to given value.

### HasDeviceMoId

`func (o *MemoryPersistentMemoryRegionRelationship) HasDeviceMoId() bool`

HasDeviceMoId returns a boolean if a field has been set.

### GetDn

`func (o *MemoryPersistentMemoryRegionRelationship) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *MemoryPersistentMemoryRegionRelationship) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *MemoryPersistentMemoryRegionRelationship) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *MemoryPersistentMemoryRegionRelationship) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetRn

`func (o *MemoryPersistentMemoryRegionRelationship) GetRn() string`

GetRn returns the Rn field if non-nil, zero value otherwise.

### GetRnOk

`func (o *MemoryPersistentMemoryRegionRelationship) GetRnOk() (*string, bool)`

GetRnOk returns a tuple with the Rn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRn

`func (o *MemoryPersistentMemoryRegionRelationship) SetRn(v string)`

SetRn sets Rn field to given value.

### HasRn

`func (o *MemoryPersistentMemoryRegionRelationship) HasRn() bool`

HasRn returns a boolean if a field has been set.

### GetFreeCapacity

`func (o *MemoryPersistentMemoryRegionRelationship) GetFreeCapacity() string`

GetFreeCapacity returns the FreeCapacity field if non-nil, zero value otherwise.

### GetFreeCapacityOk

`func (o *MemoryPersistentMemoryRegionRelationship) GetFreeCapacityOk() (*string, bool)`

GetFreeCapacityOk returns a tuple with the FreeCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeCapacity

`func (o *MemoryPersistentMemoryRegionRelationship) SetFreeCapacity(v string)`

SetFreeCapacity sets FreeCapacity field to given value.

### HasFreeCapacity

`func (o *MemoryPersistentMemoryRegionRelationship) HasFreeCapacity() bool`

HasFreeCapacity returns a boolean if a field has been set.

### GetHealthState

`func (o *MemoryPersistentMemoryRegionRelationship) GetHealthState() string`

GetHealthState returns the HealthState field if non-nil, zero value otherwise.

### GetHealthStateOk

`func (o *MemoryPersistentMemoryRegionRelationship) GetHealthStateOk() (*string, bool)`

GetHealthStateOk returns a tuple with the HealthState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthState

`func (o *MemoryPersistentMemoryRegionRelationship) SetHealthState(v string)`

SetHealthState sets HealthState field to given value.

### HasHealthState

`func (o *MemoryPersistentMemoryRegionRelationship) HasHealthState() bool`

HasHealthState returns a boolean if a field has been set.

### GetInterleavedSetId

`func (o *MemoryPersistentMemoryRegionRelationship) GetInterleavedSetId() string`

GetInterleavedSetId returns the InterleavedSetId field if non-nil, zero value otherwise.

### GetInterleavedSetIdOk

`func (o *MemoryPersistentMemoryRegionRelationship) GetInterleavedSetIdOk() (*string, bool)`

GetInterleavedSetIdOk returns a tuple with the InterleavedSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterleavedSetId

`func (o *MemoryPersistentMemoryRegionRelationship) SetInterleavedSetId(v string)`

SetInterleavedSetId sets InterleavedSetId field to given value.

### HasInterleavedSetId

`func (o *MemoryPersistentMemoryRegionRelationship) HasInterleavedSetId() bool`

HasInterleavedSetId returns a boolean if a field has been set.

### GetLocaterIds

`func (o *MemoryPersistentMemoryRegionRelationship) GetLocaterIds() string`

GetLocaterIds returns the LocaterIds field if non-nil, zero value otherwise.

### GetLocaterIdsOk

`func (o *MemoryPersistentMemoryRegionRelationship) GetLocaterIdsOk() (*string, bool)`

GetLocaterIdsOk returns a tuple with the LocaterIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocaterIds

`func (o *MemoryPersistentMemoryRegionRelationship) SetLocaterIds(v string)`

SetLocaterIds sets LocaterIds field to given value.

### HasLocaterIds

`func (o *MemoryPersistentMemoryRegionRelationship) HasLocaterIds() bool`

HasLocaterIds returns a boolean if a field has been set.

### GetPersistentMemoryType

`func (o *MemoryPersistentMemoryRegionRelationship) GetPersistentMemoryType() string`

GetPersistentMemoryType returns the PersistentMemoryType field if non-nil, zero value otherwise.

### GetPersistentMemoryTypeOk

`func (o *MemoryPersistentMemoryRegionRelationship) GetPersistentMemoryTypeOk() (*string, bool)`

GetPersistentMemoryTypeOk returns a tuple with the PersistentMemoryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentMemoryType

`func (o *MemoryPersistentMemoryRegionRelationship) SetPersistentMemoryType(v string)`

SetPersistentMemoryType sets PersistentMemoryType field to given value.

### HasPersistentMemoryType

`func (o *MemoryPersistentMemoryRegionRelationship) HasPersistentMemoryType() bool`

HasPersistentMemoryType returns a boolean if a field has been set.

### GetRegionId

`func (o *MemoryPersistentMemoryRegionRelationship) GetRegionId() string`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *MemoryPersistentMemoryRegionRelationship) GetRegionIdOk() (*string, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *MemoryPersistentMemoryRegionRelationship) SetRegionId(v string)`

SetRegionId sets RegionId field to given value.

### HasRegionId

`func (o *MemoryPersistentMemoryRegionRelationship) HasRegionId() bool`

HasRegionId returns a boolean if a field has been set.

### GetSocketId

`func (o *MemoryPersistentMemoryRegionRelationship) GetSocketId() string`

GetSocketId returns the SocketId field if non-nil, zero value otherwise.

### GetSocketIdOk

`func (o *MemoryPersistentMemoryRegionRelationship) GetSocketIdOk() (*string, bool)`

GetSocketIdOk returns a tuple with the SocketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocketId

`func (o *MemoryPersistentMemoryRegionRelationship) SetSocketId(v string)`

SetSocketId sets SocketId field to given value.

### HasSocketId

`func (o *MemoryPersistentMemoryRegionRelationship) HasSocketId() bool`

HasSocketId returns a boolean if a field has been set.

### GetSocketMemoryId

`func (o *MemoryPersistentMemoryRegionRelationship) GetSocketMemoryId() string`

GetSocketMemoryId returns the SocketMemoryId field if non-nil, zero value otherwise.

### GetSocketMemoryIdOk

`func (o *MemoryPersistentMemoryRegionRelationship) GetSocketMemoryIdOk() (*string, bool)`

GetSocketMemoryIdOk returns a tuple with the SocketMemoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocketMemoryId

`func (o *MemoryPersistentMemoryRegionRelationship) SetSocketMemoryId(v string)`

SetSocketMemoryId sets SocketMemoryId field to given value.

### HasSocketMemoryId

`func (o *MemoryPersistentMemoryRegionRelationship) HasSocketMemoryId() bool`

HasSocketMemoryId returns a boolean if a field has been set.

### GetTotalCapacity

`func (o *MemoryPersistentMemoryRegionRelationship) GetTotalCapacity() string`

GetTotalCapacity returns the TotalCapacity field if non-nil, zero value otherwise.

### GetTotalCapacityOk

`func (o *MemoryPersistentMemoryRegionRelationship) GetTotalCapacityOk() (*string, bool)`

GetTotalCapacityOk returns a tuple with the TotalCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCapacity

`func (o *MemoryPersistentMemoryRegionRelationship) SetTotalCapacity(v string)`

SetTotalCapacity sets TotalCapacity field to given value.

### HasTotalCapacity

`func (o *MemoryPersistentMemoryRegionRelationship) HasTotalCapacity() bool`

HasTotalCapacity returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *MemoryPersistentMemoryRegionRelationship) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *MemoryPersistentMemoryRegionRelationship) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *MemoryPersistentMemoryRegionRelationship) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *MemoryPersistentMemoryRegionRelationship) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetMemoryPersistentMemoryConfiguration

`func (o *MemoryPersistentMemoryRegionRelationship) GetMemoryPersistentMemoryConfiguration() MemoryPersistentMemoryConfigurationRelationship`

GetMemoryPersistentMemoryConfiguration returns the MemoryPersistentMemoryConfiguration field if non-nil, zero value otherwise.

### GetMemoryPersistentMemoryConfigurationOk

`func (o *MemoryPersistentMemoryRegionRelationship) GetMemoryPersistentMemoryConfigurationOk() (*MemoryPersistentMemoryConfigurationRelationship, bool)`

GetMemoryPersistentMemoryConfigurationOk returns a tuple with the MemoryPersistentMemoryConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryPersistentMemoryConfiguration

`func (o *MemoryPersistentMemoryRegionRelationship) SetMemoryPersistentMemoryConfiguration(v MemoryPersistentMemoryConfigurationRelationship)`

SetMemoryPersistentMemoryConfiguration sets MemoryPersistentMemoryConfiguration field to given value.

### HasMemoryPersistentMemoryConfiguration

`func (o *MemoryPersistentMemoryRegionRelationship) HasMemoryPersistentMemoryConfiguration() bool`

HasMemoryPersistentMemoryConfiguration returns a boolean if a field has been set.

### GetPersistentMemoryNamespaces

`func (o *MemoryPersistentMemoryRegionRelationship) GetPersistentMemoryNamespaces() []MemoryPersistentMemoryNamespaceRelationship`

GetPersistentMemoryNamespaces returns the PersistentMemoryNamespaces field if non-nil, zero value otherwise.

### GetPersistentMemoryNamespacesOk

`func (o *MemoryPersistentMemoryRegionRelationship) GetPersistentMemoryNamespacesOk() (*[]MemoryPersistentMemoryNamespaceRelationship, bool)`

GetPersistentMemoryNamespacesOk returns a tuple with the PersistentMemoryNamespaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentMemoryNamespaces

`func (o *MemoryPersistentMemoryRegionRelationship) SetPersistentMemoryNamespaces(v []MemoryPersistentMemoryNamespaceRelationship)`

SetPersistentMemoryNamespaces sets PersistentMemoryNamespaces field to given value.

### HasPersistentMemoryNamespaces

`func (o *MemoryPersistentMemoryRegionRelationship) HasPersistentMemoryNamespaces() bool`

HasPersistentMemoryNamespaces returns a boolean if a field has been set.

### SetPersistentMemoryNamespacesNil

`func (o *MemoryPersistentMemoryRegionRelationship) SetPersistentMemoryNamespacesNil(b bool)`

 SetPersistentMemoryNamespacesNil sets the value for PersistentMemoryNamespaces to be an explicit nil

### UnsetPersistentMemoryNamespaces
`func (o *MemoryPersistentMemoryRegionRelationship) UnsetPersistentMemoryNamespaces()`

UnsetPersistentMemoryNamespaces ensures that no value is present for PersistentMemoryNamespaces, not even an explicit nil
### GetRegisteredDevice

`func (o *MemoryPersistentMemoryRegionRelationship) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *MemoryPersistentMemoryRegionRelationship) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *MemoryPersistentMemoryRegionRelationship) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *MemoryPersistentMemoryRegionRelationship) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



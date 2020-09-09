# StoragePhysicalDiskUsageRelationship

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
**NumberOfBlocks** | Pointer to **string** | The number of blocks that are a part of the virtual drive. | [optional] [readonly] 
**PhysicalDrive** | Pointer to **string** | The physical disk for which the usage is reported. | [optional] [readonly] 
**Span** | Pointer to **string** | The span of the physical disk. | [optional] [readonly] 
**StartingBlock** | Pointer to **string** | The starting block id of the virtual drive within the physical drive. | [optional] [readonly] 
**State** | Pointer to **string** | The current state of the physical disk usage. | [optional] [readonly] 
**VirtualDrive** | Pointer to **string** | The virtual drive corresponding to the physical disk. | [optional] [readonly] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](inventory.DeviceInfo.Relationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewStoragePhysicalDiskUsageRelationship

`func NewStoragePhysicalDiskUsageRelationship(classId string, objectType string, ) *StoragePhysicalDiskUsageRelationship`

NewStoragePhysicalDiskUsageRelationship instantiates a new StoragePhysicalDiskUsageRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoragePhysicalDiskUsageRelationshipWithDefaults

`func NewStoragePhysicalDiskUsageRelationshipWithDefaults() *StoragePhysicalDiskUsageRelationship`

NewStoragePhysicalDiskUsageRelationshipWithDefaults instantiates a new StoragePhysicalDiskUsageRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountMoid

`func (o *StoragePhysicalDiskUsageRelationship) GetAccountMoid() string`

GetAccountMoid returns the AccountMoid field if non-nil, zero value otherwise.

### GetAccountMoidOk

`func (o *StoragePhysicalDiskUsageRelationship) GetAccountMoidOk() (*string, bool)`

GetAccountMoidOk returns a tuple with the AccountMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMoid

`func (o *StoragePhysicalDiskUsageRelationship) SetAccountMoid(v string)`

SetAccountMoid sets AccountMoid field to given value.

### HasAccountMoid

`func (o *StoragePhysicalDiskUsageRelationship) HasAccountMoid() bool`

HasAccountMoid returns a boolean if a field has been set.

### GetClassId

`func (o *StoragePhysicalDiskUsageRelationship) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StoragePhysicalDiskUsageRelationship) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StoragePhysicalDiskUsageRelationship) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetCreateTime

`func (o *StoragePhysicalDiskUsageRelationship) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *StoragePhysicalDiskUsageRelationship) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *StoragePhysicalDiskUsageRelationship) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *StoragePhysicalDiskUsageRelationship) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDomainGroupMoid

`func (o *StoragePhysicalDiskUsageRelationship) GetDomainGroupMoid() string`

GetDomainGroupMoid returns the DomainGroupMoid field if non-nil, zero value otherwise.

### GetDomainGroupMoidOk

`func (o *StoragePhysicalDiskUsageRelationship) GetDomainGroupMoidOk() (*string, bool)`

GetDomainGroupMoidOk returns a tuple with the DomainGroupMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGroupMoid

`func (o *StoragePhysicalDiskUsageRelationship) SetDomainGroupMoid(v string)`

SetDomainGroupMoid sets DomainGroupMoid field to given value.

### HasDomainGroupMoid

`func (o *StoragePhysicalDiskUsageRelationship) HasDomainGroupMoid() bool`

HasDomainGroupMoid returns a boolean if a field has been set.

### GetModTime

`func (o *StoragePhysicalDiskUsageRelationship) GetModTime() time.Time`

GetModTime returns the ModTime field if non-nil, zero value otherwise.

### GetModTimeOk

`func (o *StoragePhysicalDiskUsageRelationship) GetModTimeOk() (*time.Time, bool)`

GetModTimeOk returns a tuple with the ModTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModTime

`func (o *StoragePhysicalDiskUsageRelationship) SetModTime(v time.Time)`

SetModTime sets ModTime field to given value.

### HasModTime

`func (o *StoragePhysicalDiskUsageRelationship) HasModTime() bool`

HasModTime returns a boolean if a field has been set.

### GetMoid

`func (o *StoragePhysicalDiskUsageRelationship) GetMoid() string`

GetMoid returns the Moid field if non-nil, zero value otherwise.

### GetMoidOk

`func (o *StoragePhysicalDiskUsageRelationship) GetMoidOk() (*string, bool)`

GetMoidOk returns a tuple with the Moid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoid

`func (o *StoragePhysicalDiskUsageRelationship) SetMoid(v string)`

SetMoid sets Moid field to given value.

### HasMoid

`func (o *StoragePhysicalDiskUsageRelationship) HasMoid() bool`

HasMoid returns a boolean if a field has been set.

### GetObjectType

`func (o *StoragePhysicalDiskUsageRelationship) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StoragePhysicalDiskUsageRelationship) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StoragePhysicalDiskUsageRelationship) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOwners

`func (o *StoragePhysicalDiskUsageRelationship) GetOwners() []string`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *StoragePhysicalDiskUsageRelationship) GetOwnersOk() (*[]string, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *StoragePhysicalDiskUsageRelationship) SetOwners(v []string)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *StoragePhysicalDiskUsageRelationship) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetSharedScope

`func (o *StoragePhysicalDiskUsageRelationship) GetSharedScope() string`

GetSharedScope returns the SharedScope field if non-nil, zero value otherwise.

### GetSharedScopeOk

`func (o *StoragePhysicalDiskUsageRelationship) GetSharedScopeOk() (*string, bool)`

GetSharedScopeOk returns a tuple with the SharedScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedScope

`func (o *StoragePhysicalDiskUsageRelationship) SetSharedScope(v string)`

SetSharedScope sets SharedScope field to given value.

### HasSharedScope

`func (o *StoragePhysicalDiskUsageRelationship) HasSharedScope() bool`

HasSharedScope returns a boolean if a field has been set.

### GetTags

`func (o *StoragePhysicalDiskUsageRelationship) GetTags() []MoTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *StoragePhysicalDiskUsageRelationship) GetTagsOk() (*[]MoTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *StoragePhysicalDiskUsageRelationship) SetTags(v []MoTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *StoragePhysicalDiskUsageRelationship) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersionContext

`func (o *StoragePhysicalDiskUsageRelationship) GetVersionContext() MoVersionContext`

GetVersionContext returns the VersionContext field if non-nil, zero value otherwise.

### GetVersionContextOk

`func (o *StoragePhysicalDiskUsageRelationship) GetVersionContextOk() (*MoVersionContext, bool)`

GetVersionContextOk returns a tuple with the VersionContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionContext

`func (o *StoragePhysicalDiskUsageRelationship) SetVersionContext(v MoVersionContext)`

SetVersionContext sets VersionContext field to given value.

### HasVersionContext

`func (o *StoragePhysicalDiskUsageRelationship) HasVersionContext() bool`

HasVersionContext returns a boolean if a field has been set.

### GetAncestors

`func (o *StoragePhysicalDiskUsageRelationship) GetAncestors() []MoBaseMoRelationship`

GetAncestors returns the Ancestors field if non-nil, zero value otherwise.

### GetAncestorsOk

`func (o *StoragePhysicalDiskUsageRelationship) GetAncestorsOk() (*[]MoBaseMoRelationship, bool)`

GetAncestorsOk returns a tuple with the Ancestors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestors

`func (o *StoragePhysicalDiskUsageRelationship) SetAncestors(v []MoBaseMoRelationship)`

SetAncestors sets Ancestors field to given value.

### HasAncestors

`func (o *StoragePhysicalDiskUsageRelationship) HasAncestors() bool`

HasAncestors returns a boolean if a field has been set.

### SetAncestorsNil

`func (o *StoragePhysicalDiskUsageRelationship) SetAncestorsNil(b bool)`

 SetAncestorsNil sets the value for Ancestors to be an explicit nil

### UnsetAncestors
`func (o *StoragePhysicalDiskUsageRelationship) UnsetAncestors()`

UnsetAncestors ensures that no value is present for Ancestors, not even an explicit nil
### GetParent

`func (o *StoragePhysicalDiskUsageRelationship) GetParent() MoBaseMoRelationship`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *StoragePhysicalDiskUsageRelationship) GetParentOk() (*MoBaseMoRelationship, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *StoragePhysicalDiskUsageRelationship) SetParent(v MoBaseMoRelationship)`

SetParent sets Parent field to given value.

### HasParent

`func (o *StoragePhysicalDiskUsageRelationship) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPermissionResources

`func (o *StoragePhysicalDiskUsageRelationship) GetPermissionResources() []MoBaseMoRelationship`

GetPermissionResources returns the PermissionResources field if non-nil, zero value otherwise.

### GetPermissionResourcesOk

`func (o *StoragePhysicalDiskUsageRelationship) GetPermissionResourcesOk() (*[]MoBaseMoRelationship, bool)`

GetPermissionResourcesOk returns a tuple with the PermissionResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionResources

`func (o *StoragePhysicalDiskUsageRelationship) SetPermissionResources(v []MoBaseMoRelationship)`

SetPermissionResources sets PermissionResources field to given value.

### HasPermissionResources

`func (o *StoragePhysicalDiskUsageRelationship) HasPermissionResources() bool`

HasPermissionResources returns a boolean if a field has been set.

### SetPermissionResourcesNil

`func (o *StoragePhysicalDiskUsageRelationship) SetPermissionResourcesNil(b bool)`

 SetPermissionResourcesNil sets the value for PermissionResources to be an explicit nil

### UnsetPermissionResources
`func (o *StoragePhysicalDiskUsageRelationship) UnsetPermissionResources()`

UnsetPermissionResources ensures that no value is present for PermissionResources, not even an explicit nil
### GetDisplayNames

`func (o *StoragePhysicalDiskUsageRelationship) GetDisplayNames() map[string][]string`

GetDisplayNames returns the DisplayNames field if non-nil, zero value otherwise.

### GetDisplayNamesOk

`func (o *StoragePhysicalDiskUsageRelationship) GetDisplayNamesOk() (*map[string][]string, bool)`

GetDisplayNamesOk returns a tuple with the DisplayNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNames

`func (o *StoragePhysicalDiskUsageRelationship) SetDisplayNames(v map[string][]string)`

SetDisplayNames sets DisplayNames field to given value.

### HasDisplayNames

`func (o *StoragePhysicalDiskUsageRelationship) HasDisplayNames() bool`

HasDisplayNames returns a boolean if a field has been set.

### SetDisplayNamesNil

`func (o *StoragePhysicalDiskUsageRelationship) SetDisplayNamesNil(b bool)`

 SetDisplayNamesNil sets the value for DisplayNames to be an explicit nil

### UnsetDisplayNames
`func (o *StoragePhysicalDiskUsageRelationship) UnsetDisplayNames()`

UnsetDisplayNames ensures that no value is present for DisplayNames, not even an explicit nil
### GetDeviceMoId

`func (o *StoragePhysicalDiskUsageRelationship) GetDeviceMoId() string`

GetDeviceMoId returns the DeviceMoId field if non-nil, zero value otherwise.

### GetDeviceMoIdOk

`func (o *StoragePhysicalDiskUsageRelationship) GetDeviceMoIdOk() (*string, bool)`

GetDeviceMoIdOk returns a tuple with the DeviceMoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMoId

`func (o *StoragePhysicalDiskUsageRelationship) SetDeviceMoId(v string)`

SetDeviceMoId sets DeviceMoId field to given value.

### HasDeviceMoId

`func (o *StoragePhysicalDiskUsageRelationship) HasDeviceMoId() bool`

HasDeviceMoId returns a boolean if a field has been set.

### GetDn

`func (o *StoragePhysicalDiskUsageRelationship) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *StoragePhysicalDiskUsageRelationship) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *StoragePhysicalDiskUsageRelationship) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *StoragePhysicalDiskUsageRelationship) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetRn

`func (o *StoragePhysicalDiskUsageRelationship) GetRn() string`

GetRn returns the Rn field if non-nil, zero value otherwise.

### GetRnOk

`func (o *StoragePhysicalDiskUsageRelationship) GetRnOk() (*string, bool)`

GetRnOk returns a tuple with the Rn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRn

`func (o *StoragePhysicalDiskUsageRelationship) SetRn(v string)`

SetRn sets Rn field to given value.

### HasRn

`func (o *StoragePhysicalDiskUsageRelationship) HasRn() bool`

HasRn returns a boolean if a field has been set.

### GetNumberOfBlocks

`func (o *StoragePhysicalDiskUsageRelationship) GetNumberOfBlocks() string`

GetNumberOfBlocks returns the NumberOfBlocks field if non-nil, zero value otherwise.

### GetNumberOfBlocksOk

`func (o *StoragePhysicalDiskUsageRelationship) GetNumberOfBlocksOk() (*string, bool)`

GetNumberOfBlocksOk returns a tuple with the NumberOfBlocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfBlocks

`func (o *StoragePhysicalDiskUsageRelationship) SetNumberOfBlocks(v string)`

SetNumberOfBlocks sets NumberOfBlocks field to given value.

### HasNumberOfBlocks

`func (o *StoragePhysicalDiskUsageRelationship) HasNumberOfBlocks() bool`

HasNumberOfBlocks returns a boolean if a field has been set.

### GetPhysicalDrive

`func (o *StoragePhysicalDiskUsageRelationship) GetPhysicalDrive() string`

GetPhysicalDrive returns the PhysicalDrive field if non-nil, zero value otherwise.

### GetPhysicalDriveOk

`func (o *StoragePhysicalDiskUsageRelationship) GetPhysicalDriveOk() (*string, bool)`

GetPhysicalDriveOk returns a tuple with the PhysicalDrive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalDrive

`func (o *StoragePhysicalDiskUsageRelationship) SetPhysicalDrive(v string)`

SetPhysicalDrive sets PhysicalDrive field to given value.

### HasPhysicalDrive

`func (o *StoragePhysicalDiskUsageRelationship) HasPhysicalDrive() bool`

HasPhysicalDrive returns a boolean if a field has been set.

### GetSpan

`func (o *StoragePhysicalDiskUsageRelationship) GetSpan() string`

GetSpan returns the Span field if non-nil, zero value otherwise.

### GetSpanOk

`func (o *StoragePhysicalDiskUsageRelationship) GetSpanOk() (*string, bool)`

GetSpanOk returns a tuple with the Span field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpan

`func (o *StoragePhysicalDiskUsageRelationship) SetSpan(v string)`

SetSpan sets Span field to given value.

### HasSpan

`func (o *StoragePhysicalDiskUsageRelationship) HasSpan() bool`

HasSpan returns a boolean if a field has been set.

### GetStartingBlock

`func (o *StoragePhysicalDiskUsageRelationship) GetStartingBlock() string`

GetStartingBlock returns the StartingBlock field if non-nil, zero value otherwise.

### GetStartingBlockOk

`func (o *StoragePhysicalDiskUsageRelationship) GetStartingBlockOk() (*string, bool)`

GetStartingBlockOk returns a tuple with the StartingBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartingBlock

`func (o *StoragePhysicalDiskUsageRelationship) SetStartingBlock(v string)`

SetStartingBlock sets StartingBlock field to given value.

### HasStartingBlock

`func (o *StoragePhysicalDiskUsageRelationship) HasStartingBlock() bool`

HasStartingBlock returns a boolean if a field has been set.

### GetState

`func (o *StoragePhysicalDiskUsageRelationship) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *StoragePhysicalDiskUsageRelationship) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *StoragePhysicalDiskUsageRelationship) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *StoragePhysicalDiskUsageRelationship) HasState() bool`

HasState returns a boolean if a field has been set.

### GetVirtualDrive

`func (o *StoragePhysicalDiskUsageRelationship) GetVirtualDrive() string`

GetVirtualDrive returns the VirtualDrive field if non-nil, zero value otherwise.

### GetVirtualDriveOk

`func (o *StoragePhysicalDiskUsageRelationship) GetVirtualDriveOk() (*string, bool)`

GetVirtualDriveOk returns a tuple with the VirtualDrive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualDrive

`func (o *StoragePhysicalDiskUsageRelationship) SetVirtualDrive(v string)`

SetVirtualDrive sets VirtualDrive field to given value.

### HasVirtualDrive

`func (o *StoragePhysicalDiskUsageRelationship) HasVirtualDrive() bool`

HasVirtualDrive returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *StoragePhysicalDiskUsageRelationship) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *StoragePhysicalDiskUsageRelationship) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *StoragePhysicalDiskUsageRelationship) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *StoragePhysicalDiskUsageRelationship) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *StoragePhysicalDiskUsageRelationship) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StoragePhysicalDiskUsageRelationship) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StoragePhysicalDiskUsageRelationship) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StoragePhysicalDiskUsageRelationship) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



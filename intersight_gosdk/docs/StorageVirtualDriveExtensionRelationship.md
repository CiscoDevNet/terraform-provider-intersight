# StorageVirtualDriveExtensionRelationship

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
**Bootable** | Pointer to **string** | The ability to boot from the virtual drive. | [optional] [readonly] 
**ContainerId** | Pointer to **int64** | The container id of the virtual drive. | [optional] [readonly] 
**DriveState** | Pointer to **string** | The state of the virtual drive. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the Virtual drive. | [optional] [readonly] 
**OperDeviceId** | Pointer to **string** | The operational device id of the virtual drive. | [optional] [readonly] 
**Uuid** | Pointer to **string** | The UUID assigned to the virtual drive. | [optional] [readonly] 
**VendorUuid** | Pointer to **string** | The UUID value of the vendor assigned to the virtual drive. | [optional] [readonly] 
**VirtualDriveDn** | Pointer to **string** | The distinguished name of the virtual drive for which the extended data is provided. | [optional] [readonly] 
**VirtualDriveId** | Pointer to **string** | The Id of the virtual drive. | [optional] [readonly] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](inventory.DeviceInfo.Relationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 
**StorageController** | Pointer to [**StorageControllerRelationship**](storage.Controller.Relationship.md) |  | [optional] 
**VirtualDrive** | Pointer to [**StorageVirtualDriveRelationship**](storage.VirtualDrive.Relationship.md) |  | [optional] 

## Methods

### NewStorageVirtualDriveExtensionRelationship

`func NewStorageVirtualDriveExtensionRelationship(classId string, objectType string, ) *StorageVirtualDriveExtensionRelationship`

NewStorageVirtualDriveExtensionRelationship instantiates a new StorageVirtualDriveExtensionRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageVirtualDriveExtensionRelationshipWithDefaults

`func NewStorageVirtualDriveExtensionRelationshipWithDefaults() *StorageVirtualDriveExtensionRelationship`

NewStorageVirtualDriveExtensionRelationshipWithDefaults instantiates a new StorageVirtualDriveExtensionRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountMoid

`func (o *StorageVirtualDriveExtensionRelationship) GetAccountMoid() string`

GetAccountMoid returns the AccountMoid field if non-nil, zero value otherwise.

### GetAccountMoidOk

`func (o *StorageVirtualDriveExtensionRelationship) GetAccountMoidOk() (*string, bool)`

GetAccountMoidOk returns a tuple with the AccountMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMoid

`func (o *StorageVirtualDriveExtensionRelationship) SetAccountMoid(v string)`

SetAccountMoid sets AccountMoid field to given value.

### HasAccountMoid

`func (o *StorageVirtualDriveExtensionRelationship) HasAccountMoid() bool`

HasAccountMoid returns a boolean if a field has been set.

### GetClassId

`func (o *StorageVirtualDriveExtensionRelationship) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageVirtualDriveExtensionRelationship) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageVirtualDriveExtensionRelationship) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetCreateTime

`func (o *StorageVirtualDriveExtensionRelationship) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *StorageVirtualDriveExtensionRelationship) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *StorageVirtualDriveExtensionRelationship) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *StorageVirtualDriveExtensionRelationship) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDomainGroupMoid

`func (o *StorageVirtualDriveExtensionRelationship) GetDomainGroupMoid() string`

GetDomainGroupMoid returns the DomainGroupMoid field if non-nil, zero value otherwise.

### GetDomainGroupMoidOk

`func (o *StorageVirtualDriveExtensionRelationship) GetDomainGroupMoidOk() (*string, bool)`

GetDomainGroupMoidOk returns a tuple with the DomainGroupMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGroupMoid

`func (o *StorageVirtualDriveExtensionRelationship) SetDomainGroupMoid(v string)`

SetDomainGroupMoid sets DomainGroupMoid field to given value.

### HasDomainGroupMoid

`func (o *StorageVirtualDriveExtensionRelationship) HasDomainGroupMoid() bool`

HasDomainGroupMoid returns a boolean if a field has been set.

### GetModTime

`func (o *StorageVirtualDriveExtensionRelationship) GetModTime() time.Time`

GetModTime returns the ModTime field if non-nil, zero value otherwise.

### GetModTimeOk

`func (o *StorageVirtualDriveExtensionRelationship) GetModTimeOk() (*time.Time, bool)`

GetModTimeOk returns a tuple with the ModTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModTime

`func (o *StorageVirtualDriveExtensionRelationship) SetModTime(v time.Time)`

SetModTime sets ModTime field to given value.

### HasModTime

`func (o *StorageVirtualDriveExtensionRelationship) HasModTime() bool`

HasModTime returns a boolean if a field has been set.

### GetMoid

`func (o *StorageVirtualDriveExtensionRelationship) GetMoid() string`

GetMoid returns the Moid field if non-nil, zero value otherwise.

### GetMoidOk

`func (o *StorageVirtualDriveExtensionRelationship) GetMoidOk() (*string, bool)`

GetMoidOk returns a tuple with the Moid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoid

`func (o *StorageVirtualDriveExtensionRelationship) SetMoid(v string)`

SetMoid sets Moid field to given value.

### HasMoid

`func (o *StorageVirtualDriveExtensionRelationship) HasMoid() bool`

HasMoid returns a boolean if a field has been set.

### GetObjectType

`func (o *StorageVirtualDriveExtensionRelationship) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageVirtualDriveExtensionRelationship) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageVirtualDriveExtensionRelationship) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOwners

`func (o *StorageVirtualDriveExtensionRelationship) GetOwners() []string`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *StorageVirtualDriveExtensionRelationship) GetOwnersOk() (*[]string, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *StorageVirtualDriveExtensionRelationship) SetOwners(v []string)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *StorageVirtualDriveExtensionRelationship) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetSharedScope

`func (o *StorageVirtualDriveExtensionRelationship) GetSharedScope() string`

GetSharedScope returns the SharedScope field if non-nil, zero value otherwise.

### GetSharedScopeOk

`func (o *StorageVirtualDriveExtensionRelationship) GetSharedScopeOk() (*string, bool)`

GetSharedScopeOk returns a tuple with the SharedScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedScope

`func (o *StorageVirtualDriveExtensionRelationship) SetSharedScope(v string)`

SetSharedScope sets SharedScope field to given value.

### HasSharedScope

`func (o *StorageVirtualDriveExtensionRelationship) HasSharedScope() bool`

HasSharedScope returns a boolean if a field has been set.

### GetTags

`func (o *StorageVirtualDriveExtensionRelationship) GetTags() []MoTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *StorageVirtualDriveExtensionRelationship) GetTagsOk() (*[]MoTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *StorageVirtualDriveExtensionRelationship) SetTags(v []MoTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *StorageVirtualDriveExtensionRelationship) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersionContext

`func (o *StorageVirtualDriveExtensionRelationship) GetVersionContext() MoVersionContext`

GetVersionContext returns the VersionContext field if non-nil, zero value otherwise.

### GetVersionContextOk

`func (o *StorageVirtualDriveExtensionRelationship) GetVersionContextOk() (*MoVersionContext, bool)`

GetVersionContextOk returns a tuple with the VersionContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionContext

`func (o *StorageVirtualDriveExtensionRelationship) SetVersionContext(v MoVersionContext)`

SetVersionContext sets VersionContext field to given value.

### HasVersionContext

`func (o *StorageVirtualDriveExtensionRelationship) HasVersionContext() bool`

HasVersionContext returns a boolean if a field has been set.

### GetAncestors

`func (o *StorageVirtualDriveExtensionRelationship) GetAncestors() []MoBaseMoRelationship`

GetAncestors returns the Ancestors field if non-nil, zero value otherwise.

### GetAncestorsOk

`func (o *StorageVirtualDriveExtensionRelationship) GetAncestorsOk() (*[]MoBaseMoRelationship, bool)`

GetAncestorsOk returns a tuple with the Ancestors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestors

`func (o *StorageVirtualDriveExtensionRelationship) SetAncestors(v []MoBaseMoRelationship)`

SetAncestors sets Ancestors field to given value.

### HasAncestors

`func (o *StorageVirtualDriveExtensionRelationship) HasAncestors() bool`

HasAncestors returns a boolean if a field has been set.

### SetAncestorsNil

`func (o *StorageVirtualDriveExtensionRelationship) SetAncestorsNil(b bool)`

 SetAncestorsNil sets the value for Ancestors to be an explicit nil

### UnsetAncestors
`func (o *StorageVirtualDriveExtensionRelationship) UnsetAncestors()`

UnsetAncestors ensures that no value is present for Ancestors, not even an explicit nil
### GetParent

`func (o *StorageVirtualDriveExtensionRelationship) GetParent() MoBaseMoRelationship`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *StorageVirtualDriveExtensionRelationship) GetParentOk() (*MoBaseMoRelationship, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *StorageVirtualDriveExtensionRelationship) SetParent(v MoBaseMoRelationship)`

SetParent sets Parent field to given value.

### HasParent

`func (o *StorageVirtualDriveExtensionRelationship) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPermissionResources

`func (o *StorageVirtualDriveExtensionRelationship) GetPermissionResources() []MoBaseMoRelationship`

GetPermissionResources returns the PermissionResources field if non-nil, zero value otherwise.

### GetPermissionResourcesOk

`func (o *StorageVirtualDriveExtensionRelationship) GetPermissionResourcesOk() (*[]MoBaseMoRelationship, bool)`

GetPermissionResourcesOk returns a tuple with the PermissionResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionResources

`func (o *StorageVirtualDriveExtensionRelationship) SetPermissionResources(v []MoBaseMoRelationship)`

SetPermissionResources sets PermissionResources field to given value.

### HasPermissionResources

`func (o *StorageVirtualDriveExtensionRelationship) HasPermissionResources() bool`

HasPermissionResources returns a boolean if a field has been set.

### SetPermissionResourcesNil

`func (o *StorageVirtualDriveExtensionRelationship) SetPermissionResourcesNil(b bool)`

 SetPermissionResourcesNil sets the value for PermissionResources to be an explicit nil

### UnsetPermissionResources
`func (o *StorageVirtualDriveExtensionRelationship) UnsetPermissionResources()`

UnsetPermissionResources ensures that no value is present for PermissionResources, not even an explicit nil
### GetDisplayNames

`func (o *StorageVirtualDriveExtensionRelationship) GetDisplayNames() map[string][]string`

GetDisplayNames returns the DisplayNames field if non-nil, zero value otherwise.

### GetDisplayNamesOk

`func (o *StorageVirtualDriveExtensionRelationship) GetDisplayNamesOk() (*map[string][]string, bool)`

GetDisplayNamesOk returns a tuple with the DisplayNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNames

`func (o *StorageVirtualDriveExtensionRelationship) SetDisplayNames(v map[string][]string)`

SetDisplayNames sets DisplayNames field to given value.

### HasDisplayNames

`func (o *StorageVirtualDriveExtensionRelationship) HasDisplayNames() bool`

HasDisplayNames returns a boolean if a field has been set.

### SetDisplayNamesNil

`func (o *StorageVirtualDriveExtensionRelationship) SetDisplayNamesNil(b bool)`

 SetDisplayNamesNil sets the value for DisplayNames to be an explicit nil

### UnsetDisplayNames
`func (o *StorageVirtualDriveExtensionRelationship) UnsetDisplayNames()`

UnsetDisplayNames ensures that no value is present for DisplayNames, not even an explicit nil
### GetDeviceMoId

`func (o *StorageVirtualDriveExtensionRelationship) GetDeviceMoId() string`

GetDeviceMoId returns the DeviceMoId field if non-nil, zero value otherwise.

### GetDeviceMoIdOk

`func (o *StorageVirtualDriveExtensionRelationship) GetDeviceMoIdOk() (*string, bool)`

GetDeviceMoIdOk returns a tuple with the DeviceMoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMoId

`func (o *StorageVirtualDriveExtensionRelationship) SetDeviceMoId(v string)`

SetDeviceMoId sets DeviceMoId field to given value.

### HasDeviceMoId

`func (o *StorageVirtualDriveExtensionRelationship) HasDeviceMoId() bool`

HasDeviceMoId returns a boolean if a field has been set.

### GetDn

`func (o *StorageVirtualDriveExtensionRelationship) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *StorageVirtualDriveExtensionRelationship) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *StorageVirtualDriveExtensionRelationship) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *StorageVirtualDriveExtensionRelationship) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetRn

`func (o *StorageVirtualDriveExtensionRelationship) GetRn() string`

GetRn returns the Rn field if non-nil, zero value otherwise.

### GetRnOk

`func (o *StorageVirtualDriveExtensionRelationship) GetRnOk() (*string, bool)`

GetRnOk returns a tuple with the Rn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRn

`func (o *StorageVirtualDriveExtensionRelationship) SetRn(v string)`

SetRn sets Rn field to given value.

### HasRn

`func (o *StorageVirtualDriveExtensionRelationship) HasRn() bool`

HasRn returns a boolean if a field has been set.

### GetBootable

`func (o *StorageVirtualDriveExtensionRelationship) GetBootable() string`

GetBootable returns the Bootable field if non-nil, zero value otherwise.

### GetBootableOk

`func (o *StorageVirtualDriveExtensionRelationship) GetBootableOk() (*string, bool)`

GetBootableOk returns a tuple with the Bootable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootable

`func (o *StorageVirtualDriveExtensionRelationship) SetBootable(v string)`

SetBootable sets Bootable field to given value.

### HasBootable

`func (o *StorageVirtualDriveExtensionRelationship) HasBootable() bool`

HasBootable returns a boolean if a field has been set.

### GetContainerId

`func (o *StorageVirtualDriveExtensionRelationship) GetContainerId() int64`

GetContainerId returns the ContainerId field if non-nil, zero value otherwise.

### GetContainerIdOk

`func (o *StorageVirtualDriveExtensionRelationship) GetContainerIdOk() (*int64, bool)`

GetContainerIdOk returns a tuple with the ContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerId

`func (o *StorageVirtualDriveExtensionRelationship) SetContainerId(v int64)`

SetContainerId sets ContainerId field to given value.

### HasContainerId

`func (o *StorageVirtualDriveExtensionRelationship) HasContainerId() bool`

HasContainerId returns a boolean if a field has been set.

### GetDriveState

`func (o *StorageVirtualDriveExtensionRelationship) GetDriveState() string`

GetDriveState returns the DriveState field if non-nil, zero value otherwise.

### GetDriveStateOk

`func (o *StorageVirtualDriveExtensionRelationship) GetDriveStateOk() (*string, bool)`

GetDriveStateOk returns a tuple with the DriveState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriveState

`func (o *StorageVirtualDriveExtensionRelationship) SetDriveState(v string)`

SetDriveState sets DriveState field to given value.

### HasDriveState

`func (o *StorageVirtualDriveExtensionRelationship) HasDriveState() bool`

HasDriveState returns a boolean if a field has been set.

### GetName

`func (o *StorageVirtualDriveExtensionRelationship) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageVirtualDriveExtensionRelationship) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageVirtualDriveExtensionRelationship) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StorageVirtualDriveExtensionRelationship) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOperDeviceId

`func (o *StorageVirtualDriveExtensionRelationship) GetOperDeviceId() string`

GetOperDeviceId returns the OperDeviceId field if non-nil, zero value otherwise.

### GetOperDeviceIdOk

`func (o *StorageVirtualDriveExtensionRelationship) GetOperDeviceIdOk() (*string, bool)`

GetOperDeviceIdOk returns a tuple with the OperDeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperDeviceId

`func (o *StorageVirtualDriveExtensionRelationship) SetOperDeviceId(v string)`

SetOperDeviceId sets OperDeviceId field to given value.

### HasOperDeviceId

`func (o *StorageVirtualDriveExtensionRelationship) HasOperDeviceId() bool`

HasOperDeviceId returns a boolean if a field has been set.

### GetUuid

`func (o *StorageVirtualDriveExtensionRelationship) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *StorageVirtualDriveExtensionRelationship) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *StorageVirtualDriveExtensionRelationship) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *StorageVirtualDriveExtensionRelationship) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetVendorUuid

`func (o *StorageVirtualDriveExtensionRelationship) GetVendorUuid() string`

GetVendorUuid returns the VendorUuid field if non-nil, zero value otherwise.

### GetVendorUuidOk

`func (o *StorageVirtualDriveExtensionRelationship) GetVendorUuidOk() (*string, bool)`

GetVendorUuidOk returns a tuple with the VendorUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorUuid

`func (o *StorageVirtualDriveExtensionRelationship) SetVendorUuid(v string)`

SetVendorUuid sets VendorUuid field to given value.

### HasVendorUuid

`func (o *StorageVirtualDriveExtensionRelationship) HasVendorUuid() bool`

HasVendorUuid returns a boolean if a field has been set.

### GetVirtualDriveDn

`func (o *StorageVirtualDriveExtensionRelationship) GetVirtualDriveDn() string`

GetVirtualDriveDn returns the VirtualDriveDn field if non-nil, zero value otherwise.

### GetVirtualDriveDnOk

`func (o *StorageVirtualDriveExtensionRelationship) GetVirtualDriveDnOk() (*string, bool)`

GetVirtualDriveDnOk returns a tuple with the VirtualDriveDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualDriveDn

`func (o *StorageVirtualDriveExtensionRelationship) SetVirtualDriveDn(v string)`

SetVirtualDriveDn sets VirtualDriveDn field to given value.

### HasVirtualDriveDn

`func (o *StorageVirtualDriveExtensionRelationship) HasVirtualDriveDn() bool`

HasVirtualDriveDn returns a boolean if a field has been set.

### GetVirtualDriveId

`func (o *StorageVirtualDriveExtensionRelationship) GetVirtualDriveId() string`

GetVirtualDriveId returns the VirtualDriveId field if non-nil, zero value otherwise.

### GetVirtualDriveIdOk

`func (o *StorageVirtualDriveExtensionRelationship) GetVirtualDriveIdOk() (*string, bool)`

GetVirtualDriveIdOk returns a tuple with the VirtualDriveId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualDriveId

`func (o *StorageVirtualDriveExtensionRelationship) SetVirtualDriveId(v string)`

SetVirtualDriveId sets VirtualDriveId field to given value.

### HasVirtualDriveId

`func (o *StorageVirtualDriveExtensionRelationship) HasVirtualDriveId() bool`

HasVirtualDriveId returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *StorageVirtualDriveExtensionRelationship) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *StorageVirtualDriveExtensionRelationship) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *StorageVirtualDriveExtensionRelationship) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *StorageVirtualDriveExtensionRelationship) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *StorageVirtualDriveExtensionRelationship) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StorageVirtualDriveExtensionRelationship) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StorageVirtualDriveExtensionRelationship) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StorageVirtualDriveExtensionRelationship) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### GetStorageController

`func (o *StorageVirtualDriveExtensionRelationship) GetStorageController() StorageControllerRelationship`

GetStorageController returns the StorageController field if non-nil, zero value otherwise.

### GetStorageControllerOk

`func (o *StorageVirtualDriveExtensionRelationship) GetStorageControllerOk() (*StorageControllerRelationship, bool)`

GetStorageControllerOk returns a tuple with the StorageController field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageController

`func (o *StorageVirtualDriveExtensionRelationship) SetStorageController(v StorageControllerRelationship)`

SetStorageController sets StorageController field to given value.

### HasStorageController

`func (o *StorageVirtualDriveExtensionRelationship) HasStorageController() bool`

HasStorageController returns a boolean if a field has been set.

### GetVirtualDrive

`func (o *StorageVirtualDriveExtensionRelationship) GetVirtualDrive() StorageVirtualDriveRelationship`

GetVirtualDrive returns the VirtualDrive field if non-nil, zero value otherwise.

### GetVirtualDriveOk

`func (o *StorageVirtualDriveExtensionRelationship) GetVirtualDriveOk() (*StorageVirtualDriveRelationship, bool)`

GetVirtualDriveOk returns a tuple with the VirtualDrive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualDrive

`func (o *StorageVirtualDriveExtensionRelationship) SetVirtualDrive(v StorageVirtualDriveRelationship)`

SetVirtualDrive sets VirtualDrive field to given value.

### HasVirtualDrive

`func (o *StorageVirtualDriveExtensionRelationship) HasVirtualDrive() bool`

HasVirtualDrive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



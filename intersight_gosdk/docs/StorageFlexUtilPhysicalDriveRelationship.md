# StorageFlexUtilPhysicalDriveRelationship

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
**Model** | Pointer to **string** | This field identifies the model of the given component. | [optional] [readonly] 
**Revision** | Pointer to **string** | This field identifies the revision of the given component. | [optional] [readonly] 
**Serial** | Pointer to **string** | This field identifies the serial of the given component. | [optional] [readonly] 
**Vendor** | Pointer to **string** | This field identifies the vendor of the given component. | [optional] [readonly] 
**BlockSize** | Pointer to **string** | Block size of the FlexUtil Physical drive. | [optional] 
**Capacity** | Pointer to **string** | Capacity of the FlexUtil Physical drive. | [optional] 
**Controller** | Pointer to **string** | Type of the Physical Drive Controller. | [optional] 
**DrivesEnabled** | Pointer to **string** | The number of drives enabled in the FlexUtil Physical Drive. | [optional] 
**Health** | Pointer to **string** | Health of the FlexUtil Physical drive. | [optional] 
**ManufacturerDate** | Pointer to **string** | Manufacturing date of the FlexUtil Physical Drive. | [optional] 
**ManufacturerId** | Pointer to **string** | Manufacturer identity of the FlexUtil Physical Drive. | [optional] 
**OemId** | Pointer to **string** | The OEM Identifier of the FlexUtil physical drive. | [optional] 
**PartitionCount** | Pointer to **string** | The number of partitions present on the FlexUtil Physical Drive. | [optional] 
**PdStatus** | Pointer to **string** | Status of the FlexUtil Physical Drive. | [optional] 
**PhysicalDrive** | Pointer to **string** | The type of physical drive. Example - microSD. | [optional] 
**ProductName** | Pointer to **string** | Product name of the FlexUtil Physical Drive. | [optional] 
**ProductRevision** | Pointer to **string** | Product revision of the FlexUtil Physical Drive. | [optional] 
**ReadErrorCount** | Pointer to **string** | Read error count of the FlexUtil Physical Drive. | [optional] 
**ReadErrorThreshold** | Pointer to **string** | Read error threshold for FlexUtil Physical Drive. | [optional] 
**WriteEnabled** | Pointer to **string** | Write access state of the FlexUtil Physical Drive. | [optional] 
**WriteErrorCount** | Pointer to **string** | Write error count of the FlexUtil Physical Drive. | [optional] 
**WriteErrorThreshold** | Pointer to **string** | Write error threshold for FlexUtil Physical Drive. | [optional] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](inventory.DeviceInfo.Relationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 
**StorageFlexUtilController** | Pointer to [**StorageFlexUtilControllerRelationship**](storage.FlexUtilController.Relationship.md) |  | [optional] 

## Methods

### NewStorageFlexUtilPhysicalDriveRelationship

`func NewStorageFlexUtilPhysicalDriveRelationship(classId string, objectType string, ) *StorageFlexUtilPhysicalDriveRelationship`

NewStorageFlexUtilPhysicalDriveRelationship instantiates a new StorageFlexUtilPhysicalDriveRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageFlexUtilPhysicalDriveRelationshipWithDefaults

`func NewStorageFlexUtilPhysicalDriveRelationshipWithDefaults() *StorageFlexUtilPhysicalDriveRelationship`

NewStorageFlexUtilPhysicalDriveRelationshipWithDefaults instantiates a new StorageFlexUtilPhysicalDriveRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountMoid

`func (o *StorageFlexUtilPhysicalDriveRelationship) GetAccountMoid() string`

GetAccountMoid returns the AccountMoid field if non-nil, zero value otherwise.

### GetAccountMoidOk

`func (o *StorageFlexUtilPhysicalDriveRelationship) GetAccountMoidOk() (*string, bool)`

GetAccountMoidOk returns a tuple with the AccountMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMoid

`func (o *StorageFlexUtilPhysicalDriveRelationship) SetAccountMoid(v string)`

SetAccountMoid sets AccountMoid field to given value.

### HasAccountMoid

`func (o *StorageFlexUtilPhysicalDriveRelationship) HasAccountMoid() bool`

HasAccountMoid returns a boolean if a field has been set.

### GetClassId

`func (o *StorageFlexUtilPhysicalDriveRelationship) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageFlexUtilPhysicalDriveRelationship) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageFlexUtilPhysicalDriveRelationship) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetCreateTime

`func (o *StorageFlexUtilPhysicalDriveRelationship) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *StorageFlexUtilPhysicalDriveRelationship) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *StorageFlexUtilPhysicalDriveRelationship) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *StorageFlexUtilPhysicalDriveRelationship) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDomainGroupMoid

`func (o *StorageFlexUtilPhysicalDriveRelationship) GetDomainGroupMoid() string`

GetDomainGroupMoid returns the DomainGroupMoid field if non-nil, zero value otherwise.

### GetDomainGroupMoidOk

`func (o *StorageFlexUtilPhysicalDriveRelationship) GetDomainGroupMoidOk() (*string, bool)`

GetDomainGroupMoidOk returns a tuple with the DomainGroupMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGroupMoid

`func (o *StorageFlexUtilPhysicalDriveRelationship) SetDomainGroupMoid(v string)`

SetDomainGroupMoid sets DomainGroupMoid field to given value.

### HasDomainGroupMoid

`func (o *StorageFlexUtilPhysicalDriveRelationship) HasDomainGroupMoid() bool`

HasDomainGroupMoid returns a boolean if a field has been set.

### GetModTime

`func (o *StorageFlexUtilPhysicalDriveRelationship) GetModTime() time.Time`

GetModTime returns the ModTime field if non-nil, zero value otherwise.

### GetModTimeOk

`func (o *StorageFlexUtilPhysicalDriveRelationship) GetModTimeOk() (*time.Time, bool)`

GetModTimeOk returns a tuple with the ModTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModTime

`func (o *StorageFlexUtilPhysicalDriveRelationship) SetModTime(v time.Time)`

SetModTime sets ModTime field to given value.

### HasModTime

`func (o *StorageFlexUtilPhysicalDriveRelationship) HasModTime() bool`

HasModTime returns a boolean if a field has been set.

### GetMoid

`func (o *StorageFlexUtilPhysicalDriveRelationship) GetMoid() string`

GetMoid returns the Moid field if non-nil, zero value otherwise.

### GetMoidOk

`func (o *StorageFlexUtilPhysicalDriveRelationship) GetMoidOk() (*string, bool)`

GetMoidOk returns a tuple with the Moid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoid

`func (o *StorageFlexUtilPhysicalDriveRelationship) SetMoid(v string)`

SetMoid sets Moid field to given value.

### HasMoid

`func (o *StorageFlexUtilPhysicalDriveRelationship) HasMoid() bool`

HasMoid returns a boolean if a field has been set.

### GetObjectType

`func (o *StorageFlexUtilPhysicalDriveRelationship) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageFlexUtilPhysicalDriveRelationship) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageFlexUtilPhysicalDriveRelationship) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOwners

`func (o *StorageFlexUtilPhysicalDriveRelationship) GetOwners() []string`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *StorageFlexUtilPhysicalDriveRelationship) GetOwnersOk() (*[]string, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *StorageFlexUtilPhysicalDriveRelationship) SetOwners(v []string)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *StorageFlexUtilPhysicalDriveRelationship) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetSharedScope

`func (o *StorageFlexUtilPhysicalDriveRelationship) GetSharedScope() string`

GetSharedScope returns the SharedScope field if non-nil, zero value otherwise.

### GetSharedScopeOk

`func (o *StorageFlexUtilPhysicalDriveRelationship) GetSharedScopeOk() (*string, bool)`

GetSharedScopeOk returns a tuple with the SharedScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedScope

`func (o *StorageFlexUtilPhysicalDriveRelationship) SetSharedScope(v string)`

SetSharedScope sets SharedScope field to given value.

### HasSharedScope

`func (o *StorageFlexUtilPhysicalDriveRelationship) HasSharedScope() bool`

HasSharedScope returns a boolean if a field has been set.

### GetTags

`func (o *StorageFlexUtilPhysicalDriveRelationship) GetTags() []MoTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *StorageFlexUtilPhysicalDriveRelationship) GetTagsOk() (*[]MoTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *StorageFlexUtilPhysicalDriveRelationship) SetTags(v []MoTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *StorageFlexUtilPhysicalDriveRelationship) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersionContext

`func (o *StorageFlexUtilPhysicalDriveRelationship) GetVersionContext() MoVersionContext`

GetVersionContext returns the VersionContext field if non-nil, zero value otherwise.

### GetVersionContextOk

`func (o *StorageFlexUtilPhysicalDriveRelationship) GetVersionContextOk() (*MoVersionContext, bool)`

GetVersionContextOk returns a tuple with the VersionContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionContext

`func (o *StorageFlexUtilPhysicalDriveRelationship) SetVersionContext(v MoVersionContext)`

SetVersionContext sets VersionContext field to given value.

### HasVersionContext

`func (o *StorageFlexUtilPhysicalDriveRelationship) HasVersionContext() bool`

HasVersionContext returns a boolean if a field has been set.

### GetAncestors

`func (o *StorageFlexUtilPhysicalDriveRelationship) GetAncestors() []MoBaseMoRelationship`

GetAncestors returns the Ancestors field if non-nil, zero value otherwise.

### GetAncestorsOk

`func (o *StorageFlexUtilPhysicalDriveRelationship) GetAncestorsOk() (*[]MoBaseMoRelationship, bool)`

GetAncestorsOk returns a tuple with the Ancestors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestors

`func (o *StorageFlexUtilPhysicalDriveRelationship) SetAncestors(v []MoBaseMoRelationship)`

SetAncestors sets Ancestors field to given value.

### HasAncestors

`func (o *StorageFlexUtilPhysicalDriveRelationship) HasAncestors() bool`

HasAncestors returns a boolean if a field has been set.

### SetAncestorsNil

`func (o *StorageFlexUtilPhysicalDriveRelationship) SetAncestorsNil(b bool)`

 SetAncestorsNil sets the value for Ancestors to be an explicit nil

### UnsetAncestors
`func (o *StorageFlexUtilPhysicalDriveRelationship) UnsetAncestors()`

UnsetAncestors ensures that no value is present for Ancestors, not even an explicit nil
### GetParent

`func (o *StorageFlexUtilPhysicalDriveRelationship) GetParent() MoBaseMoRelationship`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *StorageFlexUtilPhysicalDriveRelationship) GetParentOk() (*MoBaseMoRelationship, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *StorageFlexUtilPhysicalDriveRelationship) SetParent(v MoBaseMoRelationship)`

SetParent sets Parent field to given value.

### HasParent

`func (o *StorageFlexUtilPhysicalDriveRelationship) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPermissionResources

`func (o *StorageFlexUtilPhysicalDriveRelationship) GetPermissionResources() []MoBaseMoRelationship`

GetPermissionResources returns the PermissionResources field if non-nil, zero value otherwise.

### GetPermissionResourcesOk

`func (o *StorageFlexUtilPhysicalDriveRelationship) GetPermissionResourcesOk() (*[]MoBaseMoRelationship, bool)`

GetPermissionResourcesOk returns a tuple with the PermissionResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionResources

`func (o *StorageFlexUtilPhysicalDriveRelationship) SetPermissionResources(v []MoBaseMoRelationship)`

SetPermissionResources sets PermissionResources field to given value.

### HasPermissionResources

`func (o *StorageFlexUtilPhysicalDriveRelationship) HasPermissionResources() bool`

HasPermissionResources returns a boolean if a field has been set.

### SetPermissionResourcesNil

`func (o *StorageFlexUtilPhysicalDriveRelationship) SetPermissionResourcesNil(b bool)`

 SetPermissionResourcesNil sets the value for PermissionResources to be an explicit nil

### UnsetPermissionResources
`func (o *StorageFlexUtilPhysicalDriveRelationship) UnsetPermissionResources()`

UnsetPermissionResources ensures that no value is present for PermissionResources, not even an explicit nil
### GetDisplayNames

`func (o *StorageFlexUtilPhysicalDriveRelationship) GetDisplayNames() map[string][]string`

GetDisplayNames returns the DisplayNames field if non-nil, zero value otherwise.

### GetDisplayNamesOk

`func (o *StorageFlexUtilPhysicalDriveRelationship) GetDisplayNamesOk() (*map[string][]string, bool)`

GetDisplayNamesOk returns a tuple with the DisplayNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNames

`func (o *StorageFlexUtilPhysicalDriveRelationship) SetDisplayNames(v map[string][]string)`

SetDisplayNames sets DisplayNames field to given value.

### HasDisplayNames

`func (o *StorageFlexUtilPhysicalDriveRelationship) HasDisplayNames() bool`

HasDisplayNames returns a boolean if a field has been set.

### SetDisplayNamesNil

`func (o *StorageFlexUtilPhysicalDriveRelationship) SetDisplayNamesNil(b bool)`

 SetDisplayNamesNil sets the value for DisplayNames to be an explicit nil

### UnsetDisplayNames
`func (o *StorageFlexUtilPhysicalDriveRelationship) UnsetDisplayNames()`

UnsetDisplayNames ensures that no value is present for DisplayNames, not even an explicit nil
### GetDeviceMoId

`func (o *StorageFlexUtilPhysicalDriveRelationship) GetDeviceMoId() string`

GetDeviceMoId returns the DeviceMoId field if non-nil, zero value otherwise.

### GetDeviceMoIdOk

`func (o *StorageFlexUtilPhysicalDriveRelationship) GetDeviceMoIdOk() (*string, bool)`

GetDeviceMoIdOk returns a tuple with the DeviceMoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMoId

`func (o *StorageFlexUtilPhysicalDriveRelationship) SetDeviceMoId(v string)`

SetDeviceMoId sets DeviceMoId field to given value.

### HasDeviceMoId

`func (o *StorageFlexUtilPhysicalDriveRelationship) HasDeviceMoId() bool`

HasDeviceMoId returns a boolean if a field has been set.

### GetDn

`func (o *StorageFlexUtilPhysicalDriveRelationship) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *StorageFlexUtilPhysicalDriveRelationship) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *StorageFlexUtilPhysicalDriveRelationship) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *StorageFlexUtilPhysicalDriveRelationship) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetRn

`func (o *StorageFlexUtilPhysicalDriveRelationship) GetRn() string`

GetRn returns the Rn field if non-nil, zero value otherwise.

### GetRnOk

`func (o *StorageFlexUtilPhysicalDriveRelationship) GetRnOk() (*string, bool)`

GetRnOk returns a tuple with the Rn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRn

`func (o *StorageFlexUtilPhysicalDriveRelationship) SetRn(v string)`

SetRn sets Rn field to given value.

### HasRn

`func (o *StorageFlexUtilPhysicalDriveRelationship) HasRn() bool`

HasRn returns a boolean if a field has been set.

### GetModel

`func (o *StorageFlexUtilPhysicalDriveRelationship) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *StorageFlexUtilPhysicalDriveRelationship) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *StorageFlexUtilPhysicalDriveRelationship) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *StorageFlexUtilPhysicalDriveRelationship) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetRevision

`func (o *StorageFlexUtilPhysicalDriveRelationship) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *StorageFlexUtilPhysicalDriveRelationship) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *StorageFlexUtilPhysicalDriveRelationship) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *StorageFlexUtilPhysicalDriveRelationship) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetSerial

`func (o *StorageFlexUtilPhysicalDriveRelationship) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *StorageFlexUtilPhysicalDriveRelationship) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *StorageFlexUtilPhysicalDriveRelationship) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *StorageFlexUtilPhysicalDriveRelationship) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetVendor

`func (o *StorageFlexUtilPhysicalDriveRelationship) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *StorageFlexUtilPhysicalDriveRelationship) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *StorageFlexUtilPhysicalDriveRelationship) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *StorageFlexUtilPhysicalDriveRelationship) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetBlockSize

`func (o *StorageFlexUtilPhysicalDriveRelationship) GetBlockSize() string`

GetBlockSize returns the BlockSize field if non-nil, zero value otherwise.

### GetBlockSizeOk

`func (o *StorageFlexUtilPhysicalDriveRelationship) GetBlockSizeOk() (*string, bool)`

GetBlockSizeOk returns a tuple with the BlockSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockSize

`func (o *StorageFlexUtilPhysicalDriveRelationship) SetBlockSize(v string)`

SetBlockSize sets BlockSize field to given value.

### HasBlockSize

`func (o *StorageFlexUtilPhysicalDriveRelationship) HasBlockSize() bool`

HasBlockSize returns a boolean if a field has been set.

### GetCapacity

`func (o *StorageFlexUtilPhysicalDriveRelationship) GetCapacity() string`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *StorageFlexUtilPhysicalDriveRelationship) GetCapacityOk() (*string, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *StorageFlexUtilPhysicalDriveRelationship) SetCapacity(v string)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *StorageFlexUtilPhysicalDriveRelationship) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### GetController

`func (o *StorageFlexUtilPhysicalDriveRelationship) GetController() string`

GetController returns the Controller field if non-nil, zero value otherwise.

### GetControllerOk

`func (o *StorageFlexUtilPhysicalDriveRelationship) GetControllerOk() (*string, bool)`

GetControllerOk returns a tuple with the Controller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetController

`func (o *StorageFlexUtilPhysicalDriveRelationship) SetController(v string)`

SetController sets Controller field to given value.

### HasController

`func (o *StorageFlexUtilPhysicalDriveRelationship) HasController() bool`

HasController returns a boolean if a field has been set.

### GetDrivesEnabled

`func (o *StorageFlexUtilPhysicalDriveRelationship) GetDrivesEnabled() string`

GetDrivesEnabled returns the DrivesEnabled field if non-nil, zero value otherwise.

### GetDrivesEnabledOk

`func (o *StorageFlexUtilPhysicalDriveRelationship) GetDrivesEnabledOk() (*string, bool)`

GetDrivesEnabledOk returns a tuple with the DrivesEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrivesEnabled

`func (o *StorageFlexUtilPhysicalDriveRelationship) SetDrivesEnabled(v string)`

SetDrivesEnabled sets DrivesEnabled field to given value.

### HasDrivesEnabled

`func (o *StorageFlexUtilPhysicalDriveRelationship) HasDrivesEnabled() bool`

HasDrivesEnabled returns a boolean if a field has been set.

### GetHealth

`func (o *StorageFlexUtilPhysicalDriveRelationship) GetHealth() string`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *StorageFlexUtilPhysicalDriveRelationship) GetHealthOk() (*string, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *StorageFlexUtilPhysicalDriveRelationship) SetHealth(v string)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *StorageFlexUtilPhysicalDriveRelationship) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetManufacturerDate

`func (o *StorageFlexUtilPhysicalDriveRelationship) GetManufacturerDate() string`

GetManufacturerDate returns the ManufacturerDate field if non-nil, zero value otherwise.

### GetManufacturerDateOk

`func (o *StorageFlexUtilPhysicalDriveRelationship) GetManufacturerDateOk() (*string, bool)`

GetManufacturerDateOk returns a tuple with the ManufacturerDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturerDate

`func (o *StorageFlexUtilPhysicalDriveRelationship) SetManufacturerDate(v string)`

SetManufacturerDate sets ManufacturerDate field to given value.

### HasManufacturerDate

`func (o *StorageFlexUtilPhysicalDriveRelationship) HasManufacturerDate() bool`

HasManufacturerDate returns a boolean if a field has been set.

### GetManufacturerId

`func (o *StorageFlexUtilPhysicalDriveRelationship) GetManufacturerId() string`

GetManufacturerId returns the ManufacturerId field if non-nil, zero value otherwise.

### GetManufacturerIdOk

`func (o *StorageFlexUtilPhysicalDriveRelationship) GetManufacturerIdOk() (*string, bool)`

GetManufacturerIdOk returns a tuple with the ManufacturerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturerId

`func (o *StorageFlexUtilPhysicalDriveRelationship) SetManufacturerId(v string)`

SetManufacturerId sets ManufacturerId field to given value.

### HasManufacturerId

`func (o *StorageFlexUtilPhysicalDriveRelationship) HasManufacturerId() bool`

HasManufacturerId returns a boolean if a field has been set.

### GetOemId

`func (o *StorageFlexUtilPhysicalDriveRelationship) GetOemId() string`

GetOemId returns the OemId field if non-nil, zero value otherwise.

### GetOemIdOk

`func (o *StorageFlexUtilPhysicalDriveRelationship) GetOemIdOk() (*string, bool)`

GetOemIdOk returns a tuple with the OemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOemId

`func (o *StorageFlexUtilPhysicalDriveRelationship) SetOemId(v string)`

SetOemId sets OemId field to given value.

### HasOemId

`func (o *StorageFlexUtilPhysicalDriveRelationship) HasOemId() bool`

HasOemId returns a boolean if a field has been set.

### GetPartitionCount

`func (o *StorageFlexUtilPhysicalDriveRelationship) GetPartitionCount() string`

GetPartitionCount returns the PartitionCount field if non-nil, zero value otherwise.

### GetPartitionCountOk

`func (o *StorageFlexUtilPhysicalDriveRelationship) GetPartitionCountOk() (*string, bool)`

GetPartitionCountOk returns a tuple with the PartitionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionCount

`func (o *StorageFlexUtilPhysicalDriveRelationship) SetPartitionCount(v string)`

SetPartitionCount sets PartitionCount field to given value.

### HasPartitionCount

`func (o *StorageFlexUtilPhysicalDriveRelationship) HasPartitionCount() bool`

HasPartitionCount returns a boolean if a field has been set.

### GetPdStatus

`func (o *StorageFlexUtilPhysicalDriveRelationship) GetPdStatus() string`

GetPdStatus returns the PdStatus field if non-nil, zero value otherwise.

### GetPdStatusOk

`func (o *StorageFlexUtilPhysicalDriveRelationship) GetPdStatusOk() (*string, bool)`

GetPdStatusOk returns a tuple with the PdStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdStatus

`func (o *StorageFlexUtilPhysicalDriveRelationship) SetPdStatus(v string)`

SetPdStatus sets PdStatus field to given value.

### HasPdStatus

`func (o *StorageFlexUtilPhysicalDriveRelationship) HasPdStatus() bool`

HasPdStatus returns a boolean if a field has been set.

### GetPhysicalDrive

`func (o *StorageFlexUtilPhysicalDriveRelationship) GetPhysicalDrive() string`

GetPhysicalDrive returns the PhysicalDrive field if non-nil, zero value otherwise.

### GetPhysicalDriveOk

`func (o *StorageFlexUtilPhysicalDriveRelationship) GetPhysicalDriveOk() (*string, bool)`

GetPhysicalDriveOk returns a tuple with the PhysicalDrive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalDrive

`func (o *StorageFlexUtilPhysicalDriveRelationship) SetPhysicalDrive(v string)`

SetPhysicalDrive sets PhysicalDrive field to given value.

### HasPhysicalDrive

`func (o *StorageFlexUtilPhysicalDriveRelationship) HasPhysicalDrive() bool`

HasPhysicalDrive returns a boolean if a field has been set.

### GetProductName

`func (o *StorageFlexUtilPhysicalDriveRelationship) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *StorageFlexUtilPhysicalDriveRelationship) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *StorageFlexUtilPhysicalDriveRelationship) SetProductName(v string)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *StorageFlexUtilPhysicalDriveRelationship) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### GetProductRevision

`func (o *StorageFlexUtilPhysicalDriveRelationship) GetProductRevision() string`

GetProductRevision returns the ProductRevision field if non-nil, zero value otherwise.

### GetProductRevisionOk

`func (o *StorageFlexUtilPhysicalDriveRelationship) GetProductRevisionOk() (*string, bool)`

GetProductRevisionOk returns a tuple with the ProductRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductRevision

`func (o *StorageFlexUtilPhysicalDriveRelationship) SetProductRevision(v string)`

SetProductRevision sets ProductRevision field to given value.

### HasProductRevision

`func (o *StorageFlexUtilPhysicalDriveRelationship) HasProductRevision() bool`

HasProductRevision returns a boolean if a field has been set.

### GetReadErrorCount

`func (o *StorageFlexUtilPhysicalDriveRelationship) GetReadErrorCount() string`

GetReadErrorCount returns the ReadErrorCount field if non-nil, zero value otherwise.

### GetReadErrorCountOk

`func (o *StorageFlexUtilPhysicalDriveRelationship) GetReadErrorCountOk() (*string, bool)`

GetReadErrorCountOk returns a tuple with the ReadErrorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadErrorCount

`func (o *StorageFlexUtilPhysicalDriveRelationship) SetReadErrorCount(v string)`

SetReadErrorCount sets ReadErrorCount field to given value.

### HasReadErrorCount

`func (o *StorageFlexUtilPhysicalDriveRelationship) HasReadErrorCount() bool`

HasReadErrorCount returns a boolean if a field has been set.

### GetReadErrorThreshold

`func (o *StorageFlexUtilPhysicalDriveRelationship) GetReadErrorThreshold() string`

GetReadErrorThreshold returns the ReadErrorThreshold field if non-nil, zero value otherwise.

### GetReadErrorThresholdOk

`func (o *StorageFlexUtilPhysicalDriveRelationship) GetReadErrorThresholdOk() (*string, bool)`

GetReadErrorThresholdOk returns a tuple with the ReadErrorThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadErrorThreshold

`func (o *StorageFlexUtilPhysicalDriveRelationship) SetReadErrorThreshold(v string)`

SetReadErrorThreshold sets ReadErrorThreshold field to given value.

### HasReadErrorThreshold

`func (o *StorageFlexUtilPhysicalDriveRelationship) HasReadErrorThreshold() bool`

HasReadErrorThreshold returns a boolean if a field has been set.

### GetWriteEnabled

`func (o *StorageFlexUtilPhysicalDriveRelationship) GetWriteEnabled() string`

GetWriteEnabled returns the WriteEnabled field if non-nil, zero value otherwise.

### GetWriteEnabledOk

`func (o *StorageFlexUtilPhysicalDriveRelationship) GetWriteEnabledOk() (*string, bool)`

GetWriteEnabledOk returns a tuple with the WriteEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteEnabled

`func (o *StorageFlexUtilPhysicalDriveRelationship) SetWriteEnabled(v string)`

SetWriteEnabled sets WriteEnabled field to given value.

### HasWriteEnabled

`func (o *StorageFlexUtilPhysicalDriveRelationship) HasWriteEnabled() bool`

HasWriteEnabled returns a boolean if a field has been set.

### GetWriteErrorCount

`func (o *StorageFlexUtilPhysicalDriveRelationship) GetWriteErrorCount() string`

GetWriteErrorCount returns the WriteErrorCount field if non-nil, zero value otherwise.

### GetWriteErrorCountOk

`func (o *StorageFlexUtilPhysicalDriveRelationship) GetWriteErrorCountOk() (*string, bool)`

GetWriteErrorCountOk returns a tuple with the WriteErrorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteErrorCount

`func (o *StorageFlexUtilPhysicalDriveRelationship) SetWriteErrorCount(v string)`

SetWriteErrorCount sets WriteErrorCount field to given value.

### HasWriteErrorCount

`func (o *StorageFlexUtilPhysicalDriveRelationship) HasWriteErrorCount() bool`

HasWriteErrorCount returns a boolean if a field has been set.

### GetWriteErrorThreshold

`func (o *StorageFlexUtilPhysicalDriveRelationship) GetWriteErrorThreshold() string`

GetWriteErrorThreshold returns the WriteErrorThreshold field if non-nil, zero value otherwise.

### GetWriteErrorThresholdOk

`func (o *StorageFlexUtilPhysicalDriveRelationship) GetWriteErrorThresholdOk() (*string, bool)`

GetWriteErrorThresholdOk returns a tuple with the WriteErrorThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteErrorThreshold

`func (o *StorageFlexUtilPhysicalDriveRelationship) SetWriteErrorThreshold(v string)`

SetWriteErrorThreshold sets WriteErrorThreshold field to given value.

### HasWriteErrorThreshold

`func (o *StorageFlexUtilPhysicalDriveRelationship) HasWriteErrorThreshold() bool`

HasWriteErrorThreshold returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *StorageFlexUtilPhysicalDriveRelationship) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *StorageFlexUtilPhysicalDriveRelationship) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *StorageFlexUtilPhysicalDriveRelationship) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *StorageFlexUtilPhysicalDriveRelationship) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *StorageFlexUtilPhysicalDriveRelationship) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StorageFlexUtilPhysicalDriveRelationship) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StorageFlexUtilPhysicalDriveRelationship) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StorageFlexUtilPhysicalDriveRelationship) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### GetStorageFlexUtilController

`func (o *StorageFlexUtilPhysicalDriveRelationship) GetStorageFlexUtilController() StorageFlexUtilControllerRelationship`

GetStorageFlexUtilController returns the StorageFlexUtilController field if non-nil, zero value otherwise.

### GetStorageFlexUtilControllerOk

`func (o *StorageFlexUtilPhysicalDriveRelationship) GetStorageFlexUtilControllerOk() (*StorageFlexUtilControllerRelationship, bool)`

GetStorageFlexUtilControllerOk returns a tuple with the StorageFlexUtilController field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageFlexUtilController

`func (o *StorageFlexUtilPhysicalDriveRelationship) SetStorageFlexUtilController(v StorageFlexUtilControllerRelationship)`

SetStorageFlexUtilController sets StorageFlexUtilController field to given value.

### HasStorageFlexUtilController

`func (o *StorageFlexUtilPhysicalDriveRelationship) HasStorageFlexUtilController() bool`

HasStorageFlexUtilController returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



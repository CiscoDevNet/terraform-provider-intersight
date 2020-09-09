# StorageFlexFlashVirtualDriveRelationship

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
**DriveScope** | Pointer to **string** | The drive scope of the flex flash virtual drive. | [optional] 
**DriveStatus** | Pointer to **string** | Status of virtual drive on the flex controller. | [optional] 
**PartitionId** | Pointer to **string** | The partition Id of the flex flash virtual Drive. | [optional] 
**ResidentImage** | Pointer to **string** | The resident image on the flex flash virtual Drive. | [optional] 
**Size** | Pointer to **string** | Size of virtual drive on the flex controller. | [optional] 
**VirtualDrive** | Pointer to **string** | Virtual drive on the flex flash controller. | [optional] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](inventory.DeviceInfo.Relationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 
**StorageFlexFlashController** | Pointer to [**StorageFlexFlashControllerRelationship**](storage.FlexFlashController.Relationship.md) |  | [optional] 

## Methods

### NewStorageFlexFlashVirtualDriveRelationship

`func NewStorageFlexFlashVirtualDriveRelationship(classId string, objectType string, ) *StorageFlexFlashVirtualDriveRelationship`

NewStorageFlexFlashVirtualDriveRelationship instantiates a new StorageFlexFlashVirtualDriveRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageFlexFlashVirtualDriveRelationshipWithDefaults

`func NewStorageFlexFlashVirtualDriveRelationshipWithDefaults() *StorageFlexFlashVirtualDriveRelationship`

NewStorageFlexFlashVirtualDriveRelationshipWithDefaults instantiates a new StorageFlexFlashVirtualDriveRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountMoid

`func (o *StorageFlexFlashVirtualDriveRelationship) GetAccountMoid() string`

GetAccountMoid returns the AccountMoid field if non-nil, zero value otherwise.

### GetAccountMoidOk

`func (o *StorageFlexFlashVirtualDriveRelationship) GetAccountMoidOk() (*string, bool)`

GetAccountMoidOk returns a tuple with the AccountMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMoid

`func (o *StorageFlexFlashVirtualDriveRelationship) SetAccountMoid(v string)`

SetAccountMoid sets AccountMoid field to given value.

### HasAccountMoid

`func (o *StorageFlexFlashVirtualDriveRelationship) HasAccountMoid() bool`

HasAccountMoid returns a boolean if a field has been set.

### GetClassId

`func (o *StorageFlexFlashVirtualDriveRelationship) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageFlexFlashVirtualDriveRelationship) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageFlexFlashVirtualDriveRelationship) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetCreateTime

`func (o *StorageFlexFlashVirtualDriveRelationship) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *StorageFlexFlashVirtualDriveRelationship) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *StorageFlexFlashVirtualDriveRelationship) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *StorageFlexFlashVirtualDriveRelationship) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDomainGroupMoid

`func (o *StorageFlexFlashVirtualDriveRelationship) GetDomainGroupMoid() string`

GetDomainGroupMoid returns the DomainGroupMoid field if non-nil, zero value otherwise.

### GetDomainGroupMoidOk

`func (o *StorageFlexFlashVirtualDriveRelationship) GetDomainGroupMoidOk() (*string, bool)`

GetDomainGroupMoidOk returns a tuple with the DomainGroupMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGroupMoid

`func (o *StorageFlexFlashVirtualDriveRelationship) SetDomainGroupMoid(v string)`

SetDomainGroupMoid sets DomainGroupMoid field to given value.

### HasDomainGroupMoid

`func (o *StorageFlexFlashVirtualDriveRelationship) HasDomainGroupMoid() bool`

HasDomainGroupMoid returns a boolean if a field has been set.

### GetModTime

`func (o *StorageFlexFlashVirtualDriveRelationship) GetModTime() time.Time`

GetModTime returns the ModTime field if non-nil, zero value otherwise.

### GetModTimeOk

`func (o *StorageFlexFlashVirtualDriveRelationship) GetModTimeOk() (*time.Time, bool)`

GetModTimeOk returns a tuple with the ModTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModTime

`func (o *StorageFlexFlashVirtualDriveRelationship) SetModTime(v time.Time)`

SetModTime sets ModTime field to given value.

### HasModTime

`func (o *StorageFlexFlashVirtualDriveRelationship) HasModTime() bool`

HasModTime returns a boolean if a field has been set.

### GetMoid

`func (o *StorageFlexFlashVirtualDriveRelationship) GetMoid() string`

GetMoid returns the Moid field if non-nil, zero value otherwise.

### GetMoidOk

`func (o *StorageFlexFlashVirtualDriveRelationship) GetMoidOk() (*string, bool)`

GetMoidOk returns a tuple with the Moid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoid

`func (o *StorageFlexFlashVirtualDriveRelationship) SetMoid(v string)`

SetMoid sets Moid field to given value.

### HasMoid

`func (o *StorageFlexFlashVirtualDriveRelationship) HasMoid() bool`

HasMoid returns a boolean if a field has been set.

### GetObjectType

`func (o *StorageFlexFlashVirtualDriveRelationship) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageFlexFlashVirtualDriveRelationship) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageFlexFlashVirtualDriveRelationship) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOwners

`func (o *StorageFlexFlashVirtualDriveRelationship) GetOwners() []string`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *StorageFlexFlashVirtualDriveRelationship) GetOwnersOk() (*[]string, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *StorageFlexFlashVirtualDriveRelationship) SetOwners(v []string)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *StorageFlexFlashVirtualDriveRelationship) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetSharedScope

`func (o *StorageFlexFlashVirtualDriveRelationship) GetSharedScope() string`

GetSharedScope returns the SharedScope field if non-nil, zero value otherwise.

### GetSharedScopeOk

`func (o *StorageFlexFlashVirtualDriveRelationship) GetSharedScopeOk() (*string, bool)`

GetSharedScopeOk returns a tuple with the SharedScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedScope

`func (o *StorageFlexFlashVirtualDriveRelationship) SetSharedScope(v string)`

SetSharedScope sets SharedScope field to given value.

### HasSharedScope

`func (o *StorageFlexFlashVirtualDriveRelationship) HasSharedScope() bool`

HasSharedScope returns a boolean if a field has been set.

### GetTags

`func (o *StorageFlexFlashVirtualDriveRelationship) GetTags() []MoTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *StorageFlexFlashVirtualDriveRelationship) GetTagsOk() (*[]MoTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *StorageFlexFlashVirtualDriveRelationship) SetTags(v []MoTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *StorageFlexFlashVirtualDriveRelationship) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersionContext

`func (o *StorageFlexFlashVirtualDriveRelationship) GetVersionContext() MoVersionContext`

GetVersionContext returns the VersionContext field if non-nil, zero value otherwise.

### GetVersionContextOk

`func (o *StorageFlexFlashVirtualDriveRelationship) GetVersionContextOk() (*MoVersionContext, bool)`

GetVersionContextOk returns a tuple with the VersionContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionContext

`func (o *StorageFlexFlashVirtualDriveRelationship) SetVersionContext(v MoVersionContext)`

SetVersionContext sets VersionContext field to given value.

### HasVersionContext

`func (o *StorageFlexFlashVirtualDriveRelationship) HasVersionContext() bool`

HasVersionContext returns a boolean if a field has been set.

### GetAncestors

`func (o *StorageFlexFlashVirtualDriveRelationship) GetAncestors() []MoBaseMoRelationship`

GetAncestors returns the Ancestors field if non-nil, zero value otherwise.

### GetAncestorsOk

`func (o *StorageFlexFlashVirtualDriveRelationship) GetAncestorsOk() (*[]MoBaseMoRelationship, bool)`

GetAncestorsOk returns a tuple with the Ancestors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestors

`func (o *StorageFlexFlashVirtualDriveRelationship) SetAncestors(v []MoBaseMoRelationship)`

SetAncestors sets Ancestors field to given value.

### HasAncestors

`func (o *StorageFlexFlashVirtualDriveRelationship) HasAncestors() bool`

HasAncestors returns a boolean if a field has been set.

### SetAncestorsNil

`func (o *StorageFlexFlashVirtualDriveRelationship) SetAncestorsNil(b bool)`

 SetAncestorsNil sets the value for Ancestors to be an explicit nil

### UnsetAncestors
`func (o *StorageFlexFlashVirtualDriveRelationship) UnsetAncestors()`

UnsetAncestors ensures that no value is present for Ancestors, not even an explicit nil
### GetParent

`func (o *StorageFlexFlashVirtualDriveRelationship) GetParent() MoBaseMoRelationship`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *StorageFlexFlashVirtualDriveRelationship) GetParentOk() (*MoBaseMoRelationship, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *StorageFlexFlashVirtualDriveRelationship) SetParent(v MoBaseMoRelationship)`

SetParent sets Parent field to given value.

### HasParent

`func (o *StorageFlexFlashVirtualDriveRelationship) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPermissionResources

`func (o *StorageFlexFlashVirtualDriveRelationship) GetPermissionResources() []MoBaseMoRelationship`

GetPermissionResources returns the PermissionResources field if non-nil, zero value otherwise.

### GetPermissionResourcesOk

`func (o *StorageFlexFlashVirtualDriveRelationship) GetPermissionResourcesOk() (*[]MoBaseMoRelationship, bool)`

GetPermissionResourcesOk returns a tuple with the PermissionResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionResources

`func (o *StorageFlexFlashVirtualDriveRelationship) SetPermissionResources(v []MoBaseMoRelationship)`

SetPermissionResources sets PermissionResources field to given value.

### HasPermissionResources

`func (o *StorageFlexFlashVirtualDriveRelationship) HasPermissionResources() bool`

HasPermissionResources returns a boolean if a field has been set.

### SetPermissionResourcesNil

`func (o *StorageFlexFlashVirtualDriveRelationship) SetPermissionResourcesNil(b bool)`

 SetPermissionResourcesNil sets the value for PermissionResources to be an explicit nil

### UnsetPermissionResources
`func (o *StorageFlexFlashVirtualDriveRelationship) UnsetPermissionResources()`

UnsetPermissionResources ensures that no value is present for PermissionResources, not even an explicit nil
### GetDisplayNames

`func (o *StorageFlexFlashVirtualDriveRelationship) GetDisplayNames() map[string][]string`

GetDisplayNames returns the DisplayNames field if non-nil, zero value otherwise.

### GetDisplayNamesOk

`func (o *StorageFlexFlashVirtualDriveRelationship) GetDisplayNamesOk() (*map[string][]string, bool)`

GetDisplayNamesOk returns a tuple with the DisplayNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNames

`func (o *StorageFlexFlashVirtualDriveRelationship) SetDisplayNames(v map[string][]string)`

SetDisplayNames sets DisplayNames field to given value.

### HasDisplayNames

`func (o *StorageFlexFlashVirtualDriveRelationship) HasDisplayNames() bool`

HasDisplayNames returns a boolean if a field has been set.

### SetDisplayNamesNil

`func (o *StorageFlexFlashVirtualDriveRelationship) SetDisplayNamesNil(b bool)`

 SetDisplayNamesNil sets the value for DisplayNames to be an explicit nil

### UnsetDisplayNames
`func (o *StorageFlexFlashVirtualDriveRelationship) UnsetDisplayNames()`

UnsetDisplayNames ensures that no value is present for DisplayNames, not even an explicit nil
### GetDeviceMoId

`func (o *StorageFlexFlashVirtualDriveRelationship) GetDeviceMoId() string`

GetDeviceMoId returns the DeviceMoId field if non-nil, zero value otherwise.

### GetDeviceMoIdOk

`func (o *StorageFlexFlashVirtualDriveRelationship) GetDeviceMoIdOk() (*string, bool)`

GetDeviceMoIdOk returns a tuple with the DeviceMoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMoId

`func (o *StorageFlexFlashVirtualDriveRelationship) SetDeviceMoId(v string)`

SetDeviceMoId sets DeviceMoId field to given value.

### HasDeviceMoId

`func (o *StorageFlexFlashVirtualDriveRelationship) HasDeviceMoId() bool`

HasDeviceMoId returns a boolean if a field has been set.

### GetDn

`func (o *StorageFlexFlashVirtualDriveRelationship) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *StorageFlexFlashVirtualDriveRelationship) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *StorageFlexFlashVirtualDriveRelationship) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *StorageFlexFlashVirtualDriveRelationship) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetRn

`func (o *StorageFlexFlashVirtualDriveRelationship) GetRn() string`

GetRn returns the Rn field if non-nil, zero value otherwise.

### GetRnOk

`func (o *StorageFlexFlashVirtualDriveRelationship) GetRnOk() (*string, bool)`

GetRnOk returns a tuple with the Rn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRn

`func (o *StorageFlexFlashVirtualDriveRelationship) SetRn(v string)`

SetRn sets Rn field to given value.

### HasRn

`func (o *StorageFlexFlashVirtualDriveRelationship) HasRn() bool`

HasRn returns a boolean if a field has been set.

### GetModel

`func (o *StorageFlexFlashVirtualDriveRelationship) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *StorageFlexFlashVirtualDriveRelationship) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *StorageFlexFlashVirtualDriveRelationship) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *StorageFlexFlashVirtualDriveRelationship) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetRevision

`func (o *StorageFlexFlashVirtualDriveRelationship) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *StorageFlexFlashVirtualDriveRelationship) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *StorageFlexFlashVirtualDriveRelationship) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *StorageFlexFlashVirtualDriveRelationship) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetSerial

`func (o *StorageFlexFlashVirtualDriveRelationship) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *StorageFlexFlashVirtualDriveRelationship) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *StorageFlexFlashVirtualDriveRelationship) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *StorageFlexFlashVirtualDriveRelationship) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetVendor

`func (o *StorageFlexFlashVirtualDriveRelationship) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *StorageFlexFlashVirtualDriveRelationship) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *StorageFlexFlashVirtualDriveRelationship) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *StorageFlexFlashVirtualDriveRelationship) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetDriveScope

`func (o *StorageFlexFlashVirtualDriveRelationship) GetDriveScope() string`

GetDriveScope returns the DriveScope field if non-nil, zero value otherwise.

### GetDriveScopeOk

`func (o *StorageFlexFlashVirtualDriveRelationship) GetDriveScopeOk() (*string, bool)`

GetDriveScopeOk returns a tuple with the DriveScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriveScope

`func (o *StorageFlexFlashVirtualDriveRelationship) SetDriveScope(v string)`

SetDriveScope sets DriveScope field to given value.

### HasDriveScope

`func (o *StorageFlexFlashVirtualDriveRelationship) HasDriveScope() bool`

HasDriveScope returns a boolean if a field has been set.

### GetDriveStatus

`func (o *StorageFlexFlashVirtualDriveRelationship) GetDriveStatus() string`

GetDriveStatus returns the DriveStatus field if non-nil, zero value otherwise.

### GetDriveStatusOk

`func (o *StorageFlexFlashVirtualDriveRelationship) GetDriveStatusOk() (*string, bool)`

GetDriveStatusOk returns a tuple with the DriveStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriveStatus

`func (o *StorageFlexFlashVirtualDriveRelationship) SetDriveStatus(v string)`

SetDriveStatus sets DriveStatus field to given value.

### HasDriveStatus

`func (o *StorageFlexFlashVirtualDriveRelationship) HasDriveStatus() bool`

HasDriveStatus returns a boolean if a field has been set.

### GetPartitionId

`func (o *StorageFlexFlashVirtualDriveRelationship) GetPartitionId() string`

GetPartitionId returns the PartitionId field if non-nil, zero value otherwise.

### GetPartitionIdOk

`func (o *StorageFlexFlashVirtualDriveRelationship) GetPartitionIdOk() (*string, bool)`

GetPartitionIdOk returns a tuple with the PartitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionId

`func (o *StorageFlexFlashVirtualDriveRelationship) SetPartitionId(v string)`

SetPartitionId sets PartitionId field to given value.

### HasPartitionId

`func (o *StorageFlexFlashVirtualDriveRelationship) HasPartitionId() bool`

HasPartitionId returns a boolean if a field has been set.

### GetResidentImage

`func (o *StorageFlexFlashVirtualDriveRelationship) GetResidentImage() string`

GetResidentImage returns the ResidentImage field if non-nil, zero value otherwise.

### GetResidentImageOk

`func (o *StorageFlexFlashVirtualDriveRelationship) GetResidentImageOk() (*string, bool)`

GetResidentImageOk returns a tuple with the ResidentImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResidentImage

`func (o *StorageFlexFlashVirtualDriveRelationship) SetResidentImage(v string)`

SetResidentImage sets ResidentImage field to given value.

### HasResidentImage

`func (o *StorageFlexFlashVirtualDriveRelationship) HasResidentImage() bool`

HasResidentImage returns a boolean if a field has been set.

### GetSize

`func (o *StorageFlexFlashVirtualDriveRelationship) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *StorageFlexFlashVirtualDriveRelationship) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *StorageFlexFlashVirtualDriveRelationship) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *StorageFlexFlashVirtualDriveRelationship) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetVirtualDrive

`func (o *StorageFlexFlashVirtualDriveRelationship) GetVirtualDrive() string`

GetVirtualDrive returns the VirtualDrive field if non-nil, zero value otherwise.

### GetVirtualDriveOk

`func (o *StorageFlexFlashVirtualDriveRelationship) GetVirtualDriveOk() (*string, bool)`

GetVirtualDriveOk returns a tuple with the VirtualDrive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualDrive

`func (o *StorageFlexFlashVirtualDriveRelationship) SetVirtualDrive(v string)`

SetVirtualDrive sets VirtualDrive field to given value.

### HasVirtualDrive

`func (o *StorageFlexFlashVirtualDriveRelationship) HasVirtualDrive() bool`

HasVirtualDrive returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *StorageFlexFlashVirtualDriveRelationship) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *StorageFlexFlashVirtualDriveRelationship) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *StorageFlexFlashVirtualDriveRelationship) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *StorageFlexFlashVirtualDriveRelationship) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *StorageFlexFlashVirtualDriveRelationship) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StorageFlexFlashVirtualDriveRelationship) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StorageFlexFlashVirtualDriveRelationship) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StorageFlexFlashVirtualDriveRelationship) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### GetStorageFlexFlashController

`func (o *StorageFlexFlashVirtualDriveRelationship) GetStorageFlexFlashController() StorageFlexFlashControllerRelationship`

GetStorageFlexFlashController returns the StorageFlexFlashController field if non-nil, zero value otherwise.

### GetStorageFlexFlashControllerOk

`func (o *StorageFlexFlashVirtualDriveRelationship) GetStorageFlexFlashControllerOk() (*StorageFlexFlashControllerRelationship, bool)`

GetStorageFlexFlashControllerOk returns a tuple with the StorageFlexFlashController field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageFlexFlashController

`func (o *StorageFlexFlashVirtualDriveRelationship) SetStorageFlexFlashController(v StorageFlexFlashControllerRelationship)`

SetStorageFlexFlashController sets StorageFlexFlashController field to given value.

### HasStorageFlexFlashController

`func (o *StorageFlexFlashVirtualDriveRelationship) HasStorageFlexFlashController() bool`

HasStorageFlexFlashController returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



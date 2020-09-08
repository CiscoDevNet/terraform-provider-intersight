# StorageFlexFlashControllerRelationship

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
**ControllerState** | Pointer to **string** | State of the Flex Flash Storage Controller. | [optional] [readonly] 
**FfControllerId** | Pointer to **string** | Identifier for the Flex Flash Storage Controller. | [optional] [readonly] 
**ComputeBoard** | Pointer to [**ComputeBoardRelationship**](compute.Board.Relationship.md) |  | [optional] 
**FlexFlashControllerProps** | Pointer to [**[]StorageFlexFlashControllerPropsRelationship**](storage.FlexFlashControllerProps.Relationship.md) | An array of relationships to storageFlexFlashControllerProps resources. | [optional] [readonly] 
**FlexFlashPhysicalDrives** | Pointer to [**[]StorageFlexFlashPhysicalDriveRelationship**](storage.FlexFlashPhysicalDrive.Relationship.md) | An array of relationships to storageFlexFlashPhysicalDrive resources. | [optional] [readonly] 
**FlexFlashVirtualDrives** | Pointer to [**[]StorageFlexFlashVirtualDriveRelationship**](storage.FlexFlashVirtualDrive.Relationship.md) | An array of relationships to storageFlexFlashVirtualDrive resources. | [optional] [readonly] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](inventory.DeviceInfo.Relationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 
**RunningFirmware** | Pointer to [**[]FirmwareRunningFirmwareRelationship**](firmware.RunningFirmware.Relationship.md) | An array of relationships to firmwareRunningFirmware resources. | [optional] [readonly] 

## Methods

### NewStorageFlexFlashControllerRelationship

`func NewStorageFlexFlashControllerRelationship(classId string, objectType string, ) *StorageFlexFlashControllerRelationship`

NewStorageFlexFlashControllerRelationship instantiates a new StorageFlexFlashControllerRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageFlexFlashControllerRelationshipWithDefaults

`func NewStorageFlexFlashControllerRelationshipWithDefaults() *StorageFlexFlashControllerRelationship`

NewStorageFlexFlashControllerRelationshipWithDefaults instantiates a new StorageFlexFlashControllerRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountMoid

`func (o *StorageFlexFlashControllerRelationship) GetAccountMoid() string`

GetAccountMoid returns the AccountMoid field if non-nil, zero value otherwise.

### GetAccountMoidOk

`func (o *StorageFlexFlashControllerRelationship) GetAccountMoidOk() (*string, bool)`

GetAccountMoidOk returns a tuple with the AccountMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMoid

`func (o *StorageFlexFlashControllerRelationship) SetAccountMoid(v string)`

SetAccountMoid sets AccountMoid field to given value.

### HasAccountMoid

`func (o *StorageFlexFlashControllerRelationship) HasAccountMoid() bool`

HasAccountMoid returns a boolean if a field has been set.

### GetClassId

`func (o *StorageFlexFlashControllerRelationship) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageFlexFlashControllerRelationship) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageFlexFlashControllerRelationship) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetCreateTime

`func (o *StorageFlexFlashControllerRelationship) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *StorageFlexFlashControllerRelationship) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *StorageFlexFlashControllerRelationship) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *StorageFlexFlashControllerRelationship) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDomainGroupMoid

`func (o *StorageFlexFlashControllerRelationship) GetDomainGroupMoid() string`

GetDomainGroupMoid returns the DomainGroupMoid field if non-nil, zero value otherwise.

### GetDomainGroupMoidOk

`func (o *StorageFlexFlashControllerRelationship) GetDomainGroupMoidOk() (*string, bool)`

GetDomainGroupMoidOk returns a tuple with the DomainGroupMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGroupMoid

`func (o *StorageFlexFlashControllerRelationship) SetDomainGroupMoid(v string)`

SetDomainGroupMoid sets DomainGroupMoid field to given value.

### HasDomainGroupMoid

`func (o *StorageFlexFlashControllerRelationship) HasDomainGroupMoid() bool`

HasDomainGroupMoid returns a boolean if a field has been set.

### GetModTime

`func (o *StorageFlexFlashControllerRelationship) GetModTime() time.Time`

GetModTime returns the ModTime field if non-nil, zero value otherwise.

### GetModTimeOk

`func (o *StorageFlexFlashControllerRelationship) GetModTimeOk() (*time.Time, bool)`

GetModTimeOk returns a tuple with the ModTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModTime

`func (o *StorageFlexFlashControllerRelationship) SetModTime(v time.Time)`

SetModTime sets ModTime field to given value.

### HasModTime

`func (o *StorageFlexFlashControllerRelationship) HasModTime() bool`

HasModTime returns a boolean if a field has been set.

### GetMoid

`func (o *StorageFlexFlashControllerRelationship) GetMoid() string`

GetMoid returns the Moid field if non-nil, zero value otherwise.

### GetMoidOk

`func (o *StorageFlexFlashControllerRelationship) GetMoidOk() (*string, bool)`

GetMoidOk returns a tuple with the Moid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoid

`func (o *StorageFlexFlashControllerRelationship) SetMoid(v string)`

SetMoid sets Moid field to given value.

### HasMoid

`func (o *StorageFlexFlashControllerRelationship) HasMoid() bool`

HasMoid returns a boolean if a field has been set.

### GetObjectType

`func (o *StorageFlexFlashControllerRelationship) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageFlexFlashControllerRelationship) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageFlexFlashControllerRelationship) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOwners

`func (o *StorageFlexFlashControllerRelationship) GetOwners() []string`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *StorageFlexFlashControllerRelationship) GetOwnersOk() (*[]string, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *StorageFlexFlashControllerRelationship) SetOwners(v []string)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *StorageFlexFlashControllerRelationship) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetSharedScope

`func (o *StorageFlexFlashControllerRelationship) GetSharedScope() string`

GetSharedScope returns the SharedScope field if non-nil, zero value otherwise.

### GetSharedScopeOk

`func (o *StorageFlexFlashControllerRelationship) GetSharedScopeOk() (*string, bool)`

GetSharedScopeOk returns a tuple with the SharedScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedScope

`func (o *StorageFlexFlashControllerRelationship) SetSharedScope(v string)`

SetSharedScope sets SharedScope field to given value.

### HasSharedScope

`func (o *StorageFlexFlashControllerRelationship) HasSharedScope() bool`

HasSharedScope returns a boolean if a field has been set.

### GetTags

`func (o *StorageFlexFlashControllerRelationship) GetTags() []MoTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *StorageFlexFlashControllerRelationship) GetTagsOk() (*[]MoTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *StorageFlexFlashControllerRelationship) SetTags(v []MoTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *StorageFlexFlashControllerRelationship) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersionContext

`func (o *StorageFlexFlashControllerRelationship) GetVersionContext() MoVersionContext`

GetVersionContext returns the VersionContext field if non-nil, zero value otherwise.

### GetVersionContextOk

`func (o *StorageFlexFlashControllerRelationship) GetVersionContextOk() (*MoVersionContext, bool)`

GetVersionContextOk returns a tuple with the VersionContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionContext

`func (o *StorageFlexFlashControllerRelationship) SetVersionContext(v MoVersionContext)`

SetVersionContext sets VersionContext field to given value.

### HasVersionContext

`func (o *StorageFlexFlashControllerRelationship) HasVersionContext() bool`

HasVersionContext returns a boolean if a field has been set.

### GetAncestors

`func (o *StorageFlexFlashControllerRelationship) GetAncestors() []MoBaseMoRelationship`

GetAncestors returns the Ancestors field if non-nil, zero value otherwise.

### GetAncestorsOk

`func (o *StorageFlexFlashControllerRelationship) GetAncestorsOk() (*[]MoBaseMoRelationship, bool)`

GetAncestorsOk returns a tuple with the Ancestors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestors

`func (o *StorageFlexFlashControllerRelationship) SetAncestors(v []MoBaseMoRelationship)`

SetAncestors sets Ancestors field to given value.

### HasAncestors

`func (o *StorageFlexFlashControllerRelationship) HasAncestors() bool`

HasAncestors returns a boolean if a field has been set.

### SetAncestorsNil

`func (o *StorageFlexFlashControllerRelationship) SetAncestorsNil(b bool)`

 SetAncestorsNil sets the value for Ancestors to be an explicit nil

### UnsetAncestors
`func (o *StorageFlexFlashControllerRelationship) UnsetAncestors()`

UnsetAncestors ensures that no value is present for Ancestors, not even an explicit nil
### GetParent

`func (o *StorageFlexFlashControllerRelationship) GetParent() MoBaseMoRelationship`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *StorageFlexFlashControllerRelationship) GetParentOk() (*MoBaseMoRelationship, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *StorageFlexFlashControllerRelationship) SetParent(v MoBaseMoRelationship)`

SetParent sets Parent field to given value.

### HasParent

`func (o *StorageFlexFlashControllerRelationship) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPermissionResources

`func (o *StorageFlexFlashControllerRelationship) GetPermissionResources() []MoBaseMoRelationship`

GetPermissionResources returns the PermissionResources field if non-nil, zero value otherwise.

### GetPermissionResourcesOk

`func (o *StorageFlexFlashControllerRelationship) GetPermissionResourcesOk() (*[]MoBaseMoRelationship, bool)`

GetPermissionResourcesOk returns a tuple with the PermissionResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionResources

`func (o *StorageFlexFlashControllerRelationship) SetPermissionResources(v []MoBaseMoRelationship)`

SetPermissionResources sets PermissionResources field to given value.

### HasPermissionResources

`func (o *StorageFlexFlashControllerRelationship) HasPermissionResources() bool`

HasPermissionResources returns a boolean if a field has been set.

### SetPermissionResourcesNil

`func (o *StorageFlexFlashControllerRelationship) SetPermissionResourcesNil(b bool)`

 SetPermissionResourcesNil sets the value for PermissionResources to be an explicit nil

### UnsetPermissionResources
`func (o *StorageFlexFlashControllerRelationship) UnsetPermissionResources()`

UnsetPermissionResources ensures that no value is present for PermissionResources, not even an explicit nil
### GetDisplayNames

`func (o *StorageFlexFlashControllerRelationship) GetDisplayNames() map[string][]string`

GetDisplayNames returns the DisplayNames field if non-nil, zero value otherwise.

### GetDisplayNamesOk

`func (o *StorageFlexFlashControllerRelationship) GetDisplayNamesOk() (*map[string][]string, bool)`

GetDisplayNamesOk returns a tuple with the DisplayNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNames

`func (o *StorageFlexFlashControllerRelationship) SetDisplayNames(v map[string][]string)`

SetDisplayNames sets DisplayNames field to given value.

### HasDisplayNames

`func (o *StorageFlexFlashControllerRelationship) HasDisplayNames() bool`

HasDisplayNames returns a boolean if a field has been set.

### SetDisplayNamesNil

`func (o *StorageFlexFlashControllerRelationship) SetDisplayNamesNil(b bool)`

 SetDisplayNamesNil sets the value for DisplayNames to be an explicit nil

### UnsetDisplayNames
`func (o *StorageFlexFlashControllerRelationship) UnsetDisplayNames()`

UnsetDisplayNames ensures that no value is present for DisplayNames, not even an explicit nil
### GetDeviceMoId

`func (o *StorageFlexFlashControllerRelationship) GetDeviceMoId() string`

GetDeviceMoId returns the DeviceMoId field if non-nil, zero value otherwise.

### GetDeviceMoIdOk

`func (o *StorageFlexFlashControllerRelationship) GetDeviceMoIdOk() (*string, bool)`

GetDeviceMoIdOk returns a tuple with the DeviceMoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMoId

`func (o *StorageFlexFlashControllerRelationship) SetDeviceMoId(v string)`

SetDeviceMoId sets DeviceMoId field to given value.

### HasDeviceMoId

`func (o *StorageFlexFlashControllerRelationship) HasDeviceMoId() bool`

HasDeviceMoId returns a boolean if a field has been set.

### GetDn

`func (o *StorageFlexFlashControllerRelationship) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *StorageFlexFlashControllerRelationship) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *StorageFlexFlashControllerRelationship) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *StorageFlexFlashControllerRelationship) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetRn

`func (o *StorageFlexFlashControllerRelationship) GetRn() string`

GetRn returns the Rn field if non-nil, zero value otherwise.

### GetRnOk

`func (o *StorageFlexFlashControllerRelationship) GetRnOk() (*string, bool)`

GetRnOk returns a tuple with the Rn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRn

`func (o *StorageFlexFlashControllerRelationship) SetRn(v string)`

SetRn sets Rn field to given value.

### HasRn

`func (o *StorageFlexFlashControllerRelationship) HasRn() bool`

HasRn returns a boolean if a field has been set.

### GetModel

`func (o *StorageFlexFlashControllerRelationship) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *StorageFlexFlashControllerRelationship) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *StorageFlexFlashControllerRelationship) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *StorageFlexFlashControllerRelationship) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetRevision

`func (o *StorageFlexFlashControllerRelationship) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *StorageFlexFlashControllerRelationship) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *StorageFlexFlashControllerRelationship) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *StorageFlexFlashControllerRelationship) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetSerial

`func (o *StorageFlexFlashControllerRelationship) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *StorageFlexFlashControllerRelationship) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *StorageFlexFlashControllerRelationship) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *StorageFlexFlashControllerRelationship) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetVendor

`func (o *StorageFlexFlashControllerRelationship) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *StorageFlexFlashControllerRelationship) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *StorageFlexFlashControllerRelationship) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *StorageFlexFlashControllerRelationship) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetControllerState

`func (o *StorageFlexFlashControllerRelationship) GetControllerState() string`

GetControllerState returns the ControllerState field if non-nil, zero value otherwise.

### GetControllerStateOk

`func (o *StorageFlexFlashControllerRelationship) GetControllerStateOk() (*string, bool)`

GetControllerStateOk returns a tuple with the ControllerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerState

`func (o *StorageFlexFlashControllerRelationship) SetControllerState(v string)`

SetControllerState sets ControllerState field to given value.

### HasControllerState

`func (o *StorageFlexFlashControllerRelationship) HasControllerState() bool`

HasControllerState returns a boolean if a field has been set.

### GetFfControllerId

`func (o *StorageFlexFlashControllerRelationship) GetFfControllerId() string`

GetFfControllerId returns the FfControllerId field if non-nil, zero value otherwise.

### GetFfControllerIdOk

`func (o *StorageFlexFlashControllerRelationship) GetFfControllerIdOk() (*string, bool)`

GetFfControllerIdOk returns a tuple with the FfControllerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFfControllerId

`func (o *StorageFlexFlashControllerRelationship) SetFfControllerId(v string)`

SetFfControllerId sets FfControllerId field to given value.

### HasFfControllerId

`func (o *StorageFlexFlashControllerRelationship) HasFfControllerId() bool`

HasFfControllerId returns a boolean if a field has been set.

### GetComputeBoard

`func (o *StorageFlexFlashControllerRelationship) GetComputeBoard() ComputeBoardRelationship`

GetComputeBoard returns the ComputeBoard field if non-nil, zero value otherwise.

### GetComputeBoardOk

`func (o *StorageFlexFlashControllerRelationship) GetComputeBoardOk() (*ComputeBoardRelationship, bool)`

GetComputeBoardOk returns a tuple with the ComputeBoard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeBoard

`func (o *StorageFlexFlashControllerRelationship) SetComputeBoard(v ComputeBoardRelationship)`

SetComputeBoard sets ComputeBoard field to given value.

### HasComputeBoard

`func (o *StorageFlexFlashControllerRelationship) HasComputeBoard() bool`

HasComputeBoard returns a boolean if a field has been set.

### GetFlexFlashControllerProps

`func (o *StorageFlexFlashControllerRelationship) GetFlexFlashControllerProps() []StorageFlexFlashControllerPropsRelationship`

GetFlexFlashControllerProps returns the FlexFlashControllerProps field if non-nil, zero value otherwise.

### GetFlexFlashControllerPropsOk

`func (o *StorageFlexFlashControllerRelationship) GetFlexFlashControllerPropsOk() (*[]StorageFlexFlashControllerPropsRelationship, bool)`

GetFlexFlashControllerPropsOk returns a tuple with the FlexFlashControllerProps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexFlashControllerProps

`func (o *StorageFlexFlashControllerRelationship) SetFlexFlashControllerProps(v []StorageFlexFlashControllerPropsRelationship)`

SetFlexFlashControllerProps sets FlexFlashControllerProps field to given value.

### HasFlexFlashControllerProps

`func (o *StorageFlexFlashControllerRelationship) HasFlexFlashControllerProps() bool`

HasFlexFlashControllerProps returns a boolean if a field has been set.

### SetFlexFlashControllerPropsNil

`func (o *StorageFlexFlashControllerRelationship) SetFlexFlashControllerPropsNil(b bool)`

 SetFlexFlashControllerPropsNil sets the value for FlexFlashControllerProps to be an explicit nil

### UnsetFlexFlashControllerProps
`func (o *StorageFlexFlashControllerRelationship) UnsetFlexFlashControllerProps()`

UnsetFlexFlashControllerProps ensures that no value is present for FlexFlashControllerProps, not even an explicit nil
### GetFlexFlashPhysicalDrives

`func (o *StorageFlexFlashControllerRelationship) GetFlexFlashPhysicalDrives() []StorageFlexFlashPhysicalDriveRelationship`

GetFlexFlashPhysicalDrives returns the FlexFlashPhysicalDrives field if non-nil, zero value otherwise.

### GetFlexFlashPhysicalDrivesOk

`func (o *StorageFlexFlashControllerRelationship) GetFlexFlashPhysicalDrivesOk() (*[]StorageFlexFlashPhysicalDriveRelationship, bool)`

GetFlexFlashPhysicalDrivesOk returns a tuple with the FlexFlashPhysicalDrives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexFlashPhysicalDrives

`func (o *StorageFlexFlashControllerRelationship) SetFlexFlashPhysicalDrives(v []StorageFlexFlashPhysicalDriveRelationship)`

SetFlexFlashPhysicalDrives sets FlexFlashPhysicalDrives field to given value.

### HasFlexFlashPhysicalDrives

`func (o *StorageFlexFlashControllerRelationship) HasFlexFlashPhysicalDrives() bool`

HasFlexFlashPhysicalDrives returns a boolean if a field has been set.

### SetFlexFlashPhysicalDrivesNil

`func (o *StorageFlexFlashControllerRelationship) SetFlexFlashPhysicalDrivesNil(b bool)`

 SetFlexFlashPhysicalDrivesNil sets the value for FlexFlashPhysicalDrives to be an explicit nil

### UnsetFlexFlashPhysicalDrives
`func (o *StorageFlexFlashControllerRelationship) UnsetFlexFlashPhysicalDrives()`

UnsetFlexFlashPhysicalDrives ensures that no value is present for FlexFlashPhysicalDrives, not even an explicit nil
### GetFlexFlashVirtualDrives

`func (o *StorageFlexFlashControllerRelationship) GetFlexFlashVirtualDrives() []StorageFlexFlashVirtualDriveRelationship`

GetFlexFlashVirtualDrives returns the FlexFlashVirtualDrives field if non-nil, zero value otherwise.

### GetFlexFlashVirtualDrivesOk

`func (o *StorageFlexFlashControllerRelationship) GetFlexFlashVirtualDrivesOk() (*[]StorageFlexFlashVirtualDriveRelationship, bool)`

GetFlexFlashVirtualDrivesOk returns a tuple with the FlexFlashVirtualDrives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexFlashVirtualDrives

`func (o *StorageFlexFlashControllerRelationship) SetFlexFlashVirtualDrives(v []StorageFlexFlashVirtualDriveRelationship)`

SetFlexFlashVirtualDrives sets FlexFlashVirtualDrives field to given value.

### HasFlexFlashVirtualDrives

`func (o *StorageFlexFlashControllerRelationship) HasFlexFlashVirtualDrives() bool`

HasFlexFlashVirtualDrives returns a boolean if a field has been set.

### SetFlexFlashVirtualDrivesNil

`func (o *StorageFlexFlashControllerRelationship) SetFlexFlashVirtualDrivesNil(b bool)`

 SetFlexFlashVirtualDrivesNil sets the value for FlexFlashVirtualDrives to be an explicit nil

### UnsetFlexFlashVirtualDrives
`func (o *StorageFlexFlashControllerRelationship) UnsetFlexFlashVirtualDrives()`

UnsetFlexFlashVirtualDrives ensures that no value is present for FlexFlashVirtualDrives, not even an explicit nil
### GetInventoryDeviceInfo

`func (o *StorageFlexFlashControllerRelationship) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *StorageFlexFlashControllerRelationship) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *StorageFlexFlashControllerRelationship) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *StorageFlexFlashControllerRelationship) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *StorageFlexFlashControllerRelationship) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StorageFlexFlashControllerRelationship) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StorageFlexFlashControllerRelationship) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StorageFlexFlashControllerRelationship) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### GetRunningFirmware

`func (o *StorageFlexFlashControllerRelationship) GetRunningFirmware() []FirmwareRunningFirmwareRelationship`

GetRunningFirmware returns the RunningFirmware field if non-nil, zero value otherwise.

### GetRunningFirmwareOk

`func (o *StorageFlexFlashControllerRelationship) GetRunningFirmwareOk() (*[]FirmwareRunningFirmwareRelationship, bool)`

GetRunningFirmwareOk returns a tuple with the RunningFirmware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningFirmware

`func (o *StorageFlexFlashControllerRelationship) SetRunningFirmware(v []FirmwareRunningFirmwareRelationship)`

SetRunningFirmware sets RunningFirmware field to given value.

### HasRunningFirmware

`func (o *StorageFlexFlashControllerRelationship) HasRunningFirmware() bool`

HasRunningFirmware returns a boolean if a field has been set.

### SetRunningFirmwareNil

`func (o *StorageFlexFlashControllerRelationship) SetRunningFirmwareNil(b bool)`

 SetRunningFirmwareNil sets the value for RunningFirmware to be an explicit nil

### UnsetRunningFirmware
`func (o *StorageFlexFlashControllerRelationship) UnsetRunningFirmware()`

UnsetRunningFirmware ensures that no value is present for RunningFirmware, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



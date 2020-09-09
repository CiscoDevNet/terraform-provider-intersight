# MemoryArrayRelationship

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
**ArrayId** | Pointer to **int64** | The instance number of the memory array. | [optional] [readonly] 
**CpuId** | Pointer to **int64** | ID of the CPU that access this memory array. | [optional] [readonly] 
**CurrentCapacity** | Pointer to **string** | Current capacity of all the memory units on a server. | [optional] [readonly] 
**ErrorCorrection** | Pointer to **string** | The primary hardware error detection or correction method supported by the memory array. | [optional] [readonly] 
**MaxCapacity** | Pointer to **string** | Maximum capacity of all the memory units on a server. | [optional] [readonly] 
**MaxDevices** | Pointer to **string** | The maximum number of slots or sockets available for memory devices in the memory array. | [optional] [readonly] 
**OperPowerState** | Pointer to **string** | The power state indicator of the memory array. | [optional] [readonly] 
**Presence** | Pointer to **string** | The presence of atleast one memory device in the array. Valid values are &#39;equipped&#39; and &#39;absent&#39;. | [optional] [readonly] 
**ComputeBlade** | Pointer to [**ComputeBladeRelationship**](compute.Blade.Relationship.md) |  | [optional] 
**ComputeBoard** | Pointer to [**ComputeBoardRelationship**](compute.Board.Relationship.md) |  | [optional] 
**ComputeRackUnit** | Pointer to [**ComputeRackUnitRelationship**](compute.RackUnit.Relationship.md) |  | [optional] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](inventory.DeviceInfo.Relationship.md) |  | [optional] 
**PersistentMemoryUnits** | Pointer to [**[]MemoryPersistentMemoryUnitRelationship**](memory.PersistentMemoryUnit.Relationship.md) | An array of relationships to memoryPersistentMemoryUnit resources. | [optional] [readonly] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 
**Units** | Pointer to [**[]MemoryUnitRelationship**](memory.Unit.Relationship.md) | An array of relationships to memoryUnit resources. | [optional] [readonly] 

## Methods

### NewMemoryArrayRelationship

`func NewMemoryArrayRelationship(classId string, objectType string, ) *MemoryArrayRelationship`

NewMemoryArrayRelationship instantiates a new MemoryArrayRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemoryArrayRelationshipWithDefaults

`func NewMemoryArrayRelationshipWithDefaults() *MemoryArrayRelationship`

NewMemoryArrayRelationshipWithDefaults instantiates a new MemoryArrayRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountMoid

`func (o *MemoryArrayRelationship) GetAccountMoid() string`

GetAccountMoid returns the AccountMoid field if non-nil, zero value otherwise.

### GetAccountMoidOk

`func (o *MemoryArrayRelationship) GetAccountMoidOk() (*string, bool)`

GetAccountMoidOk returns a tuple with the AccountMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMoid

`func (o *MemoryArrayRelationship) SetAccountMoid(v string)`

SetAccountMoid sets AccountMoid field to given value.

### HasAccountMoid

`func (o *MemoryArrayRelationship) HasAccountMoid() bool`

HasAccountMoid returns a boolean if a field has been set.

### GetClassId

`func (o *MemoryArrayRelationship) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *MemoryArrayRelationship) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *MemoryArrayRelationship) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetCreateTime

`func (o *MemoryArrayRelationship) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *MemoryArrayRelationship) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *MemoryArrayRelationship) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *MemoryArrayRelationship) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDomainGroupMoid

`func (o *MemoryArrayRelationship) GetDomainGroupMoid() string`

GetDomainGroupMoid returns the DomainGroupMoid field if non-nil, zero value otherwise.

### GetDomainGroupMoidOk

`func (o *MemoryArrayRelationship) GetDomainGroupMoidOk() (*string, bool)`

GetDomainGroupMoidOk returns a tuple with the DomainGroupMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGroupMoid

`func (o *MemoryArrayRelationship) SetDomainGroupMoid(v string)`

SetDomainGroupMoid sets DomainGroupMoid field to given value.

### HasDomainGroupMoid

`func (o *MemoryArrayRelationship) HasDomainGroupMoid() bool`

HasDomainGroupMoid returns a boolean if a field has been set.

### GetModTime

`func (o *MemoryArrayRelationship) GetModTime() time.Time`

GetModTime returns the ModTime field if non-nil, zero value otherwise.

### GetModTimeOk

`func (o *MemoryArrayRelationship) GetModTimeOk() (*time.Time, bool)`

GetModTimeOk returns a tuple with the ModTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModTime

`func (o *MemoryArrayRelationship) SetModTime(v time.Time)`

SetModTime sets ModTime field to given value.

### HasModTime

`func (o *MemoryArrayRelationship) HasModTime() bool`

HasModTime returns a boolean if a field has been set.

### GetMoid

`func (o *MemoryArrayRelationship) GetMoid() string`

GetMoid returns the Moid field if non-nil, zero value otherwise.

### GetMoidOk

`func (o *MemoryArrayRelationship) GetMoidOk() (*string, bool)`

GetMoidOk returns a tuple with the Moid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoid

`func (o *MemoryArrayRelationship) SetMoid(v string)`

SetMoid sets Moid field to given value.

### HasMoid

`func (o *MemoryArrayRelationship) HasMoid() bool`

HasMoid returns a boolean if a field has been set.

### GetObjectType

`func (o *MemoryArrayRelationship) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *MemoryArrayRelationship) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *MemoryArrayRelationship) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOwners

`func (o *MemoryArrayRelationship) GetOwners() []string`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *MemoryArrayRelationship) GetOwnersOk() (*[]string, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *MemoryArrayRelationship) SetOwners(v []string)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *MemoryArrayRelationship) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetSharedScope

`func (o *MemoryArrayRelationship) GetSharedScope() string`

GetSharedScope returns the SharedScope field if non-nil, zero value otherwise.

### GetSharedScopeOk

`func (o *MemoryArrayRelationship) GetSharedScopeOk() (*string, bool)`

GetSharedScopeOk returns a tuple with the SharedScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedScope

`func (o *MemoryArrayRelationship) SetSharedScope(v string)`

SetSharedScope sets SharedScope field to given value.

### HasSharedScope

`func (o *MemoryArrayRelationship) HasSharedScope() bool`

HasSharedScope returns a boolean if a field has been set.

### GetTags

`func (o *MemoryArrayRelationship) GetTags() []MoTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *MemoryArrayRelationship) GetTagsOk() (*[]MoTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *MemoryArrayRelationship) SetTags(v []MoTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *MemoryArrayRelationship) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersionContext

`func (o *MemoryArrayRelationship) GetVersionContext() MoVersionContext`

GetVersionContext returns the VersionContext field if non-nil, zero value otherwise.

### GetVersionContextOk

`func (o *MemoryArrayRelationship) GetVersionContextOk() (*MoVersionContext, bool)`

GetVersionContextOk returns a tuple with the VersionContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionContext

`func (o *MemoryArrayRelationship) SetVersionContext(v MoVersionContext)`

SetVersionContext sets VersionContext field to given value.

### HasVersionContext

`func (o *MemoryArrayRelationship) HasVersionContext() bool`

HasVersionContext returns a boolean if a field has been set.

### GetAncestors

`func (o *MemoryArrayRelationship) GetAncestors() []MoBaseMoRelationship`

GetAncestors returns the Ancestors field if non-nil, zero value otherwise.

### GetAncestorsOk

`func (o *MemoryArrayRelationship) GetAncestorsOk() (*[]MoBaseMoRelationship, bool)`

GetAncestorsOk returns a tuple with the Ancestors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestors

`func (o *MemoryArrayRelationship) SetAncestors(v []MoBaseMoRelationship)`

SetAncestors sets Ancestors field to given value.

### HasAncestors

`func (o *MemoryArrayRelationship) HasAncestors() bool`

HasAncestors returns a boolean if a field has been set.

### SetAncestorsNil

`func (o *MemoryArrayRelationship) SetAncestorsNil(b bool)`

 SetAncestorsNil sets the value for Ancestors to be an explicit nil

### UnsetAncestors
`func (o *MemoryArrayRelationship) UnsetAncestors()`

UnsetAncestors ensures that no value is present for Ancestors, not even an explicit nil
### GetParent

`func (o *MemoryArrayRelationship) GetParent() MoBaseMoRelationship`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *MemoryArrayRelationship) GetParentOk() (*MoBaseMoRelationship, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *MemoryArrayRelationship) SetParent(v MoBaseMoRelationship)`

SetParent sets Parent field to given value.

### HasParent

`func (o *MemoryArrayRelationship) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPermissionResources

`func (o *MemoryArrayRelationship) GetPermissionResources() []MoBaseMoRelationship`

GetPermissionResources returns the PermissionResources field if non-nil, zero value otherwise.

### GetPermissionResourcesOk

`func (o *MemoryArrayRelationship) GetPermissionResourcesOk() (*[]MoBaseMoRelationship, bool)`

GetPermissionResourcesOk returns a tuple with the PermissionResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionResources

`func (o *MemoryArrayRelationship) SetPermissionResources(v []MoBaseMoRelationship)`

SetPermissionResources sets PermissionResources field to given value.

### HasPermissionResources

`func (o *MemoryArrayRelationship) HasPermissionResources() bool`

HasPermissionResources returns a boolean if a field has been set.

### SetPermissionResourcesNil

`func (o *MemoryArrayRelationship) SetPermissionResourcesNil(b bool)`

 SetPermissionResourcesNil sets the value for PermissionResources to be an explicit nil

### UnsetPermissionResources
`func (o *MemoryArrayRelationship) UnsetPermissionResources()`

UnsetPermissionResources ensures that no value is present for PermissionResources, not even an explicit nil
### GetDisplayNames

`func (o *MemoryArrayRelationship) GetDisplayNames() map[string][]string`

GetDisplayNames returns the DisplayNames field if non-nil, zero value otherwise.

### GetDisplayNamesOk

`func (o *MemoryArrayRelationship) GetDisplayNamesOk() (*map[string][]string, bool)`

GetDisplayNamesOk returns a tuple with the DisplayNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNames

`func (o *MemoryArrayRelationship) SetDisplayNames(v map[string][]string)`

SetDisplayNames sets DisplayNames field to given value.

### HasDisplayNames

`func (o *MemoryArrayRelationship) HasDisplayNames() bool`

HasDisplayNames returns a boolean if a field has been set.

### SetDisplayNamesNil

`func (o *MemoryArrayRelationship) SetDisplayNamesNil(b bool)`

 SetDisplayNamesNil sets the value for DisplayNames to be an explicit nil

### UnsetDisplayNames
`func (o *MemoryArrayRelationship) UnsetDisplayNames()`

UnsetDisplayNames ensures that no value is present for DisplayNames, not even an explicit nil
### GetDeviceMoId

`func (o *MemoryArrayRelationship) GetDeviceMoId() string`

GetDeviceMoId returns the DeviceMoId field if non-nil, zero value otherwise.

### GetDeviceMoIdOk

`func (o *MemoryArrayRelationship) GetDeviceMoIdOk() (*string, bool)`

GetDeviceMoIdOk returns a tuple with the DeviceMoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMoId

`func (o *MemoryArrayRelationship) SetDeviceMoId(v string)`

SetDeviceMoId sets DeviceMoId field to given value.

### HasDeviceMoId

`func (o *MemoryArrayRelationship) HasDeviceMoId() bool`

HasDeviceMoId returns a boolean if a field has been set.

### GetDn

`func (o *MemoryArrayRelationship) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *MemoryArrayRelationship) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *MemoryArrayRelationship) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *MemoryArrayRelationship) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetRn

`func (o *MemoryArrayRelationship) GetRn() string`

GetRn returns the Rn field if non-nil, zero value otherwise.

### GetRnOk

`func (o *MemoryArrayRelationship) GetRnOk() (*string, bool)`

GetRnOk returns a tuple with the Rn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRn

`func (o *MemoryArrayRelationship) SetRn(v string)`

SetRn sets Rn field to given value.

### HasRn

`func (o *MemoryArrayRelationship) HasRn() bool`

HasRn returns a boolean if a field has been set.

### GetModel

`func (o *MemoryArrayRelationship) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *MemoryArrayRelationship) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *MemoryArrayRelationship) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *MemoryArrayRelationship) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetRevision

`func (o *MemoryArrayRelationship) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *MemoryArrayRelationship) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *MemoryArrayRelationship) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *MemoryArrayRelationship) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetSerial

`func (o *MemoryArrayRelationship) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *MemoryArrayRelationship) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *MemoryArrayRelationship) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *MemoryArrayRelationship) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetVendor

`func (o *MemoryArrayRelationship) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *MemoryArrayRelationship) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *MemoryArrayRelationship) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *MemoryArrayRelationship) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetArrayId

`func (o *MemoryArrayRelationship) GetArrayId() int64`

GetArrayId returns the ArrayId field if non-nil, zero value otherwise.

### GetArrayIdOk

`func (o *MemoryArrayRelationship) GetArrayIdOk() (*int64, bool)`

GetArrayIdOk returns a tuple with the ArrayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrayId

`func (o *MemoryArrayRelationship) SetArrayId(v int64)`

SetArrayId sets ArrayId field to given value.

### HasArrayId

`func (o *MemoryArrayRelationship) HasArrayId() bool`

HasArrayId returns a boolean if a field has been set.

### GetCpuId

`func (o *MemoryArrayRelationship) GetCpuId() int64`

GetCpuId returns the CpuId field if non-nil, zero value otherwise.

### GetCpuIdOk

`func (o *MemoryArrayRelationship) GetCpuIdOk() (*int64, bool)`

GetCpuIdOk returns a tuple with the CpuId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuId

`func (o *MemoryArrayRelationship) SetCpuId(v int64)`

SetCpuId sets CpuId field to given value.

### HasCpuId

`func (o *MemoryArrayRelationship) HasCpuId() bool`

HasCpuId returns a boolean if a field has been set.

### GetCurrentCapacity

`func (o *MemoryArrayRelationship) GetCurrentCapacity() string`

GetCurrentCapacity returns the CurrentCapacity field if non-nil, zero value otherwise.

### GetCurrentCapacityOk

`func (o *MemoryArrayRelationship) GetCurrentCapacityOk() (*string, bool)`

GetCurrentCapacityOk returns a tuple with the CurrentCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentCapacity

`func (o *MemoryArrayRelationship) SetCurrentCapacity(v string)`

SetCurrentCapacity sets CurrentCapacity field to given value.

### HasCurrentCapacity

`func (o *MemoryArrayRelationship) HasCurrentCapacity() bool`

HasCurrentCapacity returns a boolean if a field has been set.

### GetErrorCorrection

`func (o *MemoryArrayRelationship) GetErrorCorrection() string`

GetErrorCorrection returns the ErrorCorrection field if non-nil, zero value otherwise.

### GetErrorCorrectionOk

`func (o *MemoryArrayRelationship) GetErrorCorrectionOk() (*string, bool)`

GetErrorCorrectionOk returns a tuple with the ErrorCorrection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCorrection

`func (o *MemoryArrayRelationship) SetErrorCorrection(v string)`

SetErrorCorrection sets ErrorCorrection field to given value.

### HasErrorCorrection

`func (o *MemoryArrayRelationship) HasErrorCorrection() bool`

HasErrorCorrection returns a boolean if a field has been set.

### GetMaxCapacity

`func (o *MemoryArrayRelationship) GetMaxCapacity() string`

GetMaxCapacity returns the MaxCapacity field if non-nil, zero value otherwise.

### GetMaxCapacityOk

`func (o *MemoryArrayRelationship) GetMaxCapacityOk() (*string, bool)`

GetMaxCapacityOk returns a tuple with the MaxCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCapacity

`func (o *MemoryArrayRelationship) SetMaxCapacity(v string)`

SetMaxCapacity sets MaxCapacity field to given value.

### HasMaxCapacity

`func (o *MemoryArrayRelationship) HasMaxCapacity() bool`

HasMaxCapacity returns a boolean if a field has been set.

### GetMaxDevices

`func (o *MemoryArrayRelationship) GetMaxDevices() string`

GetMaxDevices returns the MaxDevices field if non-nil, zero value otherwise.

### GetMaxDevicesOk

`func (o *MemoryArrayRelationship) GetMaxDevicesOk() (*string, bool)`

GetMaxDevicesOk returns a tuple with the MaxDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDevices

`func (o *MemoryArrayRelationship) SetMaxDevices(v string)`

SetMaxDevices sets MaxDevices field to given value.

### HasMaxDevices

`func (o *MemoryArrayRelationship) HasMaxDevices() bool`

HasMaxDevices returns a boolean if a field has been set.

### GetOperPowerState

`func (o *MemoryArrayRelationship) GetOperPowerState() string`

GetOperPowerState returns the OperPowerState field if non-nil, zero value otherwise.

### GetOperPowerStateOk

`func (o *MemoryArrayRelationship) GetOperPowerStateOk() (*string, bool)`

GetOperPowerStateOk returns a tuple with the OperPowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperPowerState

`func (o *MemoryArrayRelationship) SetOperPowerState(v string)`

SetOperPowerState sets OperPowerState field to given value.

### HasOperPowerState

`func (o *MemoryArrayRelationship) HasOperPowerState() bool`

HasOperPowerState returns a boolean if a field has been set.

### GetPresence

`func (o *MemoryArrayRelationship) GetPresence() string`

GetPresence returns the Presence field if non-nil, zero value otherwise.

### GetPresenceOk

`func (o *MemoryArrayRelationship) GetPresenceOk() (*string, bool)`

GetPresenceOk returns a tuple with the Presence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresence

`func (o *MemoryArrayRelationship) SetPresence(v string)`

SetPresence sets Presence field to given value.

### HasPresence

`func (o *MemoryArrayRelationship) HasPresence() bool`

HasPresence returns a boolean if a field has been set.

### GetComputeBlade

`func (o *MemoryArrayRelationship) GetComputeBlade() ComputeBladeRelationship`

GetComputeBlade returns the ComputeBlade field if non-nil, zero value otherwise.

### GetComputeBladeOk

`func (o *MemoryArrayRelationship) GetComputeBladeOk() (*ComputeBladeRelationship, bool)`

GetComputeBladeOk returns a tuple with the ComputeBlade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeBlade

`func (o *MemoryArrayRelationship) SetComputeBlade(v ComputeBladeRelationship)`

SetComputeBlade sets ComputeBlade field to given value.

### HasComputeBlade

`func (o *MemoryArrayRelationship) HasComputeBlade() bool`

HasComputeBlade returns a boolean if a field has been set.

### GetComputeBoard

`func (o *MemoryArrayRelationship) GetComputeBoard() ComputeBoardRelationship`

GetComputeBoard returns the ComputeBoard field if non-nil, zero value otherwise.

### GetComputeBoardOk

`func (o *MemoryArrayRelationship) GetComputeBoardOk() (*ComputeBoardRelationship, bool)`

GetComputeBoardOk returns a tuple with the ComputeBoard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeBoard

`func (o *MemoryArrayRelationship) SetComputeBoard(v ComputeBoardRelationship)`

SetComputeBoard sets ComputeBoard field to given value.

### HasComputeBoard

`func (o *MemoryArrayRelationship) HasComputeBoard() bool`

HasComputeBoard returns a boolean if a field has been set.

### GetComputeRackUnit

`func (o *MemoryArrayRelationship) GetComputeRackUnit() ComputeRackUnitRelationship`

GetComputeRackUnit returns the ComputeRackUnit field if non-nil, zero value otherwise.

### GetComputeRackUnitOk

`func (o *MemoryArrayRelationship) GetComputeRackUnitOk() (*ComputeRackUnitRelationship, bool)`

GetComputeRackUnitOk returns a tuple with the ComputeRackUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRackUnit

`func (o *MemoryArrayRelationship) SetComputeRackUnit(v ComputeRackUnitRelationship)`

SetComputeRackUnit sets ComputeRackUnit field to given value.

### HasComputeRackUnit

`func (o *MemoryArrayRelationship) HasComputeRackUnit() bool`

HasComputeRackUnit returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *MemoryArrayRelationship) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *MemoryArrayRelationship) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *MemoryArrayRelationship) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *MemoryArrayRelationship) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetPersistentMemoryUnits

`func (o *MemoryArrayRelationship) GetPersistentMemoryUnits() []MemoryPersistentMemoryUnitRelationship`

GetPersistentMemoryUnits returns the PersistentMemoryUnits field if non-nil, zero value otherwise.

### GetPersistentMemoryUnitsOk

`func (o *MemoryArrayRelationship) GetPersistentMemoryUnitsOk() (*[]MemoryPersistentMemoryUnitRelationship, bool)`

GetPersistentMemoryUnitsOk returns a tuple with the PersistentMemoryUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentMemoryUnits

`func (o *MemoryArrayRelationship) SetPersistentMemoryUnits(v []MemoryPersistentMemoryUnitRelationship)`

SetPersistentMemoryUnits sets PersistentMemoryUnits field to given value.

### HasPersistentMemoryUnits

`func (o *MemoryArrayRelationship) HasPersistentMemoryUnits() bool`

HasPersistentMemoryUnits returns a boolean if a field has been set.

### SetPersistentMemoryUnitsNil

`func (o *MemoryArrayRelationship) SetPersistentMemoryUnitsNil(b bool)`

 SetPersistentMemoryUnitsNil sets the value for PersistentMemoryUnits to be an explicit nil

### UnsetPersistentMemoryUnits
`func (o *MemoryArrayRelationship) UnsetPersistentMemoryUnits()`

UnsetPersistentMemoryUnits ensures that no value is present for PersistentMemoryUnits, not even an explicit nil
### GetRegisteredDevice

`func (o *MemoryArrayRelationship) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *MemoryArrayRelationship) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *MemoryArrayRelationship) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *MemoryArrayRelationship) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### GetUnits

`func (o *MemoryArrayRelationship) GetUnits() []MemoryUnitRelationship`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *MemoryArrayRelationship) GetUnitsOk() (*[]MemoryUnitRelationship, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *MemoryArrayRelationship) SetUnits(v []MemoryUnitRelationship)`

SetUnits sets Units field to given value.

### HasUnits

`func (o *MemoryArrayRelationship) HasUnits() bool`

HasUnits returns a boolean if a field has been set.

### SetUnitsNil

`func (o *MemoryArrayRelationship) SetUnitsNil(b bool)`

 SetUnitsNil sets the value for Units to be an explicit nil

### UnsetUnits
`func (o *MemoryArrayRelationship) UnsetUnits()`

UnsetUnits ensures that no value is present for Units, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



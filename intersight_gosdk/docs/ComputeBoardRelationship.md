# ComputeBoardRelationship

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
**BoardId** | Pointer to **int64** | The identity of the motherboard. | [optional] [readonly] 
**CpuTypeController** | Pointer to **string** | The type of central processing unit on the mother board. | [optional] [readonly] 
**OperPowerState** | Pointer to **string** | Current power state of the mother board of the server. | [optional] [readonly] 
**Presence** | Pointer to **string** | Identifies the presence of the mother board of the server. | [optional] [readonly] 
**ComputeBlade** | Pointer to [**ComputeBladeRelationship**](compute.Blade.Relationship.md) |  | [optional] 
**ComputeRackUnit** | Pointer to [**ComputeRackUnitRelationship**](compute.RackUnit.Relationship.md) |  | [optional] 
**EquipmentTpms** | Pointer to [**[]EquipmentTpmRelationship**](equipment.Tpm.Relationship.md) | An array of relationships to equipmentTpm resources. | [optional] [readonly] 
**GraphicsCards** | Pointer to [**[]GraphicsCardRelationship**](graphics.Card.Relationship.md) | An array of relationships to graphicsCard resources. | [optional] [readonly] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](inventory.DeviceInfo.Relationship.md) |  | [optional] 
**MemoryArrays** | Pointer to [**[]MemoryArrayRelationship**](memory.Array.Relationship.md) | An array of relationships to memoryArray resources. | [optional] [readonly] 
**PciCoprocessorCards** | Pointer to [**[]PciCoprocessorCardRelationship**](pci.CoprocessorCard.Relationship.md) | An array of relationships to pciCoprocessorCard resources. | [optional] [readonly] 
**PciSwitch** | Pointer to [**[]PciSwitchRelationship**](pci.Switch.Relationship.md) | An array of relationships to pciSwitch resources. | [optional] [readonly] 
**PersistentMemoryConfiguration** | Pointer to [**MemoryPersistentMemoryConfigurationRelationship**](memory.PersistentMemoryConfiguration.Relationship.md) |  | [optional] 
**Processors** | Pointer to [**[]ProcessorUnitRelationship**](processor.Unit.Relationship.md) | An array of relationships to processorUnit resources. | [optional] [readonly] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 
**SecurityUnits** | Pointer to [**[]SecurityUnitRelationship**](security.Unit.Relationship.md) | An array of relationships to securityUnit resources. | [optional] [readonly] 
**StorageControllers** | Pointer to [**[]StorageControllerRelationship**](storage.Controller.Relationship.md) | An array of relationships to storageController resources. | [optional] [readonly] 
**StorageFlexFlashControllers** | Pointer to [**[]StorageFlexFlashControllerRelationship**](storage.FlexFlashController.Relationship.md) | An array of relationships to storageFlexFlashController resources. | [optional] [readonly] 
**StorageFlexUtilControllers** | Pointer to [**[]StorageFlexUtilControllerRelationship**](storage.FlexUtilController.Relationship.md) | An array of relationships to storageFlexUtilController resources. | [optional] [readonly] 

## Methods

### NewComputeBoardRelationship

`func NewComputeBoardRelationship(classId string, objectType string, ) *ComputeBoardRelationship`

NewComputeBoardRelationship instantiates a new ComputeBoardRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputeBoardRelationshipWithDefaults

`func NewComputeBoardRelationshipWithDefaults() *ComputeBoardRelationship`

NewComputeBoardRelationshipWithDefaults instantiates a new ComputeBoardRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountMoid

`func (o *ComputeBoardRelationship) GetAccountMoid() string`

GetAccountMoid returns the AccountMoid field if non-nil, zero value otherwise.

### GetAccountMoidOk

`func (o *ComputeBoardRelationship) GetAccountMoidOk() (*string, bool)`

GetAccountMoidOk returns a tuple with the AccountMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMoid

`func (o *ComputeBoardRelationship) SetAccountMoid(v string)`

SetAccountMoid sets AccountMoid field to given value.

### HasAccountMoid

`func (o *ComputeBoardRelationship) HasAccountMoid() bool`

HasAccountMoid returns a boolean if a field has been set.

### GetClassId

`func (o *ComputeBoardRelationship) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ComputeBoardRelationship) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ComputeBoardRelationship) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetCreateTime

`func (o *ComputeBoardRelationship) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *ComputeBoardRelationship) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *ComputeBoardRelationship) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *ComputeBoardRelationship) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDomainGroupMoid

`func (o *ComputeBoardRelationship) GetDomainGroupMoid() string`

GetDomainGroupMoid returns the DomainGroupMoid field if non-nil, zero value otherwise.

### GetDomainGroupMoidOk

`func (o *ComputeBoardRelationship) GetDomainGroupMoidOk() (*string, bool)`

GetDomainGroupMoidOk returns a tuple with the DomainGroupMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGroupMoid

`func (o *ComputeBoardRelationship) SetDomainGroupMoid(v string)`

SetDomainGroupMoid sets DomainGroupMoid field to given value.

### HasDomainGroupMoid

`func (o *ComputeBoardRelationship) HasDomainGroupMoid() bool`

HasDomainGroupMoid returns a boolean if a field has been set.

### GetModTime

`func (o *ComputeBoardRelationship) GetModTime() time.Time`

GetModTime returns the ModTime field if non-nil, zero value otherwise.

### GetModTimeOk

`func (o *ComputeBoardRelationship) GetModTimeOk() (*time.Time, bool)`

GetModTimeOk returns a tuple with the ModTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModTime

`func (o *ComputeBoardRelationship) SetModTime(v time.Time)`

SetModTime sets ModTime field to given value.

### HasModTime

`func (o *ComputeBoardRelationship) HasModTime() bool`

HasModTime returns a boolean if a field has been set.

### GetMoid

`func (o *ComputeBoardRelationship) GetMoid() string`

GetMoid returns the Moid field if non-nil, zero value otherwise.

### GetMoidOk

`func (o *ComputeBoardRelationship) GetMoidOk() (*string, bool)`

GetMoidOk returns a tuple with the Moid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoid

`func (o *ComputeBoardRelationship) SetMoid(v string)`

SetMoid sets Moid field to given value.

### HasMoid

`func (o *ComputeBoardRelationship) HasMoid() bool`

HasMoid returns a boolean if a field has been set.

### GetObjectType

`func (o *ComputeBoardRelationship) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ComputeBoardRelationship) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ComputeBoardRelationship) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOwners

`func (o *ComputeBoardRelationship) GetOwners() []string`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *ComputeBoardRelationship) GetOwnersOk() (*[]string, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *ComputeBoardRelationship) SetOwners(v []string)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *ComputeBoardRelationship) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetSharedScope

`func (o *ComputeBoardRelationship) GetSharedScope() string`

GetSharedScope returns the SharedScope field if non-nil, zero value otherwise.

### GetSharedScopeOk

`func (o *ComputeBoardRelationship) GetSharedScopeOk() (*string, bool)`

GetSharedScopeOk returns a tuple with the SharedScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedScope

`func (o *ComputeBoardRelationship) SetSharedScope(v string)`

SetSharedScope sets SharedScope field to given value.

### HasSharedScope

`func (o *ComputeBoardRelationship) HasSharedScope() bool`

HasSharedScope returns a boolean if a field has been set.

### GetTags

`func (o *ComputeBoardRelationship) GetTags() []MoTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ComputeBoardRelationship) GetTagsOk() (*[]MoTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ComputeBoardRelationship) SetTags(v []MoTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ComputeBoardRelationship) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersionContext

`func (o *ComputeBoardRelationship) GetVersionContext() MoVersionContext`

GetVersionContext returns the VersionContext field if non-nil, zero value otherwise.

### GetVersionContextOk

`func (o *ComputeBoardRelationship) GetVersionContextOk() (*MoVersionContext, bool)`

GetVersionContextOk returns a tuple with the VersionContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionContext

`func (o *ComputeBoardRelationship) SetVersionContext(v MoVersionContext)`

SetVersionContext sets VersionContext field to given value.

### HasVersionContext

`func (o *ComputeBoardRelationship) HasVersionContext() bool`

HasVersionContext returns a boolean if a field has been set.

### GetAncestors

`func (o *ComputeBoardRelationship) GetAncestors() []MoBaseMoRelationship`

GetAncestors returns the Ancestors field if non-nil, zero value otherwise.

### GetAncestorsOk

`func (o *ComputeBoardRelationship) GetAncestorsOk() (*[]MoBaseMoRelationship, bool)`

GetAncestorsOk returns a tuple with the Ancestors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestors

`func (o *ComputeBoardRelationship) SetAncestors(v []MoBaseMoRelationship)`

SetAncestors sets Ancestors field to given value.

### HasAncestors

`func (o *ComputeBoardRelationship) HasAncestors() bool`

HasAncestors returns a boolean if a field has been set.

### SetAncestorsNil

`func (o *ComputeBoardRelationship) SetAncestorsNil(b bool)`

 SetAncestorsNil sets the value for Ancestors to be an explicit nil

### UnsetAncestors
`func (o *ComputeBoardRelationship) UnsetAncestors()`

UnsetAncestors ensures that no value is present for Ancestors, not even an explicit nil
### GetParent

`func (o *ComputeBoardRelationship) GetParent() MoBaseMoRelationship`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *ComputeBoardRelationship) GetParentOk() (*MoBaseMoRelationship, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *ComputeBoardRelationship) SetParent(v MoBaseMoRelationship)`

SetParent sets Parent field to given value.

### HasParent

`func (o *ComputeBoardRelationship) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPermissionResources

`func (o *ComputeBoardRelationship) GetPermissionResources() []MoBaseMoRelationship`

GetPermissionResources returns the PermissionResources field if non-nil, zero value otherwise.

### GetPermissionResourcesOk

`func (o *ComputeBoardRelationship) GetPermissionResourcesOk() (*[]MoBaseMoRelationship, bool)`

GetPermissionResourcesOk returns a tuple with the PermissionResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionResources

`func (o *ComputeBoardRelationship) SetPermissionResources(v []MoBaseMoRelationship)`

SetPermissionResources sets PermissionResources field to given value.

### HasPermissionResources

`func (o *ComputeBoardRelationship) HasPermissionResources() bool`

HasPermissionResources returns a boolean if a field has been set.

### SetPermissionResourcesNil

`func (o *ComputeBoardRelationship) SetPermissionResourcesNil(b bool)`

 SetPermissionResourcesNil sets the value for PermissionResources to be an explicit nil

### UnsetPermissionResources
`func (o *ComputeBoardRelationship) UnsetPermissionResources()`

UnsetPermissionResources ensures that no value is present for PermissionResources, not even an explicit nil
### GetDisplayNames

`func (o *ComputeBoardRelationship) GetDisplayNames() map[string][]string`

GetDisplayNames returns the DisplayNames field if non-nil, zero value otherwise.

### GetDisplayNamesOk

`func (o *ComputeBoardRelationship) GetDisplayNamesOk() (*map[string][]string, bool)`

GetDisplayNamesOk returns a tuple with the DisplayNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNames

`func (o *ComputeBoardRelationship) SetDisplayNames(v map[string][]string)`

SetDisplayNames sets DisplayNames field to given value.

### HasDisplayNames

`func (o *ComputeBoardRelationship) HasDisplayNames() bool`

HasDisplayNames returns a boolean if a field has been set.

### SetDisplayNamesNil

`func (o *ComputeBoardRelationship) SetDisplayNamesNil(b bool)`

 SetDisplayNamesNil sets the value for DisplayNames to be an explicit nil

### UnsetDisplayNames
`func (o *ComputeBoardRelationship) UnsetDisplayNames()`

UnsetDisplayNames ensures that no value is present for DisplayNames, not even an explicit nil
### GetDeviceMoId

`func (o *ComputeBoardRelationship) GetDeviceMoId() string`

GetDeviceMoId returns the DeviceMoId field if non-nil, zero value otherwise.

### GetDeviceMoIdOk

`func (o *ComputeBoardRelationship) GetDeviceMoIdOk() (*string, bool)`

GetDeviceMoIdOk returns a tuple with the DeviceMoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMoId

`func (o *ComputeBoardRelationship) SetDeviceMoId(v string)`

SetDeviceMoId sets DeviceMoId field to given value.

### HasDeviceMoId

`func (o *ComputeBoardRelationship) HasDeviceMoId() bool`

HasDeviceMoId returns a boolean if a field has been set.

### GetDn

`func (o *ComputeBoardRelationship) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *ComputeBoardRelationship) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *ComputeBoardRelationship) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *ComputeBoardRelationship) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetRn

`func (o *ComputeBoardRelationship) GetRn() string`

GetRn returns the Rn field if non-nil, zero value otherwise.

### GetRnOk

`func (o *ComputeBoardRelationship) GetRnOk() (*string, bool)`

GetRnOk returns a tuple with the Rn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRn

`func (o *ComputeBoardRelationship) SetRn(v string)`

SetRn sets Rn field to given value.

### HasRn

`func (o *ComputeBoardRelationship) HasRn() bool`

HasRn returns a boolean if a field has been set.

### GetModel

`func (o *ComputeBoardRelationship) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ComputeBoardRelationship) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ComputeBoardRelationship) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ComputeBoardRelationship) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetRevision

`func (o *ComputeBoardRelationship) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *ComputeBoardRelationship) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *ComputeBoardRelationship) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *ComputeBoardRelationship) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetSerial

`func (o *ComputeBoardRelationship) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *ComputeBoardRelationship) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *ComputeBoardRelationship) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *ComputeBoardRelationship) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetVendor

`func (o *ComputeBoardRelationship) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *ComputeBoardRelationship) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *ComputeBoardRelationship) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *ComputeBoardRelationship) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetBoardId

`func (o *ComputeBoardRelationship) GetBoardId() int64`

GetBoardId returns the BoardId field if non-nil, zero value otherwise.

### GetBoardIdOk

`func (o *ComputeBoardRelationship) GetBoardIdOk() (*int64, bool)`

GetBoardIdOk returns a tuple with the BoardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoardId

`func (o *ComputeBoardRelationship) SetBoardId(v int64)`

SetBoardId sets BoardId field to given value.

### HasBoardId

`func (o *ComputeBoardRelationship) HasBoardId() bool`

HasBoardId returns a boolean if a field has been set.

### GetCpuTypeController

`func (o *ComputeBoardRelationship) GetCpuTypeController() string`

GetCpuTypeController returns the CpuTypeController field if non-nil, zero value otherwise.

### GetCpuTypeControllerOk

`func (o *ComputeBoardRelationship) GetCpuTypeControllerOk() (*string, bool)`

GetCpuTypeControllerOk returns a tuple with the CpuTypeController field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuTypeController

`func (o *ComputeBoardRelationship) SetCpuTypeController(v string)`

SetCpuTypeController sets CpuTypeController field to given value.

### HasCpuTypeController

`func (o *ComputeBoardRelationship) HasCpuTypeController() bool`

HasCpuTypeController returns a boolean if a field has been set.

### GetOperPowerState

`func (o *ComputeBoardRelationship) GetOperPowerState() string`

GetOperPowerState returns the OperPowerState field if non-nil, zero value otherwise.

### GetOperPowerStateOk

`func (o *ComputeBoardRelationship) GetOperPowerStateOk() (*string, bool)`

GetOperPowerStateOk returns a tuple with the OperPowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperPowerState

`func (o *ComputeBoardRelationship) SetOperPowerState(v string)`

SetOperPowerState sets OperPowerState field to given value.

### HasOperPowerState

`func (o *ComputeBoardRelationship) HasOperPowerState() bool`

HasOperPowerState returns a boolean if a field has been set.

### GetPresence

`func (o *ComputeBoardRelationship) GetPresence() string`

GetPresence returns the Presence field if non-nil, zero value otherwise.

### GetPresenceOk

`func (o *ComputeBoardRelationship) GetPresenceOk() (*string, bool)`

GetPresenceOk returns a tuple with the Presence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresence

`func (o *ComputeBoardRelationship) SetPresence(v string)`

SetPresence sets Presence field to given value.

### HasPresence

`func (o *ComputeBoardRelationship) HasPresence() bool`

HasPresence returns a boolean if a field has been set.

### GetComputeBlade

`func (o *ComputeBoardRelationship) GetComputeBlade() ComputeBladeRelationship`

GetComputeBlade returns the ComputeBlade field if non-nil, zero value otherwise.

### GetComputeBladeOk

`func (o *ComputeBoardRelationship) GetComputeBladeOk() (*ComputeBladeRelationship, bool)`

GetComputeBladeOk returns a tuple with the ComputeBlade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeBlade

`func (o *ComputeBoardRelationship) SetComputeBlade(v ComputeBladeRelationship)`

SetComputeBlade sets ComputeBlade field to given value.

### HasComputeBlade

`func (o *ComputeBoardRelationship) HasComputeBlade() bool`

HasComputeBlade returns a boolean if a field has been set.

### GetComputeRackUnit

`func (o *ComputeBoardRelationship) GetComputeRackUnit() ComputeRackUnitRelationship`

GetComputeRackUnit returns the ComputeRackUnit field if non-nil, zero value otherwise.

### GetComputeRackUnitOk

`func (o *ComputeBoardRelationship) GetComputeRackUnitOk() (*ComputeRackUnitRelationship, bool)`

GetComputeRackUnitOk returns a tuple with the ComputeRackUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRackUnit

`func (o *ComputeBoardRelationship) SetComputeRackUnit(v ComputeRackUnitRelationship)`

SetComputeRackUnit sets ComputeRackUnit field to given value.

### HasComputeRackUnit

`func (o *ComputeBoardRelationship) HasComputeRackUnit() bool`

HasComputeRackUnit returns a boolean if a field has been set.

### GetEquipmentTpms

`func (o *ComputeBoardRelationship) GetEquipmentTpms() []EquipmentTpmRelationship`

GetEquipmentTpms returns the EquipmentTpms field if non-nil, zero value otherwise.

### GetEquipmentTpmsOk

`func (o *ComputeBoardRelationship) GetEquipmentTpmsOk() (*[]EquipmentTpmRelationship, bool)`

GetEquipmentTpmsOk returns a tuple with the EquipmentTpms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentTpms

`func (o *ComputeBoardRelationship) SetEquipmentTpms(v []EquipmentTpmRelationship)`

SetEquipmentTpms sets EquipmentTpms field to given value.

### HasEquipmentTpms

`func (o *ComputeBoardRelationship) HasEquipmentTpms() bool`

HasEquipmentTpms returns a boolean if a field has been set.

### SetEquipmentTpmsNil

`func (o *ComputeBoardRelationship) SetEquipmentTpmsNil(b bool)`

 SetEquipmentTpmsNil sets the value for EquipmentTpms to be an explicit nil

### UnsetEquipmentTpms
`func (o *ComputeBoardRelationship) UnsetEquipmentTpms()`

UnsetEquipmentTpms ensures that no value is present for EquipmentTpms, not even an explicit nil
### GetGraphicsCards

`func (o *ComputeBoardRelationship) GetGraphicsCards() []GraphicsCardRelationship`

GetGraphicsCards returns the GraphicsCards field if non-nil, zero value otherwise.

### GetGraphicsCardsOk

`func (o *ComputeBoardRelationship) GetGraphicsCardsOk() (*[]GraphicsCardRelationship, bool)`

GetGraphicsCardsOk returns a tuple with the GraphicsCards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraphicsCards

`func (o *ComputeBoardRelationship) SetGraphicsCards(v []GraphicsCardRelationship)`

SetGraphicsCards sets GraphicsCards field to given value.

### HasGraphicsCards

`func (o *ComputeBoardRelationship) HasGraphicsCards() bool`

HasGraphicsCards returns a boolean if a field has been set.

### SetGraphicsCardsNil

`func (o *ComputeBoardRelationship) SetGraphicsCardsNil(b bool)`

 SetGraphicsCardsNil sets the value for GraphicsCards to be an explicit nil

### UnsetGraphicsCards
`func (o *ComputeBoardRelationship) UnsetGraphicsCards()`

UnsetGraphicsCards ensures that no value is present for GraphicsCards, not even an explicit nil
### GetInventoryDeviceInfo

`func (o *ComputeBoardRelationship) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *ComputeBoardRelationship) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *ComputeBoardRelationship) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *ComputeBoardRelationship) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetMemoryArrays

`func (o *ComputeBoardRelationship) GetMemoryArrays() []MemoryArrayRelationship`

GetMemoryArrays returns the MemoryArrays field if non-nil, zero value otherwise.

### GetMemoryArraysOk

`func (o *ComputeBoardRelationship) GetMemoryArraysOk() (*[]MemoryArrayRelationship, bool)`

GetMemoryArraysOk returns a tuple with the MemoryArrays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryArrays

`func (o *ComputeBoardRelationship) SetMemoryArrays(v []MemoryArrayRelationship)`

SetMemoryArrays sets MemoryArrays field to given value.

### HasMemoryArrays

`func (o *ComputeBoardRelationship) HasMemoryArrays() bool`

HasMemoryArrays returns a boolean if a field has been set.

### SetMemoryArraysNil

`func (o *ComputeBoardRelationship) SetMemoryArraysNil(b bool)`

 SetMemoryArraysNil sets the value for MemoryArrays to be an explicit nil

### UnsetMemoryArrays
`func (o *ComputeBoardRelationship) UnsetMemoryArrays()`

UnsetMemoryArrays ensures that no value is present for MemoryArrays, not even an explicit nil
### GetPciCoprocessorCards

`func (o *ComputeBoardRelationship) GetPciCoprocessorCards() []PciCoprocessorCardRelationship`

GetPciCoprocessorCards returns the PciCoprocessorCards field if non-nil, zero value otherwise.

### GetPciCoprocessorCardsOk

`func (o *ComputeBoardRelationship) GetPciCoprocessorCardsOk() (*[]PciCoprocessorCardRelationship, bool)`

GetPciCoprocessorCardsOk returns a tuple with the PciCoprocessorCards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciCoprocessorCards

`func (o *ComputeBoardRelationship) SetPciCoprocessorCards(v []PciCoprocessorCardRelationship)`

SetPciCoprocessorCards sets PciCoprocessorCards field to given value.

### HasPciCoprocessorCards

`func (o *ComputeBoardRelationship) HasPciCoprocessorCards() bool`

HasPciCoprocessorCards returns a boolean if a field has been set.

### SetPciCoprocessorCardsNil

`func (o *ComputeBoardRelationship) SetPciCoprocessorCardsNil(b bool)`

 SetPciCoprocessorCardsNil sets the value for PciCoprocessorCards to be an explicit nil

### UnsetPciCoprocessorCards
`func (o *ComputeBoardRelationship) UnsetPciCoprocessorCards()`

UnsetPciCoprocessorCards ensures that no value is present for PciCoprocessorCards, not even an explicit nil
### GetPciSwitch

`func (o *ComputeBoardRelationship) GetPciSwitch() []PciSwitchRelationship`

GetPciSwitch returns the PciSwitch field if non-nil, zero value otherwise.

### GetPciSwitchOk

`func (o *ComputeBoardRelationship) GetPciSwitchOk() (*[]PciSwitchRelationship, bool)`

GetPciSwitchOk returns a tuple with the PciSwitch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciSwitch

`func (o *ComputeBoardRelationship) SetPciSwitch(v []PciSwitchRelationship)`

SetPciSwitch sets PciSwitch field to given value.

### HasPciSwitch

`func (o *ComputeBoardRelationship) HasPciSwitch() bool`

HasPciSwitch returns a boolean if a field has been set.

### SetPciSwitchNil

`func (o *ComputeBoardRelationship) SetPciSwitchNil(b bool)`

 SetPciSwitchNil sets the value for PciSwitch to be an explicit nil

### UnsetPciSwitch
`func (o *ComputeBoardRelationship) UnsetPciSwitch()`

UnsetPciSwitch ensures that no value is present for PciSwitch, not even an explicit nil
### GetPersistentMemoryConfiguration

`func (o *ComputeBoardRelationship) GetPersistentMemoryConfiguration() MemoryPersistentMemoryConfigurationRelationship`

GetPersistentMemoryConfiguration returns the PersistentMemoryConfiguration field if non-nil, zero value otherwise.

### GetPersistentMemoryConfigurationOk

`func (o *ComputeBoardRelationship) GetPersistentMemoryConfigurationOk() (*MemoryPersistentMemoryConfigurationRelationship, bool)`

GetPersistentMemoryConfigurationOk returns a tuple with the PersistentMemoryConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentMemoryConfiguration

`func (o *ComputeBoardRelationship) SetPersistentMemoryConfiguration(v MemoryPersistentMemoryConfigurationRelationship)`

SetPersistentMemoryConfiguration sets PersistentMemoryConfiguration field to given value.

### HasPersistentMemoryConfiguration

`func (o *ComputeBoardRelationship) HasPersistentMemoryConfiguration() bool`

HasPersistentMemoryConfiguration returns a boolean if a field has been set.

### GetProcessors

`func (o *ComputeBoardRelationship) GetProcessors() []ProcessorUnitRelationship`

GetProcessors returns the Processors field if non-nil, zero value otherwise.

### GetProcessorsOk

`func (o *ComputeBoardRelationship) GetProcessorsOk() (*[]ProcessorUnitRelationship, bool)`

GetProcessorsOk returns a tuple with the Processors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessors

`func (o *ComputeBoardRelationship) SetProcessors(v []ProcessorUnitRelationship)`

SetProcessors sets Processors field to given value.

### HasProcessors

`func (o *ComputeBoardRelationship) HasProcessors() bool`

HasProcessors returns a boolean if a field has been set.

### SetProcessorsNil

`func (o *ComputeBoardRelationship) SetProcessorsNil(b bool)`

 SetProcessorsNil sets the value for Processors to be an explicit nil

### UnsetProcessors
`func (o *ComputeBoardRelationship) UnsetProcessors()`

UnsetProcessors ensures that no value is present for Processors, not even an explicit nil
### GetRegisteredDevice

`func (o *ComputeBoardRelationship) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *ComputeBoardRelationship) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *ComputeBoardRelationship) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *ComputeBoardRelationship) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### GetSecurityUnits

`func (o *ComputeBoardRelationship) GetSecurityUnits() []SecurityUnitRelationship`

GetSecurityUnits returns the SecurityUnits field if non-nil, zero value otherwise.

### GetSecurityUnitsOk

`func (o *ComputeBoardRelationship) GetSecurityUnitsOk() (*[]SecurityUnitRelationship, bool)`

GetSecurityUnitsOk returns a tuple with the SecurityUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityUnits

`func (o *ComputeBoardRelationship) SetSecurityUnits(v []SecurityUnitRelationship)`

SetSecurityUnits sets SecurityUnits field to given value.

### HasSecurityUnits

`func (o *ComputeBoardRelationship) HasSecurityUnits() bool`

HasSecurityUnits returns a boolean if a field has been set.

### SetSecurityUnitsNil

`func (o *ComputeBoardRelationship) SetSecurityUnitsNil(b bool)`

 SetSecurityUnitsNil sets the value for SecurityUnits to be an explicit nil

### UnsetSecurityUnits
`func (o *ComputeBoardRelationship) UnsetSecurityUnits()`

UnsetSecurityUnits ensures that no value is present for SecurityUnits, not even an explicit nil
### GetStorageControllers

`func (o *ComputeBoardRelationship) GetStorageControllers() []StorageControllerRelationship`

GetStorageControllers returns the StorageControllers field if non-nil, zero value otherwise.

### GetStorageControllersOk

`func (o *ComputeBoardRelationship) GetStorageControllersOk() (*[]StorageControllerRelationship, bool)`

GetStorageControllersOk returns a tuple with the StorageControllers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageControllers

`func (o *ComputeBoardRelationship) SetStorageControllers(v []StorageControllerRelationship)`

SetStorageControllers sets StorageControllers field to given value.

### HasStorageControllers

`func (o *ComputeBoardRelationship) HasStorageControllers() bool`

HasStorageControllers returns a boolean if a field has been set.

### SetStorageControllersNil

`func (o *ComputeBoardRelationship) SetStorageControllersNil(b bool)`

 SetStorageControllersNil sets the value for StorageControllers to be an explicit nil

### UnsetStorageControllers
`func (o *ComputeBoardRelationship) UnsetStorageControllers()`

UnsetStorageControllers ensures that no value is present for StorageControllers, not even an explicit nil
### GetStorageFlexFlashControllers

`func (o *ComputeBoardRelationship) GetStorageFlexFlashControllers() []StorageFlexFlashControllerRelationship`

GetStorageFlexFlashControllers returns the StorageFlexFlashControllers field if non-nil, zero value otherwise.

### GetStorageFlexFlashControllersOk

`func (o *ComputeBoardRelationship) GetStorageFlexFlashControllersOk() (*[]StorageFlexFlashControllerRelationship, bool)`

GetStorageFlexFlashControllersOk returns a tuple with the StorageFlexFlashControllers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageFlexFlashControllers

`func (o *ComputeBoardRelationship) SetStorageFlexFlashControllers(v []StorageFlexFlashControllerRelationship)`

SetStorageFlexFlashControllers sets StorageFlexFlashControllers field to given value.

### HasStorageFlexFlashControllers

`func (o *ComputeBoardRelationship) HasStorageFlexFlashControllers() bool`

HasStorageFlexFlashControllers returns a boolean if a field has been set.

### SetStorageFlexFlashControllersNil

`func (o *ComputeBoardRelationship) SetStorageFlexFlashControllersNil(b bool)`

 SetStorageFlexFlashControllersNil sets the value for StorageFlexFlashControllers to be an explicit nil

### UnsetStorageFlexFlashControllers
`func (o *ComputeBoardRelationship) UnsetStorageFlexFlashControllers()`

UnsetStorageFlexFlashControllers ensures that no value is present for StorageFlexFlashControllers, not even an explicit nil
### GetStorageFlexUtilControllers

`func (o *ComputeBoardRelationship) GetStorageFlexUtilControllers() []StorageFlexUtilControllerRelationship`

GetStorageFlexUtilControllers returns the StorageFlexUtilControllers field if non-nil, zero value otherwise.

### GetStorageFlexUtilControllersOk

`func (o *ComputeBoardRelationship) GetStorageFlexUtilControllersOk() (*[]StorageFlexUtilControllerRelationship, bool)`

GetStorageFlexUtilControllersOk returns a tuple with the StorageFlexUtilControllers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageFlexUtilControllers

`func (o *ComputeBoardRelationship) SetStorageFlexUtilControllers(v []StorageFlexUtilControllerRelationship)`

SetStorageFlexUtilControllers sets StorageFlexUtilControllers field to given value.

### HasStorageFlexUtilControllers

`func (o *ComputeBoardRelationship) HasStorageFlexUtilControllers() bool`

HasStorageFlexUtilControllers returns a boolean if a field has been set.

### SetStorageFlexUtilControllersNil

`func (o *ComputeBoardRelationship) SetStorageFlexUtilControllersNil(b bool)`

 SetStorageFlexUtilControllersNil sets the value for StorageFlexUtilControllers to be an explicit nil

### UnsetStorageFlexUtilControllers
`func (o *ComputeBoardRelationship) UnsetStorageFlexUtilControllers()`

UnsetStorageFlexUtilControllers ensures that no value is present for StorageFlexUtilControllers, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



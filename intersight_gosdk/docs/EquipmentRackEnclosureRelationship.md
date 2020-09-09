# EquipmentRackEnclosureRelationship

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
**EnclosureId** | Pointer to **int64** | This represents the Enclosure Identifier for Rack servers. | [optional] [readonly] 
**Fanmodules** | Pointer to [**[]EquipmentFanModuleRelationship**](equipment.FanModule.Relationship.md) | An array of relationships to equipmentFanModule resources. | [optional] [readonly] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](inventory.DeviceInfo.Relationship.md) |  | [optional] 
**Psus** | Pointer to [**[]EquipmentPsuRelationship**](equipment.Psu.Relationship.md) | An array of relationships to equipmentPsu resources. | [optional] [readonly] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 
**Slots** | Pointer to [**[]EquipmentRackEnclosureSlotRelationship**](equipment.RackEnclosureSlot.Relationship.md) | An array of relationships to equipmentRackEnclosureSlot resources. | [optional] [readonly] 

## Methods

### NewEquipmentRackEnclosureRelationship

`func NewEquipmentRackEnclosureRelationship(classId string, objectType string, ) *EquipmentRackEnclosureRelationship`

NewEquipmentRackEnclosureRelationship instantiates a new EquipmentRackEnclosureRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquipmentRackEnclosureRelationshipWithDefaults

`func NewEquipmentRackEnclosureRelationshipWithDefaults() *EquipmentRackEnclosureRelationship`

NewEquipmentRackEnclosureRelationshipWithDefaults instantiates a new EquipmentRackEnclosureRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountMoid

`func (o *EquipmentRackEnclosureRelationship) GetAccountMoid() string`

GetAccountMoid returns the AccountMoid field if non-nil, zero value otherwise.

### GetAccountMoidOk

`func (o *EquipmentRackEnclosureRelationship) GetAccountMoidOk() (*string, bool)`

GetAccountMoidOk returns a tuple with the AccountMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMoid

`func (o *EquipmentRackEnclosureRelationship) SetAccountMoid(v string)`

SetAccountMoid sets AccountMoid field to given value.

### HasAccountMoid

`func (o *EquipmentRackEnclosureRelationship) HasAccountMoid() bool`

HasAccountMoid returns a boolean if a field has been set.

### GetClassId

`func (o *EquipmentRackEnclosureRelationship) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *EquipmentRackEnclosureRelationship) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *EquipmentRackEnclosureRelationship) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetCreateTime

`func (o *EquipmentRackEnclosureRelationship) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *EquipmentRackEnclosureRelationship) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *EquipmentRackEnclosureRelationship) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *EquipmentRackEnclosureRelationship) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDomainGroupMoid

`func (o *EquipmentRackEnclosureRelationship) GetDomainGroupMoid() string`

GetDomainGroupMoid returns the DomainGroupMoid field if non-nil, zero value otherwise.

### GetDomainGroupMoidOk

`func (o *EquipmentRackEnclosureRelationship) GetDomainGroupMoidOk() (*string, bool)`

GetDomainGroupMoidOk returns a tuple with the DomainGroupMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGroupMoid

`func (o *EquipmentRackEnclosureRelationship) SetDomainGroupMoid(v string)`

SetDomainGroupMoid sets DomainGroupMoid field to given value.

### HasDomainGroupMoid

`func (o *EquipmentRackEnclosureRelationship) HasDomainGroupMoid() bool`

HasDomainGroupMoid returns a boolean if a field has been set.

### GetModTime

`func (o *EquipmentRackEnclosureRelationship) GetModTime() time.Time`

GetModTime returns the ModTime field if non-nil, zero value otherwise.

### GetModTimeOk

`func (o *EquipmentRackEnclosureRelationship) GetModTimeOk() (*time.Time, bool)`

GetModTimeOk returns a tuple with the ModTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModTime

`func (o *EquipmentRackEnclosureRelationship) SetModTime(v time.Time)`

SetModTime sets ModTime field to given value.

### HasModTime

`func (o *EquipmentRackEnclosureRelationship) HasModTime() bool`

HasModTime returns a boolean if a field has been set.

### GetMoid

`func (o *EquipmentRackEnclosureRelationship) GetMoid() string`

GetMoid returns the Moid field if non-nil, zero value otherwise.

### GetMoidOk

`func (o *EquipmentRackEnclosureRelationship) GetMoidOk() (*string, bool)`

GetMoidOk returns a tuple with the Moid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoid

`func (o *EquipmentRackEnclosureRelationship) SetMoid(v string)`

SetMoid sets Moid field to given value.

### HasMoid

`func (o *EquipmentRackEnclosureRelationship) HasMoid() bool`

HasMoid returns a boolean if a field has been set.

### GetObjectType

`func (o *EquipmentRackEnclosureRelationship) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *EquipmentRackEnclosureRelationship) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *EquipmentRackEnclosureRelationship) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOwners

`func (o *EquipmentRackEnclosureRelationship) GetOwners() []string`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *EquipmentRackEnclosureRelationship) GetOwnersOk() (*[]string, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *EquipmentRackEnclosureRelationship) SetOwners(v []string)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *EquipmentRackEnclosureRelationship) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetSharedScope

`func (o *EquipmentRackEnclosureRelationship) GetSharedScope() string`

GetSharedScope returns the SharedScope field if non-nil, zero value otherwise.

### GetSharedScopeOk

`func (o *EquipmentRackEnclosureRelationship) GetSharedScopeOk() (*string, bool)`

GetSharedScopeOk returns a tuple with the SharedScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedScope

`func (o *EquipmentRackEnclosureRelationship) SetSharedScope(v string)`

SetSharedScope sets SharedScope field to given value.

### HasSharedScope

`func (o *EquipmentRackEnclosureRelationship) HasSharedScope() bool`

HasSharedScope returns a boolean if a field has been set.

### GetTags

`func (o *EquipmentRackEnclosureRelationship) GetTags() []MoTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *EquipmentRackEnclosureRelationship) GetTagsOk() (*[]MoTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *EquipmentRackEnclosureRelationship) SetTags(v []MoTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *EquipmentRackEnclosureRelationship) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersionContext

`func (o *EquipmentRackEnclosureRelationship) GetVersionContext() MoVersionContext`

GetVersionContext returns the VersionContext field if non-nil, zero value otherwise.

### GetVersionContextOk

`func (o *EquipmentRackEnclosureRelationship) GetVersionContextOk() (*MoVersionContext, bool)`

GetVersionContextOk returns a tuple with the VersionContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionContext

`func (o *EquipmentRackEnclosureRelationship) SetVersionContext(v MoVersionContext)`

SetVersionContext sets VersionContext field to given value.

### HasVersionContext

`func (o *EquipmentRackEnclosureRelationship) HasVersionContext() bool`

HasVersionContext returns a boolean if a field has been set.

### GetAncestors

`func (o *EquipmentRackEnclosureRelationship) GetAncestors() []MoBaseMoRelationship`

GetAncestors returns the Ancestors field if non-nil, zero value otherwise.

### GetAncestorsOk

`func (o *EquipmentRackEnclosureRelationship) GetAncestorsOk() (*[]MoBaseMoRelationship, bool)`

GetAncestorsOk returns a tuple with the Ancestors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestors

`func (o *EquipmentRackEnclosureRelationship) SetAncestors(v []MoBaseMoRelationship)`

SetAncestors sets Ancestors field to given value.

### HasAncestors

`func (o *EquipmentRackEnclosureRelationship) HasAncestors() bool`

HasAncestors returns a boolean if a field has been set.

### SetAncestorsNil

`func (o *EquipmentRackEnclosureRelationship) SetAncestorsNil(b bool)`

 SetAncestorsNil sets the value for Ancestors to be an explicit nil

### UnsetAncestors
`func (o *EquipmentRackEnclosureRelationship) UnsetAncestors()`

UnsetAncestors ensures that no value is present for Ancestors, not even an explicit nil
### GetParent

`func (o *EquipmentRackEnclosureRelationship) GetParent() MoBaseMoRelationship`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *EquipmentRackEnclosureRelationship) GetParentOk() (*MoBaseMoRelationship, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *EquipmentRackEnclosureRelationship) SetParent(v MoBaseMoRelationship)`

SetParent sets Parent field to given value.

### HasParent

`func (o *EquipmentRackEnclosureRelationship) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPermissionResources

`func (o *EquipmentRackEnclosureRelationship) GetPermissionResources() []MoBaseMoRelationship`

GetPermissionResources returns the PermissionResources field if non-nil, zero value otherwise.

### GetPermissionResourcesOk

`func (o *EquipmentRackEnclosureRelationship) GetPermissionResourcesOk() (*[]MoBaseMoRelationship, bool)`

GetPermissionResourcesOk returns a tuple with the PermissionResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionResources

`func (o *EquipmentRackEnclosureRelationship) SetPermissionResources(v []MoBaseMoRelationship)`

SetPermissionResources sets PermissionResources field to given value.

### HasPermissionResources

`func (o *EquipmentRackEnclosureRelationship) HasPermissionResources() bool`

HasPermissionResources returns a boolean if a field has been set.

### SetPermissionResourcesNil

`func (o *EquipmentRackEnclosureRelationship) SetPermissionResourcesNil(b bool)`

 SetPermissionResourcesNil sets the value for PermissionResources to be an explicit nil

### UnsetPermissionResources
`func (o *EquipmentRackEnclosureRelationship) UnsetPermissionResources()`

UnsetPermissionResources ensures that no value is present for PermissionResources, not even an explicit nil
### GetDisplayNames

`func (o *EquipmentRackEnclosureRelationship) GetDisplayNames() map[string][]string`

GetDisplayNames returns the DisplayNames field if non-nil, zero value otherwise.

### GetDisplayNamesOk

`func (o *EquipmentRackEnclosureRelationship) GetDisplayNamesOk() (*map[string][]string, bool)`

GetDisplayNamesOk returns a tuple with the DisplayNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNames

`func (o *EquipmentRackEnclosureRelationship) SetDisplayNames(v map[string][]string)`

SetDisplayNames sets DisplayNames field to given value.

### HasDisplayNames

`func (o *EquipmentRackEnclosureRelationship) HasDisplayNames() bool`

HasDisplayNames returns a boolean if a field has been set.

### SetDisplayNamesNil

`func (o *EquipmentRackEnclosureRelationship) SetDisplayNamesNil(b bool)`

 SetDisplayNamesNil sets the value for DisplayNames to be an explicit nil

### UnsetDisplayNames
`func (o *EquipmentRackEnclosureRelationship) UnsetDisplayNames()`

UnsetDisplayNames ensures that no value is present for DisplayNames, not even an explicit nil
### GetDeviceMoId

`func (o *EquipmentRackEnclosureRelationship) GetDeviceMoId() string`

GetDeviceMoId returns the DeviceMoId field if non-nil, zero value otherwise.

### GetDeviceMoIdOk

`func (o *EquipmentRackEnclosureRelationship) GetDeviceMoIdOk() (*string, bool)`

GetDeviceMoIdOk returns a tuple with the DeviceMoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMoId

`func (o *EquipmentRackEnclosureRelationship) SetDeviceMoId(v string)`

SetDeviceMoId sets DeviceMoId field to given value.

### HasDeviceMoId

`func (o *EquipmentRackEnclosureRelationship) HasDeviceMoId() bool`

HasDeviceMoId returns a boolean if a field has been set.

### GetDn

`func (o *EquipmentRackEnclosureRelationship) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *EquipmentRackEnclosureRelationship) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *EquipmentRackEnclosureRelationship) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *EquipmentRackEnclosureRelationship) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetRn

`func (o *EquipmentRackEnclosureRelationship) GetRn() string`

GetRn returns the Rn field if non-nil, zero value otherwise.

### GetRnOk

`func (o *EquipmentRackEnclosureRelationship) GetRnOk() (*string, bool)`

GetRnOk returns a tuple with the Rn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRn

`func (o *EquipmentRackEnclosureRelationship) SetRn(v string)`

SetRn sets Rn field to given value.

### HasRn

`func (o *EquipmentRackEnclosureRelationship) HasRn() bool`

HasRn returns a boolean if a field has been set.

### GetModel

`func (o *EquipmentRackEnclosureRelationship) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *EquipmentRackEnclosureRelationship) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *EquipmentRackEnclosureRelationship) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *EquipmentRackEnclosureRelationship) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetRevision

`func (o *EquipmentRackEnclosureRelationship) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *EquipmentRackEnclosureRelationship) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *EquipmentRackEnclosureRelationship) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *EquipmentRackEnclosureRelationship) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetSerial

`func (o *EquipmentRackEnclosureRelationship) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *EquipmentRackEnclosureRelationship) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *EquipmentRackEnclosureRelationship) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *EquipmentRackEnclosureRelationship) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetVendor

`func (o *EquipmentRackEnclosureRelationship) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *EquipmentRackEnclosureRelationship) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *EquipmentRackEnclosureRelationship) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *EquipmentRackEnclosureRelationship) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetEnclosureId

`func (o *EquipmentRackEnclosureRelationship) GetEnclosureId() int64`

GetEnclosureId returns the EnclosureId field if non-nil, zero value otherwise.

### GetEnclosureIdOk

`func (o *EquipmentRackEnclosureRelationship) GetEnclosureIdOk() (*int64, bool)`

GetEnclosureIdOk returns a tuple with the EnclosureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnclosureId

`func (o *EquipmentRackEnclosureRelationship) SetEnclosureId(v int64)`

SetEnclosureId sets EnclosureId field to given value.

### HasEnclosureId

`func (o *EquipmentRackEnclosureRelationship) HasEnclosureId() bool`

HasEnclosureId returns a boolean if a field has been set.

### GetFanmodules

`func (o *EquipmentRackEnclosureRelationship) GetFanmodules() []EquipmentFanModuleRelationship`

GetFanmodules returns the Fanmodules field if non-nil, zero value otherwise.

### GetFanmodulesOk

`func (o *EquipmentRackEnclosureRelationship) GetFanmodulesOk() (*[]EquipmentFanModuleRelationship, bool)`

GetFanmodulesOk returns a tuple with the Fanmodules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFanmodules

`func (o *EquipmentRackEnclosureRelationship) SetFanmodules(v []EquipmentFanModuleRelationship)`

SetFanmodules sets Fanmodules field to given value.

### HasFanmodules

`func (o *EquipmentRackEnclosureRelationship) HasFanmodules() bool`

HasFanmodules returns a boolean if a field has been set.

### SetFanmodulesNil

`func (o *EquipmentRackEnclosureRelationship) SetFanmodulesNil(b bool)`

 SetFanmodulesNil sets the value for Fanmodules to be an explicit nil

### UnsetFanmodules
`func (o *EquipmentRackEnclosureRelationship) UnsetFanmodules()`

UnsetFanmodules ensures that no value is present for Fanmodules, not even an explicit nil
### GetInventoryDeviceInfo

`func (o *EquipmentRackEnclosureRelationship) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *EquipmentRackEnclosureRelationship) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *EquipmentRackEnclosureRelationship) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *EquipmentRackEnclosureRelationship) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetPsus

`func (o *EquipmentRackEnclosureRelationship) GetPsus() []EquipmentPsuRelationship`

GetPsus returns the Psus field if non-nil, zero value otherwise.

### GetPsusOk

`func (o *EquipmentRackEnclosureRelationship) GetPsusOk() (*[]EquipmentPsuRelationship, bool)`

GetPsusOk returns a tuple with the Psus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPsus

`func (o *EquipmentRackEnclosureRelationship) SetPsus(v []EquipmentPsuRelationship)`

SetPsus sets Psus field to given value.

### HasPsus

`func (o *EquipmentRackEnclosureRelationship) HasPsus() bool`

HasPsus returns a boolean if a field has been set.

### SetPsusNil

`func (o *EquipmentRackEnclosureRelationship) SetPsusNil(b bool)`

 SetPsusNil sets the value for Psus to be an explicit nil

### UnsetPsus
`func (o *EquipmentRackEnclosureRelationship) UnsetPsus()`

UnsetPsus ensures that no value is present for Psus, not even an explicit nil
### GetRegisteredDevice

`func (o *EquipmentRackEnclosureRelationship) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *EquipmentRackEnclosureRelationship) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *EquipmentRackEnclosureRelationship) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *EquipmentRackEnclosureRelationship) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### GetSlots

`func (o *EquipmentRackEnclosureRelationship) GetSlots() []EquipmentRackEnclosureSlotRelationship`

GetSlots returns the Slots field if non-nil, zero value otherwise.

### GetSlotsOk

`func (o *EquipmentRackEnclosureRelationship) GetSlotsOk() (*[]EquipmentRackEnclosureSlotRelationship, bool)`

GetSlotsOk returns a tuple with the Slots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlots

`func (o *EquipmentRackEnclosureRelationship) SetSlots(v []EquipmentRackEnclosureSlotRelationship)`

SetSlots sets Slots field to given value.

### HasSlots

`func (o *EquipmentRackEnclosureRelationship) HasSlots() bool`

HasSlots returns a boolean if a field has been set.

### SetSlotsNil

`func (o *EquipmentRackEnclosureRelationship) SetSlotsNil(b bool)`

 SetSlotsNil sets the value for Slots to be an explicit nil

### UnsetSlots
`func (o *EquipmentRackEnclosureRelationship) UnsetSlots()`

UnsetSlots ensures that no value is present for Slots, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



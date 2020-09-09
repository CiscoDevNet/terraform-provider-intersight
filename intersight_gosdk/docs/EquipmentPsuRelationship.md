# EquipmentPsuRelationship

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
**Description** | Pointer to **string** | This field is to provide description for the power supply unit. | [optional] [readonly] 
**OperState** | Pointer to **string** | This field identifies the psu operational state. | [optional] [readonly] 
**PartNumber** | Pointer to **string** | This field identifies the Part Number for this Power Supply Unit. | [optional] [readonly] 
**Pid** | Pointer to **string** | This field identifies the Product ID for the Power Supply. | [optional] [readonly] 
**Presence** | Pointer to **string** | This field identifies the presence state of the psu. | [optional] [readonly] 
**PsuFwVersion** | Pointer to **string** | This field identifies the Firmware Version of the Power Supply. | [optional] [readonly] 
**PsuId** | Pointer to **int64** | This represents power supply unit identifier in chassis/server. | [optional] [readonly] 
**PsuInputSrc** | Pointer to **string** | This field identifies the input source for the Power Supply. | [optional] [readonly] 
**PsuType** | Pointer to **string** | This field identifies the type of the Power Supply. | [optional] [readonly] 
**PsuWattage** | Pointer to **string** | This field identifies the Wattage of the Power Supply. | [optional] [readonly] 
**Sku** | Pointer to **string** | This field identifies the Stockkeeping Unit for this Power Supply. | [optional] [readonly] 
**Vid** | Pointer to **string** | This field identifies the Vendor ID for this Power Supply Unit. | [optional] [readonly] 
**Voltage** | Pointer to **string** | This field is used to indicate the Voltage for this Power Supply. | [optional] [readonly] 
**ComputeRackUnit** | Pointer to [**ComputeRackUnitRelationship**](compute.RackUnit.Relationship.md) |  | [optional] 
**EquipmentChassis** | Pointer to [**EquipmentChassisRelationship**](equipment.Chassis.Relationship.md) |  | [optional] 
**EquipmentFex** | Pointer to [**EquipmentFexRelationship**](equipment.Fex.Relationship.md) |  | [optional] 
**EquipmentRackEnclosure** | Pointer to [**EquipmentRackEnclosureRelationship**](equipment.RackEnclosure.Relationship.md) |  | [optional] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](inventory.DeviceInfo.Relationship.md) |  | [optional] 
**NetworkElement** | Pointer to [**NetworkElementRelationship**](network.Element.Relationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewEquipmentPsuRelationship

`func NewEquipmentPsuRelationship(classId string, objectType string, ) *EquipmentPsuRelationship`

NewEquipmentPsuRelationship instantiates a new EquipmentPsuRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquipmentPsuRelationshipWithDefaults

`func NewEquipmentPsuRelationshipWithDefaults() *EquipmentPsuRelationship`

NewEquipmentPsuRelationshipWithDefaults instantiates a new EquipmentPsuRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountMoid

`func (o *EquipmentPsuRelationship) GetAccountMoid() string`

GetAccountMoid returns the AccountMoid field if non-nil, zero value otherwise.

### GetAccountMoidOk

`func (o *EquipmentPsuRelationship) GetAccountMoidOk() (*string, bool)`

GetAccountMoidOk returns a tuple with the AccountMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMoid

`func (o *EquipmentPsuRelationship) SetAccountMoid(v string)`

SetAccountMoid sets AccountMoid field to given value.

### HasAccountMoid

`func (o *EquipmentPsuRelationship) HasAccountMoid() bool`

HasAccountMoid returns a boolean if a field has been set.

### GetClassId

`func (o *EquipmentPsuRelationship) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *EquipmentPsuRelationship) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *EquipmentPsuRelationship) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetCreateTime

`func (o *EquipmentPsuRelationship) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *EquipmentPsuRelationship) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *EquipmentPsuRelationship) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *EquipmentPsuRelationship) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDomainGroupMoid

`func (o *EquipmentPsuRelationship) GetDomainGroupMoid() string`

GetDomainGroupMoid returns the DomainGroupMoid field if non-nil, zero value otherwise.

### GetDomainGroupMoidOk

`func (o *EquipmentPsuRelationship) GetDomainGroupMoidOk() (*string, bool)`

GetDomainGroupMoidOk returns a tuple with the DomainGroupMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGroupMoid

`func (o *EquipmentPsuRelationship) SetDomainGroupMoid(v string)`

SetDomainGroupMoid sets DomainGroupMoid field to given value.

### HasDomainGroupMoid

`func (o *EquipmentPsuRelationship) HasDomainGroupMoid() bool`

HasDomainGroupMoid returns a boolean if a field has been set.

### GetModTime

`func (o *EquipmentPsuRelationship) GetModTime() time.Time`

GetModTime returns the ModTime field if non-nil, zero value otherwise.

### GetModTimeOk

`func (o *EquipmentPsuRelationship) GetModTimeOk() (*time.Time, bool)`

GetModTimeOk returns a tuple with the ModTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModTime

`func (o *EquipmentPsuRelationship) SetModTime(v time.Time)`

SetModTime sets ModTime field to given value.

### HasModTime

`func (o *EquipmentPsuRelationship) HasModTime() bool`

HasModTime returns a boolean if a field has been set.

### GetMoid

`func (o *EquipmentPsuRelationship) GetMoid() string`

GetMoid returns the Moid field if non-nil, zero value otherwise.

### GetMoidOk

`func (o *EquipmentPsuRelationship) GetMoidOk() (*string, bool)`

GetMoidOk returns a tuple with the Moid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoid

`func (o *EquipmentPsuRelationship) SetMoid(v string)`

SetMoid sets Moid field to given value.

### HasMoid

`func (o *EquipmentPsuRelationship) HasMoid() bool`

HasMoid returns a boolean if a field has been set.

### GetObjectType

`func (o *EquipmentPsuRelationship) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *EquipmentPsuRelationship) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *EquipmentPsuRelationship) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOwners

`func (o *EquipmentPsuRelationship) GetOwners() []string`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *EquipmentPsuRelationship) GetOwnersOk() (*[]string, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *EquipmentPsuRelationship) SetOwners(v []string)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *EquipmentPsuRelationship) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetSharedScope

`func (o *EquipmentPsuRelationship) GetSharedScope() string`

GetSharedScope returns the SharedScope field if non-nil, zero value otherwise.

### GetSharedScopeOk

`func (o *EquipmentPsuRelationship) GetSharedScopeOk() (*string, bool)`

GetSharedScopeOk returns a tuple with the SharedScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedScope

`func (o *EquipmentPsuRelationship) SetSharedScope(v string)`

SetSharedScope sets SharedScope field to given value.

### HasSharedScope

`func (o *EquipmentPsuRelationship) HasSharedScope() bool`

HasSharedScope returns a boolean if a field has been set.

### GetTags

`func (o *EquipmentPsuRelationship) GetTags() []MoTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *EquipmentPsuRelationship) GetTagsOk() (*[]MoTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *EquipmentPsuRelationship) SetTags(v []MoTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *EquipmentPsuRelationship) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersionContext

`func (o *EquipmentPsuRelationship) GetVersionContext() MoVersionContext`

GetVersionContext returns the VersionContext field if non-nil, zero value otherwise.

### GetVersionContextOk

`func (o *EquipmentPsuRelationship) GetVersionContextOk() (*MoVersionContext, bool)`

GetVersionContextOk returns a tuple with the VersionContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionContext

`func (o *EquipmentPsuRelationship) SetVersionContext(v MoVersionContext)`

SetVersionContext sets VersionContext field to given value.

### HasVersionContext

`func (o *EquipmentPsuRelationship) HasVersionContext() bool`

HasVersionContext returns a boolean if a field has been set.

### GetAncestors

`func (o *EquipmentPsuRelationship) GetAncestors() []MoBaseMoRelationship`

GetAncestors returns the Ancestors field if non-nil, zero value otherwise.

### GetAncestorsOk

`func (o *EquipmentPsuRelationship) GetAncestorsOk() (*[]MoBaseMoRelationship, bool)`

GetAncestorsOk returns a tuple with the Ancestors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestors

`func (o *EquipmentPsuRelationship) SetAncestors(v []MoBaseMoRelationship)`

SetAncestors sets Ancestors field to given value.

### HasAncestors

`func (o *EquipmentPsuRelationship) HasAncestors() bool`

HasAncestors returns a boolean if a field has been set.

### SetAncestorsNil

`func (o *EquipmentPsuRelationship) SetAncestorsNil(b bool)`

 SetAncestorsNil sets the value for Ancestors to be an explicit nil

### UnsetAncestors
`func (o *EquipmentPsuRelationship) UnsetAncestors()`

UnsetAncestors ensures that no value is present for Ancestors, not even an explicit nil
### GetParent

`func (o *EquipmentPsuRelationship) GetParent() MoBaseMoRelationship`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *EquipmentPsuRelationship) GetParentOk() (*MoBaseMoRelationship, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *EquipmentPsuRelationship) SetParent(v MoBaseMoRelationship)`

SetParent sets Parent field to given value.

### HasParent

`func (o *EquipmentPsuRelationship) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPermissionResources

`func (o *EquipmentPsuRelationship) GetPermissionResources() []MoBaseMoRelationship`

GetPermissionResources returns the PermissionResources field if non-nil, zero value otherwise.

### GetPermissionResourcesOk

`func (o *EquipmentPsuRelationship) GetPermissionResourcesOk() (*[]MoBaseMoRelationship, bool)`

GetPermissionResourcesOk returns a tuple with the PermissionResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionResources

`func (o *EquipmentPsuRelationship) SetPermissionResources(v []MoBaseMoRelationship)`

SetPermissionResources sets PermissionResources field to given value.

### HasPermissionResources

`func (o *EquipmentPsuRelationship) HasPermissionResources() bool`

HasPermissionResources returns a boolean if a field has been set.

### SetPermissionResourcesNil

`func (o *EquipmentPsuRelationship) SetPermissionResourcesNil(b bool)`

 SetPermissionResourcesNil sets the value for PermissionResources to be an explicit nil

### UnsetPermissionResources
`func (o *EquipmentPsuRelationship) UnsetPermissionResources()`

UnsetPermissionResources ensures that no value is present for PermissionResources, not even an explicit nil
### GetDisplayNames

`func (o *EquipmentPsuRelationship) GetDisplayNames() map[string][]string`

GetDisplayNames returns the DisplayNames field if non-nil, zero value otherwise.

### GetDisplayNamesOk

`func (o *EquipmentPsuRelationship) GetDisplayNamesOk() (*map[string][]string, bool)`

GetDisplayNamesOk returns a tuple with the DisplayNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNames

`func (o *EquipmentPsuRelationship) SetDisplayNames(v map[string][]string)`

SetDisplayNames sets DisplayNames field to given value.

### HasDisplayNames

`func (o *EquipmentPsuRelationship) HasDisplayNames() bool`

HasDisplayNames returns a boolean if a field has been set.

### SetDisplayNamesNil

`func (o *EquipmentPsuRelationship) SetDisplayNamesNil(b bool)`

 SetDisplayNamesNil sets the value for DisplayNames to be an explicit nil

### UnsetDisplayNames
`func (o *EquipmentPsuRelationship) UnsetDisplayNames()`

UnsetDisplayNames ensures that no value is present for DisplayNames, not even an explicit nil
### GetDeviceMoId

`func (o *EquipmentPsuRelationship) GetDeviceMoId() string`

GetDeviceMoId returns the DeviceMoId field if non-nil, zero value otherwise.

### GetDeviceMoIdOk

`func (o *EquipmentPsuRelationship) GetDeviceMoIdOk() (*string, bool)`

GetDeviceMoIdOk returns a tuple with the DeviceMoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMoId

`func (o *EquipmentPsuRelationship) SetDeviceMoId(v string)`

SetDeviceMoId sets DeviceMoId field to given value.

### HasDeviceMoId

`func (o *EquipmentPsuRelationship) HasDeviceMoId() bool`

HasDeviceMoId returns a boolean if a field has been set.

### GetDn

`func (o *EquipmentPsuRelationship) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *EquipmentPsuRelationship) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *EquipmentPsuRelationship) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *EquipmentPsuRelationship) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetRn

`func (o *EquipmentPsuRelationship) GetRn() string`

GetRn returns the Rn field if non-nil, zero value otherwise.

### GetRnOk

`func (o *EquipmentPsuRelationship) GetRnOk() (*string, bool)`

GetRnOk returns a tuple with the Rn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRn

`func (o *EquipmentPsuRelationship) SetRn(v string)`

SetRn sets Rn field to given value.

### HasRn

`func (o *EquipmentPsuRelationship) HasRn() bool`

HasRn returns a boolean if a field has been set.

### GetModel

`func (o *EquipmentPsuRelationship) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *EquipmentPsuRelationship) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *EquipmentPsuRelationship) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *EquipmentPsuRelationship) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetRevision

`func (o *EquipmentPsuRelationship) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *EquipmentPsuRelationship) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *EquipmentPsuRelationship) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *EquipmentPsuRelationship) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetSerial

`func (o *EquipmentPsuRelationship) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *EquipmentPsuRelationship) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *EquipmentPsuRelationship) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *EquipmentPsuRelationship) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetVendor

`func (o *EquipmentPsuRelationship) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *EquipmentPsuRelationship) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *EquipmentPsuRelationship) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *EquipmentPsuRelationship) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetDescription

`func (o *EquipmentPsuRelationship) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EquipmentPsuRelationship) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EquipmentPsuRelationship) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EquipmentPsuRelationship) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOperState

`func (o *EquipmentPsuRelationship) GetOperState() string`

GetOperState returns the OperState field if non-nil, zero value otherwise.

### GetOperStateOk

`func (o *EquipmentPsuRelationship) GetOperStateOk() (*string, bool)`

GetOperStateOk returns a tuple with the OperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperState

`func (o *EquipmentPsuRelationship) SetOperState(v string)`

SetOperState sets OperState field to given value.

### HasOperState

`func (o *EquipmentPsuRelationship) HasOperState() bool`

HasOperState returns a boolean if a field has been set.

### GetPartNumber

`func (o *EquipmentPsuRelationship) GetPartNumber() string`

GetPartNumber returns the PartNumber field if non-nil, zero value otherwise.

### GetPartNumberOk

`func (o *EquipmentPsuRelationship) GetPartNumberOk() (*string, bool)`

GetPartNumberOk returns a tuple with the PartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartNumber

`func (o *EquipmentPsuRelationship) SetPartNumber(v string)`

SetPartNumber sets PartNumber field to given value.

### HasPartNumber

`func (o *EquipmentPsuRelationship) HasPartNumber() bool`

HasPartNumber returns a boolean if a field has been set.

### GetPid

`func (o *EquipmentPsuRelationship) GetPid() string`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *EquipmentPsuRelationship) GetPidOk() (*string, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *EquipmentPsuRelationship) SetPid(v string)`

SetPid sets Pid field to given value.

### HasPid

`func (o *EquipmentPsuRelationship) HasPid() bool`

HasPid returns a boolean if a field has been set.

### GetPresence

`func (o *EquipmentPsuRelationship) GetPresence() string`

GetPresence returns the Presence field if non-nil, zero value otherwise.

### GetPresenceOk

`func (o *EquipmentPsuRelationship) GetPresenceOk() (*string, bool)`

GetPresenceOk returns a tuple with the Presence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresence

`func (o *EquipmentPsuRelationship) SetPresence(v string)`

SetPresence sets Presence field to given value.

### HasPresence

`func (o *EquipmentPsuRelationship) HasPresence() bool`

HasPresence returns a boolean if a field has been set.

### GetPsuFwVersion

`func (o *EquipmentPsuRelationship) GetPsuFwVersion() string`

GetPsuFwVersion returns the PsuFwVersion field if non-nil, zero value otherwise.

### GetPsuFwVersionOk

`func (o *EquipmentPsuRelationship) GetPsuFwVersionOk() (*string, bool)`

GetPsuFwVersionOk returns a tuple with the PsuFwVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPsuFwVersion

`func (o *EquipmentPsuRelationship) SetPsuFwVersion(v string)`

SetPsuFwVersion sets PsuFwVersion field to given value.

### HasPsuFwVersion

`func (o *EquipmentPsuRelationship) HasPsuFwVersion() bool`

HasPsuFwVersion returns a boolean if a field has been set.

### GetPsuId

`func (o *EquipmentPsuRelationship) GetPsuId() int64`

GetPsuId returns the PsuId field if non-nil, zero value otherwise.

### GetPsuIdOk

`func (o *EquipmentPsuRelationship) GetPsuIdOk() (*int64, bool)`

GetPsuIdOk returns a tuple with the PsuId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPsuId

`func (o *EquipmentPsuRelationship) SetPsuId(v int64)`

SetPsuId sets PsuId field to given value.

### HasPsuId

`func (o *EquipmentPsuRelationship) HasPsuId() bool`

HasPsuId returns a boolean if a field has been set.

### GetPsuInputSrc

`func (o *EquipmentPsuRelationship) GetPsuInputSrc() string`

GetPsuInputSrc returns the PsuInputSrc field if non-nil, zero value otherwise.

### GetPsuInputSrcOk

`func (o *EquipmentPsuRelationship) GetPsuInputSrcOk() (*string, bool)`

GetPsuInputSrcOk returns a tuple with the PsuInputSrc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPsuInputSrc

`func (o *EquipmentPsuRelationship) SetPsuInputSrc(v string)`

SetPsuInputSrc sets PsuInputSrc field to given value.

### HasPsuInputSrc

`func (o *EquipmentPsuRelationship) HasPsuInputSrc() bool`

HasPsuInputSrc returns a boolean if a field has been set.

### GetPsuType

`func (o *EquipmentPsuRelationship) GetPsuType() string`

GetPsuType returns the PsuType field if non-nil, zero value otherwise.

### GetPsuTypeOk

`func (o *EquipmentPsuRelationship) GetPsuTypeOk() (*string, bool)`

GetPsuTypeOk returns a tuple with the PsuType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPsuType

`func (o *EquipmentPsuRelationship) SetPsuType(v string)`

SetPsuType sets PsuType field to given value.

### HasPsuType

`func (o *EquipmentPsuRelationship) HasPsuType() bool`

HasPsuType returns a boolean if a field has been set.

### GetPsuWattage

`func (o *EquipmentPsuRelationship) GetPsuWattage() string`

GetPsuWattage returns the PsuWattage field if non-nil, zero value otherwise.

### GetPsuWattageOk

`func (o *EquipmentPsuRelationship) GetPsuWattageOk() (*string, bool)`

GetPsuWattageOk returns a tuple with the PsuWattage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPsuWattage

`func (o *EquipmentPsuRelationship) SetPsuWattage(v string)`

SetPsuWattage sets PsuWattage field to given value.

### HasPsuWattage

`func (o *EquipmentPsuRelationship) HasPsuWattage() bool`

HasPsuWattage returns a boolean if a field has been set.

### GetSku

`func (o *EquipmentPsuRelationship) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *EquipmentPsuRelationship) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *EquipmentPsuRelationship) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *EquipmentPsuRelationship) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetVid

`func (o *EquipmentPsuRelationship) GetVid() string`

GetVid returns the Vid field if non-nil, zero value otherwise.

### GetVidOk

`func (o *EquipmentPsuRelationship) GetVidOk() (*string, bool)`

GetVidOk returns a tuple with the Vid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVid

`func (o *EquipmentPsuRelationship) SetVid(v string)`

SetVid sets Vid field to given value.

### HasVid

`func (o *EquipmentPsuRelationship) HasVid() bool`

HasVid returns a boolean if a field has been set.

### GetVoltage

`func (o *EquipmentPsuRelationship) GetVoltage() string`

GetVoltage returns the Voltage field if non-nil, zero value otherwise.

### GetVoltageOk

`func (o *EquipmentPsuRelationship) GetVoltageOk() (*string, bool)`

GetVoltageOk returns a tuple with the Voltage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoltage

`func (o *EquipmentPsuRelationship) SetVoltage(v string)`

SetVoltage sets Voltage field to given value.

### HasVoltage

`func (o *EquipmentPsuRelationship) HasVoltage() bool`

HasVoltage returns a boolean if a field has been set.

### GetComputeRackUnit

`func (o *EquipmentPsuRelationship) GetComputeRackUnit() ComputeRackUnitRelationship`

GetComputeRackUnit returns the ComputeRackUnit field if non-nil, zero value otherwise.

### GetComputeRackUnitOk

`func (o *EquipmentPsuRelationship) GetComputeRackUnitOk() (*ComputeRackUnitRelationship, bool)`

GetComputeRackUnitOk returns a tuple with the ComputeRackUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRackUnit

`func (o *EquipmentPsuRelationship) SetComputeRackUnit(v ComputeRackUnitRelationship)`

SetComputeRackUnit sets ComputeRackUnit field to given value.

### HasComputeRackUnit

`func (o *EquipmentPsuRelationship) HasComputeRackUnit() bool`

HasComputeRackUnit returns a boolean if a field has been set.

### GetEquipmentChassis

`func (o *EquipmentPsuRelationship) GetEquipmentChassis() EquipmentChassisRelationship`

GetEquipmentChassis returns the EquipmentChassis field if non-nil, zero value otherwise.

### GetEquipmentChassisOk

`func (o *EquipmentPsuRelationship) GetEquipmentChassisOk() (*EquipmentChassisRelationship, bool)`

GetEquipmentChassisOk returns a tuple with the EquipmentChassis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentChassis

`func (o *EquipmentPsuRelationship) SetEquipmentChassis(v EquipmentChassisRelationship)`

SetEquipmentChassis sets EquipmentChassis field to given value.

### HasEquipmentChassis

`func (o *EquipmentPsuRelationship) HasEquipmentChassis() bool`

HasEquipmentChassis returns a boolean if a field has been set.

### GetEquipmentFex

`func (o *EquipmentPsuRelationship) GetEquipmentFex() EquipmentFexRelationship`

GetEquipmentFex returns the EquipmentFex field if non-nil, zero value otherwise.

### GetEquipmentFexOk

`func (o *EquipmentPsuRelationship) GetEquipmentFexOk() (*EquipmentFexRelationship, bool)`

GetEquipmentFexOk returns a tuple with the EquipmentFex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentFex

`func (o *EquipmentPsuRelationship) SetEquipmentFex(v EquipmentFexRelationship)`

SetEquipmentFex sets EquipmentFex field to given value.

### HasEquipmentFex

`func (o *EquipmentPsuRelationship) HasEquipmentFex() bool`

HasEquipmentFex returns a boolean if a field has been set.

### GetEquipmentRackEnclosure

`func (o *EquipmentPsuRelationship) GetEquipmentRackEnclosure() EquipmentRackEnclosureRelationship`

GetEquipmentRackEnclosure returns the EquipmentRackEnclosure field if non-nil, zero value otherwise.

### GetEquipmentRackEnclosureOk

`func (o *EquipmentPsuRelationship) GetEquipmentRackEnclosureOk() (*EquipmentRackEnclosureRelationship, bool)`

GetEquipmentRackEnclosureOk returns a tuple with the EquipmentRackEnclosure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentRackEnclosure

`func (o *EquipmentPsuRelationship) SetEquipmentRackEnclosure(v EquipmentRackEnclosureRelationship)`

SetEquipmentRackEnclosure sets EquipmentRackEnclosure field to given value.

### HasEquipmentRackEnclosure

`func (o *EquipmentPsuRelationship) HasEquipmentRackEnclosure() bool`

HasEquipmentRackEnclosure returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *EquipmentPsuRelationship) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *EquipmentPsuRelationship) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *EquipmentPsuRelationship) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *EquipmentPsuRelationship) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetNetworkElement

`func (o *EquipmentPsuRelationship) GetNetworkElement() NetworkElementRelationship`

GetNetworkElement returns the NetworkElement field if non-nil, zero value otherwise.

### GetNetworkElementOk

`func (o *EquipmentPsuRelationship) GetNetworkElementOk() (*NetworkElementRelationship, bool)`

GetNetworkElementOk returns a tuple with the NetworkElement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkElement

`func (o *EquipmentPsuRelationship) SetNetworkElement(v NetworkElementRelationship)`

SetNetworkElement sets NetworkElement field to given value.

### HasNetworkElement

`func (o *EquipmentPsuRelationship) HasNetworkElement() bool`

HasNetworkElement returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *EquipmentPsuRelationship) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *EquipmentPsuRelationship) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *EquipmentPsuRelationship) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *EquipmentPsuRelationship) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



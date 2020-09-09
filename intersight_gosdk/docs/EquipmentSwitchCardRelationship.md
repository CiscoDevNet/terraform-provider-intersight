# EquipmentSwitchCardRelationship

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
**Description** | Pointer to **string** | Detailed description of this switch hardware. | [optional] [readonly] 
**NumPorts** | Pointer to **int64** | Number of ports present in this switch hardware. | [optional] [readonly] 
**OutOfBandIpAddress** | Pointer to **string** | Field specifies this Switch&#39;s Out-of-band IP address. | [optional] [readonly] 
**OutOfBandIpGateway** | Pointer to **string** | Field specifies this Switch&#39;s default gateway for the out-of-band management interface. | [optional] [readonly] 
**Presence** | Pointer to **string** | Presence for this switch hardware. | [optional] [readonly] 
**SlotId** | Pointer to **int64** | Slot identifier of the local Switch slot Interface. | [optional] [readonly] 
**State** | Pointer to **string** | Operational state of the switch hardware. | [optional] [readonly] 
**SwitchId** | Pointer to **string** | Switch Identifier that is local to a cluster. | [optional] [readonly] 
**FcPortChannels** | Pointer to [**[]FcPortChannelRelationship**](fc.PortChannel.Relationship.md) | An array of relationships to fcPortChannel resources. | [optional] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](inventory.DeviceInfo.Relationship.md) |  | [optional] 
**NetworkElement** | Pointer to [**NetworkElementRelationship**](network.Element.Relationship.md) |  | [optional] 
**PortChannels** | Pointer to [**[]EtherPortChannelRelationship**](ether.PortChannel.Relationship.md) | An array of relationships to etherPortChannel resources. | [optional] 
**PortGroups** | Pointer to [**[]PortGroupRelationship**](port.Group.Relationship.md) | An array of relationships to portGroup resources. | [optional] [readonly] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewEquipmentSwitchCardRelationship

`func NewEquipmentSwitchCardRelationship(classId string, objectType string, ) *EquipmentSwitchCardRelationship`

NewEquipmentSwitchCardRelationship instantiates a new EquipmentSwitchCardRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquipmentSwitchCardRelationshipWithDefaults

`func NewEquipmentSwitchCardRelationshipWithDefaults() *EquipmentSwitchCardRelationship`

NewEquipmentSwitchCardRelationshipWithDefaults instantiates a new EquipmentSwitchCardRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountMoid

`func (o *EquipmentSwitchCardRelationship) GetAccountMoid() string`

GetAccountMoid returns the AccountMoid field if non-nil, zero value otherwise.

### GetAccountMoidOk

`func (o *EquipmentSwitchCardRelationship) GetAccountMoidOk() (*string, bool)`

GetAccountMoidOk returns a tuple with the AccountMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMoid

`func (o *EquipmentSwitchCardRelationship) SetAccountMoid(v string)`

SetAccountMoid sets AccountMoid field to given value.

### HasAccountMoid

`func (o *EquipmentSwitchCardRelationship) HasAccountMoid() bool`

HasAccountMoid returns a boolean if a field has been set.

### GetClassId

`func (o *EquipmentSwitchCardRelationship) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *EquipmentSwitchCardRelationship) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *EquipmentSwitchCardRelationship) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetCreateTime

`func (o *EquipmentSwitchCardRelationship) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *EquipmentSwitchCardRelationship) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *EquipmentSwitchCardRelationship) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *EquipmentSwitchCardRelationship) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDomainGroupMoid

`func (o *EquipmentSwitchCardRelationship) GetDomainGroupMoid() string`

GetDomainGroupMoid returns the DomainGroupMoid field if non-nil, zero value otherwise.

### GetDomainGroupMoidOk

`func (o *EquipmentSwitchCardRelationship) GetDomainGroupMoidOk() (*string, bool)`

GetDomainGroupMoidOk returns a tuple with the DomainGroupMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGroupMoid

`func (o *EquipmentSwitchCardRelationship) SetDomainGroupMoid(v string)`

SetDomainGroupMoid sets DomainGroupMoid field to given value.

### HasDomainGroupMoid

`func (o *EquipmentSwitchCardRelationship) HasDomainGroupMoid() bool`

HasDomainGroupMoid returns a boolean if a field has been set.

### GetModTime

`func (o *EquipmentSwitchCardRelationship) GetModTime() time.Time`

GetModTime returns the ModTime field if non-nil, zero value otherwise.

### GetModTimeOk

`func (o *EquipmentSwitchCardRelationship) GetModTimeOk() (*time.Time, bool)`

GetModTimeOk returns a tuple with the ModTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModTime

`func (o *EquipmentSwitchCardRelationship) SetModTime(v time.Time)`

SetModTime sets ModTime field to given value.

### HasModTime

`func (o *EquipmentSwitchCardRelationship) HasModTime() bool`

HasModTime returns a boolean if a field has been set.

### GetMoid

`func (o *EquipmentSwitchCardRelationship) GetMoid() string`

GetMoid returns the Moid field if non-nil, zero value otherwise.

### GetMoidOk

`func (o *EquipmentSwitchCardRelationship) GetMoidOk() (*string, bool)`

GetMoidOk returns a tuple with the Moid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoid

`func (o *EquipmentSwitchCardRelationship) SetMoid(v string)`

SetMoid sets Moid field to given value.

### HasMoid

`func (o *EquipmentSwitchCardRelationship) HasMoid() bool`

HasMoid returns a boolean if a field has been set.

### GetObjectType

`func (o *EquipmentSwitchCardRelationship) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *EquipmentSwitchCardRelationship) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *EquipmentSwitchCardRelationship) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOwners

`func (o *EquipmentSwitchCardRelationship) GetOwners() []string`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *EquipmentSwitchCardRelationship) GetOwnersOk() (*[]string, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *EquipmentSwitchCardRelationship) SetOwners(v []string)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *EquipmentSwitchCardRelationship) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetSharedScope

`func (o *EquipmentSwitchCardRelationship) GetSharedScope() string`

GetSharedScope returns the SharedScope field if non-nil, zero value otherwise.

### GetSharedScopeOk

`func (o *EquipmentSwitchCardRelationship) GetSharedScopeOk() (*string, bool)`

GetSharedScopeOk returns a tuple with the SharedScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedScope

`func (o *EquipmentSwitchCardRelationship) SetSharedScope(v string)`

SetSharedScope sets SharedScope field to given value.

### HasSharedScope

`func (o *EquipmentSwitchCardRelationship) HasSharedScope() bool`

HasSharedScope returns a boolean if a field has been set.

### GetTags

`func (o *EquipmentSwitchCardRelationship) GetTags() []MoTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *EquipmentSwitchCardRelationship) GetTagsOk() (*[]MoTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *EquipmentSwitchCardRelationship) SetTags(v []MoTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *EquipmentSwitchCardRelationship) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersionContext

`func (o *EquipmentSwitchCardRelationship) GetVersionContext() MoVersionContext`

GetVersionContext returns the VersionContext field if non-nil, zero value otherwise.

### GetVersionContextOk

`func (o *EquipmentSwitchCardRelationship) GetVersionContextOk() (*MoVersionContext, bool)`

GetVersionContextOk returns a tuple with the VersionContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionContext

`func (o *EquipmentSwitchCardRelationship) SetVersionContext(v MoVersionContext)`

SetVersionContext sets VersionContext field to given value.

### HasVersionContext

`func (o *EquipmentSwitchCardRelationship) HasVersionContext() bool`

HasVersionContext returns a boolean if a field has been set.

### GetAncestors

`func (o *EquipmentSwitchCardRelationship) GetAncestors() []MoBaseMoRelationship`

GetAncestors returns the Ancestors field if non-nil, zero value otherwise.

### GetAncestorsOk

`func (o *EquipmentSwitchCardRelationship) GetAncestorsOk() (*[]MoBaseMoRelationship, bool)`

GetAncestorsOk returns a tuple with the Ancestors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestors

`func (o *EquipmentSwitchCardRelationship) SetAncestors(v []MoBaseMoRelationship)`

SetAncestors sets Ancestors field to given value.

### HasAncestors

`func (o *EquipmentSwitchCardRelationship) HasAncestors() bool`

HasAncestors returns a boolean if a field has been set.

### SetAncestorsNil

`func (o *EquipmentSwitchCardRelationship) SetAncestorsNil(b bool)`

 SetAncestorsNil sets the value for Ancestors to be an explicit nil

### UnsetAncestors
`func (o *EquipmentSwitchCardRelationship) UnsetAncestors()`

UnsetAncestors ensures that no value is present for Ancestors, not even an explicit nil
### GetParent

`func (o *EquipmentSwitchCardRelationship) GetParent() MoBaseMoRelationship`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *EquipmentSwitchCardRelationship) GetParentOk() (*MoBaseMoRelationship, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *EquipmentSwitchCardRelationship) SetParent(v MoBaseMoRelationship)`

SetParent sets Parent field to given value.

### HasParent

`func (o *EquipmentSwitchCardRelationship) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPermissionResources

`func (o *EquipmentSwitchCardRelationship) GetPermissionResources() []MoBaseMoRelationship`

GetPermissionResources returns the PermissionResources field if non-nil, zero value otherwise.

### GetPermissionResourcesOk

`func (o *EquipmentSwitchCardRelationship) GetPermissionResourcesOk() (*[]MoBaseMoRelationship, bool)`

GetPermissionResourcesOk returns a tuple with the PermissionResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionResources

`func (o *EquipmentSwitchCardRelationship) SetPermissionResources(v []MoBaseMoRelationship)`

SetPermissionResources sets PermissionResources field to given value.

### HasPermissionResources

`func (o *EquipmentSwitchCardRelationship) HasPermissionResources() bool`

HasPermissionResources returns a boolean if a field has been set.

### SetPermissionResourcesNil

`func (o *EquipmentSwitchCardRelationship) SetPermissionResourcesNil(b bool)`

 SetPermissionResourcesNil sets the value for PermissionResources to be an explicit nil

### UnsetPermissionResources
`func (o *EquipmentSwitchCardRelationship) UnsetPermissionResources()`

UnsetPermissionResources ensures that no value is present for PermissionResources, not even an explicit nil
### GetDisplayNames

`func (o *EquipmentSwitchCardRelationship) GetDisplayNames() map[string][]string`

GetDisplayNames returns the DisplayNames field if non-nil, zero value otherwise.

### GetDisplayNamesOk

`func (o *EquipmentSwitchCardRelationship) GetDisplayNamesOk() (*map[string][]string, bool)`

GetDisplayNamesOk returns a tuple with the DisplayNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNames

`func (o *EquipmentSwitchCardRelationship) SetDisplayNames(v map[string][]string)`

SetDisplayNames sets DisplayNames field to given value.

### HasDisplayNames

`func (o *EquipmentSwitchCardRelationship) HasDisplayNames() bool`

HasDisplayNames returns a boolean if a field has been set.

### SetDisplayNamesNil

`func (o *EquipmentSwitchCardRelationship) SetDisplayNamesNil(b bool)`

 SetDisplayNamesNil sets the value for DisplayNames to be an explicit nil

### UnsetDisplayNames
`func (o *EquipmentSwitchCardRelationship) UnsetDisplayNames()`

UnsetDisplayNames ensures that no value is present for DisplayNames, not even an explicit nil
### GetDeviceMoId

`func (o *EquipmentSwitchCardRelationship) GetDeviceMoId() string`

GetDeviceMoId returns the DeviceMoId field if non-nil, zero value otherwise.

### GetDeviceMoIdOk

`func (o *EquipmentSwitchCardRelationship) GetDeviceMoIdOk() (*string, bool)`

GetDeviceMoIdOk returns a tuple with the DeviceMoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMoId

`func (o *EquipmentSwitchCardRelationship) SetDeviceMoId(v string)`

SetDeviceMoId sets DeviceMoId field to given value.

### HasDeviceMoId

`func (o *EquipmentSwitchCardRelationship) HasDeviceMoId() bool`

HasDeviceMoId returns a boolean if a field has been set.

### GetDn

`func (o *EquipmentSwitchCardRelationship) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *EquipmentSwitchCardRelationship) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *EquipmentSwitchCardRelationship) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *EquipmentSwitchCardRelationship) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetRn

`func (o *EquipmentSwitchCardRelationship) GetRn() string`

GetRn returns the Rn field if non-nil, zero value otherwise.

### GetRnOk

`func (o *EquipmentSwitchCardRelationship) GetRnOk() (*string, bool)`

GetRnOk returns a tuple with the Rn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRn

`func (o *EquipmentSwitchCardRelationship) SetRn(v string)`

SetRn sets Rn field to given value.

### HasRn

`func (o *EquipmentSwitchCardRelationship) HasRn() bool`

HasRn returns a boolean if a field has been set.

### GetModel

`func (o *EquipmentSwitchCardRelationship) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *EquipmentSwitchCardRelationship) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *EquipmentSwitchCardRelationship) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *EquipmentSwitchCardRelationship) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetRevision

`func (o *EquipmentSwitchCardRelationship) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *EquipmentSwitchCardRelationship) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *EquipmentSwitchCardRelationship) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *EquipmentSwitchCardRelationship) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetSerial

`func (o *EquipmentSwitchCardRelationship) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *EquipmentSwitchCardRelationship) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *EquipmentSwitchCardRelationship) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *EquipmentSwitchCardRelationship) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetVendor

`func (o *EquipmentSwitchCardRelationship) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *EquipmentSwitchCardRelationship) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *EquipmentSwitchCardRelationship) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *EquipmentSwitchCardRelationship) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetDescription

`func (o *EquipmentSwitchCardRelationship) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EquipmentSwitchCardRelationship) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EquipmentSwitchCardRelationship) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EquipmentSwitchCardRelationship) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetNumPorts

`func (o *EquipmentSwitchCardRelationship) GetNumPorts() int64`

GetNumPorts returns the NumPorts field if non-nil, zero value otherwise.

### GetNumPortsOk

`func (o *EquipmentSwitchCardRelationship) GetNumPortsOk() (*int64, bool)`

GetNumPortsOk returns a tuple with the NumPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumPorts

`func (o *EquipmentSwitchCardRelationship) SetNumPorts(v int64)`

SetNumPorts sets NumPorts field to given value.

### HasNumPorts

`func (o *EquipmentSwitchCardRelationship) HasNumPorts() bool`

HasNumPorts returns a boolean if a field has been set.

### GetOutOfBandIpAddress

`func (o *EquipmentSwitchCardRelationship) GetOutOfBandIpAddress() string`

GetOutOfBandIpAddress returns the OutOfBandIpAddress field if non-nil, zero value otherwise.

### GetOutOfBandIpAddressOk

`func (o *EquipmentSwitchCardRelationship) GetOutOfBandIpAddressOk() (*string, bool)`

GetOutOfBandIpAddressOk returns a tuple with the OutOfBandIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutOfBandIpAddress

`func (o *EquipmentSwitchCardRelationship) SetOutOfBandIpAddress(v string)`

SetOutOfBandIpAddress sets OutOfBandIpAddress field to given value.

### HasOutOfBandIpAddress

`func (o *EquipmentSwitchCardRelationship) HasOutOfBandIpAddress() bool`

HasOutOfBandIpAddress returns a boolean if a field has been set.

### GetOutOfBandIpGateway

`func (o *EquipmentSwitchCardRelationship) GetOutOfBandIpGateway() string`

GetOutOfBandIpGateway returns the OutOfBandIpGateway field if non-nil, zero value otherwise.

### GetOutOfBandIpGatewayOk

`func (o *EquipmentSwitchCardRelationship) GetOutOfBandIpGatewayOk() (*string, bool)`

GetOutOfBandIpGatewayOk returns a tuple with the OutOfBandIpGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutOfBandIpGateway

`func (o *EquipmentSwitchCardRelationship) SetOutOfBandIpGateway(v string)`

SetOutOfBandIpGateway sets OutOfBandIpGateway field to given value.

### HasOutOfBandIpGateway

`func (o *EquipmentSwitchCardRelationship) HasOutOfBandIpGateway() bool`

HasOutOfBandIpGateway returns a boolean if a field has been set.

### GetPresence

`func (o *EquipmentSwitchCardRelationship) GetPresence() string`

GetPresence returns the Presence field if non-nil, zero value otherwise.

### GetPresenceOk

`func (o *EquipmentSwitchCardRelationship) GetPresenceOk() (*string, bool)`

GetPresenceOk returns a tuple with the Presence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresence

`func (o *EquipmentSwitchCardRelationship) SetPresence(v string)`

SetPresence sets Presence field to given value.

### HasPresence

`func (o *EquipmentSwitchCardRelationship) HasPresence() bool`

HasPresence returns a boolean if a field has been set.

### GetSlotId

`func (o *EquipmentSwitchCardRelationship) GetSlotId() int64`

GetSlotId returns the SlotId field if non-nil, zero value otherwise.

### GetSlotIdOk

`func (o *EquipmentSwitchCardRelationship) GetSlotIdOk() (*int64, bool)`

GetSlotIdOk returns a tuple with the SlotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotId

`func (o *EquipmentSwitchCardRelationship) SetSlotId(v int64)`

SetSlotId sets SlotId field to given value.

### HasSlotId

`func (o *EquipmentSwitchCardRelationship) HasSlotId() bool`

HasSlotId returns a boolean if a field has been set.

### GetState

`func (o *EquipmentSwitchCardRelationship) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *EquipmentSwitchCardRelationship) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *EquipmentSwitchCardRelationship) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *EquipmentSwitchCardRelationship) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSwitchId

`func (o *EquipmentSwitchCardRelationship) GetSwitchId() string`

GetSwitchId returns the SwitchId field if non-nil, zero value otherwise.

### GetSwitchIdOk

`func (o *EquipmentSwitchCardRelationship) GetSwitchIdOk() (*string, bool)`

GetSwitchIdOk returns a tuple with the SwitchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchId

`func (o *EquipmentSwitchCardRelationship) SetSwitchId(v string)`

SetSwitchId sets SwitchId field to given value.

### HasSwitchId

`func (o *EquipmentSwitchCardRelationship) HasSwitchId() bool`

HasSwitchId returns a boolean if a field has been set.

### GetFcPortChannels

`func (o *EquipmentSwitchCardRelationship) GetFcPortChannels() []FcPortChannelRelationship`

GetFcPortChannels returns the FcPortChannels field if non-nil, zero value otherwise.

### GetFcPortChannelsOk

`func (o *EquipmentSwitchCardRelationship) GetFcPortChannelsOk() (*[]FcPortChannelRelationship, bool)`

GetFcPortChannelsOk returns a tuple with the FcPortChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcPortChannels

`func (o *EquipmentSwitchCardRelationship) SetFcPortChannels(v []FcPortChannelRelationship)`

SetFcPortChannels sets FcPortChannels field to given value.

### HasFcPortChannels

`func (o *EquipmentSwitchCardRelationship) HasFcPortChannels() bool`

HasFcPortChannels returns a boolean if a field has been set.

### SetFcPortChannelsNil

`func (o *EquipmentSwitchCardRelationship) SetFcPortChannelsNil(b bool)`

 SetFcPortChannelsNil sets the value for FcPortChannels to be an explicit nil

### UnsetFcPortChannels
`func (o *EquipmentSwitchCardRelationship) UnsetFcPortChannels()`

UnsetFcPortChannels ensures that no value is present for FcPortChannels, not even an explicit nil
### GetInventoryDeviceInfo

`func (o *EquipmentSwitchCardRelationship) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *EquipmentSwitchCardRelationship) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *EquipmentSwitchCardRelationship) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *EquipmentSwitchCardRelationship) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetNetworkElement

`func (o *EquipmentSwitchCardRelationship) GetNetworkElement() NetworkElementRelationship`

GetNetworkElement returns the NetworkElement field if non-nil, zero value otherwise.

### GetNetworkElementOk

`func (o *EquipmentSwitchCardRelationship) GetNetworkElementOk() (*NetworkElementRelationship, bool)`

GetNetworkElementOk returns a tuple with the NetworkElement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkElement

`func (o *EquipmentSwitchCardRelationship) SetNetworkElement(v NetworkElementRelationship)`

SetNetworkElement sets NetworkElement field to given value.

### HasNetworkElement

`func (o *EquipmentSwitchCardRelationship) HasNetworkElement() bool`

HasNetworkElement returns a boolean if a field has been set.

### GetPortChannels

`func (o *EquipmentSwitchCardRelationship) GetPortChannels() []EtherPortChannelRelationship`

GetPortChannels returns the PortChannels field if non-nil, zero value otherwise.

### GetPortChannelsOk

`func (o *EquipmentSwitchCardRelationship) GetPortChannelsOk() (*[]EtherPortChannelRelationship, bool)`

GetPortChannelsOk returns a tuple with the PortChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortChannels

`func (o *EquipmentSwitchCardRelationship) SetPortChannels(v []EtherPortChannelRelationship)`

SetPortChannels sets PortChannels field to given value.

### HasPortChannels

`func (o *EquipmentSwitchCardRelationship) HasPortChannels() bool`

HasPortChannels returns a boolean if a field has been set.

### SetPortChannelsNil

`func (o *EquipmentSwitchCardRelationship) SetPortChannelsNil(b bool)`

 SetPortChannelsNil sets the value for PortChannels to be an explicit nil

### UnsetPortChannels
`func (o *EquipmentSwitchCardRelationship) UnsetPortChannels()`

UnsetPortChannels ensures that no value is present for PortChannels, not even an explicit nil
### GetPortGroups

`func (o *EquipmentSwitchCardRelationship) GetPortGroups() []PortGroupRelationship`

GetPortGroups returns the PortGroups field if non-nil, zero value otherwise.

### GetPortGroupsOk

`func (o *EquipmentSwitchCardRelationship) GetPortGroupsOk() (*[]PortGroupRelationship, bool)`

GetPortGroupsOk returns a tuple with the PortGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortGroups

`func (o *EquipmentSwitchCardRelationship) SetPortGroups(v []PortGroupRelationship)`

SetPortGroups sets PortGroups field to given value.

### HasPortGroups

`func (o *EquipmentSwitchCardRelationship) HasPortGroups() bool`

HasPortGroups returns a boolean if a field has been set.

### SetPortGroupsNil

`func (o *EquipmentSwitchCardRelationship) SetPortGroupsNil(b bool)`

 SetPortGroupsNil sets the value for PortGroups to be an explicit nil

### UnsetPortGroups
`func (o *EquipmentSwitchCardRelationship) UnsetPortGroups()`

UnsetPortGroups ensures that no value is present for PortGroups, not even an explicit nil
### GetRegisteredDevice

`func (o *EquipmentSwitchCardRelationship) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *EquipmentSwitchCardRelationship) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *EquipmentSwitchCardRelationship) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *EquipmentSwitchCardRelationship) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



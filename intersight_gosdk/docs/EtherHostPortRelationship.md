# EtherHostPortRelationship

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
**OperState** | Pointer to **string** | Operational state of this port (enabled/disabled). | [optional] [readonly] 
**OperStateQual** | Pointer to **string** | Reason for this port&#39;s Operational state. | [optional] [readonly] 
**PortId** | Pointer to **int64** | Switch physical port identifier. | [optional] [readonly] 
**Role** | Pointer to **string** | The role assigned to this port. | [optional] [readonly] 
**SlotId** | Pointer to **int64** | Switch expansion slot module identifier. | [optional] [readonly] 
**SwitchId** | Pointer to **string** | Switch Identifier that is local to a cluster. | [optional] [readonly] 
**MacAddress** | Pointer to **string** | Mac Address of a port in the Fabric Interconnect. | [optional] [readonly] 
**Mode** | Pointer to **string** | Operating mode of this port. | [optional] [readonly] 
**OperSpeed** | Pointer to **string** | Current Operational speed for this port. | [optional] [readonly] 
**PeerDn** | Pointer to **string** | PeerDn for ethernet physical port. | [optional] [readonly] 
**PortChannelId** | Pointer to **int64** | Port channel id for port channel created on FI switch. | [optional] [readonly] 
**PortType** | Pointer to **string** | Defines the transport type for this port (ethernet OR fc). | [optional] [readonly] 
**TransceiverType** | Pointer to **string** | Transceiver model attached to a port in the Fabric Interconnect. | [optional] [readonly] 
**AcknowledgedPeerInterface** | Pointer to [**PortInterfaceBaseRelationship**](port.InterfaceBase.Relationship.md) |  | [optional] 
**PeerInterface** | Pointer to [**PortInterfaceBaseRelationship**](port.InterfaceBase.Relationship.md) |  | [optional] 
**ModuleId** | Pointer to **int64** | Fabric extender identifier for this port. | [optional] 
**Speed** | Pointer to **string** | Host Port Speed of IO card or fabric extender. | [optional] [readonly] 
**EquipmentIoCardBase** | Pointer to [**EquipmentIoCardBaseRelationship**](equipment.IoCardBase.Relationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewEtherHostPortRelationship

`func NewEtherHostPortRelationship(classId string, objectType string, ) *EtherHostPortRelationship`

NewEtherHostPortRelationship instantiates a new EtherHostPortRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEtherHostPortRelationshipWithDefaults

`func NewEtherHostPortRelationshipWithDefaults() *EtherHostPortRelationship`

NewEtherHostPortRelationshipWithDefaults instantiates a new EtherHostPortRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountMoid

`func (o *EtherHostPortRelationship) GetAccountMoid() string`

GetAccountMoid returns the AccountMoid field if non-nil, zero value otherwise.

### GetAccountMoidOk

`func (o *EtherHostPortRelationship) GetAccountMoidOk() (*string, bool)`

GetAccountMoidOk returns a tuple with the AccountMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMoid

`func (o *EtherHostPortRelationship) SetAccountMoid(v string)`

SetAccountMoid sets AccountMoid field to given value.

### HasAccountMoid

`func (o *EtherHostPortRelationship) HasAccountMoid() bool`

HasAccountMoid returns a boolean if a field has been set.

### GetClassId

`func (o *EtherHostPortRelationship) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *EtherHostPortRelationship) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *EtherHostPortRelationship) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetCreateTime

`func (o *EtherHostPortRelationship) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *EtherHostPortRelationship) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *EtherHostPortRelationship) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *EtherHostPortRelationship) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDomainGroupMoid

`func (o *EtherHostPortRelationship) GetDomainGroupMoid() string`

GetDomainGroupMoid returns the DomainGroupMoid field if non-nil, zero value otherwise.

### GetDomainGroupMoidOk

`func (o *EtherHostPortRelationship) GetDomainGroupMoidOk() (*string, bool)`

GetDomainGroupMoidOk returns a tuple with the DomainGroupMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGroupMoid

`func (o *EtherHostPortRelationship) SetDomainGroupMoid(v string)`

SetDomainGroupMoid sets DomainGroupMoid field to given value.

### HasDomainGroupMoid

`func (o *EtherHostPortRelationship) HasDomainGroupMoid() bool`

HasDomainGroupMoid returns a boolean if a field has been set.

### GetModTime

`func (o *EtherHostPortRelationship) GetModTime() time.Time`

GetModTime returns the ModTime field if non-nil, zero value otherwise.

### GetModTimeOk

`func (o *EtherHostPortRelationship) GetModTimeOk() (*time.Time, bool)`

GetModTimeOk returns a tuple with the ModTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModTime

`func (o *EtherHostPortRelationship) SetModTime(v time.Time)`

SetModTime sets ModTime field to given value.

### HasModTime

`func (o *EtherHostPortRelationship) HasModTime() bool`

HasModTime returns a boolean if a field has been set.

### GetMoid

`func (o *EtherHostPortRelationship) GetMoid() string`

GetMoid returns the Moid field if non-nil, zero value otherwise.

### GetMoidOk

`func (o *EtherHostPortRelationship) GetMoidOk() (*string, bool)`

GetMoidOk returns a tuple with the Moid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoid

`func (o *EtherHostPortRelationship) SetMoid(v string)`

SetMoid sets Moid field to given value.

### HasMoid

`func (o *EtherHostPortRelationship) HasMoid() bool`

HasMoid returns a boolean if a field has been set.

### GetObjectType

`func (o *EtherHostPortRelationship) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *EtherHostPortRelationship) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *EtherHostPortRelationship) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOwners

`func (o *EtherHostPortRelationship) GetOwners() []string`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *EtherHostPortRelationship) GetOwnersOk() (*[]string, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *EtherHostPortRelationship) SetOwners(v []string)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *EtherHostPortRelationship) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetSharedScope

`func (o *EtherHostPortRelationship) GetSharedScope() string`

GetSharedScope returns the SharedScope field if non-nil, zero value otherwise.

### GetSharedScopeOk

`func (o *EtherHostPortRelationship) GetSharedScopeOk() (*string, bool)`

GetSharedScopeOk returns a tuple with the SharedScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedScope

`func (o *EtherHostPortRelationship) SetSharedScope(v string)`

SetSharedScope sets SharedScope field to given value.

### HasSharedScope

`func (o *EtherHostPortRelationship) HasSharedScope() bool`

HasSharedScope returns a boolean if a field has been set.

### GetTags

`func (o *EtherHostPortRelationship) GetTags() []MoTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *EtherHostPortRelationship) GetTagsOk() (*[]MoTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *EtherHostPortRelationship) SetTags(v []MoTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *EtherHostPortRelationship) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersionContext

`func (o *EtherHostPortRelationship) GetVersionContext() MoVersionContext`

GetVersionContext returns the VersionContext field if non-nil, zero value otherwise.

### GetVersionContextOk

`func (o *EtherHostPortRelationship) GetVersionContextOk() (*MoVersionContext, bool)`

GetVersionContextOk returns a tuple with the VersionContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionContext

`func (o *EtherHostPortRelationship) SetVersionContext(v MoVersionContext)`

SetVersionContext sets VersionContext field to given value.

### HasVersionContext

`func (o *EtherHostPortRelationship) HasVersionContext() bool`

HasVersionContext returns a boolean if a field has been set.

### GetAncestors

`func (o *EtherHostPortRelationship) GetAncestors() []MoBaseMoRelationship`

GetAncestors returns the Ancestors field if non-nil, zero value otherwise.

### GetAncestorsOk

`func (o *EtherHostPortRelationship) GetAncestorsOk() (*[]MoBaseMoRelationship, bool)`

GetAncestorsOk returns a tuple with the Ancestors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestors

`func (o *EtherHostPortRelationship) SetAncestors(v []MoBaseMoRelationship)`

SetAncestors sets Ancestors field to given value.

### HasAncestors

`func (o *EtherHostPortRelationship) HasAncestors() bool`

HasAncestors returns a boolean if a field has been set.

### SetAncestorsNil

`func (o *EtherHostPortRelationship) SetAncestorsNil(b bool)`

 SetAncestorsNil sets the value for Ancestors to be an explicit nil

### UnsetAncestors
`func (o *EtherHostPortRelationship) UnsetAncestors()`

UnsetAncestors ensures that no value is present for Ancestors, not even an explicit nil
### GetParent

`func (o *EtherHostPortRelationship) GetParent() MoBaseMoRelationship`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *EtherHostPortRelationship) GetParentOk() (*MoBaseMoRelationship, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *EtherHostPortRelationship) SetParent(v MoBaseMoRelationship)`

SetParent sets Parent field to given value.

### HasParent

`func (o *EtherHostPortRelationship) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPermissionResources

`func (o *EtherHostPortRelationship) GetPermissionResources() []MoBaseMoRelationship`

GetPermissionResources returns the PermissionResources field if non-nil, zero value otherwise.

### GetPermissionResourcesOk

`func (o *EtherHostPortRelationship) GetPermissionResourcesOk() (*[]MoBaseMoRelationship, bool)`

GetPermissionResourcesOk returns a tuple with the PermissionResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionResources

`func (o *EtherHostPortRelationship) SetPermissionResources(v []MoBaseMoRelationship)`

SetPermissionResources sets PermissionResources field to given value.

### HasPermissionResources

`func (o *EtherHostPortRelationship) HasPermissionResources() bool`

HasPermissionResources returns a boolean if a field has been set.

### SetPermissionResourcesNil

`func (o *EtherHostPortRelationship) SetPermissionResourcesNil(b bool)`

 SetPermissionResourcesNil sets the value for PermissionResources to be an explicit nil

### UnsetPermissionResources
`func (o *EtherHostPortRelationship) UnsetPermissionResources()`

UnsetPermissionResources ensures that no value is present for PermissionResources, not even an explicit nil
### GetDisplayNames

`func (o *EtherHostPortRelationship) GetDisplayNames() map[string][]string`

GetDisplayNames returns the DisplayNames field if non-nil, zero value otherwise.

### GetDisplayNamesOk

`func (o *EtherHostPortRelationship) GetDisplayNamesOk() (*map[string][]string, bool)`

GetDisplayNamesOk returns a tuple with the DisplayNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNames

`func (o *EtherHostPortRelationship) SetDisplayNames(v map[string][]string)`

SetDisplayNames sets DisplayNames field to given value.

### HasDisplayNames

`func (o *EtherHostPortRelationship) HasDisplayNames() bool`

HasDisplayNames returns a boolean if a field has been set.

### SetDisplayNamesNil

`func (o *EtherHostPortRelationship) SetDisplayNamesNil(b bool)`

 SetDisplayNamesNil sets the value for DisplayNames to be an explicit nil

### UnsetDisplayNames
`func (o *EtherHostPortRelationship) UnsetDisplayNames()`

UnsetDisplayNames ensures that no value is present for DisplayNames, not even an explicit nil
### GetDeviceMoId

`func (o *EtherHostPortRelationship) GetDeviceMoId() string`

GetDeviceMoId returns the DeviceMoId field if non-nil, zero value otherwise.

### GetDeviceMoIdOk

`func (o *EtherHostPortRelationship) GetDeviceMoIdOk() (*string, bool)`

GetDeviceMoIdOk returns a tuple with the DeviceMoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMoId

`func (o *EtherHostPortRelationship) SetDeviceMoId(v string)`

SetDeviceMoId sets DeviceMoId field to given value.

### HasDeviceMoId

`func (o *EtherHostPortRelationship) HasDeviceMoId() bool`

HasDeviceMoId returns a boolean if a field has been set.

### GetDn

`func (o *EtherHostPortRelationship) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *EtherHostPortRelationship) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *EtherHostPortRelationship) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *EtherHostPortRelationship) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetRn

`func (o *EtherHostPortRelationship) GetRn() string`

GetRn returns the Rn field if non-nil, zero value otherwise.

### GetRnOk

`func (o *EtherHostPortRelationship) GetRnOk() (*string, bool)`

GetRnOk returns a tuple with the Rn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRn

`func (o *EtherHostPortRelationship) SetRn(v string)`

SetRn sets Rn field to given value.

### HasRn

`func (o *EtherHostPortRelationship) HasRn() bool`

HasRn returns a boolean if a field has been set.

### GetOperState

`func (o *EtherHostPortRelationship) GetOperState() string`

GetOperState returns the OperState field if non-nil, zero value otherwise.

### GetOperStateOk

`func (o *EtherHostPortRelationship) GetOperStateOk() (*string, bool)`

GetOperStateOk returns a tuple with the OperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperState

`func (o *EtherHostPortRelationship) SetOperState(v string)`

SetOperState sets OperState field to given value.

### HasOperState

`func (o *EtherHostPortRelationship) HasOperState() bool`

HasOperState returns a boolean if a field has been set.

### GetOperStateQual

`func (o *EtherHostPortRelationship) GetOperStateQual() string`

GetOperStateQual returns the OperStateQual field if non-nil, zero value otherwise.

### GetOperStateQualOk

`func (o *EtherHostPortRelationship) GetOperStateQualOk() (*string, bool)`

GetOperStateQualOk returns a tuple with the OperStateQual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperStateQual

`func (o *EtherHostPortRelationship) SetOperStateQual(v string)`

SetOperStateQual sets OperStateQual field to given value.

### HasOperStateQual

`func (o *EtherHostPortRelationship) HasOperStateQual() bool`

HasOperStateQual returns a boolean if a field has been set.

### GetPortId

`func (o *EtherHostPortRelationship) GetPortId() int64`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *EtherHostPortRelationship) GetPortIdOk() (*int64, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *EtherHostPortRelationship) SetPortId(v int64)`

SetPortId sets PortId field to given value.

### HasPortId

`func (o *EtherHostPortRelationship) HasPortId() bool`

HasPortId returns a boolean if a field has been set.

### GetRole

`func (o *EtherHostPortRelationship) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *EtherHostPortRelationship) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *EtherHostPortRelationship) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *EtherHostPortRelationship) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetSlotId

`func (o *EtherHostPortRelationship) GetSlotId() int64`

GetSlotId returns the SlotId field if non-nil, zero value otherwise.

### GetSlotIdOk

`func (o *EtherHostPortRelationship) GetSlotIdOk() (*int64, bool)`

GetSlotIdOk returns a tuple with the SlotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotId

`func (o *EtherHostPortRelationship) SetSlotId(v int64)`

SetSlotId sets SlotId field to given value.

### HasSlotId

`func (o *EtherHostPortRelationship) HasSlotId() bool`

HasSlotId returns a boolean if a field has been set.

### GetSwitchId

`func (o *EtherHostPortRelationship) GetSwitchId() string`

GetSwitchId returns the SwitchId field if non-nil, zero value otherwise.

### GetSwitchIdOk

`func (o *EtherHostPortRelationship) GetSwitchIdOk() (*string, bool)`

GetSwitchIdOk returns a tuple with the SwitchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchId

`func (o *EtherHostPortRelationship) SetSwitchId(v string)`

SetSwitchId sets SwitchId field to given value.

### HasSwitchId

`func (o *EtherHostPortRelationship) HasSwitchId() bool`

HasSwitchId returns a boolean if a field has been set.

### GetMacAddress

`func (o *EtherHostPortRelationship) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *EtherHostPortRelationship) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *EtherHostPortRelationship) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *EtherHostPortRelationship) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetMode

`func (o *EtherHostPortRelationship) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *EtherHostPortRelationship) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *EtherHostPortRelationship) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *EtherHostPortRelationship) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetOperSpeed

`func (o *EtherHostPortRelationship) GetOperSpeed() string`

GetOperSpeed returns the OperSpeed field if non-nil, zero value otherwise.

### GetOperSpeedOk

`func (o *EtherHostPortRelationship) GetOperSpeedOk() (*string, bool)`

GetOperSpeedOk returns a tuple with the OperSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperSpeed

`func (o *EtherHostPortRelationship) SetOperSpeed(v string)`

SetOperSpeed sets OperSpeed field to given value.

### HasOperSpeed

`func (o *EtherHostPortRelationship) HasOperSpeed() bool`

HasOperSpeed returns a boolean if a field has been set.

### GetPeerDn

`func (o *EtherHostPortRelationship) GetPeerDn() string`

GetPeerDn returns the PeerDn field if non-nil, zero value otherwise.

### GetPeerDnOk

`func (o *EtherHostPortRelationship) GetPeerDnOk() (*string, bool)`

GetPeerDnOk returns a tuple with the PeerDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerDn

`func (o *EtherHostPortRelationship) SetPeerDn(v string)`

SetPeerDn sets PeerDn field to given value.

### HasPeerDn

`func (o *EtherHostPortRelationship) HasPeerDn() bool`

HasPeerDn returns a boolean if a field has been set.

### GetPortChannelId

`func (o *EtherHostPortRelationship) GetPortChannelId() int64`

GetPortChannelId returns the PortChannelId field if non-nil, zero value otherwise.

### GetPortChannelIdOk

`func (o *EtherHostPortRelationship) GetPortChannelIdOk() (*int64, bool)`

GetPortChannelIdOk returns a tuple with the PortChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortChannelId

`func (o *EtherHostPortRelationship) SetPortChannelId(v int64)`

SetPortChannelId sets PortChannelId field to given value.

### HasPortChannelId

`func (o *EtherHostPortRelationship) HasPortChannelId() bool`

HasPortChannelId returns a boolean if a field has been set.

### GetPortType

`func (o *EtherHostPortRelationship) GetPortType() string`

GetPortType returns the PortType field if non-nil, zero value otherwise.

### GetPortTypeOk

`func (o *EtherHostPortRelationship) GetPortTypeOk() (*string, bool)`

GetPortTypeOk returns a tuple with the PortType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortType

`func (o *EtherHostPortRelationship) SetPortType(v string)`

SetPortType sets PortType field to given value.

### HasPortType

`func (o *EtherHostPortRelationship) HasPortType() bool`

HasPortType returns a boolean if a field has been set.

### GetTransceiverType

`func (o *EtherHostPortRelationship) GetTransceiverType() string`

GetTransceiverType returns the TransceiverType field if non-nil, zero value otherwise.

### GetTransceiverTypeOk

`func (o *EtherHostPortRelationship) GetTransceiverTypeOk() (*string, bool)`

GetTransceiverTypeOk returns a tuple with the TransceiverType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransceiverType

`func (o *EtherHostPortRelationship) SetTransceiverType(v string)`

SetTransceiverType sets TransceiverType field to given value.

### HasTransceiverType

`func (o *EtherHostPortRelationship) HasTransceiverType() bool`

HasTransceiverType returns a boolean if a field has been set.

### GetAcknowledgedPeerInterface

`func (o *EtherHostPortRelationship) GetAcknowledgedPeerInterface() PortInterfaceBaseRelationship`

GetAcknowledgedPeerInterface returns the AcknowledgedPeerInterface field if non-nil, zero value otherwise.

### GetAcknowledgedPeerInterfaceOk

`func (o *EtherHostPortRelationship) GetAcknowledgedPeerInterfaceOk() (*PortInterfaceBaseRelationship, bool)`

GetAcknowledgedPeerInterfaceOk returns a tuple with the AcknowledgedPeerInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledgedPeerInterface

`func (o *EtherHostPortRelationship) SetAcknowledgedPeerInterface(v PortInterfaceBaseRelationship)`

SetAcknowledgedPeerInterface sets AcknowledgedPeerInterface field to given value.

### HasAcknowledgedPeerInterface

`func (o *EtherHostPortRelationship) HasAcknowledgedPeerInterface() bool`

HasAcknowledgedPeerInterface returns a boolean if a field has been set.

### GetPeerInterface

`func (o *EtherHostPortRelationship) GetPeerInterface() PortInterfaceBaseRelationship`

GetPeerInterface returns the PeerInterface field if non-nil, zero value otherwise.

### GetPeerInterfaceOk

`func (o *EtherHostPortRelationship) GetPeerInterfaceOk() (*PortInterfaceBaseRelationship, bool)`

GetPeerInterfaceOk returns a tuple with the PeerInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerInterface

`func (o *EtherHostPortRelationship) SetPeerInterface(v PortInterfaceBaseRelationship)`

SetPeerInterface sets PeerInterface field to given value.

### HasPeerInterface

`func (o *EtherHostPortRelationship) HasPeerInterface() bool`

HasPeerInterface returns a boolean if a field has been set.

### GetModuleId

`func (o *EtherHostPortRelationship) GetModuleId() int64`

GetModuleId returns the ModuleId field if non-nil, zero value otherwise.

### GetModuleIdOk

`func (o *EtherHostPortRelationship) GetModuleIdOk() (*int64, bool)`

GetModuleIdOk returns a tuple with the ModuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleId

`func (o *EtherHostPortRelationship) SetModuleId(v int64)`

SetModuleId sets ModuleId field to given value.

### HasModuleId

`func (o *EtherHostPortRelationship) HasModuleId() bool`

HasModuleId returns a boolean if a field has been set.

### GetSpeed

`func (o *EtherHostPortRelationship) GetSpeed() string`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *EtherHostPortRelationship) GetSpeedOk() (*string, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *EtherHostPortRelationship) SetSpeed(v string)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *EtherHostPortRelationship) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetEquipmentIoCardBase

`func (o *EtherHostPortRelationship) GetEquipmentIoCardBase() EquipmentIoCardBaseRelationship`

GetEquipmentIoCardBase returns the EquipmentIoCardBase field if non-nil, zero value otherwise.

### GetEquipmentIoCardBaseOk

`func (o *EtherHostPortRelationship) GetEquipmentIoCardBaseOk() (*EquipmentIoCardBaseRelationship, bool)`

GetEquipmentIoCardBaseOk returns a tuple with the EquipmentIoCardBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentIoCardBase

`func (o *EtherHostPortRelationship) SetEquipmentIoCardBase(v EquipmentIoCardBaseRelationship)`

SetEquipmentIoCardBase sets EquipmentIoCardBase field to given value.

### HasEquipmentIoCardBase

`func (o *EtherHostPortRelationship) HasEquipmentIoCardBase() bool`

HasEquipmentIoCardBase returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *EtherHostPortRelationship) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *EtherHostPortRelationship) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *EtherHostPortRelationship) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *EtherHostPortRelationship) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



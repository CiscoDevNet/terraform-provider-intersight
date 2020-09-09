# EtherPhysicalPortRelationship

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
**AdminSpeed** | Pointer to **string** | Administratively configured speed for this port. | [optional] [readonly] 
**AdminState** | Pointer to **string** | Administratively configured state (enabled/disabled) for this port. | [optional] [readonly] 
**AggregatePortId** | Pointer to **int64** | Breakout port member in the Fabric Interconnect. | [optional] [readonly] 
**LicenseGrace** | Pointer to **string** | The number of days this port&#39;s license has been in Grace Period for. | [optional] [readonly] 
**LicenseState** | Pointer to **string** | The state of the port&#39;s licensing. | [optional] [readonly] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](inventory.DeviceInfo.Relationship.md) |  | [optional] 
**PortGroup** | Pointer to [**PortGroupRelationship**](port.Group.Relationship.md) |  | [optional] 
**PortSubGroup** | Pointer to [**PortSubGroupRelationship**](port.SubGroup.Relationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewEtherPhysicalPortRelationship

`func NewEtherPhysicalPortRelationship(classId string, objectType string, ) *EtherPhysicalPortRelationship`

NewEtherPhysicalPortRelationship instantiates a new EtherPhysicalPortRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEtherPhysicalPortRelationshipWithDefaults

`func NewEtherPhysicalPortRelationshipWithDefaults() *EtherPhysicalPortRelationship`

NewEtherPhysicalPortRelationshipWithDefaults instantiates a new EtherPhysicalPortRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountMoid

`func (o *EtherPhysicalPortRelationship) GetAccountMoid() string`

GetAccountMoid returns the AccountMoid field if non-nil, zero value otherwise.

### GetAccountMoidOk

`func (o *EtherPhysicalPortRelationship) GetAccountMoidOk() (*string, bool)`

GetAccountMoidOk returns a tuple with the AccountMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMoid

`func (o *EtherPhysicalPortRelationship) SetAccountMoid(v string)`

SetAccountMoid sets AccountMoid field to given value.

### HasAccountMoid

`func (o *EtherPhysicalPortRelationship) HasAccountMoid() bool`

HasAccountMoid returns a boolean if a field has been set.

### GetClassId

`func (o *EtherPhysicalPortRelationship) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *EtherPhysicalPortRelationship) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *EtherPhysicalPortRelationship) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetCreateTime

`func (o *EtherPhysicalPortRelationship) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *EtherPhysicalPortRelationship) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *EtherPhysicalPortRelationship) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *EtherPhysicalPortRelationship) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDomainGroupMoid

`func (o *EtherPhysicalPortRelationship) GetDomainGroupMoid() string`

GetDomainGroupMoid returns the DomainGroupMoid field if non-nil, zero value otherwise.

### GetDomainGroupMoidOk

`func (o *EtherPhysicalPortRelationship) GetDomainGroupMoidOk() (*string, bool)`

GetDomainGroupMoidOk returns a tuple with the DomainGroupMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGroupMoid

`func (o *EtherPhysicalPortRelationship) SetDomainGroupMoid(v string)`

SetDomainGroupMoid sets DomainGroupMoid field to given value.

### HasDomainGroupMoid

`func (o *EtherPhysicalPortRelationship) HasDomainGroupMoid() bool`

HasDomainGroupMoid returns a boolean if a field has been set.

### GetModTime

`func (o *EtherPhysicalPortRelationship) GetModTime() time.Time`

GetModTime returns the ModTime field if non-nil, zero value otherwise.

### GetModTimeOk

`func (o *EtherPhysicalPortRelationship) GetModTimeOk() (*time.Time, bool)`

GetModTimeOk returns a tuple with the ModTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModTime

`func (o *EtherPhysicalPortRelationship) SetModTime(v time.Time)`

SetModTime sets ModTime field to given value.

### HasModTime

`func (o *EtherPhysicalPortRelationship) HasModTime() bool`

HasModTime returns a boolean if a field has been set.

### GetMoid

`func (o *EtherPhysicalPortRelationship) GetMoid() string`

GetMoid returns the Moid field if non-nil, zero value otherwise.

### GetMoidOk

`func (o *EtherPhysicalPortRelationship) GetMoidOk() (*string, bool)`

GetMoidOk returns a tuple with the Moid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoid

`func (o *EtherPhysicalPortRelationship) SetMoid(v string)`

SetMoid sets Moid field to given value.

### HasMoid

`func (o *EtherPhysicalPortRelationship) HasMoid() bool`

HasMoid returns a boolean if a field has been set.

### GetObjectType

`func (o *EtherPhysicalPortRelationship) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *EtherPhysicalPortRelationship) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *EtherPhysicalPortRelationship) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOwners

`func (o *EtherPhysicalPortRelationship) GetOwners() []string`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *EtherPhysicalPortRelationship) GetOwnersOk() (*[]string, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *EtherPhysicalPortRelationship) SetOwners(v []string)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *EtherPhysicalPortRelationship) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetSharedScope

`func (o *EtherPhysicalPortRelationship) GetSharedScope() string`

GetSharedScope returns the SharedScope field if non-nil, zero value otherwise.

### GetSharedScopeOk

`func (o *EtherPhysicalPortRelationship) GetSharedScopeOk() (*string, bool)`

GetSharedScopeOk returns a tuple with the SharedScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedScope

`func (o *EtherPhysicalPortRelationship) SetSharedScope(v string)`

SetSharedScope sets SharedScope field to given value.

### HasSharedScope

`func (o *EtherPhysicalPortRelationship) HasSharedScope() bool`

HasSharedScope returns a boolean if a field has been set.

### GetTags

`func (o *EtherPhysicalPortRelationship) GetTags() []MoTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *EtherPhysicalPortRelationship) GetTagsOk() (*[]MoTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *EtherPhysicalPortRelationship) SetTags(v []MoTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *EtherPhysicalPortRelationship) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersionContext

`func (o *EtherPhysicalPortRelationship) GetVersionContext() MoVersionContext`

GetVersionContext returns the VersionContext field if non-nil, zero value otherwise.

### GetVersionContextOk

`func (o *EtherPhysicalPortRelationship) GetVersionContextOk() (*MoVersionContext, bool)`

GetVersionContextOk returns a tuple with the VersionContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionContext

`func (o *EtherPhysicalPortRelationship) SetVersionContext(v MoVersionContext)`

SetVersionContext sets VersionContext field to given value.

### HasVersionContext

`func (o *EtherPhysicalPortRelationship) HasVersionContext() bool`

HasVersionContext returns a boolean if a field has been set.

### GetAncestors

`func (o *EtherPhysicalPortRelationship) GetAncestors() []MoBaseMoRelationship`

GetAncestors returns the Ancestors field if non-nil, zero value otherwise.

### GetAncestorsOk

`func (o *EtherPhysicalPortRelationship) GetAncestorsOk() (*[]MoBaseMoRelationship, bool)`

GetAncestorsOk returns a tuple with the Ancestors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestors

`func (o *EtherPhysicalPortRelationship) SetAncestors(v []MoBaseMoRelationship)`

SetAncestors sets Ancestors field to given value.

### HasAncestors

`func (o *EtherPhysicalPortRelationship) HasAncestors() bool`

HasAncestors returns a boolean if a field has been set.

### SetAncestorsNil

`func (o *EtherPhysicalPortRelationship) SetAncestorsNil(b bool)`

 SetAncestorsNil sets the value for Ancestors to be an explicit nil

### UnsetAncestors
`func (o *EtherPhysicalPortRelationship) UnsetAncestors()`

UnsetAncestors ensures that no value is present for Ancestors, not even an explicit nil
### GetParent

`func (o *EtherPhysicalPortRelationship) GetParent() MoBaseMoRelationship`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *EtherPhysicalPortRelationship) GetParentOk() (*MoBaseMoRelationship, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *EtherPhysicalPortRelationship) SetParent(v MoBaseMoRelationship)`

SetParent sets Parent field to given value.

### HasParent

`func (o *EtherPhysicalPortRelationship) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPermissionResources

`func (o *EtherPhysicalPortRelationship) GetPermissionResources() []MoBaseMoRelationship`

GetPermissionResources returns the PermissionResources field if non-nil, zero value otherwise.

### GetPermissionResourcesOk

`func (o *EtherPhysicalPortRelationship) GetPermissionResourcesOk() (*[]MoBaseMoRelationship, bool)`

GetPermissionResourcesOk returns a tuple with the PermissionResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionResources

`func (o *EtherPhysicalPortRelationship) SetPermissionResources(v []MoBaseMoRelationship)`

SetPermissionResources sets PermissionResources field to given value.

### HasPermissionResources

`func (o *EtherPhysicalPortRelationship) HasPermissionResources() bool`

HasPermissionResources returns a boolean if a field has been set.

### SetPermissionResourcesNil

`func (o *EtherPhysicalPortRelationship) SetPermissionResourcesNil(b bool)`

 SetPermissionResourcesNil sets the value for PermissionResources to be an explicit nil

### UnsetPermissionResources
`func (o *EtherPhysicalPortRelationship) UnsetPermissionResources()`

UnsetPermissionResources ensures that no value is present for PermissionResources, not even an explicit nil
### GetDisplayNames

`func (o *EtherPhysicalPortRelationship) GetDisplayNames() map[string][]string`

GetDisplayNames returns the DisplayNames field if non-nil, zero value otherwise.

### GetDisplayNamesOk

`func (o *EtherPhysicalPortRelationship) GetDisplayNamesOk() (*map[string][]string, bool)`

GetDisplayNamesOk returns a tuple with the DisplayNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNames

`func (o *EtherPhysicalPortRelationship) SetDisplayNames(v map[string][]string)`

SetDisplayNames sets DisplayNames field to given value.

### HasDisplayNames

`func (o *EtherPhysicalPortRelationship) HasDisplayNames() bool`

HasDisplayNames returns a boolean if a field has been set.

### SetDisplayNamesNil

`func (o *EtherPhysicalPortRelationship) SetDisplayNamesNil(b bool)`

 SetDisplayNamesNil sets the value for DisplayNames to be an explicit nil

### UnsetDisplayNames
`func (o *EtherPhysicalPortRelationship) UnsetDisplayNames()`

UnsetDisplayNames ensures that no value is present for DisplayNames, not even an explicit nil
### GetDeviceMoId

`func (o *EtherPhysicalPortRelationship) GetDeviceMoId() string`

GetDeviceMoId returns the DeviceMoId field if non-nil, zero value otherwise.

### GetDeviceMoIdOk

`func (o *EtherPhysicalPortRelationship) GetDeviceMoIdOk() (*string, bool)`

GetDeviceMoIdOk returns a tuple with the DeviceMoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMoId

`func (o *EtherPhysicalPortRelationship) SetDeviceMoId(v string)`

SetDeviceMoId sets DeviceMoId field to given value.

### HasDeviceMoId

`func (o *EtherPhysicalPortRelationship) HasDeviceMoId() bool`

HasDeviceMoId returns a boolean if a field has been set.

### GetDn

`func (o *EtherPhysicalPortRelationship) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *EtherPhysicalPortRelationship) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *EtherPhysicalPortRelationship) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *EtherPhysicalPortRelationship) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetRn

`func (o *EtherPhysicalPortRelationship) GetRn() string`

GetRn returns the Rn field if non-nil, zero value otherwise.

### GetRnOk

`func (o *EtherPhysicalPortRelationship) GetRnOk() (*string, bool)`

GetRnOk returns a tuple with the Rn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRn

`func (o *EtherPhysicalPortRelationship) SetRn(v string)`

SetRn sets Rn field to given value.

### HasRn

`func (o *EtherPhysicalPortRelationship) HasRn() bool`

HasRn returns a boolean if a field has been set.

### GetOperState

`func (o *EtherPhysicalPortRelationship) GetOperState() string`

GetOperState returns the OperState field if non-nil, zero value otherwise.

### GetOperStateOk

`func (o *EtherPhysicalPortRelationship) GetOperStateOk() (*string, bool)`

GetOperStateOk returns a tuple with the OperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperState

`func (o *EtherPhysicalPortRelationship) SetOperState(v string)`

SetOperState sets OperState field to given value.

### HasOperState

`func (o *EtherPhysicalPortRelationship) HasOperState() bool`

HasOperState returns a boolean if a field has been set.

### GetOperStateQual

`func (o *EtherPhysicalPortRelationship) GetOperStateQual() string`

GetOperStateQual returns the OperStateQual field if non-nil, zero value otherwise.

### GetOperStateQualOk

`func (o *EtherPhysicalPortRelationship) GetOperStateQualOk() (*string, bool)`

GetOperStateQualOk returns a tuple with the OperStateQual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperStateQual

`func (o *EtherPhysicalPortRelationship) SetOperStateQual(v string)`

SetOperStateQual sets OperStateQual field to given value.

### HasOperStateQual

`func (o *EtherPhysicalPortRelationship) HasOperStateQual() bool`

HasOperStateQual returns a boolean if a field has been set.

### GetPortId

`func (o *EtherPhysicalPortRelationship) GetPortId() int64`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *EtherPhysicalPortRelationship) GetPortIdOk() (*int64, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *EtherPhysicalPortRelationship) SetPortId(v int64)`

SetPortId sets PortId field to given value.

### HasPortId

`func (o *EtherPhysicalPortRelationship) HasPortId() bool`

HasPortId returns a boolean if a field has been set.

### GetRole

`func (o *EtherPhysicalPortRelationship) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *EtherPhysicalPortRelationship) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *EtherPhysicalPortRelationship) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *EtherPhysicalPortRelationship) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetSlotId

`func (o *EtherPhysicalPortRelationship) GetSlotId() int64`

GetSlotId returns the SlotId field if non-nil, zero value otherwise.

### GetSlotIdOk

`func (o *EtherPhysicalPortRelationship) GetSlotIdOk() (*int64, bool)`

GetSlotIdOk returns a tuple with the SlotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotId

`func (o *EtherPhysicalPortRelationship) SetSlotId(v int64)`

SetSlotId sets SlotId field to given value.

### HasSlotId

`func (o *EtherPhysicalPortRelationship) HasSlotId() bool`

HasSlotId returns a boolean if a field has been set.

### GetSwitchId

`func (o *EtherPhysicalPortRelationship) GetSwitchId() string`

GetSwitchId returns the SwitchId field if non-nil, zero value otherwise.

### GetSwitchIdOk

`func (o *EtherPhysicalPortRelationship) GetSwitchIdOk() (*string, bool)`

GetSwitchIdOk returns a tuple with the SwitchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchId

`func (o *EtherPhysicalPortRelationship) SetSwitchId(v string)`

SetSwitchId sets SwitchId field to given value.

### HasSwitchId

`func (o *EtherPhysicalPortRelationship) HasSwitchId() bool`

HasSwitchId returns a boolean if a field has been set.

### GetMacAddress

`func (o *EtherPhysicalPortRelationship) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *EtherPhysicalPortRelationship) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *EtherPhysicalPortRelationship) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *EtherPhysicalPortRelationship) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetMode

`func (o *EtherPhysicalPortRelationship) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *EtherPhysicalPortRelationship) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *EtherPhysicalPortRelationship) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *EtherPhysicalPortRelationship) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetOperSpeed

`func (o *EtherPhysicalPortRelationship) GetOperSpeed() string`

GetOperSpeed returns the OperSpeed field if non-nil, zero value otherwise.

### GetOperSpeedOk

`func (o *EtherPhysicalPortRelationship) GetOperSpeedOk() (*string, bool)`

GetOperSpeedOk returns a tuple with the OperSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperSpeed

`func (o *EtherPhysicalPortRelationship) SetOperSpeed(v string)`

SetOperSpeed sets OperSpeed field to given value.

### HasOperSpeed

`func (o *EtherPhysicalPortRelationship) HasOperSpeed() bool`

HasOperSpeed returns a boolean if a field has been set.

### GetPeerDn

`func (o *EtherPhysicalPortRelationship) GetPeerDn() string`

GetPeerDn returns the PeerDn field if non-nil, zero value otherwise.

### GetPeerDnOk

`func (o *EtherPhysicalPortRelationship) GetPeerDnOk() (*string, bool)`

GetPeerDnOk returns a tuple with the PeerDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerDn

`func (o *EtherPhysicalPortRelationship) SetPeerDn(v string)`

SetPeerDn sets PeerDn field to given value.

### HasPeerDn

`func (o *EtherPhysicalPortRelationship) HasPeerDn() bool`

HasPeerDn returns a boolean if a field has been set.

### GetPortChannelId

`func (o *EtherPhysicalPortRelationship) GetPortChannelId() int64`

GetPortChannelId returns the PortChannelId field if non-nil, zero value otherwise.

### GetPortChannelIdOk

`func (o *EtherPhysicalPortRelationship) GetPortChannelIdOk() (*int64, bool)`

GetPortChannelIdOk returns a tuple with the PortChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortChannelId

`func (o *EtherPhysicalPortRelationship) SetPortChannelId(v int64)`

SetPortChannelId sets PortChannelId field to given value.

### HasPortChannelId

`func (o *EtherPhysicalPortRelationship) HasPortChannelId() bool`

HasPortChannelId returns a boolean if a field has been set.

### GetPortType

`func (o *EtherPhysicalPortRelationship) GetPortType() string`

GetPortType returns the PortType field if non-nil, zero value otherwise.

### GetPortTypeOk

`func (o *EtherPhysicalPortRelationship) GetPortTypeOk() (*string, bool)`

GetPortTypeOk returns a tuple with the PortType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortType

`func (o *EtherPhysicalPortRelationship) SetPortType(v string)`

SetPortType sets PortType field to given value.

### HasPortType

`func (o *EtherPhysicalPortRelationship) HasPortType() bool`

HasPortType returns a boolean if a field has been set.

### GetTransceiverType

`func (o *EtherPhysicalPortRelationship) GetTransceiverType() string`

GetTransceiverType returns the TransceiverType field if non-nil, zero value otherwise.

### GetTransceiverTypeOk

`func (o *EtherPhysicalPortRelationship) GetTransceiverTypeOk() (*string, bool)`

GetTransceiverTypeOk returns a tuple with the TransceiverType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransceiverType

`func (o *EtherPhysicalPortRelationship) SetTransceiverType(v string)`

SetTransceiverType sets TransceiverType field to given value.

### HasTransceiverType

`func (o *EtherPhysicalPortRelationship) HasTransceiverType() bool`

HasTransceiverType returns a boolean if a field has been set.

### GetAcknowledgedPeerInterface

`func (o *EtherPhysicalPortRelationship) GetAcknowledgedPeerInterface() PortInterfaceBaseRelationship`

GetAcknowledgedPeerInterface returns the AcknowledgedPeerInterface field if non-nil, zero value otherwise.

### GetAcknowledgedPeerInterfaceOk

`func (o *EtherPhysicalPortRelationship) GetAcknowledgedPeerInterfaceOk() (*PortInterfaceBaseRelationship, bool)`

GetAcknowledgedPeerInterfaceOk returns a tuple with the AcknowledgedPeerInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledgedPeerInterface

`func (o *EtherPhysicalPortRelationship) SetAcknowledgedPeerInterface(v PortInterfaceBaseRelationship)`

SetAcknowledgedPeerInterface sets AcknowledgedPeerInterface field to given value.

### HasAcknowledgedPeerInterface

`func (o *EtherPhysicalPortRelationship) HasAcknowledgedPeerInterface() bool`

HasAcknowledgedPeerInterface returns a boolean if a field has been set.

### GetPeerInterface

`func (o *EtherPhysicalPortRelationship) GetPeerInterface() PortInterfaceBaseRelationship`

GetPeerInterface returns the PeerInterface field if non-nil, zero value otherwise.

### GetPeerInterfaceOk

`func (o *EtherPhysicalPortRelationship) GetPeerInterfaceOk() (*PortInterfaceBaseRelationship, bool)`

GetPeerInterfaceOk returns a tuple with the PeerInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerInterface

`func (o *EtherPhysicalPortRelationship) SetPeerInterface(v PortInterfaceBaseRelationship)`

SetPeerInterface sets PeerInterface field to given value.

### HasPeerInterface

`func (o *EtherPhysicalPortRelationship) HasPeerInterface() bool`

HasPeerInterface returns a boolean if a field has been set.

### GetAdminSpeed

`func (o *EtherPhysicalPortRelationship) GetAdminSpeed() string`

GetAdminSpeed returns the AdminSpeed field if non-nil, zero value otherwise.

### GetAdminSpeedOk

`func (o *EtherPhysicalPortRelationship) GetAdminSpeedOk() (*string, bool)`

GetAdminSpeedOk returns a tuple with the AdminSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminSpeed

`func (o *EtherPhysicalPortRelationship) SetAdminSpeed(v string)`

SetAdminSpeed sets AdminSpeed field to given value.

### HasAdminSpeed

`func (o *EtherPhysicalPortRelationship) HasAdminSpeed() bool`

HasAdminSpeed returns a boolean if a field has been set.

### GetAdminState

`func (o *EtherPhysicalPortRelationship) GetAdminState() string`

GetAdminState returns the AdminState field if non-nil, zero value otherwise.

### GetAdminStateOk

`func (o *EtherPhysicalPortRelationship) GetAdminStateOk() (*string, bool)`

GetAdminStateOk returns a tuple with the AdminState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminState

`func (o *EtherPhysicalPortRelationship) SetAdminState(v string)`

SetAdminState sets AdminState field to given value.

### HasAdminState

`func (o *EtherPhysicalPortRelationship) HasAdminState() bool`

HasAdminState returns a boolean if a field has been set.

### GetAggregatePortId

`func (o *EtherPhysicalPortRelationship) GetAggregatePortId() int64`

GetAggregatePortId returns the AggregatePortId field if non-nil, zero value otherwise.

### GetAggregatePortIdOk

`func (o *EtherPhysicalPortRelationship) GetAggregatePortIdOk() (*int64, bool)`

GetAggregatePortIdOk returns a tuple with the AggregatePortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregatePortId

`func (o *EtherPhysicalPortRelationship) SetAggregatePortId(v int64)`

SetAggregatePortId sets AggregatePortId field to given value.

### HasAggregatePortId

`func (o *EtherPhysicalPortRelationship) HasAggregatePortId() bool`

HasAggregatePortId returns a boolean if a field has been set.

### GetLicenseGrace

`func (o *EtherPhysicalPortRelationship) GetLicenseGrace() string`

GetLicenseGrace returns the LicenseGrace field if non-nil, zero value otherwise.

### GetLicenseGraceOk

`func (o *EtherPhysicalPortRelationship) GetLicenseGraceOk() (*string, bool)`

GetLicenseGraceOk returns a tuple with the LicenseGrace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseGrace

`func (o *EtherPhysicalPortRelationship) SetLicenseGrace(v string)`

SetLicenseGrace sets LicenseGrace field to given value.

### HasLicenseGrace

`func (o *EtherPhysicalPortRelationship) HasLicenseGrace() bool`

HasLicenseGrace returns a boolean if a field has been set.

### GetLicenseState

`func (o *EtherPhysicalPortRelationship) GetLicenseState() string`

GetLicenseState returns the LicenseState field if non-nil, zero value otherwise.

### GetLicenseStateOk

`func (o *EtherPhysicalPortRelationship) GetLicenseStateOk() (*string, bool)`

GetLicenseStateOk returns a tuple with the LicenseState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseState

`func (o *EtherPhysicalPortRelationship) SetLicenseState(v string)`

SetLicenseState sets LicenseState field to given value.

### HasLicenseState

`func (o *EtherPhysicalPortRelationship) HasLicenseState() bool`

HasLicenseState returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *EtherPhysicalPortRelationship) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *EtherPhysicalPortRelationship) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *EtherPhysicalPortRelationship) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *EtherPhysicalPortRelationship) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetPortGroup

`func (o *EtherPhysicalPortRelationship) GetPortGroup() PortGroupRelationship`

GetPortGroup returns the PortGroup field if non-nil, zero value otherwise.

### GetPortGroupOk

`func (o *EtherPhysicalPortRelationship) GetPortGroupOk() (*PortGroupRelationship, bool)`

GetPortGroupOk returns a tuple with the PortGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortGroup

`func (o *EtherPhysicalPortRelationship) SetPortGroup(v PortGroupRelationship)`

SetPortGroup sets PortGroup field to given value.

### HasPortGroup

`func (o *EtherPhysicalPortRelationship) HasPortGroup() bool`

HasPortGroup returns a boolean if a field has been set.

### GetPortSubGroup

`func (o *EtherPhysicalPortRelationship) GetPortSubGroup() PortSubGroupRelationship`

GetPortSubGroup returns the PortSubGroup field if non-nil, zero value otherwise.

### GetPortSubGroupOk

`func (o *EtherPhysicalPortRelationship) GetPortSubGroupOk() (*PortSubGroupRelationship, bool)`

GetPortSubGroupOk returns a tuple with the PortSubGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortSubGroup

`func (o *EtherPhysicalPortRelationship) SetPortSubGroup(v PortSubGroupRelationship)`

SetPortSubGroup sets PortSubGroup field to given value.

### HasPortSubGroup

`func (o *EtherPhysicalPortRelationship) HasPortSubGroup() bool`

HasPortSubGroup returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *EtherPhysicalPortRelationship) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *EtherPhysicalPortRelationship) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *EtherPhysicalPortRelationship) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *EtherPhysicalPortRelationship) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# EtherNetworkPortRelationship

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
**OperState** | Pointer to **string** | Operational state of an Interface. | [optional] 
**AcknowledgedPeerInterface** | Pointer to [**EtherPhysicalPortBaseRelationship**](ether.PhysicalPortBase.Relationship.md) |  | [optional] 
**PeerInterface** | Pointer to [**EtherPhysicalPortBaseRelationship**](ether.PhysicalPortBase.Relationship.md) |  | [optional] 
**ModuleId** | Pointer to **int64** | Febric extender identifier for this port. | [optional] 
**PeerDn** | Pointer to **string** | Peer DN for network host port of fabric extender. | [optional] 
**PortId** | Pointer to **int64** | Switch physical port identifier. | [optional] 
**SlotId** | Pointer to **int64** | Switch expansion slot module identifier. | [optional] 
**Speed** | Pointer to **string** | Network Port Speed of IO card or fabric extender. | [optional] [readonly] 
**SwitchId** | Pointer to **string** | Switch Identifier that is local to a cluster. | [optional] 
**EquipmentIoCardBase** | Pointer to [**EquipmentIoCardBaseRelationship**](equipment.IoCardBase.Relationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewEtherNetworkPortRelationship

`func NewEtherNetworkPortRelationship(classId string, objectType string, ) *EtherNetworkPortRelationship`

NewEtherNetworkPortRelationship instantiates a new EtherNetworkPortRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEtherNetworkPortRelationshipWithDefaults

`func NewEtherNetworkPortRelationshipWithDefaults() *EtherNetworkPortRelationship`

NewEtherNetworkPortRelationshipWithDefaults instantiates a new EtherNetworkPortRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountMoid

`func (o *EtherNetworkPortRelationship) GetAccountMoid() string`

GetAccountMoid returns the AccountMoid field if non-nil, zero value otherwise.

### GetAccountMoidOk

`func (o *EtherNetworkPortRelationship) GetAccountMoidOk() (*string, bool)`

GetAccountMoidOk returns a tuple with the AccountMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMoid

`func (o *EtherNetworkPortRelationship) SetAccountMoid(v string)`

SetAccountMoid sets AccountMoid field to given value.

### HasAccountMoid

`func (o *EtherNetworkPortRelationship) HasAccountMoid() bool`

HasAccountMoid returns a boolean if a field has been set.

### GetClassId

`func (o *EtherNetworkPortRelationship) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *EtherNetworkPortRelationship) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *EtherNetworkPortRelationship) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetCreateTime

`func (o *EtherNetworkPortRelationship) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *EtherNetworkPortRelationship) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *EtherNetworkPortRelationship) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *EtherNetworkPortRelationship) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDomainGroupMoid

`func (o *EtherNetworkPortRelationship) GetDomainGroupMoid() string`

GetDomainGroupMoid returns the DomainGroupMoid field if non-nil, zero value otherwise.

### GetDomainGroupMoidOk

`func (o *EtherNetworkPortRelationship) GetDomainGroupMoidOk() (*string, bool)`

GetDomainGroupMoidOk returns a tuple with the DomainGroupMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGroupMoid

`func (o *EtherNetworkPortRelationship) SetDomainGroupMoid(v string)`

SetDomainGroupMoid sets DomainGroupMoid field to given value.

### HasDomainGroupMoid

`func (o *EtherNetworkPortRelationship) HasDomainGroupMoid() bool`

HasDomainGroupMoid returns a boolean if a field has been set.

### GetModTime

`func (o *EtherNetworkPortRelationship) GetModTime() time.Time`

GetModTime returns the ModTime field if non-nil, zero value otherwise.

### GetModTimeOk

`func (o *EtherNetworkPortRelationship) GetModTimeOk() (*time.Time, bool)`

GetModTimeOk returns a tuple with the ModTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModTime

`func (o *EtherNetworkPortRelationship) SetModTime(v time.Time)`

SetModTime sets ModTime field to given value.

### HasModTime

`func (o *EtherNetworkPortRelationship) HasModTime() bool`

HasModTime returns a boolean if a field has been set.

### GetMoid

`func (o *EtherNetworkPortRelationship) GetMoid() string`

GetMoid returns the Moid field if non-nil, zero value otherwise.

### GetMoidOk

`func (o *EtherNetworkPortRelationship) GetMoidOk() (*string, bool)`

GetMoidOk returns a tuple with the Moid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoid

`func (o *EtherNetworkPortRelationship) SetMoid(v string)`

SetMoid sets Moid field to given value.

### HasMoid

`func (o *EtherNetworkPortRelationship) HasMoid() bool`

HasMoid returns a boolean if a field has been set.

### GetObjectType

`func (o *EtherNetworkPortRelationship) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *EtherNetworkPortRelationship) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *EtherNetworkPortRelationship) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOwners

`func (o *EtherNetworkPortRelationship) GetOwners() []string`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *EtherNetworkPortRelationship) GetOwnersOk() (*[]string, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *EtherNetworkPortRelationship) SetOwners(v []string)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *EtherNetworkPortRelationship) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetSharedScope

`func (o *EtherNetworkPortRelationship) GetSharedScope() string`

GetSharedScope returns the SharedScope field if non-nil, zero value otherwise.

### GetSharedScopeOk

`func (o *EtherNetworkPortRelationship) GetSharedScopeOk() (*string, bool)`

GetSharedScopeOk returns a tuple with the SharedScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedScope

`func (o *EtherNetworkPortRelationship) SetSharedScope(v string)`

SetSharedScope sets SharedScope field to given value.

### HasSharedScope

`func (o *EtherNetworkPortRelationship) HasSharedScope() bool`

HasSharedScope returns a boolean if a field has been set.

### GetTags

`func (o *EtherNetworkPortRelationship) GetTags() []MoTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *EtherNetworkPortRelationship) GetTagsOk() (*[]MoTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *EtherNetworkPortRelationship) SetTags(v []MoTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *EtherNetworkPortRelationship) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersionContext

`func (o *EtherNetworkPortRelationship) GetVersionContext() MoVersionContext`

GetVersionContext returns the VersionContext field if non-nil, zero value otherwise.

### GetVersionContextOk

`func (o *EtherNetworkPortRelationship) GetVersionContextOk() (*MoVersionContext, bool)`

GetVersionContextOk returns a tuple with the VersionContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionContext

`func (o *EtherNetworkPortRelationship) SetVersionContext(v MoVersionContext)`

SetVersionContext sets VersionContext field to given value.

### HasVersionContext

`func (o *EtherNetworkPortRelationship) HasVersionContext() bool`

HasVersionContext returns a boolean if a field has been set.

### GetAncestors

`func (o *EtherNetworkPortRelationship) GetAncestors() []MoBaseMoRelationship`

GetAncestors returns the Ancestors field if non-nil, zero value otherwise.

### GetAncestorsOk

`func (o *EtherNetworkPortRelationship) GetAncestorsOk() (*[]MoBaseMoRelationship, bool)`

GetAncestorsOk returns a tuple with the Ancestors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestors

`func (o *EtherNetworkPortRelationship) SetAncestors(v []MoBaseMoRelationship)`

SetAncestors sets Ancestors field to given value.

### HasAncestors

`func (o *EtherNetworkPortRelationship) HasAncestors() bool`

HasAncestors returns a boolean if a field has been set.

### SetAncestorsNil

`func (o *EtherNetworkPortRelationship) SetAncestorsNil(b bool)`

 SetAncestorsNil sets the value for Ancestors to be an explicit nil

### UnsetAncestors
`func (o *EtherNetworkPortRelationship) UnsetAncestors()`

UnsetAncestors ensures that no value is present for Ancestors, not even an explicit nil
### GetParent

`func (o *EtherNetworkPortRelationship) GetParent() MoBaseMoRelationship`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *EtherNetworkPortRelationship) GetParentOk() (*MoBaseMoRelationship, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *EtherNetworkPortRelationship) SetParent(v MoBaseMoRelationship)`

SetParent sets Parent field to given value.

### HasParent

`func (o *EtherNetworkPortRelationship) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPermissionResources

`func (o *EtherNetworkPortRelationship) GetPermissionResources() []MoBaseMoRelationship`

GetPermissionResources returns the PermissionResources field if non-nil, zero value otherwise.

### GetPermissionResourcesOk

`func (o *EtherNetworkPortRelationship) GetPermissionResourcesOk() (*[]MoBaseMoRelationship, bool)`

GetPermissionResourcesOk returns a tuple with the PermissionResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionResources

`func (o *EtherNetworkPortRelationship) SetPermissionResources(v []MoBaseMoRelationship)`

SetPermissionResources sets PermissionResources field to given value.

### HasPermissionResources

`func (o *EtherNetworkPortRelationship) HasPermissionResources() bool`

HasPermissionResources returns a boolean if a field has been set.

### SetPermissionResourcesNil

`func (o *EtherNetworkPortRelationship) SetPermissionResourcesNil(b bool)`

 SetPermissionResourcesNil sets the value for PermissionResources to be an explicit nil

### UnsetPermissionResources
`func (o *EtherNetworkPortRelationship) UnsetPermissionResources()`

UnsetPermissionResources ensures that no value is present for PermissionResources, not even an explicit nil
### GetDisplayNames

`func (o *EtherNetworkPortRelationship) GetDisplayNames() map[string][]string`

GetDisplayNames returns the DisplayNames field if non-nil, zero value otherwise.

### GetDisplayNamesOk

`func (o *EtherNetworkPortRelationship) GetDisplayNamesOk() (*map[string][]string, bool)`

GetDisplayNamesOk returns a tuple with the DisplayNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNames

`func (o *EtherNetworkPortRelationship) SetDisplayNames(v map[string][]string)`

SetDisplayNames sets DisplayNames field to given value.

### HasDisplayNames

`func (o *EtherNetworkPortRelationship) HasDisplayNames() bool`

HasDisplayNames returns a boolean if a field has been set.

### SetDisplayNamesNil

`func (o *EtherNetworkPortRelationship) SetDisplayNamesNil(b bool)`

 SetDisplayNamesNil sets the value for DisplayNames to be an explicit nil

### UnsetDisplayNames
`func (o *EtherNetworkPortRelationship) UnsetDisplayNames()`

UnsetDisplayNames ensures that no value is present for DisplayNames, not even an explicit nil
### GetDeviceMoId

`func (o *EtherNetworkPortRelationship) GetDeviceMoId() string`

GetDeviceMoId returns the DeviceMoId field if non-nil, zero value otherwise.

### GetDeviceMoIdOk

`func (o *EtherNetworkPortRelationship) GetDeviceMoIdOk() (*string, bool)`

GetDeviceMoIdOk returns a tuple with the DeviceMoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMoId

`func (o *EtherNetworkPortRelationship) SetDeviceMoId(v string)`

SetDeviceMoId sets DeviceMoId field to given value.

### HasDeviceMoId

`func (o *EtherNetworkPortRelationship) HasDeviceMoId() bool`

HasDeviceMoId returns a boolean if a field has been set.

### GetDn

`func (o *EtherNetworkPortRelationship) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *EtherNetworkPortRelationship) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *EtherNetworkPortRelationship) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *EtherNetworkPortRelationship) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetRn

`func (o *EtherNetworkPortRelationship) GetRn() string`

GetRn returns the Rn field if non-nil, zero value otherwise.

### GetRnOk

`func (o *EtherNetworkPortRelationship) GetRnOk() (*string, bool)`

GetRnOk returns a tuple with the Rn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRn

`func (o *EtherNetworkPortRelationship) SetRn(v string)`

SetRn sets Rn field to given value.

### HasRn

`func (o *EtherNetworkPortRelationship) HasRn() bool`

HasRn returns a boolean if a field has been set.

### GetOperState

`func (o *EtherNetworkPortRelationship) GetOperState() string`

GetOperState returns the OperState field if non-nil, zero value otherwise.

### GetOperStateOk

`func (o *EtherNetworkPortRelationship) GetOperStateOk() (*string, bool)`

GetOperStateOk returns a tuple with the OperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperState

`func (o *EtherNetworkPortRelationship) SetOperState(v string)`

SetOperState sets OperState field to given value.

### HasOperState

`func (o *EtherNetworkPortRelationship) HasOperState() bool`

HasOperState returns a boolean if a field has been set.

### GetAcknowledgedPeerInterface

`func (o *EtherNetworkPortRelationship) GetAcknowledgedPeerInterface() EtherPhysicalPortBaseRelationship`

GetAcknowledgedPeerInterface returns the AcknowledgedPeerInterface field if non-nil, zero value otherwise.

### GetAcknowledgedPeerInterfaceOk

`func (o *EtherNetworkPortRelationship) GetAcknowledgedPeerInterfaceOk() (*EtherPhysicalPortBaseRelationship, bool)`

GetAcknowledgedPeerInterfaceOk returns a tuple with the AcknowledgedPeerInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledgedPeerInterface

`func (o *EtherNetworkPortRelationship) SetAcknowledgedPeerInterface(v EtherPhysicalPortBaseRelationship)`

SetAcknowledgedPeerInterface sets AcknowledgedPeerInterface field to given value.

### HasAcknowledgedPeerInterface

`func (o *EtherNetworkPortRelationship) HasAcknowledgedPeerInterface() bool`

HasAcknowledgedPeerInterface returns a boolean if a field has been set.

### GetPeerInterface

`func (o *EtherNetworkPortRelationship) GetPeerInterface() EtherPhysicalPortBaseRelationship`

GetPeerInterface returns the PeerInterface field if non-nil, zero value otherwise.

### GetPeerInterfaceOk

`func (o *EtherNetworkPortRelationship) GetPeerInterfaceOk() (*EtherPhysicalPortBaseRelationship, bool)`

GetPeerInterfaceOk returns a tuple with the PeerInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerInterface

`func (o *EtherNetworkPortRelationship) SetPeerInterface(v EtherPhysicalPortBaseRelationship)`

SetPeerInterface sets PeerInterface field to given value.

### HasPeerInterface

`func (o *EtherNetworkPortRelationship) HasPeerInterface() bool`

HasPeerInterface returns a boolean if a field has been set.

### GetModuleId

`func (o *EtherNetworkPortRelationship) GetModuleId() int64`

GetModuleId returns the ModuleId field if non-nil, zero value otherwise.

### GetModuleIdOk

`func (o *EtherNetworkPortRelationship) GetModuleIdOk() (*int64, bool)`

GetModuleIdOk returns a tuple with the ModuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleId

`func (o *EtherNetworkPortRelationship) SetModuleId(v int64)`

SetModuleId sets ModuleId field to given value.

### HasModuleId

`func (o *EtherNetworkPortRelationship) HasModuleId() bool`

HasModuleId returns a boolean if a field has been set.

### GetPeerDn

`func (o *EtherNetworkPortRelationship) GetPeerDn() string`

GetPeerDn returns the PeerDn field if non-nil, zero value otherwise.

### GetPeerDnOk

`func (o *EtherNetworkPortRelationship) GetPeerDnOk() (*string, bool)`

GetPeerDnOk returns a tuple with the PeerDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerDn

`func (o *EtherNetworkPortRelationship) SetPeerDn(v string)`

SetPeerDn sets PeerDn field to given value.

### HasPeerDn

`func (o *EtherNetworkPortRelationship) HasPeerDn() bool`

HasPeerDn returns a boolean if a field has been set.

### GetPortId

`func (o *EtherNetworkPortRelationship) GetPortId() int64`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *EtherNetworkPortRelationship) GetPortIdOk() (*int64, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *EtherNetworkPortRelationship) SetPortId(v int64)`

SetPortId sets PortId field to given value.

### HasPortId

`func (o *EtherNetworkPortRelationship) HasPortId() bool`

HasPortId returns a boolean if a field has been set.

### GetSlotId

`func (o *EtherNetworkPortRelationship) GetSlotId() int64`

GetSlotId returns the SlotId field if non-nil, zero value otherwise.

### GetSlotIdOk

`func (o *EtherNetworkPortRelationship) GetSlotIdOk() (*int64, bool)`

GetSlotIdOk returns a tuple with the SlotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotId

`func (o *EtherNetworkPortRelationship) SetSlotId(v int64)`

SetSlotId sets SlotId field to given value.

### HasSlotId

`func (o *EtherNetworkPortRelationship) HasSlotId() bool`

HasSlotId returns a boolean if a field has been set.

### GetSpeed

`func (o *EtherNetworkPortRelationship) GetSpeed() string`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *EtherNetworkPortRelationship) GetSpeedOk() (*string, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *EtherNetworkPortRelationship) SetSpeed(v string)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *EtherNetworkPortRelationship) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetSwitchId

`func (o *EtherNetworkPortRelationship) GetSwitchId() string`

GetSwitchId returns the SwitchId field if non-nil, zero value otherwise.

### GetSwitchIdOk

`func (o *EtherNetworkPortRelationship) GetSwitchIdOk() (*string, bool)`

GetSwitchIdOk returns a tuple with the SwitchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchId

`func (o *EtherNetworkPortRelationship) SetSwitchId(v string)`

SetSwitchId sets SwitchId field to given value.

### HasSwitchId

`func (o *EtherNetworkPortRelationship) HasSwitchId() bool`

HasSwitchId returns a boolean if a field has been set.

### GetEquipmentIoCardBase

`func (o *EtherNetworkPortRelationship) GetEquipmentIoCardBase() EquipmentIoCardBaseRelationship`

GetEquipmentIoCardBase returns the EquipmentIoCardBase field if non-nil, zero value otherwise.

### GetEquipmentIoCardBaseOk

`func (o *EtherNetworkPortRelationship) GetEquipmentIoCardBaseOk() (*EquipmentIoCardBaseRelationship, bool)`

GetEquipmentIoCardBaseOk returns a tuple with the EquipmentIoCardBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentIoCardBase

`func (o *EtherNetworkPortRelationship) SetEquipmentIoCardBase(v EquipmentIoCardBaseRelationship)`

SetEquipmentIoCardBase sets EquipmentIoCardBase field to given value.

### HasEquipmentIoCardBase

`func (o *EtherNetworkPortRelationship) HasEquipmentIoCardBase() bool`

HasEquipmentIoCardBase returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *EtherNetworkPortRelationship) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *EtherNetworkPortRelationship) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *EtherNetworkPortRelationship) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *EtherNetworkPortRelationship) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



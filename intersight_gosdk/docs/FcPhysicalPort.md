# FcPhysicalPort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "fc.PhysicalPort"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "fc.PhysicalPort"]
**AdminSpeed** | Pointer to **string** | Administrator configured Speed applied on the port. | [optional] [readonly] 
**AdminState** | Pointer to **string** | Administratively configured state (enabled/disabled) for this port. | [optional] [readonly] 
**AggregatePortId** | Pointer to **int64** | Breakout port member in the Fabric Interconnect. | [optional] [readonly] [default to 0]
**B2bCredit** | Pointer to **int64** | Buffer to Buffer credits of FC port. | [optional] [readonly] 
**MaxSpeed** | Pointer to **string** | Maximum Speed with which the port operates. | [optional] [readonly] 
**Mode** | Pointer to **string** | Mode information N_proxy, F or E associated to the Fibre Channel port. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the physical port of FC. | [optional] [readonly] 
**OperSpeed** | Pointer to **string** | Operational Speed with which the port operates. | [optional] [readonly] 
**PeerDn** | Pointer to **string** | PeerDn for fibre channel physical port. | [optional] [readonly] 
**PortChannelId** | Pointer to **int64** | Port channel id of FC port channel created on FI switch. | [optional] [readonly] 
**TransceiverType** | Pointer to **string** | Transceiver type of a Fibre Channel port. | [optional] [readonly] 
**Vsan** | Pointer to **int64** | Virtual San that is associated to the port. | [optional] [readonly] 
**Wwn** | Pointer to **string** | World Wide Name of a Fibre Channel port. | [optional] [readonly] 
**EquipmentSwitchCard** | Pointer to [**NullableEquipmentSwitchCardRelationship**](EquipmentSwitchCardRelationship.md) |  | [optional] 
**FcNeighbor** | Pointer to [**NullableFcNeighborRelationship**](FcNeighborRelationship.md) |  | [optional] 
**InventoryDeviceInfo** | Pointer to [**NullableInventoryDeviceInfoRelationship**](InventoryDeviceInfoRelationship.md) |  | [optional] 
**NetworkSupervisorCard** | Pointer to [**NullableNetworkSupervisorCardRelationship**](NetworkSupervisorCardRelationship.md) |  | [optional] 
**PortGroup** | Pointer to [**NullablePortGroupRelationship**](PortGroupRelationship.md) |  | [optional] 
**PortSubGroup** | Pointer to [**NullablePortSubGroupRelationship**](PortSubGroupRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewFcPhysicalPort

`func NewFcPhysicalPort(classId string, objectType string, ) *FcPhysicalPort`

NewFcPhysicalPort instantiates a new FcPhysicalPort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFcPhysicalPortWithDefaults

`func NewFcPhysicalPortWithDefaults() *FcPhysicalPort`

NewFcPhysicalPortWithDefaults instantiates a new FcPhysicalPort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FcPhysicalPort) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FcPhysicalPort) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FcPhysicalPort) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FcPhysicalPort) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FcPhysicalPort) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FcPhysicalPort) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdminSpeed

`func (o *FcPhysicalPort) GetAdminSpeed() string`

GetAdminSpeed returns the AdminSpeed field if non-nil, zero value otherwise.

### GetAdminSpeedOk

`func (o *FcPhysicalPort) GetAdminSpeedOk() (*string, bool)`

GetAdminSpeedOk returns a tuple with the AdminSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminSpeed

`func (o *FcPhysicalPort) SetAdminSpeed(v string)`

SetAdminSpeed sets AdminSpeed field to given value.

### HasAdminSpeed

`func (o *FcPhysicalPort) HasAdminSpeed() bool`

HasAdminSpeed returns a boolean if a field has been set.

### GetAdminState

`func (o *FcPhysicalPort) GetAdminState() string`

GetAdminState returns the AdminState field if non-nil, zero value otherwise.

### GetAdminStateOk

`func (o *FcPhysicalPort) GetAdminStateOk() (*string, bool)`

GetAdminStateOk returns a tuple with the AdminState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminState

`func (o *FcPhysicalPort) SetAdminState(v string)`

SetAdminState sets AdminState field to given value.

### HasAdminState

`func (o *FcPhysicalPort) HasAdminState() bool`

HasAdminState returns a boolean if a field has been set.

### GetAggregatePortId

`func (o *FcPhysicalPort) GetAggregatePortId() int64`

GetAggregatePortId returns the AggregatePortId field if non-nil, zero value otherwise.

### GetAggregatePortIdOk

`func (o *FcPhysicalPort) GetAggregatePortIdOk() (*int64, bool)`

GetAggregatePortIdOk returns a tuple with the AggregatePortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregatePortId

`func (o *FcPhysicalPort) SetAggregatePortId(v int64)`

SetAggregatePortId sets AggregatePortId field to given value.

### HasAggregatePortId

`func (o *FcPhysicalPort) HasAggregatePortId() bool`

HasAggregatePortId returns a boolean if a field has been set.

### GetB2bCredit

`func (o *FcPhysicalPort) GetB2bCredit() int64`

GetB2bCredit returns the B2bCredit field if non-nil, zero value otherwise.

### GetB2bCreditOk

`func (o *FcPhysicalPort) GetB2bCreditOk() (*int64, bool)`

GetB2bCreditOk returns a tuple with the B2bCredit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetB2bCredit

`func (o *FcPhysicalPort) SetB2bCredit(v int64)`

SetB2bCredit sets B2bCredit field to given value.

### HasB2bCredit

`func (o *FcPhysicalPort) HasB2bCredit() bool`

HasB2bCredit returns a boolean if a field has been set.

### GetMaxSpeed

`func (o *FcPhysicalPort) GetMaxSpeed() string`

GetMaxSpeed returns the MaxSpeed field if non-nil, zero value otherwise.

### GetMaxSpeedOk

`func (o *FcPhysicalPort) GetMaxSpeedOk() (*string, bool)`

GetMaxSpeedOk returns a tuple with the MaxSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSpeed

`func (o *FcPhysicalPort) SetMaxSpeed(v string)`

SetMaxSpeed sets MaxSpeed field to given value.

### HasMaxSpeed

`func (o *FcPhysicalPort) HasMaxSpeed() bool`

HasMaxSpeed returns a boolean if a field has been set.

### GetMode

`func (o *FcPhysicalPort) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *FcPhysicalPort) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *FcPhysicalPort) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *FcPhysicalPort) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetName

`func (o *FcPhysicalPort) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FcPhysicalPort) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FcPhysicalPort) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FcPhysicalPort) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOperSpeed

`func (o *FcPhysicalPort) GetOperSpeed() string`

GetOperSpeed returns the OperSpeed field if non-nil, zero value otherwise.

### GetOperSpeedOk

`func (o *FcPhysicalPort) GetOperSpeedOk() (*string, bool)`

GetOperSpeedOk returns a tuple with the OperSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperSpeed

`func (o *FcPhysicalPort) SetOperSpeed(v string)`

SetOperSpeed sets OperSpeed field to given value.

### HasOperSpeed

`func (o *FcPhysicalPort) HasOperSpeed() bool`

HasOperSpeed returns a boolean if a field has been set.

### GetPeerDn

`func (o *FcPhysicalPort) GetPeerDn() string`

GetPeerDn returns the PeerDn field if non-nil, zero value otherwise.

### GetPeerDnOk

`func (o *FcPhysicalPort) GetPeerDnOk() (*string, bool)`

GetPeerDnOk returns a tuple with the PeerDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerDn

`func (o *FcPhysicalPort) SetPeerDn(v string)`

SetPeerDn sets PeerDn field to given value.

### HasPeerDn

`func (o *FcPhysicalPort) HasPeerDn() bool`

HasPeerDn returns a boolean if a field has been set.

### GetPortChannelId

`func (o *FcPhysicalPort) GetPortChannelId() int64`

GetPortChannelId returns the PortChannelId field if non-nil, zero value otherwise.

### GetPortChannelIdOk

`func (o *FcPhysicalPort) GetPortChannelIdOk() (*int64, bool)`

GetPortChannelIdOk returns a tuple with the PortChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortChannelId

`func (o *FcPhysicalPort) SetPortChannelId(v int64)`

SetPortChannelId sets PortChannelId field to given value.

### HasPortChannelId

`func (o *FcPhysicalPort) HasPortChannelId() bool`

HasPortChannelId returns a boolean if a field has been set.

### GetTransceiverType

`func (o *FcPhysicalPort) GetTransceiverType() string`

GetTransceiverType returns the TransceiverType field if non-nil, zero value otherwise.

### GetTransceiverTypeOk

`func (o *FcPhysicalPort) GetTransceiverTypeOk() (*string, bool)`

GetTransceiverTypeOk returns a tuple with the TransceiverType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransceiverType

`func (o *FcPhysicalPort) SetTransceiverType(v string)`

SetTransceiverType sets TransceiverType field to given value.

### HasTransceiverType

`func (o *FcPhysicalPort) HasTransceiverType() bool`

HasTransceiverType returns a boolean if a field has been set.

### GetVsan

`func (o *FcPhysicalPort) GetVsan() int64`

GetVsan returns the Vsan field if non-nil, zero value otherwise.

### GetVsanOk

`func (o *FcPhysicalPort) GetVsanOk() (*int64, bool)`

GetVsanOk returns a tuple with the Vsan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsan

`func (o *FcPhysicalPort) SetVsan(v int64)`

SetVsan sets Vsan field to given value.

### HasVsan

`func (o *FcPhysicalPort) HasVsan() bool`

HasVsan returns a boolean if a field has been set.

### GetWwn

`func (o *FcPhysicalPort) GetWwn() string`

GetWwn returns the Wwn field if non-nil, zero value otherwise.

### GetWwnOk

`func (o *FcPhysicalPort) GetWwnOk() (*string, bool)`

GetWwnOk returns a tuple with the Wwn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWwn

`func (o *FcPhysicalPort) SetWwn(v string)`

SetWwn sets Wwn field to given value.

### HasWwn

`func (o *FcPhysicalPort) HasWwn() bool`

HasWwn returns a boolean if a field has been set.

### GetEquipmentSwitchCard

`func (o *FcPhysicalPort) GetEquipmentSwitchCard() EquipmentSwitchCardRelationship`

GetEquipmentSwitchCard returns the EquipmentSwitchCard field if non-nil, zero value otherwise.

### GetEquipmentSwitchCardOk

`func (o *FcPhysicalPort) GetEquipmentSwitchCardOk() (*EquipmentSwitchCardRelationship, bool)`

GetEquipmentSwitchCardOk returns a tuple with the EquipmentSwitchCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentSwitchCard

`func (o *FcPhysicalPort) SetEquipmentSwitchCard(v EquipmentSwitchCardRelationship)`

SetEquipmentSwitchCard sets EquipmentSwitchCard field to given value.

### HasEquipmentSwitchCard

`func (o *FcPhysicalPort) HasEquipmentSwitchCard() bool`

HasEquipmentSwitchCard returns a boolean if a field has been set.

### SetEquipmentSwitchCardNil

`func (o *FcPhysicalPort) SetEquipmentSwitchCardNil(b bool)`

 SetEquipmentSwitchCardNil sets the value for EquipmentSwitchCard to be an explicit nil

### UnsetEquipmentSwitchCard
`func (o *FcPhysicalPort) UnsetEquipmentSwitchCard()`

UnsetEquipmentSwitchCard ensures that no value is present for EquipmentSwitchCard, not even an explicit nil
### GetFcNeighbor

`func (o *FcPhysicalPort) GetFcNeighbor() FcNeighborRelationship`

GetFcNeighbor returns the FcNeighbor field if non-nil, zero value otherwise.

### GetFcNeighborOk

`func (o *FcPhysicalPort) GetFcNeighborOk() (*FcNeighborRelationship, bool)`

GetFcNeighborOk returns a tuple with the FcNeighbor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcNeighbor

`func (o *FcPhysicalPort) SetFcNeighbor(v FcNeighborRelationship)`

SetFcNeighbor sets FcNeighbor field to given value.

### HasFcNeighbor

`func (o *FcPhysicalPort) HasFcNeighbor() bool`

HasFcNeighbor returns a boolean if a field has been set.

### SetFcNeighborNil

`func (o *FcPhysicalPort) SetFcNeighborNil(b bool)`

 SetFcNeighborNil sets the value for FcNeighbor to be an explicit nil

### UnsetFcNeighbor
`func (o *FcPhysicalPort) UnsetFcNeighbor()`

UnsetFcNeighbor ensures that no value is present for FcNeighbor, not even an explicit nil
### GetInventoryDeviceInfo

`func (o *FcPhysicalPort) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *FcPhysicalPort) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *FcPhysicalPort) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *FcPhysicalPort) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### SetInventoryDeviceInfoNil

`func (o *FcPhysicalPort) SetInventoryDeviceInfoNil(b bool)`

 SetInventoryDeviceInfoNil sets the value for InventoryDeviceInfo to be an explicit nil

### UnsetInventoryDeviceInfo
`func (o *FcPhysicalPort) UnsetInventoryDeviceInfo()`

UnsetInventoryDeviceInfo ensures that no value is present for InventoryDeviceInfo, not even an explicit nil
### GetNetworkSupervisorCard

`func (o *FcPhysicalPort) GetNetworkSupervisorCard() NetworkSupervisorCardRelationship`

GetNetworkSupervisorCard returns the NetworkSupervisorCard field if non-nil, zero value otherwise.

### GetNetworkSupervisorCardOk

`func (o *FcPhysicalPort) GetNetworkSupervisorCardOk() (*NetworkSupervisorCardRelationship, bool)`

GetNetworkSupervisorCardOk returns a tuple with the NetworkSupervisorCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkSupervisorCard

`func (o *FcPhysicalPort) SetNetworkSupervisorCard(v NetworkSupervisorCardRelationship)`

SetNetworkSupervisorCard sets NetworkSupervisorCard field to given value.

### HasNetworkSupervisorCard

`func (o *FcPhysicalPort) HasNetworkSupervisorCard() bool`

HasNetworkSupervisorCard returns a boolean if a field has been set.

### SetNetworkSupervisorCardNil

`func (o *FcPhysicalPort) SetNetworkSupervisorCardNil(b bool)`

 SetNetworkSupervisorCardNil sets the value for NetworkSupervisorCard to be an explicit nil

### UnsetNetworkSupervisorCard
`func (o *FcPhysicalPort) UnsetNetworkSupervisorCard()`

UnsetNetworkSupervisorCard ensures that no value is present for NetworkSupervisorCard, not even an explicit nil
### GetPortGroup

`func (o *FcPhysicalPort) GetPortGroup() PortGroupRelationship`

GetPortGroup returns the PortGroup field if non-nil, zero value otherwise.

### GetPortGroupOk

`func (o *FcPhysicalPort) GetPortGroupOk() (*PortGroupRelationship, bool)`

GetPortGroupOk returns a tuple with the PortGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortGroup

`func (o *FcPhysicalPort) SetPortGroup(v PortGroupRelationship)`

SetPortGroup sets PortGroup field to given value.

### HasPortGroup

`func (o *FcPhysicalPort) HasPortGroup() bool`

HasPortGroup returns a boolean if a field has been set.

### SetPortGroupNil

`func (o *FcPhysicalPort) SetPortGroupNil(b bool)`

 SetPortGroupNil sets the value for PortGroup to be an explicit nil

### UnsetPortGroup
`func (o *FcPhysicalPort) UnsetPortGroup()`

UnsetPortGroup ensures that no value is present for PortGroup, not even an explicit nil
### GetPortSubGroup

`func (o *FcPhysicalPort) GetPortSubGroup() PortSubGroupRelationship`

GetPortSubGroup returns the PortSubGroup field if non-nil, zero value otherwise.

### GetPortSubGroupOk

`func (o *FcPhysicalPort) GetPortSubGroupOk() (*PortSubGroupRelationship, bool)`

GetPortSubGroupOk returns a tuple with the PortSubGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortSubGroup

`func (o *FcPhysicalPort) SetPortSubGroup(v PortSubGroupRelationship)`

SetPortSubGroup sets PortSubGroup field to given value.

### HasPortSubGroup

`func (o *FcPhysicalPort) HasPortSubGroup() bool`

HasPortSubGroup returns a boolean if a field has been set.

### SetPortSubGroupNil

`func (o *FcPhysicalPort) SetPortSubGroupNil(b bool)`

 SetPortSubGroupNil sets the value for PortSubGroup to be an explicit nil

### UnsetPortSubGroup
`func (o *FcPhysicalPort) UnsetPortSubGroup()`

UnsetPortSubGroup ensures that no value is present for PortSubGroup, not even an explicit nil
### GetRegisteredDevice

`func (o *FcPhysicalPort) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *FcPhysicalPort) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *FcPhysicalPort) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *FcPhysicalPort) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *FcPhysicalPort) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *FcPhysicalPort) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



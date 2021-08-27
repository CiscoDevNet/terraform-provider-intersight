# EquipmentSwitchCard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "equipment.SwitchCard"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "equipment.SwitchCard"]
**Description** | Pointer to **string** | Detailed description of this switch hardware. | [optional] [readonly] 
**EthernetSwitchingMode** | Pointer to **string** | The user configured Ethernet switching mode for this switch (End-Host or Switch). * &#x60;end-host&#x60; - In end-host mode, the fabric interconnects appear to the upstream devices as end hosts with multiple links.In this mode, the switch does not run Spanning Tree Protocol and avoids loops by following a set of rules for traffic forwarding.In case of ethernet switching mode - Ethernet end-host mode is also known as Ethernet host virtualizer. * &#x60;switch&#x60; - In switch mode, the switch runs Spanning Tree Protocol to avoid loops, and broadcast and multicast packets are handled in the traditional way.This is the traditional switch mode. | [optional] [readonly] [default to "end-host"]
**FcSwitchingMode** | Pointer to **string** | The user configured FC switching mode for this switch (End-Host or Switch). * &#x60;end-host&#x60; - In end-host mode, the fabric interconnects appear to the upstream devices as end hosts with multiple links.In this mode, the switch does not run Spanning Tree Protocol and avoids loops by following a set of rules for traffic forwarding.In case of ethernet switching mode - Ethernet end-host mode is also known as Ethernet host virtualizer. * &#x60;switch&#x60; - In switch mode, the switch runs Spanning Tree Protocol to avoid loops, and broadcast and multicast packets are handled in the traditional way.This is the traditional switch mode. | [optional] [readonly] [default to "end-host"]
**NumPorts** | Pointer to **int64** | Number of ports present in this switch hardware. | [optional] [readonly] 
**OutOfBandIpAddress** | Pointer to **string** | Field specifies this Switch&#39;s Out-of-band IP address. | [optional] [readonly] 
**OutOfBandIpGateway** | Pointer to **string** | Field specifies this Switch&#39;s default gateway for the out-of-band management interface. | [optional] [readonly] 
**OutOfBandIpMask** | Pointer to **string** | Field specifies the Netmask for this Switch&#39;s Out-of-band IP address. | [optional] 
**SlotId** | Pointer to **int64** | Slot identifier of the local Switch slot Interface. | [optional] [readonly] 
**State** | Pointer to **string** | Operational state of the switch hardware. | [optional] [readonly] 
**SwitchId** | Pointer to **string** | Switch Identifier that is local to a cluster. | [optional] [readonly] 
**Thermal** | Pointer to **string** | The Thermal status of the fabric interconnect. * &#x60;unknown&#x60; - The default state of the sensor (in case no data is received). * &#x60;ok&#x60; - State of the sensor indicating the sensor&#39;s temperature range is okay. * &#x60;upper-non-recoverable&#x60; - State of the sensor indicating that the temperature is extremely high above normal range. * &#x60;upper-critical&#x60; - State of the sensor indicating that the temperature is above normal range. * &#x60;upper-non-critical&#x60; - State of the sensor indicating that the temperature is a little above the normal range. * &#x60;lower-non-critical&#x60; - State of the sensor indicating that the temperature is a little below the normal range. * &#x60;lower-critical&#x60; - State of the sensor indicating that the temperature is below normal range. * &#x60;lower-non-recoverable&#x60; - State of the sensor indicating that the temperature is extremely below normal range. | [optional] [default to "unknown"]
**FcPortChannels** | Pointer to [**[]FcPortChannelRelationship**](FcPortChannelRelationship.md) | An array of relationships to fcPortChannel resources. | [optional] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](InventoryDeviceInfoRelationship.md) |  | [optional] 
**NetworkElement** | Pointer to [**NetworkElementRelationship**](NetworkElementRelationship.md) |  | [optional] 
**PortChannels** | Pointer to [**[]EtherPortChannelRelationship**](EtherPortChannelRelationship.md) | An array of relationships to etherPortChannel resources. | [optional] 
**PortGroups** | Pointer to [**[]PortGroupRelationship**](PortGroupRelationship.md) | An array of relationships to portGroup resources. | [optional] [readonly] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewEquipmentSwitchCard

`func NewEquipmentSwitchCard(classId string, objectType string, ) *EquipmentSwitchCard`

NewEquipmentSwitchCard instantiates a new EquipmentSwitchCard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquipmentSwitchCardWithDefaults

`func NewEquipmentSwitchCardWithDefaults() *EquipmentSwitchCard`

NewEquipmentSwitchCardWithDefaults instantiates a new EquipmentSwitchCard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *EquipmentSwitchCard) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *EquipmentSwitchCard) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *EquipmentSwitchCard) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *EquipmentSwitchCard) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *EquipmentSwitchCard) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *EquipmentSwitchCard) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDescription

`func (o *EquipmentSwitchCard) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EquipmentSwitchCard) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EquipmentSwitchCard) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EquipmentSwitchCard) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEthernetSwitchingMode

`func (o *EquipmentSwitchCard) GetEthernetSwitchingMode() string`

GetEthernetSwitchingMode returns the EthernetSwitchingMode field if non-nil, zero value otherwise.

### GetEthernetSwitchingModeOk

`func (o *EquipmentSwitchCard) GetEthernetSwitchingModeOk() (*string, bool)`

GetEthernetSwitchingModeOk returns a tuple with the EthernetSwitchingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthernetSwitchingMode

`func (o *EquipmentSwitchCard) SetEthernetSwitchingMode(v string)`

SetEthernetSwitchingMode sets EthernetSwitchingMode field to given value.

### HasEthernetSwitchingMode

`func (o *EquipmentSwitchCard) HasEthernetSwitchingMode() bool`

HasEthernetSwitchingMode returns a boolean if a field has been set.

### GetFcSwitchingMode

`func (o *EquipmentSwitchCard) GetFcSwitchingMode() string`

GetFcSwitchingMode returns the FcSwitchingMode field if non-nil, zero value otherwise.

### GetFcSwitchingModeOk

`func (o *EquipmentSwitchCard) GetFcSwitchingModeOk() (*string, bool)`

GetFcSwitchingModeOk returns a tuple with the FcSwitchingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcSwitchingMode

`func (o *EquipmentSwitchCard) SetFcSwitchingMode(v string)`

SetFcSwitchingMode sets FcSwitchingMode field to given value.

### HasFcSwitchingMode

`func (o *EquipmentSwitchCard) HasFcSwitchingMode() bool`

HasFcSwitchingMode returns a boolean if a field has been set.

### GetNumPorts

`func (o *EquipmentSwitchCard) GetNumPorts() int64`

GetNumPorts returns the NumPorts field if non-nil, zero value otherwise.

### GetNumPortsOk

`func (o *EquipmentSwitchCard) GetNumPortsOk() (*int64, bool)`

GetNumPortsOk returns a tuple with the NumPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumPorts

`func (o *EquipmentSwitchCard) SetNumPorts(v int64)`

SetNumPorts sets NumPorts field to given value.

### HasNumPorts

`func (o *EquipmentSwitchCard) HasNumPorts() bool`

HasNumPorts returns a boolean if a field has been set.

### GetOutOfBandIpAddress

`func (o *EquipmentSwitchCard) GetOutOfBandIpAddress() string`

GetOutOfBandIpAddress returns the OutOfBandIpAddress field if non-nil, zero value otherwise.

### GetOutOfBandIpAddressOk

`func (o *EquipmentSwitchCard) GetOutOfBandIpAddressOk() (*string, bool)`

GetOutOfBandIpAddressOk returns a tuple with the OutOfBandIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutOfBandIpAddress

`func (o *EquipmentSwitchCard) SetOutOfBandIpAddress(v string)`

SetOutOfBandIpAddress sets OutOfBandIpAddress field to given value.

### HasOutOfBandIpAddress

`func (o *EquipmentSwitchCard) HasOutOfBandIpAddress() bool`

HasOutOfBandIpAddress returns a boolean if a field has been set.

### GetOutOfBandIpGateway

`func (o *EquipmentSwitchCard) GetOutOfBandIpGateway() string`

GetOutOfBandIpGateway returns the OutOfBandIpGateway field if non-nil, zero value otherwise.

### GetOutOfBandIpGatewayOk

`func (o *EquipmentSwitchCard) GetOutOfBandIpGatewayOk() (*string, bool)`

GetOutOfBandIpGatewayOk returns a tuple with the OutOfBandIpGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutOfBandIpGateway

`func (o *EquipmentSwitchCard) SetOutOfBandIpGateway(v string)`

SetOutOfBandIpGateway sets OutOfBandIpGateway field to given value.

### HasOutOfBandIpGateway

`func (o *EquipmentSwitchCard) HasOutOfBandIpGateway() bool`

HasOutOfBandIpGateway returns a boolean if a field has been set.

### GetOutOfBandIpMask

`func (o *EquipmentSwitchCard) GetOutOfBandIpMask() string`

GetOutOfBandIpMask returns the OutOfBandIpMask field if non-nil, zero value otherwise.

### GetOutOfBandIpMaskOk

`func (o *EquipmentSwitchCard) GetOutOfBandIpMaskOk() (*string, bool)`

GetOutOfBandIpMaskOk returns a tuple with the OutOfBandIpMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutOfBandIpMask

`func (o *EquipmentSwitchCard) SetOutOfBandIpMask(v string)`

SetOutOfBandIpMask sets OutOfBandIpMask field to given value.

### HasOutOfBandIpMask

`func (o *EquipmentSwitchCard) HasOutOfBandIpMask() bool`

HasOutOfBandIpMask returns a boolean if a field has been set.

### GetSlotId

`func (o *EquipmentSwitchCard) GetSlotId() int64`

GetSlotId returns the SlotId field if non-nil, zero value otherwise.

### GetSlotIdOk

`func (o *EquipmentSwitchCard) GetSlotIdOk() (*int64, bool)`

GetSlotIdOk returns a tuple with the SlotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotId

`func (o *EquipmentSwitchCard) SetSlotId(v int64)`

SetSlotId sets SlotId field to given value.

### HasSlotId

`func (o *EquipmentSwitchCard) HasSlotId() bool`

HasSlotId returns a boolean if a field has been set.

### GetState

`func (o *EquipmentSwitchCard) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *EquipmentSwitchCard) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *EquipmentSwitchCard) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *EquipmentSwitchCard) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSwitchId

`func (o *EquipmentSwitchCard) GetSwitchId() string`

GetSwitchId returns the SwitchId field if non-nil, zero value otherwise.

### GetSwitchIdOk

`func (o *EquipmentSwitchCard) GetSwitchIdOk() (*string, bool)`

GetSwitchIdOk returns a tuple with the SwitchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchId

`func (o *EquipmentSwitchCard) SetSwitchId(v string)`

SetSwitchId sets SwitchId field to given value.

### HasSwitchId

`func (o *EquipmentSwitchCard) HasSwitchId() bool`

HasSwitchId returns a boolean if a field has been set.

### GetThermal

`func (o *EquipmentSwitchCard) GetThermal() string`

GetThermal returns the Thermal field if non-nil, zero value otherwise.

### GetThermalOk

`func (o *EquipmentSwitchCard) GetThermalOk() (*string, bool)`

GetThermalOk returns a tuple with the Thermal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThermal

`func (o *EquipmentSwitchCard) SetThermal(v string)`

SetThermal sets Thermal field to given value.

### HasThermal

`func (o *EquipmentSwitchCard) HasThermal() bool`

HasThermal returns a boolean if a field has been set.

### GetFcPortChannels

`func (o *EquipmentSwitchCard) GetFcPortChannels() []FcPortChannelRelationship`

GetFcPortChannels returns the FcPortChannels field if non-nil, zero value otherwise.

### GetFcPortChannelsOk

`func (o *EquipmentSwitchCard) GetFcPortChannelsOk() (*[]FcPortChannelRelationship, bool)`

GetFcPortChannelsOk returns a tuple with the FcPortChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcPortChannels

`func (o *EquipmentSwitchCard) SetFcPortChannels(v []FcPortChannelRelationship)`

SetFcPortChannels sets FcPortChannels field to given value.

### HasFcPortChannels

`func (o *EquipmentSwitchCard) HasFcPortChannels() bool`

HasFcPortChannels returns a boolean if a field has been set.

### SetFcPortChannelsNil

`func (o *EquipmentSwitchCard) SetFcPortChannelsNil(b bool)`

 SetFcPortChannelsNil sets the value for FcPortChannels to be an explicit nil

### UnsetFcPortChannels
`func (o *EquipmentSwitchCard) UnsetFcPortChannels()`

UnsetFcPortChannels ensures that no value is present for FcPortChannels, not even an explicit nil
### GetInventoryDeviceInfo

`func (o *EquipmentSwitchCard) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *EquipmentSwitchCard) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *EquipmentSwitchCard) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *EquipmentSwitchCard) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetNetworkElement

`func (o *EquipmentSwitchCard) GetNetworkElement() NetworkElementRelationship`

GetNetworkElement returns the NetworkElement field if non-nil, zero value otherwise.

### GetNetworkElementOk

`func (o *EquipmentSwitchCard) GetNetworkElementOk() (*NetworkElementRelationship, bool)`

GetNetworkElementOk returns a tuple with the NetworkElement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkElement

`func (o *EquipmentSwitchCard) SetNetworkElement(v NetworkElementRelationship)`

SetNetworkElement sets NetworkElement field to given value.

### HasNetworkElement

`func (o *EquipmentSwitchCard) HasNetworkElement() bool`

HasNetworkElement returns a boolean if a field has been set.

### GetPortChannels

`func (o *EquipmentSwitchCard) GetPortChannels() []EtherPortChannelRelationship`

GetPortChannels returns the PortChannels field if non-nil, zero value otherwise.

### GetPortChannelsOk

`func (o *EquipmentSwitchCard) GetPortChannelsOk() (*[]EtherPortChannelRelationship, bool)`

GetPortChannelsOk returns a tuple with the PortChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortChannels

`func (o *EquipmentSwitchCard) SetPortChannels(v []EtherPortChannelRelationship)`

SetPortChannels sets PortChannels field to given value.

### HasPortChannels

`func (o *EquipmentSwitchCard) HasPortChannels() bool`

HasPortChannels returns a boolean if a field has been set.

### SetPortChannelsNil

`func (o *EquipmentSwitchCard) SetPortChannelsNil(b bool)`

 SetPortChannelsNil sets the value for PortChannels to be an explicit nil

### UnsetPortChannels
`func (o *EquipmentSwitchCard) UnsetPortChannels()`

UnsetPortChannels ensures that no value is present for PortChannels, not even an explicit nil
### GetPortGroups

`func (o *EquipmentSwitchCard) GetPortGroups() []PortGroupRelationship`

GetPortGroups returns the PortGroups field if non-nil, zero value otherwise.

### GetPortGroupsOk

`func (o *EquipmentSwitchCard) GetPortGroupsOk() (*[]PortGroupRelationship, bool)`

GetPortGroupsOk returns a tuple with the PortGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortGroups

`func (o *EquipmentSwitchCard) SetPortGroups(v []PortGroupRelationship)`

SetPortGroups sets PortGroups field to given value.

### HasPortGroups

`func (o *EquipmentSwitchCard) HasPortGroups() bool`

HasPortGroups returns a boolean if a field has been set.

### SetPortGroupsNil

`func (o *EquipmentSwitchCard) SetPortGroupsNil(b bool)`

 SetPortGroupsNil sets the value for PortGroups to be an explicit nil

### UnsetPortGroups
`func (o *EquipmentSwitchCard) UnsetPortGroups()`

UnsetPortGroups ensures that no value is present for PortGroups, not even an explicit nil
### GetRegisteredDevice

`func (o *EquipmentSwitchCard) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *EquipmentSwitchCard) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *EquipmentSwitchCard) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *EquipmentSwitchCard) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



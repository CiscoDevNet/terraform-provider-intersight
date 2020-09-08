# EquipmentSwitchCardAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewEquipmentSwitchCardAllOf

`func NewEquipmentSwitchCardAllOf() *EquipmentSwitchCardAllOf`

NewEquipmentSwitchCardAllOf instantiates a new EquipmentSwitchCardAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquipmentSwitchCardAllOfWithDefaults

`func NewEquipmentSwitchCardAllOfWithDefaults() *EquipmentSwitchCardAllOf`

NewEquipmentSwitchCardAllOfWithDefaults instantiates a new EquipmentSwitchCardAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *EquipmentSwitchCardAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EquipmentSwitchCardAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EquipmentSwitchCardAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EquipmentSwitchCardAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetNumPorts

`func (o *EquipmentSwitchCardAllOf) GetNumPorts() int64`

GetNumPorts returns the NumPorts field if non-nil, zero value otherwise.

### GetNumPortsOk

`func (o *EquipmentSwitchCardAllOf) GetNumPortsOk() (*int64, bool)`

GetNumPortsOk returns a tuple with the NumPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumPorts

`func (o *EquipmentSwitchCardAllOf) SetNumPorts(v int64)`

SetNumPorts sets NumPorts field to given value.

### HasNumPorts

`func (o *EquipmentSwitchCardAllOf) HasNumPorts() bool`

HasNumPorts returns a boolean if a field has been set.

### GetOutOfBandIpAddress

`func (o *EquipmentSwitchCardAllOf) GetOutOfBandIpAddress() string`

GetOutOfBandIpAddress returns the OutOfBandIpAddress field if non-nil, zero value otherwise.

### GetOutOfBandIpAddressOk

`func (o *EquipmentSwitchCardAllOf) GetOutOfBandIpAddressOk() (*string, bool)`

GetOutOfBandIpAddressOk returns a tuple with the OutOfBandIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutOfBandIpAddress

`func (o *EquipmentSwitchCardAllOf) SetOutOfBandIpAddress(v string)`

SetOutOfBandIpAddress sets OutOfBandIpAddress field to given value.

### HasOutOfBandIpAddress

`func (o *EquipmentSwitchCardAllOf) HasOutOfBandIpAddress() bool`

HasOutOfBandIpAddress returns a boolean if a field has been set.

### GetOutOfBandIpGateway

`func (o *EquipmentSwitchCardAllOf) GetOutOfBandIpGateway() string`

GetOutOfBandIpGateway returns the OutOfBandIpGateway field if non-nil, zero value otherwise.

### GetOutOfBandIpGatewayOk

`func (o *EquipmentSwitchCardAllOf) GetOutOfBandIpGatewayOk() (*string, bool)`

GetOutOfBandIpGatewayOk returns a tuple with the OutOfBandIpGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutOfBandIpGateway

`func (o *EquipmentSwitchCardAllOf) SetOutOfBandIpGateway(v string)`

SetOutOfBandIpGateway sets OutOfBandIpGateway field to given value.

### HasOutOfBandIpGateway

`func (o *EquipmentSwitchCardAllOf) HasOutOfBandIpGateway() bool`

HasOutOfBandIpGateway returns a boolean if a field has been set.

### GetPresence

`func (o *EquipmentSwitchCardAllOf) GetPresence() string`

GetPresence returns the Presence field if non-nil, zero value otherwise.

### GetPresenceOk

`func (o *EquipmentSwitchCardAllOf) GetPresenceOk() (*string, bool)`

GetPresenceOk returns a tuple with the Presence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresence

`func (o *EquipmentSwitchCardAllOf) SetPresence(v string)`

SetPresence sets Presence field to given value.

### HasPresence

`func (o *EquipmentSwitchCardAllOf) HasPresence() bool`

HasPresence returns a boolean if a field has been set.

### GetSlotId

`func (o *EquipmentSwitchCardAllOf) GetSlotId() int64`

GetSlotId returns the SlotId field if non-nil, zero value otherwise.

### GetSlotIdOk

`func (o *EquipmentSwitchCardAllOf) GetSlotIdOk() (*int64, bool)`

GetSlotIdOk returns a tuple with the SlotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotId

`func (o *EquipmentSwitchCardAllOf) SetSlotId(v int64)`

SetSlotId sets SlotId field to given value.

### HasSlotId

`func (o *EquipmentSwitchCardAllOf) HasSlotId() bool`

HasSlotId returns a boolean if a field has been set.

### GetState

`func (o *EquipmentSwitchCardAllOf) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *EquipmentSwitchCardAllOf) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *EquipmentSwitchCardAllOf) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *EquipmentSwitchCardAllOf) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSwitchId

`func (o *EquipmentSwitchCardAllOf) GetSwitchId() string`

GetSwitchId returns the SwitchId field if non-nil, zero value otherwise.

### GetSwitchIdOk

`func (o *EquipmentSwitchCardAllOf) GetSwitchIdOk() (*string, bool)`

GetSwitchIdOk returns a tuple with the SwitchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchId

`func (o *EquipmentSwitchCardAllOf) SetSwitchId(v string)`

SetSwitchId sets SwitchId field to given value.

### HasSwitchId

`func (o *EquipmentSwitchCardAllOf) HasSwitchId() bool`

HasSwitchId returns a boolean if a field has been set.

### GetFcPortChannels

`func (o *EquipmentSwitchCardAllOf) GetFcPortChannels() []FcPortChannelRelationship`

GetFcPortChannels returns the FcPortChannels field if non-nil, zero value otherwise.

### GetFcPortChannelsOk

`func (o *EquipmentSwitchCardAllOf) GetFcPortChannelsOk() (*[]FcPortChannelRelationship, bool)`

GetFcPortChannelsOk returns a tuple with the FcPortChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcPortChannels

`func (o *EquipmentSwitchCardAllOf) SetFcPortChannels(v []FcPortChannelRelationship)`

SetFcPortChannels sets FcPortChannels field to given value.

### HasFcPortChannels

`func (o *EquipmentSwitchCardAllOf) HasFcPortChannels() bool`

HasFcPortChannels returns a boolean if a field has been set.

### SetFcPortChannelsNil

`func (o *EquipmentSwitchCardAllOf) SetFcPortChannelsNil(b bool)`

 SetFcPortChannelsNil sets the value for FcPortChannels to be an explicit nil

### UnsetFcPortChannels
`func (o *EquipmentSwitchCardAllOf) UnsetFcPortChannels()`

UnsetFcPortChannels ensures that no value is present for FcPortChannels, not even an explicit nil
### GetInventoryDeviceInfo

`func (o *EquipmentSwitchCardAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *EquipmentSwitchCardAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *EquipmentSwitchCardAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *EquipmentSwitchCardAllOf) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetNetworkElement

`func (o *EquipmentSwitchCardAllOf) GetNetworkElement() NetworkElementRelationship`

GetNetworkElement returns the NetworkElement field if non-nil, zero value otherwise.

### GetNetworkElementOk

`func (o *EquipmentSwitchCardAllOf) GetNetworkElementOk() (*NetworkElementRelationship, bool)`

GetNetworkElementOk returns a tuple with the NetworkElement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkElement

`func (o *EquipmentSwitchCardAllOf) SetNetworkElement(v NetworkElementRelationship)`

SetNetworkElement sets NetworkElement field to given value.

### HasNetworkElement

`func (o *EquipmentSwitchCardAllOf) HasNetworkElement() bool`

HasNetworkElement returns a boolean if a field has been set.

### GetPortChannels

`func (o *EquipmentSwitchCardAllOf) GetPortChannels() []EtherPortChannelRelationship`

GetPortChannels returns the PortChannels field if non-nil, zero value otherwise.

### GetPortChannelsOk

`func (o *EquipmentSwitchCardAllOf) GetPortChannelsOk() (*[]EtherPortChannelRelationship, bool)`

GetPortChannelsOk returns a tuple with the PortChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortChannels

`func (o *EquipmentSwitchCardAllOf) SetPortChannels(v []EtherPortChannelRelationship)`

SetPortChannels sets PortChannels field to given value.

### HasPortChannels

`func (o *EquipmentSwitchCardAllOf) HasPortChannels() bool`

HasPortChannels returns a boolean if a field has been set.

### SetPortChannelsNil

`func (o *EquipmentSwitchCardAllOf) SetPortChannelsNil(b bool)`

 SetPortChannelsNil sets the value for PortChannels to be an explicit nil

### UnsetPortChannels
`func (o *EquipmentSwitchCardAllOf) UnsetPortChannels()`

UnsetPortChannels ensures that no value is present for PortChannels, not even an explicit nil
### GetPortGroups

`func (o *EquipmentSwitchCardAllOf) GetPortGroups() []PortGroupRelationship`

GetPortGroups returns the PortGroups field if non-nil, zero value otherwise.

### GetPortGroupsOk

`func (o *EquipmentSwitchCardAllOf) GetPortGroupsOk() (*[]PortGroupRelationship, bool)`

GetPortGroupsOk returns a tuple with the PortGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortGroups

`func (o *EquipmentSwitchCardAllOf) SetPortGroups(v []PortGroupRelationship)`

SetPortGroups sets PortGroups field to given value.

### HasPortGroups

`func (o *EquipmentSwitchCardAllOf) HasPortGroups() bool`

HasPortGroups returns a boolean if a field has been set.

### SetPortGroupsNil

`func (o *EquipmentSwitchCardAllOf) SetPortGroupsNil(b bool)`

 SetPortGroupsNil sets the value for PortGroups to be an explicit nil

### UnsetPortGroups
`func (o *EquipmentSwitchCardAllOf) UnsetPortGroups()`

UnsetPortGroups ensures that no value is present for PortGroups, not even an explicit nil
### GetRegisteredDevice

`func (o *EquipmentSwitchCardAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *EquipmentSwitchCardAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *EquipmentSwitchCardAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *EquipmentSwitchCardAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



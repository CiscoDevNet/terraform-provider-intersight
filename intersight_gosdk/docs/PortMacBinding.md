# PortMacBinding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AggregatePortId** | Pointer to **int64** | Aggregate Port ID of the local Switch Interface. | [optional] 
**ChassisId** | Pointer to **int64** | Chassis/FEX device idetifier that is local to a cluster. | [optional] 
**DeviceMac** | Pointer to **string** | Device ID value that is advertised and available as a part of LLDP TLV. | [optional] 
**PortId** | Pointer to **int64** | Port ID of the local Switch Interface. | [optional] 
**PortMac** | Pointer to **string** | Port ID value that is advertised and available as a part of LLDP TLV. | [optional] 
**SlotId** | Pointer to **int64** | Slot ID of the local Switch slot Interface. | [optional] 
**SwitchId** | Pointer to **int64** | Switch Identifier that is local to a cluster. | [optional] 
**NetworkElement** | Pointer to [**NetworkElementRelationship**](network.Element.Relationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewPortMacBinding

`func NewPortMacBinding() *PortMacBinding`

NewPortMacBinding instantiates a new PortMacBinding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortMacBindingWithDefaults

`func NewPortMacBindingWithDefaults() *PortMacBinding`

NewPortMacBindingWithDefaults instantiates a new PortMacBinding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregatePortId

`func (o *PortMacBinding) GetAggregatePortId() int64`

GetAggregatePortId returns the AggregatePortId field if non-nil, zero value otherwise.

### GetAggregatePortIdOk

`func (o *PortMacBinding) GetAggregatePortIdOk() (*int64, bool)`

GetAggregatePortIdOk returns a tuple with the AggregatePortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregatePortId

`func (o *PortMacBinding) SetAggregatePortId(v int64)`

SetAggregatePortId sets AggregatePortId field to given value.

### HasAggregatePortId

`func (o *PortMacBinding) HasAggregatePortId() bool`

HasAggregatePortId returns a boolean if a field has been set.

### GetChassisId

`func (o *PortMacBinding) GetChassisId() int64`

GetChassisId returns the ChassisId field if non-nil, zero value otherwise.

### GetChassisIdOk

`func (o *PortMacBinding) GetChassisIdOk() (*int64, bool)`

GetChassisIdOk returns a tuple with the ChassisId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassisId

`func (o *PortMacBinding) SetChassisId(v int64)`

SetChassisId sets ChassisId field to given value.

### HasChassisId

`func (o *PortMacBinding) HasChassisId() bool`

HasChassisId returns a boolean if a field has been set.

### GetDeviceMac

`func (o *PortMacBinding) GetDeviceMac() string`

GetDeviceMac returns the DeviceMac field if non-nil, zero value otherwise.

### GetDeviceMacOk

`func (o *PortMacBinding) GetDeviceMacOk() (*string, bool)`

GetDeviceMacOk returns a tuple with the DeviceMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMac

`func (o *PortMacBinding) SetDeviceMac(v string)`

SetDeviceMac sets DeviceMac field to given value.

### HasDeviceMac

`func (o *PortMacBinding) HasDeviceMac() bool`

HasDeviceMac returns a boolean if a field has been set.

### GetPortId

`func (o *PortMacBinding) GetPortId() int64`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *PortMacBinding) GetPortIdOk() (*int64, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *PortMacBinding) SetPortId(v int64)`

SetPortId sets PortId field to given value.

### HasPortId

`func (o *PortMacBinding) HasPortId() bool`

HasPortId returns a boolean if a field has been set.

### GetPortMac

`func (o *PortMacBinding) GetPortMac() string`

GetPortMac returns the PortMac field if non-nil, zero value otherwise.

### GetPortMacOk

`func (o *PortMacBinding) GetPortMacOk() (*string, bool)`

GetPortMacOk returns a tuple with the PortMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortMac

`func (o *PortMacBinding) SetPortMac(v string)`

SetPortMac sets PortMac field to given value.

### HasPortMac

`func (o *PortMacBinding) HasPortMac() bool`

HasPortMac returns a boolean if a field has been set.

### GetSlotId

`func (o *PortMacBinding) GetSlotId() int64`

GetSlotId returns the SlotId field if non-nil, zero value otherwise.

### GetSlotIdOk

`func (o *PortMacBinding) GetSlotIdOk() (*int64, bool)`

GetSlotIdOk returns a tuple with the SlotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotId

`func (o *PortMacBinding) SetSlotId(v int64)`

SetSlotId sets SlotId field to given value.

### HasSlotId

`func (o *PortMacBinding) HasSlotId() bool`

HasSlotId returns a boolean if a field has been set.

### GetSwitchId

`func (o *PortMacBinding) GetSwitchId() int64`

GetSwitchId returns the SwitchId field if non-nil, zero value otherwise.

### GetSwitchIdOk

`func (o *PortMacBinding) GetSwitchIdOk() (*int64, bool)`

GetSwitchIdOk returns a tuple with the SwitchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchId

`func (o *PortMacBinding) SetSwitchId(v int64)`

SetSwitchId sets SwitchId field to given value.

### HasSwitchId

`func (o *PortMacBinding) HasSwitchId() bool`

HasSwitchId returns a boolean if a field has been set.

### GetNetworkElement

`func (o *PortMacBinding) GetNetworkElement() NetworkElementRelationship`

GetNetworkElement returns the NetworkElement field if non-nil, zero value otherwise.

### GetNetworkElementOk

`func (o *PortMacBinding) GetNetworkElementOk() (*NetworkElementRelationship, bool)`

GetNetworkElementOk returns a tuple with the NetworkElement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkElement

`func (o *PortMacBinding) SetNetworkElement(v NetworkElementRelationship)`

SetNetworkElement sets NetworkElement field to given value.

### HasNetworkElement

`func (o *PortMacBinding) HasNetworkElement() bool`

HasNetworkElement returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *PortMacBinding) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *PortMacBinding) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *PortMacBinding) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *PortMacBinding) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



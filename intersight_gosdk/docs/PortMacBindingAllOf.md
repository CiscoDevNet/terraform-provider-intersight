# PortMacBindingAllOf

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

### NewPortMacBindingAllOf

`func NewPortMacBindingAllOf() *PortMacBindingAllOf`

NewPortMacBindingAllOf instantiates a new PortMacBindingAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortMacBindingAllOfWithDefaults

`func NewPortMacBindingAllOfWithDefaults() *PortMacBindingAllOf`

NewPortMacBindingAllOfWithDefaults instantiates a new PortMacBindingAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregatePortId

`func (o *PortMacBindingAllOf) GetAggregatePortId() int64`

GetAggregatePortId returns the AggregatePortId field if non-nil, zero value otherwise.

### GetAggregatePortIdOk

`func (o *PortMacBindingAllOf) GetAggregatePortIdOk() (*int64, bool)`

GetAggregatePortIdOk returns a tuple with the AggregatePortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregatePortId

`func (o *PortMacBindingAllOf) SetAggregatePortId(v int64)`

SetAggregatePortId sets AggregatePortId field to given value.

### HasAggregatePortId

`func (o *PortMacBindingAllOf) HasAggregatePortId() bool`

HasAggregatePortId returns a boolean if a field has been set.

### GetChassisId

`func (o *PortMacBindingAllOf) GetChassisId() int64`

GetChassisId returns the ChassisId field if non-nil, zero value otherwise.

### GetChassisIdOk

`func (o *PortMacBindingAllOf) GetChassisIdOk() (*int64, bool)`

GetChassisIdOk returns a tuple with the ChassisId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassisId

`func (o *PortMacBindingAllOf) SetChassisId(v int64)`

SetChassisId sets ChassisId field to given value.

### HasChassisId

`func (o *PortMacBindingAllOf) HasChassisId() bool`

HasChassisId returns a boolean if a field has been set.

### GetDeviceMac

`func (o *PortMacBindingAllOf) GetDeviceMac() string`

GetDeviceMac returns the DeviceMac field if non-nil, zero value otherwise.

### GetDeviceMacOk

`func (o *PortMacBindingAllOf) GetDeviceMacOk() (*string, bool)`

GetDeviceMacOk returns a tuple with the DeviceMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMac

`func (o *PortMacBindingAllOf) SetDeviceMac(v string)`

SetDeviceMac sets DeviceMac field to given value.

### HasDeviceMac

`func (o *PortMacBindingAllOf) HasDeviceMac() bool`

HasDeviceMac returns a boolean if a field has been set.

### GetPortId

`func (o *PortMacBindingAllOf) GetPortId() int64`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *PortMacBindingAllOf) GetPortIdOk() (*int64, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *PortMacBindingAllOf) SetPortId(v int64)`

SetPortId sets PortId field to given value.

### HasPortId

`func (o *PortMacBindingAllOf) HasPortId() bool`

HasPortId returns a boolean if a field has been set.

### GetPortMac

`func (o *PortMacBindingAllOf) GetPortMac() string`

GetPortMac returns the PortMac field if non-nil, zero value otherwise.

### GetPortMacOk

`func (o *PortMacBindingAllOf) GetPortMacOk() (*string, bool)`

GetPortMacOk returns a tuple with the PortMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortMac

`func (o *PortMacBindingAllOf) SetPortMac(v string)`

SetPortMac sets PortMac field to given value.

### HasPortMac

`func (o *PortMacBindingAllOf) HasPortMac() bool`

HasPortMac returns a boolean if a field has been set.

### GetSlotId

`func (o *PortMacBindingAllOf) GetSlotId() int64`

GetSlotId returns the SlotId field if non-nil, zero value otherwise.

### GetSlotIdOk

`func (o *PortMacBindingAllOf) GetSlotIdOk() (*int64, bool)`

GetSlotIdOk returns a tuple with the SlotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotId

`func (o *PortMacBindingAllOf) SetSlotId(v int64)`

SetSlotId sets SlotId field to given value.

### HasSlotId

`func (o *PortMacBindingAllOf) HasSlotId() bool`

HasSlotId returns a boolean if a field has been set.

### GetSwitchId

`func (o *PortMacBindingAllOf) GetSwitchId() int64`

GetSwitchId returns the SwitchId field if non-nil, zero value otherwise.

### GetSwitchIdOk

`func (o *PortMacBindingAllOf) GetSwitchIdOk() (*int64, bool)`

GetSwitchIdOk returns a tuple with the SwitchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchId

`func (o *PortMacBindingAllOf) SetSwitchId(v int64)`

SetSwitchId sets SwitchId field to given value.

### HasSwitchId

`func (o *PortMacBindingAllOf) HasSwitchId() bool`

HasSwitchId returns a boolean if a field has been set.

### GetNetworkElement

`func (o *PortMacBindingAllOf) GetNetworkElement() NetworkElementRelationship`

GetNetworkElement returns the NetworkElement field if non-nil, zero value otherwise.

### GetNetworkElementOk

`func (o *PortMacBindingAllOf) GetNetworkElementOk() (*NetworkElementRelationship, bool)`

GetNetworkElementOk returns a tuple with the NetworkElement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkElement

`func (o *PortMacBindingAllOf) SetNetworkElement(v NetworkElementRelationship)`

SetNetworkElement sets NetworkElement field to given value.

### HasNetworkElement

`func (o *PortMacBindingAllOf) HasNetworkElement() bool`

HasNetworkElement returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *PortMacBindingAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *PortMacBindingAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *PortMacBindingAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *PortMacBindingAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



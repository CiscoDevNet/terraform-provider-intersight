# FabricPortModeAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomMode** | Pointer to **string** | Custom Port Mode specified for the port range. * &#x60;FibreChannel&#x60; - Fibre Channel Port Types. * &#x60;BreakoutEthernet10G&#x60; - Breakout Ethernet 10G Port Type. * &#x60;BreakoutEthernet25G&#x60; - Breakout Ethernet 25G Port Type. | [optional] [default to "FibreChannel"]
**PortIdEnd** | Pointer to **int64** | Ending range of the Port Identifier. | [optional] 
**PortIdStart** | Pointer to **int64** | Starting range of the Port Identifier. | [optional] 
**SlotId** | Pointer to **int64** | Slot Identifier of the switch. | [optional] 
**PortPolicy** | Pointer to [**FabricPortPolicyRelationship**](fabric.PortPolicy.Relationship.md) |  | [optional] 

## Methods

### NewFabricPortModeAllOf

`func NewFabricPortModeAllOf() *FabricPortModeAllOf`

NewFabricPortModeAllOf instantiates a new FabricPortModeAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricPortModeAllOfWithDefaults

`func NewFabricPortModeAllOfWithDefaults() *FabricPortModeAllOf`

NewFabricPortModeAllOfWithDefaults instantiates a new FabricPortModeAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomMode

`func (o *FabricPortModeAllOf) GetCustomMode() string`

GetCustomMode returns the CustomMode field if non-nil, zero value otherwise.

### GetCustomModeOk

`func (o *FabricPortModeAllOf) GetCustomModeOk() (*string, bool)`

GetCustomModeOk returns a tuple with the CustomMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMode

`func (o *FabricPortModeAllOf) SetCustomMode(v string)`

SetCustomMode sets CustomMode field to given value.

### HasCustomMode

`func (o *FabricPortModeAllOf) HasCustomMode() bool`

HasCustomMode returns a boolean if a field has been set.

### GetPortIdEnd

`func (o *FabricPortModeAllOf) GetPortIdEnd() int64`

GetPortIdEnd returns the PortIdEnd field if non-nil, zero value otherwise.

### GetPortIdEndOk

`func (o *FabricPortModeAllOf) GetPortIdEndOk() (*int64, bool)`

GetPortIdEndOk returns a tuple with the PortIdEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortIdEnd

`func (o *FabricPortModeAllOf) SetPortIdEnd(v int64)`

SetPortIdEnd sets PortIdEnd field to given value.

### HasPortIdEnd

`func (o *FabricPortModeAllOf) HasPortIdEnd() bool`

HasPortIdEnd returns a boolean if a field has been set.

### GetPortIdStart

`func (o *FabricPortModeAllOf) GetPortIdStart() int64`

GetPortIdStart returns the PortIdStart field if non-nil, zero value otherwise.

### GetPortIdStartOk

`func (o *FabricPortModeAllOf) GetPortIdStartOk() (*int64, bool)`

GetPortIdStartOk returns a tuple with the PortIdStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortIdStart

`func (o *FabricPortModeAllOf) SetPortIdStart(v int64)`

SetPortIdStart sets PortIdStart field to given value.

### HasPortIdStart

`func (o *FabricPortModeAllOf) HasPortIdStart() bool`

HasPortIdStart returns a boolean if a field has been set.

### GetSlotId

`func (o *FabricPortModeAllOf) GetSlotId() int64`

GetSlotId returns the SlotId field if non-nil, zero value otherwise.

### GetSlotIdOk

`func (o *FabricPortModeAllOf) GetSlotIdOk() (*int64, bool)`

GetSlotIdOk returns a tuple with the SlotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotId

`func (o *FabricPortModeAllOf) SetSlotId(v int64)`

SetSlotId sets SlotId field to given value.

### HasSlotId

`func (o *FabricPortModeAllOf) HasSlotId() bool`

HasSlotId returns a boolean if a field has been set.

### GetPortPolicy

`func (o *FabricPortModeAllOf) GetPortPolicy() FabricPortPolicyRelationship`

GetPortPolicy returns the PortPolicy field if non-nil, zero value otherwise.

### GetPortPolicyOk

`func (o *FabricPortModeAllOf) GetPortPolicyOk() (*FabricPortPolicyRelationship, bool)`

GetPortPolicyOk returns a tuple with the PortPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortPolicy

`func (o *FabricPortModeAllOf) SetPortPolicy(v FabricPortPolicyRelationship)`

SetPortPolicy sets PortPolicy field to given value.

### HasPortPolicy

`func (o *FabricPortModeAllOf) HasPortPolicy() bool`

HasPortPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



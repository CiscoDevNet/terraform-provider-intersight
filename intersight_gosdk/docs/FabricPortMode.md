# FabricPortMode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomMode** | Pointer to **string** | Custom Port Mode specified for the port range. * &#x60;FibreChannel&#x60; - Fibre Channel Port Types. * &#x60;BreakoutEthernet10G&#x60; - Breakout Ethernet 10G Port Type. * &#x60;BreakoutEthernet25G&#x60; - Breakout Ethernet 25G Port Type. | [optional] [default to "FibreChannel"]
**PortIdEnd** | Pointer to **int64** | Ending range of the Port Identifier. | [optional] 
**PortIdStart** | Pointer to **int64** | Starting range of the Port Identifier. | [optional] 
**SlotId** | Pointer to **int64** | Slot Identifier of the switch. | [optional] 
**PortPolicy** | Pointer to [**FabricPortPolicyRelationship**](fabric.PortPolicy.Relationship.md) |  | [optional] 

## Methods

### NewFabricPortMode

`func NewFabricPortMode() *FabricPortMode`

NewFabricPortMode instantiates a new FabricPortMode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricPortModeWithDefaults

`func NewFabricPortModeWithDefaults() *FabricPortMode`

NewFabricPortModeWithDefaults instantiates a new FabricPortMode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomMode

`func (o *FabricPortMode) GetCustomMode() string`

GetCustomMode returns the CustomMode field if non-nil, zero value otherwise.

### GetCustomModeOk

`func (o *FabricPortMode) GetCustomModeOk() (*string, bool)`

GetCustomModeOk returns a tuple with the CustomMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMode

`func (o *FabricPortMode) SetCustomMode(v string)`

SetCustomMode sets CustomMode field to given value.

### HasCustomMode

`func (o *FabricPortMode) HasCustomMode() bool`

HasCustomMode returns a boolean if a field has been set.

### GetPortIdEnd

`func (o *FabricPortMode) GetPortIdEnd() int64`

GetPortIdEnd returns the PortIdEnd field if non-nil, zero value otherwise.

### GetPortIdEndOk

`func (o *FabricPortMode) GetPortIdEndOk() (*int64, bool)`

GetPortIdEndOk returns a tuple with the PortIdEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortIdEnd

`func (o *FabricPortMode) SetPortIdEnd(v int64)`

SetPortIdEnd sets PortIdEnd field to given value.

### HasPortIdEnd

`func (o *FabricPortMode) HasPortIdEnd() bool`

HasPortIdEnd returns a boolean if a field has been set.

### GetPortIdStart

`func (o *FabricPortMode) GetPortIdStart() int64`

GetPortIdStart returns the PortIdStart field if non-nil, zero value otherwise.

### GetPortIdStartOk

`func (o *FabricPortMode) GetPortIdStartOk() (*int64, bool)`

GetPortIdStartOk returns a tuple with the PortIdStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortIdStart

`func (o *FabricPortMode) SetPortIdStart(v int64)`

SetPortIdStart sets PortIdStart field to given value.

### HasPortIdStart

`func (o *FabricPortMode) HasPortIdStart() bool`

HasPortIdStart returns a boolean if a field has been set.

### GetSlotId

`func (o *FabricPortMode) GetSlotId() int64`

GetSlotId returns the SlotId field if non-nil, zero value otherwise.

### GetSlotIdOk

`func (o *FabricPortMode) GetSlotIdOk() (*int64, bool)`

GetSlotIdOk returns a tuple with the SlotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotId

`func (o *FabricPortMode) SetSlotId(v int64)`

SetSlotId sets SlotId field to given value.

### HasSlotId

`func (o *FabricPortMode) HasSlotId() bool`

HasSlotId returns a boolean if a field has been set.

### GetPortPolicy

`func (o *FabricPortMode) GetPortPolicy() FabricPortPolicyRelationship`

GetPortPolicy returns the PortPolicy field if non-nil, zero value otherwise.

### GetPortPolicyOk

`func (o *FabricPortMode) GetPortPolicyOk() (*FabricPortPolicyRelationship, bool)`

GetPortPolicyOk returns a tuple with the PortPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortPolicy

`func (o *FabricPortMode) SetPortPolicy(v FabricPortPolicyRelationship)`

SetPortPolicy sets PortPolicy field to given value.

### HasPortPolicy

`func (o *FabricPortMode) HasPortPolicy() bool`

HasPortPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



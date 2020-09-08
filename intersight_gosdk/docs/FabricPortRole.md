# FabricPortRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AggregatePortId** | Pointer to **int64** | Breakout port Identifier of the Switch Interface. | [optional] 
**PortId** | Pointer to **int64** | Port Identifier of the Switch/FEX/Chassis Interface. | [optional] 
**SlotId** | Pointer to **int64** | Slot Identifier of the Switch/FEX/Chassis Interface. | [optional] 
**PortPolicy** | Pointer to [**FabricPortPolicyRelationship**](fabric.PortPolicy.Relationship.md) |  | [optional] 

## Methods

### NewFabricPortRole

`func NewFabricPortRole() *FabricPortRole`

NewFabricPortRole instantiates a new FabricPortRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricPortRoleWithDefaults

`func NewFabricPortRoleWithDefaults() *FabricPortRole`

NewFabricPortRoleWithDefaults instantiates a new FabricPortRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregatePortId

`func (o *FabricPortRole) GetAggregatePortId() int64`

GetAggregatePortId returns the AggregatePortId field if non-nil, zero value otherwise.

### GetAggregatePortIdOk

`func (o *FabricPortRole) GetAggregatePortIdOk() (*int64, bool)`

GetAggregatePortIdOk returns a tuple with the AggregatePortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregatePortId

`func (o *FabricPortRole) SetAggregatePortId(v int64)`

SetAggregatePortId sets AggregatePortId field to given value.

### HasAggregatePortId

`func (o *FabricPortRole) HasAggregatePortId() bool`

HasAggregatePortId returns a boolean if a field has been set.

### GetPortId

`func (o *FabricPortRole) GetPortId() int64`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *FabricPortRole) GetPortIdOk() (*int64, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *FabricPortRole) SetPortId(v int64)`

SetPortId sets PortId field to given value.

### HasPortId

`func (o *FabricPortRole) HasPortId() bool`

HasPortId returns a boolean if a field has been set.

### GetSlotId

`func (o *FabricPortRole) GetSlotId() int64`

GetSlotId returns the SlotId field if non-nil, zero value otherwise.

### GetSlotIdOk

`func (o *FabricPortRole) GetSlotIdOk() (*int64, bool)`

GetSlotIdOk returns a tuple with the SlotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotId

`func (o *FabricPortRole) SetSlotId(v int64)`

SetSlotId sets SlotId field to given value.

### HasSlotId

`func (o *FabricPortRole) HasSlotId() bool`

HasSlotId returns a boolean if a field has been set.

### GetPortPolicy

`func (o *FabricPortRole) GetPortPolicy() FabricPortPolicyRelationship`

GetPortPolicy returns the PortPolicy field if non-nil, zero value otherwise.

### GetPortPolicyOk

`func (o *FabricPortRole) GetPortPolicyOk() (*FabricPortPolicyRelationship, bool)`

GetPortPolicyOk returns a tuple with the PortPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortPolicy

`func (o *FabricPortRole) SetPortPolicy(v FabricPortPolicyRelationship)`

SetPortPolicy sets PortPolicy field to given value.

### HasPortPolicy

`func (o *FabricPortRole) HasPortPolicy() bool`

HasPortPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



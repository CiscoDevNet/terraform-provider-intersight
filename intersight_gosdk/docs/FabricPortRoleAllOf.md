# FabricPortRoleAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AggregatePortId** | Pointer to **int64** | Breakout port Identifier of the Switch Interface. | [optional] 
**PortId** | Pointer to **int64** | Port Identifier of the Switch/FEX/Chassis Interface. | [optional] 
**SlotId** | Pointer to **int64** | Slot Identifier of the Switch/FEX/Chassis Interface. | [optional] 
**PortPolicy** | Pointer to [**FabricPortPolicyRelationship**](fabric.PortPolicy.Relationship.md) |  | [optional] 

## Methods

### NewFabricPortRoleAllOf

`func NewFabricPortRoleAllOf() *FabricPortRoleAllOf`

NewFabricPortRoleAllOf instantiates a new FabricPortRoleAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricPortRoleAllOfWithDefaults

`func NewFabricPortRoleAllOfWithDefaults() *FabricPortRoleAllOf`

NewFabricPortRoleAllOfWithDefaults instantiates a new FabricPortRoleAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregatePortId

`func (o *FabricPortRoleAllOf) GetAggregatePortId() int64`

GetAggregatePortId returns the AggregatePortId field if non-nil, zero value otherwise.

### GetAggregatePortIdOk

`func (o *FabricPortRoleAllOf) GetAggregatePortIdOk() (*int64, bool)`

GetAggregatePortIdOk returns a tuple with the AggregatePortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregatePortId

`func (o *FabricPortRoleAllOf) SetAggregatePortId(v int64)`

SetAggregatePortId sets AggregatePortId field to given value.

### HasAggregatePortId

`func (o *FabricPortRoleAllOf) HasAggregatePortId() bool`

HasAggregatePortId returns a boolean if a field has been set.

### GetPortId

`func (o *FabricPortRoleAllOf) GetPortId() int64`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *FabricPortRoleAllOf) GetPortIdOk() (*int64, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *FabricPortRoleAllOf) SetPortId(v int64)`

SetPortId sets PortId field to given value.

### HasPortId

`func (o *FabricPortRoleAllOf) HasPortId() bool`

HasPortId returns a boolean if a field has been set.

### GetSlotId

`func (o *FabricPortRoleAllOf) GetSlotId() int64`

GetSlotId returns the SlotId field if non-nil, zero value otherwise.

### GetSlotIdOk

`func (o *FabricPortRoleAllOf) GetSlotIdOk() (*int64, bool)`

GetSlotIdOk returns a tuple with the SlotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotId

`func (o *FabricPortRoleAllOf) SetSlotId(v int64)`

SetSlotId sets SlotId field to given value.

### HasSlotId

`func (o *FabricPortRoleAllOf) HasSlotId() bool`

HasSlotId returns a boolean if a field has been set.

### GetPortPolicy

`func (o *FabricPortRoleAllOf) GetPortPolicy() FabricPortPolicyRelationship`

GetPortPolicy returns the PortPolicy field if non-nil, zero value otherwise.

### GetPortPolicyOk

`func (o *FabricPortRoleAllOf) GetPortPolicyOk() (*FabricPortPolicyRelationship, bool)`

GetPortPolicyOk returns a tuple with the PortPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortPolicy

`func (o *FabricPortRoleAllOf) SetPortPolicy(v FabricPortPolicyRelationship)`

SetPortPolicy sets PortPolicy field to given value.

### HasPortPolicy

`func (o *FabricPortRoleAllOf) HasPortPolicy() bool`

HasPortPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



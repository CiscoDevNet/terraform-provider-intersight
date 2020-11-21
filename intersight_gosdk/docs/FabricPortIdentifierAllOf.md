# FabricPortIdentifierAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "fabric.PortIdentifier"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "fabric.PortIdentifier"]
**AggregatePortId** | Pointer to **int64** | Breakout port Identifier of the Switch Interface. When a port is not configured as a breakout port, the aggregatePortId is set to 0, and unused. When a port is configured as a breakout port, the &#39;aggregatePortId&#39; port number as labeled on the equipment, e.g. the id of the port on the switch. | [optional] 
**PortId** | Pointer to **int64** | Port Identifier of the Switch/FEX/Chassis Interface. When a port is not configured as a breakout port, the portId is the port number as labeled on the equipment, e.g. the id of the port on the switch, FEX or chassis. When a port is configured as a breakout port, the &#39;portId&#39; represents the port id on the fanout side of the breakout cable. | [optional] 
**SlotId** | Pointer to **int64** | Slot Identifier of the Switch/FEX/Chassis Interface. | [optional] 

## Methods

### NewFabricPortIdentifierAllOf

`func NewFabricPortIdentifierAllOf(classId string, objectType string, ) *FabricPortIdentifierAllOf`

NewFabricPortIdentifierAllOf instantiates a new FabricPortIdentifierAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricPortIdentifierAllOfWithDefaults

`func NewFabricPortIdentifierAllOfWithDefaults() *FabricPortIdentifierAllOf`

NewFabricPortIdentifierAllOfWithDefaults instantiates a new FabricPortIdentifierAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FabricPortIdentifierAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FabricPortIdentifierAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FabricPortIdentifierAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FabricPortIdentifierAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FabricPortIdentifierAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FabricPortIdentifierAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAggregatePortId

`func (o *FabricPortIdentifierAllOf) GetAggregatePortId() int64`

GetAggregatePortId returns the AggregatePortId field if non-nil, zero value otherwise.

### GetAggregatePortIdOk

`func (o *FabricPortIdentifierAllOf) GetAggregatePortIdOk() (*int64, bool)`

GetAggregatePortIdOk returns a tuple with the AggregatePortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregatePortId

`func (o *FabricPortIdentifierAllOf) SetAggregatePortId(v int64)`

SetAggregatePortId sets AggregatePortId field to given value.

### HasAggregatePortId

`func (o *FabricPortIdentifierAllOf) HasAggregatePortId() bool`

HasAggregatePortId returns a boolean if a field has been set.

### GetPortId

`func (o *FabricPortIdentifierAllOf) GetPortId() int64`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *FabricPortIdentifierAllOf) GetPortIdOk() (*int64, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *FabricPortIdentifierAllOf) SetPortId(v int64)`

SetPortId sets PortId field to given value.

### HasPortId

`func (o *FabricPortIdentifierAllOf) HasPortId() bool`

HasPortId returns a boolean if a field has been set.

### GetSlotId

`func (o *FabricPortIdentifierAllOf) GetSlotId() int64`

GetSlotId returns the SlotId field if non-nil, zero value otherwise.

### GetSlotIdOk

`func (o *FabricPortIdentifierAllOf) GetSlotIdOk() (*int64, bool)`

GetSlotIdOk returns a tuple with the SlotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotId

`func (o *FabricPortIdentifierAllOf) SetSlotId(v int64)`

SetSlotId sets SlotId field to given value.

### HasSlotId

`func (o *FabricPortIdentifierAllOf) HasSlotId() bool`

HasSlotId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



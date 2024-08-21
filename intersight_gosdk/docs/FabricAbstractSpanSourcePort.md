# FabricAbstractSpanSourcePort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "fabric.SpanSourceEthPort"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "fabric.SpanSourceEthPort"]
**AggregatePortId** | Pointer to **int64** | Breakout port Identifier of the Switch Interface. When a port is not configured as a breakout port, the aggregatePortId is set to 0, and unused. When a port is configured as a breakout port, the &#39;aggregatePortId&#39; port number as labeled on the equipment, e.g. the id of the port on the switch. | [optional] 
**PortId** | Pointer to **int64** | Port Identifier of the Switch Interface. When a port is not configured as a breakout port, the portId is the port number as labeled on the equipment, e.g. the id of the port on the switch, FEX or chassis. When a port is configured as a breakout port, the &#39;portId&#39; represents the port id on the fanout side of the breakout cable. | [optional] 
**SlotId** | Pointer to **int64** | Slot Identifier of the Switch Interface. | [optional] 
**SourceRole** | Pointer to **string** | Role of the port configured as SPAN source. * &#x60;Uplink&#x60; - Uplink Role corresponding to PortRole in PortPolicy. * &#x60;FcoeUplink&#x60; - FcoeUplink Role corresponding to PortRole in PortPolicy. * &#x60;FcUplink&#x60; - FcoeUplink Role corresponding to PortRole in PortPolicy. * &#x60;Appliance&#x60; - FcoeUplink Role corresponding to PortRole in PortPolicy. | [optional] [readonly] [default to "Uplink"]
**PhysicalPort** | Pointer to [**NullableInventoryInterfaceRelationship**](InventoryInterfaceRelationship.md) |  | [optional] 
**SpanSession** | Pointer to [**NullableFabricSpanSessionRelationship**](FabricSpanSessionRelationship.md) |  | [optional] 

## Methods

### NewFabricAbstractSpanSourcePort

`func NewFabricAbstractSpanSourcePort(classId string, objectType string, ) *FabricAbstractSpanSourcePort`

NewFabricAbstractSpanSourcePort instantiates a new FabricAbstractSpanSourcePort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricAbstractSpanSourcePortWithDefaults

`func NewFabricAbstractSpanSourcePortWithDefaults() *FabricAbstractSpanSourcePort`

NewFabricAbstractSpanSourcePortWithDefaults instantiates a new FabricAbstractSpanSourcePort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FabricAbstractSpanSourcePort) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FabricAbstractSpanSourcePort) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FabricAbstractSpanSourcePort) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FabricAbstractSpanSourcePort) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FabricAbstractSpanSourcePort) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FabricAbstractSpanSourcePort) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAggregatePortId

`func (o *FabricAbstractSpanSourcePort) GetAggregatePortId() int64`

GetAggregatePortId returns the AggregatePortId field if non-nil, zero value otherwise.

### GetAggregatePortIdOk

`func (o *FabricAbstractSpanSourcePort) GetAggregatePortIdOk() (*int64, bool)`

GetAggregatePortIdOk returns a tuple with the AggregatePortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregatePortId

`func (o *FabricAbstractSpanSourcePort) SetAggregatePortId(v int64)`

SetAggregatePortId sets AggregatePortId field to given value.

### HasAggregatePortId

`func (o *FabricAbstractSpanSourcePort) HasAggregatePortId() bool`

HasAggregatePortId returns a boolean if a field has been set.

### GetPortId

`func (o *FabricAbstractSpanSourcePort) GetPortId() int64`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *FabricAbstractSpanSourcePort) GetPortIdOk() (*int64, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *FabricAbstractSpanSourcePort) SetPortId(v int64)`

SetPortId sets PortId field to given value.

### HasPortId

`func (o *FabricAbstractSpanSourcePort) HasPortId() bool`

HasPortId returns a boolean if a field has been set.

### GetSlotId

`func (o *FabricAbstractSpanSourcePort) GetSlotId() int64`

GetSlotId returns the SlotId field if non-nil, zero value otherwise.

### GetSlotIdOk

`func (o *FabricAbstractSpanSourcePort) GetSlotIdOk() (*int64, bool)`

GetSlotIdOk returns a tuple with the SlotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotId

`func (o *FabricAbstractSpanSourcePort) SetSlotId(v int64)`

SetSlotId sets SlotId field to given value.

### HasSlotId

`func (o *FabricAbstractSpanSourcePort) HasSlotId() bool`

HasSlotId returns a boolean if a field has been set.

### GetSourceRole

`func (o *FabricAbstractSpanSourcePort) GetSourceRole() string`

GetSourceRole returns the SourceRole field if non-nil, zero value otherwise.

### GetSourceRoleOk

`func (o *FabricAbstractSpanSourcePort) GetSourceRoleOk() (*string, bool)`

GetSourceRoleOk returns a tuple with the SourceRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceRole

`func (o *FabricAbstractSpanSourcePort) SetSourceRole(v string)`

SetSourceRole sets SourceRole field to given value.

### HasSourceRole

`func (o *FabricAbstractSpanSourcePort) HasSourceRole() bool`

HasSourceRole returns a boolean if a field has been set.

### GetPhysicalPort

`func (o *FabricAbstractSpanSourcePort) GetPhysicalPort() InventoryInterfaceRelationship`

GetPhysicalPort returns the PhysicalPort field if non-nil, zero value otherwise.

### GetPhysicalPortOk

`func (o *FabricAbstractSpanSourcePort) GetPhysicalPortOk() (*InventoryInterfaceRelationship, bool)`

GetPhysicalPortOk returns a tuple with the PhysicalPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalPort

`func (o *FabricAbstractSpanSourcePort) SetPhysicalPort(v InventoryInterfaceRelationship)`

SetPhysicalPort sets PhysicalPort field to given value.

### HasPhysicalPort

`func (o *FabricAbstractSpanSourcePort) HasPhysicalPort() bool`

HasPhysicalPort returns a boolean if a field has been set.

### SetPhysicalPortNil

`func (o *FabricAbstractSpanSourcePort) SetPhysicalPortNil(b bool)`

 SetPhysicalPortNil sets the value for PhysicalPort to be an explicit nil

### UnsetPhysicalPort
`func (o *FabricAbstractSpanSourcePort) UnsetPhysicalPort()`

UnsetPhysicalPort ensures that no value is present for PhysicalPort, not even an explicit nil
### GetSpanSession

`func (o *FabricAbstractSpanSourcePort) GetSpanSession() FabricSpanSessionRelationship`

GetSpanSession returns the SpanSession field if non-nil, zero value otherwise.

### GetSpanSessionOk

`func (o *FabricAbstractSpanSourcePort) GetSpanSessionOk() (*FabricSpanSessionRelationship, bool)`

GetSpanSessionOk returns a tuple with the SpanSession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanSession

`func (o *FabricAbstractSpanSourcePort) SetSpanSession(v FabricSpanSessionRelationship)`

SetSpanSession sets SpanSession field to given value.

### HasSpanSession

`func (o *FabricAbstractSpanSourcePort) HasSpanSession() bool`

HasSpanSession returns a boolean if a field has been set.

### SetSpanSessionNil

`func (o *FabricAbstractSpanSourcePort) SetSpanSessionNil(b bool)`

 SetSpanSessionNil sets the value for SpanSession to be an explicit nil

### UnsetSpanSession
`func (o *FabricAbstractSpanSourcePort) UnsetSpanSession()`

UnsetSpanSession ensures that no value is present for SpanSession, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



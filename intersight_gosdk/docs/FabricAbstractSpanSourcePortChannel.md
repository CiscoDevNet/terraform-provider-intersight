# FabricAbstractSpanSourcePortChannel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "fabric.SpanSourceEthPortChannel"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "fabric.SpanSourceEthPortChannel"]
**PcId** | Pointer to **int64** | Port-channel ID of SPAN source. | [optional] 
**SourceRole** | Pointer to **string** | Role of the port-channel configured as SPAN source. * &#x60;Uplink&#x60; - Uplink Role corresponding to PortRole in PortPolicy. * &#x60;FcoeUplink&#x60; - FcoeUplink Role corresponding to PortRole in PortPolicy. * &#x60;FcUplink&#x60; - FcoeUplink Role corresponding to PortRole in PortPolicy. * &#x60;Appliance&#x60; - FcoeUplink Role corresponding to PortRole in PortPolicy. | [optional] [readonly] [default to "Uplink"]
**PhysicalPortChannel** | Pointer to [**NullableInventoryInterfaceRelationship**](InventoryInterfaceRelationship.md) |  | [optional] 
**SpanSession** | Pointer to [**NullableFabricSpanSessionRelationship**](FabricSpanSessionRelationship.md) |  | [optional] 

## Methods

### NewFabricAbstractSpanSourcePortChannel

`func NewFabricAbstractSpanSourcePortChannel(classId string, objectType string, ) *FabricAbstractSpanSourcePortChannel`

NewFabricAbstractSpanSourcePortChannel instantiates a new FabricAbstractSpanSourcePortChannel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricAbstractSpanSourcePortChannelWithDefaults

`func NewFabricAbstractSpanSourcePortChannelWithDefaults() *FabricAbstractSpanSourcePortChannel`

NewFabricAbstractSpanSourcePortChannelWithDefaults instantiates a new FabricAbstractSpanSourcePortChannel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FabricAbstractSpanSourcePortChannel) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FabricAbstractSpanSourcePortChannel) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FabricAbstractSpanSourcePortChannel) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FabricAbstractSpanSourcePortChannel) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FabricAbstractSpanSourcePortChannel) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FabricAbstractSpanSourcePortChannel) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetPcId

`func (o *FabricAbstractSpanSourcePortChannel) GetPcId() int64`

GetPcId returns the PcId field if non-nil, zero value otherwise.

### GetPcIdOk

`func (o *FabricAbstractSpanSourcePortChannel) GetPcIdOk() (*int64, bool)`

GetPcIdOk returns a tuple with the PcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcId

`func (o *FabricAbstractSpanSourcePortChannel) SetPcId(v int64)`

SetPcId sets PcId field to given value.

### HasPcId

`func (o *FabricAbstractSpanSourcePortChannel) HasPcId() bool`

HasPcId returns a boolean if a field has been set.

### GetSourceRole

`func (o *FabricAbstractSpanSourcePortChannel) GetSourceRole() string`

GetSourceRole returns the SourceRole field if non-nil, zero value otherwise.

### GetSourceRoleOk

`func (o *FabricAbstractSpanSourcePortChannel) GetSourceRoleOk() (*string, bool)`

GetSourceRoleOk returns a tuple with the SourceRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceRole

`func (o *FabricAbstractSpanSourcePortChannel) SetSourceRole(v string)`

SetSourceRole sets SourceRole field to given value.

### HasSourceRole

`func (o *FabricAbstractSpanSourcePortChannel) HasSourceRole() bool`

HasSourceRole returns a boolean if a field has been set.

### GetPhysicalPortChannel

`func (o *FabricAbstractSpanSourcePortChannel) GetPhysicalPortChannel() InventoryInterfaceRelationship`

GetPhysicalPortChannel returns the PhysicalPortChannel field if non-nil, zero value otherwise.

### GetPhysicalPortChannelOk

`func (o *FabricAbstractSpanSourcePortChannel) GetPhysicalPortChannelOk() (*InventoryInterfaceRelationship, bool)`

GetPhysicalPortChannelOk returns a tuple with the PhysicalPortChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalPortChannel

`func (o *FabricAbstractSpanSourcePortChannel) SetPhysicalPortChannel(v InventoryInterfaceRelationship)`

SetPhysicalPortChannel sets PhysicalPortChannel field to given value.

### HasPhysicalPortChannel

`func (o *FabricAbstractSpanSourcePortChannel) HasPhysicalPortChannel() bool`

HasPhysicalPortChannel returns a boolean if a field has been set.

### SetPhysicalPortChannelNil

`func (o *FabricAbstractSpanSourcePortChannel) SetPhysicalPortChannelNil(b bool)`

 SetPhysicalPortChannelNil sets the value for PhysicalPortChannel to be an explicit nil

### UnsetPhysicalPortChannel
`func (o *FabricAbstractSpanSourcePortChannel) UnsetPhysicalPortChannel()`

UnsetPhysicalPortChannel ensures that no value is present for PhysicalPortChannel, not even an explicit nil
### GetSpanSession

`func (o *FabricAbstractSpanSourcePortChannel) GetSpanSession() FabricSpanSessionRelationship`

GetSpanSession returns the SpanSession field if non-nil, zero value otherwise.

### GetSpanSessionOk

`func (o *FabricAbstractSpanSourcePortChannel) GetSpanSessionOk() (*FabricSpanSessionRelationship, bool)`

GetSpanSessionOk returns a tuple with the SpanSession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanSession

`func (o *FabricAbstractSpanSourcePortChannel) SetSpanSession(v FabricSpanSessionRelationship)`

SetSpanSession sets SpanSession field to given value.

### HasSpanSession

`func (o *FabricAbstractSpanSourcePortChannel) HasSpanSession() bool`

HasSpanSession returns a boolean if a field has been set.

### SetSpanSessionNil

`func (o *FabricAbstractSpanSourcePortChannel) SetSpanSessionNil(b bool)`

 SetSpanSessionNil sets the value for SpanSession to be an explicit nil

### UnsetSpanSession
`func (o *FabricAbstractSpanSourcePortChannel) UnsetSpanSession()`

UnsetSpanSession ensures that no value is present for SpanSession, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# FabricPortChannelRoleAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PcId** | Pointer to **int64** | Unique Identifier of the port-channel, local to this switch. | [optional] 
**Ports** | Pointer to [**[]FabricPortIdentifier**](fabric.PortIdentifier.md) |  | [optional] 
**PortPolicy** | Pointer to [**FabricPortPolicyRelationship**](fabric.PortPolicy.Relationship.md) |  | [optional] 

## Methods

### NewFabricPortChannelRoleAllOf

`func NewFabricPortChannelRoleAllOf() *FabricPortChannelRoleAllOf`

NewFabricPortChannelRoleAllOf instantiates a new FabricPortChannelRoleAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricPortChannelRoleAllOfWithDefaults

`func NewFabricPortChannelRoleAllOfWithDefaults() *FabricPortChannelRoleAllOf`

NewFabricPortChannelRoleAllOfWithDefaults instantiates a new FabricPortChannelRoleAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPcId

`func (o *FabricPortChannelRoleAllOf) GetPcId() int64`

GetPcId returns the PcId field if non-nil, zero value otherwise.

### GetPcIdOk

`func (o *FabricPortChannelRoleAllOf) GetPcIdOk() (*int64, bool)`

GetPcIdOk returns a tuple with the PcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcId

`func (o *FabricPortChannelRoleAllOf) SetPcId(v int64)`

SetPcId sets PcId field to given value.

### HasPcId

`func (o *FabricPortChannelRoleAllOf) HasPcId() bool`

HasPcId returns a boolean if a field has been set.

### GetPorts

`func (o *FabricPortChannelRoleAllOf) GetPorts() []FabricPortIdentifier`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *FabricPortChannelRoleAllOf) GetPortsOk() (*[]FabricPortIdentifier, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *FabricPortChannelRoleAllOf) SetPorts(v []FabricPortIdentifier)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *FabricPortChannelRoleAllOf) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetPortPolicy

`func (o *FabricPortChannelRoleAllOf) GetPortPolicy() FabricPortPolicyRelationship`

GetPortPolicy returns the PortPolicy field if non-nil, zero value otherwise.

### GetPortPolicyOk

`func (o *FabricPortChannelRoleAllOf) GetPortPolicyOk() (*FabricPortPolicyRelationship, bool)`

GetPortPolicyOk returns a tuple with the PortPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortPolicy

`func (o *FabricPortChannelRoleAllOf) SetPortPolicy(v FabricPortPolicyRelationship)`

SetPortPolicy sets PortPolicy field to given value.

### HasPortPolicy

`func (o *FabricPortChannelRoleAllOf) HasPortPolicy() bool`

HasPortPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



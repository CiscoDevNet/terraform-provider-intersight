# FabricPortChannelRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**PcId** | Pointer to **int64** | Unique Identifier of the port-channel, local to this switch. | [optional] 
**Ports** | Pointer to [**[]FabricPortIdentifier**](FabricPortIdentifier.md) |  | [optional] 
**PortPolicy** | Pointer to [**FabricPortPolicyRelationship**](FabricPortPolicyRelationship.md) |  | [optional] 

## Methods

### NewFabricPortChannelRole

`func NewFabricPortChannelRole(classId string, objectType string, ) *FabricPortChannelRole`

NewFabricPortChannelRole instantiates a new FabricPortChannelRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricPortChannelRoleWithDefaults

`func NewFabricPortChannelRoleWithDefaults() *FabricPortChannelRole`

NewFabricPortChannelRoleWithDefaults instantiates a new FabricPortChannelRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FabricPortChannelRole) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FabricPortChannelRole) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FabricPortChannelRole) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FabricPortChannelRole) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FabricPortChannelRole) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FabricPortChannelRole) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetPcId

`func (o *FabricPortChannelRole) GetPcId() int64`

GetPcId returns the PcId field if non-nil, zero value otherwise.

### GetPcIdOk

`func (o *FabricPortChannelRole) GetPcIdOk() (*int64, bool)`

GetPcIdOk returns a tuple with the PcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcId

`func (o *FabricPortChannelRole) SetPcId(v int64)`

SetPcId sets PcId field to given value.

### HasPcId

`func (o *FabricPortChannelRole) HasPcId() bool`

HasPcId returns a boolean if a field has been set.

### GetPorts

`func (o *FabricPortChannelRole) GetPorts() []FabricPortIdentifier`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *FabricPortChannelRole) GetPortsOk() (*[]FabricPortIdentifier, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *FabricPortChannelRole) SetPorts(v []FabricPortIdentifier)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *FabricPortChannelRole) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### SetPortsNil

`func (o *FabricPortChannelRole) SetPortsNil(b bool)`

 SetPortsNil sets the value for Ports to be an explicit nil

### UnsetPorts
`func (o *FabricPortChannelRole) UnsetPorts()`

UnsetPorts ensures that no value is present for Ports, not even an explicit nil
### GetPortPolicy

`func (o *FabricPortChannelRole) GetPortPolicy() FabricPortPolicyRelationship`

GetPortPolicy returns the PortPolicy field if non-nil, zero value otherwise.

### GetPortPolicyOk

`func (o *FabricPortChannelRole) GetPortPolicyOk() (*FabricPortPolicyRelationship, bool)`

GetPortPolicyOk returns a tuple with the PortPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortPolicy

`func (o *FabricPortChannelRole) SetPortPolicy(v FabricPortPolicyRelationship)`

SetPortPolicy sets PortPolicy field to given value.

### HasPortPolicy

`func (o *FabricPortChannelRole) HasPortPolicy() bool`

HasPortPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



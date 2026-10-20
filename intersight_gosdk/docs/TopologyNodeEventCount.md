# TopologyNodeEventCount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "topology.NodeEventCount"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "topology.NodeEventCount"]
**EventCount** | Pointer to [**NullableTopologyEventCount**](TopologyEventCount.md) |  | [optional] 
**NodeId** | Pointer to **string** | Identifier of the topology node these counts belong to. | [optional] [readonly] 

## Methods

### NewTopologyNodeEventCount

`func NewTopologyNodeEventCount(classId string, objectType string, ) *TopologyNodeEventCount`

NewTopologyNodeEventCount instantiates a new TopologyNodeEventCount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopologyNodeEventCountWithDefaults

`func NewTopologyNodeEventCountWithDefaults() *TopologyNodeEventCount`

NewTopologyNodeEventCountWithDefaults instantiates a new TopologyNodeEventCount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *TopologyNodeEventCount) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *TopologyNodeEventCount) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *TopologyNodeEventCount) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *TopologyNodeEventCount) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *TopologyNodeEventCount) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *TopologyNodeEventCount) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEventCount

`func (o *TopologyNodeEventCount) GetEventCount() TopologyEventCount`

GetEventCount returns the EventCount field if non-nil, zero value otherwise.

### GetEventCountOk

`func (o *TopologyNodeEventCount) GetEventCountOk() (*TopologyEventCount, bool)`

GetEventCountOk returns a tuple with the EventCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventCount

`func (o *TopologyNodeEventCount) SetEventCount(v TopologyEventCount)`

SetEventCount sets EventCount field to given value.

### HasEventCount

`func (o *TopologyNodeEventCount) HasEventCount() bool`

HasEventCount returns a boolean if a field has been set.

### SetEventCountNil

`func (o *TopologyNodeEventCount) SetEventCountNil(b bool)`

 SetEventCountNil sets the value for EventCount to be an explicit nil

### UnsetEventCount
`func (o *TopologyNodeEventCount) UnsetEventCount()`

UnsetEventCount ensures that no value is present for EventCount, not even an explicit nil
### GetNodeId

`func (o *TopologyNodeEventCount) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *TopologyNodeEventCount) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *TopologyNodeEventCount) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *TopologyNodeEventCount) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# TopologyView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "topology.View"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "topology.View"]
**Edges** | Pointer to [**[]TopologyEdge**](TopologyEdge.md) |  | [optional] 
**Hierarchy** | Pointer to [**NullableTopologyHierarchy**](TopologyHierarchy.md) |  | [optional] 
**NodeEventCounts** | Pointer to [**[]TopologyNodeEventCount**](TopologyNodeEventCount.md) |  | [optional] 
**Nodes** | Pointer to [**[]TopologyNode**](TopologyNode.md) |  | [optional] 

## Methods

### NewTopologyView

`func NewTopologyView(classId string, objectType string, ) *TopologyView`

NewTopologyView instantiates a new TopologyView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopologyViewWithDefaults

`func NewTopologyViewWithDefaults() *TopologyView`

NewTopologyViewWithDefaults instantiates a new TopologyView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *TopologyView) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *TopologyView) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *TopologyView) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *TopologyView) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *TopologyView) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *TopologyView) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEdges

`func (o *TopologyView) GetEdges() []TopologyEdge`

GetEdges returns the Edges field if non-nil, zero value otherwise.

### GetEdgesOk

`func (o *TopologyView) GetEdgesOk() (*[]TopologyEdge, bool)`

GetEdgesOk returns a tuple with the Edges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdges

`func (o *TopologyView) SetEdges(v []TopologyEdge)`

SetEdges sets Edges field to given value.

### HasEdges

`func (o *TopologyView) HasEdges() bool`

HasEdges returns a boolean if a field has been set.

### SetEdgesNil

`func (o *TopologyView) SetEdgesNil(b bool)`

 SetEdgesNil sets the value for Edges to be an explicit nil

### UnsetEdges
`func (o *TopologyView) UnsetEdges()`

UnsetEdges ensures that no value is present for Edges, not even an explicit nil
### GetHierarchy

`func (o *TopologyView) GetHierarchy() TopologyHierarchy`

GetHierarchy returns the Hierarchy field if non-nil, zero value otherwise.

### GetHierarchyOk

`func (o *TopologyView) GetHierarchyOk() (*TopologyHierarchy, bool)`

GetHierarchyOk returns a tuple with the Hierarchy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHierarchy

`func (o *TopologyView) SetHierarchy(v TopologyHierarchy)`

SetHierarchy sets Hierarchy field to given value.

### HasHierarchy

`func (o *TopologyView) HasHierarchy() bool`

HasHierarchy returns a boolean if a field has been set.

### SetHierarchyNil

`func (o *TopologyView) SetHierarchyNil(b bool)`

 SetHierarchyNil sets the value for Hierarchy to be an explicit nil

### UnsetHierarchy
`func (o *TopologyView) UnsetHierarchy()`

UnsetHierarchy ensures that no value is present for Hierarchy, not even an explicit nil
### GetNodeEventCounts

`func (o *TopologyView) GetNodeEventCounts() []TopologyNodeEventCount`

GetNodeEventCounts returns the NodeEventCounts field if non-nil, zero value otherwise.

### GetNodeEventCountsOk

`func (o *TopologyView) GetNodeEventCountsOk() (*[]TopologyNodeEventCount, bool)`

GetNodeEventCountsOk returns a tuple with the NodeEventCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeEventCounts

`func (o *TopologyView) SetNodeEventCounts(v []TopologyNodeEventCount)`

SetNodeEventCounts sets NodeEventCounts field to given value.

### HasNodeEventCounts

`func (o *TopologyView) HasNodeEventCounts() bool`

HasNodeEventCounts returns a boolean if a field has been set.

### SetNodeEventCountsNil

`func (o *TopologyView) SetNodeEventCountsNil(b bool)`

 SetNodeEventCountsNil sets the value for NodeEventCounts to be an explicit nil

### UnsetNodeEventCounts
`func (o *TopologyView) UnsetNodeEventCounts()`

UnsetNodeEventCounts ensures that no value is present for NodeEventCounts, not even an explicit nil
### GetNodes

`func (o *TopologyView) GetNodes() []TopologyNode`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *TopologyView) GetNodesOk() (*[]TopologyNode, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *TopologyView) SetNodes(v []TopologyNode)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *TopologyView) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### SetNodesNil

`func (o *TopologyView) SetNodesNil(b bool)`

 SetNodesNil sets the value for Nodes to be an explicit nil

### UnsetNodes
`func (o *TopologyView) UnsetNodes()`

UnsetNodes ensures that no value is present for Nodes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



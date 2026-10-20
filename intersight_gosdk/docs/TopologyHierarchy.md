# TopologyHierarchy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "topology.Hierarchy"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "topology.Hierarchy"]
**HierarchyNodes** | Pointer to [**[]TopologyHierarchyNode**](TopologyHierarchyNode.md) |  | [optional] 

## Methods

### NewTopologyHierarchy

`func NewTopologyHierarchy(classId string, objectType string, ) *TopologyHierarchy`

NewTopologyHierarchy instantiates a new TopologyHierarchy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopologyHierarchyWithDefaults

`func NewTopologyHierarchyWithDefaults() *TopologyHierarchy`

NewTopologyHierarchyWithDefaults instantiates a new TopologyHierarchy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *TopologyHierarchy) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *TopologyHierarchy) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *TopologyHierarchy) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *TopologyHierarchy) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *TopologyHierarchy) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *TopologyHierarchy) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetHierarchyNodes

`func (o *TopologyHierarchy) GetHierarchyNodes() []TopologyHierarchyNode`

GetHierarchyNodes returns the HierarchyNodes field if non-nil, zero value otherwise.

### GetHierarchyNodesOk

`func (o *TopologyHierarchy) GetHierarchyNodesOk() (*[]TopologyHierarchyNode, bool)`

GetHierarchyNodesOk returns a tuple with the HierarchyNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHierarchyNodes

`func (o *TopologyHierarchy) SetHierarchyNodes(v []TopologyHierarchyNode)`

SetHierarchyNodes sets HierarchyNodes field to given value.

### HasHierarchyNodes

`func (o *TopologyHierarchy) HasHierarchyNodes() bool`

HasHierarchyNodes returns a boolean if a field has been set.

### SetHierarchyNodesNil

`func (o *TopologyHierarchy) SetHierarchyNodesNil(b bool)`

 SetHierarchyNodesNil sets the value for HierarchyNodes to be an explicit nil

### UnsetHierarchyNodes
`func (o *TopologyHierarchy) UnsetHierarchyNodes()`

UnsetHierarchyNodes ensures that no value is present for HierarchyNodes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



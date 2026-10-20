# TopologyHierarchyNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "topology.HierarchyNode"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "topology.HierarchyNode"]
**Children** | Pointer to [**[]TopologyHierarchyNode**](TopologyHierarchyNode.md) |  | [optional] 
**Parent** | Pointer to **string** | Identifier of the parent node in the containment hierarchy. | [optional] [readonly] 

## Methods

### NewTopologyHierarchyNode

`func NewTopologyHierarchyNode(classId string, objectType string, ) *TopologyHierarchyNode`

NewTopologyHierarchyNode instantiates a new TopologyHierarchyNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopologyHierarchyNodeWithDefaults

`func NewTopologyHierarchyNodeWithDefaults() *TopologyHierarchyNode`

NewTopologyHierarchyNodeWithDefaults instantiates a new TopologyHierarchyNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *TopologyHierarchyNode) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *TopologyHierarchyNode) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *TopologyHierarchyNode) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *TopologyHierarchyNode) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *TopologyHierarchyNode) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *TopologyHierarchyNode) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetChildren

`func (o *TopologyHierarchyNode) GetChildren() []TopologyHierarchyNode`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *TopologyHierarchyNode) GetChildrenOk() (*[]TopologyHierarchyNode, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *TopologyHierarchyNode) SetChildren(v []TopologyHierarchyNode)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *TopologyHierarchyNode) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### SetChildrenNil

`func (o *TopologyHierarchyNode) SetChildrenNil(b bool)`

 SetChildrenNil sets the value for Children to be an explicit nil

### UnsetChildren
`func (o *TopologyHierarchyNode) UnsetChildren()`

UnsetChildren ensures that no value is present for Children, not even an explicit nil
### GetParent

`func (o *TopologyHierarchyNode) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *TopologyHierarchyNode) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *TopologyHierarchyNode) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *TopologyHierarchyNode) HasParent() bool`

HasParent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



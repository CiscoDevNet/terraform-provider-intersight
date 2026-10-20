# TopologyNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "topology.Node"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "topology.Node"]
**Id** | Pointer to **string** | Stable identifier for the topology node. | [optional] [readonly] 
**Meta** | Pointer to **interface{}** | Additional node metadata returned as a JSON document. | [optional] [readonly] 
**Model** | Pointer to **string** | Hardware or software model associated with the topology node. | [optional] [readonly] 
**Moid** | Pointer to **string** | Moid of the backing MO represented by the topology node. | [optional] [readonly] 
**Name** | Pointer to **string** | Display name of the topology node. | [optional] [readonly] 
**Type** | Pointer to **string** | Topology-specific node type used by the UI. | [optional] [readonly] 

## Methods

### NewTopologyNode

`func NewTopologyNode(classId string, objectType string, ) *TopologyNode`

NewTopologyNode instantiates a new TopologyNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopologyNodeWithDefaults

`func NewTopologyNodeWithDefaults() *TopologyNode`

NewTopologyNodeWithDefaults instantiates a new TopologyNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *TopologyNode) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *TopologyNode) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *TopologyNode) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *TopologyNode) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *TopologyNode) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *TopologyNode) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetId

`func (o *TopologyNode) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TopologyNode) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TopologyNode) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TopologyNode) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMeta

`func (o *TopologyNode) GetMeta() interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *TopologyNode) GetMetaOk() (*interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *TopologyNode) SetMeta(v interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *TopologyNode) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### SetMetaNil

`func (o *TopologyNode) SetMetaNil(b bool)`

 SetMetaNil sets the value for Meta to be an explicit nil

### UnsetMeta
`func (o *TopologyNode) UnsetMeta()`

UnsetMeta ensures that no value is present for Meta, not even an explicit nil
### GetModel

`func (o *TopologyNode) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *TopologyNode) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *TopologyNode) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *TopologyNode) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetMoid

`func (o *TopologyNode) GetMoid() string`

GetMoid returns the Moid field if non-nil, zero value otherwise.

### GetMoidOk

`func (o *TopologyNode) GetMoidOk() (*string, bool)`

GetMoidOk returns a tuple with the Moid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoid

`func (o *TopologyNode) SetMoid(v string)`

SetMoid sets Moid field to given value.

### HasMoid

`func (o *TopologyNode) HasMoid() bool`

HasMoid returns a boolean if a field has been set.

### GetName

`func (o *TopologyNode) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TopologyNode) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TopologyNode) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TopologyNode) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *TopologyNode) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TopologyNode) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TopologyNode) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TopologyNode) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



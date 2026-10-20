# TopologyEdge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "topology.Edge"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "topology.Edge"]
**From** | Pointer to **string** | Identifier of the source node or port for the edge. | [optional] [readonly] 
**Id** | Pointer to **string** | Stable identifier for the topology edge. | [optional] [readonly] 
**Speed** | Pointer to **string** | Link speed reported for the edge. | [optional] [readonly] 
**Status** | Pointer to **string** | Current operational status associated with the edge. | [optional] [readonly] 
**To** | Pointer to **string** | Identifier of the destination node or port for the edge. | [optional] [readonly] 

## Methods

### NewTopologyEdge

`func NewTopologyEdge(classId string, objectType string, ) *TopologyEdge`

NewTopologyEdge instantiates a new TopologyEdge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopologyEdgeWithDefaults

`func NewTopologyEdgeWithDefaults() *TopologyEdge`

NewTopologyEdgeWithDefaults instantiates a new TopologyEdge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *TopologyEdge) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *TopologyEdge) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *TopologyEdge) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *TopologyEdge) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *TopologyEdge) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *TopologyEdge) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetFrom

`func (o *TopologyEdge) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *TopologyEdge) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *TopologyEdge) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *TopologyEdge) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetId

`func (o *TopologyEdge) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TopologyEdge) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TopologyEdge) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TopologyEdge) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSpeed

`func (o *TopologyEdge) GetSpeed() string`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *TopologyEdge) GetSpeedOk() (*string, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *TopologyEdge) SetSpeed(v string)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *TopologyEdge) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetStatus

`func (o *TopologyEdge) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TopologyEdge) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TopologyEdge) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TopologyEdge) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTo

`func (o *TopologyEdge) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *TopologyEdge) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *TopologyEdge) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *TopologyEdge) HasTo() bool`

HasTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



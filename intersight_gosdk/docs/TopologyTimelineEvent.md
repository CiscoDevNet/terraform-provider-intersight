# TopologyTimelineEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "topology.TimelineEvent"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "topology.TimelineEvent"]
**EventCounts** | Pointer to [**[]TopologyEventCount**](TopologyEventCount.md) |  | [optional] 

## Methods

### NewTopologyTimelineEvent

`func NewTopologyTimelineEvent(classId string, objectType string, ) *TopologyTimelineEvent`

NewTopologyTimelineEvent instantiates a new TopologyTimelineEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopologyTimelineEventWithDefaults

`func NewTopologyTimelineEventWithDefaults() *TopologyTimelineEvent`

NewTopologyTimelineEventWithDefaults instantiates a new TopologyTimelineEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *TopologyTimelineEvent) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *TopologyTimelineEvent) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *TopologyTimelineEvent) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *TopologyTimelineEvent) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *TopologyTimelineEvent) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *TopologyTimelineEvent) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEventCounts

`func (o *TopologyTimelineEvent) GetEventCounts() []TopologyEventCount`

GetEventCounts returns the EventCounts field if non-nil, zero value otherwise.

### GetEventCountsOk

`func (o *TopologyTimelineEvent) GetEventCountsOk() (*[]TopologyEventCount, bool)`

GetEventCountsOk returns a tuple with the EventCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventCounts

`func (o *TopologyTimelineEvent) SetEventCounts(v []TopologyEventCount)`

SetEventCounts sets EventCounts field to given value.

### HasEventCounts

`func (o *TopologyTimelineEvent) HasEventCounts() bool`

HasEventCounts returns a boolean if a field has been set.

### SetEventCountsNil

`func (o *TopologyTimelineEvent) SetEventCountsNil(b bool)`

 SetEventCountsNil sets the value for EventCounts to be an explicit nil

### UnsetEventCounts
`func (o *TopologyTimelineEvent) UnsetEventCounts()`

UnsetEventCounts ensures that no value is present for EventCounts, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



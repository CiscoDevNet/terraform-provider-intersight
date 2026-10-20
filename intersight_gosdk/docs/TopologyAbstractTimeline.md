# TopologyAbstractTimeline

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**EndTime** | Pointer to **time.Time** | Exclusive historical window end. | [optional] 
**Granularity** | Pointer to **string** | Used only for current topology request overlays. | [optional] 
**StartTime** | Pointer to **time.Time** | Inclusive historical window start. | [optional] 
**TopologyMoid** | Pointer to **string** | Moid of the root topology object the request is scoped to. | [optional] 
**TopologyObjectType** | Pointer to **string** | ObjectType of the root topology object the request is scoped to. | [optional] 
**ViewType** | Pointer to **string** | Root topology view type selected by the UI. | [optional] 

## Methods

### NewTopologyAbstractTimeline

`func NewTopologyAbstractTimeline(classId string, objectType string, ) *TopologyAbstractTimeline`

NewTopologyAbstractTimeline instantiates a new TopologyAbstractTimeline object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopologyAbstractTimelineWithDefaults

`func NewTopologyAbstractTimelineWithDefaults() *TopologyAbstractTimeline`

NewTopologyAbstractTimelineWithDefaults instantiates a new TopologyAbstractTimeline object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *TopologyAbstractTimeline) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *TopologyAbstractTimeline) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *TopologyAbstractTimeline) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *TopologyAbstractTimeline) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *TopologyAbstractTimeline) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *TopologyAbstractTimeline) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEndTime

`func (o *TopologyAbstractTimeline) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *TopologyAbstractTimeline) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *TopologyAbstractTimeline) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *TopologyAbstractTimeline) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetGranularity

`func (o *TopologyAbstractTimeline) GetGranularity() string`

GetGranularity returns the Granularity field if non-nil, zero value otherwise.

### GetGranularityOk

`func (o *TopologyAbstractTimeline) GetGranularityOk() (*string, bool)`

GetGranularityOk returns a tuple with the Granularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGranularity

`func (o *TopologyAbstractTimeline) SetGranularity(v string)`

SetGranularity sets Granularity field to given value.

### HasGranularity

`func (o *TopologyAbstractTimeline) HasGranularity() bool`

HasGranularity returns a boolean if a field has been set.

### GetStartTime

`func (o *TopologyAbstractTimeline) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *TopologyAbstractTimeline) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *TopologyAbstractTimeline) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *TopologyAbstractTimeline) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetTopologyMoid

`func (o *TopologyAbstractTimeline) GetTopologyMoid() string`

GetTopologyMoid returns the TopologyMoid field if non-nil, zero value otherwise.

### GetTopologyMoidOk

`func (o *TopologyAbstractTimeline) GetTopologyMoidOk() (*string, bool)`

GetTopologyMoidOk returns a tuple with the TopologyMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopologyMoid

`func (o *TopologyAbstractTimeline) SetTopologyMoid(v string)`

SetTopologyMoid sets TopologyMoid field to given value.

### HasTopologyMoid

`func (o *TopologyAbstractTimeline) HasTopologyMoid() bool`

HasTopologyMoid returns a boolean if a field has been set.

### GetTopologyObjectType

`func (o *TopologyAbstractTimeline) GetTopologyObjectType() string`

GetTopologyObjectType returns the TopologyObjectType field if non-nil, zero value otherwise.

### GetTopologyObjectTypeOk

`func (o *TopologyAbstractTimeline) GetTopologyObjectTypeOk() (*string, bool)`

GetTopologyObjectTypeOk returns a tuple with the TopologyObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopologyObjectType

`func (o *TopologyAbstractTimeline) SetTopologyObjectType(v string)`

SetTopologyObjectType sets TopologyObjectType field to given value.

### HasTopologyObjectType

`func (o *TopologyAbstractTimeline) HasTopologyObjectType() bool`

HasTopologyObjectType returns a boolean if a field has been set.

### GetViewType

`func (o *TopologyAbstractTimeline) GetViewType() string`

GetViewType returns the ViewType field if non-nil, zero value otherwise.

### GetViewTypeOk

`func (o *TopologyAbstractTimeline) GetViewTypeOk() (*string, bool)`

GetViewTypeOk returns a tuple with the ViewType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewType

`func (o *TopologyAbstractTimeline) SetViewType(v string)`

SetViewType sets ViewType field to given value.

### HasViewType

`func (o *TopologyAbstractTimeline) HasViewType() bool`

HasViewType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



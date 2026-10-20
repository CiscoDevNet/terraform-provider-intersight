# TopologyEventCount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "topology.EventCount"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "topology.EventCount"]
**AlarmCount** | Pointer to [**NullableTopologyAlarmCount**](TopologyAlarmCount.md) |  | [optional] 
**ErrorMetricsCount** | Pointer to **int32** | Number of error metric events associated with the interval. | [optional] [readonly] 
**RequestCount** | Pointer to **int32** | Number of workflow requests associated with the interval. | [optional] [readonly] 
**StartTime** | Pointer to **time.Time** | Inclusive start timestamp of the interval represented by this count. | [optional] [readonly] 

## Methods

### NewTopologyEventCount

`func NewTopologyEventCount(classId string, objectType string, ) *TopologyEventCount`

NewTopologyEventCount instantiates a new TopologyEventCount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopologyEventCountWithDefaults

`func NewTopologyEventCountWithDefaults() *TopologyEventCount`

NewTopologyEventCountWithDefaults instantiates a new TopologyEventCount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *TopologyEventCount) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *TopologyEventCount) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *TopologyEventCount) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *TopologyEventCount) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *TopologyEventCount) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *TopologyEventCount) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAlarmCount

`func (o *TopologyEventCount) GetAlarmCount() TopologyAlarmCount`

GetAlarmCount returns the AlarmCount field if non-nil, zero value otherwise.

### GetAlarmCountOk

`func (o *TopologyEventCount) GetAlarmCountOk() (*TopologyAlarmCount, bool)`

GetAlarmCountOk returns a tuple with the AlarmCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmCount

`func (o *TopologyEventCount) SetAlarmCount(v TopologyAlarmCount)`

SetAlarmCount sets AlarmCount field to given value.

### HasAlarmCount

`func (o *TopologyEventCount) HasAlarmCount() bool`

HasAlarmCount returns a boolean if a field has been set.

### SetAlarmCountNil

`func (o *TopologyEventCount) SetAlarmCountNil(b bool)`

 SetAlarmCountNil sets the value for AlarmCount to be an explicit nil

### UnsetAlarmCount
`func (o *TopologyEventCount) UnsetAlarmCount()`

UnsetAlarmCount ensures that no value is present for AlarmCount, not even an explicit nil
### GetErrorMetricsCount

`func (o *TopologyEventCount) GetErrorMetricsCount() int32`

GetErrorMetricsCount returns the ErrorMetricsCount field if non-nil, zero value otherwise.

### GetErrorMetricsCountOk

`func (o *TopologyEventCount) GetErrorMetricsCountOk() (*int32, bool)`

GetErrorMetricsCountOk returns a tuple with the ErrorMetricsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMetricsCount

`func (o *TopologyEventCount) SetErrorMetricsCount(v int32)`

SetErrorMetricsCount sets ErrorMetricsCount field to given value.

### HasErrorMetricsCount

`func (o *TopologyEventCount) HasErrorMetricsCount() bool`

HasErrorMetricsCount returns a boolean if a field has been set.

### GetRequestCount

`func (o *TopologyEventCount) GetRequestCount() int32`

GetRequestCount returns the RequestCount field if non-nil, zero value otherwise.

### GetRequestCountOk

`func (o *TopologyEventCount) GetRequestCountOk() (*int32, bool)`

GetRequestCountOk returns a tuple with the RequestCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestCount

`func (o *TopologyEventCount) SetRequestCount(v int32)`

SetRequestCount sets RequestCount field to given value.

### HasRequestCount

`func (o *TopologyEventCount) HasRequestCount() bool`

HasRequestCount returns a boolean if a field has been set.

### GetStartTime

`func (o *TopologyEventCount) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *TopologyEventCount) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *TopologyEventCount) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *TopologyEventCount) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# SchedulerWeeklyCadenceParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "scheduler.WeeklyCadenceParams"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "scheduler.WeeklyCadenceParams"]
**DayOfWeek** | Pointer to **[]string** |  | [optional] 
**RunEvery** | Pointer to **int64** | A weekly interval for a task execution. If an interval is not explicitly specified, the task will be executed once every week by default. | [optional] [default to 1]

## Methods

### NewSchedulerWeeklyCadenceParams

`func NewSchedulerWeeklyCadenceParams(classId string, objectType string, ) *SchedulerWeeklyCadenceParams`

NewSchedulerWeeklyCadenceParams instantiates a new SchedulerWeeklyCadenceParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchedulerWeeklyCadenceParamsWithDefaults

`func NewSchedulerWeeklyCadenceParamsWithDefaults() *SchedulerWeeklyCadenceParams`

NewSchedulerWeeklyCadenceParamsWithDefaults instantiates a new SchedulerWeeklyCadenceParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *SchedulerWeeklyCadenceParams) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *SchedulerWeeklyCadenceParams) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *SchedulerWeeklyCadenceParams) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *SchedulerWeeklyCadenceParams) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *SchedulerWeeklyCadenceParams) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *SchedulerWeeklyCadenceParams) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDayOfWeek

`func (o *SchedulerWeeklyCadenceParams) GetDayOfWeek() []string`

GetDayOfWeek returns the DayOfWeek field if non-nil, zero value otherwise.

### GetDayOfWeekOk

`func (o *SchedulerWeeklyCadenceParams) GetDayOfWeekOk() (*[]string, bool)`

GetDayOfWeekOk returns a tuple with the DayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfWeek

`func (o *SchedulerWeeklyCadenceParams) SetDayOfWeek(v []string)`

SetDayOfWeek sets DayOfWeek field to given value.

### HasDayOfWeek

`func (o *SchedulerWeeklyCadenceParams) HasDayOfWeek() bool`

HasDayOfWeek returns a boolean if a field has been set.

### SetDayOfWeekNil

`func (o *SchedulerWeeklyCadenceParams) SetDayOfWeekNil(b bool)`

 SetDayOfWeekNil sets the value for DayOfWeek to be an explicit nil

### UnsetDayOfWeek
`func (o *SchedulerWeeklyCadenceParams) UnsetDayOfWeek()`

UnsetDayOfWeek ensures that no value is present for DayOfWeek, not even an explicit nil
### GetRunEvery

`func (o *SchedulerWeeklyCadenceParams) GetRunEvery() int64`

GetRunEvery returns the RunEvery field if non-nil, zero value otherwise.

### GetRunEveryOk

`func (o *SchedulerWeeklyCadenceParams) GetRunEveryOk() (*int64, bool)`

GetRunEveryOk returns a tuple with the RunEvery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunEvery

`func (o *SchedulerWeeklyCadenceParams) SetRunEvery(v int64)`

SetRunEvery sets RunEvery field to given value.

### HasRunEvery

`func (o *SchedulerWeeklyCadenceParams) HasRunEvery() bool`

HasRunEvery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



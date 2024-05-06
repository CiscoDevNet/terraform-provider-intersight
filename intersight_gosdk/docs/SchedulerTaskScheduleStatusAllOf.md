# SchedulerTaskScheduleStatusAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "scheduler.TaskScheduleStatus"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "scheduler.TaskScheduleStatus"]
**Count** | Pointer to **int64** | The task completion count, which includes both successful executions and any failures. | [optional] [readonly] 
**CurrentStatus** | Pointer to **string** | The status of the current task. * &#x60;None&#x60; - No status is set (default). * &#x60;Scheduled&#x60; - The status is set when a task is scheduled. * &#x60;Running&#x60; - The status is set when a task is running. * &#x60;Completed&#x60; - The status is set when a task is complete. * &#x60;Failed&#x60; - The status is set when a task fails. * &#x60;Suspended&#x60; - The status is set when a task is suspended. * &#x60;Skipped&#x60; - The status is set when a task is skipped because the previous task is still running. | [optional] [readonly] [default to "None"]
**IsSystemSuspended** | Pointer to **bool** | Indicates if this task was suspended by the system. | [optional] [readonly] 
**NextRunStartTime** | Pointer to **time.Time** | The next run time for a recurrently scheduled the task. | [optional] [readonly] 
**PrevRunEndTime** | Pointer to **time.Time** | The time when the last occurrence of scheduled task completed. | [optional] [readonly] 
**PrevRunStartTime** | Pointer to **time.Time** | The previous time the scheduled task was run. | [optional] [readonly] 
**Reason** | Pointer to **string** | The reason why the task failed or suspended. | [optional] [readonly] 

## Methods

### NewSchedulerTaskScheduleStatusAllOf

`func NewSchedulerTaskScheduleStatusAllOf(classId string, objectType string, ) *SchedulerTaskScheduleStatusAllOf`

NewSchedulerTaskScheduleStatusAllOf instantiates a new SchedulerTaskScheduleStatusAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchedulerTaskScheduleStatusAllOfWithDefaults

`func NewSchedulerTaskScheduleStatusAllOfWithDefaults() *SchedulerTaskScheduleStatusAllOf`

NewSchedulerTaskScheduleStatusAllOfWithDefaults instantiates a new SchedulerTaskScheduleStatusAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *SchedulerTaskScheduleStatusAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *SchedulerTaskScheduleStatusAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *SchedulerTaskScheduleStatusAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *SchedulerTaskScheduleStatusAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *SchedulerTaskScheduleStatusAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *SchedulerTaskScheduleStatusAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCount

`func (o *SchedulerTaskScheduleStatusAllOf) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *SchedulerTaskScheduleStatusAllOf) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *SchedulerTaskScheduleStatusAllOf) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *SchedulerTaskScheduleStatusAllOf) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetCurrentStatus

`func (o *SchedulerTaskScheduleStatusAllOf) GetCurrentStatus() string`

GetCurrentStatus returns the CurrentStatus field if non-nil, zero value otherwise.

### GetCurrentStatusOk

`func (o *SchedulerTaskScheduleStatusAllOf) GetCurrentStatusOk() (*string, bool)`

GetCurrentStatusOk returns a tuple with the CurrentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentStatus

`func (o *SchedulerTaskScheduleStatusAllOf) SetCurrentStatus(v string)`

SetCurrentStatus sets CurrentStatus field to given value.

### HasCurrentStatus

`func (o *SchedulerTaskScheduleStatusAllOf) HasCurrentStatus() bool`

HasCurrentStatus returns a boolean if a field has been set.

### GetIsSystemSuspended

`func (o *SchedulerTaskScheduleStatusAllOf) GetIsSystemSuspended() bool`

GetIsSystemSuspended returns the IsSystemSuspended field if non-nil, zero value otherwise.

### GetIsSystemSuspendedOk

`func (o *SchedulerTaskScheduleStatusAllOf) GetIsSystemSuspendedOk() (*bool, bool)`

GetIsSystemSuspendedOk returns a tuple with the IsSystemSuspended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSystemSuspended

`func (o *SchedulerTaskScheduleStatusAllOf) SetIsSystemSuspended(v bool)`

SetIsSystemSuspended sets IsSystemSuspended field to given value.

### HasIsSystemSuspended

`func (o *SchedulerTaskScheduleStatusAllOf) HasIsSystemSuspended() bool`

HasIsSystemSuspended returns a boolean if a field has been set.

### GetNextRunStartTime

`func (o *SchedulerTaskScheduleStatusAllOf) GetNextRunStartTime() time.Time`

GetNextRunStartTime returns the NextRunStartTime field if non-nil, zero value otherwise.

### GetNextRunStartTimeOk

`func (o *SchedulerTaskScheduleStatusAllOf) GetNextRunStartTimeOk() (*time.Time, bool)`

GetNextRunStartTimeOk returns a tuple with the NextRunStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextRunStartTime

`func (o *SchedulerTaskScheduleStatusAllOf) SetNextRunStartTime(v time.Time)`

SetNextRunStartTime sets NextRunStartTime field to given value.

### HasNextRunStartTime

`func (o *SchedulerTaskScheduleStatusAllOf) HasNextRunStartTime() bool`

HasNextRunStartTime returns a boolean if a field has been set.

### GetPrevRunEndTime

`func (o *SchedulerTaskScheduleStatusAllOf) GetPrevRunEndTime() time.Time`

GetPrevRunEndTime returns the PrevRunEndTime field if non-nil, zero value otherwise.

### GetPrevRunEndTimeOk

`func (o *SchedulerTaskScheduleStatusAllOf) GetPrevRunEndTimeOk() (*time.Time, bool)`

GetPrevRunEndTimeOk returns a tuple with the PrevRunEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevRunEndTime

`func (o *SchedulerTaskScheduleStatusAllOf) SetPrevRunEndTime(v time.Time)`

SetPrevRunEndTime sets PrevRunEndTime field to given value.

### HasPrevRunEndTime

`func (o *SchedulerTaskScheduleStatusAllOf) HasPrevRunEndTime() bool`

HasPrevRunEndTime returns a boolean if a field has been set.

### GetPrevRunStartTime

`func (o *SchedulerTaskScheduleStatusAllOf) GetPrevRunStartTime() time.Time`

GetPrevRunStartTime returns the PrevRunStartTime field if non-nil, zero value otherwise.

### GetPrevRunStartTimeOk

`func (o *SchedulerTaskScheduleStatusAllOf) GetPrevRunStartTimeOk() (*time.Time, bool)`

GetPrevRunStartTimeOk returns a tuple with the PrevRunStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevRunStartTime

`func (o *SchedulerTaskScheduleStatusAllOf) SetPrevRunStartTime(v time.Time)`

SetPrevRunStartTime sets PrevRunStartTime field to given value.

### HasPrevRunStartTime

`func (o *SchedulerTaskScheduleStatusAllOf) HasPrevRunStartTime() bool`

HasPrevRunStartTime returns a boolean if a field has been set.

### GetReason

`func (o *SchedulerTaskScheduleStatusAllOf) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *SchedulerTaskScheduleStatusAllOf) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *SchedulerTaskScheduleStatusAllOf) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *SchedulerTaskScheduleStatusAllOf) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# SchedulerTaskScheduleStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "scheduler.TaskScheduleStatus"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "scheduler.TaskScheduleStatus"]
**ConsecutiveFailures** | Pointer to **int64** | The number of consecutive times the task has failed. | [optional] [readonly] 
**Count** | Pointer to **int64** | The task completion count, which includes both successful executions and any failures. | [optional] [readonly] 
**CurrentStatus** | Pointer to **string** | The status of the current task. * &#x60;None&#x60; - No status is set (default). * &#x60;Scheduled&#x60; - The status is set when a task is scheduled. * &#x60;Running&#x60; - The status is set when a task is running. * &#x60;Completed&#x60; - The status is set when a task is complete. * &#x60;Failed&#x60; - The status is set when a task fails. * &#x60;Suspended&#x60; - The status is set when a task is suspended. * &#x60;Skipped&#x60; - The status is set when a task is skipped because the previous task is still running. | [optional] [readonly] [default to "None"]
**IsSystemSuspended** | Pointer to **bool** | Indicates if this task was suspended by the system. | [optional] [readonly] 
**LastRunStatus** | Pointer to **string** | The last task completion status, which includes both successful executions and any failures. * &#x60;None&#x60; - No status is set (default). * &#x60;Scheduled&#x60; - The status is set when a task is scheduled. * &#x60;Running&#x60; - The status is set when a task is running. * &#x60;Completed&#x60; - The status is set when a task is complete. * &#x60;Failed&#x60; - The status is set when a task fails. * &#x60;Suspended&#x60; - The status is set when a task is suspended. * &#x60;Skipped&#x60; - The status is set when a task is skipped because the previous task is still running. | [optional] [readonly] [default to "None"]
**NextRunStartTime** | Pointer to **time.Time** | The next run time for a recurrently scheduled the task. | [optional] [readonly] 
**PrevRunEndTime** | Pointer to **time.Time** | The time when the last occurrence of scheduled task completed. | [optional] [readonly] 
**PrevRunStartTime** | Pointer to **time.Time** | The previous time the scheduled task was run. | [optional] [readonly] 
**Reason** | Pointer to **string** | The reason why the task failed or suspended. | [optional] [readonly] 

## Methods

### NewSchedulerTaskScheduleStatus

`func NewSchedulerTaskScheduleStatus(classId string, objectType string, ) *SchedulerTaskScheduleStatus`

NewSchedulerTaskScheduleStatus instantiates a new SchedulerTaskScheduleStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchedulerTaskScheduleStatusWithDefaults

`func NewSchedulerTaskScheduleStatusWithDefaults() *SchedulerTaskScheduleStatus`

NewSchedulerTaskScheduleStatusWithDefaults instantiates a new SchedulerTaskScheduleStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *SchedulerTaskScheduleStatus) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *SchedulerTaskScheduleStatus) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *SchedulerTaskScheduleStatus) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *SchedulerTaskScheduleStatus) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *SchedulerTaskScheduleStatus) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *SchedulerTaskScheduleStatus) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetConsecutiveFailures

`func (o *SchedulerTaskScheduleStatus) GetConsecutiveFailures() int64`

GetConsecutiveFailures returns the ConsecutiveFailures field if non-nil, zero value otherwise.

### GetConsecutiveFailuresOk

`func (o *SchedulerTaskScheduleStatus) GetConsecutiveFailuresOk() (*int64, bool)`

GetConsecutiveFailuresOk returns a tuple with the ConsecutiveFailures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsecutiveFailures

`func (o *SchedulerTaskScheduleStatus) SetConsecutiveFailures(v int64)`

SetConsecutiveFailures sets ConsecutiveFailures field to given value.

### HasConsecutiveFailures

`func (o *SchedulerTaskScheduleStatus) HasConsecutiveFailures() bool`

HasConsecutiveFailures returns a boolean if a field has been set.

### GetCount

`func (o *SchedulerTaskScheduleStatus) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *SchedulerTaskScheduleStatus) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *SchedulerTaskScheduleStatus) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *SchedulerTaskScheduleStatus) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetCurrentStatus

`func (o *SchedulerTaskScheduleStatus) GetCurrentStatus() string`

GetCurrentStatus returns the CurrentStatus field if non-nil, zero value otherwise.

### GetCurrentStatusOk

`func (o *SchedulerTaskScheduleStatus) GetCurrentStatusOk() (*string, bool)`

GetCurrentStatusOk returns a tuple with the CurrentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentStatus

`func (o *SchedulerTaskScheduleStatus) SetCurrentStatus(v string)`

SetCurrentStatus sets CurrentStatus field to given value.

### HasCurrentStatus

`func (o *SchedulerTaskScheduleStatus) HasCurrentStatus() bool`

HasCurrentStatus returns a boolean if a field has been set.

### GetIsSystemSuspended

`func (o *SchedulerTaskScheduleStatus) GetIsSystemSuspended() bool`

GetIsSystemSuspended returns the IsSystemSuspended field if non-nil, zero value otherwise.

### GetIsSystemSuspendedOk

`func (o *SchedulerTaskScheduleStatus) GetIsSystemSuspendedOk() (*bool, bool)`

GetIsSystemSuspendedOk returns a tuple with the IsSystemSuspended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSystemSuspended

`func (o *SchedulerTaskScheduleStatus) SetIsSystemSuspended(v bool)`

SetIsSystemSuspended sets IsSystemSuspended field to given value.

### HasIsSystemSuspended

`func (o *SchedulerTaskScheduleStatus) HasIsSystemSuspended() bool`

HasIsSystemSuspended returns a boolean if a field has been set.

### GetLastRunStatus

`func (o *SchedulerTaskScheduleStatus) GetLastRunStatus() string`

GetLastRunStatus returns the LastRunStatus field if non-nil, zero value otherwise.

### GetLastRunStatusOk

`func (o *SchedulerTaskScheduleStatus) GetLastRunStatusOk() (*string, bool)`

GetLastRunStatusOk returns a tuple with the LastRunStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRunStatus

`func (o *SchedulerTaskScheduleStatus) SetLastRunStatus(v string)`

SetLastRunStatus sets LastRunStatus field to given value.

### HasLastRunStatus

`func (o *SchedulerTaskScheduleStatus) HasLastRunStatus() bool`

HasLastRunStatus returns a boolean if a field has been set.

### GetNextRunStartTime

`func (o *SchedulerTaskScheduleStatus) GetNextRunStartTime() time.Time`

GetNextRunStartTime returns the NextRunStartTime field if non-nil, zero value otherwise.

### GetNextRunStartTimeOk

`func (o *SchedulerTaskScheduleStatus) GetNextRunStartTimeOk() (*time.Time, bool)`

GetNextRunStartTimeOk returns a tuple with the NextRunStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextRunStartTime

`func (o *SchedulerTaskScheduleStatus) SetNextRunStartTime(v time.Time)`

SetNextRunStartTime sets NextRunStartTime field to given value.

### HasNextRunStartTime

`func (o *SchedulerTaskScheduleStatus) HasNextRunStartTime() bool`

HasNextRunStartTime returns a boolean if a field has been set.

### GetPrevRunEndTime

`func (o *SchedulerTaskScheduleStatus) GetPrevRunEndTime() time.Time`

GetPrevRunEndTime returns the PrevRunEndTime field if non-nil, zero value otherwise.

### GetPrevRunEndTimeOk

`func (o *SchedulerTaskScheduleStatus) GetPrevRunEndTimeOk() (*time.Time, bool)`

GetPrevRunEndTimeOk returns a tuple with the PrevRunEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevRunEndTime

`func (o *SchedulerTaskScheduleStatus) SetPrevRunEndTime(v time.Time)`

SetPrevRunEndTime sets PrevRunEndTime field to given value.

### HasPrevRunEndTime

`func (o *SchedulerTaskScheduleStatus) HasPrevRunEndTime() bool`

HasPrevRunEndTime returns a boolean if a field has been set.

### GetPrevRunStartTime

`func (o *SchedulerTaskScheduleStatus) GetPrevRunStartTime() time.Time`

GetPrevRunStartTime returns the PrevRunStartTime field if non-nil, zero value otherwise.

### GetPrevRunStartTimeOk

`func (o *SchedulerTaskScheduleStatus) GetPrevRunStartTimeOk() (*time.Time, bool)`

GetPrevRunStartTimeOk returns a tuple with the PrevRunStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevRunStartTime

`func (o *SchedulerTaskScheduleStatus) SetPrevRunStartTime(v time.Time)`

SetPrevRunStartTime sets PrevRunStartTime field to given value.

### HasPrevRunStartTime

`func (o *SchedulerTaskScheduleStatus) HasPrevRunStartTime() bool`

HasPrevRunStartTime returns a boolean if a field has been set.

### GetReason

`func (o *SchedulerTaskScheduleStatus) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *SchedulerTaskScheduleStatus) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *SchedulerTaskScheduleStatus) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *SchedulerTaskScheduleStatus) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



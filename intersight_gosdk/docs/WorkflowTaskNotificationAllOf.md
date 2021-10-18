# WorkflowTaskNotificationAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.TaskNotification"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.TaskNotification"]
**CorrelationId** | Pointer to **string** | The correlation id of the scheduled task. | [optional] 
**EndTime** | Pointer to **string** | The end time of the scheduled task. | [optional] 
**Input** | Pointer to **string** | The input of the scheduled task. | [optional] 
**Output** | Pointer to **string** | The output of the scheduled task. | [optional] 
**ReasonForIncompletion** | Pointer to **string** | The reason for incompletion status of the task. | [optional] 
**ReferenceTaskName** | Pointer to **string** | The task reference name of the scheduled task. | [optional] 
**RetryCount** | Pointer to **float32** | The number of times the task retries on failure. | [optional] 
**ScheduledTime** | Pointer to **string** | The scheduled time of the task. | [optional] 
**StartTime** | Pointer to **string** | The start time of the scheduled task. | [optional] 
**Status** | Pointer to **string** | The status of the scheduled task. | [optional] 
**TaskDefName** | Pointer to **string** | The definition of the task explains about the task. | [optional] 
**TaskDescription** | Pointer to **string** | The description of the task explains about the task. | [optional] 
**TaskId** | Pointer to **string** | Unique id of the scheduled task. | [optional] 
**TaskType** | Pointer to **string** | The type of the scheduled task. | [optional] 
**UpdateTime** | Pointer to **string** | The update time of the scheduled task. | [optional] 
**WorkflowId** | Pointer to **string** | The unique id of the running workflow containing this scheduled task. | [optional] 
**WorkflowTaskType** | Pointer to **string** | The type of the workflow task. | [optional] 
**WorkflowType** | Pointer to **string** | The type of workflow containing this scheduled task. | [optional] 

## Methods

### NewWorkflowTaskNotificationAllOf

`func NewWorkflowTaskNotificationAllOf(classId string, objectType string, ) *WorkflowTaskNotificationAllOf`

NewWorkflowTaskNotificationAllOf instantiates a new WorkflowTaskNotificationAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowTaskNotificationAllOfWithDefaults

`func NewWorkflowTaskNotificationAllOfWithDefaults() *WorkflowTaskNotificationAllOf`

NewWorkflowTaskNotificationAllOfWithDefaults instantiates a new WorkflowTaskNotificationAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowTaskNotificationAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowTaskNotificationAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowTaskNotificationAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowTaskNotificationAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowTaskNotificationAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowTaskNotificationAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCorrelationId

`func (o *WorkflowTaskNotificationAllOf) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *WorkflowTaskNotificationAllOf) GetCorrelationIdOk() (*string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationId

`func (o *WorkflowTaskNotificationAllOf) SetCorrelationId(v string)`

SetCorrelationId sets CorrelationId field to given value.

### HasCorrelationId

`func (o *WorkflowTaskNotificationAllOf) HasCorrelationId() bool`

HasCorrelationId returns a boolean if a field has been set.

### GetEndTime

`func (o *WorkflowTaskNotificationAllOf) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *WorkflowTaskNotificationAllOf) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *WorkflowTaskNotificationAllOf) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *WorkflowTaskNotificationAllOf) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetInput

`func (o *WorkflowTaskNotificationAllOf) GetInput() string`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *WorkflowTaskNotificationAllOf) GetInputOk() (*string, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *WorkflowTaskNotificationAllOf) SetInput(v string)`

SetInput sets Input field to given value.

### HasInput

`func (o *WorkflowTaskNotificationAllOf) HasInput() bool`

HasInput returns a boolean if a field has been set.

### GetOutput

`func (o *WorkflowTaskNotificationAllOf) GetOutput() string`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *WorkflowTaskNotificationAllOf) GetOutputOk() (*string, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *WorkflowTaskNotificationAllOf) SetOutput(v string)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *WorkflowTaskNotificationAllOf) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### GetReasonForIncompletion

`func (o *WorkflowTaskNotificationAllOf) GetReasonForIncompletion() string`

GetReasonForIncompletion returns the ReasonForIncompletion field if non-nil, zero value otherwise.

### GetReasonForIncompletionOk

`func (o *WorkflowTaskNotificationAllOf) GetReasonForIncompletionOk() (*string, bool)`

GetReasonForIncompletionOk returns a tuple with the ReasonForIncompletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonForIncompletion

`func (o *WorkflowTaskNotificationAllOf) SetReasonForIncompletion(v string)`

SetReasonForIncompletion sets ReasonForIncompletion field to given value.

### HasReasonForIncompletion

`func (o *WorkflowTaskNotificationAllOf) HasReasonForIncompletion() bool`

HasReasonForIncompletion returns a boolean if a field has been set.

### GetReferenceTaskName

`func (o *WorkflowTaskNotificationAllOf) GetReferenceTaskName() string`

GetReferenceTaskName returns the ReferenceTaskName field if non-nil, zero value otherwise.

### GetReferenceTaskNameOk

`func (o *WorkflowTaskNotificationAllOf) GetReferenceTaskNameOk() (*string, bool)`

GetReferenceTaskNameOk returns a tuple with the ReferenceTaskName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceTaskName

`func (o *WorkflowTaskNotificationAllOf) SetReferenceTaskName(v string)`

SetReferenceTaskName sets ReferenceTaskName field to given value.

### HasReferenceTaskName

`func (o *WorkflowTaskNotificationAllOf) HasReferenceTaskName() bool`

HasReferenceTaskName returns a boolean if a field has been set.

### GetRetryCount

`func (o *WorkflowTaskNotificationAllOf) GetRetryCount() float32`

GetRetryCount returns the RetryCount field if non-nil, zero value otherwise.

### GetRetryCountOk

`func (o *WorkflowTaskNotificationAllOf) GetRetryCountOk() (*float32, bool)`

GetRetryCountOk returns a tuple with the RetryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryCount

`func (o *WorkflowTaskNotificationAllOf) SetRetryCount(v float32)`

SetRetryCount sets RetryCount field to given value.

### HasRetryCount

`func (o *WorkflowTaskNotificationAllOf) HasRetryCount() bool`

HasRetryCount returns a boolean if a field has been set.

### GetScheduledTime

`func (o *WorkflowTaskNotificationAllOf) GetScheduledTime() string`

GetScheduledTime returns the ScheduledTime field if non-nil, zero value otherwise.

### GetScheduledTimeOk

`func (o *WorkflowTaskNotificationAllOf) GetScheduledTimeOk() (*string, bool)`

GetScheduledTimeOk returns a tuple with the ScheduledTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledTime

`func (o *WorkflowTaskNotificationAllOf) SetScheduledTime(v string)`

SetScheduledTime sets ScheduledTime field to given value.

### HasScheduledTime

`func (o *WorkflowTaskNotificationAllOf) HasScheduledTime() bool`

HasScheduledTime returns a boolean if a field has been set.

### GetStartTime

`func (o *WorkflowTaskNotificationAllOf) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *WorkflowTaskNotificationAllOf) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *WorkflowTaskNotificationAllOf) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *WorkflowTaskNotificationAllOf) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStatus

`func (o *WorkflowTaskNotificationAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WorkflowTaskNotificationAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WorkflowTaskNotificationAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WorkflowTaskNotificationAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTaskDefName

`func (o *WorkflowTaskNotificationAllOf) GetTaskDefName() string`

GetTaskDefName returns the TaskDefName field if non-nil, zero value otherwise.

### GetTaskDefNameOk

`func (o *WorkflowTaskNotificationAllOf) GetTaskDefNameOk() (*string, bool)`

GetTaskDefNameOk returns a tuple with the TaskDefName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskDefName

`func (o *WorkflowTaskNotificationAllOf) SetTaskDefName(v string)`

SetTaskDefName sets TaskDefName field to given value.

### HasTaskDefName

`func (o *WorkflowTaskNotificationAllOf) HasTaskDefName() bool`

HasTaskDefName returns a boolean if a field has been set.

### GetTaskDescription

`func (o *WorkflowTaskNotificationAllOf) GetTaskDescription() string`

GetTaskDescription returns the TaskDescription field if non-nil, zero value otherwise.

### GetTaskDescriptionOk

`func (o *WorkflowTaskNotificationAllOf) GetTaskDescriptionOk() (*string, bool)`

GetTaskDescriptionOk returns a tuple with the TaskDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskDescription

`func (o *WorkflowTaskNotificationAllOf) SetTaskDescription(v string)`

SetTaskDescription sets TaskDescription field to given value.

### HasTaskDescription

`func (o *WorkflowTaskNotificationAllOf) HasTaskDescription() bool`

HasTaskDescription returns a boolean if a field has been set.

### GetTaskId

`func (o *WorkflowTaskNotificationAllOf) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *WorkflowTaskNotificationAllOf) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *WorkflowTaskNotificationAllOf) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.

### HasTaskId

`func (o *WorkflowTaskNotificationAllOf) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.

### GetTaskType

`func (o *WorkflowTaskNotificationAllOf) GetTaskType() string`

GetTaskType returns the TaskType field if non-nil, zero value otherwise.

### GetTaskTypeOk

`func (o *WorkflowTaskNotificationAllOf) GetTaskTypeOk() (*string, bool)`

GetTaskTypeOk returns a tuple with the TaskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskType

`func (o *WorkflowTaskNotificationAllOf) SetTaskType(v string)`

SetTaskType sets TaskType field to given value.

### HasTaskType

`func (o *WorkflowTaskNotificationAllOf) HasTaskType() bool`

HasTaskType returns a boolean if a field has been set.

### GetUpdateTime

`func (o *WorkflowTaskNotificationAllOf) GetUpdateTime() string`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *WorkflowTaskNotificationAllOf) GetUpdateTimeOk() (*string, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *WorkflowTaskNotificationAllOf) SetUpdateTime(v string)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *WorkflowTaskNotificationAllOf) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetWorkflowId

`func (o *WorkflowTaskNotificationAllOf) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *WorkflowTaskNotificationAllOf) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *WorkflowTaskNotificationAllOf) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.

### HasWorkflowId

`func (o *WorkflowTaskNotificationAllOf) HasWorkflowId() bool`

HasWorkflowId returns a boolean if a field has been set.

### GetWorkflowTaskType

`func (o *WorkflowTaskNotificationAllOf) GetWorkflowTaskType() string`

GetWorkflowTaskType returns the WorkflowTaskType field if non-nil, zero value otherwise.

### GetWorkflowTaskTypeOk

`func (o *WorkflowTaskNotificationAllOf) GetWorkflowTaskTypeOk() (*string, bool)`

GetWorkflowTaskTypeOk returns a tuple with the WorkflowTaskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowTaskType

`func (o *WorkflowTaskNotificationAllOf) SetWorkflowTaskType(v string)`

SetWorkflowTaskType sets WorkflowTaskType field to given value.

### HasWorkflowTaskType

`func (o *WorkflowTaskNotificationAllOf) HasWorkflowTaskType() bool`

HasWorkflowTaskType returns a boolean if a field has been set.

### GetWorkflowType

`func (o *WorkflowTaskNotificationAllOf) GetWorkflowType() string`

GetWorkflowType returns the WorkflowType field if non-nil, zero value otherwise.

### GetWorkflowTypeOk

`func (o *WorkflowTaskNotificationAllOf) GetWorkflowTypeOk() (*string, bool)`

GetWorkflowTypeOk returns a tuple with the WorkflowType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowType

`func (o *WorkflowTaskNotificationAllOf) SetWorkflowType(v string)`

SetWorkflowType sets WorkflowType field to given value.

### HasWorkflowType

`func (o *WorkflowTaskNotificationAllOf) HasWorkflowType() bool`

HasWorkflowType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



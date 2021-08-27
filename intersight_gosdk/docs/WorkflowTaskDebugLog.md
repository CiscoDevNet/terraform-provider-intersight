# WorkflowTaskDebugLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.TaskDebugLog"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.TaskDebugLog"]
**RetryCount** | Pointer to **int64** | A counter for number of retries. | [optional] [readonly] 
**TaskDebugLogEntries** | Pointer to **interface{}** | Holds information helpful in isolating task failures. | [optional] [readonly] 
**TaskInstId** | Pointer to **string** | The unique identifier for task instance. | [optional] [readonly] 
**TaskInfo** | Pointer to [**WorkflowTaskInfoRelationship**](WorkflowTaskInfoRelationship.md) |  | [optional] 
**WorkflowInfo** | Pointer to [**WorkflowWorkflowInfoRelationship**](WorkflowWorkflowInfoRelationship.md) |  | [optional] 

## Methods

### NewWorkflowTaskDebugLog

`func NewWorkflowTaskDebugLog(classId string, objectType string, ) *WorkflowTaskDebugLog`

NewWorkflowTaskDebugLog instantiates a new WorkflowTaskDebugLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowTaskDebugLogWithDefaults

`func NewWorkflowTaskDebugLogWithDefaults() *WorkflowTaskDebugLog`

NewWorkflowTaskDebugLogWithDefaults instantiates a new WorkflowTaskDebugLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowTaskDebugLog) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowTaskDebugLog) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowTaskDebugLog) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowTaskDebugLog) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowTaskDebugLog) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowTaskDebugLog) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetRetryCount

`func (o *WorkflowTaskDebugLog) GetRetryCount() int64`

GetRetryCount returns the RetryCount field if non-nil, zero value otherwise.

### GetRetryCountOk

`func (o *WorkflowTaskDebugLog) GetRetryCountOk() (*int64, bool)`

GetRetryCountOk returns a tuple with the RetryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryCount

`func (o *WorkflowTaskDebugLog) SetRetryCount(v int64)`

SetRetryCount sets RetryCount field to given value.

### HasRetryCount

`func (o *WorkflowTaskDebugLog) HasRetryCount() bool`

HasRetryCount returns a boolean if a field has been set.

### GetTaskDebugLogEntries

`func (o *WorkflowTaskDebugLog) GetTaskDebugLogEntries() interface{}`

GetTaskDebugLogEntries returns the TaskDebugLogEntries field if non-nil, zero value otherwise.

### GetTaskDebugLogEntriesOk

`func (o *WorkflowTaskDebugLog) GetTaskDebugLogEntriesOk() (*interface{}, bool)`

GetTaskDebugLogEntriesOk returns a tuple with the TaskDebugLogEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskDebugLogEntries

`func (o *WorkflowTaskDebugLog) SetTaskDebugLogEntries(v interface{})`

SetTaskDebugLogEntries sets TaskDebugLogEntries field to given value.

### HasTaskDebugLogEntries

`func (o *WorkflowTaskDebugLog) HasTaskDebugLogEntries() bool`

HasTaskDebugLogEntries returns a boolean if a field has been set.

### SetTaskDebugLogEntriesNil

`func (o *WorkflowTaskDebugLog) SetTaskDebugLogEntriesNil(b bool)`

 SetTaskDebugLogEntriesNil sets the value for TaskDebugLogEntries to be an explicit nil

### UnsetTaskDebugLogEntries
`func (o *WorkflowTaskDebugLog) UnsetTaskDebugLogEntries()`

UnsetTaskDebugLogEntries ensures that no value is present for TaskDebugLogEntries, not even an explicit nil
### GetTaskInstId

`func (o *WorkflowTaskDebugLog) GetTaskInstId() string`

GetTaskInstId returns the TaskInstId field if non-nil, zero value otherwise.

### GetTaskInstIdOk

`func (o *WorkflowTaskDebugLog) GetTaskInstIdOk() (*string, bool)`

GetTaskInstIdOk returns a tuple with the TaskInstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskInstId

`func (o *WorkflowTaskDebugLog) SetTaskInstId(v string)`

SetTaskInstId sets TaskInstId field to given value.

### HasTaskInstId

`func (o *WorkflowTaskDebugLog) HasTaskInstId() bool`

HasTaskInstId returns a boolean if a field has been set.

### GetTaskInfo

`func (o *WorkflowTaskDebugLog) GetTaskInfo() WorkflowTaskInfoRelationship`

GetTaskInfo returns the TaskInfo field if non-nil, zero value otherwise.

### GetTaskInfoOk

`func (o *WorkflowTaskDebugLog) GetTaskInfoOk() (*WorkflowTaskInfoRelationship, bool)`

GetTaskInfoOk returns a tuple with the TaskInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskInfo

`func (o *WorkflowTaskDebugLog) SetTaskInfo(v WorkflowTaskInfoRelationship)`

SetTaskInfo sets TaskInfo field to given value.

### HasTaskInfo

`func (o *WorkflowTaskDebugLog) HasTaskInfo() bool`

HasTaskInfo returns a boolean if a field has been set.

### GetWorkflowInfo

`func (o *WorkflowTaskDebugLog) GetWorkflowInfo() WorkflowWorkflowInfoRelationship`

GetWorkflowInfo returns the WorkflowInfo field if non-nil, zero value otherwise.

### GetWorkflowInfoOk

`func (o *WorkflowTaskDebugLog) GetWorkflowInfoOk() (*WorkflowWorkflowInfoRelationship, bool)`

GetWorkflowInfoOk returns a tuple with the WorkflowInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowInfo

`func (o *WorkflowTaskDebugLog) SetWorkflowInfo(v WorkflowWorkflowInfoRelationship)`

SetWorkflowInfo sets WorkflowInfo field to given value.

### HasWorkflowInfo

`func (o *WorkflowTaskDebugLog) HasWorkflowInfo() bool`

HasWorkflowInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# WorkflowForkTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ForkedTasks** | Pointer to **[]string** |  | [optional] 
**JoinTask** | Pointer to **string** | Task name for the join control task that must follow a fork control task. | [optional] 

## Methods

### NewWorkflowForkTask

`func NewWorkflowForkTask() *WorkflowForkTask`

NewWorkflowForkTask instantiates a new WorkflowForkTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowForkTaskWithDefaults

`func NewWorkflowForkTaskWithDefaults() *WorkflowForkTask`

NewWorkflowForkTaskWithDefaults instantiates a new WorkflowForkTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForkedTasks

`func (o *WorkflowForkTask) GetForkedTasks() []string`

GetForkedTasks returns the ForkedTasks field if non-nil, zero value otherwise.

### GetForkedTasksOk

`func (o *WorkflowForkTask) GetForkedTasksOk() (*[]string, bool)`

GetForkedTasksOk returns a tuple with the ForkedTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForkedTasks

`func (o *WorkflowForkTask) SetForkedTasks(v []string)`

SetForkedTasks sets ForkedTasks field to given value.

### HasForkedTasks

`func (o *WorkflowForkTask) HasForkedTasks() bool`

HasForkedTasks returns a boolean if a field has been set.

### GetJoinTask

`func (o *WorkflowForkTask) GetJoinTask() string`

GetJoinTask returns the JoinTask field if non-nil, zero value otherwise.

### GetJoinTaskOk

`func (o *WorkflowForkTask) GetJoinTaskOk() (*string, bool)`

GetJoinTaskOk returns a tuple with the JoinTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinTask

`func (o *WorkflowForkTask) SetJoinTask(v string)`

SetJoinTask sets JoinTask field to given value.

### HasJoinTask

`func (o *WorkflowForkTask) HasJoinTask() bool`

HasJoinTask returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



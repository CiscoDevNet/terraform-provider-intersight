# WorkflowDynamicWorkflowActionTaskList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | The action of the Dynamic Workflow. | [optional] 
**Tasks** | Pointer to **interface{}** | The task list that has precedence which dictates how the workflow should be constructed. | [optional] 

## Methods

### NewWorkflowDynamicWorkflowActionTaskList

`func NewWorkflowDynamicWorkflowActionTaskList() *WorkflowDynamicWorkflowActionTaskList`

NewWorkflowDynamicWorkflowActionTaskList instantiates a new WorkflowDynamicWorkflowActionTaskList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowDynamicWorkflowActionTaskListWithDefaults

`func NewWorkflowDynamicWorkflowActionTaskListWithDefaults() *WorkflowDynamicWorkflowActionTaskList`

NewWorkflowDynamicWorkflowActionTaskListWithDefaults instantiates a new WorkflowDynamicWorkflowActionTaskList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *WorkflowDynamicWorkflowActionTaskList) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *WorkflowDynamicWorkflowActionTaskList) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *WorkflowDynamicWorkflowActionTaskList) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *WorkflowDynamicWorkflowActionTaskList) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetTasks

`func (o *WorkflowDynamicWorkflowActionTaskList) GetTasks() interface{}`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *WorkflowDynamicWorkflowActionTaskList) GetTasksOk() (*interface{}, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *WorkflowDynamicWorkflowActionTaskList) SetTasks(v interface{})`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *WorkflowDynamicWorkflowActionTaskList) HasTasks() bool`

HasTasks returns a boolean if a field has been set.

### SetTasksNil

`func (o *WorkflowDynamicWorkflowActionTaskList) SetTasksNil(b bool)`

 SetTasksNil sets the value for Tasks to be an explicit nil

### UnsetTasks
`func (o *WorkflowDynamicWorkflowActionTaskList) UnsetTasks()`

UnsetTasks ensures that no value is present for Tasks, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



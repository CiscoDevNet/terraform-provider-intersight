# WorkflowJoinTaskAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JoinOnTasks** | Pointer to **[]string** |  | [optional] 
**OnSuccess** | Pointer to **string** | Name of the next task to run if all fork path specified in the JoinOnTasks list succeeds which is the unique name given to the task instance within the workflow. In a graph model, denotes an edge to another Task Node. | [optional] 

## Methods

### NewWorkflowJoinTaskAllOf

`func NewWorkflowJoinTaskAllOf() *WorkflowJoinTaskAllOf`

NewWorkflowJoinTaskAllOf instantiates a new WorkflowJoinTaskAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowJoinTaskAllOfWithDefaults

`func NewWorkflowJoinTaskAllOfWithDefaults() *WorkflowJoinTaskAllOf`

NewWorkflowJoinTaskAllOfWithDefaults instantiates a new WorkflowJoinTaskAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJoinOnTasks

`func (o *WorkflowJoinTaskAllOf) GetJoinOnTasks() []string`

GetJoinOnTasks returns the JoinOnTasks field if non-nil, zero value otherwise.

### GetJoinOnTasksOk

`func (o *WorkflowJoinTaskAllOf) GetJoinOnTasksOk() (*[]string, bool)`

GetJoinOnTasksOk returns a tuple with the JoinOnTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinOnTasks

`func (o *WorkflowJoinTaskAllOf) SetJoinOnTasks(v []string)`

SetJoinOnTasks sets JoinOnTasks field to given value.

### HasJoinOnTasks

`func (o *WorkflowJoinTaskAllOf) HasJoinOnTasks() bool`

HasJoinOnTasks returns a boolean if a field has been set.

### GetOnSuccess

`func (o *WorkflowJoinTaskAllOf) GetOnSuccess() string`

GetOnSuccess returns the OnSuccess field if non-nil, zero value otherwise.

### GetOnSuccessOk

`func (o *WorkflowJoinTaskAllOf) GetOnSuccessOk() (*string, bool)`

GetOnSuccessOk returns a tuple with the OnSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnSuccess

`func (o *WorkflowJoinTaskAllOf) SetOnSuccess(v string)`

SetOnSuccess sets OnSuccess field to given value.

### HasOnSuccess

`func (o *WorkflowJoinTaskAllOf) HasOnSuccess() bool`

HasOnSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# WorkflowForkTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.ForkTask"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.ForkTask"]
**ForkedTasks** | Pointer to **[]string** |  | [optional] 
**JoinTask** | Pointer to **string** | Task name for the join control task that must follow a fork control task. | [optional] 

## Methods

### NewWorkflowForkTask

`func NewWorkflowForkTask(classId string, objectType string, ) *WorkflowForkTask`

NewWorkflowForkTask instantiates a new WorkflowForkTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowForkTaskWithDefaults

`func NewWorkflowForkTaskWithDefaults() *WorkflowForkTask`

NewWorkflowForkTaskWithDefaults instantiates a new WorkflowForkTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowForkTask) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowForkTask) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowForkTask) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowForkTask) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowForkTask) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowForkTask) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


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

### SetForkedTasksNil

`func (o *WorkflowForkTask) SetForkedTasksNil(b bool)`

 SetForkedTasksNil sets the value for ForkedTasks to be an explicit nil

### UnsetForkedTasks
`func (o *WorkflowForkTask) UnsetForkedTasks()`

UnsetForkedTasks ensures that no value is present for ForkedTasks, not even an explicit nil
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



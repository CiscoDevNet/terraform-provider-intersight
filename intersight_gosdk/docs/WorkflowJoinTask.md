# WorkflowJoinTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.JoinTask"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.JoinTask"]
**JoinOnTasks** | Pointer to **[]string** |  | [optional] 
**OnSuccess** | Pointer to **string** | Name of the next task to run if all fork path specified in the JoinOnTasks list succeeds which is the unique name given to the task instance within the workflow. In a graph model, denotes an edge to another Task Node. | [optional] 

## Methods

### NewWorkflowJoinTask

`func NewWorkflowJoinTask(classId string, objectType string, ) *WorkflowJoinTask`

NewWorkflowJoinTask instantiates a new WorkflowJoinTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowJoinTaskWithDefaults

`func NewWorkflowJoinTaskWithDefaults() *WorkflowJoinTask`

NewWorkflowJoinTaskWithDefaults instantiates a new WorkflowJoinTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowJoinTask) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowJoinTask) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowJoinTask) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowJoinTask) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowJoinTask) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowJoinTask) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetJoinOnTasks

`func (o *WorkflowJoinTask) GetJoinOnTasks() []string`

GetJoinOnTasks returns the JoinOnTasks field if non-nil, zero value otherwise.

### GetJoinOnTasksOk

`func (o *WorkflowJoinTask) GetJoinOnTasksOk() (*[]string, bool)`

GetJoinOnTasksOk returns a tuple with the JoinOnTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinOnTasks

`func (o *WorkflowJoinTask) SetJoinOnTasks(v []string)`

SetJoinOnTasks sets JoinOnTasks field to given value.

### HasJoinOnTasks

`func (o *WorkflowJoinTask) HasJoinOnTasks() bool`

HasJoinOnTasks returns a boolean if a field has been set.

### SetJoinOnTasksNil

`func (o *WorkflowJoinTask) SetJoinOnTasksNil(b bool)`

 SetJoinOnTasksNil sets the value for JoinOnTasks to be an explicit nil

### UnsetJoinOnTasks
`func (o *WorkflowJoinTask) UnsetJoinOnTasks()`

UnsetJoinOnTasks ensures that no value is present for JoinOnTasks, not even an explicit nil
### GetOnSuccess

`func (o *WorkflowJoinTask) GetOnSuccess() string`

GetOnSuccess returns the OnSuccess field if non-nil, zero value otherwise.

### GetOnSuccessOk

`func (o *WorkflowJoinTask) GetOnSuccessOk() (*string, bool)`

GetOnSuccessOk returns a tuple with the OnSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnSuccess

`func (o *WorkflowJoinTask) SetOnSuccess(v string)`

SetOnSuccess sets OnSuccess field to given value.

### HasOnSuccess

`func (o *WorkflowJoinTask) HasOnSuccess() bool`

HasOnSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# WorkflowDynamicWorkflowActionTaskList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.DynamicWorkflowActionTaskList"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.DynamicWorkflowActionTaskList"]
**Action** | Pointer to **string** | The action of the Dynamic Workflow. | [optional] 
**Tasks** | Pointer to **interface{}** | The task list that has precedence which dictates how the workflow should be constructed. | [optional] 

## Methods

### NewWorkflowDynamicWorkflowActionTaskList

`func NewWorkflowDynamicWorkflowActionTaskList(classId string, objectType string, ) *WorkflowDynamicWorkflowActionTaskList`

NewWorkflowDynamicWorkflowActionTaskList instantiates a new WorkflowDynamicWorkflowActionTaskList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowDynamicWorkflowActionTaskListWithDefaults

`func NewWorkflowDynamicWorkflowActionTaskListWithDefaults() *WorkflowDynamicWorkflowActionTaskList`

NewWorkflowDynamicWorkflowActionTaskListWithDefaults instantiates a new WorkflowDynamicWorkflowActionTaskList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowDynamicWorkflowActionTaskList) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowDynamicWorkflowActionTaskList) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowDynamicWorkflowActionTaskList) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowDynamicWorkflowActionTaskList) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowDynamicWorkflowActionTaskList) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowDynamicWorkflowActionTaskList) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


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



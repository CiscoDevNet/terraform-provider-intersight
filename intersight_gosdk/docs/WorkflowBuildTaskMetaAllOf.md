# WorkflowBuildTaskMetaAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.BuildTaskMeta"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.BuildTaskMeta"]
**Name** | Pointer to **string** | Name for the BuildTaskMeta instance used to created a dynamic workflow. | [optional] [readonly] 
**Src** | Pointer to **string** | Microservice owner for the task in this workflow. | [optional] [readonly] 
**TaskList** | Pointer to **interface{}** | Task list used to build the dynamic workflow. | [optional] [readonly] 
**TaskType** | Pointer to **string** | The type of the task within this workflow. | [optional] [readonly] 
**WorkflowType** | Pointer to **string** | The type for the dynamic workflow. | [optional] [readonly] 

## Methods

### NewWorkflowBuildTaskMetaAllOf

`func NewWorkflowBuildTaskMetaAllOf(classId string, objectType string, ) *WorkflowBuildTaskMetaAllOf`

NewWorkflowBuildTaskMetaAllOf instantiates a new WorkflowBuildTaskMetaAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowBuildTaskMetaAllOfWithDefaults

`func NewWorkflowBuildTaskMetaAllOfWithDefaults() *WorkflowBuildTaskMetaAllOf`

NewWorkflowBuildTaskMetaAllOfWithDefaults instantiates a new WorkflowBuildTaskMetaAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowBuildTaskMetaAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowBuildTaskMetaAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowBuildTaskMetaAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowBuildTaskMetaAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowBuildTaskMetaAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowBuildTaskMetaAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetName

`func (o *WorkflowBuildTaskMetaAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowBuildTaskMetaAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowBuildTaskMetaAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowBuildTaskMetaAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSrc

`func (o *WorkflowBuildTaskMetaAllOf) GetSrc() string`

GetSrc returns the Src field if non-nil, zero value otherwise.

### GetSrcOk

`func (o *WorkflowBuildTaskMetaAllOf) GetSrcOk() (*string, bool)`

GetSrcOk returns a tuple with the Src field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrc

`func (o *WorkflowBuildTaskMetaAllOf) SetSrc(v string)`

SetSrc sets Src field to given value.

### HasSrc

`func (o *WorkflowBuildTaskMetaAllOf) HasSrc() bool`

HasSrc returns a boolean if a field has been set.

### GetTaskList

`func (o *WorkflowBuildTaskMetaAllOf) GetTaskList() interface{}`

GetTaskList returns the TaskList field if non-nil, zero value otherwise.

### GetTaskListOk

`func (o *WorkflowBuildTaskMetaAllOf) GetTaskListOk() (*interface{}, bool)`

GetTaskListOk returns a tuple with the TaskList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskList

`func (o *WorkflowBuildTaskMetaAllOf) SetTaskList(v interface{})`

SetTaskList sets TaskList field to given value.

### HasTaskList

`func (o *WorkflowBuildTaskMetaAllOf) HasTaskList() bool`

HasTaskList returns a boolean if a field has been set.

### SetTaskListNil

`func (o *WorkflowBuildTaskMetaAllOf) SetTaskListNil(b bool)`

 SetTaskListNil sets the value for TaskList to be an explicit nil

### UnsetTaskList
`func (o *WorkflowBuildTaskMetaAllOf) UnsetTaskList()`

UnsetTaskList ensures that no value is present for TaskList, not even an explicit nil
### GetTaskType

`func (o *WorkflowBuildTaskMetaAllOf) GetTaskType() string`

GetTaskType returns the TaskType field if non-nil, zero value otherwise.

### GetTaskTypeOk

`func (o *WorkflowBuildTaskMetaAllOf) GetTaskTypeOk() (*string, bool)`

GetTaskTypeOk returns a tuple with the TaskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskType

`func (o *WorkflowBuildTaskMetaAllOf) SetTaskType(v string)`

SetTaskType sets TaskType field to given value.

### HasTaskType

`func (o *WorkflowBuildTaskMetaAllOf) HasTaskType() bool`

HasTaskType returns a boolean if a field has been set.

### GetWorkflowType

`func (o *WorkflowBuildTaskMetaAllOf) GetWorkflowType() string`

GetWorkflowType returns the WorkflowType field if non-nil, zero value otherwise.

### GetWorkflowTypeOk

`func (o *WorkflowBuildTaskMetaAllOf) GetWorkflowTypeOk() (*string, bool)`

GetWorkflowTypeOk returns a tuple with the WorkflowType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowType

`func (o *WorkflowBuildTaskMetaAllOf) SetWorkflowType(v string)`

SetWorkflowType sets WorkflowType field to given value.

### HasWorkflowType

`func (o *WorkflowBuildTaskMetaAllOf) HasWorkflowType() bool`

HasWorkflowType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# WorkflowRollbackWorkflowTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.RollbackWorkflowTask"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.RollbackWorkflowTask"]
**Description** | Pointer to **string** | Description of the rollback task. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of TaskInfo that needs to be rolled back. | [optional] 
**RefName** | Pointer to **string** | Reference name of TaskInfo that need to be rolled back. | [optional] 
**RollbackCompleted** | Pointer to **bool** | Status of the rollback operation for the task. | [optional] [readonly] 
**RollbackTaskName** | Pointer to **string** | Name of TaskInfo that performs the rollback operation. | [optional] [readonly] 
**Status** | Pointer to **string** | Status of the rollback task. By default, task status will be not started. Task status will be set to completed on successful execution, otherwise it will be set to failed. * &#x60;NotStarted&#x60; - Status of rollback task when it is not started rollback. * &#x60;NotSupported&#x60; - Status of task when it is not supporting rollback. * &#x60;Completed&#x60; - Status of rollback task once execution is successful. * &#x60;Failed&#x60; - Status of rollback task when it is failed. | [optional] [readonly] [default to "NotStarted"]
**TaskInfoMoid** | Pointer to **string** | Moid of TaskInfo that supports rollback operation. | [optional] 
**TaskPath** | Pointer to **string** | Path of rollback task if it is inside sub-workflow. | [optional] [readonly] 

## Methods

### NewWorkflowRollbackWorkflowTask

`func NewWorkflowRollbackWorkflowTask(classId string, objectType string, ) *WorkflowRollbackWorkflowTask`

NewWorkflowRollbackWorkflowTask instantiates a new WorkflowRollbackWorkflowTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowRollbackWorkflowTaskWithDefaults

`func NewWorkflowRollbackWorkflowTaskWithDefaults() *WorkflowRollbackWorkflowTask`

NewWorkflowRollbackWorkflowTaskWithDefaults instantiates a new WorkflowRollbackWorkflowTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowRollbackWorkflowTask) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowRollbackWorkflowTask) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowRollbackWorkflowTask) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowRollbackWorkflowTask) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowRollbackWorkflowTask) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowRollbackWorkflowTask) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDescription

`func (o *WorkflowRollbackWorkflowTask) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkflowRollbackWorkflowTask) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkflowRollbackWorkflowTask) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkflowRollbackWorkflowTask) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *WorkflowRollbackWorkflowTask) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowRollbackWorkflowTask) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowRollbackWorkflowTask) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowRollbackWorkflowTask) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRefName

`func (o *WorkflowRollbackWorkflowTask) GetRefName() string`

GetRefName returns the RefName field if non-nil, zero value otherwise.

### GetRefNameOk

`func (o *WorkflowRollbackWorkflowTask) GetRefNameOk() (*string, bool)`

GetRefNameOk returns a tuple with the RefName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefName

`func (o *WorkflowRollbackWorkflowTask) SetRefName(v string)`

SetRefName sets RefName field to given value.

### HasRefName

`func (o *WorkflowRollbackWorkflowTask) HasRefName() bool`

HasRefName returns a boolean if a field has been set.

### GetRollbackCompleted

`func (o *WorkflowRollbackWorkflowTask) GetRollbackCompleted() bool`

GetRollbackCompleted returns the RollbackCompleted field if non-nil, zero value otherwise.

### GetRollbackCompletedOk

`func (o *WorkflowRollbackWorkflowTask) GetRollbackCompletedOk() (*bool, bool)`

GetRollbackCompletedOk returns a tuple with the RollbackCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollbackCompleted

`func (o *WorkflowRollbackWorkflowTask) SetRollbackCompleted(v bool)`

SetRollbackCompleted sets RollbackCompleted field to given value.

### HasRollbackCompleted

`func (o *WorkflowRollbackWorkflowTask) HasRollbackCompleted() bool`

HasRollbackCompleted returns a boolean if a field has been set.

### GetRollbackTaskName

`func (o *WorkflowRollbackWorkflowTask) GetRollbackTaskName() string`

GetRollbackTaskName returns the RollbackTaskName field if non-nil, zero value otherwise.

### GetRollbackTaskNameOk

`func (o *WorkflowRollbackWorkflowTask) GetRollbackTaskNameOk() (*string, bool)`

GetRollbackTaskNameOk returns a tuple with the RollbackTaskName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollbackTaskName

`func (o *WorkflowRollbackWorkflowTask) SetRollbackTaskName(v string)`

SetRollbackTaskName sets RollbackTaskName field to given value.

### HasRollbackTaskName

`func (o *WorkflowRollbackWorkflowTask) HasRollbackTaskName() bool`

HasRollbackTaskName returns a boolean if a field has been set.

### GetStatus

`func (o *WorkflowRollbackWorkflowTask) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WorkflowRollbackWorkflowTask) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WorkflowRollbackWorkflowTask) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WorkflowRollbackWorkflowTask) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTaskInfoMoid

`func (o *WorkflowRollbackWorkflowTask) GetTaskInfoMoid() string`

GetTaskInfoMoid returns the TaskInfoMoid field if non-nil, zero value otherwise.

### GetTaskInfoMoidOk

`func (o *WorkflowRollbackWorkflowTask) GetTaskInfoMoidOk() (*string, bool)`

GetTaskInfoMoidOk returns a tuple with the TaskInfoMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskInfoMoid

`func (o *WorkflowRollbackWorkflowTask) SetTaskInfoMoid(v string)`

SetTaskInfoMoid sets TaskInfoMoid field to given value.

### HasTaskInfoMoid

`func (o *WorkflowRollbackWorkflowTask) HasTaskInfoMoid() bool`

HasTaskInfoMoid returns a boolean if a field has been set.

### GetTaskPath

`func (o *WorkflowRollbackWorkflowTask) GetTaskPath() string`

GetTaskPath returns the TaskPath field if non-nil, zero value otherwise.

### GetTaskPathOk

`func (o *WorkflowRollbackWorkflowTask) GetTaskPathOk() (*string, bool)`

GetTaskPathOk returns a tuple with the TaskPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskPath

`func (o *WorkflowRollbackWorkflowTask) SetTaskPath(v string)`

SetTaskPath sets TaskPath field to given value.

### HasTaskPath

`func (o *WorkflowRollbackWorkflowTask) HasTaskPath() bool`

HasTaskPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



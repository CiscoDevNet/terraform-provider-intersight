# WorkflowRollbackWorkflowTaskAllOf

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

### NewWorkflowRollbackWorkflowTaskAllOf

`func NewWorkflowRollbackWorkflowTaskAllOf(classId string, objectType string, ) *WorkflowRollbackWorkflowTaskAllOf`

NewWorkflowRollbackWorkflowTaskAllOf instantiates a new WorkflowRollbackWorkflowTaskAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowRollbackWorkflowTaskAllOfWithDefaults

`func NewWorkflowRollbackWorkflowTaskAllOfWithDefaults() *WorkflowRollbackWorkflowTaskAllOf`

NewWorkflowRollbackWorkflowTaskAllOfWithDefaults instantiates a new WorkflowRollbackWorkflowTaskAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowRollbackWorkflowTaskAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowRollbackWorkflowTaskAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowRollbackWorkflowTaskAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowRollbackWorkflowTaskAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowRollbackWorkflowTaskAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowRollbackWorkflowTaskAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDescription

`func (o *WorkflowRollbackWorkflowTaskAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkflowRollbackWorkflowTaskAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkflowRollbackWorkflowTaskAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkflowRollbackWorkflowTaskAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *WorkflowRollbackWorkflowTaskAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowRollbackWorkflowTaskAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowRollbackWorkflowTaskAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowRollbackWorkflowTaskAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRefName

`func (o *WorkflowRollbackWorkflowTaskAllOf) GetRefName() string`

GetRefName returns the RefName field if non-nil, zero value otherwise.

### GetRefNameOk

`func (o *WorkflowRollbackWorkflowTaskAllOf) GetRefNameOk() (*string, bool)`

GetRefNameOk returns a tuple with the RefName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefName

`func (o *WorkflowRollbackWorkflowTaskAllOf) SetRefName(v string)`

SetRefName sets RefName field to given value.

### HasRefName

`func (o *WorkflowRollbackWorkflowTaskAllOf) HasRefName() bool`

HasRefName returns a boolean if a field has been set.

### GetRollbackCompleted

`func (o *WorkflowRollbackWorkflowTaskAllOf) GetRollbackCompleted() bool`

GetRollbackCompleted returns the RollbackCompleted field if non-nil, zero value otherwise.

### GetRollbackCompletedOk

`func (o *WorkflowRollbackWorkflowTaskAllOf) GetRollbackCompletedOk() (*bool, bool)`

GetRollbackCompletedOk returns a tuple with the RollbackCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollbackCompleted

`func (o *WorkflowRollbackWorkflowTaskAllOf) SetRollbackCompleted(v bool)`

SetRollbackCompleted sets RollbackCompleted field to given value.

### HasRollbackCompleted

`func (o *WorkflowRollbackWorkflowTaskAllOf) HasRollbackCompleted() bool`

HasRollbackCompleted returns a boolean if a field has been set.

### GetRollbackTaskName

`func (o *WorkflowRollbackWorkflowTaskAllOf) GetRollbackTaskName() string`

GetRollbackTaskName returns the RollbackTaskName field if non-nil, zero value otherwise.

### GetRollbackTaskNameOk

`func (o *WorkflowRollbackWorkflowTaskAllOf) GetRollbackTaskNameOk() (*string, bool)`

GetRollbackTaskNameOk returns a tuple with the RollbackTaskName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollbackTaskName

`func (o *WorkflowRollbackWorkflowTaskAllOf) SetRollbackTaskName(v string)`

SetRollbackTaskName sets RollbackTaskName field to given value.

### HasRollbackTaskName

`func (o *WorkflowRollbackWorkflowTaskAllOf) HasRollbackTaskName() bool`

HasRollbackTaskName returns a boolean if a field has been set.

### GetStatus

`func (o *WorkflowRollbackWorkflowTaskAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WorkflowRollbackWorkflowTaskAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WorkflowRollbackWorkflowTaskAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WorkflowRollbackWorkflowTaskAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTaskInfoMoid

`func (o *WorkflowRollbackWorkflowTaskAllOf) GetTaskInfoMoid() string`

GetTaskInfoMoid returns the TaskInfoMoid field if non-nil, zero value otherwise.

### GetTaskInfoMoidOk

`func (o *WorkflowRollbackWorkflowTaskAllOf) GetTaskInfoMoidOk() (*string, bool)`

GetTaskInfoMoidOk returns a tuple with the TaskInfoMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskInfoMoid

`func (o *WorkflowRollbackWorkflowTaskAllOf) SetTaskInfoMoid(v string)`

SetTaskInfoMoid sets TaskInfoMoid field to given value.

### HasTaskInfoMoid

`func (o *WorkflowRollbackWorkflowTaskAllOf) HasTaskInfoMoid() bool`

HasTaskInfoMoid returns a boolean if a field has been set.

### GetTaskPath

`func (o *WorkflowRollbackWorkflowTaskAllOf) GetTaskPath() string`

GetTaskPath returns the TaskPath field if non-nil, zero value otherwise.

### GetTaskPathOk

`func (o *WorkflowRollbackWorkflowTaskAllOf) GetTaskPathOk() (*string, bool)`

GetTaskPathOk returns a tuple with the TaskPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskPath

`func (o *WorkflowRollbackWorkflowTaskAllOf) SetTaskPath(v string)`

SetTaskPath sets TaskPath field to given value.

### HasTaskPath

`func (o *WorkflowRollbackWorkflowTaskAllOf) HasTaskPath() bool`

HasTaskPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



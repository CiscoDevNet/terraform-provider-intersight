# WorkflowRollbackTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.RollbackTask"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.RollbackTask"]
**CatalogMoid** | Pointer to **string** | The catalog under which the task definition has been added. | [optional] 
**Description** | Pointer to **string** | Description of rollback task definition. | [optional] 
**InputParameters** | Pointer to **interface{}** | Input parameters mapping for rollback task from the input or output of the main task definition. | [optional] 
**Name** | Pointer to **string** | Name of the task definition which is capable of doing rollback of this task. | [optional] 
**TaskMoid** | Pointer to **string** | The resolved referenced rollback task definition managed object. | [optional] 
**Version** | Pointer to **int64** | The version of the task definition. | [optional] 

## Methods

### NewWorkflowRollbackTask

`func NewWorkflowRollbackTask(classId string, objectType string, ) *WorkflowRollbackTask`

NewWorkflowRollbackTask instantiates a new WorkflowRollbackTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowRollbackTaskWithDefaults

`func NewWorkflowRollbackTaskWithDefaults() *WorkflowRollbackTask`

NewWorkflowRollbackTaskWithDefaults instantiates a new WorkflowRollbackTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowRollbackTask) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowRollbackTask) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowRollbackTask) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowRollbackTask) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowRollbackTask) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowRollbackTask) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCatalogMoid

`func (o *WorkflowRollbackTask) GetCatalogMoid() string`

GetCatalogMoid returns the CatalogMoid field if non-nil, zero value otherwise.

### GetCatalogMoidOk

`func (o *WorkflowRollbackTask) GetCatalogMoidOk() (*string, bool)`

GetCatalogMoidOk returns a tuple with the CatalogMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogMoid

`func (o *WorkflowRollbackTask) SetCatalogMoid(v string)`

SetCatalogMoid sets CatalogMoid field to given value.

### HasCatalogMoid

`func (o *WorkflowRollbackTask) HasCatalogMoid() bool`

HasCatalogMoid returns a boolean if a field has been set.

### GetDescription

`func (o *WorkflowRollbackTask) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkflowRollbackTask) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkflowRollbackTask) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkflowRollbackTask) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInputParameters

`func (o *WorkflowRollbackTask) GetInputParameters() interface{}`

GetInputParameters returns the InputParameters field if non-nil, zero value otherwise.

### GetInputParametersOk

`func (o *WorkflowRollbackTask) GetInputParametersOk() (*interface{}, bool)`

GetInputParametersOk returns a tuple with the InputParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputParameters

`func (o *WorkflowRollbackTask) SetInputParameters(v interface{})`

SetInputParameters sets InputParameters field to given value.

### HasInputParameters

`func (o *WorkflowRollbackTask) HasInputParameters() bool`

HasInputParameters returns a boolean if a field has been set.

### SetInputParametersNil

`func (o *WorkflowRollbackTask) SetInputParametersNil(b bool)`

 SetInputParametersNil sets the value for InputParameters to be an explicit nil

### UnsetInputParameters
`func (o *WorkflowRollbackTask) UnsetInputParameters()`

UnsetInputParameters ensures that no value is present for InputParameters, not even an explicit nil
### GetName

`func (o *WorkflowRollbackTask) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowRollbackTask) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowRollbackTask) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowRollbackTask) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTaskMoid

`func (o *WorkflowRollbackTask) GetTaskMoid() string`

GetTaskMoid returns the TaskMoid field if non-nil, zero value otherwise.

### GetTaskMoidOk

`func (o *WorkflowRollbackTask) GetTaskMoidOk() (*string, bool)`

GetTaskMoidOk returns a tuple with the TaskMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskMoid

`func (o *WorkflowRollbackTask) SetTaskMoid(v string)`

SetTaskMoid sets TaskMoid field to given value.

### HasTaskMoid

`func (o *WorkflowRollbackTask) HasTaskMoid() bool`

HasTaskMoid returns a boolean if a field has been set.

### GetVersion

`func (o *WorkflowRollbackTask) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *WorkflowRollbackTask) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *WorkflowRollbackTask) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *WorkflowRollbackTask) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



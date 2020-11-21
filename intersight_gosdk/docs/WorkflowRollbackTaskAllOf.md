# WorkflowRollbackTaskAllOf

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

### NewWorkflowRollbackTaskAllOf

`func NewWorkflowRollbackTaskAllOf(classId string, objectType string, ) *WorkflowRollbackTaskAllOf`

NewWorkflowRollbackTaskAllOf instantiates a new WorkflowRollbackTaskAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowRollbackTaskAllOfWithDefaults

`func NewWorkflowRollbackTaskAllOfWithDefaults() *WorkflowRollbackTaskAllOf`

NewWorkflowRollbackTaskAllOfWithDefaults instantiates a new WorkflowRollbackTaskAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowRollbackTaskAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowRollbackTaskAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowRollbackTaskAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowRollbackTaskAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowRollbackTaskAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowRollbackTaskAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCatalogMoid

`func (o *WorkflowRollbackTaskAllOf) GetCatalogMoid() string`

GetCatalogMoid returns the CatalogMoid field if non-nil, zero value otherwise.

### GetCatalogMoidOk

`func (o *WorkflowRollbackTaskAllOf) GetCatalogMoidOk() (*string, bool)`

GetCatalogMoidOk returns a tuple with the CatalogMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogMoid

`func (o *WorkflowRollbackTaskAllOf) SetCatalogMoid(v string)`

SetCatalogMoid sets CatalogMoid field to given value.

### HasCatalogMoid

`func (o *WorkflowRollbackTaskAllOf) HasCatalogMoid() bool`

HasCatalogMoid returns a boolean if a field has been set.

### GetDescription

`func (o *WorkflowRollbackTaskAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkflowRollbackTaskAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkflowRollbackTaskAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkflowRollbackTaskAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInputParameters

`func (o *WorkflowRollbackTaskAllOf) GetInputParameters() interface{}`

GetInputParameters returns the InputParameters field if non-nil, zero value otherwise.

### GetInputParametersOk

`func (o *WorkflowRollbackTaskAllOf) GetInputParametersOk() (*interface{}, bool)`

GetInputParametersOk returns a tuple with the InputParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputParameters

`func (o *WorkflowRollbackTaskAllOf) SetInputParameters(v interface{})`

SetInputParameters sets InputParameters field to given value.

### HasInputParameters

`func (o *WorkflowRollbackTaskAllOf) HasInputParameters() bool`

HasInputParameters returns a boolean if a field has been set.

### SetInputParametersNil

`func (o *WorkflowRollbackTaskAllOf) SetInputParametersNil(b bool)`

 SetInputParametersNil sets the value for InputParameters to be an explicit nil

### UnsetInputParameters
`func (o *WorkflowRollbackTaskAllOf) UnsetInputParameters()`

UnsetInputParameters ensures that no value is present for InputParameters, not even an explicit nil
### GetName

`func (o *WorkflowRollbackTaskAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowRollbackTaskAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowRollbackTaskAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowRollbackTaskAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTaskMoid

`func (o *WorkflowRollbackTaskAllOf) GetTaskMoid() string`

GetTaskMoid returns the TaskMoid field if non-nil, zero value otherwise.

### GetTaskMoidOk

`func (o *WorkflowRollbackTaskAllOf) GetTaskMoidOk() (*string, bool)`

GetTaskMoidOk returns a tuple with the TaskMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskMoid

`func (o *WorkflowRollbackTaskAllOf) SetTaskMoid(v string)`

SetTaskMoid sets TaskMoid field to given value.

### HasTaskMoid

`func (o *WorkflowRollbackTaskAllOf) HasTaskMoid() bool`

HasTaskMoid returns a boolean if a field has been set.

### GetVersion

`func (o *WorkflowRollbackTaskAllOf) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *WorkflowRollbackTaskAllOf) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *WorkflowRollbackTaskAllOf) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *WorkflowRollbackTaskAllOf) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



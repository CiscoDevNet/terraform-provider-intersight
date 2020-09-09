# WorkflowWorkerTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CatalogMoid** | Pointer to **string** | Specify the catalog moid that this task belongs. | [optional] 
**TaskDefinitionId** | Pointer to **string** | The resolved referenced task definition managed object. | [optional] [readonly] 
**TaskDefinitionName** | Pointer to **string** | The qualified name of task that should be executed. | [optional] 
**Version** | Pointer to **int64** | The task definition version to use in this workflow. When no version is specified then the default version of the task at the time of creating or updating this workflow is used. | [optional] 

## Methods

### NewWorkflowWorkerTask

`func NewWorkflowWorkerTask() *WorkflowWorkerTask`

NewWorkflowWorkerTask instantiates a new WorkflowWorkerTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowWorkerTaskWithDefaults

`func NewWorkflowWorkerTaskWithDefaults() *WorkflowWorkerTask`

NewWorkflowWorkerTaskWithDefaults instantiates a new WorkflowWorkerTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCatalogMoid

`func (o *WorkflowWorkerTask) GetCatalogMoid() string`

GetCatalogMoid returns the CatalogMoid field if non-nil, zero value otherwise.

### GetCatalogMoidOk

`func (o *WorkflowWorkerTask) GetCatalogMoidOk() (*string, bool)`

GetCatalogMoidOk returns a tuple with the CatalogMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogMoid

`func (o *WorkflowWorkerTask) SetCatalogMoid(v string)`

SetCatalogMoid sets CatalogMoid field to given value.

### HasCatalogMoid

`func (o *WorkflowWorkerTask) HasCatalogMoid() bool`

HasCatalogMoid returns a boolean if a field has been set.

### GetTaskDefinitionId

`func (o *WorkflowWorkerTask) GetTaskDefinitionId() string`

GetTaskDefinitionId returns the TaskDefinitionId field if non-nil, zero value otherwise.

### GetTaskDefinitionIdOk

`func (o *WorkflowWorkerTask) GetTaskDefinitionIdOk() (*string, bool)`

GetTaskDefinitionIdOk returns a tuple with the TaskDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskDefinitionId

`func (o *WorkflowWorkerTask) SetTaskDefinitionId(v string)`

SetTaskDefinitionId sets TaskDefinitionId field to given value.

### HasTaskDefinitionId

`func (o *WorkflowWorkerTask) HasTaskDefinitionId() bool`

HasTaskDefinitionId returns a boolean if a field has been set.

### GetTaskDefinitionName

`func (o *WorkflowWorkerTask) GetTaskDefinitionName() string`

GetTaskDefinitionName returns the TaskDefinitionName field if non-nil, zero value otherwise.

### GetTaskDefinitionNameOk

`func (o *WorkflowWorkerTask) GetTaskDefinitionNameOk() (*string, bool)`

GetTaskDefinitionNameOk returns a tuple with the TaskDefinitionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskDefinitionName

`func (o *WorkflowWorkerTask) SetTaskDefinitionName(v string)`

SetTaskDefinitionName sets TaskDefinitionName field to given value.

### HasTaskDefinitionName

`func (o *WorkflowWorkerTask) HasTaskDefinitionName() bool`

HasTaskDefinitionName returns a boolean if a field has been set.

### GetVersion

`func (o *WorkflowWorkerTask) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *WorkflowWorkerTask) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *WorkflowWorkerTask) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *WorkflowWorkerTask) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



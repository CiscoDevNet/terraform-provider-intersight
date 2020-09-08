# WorkflowSubWorkflowTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CatalogMoid** | Pointer to **string** | Specify the catalog moid that this task belongs. | [optional] 
**Version** | Pointer to **int64** | The workflow definition version to use as subworkflow. When no version is specified then the default version of the workflow at the time of creating or updating this workflow is used. | [optional] 
**WorkflowDefinitionId** | Pointer to **string** | The resolved referenced workflow definition managed object. | [optional] [readonly] 
**WorkflowDefinitionName** | Pointer to **string** | The qualified name of workflow that should be executed as a task. | [optional] 

## Methods

### NewWorkflowSubWorkflowTask

`func NewWorkflowSubWorkflowTask() *WorkflowSubWorkflowTask`

NewWorkflowSubWorkflowTask instantiates a new WorkflowSubWorkflowTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowSubWorkflowTaskWithDefaults

`func NewWorkflowSubWorkflowTaskWithDefaults() *WorkflowSubWorkflowTask`

NewWorkflowSubWorkflowTaskWithDefaults instantiates a new WorkflowSubWorkflowTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCatalogMoid

`func (o *WorkflowSubWorkflowTask) GetCatalogMoid() string`

GetCatalogMoid returns the CatalogMoid field if non-nil, zero value otherwise.

### GetCatalogMoidOk

`func (o *WorkflowSubWorkflowTask) GetCatalogMoidOk() (*string, bool)`

GetCatalogMoidOk returns a tuple with the CatalogMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogMoid

`func (o *WorkflowSubWorkflowTask) SetCatalogMoid(v string)`

SetCatalogMoid sets CatalogMoid field to given value.

### HasCatalogMoid

`func (o *WorkflowSubWorkflowTask) HasCatalogMoid() bool`

HasCatalogMoid returns a boolean if a field has been set.

### GetVersion

`func (o *WorkflowSubWorkflowTask) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *WorkflowSubWorkflowTask) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *WorkflowSubWorkflowTask) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *WorkflowSubWorkflowTask) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetWorkflowDefinitionId

`func (o *WorkflowSubWorkflowTask) GetWorkflowDefinitionId() string`

GetWorkflowDefinitionId returns the WorkflowDefinitionId field if non-nil, zero value otherwise.

### GetWorkflowDefinitionIdOk

`func (o *WorkflowSubWorkflowTask) GetWorkflowDefinitionIdOk() (*string, bool)`

GetWorkflowDefinitionIdOk returns a tuple with the WorkflowDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowDefinitionId

`func (o *WorkflowSubWorkflowTask) SetWorkflowDefinitionId(v string)`

SetWorkflowDefinitionId sets WorkflowDefinitionId field to given value.

### HasWorkflowDefinitionId

`func (o *WorkflowSubWorkflowTask) HasWorkflowDefinitionId() bool`

HasWorkflowDefinitionId returns a boolean if a field has been set.

### GetWorkflowDefinitionName

`func (o *WorkflowSubWorkflowTask) GetWorkflowDefinitionName() string`

GetWorkflowDefinitionName returns the WorkflowDefinitionName field if non-nil, zero value otherwise.

### GetWorkflowDefinitionNameOk

`func (o *WorkflowSubWorkflowTask) GetWorkflowDefinitionNameOk() (*string, bool)`

GetWorkflowDefinitionNameOk returns a tuple with the WorkflowDefinitionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowDefinitionName

`func (o *WorkflowSubWorkflowTask) SetWorkflowDefinitionName(v string)`

SetWorkflowDefinitionName sets WorkflowDefinitionName field to given value.

### HasWorkflowDefinitionName

`func (o *WorkflowSubWorkflowTask) HasWorkflowDefinitionName() bool`

HasWorkflowDefinitionName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



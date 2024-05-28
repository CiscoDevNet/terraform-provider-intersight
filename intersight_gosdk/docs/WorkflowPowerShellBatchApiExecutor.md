# WorkflowPowerShellBatchApiExecutor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.PowerShellBatchApiExecutor"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.PowerShellBatchApiExecutor"]
**ErrorResponseHandler** | Pointer to [**NullableWorkflowErrorResponseHandlerRelationship**](WorkflowErrorResponseHandlerRelationship.md) |  | [optional] 
**TaskDefinition** | Pointer to [**NullableWorkflowTaskDefinitionRelationship**](WorkflowTaskDefinitionRelationship.md) |  | [optional] 

## Methods

### NewWorkflowPowerShellBatchApiExecutor

`func NewWorkflowPowerShellBatchApiExecutor(classId string, objectType string, ) *WorkflowPowerShellBatchApiExecutor`

NewWorkflowPowerShellBatchApiExecutor instantiates a new WorkflowPowerShellBatchApiExecutor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowPowerShellBatchApiExecutorWithDefaults

`func NewWorkflowPowerShellBatchApiExecutorWithDefaults() *WorkflowPowerShellBatchApiExecutor`

NewWorkflowPowerShellBatchApiExecutorWithDefaults instantiates a new WorkflowPowerShellBatchApiExecutor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowPowerShellBatchApiExecutor) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowPowerShellBatchApiExecutor) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowPowerShellBatchApiExecutor) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowPowerShellBatchApiExecutor) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowPowerShellBatchApiExecutor) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowPowerShellBatchApiExecutor) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetErrorResponseHandler

`func (o *WorkflowPowerShellBatchApiExecutor) GetErrorResponseHandler() WorkflowErrorResponseHandlerRelationship`

GetErrorResponseHandler returns the ErrorResponseHandler field if non-nil, zero value otherwise.

### GetErrorResponseHandlerOk

`func (o *WorkflowPowerShellBatchApiExecutor) GetErrorResponseHandlerOk() (*WorkflowErrorResponseHandlerRelationship, bool)`

GetErrorResponseHandlerOk returns a tuple with the ErrorResponseHandler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorResponseHandler

`func (o *WorkflowPowerShellBatchApiExecutor) SetErrorResponseHandler(v WorkflowErrorResponseHandlerRelationship)`

SetErrorResponseHandler sets ErrorResponseHandler field to given value.

### HasErrorResponseHandler

`func (o *WorkflowPowerShellBatchApiExecutor) HasErrorResponseHandler() bool`

HasErrorResponseHandler returns a boolean if a field has been set.

### SetErrorResponseHandlerNil

`func (o *WorkflowPowerShellBatchApiExecutor) SetErrorResponseHandlerNil(b bool)`

 SetErrorResponseHandlerNil sets the value for ErrorResponseHandler to be an explicit nil

### UnsetErrorResponseHandler
`func (o *WorkflowPowerShellBatchApiExecutor) UnsetErrorResponseHandler()`

UnsetErrorResponseHandler ensures that no value is present for ErrorResponseHandler, not even an explicit nil
### GetTaskDefinition

`func (o *WorkflowPowerShellBatchApiExecutor) GetTaskDefinition() WorkflowTaskDefinitionRelationship`

GetTaskDefinition returns the TaskDefinition field if non-nil, zero value otherwise.

### GetTaskDefinitionOk

`func (o *WorkflowPowerShellBatchApiExecutor) GetTaskDefinitionOk() (*WorkflowTaskDefinitionRelationship, bool)`

GetTaskDefinitionOk returns a tuple with the TaskDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskDefinition

`func (o *WorkflowPowerShellBatchApiExecutor) SetTaskDefinition(v WorkflowTaskDefinitionRelationship)`

SetTaskDefinition sets TaskDefinition field to given value.

### HasTaskDefinition

`func (o *WorkflowPowerShellBatchApiExecutor) HasTaskDefinition() bool`

HasTaskDefinition returns a boolean if a field has been set.

### SetTaskDefinitionNil

`func (o *WorkflowPowerShellBatchApiExecutor) SetTaskDefinitionNil(b bool)`

 SetTaskDefinitionNil sets the value for TaskDefinition to be an explicit nil

### UnsetTaskDefinition
`func (o *WorkflowPowerShellBatchApiExecutor) UnsetTaskDefinition()`

UnsetTaskDefinition ensures that no value is present for TaskDefinition, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



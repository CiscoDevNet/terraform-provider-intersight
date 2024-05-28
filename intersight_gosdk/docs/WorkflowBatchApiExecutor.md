# WorkflowBatchApiExecutor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.BatchApiExecutor"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.BatchApiExecutor"]
**ErrorResponseHandler** | Pointer to [**NullableWorkflowErrorResponseHandlerRelationship**](WorkflowErrorResponseHandlerRelationship.md) |  | [optional] 
**TaskDefinition** | Pointer to [**NullableWorkflowTaskDefinitionRelationship**](WorkflowTaskDefinitionRelationship.md) |  | [optional] 

## Methods

### NewWorkflowBatchApiExecutor

`func NewWorkflowBatchApiExecutor(classId string, objectType string, ) *WorkflowBatchApiExecutor`

NewWorkflowBatchApiExecutor instantiates a new WorkflowBatchApiExecutor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowBatchApiExecutorWithDefaults

`func NewWorkflowBatchApiExecutorWithDefaults() *WorkflowBatchApiExecutor`

NewWorkflowBatchApiExecutorWithDefaults instantiates a new WorkflowBatchApiExecutor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowBatchApiExecutor) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowBatchApiExecutor) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowBatchApiExecutor) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowBatchApiExecutor) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowBatchApiExecutor) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowBatchApiExecutor) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetErrorResponseHandler

`func (o *WorkflowBatchApiExecutor) GetErrorResponseHandler() WorkflowErrorResponseHandlerRelationship`

GetErrorResponseHandler returns the ErrorResponseHandler field if non-nil, zero value otherwise.

### GetErrorResponseHandlerOk

`func (o *WorkflowBatchApiExecutor) GetErrorResponseHandlerOk() (*WorkflowErrorResponseHandlerRelationship, bool)`

GetErrorResponseHandlerOk returns a tuple with the ErrorResponseHandler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorResponseHandler

`func (o *WorkflowBatchApiExecutor) SetErrorResponseHandler(v WorkflowErrorResponseHandlerRelationship)`

SetErrorResponseHandler sets ErrorResponseHandler field to given value.

### HasErrorResponseHandler

`func (o *WorkflowBatchApiExecutor) HasErrorResponseHandler() bool`

HasErrorResponseHandler returns a boolean if a field has been set.

### SetErrorResponseHandlerNil

`func (o *WorkflowBatchApiExecutor) SetErrorResponseHandlerNil(b bool)`

 SetErrorResponseHandlerNil sets the value for ErrorResponseHandler to be an explicit nil

### UnsetErrorResponseHandler
`func (o *WorkflowBatchApiExecutor) UnsetErrorResponseHandler()`

UnsetErrorResponseHandler ensures that no value is present for ErrorResponseHandler, not even an explicit nil
### GetTaskDefinition

`func (o *WorkflowBatchApiExecutor) GetTaskDefinition() WorkflowTaskDefinitionRelationship`

GetTaskDefinition returns the TaskDefinition field if non-nil, zero value otherwise.

### GetTaskDefinitionOk

`func (o *WorkflowBatchApiExecutor) GetTaskDefinitionOk() (*WorkflowTaskDefinitionRelationship, bool)`

GetTaskDefinitionOk returns a tuple with the TaskDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskDefinition

`func (o *WorkflowBatchApiExecutor) SetTaskDefinition(v WorkflowTaskDefinitionRelationship)`

SetTaskDefinition sets TaskDefinition field to given value.

### HasTaskDefinition

`func (o *WorkflowBatchApiExecutor) HasTaskDefinition() bool`

HasTaskDefinition returns a boolean if a field has been set.

### SetTaskDefinitionNil

`func (o *WorkflowBatchApiExecutor) SetTaskDefinitionNil(b bool)`

 SetTaskDefinitionNil sets the value for TaskDefinition to be an explicit nil

### UnsetTaskDefinition
`func (o *WorkflowBatchApiExecutor) UnsetTaskDefinition()`

UnsetTaskDefinition ensures that no value is present for TaskDefinition, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



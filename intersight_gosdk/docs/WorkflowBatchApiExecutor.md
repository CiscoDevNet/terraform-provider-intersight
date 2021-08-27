# WorkflowBatchApiExecutor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.BatchApiExecutor"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.BatchApiExecutor"]
**Batch** | Pointer to [**[]WorkflowApi**](WorkflowApi.md) |  | [optional] 
**Constraints** | Pointer to [**NullableWorkflowTaskConstraints**](WorkflowTaskConstraints.md) |  | [optional] 
**Description** | Pointer to **string** | A detailed description about the batch APIs. | [optional] 
**Name** | Pointer to **string** | Name for the batch API task. | [optional] 
**Outcomes** | Pointer to **interface{}** | All the possible outcomes of this task are captured here. Outcomes property is a collection property of type workflow.Outcome objects. The outcomes can be mapped to the message to be shown. The outcomes are evaluated in the order they are given. At the end of the outcomes list, an catchall success/fail outcome can be added with condition as &#39;true&#39;. This is an optional property and if not specified the task will be marked as success. | [optional] 
**Output** | Pointer to **interface{}** | Intersight Orchestrator allows the extraction of required values from API responses using the API response grammar. These extracted values can be mapped to task output parameters defined in task definition. The mapping of API output parameters to the task output parameters is provided as JSON in this property. | [optional] 
**RetryFromFailedApi** | Pointer to **bool** | When an execution of a nth API in the Batch fails, Retry from falied API flag indicates if the execution should start from the nth API or the first API during task retry. By default the value is set to false. | [optional] 
**SkipOnCondition** | Pointer to **string** | The skip expression, if provided, allows the batch API executor to skip the task execution when the given expression evaluates to true. The expression is given as such a golang template that has to be evaluated to a final content true/false. The expression is an optional and in case not provided, the API will always be executed. | [optional] 
**ErrorResponseHandler** | Pointer to [**WorkflowErrorResponseHandlerRelationship**](WorkflowErrorResponseHandlerRelationship.md) |  | [optional] 
**TaskDefinition** | Pointer to [**WorkflowTaskDefinitionRelationship**](WorkflowTaskDefinitionRelationship.md) |  | [optional] 

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


### GetBatch

`func (o *WorkflowBatchApiExecutor) GetBatch() []WorkflowApi`

GetBatch returns the Batch field if non-nil, zero value otherwise.

### GetBatchOk

`func (o *WorkflowBatchApiExecutor) GetBatchOk() (*[]WorkflowApi, bool)`

GetBatchOk returns a tuple with the Batch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatch

`func (o *WorkflowBatchApiExecutor) SetBatch(v []WorkflowApi)`

SetBatch sets Batch field to given value.

### HasBatch

`func (o *WorkflowBatchApiExecutor) HasBatch() bool`

HasBatch returns a boolean if a field has been set.

### SetBatchNil

`func (o *WorkflowBatchApiExecutor) SetBatchNil(b bool)`

 SetBatchNil sets the value for Batch to be an explicit nil

### UnsetBatch
`func (o *WorkflowBatchApiExecutor) UnsetBatch()`

UnsetBatch ensures that no value is present for Batch, not even an explicit nil
### GetConstraints

`func (o *WorkflowBatchApiExecutor) GetConstraints() WorkflowTaskConstraints`

GetConstraints returns the Constraints field if non-nil, zero value otherwise.

### GetConstraintsOk

`func (o *WorkflowBatchApiExecutor) GetConstraintsOk() (*WorkflowTaskConstraints, bool)`

GetConstraintsOk returns a tuple with the Constraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraints

`func (o *WorkflowBatchApiExecutor) SetConstraints(v WorkflowTaskConstraints)`

SetConstraints sets Constraints field to given value.

### HasConstraints

`func (o *WorkflowBatchApiExecutor) HasConstraints() bool`

HasConstraints returns a boolean if a field has been set.

### SetConstraintsNil

`func (o *WorkflowBatchApiExecutor) SetConstraintsNil(b bool)`

 SetConstraintsNil sets the value for Constraints to be an explicit nil

### UnsetConstraints
`func (o *WorkflowBatchApiExecutor) UnsetConstraints()`

UnsetConstraints ensures that no value is present for Constraints, not even an explicit nil
### GetDescription

`func (o *WorkflowBatchApiExecutor) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkflowBatchApiExecutor) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkflowBatchApiExecutor) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkflowBatchApiExecutor) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *WorkflowBatchApiExecutor) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowBatchApiExecutor) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowBatchApiExecutor) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowBatchApiExecutor) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOutcomes

`func (o *WorkflowBatchApiExecutor) GetOutcomes() interface{}`

GetOutcomes returns the Outcomes field if non-nil, zero value otherwise.

### GetOutcomesOk

`func (o *WorkflowBatchApiExecutor) GetOutcomesOk() (*interface{}, bool)`

GetOutcomesOk returns a tuple with the Outcomes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcomes

`func (o *WorkflowBatchApiExecutor) SetOutcomes(v interface{})`

SetOutcomes sets Outcomes field to given value.

### HasOutcomes

`func (o *WorkflowBatchApiExecutor) HasOutcomes() bool`

HasOutcomes returns a boolean if a field has been set.

### SetOutcomesNil

`func (o *WorkflowBatchApiExecutor) SetOutcomesNil(b bool)`

 SetOutcomesNil sets the value for Outcomes to be an explicit nil

### UnsetOutcomes
`func (o *WorkflowBatchApiExecutor) UnsetOutcomes()`

UnsetOutcomes ensures that no value is present for Outcomes, not even an explicit nil
### GetOutput

`func (o *WorkflowBatchApiExecutor) GetOutput() interface{}`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *WorkflowBatchApiExecutor) GetOutputOk() (*interface{}, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *WorkflowBatchApiExecutor) SetOutput(v interface{})`

SetOutput sets Output field to given value.

### HasOutput

`func (o *WorkflowBatchApiExecutor) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### SetOutputNil

`func (o *WorkflowBatchApiExecutor) SetOutputNil(b bool)`

 SetOutputNil sets the value for Output to be an explicit nil

### UnsetOutput
`func (o *WorkflowBatchApiExecutor) UnsetOutput()`

UnsetOutput ensures that no value is present for Output, not even an explicit nil
### GetRetryFromFailedApi

`func (o *WorkflowBatchApiExecutor) GetRetryFromFailedApi() bool`

GetRetryFromFailedApi returns the RetryFromFailedApi field if non-nil, zero value otherwise.

### GetRetryFromFailedApiOk

`func (o *WorkflowBatchApiExecutor) GetRetryFromFailedApiOk() (*bool, bool)`

GetRetryFromFailedApiOk returns a tuple with the RetryFromFailedApi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryFromFailedApi

`func (o *WorkflowBatchApiExecutor) SetRetryFromFailedApi(v bool)`

SetRetryFromFailedApi sets RetryFromFailedApi field to given value.

### HasRetryFromFailedApi

`func (o *WorkflowBatchApiExecutor) HasRetryFromFailedApi() bool`

HasRetryFromFailedApi returns a boolean if a field has been set.

### GetSkipOnCondition

`func (o *WorkflowBatchApiExecutor) GetSkipOnCondition() string`

GetSkipOnCondition returns the SkipOnCondition field if non-nil, zero value otherwise.

### GetSkipOnConditionOk

`func (o *WorkflowBatchApiExecutor) GetSkipOnConditionOk() (*string, bool)`

GetSkipOnConditionOk returns a tuple with the SkipOnCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipOnCondition

`func (o *WorkflowBatchApiExecutor) SetSkipOnCondition(v string)`

SetSkipOnCondition sets SkipOnCondition field to given value.

### HasSkipOnCondition

`func (o *WorkflowBatchApiExecutor) HasSkipOnCondition() bool`

HasSkipOnCondition returns a boolean if a field has been set.

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



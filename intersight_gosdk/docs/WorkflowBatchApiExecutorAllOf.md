# WorkflowBatchApiExecutorAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Batch** | Pointer to [**[]WorkflowApi**](workflow.Api.md) |  | [optional] 
**Constraints** | Pointer to [**WorkflowTaskConstraints**](workflow.TaskConstraints.md) |  | [optional] 
**Description** | Pointer to **string** | A detailed description about the batch APIs. | [optional] 
**Name** | Pointer to **string** | Name for the batch API task. | [optional] 
**Outcomes** | Pointer to **interface{}** | All the possible outcomes of this task are captured here. Outcomes property is a collection property of type workflow.Outcome objects. The outcomes can be mapped to the message to be shown. The outcomes are evaluated in the order they are given. At the end of the outcomes list, an catchall success/fail outcome can be added with condition as &#39;true&#39;. This is an optional property and if not specified the task will be marked as success. | [optional] 
**Output** | Pointer to **interface{}** | Intersight Orchestrator allows the extraction of required values from API responses using the API response grammar. These extracted values can be mapped to task output parameters defined in task definition. The mapping of API output parameters to the task output parameters is provided as JSON in this property. | [optional] 
**RetryFromFailedApi** | Pointer to **bool** | When an execution of a nth API in the Batch fails, Retry from falied API flag indicates if the execution should start from the nth API or the first API during task retry. By default the value is set to false. | [optional] 
**SkipOnCondition** | Pointer to **string** | The skip expression, if provided, allows the batch API executor to skip the task execution when the given expression evaluates to true. The expression is given as such a golang template that has to be evaluated to a final content true/false. The expression is an optional and in case not provided, the API will always be executed. | [optional] 
**ErrorResponseHandler** | Pointer to [**WorkflowErrorResponseHandlerRelationship**](workflow.ErrorResponseHandler.Relationship.md) |  | [optional] 
**TaskDefinition** | Pointer to [**WorkflowTaskDefinitionRelationship**](workflow.TaskDefinition.Relationship.md) |  | [optional] 

## Methods

### NewWorkflowBatchApiExecutorAllOf

`func NewWorkflowBatchApiExecutorAllOf() *WorkflowBatchApiExecutorAllOf`

NewWorkflowBatchApiExecutorAllOf instantiates a new WorkflowBatchApiExecutorAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowBatchApiExecutorAllOfWithDefaults

`func NewWorkflowBatchApiExecutorAllOfWithDefaults() *WorkflowBatchApiExecutorAllOf`

NewWorkflowBatchApiExecutorAllOfWithDefaults instantiates a new WorkflowBatchApiExecutorAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBatch

`func (o *WorkflowBatchApiExecutorAllOf) GetBatch() []WorkflowApi`

GetBatch returns the Batch field if non-nil, zero value otherwise.

### GetBatchOk

`func (o *WorkflowBatchApiExecutorAllOf) GetBatchOk() (*[]WorkflowApi, bool)`

GetBatchOk returns a tuple with the Batch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatch

`func (o *WorkflowBatchApiExecutorAllOf) SetBatch(v []WorkflowApi)`

SetBatch sets Batch field to given value.

### HasBatch

`func (o *WorkflowBatchApiExecutorAllOf) HasBatch() bool`

HasBatch returns a boolean if a field has been set.

### GetConstraints

`func (o *WorkflowBatchApiExecutorAllOf) GetConstraints() WorkflowTaskConstraints`

GetConstraints returns the Constraints field if non-nil, zero value otherwise.

### GetConstraintsOk

`func (o *WorkflowBatchApiExecutorAllOf) GetConstraintsOk() (*WorkflowTaskConstraints, bool)`

GetConstraintsOk returns a tuple with the Constraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraints

`func (o *WorkflowBatchApiExecutorAllOf) SetConstraints(v WorkflowTaskConstraints)`

SetConstraints sets Constraints field to given value.

### HasConstraints

`func (o *WorkflowBatchApiExecutorAllOf) HasConstraints() bool`

HasConstraints returns a boolean if a field has been set.

### GetDescription

`func (o *WorkflowBatchApiExecutorAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkflowBatchApiExecutorAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkflowBatchApiExecutorAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkflowBatchApiExecutorAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *WorkflowBatchApiExecutorAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowBatchApiExecutorAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowBatchApiExecutorAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowBatchApiExecutorAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOutcomes

`func (o *WorkflowBatchApiExecutorAllOf) GetOutcomes() interface{}`

GetOutcomes returns the Outcomes field if non-nil, zero value otherwise.

### GetOutcomesOk

`func (o *WorkflowBatchApiExecutorAllOf) GetOutcomesOk() (*interface{}, bool)`

GetOutcomesOk returns a tuple with the Outcomes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcomes

`func (o *WorkflowBatchApiExecutorAllOf) SetOutcomes(v interface{})`

SetOutcomes sets Outcomes field to given value.

### HasOutcomes

`func (o *WorkflowBatchApiExecutorAllOf) HasOutcomes() bool`

HasOutcomes returns a boolean if a field has been set.

### SetOutcomesNil

`func (o *WorkflowBatchApiExecutorAllOf) SetOutcomesNil(b bool)`

 SetOutcomesNil sets the value for Outcomes to be an explicit nil

### UnsetOutcomes
`func (o *WorkflowBatchApiExecutorAllOf) UnsetOutcomes()`

UnsetOutcomes ensures that no value is present for Outcomes, not even an explicit nil
### GetOutput

`func (o *WorkflowBatchApiExecutorAllOf) GetOutput() interface{}`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *WorkflowBatchApiExecutorAllOf) GetOutputOk() (*interface{}, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *WorkflowBatchApiExecutorAllOf) SetOutput(v interface{})`

SetOutput sets Output field to given value.

### HasOutput

`func (o *WorkflowBatchApiExecutorAllOf) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### SetOutputNil

`func (o *WorkflowBatchApiExecutorAllOf) SetOutputNil(b bool)`

 SetOutputNil sets the value for Output to be an explicit nil

### UnsetOutput
`func (o *WorkflowBatchApiExecutorAllOf) UnsetOutput()`

UnsetOutput ensures that no value is present for Output, not even an explicit nil
### GetRetryFromFailedApi

`func (o *WorkflowBatchApiExecutorAllOf) GetRetryFromFailedApi() bool`

GetRetryFromFailedApi returns the RetryFromFailedApi field if non-nil, zero value otherwise.

### GetRetryFromFailedApiOk

`func (o *WorkflowBatchApiExecutorAllOf) GetRetryFromFailedApiOk() (*bool, bool)`

GetRetryFromFailedApiOk returns a tuple with the RetryFromFailedApi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryFromFailedApi

`func (o *WorkflowBatchApiExecutorAllOf) SetRetryFromFailedApi(v bool)`

SetRetryFromFailedApi sets RetryFromFailedApi field to given value.

### HasRetryFromFailedApi

`func (o *WorkflowBatchApiExecutorAllOf) HasRetryFromFailedApi() bool`

HasRetryFromFailedApi returns a boolean if a field has been set.

### GetSkipOnCondition

`func (o *WorkflowBatchApiExecutorAllOf) GetSkipOnCondition() string`

GetSkipOnCondition returns the SkipOnCondition field if non-nil, zero value otherwise.

### GetSkipOnConditionOk

`func (o *WorkflowBatchApiExecutorAllOf) GetSkipOnConditionOk() (*string, bool)`

GetSkipOnConditionOk returns a tuple with the SkipOnCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipOnCondition

`func (o *WorkflowBatchApiExecutorAllOf) SetSkipOnCondition(v string)`

SetSkipOnCondition sets SkipOnCondition field to given value.

### HasSkipOnCondition

`func (o *WorkflowBatchApiExecutorAllOf) HasSkipOnCondition() bool`

HasSkipOnCondition returns a boolean if a field has been set.

### GetErrorResponseHandler

`func (o *WorkflowBatchApiExecutorAllOf) GetErrorResponseHandler() WorkflowErrorResponseHandlerRelationship`

GetErrorResponseHandler returns the ErrorResponseHandler field if non-nil, zero value otherwise.

### GetErrorResponseHandlerOk

`func (o *WorkflowBatchApiExecutorAllOf) GetErrorResponseHandlerOk() (*WorkflowErrorResponseHandlerRelationship, bool)`

GetErrorResponseHandlerOk returns a tuple with the ErrorResponseHandler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorResponseHandler

`func (o *WorkflowBatchApiExecutorAllOf) SetErrorResponseHandler(v WorkflowErrorResponseHandlerRelationship)`

SetErrorResponseHandler sets ErrorResponseHandler field to given value.

### HasErrorResponseHandler

`func (o *WorkflowBatchApiExecutorAllOf) HasErrorResponseHandler() bool`

HasErrorResponseHandler returns a boolean if a field has been set.

### GetTaskDefinition

`func (o *WorkflowBatchApiExecutorAllOf) GetTaskDefinition() WorkflowTaskDefinitionRelationship`

GetTaskDefinition returns the TaskDefinition field if non-nil, zero value otherwise.

### GetTaskDefinitionOk

`func (o *WorkflowBatchApiExecutorAllOf) GetTaskDefinitionOk() (*WorkflowTaskDefinitionRelationship, bool)`

GetTaskDefinitionOk returns a tuple with the TaskDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskDefinition

`func (o *WorkflowBatchApiExecutorAllOf) SetTaskDefinition(v WorkflowTaskDefinitionRelationship)`

SetTaskDefinition sets TaskDefinition field to given value.

### HasTaskDefinition

`func (o *WorkflowBatchApiExecutorAllOf) HasTaskDefinition() bool`

HasTaskDefinition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



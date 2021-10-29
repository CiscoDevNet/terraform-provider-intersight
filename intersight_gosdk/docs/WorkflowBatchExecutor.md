# WorkflowBatchExecutor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "workflow.BatchApiExecutor"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "workflow.BatchApiExecutor"]
**Batch** | Pointer to [**[]WorkflowApi**](WorkflowApi.md) |  | [optional] 
**Constraints** | Pointer to [**NullableWorkflowTaskConstraints**](WorkflowTaskConstraints.md) |  | [optional] 
**Description** | Pointer to **string** | A detailed description about the batch APIs. | [optional] 
**Name** | Pointer to **string** | Name for the batch API task. | [optional] 
**Outcomes** | Pointer to **interface{}** | All the possible outcomes of this task are captured here. Outcomes property is a collection property of type workflow.Outcome objects. The outcomes can be mapped to the message to be shown. The outcomes are evaluated in the order they are given. At the end of the outcomes list, an catchall success/fail outcome can be added with condition as &#39;true&#39;. This is an optional property and if not specified the task will be marked as success. | [optional] 
**Output** | Pointer to **interface{}** | Intersight Orchestrator allows the extraction of required values from API responses using the API response grammar. These extracted values can be mapped to task output parameters defined in task definition. The mapping of API output parameters to the task output parameters is provided as JSON in this property. | [optional] 
**RetryFromFailedApi** | Pointer to **bool** | When an execution of a nth API in the Batch fails, Retry from falied API flag indicates if the execution should start from the nth API or the first API during task retry. By default the value is set to false. | [optional] 
**SkipOnCondition** | Pointer to **string** | The skip expression, if provided, allows the batch API executor to skip the task execution when the given expression evaluates to true. The expression is given as such a golang template that has to be evaluated to a final content true/false. The expression is an optional and in case not provided, the API will always be executed. | [optional] 

## Methods

### NewWorkflowBatchExecutor

`func NewWorkflowBatchExecutor(classId string, objectType string, ) *WorkflowBatchExecutor`

NewWorkflowBatchExecutor instantiates a new WorkflowBatchExecutor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowBatchExecutorWithDefaults

`func NewWorkflowBatchExecutorWithDefaults() *WorkflowBatchExecutor`

NewWorkflowBatchExecutorWithDefaults instantiates a new WorkflowBatchExecutor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowBatchExecutor) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowBatchExecutor) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowBatchExecutor) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowBatchExecutor) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowBatchExecutor) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowBatchExecutor) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBatch

`func (o *WorkflowBatchExecutor) GetBatch() []WorkflowApi`

GetBatch returns the Batch field if non-nil, zero value otherwise.

### GetBatchOk

`func (o *WorkflowBatchExecutor) GetBatchOk() (*[]WorkflowApi, bool)`

GetBatchOk returns a tuple with the Batch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatch

`func (o *WorkflowBatchExecutor) SetBatch(v []WorkflowApi)`

SetBatch sets Batch field to given value.

### HasBatch

`func (o *WorkflowBatchExecutor) HasBatch() bool`

HasBatch returns a boolean if a field has been set.

### SetBatchNil

`func (o *WorkflowBatchExecutor) SetBatchNil(b bool)`

 SetBatchNil sets the value for Batch to be an explicit nil

### UnsetBatch
`func (o *WorkflowBatchExecutor) UnsetBatch()`

UnsetBatch ensures that no value is present for Batch, not even an explicit nil
### GetConstraints

`func (o *WorkflowBatchExecutor) GetConstraints() WorkflowTaskConstraints`

GetConstraints returns the Constraints field if non-nil, zero value otherwise.

### GetConstraintsOk

`func (o *WorkflowBatchExecutor) GetConstraintsOk() (*WorkflowTaskConstraints, bool)`

GetConstraintsOk returns a tuple with the Constraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraints

`func (o *WorkflowBatchExecutor) SetConstraints(v WorkflowTaskConstraints)`

SetConstraints sets Constraints field to given value.

### HasConstraints

`func (o *WorkflowBatchExecutor) HasConstraints() bool`

HasConstraints returns a boolean if a field has been set.

### SetConstraintsNil

`func (o *WorkflowBatchExecutor) SetConstraintsNil(b bool)`

 SetConstraintsNil sets the value for Constraints to be an explicit nil

### UnsetConstraints
`func (o *WorkflowBatchExecutor) UnsetConstraints()`

UnsetConstraints ensures that no value is present for Constraints, not even an explicit nil
### GetDescription

`func (o *WorkflowBatchExecutor) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkflowBatchExecutor) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkflowBatchExecutor) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkflowBatchExecutor) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *WorkflowBatchExecutor) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowBatchExecutor) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowBatchExecutor) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowBatchExecutor) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOutcomes

`func (o *WorkflowBatchExecutor) GetOutcomes() interface{}`

GetOutcomes returns the Outcomes field if non-nil, zero value otherwise.

### GetOutcomesOk

`func (o *WorkflowBatchExecutor) GetOutcomesOk() (*interface{}, bool)`

GetOutcomesOk returns a tuple with the Outcomes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcomes

`func (o *WorkflowBatchExecutor) SetOutcomes(v interface{})`

SetOutcomes sets Outcomes field to given value.

### HasOutcomes

`func (o *WorkflowBatchExecutor) HasOutcomes() bool`

HasOutcomes returns a boolean if a field has been set.

### SetOutcomesNil

`func (o *WorkflowBatchExecutor) SetOutcomesNil(b bool)`

 SetOutcomesNil sets the value for Outcomes to be an explicit nil

### UnsetOutcomes
`func (o *WorkflowBatchExecutor) UnsetOutcomes()`

UnsetOutcomes ensures that no value is present for Outcomes, not even an explicit nil
### GetOutput

`func (o *WorkflowBatchExecutor) GetOutput() interface{}`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *WorkflowBatchExecutor) GetOutputOk() (*interface{}, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *WorkflowBatchExecutor) SetOutput(v interface{})`

SetOutput sets Output field to given value.

### HasOutput

`func (o *WorkflowBatchExecutor) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### SetOutputNil

`func (o *WorkflowBatchExecutor) SetOutputNil(b bool)`

 SetOutputNil sets the value for Output to be an explicit nil

### UnsetOutput
`func (o *WorkflowBatchExecutor) UnsetOutput()`

UnsetOutput ensures that no value is present for Output, not even an explicit nil
### GetRetryFromFailedApi

`func (o *WorkflowBatchExecutor) GetRetryFromFailedApi() bool`

GetRetryFromFailedApi returns the RetryFromFailedApi field if non-nil, zero value otherwise.

### GetRetryFromFailedApiOk

`func (o *WorkflowBatchExecutor) GetRetryFromFailedApiOk() (*bool, bool)`

GetRetryFromFailedApiOk returns a tuple with the RetryFromFailedApi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryFromFailedApi

`func (o *WorkflowBatchExecutor) SetRetryFromFailedApi(v bool)`

SetRetryFromFailedApi sets RetryFromFailedApi field to given value.

### HasRetryFromFailedApi

`func (o *WorkflowBatchExecutor) HasRetryFromFailedApi() bool`

HasRetryFromFailedApi returns a boolean if a field has been set.

### GetSkipOnCondition

`func (o *WorkflowBatchExecutor) GetSkipOnCondition() string`

GetSkipOnCondition returns the SkipOnCondition field if non-nil, zero value otherwise.

### GetSkipOnConditionOk

`func (o *WorkflowBatchExecutor) GetSkipOnConditionOk() (*string, bool)`

GetSkipOnConditionOk returns a tuple with the SkipOnCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipOnCondition

`func (o *WorkflowBatchExecutor) SetSkipOnCondition(v string)`

SetSkipOnCondition sets SkipOnCondition field to given value.

### HasSkipOnCondition

`func (o *WorkflowBatchExecutor) HasSkipOnCondition() bool`

HasSkipOnCondition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



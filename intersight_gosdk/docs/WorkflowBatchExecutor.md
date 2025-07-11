# WorkflowBatchExecutor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**Batch** | Pointer to [**[]WorkflowApi**](WorkflowApi.md) |  | [optional] 
**CancelAction** | Pointer to [**[]WorkflowApi**](WorkflowApi.md) |  | [optional] 
**Constraints** | Pointer to [**NullableWorkflowTaskConstraints**](WorkflowTaskConstraints.md) |  | [optional] 
**Description** | Pointer to **string** | Detailed description of the batch APIs. | [optional] 
**Name** | Pointer to **string** | Name of the batch API task. | [optional] 
**Outcomes** | Pointer to **interface{}** | Collection of possible task outcomes, represented as workflow.Outcome objects. Outcomes can be mapped to messages and are evaluated in the given order.  A catch-all success or failure outcome with condition &#39;true&#39; can be included at the end. Optional property; if not specified, the task defaults to success. | [optional] 
**Output** | Pointer to **interface{}** | JSON mapping of extracted API response values to task output parameters, using API response grammar defined in Intersight Orchestrator. | [optional] 
**RetryFromFailedApi** | Pointer to **bool** | Flag indicating if the retry task should from the failed API or the first API in the batch execution; default value is false. | [optional] 
**SkipOnCondition** | Pointer to **string** | Optional skip expression allowing the batch API executor to skip task execution when the provided Go template expression evaluates to true.  If not specified, the API will always be executed. | [optional] 
**UiRenderingData** | Pointer to **interface{}** | Data required for rendering the task in the user interface. | [optional] 

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
### GetCancelAction

`func (o *WorkflowBatchExecutor) GetCancelAction() []WorkflowApi`

GetCancelAction returns the CancelAction field if non-nil, zero value otherwise.

### GetCancelActionOk

`func (o *WorkflowBatchExecutor) GetCancelActionOk() (*[]WorkflowApi, bool)`

GetCancelActionOk returns a tuple with the CancelAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelAction

`func (o *WorkflowBatchExecutor) SetCancelAction(v []WorkflowApi)`

SetCancelAction sets CancelAction field to given value.

### HasCancelAction

`func (o *WorkflowBatchExecutor) HasCancelAction() bool`

HasCancelAction returns a boolean if a field has been set.

### SetCancelActionNil

`func (o *WorkflowBatchExecutor) SetCancelActionNil(b bool)`

 SetCancelActionNil sets the value for CancelAction to be an explicit nil

### UnsetCancelAction
`func (o *WorkflowBatchExecutor) UnsetCancelAction()`

UnsetCancelAction ensures that no value is present for CancelAction, not even an explicit nil
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

### GetUiRenderingData

`func (o *WorkflowBatchExecutor) GetUiRenderingData() interface{}`

GetUiRenderingData returns the UiRenderingData field if non-nil, zero value otherwise.

### GetUiRenderingDataOk

`func (o *WorkflowBatchExecutor) GetUiRenderingDataOk() (*interface{}, bool)`

GetUiRenderingDataOk returns a tuple with the UiRenderingData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiRenderingData

`func (o *WorkflowBatchExecutor) SetUiRenderingData(v interface{})`

SetUiRenderingData sets UiRenderingData field to given value.

### HasUiRenderingData

`func (o *WorkflowBatchExecutor) HasUiRenderingData() bool`

HasUiRenderingData returns a boolean if a field has been set.

### SetUiRenderingDataNil

`func (o *WorkflowBatchExecutor) SetUiRenderingDataNil(b bool)`

 SetUiRenderingDataNil sets the value for UiRenderingData to be an explicit nil

### UnsetUiRenderingData
`func (o *WorkflowBatchExecutor) UnsetUiRenderingData()`

UnsetUiRenderingData ensures that no value is present for UiRenderingData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



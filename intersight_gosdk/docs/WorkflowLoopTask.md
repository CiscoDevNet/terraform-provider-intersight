# WorkflowLoopTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.LoopTask"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.LoopTask"]
**FailurePolicy** | Pointer to **string** | The policy to handle the failure of an iteration within a parallel loop. * &#x60;FailOnFirstFailure&#x60; - The enum specifies the option as FailOnFirstFailure where the loop task will fail if one of the iteration in the loop fails. The running iterations will be cancelled on first failure and the loop will be marked as failed. * &#x60;ContinueOnFailure&#x60; - The enum specifies the option as ContinueOnFailure where the loop task will continue with all iterations, even if one fails. Running iterations will not be canceled, but the loop will be marked as failed after all iterations are complete. | [optional] [default to "FailOnFirstFailure"]
**NumberOfBatches** | Pointer to **int64** | All iterations of the loop run in parallel within a single batch, with a maximum of 100 iterations. To run more than 100 iterations, you can increase the number of batches. The configuration is acceptable as long as the total number of iterations divided by the number of batches is less than 100. Adjusting the number of batches also allows you to control how many iterations run in parallel. For example, if the total count is 100 and you set the number of batches to 5, then 20 tasks will run in parallel across the 5 batches. It&#39;s important to note that the number of batches must be less than the total count. If the count is dynamic and falls below the number of batches, this may result in empty batches with no tasks. | [optional] [default to 1]
**Parallel** | Pointer to **bool** | This field is deprecated. Always set to true for parallel loop. | [optional] [default to true]

## Methods

### NewWorkflowLoopTask

`func NewWorkflowLoopTask(classId string, objectType string, ) *WorkflowLoopTask`

NewWorkflowLoopTask instantiates a new WorkflowLoopTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowLoopTaskWithDefaults

`func NewWorkflowLoopTaskWithDefaults() *WorkflowLoopTask`

NewWorkflowLoopTaskWithDefaults instantiates a new WorkflowLoopTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowLoopTask) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowLoopTask) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowLoopTask) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowLoopTask) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowLoopTask) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowLoopTask) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetFailurePolicy

`func (o *WorkflowLoopTask) GetFailurePolicy() string`

GetFailurePolicy returns the FailurePolicy field if non-nil, zero value otherwise.

### GetFailurePolicyOk

`func (o *WorkflowLoopTask) GetFailurePolicyOk() (*string, bool)`

GetFailurePolicyOk returns a tuple with the FailurePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailurePolicy

`func (o *WorkflowLoopTask) SetFailurePolicy(v string)`

SetFailurePolicy sets FailurePolicy field to given value.

### HasFailurePolicy

`func (o *WorkflowLoopTask) HasFailurePolicy() bool`

HasFailurePolicy returns a boolean if a field has been set.

### GetNumberOfBatches

`func (o *WorkflowLoopTask) GetNumberOfBatches() int64`

GetNumberOfBatches returns the NumberOfBatches field if non-nil, zero value otherwise.

### GetNumberOfBatchesOk

`func (o *WorkflowLoopTask) GetNumberOfBatchesOk() (*int64, bool)`

GetNumberOfBatchesOk returns a tuple with the NumberOfBatches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfBatches

`func (o *WorkflowLoopTask) SetNumberOfBatches(v int64)`

SetNumberOfBatches sets NumberOfBatches field to given value.

### HasNumberOfBatches

`func (o *WorkflowLoopTask) HasNumberOfBatches() bool`

HasNumberOfBatches returns a boolean if a field has been set.

### GetParallel

`func (o *WorkflowLoopTask) GetParallel() bool`

GetParallel returns the Parallel field if non-nil, zero value otherwise.

### GetParallelOk

`func (o *WorkflowLoopTask) GetParallelOk() (*bool, bool)`

GetParallelOk returns a tuple with the Parallel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParallel

`func (o *WorkflowLoopTask) SetParallel(v bool)`

SetParallel sets Parallel field to given value.

### HasParallel

`func (o *WorkflowLoopTask) HasParallel() bool`

HasParallel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



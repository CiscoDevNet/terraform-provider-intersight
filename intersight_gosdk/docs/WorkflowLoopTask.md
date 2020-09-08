# WorkflowLoopTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **string** | Count value for the loop, this can be a constant or an expression which will evaluate to an integer value. Example, Use the length of the input A which is an array. Count must be less than or equal to 500. | [optional] 
**LoopStartTask** | Pointer to **string** | Start task where the list of tasks will be executed multiple times based on the count value. | [optional] 
**NumberOfBatches** | Pointer to **int64** | When tasks are run in parallel and the count is large, the actual number of task run in parallel can be controlled by this property. If count is 100 and numberOfBatches is 5 then 20 tasks are run in parallel 5 times. Parallel batch size must be less than the count. In cases where count is dynamic and depends on input given during workflow execution, if that count is less than batch then empty batches might get created which do not have any tasks under them. | [optional] 
**OnSuccess** | Pointer to **string** | This specifies the name of the next task to run if all iterations of the loop task succeeds.  This is the unique name given to the task instance within the workflow. In a graph model, denotes an edge to another Task Node. | [optional] 
**Parallel** | Pointer to **bool** | When set to true the loop will run in parallel else it will run in a serial fashion. Only one task is supported inside the loop task when the loop is run in parallel. Subworkflow can be used inside the single loop task to build complex conditions. | [optional] 

## Methods

### NewWorkflowLoopTask

`func NewWorkflowLoopTask() *WorkflowLoopTask`

NewWorkflowLoopTask instantiates a new WorkflowLoopTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowLoopTaskWithDefaults

`func NewWorkflowLoopTaskWithDefaults() *WorkflowLoopTask`

NewWorkflowLoopTaskWithDefaults instantiates a new WorkflowLoopTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *WorkflowLoopTask) GetCount() string`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *WorkflowLoopTask) GetCountOk() (*string, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *WorkflowLoopTask) SetCount(v string)`

SetCount sets Count field to given value.

### HasCount

`func (o *WorkflowLoopTask) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetLoopStartTask

`func (o *WorkflowLoopTask) GetLoopStartTask() string`

GetLoopStartTask returns the LoopStartTask field if non-nil, zero value otherwise.

### GetLoopStartTaskOk

`func (o *WorkflowLoopTask) GetLoopStartTaskOk() (*string, bool)`

GetLoopStartTaskOk returns a tuple with the LoopStartTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoopStartTask

`func (o *WorkflowLoopTask) SetLoopStartTask(v string)`

SetLoopStartTask sets LoopStartTask field to given value.

### HasLoopStartTask

`func (o *WorkflowLoopTask) HasLoopStartTask() bool`

HasLoopStartTask returns a boolean if a field has been set.

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

### GetOnSuccess

`func (o *WorkflowLoopTask) GetOnSuccess() string`

GetOnSuccess returns the OnSuccess field if non-nil, zero value otherwise.

### GetOnSuccessOk

`func (o *WorkflowLoopTask) GetOnSuccessOk() (*string, bool)`

GetOnSuccessOk returns a tuple with the OnSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnSuccess

`func (o *WorkflowLoopTask) SetOnSuccess(v string)`

SetOnSuccess sets OnSuccess field to given value.

### HasOnSuccess

`func (o *WorkflowLoopTask) HasOnSuccess() bool`

HasOnSuccess returns a boolean if a field has been set.

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



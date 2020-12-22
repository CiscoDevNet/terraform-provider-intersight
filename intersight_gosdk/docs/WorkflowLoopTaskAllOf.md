# WorkflowLoopTaskAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.LoopTask"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.LoopTask"]
**Count** | Pointer to **string** | Count value for the loop, this can be a constant or an expression which will evaluate to an integer value. Example, Use the length of the input A which is an array. Count must be less than or equal to 500. | [optional] 
**LoopStartTask** | Pointer to **string** | Start task where the list of tasks will be executed multiple times based on the count value. | [optional] 
**NumberOfBatches** | Pointer to **int64** | When tasks are run in parallel and the count is large, the actual number of task run in parallel can be controlled by this property. If count is 100 and numberOfBatches is 5 then 20 tasks are run in parallel 5 times. Parallel batch size must be less than the count. In cases where count is dynamic and depends on input given during workflow execution, if that count is less than batch then empty batches might get created which do not have any tasks under them. | [optional] [default to 1]
**OnSuccess** | Pointer to **string** | This specifies the name of the next task to run if all iterations of the loop task succeeds. The unique name given to the task instance within the workflow must be provided here. In a graph model, denotes an edge to another Task Node. | [optional] 
**Parallel** | Pointer to **bool** | When set to true the loop will run in parallel else it will run in a serial fashion. Only one task is supported inside the loop task when the loop is run in parallel. Subworkflow can be used inside the single loop task to build complex conditions. | [optional] [default to true]

## Methods

### NewWorkflowLoopTaskAllOf

`func NewWorkflowLoopTaskAllOf(classId string, objectType string, ) *WorkflowLoopTaskAllOf`

NewWorkflowLoopTaskAllOf instantiates a new WorkflowLoopTaskAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowLoopTaskAllOfWithDefaults

`func NewWorkflowLoopTaskAllOfWithDefaults() *WorkflowLoopTaskAllOf`

NewWorkflowLoopTaskAllOfWithDefaults instantiates a new WorkflowLoopTaskAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowLoopTaskAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowLoopTaskAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowLoopTaskAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowLoopTaskAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowLoopTaskAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowLoopTaskAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCount

`func (o *WorkflowLoopTaskAllOf) GetCount() string`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *WorkflowLoopTaskAllOf) GetCountOk() (*string, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *WorkflowLoopTaskAllOf) SetCount(v string)`

SetCount sets Count field to given value.

### HasCount

`func (o *WorkflowLoopTaskAllOf) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetLoopStartTask

`func (o *WorkflowLoopTaskAllOf) GetLoopStartTask() string`

GetLoopStartTask returns the LoopStartTask field if non-nil, zero value otherwise.

### GetLoopStartTaskOk

`func (o *WorkflowLoopTaskAllOf) GetLoopStartTaskOk() (*string, bool)`

GetLoopStartTaskOk returns a tuple with the LoopStartTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoopStartTask

`func (o *WorkflowLoopTaskAllOf) SetLoopStartTask(v string)`

SetLoopStartTask sets LoopStartTask field to given value.

### HasLoopStartTask

`func (o *WorkflowLoopTaskAllOf) HasLoopStartTask() bool`

HasLoopStartTask returns a boolean if a field has been set.

### GetNumberOfBatches

`func (o *WorkflowLoopTaskAllOf) GetNumberOfBatches() int64`

GetNumberOfBatches returns the NumberOfBatches field if non-nil, zero value otherwise.

### GetNumberOfBatchesOk

`func (o *WorkflowLoopTaskAllOf) GetNumberOfBatchesOk() (*int64, bool)`

GetNumberOfBatchesOk returns a tuple with the NumberOfBatches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfBatches

`func (o *WorkflowLoopTaskAllOf) SetNumberOfBatches(v int64)`

SetNumberOfBatches sets NumberOfBatches field to given value.

### HasNumberOfBatches

`func (o *WorkflowLoopTaskAllOf) HasNumberOfBatches() bool`

HasNumberOfBatches returns a boolean if a field has been set.

### GetOnSuccess

`func (o *WorkflowLoopTaskAllOf) GetOnSuccess() string`

GetOnSuccess returns the OnSuccess field if non-nil, zero value otherwise.

### GetOnSuccessOk

`func (o *WorkflowLoopTaskAllOf) GetOnSuccessOk() (*string, bool)`

GetOnSuccessOk returns a tuple with the OnSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnSuccess

`func (o *WorkflowLoopTaskAllOf) SetOnSuccess(v string)`

SetOnSuccess sets OnSuccess field to given value.

### HasOnSuccess

`func (o *WorkflowLoopTaskAllOf) HasOnSuccess() bool`

HasOnSuccess returns a boolean if a field has been set.

### GetParallel

`func (o *WorkflowLoopTaskAllOf) GetParallel() bool`

GetParallel returns the Parallel field if non-nil, zero value otherwise.

### GetParallelOk

`func (o *WorkflowLoopTaskAllOf) GetParallelOk() (*bool, bool)`

GetParallelOk returns a tuple with the Parallel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParallel

`func (o *WorkflowLoopTaskAllOf) SetParallel(v bool)`

SetParallel sets Parallel field to given value.

### HasParallel

`func (o *WorkflowLoopTaskAllOf) HasParallel() bool`

HasParallel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



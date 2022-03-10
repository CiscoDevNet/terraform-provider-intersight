# WorkflowLoopTaskAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.LoopTask"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.LoopTask"]
**NumberOfBatches** | Pointer to **int64** | When tasks are run in parallel and the count is large, the actual number of task run in parallel can be controlled by this property. If count is 100 and numberOfBatches is 5 then 20 tasks are run in parallel 5 times. Parallel batch size must be less than the count. In cases where count is dynamic and depends on input given during workflow execution, if that count is less than batch then empty batches might get created which do not have any tasks under them. | [optional] [default to 1]
**Parallel** | Pointer to **bool** | This field is deprecated. Always set to true for parallel loop. | [optional] [default to true]

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



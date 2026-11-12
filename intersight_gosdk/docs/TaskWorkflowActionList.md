# TaskWorkflowActionList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | The total number of &#39;task.WorkflowAction&#39; resources matching the request, accross all pages. The &#39;Count&#39; attribute is included when the HTTP GET request includes the &#39;$inlinecount&#39; parameter. | [optional] 
**Results** | Pointer to [**[]TaskWorkflowAction**](TaskWorkflowAction.md) | The array of &#39;task.WorkflowAction&#39; resources matching the request. | [optional] 

## Methods

### NewTaskWorkflowActionList

`func NewTaskWorkflowActionList() *TaskWorkflowActionList`

NewTaskWorkflowActionList instantiates a new TaskWorkflowActionList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskWorkflowActionListWithDefaults

`func NewTaskWorkflowActionListWithDefaults() *TaskWorkflowActionList`

NewTaskWorkflowActionListWithDefaults instantiates a new TaskWorkflowActionList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *TaskWorkflowActionList) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *TaskWorkflowActionList) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *TaskWorkflowActionList) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *TaskWorkflowActionList) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetResults

`func (o *TaskWorkflowActionList) GetResults() []TaskWorkflowAction`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *TaskWorkflowActionList) GetResultsOk() (*[]TaskWorkflowAction, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *TaskWorkflowActionList) SetResults(v []TaskWorkflowAction)`

SetResults sets Results field to given value.

### HasResults

`func (o *TaskWorkflowActionList) HasResults() bool`

HasResults returns a boolean if a field has been set.

### SetResultsNil

`func (o *TaskWorkflowActionList) SetResultsNil(b bool)`

 SetResultsNil sets the value for Results to be an explicit nil

### UnsetResults
`func (o *TaskWorkflowActionList) UnsetResults()`

UnsetResults ensures that no value is present for Results, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



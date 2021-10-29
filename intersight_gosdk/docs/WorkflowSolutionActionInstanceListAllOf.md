# WorkflowSolutionActionInstanceListAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | The total number of &#39;workflow.SolutionActionInstance&#39; resources matching the request, accross all pages. The &#39;Count&#39; attribute is included when the HTTP GET request includes the &#39;$inlinecount&#39; parameter. | [optional] 
**Results** | Pointer to [**[]WorkflowSolutionActionInstance**](WorkflowSolutionActionInstance.md) | The array of &#39;workflow.SolutionActionInstance&#39; resources matching the request. | [optional] 

## Methods

### NewWorkflowSolutionActionInstanceListAllOf

`func NewWorkflowSolutionActionInstanceListAllOf() *WorkflowSolutionActionInstanceListAllOf`

NewWorkflowSolutionActionInstanceListAllOf instantiates a new WorkflowSolutionActionInstanceListAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowSolutionActionInstanceListAllOfWithDefaults

`func NewWorkflowSolutionActionInstanceListAllOfWithDefaults() *WorkflowSolutionActionInstanceListAllOf`

NewWorkflowSolutionActionInstanceListAllOfWithDefaults instantiates a new WorkflowSolutionActionInstanceListAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *WorkflowSolutionActionInstanceListAllOf) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *WorkflowSolutionActionInstanceListAllOf) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *WorkflowSolutionActionInstanceListAllOf) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *WorkflowSolutionActionInstanceListAllOf) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetResults

`func (o *WorkflowSolutionActionInstanceListAllOf) GetResults() []WorkflowSolutionActionInstance`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *WorkflowSolutionActionInstanceListAllOf) GetResultsOk() (*[]WorkflowSolutionActionInstance, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *WorkflowSolutionActionInstanceListAllOf) SetResults(v []WorkflowSolutionActionInstance)`

SetResults sets Results field to given value.

### HasResults

`func (o *WorkflowSolutionActionInstanceListAllOf) HasResults() bool`

HasResults returns a boolean if a field has been set.

### SetResultsNil

`func (o *WorkflowSolutionActionInstanceListAllOf) SetResultsNil(b bool)`

 SetResultsNil sets the value for Results to be an explicit nil

### UnsetResults
`func (o *WorkflowSolutionActionInstanceListAllOf) UnsetResults()`

UnsetResults ensures that no value is present for Results, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



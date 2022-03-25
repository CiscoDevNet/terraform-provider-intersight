# WorkflowServiceItemDefinitionList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | The total number of &#39;workflow.ServiceItemDefinition&#39; resources matching the request, accross all pages. The &#39;Count&#39; attribute is included when the HTTP GET request includes the &#39;$inlinecount&#39; parameter. | [optional] 
**Results** | Pointer to [**[]WorkflowServiceItemDefinition**](WorkflowServiceItemDefinition.md) | The array of &#39;workflow.ServiceItemDefinition&#39; resources matching the request. | [optional] 

## Methods

### NewWorkflowServiceItemDefinitionList

`func NewWorkflowServiceItemDefinitionList() *WorkflowServiceItemDefinitionList`

NewWorkflowServiceItemDefinitionList instantiates a new WorkflowServiceItemDefinitionList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowServiceItemDefinitionListWithDefaults

`func NewWorkflowServiceItemDefinitionListWithDefaults() *WorkflowServiceItemDefinitionList`

NewWorkflowServiceItemDefinitionListWithDefaults instantiates a new WorkflowServiceItemDefinitionList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *WorkflowServiceItemDefinitionList) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *WorkflowServiceItemDefinitionList) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *WorkflowServiceItemDefinitionList) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *WorkflowServiceItemDefinitionList) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetResults

`func (o *WorkflowServiceItemDefinitionList) GetResults() []WorkflowServiceItemDefinition`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *WorkflowServiceItemDefinitionList) GetResultsOk() (*[]WorkflowServiceItemDefinition, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *WorkflowServiceItemDefinitionList) SetResults(v []WorkflowServiceItemDefinition)`

SetResults sets Results field to given value.

### HasResults

`func (o *WorkflowServiceItemDefinitionList) HasResults() bool`

HasResults returns a boolean if a field has been set.

### SetResultsNil

`func (o *WorkflowServiceItemDefinitionList) SetResultsNil(b bool)`

 SetResultsNil sets the value for Results to be an explicit nil

### UnsetResults
`func (o *WorkflowServiceItemDefinitionList) UnsetResults()`

UnsetResults ensures that no value is present for Results, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



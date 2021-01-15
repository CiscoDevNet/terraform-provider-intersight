# HyperflexVmRestoreOperationListAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | The total number of &#39;hyperflex.VmRestoreOperation&#39; resources matching the request, accross all pages. The &#39;Count&#39; attribute is included when the HTTP GET request includes the &#39;$inlinecount&#39; parameter. | [optional] 
**Results** | Pointer to [**[]HyperflexVmRestoreOperation**](HyperflexVmRestoreOperation.md) | The array of &#39;hyperflex.VmRestoreOperation&#39; resources matching the request. | [optional] 

## Methods

### NewHyperflexVmRestoreOperationListAllOf

`func NewHyperflexVmRestoreOperationListAllOf() *HyperflexVmRestoreOperationListAllOf`

NewHyperflexVmRestoreOperationListAllOf instantiates a new HyperflexVmRestoreOperationListAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexVmRestoreOperationListAllOfWithDefaults

`func NewHyperflexVmRestoreOperationListAllOfWithDefaults() *HyperflexVmRestoreOperationListAllOf`

NewHyperflexVmRestoreOperationListAllOfWithDefaults instantiates a new HyperflexVmRestoreOperationListAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *HyperflexVmRestoreOperationListAllOf) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *HyperflexVmRestoreOperationListAllOf) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *HyperflexVmRestoreOperationListAllOf) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *HyperflexVmRestoreOperationListAllOf) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetResults

`func (o *HyperflexVmRestoreOperationListAllOf) GetResults() []HyperflexVmRestoreOperation`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *HyperflexVmRestoreOperationListAllOf) GetResultsOk() (*[]HyperflexVmRestoreOperation, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *HyperflexVmRestoreOperationListAllOf) SetResults(v []HyperflexVmRestoreOperation)`

SetResults sets Results field to given value.

### HasResults

`func (o *HyperflexVmRestoreOperationListAllOf) HasResults() bool`

HasResults returns a boolean if a field has been set.

### SetResultsNil

`func (o *HyperflexVmRestoreOperationListAllOf) SetResultsNil(b bool)`

 SetResultsNil sets the value for Results to be an explicit nil

### UnsetResults
`func (o *HyperflexVmRestoreOperationListAllOf) UnsetResults()`

UnsetResults ensures that no value is present for Results, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



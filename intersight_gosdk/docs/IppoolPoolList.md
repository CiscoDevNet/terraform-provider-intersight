# IppoolPoolList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | The total number of &#39;ippool.Pool&#39; resources matching the request, accross all pages. The &#39;Count&#39; attribute is included when the HTTP GET request includes the &#39;$inlinecount&#39; parameter. | [optional] 
**Results** | Pointer to [**[]IppoolPool**](ippool.Pool.md) | The array of &#39;ippool.Pool&#39; resources matching the request. | [optional] 

## Methods

### NewIppoolPoolList

`func NewIppoolPoolList() *IppoolPoolList`

NewIppoolPoolList instantiates a new IppoolPoolList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIppoolPoolListWithDefaults

`func NewIppoolPoolListWithDefaults() *IppoolPoolList`

NewIppoolPoolListWithDefaults instantiates a new IppoolPoolList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *IppoolPoolList) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *IppoolPoolList) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *IppoolPoolList) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *IppoolPoolList) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetResults

`func (o *IppoolPoolList) GetResults() []IppoolPool`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *IppoolPoolList) GetResultsOk() (*[]IppoolPool, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *IppoolPoolList) SetResults(v []IppoolPool)`

SetResults sets Results field to given value.

### HasResults

`func (o *IppoolPoolList) HasResults() bool`

HasResults returns a boolean if a field has been set.

### SetResultsNil

`func (o *IppoolPoolList) SetResultsNil(b bool)`

 SetResultsNil sets the value for Results to be an explicit nil

### UnsetResults
`func (o *IppoolPoolList) UnsetResults()`

UnsetResults ensures that no value is present for Results, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



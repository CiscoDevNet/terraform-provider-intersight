# KubernetesConfigResultEntryList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | The total number of &#39;kubernetes.ConfigResultEntry&#39; resources matching the request, accross all pages. The &#39;Count&#39; attribute is included when the HTTP GET request includes the &#39;$inlinecount&#39; parameter. | [optional] 
**Results** | Pointer to [**[]KubernetesConfigResultEntry**](KubernetesConfigResultEntry.md) | The array of &#39;kubernetes.ConfigResultEntry&#39; resources matching the request. | [optional] 

## Methods

### NewKubernetesConfigResultEntryList

`func NewKubernetesConfigResultEntryList() *KubernetesConfigResultEntryList`

NewKubernetesConfigResultEntryList instantiates a new KubernetesConfigResultEntryList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesConfigResultEntryListWithDefaults

`func NewKubernetesConfigResultEntryListWithDefaults() *KubernetesConfigResultEntryList`

NewKubernetesConfigResultEntryListWithDefaults instantiates a new KubernetesConfigResultEntryList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *KubernetesConfigResultEntryList) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *KubernetesConfigResultEntryList) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *KubernetesConfigResultEntryList) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *KubernetesConfigResultEntryList) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetResults

`func (o *KubernetesConfigResultEntryList) GetResults() []KubernetesConfigResultEntry`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *KubernetesConfigResultEntryList) GetResultsOk() (*[]KubernetesConfigResultEntry, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *KubernetesConfigResultEntryList) SetResults(v []KubernetesConfigResultEntry)`

SetResults sets Results field to given value.

### HasResults

`func (o *KubernetesConfigResultEntryList) HasResults() bool`

HasResults returns a boolean if a field has been set.

### SetResultsNil

`func (o *KubernetesConfigResultEntryList) SetResultsNil(b bool)`

 SetResultsNil sets the value for Results to be an explicit nil

### UnsetResults
`func (o *KubernetesConfigResultEntryList) UnsetResults()`

UnsetResults ensures that no value is present for Results, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



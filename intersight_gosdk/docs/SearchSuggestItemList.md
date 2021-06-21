# SearchSuggestItemList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | The total number of &#39;search.SuggestItem&#39; resources matching the request, accross all pages. The &#39;Count&#39; attribute is included when the HTTP GET request includes the &#39;$inlinecount&#39; parameter. | [optional] 
**Results** | Pointer to **[]map[string]interface{}** | The array of &#39;search.SuggestItem&#39; resources matching the request. | [optional] 

## Methods

### NewSearchSuggestItemList

`func NewSearchSuggestItemList() *SearchSuggestItemList`

NewSearchSuggestItemList instantiates a new SearchSuggestItemList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchSuggestItemListWithDefaults

`func NewSearchSuggestItemListWithDefaults() *SearchSuggestItemList`

NewSearchSuggestItemListWithDefaults instantiates a new SearchSuggestItemList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *SearchSuggestItemList) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *SearchSuggestItemList) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *SearchSuggestItemList) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *SearchSuggestItemList) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetResults

`func (o *SearchSuggestItemList) GetResults() []map[string]interface{}`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *SearchSuggestItemList) GetResultsOk() (*[]map[string]interface{}, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *SearchSuggestItemList) SetResults(v []map[string]interface{})`

SetResults sets Results field to given value.

### HasResults

`func (o *SearchSuggestItemList) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# SearchTagItemListAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | The total number of &#39;search.TagItem&#39; resources matching the request, accross all pages. The &#39;Count&#39; attribute is included when the HTTP GET request includes the &#39;$inlinecount&#39; parameter. | [optional] 
**Results** | Pointer to [**[]SearchTagItem**](search.TagItem.md) | The array of &#39;search.TagItem&#39; resources matching the request. | [optional] 

## Methods

### NewSearchTagItemListAllOf

`func NewSearchTagItemListAllOf() *SearchTagItemListAllOf`

NewSearchTagItemListAllOf instantiates a new SearchTagItemListAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchTagItemListAllOfWithDefaults

`func NewSearchTagItemListAllOfWithDefaults() *SearchTagItemListAllOf`

NewSearchTagItemListAllOfWithDefaults instantiates a new SearchTagItemListAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *SearchTagItemListAllOf) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *SearchTagItemListAllOf) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *SearchTagItemListAllOf) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *SearchTagItemListAllOf) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetResults

`func (o *SearchTagItemListAllOf) GetResults() []SearchTagItem`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *SearchTagItemListAllOf) GetResultsOk() (*[]SearchTagItem, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *SearchTagItemListAllOf) SetResults(v []SearchTagItem)`

SetResults sets Results field to given value.

### HasResults

`func (o *SearchTagItemListAllOf) HasResults() bool`

HasResults returns a boolean if a field has been set.

### SetResultsNil

`func (o *SearchTagItemListAllOf) SetResultsNil(b bool)`

 SetResultsNil sets the value for Results to be an explicit nil

### UnsetResults
`func (o *SearchTagItemListAllOf) UnsetResults()`

UnsetResults ensures that no value is present for Results, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



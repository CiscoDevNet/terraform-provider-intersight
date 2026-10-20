# DnacTemplateList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | The total number of &#39;dnac.Template&#39; resources matching the request, accross all pages. The &#39;Count&#39; attribute is included when the HTTP GET request includes the &#39;$inlinecount&#39; parameter. | [optional] 
**Results** | Pointer to [**[]DnacTemplate**](DnacTemplate.md) | The array of &#39;dnac.Template&#39; resources matching the request. | [optional] 

## Methods

### NewDnacTemplateList

`func NewDnacTemplateList() *DnacTemplateList`

NewDnacTemplateList instantiates a new DnacTemplateList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnacTemplateListWithDefaults

`func NewDnacTemplateListWithDefaults() *DnacTemplateList`

NewDnacTemplateListWithDefaults instantiates a new DnacTemplateList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *DnacTemplateList) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *DnacTemplateList) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *DnacTemplateList) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *DnacTemplateList) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetResults

`func (o *DnacTemplateList) GetResults() []DnacTemplate`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *DnacTemplateList) GetResultsOk() (*[]DnacTemplate, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *DnacTemplateList) SetResults(v []DnacTemplate)`

SetResults sets Results field to given value.

### HasResults

`func (o *DnacTemplateList) HasResults() bool`

HasResults returns a boolean if a field has been set.

### SetResultsNil

`func (o *DnacTemplateList) SetResultsNil(b bool)`

 SetResultsNil sets the value for Results to be an explicit nil

### UnsetResults
`func (o *DnacTemplateList) UnsetResults()`

UnsetResults ensures that no value is present for Results, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



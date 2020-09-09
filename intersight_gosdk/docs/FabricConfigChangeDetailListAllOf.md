# FabricConfigChangeDetailListAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | The total number of &#39;fabric.ConfigChangeDetail&#39; resources matching the request, accross all pages. The &#39;Count&#39; attribute is included when the HTTP GET request includes the &#39;$inlinecount&#39; parameter. | [optional] 
**Results** | Pointer to [**[]FabricConfigChangeDetail**](fabric.ConfigChangeDetail.md) | The array of &#39;fabric.ConfigChangeDetail&#39; resources matching the request. | [optional] 

## Methods

### NewFabricConfigChangeDetailListAllOf

`func NewFabricConfigChangeDetailListAllOf() *FabricConfigChangeDetailListAllOf`

NewFabricConfigChangeDetailListAllOf instantiates a new FabricConfigChangeDetailListAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricConfigChangeDetailListAllOfWithDefaults

`func NewFabricConfigChangeDetailListAllOfWithDefaults() *FabricConfigChangeDetailListAllOf`

NewFabricConfigChangeDetailListAllOfWithDefaults instantiates a new FabricConfigChangeDetailListAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *FabricConfigChangeDetailListAllOf) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *FabricConfigChangeDetailListAllOf) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *FabricConfigChangeDetailListAllOf) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *FabricConfigChangeDetailListAllOf) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetResults

`func (o *FabricConfigChangeDetailListAllOf) GetResults() []FabricConfigChangeDetail`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *FabricConfigChangeDetailListAllOf) GetResultsOk() (*[]FabricConfigChangeDetail, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *FabricConfigChangeDetailListAllOf) SetResults(v []FabricConfigChangeDetail)`

SetResults sets Results field to given value.

### HasResults

`func (o *FabricConfigChangeDetailListAllOf) HasResults() bool`

HasResults returns a boolean if a field has been set.

### SetResultsNil

`func (o *FabricConfigChangeDetailListAllOf) SetResultsNil(b bool)`

 SetResultsNil sets the value for Results to be an explicit nil

### UnsetResults
`func (o *FabricConfigChangeDetailListAllOf) UnsetResults()`

UnsetResults ensures that no value is present for Results, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



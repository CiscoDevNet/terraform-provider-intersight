# NiaapiDcnmReleaseRecommendListAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | The total number of &#39;niaapi.DcnmReleaseRecommend&#39; resources matching the request, accross all pages. The &#39;Count&#39; attribute is included when the HTTP GET request includes the &#39;$inlinecount&#39; parameter. | [optional] 
**Results** | Pointer to [**[]NiaapiDcnmReleaseRecommend**](NiaapiDcnmReleaseRecommend.md) | The array of &#39;niaapi.DcnmReleaseRecommend&#39; resources matching the request. | [optional] 

## Methods

### NewNiaapiDcnmReleaseRecommendListAllOf

`func NewNiaapiDcnmReleaseRecommendListAllOf() *NiaapiDcnmReleaseRecommendListAllOf`

NewNiaapiDcnmReleaseRecommendListAllOf instantiates a new NiaapiDcnmReleaseRecommendListAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiaapiDcnmReleaseRecommendListAllOfWithDefaults

`func NewNiaapiDcnmReleaseRecommendListAllOfWithDefaults() *NiaapiDcnmReleaseRecommendListAllOf`

NewNiaapiDcnmReleaseRecommendListAllOfWithDefaults instantiates a new NiaapiDcnmReleaseRecommendListAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *NiaapiDcnmReleaseRecommendListAllOf) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *NiaapiDcnmReleaseRecommendListAllOf) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *NiaapiDcnmReleaseRecommendListAllOf) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *NiaapiDcnmReleaseRecommendListAllOf) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetResults

`func (o *NiaapiDcnmReleaseRecommendListAllOf) GetResults() []NiaapiDcnmReleaseRecommend`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *NiaapiDcnmReleaseRecommendListAllOf) GetResultsOk() (*[]NiaapiDcnmReleaseRecommend, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *NiaapiDcnmReleaseRecommendListAllOf) SetResults(v []NiaapiDcnmReleaseRecommend)`

SetResults sets Results field to given value.

### HasResults

`func (o *NiaapiDcnmReleaseRecommendListAllOf) HasResults() bool`

HasResults returns a boolean if a field has been set.

### SetResultsNil

`func (o *NiaapiDcnmReleaseRecommendListAllOf) SetResultsNil(b bool)`

 SetResultsNil sets the value for Results to be an explicit nil

### UnsetResults
`func (o *NiaapiDcnmReleaseRecommendListAllOf) UnsetResults()`

UnsetResults ensures that no value is present for Results, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



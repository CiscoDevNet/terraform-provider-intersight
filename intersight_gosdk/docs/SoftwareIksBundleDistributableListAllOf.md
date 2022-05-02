# SoftwareIksBundleDistributableListAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | The total number of &#39;software.IksBundleDistributable&#39; resources matching the request, accross all pages. The &#39;Count&#39; attribute is included when the HTTP GET request includes the &#39;$inlinecount&#39; parameter. | [optional] 
**Results** | Pointer to [**[]SoftwareIksBundleDistributable**](SoftwareIksBundleDistributable.md) | The array of &#39;software.IksBundleDistributable&#39; resources matching the request. | [optional] 

## Methods

### NewSoftwareIksBundleDistributableListAllOf

`func NewSoftwareIksBundleDistributableListAllOf() *SoftwareIksBundleDistributableListAllOf`

NewSoftwareIksBundleDistributableListAllOf instantiates a new SoftwareIksBundleDistributableListAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoftwareIksBundleDistributableListAllOfWithDefaults

`func NewSoftwareIksBundleDistributableListAllOfWithDefaults() *SoftwareIksBundleDistributableListAllOf`

NewSoftwareIksBundleDistributableListAllOfWithDefaults instantiates a new SoftwareIksBundleDistributableListAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *SoftwareIksBundleDistributableListAllOf) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *SoftwareIksBundleDistributableListAllOf) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *SoftwareIksBundleDistributableListAllOf) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *SoftwareIksBundleDistributableListAllOf) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetResults

`func (o *SoftwareIksBundleDistributableListAllOf) GetResults() []SoftwareIksBundleDistributable`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *SoftwareIksBundleDistributableListAllOf) GetResultsOk() (*[]SoftwareIksBundleDistributable, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *SoftwareIksBundleDistributableListAllOf) SetResults(v []SoftwareIksBundleDistributable)`

SetResults sets Results field to given value.

### HasResults

`func (o *SoftwareIksBundleDistributableListAllOf) HasResults() bool`

HasResults returns a boolean if a field has been set.

### SetResultsNil

`func (o *SoftwareIksBundleDistributableListAllOf) SetResultsNil(b bool)`

 SetResultsNil sets the value for Results to be an explicit nil

### UnsetResults
`func (o *SoftwareIksBundleDistributableListAllOf) UnsetResults()`

UnsetResults ensures that no value is present for Results, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# CatalystsdwanVedgeDeviceList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | The total number of &#39;catalystsdwan.VedgeDevice&#39; resources matching the request, accross all pages. The &#39;Count&#39; attribute is included when the HTTP GET request includes the &#39;$inlinecount&#39; parameter. | [optional] 
**Results** | Pointer to [**[]CatalystsdwanVedgeDevice**](CatalystsdwanVedgeDevice.md) | The array of &#39;catalystsdwan.VedgeDevice&#39; resources matching the request. | [optional] 

## Methods

### NewCatalystsdwanVedgeDeviceList

`func NewCatalystsdwanVedgeDeviceList() *CatalystsdwanVedgeDeviceList`

NewCatalystsdwanVedgeDeviceList instantiates a new CatalystsdwanVedgeDeviceList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalystsdwanVedgeDeviceListWithDefaults

`func NewCatalystsdwanVedgeDeviceListWithDefaults() *CatalystsdwanVedgeDeviceList`

NewCatalystsdwanVedgeDeviceListWithDefaults instantiates a new CatalystsdwanVedgeDeviceList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *CatalystsdwanVedgeDeviceList) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *CatalystsdwanVedgeDeviceList) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *CatalystsdwanVedgeDeviceList) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *CatalystsdwanVedgeDeviceList) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetResults

`func (o *CatalystsdwanVedgeDeviceList) GetResults() []CatalystsdwanVedgeDevice`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *CatalystsdwanVedgeDeviceList) GetResultsOk() (*[]CatalystsdwanVedgeDevice, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *CatalystsdwanVedgeDeviceList) SetResults(v []CatalystsdwanVedgeDevice)`

SetResults sets Results field to given value.

### HasResults

`func (o *CatalystsdwanVedgeDeviceList) HasResults() bool`

HasResults returns a boolean if a field has been set.

### SetResultsNil

`func (o *CatalystsdwanVedgeDeviceList) SetResultsNil(b bool)`

 SetResultsNil sets the value for Results to be an explicit nil

### UnsetResults
`func (o *CatalystsdwanVedgeDeviceList) UnsetResults()`

UnsetResults ensures that no value is present for Results, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



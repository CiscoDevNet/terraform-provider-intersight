# NiatelemetryApicPsuDetailsListAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | The total number of &#39;niatelemetry.ApicPsuDetails&#39; resources matching the request, accross all pages. The &#39;Count&#39; attribute is included when the HTTP GET request includes the &#39;$inlinecount&#39; parameter. | [optional] 
**Results** | Pointer to [**[]NiatelemetryApicPsuDetails**](NiatelemetryApicPsuDetails.md) | The array of &#39;niatelemetry.ApicPsuDetails&#39; resources matching the request. | [optional] 

## Methods

### NewNiatelemetryApicPsuDetailsListAllOf

`func NewNiatelemetryApicPsuDetailsListAllOf() *NiatelemetryApicPsuDetailsListAllOf`

NewNiatelemetryApicPsuDetailsListAllOf instantiates a new NiatelemetryApicPsuDetailsListAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryApicPsuDetailsListAllOfWithDefaults

`func NewNiatelemetryApicPsuDetailsListAllOfWithDefaults() *NiatelemetryApicPsuDetailsListAllOf`

NewNiatelemetryApicPsuDetailsListAllOfWithDefaults instantiates a new NiatelemetryApicPsuDetailsListAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *NiatelemetryApicPsuDetailsListAllOf) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *NiatelemetryApicPsuDetailsListAllOf) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *NiatelemetryApicPsuDetailsListAllOf) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *NiatelemetryApicPsuDetailsListAllOf) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetResults

`func (o *NiatelemetryApicPsuDetailsListAllOf) GetResults() []NiatelemetryApicPsuDetails`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *NiatelemetryApicPsuDetailsListAllOf) GetResultsOk() (*[]NiatelemetryApicPsuDetails, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *NiatelemetryApicPsuDetailsListAllOf) SetResults(v []NiatelemetryApicPsuDetails)`

SetResults sets Results field to given value.

### HasResults

`func (o *NiatelemetryApicPsuDetailsListAllOf) HasResults() bool`

HasResults returns a boolean if a field has been set.

### SetResultsNil

`func (o *NiatelemetryApicPsuDetailsListAllOf) SetResultsNil(b bool)`

 SetResultsNil sets the value for Results to be an explicit nil

### UnsetResults
`func (o *NiatelemetryApicPsuDetailsListAllOf) UnsetResults()`

UnsetResults ensures that no value is present for Results, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



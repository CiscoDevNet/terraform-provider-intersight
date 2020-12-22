# NiatelemetryAppDetailsList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | The total number of &#39;niatelemetry.AppDetails&#39; resources matching the request, accross all pages. The &#39;Count&#39; attribute is included when the HTTP GET request includes the &#39;$inlinecount&#39; parameter. | [optional] 
**Results** | Pointer to [**[]NiatelemetryAppDetails**](NiatelemetryAppDetails.md) | The array of &#39;niatelemetry.AppDetails&#39; resources matching the request. | [optional] 

## Methods

### NewNiatelemetryAppDetailsList

`func NewNiatelemetryAppDetailsList() *NiatelemetryAppDetailsList`

NewNiatelemetryAppDetailsList instantiates a new NiatelemetryAppDetailsList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryAppDetailsListWithDefaults

`func NewNiatelemetryAppDetailsListWithDefaults() *NiatelemetryAppDetailsList`

NewNiatelemetryAppDetailsListWithDefaults instantiates a new NiatelemetryAppDetailsList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *NiatelemetryAppDetailsList) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *NiatelemetryAppDetailsList) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *NiatelemetryAppDetailsList) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *NiatelemetryAppDetailsList) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetResults

`func (o *NiatelemetryAppDetailsList) GetResults() []NiatelemetryAppDetails`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *NiatelemetryAppDetailsList) GetResultsOk() (*[]NiatelemetryAppDetails, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *NiatelemetryAppDetailsList) SetResults(v []NiatelemetryAppDetails)`

SetResults sets Results field to given value.

### HasResults

`func (o *NiatelemetryAppDetailsList) HasResults() bool`

HasResults returns a boolean if a field has been set.

### SetResultsNil

`func (o *NiatelemetryAppDetailsList) SetResultsNil(b bool)`

 SetResultsNil sets the value for Results to be an explicit nil

### UnsetResults
`func (o *NiatelemetryAppDetailsList) UnsetResults()`

UnsetResults ensures that no value is present for Results, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



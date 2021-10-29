# NiatelemetryApicAppPluginDetailsList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | The total number of &#39;niatelemetry.ApicAppPluginDetails&#39; resources matching the request, accross all pages. The &#39;Count&#39; attribute is included when the HTTP GET request includes the &#39;$inlinecount&#39; parameter. | [optional] 
**Results** | Pointer to [**[]NiatelemetryApicAppPluginDetails**](NiatelemetryApicAppPluginDetails.md) | The array of &#39;niatelemetry.ApicAppPluginDetails&#39; resources matching the request. | [optional] 

## Methods

### NewNiatelemetryApicAppPluginDetailsList

`func NewNiatelemetryApicAppPluginDetailsList() *NiatelemetryApicAppPluginDetailsList`

NewNiatelemetryApicAppPluginDetailsList instantiates a new NiatelemetryApicAppPluginDetailsList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryApicAppPluginDetailsListWithDefaults

`func NewNiatelemetryApicAppPluginDetailsListWithDefaults() *NiatelemetryApicAppPluginDetailsList`

NewNiatelemetryApicAppPluginDetailsListWithDefaults instantiates a new NiatelemetryApicAppPluginDetailsList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *NiatelemetryApicAppPluginDetailsList) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *NiatelemetryApicAppPluginDetailsList) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *NiatelemetryApicAppPluginDetailsList) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *NiatelemetryApicAppPluginDetailsList) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetResults

`func (o *NiatelemetryApicAppPluginDetailsList) GetResults() []NiatelemetryApicAppPluginDetails`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *NiatelemetryApicAppPluginDetailsList) GetResultsOk() (*[]NiatelemetryApicAppPluginDetails, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *NiatelemetryApicAppPluginDetailsList) SetResults(v []NiatelemetryApicAppPluginDetails)`

SetResults sets Results field to given value.

### HasResults

`func (o *NiatelemetryApicAppPluginDetailsList) HasResults() bool`

HasResults returns a boolean if a field has been set.

### SetResultsNil

`func (o *NiatelemetryApicAppPluginDetailsList) SetResultsNil(b bool)`

 SetResultsNil sets the value for Results to be an explicit nil

### UnsetResults
`func (o *NiatelemetryApicAppPluginDetailsList) UnsetResults()`

UnsetResults ensures that no value is present for Results, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



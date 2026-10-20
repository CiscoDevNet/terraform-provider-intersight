# NiatelemetryHttpsAclFilterDetailsList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | The total number of &#39;niatelemetry.HttpsAclFilterDetails&#39; resources matching the request, accross all pages. The &#39;Count&#39; attribute is included when the HTTP GET request includes the &#39;$inlinecount&#39; parameter. | [optional] 
**Results** | Pointer to [**[]NiatelemetryHttpsAclFilterDetails**](NiatelemetryHttpsAclFilterDetails.md) | The array of &#39;niatelemetry.HttpsAclFilterDetails&#39; resources matching the request. | [optional] 

## Methods

### NewNiatelemetryHttpsAclFilterDetailsList

`func NewNiatelemetryHttpsAclFilterDetailsList() *NiatelemetryHttpsAclFilterDetailsList`

NewNiatelemetryHttpsAclFilterDetailsList instantiates a new NiatelemetryHttpsAclFilterDetailsList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryHttpsAclFilterDetailsListWithDefaults

`func NewNiatelemetryHttpsAclFilterDetailsListWithDefaults() *NiatelemetryHttpsAclFilterDetailsList`

NewNiatelemetryHttpsAclFilterDetailsListWithDefaults instantiates a new NiatelemetryHttpsAclFilterDetailsList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *NiatelemetryHttpsAclFilterDetailsList) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *NiatelemetryHttpsAclFilterDetailsList) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *NiatelemetryHttpsAclFilterDetailsList) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *NiatelemetryHttpsAclFilterDetailsList) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetResults

`func (o *NiatelemetryHttpsAclFilterDetailsList) GetResults() []NiatelemetryHttpsAclFilterDetails`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *NiatelemetryHttpsAclFilterDetailsList) GetResultsOk() (*[]NiatelemetryHttpsAclFilterDetails, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *NiatelemetryHttpsAclFilterDetailsList) SetResults(v []NiatelemetryHttpsAclFilterDetails)`

SetResults sets Results field to given value.

### HasResults

`func (o *NiatelemetryHttpsAclFilterDetailsList) HasResults() bool`

HasResults returns a boolean if a field has been set.

### SetResultsNil

`func (o *NiatelemetryHttpsAclFilterDetailsList) SetResultsNil(b bool)`

 SetResultsNil sets the value for Results to be an explicit nil

### UnsetResults
`func (o *NiatelemetryHttpsAclFilterDetailsList) UnsetResults()`

UnsetResults ensures that no value is present for Results, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# MerakiDeviceList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | The total number of &#39;meraki.Device&#39; resources matching the request, accross all pages. The &#39;Count&#39; attribute is included when the HTTP GET request includes the &#39;$inlinecount&#39; parameter. | [optional] 
**Results** | Pointer to [**[]MerakiDevice**](MerakiDevice.md) | The array of &#39;meraki.Device&#39; resources matching the request. | [optional] 

## Methods

### NewMerakiDeviceList

`func NewMerakiDeviceList() *MerakiDeviceList`

NewMerakiDeviceList instantiates a new MerakiDeviceList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerakiDeviceListWithDefaults

`func NewMerakiDeviceListWithDefaults() *MerakiDeviceList`

NewMerakiDeviceListWithDefaults instantiates a new MerakiDeviceList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *MerakiDeviceList) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *MerakiDeviceList) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *MerakiDeviceList) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *MerakiDeviceList) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetResults

`func (o *MerakiDeviceList) GetResults() []MerakiDevice`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *MerakiDeviceList) GetResultsOk() (*[]MerakiDevice, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *MerakiDeviceList) SetResults(v []MerakiDevice)`

SetResults sets Results field to given value.

### HasResults

`func (o *MerakiDeviceList) HasResults() bool`

HasResults returns a boolean if a field has been set.

### SetResultsNil

`func (o *MerakiDeviceList) SetResultsNil(b bool)`

 SetResultsNil sets the value for Results to be an explicit nil

### UnsetResults
`func (o *MerakiDeviceList) UnsetResults()`

UnsetResults ensures that no value is present for Results, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



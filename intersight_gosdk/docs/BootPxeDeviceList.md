# BootPxeDeviceList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | The total number of &#39;boot.PxeDevice&#39; resources matching the request, accross all pages. The &#39;Count&#39; attribute is included when the HTTP GET request includes the &#39;$inlinecount&#39; parameter. | [optional] 
**Results** | Pointer to [**[]BootPxeDevice**](boot.PxeDevice.md) | The array of &#39;boot.PxeDevice&#39; resources matching the request. | [optional] 

## Methods

### NewBootPxeDeviceList

`func NewBootPxeDeviceList() *BootPxeDeviceList`

NewBootPxeDeviceList instantiates a new BootPxeDeviceList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBootPxeDeviceListWithDefaults

`func NewBootPxeDeviceListWithDefaults() *BootPxeDeviceList`

NewBootPxeDeviceListWithDefaults instantiates a new BootPxeDeviceList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *BootPxeDeviceList) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *BootPxeDeviceList) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *BootPxeDeviceList) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *BootPxeDeviceList) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetResults

`func (o *BootPxeDeviceList) GetResults() []BootPxeDevice`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *BootPxeDeviceList) GetResultsOk() (*[]BootPxeDevice, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *BootPxeDeviceList) SetResults(v []BootPxeDevice)`

SetResults sets Results field to given value.

### HasResults

`func (o *BootPxeDeviceList) HasResults() bool`

HasResults returns a boolean if a field has been set.

### SetResultsNil

`func (o *BootPxeDeviceList) SetResultsNil(b bool)`

 SetResultsNil sets the value for Results to be an explicit nil

### UnsetResults
`func (o *BootPxeDeviceList) UnsetResults()`

UnsetResults ensures that no value is present for Results, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



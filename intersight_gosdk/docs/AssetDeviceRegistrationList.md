# AssetDeviceRegistrationList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | The total number of &#39;asset.DeviceRegistration&#39; resources matching the request, accross all pages. The &#39;Count&#39; attribute is included when the HTTP GET request includes the &#39;$inlinecount&#39; parameter. | [optional] 
**Results** | Pointer to [**[]AssetDeviceRegistration**](asset.DeviceRegistration.md) | The array of &#39;asset.DeviceRegistration&#39; resources matching the request. | [optional] 

## Methods

### NewAssetDeviceRegistrationList

`func NewAssetDeviceRegistrationList() *AssetDeviceRegistrationList`

NewAssetDeviceRegistrationList instantiates a new AssetDeviceRegistrationList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetDeviceRegistrationListWithDefaults

`func NewAssetDeviceRegistrationListWithDefaults() *AssetDeviceRegistrationList`

NewAssetDeviceRegistrationListWithDefaults instantiates a new AssetDeviceRegistrationList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *AssetDeviceRegistrationList) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *AssetDeviceRegistrationList) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *AssetDeviceRegistrationList) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *AssetDeviceRegistrationList) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetResults

`func (o *AssetDeviceRegistrationList) GetResults() []AssetDeviceRegistration`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *AssetDeviceRegistrationList) GetResultsOk() (*[]AssetDeviceRegistration, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *AssetDeviceRegistrationList) SetResults(v []AssetDeviceRegistration)`

SetResults sets Results field to given value.

### HasResults

`func (o *AssetDeviceRegistrationList) HasResults() bool`

HasResults returns a boolean if a field has been set.

### SetResultsNil

`func (o *AssetDeviceRegistrationList) SetResultsNil(b bool)`

 SetResultsNil sets the value for Results to be an explicit nil

### UnsetResults
`func (o *AssetDeviceRegistrationList) UnsetResults()`

UnsetResults ensures that no value is present for Results, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



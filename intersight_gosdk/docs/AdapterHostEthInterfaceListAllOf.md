# AdapterHostEthInterfaceListAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | The total number of &#39;adapter.HostEthInterface&#39; resources matching the request, accross all pages. The &#39;Count&#39; attribute is included when the HTTP GET request includes the &#39;$inlinecount&#39; parameter. | [optional] 
**Results** | Pointer to [**[]AdapterHostEthInterface**](adapter.HostEthInterface.md) | The array of &#39;adapter.HostEthInterface&#39; resources matching the request. | [optional] 

## Methods

### NewAdapterHostEthInterfaceListAllOf

`func NewAdapterHostEthInterfaceListAllOf() *AdapterHostEthInterfaceListAllOf`

NewAdapterHostEthInterfaceListAllOf instantiates a new AdapterHostEthInterfaceListAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdapterHostEthInterfaceListAllOfWithDefaults

`func NewAdapterHostEthInterfaceListAllOfWithDefaults() *AdapterHostEthInterfaceListAllOf`

NewAdapterHostEthInterfaceListAllOfWithDefaults instantiates a new AdapterHostEthInterfaceListAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *AdapterHostEthInterfaceListAllOf) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *AdapterHostEthInterfaceListAllOf) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *AdapterHostEthInterfaceListAllOf) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *AdapterHostEthInterfaceListAllOf) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetResults

`func (o *AdapterHostEthInterfaceListAllOf) GetResults() []AdapterHostEthInterface`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *AdapterHostEthInterfaceListAllOf) GetResultsOk() (*[]AdapterHostEthInterface, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *AdapterHostEthInterfaceListAllOf) SetResults(v []AdapterHostEthInterface)`

SetResults sets Results field to given value.

### HasResults

`func (o *AdapterHostEthInterfaceListAllOf) HasResults() bool`

HasResults returns a boolean if a field has been set.

### SetResultsNil

`func (o *AdapterHostEthInterfaceListAllOf) SetResultsNil(b bool)`

 SetResultsNil sets the value for Results to be an explicit nil

### UnsetResults
`func (o *AdapterHostEthInterfaceListAllOf) UnsetResults()`

UnsetResults ensures that no value is present for Results, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



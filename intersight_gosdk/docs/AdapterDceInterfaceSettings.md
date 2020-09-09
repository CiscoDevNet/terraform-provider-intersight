# AdapterDceInterfaceSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FecMode** | Pointer to **string** | Forward Error Correction (FEC) mode setting for the DCE interfaces of the adapter. FEC mode setting is supported only for Cisco VIC 14xx adapters. FEC mode &#39;cl74&#39; is unsupported for Cisco VIC 1495/1497. This setting will be ignored for unsupported adapters and for unavailable DCE interfaces. * &#x60;Auto&#x60; - Use auto-negotiation with peer switch to arrive at common negotiated FEC mode. * &#x60;cl91&#x60; - Use cl91 standard as FEC mode setting. &#39;Clause 91&#39; aka RS-FEC (&#39;ReedSolomon&#39; FEC) offers better error protection against bursty and random errors but adds latency. * &#x60;cl74&#x60; - Use cl74 standard as FEC mode setting. &#39;Clause 74&#39; aka FC-FEC (&#39;FireCode&#39; FEC) offers simple, low-latency protection against 1 burst/sparse bit error, but it is not good for random errors. * &#x60;Off&#x60; - Disable FEC mode on the DCE Interface. | [optional] [default to "Auto"]
**InterfaceId** | Pointer to **int64** | DCE interface id on which settings needs to be configured. Supported values are (0-3). | [optional] 

## Methods

### NewAdapterDceInterfaceSettings

`func NewAdapterDceInterfaceSettings() *AdapterDceInterfaceSettings`

NewAdapterDceInterfaceSettings instantiates a new AdapterDceInterfaceSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdapterDceInterfaceSettingsWithDefaults

`func NewAdapterDceInterfaceSettingsWithDefaults() *AdapterDceInterfaceSettings`

NewAdapterDceInterfaceSettingsWithDefaults instantiates a new AdapterDceInterfaceSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFecMode

`func (o *AdapterDceInterfaceSettings) GetFecMode() string`

GetFecMode returns the FecMode field if non-nil, zero value otherwise.

### GetFecModeOk

`func (o *AdapterDceInterfaceSettings) GetFecModeOk() (*string, bool)`

GetFecModeOk returns a tuple with the FecMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFecMode

`func (o *AdapterDceInterfaceSettings) SetFecMode(v string)`

SetFecMode sets FecMode field to given value.

### HasFecMode

`func (o *AdapterDceInterfaceSettings) HasFecMode() bool`

HasFecMode returns a boolean if a field has been set.

### GetInterfaceId

`func (o *AdapterDceInterfaceSettings) GetInterfaceId() int64`

GetInterfaceId returns the InterfaceId field if non-nil, zero value otherwise.

### GetInterfaceIdOk

`func (o *AdapterDceInterfaceSettings) GetInterfaceIdOk() (*int64, bool)`

GetInterfaceIdOk returns a tuple with the InterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceId

`func (o *AdapterDceInterfaceSettings) SetInterfaceId(v int64)`

SetInterfaceId sets InterfaceId field to given value.

### HasInterfaceId

`func (o *AdapterDceInterfaceSettings) HasInterfaceId() bool`

HasInterfaceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



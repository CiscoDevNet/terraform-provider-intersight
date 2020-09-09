# AdapterAdapterConfigAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DceInterfaceSettings** | Pointer to [**[]AdapterDceInterfaceSettings**](adapter.DceInterfaceSettings.md) |  | [optional] 
**EthSettings** | Pointer to [**AdapterEthSettings**](adapter.EthSettings.md) |  | [optional] 
**FcSettings** | Pointer to [**AdapterFcSettings**](adapter.FcSettings.md) |  | [optional] 
**PortChannelSettings** | Pointer to [**AdapterPortChannelSettings**](adapter.PortChannelSettings.md) |  | [optional] 
**SlotId** | Pointer to **string** | PCIe slot where the VIC adapter is installed. Supported values are (1-15) and MLOM. | [optional] 

## Methods

### NewAdapterAdapterConfigAllOf

`func NewAdapterAdapterConfigAllOf() *AdapterAdapterConfigAllOf`

NewAdapterAdapterConfigAllOf instantiates a new AdapterAdapterConfigAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdapterAdapterConfigAllOfWithDefaults

`func NewAdapterAdapterConfigAllOfWithDefaults() *AdapterAdapterConfigAllOf`

NewAdapterAdapterConfigAllOfWithDefaults instantiates a new AdapterAdapterConfigAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDceInterfaceSettings

`func (o *AdapterAdapterConfigAllOf) GetDceInterfaceSettings() []AdapterDceInterfaceSettings`

GetDceInterfaceSettings returns the DceInterfaceSettings field if non-nil, zero value otherwise.

### GetDceInterfaceSettingsOk

`func (o *AdapterAdapterConfigAllOf) GetDceInterfaceSettingsOk() (*[]AdapterDceInterfaceSettings, bool)`

GetDceInterfaceSettingsOk returns a tuple with the DceInterfaceSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDceInterfaceSettings

`func (o *AdapterAdapterConfigAllOf) SetDceInterfaceSettings(v []AdapterDceInterfaceSettings)`

SetDceInterfaceSettings sets DceInterfaceSettings field to given value.

### HasDceInterfaceSettings

`func (o *AdapterAdapterConfigAllOf) HasDceInterfaceSettings() bool`

HasDceInterfaceSettings returns a boolean if a field has been set.

### GetEthSettings

`func (o *AdapterAdapterConfigAllOf) GetEthSettings() AdapterEthSettings`

GetEthSettings returns the EthSettings field if non-nil, zero value otherwise.

### GetEthSettingsOk

`func (o *AdapterAdapterConfigAllOf) GetEthSettingsOk() (*AdapterEthSettings, bool)`

GetEthSettingsOk returns a tuple with the EthSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthSettings

`func (o *AdapterAdapterConfigAllOf) SetEthSettings(v AdapterEthSettings)`

SetEthSettings sets EthSettings field to given value.

### HasEthSettings

`func (o *AdapterAdapterConfigAllOf) HasEthSettings() bool`

HasEthSettings returns a boolean if a field has been set.

### GetFcSettings

`func (o *AdapterAdapterConfigAllOf) GetFcSettings() AdapterFcSettings`

GetFcSettings returns the FcSettings field if non-nil, zero value otherwise.

### GetFcSettingsOk

`func (o *AdapterAdapterConfigAllOf) GetFcSettingsOk() (*AdapterFcSettings, bool)`

GetFcSettingsOk returns a tuple with the FcSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcSettings

`func (o *AdapterAdapterConfigAllOf) SetFcSettings(v AdapterFcSettings)`

SetFcSettings sets FcSettings field to given value.

### HasFcSettings

`func (o *AdapterAdapterConfigAllOf) HasFcSettings() bool`

HasFcSettings returns a boolean if a field has been set.

### GetPortChannelSettings

`func (o *AdapterAdapterConfigAllOf) GetPortChannelSettings() AdapterPortChannelSettings`

GetPortChannelSettings returns the PortChannelSettings field if non-nil, zero value otherwise.

### GetPortChannelSettingsOk

`func (o *AdapterAdapterConfigAllOf) GetPortChannelSettingsOk() (*AdapterPortChannelSettings, bool)`

GetPortChannelSettingsOk returns a tuple with the PortChannelSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortChannelSettings

`func (o *AdapterAdapterConfigAllOf) SetPortChannelSettings(v AdapterPortChannelSettings)`

SetPortChannelSettings sets PortChannelSettings field to given value.

### HasPortChannelSettings

`func (o *AdapterAdapterConfigAllOf) HasPortChannelSettings() bool`

HasPortChannelSettings returns a boolean if a field has been set.

### GetSlotId

`func (o *AdapterAdapterConfigAllOf) GetSlotId() string`

GetSlotId returns the SlotId field if non-nil, zero value otherwise.

### GetSlotIdOk

`func (o *AdapterAdapterConfigAllOf) GetSlotIdOk() (*string, bool)`

GetSlotIdOk returns a tuple with the SlotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotId

`func (o *AdapterAdapterConfigAllOf) SetSlotId(v string)`

SetSlotId sets SlotId field to given value.

### HasSlotId

`func (o *AdapterAdapterConfigAllOf) HasSlotId() bool`

HasSlotId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



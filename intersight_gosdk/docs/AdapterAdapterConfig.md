# AdapterAdapterConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DceInterfaceSettings** | Pointer to [**[]AdapterDceInterfaceSettings**](adapter.DceInterfaceSettings.md) |  | [optional] 
**EthSettings** | Pointer to [**AdapterEthSettings**](adapter.EthSettings.md) |  | [optional] 
**FcSettings** | Pointer to [**AdapterFcSettings**](adapter.FcSettings.md) |  | [optional] 
**PortChannelSettings** | Pointer to [**AdapterPortChannelSettings**](adapter.PortChannelSettings.md) |  | [optional] 
**SlotId** | Pointer to **string** | PCIe slot where the VIC adapter is installed. Supported values are (1-15) and MLOM. | [optional] 

## Methods

### NewAdapterAdapterConfig

`func NewAdapterAdapterConfig() *AdapterAdapterConfig`

NewAdapterAdapterConfig instantiates a new AdapterAdapterConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdapterAdapterConfigWithDefaults

`func NewAdapterAdapterConfigWithDefaults() *AdapterAdapterConfig`

NewAdapterAdapterConfigWithDefaults instantiates a new AdapterAdapterConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDceInterfaceSettings

`func (o *AdapterAdapterConfig) GetDceInterfaceSettings() []AdapterDceInterfaceSettings`

GetDceInterfaceSettings returns the DceInterfaceSettings field if non-nil, zero value otherwise.

### GetDceInterfaceSettingsOk

`func (o *AdapterAdapterConfig) GetDceInterfaceSettingsOk() (*[]AdapterDceInterfaceSettings, bool)`

GetDceInterfaceSettingsOk returns a tuple with the DceInterfaceSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDceInterfaceSettings

`func (o *AdapterAdapterConfig) SetDceInterfaceSettings(v []AdapterDceInterfaceSettings)`

SetDceInterfaceSettings sets DceInterfaceSettings field to given value.

### HasDceInterfaceSettings

`func (o *AdapterAdapterConfig) HasDceInterfaceSettings() bool`

HasDceInterfaceSettings returns a boolean if a field has been set.

### GetEthSettings

`func (o *AdapterAdapterConfig) GetEthSettings() AdapterEthSettings`

GetEthSettings returns the EthSettings field if non-nil, zero value otherwise.

### GetEthSettingsOk

`func (o *AdapterAdapterConfig) GetEthSettingsOk() (*AdapterEthSettings, bool)`

GetEthSettingsOk returns a tuple with the EthSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthSettings

`func (o *AdapterAdapterConfig) SetEthSettings(v AdapterEthSettings)`

SetEthSettings sets EthSettings field to given value.

### HasEthSettings

`func (o *AdapterAdapterConfig) HasEthSettings() bool`

HasEthSettings returns a boolean if a field has been set.

### GetFcSettings

`func (o *AdapterAdapterConfig) GetFcSettings() AdapterFcSettings`

GetFcSettings returns the FcSettings field if non-nil, zero value otherwise.

### GetFcSettingsOk

`func (o *AdapterAdapterConfig) GetFcSettingsOk() (*AdapterFcSettings, bool)`

GetFcSettingsOk returns a tuple with the FcSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcSettings

`func (o *AdapterAdapterConfig) SetFcSettings(v AdapterFcSettings)`

SetFcSettings sets FcSettings field to given value.

### HasFcSettings

`func (o *AdapterAdapterConfig) HasFcSettings() bool`

HasFcSettings returns a boolean if a field has been set.

### GetPortChannelSettings

`func (o *AdapterAdapterConfig) GetPortChannelSettings() AdapterPortChannelSettings`

GetPortChannelSettings returns the PortChannelSettings field if non-nil, zero value otherwise.

### GetPortChannelSettingsOk

`func (o *AdapterAdapterConfig) GetPortChannelSettingsOk() (*AdapterPortChannelSettings, bool)`

GetPortChannelSettingsOk returns a tuple with the PortChannelSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortChannelSettings

`func (o *AdapterAdapterConfig) SetPortChannelSettings(v AdapterPortChannelSettings)`

SetPortChannelSettings sets PortChannelSettings field to given value.

### HasPortChannelSettings

`func (o *AdapterAdapterConfig) HasPortChannelSettings() bool`

HasPortChannelSettings returns a boolean if a field has been set.

### GetSlotId

`func (o *AdapterAdapterConfig) GetSlotId() string`

GetSlotId returns the SlotId field if non-nil, zero value otherwise.

### GetSlotIdOk

`func (o *AdapterAdapterConfig) GetSlotIdOk() (*string, bool)`

GetSlotIdOk returns a tuple with the SlotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotId

`func (o *AdapterAdapterConfig) SetSlotId(v string)`

SetSlotId sets SlotId field to given value.

### HasSlotId

`func (o *AdapterAdapterConfig) HasSlotId() bool`

HasSlotId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



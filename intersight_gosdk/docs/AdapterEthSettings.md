# AdapterEthSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LldpEnabled** | Pointer to **bool** | Status of LLDP protocol on the adapter interfaces. | [optional] 

## Methods

### NewAdapterEthSettings

`func NewAdapterEthSettings() *AdapterEthSettings`

NewAdapterEthSettings instantiates a new AdapterEthSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdapterEthSettingsWithDefaults

`func NewAdapterEthSettingsWithDefaults() *AdapterEthSettings`

NewAdapterEthSettingsWithDefaults instantiates a new AdapterEthSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLldpEnabled

`func (o *AdapterEthSettings) GetLldpEnabled() bool`

GetLldpEnabled returns the LldpEnabled field if non-nil, zero value otherwise.

### GetLldpEnabledOk

`func (o *AdapterEthSettings) GetLldpEnabledOk() (*bool, bool)`

GetLldpEnabledOk returns a tuple with the LldpEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLldpEnabled

`func (o *AdapterEthSettings) SetLldpEnabled(v bool)`

SetLldpEnabled sets LldpEnabled field to given value.

### HasLldpEnabled

`func (o *AdapterEthSettings) HasLldpEnabled() bool`

HasLldpEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



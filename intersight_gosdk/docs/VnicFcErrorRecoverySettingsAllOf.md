# VnicFcErrorRecoverySettingsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Enables Fibre Channel Error recovery. | [optional] 
**IoRetryCount** | Pointer to **int64** | The number of times an I/O request to a port is retried because the port is busy before the system decides the port is unavailable. | [optional] 
**IoRetryTimeout** | Pointer to **int64** | The number of seconds the adapter waits before aborting the pending command and resending the same IO request. | [optional] 
**LinkDownTimeout** | Pointer to **int64** | The number of milliseconds the port should actually be down before it is marked down and fabric connectivity is lost. | [optional] 
**PortDownTimeout** | Pointer to **int64** | The number of milliseconds a remote Fibre Channel port should be offline before informing the SCSI upper layer that the port is unavailable. For a server with a VIC adapter running ESXi, the recommended value is 10000. For a server with a port used to boot a Windows OS from the SAN, the recommended value is 5000 milliseconds. | [optional] 

## Methods

### NewVnicFcErrorRecoverySettingsAllOf

`func NewVnicFcErrorRecoverySettingsAllOf() *VnicFcErrorRecoverySettingsAllOf`

NewVnicFcErrorRecoverySettingsAllOf instantiates a new VnicFcErrorRecoverySettingsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVnicFcErrorRecoverySettingsAllOfWithDefaults

`func NewVnicFcErrorRecoverySettingsAllOfWithDefaults() *VnicFcErrorRecoverySettingsAllOf`

NewVnicFcErrorRecoverySettingsAllOfWithDefaults instantiates a new VnicFcErrorRecoverySettingsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *VnicFcErrorRecoverySettingsAllOf) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *VnicFcErrorRecoverySettingsAllOf) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *VnicFcErrorRecoverySettingsAllOf) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *VnicFcErrorRecoverySettingsAllOf) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetIoRetryCount

`func (o *VnicFcErrorRecoverySettingsAllOf) GetIoRetryCount() int64`

GetIoRetryCount returns the IoRetryCount field if non-nil, zero value otherwise.

### GetIoRetryCountOk

`func (o *VnicFcErrorRecoverySettingsAllOf) GetIoRetryCountOk() (*int64, bool)`

GetIoRetryCountOk returns a tuple with the IoRetryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIoRetryCount

`func (o *VnicFcErrorRecoverySettingsAllOf) SetIoRetryCount(v int64)`

SetIoRetryCount sets IoRetryCount field to given value.

### HasIoRetryCount

`func (o *VnicFcErrorRecoverySettingsAllOf) HasIoRetryCount() bool`

HasIoRetryCount returns a boolean if a field has been set.

### GetIoRetryTimeout

`func (o *VnicFcErrorRecoverySettingsAllOf) GetIoRetryTimeout() int64`

GetIoRetryTimeout returns the IoRetryTimeout field if non-nil, zero value otherwise.

### GetIoRetryTimeoutOk

`func (o *VnicFcErrorRecoverySettingsAllOf) GetIoRetryTimeoutOk() (*int64, bool)`

GetIoRetryTimeoutOk returns a tuple with the IoRetryTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIoRetryTimeout

`func (o *VnicFcErrorRecoverySettingsAllOf) SetIoRetryTimeout(v int64)`

SetIoRetryTimeout sets IoRetryTimeout field to given value.

### HasIoRetryTimeout

`func (o *VnicFcErrorRecoverySettingsAllOf) HasIoRetryTimeout() bool`

HasIoRetryTimeout returns a boolean if a field has been set.

### GetLinkDownTimeout

`func (o *VnicFcErrorRecoverySettingsAllOf) GetLinkDownTimeout() int64`

GetLinkDownTimeout returns the LinkDownTimeout field if non-nil, zero value otherwise.

### GetLinkDownTimeoutOk

`func (o *VnicFcErrorRecoverySettingsAllOf) GetLinkDownTimeoutOk() (*int64, bool)`

GetLinkDownTimeoutOk returns a tuple with the LinkDownTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkDownTimeout

`func (o *VnicFcErrorRecoverySettingsAllOf) SetLinkDownTimeout(v int64)`

SetLinkDownTimeout sets LinkDownTimeout field to given value.

### HasLinkDownTimeout

`func (o *VnicFcErrorRecoverySettingsAllOf) HasLinkDownTimeout() bool`

HasLinkDownTimeout returns a boolean if a field has been set.

### GetPortDownTimeout

`func (o *VnicFcErrorRecoverySettingsAllOf) GetPortDownTimeout() int64`

GetPortDownTimeout returns the PortDownTimeout field if non-nil, zero value otherwise.

### GetPortDownTimeoutOk

`func (o *VnicFcErrorRecoverySettingsAllOf) GetPortDownTimeoutOk() (*int64, bool)`

GetPortDownTimeoutOk returns a tuple with the PortDownTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortDownTimeout

`func (o *VnicFcErrorRecoverySettingsAllOf) SetPortDownTimeout(v int64)`

SetPortDownTimeout sets PortDownTimeout field to given value.

### HasPortDownTimeout

`func (o *VnicFcErrorRecoverySettingsAllOf) HasPortDownTimeout() bool`

HasPortDownTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



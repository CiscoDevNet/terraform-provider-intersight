# VnicFcErrorRecoverySettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Enables Fibre Channel Error recovery. | [optional] 
**IoRetryCount** | Pointer to **int64** | The number of times an I/O request to a port is retried because the port is busy before the system decides the port is unavailable. | [optional] 
**IoRetryTimeout** | Pointer to **int64** | The number of seconds the adapter waits before aborting the pending command and resending the same IO request. | [optional] 
**LinkDownTimeout** | Pointer to **int64** | The number of milliseconds the port should actually be down before it is marked down and fabric connectivity is lost. | [optional] 
**PortDownTimeout** | Pointer to **int64** | The number of milliseconds a remote Fibre Channel port should be offline before informing the SCSI upper layer that the port is unavailable. For a server with a VIC adapter running ESXi, the recommended value is 10000. For a server with a port used to boot a Windows OS from the SAN, the recommended value is 5000 milliseconds. | [optional] 

## Methods

### NewVnicFcErrorRecoverySettings

`func NewVnicFcErrorRecoverySettings() *VnicFcErrorRecoverySettings`

NewVnicFcErrorRecoverySettings instantiates a new VnicFcErrorRecoverySettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVnicFcErrorRecoverySettingsWithDefaults

`func NewVnicFcErrorRecoverySettingsWithDefaults() *VnicFcErrorRecoverySettings`

NewVnicFcErrorRecoverySettingsWithDefaults instantiates a new VnicFcErrorRecoverySettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *VnicFcErrorRecoverySettings) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *VnicFcErrorRecoverySettings) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *VnicFcErrorRecoverySettings) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *VnicFcErrorRecoverySettings) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetIoRetryCount

`func (o *VnicFcErrorRecoverySettings) GetIoRetryCount() int64`

GetIoRetryCount returns the IoRetryCount field if non-nil, zero value otherwise.

### GetIoRetryCountOk

`func (o *VnicFcErrorRecoverySettings) GetIoRetryCountOk() (*int64, bool)`

GetIoRetryCountOk returns a tuple with the IoRetryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIoRetryCount

`func (o *VnicFcErrorRecoverySettings) SetIoRetryCount(v int64)`

SetIoRetryCount sets IoRetryCount field to given value.

### HasIoRetryCount

`func (o *VnicFcErrorRecoverySettings) HasIoRetryCount() bool`

HasIoRetryCount returns a boolean if a field has been set.

### GetIoRetryTimeout

`func (o *VnicFcErrorRecoverySettings) GetIoRetryTimeout() int64`

GetIoRetryTimeout returns the IoRetryTimeout field if non-nil, zero value otherwise.

### GetIoRetryTimeoutOk

`func (o *VnicFcErrorRecoverySettings) GetIoRetryTimeoutOk() (*int64, bool)`

GetIoRetryTimeoutOk returns a tuple with the IoRetryTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIoRetryTimeout

`func (o *VnicFcErrorRecoverySettings) SetIoRetryTimeout(v int64)`

SetIoRetryTimeout sets IoRetryTimeout field to given value.

### HasIoRetryTimeout

`func (o *VnicFcErrorRecoverySettings) HasIoRetryTimeout() bool`

HasIoRetryTimeout returns a boolean if a field has been set.

### GetLinkDownTimeout

`func (o *VnicFcErrorRecoverySettings) GetLinkDownTimeout() int64`

GetLinkDownTimeout returns the LinkDownTimeout field if non-nil, zero value otherwise.

### GetLinkDownTimeoutOk

`func (o *VnicFcErrorRecoverySettings) GetLinkDownTimeoutOk() (*int64, bool)`

GetLinkDownTimeoutOk returns a tuple with the LinkDownTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkDownTimeout

`func (o *VnicFcErrorRecoverySettings) SetLinkDownTimeout(v int64)`

SetLinkDownTimeout sets LinkDownTimeout field to given value.

### HasLinkDownTimeout

`func (o *VnicFcErrorRecoverySettings) HasLinkDownTimeout() bool`

HasLinkDownTimeout returns a boolean if a field has been set.

### GetPortDownTimeout

`func (o *VnicFcErrorRecoverySettings) GetPortDownTimeout() int64`

GetPortDownTimeout returns the PortDownTimeout field if non-nil, zero value otherwise.

### GetPortDownTimeoutOk

`func (o *VnicFcErrorRecoverySettings) GetPortDownTimeoutOk() (*int64, bool)`

GetPortDownTimeoutOk returns a tuple with the PortDownTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortDownTimeout

`func (o *VnicFcErrorRecoverySettings) SetPortDownTimeout(v int64)`

SetPortDownTimeout sets PortDownTimeout field to given value.

### HasPortDownTimeout

`func (o *VnicFcErrorRecoverySettings) HasPortDownTimeout() bool`

HasPortDownTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



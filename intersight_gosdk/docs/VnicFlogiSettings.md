# VnicFlogiSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Retries** | Pointer to **int64** | The number of times that the system tries to log in to the fabric after the first failure. | [optional] 
**Timeout** | Pointer to **int64** | The number of milliseconds that the system waits before it tries to log in again. | [optional] 

## Methods

### NewVnicFlogiSettings

`func NewVnicFlogiSettings() *VnicFlogiSettings`

NewVnicFlogiSettings instantiates a new VnicFlogiSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVnicFlogiSettingsWithDefaults

`func NewVnicFlogiSettingsWithDefaults() *VnicFlogiSettings`

NewVnicFlogiSettingsWithDefaults instantiates a new VnicFlogiSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRetries

`func (o *VnicFlogiSettings) GetRetries() int64`

GetRetries returns the Retries field if non-nil, zero value otherwise.

### GetRetriesOk

`func (o *VnicFlogiSettings) GetRetriesOk() (*int64, bool)`

GetRetriesOk returns a tuple with the Retries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetries

`func (o *VnicFlogiSettings) SetRetries(v int64)`

SetRetries sets Retries field to given value.

### HasRetries

`func (o *VnicFlogiSettings) HasRetries() bool`

HasRetries returns a boolean if a field has been set.

### GetTimeout

`func (o *VnicFlogiSettings) GetTimeout() int64`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *VnicFlogiSettings) GetTimeoutOk() (*int64, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *VnicFlogiSettings) SetTimeout(v int64)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *VnicFlogiSettings) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



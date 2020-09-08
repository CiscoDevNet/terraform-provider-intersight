# VnicFlogiSettingsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Retries** | Pointer to **int64** | The number of times that the system tries to log in to the fabric after the first failure. | [optional] 
**Timeout** | Pointer to **int64** | The number of milliseconds that the system waits before it tries to log in again. | [optional] 

## Methods

### NewVnicFlogiSettingsAllOf

`func NewVnicFlogiSettingsAllOf() *VnicFlogiSettingsAllOf`

NewVnicFlogiSettingsAllOf instantiates a new VnicFlogiSettingsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVnicFlogiSettingsAllOfWithDefaults

`func NewVnicFlogiSettingsAllOfWithDefaults() *VnicFlogiSettingsAllOf`

NewVnicFlogiSettingsAllOfWithDefaults instantiates a new VnicFlogiSettingsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRetries

`func (o *VnicFlogiSettingsAllOf) GetRetries() int64`

GetRetries returns the Retries field if non-nil, zero value otherwise.

### GetRetriesOk

`func (o *VnicFlogiSettingsAllOf) GetRetriesOk() (*int64, bool)`

GetRetriesOk returns a tuple with the Retries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetries

`func (o *VnicFlogiSettingsAllOf) SetRetries(v int64)`

SetRetries sets Retries field to given value.

### HasRetries

`func (o *VnicFlogiSettingsAllOf) HasRetries() bool`

HasRetries returns a boolean if a field has been set.

### GetTimeout

`func (o *VnicFlogiSettingsAllOf) GetTimeout() int64`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *VnicFlogiSettingsAllOf) GetTimeoutOk() (*int64, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *VnicFlogiSettingsAllOf) SetTimeout(v int64)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *VnicFlogiSettingsAllOf) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



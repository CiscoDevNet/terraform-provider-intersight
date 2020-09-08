# VnicPlogiSettingsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Retries** | Pointer to **int64** | The number of times that the system tries to log in to a port after the first failure. | [optional] 
**Timeout** | Pointer to **int64** | The number of milliseconds that the system waits before it tries to log in again. | [optional] 

## Methods

### NewVnicPlogiSettingsAllOf

`func NewVnicPlogiSettingsAllOf() *VnicPlogiSettingsAllOf`

NewVnicPlogiSettingsAllOf instantiates a new VnicPlogiSettingsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVnicPlogiSettingsAllOfWithDefaults

`func NewVnicPlogiSettingsAllOfWithDefaults() *VnicPlogiSettingsAllOf`

NewVnicPlogiSettingsAllOfWithDefaults instantiates a new VnicPlogiSettingsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRetries

`func (o *VnicPlogiSettingsAllOf) GetRetries() int64`

GetRetries returns the Retries field if non-nil, zero value otherwise.

### GetRetriesOk

`func (o *VnicPlogiSettingsAllOf) GetRetriesOk() (*int64, bool)`

GetRetriesOk returns a tuple with the Retries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetries

`func (o *VnicPlogiSettingsAllOf) SetRetries(v int64)`

SetRetries sets Retries field to given value.

### HasRetries

`func (o *VnicPlogiSettingsAllOf) HasRetries() bool`

HasRetries returns a boolean if a field has been set.

### GetTimeout

`func (o *VnicPlogiSettingsAllOf) GetTimeout() int64`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *VnicPlogiSettingsAllOf) GetTimeoutOk() (*int64, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *VnicPlogiSettingsAllOf) SetTimeout(v int64)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *VnicPlogiSettingsAllOf) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



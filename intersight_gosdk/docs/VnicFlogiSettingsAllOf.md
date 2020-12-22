# VnicFlogiSettingsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "vnic.FlogiSettings"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "vnic.FlogiSettings"]
**Retries** | Pointer to **int64** | The number of times that the system tries to log in to the fabric after the first failure. | [optional] [default to 8]
**Timeout** | Pointer to **int64** | The number of milliseconds that the system waits before it tries to log in again. | [optional] [default to 4000]

## Methods

### NewVnicFlogiSettingsAllOf

`func NewVnicFlogiSettingsAllOf(classId string, objectType string, ) *VnicFlogiSettingsAllOf`

NewVnicFlogiSettingsAllOf instantiates a new VnicFlogiSettingsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVnicFlogiSettingsAllOfWithDefaults

`func NewVnicFlogiSettingsAllOfWithDefaults() *VnicFlogiSettingsAllOf`

NewVnicFlogiSettingsAllOfWithDefaults instantiates a new VnicFlogiSettingsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VnicFlogiSettingsAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VnicFlogiSettingsAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VnicFlogiSettingsAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VnicFlogiSettingsAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VnicFlogiSettingsAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VnicFlogiSettingsAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


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



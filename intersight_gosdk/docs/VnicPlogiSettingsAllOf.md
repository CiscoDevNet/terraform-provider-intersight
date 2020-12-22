# VnicPlogiSettingsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "vnic.PlogiSettings"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "vnic.PlogiSettings"]
**Retries** | Pointer to **int64** | The number of times that the system tries to log in to a port after the first failure. | [optional] [default to 8]
**Timeout** | Pointer to **int64** | The number of milliseconds that the system waits before it tries to log in again. | [optional] [default to 20000]

## Methods

### NewVnicPlogiSettingsAllOf

`func NewVnicPlogiSettingsAllOf(classId string, objectType string, ) *VnicPlogiSettingsAllOf`

NewVnicPlogiSettingsAllOf instantiates a new VnicPlogiSettingsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVnicPlogiSettingsAllOfWithDefaults

`func NewVnicPlogiSettingsAllOfWithDefaults() *VnicPlogiSettingsAllOf`

NewVnicPlogiSettingsAllOfWithDefaults instantiates a new VnicPlogiSettingsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VnicPlogiSettingsAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VnicPlogiSettingsAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VnicPlogiSettingsAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VnicPlogiSettingsAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VnicPlogiSettingsAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VnicPlogiSettingsAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


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



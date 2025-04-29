# FabricSecKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "fabric.SecKey"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "fabric.SecKey"]
**CryptographicAlgorithm** | Pointer to **string** | The cryptographic algorithm that employs the cipher-based message authentication code (CMAC) mode of operation with advanced encryption standard (AES). * &#x60;AES_256_CMAC&#x60; - Uses the AES (Advanced Encryption Standard) algorithm with a 256-bit key to generate a CMAC. * &#x60;AES_128_CMAC&#x60; - Uses the AES (Advanced Encryption Standard) algorithm with a 128-bit key to generate a CMAC. | [optional] [default to "AES_256_CMAC"]
**Id** | Pointer to **string** | Must have an even number of hexadecimal characters (including 0-9 and A-F, only) with a length between 2 and 64 characters. For example, \&quot;10\&quot;, \&quot;2000\&quot;, \&quot;ABCD1234\&quot;. | [optional] 
**IsOctetStringSet** | Pointer to **bool** | Indicates whether the value of the &#39;octetString&#39; property has been set. | [optional] [readonly] [default to false]
**KeyType** | Pointer to **string** | The type of encryption used for the specified key. * &#x60;Type-0&#x60; - No encryption for the specified octetString. * &#x60;Type-6&#x60; - Proprietary advanced encryption standard for the specified octetString. * &#x60;Type-7&#x60; - Proprietary insecure encryption for the specified octetString. | [optional] [default to "Type-0"]
**OctetString** | Pointer to **string** | The key octet string is a shared secret used in cryptographic operations. The valid size and format of the octet string depend on the selected KeyCryptographicAlgorithm and KeyEncryptionType. It should start with the character &#39;J&#39;. | [optional] 
**SendLifetimeDuration** | Pointer to **int64** | The key lifetime duration in seconds after the start time. If a non-zero value is configured for the duration, the end time configuration for the key lifetime is ignored. | [optional] 
**SendLifetimeEndTime** | Pointer to **time.Time** | The time of day and date when the key becomes inactive. | [optional] 
**SendLifetimeInfinite** | Pointer to **bool** | Indicates that the key remains active indefinitely after the specified start time. When this parameter is set, the end time and duration configurations for the key lifetime are ignored. | [optional] [default to false]
**SendLifetimeStartTime** | Pointer to **time.Time** | The time of day and date when the key becomes active. | [optional] 
**SendLifetimeTimeZone** | Pointer to **string** | The time zone used for key lifetime configurations. * &#x60;UTC&#x60; - The Universal Time (UTC) for key lifetime configurations. * &#x60;Local&#x60; - The local time zone of the device for key lifetime configurations. | [optional] [default to "UTC"]
**SendLifetimeUnlimited** | Pointer to **bool** | Indicates that the key is always active. When this parameter is set, all other key lifetime configurations are ignored. | [optional] [default to true]

## Methods

### NewFabricSecKey

`func NewFabricSecKey(classId string, objectType string, ) *FabricSecKey`

NewFabricSecKey instantiates a new FabricSecKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricSecKeyWithDefaults

`func NewFabricSecKeyWithDefaults() *FabricSecKey`

NewFabricSecKeyWithDefaults instantiates a new FabricSecKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FabricSecKey) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FabricSecKey) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FabricSecKey) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FabricSecKey) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FabricSecKey) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FabricSecKey) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCryptographicAlgorithm

`func (o *FabricSecKey) GetCryptographicAlgorithm() string`

GetCryptographicAlgorithm returns the CryptographicAlgorithm field if non-nil, zero value otherwise.

### GetCryptographicAlgorithmOk

`func (o *FabricSecKey) GetCryptographicAlgorithmOk() (*string, bool)`

GetCryptographicAlgorithmOk returns a tuple with the CryptographicAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptographicAlgorithm

`func (o *FabricSecKey) SetCryptographicAlgorithm(v string)`

SetCryptographicAlgorithm sets CryptographicAlgorithm field to given value.

### HasCryptographicAlgorithm

`func (o *FabricSecKey) HasCryptographicAlgorithm() bool`

HasCryptographicAlgorithm returns a boolean if a field has been set.

### GetId

`func (o *FabricSecKey) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FabricSecKey) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FabricSecKey) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FabricSecKey) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsOctetStringSet

`func (o *FabricSecKey) GetIsOctetStringSet() bool`

GetIsOctetStringSet returns the IsOctetStringSet field if non-nil, zero value otherwise.

### GetIsOctetStringSetOk

`func (o *FabricSecKey) GetIsOctetStringSetOk() (*bool, bool)`

GetIsOctetStringSetOk returns a tuple with the IsOctetStringSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOctetStringSet

`func (o *FabricSecKey) SetIsOctetStringSet(v bool)`

SetIsOctetStringSet sets IsOctetStringSet field to given value.

### HasIsOctetStringSet

`func (o *FabricSecKey) HasIsOctetStringSet() bool`

HasIsOctetStringSet returns a boolean if a field has been set.

### GetKeyType

`func (o *FabricSecKey) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *FabricSecKey) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *FabricSecKey) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.

### HasKeyType

`func (o *FabricSecKey) HasKeyType() bool`

HasKeyType returns a boolean if a field has been set.

### GetOctetString

`func (o *FabricSecKey) GetOctetString() string`

GetOctetString returns the OctetString field if non-nil, zero value otherwise.

### GetOctetStringOk

`func (o *FabricSecKey) GetOctetStringOk() (*string, bool)`

GetOctetStringOk returns a tuple with the OctetString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOctetString

`func (o *FabricSecKey) SetOctetString(v string)`

SetOctetString sets OctetString field to given value.

### HasOctetString

`func (o *FabricSecKey) HasOctetString() bool`

HasOctetString returns a boolean if a field has been set.

### GetSendLifetimeDuration

`func (o *FabricSecKey) GetSendLifetimeDuration() int64`

GetSendLifetimeDuration returns the SendLifetimeDuration field if non-nil, zero value otherwise.

### GetSendLifetimeDurationOk

`func (o *FabricSecKey) GetSendLifetimeDurationOk() (*int64, bool)`

GetSendLifetimeDurationOk returns a tuple with the SendLifetimeDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendLifetimeDuration

`func (o *FabricSecKey) SetSendLifetimeDuration(v int64)`

SetSendLifetimeDuration sets SendLifetimeDuration field to given value.

### HasSendLifetimeDuration

`func (o *FabricSecKey) HasSendLifetimeDuration() bool`

HasSendLifetimeDuration returns a boolean if a field has been set.

### GetSendLifetimeEndTime

`func (o *FabricSecKey) GetSendLifetimeEndTime() time.Time`

GetSendLifetimeEndTime returns the SendLifetimeEndTime field if non-nil, zero value otherwise.

### GetSendLifetimeEndTimeOk

`func (o *FabricSecKey) GetSendLifetimeEndTimeOk() (*time.Time, bool)`

GetSendLifetimeEndTimeOk returns a tuple with the SendLifetimeEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendLifetimeEndTime

`func (o *FabricSecKey) SetSendLifetimeEndTime(v time.Time)`

SetSendLifetimeEndTime sets SendLifetimeEndTime field to given value.

### HasSendLifetimeEndTime

`func (o *FabricSecKey) HasSendLifetimeEndTime() bool`

HasSendLifetimeEndTime returns a boolean if a field has been set.

### GetSendLifetimeInfinite

`func (o *FabricSecKey) GetSendLifetimeInfinite() bool`

GetSendLifetimeInfinite returns the SendLifetimeInfinite field if non-nil, zero value otherwise.

### GetSendLifetimeInfiniteOk

`func (o *FabricSecKey) GetSendLifetimeInfiniteOk() (*bool, bool)`

GetSendLifetimeInfiniteOk returns a tuple with the SendLifetimeInfinite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendLifetimeInfinite

`func (o *FabricSecKey) SetSendLifetimeInfinite(v bool)`

SetSendLifetimeInfinite sets SendLifetimeInfinite field to given value.

### HasSendLifetimeInfinite

`func (o *FabricSecKey) HasSendLifetimeInfinite() bool`

HasSendLifetimeInfinite returns a boolean if a field has been set.

### GetSendLifetimeStartTime

`func (o *FabricSecKey) GetSendLifetimeStartTime() time.Time`

GetSendLifetimeStartTime returns the SendLifetimeStartTime field if non-nil, zero value otherwise.

### GetSendLifetimeStartTimeOk

`func (o *FabricSecKey) GetSendLifetimeStartTimeOk() (*time.Time, bool)`

GetSendLifetimeStartTimeOk returns a tuple with the SendLifetimeStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendLifetimeStartTime

`func (o *FabricSecKey) SetSendLifetimeStartTime(v time.Time)`

SetSendLifetimeStartTime sets SendLifetimeStartTime field to given value.

### HasSendLifetimeStartTime

`func (o *FabricSecKey) HasSendLifetimeStartTime() bool`

HasSendLifetimeStartTime returns a boolean if a field has been set.

### GetSendLifetimeTimeZone

`func (o *FabricSecKey) GetSendLifetimeTimeZone() string`

GetSendLifetimeTimeZone returns the SendLifetimeTimeZone field if non-nil, zero value otherwise.

### GetSendLifetimeTimeZoneOk

`func (o *FabricSecKey) GetSendLifetimeTimeZoneOk() (*string, bool)`

GetSendLifetimeTimeZoneOk returns a tuple with the SendLifetimeTimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendLifetimeTimeZone

`func (o *FabricSecKey) SetSendLifetimeTimeZone(v string)`

SetSendLifetimeTimeZone sets SendLifetimeTimeZone field to given value.

### HasSendLifetimeTimeZone

`func (o *FabricSecKey) HasSendLifetimeTimeZone() bool`

HasSendLifetimeTimeZone returns a boolean if a field has been set.

### GetSendLifetimeUnlimited

`func (o *FabricSecKey) GetSendLifetimeUnlimited() bool`

GetSendLifetimeUnlimited returns the SendLifetimeUnlimited field if non-nil, zero value otherwise.

### GetSendLifetimeUnlimitedOk

`func (o *FabricSecKey) GetSendLifetimeUnlimitedOk() (*bool, bool)`

GetSendLifetimeUnlimitedOk returns a tuple with the SendLifetimeUnlimited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendLifetimeUnlimited

`func (o *FabricSecKey) SetSendLifetimeUnlimited(v bool)`

SetSendLifetimeUnlimited sets SendLifetimeUnlimited field to given value.

### HasSendLifetimeUnlimited

`func (o *FabricSecKey) HasSendLifetimeUnlimited() bool`

HasSendLifetimeUnlimited returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



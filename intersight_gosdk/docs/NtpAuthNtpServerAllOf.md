# NtpAuthNtpServerAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "ntp.AuthNtpServer"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "ntp.AuthNtpServer"]
**KeyType** | Pointer to **string** | Type of symmetric key to use for this server. * &#x60;SHA1&#x60; - Key type used by the authentication is SHA1. | [optional] [default to "SHA1"]
**ServerName** | Pointer to **string** | Server hostname or IP address. | [optional] 
**SymKeyId** | Pointer to **int64** | The key ID is a positive integer that identifies a cryptographic key used to authenticate NTP messages. | [optional] 
**SymKeyValue** | Pointer to **string** | The value of the symmetric key. | [optional] 

## Methods

### NewNtpAuthNtpServerAllOf

`func NewNtpAuthNtpServerAllOf(classId string, objectType string, ) *NtpAuthNtpServerAllOf`

NewNtpAuthNtpServerAllOf instantiates a new NtpAuthNtpServerAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNtpAuthNtpServerAllOfWithDefaults

`func NewNtpAuthNtpServerAllOfWithDefaults() *NtpAuthNtpServerAllOf`

NewNtpAuthNtpServerAllOfWithDefaults instantiates a new NtpAuthNtpServerAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NtpAuthNtpServerAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NtpAuthNtpServerAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NtpAuthNtpServerAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NtpAuthNtpServerAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NtpAuthNtpServerAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NtpAuthNtpServerAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetKeyType

`func (o *NtpAuthNtpServerAllOf) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *NtpAuthNtpServerAllOf) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *NtpAuthNtpServerAllOf) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.

### HasKeyType

`func (o *NtpAuthNtpServerAllOf) HasKeyType() bool`

HasKeyType returns a boolean if a field has been set.

### GetServerName

`func (o *NtpAuthNtpServerAllOf) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *NtpAuthNtpServerAllOf) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *NtpAuthNtpServerAllOf) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *NtpAuthNtpServerAllOf) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### GetSymKeyId

`func (o *NtpAuthNtpServerAllOf) GetSymKeyId() int64`

GetSymKeyId returns the SymKeyId field if non-nil, zero value otherwise.

### GetSymKeyIdOk

`func (o *NtpAuthNtpServerAllOf) GetSymKeyIdOk() (*int64, bool)`

GetSymKeyIdOk returns a tuple with the SymKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymKeyId

`func (o *NtpAuthNtpServerAllOf) SetSymKeyId(v int64)`

SetSymKeyId sets SymKeyId field to given value.

### HasSymKeyId

`func (o *NtpAuthNtpServerAllOf) HasSymKeyId() bool`

HasSymKeyId returns a boolean if a field has been set.

### GetSymKeyValue

`func (o *NtpAuthNtpServerAllOf) GetSymKeyValue() string`

GetSymKeyValue returns the SymKeyValue field if non-nil, zero value otherwise.

### GetSymKeyValueOk

`func (o *NtpAuthNtpServerAllOf) GetSymKeyValueOk() (*string, bool)`

GetSymKeyValueOk returns a tuple with the SymKeyValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymKeyValue

`func (o *NtpAuthNtpServerAllOf) SetSymKeyValue(v string)`

SetSymKeyValue sets SymKeyValue field to given value.

### HasSymKeyValue

`func (o *NtpAuthNtpServerAllOf) HasSymKeyValue() bool`

HasSymKeyValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# MemoryPersistentMemoryLocalSecurityAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "memory.PersistentMemoryLocalSecurity"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "memory.PersistentMemoryLocalSecurity"]
**Enabled** | Pointer to **bool** | Persistent Memory Security state. | [optional] [default to false]
**IsSecurePassphraseSet** | Pointer to **bool** | Indicates whether the value of the &#39;securePassphrase&#39; property has been set. | [optional] [readonly] [default to false]
**SecurePassphrase** | Pointer to **string** | Secure passphrase to be applied on the Persistent Memory Modules on the server. The allowed characters are a-z, A to Z, 0-9, and special characters &#x3D;, \\u0021, &amp;, \\#, $, %, +, ^, @, _, *, -. | [optional] 

## Methods

### NewMemoryPersistentMemoryLocalSecurityAllOf

`func NewMemoryPersistentMemoryLocalSecurityAllOf(classId string, objectType string, ) *MemoryPersistentMemoryLocalSecurityAllOf`

NewMemoryPersistentMemoryLocalSecurityAllOf instantiates a new MemoryPersistentMemoryLocalSecurityAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemoryPersistentMemoryLocalSecurityAllOfWithDefaults

`func NewMemoryPersistentMemoryLocalSecurityAllOfWithDefaults() *MemoryPersistentMemoryLocalSecurityAllOf`

NewMemoryPersistentMemoryLocalSecurityAllOfWithDefaults instantiates a new MemoryPersistentMemoryLocalSecurityAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *MemoryPersistentMemoryLocalSecurityAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *MemoryPersistentMemoryLocalSecurityAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *MemoryPersistentMemoryLocalSecurityAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *MemoryPersistentMemoryLocalSecurityAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *MemoryPersistentMemoryLocalSecurityAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *MemoryPersistentMemoryLocalSecurityAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEnabled

`func (o *MemoryPersistentMemoryLocalSecurityAllOf) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *MemoryPersistentMemoryLocalSecurityAllOf) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *MemoryPersistentMemoryLocalSecurityAllOf) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *MemoryPersistentMemoryLocalSecurityAllOf) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetIsSecurePassphraseSet

`func (o *MemoryPersistentMemoryLocalSecurityAllOf) GetIsSecurePassphraseSet() bool`

GetIsSecurePassphraseSet returns the IsSecurePassphraseSet field if non-nil, zero value otherwise.

### GetIsSecurePassphraseSetOk

`func (o *MemoryPersistentMemoryLocalSecurityAllOf) GetIsSecurePassphraseSetOk() (*bool, bool)`

GetIsSecurePassphraseSetOk returns a tuple with the IsSecurePassphraseSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSecurePassphraseSet

`func (o *MemoryPersistentMemoryLocalSecurityAllOf) SetIsSecurePassphraseSet(v bool)`

SetIsSecurePassphraseSet sets IsSecurePassphraseSet field to given value.

### HasIsSecurePassphraseSet

`func (o *MemoryPersistentMemoryLocalSecurityAllOf) HasIsSecurePassphraseSet() bool`

HasIsSecurePassphraseSet returns a boolean if a field has been set.

### GetSecurePassphrase

`func (o *MemoryPersistentMemoryLocalSecurityAllOf) GetSecurePassphrase() string`

GetSecurePassphrase returns the SecurePassphrase field if non-nil, zero value otherwise.

### GetSecurePassphraseOk

`func (o *MemoryPersistentMemoryLocalSecurityAllOf) GetSecurePassphraseOk() (*string, bool)`

GetSecurePassphraseOk returns a tuple with the SecurePassphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurePassphrase

`func (o *MemoryPersistentMemoryLocalSecurityAllOf) SetSecurePassphrase(v string)`

SetSecurePassphrase sets SecurePassphrase field to given value.

### HasSecurePassphrase

`func (o *MemoryPersistentMemoryLocalSecurityAllOf) HasSecurePassphrase() bool`

HasSecurePassphrase returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



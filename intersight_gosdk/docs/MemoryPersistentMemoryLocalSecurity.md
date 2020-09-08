# MemoryPersistentMemoryLocalSecurity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Persistent Memory Security state. | [optional] 
**IsSecurePassphraseSet** | Pointer to **bool** | Indicates whether the value of the &#39;securePassphrase&#39; property has been set. | [optional] [readonly] 
**SecurePassphrase** | Pointer to **string** | Secure passphrase to be applied on the Persistent Memory Modules on the server. The allowed characters are a-z, A to Z, 0-9, and special characters &#x3D;, \\u0021, &amp;, \\#, $, %, +, ^, @, _, *, -. | [optional] 

## Methods

### NewMemoryPersistentMemoryLocalSecurity

`func NewMemoryPersistentMemoryLocalSecurity() *MemoryPersistentMemoryLocalSecurity`

NewMemoryPersistentMemoryLocalSecurity instantiates a new MemoryPersistentMemoryLocalSecurity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemoryPersistentMemoryLocalSecurityWithDefaults

`func NewMemoryPersistentMemoryLocalSecurityWithDefaults() *MemoryPersistentMemoryLocalSecurity`

NewMemoryPersistentMemoryLocalSecurityWithDefaults instantiates a new MemoryPersistentMemoryLocalSecurity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *MemoryPersistentMemoryLocalSecurity) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *MemoryPersistentMemoryLocalSecurity) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *MemoryPersistentMemoryLocalSecurity) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *MemoryPersistentMemoryLocalSecurity) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetIsSecurePassphraseSet

`func (o *MemoryPersistentMemoryLocalSecurity) GetIsSecurePassphraseSet() bool`

GetIsSecurePassphraseSet returns the IsSecurePassphraseSet field if non-nil, zero value otherwise.

### GetIsSecurePassphraseSetOk

`func (o *MemoryPersistentMemoryLocalSecurity) GetIsSecurePassphraseSetOk() (*bool, bool)`

GetIsSecurePassphraseSetOk returns a tuple with the IsSecurePassphraseSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSecurePassphraseSet

`func (o *MemoryPersistentMemoryLocalSecurity) SetIsSecurePassphraseSet(v bool)`

SetIsSecurePassphraseSet sets IsSecurePassphraseSet field to given value.

### HasIsSecurePassphraseSet

`func (o *MemoryPersistentMemoryLocalSecurity) HasIsSecurePassphraseSet() bool`

HasIsSecurePassphraseSet returns a boolean if a field has been set.

### GetSecurePassphrase

`func (o *MemoryPersistentMemoryLocalSecurity) GetSecurePassphrase() string`

GetSecurePassphrase returns the SecurePassphrase field if non-nil, zero value otherwise.

### GetSecurePassphraseOk

`func (o *MemoryPersistentMemoryLocalSecurity) GetSecurePassphraseOk() (*string, bool)`

GetSecurePassphraseOk returns a tuple with the SecurePassphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurePassphrase

`func (o *MemoryPersistentMemoryLocalSecurity) SetSecurePassphrase(v string)`

SetSecurePassphrase sets SecurePassphrase field to given value.

### HasSecurePassphrase

`func (o *MemoryPersistentMemoryLocalSecurity) HasSecurePassphrase() bool`

HasSecurePassphrase returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



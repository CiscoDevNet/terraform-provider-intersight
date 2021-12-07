# AssetPrivateKeyCredential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "asset.PrivateKeyCredential"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "asset.PrivateKeyCredential"]
**IsPassphraseSet** | Pointer to **bool** | Indicates whether the value of the &#39;passphrase&#39; property has been set. | [optional] [readonly] [default to false]
**IsPrivateKeySet** | Pointer to **bool** | Indicates whether the value of the &#39;privateKey&#39; property has been set. | [optional] [readonly] [default to false]
**Passphrase** | Pointer to **string** | The passphrase associated with the private key - Optional. | [optional] 
**PrivateKey** | Pointer to **string** | The private key used to authenticate with a managed target. The corresponding public key needs to be added in the auth list of the remote endpoint. | [optional] 
**Username** | Pointer to **string** | The username used to authenticate with a managed target. | [optional] 

## Methods

### NewAssetPrivateKeyCredential

`func NewAssetPrivateKeyCredential(classId string, objectType string, ) *AssetPrivateKeyCredential`

NewAssetPrivateKeyCredential instantiates a new AssetPrivateKeyCredential object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetPrivateKeyCredentialWithDefaults

`func NewAssetPrivateKeyCredentialWithDefaults() *AssetPrivateKeyCredential`

NewAssetPrivateKeyCredentialWithDefaults instantiates a new AssetPrivateKeyCredential object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AssetPrivateKeyCredential) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AssetPrivateKeyCredential) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AssetPrivateKeyCredential) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AssetPrivateKeyCredential) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AssetPrivateKeyCredential) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AssetPrivateKeyCredential) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIsPassphraseSet

`func (o *AssetPrivateKeyCredential) GetIsPassphraseSet() bool`

GetIsPassphraseSet returns the IsPassphraseSet field if non-nil, zero value otherwise.

### GetIsPassphraseSetOk

`func (o *AssetPrivateKeyCredential) GetIsPassphraseSetOk() (*bool, bool)`

GetIsPassphraseSetOk returns a tuple with the IsPassphraseSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPassphraseSet

`func (o *AssetPrivateKeyCredential) SetIsPassphraseSet(v bool)`

SetIsPassphraseSet sets IsPassphraseSet field to given value.

### HasIsPassphraseSet

`func (o *AssetPrivateKeyCredential) HasIsPassphraseSet() bool`

HasIsPassphraseSet returns a boolean if a field has been set.

### GetIsPrivateKeySet

`func (o *AssetPrivateKeyCredential) GetIsPrivateKeySet() bool`

GetIsPrivateKeySet returns the IsPrivateKeySet field if non-nil, zero value otherwise.

### GetIsPrivateKeySetOk

`func (o *AssetPrivateKeyCredential) GetIsPrivateKeySetOk() (*bool, bool)`

GetIsPrivateKeySetOk returns a tuple with the IsPrivateKeySet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrivateKeySet

`func (o *AssetPrivateKeyCredential) SetIsPrivateKeySet(v bool)`

SetIsPrivateKeySet sets IsPrivateKeySet field to given value.

### HasIsPrivateKeySet

`func (o *AssetPrivateKeyCredential) HasIsPrivateKeySet() bool`

HasIsPrivateKeySet returns a boolean if a field has been set.

### GetPassphrase

`func (o *AssetPrivateKeyCredential) GetPassphrase() string`

GetPassphrase returns the Passphrase field if non-nil, zero value otherwise.

### GetPassphraseOk

`func (o *AssetPrivateKeyCredential) GetPassphraseOk() (*string, bool)`

GetPassphraseOk returns a tuple with the Passphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphrase

`func (o *AssetPrivateKeyCredential) SetPassphrase(v string)`

SetPassphrase sets Passphrase field to given value.

### HasPassphrase

`func (o *AssetPrivateKeyCredential) HasPassphrase() bool`

HasPassphrase returns a boolean if a field has been set.

### GetPrivateKey

`func (o *AssetPrivateKeyCredential) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *AssetPrivateKeyCredential) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *AssetPrivateKeyCredential) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *AssetPrivateKeyCredential) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### GetUsername

`func (o *AssetPrivateKeyCredential) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *AssetPrivateKeyCredential) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *AssetPrivateKeyCredential) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *AssetPrivateKeyCredential) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



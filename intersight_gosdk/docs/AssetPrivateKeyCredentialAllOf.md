# AssetPrivateKeyCredentialAllOf

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

### NewAssetPrivateKeyCredentialAllOf

`func NewAssetPrivateKeyCredentialAllOf(classId string, objectType string, ) *AssetPrivateKeyCredentialAllOf`

NewAssetPrivateKeyCredentialAllOf instantiates a new AssetPrivateKeyCredentialAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetPrivateKeyCredentialAllOfWithDefaults

`func NewAssetPrivateKeyCredentialAllOfWithDefaults() *AssetPrivateKeyCredentialAllOf`

NewAssetPrivateKeyCredentialAllOfWithDefaults instantiates a new AssetPrivateKeyCredentialAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AssetPrivateKeyCredentialAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AssetPrivateKeyCredentialAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AssetPrivateKeyCredentialAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AssetPrivateKeyCredentialAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AssetPrivateKeyCredentialAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AssetPrivateKeyCredentialAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIsPassphraseSet

`func (o *AssetPrivateKeyCredentialAllOf) GetIsPassphraseSet() bool`

GetIsPassphraseSet returns the IsPassphraseSet field if non-nil, zero value otherwise.

### GetIsPassphraseSetOk

`func (o *AssetPrivateKeyCredentialAllOf) GetIsPassphraseSetOk() (*bool, bool)`

GetIsPassphraseSetOk returns a tuple with the IsPassphraseSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPassphraseSet

`func (o *AssetPrivateKeyCredentialAllOf) SetIsPassphraseSet(v bool)`

SetIsPassphraseSet sets IsPassphraseSet field to given value.

### HasIsPassphraseSet

`func (o *AssetPrivateKeyCredentialAllOf) HasIsPassphraseSet() bool`

HasIsPassphraseSet returns a boolean if a field has been set.

### GetIsPrivateKeySet

`func (o *AssetPrivateKeyCredentialAllOf) GetIsPrivateKeySet() bool`

GetIsPrivateKeySet returns the IsPrivateKeySet field if non-nil, zero value otherwise.

### GetIsPrivateKeySetOk

`func (o *AssetPrivateKeyCredentialAllOf) GetIsPrivateKeySetOk() (*bool, bool)`

GetIsPrivateKeySetOk returns a tuple with the IsPrivateKeySet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrivateKeySet

`func (o *AssetPrivateKeyCredentialAllOf) SetIsPrivateKeySet(v bool)`

SetIsPrivateKeySet sets IsPrivateKeySet field to given value.

### HasIsPrivateKeySet

`func (o *AssetPrivateKeyCredentialAllOf) HasIsPrivateKeySet() bool`

HasIsPrivateKeySet returns a boolean if a field has been set.

### GetPassphrase

`func (o *AssetPrivateKeyCredentialAllOf) GetPassphrase() string`

GetPassphrase returns the Passphrase field if non-nil, zero value otherwise.

### GetPassphraseOk

`func (o *AssetPrivateKeyCredentialAllOf) GetPassphraseOk() (*string, bool)`

GetPassphraseOk returns a tuple with the Passphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphrase

`func (o *AssetPrivateKeyCredentialAllOf) SetPassphrase(v string)`

SetPassphrase sets Passphrase field to given value.

### HasPassphrase

`func (o *AssetPrivateKeyCredentialAllOf) HasPassphrase() bool`

HasPassphrase returns a boolean if a field has been set.

### GetPrivateKey

`func (o *AssetPrivateKeyCredentialAllOf) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *AssetPrivateKeyCredentialAllOf) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *AssetPrivateKeyCredentialAllOf) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *AssetPrivateKeyCredentialAllOf) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### GetUsername

`func (o *AssetPrivateKeyCredentialAllOf) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *AssetPrivateKeyCredentialAllOf) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *AssetPrivateKeyCredentialAllOf) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *AssetPrivateKeyCredentialAllOf) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



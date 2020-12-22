# AssetOauthClientIdSecretCredential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "asset.OauthClientIdSecretCredential"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "asset.OauthClientIdSecretCredential"]
**ClientId** | Pointer to **string** | The client ID used to authenticate with a managed target. | [optional] 
**ClientSecret** | Pointer to **string** | The client secret used to authenticate with a managed target. | [optional] 
**IsClientSecretSet** | Pointer to **bool** | Indicates whether the value of the &#39;clientSecret&#39; property has been set. | [optional] [readonly] [default to false]

## Methods

### NewAssetOauthClientIdSecretCredential

`func NewAssetOauthClientIdSecretCredential(classId string, objectType string, ) *AssetOauthClientIdSecretCredential`

NewAssetOauthClientIdSecretCredential instantiates a new AssetOauthClientIdSecretCredential object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetOauthClientIdSecretCredentialWithDefaults

`func NewAssetOauthClientIdSecretCredentialWithDefaults() *AssetOauthClientIdSecretCredential`

NewAssetOauthClientIdSecretCredentialWithDefaults instantiates a new AssetOauthClientIdSecretCredential object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AssetOauthClientIdSecretCredential) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AssetOauthClientIdSecretCredential) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AssetOauthClientIdSecretCredential) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AssetOauthClientIdSecretCredential) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AssetOauthClientIdSecretCredential) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AssetOauthClientIdSecretCredential) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetClientId

`func (o *AssetOauthClientIdSecretCredential) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *AssetOauthClientIdSecretCredential) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *AssetOauthClientIdSecretCredential) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *AssetOauthClientIdSecretCredential) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *AssetOauthClientIdSecretCredential) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *AssetOauthClientIdSecretCredential) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *AssetOauthClientIdSecretCredential) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *AssetOauthClientIdSecretCredential) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetIsClientSecretSet

`func (o *AssetOauthClientIdSecretCredential) GetIsClientSecretSet() bool`

GetIsClientSecretSet returns the IsClientSecretSet field if non-nil, zero value otherwise.

### GetIsClientSecretSetOk

`func (o *AssetOauthClientIdSecretCredential) GetIsClientSecretSetOk() (*bool, bool)`

GetIsClientSecretSetOk returns a tuple with the IsClientSecretSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsClientSecretSet

`func (o *AssetOauthClientIdSecretCredential) SetIsClientSecretSet(v bool)`

SetIsClientSecretSet sets IsClientSecretSet field to given value.

### HasIsClientSecretSet

`func (o *AssetOauthClientIdSecretCredential) HasIsClientSecretSet() bool`

HasIsClientSecretSet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



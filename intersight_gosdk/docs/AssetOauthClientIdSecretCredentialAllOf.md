# AssetOauthClientIdSecretCredentialAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "asset.OauthClientIdSecretCredential"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "asset.OauthClientIdSecretCredential"]
**ClientId** | Pointer to **string** | The client ID used to authenticate with a managed target. | [optional] 
**ClientSecret** | Pointer to **string** | The client secret used to authenticate with a managed target. | [optional] 
**IsClientSecretSet** | Pointer to **bool** | Indicates whether the value of the &#39;clientSecret&#39; property has been set. | [optional] [readonly] [default to false]

## Methods

### NewAssetOauthClientIdSecretCredentialAllOf

`func NewAssetOauthClientIdSecretCredentialAllOf(classId string, objectType string, ) *AssetOauthClientIdSecretCredentialAllOf`

NewAssetOauthClientIdSecretCredentialAllOf instantiates a new AssetOauthClientIdSecretCredentialAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetOauthClientIdSecretCredentialAllOfWithDefaults

`func NewAssetOauthClientIdSecretCredentialAllOfWithDefaults() *AssetOauthClientIdSecretCredentialAllOf`

NewAssetOauthClientIdSecretCredentialAllOfWithDefaults instantiates a new AssetOauthClientIdSecretCredentialAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AssetOauthClientIdSecretCredentialAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AssetOauthClientIdSecretCredentialAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AssetOauthClientIdSecretCredentialAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AssetOauthClientIdSecretCredentialAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AssetOauthClientIdSecretCredentialAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AssetOauthClientIdSecretCredentialAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetClientId

`func (o *AssetOauthClientIdSecretCredentialAllOf) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *AssetOauthClientIdSecretCredentialAllOf) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *AssetOauthClientIdSecretCredentialAllOf) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *AssetOauthClientIdSecretCredentialAllOf) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *AssetOauthClientIdSecretCredentialAllOf) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *AssetOauthClientIdSecretCredentialAllOf) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *AssetOauthClientIdSecretCredentialAllOf) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *AssetOauthClientIdSecretCredentialAllOf) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetIsClientSecretSet

`func (o *AssetOauthClientIdSecretCredentialAllOf) GetIsClientSecretSet() bool`

GetIsClientSecretSet returns the IsClientSecretSet field if non-nil, zero value otherwise.

### GetIsClientSecretSetOk

`func (o *AssetOauthClientIdSecretCredentialAllOf) GetIsClientSecretSetOk() (*bool, bool)`

GetIsClientSecretSetOk returns a tuple with the IsClientSecretSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsClientSecretSet

`func (o *AssetOauthClientIdSecretCredentialAllOf) SetIsClientSecretSet(v bool)`

SetIsClientSecretSet sets IsClientSecretSet field to given value.

### HasIsClientSecretSet

`func (o *AssetOauthClientIdSecretCredentialAllOf) HasIsClientSecretSet() bool`

HasIsClientSecretSet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# AssetServiceNowCredential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "asset.ServiceNowCredential"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "asset.ServiceNowCredential"]
**OauthAuthenticationCredential** | Pointer to [**AssetOauthClientIdSecretCredential**](AssetOauthClientIdSecretCredential.md) |  | [optional] 
**UserPasswordCredential** | Pointer to [**AssetUsernamePasswordCredential**](AssetUsernamePasswordCredential.md) |  | [optional] 

## Methods

### NewAssetServiceNowCredential

`func NewAssetServiceNowCredential(classId string, objectType string, ) *AssetServiceNowCredential`

NewAssetServiceNowCredential instantiates a new AssetServiceNowCredential object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetServiceNowCredentialWithDefaults

`func NewAssetServiceNowCredentialWithDefaults() *AssetServiceNowCredential`

NewAssetServiceNowCredentialWithDefaults instantiates a new AssetServiceNowCredential object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AssetServiceNowCredential) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AssetServiceNowCredential) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AssetServiceNowCredential) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AssetServiceNowCredential) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AssetServiceNowCredential) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AssetServiceNowCredential) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOauthAuthenticationCredential

`func (o *AssetServiceNowCredential) GetOauthAuthenticationCredential() AssetOauthClientIdSecretCredential`

GetOauthAuthenticationCredential returns the OauthAuthenticationCredential field if non-nil, zero value otherwise.

### GetOauthAuthenticationCredentialOk

`func (o *AssetServiceNowCredential) GetOauthAuthenticationCredentialOk() (*AssetOauthClientIdSecretCredential, bool)`

GetOauthAuthenticationCredentialOk returns a tuple with the OauthAuthenticationCredential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthAuthenticationCredential

`func (o *AssetServiceNowCredential) SetOauthAuthenticationCredential(v AssetOauthClientIdSecretCredential)`

SetOauthAuthenticationCredential sets OauthAuthenticationCredential field to given value.

### HasOauthAuthenticationCredential

`func (o *AssetServiceNowCredential) HasOauthAuthenticationCredential() bool`

HasOauthAuthenticationCredential returns a boolean if a field has been set.

### GetUserPasswordCredential

`func (o *AssetServiceNowCredential) GetUserPasswordCredential() AssetUsernamePasswordCredential`

GetUserPasswordCredential returns the UserPasswordCredential field if non-nil, zero value otherwise.

### GetUserPasswordCredentialOk

`func (o *AssetServiceNowCredential) GetUserPasswordCredentialOk() (*AssetUsernamePasswordCredential, bool)`

GetUserPasswordCredentialOk returns a tuple with the UserPasswordCredential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPasswordCredential

`func (o *AssetServiceNowCredential) SetUserPasswordCredential(v AssetUsernamePasswordCredential)`

SetUserPasswordCredential sets UserPasswordCredential field to given value.

### HasUserPasswordCredential

`func (o *AssetServiceNowCredential) HasUserPasswordCredential() bool`

HasUserPasswordCredential returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



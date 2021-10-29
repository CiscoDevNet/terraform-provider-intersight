# AssetServiceNowCredentialAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "asset.ServiceNowCredential"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "asset.ServiceNowCredential"]
**OauthAuthenticationCredential** | Pointer to [**AssetOauthClientIdSecretCredential**](AssetOauthClientIdSecretCredential.md) |  | [optional] 
**UserPasswordCredential** | Pointer to [**AssetUsernamePasswordCredential**](AssetUsernamePasswordCredential.md) |  | [optional] 

## Methods

### NewAssetServiceNowCredentialAllOf

`func NewAssetServiceNowCredentialAllOf(classId string, objectType string, ) *AssetServiceNowCredentialAllOf`

NewAssetServiceNowCredentialAllOf instantiates a new AssetServiceNowCredentialAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetServiceNowCredentialAllOfWithDefaults

`func NewAssetServiceNowCredentialAllOfWithDefaults() *AssetServiceNowCredentialAllOf`

NewAssetServiceNowCredentialAllOfWithDefaults instantiates a new AssetServiceNowCredentialAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AssetServiceNowCredentialAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AssetServiceNowCredentialAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AssetServiceNowCredentialAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AssetServiceNowCredentialAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AssetServiceNowCredentialAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AssetServiceNowCredentialAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOauthAuthenticationCredential

`func (o *AssetServiceNowCredentialAllOf) GetOauthAuthenticationCredential() AssetOauthClientIdSecretCredential`

GetOauthAuthenticationCredential returns the OauthAuthenticationCredential field if non-nil, zero value otherwise.

### GetOauthAuthenticationCredentialOk

`func (o *AssetServiceNowCredentialAllOf) GetOauthAuthenticationCredentialOk() (*AssetOauthClientIdSecretCredential, bool)`

GetOauthAuthenticationCredentialOk returns a tuple with the OauthAuthenticationCredential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthAuthenticationCredential

`func (o *AssetServiceNowCredentialAllOf) SetOauthAuthenticationCredential(v AssetOauthClientIdSecretCredential)`

SetOauthAuthenticationCredential sets OauthAuthenticationCredential field to given value.

### HasOauthAuthenticationCredential

`func (o *AssetServiceNowCredentialAllOf) HasOauthAuthenticationCredential() bool`

HasOauthAuthenticationCredential returns a boolean if a field has been set.

### GetUserPasswordCredential

`func (o *AssetServiceNowCredentialAllOf) GetUserPasswordCredential() AssetUsernamePasswordCredential`

GetUserPasswordCredential returns the UserPasswordCredential field if non-nil, zero value otherwise.

### GetUserPasswordCredentialOk

`func (o *AssetServiceNowCredentialAllOf) GetUserPasswordCredentialOk() (*AssetUsernamePasswordCredential, bool)`

GetUserPasswordCredentialOk returns a tuple with the UserPasswordCredential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPasswordCredential

`func (o *AssetServiceNowCredentialAllOf) SetUserPasswordCredential(v AssetUsernamePasswordCredential)`

SetUserPasswordCredential sets UserPasswordCredential field to given value.

### HasUserPasswordCredential

`func (o *AssetServiceNowCredentialAllOf) HasUserPasswordCredential() bool`

HasUserPasswordCredential returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



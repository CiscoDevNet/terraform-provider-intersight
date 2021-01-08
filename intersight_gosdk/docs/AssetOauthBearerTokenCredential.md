# AssetOauthBearerTokenCredential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "asset.OauthBearerTokenCredential"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "asset.OauthBearerTokenCredential"]
**IsTokenSet** | Pointer to **bool** | Indicates whether the value of the &#39;token&#39; property has been set. | [optional] [readonly] [default to false]
**ScopeType** | Pointer to **string** | Scope type for the crendetial i.e. User, Organization, Team. | [optional] 
**ScopeValue** | Pointer to **string** | Scope value for the credential i.e. username, organization name or team name. | [optional] 
**Token** | Pointer to **string** | The token used to authenticate with a managed target. | [optional] 

## Methods

### NewAssetOauthBearerTokenCredential

`func NewAssetOauthBearerTokenCredential(classId string, objectType string, ) *AssetOauthBearerTokenCredential`

NewAssetOauthBearerTokenCredential instantiates a new AssetOauthBearerTokenCredential object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetOauthBearerTokenCredentialWithDefaults

`func NewAssetOauthBearerTokenCredentialWithDefaults() *AssetOauthBearerTokenCredential`

NewAssetOauthBearerTokenCredentialWithDefaults instantiates a new AssetOauthBearerTokenCredential object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AssetOauthBearerTokenCredential) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AssetOauthBearerTokenCredential) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AssetOauthBearerTokenCredential) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AssetOauthBearerTokenCredential) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AssetOauthBearerTokenCredential) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AssetOauthBearerTokenCredential) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIsTokenSet

`func (o *AssetOauthBearerTokenCredential) GetIsTokenSet() bool`

GetIsTokenSet returns the IsTokenSet field if non-nil, zero value otherwise.

### GetIsTokenSetOk

`func (o *AssetOauthBearerTokenCredential) GetIsTokenSetOk() (*bool, bool)`

GetIsTokenSetOk returns a tuple with the IsTokenSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTokenSet

`func (o *AssetOauthBearerTokenCredential) SetIsTokenSet(v bool)`

SetIsTokenSet sets IsTokenSet field to given value.

### HasIsTokenSet

`func (o *AssetOauthBearerTokenCredential) HasIsTokenSet() bool`

HasIsTokenSet returns a boolean if a field has been set.

### GetScopeType

`func (o *AssetOauthBearerTokenCredential) GetScopeType() string`

GetScopeType returns the ScopeType field if non-nil, zero value otherwise.

### GetScopeTypeOk

`func (o *AssetOauthBearerTokenCredential) GetScopeTypeOk() (*string, bool)`

GetScopeTypeOk returns a tuple with the ScopeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeType

`func (o *AssetOauthBearerTokenCredential) SetScopeType(v string)`

SetScopeType sets ScopeType field to given value.

### HasScopeType

`func (o *AssetOauthBearerTokenCredential) HasScopeType() bool`

HasScopeType returns a boolean if a field has been set.

### GetScopeValue

`func (o *AssetOauthBearerTokenCredential) GetScopeValue() string`

GetScopeValue returns the ScopeValue field if non-nil, zero value otherwise.

### GetScopeValueOk

`func (o *AssetOauthBearerTokenCredential) GetScopeValueOk() (*string, bool)`

GetScopeValueOk returns a tuple with the ScopeValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeValue

`func (o *AssetOauthBearerTokenCredential) SetScopeValue(v string)`

SetScopeValue sets ScopeValue field to given value.

### HasScopeValue

`func (o *AssetOauthBearerTokenCredential) HasScopeValue() bool`

HasScopeValue returns a boolean if a field has been set.

### GetToken

`func (o *AssetOauthBearerTokenCredential) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *AssetOauthBearerTokenCredential) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *AssetOauthBearerTokenCredential) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *AssetOauthBearerTokenCredential) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



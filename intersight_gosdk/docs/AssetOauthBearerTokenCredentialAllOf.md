# AssetOauthBearerTokenCredentialAllOf

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

### NewAssetOauthBearerTokenCredentialAllOf

`func NewAssetOauthBearerTokenCredentialAllOf(classId string, objectType string, ) *AssetOauthBearerTokenCredentialAllOf`

NewAssetOauthBearerTokenCredentialAllOf instantiates a new AssetOauthBearerTokenCredentialAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetOauthBearerTokenCredentialAllOfWithDefaults

`func NewAssetOauthBearerTokenCredentialAllOfWithDefaults() *AssetOauthBearerTokenCredentialAllOf`

NewAssetOauthBearerTokenCredentialAllOfWithDefaults instantiates a new AssetOauthBearerTokenCredentialAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AssetOauthBearerTokenCredentialAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AssetOauthBearerTokenCredentialAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AssetOauthBearerTokenCredentialAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AssetOauthBearerTokenCredentialAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AssetOauthBearerTokenCredentialAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AssetOauthBearerTokenCredentialAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIsTokenSet

`func (o *AssetOauthBearerTokenCredentialAllOf) GetIsTokenSet() bool`

GetIsTokenSet returns the IsTokenSet field if non-nil, zero value otherwise.

### GetIsTokenSetOk

`func (o *AssetOauthBearerTokenCredentialAllOf) GetIsTokenSetOk() (*bool, bool)`

GetIsTokenSetOk returns a tuple with the IsTokenSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTokenSet

`func (o *AssetOauthBearerTokenCredentialAllOf) SetIsTokenSet(v bool)`

SetIsTokenSet sets IsTokenSet field to given value.

### HasIsTokenSet

`func (o *AssetOauthBearerTokenCredentialAllOf) HasIsTokenSet() bool`

HasIsTokenSet returns a boolean if a field has been set.

### GetScopeType

`func (o *AssetOauthBearerTokenCredentialAllOf) GetScopeType() string`

GetScopeType returns the ScopeType field if non-nil, zero value otherwise.

### GetScopeTypeOk

`func (o *AssetOauthBearerTokenCredentialAllOf) GetScopeTypeOk() (*string, bool)`

GetScopeTypeOk returns a tuple with the ScopeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeType

`func (o *AssetOauthBearerTokenCredentialAllOf) SetScopeType(v string)`

SetScopeType sets ScopeType field to given value.

### HasScopeType

`func (o *AssetOauthBearerTokenCredentialAllOf) HasScopeType() bool`

HasScopeType returns a boolean if a field has been set.

### GetScopeValue

`func (o *AssetOauthBearerTokenCredentialAllOf) GetScopeValue() string`

GetScopeValue returns the ScopeValue field if non-nil, zero value otherwise.

### GetScopeValueOk

`func (o *AssetOauthBearerTokenCredentialAllOf) GetScopeValueOk() (*string, bool)`

GetScopeValueOk returns a tuple with the ScopeValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeValue

`func (o *AssetOauthBearerTokenCredentialAllOf) SetScopeValue(v string)`

SetScopeValue sets ScopeValue field to given value.

### HasScopeValue

`func (o *AssetOauthBearerTokenCredentialAllOf) HasScopeValue() bool`

HasScopeValue returns a boolean if a field has been set.

### GetToken

`func (o *AssetOauthBearerTokenCredentialAllOf) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *AssetOauthBearerTokenCredentialAllOf) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *AssetOauthBearerTokenCredentialAllOf) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *AssetOauthBearerTokenCredentialAllOf) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# OauthAuthorization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "oauth.Authorization"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "oauth.Authorization"]
**ApiType** | Pointer to **string** | Type of OAuth Api. For example, Smart-licensing-API. * &#x60;Unknown&#x60; - Unknown is the default API type. * &#x60;SmartLicensing-API&#x60; - Smart licensing API type. | [optional] [default to "Unknown"]
**IsPasswordSet** | Pointer to **bool** | Indicates whether the value of the &#39;password&#39; property has been set. | [optional] [readonly] [default to false]
**IsUserIdSet** | Pointer to **bool** | Indicates whether the value of the &#39;userId&#39; property has been set. | [optional] [readonly] [default to false]
**Password** | Pointer to **string** | The password that will be used by Intersight to create OAuth2 tokens for interacting with the external repository, on the user account&#39;s behalf. | [optional] 
**UserId** | Pointer to **string** | The username that will be used by Intersight to create OAuth2 tokens for interacting with the external repository, on the user account&#39;s behalf. | [optional] 
**AccessToken** | Pointer to [**OauthAccessTokenRelationship**](OauthAccessTokenRelationship.md) |  | [optional] 
**Account** | Pointer to [**IamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 

## Methods

### NewOauthAuthorization

`func NewOauthAuthorization(classId string, objectType string, ) *OauthAuthorization`

NewOauthAuthorization instantiates a new OauthAuthorization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOauthAuthorizationWithDefaults

`func NewOauthAuthorizationWithDefaults() *OauthAuthorization`

NewOauthAuthorizationWithDefaults instantiates a new OauthAuthorization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *OauthAuthorization) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *OauthAuthorization) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *OauthAuthorization) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *OauthAuthorization) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *OauthAuthorization) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *OauthAuthorization) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetApiType

`func (o *OauthAuthorization) GetApiType() string`

GetApiType returns the ApiType field if non-nil, zero value otherwise.

### GetApiTypeOk

`func (o *OauthAuthorization) GetApiTypeOk() (*string, bool)`

GetApiTypeOk returns a tuple with the ApiType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiType

`func (o *OauthAuthorization) SetApiType(v string)`

SetApiType sets ApiType field to given value.

### HasApiType

`func (o *OauthAuthorization) HasApiType() bool`

HasApiType returns a boolean if a field has been set.

### GetIsPasswordSet

`func (o *OauthAuthorization) GetIsPasswordSet() bool`

GetIsPasswordSet returns the IsPasswordSet field if non-nil, zero value otherwise.

### GetIsPasswordSetOk

`func (o *OauthAuthorization) GetIsPasswordSetOk() (*bool, bool)`

GetIsPasswordSetOk returns a tuple with the IsPasswordSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPasswordSet

`func (o *OauthAuthorization) SetIsPasswordSet(v bool)`

SetIsPasswordSet sets IsPasswordSet field to given value.

### HasIsPasswordSet

`func (o *OauthAuthorization) HasIsPasswordSet() bool`

HasIsPasswordSet returns a boolean if a field has been set.

### GetIsUserIdSet

`func (o *OauthAuthorization) GetIsUserIdSet() bool`

GetIsUserIdSet returns the IsUserIdSet field if non-nil, zero value otherwise.

### GetIsUserIdSetOk

`func (o *OauthAuthorization) GetIsUserIdSetOk() (*bool, bool)`

GetIsUserIdSetOk returns a tuple with the IsUserIdSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUserIdSet

`func (o *OauthAuthorization) SetIsUserIdSet(v bool)`

SetIsUserIdSet sets IsUserIdSet field to given value.

### HasIsUserIdSet

`func (o *OauthAuthorization) HasIsUserIdSet() bool`

HasIsUserIdSet returns a boolean if a field has been set.

### GetPassword

`func (o *OauthAuthorization) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *OauthAuthorization) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *OauthAuthorization) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *OauthAuthorization) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetUserId

`func (o *OauthAuthorization) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *OauthAuthorization) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *OauthAuthorization) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *OauthAuthorization) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetAccessToken

`func (o *OauthAuthorization) GetAccessToken() OauthAccessTokenRelationship`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *OauthAuthorization) GetAccessTokenOk() (*OauthAccessTokenRelationship, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *OauthAuthorization) SetAccessToken(v OauthAccessTokenRelationship)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *OauthAuthorization) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### GetAccount

`func (o *OauthAuthorization) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *OauthAuthorization) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *OauthAuthorization) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *OauthAuthorization) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



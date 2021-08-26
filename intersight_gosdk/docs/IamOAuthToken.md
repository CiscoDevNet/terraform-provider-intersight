# IamOAuthToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iam.OAuthToken"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iam.OAuthToken"]
**AccessExpirationTime** | Pointer to **time.Time** | Expiration time for the JWT token to which it can be used for api calls. | [optional] [readonly] 
**ClientId** | Pointer to **string** | The identifier of the registered application to which the token belongs. | [optional] 
**ClientIpAddress** | Pointer to **string** | The user agent IP address from which the auth token is launched. | [optional] [readonly] 
**ClientName** | Pointer to **string** | The name of the registered application to which the token belongs. | [optional] 
**ExpirationTime** | Pointer to **time.Time** | Expiration time for the JWT token to which it can be refreshed. | [optional] [readonly] 
**LastLoginClient** | Pointer to **string** | The client address from which last login is initiated. | [optional] [readonly] 
**LastLoginTime** | Pointer to **time.Time** | The last login time for user. | [optional] [readonly] 
**TokenId** | Pointer to **string** | Token identifier. Not the Access Token itself. | [optional] [readonly] 
**UserMeta** | Pointer to [**NullableIamClientMeta**](IamClientMeta.md) |  | [optional] 
**AppRegistration** | Pointer to [**IamAppRegistrationRelationship**](IamAppRegistrationRelationship.md) |  | [optional] 
**Permission** | Pointer to [**IamPermissionRelationship**](IamPermissionRelationship.md) |  | [optional] 
**User** | Pointer to [**IamUserRelationship**](IamUserRelationship.md) |  | [optional] 

## Methods

### NewIamOAuthToken

`func NewIamOAuthToken(classId string, objectType string, ) *IamOAuthToken`

NewIamOAuthToken instantiates a new IamOAuthToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamOAuthTokenWithDefaults

`func NewIamOAuthTokenWithDefaults() *IamOAuthToken`

NewIamOAuthTokenWithDefaults instantiates a new IamOAuthToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IamOAuthToken) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamOAuthToken) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamOAuthToken) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IamOAuthToken) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamOAuthToken) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamOAuthToken) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAccessExpirationTime

`func (o *IamOAuthToken) GetAccessExpirationTime() time.Time`

GetAccessExpirationTime returns the AccessExpirationTime field if non-nil, zero value otherwise.

### GetAccessExpirationTimeOk

`func (o *IamOAuthToken) GetAccessExpirationTimeOk() (*time.Time, bool)`

GetAccessExpirationTimeOk returns a tuple with the AccessExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessExpirationTime

`func (o *IamOAuthToken) SetAccessExpirationTime(v time.Time)`

SetAccessExpirationTime sets AccessExpirationTime field to given value.

### HasAccessExpirationTime

`func (o *IamOAuthToken) HasAccessExpirationTime() bool`

HasAccessExpirationTime returns a boolean if a field has been set.

### GetClientId

`func (o *IamOAuthToken) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *IamOAuthToken) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *IamOAuthToken) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *IamOAuthToken) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientIpAddress

`func (o *IamOAuthToken) GetClientIpAddress() string`

GetClientIpAddress returns the ClientIpAddress field if non-nil, zero value otherwise.

### GetClientIpAddressOk

`func (o *IamOAuthToken) GetClientIpAddressOk() (*string, bool)`

GetClientIpAddressOk returns a tuple with the ClientIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIpAddress

`func (o *IamOAuthToken) SetClientIpAddress(v string)`

SetClientIpAddress sets ClientIpAddress field to given value.

### HasClientIpAddress

`func (o *IamOAuthToken) HasClientIpAddress() bool`

HasClientIpAddress returns a boolean if a field has been set.

### GetClientName

`func (o *IamOAuthToken) GetClientName() string`

GetClientName returns the ClientName field if non-nil, zero value otherwise.

### GetClientNameOk

`func (o *IamOAuthToken) GetClientNameOk() (*string, bool)`

GetClientNameOk returns a tuple with the ClientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientName

`func (o *IamOAuthToken) SetClientName(v string)`

SetClientName sets ClientName field to given value.

### HasClientName

`func (o *IamOAuthToken) HasClientName() bool`

HasClientName returns a boolean if a field has been set.

### GetExpirationTime

`func (o *IamOAuthToken) GetExpirationTime() time.Time`

GetExpirationTime returns the ExpirationTime field if non-nil, zero value otherwise.

### GetExpirationTimeOk

`func (o *IamOAuthToken) GetExpirationTimeOk() (*time.Time, bool)`

GetExpirationTimeOk returns a tuple with the ExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTime

`func (o *IamOAuthToken) SetExpirationTime(v time.Time)`

SetExpirationTime sets ExpirationTime field to given value.

### HasExpirationTime

`func (o *IamOAuthToken) HasExpirationTime() bool`

HasExpirationTime returns a boolean if a field has been set.

### GetLastLoginClient

`func (o *IamOAuthToken) GetLastLoginClient() string`

GetLastLoginClient returns the LastLoginClient field if non-nil, zero value otherwise.

### GetLastLoginClientOk

`func (o *IamOAuthToken) GetLastLoginClientOk() (*string, bool)`

GetLastLoginClientOk returns a tuple with the LastLoginClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLoginClient

`func (o *IamOAuthToken) SetLastLoginClient(v string)`

SetLastLoginClient sets LastLoginClient field to given value.

### HasLastLoginClient

`func (o *IamOAuthToken) HasLastLoginClient() bool`

HasLastLoginClient returns a boolean if a field has been set.

### GetLastLoginTime

`func (o *IamOAuthToken) GetLastLoginTime() time.Time`

GetLastLoginTime returns the LastLoginTime field if non-nil, zero value otherwise.

### GetLastLoginTimeOk

`func (o *IamOAuthToken) GetLastLoginTimeOk() (*time.Time, bool)`

GetLastLoginTimeOk returns a tuple with the LastLoginTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLoginTime

`func (o *IamOAuthToken) SetLastLoginTime(v time.Time)`

SetLastLoginTime sets LastLoginTime field to given value.

### HasLastLoginTime

`func (o *IamOAuthToken) HasLastLoginTime() bool`

HasLastLoginTime returns a boolean if a field has been set.

### GetTokenId

`func (o *IamOAuthToken) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *IamOAuthToken) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *IamOAuthToken) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *IamOAuthToken) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.

### GetUserMeta

`func (o *IamOAuthToken) GetUserMeta() IamClientMeta`

GetUserMeta returns the UserMeta field if non-nil, zero value otherwise.

### GetUserMetaOk

`func (o *IamOAuthToken) GetUserMetaOk() (*IamClientMeta, bool)`

GetUserMetaOk returns a tuple with the UserMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMeta

`func (o *IamOAuthToken) SetUserMeta(v IamClientMeta)`

SetUserMeta sets UserMeta field to given value.

### HasUserMeta

`func (o *IamOAuthToken) HasUserMeta() bool`

HasUserMeta returns a boolean if a field has been set.

### SetUserMetaNil

`func (o *IamOAuthToken) SetUserMetaNil(b bool)`

 SetUserMetaNil sets the value for UserMeta to be an explicit nil

### UnsetUserMeta
`func (o *IamOAuthToken) UnsetUserMeta()`

UnsetUserMeta ensures that no value is present for UserMeta, not even an explicit nil
### GetAppRegistration

`func (o *IamOAuthToken) GetAppRegistration() IamAppRegistrationRelationship`

GetAppRegistration returns the AppRegistration field if non-nil, zero value otherwise.

### GetAppRegistrationOk

`func (o *IamOAuthToken) GetAppRegistrationOk() (*IamAppRegistrationRelationship, bool)`

GetAppRegistrationOk returns a tuple with the AppRegistration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppRegistration

`func (o *IamOAuthToken) SetAppRegistration(v IamAppRegistrationRelationship)`

SetAppRegistration sets AppRegistration field to given value.

### HasAppRegistration

`func (o *IamOAuthToken) HasAppRegistration() bool`

HasAppRegistration returns a boolean if a field has been set.

### GetPermission

`func (o *IamOAuthToken) GetPermission() IamPermissionRelationship`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *IamOAuthToken) GetPermissionOk() (*IamPermissionRelationship, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *IamOAuthToken) SetPermission(v IamPermissionRelationship)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *IamOAuthToken) HasPermission() bool`

HasPermission returns a boolean if a field has been set.

### GetUser

`func (o *IamOAuthToken) GetUser() IamUserRelationship`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *IamOAuthToken) GetUserOk() (*IamUserRelationship, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *IamOAuthToken) SetUser(v IamUserRelationship)`

SetUser sets User field to given value.

### HasUser

`func (o *IamOAuthToken) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



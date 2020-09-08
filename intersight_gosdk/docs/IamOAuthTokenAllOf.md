# IamOAuthTokenAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessExpirationTime** | Pointer to [**time.Time**](time.Time.md) | Expiration time for the JWT token to which it can be used for api calls. | [optional] [readonly] 
**ClientId** | Pointer to **string** | The identifier of the registered application to which the token belongs. | [optional] 
**ClientIpAddress** | Pointer to **string** | The user agent IP address from which the auth token is launched. | [optional] [readonly] 
**ClientName** | Pointer to **string** | The name of the registered application to which the token belongs. | [optional] 
**ExpirationTime** | Pointer to [**time.Time**](time.Time.md) | Expiration time for the JWT token to which it can be refreshed. | [optional] [readonly] 
**LastLoginClient** | Pointer to **string** | The client address from which last login is initiated. | [optional] [readonly] 
**LastLoginTime** | Pointer to [**time.Time**](time.Time.md) | The last login time for user. | [optional] [readonly] 
**TokenId** | Pointer to **string** | Token identifier. Not the Access Token itself. | [optional] [readonly] 
**UserMeta** | Pointer to [**IamClientMeta**](iam.ClientMeta.md) |  | [optional] 
**AppRegistration** | Pointer to [**IamAppRegistrationRelationship**](iam.AppRegistration.Relationship.md) |  | [optional] 
**Permission** | Pointer to [**IamPermissionRelationship**](iam.Permission.Relationship.md) |  | [optional] 
**User** | Pointer to [**IamUserRelationship**](iam.User.Relationship.md) |  | [optional] 

## Methods

### NewIamOAuthTokenAllOf

`func NewIamOAuthTokenAllOf() *IamOAuthTokenAllOf`

NewIamOAuthTokenAllOf instantiates a new IamOAuthTokenAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamOAuthTokenAllOfWithDefaults

`func NewIamOAuthTokenAllOfWithDefaults() *IamOAuthTokenAllOf`

NewIamOAuthTokenAllOfWithDefaults instantiates a new IamOAuthTokenAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessExpirationTime

`func (o *IamOAuthTokenAllOf) GetAccessExpirationTime() time.Time`

GetAccessExpirationTime returns the AccessExpirationTime field if non-nil, zero value otherwise.

### GetAccessExpirationTimeOk

`func (o *IamOAuthTokenAllOf) GetAccessExpirationTimeOk() (*time.Time, bool)`

GetAccessExpirationTimeOk returns a tuple with the AccessExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessExpirationTime

`func (o *IamOAuthTokenAllOf) SetAccessExpirationTime(v time.Time)`

SetAccessExpirationTime sets AccessExpirationTime field to given value.

### HasAccessExpirationTime

`func (o *IamOAuthTokenAllOf) HasAccessExpirationTime() bool`

HasAccessExpirationTime returns a boolean if a field has been set.

### GetClientId

`func (o *IamOAuthTokenAllOf) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *IamOAuthTokenAllOf) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *IamOAuthTokenAllOf) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *IamOAuthTokenAllOf) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientIpAddress

`func (o *IamOAuthTokenAllOf) GetClientIpAddress() string`

GetClientIpAddress returns the ClientIpAddress field if non-nil, zero value otherwise.

### GetClientIpAddressOk

`func (o *IamOAuthTokenAllOf) GetClientIpAddressOk() (*string, bool)`

GetClientIpAddressOk returns a tuple with the ClientIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIpAddress

`func (o *IamOAuthTokenAllOf) SetClientIpAddress(v string)`

SetClientIpAddress sets ClientIpAddress field to given value.

### HasClientIpAddress

`func (o *IamOAuthTokenAllOf) HasClientIpAddress() bool`

HasClientIpAddress returns a boolean if a field has been set.

### GetClientName

`func (o *IamOAuthTokenAllOf) GetClientName() string`

GetClientName returns the ClientName field if non-nil, zero value otherwise.

### GetClientNameOk

`func (o *IamOAuthTokenAllOf) GetClientNameOk() (*string, bool)`

GetClientNameOk returns a tuple with the ClientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientName

`func (o *IamOAuthTokenAllOf) SetClientName(v string)`

SetClientName sets ClientName field to given value.

### HasClientName

`func (o *IamOAuthTokenAllOf) HasClientName() bool`

HasClientName returns a boolean if a field has been set.

### GetExpirationTime

`func (o *IamOAuthTokenAllOf) GetExpirationTime() time.Time`

GetExpirationTime returns the ExpirationTime field if non-nil, zero value otherwise.

### GetExpirationTimeOk

`func (o *IamOAuthTokenAllOf) GetExpirationTimeOk() (*time.Time, bool)`

GetExpirationTimeOk returns a tuple with the ExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTime

`func (o *IamOAuthTokenAllOf) SetExpirationTime(v time.Time)`

SetExpirationTime sets ExpirationTime field to given value.

### HasExpirationTime

`func (o *IamOAuthTokenAllOf) HasExpirationTime() bool`

HasExpirationTime returns a boolean if a field has been set.

### GetLastLoginClient

`func (o *IamOAuthTokenAllOf) GetLastLoginClient() string`

GetLastLoginClient returns the LastLoginClient field if non-nil, zero value otherwise.

### GetLastLoginClientOk

`func (o *IamOAuthTokenAllOf) GetLastLoginClientOk() (*string, bool)`

GetLastLoginClientOk returns a tuple with the LastLoginClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLoginClient

`func (o *IamOAuthTokenAllOf) SetLastLoginClient(v string)`

SetLastLoginClient sets LastLoginClient field to given value.

### HasLastLoginClient

`func (o *IamOAuthTokenAllOf) HasLastLoginClient() bool`

HasLastLoginClient returns a boolean if a field has been set.

### GetLastLoginTime

`func (o *IamOAuthTokenAllOf) GetLastLoginTime() time.Time`

GetLastLoginTime returns the LastLoginTime field if non-nil, zero value otherwise.

### GetLastLoginTimeOk

`func (o *IamOAuthTokenAllOf) GetLastLoginTimeOk() (*time.Time, bool)`

GetLastLoginTimeOk returns a tuple with the LastLoginTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLoginTime

`func (o *IamOAuthTokenAllOf) SetLastLoginTime(v time.Time)`

SetLastLoginTime sets LastLoginTime field to given value.

### HasLastLoginTime

`func (o *IamOAuthTokenAllOf) HasLastLoginTime() bool`

HasLastLoginTime returns a boolean if a field has been set.

### GetTokenId

`func (o *IamOAuthTokenAllOf) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *IamOAuthTokenAllOf) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *IamOAuthTokenAllOf) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *IamOAuthTokenAllOf) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.

### GetUserMeta

`func (o *IamOAuthTokenAllOf) GetUserMeta() IamClientMeta`

GetUserMeta returns the UserMeta field if non-nil, zero value otherwise.

### GetUserMetaOk

`func (o *IamOAuthTokenAllOf) GetUserMetaOk() (*IamClientMeta, bool)`

GetUserMetaOk returns a tuple with the UserMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMeta

`func (o *IamOAuthTokenAllOf) SetUserMeta(v IamClientMeta)`

SetUserMeta sets UserMeta field to given value.

### HasUserMeta

`func (o *IamOAuthTokenAllOf) HasUserMeta() bool`

HasUserMeta returns a boolean if a field has been set.

### GetAppRegistration

`func (o *IamOAuthTokenAllOf) GetAppRegistration() IamAppRegistrationRelationship`

GetAppRegistration returns the AppRegistration field if non-nil, zero value otherwise.

### GetAppRegistrationOk

`func (o *IamOAuthTokenAllOf) GetAppRegistrationOk() (*IamAppRegistrationRelationship, bool)`

GetAppRegistrationOk returns a tuple with the AppRegistration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppRegistration

`func (o *IamOAuthTokenAllOf) SetAppRegistration(v IamAppRegistrationRelationship)`

SetAppRegistration sets AppRegistration field to given value.

### HasAppRegistration

`func (o *IamOAuthTokenAllOf) HasAppRegistration() bool`

HasAppRegistration returns a boolean if a field has been set.

### GetPermission

`func (o *IamOAuthTokenAllOf) GetPermission() IamPermissionRelationship`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *IamOAuthTokenAllOf) GetPermissionOk() (*IamPermissionRelationship, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *IamOAuthTokenAllOf) SetPermission(v IamPermissionRelationship)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *IamOAuthTokenAllOf) HasPermission() bool`

HasPermission returns a boolean if a field has been set.

### GetUser

`func (o *IamOAuthTokenAllOf) GetUser() IamUserRelationship`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *IamOAuthTokenAllOf) GetUserOk() (*IamUserRelationship, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *IamOAuthTokenAllOf) SetUser(v IamUserRelationship)`

SetUser sets User field to given value.

### HasUser

`func (o *IamOAuthTokenAllOf) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



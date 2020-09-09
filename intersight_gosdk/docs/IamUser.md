# IamUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientIpAddress** | Pointer to **string** | IP address from which the user last logged in to Intersight. | [optional] [readonly] 
**Email** | Pointer to **string** | Email of the user. Users are added to Intersight using the email configured in the IdP. | [optional] 
**FirstName** | Pointer to **string** | First name of the user. This field is populated from the IdP attributes received after authentication. | [optional] [readonly] 
**LastLoginTime** | Pointer to [**time.Time**](time.Time.md) | Last successful login time for user. | [optional] [readonly] 
**LastName** | Pointer to **string** | Last name of the user. This field is populated from the IdP attributes received after authentication. | [optional] [readonly] 
**Name** | Pointer to **string** | Name as configured in the IdP. | [optional] [readonly] 
**UserIdOrEmail** | Pointer to **string** | UserID or email as configured in the IdP. | [optional] 
**UserType** | Pointer to **string** | Type of the User. If a user is added manually by specifying the email address, or has logged in using groups, based on the IdP attributes received during authentication. If added manually, the user type will be static, otherwise dynamic. | [optional] [readonly] 
**UserUniqueIdentifier** | Pointer to **string** | Unique id of the user used by the identity provider to store the user. | [optional] [readonly] 
**ApiKeys** | Pointer to [**[]IamApiKeyRelationship**](iam.ApiKey.Relationship.md) | An array of relationships to iamApiKey resources. | [optional] [readonly] 
**AppRegistrations** | Pointer to [**[]IamAppRegistrationRelationship**](iam.AppRegistration.Relationship.md) | An array of relationships to iamAppRegistration resources. | [optional] [readonly] 
**Idp** | Pointer to [**IamIdpRelationship**](iam.Idp.Relationship.md) |  | [optional] 
**Idpreference** | Pointer to [**IamIdpReferenceRelationship**](iam.IdpReference.Relationship.md) |  | [optional] 
**LocalUserPassword** | Pointer to [**IamLocalUserPasswordRelationship**](iam.LocalUserPassword.Relationship.md) |  | [optional] 
**OauthTokens** | Pointer to [**[]IamOAuthTokenRelationship**](iam.OAuthToken.Relationship.md) | An array of relationships to iamOAuthToken resources. | [optional] [readonly] 
**Permissions** | Pointer to [**[]IamPermissionRelationship**](iam.Permission.Relationship.md) | An array of relationships to iamPermission resources. | [optional] 
**Sessions** | Pointer to [**[]IamSessionRelationship**](iam.Session.Relationship.md) | An array of relationships to iamSession resources. | [optional] [readonly] 

## Methods

### NewIamUser

`func NewIamUser() *IamUser`

NewIamUser instantiates a new IamUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamUserWithDefaults

`func NewIamUserWithDefaults() *IamUser`

NewIamUserWithDefaults instantiates a new IamUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientIpAddress

`func (o *IamUser) GetClientIpAddress() string`

GetClientIpAddress returns the ClientIpAddress field if non-nil, zero value otherwise.

### GetClientIpAddressOk

`func (o *IamUser) GetClientIpAddressOk() (*string, bool)`

GetClientIpAddressOk returns a tuple with the ClientIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIpAddress

`func (o *IamUser) SetClientIpAddress(v string)`

SetClientIpAddress sets ClientIpAddress field to given value.

### HasClientIpAddress

`func (o *IamUser) HasClientIpAddress() bool`

HasClientIpAddress returns a boolean if a field has been set.

### GetEmail

`func (o *IamUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *IamUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *IamUser) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *IamUser) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetFirstName

`func (o *IamUser) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *IamUser) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *IamUser) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *IamUser) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastLoginTime

`func (o *IamUser) GetLastLoginTime() time.Time`

GetLastLoginTime returns the LastLoginTime field if non-nil, zero value otherwise.

### GetLastLoginTimeOk

`func (o *IamUser) GetLastLoginTimeOk() (*time.Time, bool)`

GetLastLoginTimeOk returns a tuple with the LastLoginTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLoginTime

`func (o *IamUser) SetLastLoginTime(v time.Time)`

SetLastLoginTime sets LastLoginTime field to given value.

### HasLastLoginTime

`func (o *IamUser) HasLastLoginTime() bool`

HasLastLoginTime returns a boolean if a field has been set.

### GetLastName

`func (o *IamUser) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *IamUser) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *IamUser) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *IamUser) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetName

`func (o *IamUser) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IamUser) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IamUser) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IamUser) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserIdOrEmail

`func (o *IamUser) GetUserIdOrEmail() string`

GetUserIdOrEmail returns the UserIdOrEmail field if non-nil, zero value otherwise.

### GetUserIdOrEmailOk

`func (o *IamUser) GetUserIdOrEmailOk() (*string, bool)`

GetUserIdOrEmailOk returns a tuple with the UserIdOrEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIdOrEmail

`func (o *IamUser) SetUserIdOrEmail(v string)`

SetUserIdOrEmail sets UserIdOrEmail field to given value.

### HasUserIdOrEmail

`func (o *IamUser) HasUserIdOrEmail() bool`

HasUserIdOrEmail returns a boolean if a field has been set.

### GetUserType

`func (o *IamUser) GetUserType() string`

GetUserType returns the UserType field if non-nil, zero value otherwise.

### GetUserTypeOk

`func (o *IamUser) GetUserTypeOk() (*string, bool)`

GetUserTypeOk returns a tuple with the UserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserType

`func (o *IamUser) SetUserType(v string)`

SetUserType sets UserType field to given value.

### HasUserType

`func (o *IamUser) HasUserType() bool`

HasUserType returns a boolean if a field has been set.

### GetUserUniqueIdentifier

`func (o *IamUser) GetUserUniqueIdentifier() string`

GetUserUniqueIdentifier returns the UserUniqueIdentifier field if non-nil, zero value otherwise.

### GetUserUniqueIdentifierOk

`func (o *IamUser) GetUserUniqueIdentifierOk() (*string, bool)`

GetUserUniqueIdentifierOk returns a tuple with the UserUniqueIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserUniqueIdentifier

`func (o *IamUser) SetUserUniqueIdentifier(v string)`

SetUserUniqueIdentifier sets UserUniqueIdentifier field to given value.

### HasUserUniqueIdentifier

`func (o *IamUser) HasUserUniqueIdentifier() bool`

HasUserUniqueIdentifier returns a boolean if a field has been set.

### GetApiKeys

`func (o *IamUser) GetApiKeys() []IamApiKeyRelationship`

GetApiKeys returns the ApiKeys field if non-nil, zero value otherwise.

### GetApiKeysOk

`func (o *IamUser) GetApiKeysOk() (*[]IamApiKeyRelationship, bool)`

GetApiKeysOk returns a tuple with the ApiKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeys

`func (o *IamUser) SetApiKeys(v []IamApiKeyRelationship)`

SetApiKeys sets ApiKeys field to given value.

### HasApiKeys

`func (o *IamUser) HasApiKeys() bool`

HasApiKeys returns a boolean if a field has been set.

### SetApiKeysNil

`func (o *IamUser) SetApiKeysNil(b bool)`

 SetApiKeysNil sets the value for ApiKeys to be an explicit nil

### UnsetApiKeys
`func (o *IamUser) UnsetApiKeys()`

UnsetApiKeys ensures that no value is present for ApiKeys, not even an explicit nil
### GetAppRegistrations

`func (o *IamUser) GetAppRegistrations() []IamAppRegistrationRelationship`

GetAppRegistrations returns the AppRegistrations field if non-nil, zero value otherwise.

### GetAppRegistrationsOk

`func (o *IamUser) GetAppRegistrationsOk() (*[]IamAppRegistrationRelationship, bool)`

GetAppRegistrationsOk returns a tuple with the AppRegistrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppRegistrations

`func (o *IamUser) SetAppRegistrations(v []IamAppRegistrationRelationship)`

SetAppRegistrations sets AppRegistrations field to given value.

### HasAppRegistrations

`func (o *IamUser) HasAppRegistrations() bool`

HasAppRegistrations returns a boolean if a field has been set.

### SetAppRegistrationsNil

`func (o *IamUser) SetAppRegistrationsNil(b bool)`

 SetAppRegistrationsNil sets the value for AppRegistrations to be an explicit nil

### UnsetAppRegistrations
`func (o *IamUser) UnsetAppRegistrations()`

UnsetAppRegistrations ensures that no value is present for AppRegistrations, not even an explicit nil
### GetIdp

`func (o *IamUser) GetIdp() IamIdpRelationship`

GetIdp returns the Idp field if non-nil, zero value otherwise.

### GetIdpOk

`func (o *IamUser) GetIdpOk() (*IamIdpRelationship, bool)`

GetIdpOk returns a tuple with the Idp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdp

`func (o *IamUser) SetIdp(v IamIdpRelationship)`

SetIdp sets Idp field to given value.

### HasIdp

`func (o *IamUser) HasIdp() bool`

HasIdp returns a boolean if a field has been set.

### GetIdpreference

`func (o *IamUser) GetIdpreference() IamIdpReferenceRelationship`

GetIdpreference returns the Idpreference field if non-nil, zero value otherwise.

### GetIdpreferenceOk

`func (o *IamUser) GetIdpreferenceOk() (*IamIdpReferenceRelationship, bool)`

GetIdpreferenceOk returns a tuple with the Idpreference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpreference

`func (o *IamUser) SetIdpreference(v IamIdpReferenceRelationship)`

SetIdpreference sets Idpreference field to given value.

### HasIdpreference

`func (o *IamUser) HasIdpreference() bool`

HasIdpreference returns a boolean if a field has been set.

### GetLocalUserPassword

`func (o *IamUser) GetLocalUserPassword() IamLocalUserPasswordRelationship`

GetLocalUserPassword returns the LocalUserPassword field if non-nil, zero value otherwise.

### GetLocalUserPasswordOk

`func (o *IamUser) GetLocalUserPasswordOk() (*IamLocalUserPasswordRelationship, bool)`

GetLocalUserPasswordOk returns a tuple with the LocalUserPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalUserPassword

`func (o *IamUser) SetLocalUserPassword(v IamLocalUserPasswordRelationship)`

SetLocalUserPassword sets LocalUserPassword field to given value.

### HasLocalUserPassword

`func (o *IamUser) HasLocalUserPassword() bool`

HasLocalUserPassword returns a boolean if a field has been set.

### GetOauthTokens

`func (o *IamUser) GetOauthTokens() []IamOAuthTokenRelationship`

GetOauthTokens returns the OauthTokens field if non-nil, zero value otherwise.

### GetOauthTokensOk

`func (o *IamUser) GetOauthTokensOk() (*[]IamOAuthTokenRelationship, bool)`

GetOauthTokensOk returns a tuple with the OauthTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthTokens

`func (o *IamUser) SetOauthTokens(v []IamOAuthTokenRelationship)`

SetOauthTokens sets OauthTokens field to given value.

### HasOauthTokens

`func (o *IamUser) HasOauthTokens() bool`

HasOauthTokens returns a boolean if a field has been set.

### SetOauthTokensNil

`func (o *IamUser) SetOauthTokensNil(b bool)`

 SetOauthTokensNil sets the value for OauthTokens to be an explicit nil

### UnsetOauthTokens
`func (o *IamUser) UnsetOauthTokens()`

UnsetOauthTokens ensures that no value is present for OauthTokens, not even an explicit nil
### GetPermissions

`func (o *IamUser) GetPermissions() []IamPermissionRelationship`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *IamUser) GetPermissionsOk() (*[]IamPermissionRelationship, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *IamUser) SetPermissions(v []IamPermissionRelationship)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *IamUser) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### SetPermissionsNil

`func (o *IamUser) SetPermissionsNil(b bool)`

 SetPermissionsNil sets the value for Permissions to be an explicit nil

### UnsetPermissions
`func (o *IamUser) UnsetPermissions()`

UnsetPermissions ensures that no value is present for Permissions, not even an explicit nil
### GetSessions

`func (o *IamUser) GetSessions() []IamSessionRelationship`

GetSessions returns the Sessions field if non-nil, zero value otherwise.

### GetSessionsOk

`func (o *IamUser) GetSessionsOk() (*[]IamSessionRelationship, bool)`

GetSessionsOk returns a tuple with the Sessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessions

`func (o *IamUser) SetSessions(v []IamSessionRelationship)`

SetSessions sets Sessions field to given value.

### HasSessions

`func (o *IamUser) HasSessions() bool`

HasSessions returns a boolean if a field has been set.

### SetSessionsNil

`func (o *IamUser) SetSessionsNil(b bool)`

 SetSessionsNil sets the value for Sessions to be an explicit nil

### UnsetSessions
`func (o *IamUser) UnsetSessions()`

UnsetSessions ensures that no value is present for Sessions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



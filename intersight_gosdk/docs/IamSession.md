# IamSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountPermissions** | Pointer to [**[]IamAccountPermissions**](iam.AccountPermissions.md) |  | [optional] 
**ClientIpAddress** | Pointer to **string** | The user agent IP address from which the session is launched. | [optional] [readonly] 
**Expiration** | Pointer to [**time.Time**](time.Time.md) | Expiration time for the session. | [optional] [readonly] 
**IdleTimeExpiration** | Pointer to [**time.Time**](time.Time.md) | Idle time expiration for the session. | [optional] [readonly] 
**LastLoginClient** | Pointer to **string** | The client address from which last login is initiated. | [optional] [readonly] 
**LastLoginTime** | Pointer to [**time.Time**](time.Time.md) | The last login time for user. | [optional] [readonly] 
**Permission** | Pointer to [**IamPermissionRelationship**](iam.Permission.Relationship.md) |  | [optional] 
**User** | Pointer to [**IamUserRelationship**](iam.User.Relationship.md) |  | [optional] 

## Methods

### NewIamSession

`func NewIamSession() *IamSession`

NewIamSession instantiates a new IamSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamSessionWithDefaults

`func NewIamSessionWithDefaults() *IamSession`

NewIamSessionWithDefaults instantiates a new IamSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountPermissions

`func (o *IamSession) GetAccountPermissions() []IamAccountPermissions`

GetAccountPermissions returns the AccountPermissions field if non-nil, zero value otherwise.

### GetAccountPermissionsOk

`func (o *IamSession) GetAccountPermissionsOk() (*[]IamAccountPermissions, bool)`

GetAccountPermissionsOk returns a tuple with the AccountPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountPermissions

`func (o *IamSession) SetAccountPermissions(v []IamAccountPermissions)`

SetAccountPermissions sets AccountPermissions field to given value.

### HasAccountPermissions

`func (o *IamSession) HasAccountPermissions() bool`

HasAccountPermissions returns a boolean if a field has been set.

### GetClientIpAddress

`func (o *IamSession) GetClientIpAddress() string`

GetClientIpAddress returns the ClientIpAddress field if non-nil, zero value otherwise.

### GetClientIpAddressOk

`func (o *IamSession) GetClientIpAddressOk() (*string, bool)`

GetClientIpAddressOk returns a tuple with the ClientIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIpAddress

`func (o *IamSession) SetClientIpAddress(v string)`

SetClientIpAddress sets ClientIpAddress field to given value.

### HasClientIpAddress

`func (o *IamSession) HasClientIpAddress() bool`

HasClientIpAddress returns a boolean if a field has been set.

### GetExpiration

`func (o *IamSession) GetExpiration() time.Time`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *IamSession) GetExpirationOk() (*time.Time, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *IamSession) SetExpiration(v time.Time)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *IamSession) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.

### GetIdleTimeExpiration

`func (o *IamSession) GetIdleTimeExpiration() time.Time`

GetIdleTimeExpiration returns the IdleTimeExpiration field if non-nil, zero value otherwise.

### GetIdleTimeExpirationOk

`func (o *IamSession) GetIdleTimeExpirationOk() (*time.Time, bool)`

GetIdleTimeExpirationOk returns a tuple with the IdleTimeExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleTimeExpiration

`func (o *IamSession) SetIdleTimeExpiration(v time.Time)`

SetIdleTimeExpiration sets IdleTimeExpiration field to given value.

### HasIdleTimeExpiration

`func (o *IamSession) HasIdleTimeExpiration() bool`

HasIdleTimeExpiration returns a boolean if a field has been set.

### GetLastLoginClient

`func (o *IamSession) GetLastLoginClient() string`

GetLastLoginClient returns the LastLoginClient field if non-nil, zero value otherwise.

### GetLastLoginClientOk

`func (o *IamSession) GetLastLoginClientOk() (*string, bool)`

GetLastLoginClientOk returns a tuple with the LastLoginClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLoginClient

`func (o *IamSession) SetLastLoginClient(v string)`

SetLastLoginClient sets LastLoginClient field to given value.

### HasLastLoginClient

`func (o *IamSession) HasLastLoginClient() bool`

HasLastLoginClient returns a boolean if a field has been set.

### GetLastLoginTime

`func (o *IamSession) GetLastLoginTime() time.Time`

GetLastLoginTime returns the LastLoginTime field if non-nil, zero value otherwise.

### GetLastLoginTimeOk

`func (o *IamSession) GetLastLoginTimeOk() (*time.Time, bool)`

GetLastLoginTimeOk returns a tuple with the LastLoginTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLoginTime

`func (o *IamSession) SetLastLoginTime(v time.Time)`

SetLastLoginTime sets LastLoginTime field to given value.

### HasLastLoginTime

`func (o *IamSession) HasLastLoginTime() bool`

HasLastLoginTime returns a boolean if a field has been set.

### GetPermission

`func (o *IamSession) GetPermission() IamPermissionRelationship`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *IamSession) GetPermissionOk() (*IamPermissionRelationship, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *IamSession) SetPermission(v IamPermissionRelationship)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *IamSession) HasPermission() bool`

HasPermission returns a boolean if a field has been set.

### GetUser

`func (o *IamSession) GetUser() IamUserRelationship`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *IamSession) GetUserOk() (*IamUserRelationship, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *IamSession) SetUser(v IamUserRelationship)`

SetUser sets User field to given value.

### HasUser

`func (o *IamSession) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



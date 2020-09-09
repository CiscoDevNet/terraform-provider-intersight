# IamSessionAllOf

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

### NewIamSessionAllOf

`func NewIamSessionAllOf() *IamSessionAllOf`

NewIamSessionAllOf instantiates a new IamSessionAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamSessionAllOfWithDefaults

`func NewIamSessionAllOfWithDefaults() *IamSessionAllOf`

NewIamSessionAllOfWithDefaults instantiates a new IamSessionAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountPermissions

`func (o *IamSessionAllOf) GetAccountPermissions() []IamAccountPermissions`

GetAccountPermissions returns the AccountPermissions field if non-nil, zero value otherwise.

### GetAccountPermissionsOk

`func (o *IamSessionAllOf) GetAccountPermissionsOk() (*[]IamAccountPermissions, bool)`

GetAccountPermissionsOk returns a tuple with the AccountPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountPermissions

`func (o *IamSessionAllOf) SetAccountPermissions(v []IamAccountPermissions)`

SetAccountPermissions sets AccountPermissions field to given value.

### HasAccountPermissions

`func (o *IamSessionAllOf) HasAccountPermissions() bool`

HasAccountPermissions returns a boolean if a field has been set.

### GetClientIpAddress

`func (o *IamSessionAllOf) GetClientIpAddress() string`

GetClientIpAddress returns the ClientIpAddress field if non-nil, zero value otherwise.

### GetClientIpAddressOk

`func (o *IamSessionAllOf) GetClientIpAddressOk() (*string, bool)`

GetClientIpAddressOk returns a tuple with the ClientIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIpAddress

`func (o *IamSessionAllOf) SetClientIpAddress(v string)`

SetClientIpAddress sets ClientIpAddress field to given value.

### HasClientIpAddress

`func (o *IamSessionAllOf) HasClientIpAddress() bool`

HasClientIpAddress returns a boolean if a field has been set.

### GetExpiration

`func (o *IamSessionAllOf) GetExpiration() time.Time`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *IamSessionAllOf) GetExpirationOk() (*time.Time, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *IamSessionAllOf) SetExpiration(v time.Time)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *IamSessionAllOf) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.

### GetIdleTimeExpiration

`func (o *IamSessionAllOf) GetIdleTimeExpiration() time.Time`

GetIdleTimeExpiration returns the IdleTimeExpiration field if non-nil, zero value otherwise.

### GetIdleTimeExpirationOk

`func (o *IamSessionAllOf) GetIdleTimeExpirationOk() (*time.Time, bool)`

GetIdleTimeExpirationOk returns a tuple with the IdleTimeExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleTimeExpiration

`func (o *IamSessionAllOf) SetIdleTimeExpiration(v time.Time)`

SetIdleTimeExpiration sets IdleTimeExpiration field to given value.

### HasIdleTimeExpiration

`func (o *IamSessionAllOf) HasIdleTimeExpiration() bool`

HasIdleTimeExpiration returns a boolean if a field has been set.

### GetLastLoginClient

`func (o *IamSessionAllOf) GetLastLoginClient() string`

GetLastLoginClient returns the LastLoginClient field if non-nil, zero value otherwise.

### GetLastLoginClientOk

`func (o *IamSessionAllOf) GetLastLoginClientOk() (*string, bool)`

GetLastLoginClientOk returns a tuple with the LastLoginClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLoginClient

`func (o *IamSessionAllOf) SetLastLoginClient(v string)`

SetLastLoginClient sets LastLoginClient field to given value.

### HasLastLoginClient

`func (o *IamSessionAllOf) HasLastLoginClient() bool`

HasLastLoginClient returns a boolean if a field has been set.

### GetLastLoginTime

`func (o *IamSessionAllOf) GetLastLoginTime() time.Time`

GetLastLoginTime returns the LastLoginTime field if non-nil, zero value otherwise.

### GetLastLoginTimeOk

`func (o *IamSessionAllOf) GetLastLoginTimeOk() (*time.Time, bool)`

GetLastLoginTimeOk returns a tuple with the LastLoginTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLoginTime

`func (o *IamSessionAllOf) SetLastLoginTime(v time.Time)`

SetLastLoginTime sets LastLoginTime field to given value.

### HasLastLoginTime

`func (o *IamSessionAllOf) HasLastLoginTime() bool`

HasLastLoginTime returns a boolean if a field has been set.

### GetPermission

`func (o *IamSessionAllOf) GetPermission() IamPermissionRelationship`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *IamSessionAllOf) GetPermissionOk() (*IamPermissionRelationship, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *IamSessionAllOf) SetPermission(v IamPermissionRelationship)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *IamSessionAllOf) HasPermission() bool`

HasPermission returns a boolean if a field has been set.

### GetUser

`func (o *IamSessionAllOf) GetUser() IamUserRelationship`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *IamSessionAllOf) GetUserOk() (*IamUserRelationship, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *IamSessionAllOf) SetUser(v IamUserRelationship)`

SetUser sets User field to given value.

### HasUser

`func (o *IamSessionAllOf) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



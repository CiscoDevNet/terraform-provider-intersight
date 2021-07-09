# IamSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iam.Session"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iam.Session"]
**AccountPermissions** | Pointer to [**[]IamAccountPermissions**](IamAccountPermissions.md) |  | [optional] 
**Expiration** | Pointer to **time.Time** | Expiration time for the session. | [optional] [readonly] 
**FailedLogins** | Pointer to **int64** | Failed logins since last login for admin user. | [optional] [readonly] 
**IdleTimeExpiration** | Pointer to **time.Time** | Idle time expiration for the session. | [optional] [readonly] 
**LastLoginClient** | Pointer to **string** | The client address from which last login is initiated. | [optional] [readonly] 
**LastLoginTime** | Pointer to **time.Time** | The last login time for user. | [optional] [readonly] 
**SessionId** | Pointer to **string** | Session token shared with the user agent which is used to identify the user session when API requests are received to perform authorization. | [optional] 
**Permission** | Pointer to [**IamPermissionRelationship**](iam.Permission.Relationship.md) |  | [optional] 
**User** | Pointer to [**IamUserRelationship**](iam.User.Relationship.md) |  | [optional] 

## Methods

### NewIamSession

`func NewIamSession(classId string, objectType string, ) *IamSession`

NewIamSession instantiates a new IamSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamSessionWithDefaults

`func NewIamSessionWithDefaults() *IamSession`

NewIamSessionWithDefaults instantiates a new IamSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IamSession) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamSession) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamSession) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IamSession) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamSession) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamSession) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


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

### SetAccountPermissionsNil

`func (o *IamSession) SetAccountPermissionsNil(b bool)`

 SetAccountPermissionsNil sets the value for AccountPermissions to be an explicit nil

### UnsetAccountPermissions
`func (o *IamSession) UnsetAccountPermissions()`

UnsetAccountPermissions ensures that no value is present for AccountPermissions, not even an explicit nil
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

### GetFailedLogins

`func (o *IamSession) GetFailedLogins() int64`

GetFailedLogins returns the FailedLogins field if non-nil, zero value otherwise.

### GetFailedLoginsOk

`func (o *IamSession) GetFailedLoginsOk() (*int64, bool)`

GetFailedLoginsOk returns a tuple with the FailedLogins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedLogins

`func (o *IamSession) SetFailedLogins(v int64)`

SetFailedLogins sets FailedLogins field to given value.

### HasFailedLogins

`func (o *IamSession) HasFailedLogins() bool`

HasFailedLogins returns a boolean if a field has been set.

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

### GetSessionId

`func (o *IamSession) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *IamSession) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *IamSession) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *IamSession) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

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



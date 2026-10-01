# IamUserGroupMembership

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iam.UserGroupMembership"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iam.UserGroupMembership"]
**ExpirationTime** | Pointer to **time.Time** | Time after which this retained membership cannot authorize a new C3 session. | [optional] [readonly] 
**LastValidatedTime** | Pointer to **time.Time** | Time at which a regular CUI login most recently validated this UserGroup membership. | [optional] [readonly] 
**User** | Pointer to [**NullableIamUserRelationship**](IamUserRelationship.md) |  | [optional] 
**UserGroup** | Pointer to [**NullableIamUserGroupRelationship**](IamUserGroupRelationship.md) |  | [optional] 

## Methods

### NewIamUserGroupMembership

`func NewIamUserGroupMembership(classId string, objectType string, ) *IamUserGroupMembership`

NewIamUserGroupMembership instantiates a new IamUserGroupMembership object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamUserGroupMembershipWithDefaults

`func NewIamUserGroupMembershipWithDefaults() *IamUserGroupMembership`

NewIamUserGroupMembershipWithDefaults instantiates a new IamUserGroupMembership object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IamUserGroupMembership) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamUserGroupMembership) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamUserGroupMembership) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IamUserGroupMembership) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamUserGroupMembership) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamUserGroupMembership) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetExpirationTime

`func (o *IamUserGroupMembership) GetExpirationTime() time.Time`

GetExpirationTime returns the ExpirationTime field if non-nil, zero value otherwise.

### GetExpirationTimeOk

`func (o *IamUserGroupMembership) GetExpirationTimeOk() (*time.Time, bool)`

GetExpirationTimeOk returns a tuple with the ExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTime

`func (o *IamUserGroupMembership) SetExpirationTime(v time.Time)`

SetExpirationTime sets ExpirationTime field to given value.

### HasExpirationTime

`func (o *IamUserGroupMembership) HasExpirationTime() bool`

HasExpirationTime returns a boolean if a field has been set.

### GetLastValidatedTime

`func (o *IamUserGroupMembership) GetLastValidatedTime() time.Time`

GetLastValidatedTime returns the LastValidatedTime field if non-nil, zero value otherwise.

### GetLastValidatedTimeOk

`func (o *IamUserGroupMembership) GetLastValidatedTimeOk() (*time.Time, bool)`

GetLastValidatedTimeOk returns a tuple with the LastValidatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastValidatedTime

`func (o *IamUserGroupMembership) SetLastValidatedTime(v time.Time)`

SetLastValidatedTime sets LastValidatedTime field to given value.

### HasLastValidatedTime

`func (o *IamUserGroupMembership) HasLastValidatedTime() bool`

HasLastValidatedTime returns a boolean if a field has been set.

### GetUser

`func (o *IamUserGroupMembership) GetUser() IamUserRelationship`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *IamUserGroupMembership) GetUserOk() (*IamUserRelationship, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *IamUserGroupMembership) SetUser(v IamUserRelationship)`

SetUser sets User field to given value.

### HasUser

`func (o *IamUserGroupMembership) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *IamUserGroupMembership) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *IamUserGroupMembership) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetUserGroup

`func (o *IamUserGroupMembership) GetUserGroup() IamUserGroupRelationship`

GetUserGroup returns the UserGroup field if non-nil, zero value otherwise.

### GetUserGroupOk

`func (o *IamUserGroupMembership) GetUserGroupOk() (*IamUserGroupRelationship, bool)`

GetUserGroupOk returns a tuple with the UserGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroup

`func (o *IamUserGroupMembership) SetUserGroup(v IamUserGroupRelationship)`

SetUserGroup sets UserGroup field to given value.

### HasUserGroup

`func (o *IamUserGroupMembership) HasUserGroup() bool`

HasUserGroup returns a boolean if a field has been set.

### SetUserGroupNil

`func (o *IamUserGroupMembership) SetUserGroupNil(b bool)`

 SetUserGroupNil sets the value for UserGroup to be an explicit nil

### UnsetUserGroup
`func (o *IamUserGroupMembership) UnsetUserGroup()`

UnsetUserGroup ensures that no value is present for UserGroup, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



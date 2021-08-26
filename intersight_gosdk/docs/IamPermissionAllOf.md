# IamPermissionAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iam.Permission"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iam.Permission"]
**Description** | Pointer to **string** | The informative description about each permission. | [optional] 
**Name** | Pointer to **string** | The name of the permission which has to be granted to user. | [optional] 
**Account** | Pointer to [**IamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 
**EndPointRoles** | Pointer to [**[]IamEndPointRoleRelationship**](IamEndPointRoleRelationship.md) | An array of relationships to iamEndPointRole resources. | [optional] [readonly] 
**PrivilegeSets** | Pointer to [**[]IamPrivilegeSetRelationship**](IamPrivilegeSetRelationship.md) | An array of relationships to iamPrivilegeSet resources. | [optional] [readonly] 
**ResourceRoles** | Pointer to [**[]IamResourceRolesRelationship**](IamResourceRolesRelationship.md) | An array of relationships to iamResourceRoles resources. | [optional] 
**Roles** | Pointer to [**[]IamRoleRelationship**](IamRoleRelationship.md) | An array of relationships to iamRole resources. | [optional] 
**SessionLimits** | Pointer to [**IamSessionLimitsRelationship**](IamSessionLimitsRelationship.md) |  | [optional] 
**UserGroups** | Pointer to [**[]IamUserGroupRelationship**](IamUserGroupRelationship.md) | An array of relationships to iamUserGroup resources. | [optional] 
**Users** | Pointer to [**[]IamUserRelationship**](IamUserRelationship.md) | An array of relationships to iamUser resources. | [optional] 

## Methods

### NewIamPermissionAllOf

`func NewIamPermissionAllOf(classId string, objectType string, ) *IamPermissionAllOf`

NewIamPermissionAllOf instantiates a new IamPermissionAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamPermissionAllOfWithDefaults

`func NewIamPermissionAllOfWithDefaults() *IamPermissionAllOf`

NewIamPermissionAllOfWithDefaults instantiates a new IamPermissionAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IamPermissionAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamPermissionAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamPermissionAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IamPermissionAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamPermissionAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamPermissionAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDescription

`func (o *IamPermissionAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IamPermissionAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IamPermissionAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IamPermissionAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *IamPermissionAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IamPermissionAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IamPermissionAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IamPermissionAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAccount

`func (o *IamPermissionAllOf) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *IamPermissionAllOf) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *IamPermissionAllOf) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *IamPermissionAllOf) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetEndPointRoles

`func (o *IamPermissionAllOf) GetEndPointRoles() []IamEndPointRoleRelationship`

GetEndPointRoles returns the EndPointRoles field if non-nil, zero value otherwise.

### GetEndPointRolesOk

`func (o *IamPermissionAllOf) GetEndPointRolesOk() (*[]IamEndPointRoleRelationship, bool)`

GetEndPointRolesOk returns a tuple with the EndPointRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPointRoles

`func (o *IamPermissionAllOf) SetEndPointRoles(v []IamEndPointRoleRelationship)`

SetEndPointRoles sets EndPointRoles field to given value.

### HasEndPointRoles

`func (o *IamPermissionAllOf) HasEndPointRoles() bool`

HasEndPointRoles returns a boolean if a field has been set.

### SetEndPointRolesNil

`func (o *IamPermissionAllOf) SetEndPointRolesNil(b bool)`

 SetEndPointRolesNil sets the value for EndPointRoles to be an explicit nil

### UnsetEndPointRoles
`func (o *IamPermissionAllOf) UnsetEndPointRoles()`

UnsetEndPointRoles ensures that no value is present for EndPointRoles, not even an explicit nil
### GetPrivilegeSets

`func (o *IamPermissionAllOf) GetPrivilegeSets() []IamPrivilegeSetRelationship`

GetPrivilegeSets returns the PrivilegeSets field if non-nil, zero value otherwise.

### GetPrivilegeSetsOk

`func (o *IamPermissionAllOf) GetPrivilegeSetsOk() (*[]IamPrivilegeSetRelationship, bool)`

GetPrivilegeSetsOk returns a tuple with the PrivilegeSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivilegeSets

`func (o *IamPermissionAllOf) SetPrivilegeSets(v []IamPrivilegeSetRelationship)`

SetPrivilegeSets sets PrivilegeSets field to given value.

### HasPrivilegeSets

`func (o *IamPermissionAllOf) HasPrivilegeSets() bool`

HasPrivilegeSets returns a boolean if a field has been set.

### SetPrivilegeSetsNil

`func (o *IamPermissionAllOf) SetPrivilegeSetsNil(b bool)`

 SetPrivilegeSetsNil sets the value for PrivilegeSets to be an explicit nil

### UnsetPrivilegeSets
`func (o *IamPermissionAllOf) UnsetPrivilegeSets()`

UnsetPrivilegeSets ensures that no value is present for PrivilegeSets, not even an explicit nil
### GetResourceRoles

`func (o *IamPermissionAllOf) GetResourceRoles() []IamResourceRolesRelationship`

GetResourceRoles returns the ResourceRoles field if non-nil, zero value otherwise.

### GetResourceRolesOk

`func (o *IamPermissionAllOf) GetResourceRolesOk() (*[]IamResourceRolesRelationship, bool)`

GetResourceRolesOk returns a tuple with the ResourceRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceRoles

`func (o *IamPermissionAllOf) SetResourceRoles(v []IamResourceRolesRelationship)`

SetResourceRoles sets ResourceRoles field to given value.

### HasResourceRoles

`func (o *IamPermissionAllOf) HasResourceRoles() bool`

HasResourceRoles returns a boolean if a field has been set.

### SetResourceRolesNil

`func (o *IamPermissionAllOf) SetResourceRolesNil(b bool)`

 SetResourceRolesNil sets the value for ResourceRoles to be an explicit nil

### UnsetResourceRoles
`func (o *IamPermissionAllOf) UnsetResourceRoles()`

UnsetResourceRoles ensures that no value is present for ResourceRoles, not even an explicit nil
### GetRoles

`func (o *IamPermissionAllOf) GetRoles() []IamRoleRelationship`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *IamPermissionAllOf) GetRolesOk() (*[]IamRoleRelationship, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *IamPermissionAllOf) SetRoles(v []IamRoleRelationship)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *IamPermissionAllOf) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### SetRolesNil

`func (o *IamPermissionAllOf) SetRolesNil(b bool)`

 SetRolesNil sets the value for Roles to be an explicit nil

### UnsetRoles
`func (o *IamPermissionAllOf) UnsetRoles()`

UnsetRoles ensures that no value is present for Roles, not even an explicit nil
### GetSessionLimits

`func (o *IamPermissionAllOf) GetSessionLimits() IamSessionLimitsRelationship`

GetSessionLimits returns the SessionLimits field if non-nil, zero value otherwise.

### GetSessionLimitsOk

`func (o *IamPermissionAllOf) GetSessionLimitsOk() (*IamSessionLimitsRelationship, bool)`

GetSessionLimitsOk returns a tuple with the SessionLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionLimits

`func (o *IamPermissionAllOf) SetSessionLimits(v IamSessionLimitsRelationship)`

SetSessionLimits sets SessionLimits field to given value.

### HasSessionLimits

`func (o *IamPermissionAllOf) HasSessionLimits() bool`

HasSessionLimits returns a boolean if a field has been set.

### GetUserGroups

`func (o *IamPermissionAllOf) GetUserGroups() []IamUserGroupRelationship`

GetUserGroups returns the UserGroups field if non-nil, zero value otherwise.

### GetUserGroupsOk

`func (o *IamPermissionAllOf) GetUserGroupsOk() (*[]IamUserGroupRelationship, bool)`

GetUserGroupsOk returns a tuple with the UserGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroups

`func (o *IamPermissionAllOf) SetUserGroups(v []IamUserGroupRelationship)`

SetUserGroups sets UserGroups field to given value.

### HasUserGroups

`func (o *IamPermissionAllOf) HasUserGroups() bool`

HasUserGroups returns a boolean if a field has been set.

### SetUserGroupsNil

`func (o *IamPermissionAllOf) SetUserGroupsNil(b bool)`

 SetUserGroupsNil sets the value for UserGroups to be an explicit nil

### UnsetUserGroups
`func (o *IamPermissionAllOf) UnsetUserGroups()`

UnsetUserGroups ensures that no value is present for UserGroups, not even an explicit nil
### GetUsers

`func (o *IamPermissionAllOf) GetUsers() []IamUserRelationship`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *IamPermissionAllOf) GetUsersOk() (*[]IamUserRelationship, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *IamPermissionAllOf) SetUsers(v []IamUserRelationship)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *IamPermissionAllOf) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### SetUsersNil

`func (o *IamPermissionAllOf) SetUsersNil(b bool)`

 SetUsersNil sets the value for Users to be an explicit nil

### UnsetUsers
`func (o *IamPermissionAllOf) UnsetUsers()`

UnsetUsers ensures that no value is present for Users, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



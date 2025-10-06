# IamUserGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iam.UserGroup"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iam.UserGroup"]
**AccessActivationTime** | Pointer to **time.Time** | AccessActivationTime indicates the activation time for the guest user&#39;s access to the Account.  Before this time, if guest user tries to login to the account, access the account will be denied. | [optional] 
**AccessExpiryTime** | Pointer to **time.Time** | AccessExpiryTime indicates the expiration time for the guest user&#39;s access to the Account. Its value can only be  assigned a date that falls within the range determined by the maximum expiration time configured for the  API entries. The AccessExpiry date can be edited to be earlier or later. | [optional] 
**AccessLink** | Pointer to **string** | AccessLink using which the guest user uses to log in to Intersight. | [optional] [readonly] 
**GroupType** | Pointer to **string** | Group type determines the type of groups that is being associated with users. By default, Default User group will be used for associating dynamic user login. If the value of the User Group is set to guest, then this type of user group will be used for guest user login. * &#x60;Default&#x60; - Default User Group Type used for dynamic users login. * &#x60;Guest&#x60; - Guest User Group type used for guest users login. | [optional] [default to "Default"]
**Instruction** | Pointer to **string** | Instruction property holds detailed guidance and information intended for individuals  accessing the system as guest users. It holds the information to assist guests in navigating the platform,  understanding policies, and performing necessary actions to ensure a seamless and secure user experience. | [optional] 
**Name** | Pointer to **string** | The name of the user group which the dynamic/or guest user belongs to. | [optional] 
**UniqueReferenceId** | Pointer to **string** | A random mixed character string which is unique per user groups. UniqueReferenceId is used as key for identifying the guest user groups. | [optional] [readonly] 
**Idp** | Pointer to [**NullableIamIdpRelationship**](IamIdpRelationship.md) |  | [optional] 
**Idpreference** | Pointer to [**NullableIamIdpReferenceRelationship**](IamIdpReferenceRelationship.md) |  | [optional] 
**Permissions** | Pointer to [**[]IamPermissionRelationship**](IamPermissionRelationship.md) | An array of relationships to iamPermission resources. | [optional] 
**Qualifier** | Pointer to [**NullableIamAbstractQualifierRelationship**](IamAbstractQualifierRelationship.md) |  | [optional] 
**Users** | Pointer to [**[]IamUserRelationship**](IamUserRelationship.md) | An array of relationships to iamUser resources. | [optional] [readonly] 

## Methods

### NewIamUserGroup

`func NewIamUserGroup(classId string, objectType string, ) *IamUserGroup`

NewIamUserGroup instantiates a new IamUserGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamUserGroupWithDefaults

`func NewIamUserGroupWithDefaults() *IamUserGroup`

NewIamUserGroupWithDefaults instantiates a new IamUserGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IamUserGroup) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamUserGroup) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamUserGroup) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IamUserGroup) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamUserGroup) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamUserGroup) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAccessActivationTime

`func (o *IamUserGroup) GetAccessActivationTime() time.Time`

GetAccessActivationTime returns the AccessActivationTime field if non-nil, zero value otherwise.

### GetAccessActivationTimeOk

`func (o *IamUserGroup) GetAccessActivationTimeOk() (*time.Time, bool)`

GetAccessActivationTimeOk returns a tuple with the AccessActivationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessActivationTime

`func (o *IamUserGroup) SetAccessActivationTime(v time.Time)`

SetAccessActivationTime sets AccessActivationTime field to given value.

### HasAccessActivationTime

`func (o *IamUserGroup) HasAccessActivationTime() bool`

HasAccessActivationTime returns a boolean if a field has been set.

### GetAccessExpiryTime

`func (o *IamUserGroup) GetAccessExpiryTime() time.Time`

GetAccessExpiryTime returns the AccessExpiryTime field if non-nil, zero value otherwise.

### GetAccessExpiryTimeOk

`func (o *IamUserGroup) GetAccessExpiryTimeOk() (*time.Time, bool)`

GetAccessExpiryTimeOk returns a tuple with the AccessExpiryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessExpiryTime

`func (o *IamUserGroup) SetAccessExpiryTime(v time.Time)`

SetAccessExpiryTime sets AccessExpiryTime field to given value.

### HasAccessExpiryTime

`func (o *IamUserGroup) HasAccessExpiryTime() bool`

HasAccessExpiryTime returns a boolean if a field has been set.

### GetAccessLink

`func (o *IamUserGroup) GetAccessLink() string`

GetAccessLink returns the AccessLink field if non-nil, zero value otherwise.

### GetAccessLinkOk

`func (o *IamUserGroup) GetAccessLinkOk() (*string, bool)`

GetAccessLinkOk returns a tuple with the AccessLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessLink

`func (o *IamUserGroup) SetAccessLink(v string)`

SetAccessLink sets AccessLink field to given value.

### HasAccessLink

`func (o *IamUserGroup) HasAccessLink() bool`

HasAccessLink returns a boolean if a field has been set.

### GetGroupType

`func (o *IamUserGroup) GetGroupType() string`

GetGroupType returns the GroupType field if non-nil, zero value otherwise.

### GetGroupTypeOk

`func (o *IamUserGroup) GetGroupTypeOk() (*string, bool)`

GetGroupTypeOk returns a tuple with the GroupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupType

`func (o *IamUserGroup) SetGroupType(v string)`

SetGroupType sets GroupType field to given value.

### HasGroupType

`func (o *IamUserGroup) HasGroupType() bool`

HasGroupType returns a boolean if a field has been set.

### GetInstruction

`func (o *IamUserGroup) GetInstruction() string`

GetInstruction returns the Instruction field if non-nil, zero value otherwise.

### GetInstructionOk

`func (o *IamUserGroup) GetInstructionOk() (*string, bool)`

GetInstructionOk returns a tuple with the Instruction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstruction

`func (o *IamUserGroup) SetInstruction(v string)`

SetInstruction sets Instruction field to given value.

### HasInstruction

`func (o *IamUserGroup) HasInstruction() bool`

HasInstruction returns a boolean if a field has been set.

### GetName

`func (o *IamUserGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IamUserGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IamUserGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IamUserGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUniqueReferenceId

`func (o *IamUserGroup) GetUniqueReferenceId() string`

GetUniqueReferenceId returns the UniqueReferenceId field if non-nil, zero value otherwise.

### GetUniqueReferenceIdOk

`func (o *IamUserGroup) GetUniqueReferenceIdOk() (*string, bool)`

GetUniqueReferenceIdOk returns a tuple with the UniqueReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueReferenceId

`func (o *IamUserGroup) SetUniqueReferenceId(v string)`

SetUniqueReferenceId sets UniqueReferenceId field to given value.

### HasUniqueReferenceId

`func (o *IamUserGroup) HasUniqueReferenceId() bool`

HasUniqueReferenceId returns a boolean if a field has been set.

### GetIdp

`func (o *IamUserGroup) GetIdp() IamIdpRelationship`

GetIdp returns the Idp field if non-nil, zero value otherwise.

### GetIdpOk

`func (o *IamUserGroup) GetIdpOk() (*IamIdpRelationship, bool)`

GetIdpOk returns a tuple with the Idp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdp

`func (o *IamUserGroup) SetIdp(v IamIdpRelationship)`

SetIdp sets Idp field to given value.

### HasIdp

`func (o *IamUserGroup) HasIdp() bool`

HasIdp returns a boolean if a field has been set.

### SetIdpNil

`func (o *IamUserGroup) SetIdpNil(b bool)`

 SetIdpNil sets the value for Idp to be an explicit nil

### UnsetIdp
`func (o *IamUserGroup) UnsetIdp()`

UnsetIdp ensures that no value is present for Idp, not even an explicit nil
### GetIdpreference

`func (o *IamUserGroup) GetIdpreference() IamIdpReferenceRelationship`

GetIdpreference returns the Idpreference field if non-nil, zero value otherwise.

### GetIdpreferenceOk

`func (o *IamUserGroup) GetIdpreferenceOk() (*IamIdpReferenceRelationship, bool)`

GetIdpreferenceOk returns a tuple with the Idpreference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpreference

`func (o *IamUserGroup) SetIdpreference(v IamIdpReferenceRelationship)`

SetIdpreference sets Idpreference field to given value.

### HasIdpreference

`func (o *IamUserGroup) HasIdpreference() bool`

HasIdpreference returns a boolean if a field has been set.

### SetIdpreferenceNil

`func (o *IamUserGroup) SetIdpreferenceNil(b bool)`

 SetIdpreferenceNil sets the value for Idpreference to be an explicit nil

### UnsetIdpreference
`func (o *IamUserGroup) UnsetIdpreference()`

UnsetIdpreference ensures that no value is present for Idpreference, not even an explicit nil
### GetPermissions

`func (o *IamUserGroup) GetPermissions() []IamPermissionRelationship`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *IamUserGroup) GetPermissionsOk() (*[]IamPermissionRelationship, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *IamUserGroup) SetPermissions(v []IamPermissionRelationship)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *IamUserGroup) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### SetPermissionsNil

`func (o *IamUserGroup) SetPermissionsNil(b bool)`

 SetPermissionsNil sets the value for Permissions to be an explicit nil

### UnsetPermissions
`func (o *IamUserGroup) UnsetPermissions()`

UnsetPermissions ensures that no value is present for Permissions, not even an explicit nil
### GetQualifier

`func (o *IamUserGroup) GetQualifier() IamAbstractQualifierRelationship`

GetQualifier returns the Qualifier field if non-nil, zero value otherwise.

### GetQualifierOk

`func (o *IamUserGroup) GetQualifierOk() (*IamAbstractQualifierRelationship, bool)`

GetQualifierOk returns a tuple with the Qualifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualifier

`func (o *IamUserGroup) SetQualifier(v IamAbstractQualifierRelationship)`

SetQualifier sets Qualifier field to given value.

### HasQualifier

`func (o *IamUserGroup) HasQualifier() bool`

HasQualifier returns a boolean if a field has been set.

### SetQualifierNil

`func (o *IamUserGroup) SetQualifierNil(b bool)`

 SetQualifierNil sets the value for Qualifier to be an explicit nil

### UnsetQualifier
`func (o *IamUserGroup) UnsetQualifier()`

UnsetQualifier ensures that no value is present for Qualifier, not even an explicit nil
### GetUsers

`func (o *IamUserGroup) GetUsers() []IamUserRelationship`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *IamUserGroup) GetUsersOk() (*[]IamUserRelationship, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *IamUserGroup) SetUsers(v []IamUserRelationship)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *IamUserGroup) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### SetUsersNil

`func (o *IamUserGroup) SetUsersNil(b bool)`

 SetUsersNil sets the value for Users to be an explicit nil

### UnsetUsers
`func (o *IamUserGroup) UnsetUsers()`

UnsetUsers ensures that no value is present for Users, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



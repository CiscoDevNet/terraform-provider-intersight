# IamUserGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iam.UserGroup"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iam.UserGroup"]
**Name** | Pointer to **string** | The name of the user group which the dynamic user belongs to. | [optional] 
**Idp** | Pointer to [**IamIdpRelationship**](IamIdpRelationship.md) |  | [optional] 
**Idpreference** | Pointer to [**IamIdpReferenceRelationship**](IamIdpReferenceRelationship.md) |  | [optional] 
**Permissions** | Pointer to [**[]IamPermissionRelationship**](IamPermissionRelationship.md) | An array of relationships to iamPermission resources. | [optional] 
**Qualifier** | Pointer to [**IamQualifierRelationship**](IamQualifierRelationship.md) |  | [optional] 
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

`func (o *IamUserGroup) GetQualifier() IamQualifierRelationship`

GetQualifier returns the Qualifier field if non-nil, zero value otherwise.

### GetQualifierOk

`func (o *IamUserGroup) GetQualifierOk() (*IamQualifierRelationship, bool)`

GetQualifierOk returns a tuple with the Qualifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualifier

`func (o *IamUserGroup) SetQualifier(v IamQualifierRelationship)`

SetQualifier sets Qualifier field to given value.

### HasQualifier

`func (o *IamUserGroup) HasQualifier() bool`

HasQualifier returns a boolean if a field has been set.

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



# IamUserGroupAllOf

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

### NewIamUserGroupAllOf

`func NewIamUserGroupAllOf(classId string, objectType string, ) *IamUserGroupAllOf`

NewIamUserGroupAllOf instantiates a new IamUserGroupAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamUserGroupAllOfWithDefaults

`func NewIamUserGroupAllOfWithDefaults() *IamUserGroupAllOf`

NewIamUserGroupAllOfWithDefaults instantiates a new IamUserGroupAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IamUserGroupAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamUserGroupAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamUserGroupAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IamUserGroupAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamUserGroupAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamUserGroupAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetName

`func (o *IamUserGroupAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IamUserGroupAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IamUserGroupAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IamUserGroupAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIdp

`func (o *IamUserGroupAllOf) GetIdp() IamIdpRelationship`

GetIdp returns the Idp field if non-nil, zero value otherwise.

### GetIdpOk

`func (o *IamUserGroupAllOf) GetIdpOk() (*IamIdpRelationship, bool)`

GetIdpOk returns a tuple with the Idp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdp

`func (o *IamUserGroupAllOf) SetIdp(v IamIdpRelationship)`

SetIdp sets Idp field to given value.

### HasIdp

`func (o *IamUserGroupAllOf) HasIdp() bool`

HasIdp returns a boolean if a field has been set.

### GetIdpreference

`func (o *IamUserGroupAllOf) GetIdpreference() IamIdpReferenceRelationship`

GetIdpreference returns the Idpreference field if non-nil, zero value otherwise.

### GetIdpreferenceOk

`func (o *IamUserGroupAllOf) GetIdpreferenceOk() (*IamIdpReferenceRelationship, bool)`

GetIdpreferenceOk returns a tuple with the Idpreference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpreference

`func (o *IamUserGroupAllOf) SetIdpreference(v IamIdpReferenceRelationship)`

SetIdpreference sets Idpreference field to given value.

### HasIdpreference

`func (o *IamUserGroupAllOf) HasIdpreference() bool`

HasIdpreference returns a boolean if a field has been set.

### GetPermissions

`func (o *IamUserGroupAllOf) GetPermissions() []IamPermissionRelationship`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *IamUserGroupAllOf) GetPermissionsOk() (*[]IamPermissionRelationship, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *IamUserGroupAllOf) SetPermissions(v []IamPermissionRelationship)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *IamUserGroupAllOf) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### SetPermissionsNil

`func (o *IamUserGroupAllOf) SetPermissionsNil(b bool)`

 SetPermissionsNil sets the value for Permissions to be an explicit nil

### UnsetPermissions
`func (o *IamUserGroupAllOf) UnsetPermissions()`

UnsetPermissions ensures that no value is present for Permissions, not even an explicit nil
### GetQualifier

`func (o *IamUserGroupAllOf) GetQualifier() IamQualifierRelationship`

GetQualifier returns the Qualifier field if non-nil, zero value otherwise.

### GetQualifierOk

`func (o *IamUserGroupAllOf) GetQualifierOk() (*IamQualifierRelationship, bool)`

GetQualifierOk returns a tuple with the Qualifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualifier

`func (o *IamUserGroupAllOf) SetQualifier(v IamQualifierRelationship)`

SetQualifier sets Qualifier field to given value.

### HasQualifier

`func (o *IamUserGroupAllOf) HasQualifier() bool`

HasQualifier returns a boolean if a field has been set.

### GetUsers

`func (o *IamUserGroupAllOf) GetUsers() []IamUserRelationship`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *IamUserGroupAllOf) GetUsersOk() (*[]IamUserRelationship, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *IamUserGroupAllOf) SetUsers(v []IamUserRelationship)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *IamUserGroupAllOf) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### SetUsersNil

`func (o *IamUserGroupAllOf) SetUsersNil(b bool)`

 SetUsersNil sets the value for Users to be an explicit nil

### UnsetUsers
`func (o *IamUserGroupAllOf) UnsetUsers()`

UnsetUsers ensures that no value is present for Users, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



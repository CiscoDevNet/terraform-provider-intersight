# IamPermissionToRoles

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iam.PermissionToRoles"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iam.PermissionToRoles"]
**Permission** | Pointer to [**MoMoRef**](MoMoRef.md) |  | [optional] 
**Roles** | Pointer to [**[]MoMoRef**](MoMoRef.md) |  | [optional] 

## Methods

### NewIamPermissionToRoles

`func NewIamPermissionToRoles(classId string, objectType string, ) *IamPermissionToRoles`

NewIamPermissionToRoles instantiates a new IamPermissionToRoles object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamPermissionToRolesWithDefaults

`func NewIamPermissionToRolesWithDefaults() *IamPermissionToRoles`

NewIamPermissionToRolesWithDefaults instantiates a new IamPermissionToRoles object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IamPermissionToRoles) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamPermissionToRoles) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamPermissionToRoles) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IamPermissionToRoles) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamPermissionToRoles) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamPermissionToRoles) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetPermission

`func (o *IamPermissionToRoles) GetPermission() MoMoRef`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *IamPermissionToRoles) GetPermissionOk() (*MoMoRef, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *IamPermissionToRoles) SetPermission(v MoMoRef)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *IamPermissionToRoles) HasPermission() bool`

HasPermission returns a boolean if a field has been set.

### GetRoles

`func (o *IamPermissionToRoles) GetRoles() []MoMoRef`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *IamPermissionToRoles) GetRolesOk() (*[]MoMoRef, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *IamPermissionToRoles) SetRoles(v []MoMoRef)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *IamPermissionToRoles) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### SetRolesNil

`func (o *IamPermissionToRoles) SetRolesNil(b bool)`

 SetRolesNil sets the value for Roles to be an explicit nil

### UnsetRoles
`func (o *IamPermissionToRoles) UnsetRoles()`

UnsetRoles ensures that no value is present for Roles, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



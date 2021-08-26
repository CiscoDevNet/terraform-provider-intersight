# IamResourceRoles

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iam.ResourceRoles"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iam.ResourceRoles"]
**EndPointRoles** | Pointer to [**[]IamEndPointRoleRelationship**](IamEndPointRoleRelationship.md) | An array of relationships to iamEndPointRole resources. | [optional] [readonly] 
**Permission** | Pointer to [**IamPermissionRelationship**](IamPermissionRelationship.md) |  | [optional] 
**PrivilegeSets** | Pointer to [**[]IamPrivilegeSetRelationship**](IamPrivilegeSetRelationship.md) | An array of relationships to iamPrivilegeSet resources. | [optional] [readonly] 
**Resource** | Pointer to [**MoBaseMoRelationship**](MoBaseMoRelationship.md) |  | [optional] 
**Roles** | Pointer to [**[]IamRoleRelationship**](IamRoleRelationship.md) | An array of relationships to iamRole resources. | [optional] 

## Methods

### NewIamResourceRoles

`func NewIamResourceRoles(classId string, objectType string, ) *IamResourceRoles`

NewIamResourceRoles instantiates a new IamResourceRoles object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamResourceRolesWithDefaults

`func NewIamResourceRolesWithDefaults() *IamResourceRoles`

NewIamResourceRolesWithDefaults instantiates a new IamResourceRoles object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IamResourceRoles) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamResourceRoles) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamResourceRoles) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IamResourceRoles) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamResourceRoles) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamResourceRoles) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEndPointRoles

`func (o *IamResourceRoles) GetEndPointRoles() []IamEndPointRoleRelationship`

GetEndPointRoles returns the EndPointRoles field if non-nil, zero value otherwise.

### GetEndPointRolesOk

`func (o *IamResourceRoles) GetEndPointRolesOk() (*[]IamEndPointRoleRelationship, bool)`

GetEndPointRolesOk returns a tuple with the EndPointRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPointRoles

`func (o *IamResourceRoles) SetEndPointRoles(v []IamEndPointRoleRelationship)`

SetEndPointRoles sets EndPointRoles field to given value.

### HasEndPointRoles

`func (o *IamResourceRoles) HasEndPointRoles() bool`

HasEndPointRoles returns a boolean if a field has been set.

### SetEndPointRolesNil

`func (o *IamResourceRoles) SetEndPointRolesNil(b bool)`

 SetEndPointRolesNil sets the value for EndPointRoles to be an explicit nil

### UnsetEndPointRoles
`func (o *IamResourceRoles) UnsetEndPointRoles()`

UnsetEndPointRoles ensures that no value is present for EndPointRoles, not even an explicit nil
### GetPermission

`func (o *IamResourceRoles) GetPermission() IamPermissionRelationship`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *IamResourceRoles) GetPermissionOk() (*IamPermissionRelationship, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *IamResourceRoles) SetPermission(v IamPermissionRelationship)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *IamResourceRoles) HasPermission() bool`

HasPermission returns a boolean if a field has been set.

### GetPrivilegeSets

`func (o *IamResourceRoles) GetPrivilegeSets() []IamPrivilegeSetRelationship`

GetPrivilegeSets returns the PrivilegeSets field if non-nil, zero value otherwise.

### GetPrivilegeSetsOk

`func (o *IamResourceRoles) GetPrivilegeSetsOk() (*[]IamPrivilegeSetRelationship, bool)`

GetPrivilegeSetsOk returns a tuple with the PrivilegeSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivilegeSets

`func (o *IamResourceRoles) SetPrivilegeSets(v []IamPrivilegeSetRelationship)`

SetPrivilegeSets sets PrivilegeSets field to given value.

### HasPrivilegeSets

`func (o *IamResourceRoles) HasPrivilegeSets() bool`

HasPrivilegeSets returns a boolean if a field has been set.

### SetPrivilegeSetsNil

`func (o *IamResourceRoles) SetPrivilegeSetsNil(b bool)`

 SetPrivilegeSetsNil sets the value for PrivilegeSets to be an explicit nil

### UnsetPrivilegeSets
`func (o *IamResourceRoles) UnsetPrivilegeSets()`

UnsetPrivilegeSets ensures that no value is present for PrivilegeSets, not even an explicit nil
### GetResource

`func (o *IamResourceRoles) GetResource() MoBaseMoRelationship`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *IamResourceRoles) GetResourceOk() (*MoBaseMoRelationship, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *IamResourceRoles) SetResource(v MoBaseMoRelationship)`

SetResource sets Resource field to given value.

### HasResource

`func (o *IamResourceRoles) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetRoles

`func (o *IamResourceRoles) GetRoles() []IamRoleRelationship`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *IamResourceRoles) GetRolesOk() (*[]IamRoleRelationship, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *IamResourceRoles) SetRoles(v []IamRoleRelationship)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *IamResourceRoles) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### SetRolesNil

`func (o *IamResourceRoles) SetRolesNil(b bool)`

 SetRolesNil sets the value for Roles to be an explicit nil

### UnsetRoles
`func (o *IamResourceRoles) UnsetRoles()`

UnsetRoles ensures that no value is present for Roles, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# IamResourceRoles

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndPointRoles** | Pointer to [**[]IamEndPointRoleRelationship**](iam.EndPointRole.Relationship.md) | An array of relationships to iamEndPointRole resources. | [optional] [readonly] 
**Permission** | Pointer to [**IamPermissionRelationship**](iam.Permission.Relationship.md) |  | [optional] 
**Resource** | Pointer to [**MoBaseMoRelationship**](mo.BaseMo.Relationship.md) |  | [optional] 
**Roles** | Pointer to [**[]IamRoleRelationship**](iam.Role.Relationship.md) | An array of relationships to iamRole resources. | [optional] 

## Methods

### NewIamResourceRoles

`func NewIamResourceRoles() *IamResourceRoles`

NewIamResourceRoles instantiates a new IamResourceRoles object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamResourceRolesWithDefaults

`func NewIamResourceRolesWithDefaults() *IamResourceRoles`

NewIamResourceRolesWithDefaults instantiates a new IamResourceRoles object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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



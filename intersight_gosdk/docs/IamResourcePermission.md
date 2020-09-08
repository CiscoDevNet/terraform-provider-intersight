# IamResourcePermission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PermissionRoles** | Pointer to [**[]IamPermissionToRoles**](iam.PermissionToRoles.md) |  | [optional] 
**TargetApp** | Pointer to **string** | Name of the service owning the resource. | [optional] [readonly] 
**Holder** | Pointer to [**IamSecurityHolderRelationship**](iam.SecurityHolder.Relationship.md) |  | [optional] 
**Resource** | Pointer to [**MoBaseMoRelationship**](mo.BaseMo.Relationship.md) |  | [optional] 

## Methods

### NewIamResourcePermission

`func NewIamResourcePermission() *IamResourcePermission`

NewIamResourcePermission instantiates a new IamResourcePermission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamResourcePermissionWithDefaults

`func NewIamResourcePermissionWithDefaults() *IamResourcePermission`

NewIamResourcePermissionWithDefaults instantiates a new IamResourcePermission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermissionRoles

`func (o *IamResourcePermission) GetPermissionRoles() []IamPermissionToRoles`

GetPermissionRoles returns the PermissionRoles field if non-nil, zero value otherwise.

### GetPermissionRolesOk

`func (o *IamResourcePermission) GetPermissionRolesOk() (*[]IamPermissionToRoles, bool)`

GetPermissionRolesOk returns a tuple with the PermissionRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionRoles

`func (o *IamResourcePermission) SetPermissionRoles(v []IamPermissionToRoles)`

SetPermissionRoles sets PermissionRoles field to given value.

### HasPermissionRoles

`func (o *IamResourcePermission) HasPermissionRoles() bool`

HasPermissionRoles returns a boolean if a field has been set.

### GetTargetApp

`func (o *IamResourcePermission) GetTargetApp() string`

GetTargetApp returns the TargetApp field if non-nil, zero value otherwise.

### GetTargetAppOk

`func (o *IamResourcePermission) GetTargetAppOk() (*string, bool)`

GetTargetAppOk returns a tuple with the TargetApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetApp

`func (o *IamResourcePermission) SetTargetApp(v string)`

SetTargetApp sets TargetApp field to given value.

### HasTargetApp

`func (o *IamResourcePermission) HasTargetApp() bool`

HasTargetApp returns a boolean if a field has been set.

### GetHolder

`func (o *IamResourcePermission) GetHolder() IamSecurityHolderRelationship`

GetHolder returns the Holder field if non-nil, zero value otherwise.

### GetHolderOk

`func (o *IamResourcePermission) GetHolderOk() (*IamSecurityHolderRelationship, bool)`

GetHolderOk returns a tuple with the Holder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolder

`func (o *IamResourcePermission) SetHolder(v IamSecurityHolderRelationship)`

SetHolder sets Holder field to given value.

### HasHolder

`func (o *IamResourcePermission) HasHolder() bool`

HasHolder returns a boolean if a field has been set.

### GetResource

`func (o *IamResourcePermission) GetResource() MoBaseMoRelationship`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *IamResourcePermission) GetResourceOk() (*MoBaseMoRelationship, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *IamResourcePermission) SetResource(v MoBaseMoRelationship)`

SetResource sets Resource field to given value.

### HasResource

`func (o *IamResourcePermission) HasResource() bool`

HasResource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# IamResourcePermissionAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PermissionRoles** | Pointer to [**[]IamPermissionToRoles**](iam.PermissionToRoles.md) |  | [optional] 
**TargetApp** | Pointer to **string** | Name of the service owning the resource. | [optional] [readonly] 
**Holder** | Pointer to [**IamSecurityHolderRelationship**](iam.SecurityHolder.Relationship.md) |  | [optional] 
**Resource** | Pointer to [**MoBaseMoRelationship**](mo.BaseMo.Relationship.md) |  | [optional] 

## Methods

### NewIamResourcePermissionAllOf

`func NewIamResourcePermissionAllOf() *IamResourcePermissionAllOf`

NewIamResourcePermissionAllOf instantiates a new IamResourcePermissionAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamResourcePermissionAllOfWithDefaults

`func NewIamResourcePermissionAllOfWithDefaults() *IamResourcePermissionAllOf`

NewIamResourcePermissionAllOfWithDefaults instantiates a new IamResourcePermissionAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermissionRoles

`func (o *IamResourcePermissionAllOf) GetPermissionRoles() []IamPermissionToRoles`

GetPermissionRoles returns the PermissionRoles field if non-nil, zero value otherwise.

### GetPermissionRolesOk

`func (o *IamResourcePermissionAllOf) GetPermissionRolesOk() (*[]IamPermissionToRoles, bool)`

GetPermissionRolesOk returns a tuple with the PermissionRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionRoles

`func (o *IamResourcePermissionAllOf) SetPermissionRoles(v []IamPermissionToRoles)`

SetPermissionRoles sets PermissionRoles field to given value.

### HasPermissionRoles

`func (o *IamResourcePermissionAllOf) HasPermissionRoles() bool`

HasPermissionRoles returns a boolean if a field has been set.

### GetTargetApp

`func (o *IamResourcePermissionAllOf) GetTargetApp() string`

GetTargetApp returns the TargetApp field if non-nil, zero value otherwise.

### GetTargetAppOk

`func (o *IamResourcePermissionAllOf) GetTargetAppOk() (*string, bool)`

GetTargetAppOk returns a tuple with the TargetApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetApp

`func (o *IamResourcePermissionAllOf) SetTargetApp(v string)`

SetTargetApp sets TargetApp field to given value.

### HasTargetApp

`func (o *IamResourcePermissionAllOf) HasTargetApp() bool`

HasTargetApp returns a boolean if a field has been set.

### GetHolder

`func (o *IamResourcePermissionAllOf) GetHolder() IamSecurityHolderRelationship`

GetHolder returns the Holder field if non-nil, zero value otherwise.

### GetHolderOk

`func (o *IamResourcePermissionAllOf) GetHolderOk() (*IamSecurityHolderRelationship, bool)`

GetHolderOk returns a tuple with the Holder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolder

`func (o *IamResourcePermissionAllOf) SetHolder(v IamSecurityHolderRelationship)`

SetHolder sets Holder field to given value.

### HasHolder

`func (o *IamResourcePermissionAllOf) HasHolder() bool`

HasHolder returns a boolean if a field has been set.

### GetResource

`func (o *IamResourcePermissionAllOf) GetResource() MoBaseMoRelationship`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *IamResourcePermissionAllOf) GetResourceOk() (*MoBaseMoRelationship, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *IamResourcePermissionAllOf) SetResource(v MoBaseMoRelationship)`

SetResource sets Resource field to given value.

### HasResource

`func (o *IamResourcePermissionAllOf) HasResource() bool`

HasResource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



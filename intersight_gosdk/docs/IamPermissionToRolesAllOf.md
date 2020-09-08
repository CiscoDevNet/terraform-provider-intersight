# IamPermissionToRolesAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Permission** | Pointer to [**CmrfCmRf**](cmrf.CmRf.md) |  | [optional] 
**Roles** | Pointer to [**[]CmrfCmRf**](cmrf.CmRf.md) |  | [optional] 

## Methods

### NewIamPermissionToRolesAllOf

`func NewIamPermissionToRolesAllOf() *IamPermissionToRolesAllOf`

NewIamPermissionToRolesAllOf instantiates a new IamPermissionToRolesAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamPermissionToRolesAllOfWithDefaults

`func NewIamPermissionToRolesAllOfWithDefaults() *IamPermissionToRolesAllOf`

NewIamPermissionToRolesAllOfWithDefaults instantiates a new IamPermissionToRolesAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermission

`func (o *IamPermissionToRolesAllOf) GetPermission() CmrfCmRf`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *IamPermissionToRolesAllOf) GetPermissionOk() (*CmrfCmRf, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *IamPermissionToRolesAllOf) SetPermission(v CmrfCmRf)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *IamPermissionToRolesAllOf) HasPermission() bool`

HasPermission returns a boolean if a field has been set.

### GetRoles

`func (o *IamPermissionToRolesAllOf) GetRoles() []CmrfCmRf`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *IamPermissionToRolesAllOf) GetRolesOk() (*[]CmrfCmRf, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *IamPermissionToRolesAllOf) SetRoles(v []CmrfCmRf)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *IamPermissionToRolesAllOf) HasRoles() bool`

HasRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



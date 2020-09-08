# IamGroupPermissionToRoles

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | Pointer to [**CmrfCmRf**](cmrf.CmRf.md) |  | [optional] 
**Orgs** | Pointer to [**[]CmrfCmRf**](cmrf.CmRf.md) |  | [optional] 

## Methods

### NewIamGroupPermissionToRoles

`func NewIamGroupPermissionToRoles() *IamGroupPermissionToRoles`

NewIamGroupPermissionToRoles instantiates a new IamGroupPermissionToRoles object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamGroupPermissionToRolesWithDefaults

`func NewIamGroupPermissionToRolesWithDefaults() *IamGroupPermissionToRoles`

NewIamGroupPermissionToRolesWithDefaults instantiates a new IamGroupPermissionToRoles object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *IamGroupPermissionToRoles) GetGroup() CmrfCmRf`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *IamGroupPermissionToRoles) GetGroupOk() (*CmrfCmRf, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *IamGroupPermissionToRoles) SetGroup(v CmrfCmRf)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *IamGroupPermissionToRoles) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetOrgs

`func (o *IamGroupPermissionToRoles) GetOrgs() []CmrfCmRf`

GetOrgs returns the Orgs field if non-nil, zero value otherwise.

### GetOrgsOk

`func (o *IamGroupPermissionToRoles) GetOrgsOk() (*[]CmrfCmRf, bool)`

GetOrgsOk returns a tuple with the Orgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgs

`func (o *IamGroupPermissionToRoles) SetOrgs(v []CmrfCmRf)`

SetOrgs sets Orgs field to given value.

### HasOrgs

`func (o *IamGroupPermissionToRoles) HasOrgs() bool`

HasOrgs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



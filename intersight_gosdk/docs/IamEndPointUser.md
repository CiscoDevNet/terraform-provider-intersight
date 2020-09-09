# IamEndPointUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the user to be created on the endpoint. It can be any string that adheres to the following constraints. It can have alphanumeric characters, dots, underscores and hyphen. It cannot be more than 16 characters. | [optional] 
**EndPointUserRole** | Pointer to [**[]IamEndPointUserRoleRelationship**](iam.EndPointUserRole.Relationship.md) | An array of relationships to iamEndPointUserRole resources. | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](organization.Organization.Relationship.md) |  | [optional] 

## Methods

### NewIamEndPointUser

`func NewIamEndPointUser() *IamEndPointUser`

NewIamEndPointUser instantiates a new IamEndPointUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamEndPointUserWithDefaults

`func NewIamEndPointUserWithDefaults() *IamEndPointUser`

NewIamEndPointUserWithDefaults instantiates a new IamEndPointUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *IamEndPointUser) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IamEndPointUser) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IamEndPointUser) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IamEndPointUser) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEndPointUserRole

`func (o *IamEndPointUser) GetEndPointUserRole() []IamEndPointUserRoleRelationship`

GetEndPointUserRole returns the EndPointUserRole field if non-nil, zero value otherwise.

### GetEndPointUserRoleOk

`func (o *IamEndPointUser) GetEndPointUserRoleOk() (*[]IamEndPointUserRoleRelationship, bool)`

GetEndPointUserRoleOk returns a tuple with the EndPointUserRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPointUserRole

`func (o *IamEndPointUser) SetEndPointUserRole(v []IamEndPointUserRoleRelationship)`

SetEndPointUserRole sets EndPointUserRole field to given value.

### HasEndPointUserRole

`func (o *IamEndPointUser) HasEndPointUserRole() bool`

HasEndPointUserRole returns a boolean if a field has been set.

### SetEndPointUserRoleNil

`func (o *IamEndPointUser) SetEndPointUserRoleNil(b bool)`

 SetEndPointUserRoleNil sets the value for EndPointUserRole to be an explicit nil

### UnsetEndPointUserRole
`func (o *IamEndPointUser) UnsetEndPointUserRole()`

UnsetEndPointUserRole ensures that no value is present for EndPointUserRole, not even an explicit nil
### GetOrganization

`func (o *IamEndPointUser) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *IamEndPointUser) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *IamEndPointUser) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *IamEndPointUser) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# IamEndPointUserPolicyAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PasswordProperties** | Pointer to [**IamEndPointPasswordProperties**](iam.EndPointPasswordProperties.md) |  | [optional] 
**EndPointUserRoles** | Pointer to [**[]IamEndPointUserRoleRelationship**](iam.EndPointUserRole.Relationship.md) | An array of relationships to iamEndPointUserRole resources. | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](organization.Organization.Relationship.md) |  | [optional] 
**Profiles** | Pointer to [**[]PolicyAbstractConfigProfileRelationship**](policy.AbstractConfigProfile.Relationship.md) | An array of relationships to policyAbstractConfigProfile resources. | [optional] 

## Methods

### NewIamEndPointUserPolicyAllOf

`func NewIamEndPointUserPolicyAllOf() *IamEndPointUserPolicyAllOf`

NewIamEndPointUserPolicyAllOf instantiates a new IamEndPointUserPolicyAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamEndPointUserPolicyAllOfWithDefaults

`func NewIamEndPointUserPolicyAllOfWithDefaults() *IamEndPointUserPolicyAllOf`

NewIamEndPointUserPolicyAllOfWithDefaults instantiates a new IamEndPointUserPolicyAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPasswordProperties

`func (o *IamEndPointUserPolicyAllOf) GetPasswordProperties() IamEndPointPasswordProperties`

GetPasswordProperties returns the PasswordProperties field if non-nil, zero value otherwise.

### GetPasswordPropertiesOk

`func (o *IamEndPointUserPolicyAllOf) GetPasswordPropertiesOk() (*IamEndPointPasswordProperties, bool)`

GetPasswordPropertiesOk returns a tuple with the PasswordProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordProperties

`func (o *IamEndPointUserPolicyAllOf) SetPasswordProperties(v IamEndPointPasswordProperties)`

SetPasswordProperties sets PasswordProperties field to given value.

### HasPasswordProperties

`func (o *IamEndPointUserPolicyAllOf) HasPasswordProperties() bool`

HasPasswordProperties returns a boolean if a field has been set.

### GetEndPointUserRoles

`func (o *IamEndPointUserPolicyAllOf) GetEndPointUserRoles() []IamEndPointUserRoleRelationship`

GetEndPointUserRoles returns the EndPointUserRoles field if non-nil, zero value otherwise.

### GetEndPointUserRolesOk

`func (o *IamEndPointUserPolicyAllOf) GetEndPointUserRolesOk() (*[]IamEndPointUserRoleRelationship, bool)`

GetEndPointUserRolesOk returns a tuple with the EndPointUserRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPointUserRoles

`func (o *IamEndPointUserPolicyAllOf) SetEndPointUserRoles(v []IamEndPointUserRoleRelationship)`

SetEndPointUserRoles sets EndPointUserRoles field to given value.

### HasEndPointUserRoles

`func (o *IamEndPointUserPolicyAllOf) HasEndPointUserRoles() bool`

HasEndPointUserRoles returns a boolean if a field has been set.

### SetEndPointUserRolesNil

`func (o *IamEndPointUserPolicyAllOf) SetEndPointUserRolesNil(b bool)`

 SetEndPointUserRolesNil sets the value for EndPointUserRoles to be an explicit nil

### UnsetEndPointUserRoles
`func (o *IamEndPointUserPolicyAllOf) UnsetEndPointUserRoles()`

UnsetEndPointUserRoles ensures that no value is present for EndPointUserRoles, not even an explicit nil
### GetOrganization

`func (o *IamEndPointUserPolicyAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *IamEndPointUserPolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *IamEndPointUserPolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *IamEndPointUserPolicyAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetProfiles

`func (o *IamEndPointUserPolicyAllOf) GetProfiles() []PolicyAbstractConfigProfileRelationship`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *IamEndPointUserPolicyAllOf) GetProfilesOk() (*[]PolicyAbstractConfigProfileRelationship, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *IamEndPointUserPolicyAllOf) SetProfiles(v []PolicyAbstractConfigProfileRelationship)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *IamEndPointUserPolicyAllOf) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.

### SetProfilesNil

`func (o *IamEndPointUserPolicyAllOf) SetProfilesNil(b bool)`

 SetProfilesNil sets the value for Profiles to be an explicit nil

### UnsetProfiles
`func (o *IamEndPointUserPolicyAllOf) UnsetProfiles()`

UnsetProfiles ensures that no value is present for Profiles, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



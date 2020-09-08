# AccessPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InbandVlan** | Pointer to **int64** | VLAN to be used for server access over Inband network. | [optional] 
**InbandIpPool** | Pointer to [**IppoolPoolRelationship**](ippool.Pool.Relationship.md) |  | [optional] 
**InbandVrf** | Pointer to [**VrfVrfRelationship**](vrf.Vrf.Relationship.md) |  | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](organization.Organization.Relationship.md) |  | [optional] 
**Profiles** | Pointer to [**[]PolicyAbstractConfigProfileRelationship**](policy.AbstractConfigProfile.Relationship.md) | An array of relationships to policyAbstractConfigProfile resources. | [optional] 

## Methods

### NewAccessPolicy

`func NewAccessPolicy() *AccessPolicy`

NewAccessPolicy instantiates a new AccessPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessPolicyWithDefaults

`func NewAccessPolicyWithDefaults() *AccessPolicy`

NewAccessPolicyWithDefaults instantiates a new AccessPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInbandVlan

`func (o *AccessPolicy) GetInbandVlan() int64`

GetInbandVlan returns the InbandVlan field if non-nil, zero value otherwise.

### GetInbandVlanOk

`func (o *AccessPolicy) GetInbandVlanOk() (*int64, bool)`

GetInbandVlanOk returns a tuple with the InbandVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInbandVlan

`func (o *AccessPolicy) SetInbandVlan(v int64)`

SetInbandVlan sets InbandVlan field to given value.

### HasInbandVlan

`func (o *AccessPolicy) HasInbandVlan() bool`

HasInbandVlan returns a boolean if a field has been set.

### GetInbandIpPool

`func (o *AccessPolicy) GetInbandIpPool() IppoolPoolRelationship`

GetInbandIpPool returns the InbandIpPool field if non-nil, zero value otherwise.

### GetInbandIpPoolOk

`func (o *AccessPolicy) GetInbandIpPoolOk() (*IppoolPoolRelationship, bool)`

GetInbandIpPoolOk returns a tuple with the InbandIpPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInbandIpPool

`func (o *AccessPolicy) SetInbandIpPool(v IppoolPoolRelationship)`

SetInbandIpPool sets InbandIpPool field to given value.

### HasInbandIpPool

`func (o *AccessPolicy) HasInbandIpPool() bool`

HasInbandIpPool returns a boolean if a field has been set.

### GetInbandVrf

`func (o *AccessPolicy) GetInbandVrf() VrfVrfRelationship`

GetInbandVrf returns the InbandVrf field if non-nil, zero value otherwise.

### GetInbandVrfOk

`func (o *AccessPolicy) GetInbandVrfOk() (*VrfVrfRelationship, bool)`

GetInbandVrfOk returns a tuple with the InbandVrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInbandVrf

`func (o *AccessPolicy) SetInbandVrf(v VrfVrfRelationship)`

SetInbandVrf sets InbandVrf field to given value.

### HasInbandVrf

`func (o *AccessPolicy) HasInbandVrf() bool`

HasInbandVrf returns a boolean if a field has been set.

### GetOrganization

`func (o *AccessPolicy) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *AccessPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *AccessPolicy) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *AccessPolicy) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetProfiles

`func (o *AccessPolicy) GetProfiles() []PolicyAbstractConfigProfileRelationship`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *AccessPolicy) GetProfilesOk() (*[]PolicyAbstractConfigProfileRelationship, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *AccessPolicy) SetProfiles(v []PolicyAbstractConfigProfileRelationship)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *AccessPolicy) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.

### SetProfilesNil

`func (o *AccessPolicy) SetProfilesNil(b bool)`

 SetProfilesNil sets the value for Profiles to be an explicit nil

### UnsetProfiles
`func (o *AccessPolicy) UnsetProfiles()`

UnsetProfiles ensures that no value is present for Profiles, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# AccessPolicyAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InbandVlan** | Pointer to **int64** | VLAN to be used for server access over Inband network. | [optional] 
**InbandIpPool** | Pointer to [**IppoolPoolRelationship**](ippool.Pool.Relationship.md) |  | [optional] 
**InbandVrf** | Pointer to [**VrfVrfRelationship**](vrf.Vrf.Relationship.md) |  | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](organization.Organization.Relationship.md) |  | [optional] 
**Profiles** | Pointer to [**[]PolicyAbstractConfigProfileRelationship**](policy.AbstractConfigProfile.Relationship.md) | An array of relationships to policyAbstractConfigProfile resources. | [optional] 

## Methods

### NewAccessPolicyAllOf

`func NewAccessPolicyAllOf() *AccessPolicyAllOf`

NewAccessPolicyAllOf instantiates a new AccessPolicyAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessPolicyAllOfWithDefaults

`func NewAccessPolicyAllOfWithDefaults() *AccessPolicyAllOf`

NewAccessPolicyAllOfWithDefaults instantiates a new AccessPolicyAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInbandVlan

`func (o *AccessPolicyAllOf) GetInbandVlan() int64`

GetInbandVlan returns the InbandVlan field if non-nil, zero value otherwise.

### GetInbandVlanOk

`func (o *AccessPolicyAllOf) GetInbandVlanOk() (*int64, bool)`

GetInbandVlanOk returns a tuple with the InbandVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInbandVlan

`func (o *AccessPolicyAllOf) SetInbandVlan(v int64)`

SetInbandVlan sets InbandVlan field to given value.

### HasInbandVlan

`func (o *AccessPolicyAllOf) HasInbandVlan() bool`

HasInbandVlan returns a boolean if a field has been set.

### GetInbandIpPool

`func (o *AccessPolicyAllOf) GetInbandIpPool() IppoolPoolRelationship`

GetInbandIpPool returns the InbandIpPool field if non-nil, zero value otherwise.

### GetInbandIpPoolOk

`func (o *AccessPolicyAllOf) GetInbandIpPoolOk() (*IppoolPoolRelationship, bool)`

GetInbandIpPoolOk returns a tuple with the InbandIpPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInbandIpPool

`func (o *AccessPolicyAllOf) SetInbandIpPool(v IppoolPoolRelationship)`

SetInbandIpPool sets InbandIpPool field to given value.

### HasInbandIpPool

`func (o *AccessPolicyAllOf) HasInbandIpPool() bool`

HasInbandIpPool returns a boolean if a field has been set.

### GetInbandVrf

`func (o *AccessPolicyAllOf) GetInbandVrf() VrfVrfRelationship`

GetInbandVrf returns the InbandVrf field if non-nil, zero value otherwise.

### GetInbandVrfOk

`func (o *AccessPolicyAllOf) GetInbandVrfOk() (*VrfVrfRelationship, bool)`

GetInbandVrfOk returns a tuple with the InbandVrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInbandVrf

`func (o *AccessPolicyAllOf) SetInbandVrf(v VrfVrfRelationship)`

SetInbandVrf sets InbandVrf field to given value.

### HasInbandVrf

`func (o *AccessPolicyAllOf) HasInbandVrf() bool`

HasInbandVrf returns a boolean if a field has been set.

### GetOrganization

`func (o *AccessPolicyAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *AccessPolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *AccessPolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *AccessPolicyAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetProfiles

`func (o *AccessPolicyAllOf) GetProfiles() []PolicyAbstractConfigProfileRelationship`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *AccessPolicyAllOf) GetProfilesOk() (*[]PolicyAbstractConfigProfileRelationship, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *AccessPolicyAllOf) SetProfiles(v []PolicyAbstractConfigProfileRelationship)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *AccessPolicyAllOf) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.

### SetProfilesNil

`func (o *AccessPolicyAllOf) SetProfilesNil(b bool)`

 SetProfilesNil sets the value for Profiles to be an explicit nil

### UnsetProfiles
`func (o *AccessPolicyAllOf) UnsetProfiles()`

UnsetProfiles ensures that no value is present for Profiles, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



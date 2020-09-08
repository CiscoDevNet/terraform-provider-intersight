# FabricFcNetworkPolicyAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableTrunking** | Pointer to **bool** | Enable or Disable Trunking on all of configured FC uplink ports. | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](organization.Organization.Relationship.md) |  | [optional] 
**Profiles** | Pointer to [**[]FabricSwitchProfileRelationship**](fabric.SwitchProfile.Relationship.md) | An array of relationships to fabricSwitchProfile resources. | [optional] 

## Methods

### NewFabricFcNetworkPolicyAllOf

`func NewFabricFcNetworkPolicyAllOf() *FabricFcNetworkPolicyAllOf`

NewFabricFcNetworkPolicyAllOf instantiates a new FabricFcNetworkPolicyAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricFcNetworkPolicyAllOfWithDefaults

`func NewFabricFcNetworkPolicyAllOfWithDefaults() *FabricFcNetworkPolicyAllOf`

NewFabricFcNetworkPolicyAllOfWithDefaults instantiates a new FabricFcNetworkPolicyAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableTrunking

`func (o *FabricFcNetworkPolicyAllOf) GetEnableTrunking() bool`

GetEnableTrunking returns the EnableTrunking field if non-nil, zero value otherwise.

### GetEnableTrunkingOk

`func (o *FabricFcNetworkPolicyAllOf) GetEnableTrunkingOk() (*bool, bool)`

GetEnableTrunkingOk returns a tuple with the EnableTrunking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableTrunking

`func (o *FabricFcNetworkPolicyAllOf) SetEnableTrunking(v bool)`

SetEnableTrunking sets EnableTrunking field to given value.

### HasEnableTrunking

`func (o *FabricFcNetworkPolicyAllOf) HasEnableTrunking() bool`

HasEnableTrunking returns a boolean if a field has been set.

### GetOrganization

`func (o *FabricFcNetworkPolicyAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *FabricFcNetworkPolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *FabricFcNetworkPolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *FabricFcNetworkPolicyAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetProfiles

`func (o *FabricFcNetworkPolicyAllOf) GetProfiles() []FabricSwitchProfileRelationship`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *FabricFcNetworkPolicyAllOf) GetProfilesOk() (*[]FabricSwitchProfileRelationship, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *FabricFcNetworkPolicyAllOf) SetProfiles(v []FabricSwitchProfileRelationship)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *FabricFcNetworkPolicyAllOf) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.

### SetProfilesNil

`func (o *FabricFcNetworkPolicyAllOf) SetProfilesNil(b bool)`

 SetProfilesNil sets the value for Profiles to be an explicit nil

### UnsetProfiles
`func (o *FabricFcNetworkPolicyAllOf) UnsetProfiles()`

UnsetProfiles ensures that no value is present for Profiles, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



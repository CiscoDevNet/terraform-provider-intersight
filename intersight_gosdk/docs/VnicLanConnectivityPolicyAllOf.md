# VnicLanConnectivityPolicyAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlacementMode** | Pointer to **string** | The mode used for placement of vnics on network adapters. It can either be Auto or Custom. * &#x60;custom&#x60; - The placement of the vNICs / vHBAs on network adapters is manually chosen by the user. * &#x60;auto&#x60; - The placement of the vNICs / vHBAs on network adapters is automatically determined by the system. | [optional] [default to "custom"]
**TargetPlatform** | Pointer to **string** | The platform for which the server profile is applicable. It can either be a server that is operating in standalone mode or which is attached to a Fabric Interconnect managed by Intersight. * &#x60;Standalone&#x60; - Servers which are operating in standalone mode i.e. not connected to a Fabric Interconnected. * &#x60;FIAttached&#x60; - Servers which are connected to a Fabric Interconnect that is managed by Intersight. | [optional] [default to "Standalone"]
**EthIfs** | Pointer to [**[]VnicEthIfRelationship**](vnic.EthIf.Relationship.md) | An array of relationships to vnicEthIf resources. | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](organization.Organization.Relationship.md) |  | [optional] 
**Profiles** | Pointer to [**[]PolicyAbstractConfigProfileRelationship**](policy.AbstractConfigProfile.Relationship.md) | An array of relationships to policyAbstractConfigProfile resources. | [optional] 

## Methods

### NewVnicLanConnectivityPolicyAllOf

`func NewVnicLanConnectivityPolicyAllOf() *VnicLanConnectivityPolicyAllOf`

NewVnicLanConnectivityPolicyAllOf instantiates a new VnicLanConnectivityPolicyAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVnicLanConnectivityPolicyAllOfWithDefaults

`func NewVnicLanConnectivityPolicyAllOfWithDefaults() *VnicLanConnectivityPolicyAllOf`

NewVnicLanConnectivityPolicyAllOfWithDefaults instantiates a new VnicLanConnectivityPolicyAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlacementMode

`func (o *VnicLanConnectivityPolicyAllOf) GetPlacementMode() string`

GetPlacementMode returns the PlacementMode field if non-nil, zero value otherwise.

### GetPlacementModeOk

`func (o *VnicLanConnectivityPolicyAllOf) GetPlacementModeOk() (*string, bool)`

GetPlacementModeOk returns a tuple with the PlacementMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacementMode

`func (o *VnicLanConnectivityPolicyAllOf) SetPlacementMode(v string)`

SetPlacementMode sets PlacementMode field to given value.

### HasPlacementMode

`func (o *VnicLanConnectivityPolicyAllOf) HasPlacementMode() bool`

HasPlacementMode returns a boolean if a field has been set.

### GetTargetPlatform

`func (o *VnicLanConnectivityPolicyAllOf) GetTargetPlatform() string`

GetTargetPlatform returns the TargetPlatform field if non-nil, zero value otherwise.

### GetTargetPlatformOk

`func (o *VnicLanConnectivityPolicyAllOf) GetTargetPlatformOk() (*string, bool)`

GetTargetPlatformOk returns a tuple with the TargetPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetPlatform

`func (o *VnicLanConnectivityPolicyAllOf) SetTargetPlatform(v string)`

SetTargetPlatform sets TargetPlatform field to given value.

### HasTargetPlatform

`func (o *VnicLanConnectivityPolicyAllOf) HasTargetPlatform() bool`

HasTargetPlatform returns a boolean if a field has been set.

### GetEthIfs

`func (o *VnicLanConnectivityPolicyAllOf) GetEthIfs() []VnicEthIfRelationship`

GetEthIfs returns the EthIfs field if non-nil, zero value otherwise.

### GetEthIfsOk

`func (o *VnicLanConnectivityPolicyAllOf) GetEthIfsOk() (*[]VnicEthIfRelationship, bool)`

GetEthIfsOk returns a tuple with the EthIfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthIfs

`func (o *VnicLanConnectivityPolicyAllOf) SetEthIfs(v []VnicEthIfRelationship)`

SetEthIfs sets EthIfs field to given value.

### HasEthIfs

`func (o *VnicLanConnectivityPolicyAllOf) HasEthIfs() bool`

HasEthIfs returns a boolean if a field has been set.

### SetEthIfsNil

`func (o *VnicLanConnectivityPolicyAllOf) SetEthIfsNil(b bool)`

 SetEthIfsNil sets the value for EthIfs to be an explicit nil

### UnsetEthIfs
`func (o *VnicLanConnectivityPolicyAllOf) UnsetEthIfs()`

UnsetEthIfs ensures that no value is present for EthIfs, not even an explicit nil
### GetOrganization

`func (o *VnicLanConnectivityPolicyAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *VnicLanConnectivityPolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *VnicLanConnectivityPolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *VnicLanConnectivityPolicyAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetProfiles

`func (o *VnicLanConnectivityPolicyAllOf) GetProfiles() []PolicyAbstractConfigProfileRelationship`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *VnicLanConnectivityPolicyAllOf) GetProfilesOk() (*[]PolicyAbstractConfigProfileRelationship, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *VnicLanConnectivityPolicyAllOf) SetProfiles(v []PolicyAbstractConfigProfileRelationship)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *VnicLanConnectivityPolicyAllOf) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.

### SetProfilesNil

`func (o *VnicLanConnectivityPolicyAllOf) SetProfilesNil(b bool)`

 SetProfilesNil sets the value for Profiles to be an explicit nil

### UnsetProfiles
`func (o *VnicLanConnectivityPolicyAllOf) UnsetProfiles()`

UnsetProfiles ensures that no value is present for Profiles, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



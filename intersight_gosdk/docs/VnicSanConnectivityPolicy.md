# VnicSanConnectivityPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlacementMode** | Pointer to **string** | The mode used for placement of vnics on network adapters. It can either be Auto or Custom. * &#x60;custom&#x60; - The placement of the vNICs / vHBAs on network adapters is manually chosen by the user. * &#x60;auto&#x60; - The placement of the vNICs / vHBAs on network adapters is automatically determined by the system. | [optional] [default to "custom"]
**TargetPlatform** | Pointer to **string** | The platform for which the server profile is applicable. It can either be a server that is operating in standalone mode or which is attached to a Fabric Interconnect managed by Intersight. * &#x60;Standalone&#x60; - Servers which are operating in standalone mode i.e. not connected to a Fabric Interconnected. * &#x60;FIAttached&#x60; - Servers which are connected to a Fabric Interconnect that is managed by Intersight. | [optional] [default to "Standalone"]
**FcIfs** | Pointer to [**[]VnicFcIfRelationship**](vnic.FcIf.Relationship.md) | An array of relationships to vnicFcIf resources. | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](organization.Organization.Relationship.md) |  | [optional] 
**Profiles** | Pointer to [**[]PolicyAbstractConfigProfileRelationship**](policy.AbstractConfigProfile.Relationship.md) | An array of relationships to policyAbstractConfigProfile resources. | [optional] 
**WwnnPool** | Pointer to [**FcpoolPoolRelationship**](fcpool.Pool.Relationship.md) |  | [optional] 

## Methods

### NewVnicSanConnectivityPolicy

`func NewVnicSanConnectivityPolicy() *VnicSanConnectivityPolicy`

NewVnicSanConnectivityPolicy instantiates a new VnicSanConnectivityPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVnicSanConnectivityPolicyWithDefaults

`func NewVnicSanConnectivityPolicyWithDefaults() *VnicSanConnectivityPolicy`

NewVnicSanConnectivityPolicyWithDefaults instantiates a new VnicSanConnectivityPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlacementMode

`func (o *VnicSanConnectivityPolicy) GetPlacementMode() string`

GetPlacementMode returns the PlacementMode field if non-nil, zero value otherwise.

### GetPlacementModeOk

`func (o *VnicSanConnectivityPolicy) GetPlacementModeOk() (*string, bool)`

GetPlacementModeOk returns a tuple with the PlacementMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacementMode

`func (o *VnicSanConnectivityPolicy) SetPlacementMode(v string)`

SetPlacementMode sets PlacementMode field to given value.

### HasPlacementMode

`func (o *VnicSanConnectivityPolicy) HasPlacementMode() bool`

HasPlacementMode returns a boolean if a field has been set.

### GetTargetPlatform

`func (o *VnicSanConnectivityPolicy) GetTargetPlatform() string`

GetTargetPlatform returns the TargetPlatform field if non-nil, zero value otherwise.

### GetTargetPlatformOk

`func (o *VnicSanConnectivityPolicy) GetTargetPlatformOk() (*string, bool)`

GetTargetPlatformOk returns a tuple with the TargetPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetPlatform

`func (o *VnicSanConnectivityPolicy) SetTargetPlatform(v string)`

SetTargetPlatform sets TargetPlatform field to given value.

### HasTargetPlatform

`func (o *VnicSanConnectivityPolicy) HasTargetPlatform() bool`

HasTargetPlatform returns a boolean if a field has been set.

### GetFcIfs

`func (o *VnicSanConnectivityPolicy) GetFcIfs() []VnicFcIfRelationship`

GetFcIfs returns the FcIfs field if non-nil, zero value otherwise.

### GetFcIfsOk

`func (o *VnicSanConnectivityPolicy) GetFcIfsOk() (*[]VnicFcIfRelationship, bool)`

GetFcIfsOk returns a tuple with the FcIfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcIfs

`func (o *VnicSanConnectivityPolicy) SetFcIfs(v []VnicFcIfRelationship)`

SetFcIfs sets FcIfs field to given value.

### HasFcIfs

`func (o *VnicSanConnectivityPolicy) HasFcIfs() bool`

HasFcIfs returns a boolean if a field has been set.

### SetFcIfsNil

`func (o *VnicSanConnectivityPolicy) SetFcIfsNil(b bool)`

 SetFcIfsNil sets the value for FcIfs to be an explicit nil

### UnsetFcIfs
`func (o *VnicSanConnectivityPolicy) UnsetFcIfs()`

UnsetFcIfs ensures that no value is present for FcIfs, not even an explicit nil
### GetOrganization

`func (o *VnicSanConnectivityPolicy) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *VnicSanConnectivityPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *VnicSanConnectivityPolicy) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *VnicSanConnectivityPolicy) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetProfiles

`func (o *VnicSanConnectivityPolicy) GetProfiles() []PolicyAbstractConfigProfileRelationship`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *VnicSanConnectivityPolicy) GetProfilesOk() (*[]PolicyAbstractConfigProfileRelationship, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *VnicSanConnectivityPolicy) SetProfiles(v []PolicyAbstractConfigProfileRelationship)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *VnicSanConnectivityPolicy) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.

### SetProfilesNil

`func (o *VnicSanConnectivityPolicy) SetProfilesNil(b bool)`

 SetProfilesNil sets the value for Profiles to be an explicit nil

### UnsetProfiles
`func (o *VnicSanConnectivityPolicy) UnsetProfiles()`

UnsetProfiles ensures that no value is present for Profiles, not even an explicit nil
### GetWwnnPool

`func (o *VnicSanConnectivityPolicy) GetWwnnPool() FcpoolPoolRelationship`

GetWwnnPool returns the WwnnPool field if non-nil, zero value otherwise.

### GetWwnnPoolOk

`func (o *VnicSanConnectivityPolicy) GetWwnnPoolOk() (*FcpoolPoolRelationship, bool)`

GetWwnnPoolOk returns a tuple with the WwnnPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWwnnPool

`func (o *VnicSanConnectivityPolicy) SetWwnnPool(v FcpoolPoolRelationship)`

SetWwnnPool sets WwnnPool field to given value.

### HasWwnnPool

`func (o *VnicSanConnectivityPolicy) HasWwnnPool() bool`

HasWwnnPool returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



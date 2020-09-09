# SdwanRouterPolicyAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentSize** | Pointer to **string** | Scale of the SD-WAN router virtual machine deployment. * &#x60;Typical&#x60; - Typical deployment configuration with 4 vCPUs and 4GB RAM. * &#x60;Minimal&#x60; - Minimal deployment configuration with 2 vCPUs and 4GB RAM. | [optional] [default to "Typical"]
**WanCount** | Pointer to **int64** | Number of WAN connections across the SD-WAN site. | [optional] 
**WanTerminationType** | Pointer to **string** | Defines if the WAN networks are singly or dually terminated. Dually terminated WANs are configured on all the SD-WAN routers. Singly terminated WANs are configured only on one of the SD-WAN routers. * &#x60;Single&#x60; - Singly terminated WANs ar evenly distributed across SD-WAN router nodes. A given WAN connection is available only on one of the router nodes. * &#x60;Dual&#x60; - Dually terminated WANs are configured on all the SD-WAN routers. A given WAN connection is available on multiple router nodes. | [optional] [default to "Single"]
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](organization.Organization.Relationship.md) |  | [optional] 
**Profiles** | Pointer to [**[]SdwanProfileRelationship**](sdwan.Profile.Relationship.md) | An array of relationships to sdwanProfile resources. | [optional] 
**SolutionImage** | Pointer to [**SoftwareSolutionDistributableRelationship**](software.SolutionDistributable.Relationship.md) |  | [optional] 

## Methods

### NewSdwanRouterPolicyAllOf

`func NewSdwanRouterPolicyAllOf() *SdwanRouterPolicyAllOf`

NewSdwanRouterPolicyAllOf instantiates a new SdwanRouterPolicyAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSdwanRouterPolicyAllOfWithDefaults

`func NewSdwanRouterPolicyAllOfWithDefaults() *SdwanRouterPolicyAllOf`

NewSdwanRouterPolicyAllOfWithDefaults instantiates a new SdwanRouterPolicyAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentSize

`func (o *SdwanRouterPolicyAllOf) GetDeploymentSize() string`

GetDeploymentSize returns the DeploymentSize field if non-nil, zero value otherwise.

### GetDeploymentSizeOk

`func (o *SdwanRouterPolicyAllOf) GetDeploymentSizeOk() (*string, bool)`

GetDeploymentSizeOk returns a tuple with the DeploymentSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentSize

`func (o *SdwanRouterPolicyAllOf) SetDeploymentSize(v string)`

SetDeploymentSize sets DeploymentSize field to given value.

### HasDeploymentSize

`func (o *SdwanRouterPolicyAllOf) HasDeploymentSize() bool`

HasDeploymentSize returns a boolean if a field has been set.

### GetWanCount

`func (o *SdwanRouterPolicyAllOf) GetWanCount() int64`

GetWanCount returns the WanCount field if non-nil, zero value otherwise.

### GetWanCountOk

`func (o *SdwanRouterPolicyAllOf) GetWanCountOk() (*int64, bool)`

GetWanCountOk returns a tuple with the WanCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanCount

`func (o *SdwanRouterPolicyAllOf) SetWanCount(v int64)`

SetWanCount sets WanCount field to given value.

### HasWanCount

`func (o *SdwanRouterPolicyAllOf) HasWanCount() bool`

HasWanCount returns a boolean if a field has been set.

### GetWanTerminationType

`func (o *SdwanRouterPolicyAllOf) GetWanTerminationType() string`

GetWanTerminationType returns the WanTerminationType field if non-nil, zero value otherwise.

### GetWanTerminationTypeOk

`func (o *SdwanRouterPolicyAllOf) GetWanTerminationTypeOk() (*string, bool)`

GetWanTerminationTypeOk returns a tuple with the WanTerminationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanTerminationType

`func (o *SdwanRouterPolicyAllOf) SetWanTerminationType(v string)`

SetWanTerminationType sets WanTerminationType field to given value.

### HasWanTerminationType

`func (o *SdwanRouterPolicyAllOf) HasWanTerminationType() bool`

HasWanTerminationType returns a boolean if a field has been set.

### GetOrganization

`func (o *SdwanRouterPolicyAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *SdwanRouterPolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *SdwanRouterPolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *SdwanRouterPolicyAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetProfiles

`func (o *SdwanRouterPolicyAllOf) GetProfiles() []SdwanProfileRelationship`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *SdwanRouterPolicyAllOf) GetProfilesOk() (*[]SdwanProfileRelationship, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *SdwanRouterPolicyAllOf) SetProfiles(v []SdwanProfileRelationship)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *SdwanRouterPolicyAllOf) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.

### SetProfilesNil

`func (o *SdwanRouterPolicyAllOf) SetProfilesNil(b bool)`

 SetProfilesNil sets the value for Profiles to be an explicit nil

### UnsetProfiles
`func (o *SdwanRouterPolicyAllOf) UnsetProfiles()`

UnsetProfiles ensures that no value is present for Profiles, not even an explicit nil
### GetSolutionImage

`func (o *SdwanRouterPolicyAllOf) GetSolutionImage() SoftwareSolutionDistributableRelationship`

GetSolutionImage returns the SolutionImage field if non-nil, zero value otherwise.

### GetSolutionImageOk

`func (o *SdwanRouterPolicyAllOf) GetSolutionImageOk() (*SoftwareSolutionDistributableRelationship, bool)`

GetSolutionImageOk returns a tuple with the SolutionImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSolutionImage

`func (o *SdwanRouterPolicyAllOf) SetSolutionImage(v SoftwareSolutionDistributableRelationship)`

SetSolutionImage sets SolutionImage field to given value.

### HasSolutionImage

`func (o *SdwanRouterPolicyAllOf) HasSolutionImage() bool`

HasSolutionImage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



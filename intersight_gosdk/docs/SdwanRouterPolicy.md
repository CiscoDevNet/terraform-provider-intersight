# SdwanRouterPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "sdwan.RouterPolicy"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "sdwan.RouterPolicy"]
**DeploymentSize** | Pointer to **string** | Scale of the SD-WAN router virtual machine deployment. * &#x60;Typical&#x60; - Typical deployment configuration with 4 vCPUs and 4GB RAM. * &#x60;Minimal&#x60; - Minimal deployment configuration with 2 vCPUs and 4GB RAM. | [optional] [default to "Typical"]
**WanCount** | Pointer to **int64** | Number of WAN connections across the SD-WAN site. | [optional] [default to 2]
**WanTerminationType** | Pointer to **string** | Defines if the WAN networks are singly or dually terminated. Dually terminated WANs are configured on all the SD-WAN routers. Singly terminated WANs are configured only on one of the SD-WAN routers. * &#x60;Single&#x60; - Singly terminated WANs ar evenly distributed across SD-WAN router nodes. A given WAN connection is available only on one of the router nodes. * &#x60;Dual&#x60; - Dually terminated WANs are configured on all the SD-WAN routers. A given WAN connection is available on multiple router nodes. | [optional] [default to "Single"]
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 
**Profiles** | Pointer to [**[]SdwanProfileRelationship**](SdwanProfileRelationship.md) | An array of relationships to sdwanProfile resources. | [optional] 
**SolutionImage** | Pointer to [**SoftwareSolutionDistributableRelationship**](SoftwareSolutionDistributableRelationship.md) |  | [optional] 

## Methods

### NewSdwanRouterPolicy

`func NewSdwanRouterPolicy(classId string, objectType string, ) *SdwanRouterPolicy`

NewSdwanRouterPolicy instantiates a new SdwanRouterPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSdwanRouterPolicyWithDefaults

`func NewSdwanRouterPolicyWithDefaults() *SdwanRouterPolicy`

NewSdwanRouterPolicyWithDefaults instantiates a new SdwanRouterPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *SdwanRouterPolicy) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *SdwanRouterPolicy) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *SdwanRouterPolicy) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *SdwanRouterPolicy) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *SdwanRouterPolicy) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *SdwanRouterPolicy) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDeploymentSize

`func (o *SdwanRouterPolicy) GetDeploymentSize() string`

GetDeploymentSize returns the DeploymentSize field if non-nil, zero value otherwise.

### GetDeploymentSizeOk

`func (o *SdwanRouterPolicy) GetDeploymentSizeOk() (*string, bool)`

GetDeploymentSizeOk returns a tuple with the DeploymentSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentSize

`func (o *SdwanRouterPolicy) SetDeploymentSize(v string)`

SetDeploymentSize sets DeploymentSize field to given value.

### HasDeploymentSize

`func (o *SdwanRouterPolicy) HasDeploymentSize() bool`

HasDeploymentSize returns a boolean if a field has been set.

### GetWanCount

`func (o *SdwanRouterPolicy) GetWanCount() int64`

GetWanCount returns the WanCount field if non-nil, zero value otherwise.

### GetWanCountOk

`func (o *SdwanRouterPolicy) GetWanCountOk() (*int64, bool)`

GetWanCountOk returns a tuple with the WanCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanCount

`func (o *SdwanRouterPolicy) SetWanCount(v int64)`

SetWanCount sets WanCount field to given value.

### HasWanCount

`func (o *SdwanRouterPolicy) HasWanCount() bool`

HasWanCount returns a boolean if a field has been set.

### GetWanTerminationType

`func (o *SdwanRouterPolicy) GetWanTerminationType() string`

GetWanTerminationType returns the WanTerminationType field if non-nil, zero value otherwise.

### GetWanTerminationTypeOk

`func (o *SdwanRouterPolicy) GetWanTerminationTypeOk() (*string, bool)`

GetWanTerminationTypeOk returns a tuple with the WanTerminationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanTerminationType

`func (o *SdwanRouterPolicy) SetWanTerminationType(v string)`

SetWanTerminationType sets WanTerminationType field to given value.

### HasWanTerminationType

`func (o *SdwanRouterPolicy) HasWanTerminationType() bool`

HasWanTerminationType returns a boolean if a field has been set.

### GetOrganization

`func (o *SdwanRouterPolicy) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *SdwanRouterPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *SdwanRouterPolicy) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *SdwanRouterPolicy) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetProfiles

`func (o *SdwanRouterPolicy) GetProfiles() []SdwanProfileRelationship`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *SdwanRouterPolicy) GetProfilesOk() (*[]SdwanProfileRelationship, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *SdwanRouterPolicy) SetProfiles(v []SdwanProfileRelationship)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *SdwanRouterPolicy) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.

### SetProfilesNil

`func (o *SdwanRouterPolicy) SetProfilesNil(b bool)`

 SetProfilesNil sets the value for Profiles to be an explicit nil

### UnsetProfiles
`func (o *SdwanRouterPolicy) UnsetProfiles()`

UnsetProfiles ensures that no value is present for Profiles, not even an explicit nil
### GetSolutionImage

`func (o *SdwanRouterPolicy) GetSolutionImage() SoftwareSolutionDistributableRelationship`

GetSolutionImage returns the SolutionImage field if non-nil, zero value otherwise.

### GetSolutionImageOk

`func (o *SdwanRouterPolicy) GetSolutionImageOk() (*SoftwareSolutionDistributableRelationship, bool)`

GetSolutionImageOk returns a tuple with the SolutionImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSolutionImage

`func (o *SdwanRouterPolicy) SetSolutionImage(v SoftwareSolutionDistributableRelationship)`

SetSolutionImage sets SolutionImage field to given value.

### HasSolutionImage

`func (o *SdwanRouterPolicy) HasSolutionImage() bool`

HasSolutionImage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



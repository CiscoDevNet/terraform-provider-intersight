# FabricSwitchClusterProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "fabric.SwitchClusterProfile"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "fabric.SwitchClusterProfile"]
**ClusterAssignments** | Pointer to [**[]FabricClusterAssignment**](FabricClusterAssignment.md) |  | [optional] 
**ConfigContext** | Pointer to [**NullablePolicyConfigContext**](PolicyConfigContext.md) |  | [optional] 
**DeployStatus** | Pointer to **string** | Deploy status of the switch cluster profile indicating if deployment has been initiated on all the members of the cluster profile. * &#x60;None&#x60; - Switch profiles not deployed on either of the switches. * &#x60;Complete&#x60; - Both switch profiles of the cluster profile are deployed. * &#x60;Partial&#x60; - Only one of the switch profiles of the cluster profile is deployed. | [optional] [readonly] [default to "None"]
**DeployedSwitches** | Pointer to **string** | Values indicating the switches on which the cluster profile has been deployed. 0 indicates that the profile has not been deployed on any switch, 1 indicates that the profile has been deployed on A, 2 indicates that it is deployed on B and 3 indicates that it is deployed on both. * &#x60;None&#x60; - Switch profiles not deployed on either of the fabric interconnects. * &#x60;A&#x60; - Switch profiles deployed only on fabric interconnect A. * &#x60;B&#x60; - Switch profiles deployed only on fabric interconnect B. * &#x60;AB&#x60; - Switch profiles deployed on both fabric interconnect A and B. | [optional] [readonly] [default to "None"]
**UserLabel** | Pointer to **string** | The user defined label assigned to the switch profile. | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 
**SwitchProfiles** | Pointer to [**[]FabricSwitchProfileRelationship**](FabricSwitchProfileRelationship.md) | An array of relationships to fabricSwitchProfile resources. | [optional] 

## Methods

### NewFabricSwitchClusterProfile

`func NewFabricSwitchClusterProfile(classId string, objectType string, ) *FabricSwitchClusterProfile`

NewFabricSwitchClusterProfile instantiates a new FabricSwitchClusterProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricSwitchClusterProfileWithDefaults

`func NewFabricSwitchClusterProfileWithDefaults() *FabricSwitchClusterProfile`

NewFabricSwitchClusterProfileWithDefaults instantiates a new FabricSwitchClusterProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FabricSwitchClusterProfile) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FabricSwitchClusterProfile) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FabricSwitchClusterProfile) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FabricSwitchClusterProfile) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FabricSwitchClusterProfile) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FabricSwitchClusterProfile) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetClusterAssignments

`func (o *FabricSwitchClusterProfile) GetClusterAssignments() []FabricClusterAssignment`

GetClusterAssignments returns the ClusterAssignments field if non-nil, zero value otherwise.

### GetClusterAssignmentsOk

`func (o *FabricSwitchClusterProfile) GetClusterAssignmentsOk() (*[]FabricClusterAssignment, bool)`

GetClusterAssignmentsOk returns a tuple with the ClusterAssignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterAssignments

`func (o *FabricSwitchClusterProfile) SetClusterAssignments(v []FabricClusterAssignment)`

SetClusterAssignments sets ClusterAssignments field to given value.

### HasClusterAssignments

`func (o *FabricSwitchClusterProfile) HasClusterAssignments() bool`

HasClusterAssignments returns a boolean if a field has been set.

### SetClusterAssignmentsNil

`func (o *FabricSwitchClusterProfile) SetClusterAssignmentsNil(b bool)`

 SetClusterAssignmentsNil sets the value for ClusterAssignments to be an explicit nil

### UnsetClusterAssignments
`func (o *FabricSwitchClusterProfile) UnsetClusterAssignments()`

UnsetClusterAssignments ensures that no value is present for ClusterAssignments, not even an explicit nil
### GetConfigContext

`func (o *FabricSwitchClusterProfile) GetConfigContext() PolicyConfigContext`

GetConfigContext returns the ConfigContext field if non-nil, zero value otherwise.

### GetConfigContextOk

`func (o *FabricSwitchClusterProfile) GetConfigContextOk() (*PolicyConfigContext, bool)`

GetConfigContextOk returns a tuple with the ConfigContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigContext

`func (o *FabricSwitchClusterProfile) SetConfigContext(v PolicyConfigContext)`

SetConfigContext sets ConfigContext field to given value.

### HasConfigContext

`func (o *FabricSwitchClusterProfile) HasConfigContext() bool`

HasConfigContext returns a boolean if a field has been set.

### SetConfigContextNil

`func (o *FabricSwitchClusterProfile) SetConfigContextNil(b bool)`

 SetConfigContextNil sets the value for ConfigContext to be an explicit nil

### UnsetConfigContext
`func (o *FabricSwitchClusterProfile) UnsetConfigContext()`

UnsetConfigContext ensures that no value is present for ConfigContext, not even an explicit nil
### GetDeployStatus

`func (o *FabricSwitchClusterProfile) GetDeployStatus() string`

GetDeployStatus returns the DeployStatus field if non-nil, zero value otherwise.

### GetDeployStatusOk

`func (o *FabricSwitchClusterProfile) GetDeployStatusOk() (*string, bool)`

GetDeployStatusOk returns a tuple with the DeployStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployStatus

`func (o *FabricSwitchClusterProfile) SetDeployStatus(v string)`

SetDeployStatus sets DeployStatus field to given value.

### HasDeployStatus

`func (o *FabricSwitchClusterProfile) HasDeployStatus() bool`

HasDeployStatus returns a boolean if a field has been set.

### GetDeployedSwitches

`func (o *FabricSwitchClusterProfile) GetDeployedSwitches() string`

GetDeployedSwitches returns the DeployedSwitches field if non-nil, zero value otherwise.

### GetDeployedSwitchesOk

`func (o *FabricSwitchClusterProfile) GetDeployedSwitchesOk() (*string, bool)`

GetDeployedSwitchesOk returns a tuple with the DeployedSwitches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployedSwitches

`func (o *FabricSwitchClusterProfile) SetDeployedSwitches(v string)`

SetDeployedSwitches sets DeployedSwitches field to given value.

### HasDeployedSwitches

`func (o *FabricSwitchClusterProfile) HasDeployedSwitches() bool`

HasDeployedSwitches returns a boolean if a field has been set.

### GetUserLabel

`func (o *FabricSwitchClusterProfile) GetUserLabel() string`

GetUserLabel returns the UserLabel field if non-nil, zero value otherwise.

### GetUserLabelOk

`func (o *FabricSwitchClusterProfile) GetUserLabelOk() (*string, bool)`

GetUserLabelOk returns a tuple with the UserLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLabel

`func (o *FabricSwitchClusterProfile) SetUserLabel(v string)`

SetUserLabel sets UserLabel field to given value.

### HasUserLabel

`func (o *FabricSwitchClusterProfile) HasUserLabel() bool`

HasUserLabel returns a boolean if a field has been set.

### GetOrganization

`func (o *FabricSwitchClusterProfile) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *FabricSwitchClusterProfile) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *FabricSwitchClusterProfile) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *FabricSwitchClusterProfile) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetSwitchProfiles

`func (o *FabricSwitchClusterProfile) GetSwitchProfiles() []FabricSwitchProfileRelationship`

GetSwitchProfiles returns the SwitchProfiles field if non-nil, zero value otherwise.

### GetSwitchProfilesOk

`func (o *FabricSwitchClusterProfile) GetSwitchProfilesOk() (*[]FabricSwitchProfileRelationship, bool)`

GetSwitchProfilesOk returns a tuple with the SwitchProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchProfiles

`func (o *FabricSwitchClusterProfile) SetSwitchProfiles(v []FabricSwitchProfileRelationship)`

SetSwitchProfiles sets SwitchProfiles field to given value.

### HasSwitchProfiles

`func (o *FabricSwitchClusterProfile) HasSwitchProfiles() bool`

HasSwitchProfiles returns a boolean if a field has been set.

### SetSwitchProfilesNil

`func (o *FabricSwitchClusterProfile) SetSwitchProfilesNil(b bool)`

 SetSwitchProfilesNil sets the value for SwitchProfiles to be an explicit nil

### UnsetSwitchProfiles
`func (o *FabricSwitchClusterProfile) UnsetSwitchProfiles()`

UnsetSwitchProfiles ensures that no value is present for SwitchProfiles, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



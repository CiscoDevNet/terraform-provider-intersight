# PowerPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "power.Policy"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "power.Policy"]
**AllocatedBudget** | Pointer to **int64** | Sets the Allocated Power Budget of the System (in Watts). This field is only supported for Cisco UCS X series Chassis. | [optional] [default to 0]
**PowerProfiling** | Pointer to **string** | Sets the Power Profiling of the Server. This field is only supported for Cisco UCS X series servers. * &#x60;Enabled&#x60; - Set the value to Enabled. * &#x60;Disabled&#x60; - Set the value to Disabled. | [optional] [default to "Enabled"]
**PowerRestoreState** | Pointer to **string** | Sets the Power Restore State of the Server. * &#x60;AlwaysOff&#x60; - Set the Power Restore Mode to Off. * &#x60;AlwaysOn&#x60; - Set the Power Restore Mode to On. * &#x60;LastState&#x60; - Set the Power Restore Mode to LastState. | [optional] [default to "AlwaysOff"]
**RedundancyMode** | Pointer to **string** | Sets the Power Redundancy of the System. N+2 mode is only supported for Cisco UCS X series Chassis. * &#x60;Grid&#x60; - Grid Mode requires two power sources. If one source fails, the surviving PSUs connected to the other source provides power to the chassis. * &#x60;NotRedundant&#x60; - Power Manager turns on the minimum number of PSUs required to support chassis power requirements. No Redundant PSUs are maintained. * &#x60;N+1&#x60; - Power Manager turns on the minimum number of PSUs required to support chassis power requirements plus one additional PSU for redundancy. * &#x60;N+2&#x60; - Power Manager turns on the minimum number of PSUs required to support chassis power requirements plus two additional PSU for redundancy. This Mode is only supported for UCS X series Chassis. | [optional] [default to "Grid"]
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](organization.Organization.Relationship.md) |  | [optional] 
**Profiles** | Pointer to [**[]PolicyAbstractConfigProfileRelationship**](PolicyAbstractConfigProfileRelationship.md) | An array of relationships to policyAbstractConfigProfile resources. | [optional] 

## Methods

### NewPowerPolicy

`func NewPowerPolicy(classId string, objectType string, ) *PowerPolicy`

NewPowerPolicy instantiates a new PowerPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPowerPolicyWithDefaults

`func NewPowerPolicyWithDefaults() *PowerPolicy`

NewPowerPolicyWithDefaults instantiates a new PowerPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *PowerPolicy) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *PowerPolicy) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *PowerPolicy) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *PowerPolicy) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PowerPolicy) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PowerPolicy) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAllocatedBudget

`func (o *PowerPolicy) GetAllocatedBudget() int64`

GetAllocatedBudget returns the AllocatedBudget field if non-nil, zero value otherwise.

### GetAllocatedBudgetOk

`func (o *PowerPolicy) GetAllocatedBudgetOk() (*int64, bool)`

GetAllocatedBudgetOk returns a tuple with the AllocatedBudget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedBudget

`func (o *PowerPolicy) SetAllocatedBudget(v int64)`

SetAllocatedBudget sets AllocatedBudget field to given value.

### HasAllocatedBudget

`func (o *PowerPolicy) HasAllocatedBudget() bool`

HasAllocatedBudget returns a boolean if a field has been set.

### GetPowerProfiling

`func (o *PowerPolicy) GetPowerProfiling() string`

GetPowerProfiling returns the PowerProfiling field if non-nil, zero value otherwise.

### GetPowerProfilingOk

`func (o *PowerPolicy) GetPowerProfilingOk() (*string, bool)`

GetPowerProfilingOk returns a tuple with the PowerProfiling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerProfiling

`func (o *PowerPolicy) SetPowerProfiling(v string)`

SetPowerProfiling sets PowerProfiling field to given value.

### HasPowerProfiling

`func (o *PowerPolicy) HasPowerProfiling() bool`

HasPowerProfiling returns a boolean if a field has been set.

### GetPowerRestoreState

`func (o *PowerPolicy) GetPowerRestoreState() string`

GetPowerRestoreState returns the PowerRestoreState field if non-nil, zero value otherwise.

### GetPowerRestoreStateOk

`func (o *PowerPolicy) GetPowerRestoreStateOk() (*string, bool)`

GetPowerRestoreStateOk returns a tuple with the PowerRestoreState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerRestoreState

`func (o *PowerPolicy) SetPowerRestoreState(v string)`

SetPowerRestoreState sets PowerRestoreState field to given value.

### HasPowerRestoreState

`func (o *PowerPolicy) HasPowerRestoreState() bool`

HasPowerRestoreState returns a boolean if a field has been set.

### GetRedundancyMode

`func (o *PowerPolicy) GetRedundancyMode() string`

GetRedundancyMode returns the RedundancyMode field if non-nil, zero value otherwise.

### GetRedundancyModeOk

`func (o *PowerPolicy) GetRedundancyModeOk() (*string, bool)`

GetRedundancyModeOk returns a tuple with the RedundancyMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedundancyMode

`func (o *PowerPolicy) SetRedundancyMode(v string)`

SetRedundancyMode sets RedundancyMode field to given value.

### HasRedundancyMode

`func (o *PowerPolicy) HasRedundancyMode() bool`

HasRedundancyMode returns a boolean if a field has been set.

### GetOrganization

`func (o *PowerPolicy) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *PowerPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *PowerPolicy) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *PowerPolicy) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetProfiles

`func (o *PowerPolicy) GetProfiles() []PolicyAbstractConfigProfileRelationship`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *PowerPolicy) GetProfilesOk() (*[]PolicyAbstractConfigProfileRelationship, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *PowerPolicy) SetProfiles(v []PolicyAbstractConfigProfileRelationship)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *PowerPolicy) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.

### SetProfilesNil

`func (o *PowerPolicy) SetProfilesNil(b bool)`

 SetProfilesNil sets the value for Profiles to be an explicit nil

### UnsetProfiles
`func (o *PowerPolicy) UnsetProfiles()`

UnsetProfiles ensures that no value is present for Profiles, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



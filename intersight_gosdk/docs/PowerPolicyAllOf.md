# PowerPolicyAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "power.Policy"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "power.Policy"]
**AllocatedBudget** | Pointer to **int64** | Sets the allocated power budget of the chassis (in Watts). | [optional] [default to 0]
**DynamicRebalancing** | Pointer to **string** | Sets the dynamic power rebalancing mode of the chassis. If enabled, this mode allows the chassis to dynamically reallocate the power between servers depending on their power usage. * &#x60;Enabled&#x60; - Set the value to Enabled. * &#x60;Disabled&#x60; - Set the value to Disabled. | [optional] [default to "Enabled"]
**ExtendedPowerCapacity** | Pointer to **string** | Sets the Extended Power Capacity of the Chassis. If Enabled, this mode allows chassis available power to be increased by borrowing power from redundant power supplies.  This option is only supported for Cisco UCS X series Chassis. * &#x60;Enabled&#x60; - Set the value to Enabled. * &#x60;Disabled&#x60; - Set the value to Disabled. | [optional] [default to "Enabled"]
**PowerPriority** | Pointer to **string** | Sets the Power Priority of the Server. This priority is used to determine the initial power allocation for servers. This field is only supported for Cisco UCS B series and X series servers. * &#x60;Low&#x60; - Set the Power Priority to Low. * &#x60;Medium&#x60; - Set the Power Priority to Medium. * &#x60;High&#x60; - Set the Power Priority to High. | [optional] [default to "Low"]
**PowerProfiling** | Pointer to **string** | Sets the Power Profiling of the Server. If Enabled, this field allows the power manager to run power profiling  utility to determine the power needs of the server.  This field is only supported for Cisco UCS X series servers. * &#x60;Enabled&#x60; - Set the value to Enabled. * &#x60;Disabled&#x60; - Set the value to Disabled. | [optional] [default to "Enabled"]
**PowerRestoreState** | Pointer to **string** | Sets the Power Restore State of the Server. In the absence of Intersight connectivity, the chassis will use this policy  to recover the host power after a power loss event.  This field is only supported for Cisco UCS B series and X series servers. * &#x60;AlwaysOff&#x60; - Set the Power Restore Mode to Off. * &#x60;AlwaysOn&#x60; - Set the Power Restore Mode to On. * &#x60;LastState&#x60; - Set the Power Restore Mode to LastState. | [optional] [default to "AlwaysOff"]
**PowerSaveMode** | Pointer to **string** | Sets the power save mode of the chassis. If the requested power budget is less than available power capacity,  the additional PSUs not required to comply with redundancy policy are placed in power save mode. * &#x60;Enabled&#x60; - Set the value to Enabled. * &#x60;Disabled&#x60; - Set the value to Disabled. | [optional] [default to "Enabled"]
**RedundancyMode** | Pointer to **string** | Sets the Power Redundancy Mode of the Chassis.  Redundancy Mode determines the number of PSUs the chassis keeps as redundant.  N+2 mode is only supported for Cisco UCS X series Chassis. * &#x60;Grid&#x60; - Grid Mode requires two power sources. If one source fails, the surviving PSUs connected to the other source provides power to the chassis. * &#x60;NotRedundant&#x60; - Power Manager turns on the minimum number of PSUs required to support chassis power requirements. No Redundant PSUs are maintained. * &#x60;N+1&#x60; - Power Manager turns on the minimum number of PSUs required to support chassis power requirements plus one additional PSU for redundancy. * &#x60;N+2&#x60; - Power Manager turns on the minimum number of PSUs required to support chassis power requirements plus two additional PSU for redundancy. This Mode is only supported for UCS X series Chassis. | [optional] [default to "Grid"]
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 
**Profiles** | Pointer to [**[]PolicyAbstractConfigProfileRelationship**](PolicyAbstractConfigProfileRelationship.md) | An array of relationships to policyAbstractConfigProfile resources. | [optional] 

## Methods

### NewPowerPolicyAllOf

`func NewPowerPolicyAllOf(classId string, objectType string, ) *PowerPolicyAllOf`

NewPowerPolicyAllOf instantiates a new PowerPolicyAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPowerPolicyAllOfWithDefaults

`func NewPowerPolicyAllOfWithDefaults() *PowerPolicyAllOf`

NewPowerPolicyAllOfWithDefaults instantiates a new PowerPolicyAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *PowerPolicyAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *PowerPolicyAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *PowerPolicyAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *PowerPolicyAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PowerPolicyAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PowerPolicyAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAllocatedBudget

`func (o *PowerPolicyAllOf) GetAllocatedBudget() int64`

GetAllocatedBudget returns the AllocatedBudget field if non-nil, zero value otherwise.

### GetAllocatedBudgetOk

`func (o *PowerPolicyAllOf) GetAllocatedBudgetOk() (*int64, bool)`

GetAllocatedBudgetOk returns a tuple with the AllocatedBudget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedBudget

`func (o *PowerPolicyAllOf) SetAllocatedBudget(v int64)`

SetAllocatedBudget sets AllocatedBudget field to given value.

### HasAllocatedBudget

`func (o *PowerPolicyAllOf) HasAllocatedBudget() bool`

HasAllocatedBudget returns a boolean if a field has been set.

### GetDynamicRebalancing

`func (o *PowerPolicyAllOf) GetDynamicRebalancing() string`

GetDynamicRebalancing returns the DynamicRebalancing field if non-nil, zero value otherwise.

### GetDynamicRebalancingOk

`func (o *PowerPolicyAllOf) GetDynamicRebalancingOk() (*string, bool)`

GetDynamicRebalancingOk returns a tuple with the DynamicRebalancing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicRebalancing

`func (o *PowerPolicyAllOf) SetDynamicRebalancing(v string)`

SetDynamicRebalancing sets DynamicRebalancing field to given value.

### HasDynamicRebalancing

`func (o *PowerPolicyAllOf) HasDynamicRebalancing() bool`

HasDynamicRebalancing returns a boolean if a field has been set.

### GetExtendedPowerCapacity

`func (o *PowerPolicyAllOf) GetExtendedPowerCapacity() string`

GetExtendedPowerCapacity returns the ExtendedPowerCapacity field if non-nil, zero value otherwise.

### GetExtendedPowerCapacityOk

`func (o *PowerPolicyAllOf) GetExtendedPowerCapacityOk() (*string, bool)`

GetExtendedPowerCapacityOk returns a tuple with the ExtendedPowerCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedPowerCapacity

`func (o *PowerPolicyAllOf) SetExtendedPowerCapacity(v string)`

SetExtendedPowerCapacity sets ExtendedPowerCapacity field to given value.

### HasExtendedPowerCapacity

`func (o *PowerPolicyAllOf) HasExtendedPowerCapacity() bool`

HasExtendedPowerCapacity returns a boolean if a field has been set.

### GetPowerPriority

`func (o *PowerPolicyAllOf) GetPowerPriority() string`

GetPowerPriority returns the PowerPriority field if non-nil, zero value otherwise.

### GetPowerPriorityOk

`func (o *PowerPolicyAllOf) GetPowerPriorityOk() (*string, bool)`

GetPowerPriorityOk returns a tuple with the PowerPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerPriority

`func (o *PowerPolicyAllOf) SetPowerPriority(v string)`

SetPowerPriority sets PowerPriority field to given value.

### HasPowerPriority

`func (o *PowerPolicyAllOf) HasPowerPriority() bool`

HasPowerPriority returns a boolean if a field has been set.

### GetPowerProfiling

`func (o *PowerPolicyAllOf) GetPowerProfiling() string`

GetPowerProfiling returns the PowerProfiling field if non-nil, zero value otherwise.

### GetPowerProfilingOk

`func (o *PowerPolicyAllOf) GetPowerProfilingOk() (*string, bool)`

GetPowerProfilingOk returns a tuple with the PowerProfiling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerProfiling

`func (o *PowerPolicyAllOf) SetPowerProfiling(v string)`

SetPowerProfiling sets PowerProfiling field to given value.

### HasPowerProfiling

`func (o *PowerPolicyAllOf) HasPowerProfiling() bool`

HasPowerProfiling returns a boolean if a field has been set.

### GetPowerRestoreState

`func (o *PowerPolicyAllOf) GetPowerRestoreState() string`

GetPowerRestoreState returns the PowerRestoreState field if non-nil, zero value otherwise.

### GetPowerRestoreStateOk

`func (o *PowerPolicyAllOf) GetPowerRestoreStateOk() (*string, bool)`

GetPowerRestoreStateOk returns a tuple with the PowerRestoreState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerRestoreState

`func (o *PowerPolicyAllOf) SetPowerRestoreState(v string)`

SetPowerRestoreState sets PowerRestoreState field to given value.

### HasPowerRestoreState

`func (o *PowerPolicyAllOf) HasPowerRestoreState() bool`

HasPowerRestoreState returns a boolean if a field has been set.

### GetPowerSaveMode

`func (o *PowerPolicyAllOf) GetPowerSaveMode() string`

GetPowerSaveMode returns the PowerSaveMode field if non-nil, zero value otherwise.

### GetPowerSaveModeOk

`func (o *PowerPolicyAllOf) GetPowerSaveModeOk() (*string, bool)`

GetPowerSaveModeOk returns a tuple with the PowerSaveMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerSaveMode

`func (o *PowerPolicyAllOf) SetPowerSaveMode(v string)`

SetPowerSaveMode sets PowerSaveMode field to given value.

### HasPowerSaveMode

`func (o *PowerPolicyAllOf) HasPowerSaveMode() bool`

HasPowerSaveMode returns a boolean if a field has been set.

### GetRedundancyMode

`func (o *PowerPolicyAllOf) GetRedundancyMode() string`

GetRedundancyMode returns the RedundancyMode field if non-nil, zero value otherwise.

### GetRedundancyModeOk

`func (o *PowerPolicyAllOf) GetRedundancyModeOk() (*string, bool)`

GetRedundancyModeOk returns a tuple with the RedundancyMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedundancyMode

`func (o *PowerPolicyAllOf) SetRedundancyMode(v string)`

SetRedundancyMode sets RedundancyMode field to given value.

### HasRedundancyMode

`func (o *PowerPolicyAllOf) HasRedundancyMode() bool`

HasRedundancyMode returns a boolean if a field has been set.

### GetOrganization

`func (o *PowerPolicyAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *PowerPolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *PowerPolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *PowerPolicyAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetProfiles

`func (o *PowerPolicyAllOf) GetProfiles() []PolicyAbstractConfigProfileRelationship`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *PowerPolicyAllOf) GetProfilesOk() (*[]PolicyAbstractConfigProfileRelationship, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *PowerPolicyAllOf) SetProfiles(v []PolicyAbstractConfigProfileRelationship)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *PowerPolicyAllOf) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.

### SetProfilesNil

`func (o *PowerPolicyAllOf) SetProfilesNil(b bool)`

 SetProfilesNil sets the value for Profiles to be an explicit nil

### UnsetProfiles
`func (o *PowerPolicyAllOf) UnsetProfiles()`

UnsetProfiles ensures that no value is present for Profiles, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



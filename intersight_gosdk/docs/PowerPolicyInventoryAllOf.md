# PowerPolicyInventoryAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "power.PolicyInventory"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "power.PolicyInventory"]
**AllocatedBudget** | Pointer to **int64** | Sets the Allocated Power Budget of the Chassis (in Watts). This field is only supported for Cisco UCS X series Chassis. | [optional] [readonly] [default to 0]
**DynamicRebalancing** | Pointer to **string** | Sets the Dynamic Power Rebalancing mode of the Chassis. If enabled, this mode allows the chassis to dynamically reallocate the power between servers depending on their power usage. This option is only supported for Cisco UCS X series Chassis. * &#x60;Enabled&#x60; - Set the value to Enabled. * &#x60;Disabled&#x60; - Set the value to Disabled. | [optional] [readonly] [default to "Enabled"]
**ExtendedPowerCapacity** | Pointer to **string** | Sets the Extended Power Capacity of the Chassis. If Enabled, this mode allows chassis available power to be increased by borrowing power from redundant power supplies.  This option is only supported for Cisco UCS X series Chassis. * &#x60;Enabled&#x60; - Set the value to Enabled. * &#x60;Disabled&#x60; - Set the value to Disabled. | [optional] [readonly] [default to "Enabled"]
**PowerPriority** | Pointer to **string** | Sets the Power Priority of the Server. This priority is used to determine the initial power allocation for servers. This field is only supported for Cisco UCS B series and X series servers. * &#x60;Low&#x60; - Set the Power Priority to Low. * &#x60;Medium&#x60; - Set the Power Priority to Medium. * &#x60;High&#x60; - Set the Power Priority to High. | [optional] [readonly] [default to "Low"]
**PowerProfiling** | Pointer to **string** | Sets the Power Profiling of the Server. If Enabled, this field allows the power manager to run power profiling  utility to determine the power needs of the server.  This field is only supported for Cisco UCS X series servers. * &#x60;Enabled&#x60; - Set the value to Enabled. * &#x60;Disabled&#x60; - Set the value to Disabled. | [optional] [readonly] [default to "Enabled"]
**PowerRestoreState** | Pointer to **string** | Sets the Power Restore State of the Server. In the absence of Intersight connectivity, the chassis will use this policy  to recover the host power after a power loss event.  This field is only supported for Cisco UCS B series and X series servers. * &#x60;AlwaysOff&#x60; - Set the Power Restore Mode to Off. * &#x60;AlwaysOn&#x60; - Set the Power Restore Mode to On. * &#x60;LastState&#x60; - Set the Power Restore Mode to LastState. | [optional] [readonly] [default to "AlwaysOff"]
**PowerSaveMode** | Pointer to **string** | Sets the Power Save mode of the Chassis. If the requested power budget is less than available power capacity,  the additional PSUs not required to comply with redundancy policy are placed in Power Save mode. This option is only supported for Cisco UCS X series Chassis. * &#x60;Enabled&#x60; - Set the value to Enabled. * &#x60;Disabled&#x60; - Set the value to Disabled. | [optional] [readonly] [default to "Enabled"]
**RedundancyMode** | Pointer to **string** | Sets the Power Redundancy Mode of the Chassis.  Redundancy Mode determines the number of PSUs the chassis keeps as redundant.  N+2 mode is only supported for Cisco UCS X series Chassis. * &#x60;Grid&#x60; - Grid Mode requires two power sources. If one source fails, the surviving PSUs connected to the other source provides power to the chassis. * &#x60;NotRedundant&#x60; - Power Manager turns on the minimum number of PSUs required to support chassis power requirements. No Redundant PSUs are maintained. * &#x60;N+1&#x60; - Power Manager turns on the minimum number of PSUs required to support chassis power requirements plus one additional PSU for redundancy. * &#x60;N+2&#x60; - Power Manager turns on the minimum number of PSUs required to support chassis power requirements plus two additional PSU for redundancy. This Mode is only supported for UCS X series Chassis. | [optional] [readonly] [default to "Grid"]
**TargetMo** | Pointer to [**MoBaseMoRelationship**](MoBaseMoRelationship.md) |  | [optional] 

## Methods

### NewPowerPolicyInventoryAllOf

`func NewPowerPolicyInventoryAllOf(classId string, objectType string, ) *PowerPolicyInventoryAllOf`

NewPowerPolicyInventoryAllOf instantiates a new PowerPolicyInventoryAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPowerPolicyInventoryAllOfWithDefaults

`func NewPowerPolicyInventoryAllOfWithDefaults() *PowerPolicyInventoryAllOf`

NewPowerPolicyInventoryAllOfWithDefaults instantiates a new PowerPolicyInventoryAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *PowerPolicyInventoryAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *PowerPolicyInventoryAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *PowerPolicyInventoryAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *PowerPolicyInventoryAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PowerPolicyInventoryAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PowerPolicyInventoryAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAllocatedBudget

`func (o *PowerPolicyInventoryAllOf) GetAllocatedBudget() int64`

GetAllocatedBudget returns the AllocatedBudget field if non-nil, zero value otherwise.

### GetAllocatedBudgetOk

`func (o *PowerPolicyInventoryAllOf) GetAllocatedBudgetOk() (*int64, bool)`

GetAllocatedBudgetOk returns a tuple with the AllocatedBudget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedBudget

`func (o *PowerPolicyInventoryAllOf) SetAllocatedBudget(v int64)`

SetAllocatedBudget sets AllocatedBudget field to given value.

### HasAllocatedBudget

`func (o *PowerPolicyInventoryAllOf) HasAllocatedBudget() bool`

HasAllocatedBudget returns a boolean if a field has been set.

### GetDynamicRebalancing

`func (o *PowerPolicyInventoryAllOf) GetDynamicRebalancing() string`

GetDynamicRebalancing returns the DynamicRebalancing field if non-nil, zero value otherwise.

### GetDynamicRebalancingOk

`func (o *PowerPolicyInventoryAllOf) GetDynamicRebalancingOk() (*string, bool)`

GetDynamicRebalancingOk returns a tuple with the DynamicRebalancing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicRebalancing

`func (o *PowerPolicyInventoryAllOf) SetDynamicRebalancing(v string)`

SetDynamicRebalancing sets DynamicRebalancing field to given value.

### HasDynamicRebalancing

`func (o *PowerPolicyInventoryAllOf) HasDynamicRebalancing() bool`

HasDynamicRebalancing returns a boolean if a field has been set.

### GetExtendedPowerCapacity

`func (o *PowerPolicyInventoryAllOf) GetExtendedPowerCapacity() string`

GetExtendedPowerCapacity returns the ExtendedPowerCapacity field if non-nil, zero value otherwise.

### GetExtendedPowerCapacityOk

`func (o *PowerPolicyInventoryAllOf) GetExtendedPowerCapacityOk() (*string, bool)`

GetExtendedPowerCapacityOk returns a tuple with the ExtendedPowerCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedPowerCapacity

`func (o *PowerPolicyInventoryAllOf) SetExtendedPowerCapacity(v string)`

SetExtendedPowerCapacity sets ExtendedPowerCapacity field to given value.

### HasExtendedPowerCapacity

`func (o *PowerPolicyInventoryAllOf) HasExtendedPowerCapacity() bool`

HasExtendedPowerCapacity returns a boolean if a field has been set.

### GetPowerPriority

`func (o *PowerPolicyInventoryAllOf) GetPowerPriority() string`

GetPowerPriority returns the PowerPriority field if non-nil, zero value otherwise.

### GetPowerPriorityOk

`func (o *PowerPolicyInventoryAllOf) GetPowerPriorityOk() (*string, bool)`

GetPowerPriorityOk returns a tuple with the PowerPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerPriority

`func (o *PowerPolicyInventoryAllOf) SetPowerPriority(v string)`

SetPowerPriority sets PowerPriority field to given value.

### HasPowerPriority

`func (o *PowerPolicyInventoryAllOf) HasPowerPriority() bool`

HasPowerPriority returns a boolean if a field has been set.

### GetPowerProfiling

`func (o *PowerPolicyInventoryAllOf) GetPowerProfiling() string`

GetPowerProfiling returns the PowerProfiling field if non-nil, zero value otherwise.

### GetPowerProfilingOk

`func (o *PowerPolicyInventoryAllOf) GetPowerProfilingOk() (*string, bool)`

GetPowerProfilingOk returns a tuple with the PowerProfiling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerProfiling

`func (o *PowerPolicyInventoryAllOf) SetPowerProfiling(v string)`

SetPowerProfiling sets PowerProfiling field to given value.

### HasPowerProfiling

`func (o *PowerPolicyInventoryAllOf) HasPowerProfiling() bool`

HasPowerProfiling returns a boolean if a field has been set.

### GetPowerRestoreState

`func (o *PowerPolicyInventoryAllOf) GetPowerRestoreState() string`

GetPowerRestoreState returns the PowerRestoreState field if non-nil, zero value otherwise.

### GetPowerRestoreStateOk

`func (o *PowerPolicyInventoryAllOf) GetPowerRestoreStateOk() (*string, bool)`

GetPowerRestoreStateOk returns a tuple with the PowerRestoreState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerRestoreState

`func (o *PowerPolicyInventoryAllOf) SetPowerRestoreState(v string)`

SetPowerRestoreState sets PowerRestoreState field to given value.

### HasPowerRestoreState

`func (o *PowerPolicyInventoryAllOf) HasPowerRestoreState() bool`

HasPowerRestoreState returns a boolean if a field has been set.

### GetPowerSaveMode

`func (o *PowerPolicyInventoryAllOf) GetPowerSaveMode() string`

GetPowerSaveMode returns the PowerSaveMode field if non-nil, zero value otherwise.

### GetPowerSaveModeOk

`func (o *PowerPolicyInventoryAllOf) GetPowerSaveModeOk() (*string, bool)`

GetPowerSaveModeOk returns a tuple with the PowerSaveMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerSaveMode

`func (o *PowerPolicyInventoryAllOf) SetPowerSaveMode(v string)`

SetPowerSaveMode sets PowerSaveMode field to given value.

### HasPowerSaveMode

`func (o *PowerPolicyInventoryAllOf) HasPowerSaveMode() bool`

HasPowerSaveMode returns a boolean if a field has been set.

### GetRedundancyMode

`func (o *PowerPolicyInventoryAllOf) GetRedundancyMode() string`

GetRedundancyMode returns the RedundancyMode field if non-nil, zero value otherwise.

### GetRedundancyModeOk

`func (o *PowerPolicyInventoryAllOf) GetRedundancyModeOk() (*string, bool)`

GetRedundancyModeOk returns a tuple with the RedundancyMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedundancyMode

`func (o *PowerPolicyInventoryAllOf) SetRedundancyMode(v string)`

SetRedundancyMode sets RedundancyMode field to given value.

### HasRedundancyMode

`func (o *PowerPolicyInventoryAllOf) HasRedundancyMode() bool`

HasRedundancyMode returns a boolean if a field has been set.

### GetTargetMo

`func (o *PowerPolicyInventoryAllOf) GetTargetMo() MoBaseMoRelationship`

GetTargetMo returns the TargetMo field if non-nil, zero value otherwise.

### GetTargetMoOk

`func (o *PowerPolicyInventoryAllOf) GetTargetMoOk() (*MoBaseMoRelationship, bool)`

GetTargetMoOk returns a tuple with the TargetMo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetMo

`func (o *PowerPolicyInventoryAllOf) SetTargetMo(v MoBaseMoRelationship)`

SetTargetMo sets TargetMo field to given value.

### HasTargetMo

`func (o *PowerPolicyInventoryAllOf) HasTargetMo() bool`

HasTargetMo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



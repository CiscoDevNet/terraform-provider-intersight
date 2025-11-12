# ComputeServerPowerParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "compute.ServerPowerParameters"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "compute.ServerPowerParameters"]
**PowerAllocation** | Pointer to **int64** | This field identifies the maximum power that has been allocated to the blade by CMC in Watts. Power budget for the chassis is configured by the power policy. That budget is then divided among the blades in the chassis by CMC. | [optional] [readonly] [default to 0]
**PowerPriority** | Pointer to **string** | Power Priority level of the Server. This priority is used to determine the initial power allocation for servers. This field is only supported for Cisco UCS B series and X series servers. * &#x60;Unknown&#x60; - Power allocation priority of server is either unknown or not supported on CMC firmware version. * &#x60;Low&#x60; - Power allocation priority of server is low. * &#x60;Medium&#x60; - Power allocation priority of server is medium. * &#x60;High&#x60; - Power allocation priority of server is high. | [optional] [readonly] [default to "Unknown"]
**PowerProfiling** | Pointer to **string** | Status of Power Profiling setting of the Server. If Enabled, this field allows the power manager to run power profiling utility to determine the power needs of the server. This field is only supported for Cisco UCS X series servers. * &#x60;Unknown&#x60; - Power Profiling state is either unknown for the server or not supported on BMC firmware version. * &#x60;Enabled&#x60; - Power Profiling is enabled for the server. * &#x60;Disabled&#x60; - Power Profiling is either disabled for the server or not supported on BMC firmware version. | [optional] [readonly] [default to "Unknown"]
**PowerRestore** | Pointer to **string** | Value of the power restore policy for the server. In the absence of Intersight connectivity, the server will use this state to recover the host power after a power loss event. * &#x60;Unknown&#x60; - Power restore state for the server is either Unknown or not supported on BMC firmware version. * &#x60;Always On&#x60; - Power restore state for server is set to Always On. * &#x60;Always Off&#x60; - Power restore state for server is set to Always Off. * &#x60;Last State&#x60; - Power restore state for server is set to Last State. | [optional] [readonly] [default to "Unknown"]
**ProcessorPackagePowerLimit** | Pointer to **string** | Processor Package Power Limit (PPL) of a server. PPL refers to the amount of power that a CPU can draw from the power supply. The Processor Package Power Limit (PPL) feature is currently available exclusively on Cisco UCS C225/C245 M8 servers. * &#x60;Unknown&#x60; - Processor package power limit is either unknown for the server or not supported on BMC firmware version. * &#x60;Default&#x60; - Processor package power limit is default for the server. * &#x60;Maximum&#x60; - Processor package power limit is maximum for the server. * &#x60;Minimum&#x60; - Processor package power limit is minimum for the server. | [optional] [readonly] [default to "Unknown"]
**ComputeBlade** | Pointer to [**NullableComputeBladeRelationship**](ComputeBladeRelationship.md) |  | [optional] 
**ComputeRackUnit** | Pointer to [**NullableComputeRackUnitRelationship**](ComputeRackUnitRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewComputeServerPowerParameters

`func NewComputeServerPowerParameters(classId string, objectType string, ) *ComputeServerPowerParameters`

NewComputeServerPowerParameters instantiates a new ComputeServerPowerParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputeServerPowerParametersWithDefaults

`func NewComputeServerPowerParametersWithDefaults() *ComputeServerPowerParameters`

NewComputeServerPowerParametersWithDefaults instantiates a new ComputeServerPowerParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ComputeServerPowerParameters) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ComputeServerPowerParameters) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ComputeServerPowerParameters) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ComputeServerPowerParameters) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ComputeServerPowerParameters) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ComputeServerPowerParameters) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetPowerAllocation

`func (o *ComputeServerPowerParameters) GetPowerAllocation() int64`

GetPowerAllocation returns the PowerAllocation field if non-nil, zero value otherwise.

### GetPowerAllocationOk

`func (o *ComputeServerPowerParameters) GetPowerAllocationOk() (*int64, bool)`

GetPowerAllocationOk returns a tuple with the PowerAllocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerAllocation

`func (o *ComputeServerPowerParameters) SetPowerAllocation(v int64)`

SetPowerAllocation sets PowerAllocation field to given value.

### HasPowerAllocation

`func (o *ComputeServerPowerParameters) HasPowerAllocation() bool`

HasPowerAllocation returns a boolean if a field has been set.

### GetPowerPriority

`func (o *ComputeServerPowerParameters) GetPowerPriority() string`

GetPowerPriority returns the PowerPriority field if non-nil, zero value otherwise.

### GetPowerPriorityOk

`func (o *ComputeServerPowerParameters) GetPowerPriorityOk() (*string, bool)`

GetPowerPriorityOk returns a tuple with the PowerPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerPriority

`func (o *ComputeServerPowerParameters) SetPowerPriority(v string)`

SetPowerPriority sets PowerPriority field to given value.

### HasPowerPriority

`func (o *ComputeServerPowerParameters) HasPowerPriority() bool`

HasPowerPriority returns a boolean if a field has been set.

### GetPowerProfiling

`func (o *ComputeServerPowerParameters) GetPowerProfiling() string`

GetPowerProfiling returns the PowerProfiling field if non-nil, zero value otherwise.

### GetPowerProfilingOk

`func (o *ComputeServerPowerParameters) GetPowerProfilingOk() (*string, bool)`

GetPowerProfilingOk returns a tuple with the PowerProfiling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerProfiling

`func (o *ComputeServerPowerParameters) SetPowerProfiling(v string)`

SetPowerProfiling sets PowerProfiling field to given value.

### HasPowerProfiling

`func (o *ComputeServerPowerParameters) HasPowerProfiling() bool`

HasPowerProfiling returns a boolean if a field has been set.

### GetPowerRestore

`func (o *ComputeServerPowerParameters) GetPowerRestore() string`

GetPowerRestore returns the PowerRestore field if non-nil, zero value otherwise.

### GetPowerRestoreOk

`func (o *ComputeServerPowerParameters) GetPowerRestoreOk() (*string, bool)`

GetPowerRestoreOk returns a tuple with the PowerRestore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerRestore

`func (o *ComputeServerPowerParameters) SetPowerRestore(v string)`

SetPowerRestore sets PowerRestore field to given value.

### HasPowerRestore

`func (o *ComputeServerPowerParameters) HasPowerRestore() bool`

HasPowerRestore returns a boolean if a field has been set.

### GetProcessorPackagePowerLimit

`func (o *ComputeServerPowerParameters) GetProcessorPackagePowerLimit() string`

GetProcessorPackagePowerLimit returns the ProcessorPackagePowerLimit field if non-nil, zero value otherwise.

### GetProcessorPackagePowerLimitOk

`func (o *ComputeServerPowerParameters) GetProcessorPackagePowerLimitOk() (*string, bool)`

GetProcessorPackagePowerLimitOk returns a tuple with the ProcessorPackagePowerLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorPackagePowerLimit

`func (o *ComputeServerPowerParameters) SetProcessorPackagePowerLimit(v string)`

SetProcessorPackagePowerLimit sets ProcessorPackagePowerLimit field to given value.

### HasProcessorPackagePowerLimit

`func (o *ComputeServerPowerParameters) HasProcessorPackagePowerLimit() bool`

HasProcessorPackagePowerLimit returns a boolean if a field has been set.

### GetComputeBlade

`func (o *ComputeServerPowerParameters) GetComputeBlade() ComputeBladeRelationship`

GetComputeBlade returns the ComputeBlade field if non-nil, zero value otherwise.

### GetComputeBladeOk

`func (o *ComputeServerPowerParameters) GetComputeBladeOk() (*ComputeBladeRelationship, bool)`

GetComputeBladeOk returns a tuple with the ComputeBlade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeBlade

`func (o *ComputeServerPowerParameters) SetComputeBlade(v ComputeBladeRelationship)`

SetComputeBlade sets ComputeBlade field to given value.

### HasComputeBlade

`func (o *ComputeServerPowerParameters) HasComputeBlade() bool`

HasComputeBlade returns a boolean if a field has been set.

### SetComputeBladeNil

`func (o *ComputeServerPowerParameters) SetComputeBladeNil(b bool)`

 SetComputeBladeNil sets the value for ComputeBlade to be an explicit nil

### UnsetComputeBlade
`func (o *ComputeServerPowerParameters) UnsetComputeBlade()`

UnsetComputeBlade ensures that no value is present for ComputeBlade, not even an explicit nil
### GetComputeRackUnit

`func (o *ComputeServerPowerParameters) GetComputeRackUnit() ComputeRackUnitRelationship`

GetComputeRackUnit returns the ComputeRackUnit field if non-nil, zero value otherwise.

### GetComputeRackUnitOk

`func (o *ComputeServerPowerParameters) GetComputeRackUnitOk() (*ComputeRackUnitRelationship, bool)`

GetComputeRackUnitOk returns a tuple with the ComputeRackUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRackUnit

`func (o *ComputeServerPowerParameters) SetComputeRackUnit(v ComputeRackUnitRelationship)`

SetComputeRackUnit sets ComputeRackUnit field to given value.

### HasComputeRackUnit

`func (o *ComputeServerPowerParameters) HasComputeRackUnit() bool`

HasComputeRackUnit returns a boolean if a field has been set.

### SetComputeRackUnitNil

`func (o *ComputeServerPowerParameters) SetComputeRackUnitNil(b bool)`

 SetComputeRackUnitNil sets the value for ComputeRackUnit to be an explicit nil

### UnsetComputeRackUnit
`func (o *ComputeServerPowerParameters) UnsetComputeRackUnit()`

UnsetComputeRackUnit ensures that no value is present for ComputeRackUnit, not even an explicit nil
### GetRegisteredDevice

`func (o *ComputeServerPowerParameters) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *ComputeServerPowerParameters) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *ComputeServerPowerParameters) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *ComputeServerPowerParameters) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *ComputeServerPowerParameters) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *ComputeServerPowerParameters) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



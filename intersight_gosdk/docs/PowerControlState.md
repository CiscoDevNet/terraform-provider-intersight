# PowerControlState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "power.ControlState"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "power.ControlState"]
**AllocatedPower** | Pointer to **int64** | This field identifies the allocated power on the chassis in Watts. | [optional] [readonly] 
**ExtendedPowerCapacity** | Pointer to **string** | The status of extended power capacity mode of the chassis. If Enabled, this mode allows chassis available power to be increased by borrowing power from redundant power supplies. * &#x60;Enabled&#x60; - Set the value to Enabled. * &#x60;Disabled&#x60; - Set the value to Disabled. * &#x60;Not Applicable&#x60; - Set the value to Not Applicable. | [optional] [readonly] [default to "Enabled"]
**GridMaxPower** | Pointer to **int64** | This field identifies the available power when PSUs are in grid mode in Watts. | [optional] [readonly] 
**MaxRequiredPower** | Pointer to **int64** | This field identifies the maximum power required by the endpoint in Watts. | [optional] [readonly] 
**MinRequiredPower** | Pointer to **int64** | This field identifies the minimum power required by the endpoint in Watts. | [optional] [readonly] 
**N1MaxPower** | Pointer to **int64** | This field identifies the available power when PSUs are in N+1 mode in Watts. | [optional] [readonly] 
**N2MaxPower** | Pointer to **int64** | This field identifies the available power when PSUs are in N+2 mode in Watts. | [optional] [readonly] 
**NonRedundantMaxPower** | Pointer to **int64** | This field identifies the available power when PSUs are in non-redundant mode in Watts. | [optional] [readonly] 
**PowerRebalancing** | Pointer to **string** | The status of power rebalancing mode of the chassis. If enabled, this mode allows the system to dynamically reallocate the power between servers depending on their power usage.  * &#x60;Enabled&#x60; - Set the value to Enabled. * &#x60;Disabled&#x60; - Set the value to Disabled. | [optional] [readonly] [default to "Enabled"]
**PowerSaveMode** | Pointer to **string** | The status of power save mode of the chassis. If Enabled and the requested power budget is less than available power capacity,  the additional PSUs not required to comply with redundancy policy are placed in Power Save mode.  * &#x60;Enabled&#x60; - Set the value to Enabled. * &#x60;Disabled&#x60; - Set the value to Disabled. | [optional] [readonly] [default to "Enabled"]
**EquipmentChassis** | Pointer to [**NullableEquipmentChassisRelationship**](EquipmentChassisRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewPowerControlState

`func NewPowerControlState(classId string, objectType string, ) *PowerControlState`

NewPowerControlState instantiates a new PowerControlState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPowerControlStateWithDefaults

`func NewPowerControlStateWithDefaults() *PowerControlState`

NewPowerControlStateWithDefaults instantiates a new PowerControlState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *PowerControlState) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *PowerControlState) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *PowerControlState) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *PowerControlState) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PowerControlState) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PowerControlState) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAllocatedPower

`func (o *PowerControlState) GetAllocatedPower() int64`

GetAllocatedPower returns the AllocatedPower field if non-nil, zero value otherwise.

### GetAllocatedPowerOk

`func (o *PowerControlState) GetAllocatedPowerOk() (*int64, bool)`

GetAllocatedPowerOk returns a tuple with the AllocatedPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedPower

`func (o *PowerControlState) SetAllocatedPower(v int64)`

SetAllocatedPower sets AllocatedPower field to given value.

### HasAllocatedPower

`func (o *PowerControlState) HasAllocatedPower() bool`

HasAllocatedPower returns a boolean if a field has been set.

### GetExtendedPowerCapacity

`func (o *PowerControlState) GetExtendedPowerCapacity() string`

GetExtendedPowerCapacity returns the ExtendedPowerCapacity field if non-nil, zero value otherwise.

### GetExtendedPowerCapacityOk

`func (o *PowerControlState) GetExtendedPowerCapacityOk() (*string, bool)`

GetExtendedPowerCapacityOk returns a tuple with the ExtendedPowerCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedPowerCapacity

`func (o *PowerControlState) SetExtendedPowerCapacity(v string)`

SetExtendedPowerCapacity sets ExtendedPowerCapacity field to given value.

### HasExtendedPowerCapacity

`func (o *PowerControlState) HasExtendedPowerCapacity() bool`

HasExtendedPowerCapacity returns a boolean if a field has been set.

### GetGridMaxPower

`func (o *PowerControlState) GetGridMaxPower() int64`

GetGridMaxPower returns the GridMaxPower field if non-nil, zero value otherwise.

### GetGridMaxPowerOk

`func (o *PowerControlState) GetGridMaxPowerOk() (*int64, bool)`

GetGridMaxPowerOk returns a tuple with the GridMaxPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGridMaxPower

`func (o *PowerControlState) SetGridMaxPower(v int64)`

SetGridMaxPower sets GridMaxPower field to given value.

### HasGridMaxPower

`func (o *PowerControlState) HasGridMaxPower() bool`

HasGridMaxPower returns a boolean if a field has been set.

### GetMaxRequiredPower

`func (o *PowerControlState) GetMaxRequiredPower() int64`

GetMaxRequiredPower returns the MaxRequiredPower field if non-nil, zero value otherwise.

### GetMaxRequiredPowerOk

`func (o *PowerControlState) GetMaxRequiredPowerOk() (*int64, bool)`

GetMaxRequiredPowerOk returns a tuple with the MaxRequiredPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRequiredPower

`func (o *PowerControlState) SetMaxRequiredPower(v int64)`

SetMaxRequiredPower sets MaxRequiredPower field to given value.

### HasMaxRequiredPower

`func (o *PowerControlState) HasMaxRequiredPower() bool`

HasMaxRequiredPower returns a boolean if a field has been set.

### GetMinRequiredPower

`func (o *PowerControlState) GetMinRequiredPower() int64`

GetMinRequiredPower returns the MinRequiredPower field if non-nil, zero value otherwise.

### GetMinRequiredPowerOk

`func (o *PowerControlState) GetMinRequiredPowerOk() (*int64, bool)`

GetMinRequiredPowerOk returns a tuple with the MinRequiredPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinRequiredPower

`func (o *PowerControlState) SetMinRequiredPower(v int64)`

SetMinRequiredPower sets MinRequiredPower field to given value.

### HasMinRequiredPower

`func (o *PowerControlState) HasMinRequiredPower() bool`

HasMinRequiredPower returns a boolean if a field has been set.

### GetN1MaxPower

`func (o *PowerControlState) GetN1MaxPower() int64`

GetN1MaxPower returns the N1MaxPower field if non-nil, zero value otherwise.

### GetN1MaxPowerOk

`func (o *PowerControlState) GetN1MaxPowerOk() (*int64, bool)`

GetN1MaxPowerOk returns a tuple with the N1MaxPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN1MaxPower

`func (o *PowerControlState) SetN1MaxPower(v int64)`

SetN1MaxPower sets N1MaxPower field to given value.

### HasN1MaxPower

`func (o *PowerControlState) HasN1MaxPower() bool`

HasN1MaxPower returns a boolean if a field has been set.

### GetN2MaxPower

`func (o *PowerControlState) GetN2MaxPower() int64`

GetN2MaxPower returns the N2MaxPower field if non-nil, zero value otherwise.

### GetN2MaxPowerOk

`func (o *PowerControlState) GetN2MaxPowerOk() (*int64, bool)`

GetN2MaxPowerOk returns a tuple with the N2MaxPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN2MaxPower

`func (o *PowerControlState) SetN2MaxPower(v int64)`

SetN2MaxPower sets N2MaxPower field to given value.

### HasN2MaxPower

`func (o *PowerControlState) HasN2MaxPower() bool`

HasN2MaxPower returns a boolean if a field has been set.

### GetNonRedundantMaxPower

`func (o *PowerControlState) GetNonRedundantMaxPower() int64`

GetNonRedundantMaxPower returns the NonRedundantMaxPower field if non-nil, zero value otherwise.

### GetNonRedundantMaxPowerOk

`func (o *PowerControlState) GetNonRedundantMaxPowerOk() (*int64, bool)`

GetNonRedundantMaxPowerOk returns a tuple with the NonRedundantMaxPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonRedundantMaxPower

`func (o *PowerControlState) SetNonRedundantMaxPower(v int64)`

SetNonRedundantMaxPower sets NonRedundantMaxPower field to given value.

### HasNonRedundantMaxPower

`func (o *PowerControlState) HasNonRedundantMaxPower() bool`

HasNonRedundantMaxPower returns a boolean if a field has been set.

### GetPowerRebalancing

`func (o *PowerControlState) GetPowerRebalancing() string`

GetPowerRebalancing returns the PowerRebalancing field if non-nil, zero value otherwise.

### GetPowerRebalancingOk

`func (o *PowerControlState) GetPowerRebalancingOk() (*string, bool)`

GetPowerRebalancingOk returns a tuple with the PowerRebalancing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerRebalancing

`func (o *PowerControlState) SetPowerRebalancing(v string)`

SetPowerRebalancing sets PowerRebalancing field to given value.

### HasPowerRebalancing

`func (o *PowerControlState) HasPowerRebalancing() bool`

HasPowerRebalancing returns a boolean if a field has been set.

### GetPowerSaveMode

`func (o *PowerControlState) GetPowerSaveMode() string`

GetPowerSaveMode returns the PowerSaveMode field if non-nil, zero value otherwise.

### GetPowerSaveModeOk

`func (o *PowerControlState) GetPowerSaveModeOk() (*string, bool)`

GetPowerSaveModeOk returns a tuple with the PowerSaveMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerSaveMode

`func (o *PowerControlState) SetPowerSaveMode(v string)`

SetPowerSaveMode sets PowerSaveMode field to given value.

### HasPowerSaveMode

`func (o *PowerControlState) HasPowerSaveMode() bool`

HasPowerSaveMode returns a boolean if a field has been set.

### GetEquipmentChassis

`func (o *PowerControlState) GetEquipmentChassis() EquipmentChassisRelationship`

GetEquipmentChassis returns the EquipmentChassis field if non-nil, zero value otherwise.

### GetEquipmentChassisOk

`func (o *PowerControlState) GetEquipmentChassisOk() (*EquipmentChassisRelationship, bool)`

GetEquipmentChassisOk returns a tuple with the EquipmentChassis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentChassis

`func (o *PowerControlState) SetEquipmentChassis(v EquipmentChassisRelationship)`

SetEquipmentChassis sets EquipmentChassis field to given value.

### HasEquipmentChassis

`func (o *PowerControlState) HasEquipmentChassis() bool`

HasEquipmentChassis returns a boolean if a field has been set.

### SetEquipmentChassisNil

`func (o *PowerControlState) SetEquipmentChassisNil(b bool)`

 SetEquipmentChassisNil sets the value for EquipmentChassis to be an explicit nil

### UnsetEquipmentChassis
`func (o *PowerControlState) UnsetEquipmentChassis()`

UnsetEquipmentChassis ensures that no value is present for EquipmentChassis, not even an explicit nil
### GetRegisteredDevice

`func (o *PowerControlState) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *PowerControlState) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *PowerControlState) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *PowerControlState) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *PowerControlState) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *PowerControlState) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



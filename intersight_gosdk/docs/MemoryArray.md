# MemoryArray

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArrayId** | Pointer to **int64** | The instance number of the memory array. | [optional] [readonly] 
**CpuId** | Pointer to **int64** | ID of the CPU that access this memory array. | [optional] [readonly] 
**CurrentCapacity** | Pointer to **string** | Current capacity of all the memory units on a server. | [optional] [readonly] 
**ErrorCorrection** | Pointer to **string** | The primary hardware error detection or correction method supported by the memory array. | [optional] [readonly] 
**MaxCapacity** | Pointer to **string** | Maximum capacity of all the memory units on a server. | [optional] [readonly] 
**MaxDevices** | Pointer to **string** | The maximum number of slots or sockets available for memory devices in the memory array. | [optional] [readonly] 
**OperPowerState** | Pointer to **string** | The power state indicator of the memory array. | [optional] [readonly] 
**Presence** | Pointer to **string** | The presence of atleast one memory device in the array. Valid values are &#39;equipped&#39; and &#39;absent&#39;. | [optional] [readonly] 
**ComputeBlade** | Pointer to [**ComputeBladeRelationship**](compute.Blade.Relationship.md) |  | [optional] 
**ComputeBoard** | Pointer to [**ComputeBoardRelationship**](compute.Board.Relationship.md) |  | [optional] 
**ComputeRackUnit** | Pointer to [**ComputeRackUnitRelationship**](compute.RackUnit.Relationship.md) |  | [optional] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](inventory.DeviceInfo.Relationship.md) |  | [optional] 
**PersistentMemoryUnits** | Pointer to [**[]MemoryPersistentMemoryUnitRelationship**](memory.PersistentMemoryUnit.Relationship.md) | An array of relationships to memoryPersistentMemoryUnit resources. | [optional] [readonly] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 
**Units** | Pointer to [**[]MemoryUnitRelationship**](memory.Unit.Relationship.md) | An array of relationships to memoryUnit resources. | [optional] [readonly] 

## Methods

### NewMemoryArray

`func NewMemoryArray() *MemoryArray`

NewMemoryArray instantiates a new MemoryArray object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemoryArrayWithDefaults

`func NewMemoryArrayWithDefaults() *MemoryArray`

NewMemoryArrayWithDefaults instantiates a new MemoryArray object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArrayId

`func (o *MemoryArray) GetArrayId() int64`

GetArrayId returns the ArrayId field if non-nil, zero value otherwise.

### GetArrayIdOk

`func (o *MemoryArray) GetArrayIdOk() (*int64, bool)`

GetArrayIdOk returns a tuple with the ArrayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrayId

`func (o *MemoryArray) SetArrayId(v int64)`

SetArrayId sets ArrayId field to given value.

### HasArrayId

`func (o *MemoryArray) HasArrayId() bool`

HasArrayId returns a boolean if a field has been set.

### GetCpuId

`func (o *MemoryArray) GetCpuId() int64`

GetCpuId returns the CpuId field if non-nil, zero value otherwise.

### GetCpuIdOk

`func (o *MemoryArray) GetCpuIdOk() (*int64, bool)`

GetCpuIdOk returns a tuple with the CpuId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuId

`func (o *MemoryArray) SetCpuId(v int64)`

SetCpuId sets CpuId field to given value.

### HasCpuId

`func (o *MemoryArray) HasCpuId() bool`

HasCpuId returns a boolean if a field has been set.

### GetCurrentCapacity

`func (o *MemoryArray) GetCurrentCapacity() string`

GetCurrentCapacity returns the CurrentCapacity field if non-nil, zero value otherwise.

### GetCurrentCapacityOk

`func (o *MemoryArray) GetCurrentCapacityOk() (*string, bool)`

GetCurrentCapacityOk returns a tuple with the CurrentCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentCapacity

`func (o *MemoryArray) SetCurrentCapacity(v string)`

SetCurrentCapacity sets CurrentCapacity field to given value.

### HasCurrentCapacity

`func (o *MemoryArray) HasCurrentCapacity() bool`

HasCurrentCapacity returns a boolean if a field has been set.

### GetErrorCorrection

`func (o *MemoryArray) GetErrorCorrection() string`

GetErrorCorrection returns the ErrorCorrection field if non-nil, zero value otherwise.

### GetErrorCorrectionOk

`func (o *MemoryArray) GetErrorCorrectionOk() (*string, bool)`

GetErrorCorrectionOk returns a tuple with the ErrorCorrection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCorrection

`func (o *MemoryArray) SetErrorCorrection(v string)`

SetErrorCorrection sets ErrorCorrection field to given value.

### HasErrorCorrection

`func (o *MemoryArray) HasErrorCorrection() bool`

HasErrorCorrection returns a boolean if a field has been set.

### GetMaxCapacity

`func (o *MemoryArray) GetMaxCapacity() string`

GetMaxCapacity returns the MaxCapacity field if non-nil, zero value otherwise.

### GetMaxCapacityOk

`func (o *MemoryArray) GetMaxCapacityOk() (*string, bool)`

GetMaxCapacityOk returns a tuple with the MaxCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCapacity

`func (o *MemoryArray) SetMaxCapacity(v string)`

SetMaxCapacity sets MaxCapacity field to given value.

### HasMaxCapacity

`func (o *MemoryArray) HasMaxCapacity() bool`

HasMaxCapacity returns a boolean if a field has been set.

### GetMaxDevices

`func (o *MemoryArray) GetMaxDevices() string`

GetMaxDevices returns the MaxDevices field if non-nil, zero value otherwise.

### GetMaxDevicesOk

`func (o *MemoryArray) GetMaxDevicesOk() (*string, bool)`

GetMaxDevicesOk returns a tuple with the MaxDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDevices

`func (o *MemoryArray) SetMaxDevices(v string)`

SetMaxDevices sets MaxDevices field to given value.

### HasMaxDevices

`func (o *MemoryArray) HasMaxDevices() bool`

HasMaxDevices returns a boolean if a field has been set.

### GetOperPowerState

`func (o *MemoryArray) GetOperPowerState() string`

GetOperPowerState returns the OperPowerState field if non-nil, zero value otherwise.

### GetOperPowerStateOk

`func (o *MemoryArray) GetOperPowerStateOk() (*string, bool)`

GetOperPowerStateOk returns a tuple with the OperPowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperPowerState

`func (o *MemoryArray) SetOperPowerState(v string)`

SetOperPowerState sets OperPowerState field to given value.

### HasOperPowerState

`func (o *MemoryArray) HasOperPowerState() bool`

HasOperPowerState returns a boolean if a field has been set.

### GetPresence

`func (o *MemoryArray) GetPresence() string`

GetPresence returns the Presence field if non-nil, zero value otherwise.

### GetPresenceOk

`func (o *MemoryArray) GetPresenceOk() (*string, bool)`

GetPresenceOk returns a tuple with the Presence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresence

`func (o *MemoryArray) SetPresence(v string)`

SetPresence sets Presence field to given value.

### HasPresence

`func (o *MemoryArray) HasPresence() bool`

HasPresence returns a boolean if a field has been set.

### GetComputeBlade

`func (o *MemoryArray) GetComputeBlade() ComputeBladeRelationship`

GetComputeBlade returns the ComputeBlade field if non-nil, zero value otherwise.

### GetComputeBladeOk

`func (o *MemoryArray) GetComputeBladeOk() (*ComputeBladeRelationship, bool)`

GetComputeBladeOk returns a tuple with the ComputeBlade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeBlade

`func (o *MemoryArray) SetComputeBlade(v ComputeBladeRelationship)`

SetComputeBlade sets ComputeBlade field to given value.

### HasComputeBlade

`func (o *MemoryArray) HasComputeBlade() bool`

HasComputeBlade returns a boolean if a field has been set.

### GetComputeBoard

`func (o *MemoryArray) GetComputeBoard() ComputeBoardRelationship`

GetComputeBoard returns the ComputeBoard field if non-nil, zero value otherwise.

### GetComputeBoardOk

`func (o *MemoryArray) GetComputeBoardOk() (*ComputeBoardRelationship, bool)`

GetComputeBoardOk returns a tuple with the ComputeBoard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeBoard

`func (o *MemoryArray) SetComputeBoard(v ComputeBoardRelationship)`

SetComputeBoard sets ComputeBoard field to given value.

### HasComputeBoard

`func (o *MemoryArray) HasComputeBoard() bool`

HasComputeBoard returns a boolean if a field has been set.

### GetComputeRackUnit

`func (o *MemoryArray) GetComputeRackUnit() ComputeRackUnitRelationship`

GetComputeRackUnit returns the ComputeRackUnit field if non-nil, zero value otherwise.

### GetComputeRackUnitOk

`func (o *MemoryArray) GetComputeRackUnitOk() (*ComputeRackUnitRelationship, bool)`

GetComputeRackUnitOk returns a tuple with the ComputeRackUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRackUnit

`func (o *MemoryArray) SetComputeRackUnit(v ComputeRackUnitRelationship)`

SetComputeRackUnit sets ComputeRackUnit field to given value.

### HasComputeRackUnit

`func (o *MemoryArray) HasComputeRackUnit() bool`

HasComputeRackUnit returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *MemoryArray) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *MemoryArray) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *MemoryArray) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *MemoryArray) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetPersistentMemoryUnits

`func (o *MemoryArray) GetPersistentMemoryUnits() []MemoryPersistentMemoryUnitRelationship`

GetPersistentMemoryUnits returns the PersistentMemoryUnits field if non-nil, zero value otherwise.

### GetPersistentMemoryUnitsOk

`func (o *MemoryArray) GetPersistentMemoryUnitsOk() (*[]MemoryPersistentMemoryUnitRelationship, bool)`

GetPersistentMemoryUnitsOk returns a tuple with the PersistentMemoryUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentMemoryUnits

`func (o *MemoryArray) SetPersistentMemoryUnits(v []MemoryPersistentMemoryUnitRelationship)`

SetPersistentMemoryUnits sets PersistentMemoryUnits field to given value.

### HasPersistentMemoryUnits

`func (o *MemoryArray) HasPersistentMemoryUnits() bool`

HasPersistentMemoryUnits returns a boolean if a field has been set.

### SetPersistentMemoryUnitsNil

`func (o *MemoryArray) SetPersistentMemoryUnitsNil(b bool)`

 SetPersistentMemoryUnitsNil sets the value for PersistentMemoryUnits to be an explicit nil

### UnsetPersistentMemoryUnits
`func (o *MemoryArray) UnsetPersistentMemoryUnits()`

UnsetPersistentMemoryUnits ensures that no value is present for PersistentMemoryUnits, not even an explicit nil
### GetRegisteredDevice

`func (o *MemoryArray) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *MemoryArray) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *MemoryArray) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *MemoryArray) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### GetUnits

`func (o *MemoryArray) GetUnits() []MemoryUnitRelationship`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *MemoryArray) GetUnitsOk() (*[]MemoryUnitRelationship, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *MemoryArray) SetUnits(v []MemoryUnitRelationship)`

SetUnits sets Units field to given value.

### HasUnits

`func (o *MemoryArray) HasUnits() bool`

HasUnits returns a boolean if a field has been set.

### SetUnitsNil

`func (o *MemoryArray) SetUnitsNil(b bool)`

 SetUnitsNil sets the value for Units to be an explicit nil

### UnsetUnits
`func (o *MemoryArray) UnsetUnits()`

UnsetUnits ensures that no value is present for Units, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



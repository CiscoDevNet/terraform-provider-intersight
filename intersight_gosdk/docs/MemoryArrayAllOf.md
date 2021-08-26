# MemoryArrayAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "memory.Array"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "memory.Array"]
**ArrayId** | Pointer to **int64** | The instance number of the memory array. | [optional] [readonly] 
**CpuId** | Pointer to **int64** | ID of the CPU that access this memory array. | [optional] [readonly] 
**CurrentCapacity** | Pointer to **string** | Current capacity of all the memory units on a server. | [optional] [readonly] 
**ErrorCorrection** | Pointer to **string** | The primary hardware error detection or correction method supported by the memory array. | [optional] [readonly] 
**MaxCapacity** | Pointer to **string** | Maximum capacity of all the memory units on a server. | [optional] [readonly] 
**MaxDevices** | Pointer to **string** | The maximum number of slots or sockets available for memory devices in the memory array. | [optional] [readonly] 
**OperPowerState** | Pointer to **string** | The power state indicator of the memory array. | [optional] [readonly] 
**ComputeBlade** | Pointer to [**ComputeBladeRelationship**](ComputeBladeRelationship.md) |  | [optional] 
**ComputeBoard** | Pointer to [**ComputeBoardRelationship**](ComputeBoardRelationship.md) |  | [optional] 
**ComputeRackUnit** | Pointer to [**ComputeRackUnitRelationship**](ComputeRackUnitRelationship.md) |  | [optional] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](InventoryDeviceInfoRelationship.md) |  | [optional] 
**PersistentMemoryUnits** | Pointer to [**[]MemoryPersistentMemoryUnitRelationship**](MemoryPersistentMemoryUnitRelationship.md) | An array of relationships to memoryPersistentMemoryUnit resources. | [optional] [readonly] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 
**Units** | Pointer to [**[]MemoryUnitRelationship**](MemoryUnitRelationship.md) | An array of relationships to memoryUnit resources. | [optional] [readonly] 

## Methods

### NewMemoryArrayAllOf

`func NewMemoryArrayAllOf(classId string, objectType string, ) *MemoryArrayAllOf`

NewMemoryArrayAllOf instantiates a new MemoryArrayAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemoryArrayAllOfWithDefaults

`func NewMemoryArrayAllOfWithDefaults() *MemoryArrayAllOf`

NewMemoryArrayAllOfWithDefaults instantiates a new MemoryArrayAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *MemoryArrayAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *MemoryArrayAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *MemoryArrayAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *MemoryArrayAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *MemoryArrayAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *MemoryArrayAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetArrayId

`func (o *MemoryArrayAllOf) GetArrayId() int64`

GetArrayId returns the ArrayId field if non-nil, zero value otherwise.

### GetArrayIdOk

`func (o *MemoryArrayAllOf) GetArrayIdOk() (*int64, bool)`

GetArrayIdOk returns a tuple with the ArrayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrayId

`func (o *MemoryArrayAllOf) SetArrayId(v int64)`

SetArrayId sets ArrayId field to given value.

### HasArrayId

`func (o *MemoryArrayAllOf) HasArrayId() bool`

HasArrayId returns a boolean if a field has been set.

### GetCpuId

`func (o *MemoryArrayAllOf) GetCpuId() int64`

GetCpuId returns the CpuId field if non-nil, zero value otherwise.

### GetCpuIdOk

`func (o *MemoryArrayAllOf) GetCpuIdOk() (*int64, bool)`

GetCpuIdOk returns a tuple with the CpuId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuId

`func (o *MemoryArrayAllOf) SetCpuId(v int64)`

SetCpuId sets CpuId field to given value.

### HasCpuId

`func (o *MemoryArrayAllOf) HasCpuId() bool`

HasCpuId returns a boolean if a field has been set.

### GetCurrentCapacity

`func (o *MemoryArrayAllOf) GetCurrentCapacity() string`

GetCurrentCapacity returns the CurrentCapacity field if non-nil, zero value otherwise.

### GetCurrentCapacityOk

`func (o *MemoryArrayAllOf) GetCurrentCapacityOk() (*string, bool)`

GetCurrentCapacityOk returns a tuple with the CurrentCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentCapacity

`func (o *MemoryArrayAllOf) SetCurrentCapacity(v string)`

SetCurrentCapacity sets CurrentCapacity field to given value.

### HasCurrentCapacity

`func (o *MemoryArrayAllOf) HasCurrentCapacity() bool`

HasCurrentCapacity returns a boolean if a field has been set.

### GetErrorCorrection

`func (o *MemoryArrayAllOf) GetErrorCorrection() string`

GetErrorCorrection returns the ErrorCorrection field if non-nil, zero value otherwise.

### GetErrorCorrectionOk

`func (o *MemoryArrayAllOf) GetErrorCorrectionOk() (*string, bool)`

GetErrorCorrectionOk returns a tuple with the ErrorCorrection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCorrection

`func (o *MemoryArrayAllOf) SetErrorCorrection(v string)`

SetErrorCorrection sets ErrorCorrection field to given value.

### HasErrorCorrection

`func (o *MemoryArrayAllOf) HasErrorCorrection() bool`

HasErrorCorrection returns a boolean if a field has been set.

### GetMaxCapacity

`func (o *MemoryArrayAllOf) GetMaxCapacity() string`

GetMaxCapacity returns the MaxCapacity field if non-nil, zero value otherwise.

### GetMaxCapacityOk

`func (o *MemoryArrayAllOf) GetMaxCapacityOk() (*string, bool)`

GetMaxCapacityOk returns a tuple with the MaxCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCapacity

`func (o *MemoryArrayAllOf) SetMaxCapacity(v string)`

SetMaxCapacity sets MaxCapacity field to given value.

### HasMaxCapacity

`func (o *MemoryArrayAllOf) HasMaxCapacity() bool`

HasMaxCapacity returns a boolean if a field has been set.

### GetMaxDevices

`func (o *MemoryArrayAllOf) GetMaxDevices() string`

GetMaxDevices returns the MaxDevices field if non-nil, zero value otherwise.

### GetMaxDevicesOk

`func (o *MemoryArrayAllOf) GetMaxDevicesOk() (*string, bool)`

GetMaxDevicesOk returns a tuple with the MaxDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDevices

`func (o *MemoryArrayAllOf) SetMaxDevices(v string)`

SetMaxDevices sets MaxDevices field to given value.

### HasMaxDevices

`func (o *MemoryArrayAllOf) HasMaxDevices() bool`

HasMaxDevices returns a boolean if a field has been set.

### GetOperPowerState

`func (o *MemoryArrayAllOf) GetOperPowerState() string`

GetOperPowerState returns the OperPowerState field if non-nil, zero value otherwise.

### GetOperPowerStateOk

`func (o *MemoryArrayAllOf) GetOperPowerStateOk() (*string, bool)`

GetOperPowerStateOk returns a tuple with the OperPowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperPowerState

`func (o *MemoryArrayAllOf) SetOperPowerState(v string)`

SetOperPowerState sets OperPowerState field to given value.

### HasOperPowerState

`func (o *MemoryArrayAllOf) HasOperPowerState() bool`

HasOperPowerState returns a boolean if a field has been set.

### GetComputeBlade

`func (o *MemoryArrayAllOf) GetComputeBlade() ComputeBladeRelationship`

GetComputeBlade returns the ComputeBlade field if non-nil, zero value otherwise.

### GetComputeBladeOk

`func (o *MemoryArrayAllOf) GetComputeBladeOk() (*ComputeBladeRelationship, bool)`

GetComputeBladeOk returns a tuple with the ComputeBlade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeBlade

`func (o *MemoryArrayAllOf) SetComputeBlade(v ComputeBladeRelationship)`

SetComputeBlade sets ComputeBlade field to given value.

### HasComputeBlade

`func (o *MemoryArrayAllOf) HasComputeBlade() bool`

HasComputeBlade returns a boolean if a field has been set.

### GetComputeBoard

`func (o *MemoryArrayAllOf) GetComputeBoard() ComputeBoardRelationship`

GetComputeBoard returns the ComputeBoard field if non-nil, zero value otherwise.

### GetComputeBoardOk

`func (o *MemoryArrayAllOf) GetComputeBoardOk() (*ComputeBoardRelationship, bool)`

GetComputeBoardOk returns a tuple with the ComputeBoard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeBoard

`func (o *MemoryArrayAllOf) SetComputeBoard(v ComputeBoardRelationship)`

SetComputeBoard sets ComputeBoard field to given value.

### HasComputeBoard

`func (o *MemoryArrayAllOf) HasComputeBoard() bool`

HasComputeBoard returns a boolean if a field has been set.

### GetComputeRackUnit

`func (o *MemoryArrayAllOf) GetComputeRackUnit() ComputeRackUnitRelationship`

GetComputeRackUnit returns the ComputeRackUnit field if non-nil, zero value otherwise.

### GetComputeRackUnitOk

`func (o *MemoryArrayAllOf) GetComputeRackUnitOk() (*ComputeRackUnitRelationship, bool)`

GetComputeRackUnitOk returns a tuple with the ComputeRackUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRackUnit

`func (o *MemoryArrayAllOf) SetComputeRackUnit(v ComputeRackUnitRelationship)`

SetComputeRackUnit sets ComputeRackUnit field to given value.

### HasComputeRackUnit

`func (o *MemoryArrayAllOf) HasComputeRackUnit() bool`

HasComputeRackUnit returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *MemoryArrayAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *MemoryArrayAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *MemoryArrayAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *MemoryArrayAllOf) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetPersistentMemoryUnits

`func (o *MemoryArrayAllOf) GetPersistentMemoryUnits() []MemoryPersistentMemoryUnitRelationship`

GetPersistentMemoryUnits returns the PersistentMemoryUnits field if non-nil, zero value otherwise.

### GetPersistentMemoryUnitsOk

`func (o *MemoryArrayAllOf) GetPersistentMemoryUnitsOk() (*[]MemoryPersistentMemoryUnitRelationship, bool)`

GetPersistentMemoryUnitsOk returns a tuple with the PersistentMemoryUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentMemoryUnits

`func (o *MemoryArrayAllOf) SetPersistentMemoryUnits(v []MemoryPersistentMemoryUnitRelationship)`

SetPersistentMemoryUnits sets PersistentMemoryUnits field to given value.

### HasPersistentMemoryUnits

`func (o *MemoryArrayAllOf) HasPersistentMemoryUnits() bool`

HasPersistentMemoryUnits returns a boolean if a field has been set.

### SetPersistentMemoryUnitsNil

`func (o *MemoryArrayAllOf) SetPersistentMemoryUnitsNil(b bool)`

 SetPersistentMemoryUnitsNil sets the value for PersistentMemoryUnits to be an explicit nil

### UnsetPersistentMemoryUnits
`func (o *MemoryArrayAllOf) UnsetPersistentMemoryUnits()`

UnsetPersistentMemoryUnits ensures that no value is present for PersistentMemoryUnits, not even an explicit nil
### GetRegisteredDevice

`func (o *MemoryArrayAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *MemoryArrayAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *MemoryArrayAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *MemoryArrayAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### GetUnits

`func (o *MemoryArrayAllOf) GetUnits() []MemoryUnitRelationship`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *MemoryArrayAllOf) GetUnitsOk() (*[]MemoryUnitRelationship, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *MemoryArrayAllOf) SetUnits(v []MemoryUnitRelationship)`

SetUnits sets Units field to given value.

### HasUnits

`func (o *MemoryArrayAllOf) HasUnits() bool`

HasUnits returns a boolean if a field has been set.

### SetUnitsNil

`func (o *MemoryArrayAllOf) SetUnitsNil(b bool)`

 SetUnitsNil sets the value for Units to be an explicit nil

### UnsetUnits
`func (o *MemoryArrayAllOf) UnsetUnits()`

UnsetUnits ensures that no value is present for Units, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



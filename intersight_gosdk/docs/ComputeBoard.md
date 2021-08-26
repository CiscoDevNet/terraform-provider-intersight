# ComputeBoard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "compute.Board"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "compute.Board"]
**BoardId** | Pointer to **int64** | Unique identifier of the mother board present in the server. | [optional] [readonly] 
**CpuTypeController** | Pointer to **string** | The type of central processing unit on the mother board. | [optional] [readonly] 
**OperPowerState** | Pointer to **string** | Current power state of the mother board of the server. | [optional] [readonly] 
**OperReason** | Pointer to **[]string** |  | [optional] 
**ComputeBlade** | Pointer to [**ComputeBladeRelationship**](ComputeBladeRelationship.md) |  | [optional] 
**ComputeRackUnit** | Pointer to [**ComputeRackUnitRelationship**](ComputeRackUnitRelationship.md) |  | [optional] 
**EquipmentTpms** | Pointer to [**[]EquipmentTpmRelationship**](EquipmentTpmRelationship.md) | An array of relationships to equipmentTpm resources. | [optional] [readonly] 
**GraphicsCards** | Pointer to [**[]GraphicsCardRelationship**](GraphicsCardRelationship.md) | An array of relationships to graphicsCard resources. | [optional] [readonly] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](InventoryDeviceInfoRelationship.md) |  | [optional] 
**MemoryArrays** | Pointer to [**[]MemoryArrayRelationship**](MemoryArrayRelationship.md) | An array of relationships to memoryArray resources. | [optional] [readonly] 
**PciCoprocessorCards** | Pointer to [**[]PciCoprocessorCardRelationship**](PciCoprocessorCardRelationship.md) | An array of relationships to pciCoprocessorCard resources. | [optional] [readonly] 
**PciSwitch** | Pointer to [**[]PciSwitchRelationship**](PciSwitchRelationship.md) | An array of relationships to pciSwitch resources. | [optional] [readonly] 
**PersistentMemoryConfiguration** | Pointer to [**MemoryPersistentMemoryConfigurationRelationship**](MemoryPersistentMemoryConfigurationRelationship.md) |  | [optional] 
**Processors** | Pointer to [**[]ProcessorUnitRelationship**](ProcessorUnitRelationship.md) | An array of relationships to processorUnit resources. | [optional] [readonly] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 
**SecurityUnits** | Pointer to [**[]SecurityUnitRelationship**](SecurityUnitRelationship.md) | An array of relationships to securityUnit resources. | [optional] [readonly] 
**StorageControllers** | Pointer to [**[]StorageControllerRelationship**](StorageControllerRelationship.md) | An array of relationships to storageController resources. | [optional] [readonly] 
**StorageFlexFlashControllers** | Pointer to [**[]StorageFlexFlashControllerRelationship**](StorageFlexFlashControllerRelationship.md) | An array of relationships to storageFlexFlashController resources. | [optional] [readonly] 
**StorageFlexUtilControllers** | Pointer to [**[]StorageFlexUtilControllerRelationship**](StorageFlexUtilControllerRelationship.md) | An array of relationships to storageFlexUtilController resources. | [optional] [readonly] 

## Methods

### NewComputeBoard

`func NewComputeBoard(classId string, objectType string, ) *ComputeBoard`

NewComputeBoard instantiates a new ComputeBoard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputeBoardWithDefaults

`func NewComputeBoardWithDefaults() *ComputeBoard`

NewComputeBoardWithDefaults instantiates a new ComputeBoard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ComputeBoard) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ComputeBoard) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ComputeBoard) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ComputeBoard) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ComputeBoard) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ComputeBoard) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBoardId

`func (o *ComputeBoard) GetBoardId() int64`

GetBoardId returns the BoardId field if non-nil, zero value otherwise.

### GetBoardIdOk

`func (o *ComputeBoard) GetBoardIdOk() (*int64, bool)`

GetBoardIdOk returns a tuple with the BoardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoardId

`func (o *ComputeBoard) SetBoardId(v int64)`

SetBoardId sets BoardId field to given value.

### HasBoardId

`func (o *ComputeBoard) HasBoardId() bool`

HasBoardId returns a boolean if a field has been set.

### GetCpuTypeController

`func (o *ComputeBoard) GetCpuTypeController() string`

GetCpuTypeController returns the CpuTypeController field if non-nil, zero value otherwise.

### GetCpuTypeControllerOk

`func (o *ComputeBoard) GetCpuTypeControllerOk() (*string, bool)`

GetCpuTypeControllerOk returns a tuple with the CpuTypeController field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuTypeController

`func (o *ComputeBoard) SetCpuTypeController(v string)`

SetCpuTypeController sets CpuTypeController field to given value.

### HasCpuTypeController

`func (o *ComputeBoard) HasCpuTypeController() bool`

HasCpuTypeController returns a boolean if a field has been set.

### GetOperPowerState

`func (o *ComputeBoard) GetOperPowerState() string`

GetOperPowerState returns the OperPowerState field if non-nil, zero value otherwise.

### GetOperPowerStateOk

`func (o *ComputeBoard) GetOperPowerStateOk() (*string, bool)`

GetOperPowerStateOk returns a tuple with the OperPowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperPowerState

`func (o *ComputeBoard) SetOperPowerState(v string)`

SetOperPowerState sets OperPowerState field to given value.

### HasOperPowerState

`func (o *ComputeBoard) HasOperPowerState() bool`

HasOperPowerState returns a boolean if a field has been set.

### GetOperReason

`func (o *ComputeBoard) GetOperReason() []string`

GetOperReason returns the OperReason field if non-nil, zero value otherwise.

### GetOperReasonOk

`func (o *ComputeBoard) GetOperReasonOk() (*[]string, bool)`

GetOperReasonOk returns a tuple with the OperReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperReason

`func (o *ComputeBoard) SetOperReason(v []string)`

SetOperReason sets OperReason field to given value.

### HasOperReason

`func (o *ComputeBoard) HasOperReason() bool`

HasOperReason returns a boolean if a field has been set.

### SetOperReasonNil

`func (o *ComputeBoard) SetOperReasonNil(b bool)`

 SetOperReasonNil sets the value for OperReason to be an explicit nil

### UnsetOperReason
`func (o *ComputeBoard) UnsetOperReason()`

UnsetOperReason ensures that no value is present for OperReason, not even an explicit nil
### GetComputeBlade

`func (o *ComputeBoard) GetComputeBlade() ComputeBladeRelationship`

GetComputeBlade returns the ComputeBlade field if non-nil, zero value otherwise.

### GetComputeBladeOk

`func (o *ComputeBoard) GetComputeBladeOk() (*ComputeBladeRelationship, bool)`

GetComputeBladeOk returns a tuple with the ComputeBlade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeBlade

`func (o *ComputeBoard) SetComputeBlade(v ComputeBladeRelationship)`

SetComputeBlade sets ComputeBlade field to given value.

### HasComputeBlade

`func (o *ComputeBoard) HasComputeBlade() bool`

HasComputeBlade returns a boolean if a field has been set.

### GetComputeRackUnit

`func (o *ComputeBoard) GetComputeRackUnit() ComputeRackUnitRelationship`

GetComputeRackUnit returns the ComputeRackUnit field if non-nil, zero value otherwise.

### GetComputeRackUnitOk

`func (o *ComputeBoard) GetComputeRackUnitOk() (*ComputeRackUnitRelationship, bool)`

GetComputeRackUnitOk returns a tuple with the ComputeRackUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRackUnit

`func (o *ComputeBoard) SetComputeRackUnit(v ComputeRackUnitRelationship)`

SetComputeRackUnit sets ComputeRackUnit field to given value.

### HasComputeRackUnit

`func (o *ComputeBoard) HasComputeRackUnit() bool`

HasComputeRackUnit returns a boolean if a field has been set.

### GetEquipmentTpms

`func (o *ComputeBoard) GetEquipmentTpms() []EquipmentTpmRelationship`

GetEquipmentTpms returns the EquipmentTpms field if non-nil, zero value otherwise.

### GetEquipmentTpmsOk

`func (o *ComputeBoard) GetEquipmentTpmsOk() (*[]EquipmentTpmRelationship, bool)`

GetEquipmentTpmsOk returns a tuple with the EquipmentTpms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentTpms

`func (o *ComputeBoard) SetEquipmentTpms(v []EquipmentTpmRelationship)`

SetEquipmentTpms sets EquipmentTpms field to given value.

### HasEquipmentTpms

`func (o *ComputeBoard) HasEquipmentTpms() bool`

HasEquipmentTpms returns a boolean if a field has been set.

### SetEquipmentTpmsNil

`func (o *ComputeBoard) SetEquipmentTpmsNil(b bool)`

 SetEquipmentTpmsNil sets the value for EquipmentTpms to be an explicit nil

### UnsetEquipmentTpms
`func (o *ComputeBoard) UnsetEquipmentTpms()`

UnsetEquipmentTpms ensures that no value is present for EquipmentTpms, not even an explicit nil
### GetGraphicsCards

`func (o *ComputeBoard) GetGraphicsCards() []GraphicsCardRelationship`

GetGraphicsCards returns the GraphicsCards field if non-nil, zero value otherwise.

### GetGraphicsCardsOk

`func (o *ComputeBoard) GetGraphicsCardsOk() (*[]GraphicsCardRelationship, bool)`

GetGraphicsCardsOk returns a tuple with the GraphicsCards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraphicsCards

`func (o *ComputeBoard) SetGraphicsCards(v []GraphicsCardRelationship)`

SetGraphicsCards sets GraphicsCards field to given value.

### HasGraphicsCards

`func (o *ComputeBoard) HasGraphicsCards() bool`

HasGraphicsCards returns a boolean if a field has been set.

### SetGraphicsCardsNil

`func (o *ComputeBoard) SetGraphicsCardsNil(b bool)`

 SetGraphicsCardsNil sets the value for GraphicsCards to be an explicit nil

### UnsetGraphicsCards
`func (o *ComputeBoard) UnsetGraphicsCards()`

UnsetGraphicsCards ensures that no value is present for GraphicsCards, not even an explicit nil
### GetInventoryDeviceInfo

`func (o *ComputeBoard) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *ComputeBoard) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *ComputeBoard) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *ComputeBoard) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetMemoryArrays

`func (o *ComputeBoard) GetMemoryArrays() []MemoryArrayRelationship`

GetMemoryArrays returns the MemoryArrays field if non-nil, zero value otherwise.

### GetMemoryArraysOk

`func (o *ComputeBoard) GetMemoryArraysOk() (*[]MemoryArrayRelationship, bool)`

GetMemoryArraysOk returns a tuple with the MemoryArrays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryArrays

`func (o *ComputeBoard) SetMemoryArrays(v []MemoryArrayRelationship)`

SetMemoryArrays sets MemoryArrays field to given value.

### HasMemoryArrays

`func (o *ComputeBoard) HasMemoryArrays() bool`

HasMemoryArrays returns a boolean if a field has been set.

### SetMemoryArraysNil

`func (o *ComputeBoard) SetMemoryArraysNil(b bool)`

 SetMemoryArraysNil sets the value for MemoryArrays to be an explicit nil

### UnsetMemoryArrays
`func (o *ComputeBoard) UnsetMemoryArrays()`

UnsetMemoryArrays ensures that no value is present for MemoryArrays, not even an explicit nil
### GetPciCoprocessorCards

`func (o *ComputeBoard) GetPciCoprocessorCards() []PciCoprocessorCardRelationship`

GetPciCoprocessorCards returns the PciCoprocessorCards field if non-nil, zero value otherwise.

### GetPciCoprocessorCardsOk

`func (o *ComputeBoard) GetPciCoprocessorCardsOk() (*[]PciCoprocessorCardRelationship, bool)`

GetPciCoprocessorCardsOk returns a tuple with the PciCoprocessorCards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciCoprocessorCards

`func (o *ComputeBoard) SetPciCoprocessorCards(v []PciCoprocessorCardRelationship)`

SetPciCoprocessorCards sets PciCoprocessorCards field to given value.

### HasPciCoprocessorCards

`func (o *ComputeBoard) HasPciCoprocessorCards() bool`

HasPciCoprocessorCards returns a boolean if a field has been set.

### SetPciCoprocessorCardsNil

`func (o *ComputeBoard) SetPciCoprocessorCardsNil(b bool)`

 SetPciCoprocessorCardsNil sets the value for PciCoprocessorCards to be an explicit nil

### UnsetPciCoprocessorCards
`func (o *ComputeBoard) UnsetPciCoprocessorCards()`

UnsetPciCoprocessorCards ensures that no value is present for PciCoprocessorCards, not even an explicit nil
### GetPciSwitch

`func (o *ComputeBoard) GetPciSwitch() []PciSwitchRelationship`

GetPciSwitch returns the PciSwitch field if non-nil, zero value otherwise.

### GetPciSwitchOk

`func (o *ComputeBoard) GetPciSwitchOk() (*[]PciSwitchRelationship, bool)`

GetPciSwitchOk returns a tuple with the PciSwitch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciSwitch

`func (o *ComputeBoard) SetPciSwitch(v []PciSwitchRelationship)`

SetPciSwitch sets PciSwitch field to given value.

### HasPciSwitch

`func (o *ComputeBoard) HasPciSwitch() bool`

HasPciSwitch returns a boolean if a field has been set.

### SetPciSwitchNil

`func (o *ComputeBoard) SetPciSwitchNil(b bool)`

 SetPciSwitchNil sets the value for PciSwitch to be an explicit nil

### UnsetPciSwitch
`func (o *ComputeBoard) UnsetPciSwitch()`

UnsetPciSwitch ensures that no value is present for PciSwitch, not even an explicit nil
### GetPersistentMemoryConfiguration

`func (o *ComputeBoard) GetPersistentMemoryConfiguration() MemoryPersistentMemoryConfigurationRelationship`

GetPersistentMemoryConfiguration returns the PersistentMemoryConfiguration field if non-nil, zero value otherwise.

### GetPersistentMemoryConfigurationOk

`func (o *ComputeBoard) GetPersistentMemoryConfigurationOk() (*MemoryPersistentMemoryConfigurationRelationship, bool)`

GetPersistentMemoryConfigurationOk returns a tuple with the PersistentMemoryConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentMemoryConfiguration

`func (o *ComputeBoard) SetPersistentMemoryConfiguration(v MemoryPersistentMemoryConfigurationRelationship)`

SetPersistentMemoryConfiguration sets PersistentMemoryConfiguration field to given value.

### HasPersistentMemoryConfiguration

`func (o *ComputeBoard) HasPersistentMemoryConfiguration() bool`

HasPersistentMemoryConfiguration returns a boolean if a field has been set.

### GetProcessors

`func (o *ComputeBoard) GetProcessors() []ProcessorUnitRelationship`

GetProcessors returns the Processors field if non-nil, zero value otherwise.

### GetProcessorsOk

`func (o *ComputeBoard) GetProcessorsOk() (*[]ProcessorUnitRelationship, bool)`

GetProcessorsOk returns a tuple with the Processors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessors

`func (o *ComputeBoard) SetProcessors(v []ProcessorUnitRelationship)`

SetProcessors sets Processors field to given value.

### HasProcessors

`func (o *ComputeBoard) HasProcessors() bool`

HasProcessors returns a boolean if a field has been set.

### SetProcessorsNil

`func (o *ComputeBoard) SetProcessorsNil(b bool)`

 SetProcessorsNil sets the value for Processors to be an explicit nil

### UnsetProcessors
`func (o *ComputeBoard) UnsetProcessors()`

UnsetProcessors ensures that no value is present for Processors, not even an explicit nil
### GetRegisteredDevice

`func (o *ComputeBoard) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *ComputeBoard) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *ComputeBoard) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *ComputeBoard) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### GetSecurityUnits

`func (o *ComputeBoard) GetSecurityUnits() []SecurityUnitRelationship`

GetSecurityUnits returns the SecurityUnits field if non-nil, zero value otherwise.

### GetSecurityUnitsOk

`func (o *ComputeBoard) GetSecurityUnitsOk() (*[]SecurityUnitRelationship, bool)`

GetSecurityUnitsOk returns a tuple with the SecurityUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityUnits

`func (o *ComputeBoard) SetSecurityUnits(v []SecurityUnitRelationship)`

SetSecurityUnits sets SecurityUnits field to given value.

### HasSecurityUnits

`func (o *ComputeBoard) HasSecurityUnits() bool`

HasSecurityUnits returns a boolean if a field has been set.

### SetSecurityUnitsNil

`func (o *ComputeBoard) SetSecurityUnitsNil(b bool)`

 SetSecurityUnitsNil sets the value for SecurityUnits to be an explicit nil

### UnsetSecurityUnits
`func (o *ComputeBoard) UnsetSecurityUnits()`

UnsetSecurityUnits ensures that no value is present for SecurityUnits, not even an explicit nil
### GetStorageControllers

`func (o *ComputeBoard) GetStorageControllers() []StorageControllerRelationship`

GetStorageControllers returns the StorageControllers field if non-nil, zero value otherwise.

### GetStorageControllersOk

`func (o *ComputeBoard) GetStorageControllersOk() (*[]StorageControllerRelationship, bool)`

GetStorageControllersOk returns a tuple with the StorageControllers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageControllers

`func (o *ComputeBoard) SetStorageControllers(v []StorageControllerRelationship)`

SetStorageControllers sets StorageControllers field to given value.

### HasStorageControllers

`func (o *ComputeBoard) HasStorageControllers() bool`

HasStorageControllers returns a boolean if a field has been set.

### SetStorageControllersNil

`func (o *ComputeBoard) SetStorageControllersNil(b bool)`

 SetStorageControllersNil sets the value for StorageControllers to be an explicit nil

### UnsetStorageControllers
`func (o *ComputeBoard) UnsetStorageControllers()`

UnsetStorageControllers ensures that no value is present for StorageControllers, not even an explicit nil
### GetStorageFlexFlashControllers

`func (o *ComputeBoard) GetStorageFlexFlashControllers() []StorageFlexFlashControllerRelationship`

GetStorageFlexFlashControllers returns the StorageFlexFlashControllers field if non-nil, zero value otherwise.

### GetStorageFlexFlashControllersOk

`func (o *ComputeBoard) GetStorageFlexFlashControllersOk() (*[]StorageFlexFlashControllerRelationship, bool)`

GetStorageFlexFlashControllersOk returns a tuple with the StorageFlexFlashControllers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageFlexFlashControllers

`func (o *ComputeBoard) SetStorageFlexFlashControllers(v []StorageFlexFlashControllerRelationship)`

SetStorageFlexFlashControllers sets StorageFlexFlashControllers field to given value.

### HasStorageFlexFlashControllers

`func (o *ComputeBoard) HasStorageFlexFlashControllers() bool`

HasStorageFlexFlashControllers returns a boolean if a field has been set.

### SetStorageFlexFlashControllersNil

`func (o *ComputeBoard) SetStorageFlexFlashControllersNil(b bool)`

 SetStorageFlexFlashControllersNil sets the value for StorageFlexFlashControllers to be an explicit nil

### UnsetStorageFlexFlashControllers
`func (o *ComputeBoard) UnsetStorageFlexFlashControllers()`

UnsetStorageFlexFlashControllers ensures that no value is present for StorageFlexFlashControllers, not even an explicit nil
### GetStorageFlexUtilControllers

`func (o *ComputeBoard) GetStorageFlexUtilControllers() []StorageFlexUtilControllerRelationship`

GetStorageFlexUtilControllers returns the StorageFlexUtilControllers field if non-nil, zero value otherwise.

### GetStorageFlexUtilControllersOk

`func (o *ComputeBoard) GetStorageFlexUtilControllersOk() (*[]StorageFlexUtilControllerRelationship, bool)`

GetStorageFlexUtilControllersOk returns a tuple with the StorageFlexUtilControllers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageFlexUtilControllers

`func (o *ComputeBoard) SetStorageFlexUtilControllers(v []StorageFlexUtilControllerRelationship)`

SetStorageFlexUtilControllers sets StorageFlexUtilControllers field to given value.

### HasStorageFlexUtilControllers

`func (o *ComputeBoard) HasStorageFlexUtilControllers() bool`

HasStorageFlexUtilControllers returns a boolean if a field has been set.

### SetStorageFlexUtilControllersNil

`func (o *ComputeBoard) SetStorageFlexUtilControllersNil(b bool)`

 SetStorageFlexUtilControllersNil sets the value for StorageFlexUtilControllers to be an explicit nil

### UnsetStorageFlexUtilControllers
`func (o *ComputeBoard) UnsetStorageFlexUtilControllers()`

UnsetStorageFlexUtilControllers ensures that no value is present for StorageFlexUtilControllers, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



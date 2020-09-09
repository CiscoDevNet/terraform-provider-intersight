# ComputeBladeAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChassisId** | Pointer to **string** | The id of the chassis that the blade is located in. | [optional] [readonly] 
**ScaledMode** | Pointer to **string** | The mode of the server that determines it is scaled. | [optional] [readonly] 
**SlotId** | Pointer to **int64** | The slot number in the chassis that the blade is located in. | [optional] [readonly] 
**Adapters** | Pointer to [**[]AdapterUnitRelationship**](adapter.Unit.Relationship.md) | An array of relationships to adapterUnit resources. | [optional] [readonly] 
**BiosUnits** | Pointer to [**[]BiosUnitRelationship**](bios.Unit.Relationship.md) | An array of relationships to biosUnit resources. | [optional] [readonly] 
**Bmc** | Pointer to [**ManagementControllerRelationship**](management.Controller.Relationship.md) |  | [optional] 
**Board** | Pointer to [**ComputeBoardRelationship**](compute.Board.Relationship.md) |  | [optional] 
**EquipmentChassis** | Pointer to [**EquipmentChassisRelationship**](equipment.Chassis.Relationship.md) |  | [optional] 
**EquipmentIoExpanders** | Pointer to [**[]EquipmentIoExpanderRelationship**](equipment.IoExpander.Relationship.md) | An array of relationships to equipmentIoExpander resources. | [optional] [readonly] 
**GenericInventoryHolders** | Pointer to [**[]InventoryGenericInventoryHolderRelationship**](inventory.GenericInventoryHolder.Relationship.md) | An array of relationships to inventoryGenericInventoryHolder resources. | [optional] [readonly] 
**GraphicsCards** | Pointer to [**[]GraphicsCardRelationship**](graphics.Card.Relationship.md) | An array of relationships to graphicsCard resources. | [optional] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](inventory.DeviceInfo.Relationship.md) |  | [optional] 
**LocatorLed** | Pointer to [**EquipmentLocatorLedRelationship**](equipment.LocatorLed.Relationship.md) |  | [optional] 
**MemoryArrays** | Pointer to [**[]MemoryArrayRelationship**](memory.Array.Relationship.md) | An array of relationships to memoryArray resources. | [optional] 
**PciDevices** | Pointer to [**[]PciDeviceRelationship**](pci.Device.Relationship.md) | An array of relationships to pciDevice resources. | [optional] [readonly] 
**Processors** | Pointer to [**[]ProcessorUnitRelationship**](processor.Unit.Relationship.md) | An array of relationships to processorUnit resources. | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 
**StorageControllers** | Pointer to [**[]StorageControllerRelationship**](storage.Controller.Relationship.md) | An array of relationships to storageController resources. | [optional] 
**StorageEnclosures** | Pointer to [**[]StorageEnclosureRelationship**](storage.Enclosure.Relationship.md) | An array of relationships to storageEnclosure resources. | [optional] [readonly] 
**TopSystem** | Pointer to [**TopSystemRelationship**](top.System.Relationship.md) |  | [optional] 

## Methods

### NewComputeBladeAllOf

`func NewComputeBladeAllOf() *ComputeBladeAllOf`

NewComputeBladeAllOf instantiates a new ComputeBladeAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputeBladeAllOfWithDefaults

`func NewComputeBladeAllOfWithDefaults() *ComputeBladeAllOf`

NewComputeBladeAllOfWithDefaults instantiates a new ComputeBladeAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChassisId

`func (o *ComputeBladeAllOf) GetChassisId() string`

GetChassisId returns the ChassisId field if non-nil, zero value otherwise.

### GetChassisIdOk

`func (o *ComputeBladeAllOf) GetChassisIdOk() (*string, bool)`

GetChassisIdOk returns a tuple with the ChassisId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassisId

`func (o *ComputeBladeAllOf) SetChassisId(v string)`

SetChassisId sets ChassisId field to given value.

### HasChassisId

`func (o *ComputeBladeAllOf) HasChassisId() bool`

HasChassisId returns a boolean if a field has been set.

### GetScaledMode

`func (o *ComputeBladeAllOf) GetScaledMode() string`

GetScaledMode returns the ScaledMode field if non-nil, zero value otherwise.

### GetScaledModeOk

`func (o *ComputeBladeAllOf) GetScaledModeOk() (*string, bool)`

GetScaledModeOk returns a tuple with the ScaledMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaledMode

`func (o *ComputeBladeAllOf) SetScaledMode(v string)`

SetScaledMode sets ScaledMode field to given value.

### HasScaledMode

`func (o *ComputeBladeAllOf) HasScaledMode() bool`

HasScaledMode returns a boolean if a field has been set.

### GetSlotId

`func (o *ComputeBladeAllOf) GetSlotId() int64`

GetSlotId returns the SlotId field if non-nil, zero value otherwise.

### GetSlotIdOk

`func (o *ComputeBladeAllOf) GetSlotIdOk() (*int64, bool)`

GetSlotIdOk returns a tuple with the SlotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotId

`func (o *ComputeBladeAllOf) SetSlotId(v int64)`

SetSlotId sets SlotId field to given value.

### HasSlotId

`func (o *ComputeBladeAllOf) HasSlotId() bool`

HasSlotId returns a boolean if a field has been set.

### GetAdapters

`func (o *ComputeBladeAllOf) GetAdapters() []AdapterUnitRelationship`

GetAdapters returns the Adapters field if non-nil, zero value otherwise.

### GetAdaptersOk

`func (o *ComputeBladeAllOf) GetAdaptersOk() (*[]AdapterUnitRelationship, bool)`

GetAdaptersOk returns a tuple with the Adapters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapters

`func (o *ComputeBladeAllOf) SetAdapters(v []AdapterUnitRelationship)`

SetAdapters sets Adapters field to given value.

### HasAdapters

`func (o *ComputeBladeAllOf) HasAdapters() bool`

HasAdapters returns a boolean if a field has been set.

### SetAdaptersNil

`func (o *ComputeBladeAllOf) SetAdaptersNil(b bool)`

 SetAdaptersNil sets the value for Adapters to be an explicit nil

### UnsetAdapters
`func (o *ComputeBladeAllOf) UnsetAdapters()`

UnsetAdapters ensures that no value is present for Adapters, not even an explicit nil
### GetBiosUnits

`func (o *ComputeBladeAllOf) GetBiosUnits() []BiosUnitRelationship`

GetBiosUnits returns the BiosUnits field if non-nil, zero value otherwise.

### GetBiosUnitsOk

`func (o *ComputeBladeAllOf) GetBiosUnitsOk() (*[]BiosUnitRelationship, bool)`

GetBiosUnitsOk returns a tuple with the BiosUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiosUnits

`func (o *ComputeBladeAllOf) SetBiosUnits(v []BiosUnitRelationship)`

SetBiosUnits sets BiosUnits field to given value.

### HasBiosUnits

`func (o *ComputeBladeAllOf) HasBiosUnits() bool`

HasBiosUnits returns a boolean if a field has been set.

### SetBiosUnitsNil

`func (o *ComputeBladeAllOf) SetBiosUnitsNil(b bool)`

 SetBiosUnitsNil sets the value for BiosUnits to be an explicit nil

### UnsetBiosUnits
`func (o *ComputeBladeAllOf) UnsetBiosUnits()`

UnsetBiosUnits ensures that no value is present for BiosUnits, not even an explicit nil
### GetBmc

`func (o *ComputeBladeAllOf) GetBmc() ManagementControllerRelationship`

GetBmc returns the Bmc field if non-nil, zero value otherwise.

### GetBmcOk

`func (o *ComputeBladeAllOf) GetBmcOk() (*ManagementControllerRelationship, bool)`

GetBmcOk returns a tuple with the Bmc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBmc

`func (o *ComputeBladeAllOf) SetBmc(v ManagementControllerRelationship)`

SetBmc sets Bmc field to given value.

### HasBmc

`func (o *ComputeBladeAllOf) HasBmc() bool`

HasBmc returns a boolean if a field has been set.

### GetBoard

`func (o *ComputeBladeAllOf) GetBoard() ComputeBoardRelationship`

GetBoard returns the Board field if non-nil, zero value otherwise.

### GetBoardOk

`func (o *ComputeBladeAllOf) GetBoardOk() (*ComputeBoardRelationship, bool)`

GetBoardOk returns a tuple with the Board field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoard

`func (o *ComputeBladeAllOf) SetBoard(v ComputeBoardRelationship)`

SetBoard sets Board field to given value.

### HasBoard

`func (o *ComputeBladeAllOf) HasBoard() bool`

HasBoard returns a boolean if a field has been set.

### GetEquipmentChassis

`func (o *ComputeBladeAllOf) GetEquipmentChassis() EquipmentChassisRelationship`

GetEquipmentChassis returns the EquipmentChassis field if non-nil, zero value otherwise.

### GetEquipmentChassisOk

`func (o *ComputeBladeAllOf) GetEquipmentChassisOk() (*EquipmentChassisRelationship, bool)`

GetEquipmentChassisOk returns a tuple with the EquipmentChassis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentChassis

`func (o *ComputeBladeAllOf) SetEquipmentChassis(v EquipmentChassisRelationship)`

SetEquipmentChassis sets EquipmentChassis field to given value.

### HasEquipmentChassis

`func (o *ComputeBladeAllOf) HasEquipmentChassis() bool`

HasEquipmentChassis returns a boolean if a field has been set.

### GetEquipmentIoExpanders

`func (o *ComputeBladeAllOf) GetEquipmentIoExpanders() []EquipmentIoExpanderRelationship`

GetEquipmentIoExpanders returns the EquipmentIoExpanders field if non-nil, zero value otherwise.

### GetEquipmentIoExpandersOk

`func (o *ComputeBladeAllOf) GetEquipmentIoExpandersOk() (*[]EquipmentIoExpanderRelationship, bool)`

GetEquipmentIoExpandersOk returns a tuple with the EquipmentIoExpanders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentIoExpanders

`func (o *ComputeBladeAllOf) SetEquipmentIoExpanders(v []EquipmentIoExpanderRelationship)`

SetEquipmentIoExpanders sets EquipmentIoExpanders field to given value.

### HasEquipmentIoExpanders

`func (o *ComputeBladeAllOf) HasEquipmentIoExpanders() bool`

HasEquipmentIoExpanders returns a boolean if a field has been set.

### SetEquipmentIoExpandersNil

`func (o *ComputeBladeAllOf) SetEquipmentIoExpandersNil(b bool)`

 SetEquipmentIoExpandersNil sets the value for EquipmentIoExpanders to be an explicit nil

### UnsetEquipmentIoExpanders
`func (o *ComputeBladeAllOf) UnsetEquipmentIoExpanders()`

UnsetEquipmentIoExpanders ensures that no value is present for EquipmentIoExpanders, not even an explicit nil
### GetGenericInventoryHolders

`func (o *ComputeBladeAllOf) GetGenericInventoryHolders() []InventoryGenericInventoryHolderRelationship`

GetGenericInventoryHolders returns the GenericInventoryHolders field if non-nil, zero value otherwise.

### GetGenericInventoryHoldersOk

`func (o *ComputeBladeAllOf) GetGenericInventoryHoldersOk() (*[]InventoryGenericInventoryHolderRelationship, bool)`

GetGenericInventoryHoldersOk returns a tuple with the GenericInventoryHolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenericInventoryHolders

`func (o *ComputeBladeAllOf) SetGenericInventoryHolders(v []InventoryGenericInventoryHolderRelationship)`

SetGenericInventoryHolders sets GenericInventoryHolders field to given value.

### HasGenericInventoryHolders

`func (o *ComputeBladeAllOf) HasGenericInventoryHolders() bool`

HasGenericInventoryHolders returns a boolean if a field has been set.

### SetGenericInventoryHoldersNil

`func (o *ComputeBladeAllOf) SetGenericInventoryHoldersNil(b bool)`

 SetGenericInventoryHoldersNil sets the value for GenericInventoryHolders to be an explicit nil

### UnsetGenericInventoryHolders
`func (o *ComputeBladeAllOf) UnsetGenericInventoryHolders()`

UnsetGenericInventoryHolders ensures that no value is present for GenericInventoryHolders, not even an explicit nil
### GetGraphicsCards

`func (o *ComputeBladeAllOf) GetGraphicsCards() []GraphicsCardRelationship`

GetGraphicsCards returns the GraphicsCards field if non-nil, zero value otherwise.

### GetGraphicsCardsOk

`func (o *ComputeBladeAllOf) GetGraphicsCardsOk() (*[]GraphicsCardRelationship, bool)`

GetGraphicsCardsOk returns a tuple with the GraphicsCards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraphicsCards

`func (o *ComputeBladeAllOf) SetGraphicsCards(v []GraphicsCardRelationship)`

SetGraphicsCards sets GraphicsCards field to given value.

### HasGraphicsCards

`func (o *ComputeBladeAllOf) HasGraphicsCards() bool`

HasGraphicsCards returns a boolean if a field has been set.

### SetGraphicsCardsNil

`func (o *ComputeBladeAllOf) SetGraphicsCardsNil(b bool)`

 SetGraphicsCardsNil sets the value for GraphicsCards to be an explicit nil

### UnsetGraphicsCards
`func (o *ComputeBladeAllOf) UnsetGraphicsCards()`

UnsetGraphicsCards ensures that no value is present for GraphicsCards, not even an explicit nil
### GetInventoryDeviceInfo

`func (o *ComputeBladeAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *ComputeBladeAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *ComputeBladeAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *ComputeBladeAllOf) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetLocatorLed

`func (o *ComputeBladeAllOf) GetLocatorLed() EquipmentLocatorLedRelationship`

GetLocatorLed returns the LocatorLed field if non-nil, zero value otherwise.

### GetLocatorLedOk

`func (o *ComputeBladeAllOf) GetLocatorLedOk() (*EquipmentLocatorLedRelationship, bool)`

GetLocatorLedOk returns a tuple with the LocatorLed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocatorLed

`func (o *ComputeBladeAllOf) SetLocatorLed(v EquipmentLocatorLedRelationship)`

SetLocatorLed sets LocatorLed field to given value.

### HasLocatorLed

`func (o *ComputeBladeAllOf) HasLocatorLed() bool`

HasLocatorLed returns a boolean if a field has been set.

### GetMemoryArrays

`func (o *ComputeBladeAllOf) GetMemoryArrays() []MemoryArrayRelationship`

GetMemoryArrays returns the MemoryArrays field if non-nil, zero value otherwise.

### GetMemoryArraysOk

`func (o *ComputeBladeAllOf) GetMemoryArraysOk() (*[]MemoryArrayRelationship, bool)`

GetMemoryArraysOk returns a tuple with the MemoryArrays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryArrays

`func (o *ComputeBladeAllOf) SetMemoryArrays(v []MemoryArrayRelationship)`

SetMemoryArrays sets MemoryArrays field to given value.

### HasMemoryArrays

`func (o *ComputeBladeAllOf) HasMemoryArrays() bool`

HasMemoryArrays returns a boolean if a field has been set.

### SetMemoryArraysNil

`func (o *ComputeBladeAllOf) SetMemoryArraysNil(b bool)`

 SetMemoryArraysNil sets the value for MemoryArrays to be an explicit nil

### UnsetMemoryArrays
`func (o *ComputeBladeAllOf) UnsetMemoryArrays()`

UnsetMemoryArrays ensures that no value is present for MemoryArrays, not even an explicit nil
### GetPciDevices

`func (o *ComputeBladeAllOf) GetPciDevices() []PciDeviceRelationship`

GetPciDevices returns the PciDevices field if non-nil, zero value otherwise.

### GetPciDevicesOk

`func (o *ComputeBladeAllOf) GetPciDevicesOk() (*[]PciDeviceRelationship, bool)`

GetPciDevicesOk returns a tuple with the PciDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciDevices

`func (o *ComputeBladeAllOf) SetPciDevices(v []PciDeviceRelationship)`

SetPciDevices sets PciDevices field to given value.

### HasPciDevices

`func (o *ComputeBladeAllOf) HasPciDevices() bool`

HasPciDevices returns a boolean if a field has been set.

### SetPciDevicesNil

`func (o *ComputeBladeAllOf) SetPciDevicesNil(b bool)`

 SetPciDevicesNil sets the value for PciDevices to be an explicit nil

### UnsetPciDevices
`func (o *ComputeBladeAllOf) UnsetPciDevices()`

UnsetPciDevices ensures that no value is present for PciDevices, not even an explicit nil
### GetProcessors

`func (o *ComputeBladeAllOf) GetProcessors() []ProcessorUnitRelationship`

GetProcessors returns the Processors field if non-nil, zero value otherwise.

### GetProcessorsOk

`func (o *ComputeBladeAllOf) GetProcessorsOk() (*[]ProcessorUnitRelationship, bool)`

GetProcessorsOk returns a tuple with the Processors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessors

`func (o *ComputeBladeAllOf) SetProcessors(v []ProcessorUnitRelationship)`

SetProcessors sets Processors field to given value.

### HasProcessors

`func (o *ComputeBladeAllOf) HasProcessors() bool`

HasProcessors returns a boolean if a field has been set.

### SetProcessorsNil

`func (o *ComputeBladeAllOf) SetProcessorsNil(b bool)`

 SetProcessorsNil sets the value for Processors to be an explicit nil

### UnsetProcessors
`func (o *ComputeBladeAllOf) UnsetProcessors()`

UnsetProcessors ensures that no value is present for Processors, not even an explicit nil
### GetRegisteredDevice

`func (o *ComputeBladeAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *ComputeBladeAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *ComputeBladeAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *ComputeBladeAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### GetStorageControllers

`func (o *ComputeBladeAllOf) GetStorageControllers() []StorageControllerRelationship`

GetStorageControllers returns the StorageControllers field if non-nil, zero value otherwise.

### GetStorageControllersOk

`func (o *ComputeBladeAllOf) GetStorageControllersOk() (*[]StorageControllerRelationship, bool)`

GetStorageControllersOk returns a tuple with the StorageControllers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageControllers

`func (o *ComputeBladeAllOf) SetStorageControllers(v []StorageControllerRelationship)`

SetStorageControllers sets StorageControllers field to given value.

### HasStorageControllers

`func (o *ComputeBladeAllOf) HasStorageControllers() bool`

HasStorageControllers returns a boolean if a field has been set.

### SetStorageControllersNil

`func (o *ComputeBladeAllOf) SetStorageControllersNil(b bool)`

 SetStorageControllersNil sets the value for StorageControllers to be an explicit nil

### UnsetStorageControllers
`func (o *ComputeBladeAllOf) UnsetStorageControllers()`

UnsetStorageControllers ensures that no value is present for StorageControllers, not even an explicit nil
### GetStorageEnclosures

`func (o *ComputeBladeAllOf) GetStorageEnclosures() []StorageEnclosureRelationship`

GetStorageEnclosures returns the StorageEnclosures field if non-nil, zero value otherwise.

### GetStorageEnclosuresOk

`func (o *ComputeBladeAllOf) GetStorageEnclosuresOk() (*[]StorageEnclosureRelationship, bool)`

GetStorageEnclosuresOk returns a tuple with the StorageEnclosures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageEnclosures

`func (o *ComputeBladeAllOf) SetStorageEnclosures(v []StorageEnclosureRelationship)`

SetStorageEnclosures sets StorageEnclosures field to given value.

### HasStorageEnclosures

`func (o *ComputeBladeAllOf) HasStorageEnclosures() bool`

HasStorageEnclosures returns a boolean if a field has been set.

### SetStorageEnclosuresNil

`func (o *ComputeBladeAllOf) SetStorageEnclosuresNil(b bool)`

 SetStorageEnclosuresNil sets the value for StorageEnclosures to be an explicit nil

### UnsetStorageEnclosures
`func (o *ComputeBladeAllOf) UnsetStorageEnclosures()`

UnsetStorageEnclosures ensures that no value is present for StorageEnclosures, not even an explicit nil
### GetTopSystem

`func (o *ComputeBladeAllOf) GetTopSystem() TopSystemRelationship`

GetTopSystem returns the TopSystem field if non-nil, zero value otherwise.

### GetTopSystemOk

`func (o *ComputeBladeAllOf) GetTopSystemOk() (*TopSystemRelationship, bool)`

GetTopSystemOk returns a tuple with the TopSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopSystem

`func (o *ComputeBladeAllOf) SetTopSystem(v TopSystemRelationship)`

SetTopSystem sets TopSystem field to given value.

### HasTopSystem

`func (o *ComputeBladeAllOf) HasTopSystem() bool`

HasTopSystem returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



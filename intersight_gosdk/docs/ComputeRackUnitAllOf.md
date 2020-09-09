# ComputeRackUnitAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionStatus** | Pointer to **string** | Connectivity Status of RackUnit to Switch - A or B or AB. | [optional] [readonly] 
**ServerId** | Pointer to **int64** | RackUnit ID that uniquely identifies the server. | [optional] [readonly] 
**TopologyScanStatus** | Pointer to **string** | To maintain the Topology workflow run status. | [optional] 
**Adapters** | Pointer to [**[]AdapterUnitRelationship**](adapter.Unit.Relationship.md) | An array of relationships to adapterUnit resources. | [optional] [readonly] 
**BiosBootmode** | Pointer to [**BiosBootModeRelationship**](bios.BootMode.Relationship.md) |  | [optional] 
**Biosunits** | Pointer to [**[]BiosUnitRelationship**](bios.Unit.Relationship.md) | An array of relationships to biosUnit resources. | [optional] [readonly] 
**Bmc** | Pointer to [**ManagementControllerRelationship**](management.Controller.Relationship.md) |  | [optional] 
**Board** | Pointer to [**ComputeBoardRelationship**](compute.Board.Relationship.md) |  | [optional] 
**BootDeviceBootmode** | Pointer to [**BootDeviceBootModeRelationship**](boot.DeviceBootMode.Relationship.md) |  | [optional] 
**Fanmodules** | Pointer to [**[]EquipmentFanModuleRelationship**](equipment.FanModule.Relationship.md) | An array of relationships to equipmentFanModule resources. | [optional] [readonly] 
**GenericInventoryHolders** | Pointer to [**[]InventoryGenericInventoryHolderRelationship**](inventory.GenericInventoryHolder.Relationship.md) | An array of relationships to inventoryGenericInventoryHolder resources. | [optional] [readonly] 
**GraphicsCards** | Pointer to [**[]GraphicsCardRelationship**](graphics.Card.Relationship.md) | An array of relationships to graphicsCard resources. | [optional] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](inventory.DeviceInfo.Relationship.md) |  | [optional] 
**LocatorLed** | Pointer to [**EquipmentLocatorLedRelationship**](equipment.LocatorLed.Relationship.md) |  | [optional] 
**MemoryArrays** | Pointer to [**[]MemoryArrayRelationship**](memory.Array.Relationship.md) | An array of relationships to memoryArray resources. | [optional] 
**PciDevices** | Pointer to [**[]PciDeviceRelationship**](pci.Device.Relationship.md) | An array of relationships to pciDevice resources. | [optional] [readonly] 
**Processors** | Pointer to [**[]ProcessorUnitRelationship**](processor.Unit.Relationship.md) | An array of relationships to processorUnit resources. | [optional] 
**Psus** | Pointer to [**[]EquipmentPsuRelationship**](equipment.Psu.Relationship.md) | An array of relationships to equipmentPsu resources. | [optional] [readonly] 
**RackEnclosureSlot** | Pointer to [**EquipmentRackEnclosureSlotRelationship**](equipment.RackEnclosureSlot.Relationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 
**SasExpanders** | Pointer to [**[]StorageSasExpanderRelationship**](storage.SasExpander.Relationship.md) | An array of relationships to storageSasExpander resources. | [optional] 
**StorageControllers** | Pointer to [**[]StorageControllerRelationship**](storage.Controller.Relationship.md) | An array of relationships to storageController resources. | [optional] 
**StorageEnclosures** | Pointer to [**[]StorageEnclosureRelationship**](storage.Enclosure.Relationship.md) | An array of relationships to storageEnclosure resources. | [optional] [readonly] 
**TopSystem** | Pointer to [**TopSystemRelationship**](top.System.Relationship.md) |  | [optional] 

## Methods

### NewComputeRackUnitAllOf

`func NewComputeRackUnitAllOf() *ComputeRackUnitAllOf`

NewComputeRackUnitAllOf instantiates a new ComputeRackUnitAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputeRackUnitAllOfWithDefaults

`func NewComputeRackUnitAllOfWithDefaults() *ComputeRackUnitAllOf`

NewComputeRackUnitAllOfWithDefaults instantiates a new ComputeRackUnitAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionStatus

`func (o *ComputeRackUnitAllOf) GetConnectionStatus() string`

GetConnectionStatus returns the ConnectionStatus field if non-nil, zero value otherwise.

### GetConnectionStatusOk

`func (o *ComputeRackUnitAllOf) GetConnectionStatusOk() (*string, bool)`

GetConnectionStatusOk returns a tuple with the ConnectionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionStatus

`func (o *ComputeRackUnitAllOf) SetConnectionStatus(v string)`

SetConnectionStatus sets ConnectionStatus field to given value.

### HasConnectionStatus

`func (o *ComputeRackUnitAllOf) HasConnectionStatus() bool`

HasConnectionStatus returns a boolean if a field has been set.

### GetServerId

`func (o *ComputeRackUnitAllOf) GetServerId() int64`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *ComputeRackUnitAllOf) GetServerIdOk() (*int64, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *ComputeRackUnitAllOf) SetServerId(v int64)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *ComputeRackUnitAllOf) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetTopologyScanStatus

`func (o *ComputeRackUnitAllOf) GetTopologyScanStatus() string`

GetTopologyScanStatus returns the TopologyScanStatus field if non-nil, zero value otherwise.

### GetTopologyScanStatusOk

`func (o *ComputeRackUnitAllOf) GetTopologyScanStatusOk() (*string, bool)`

GetTopologyScanStatusOk returns a tuple with the TopologyScanStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopologyScanStatus

`func (o *ComputeRackUnitAllOf) SetTopologyScanStatus(v string)`

SetTopologyScanStatus sets TopologyScanStatus field to given value.

### HasTopologyScanStatus

`func (o *ComputeRackUnitAllOf) HasTopologyScanStatus() bool`

HasTopologyScanStatus returns a boolean if a field has been set.

### GetAdapters

`func (o *ComputeRackUnitAllOf) GetAdapters() []AdapterUnitRelationship`

GetAdapters returns the Adapters field if non-nil, zero value otherwise.

### GetAdaptersOk

`func (o *ComputeRackUnitAllOf) GetAdaptersOk() (*[]AdapterUnitRelationship, bool)`

GetAdaptersOk returns a tuple with the Adapters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapters

`func (o *ComputeRackUnitAllOf) SetAdapters(v []AdapterUnitRelationship)`

SetAdapters sets Adapters field to given value.

### HasAdapters

`func (o *ComputeRackUnitAllOf) HasAdapters() bool`

HasAdapters returns a boolean if a field has been set.

### SetAdaptersNil

`func (o *ComputeRackUnitAllOf) SetAdaptersNil(b bool)`

 SetAdaptersNil sets the value for Adapters to be an explicit nil

### UnsetAdapters
`func (o *ComputeRackUnitAllOf) UnsetAdapters()`

UnsetAdapters ensures that no value is present for Adapters, not even an explicit nil
### GetBiosBootmode

`func (o *ComputeRackUnitAllOf) GetBiosBootmode() BiosBootModeRelationship`

GetBiosBootmode returns the BiosBootmode field if non-nil, zero value otherwise.

### GetBiosBootmodeOk

`func (o *ComputeRackUnitAllOf) GetBiosBootmodeOk() (*BiosBootModeRelationship, bool)`

GetBiosBootmodeOk returns a tuple with the BiosBootmode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiosBootmode

`func (o *ComputeRackUnitAllOf) SetBiosBootmode(v BiosBootModeRelationship)`

SetBiosBootmode sets BiosBootmode field to given value.

### HasBiosBootmode

`func (o *ComputeRackUnitAllOf) HasBiosBootmode() bool`

HasBiosBootmode returns a boolean if a field has been set.

### GetBiosunits

`func (o *ComputeRackUnitAllOf) GetBiosunits() []BiosUnitRelationship`

GetBiosunits returns the Biosunits field if non-nil, zero value otherwise.

### GetBiosunitsOk

`func (o *ComputeRackUnitAllOf) GetBiosunitsOk() (*[]BiosUnitRelationship, bool)`

GetBiosunitsOk returns a tuple with the Biosunits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiosunits

`func (o *ComputeRackUnitAllOf) SetBiosunits(v []BiosUnitRelationship)`

SetBiosunits sets Biosunits field to given value.

### HasBiosunits

`func (o *ComputeRackUnitAllOf) HasBiosunits() bool`

HasBiosunits returns a boolean if a field has been set.

### SetBiosunitsNil

`func (o *ComputeRackUnitAllOf) SetBiosunitsNil(b bool)`

 SetBiosunitsNil sets the value for Biosunits to be an explicit nil

### UnsetBiosunits
`func (o *ComputeRackUnitAllOf) UnsetBiosunits()`

UnsetBiosunits ensures that no value is present for Biosunits, not even an explicit nil
### GetBmc

`func (o *ComputeRackUnitAllOf) GetBmc() ManagementControllerRelationship`

GetBmc returns the Bmc field if non-nil, zero value otherwise.

### GetBmcOk

`func (o *ComputeRackUnitAllOf) GetBmcOk() (*ManagementControllerRelationship, bool)`

GetBmcOk returns a tuple with the Bmc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBmc

`func (o *ComputeRackUnitAllOf) SetBmc(v ManagementControllerRelationship)`

SetBmc sets Bmc field to given value.

### HasBmc

`func (o *ComputeRackUnitAllOf) HasBmc() bool`

HasBmc returns a boolean if a field has been set.

### GetBoard

`func (o *ComputeRackUnitAllOf) GetBoard() ComputeBoardRelationship`

GetBoard returns the Board field if non-nil, zero value otherwise.

### GetBoardOk

`func (o *ComputeRackUnitAllOf) GetBoardOk() (*ComputeBoardRelationship, bool)`

GetBoardOk returns a tuple with the Board field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoard

`func (o *ComputeRackUnitAllOf) SetBoard(v ComputeBoardRelationship)`

SetBoard sets Board field to given value.

### HasBoard

`func (o *ComputeRackUnitAllOf) HasBoard() bool`

HasBoard returns a boolean if a field has been set.

### GetBootDeviceBootmode

`func (o *ComputeRackUnitAllOf) GetBootDeviceBootmode() BootDeviceBootModeRelationship`

GetBootDeviceBootmode returns the BootDeviceBootmode field if non-nil, zero value otherwise.

### GetBootDeviceBootmodeOk

`func (o *ComputeRackUnitAllOf) GetBootDeviceBootmodeOk() (*BootDeviceBootModeRelationship, bool)`

GetBootDeviceBootmodeOk returns a tuple with the BootDeviceBootmode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootDeviceBootmode

`func (o *ComputeRackUnitAllOf) SetBootDeviceBootmode(v BootDeviceBootModeRelationship)`

SetBootDeviceBootmode sets BootDeviceBootmode field to given value.

### HasBootDeviceBootmode

`func (o *ComputeRackUnitAllOf) HasBootDeviceBootmode() bool`

HasBootDeviceBootmode returns a boolean if a field has been set.

### GetFanmodules

`func (o *ComputeRackUnitAllOf) GetFanmodules() []EquipmentFanModuleRelationship`

GetFanmodules returns the Fanmodules field if non-nil, zero value otherwise.

### GetFanmodulesOk

`func (o *ComputeRackUnitAllOf) GetFanmodulesOk() (*[]EquipmentFanModuleRelationship, bool)`

GetFanmodulesOk returns a tuple with the Fanmodules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFanmodules

`func (o *ComputeRackUnitAllOf) SetFanmodules(v []EquipmentFanModuleRelationship)`

SetFanmodules sets Fanmodules field to given value.

### HasFanmodules

`func (o *ComputeRackUnitAllOf) HasFanmodules() bool`

HasFanmodules returns a boolean if a field has been set.

### SetFanmodulesNil

`func (o *ComputeRackUnitAllOf) SetFanmodulesNil(b bool)`

 SetFanmodulesNil sets the value for Fanmodules to be an explicit nil

### UnsetFanmodules
`func (o *ComputeRackUnitAllOf) UnsetFanmodules()`

UnsetFanmodules ensures that no value is present for Fanmodules, not even an explicit nil
### GetGenericInventoryHolders

`func (o *ComputeRackUnitAllOf) GetGenericInventoryHolders() []InventoryGenericInventoryHolderRelationship`

GetGenericInventoryHolders returns the GenericInventoryHolders field if non-nil, zero value otherwise.

### GetGenericInventoryHoldersOk

`func (o *ComputeRackUnitAllOf) GetGenericInventoryHoldersOk() (*[]InventoryGenericInventoryHolderRelationship, bool)`

GetGenericInventoryHoldersOk returns a tuple with the GenericInventoryHolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenericInventoryHolders

`func (o *ComputeRackUnitAllOf) SetGenericInventoryHolders(v []InventoryGenericInventoryHolderRelationship)`

SetGenericInventoryHolders sets GenericInventoryHolders field to given value.

### HasGenericInventoryHolders

`func (o *ComputeRackUnitAllOf) HasGenericInventoryHolders() bool`

HasGenericInventoryHolders returns a boolean if a field has been set.

### SetGenericInventoryHoldersNil

`func (o *ComputeRackUnitAllOf) SetGenericInventoryHoldersNil(b bool)`

 SetGenericInventoryHoldersNil sets the value for GenericInventoryHolders to be an explicit nil

### UnsetGenericInventoryHolders
`func (o *ComputeRackUnitAllOf) UnsetGenericInventoryHolders()`

UnsetGenericInventoryHolders ensures that no value is present for GenericInventoryHolders, not even an explicit nil
### GetGraphicsCards

`func (o *ComputeRackUnitAllOf) GetGraphicsCards() []GraphicsCardRelationship`

GetGraphicsCards returns the GraphicsCards field if non-nil, zero value otherwise.

### GetGraphicsCardsOk

`func (o *ComputeRackUnitAllOf) GetGraphicsCardsOk() (*[]GraphicsCardRelationship, bool)`

GetGraphicsCardsOk returns a tuple with the GraphicsCards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraphicsCards

`func (o *ComputeRackUnitAllOf) SetGraphicsCards(v []GraphicsCardRelationship)`

SetGraphicsCards sets GraphicsCards field to given value.

### HasGraphicsCards

`func (o *ComputeRackUnitAllOf) HasGraphicsCards() bool`

HasGraphicsCards returns a boolean if a field has been set.

### SetGraphicsCardsNil

`func (o *ComputeRackUnitAllOf) SetGraphicsCardsNil(b bool)`

 SetGraphicsCardsNil sets the value for GraphicsCards to be an explicit nil

### UnsetGraphicsCards
`func (o *ComputeRackUnitAllOf) UnsetGraphicsCards()`

UnsetGraphicsCards ensures that no value is present for GraphicsCards, not even an explicit nil
### GetInventoryDeviceInfo

`func (o *ComputeRackUnitAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *ComputeRackUnitAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *ComputeRackUnitAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *ComputeRackUnitAllOf) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetLocatorLed

`func (o *ComputeRackUnitAllOf) GetLocatorLed() EquipmentLocatorLedRelationship`

GetLocatorLed returns the LocatorLed field if non-nil, zero value otherwise.

### GetLocatorLedOk

`func (o *ComputeRackUnitAllOf) GetLocatorLedOk() (*EquipmentLocatorLedRelationship, bool)`

GetLocatorLedOk returns a tuple with the LocatorLed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocatorLed

`func (o *ComputeRackUnitAllOf) SetLocatorLed(v EquipmentLocatorLedRelationship)`

SetLocatorLed sets LocatorLed field to given value.

### HasLocatorLed

`func (o *ComputeRackUnitAllOf) HasLocatorLed() bool`

HasLocatorLed returns a boolean if a field has been set.

### GetMemoryArrays

`func (o *ComputeRackUnitAllOf) GetMemoryArrays() []MemoryArrayRelationship`

GetMemoryArrays returns the MemoryArrays field if non-nil, zero value otherwise.

### GetMemoryArraysOk

`func (o *ComputeRackUnitAllOf) GetMemoryArraysOk() (*[]MemoryArrayRelationship, bool)`

GetMemoryArraysOk returns a tuple with the MemoryArrays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryArrays

`func (o *ComputeRackUnitAllOf) SetMemoryArrays(v []MemoryArrayRelationship)`

SetMemoryArrays sets MemoryArrays field to given value.

### HasMemoryArrays

`func (o *ComputeRackUnitAllOf) HasMemoryArrays() bool`

HasMemoryArrays returns a boolean if a field has been set.

### SetMemoryArraysNil

`func (o *ComputeRackUnitAllOf) SetMemoryArraysNil(b bool)`

 SetMemoryArraysNil sets the value for MemoryArrays to be an explicit nil

### UnsetMemoryArrays
`func (o *ComputeRackUnitAllOf) UnsetMemoryArrays()`

UnsetMemoryArrays ensures that no value is present for MemoryArrays, not even an explicit nil
### GetPciDevices

`func (o *ComputeRackUnitAllOf) GetPciDevices() []PciDeviceRelationship`

GetPciDevices returns the PciDevices field if non-nil, zero value otherwise.

### GetPciDevicesOk

`func (o *ComputeRackUnitAllOf) GetPciDevicesOk() (*[]PciDeviceRelationship, bool)`

GetPciDevicesOk returns a tuple with the PciDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciDevices

`func (o *ComputeRackUnitAllOf) SetPciDevices(v []PciDeviceRelationship)`

SetPciDevices sets PciDevices field to given value.

### HasPciDevices

`func (o *ComputeRackUnitAllOf) HasPciDevices() bool`

HasPciDevices returns a boolean if a field has been set.

### SetPciDevicesNil

`func (o *ComputeRackUnitAllOf) SetPciDevicesNil(b bool)`

 SetPciDevicesNil sets the value for PciDevices to be an explicit nil

### UnsetPciDevices
`func (o *ComputeRackUnitAllOf) UnsetPciDevices()`

UnsetPciDevices ensures that no value is present for PciDevices, not even an explicit nil
### GetProcessors

`func (o *ComputeRackUnitAllOf) GetProcessors() []ProcessorUnitRelationship`

GetProcessors returns the Processors field if non-nil, zero value otherwise.

### GetProcessorsOk

`func (o *ComputeRackUnitAllOf) GetProcessorsOk() (*[]ProcessorUnitRelationship, bool)`

GetProcessorsOk returns a tuple with the Processors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessors

`func (o *ComputeRackUnitAllOf) SetProcessors(v []ProcessorUnitRelationship)`

SetProcessors sets Processors field to given value.

### HasProcessors

`func (o *ComputeRackUnitAllOf) HasProcessors() bool`

HasProcessors returns a boolean if a field has been set.

### SetProcessorsNil

`func (o *ComputeRackUnitAllOf) SetProcessorsNil(b bool)`

 SetProcessorsNil sets the value for Processors to be an explicit nil

### UnsetProcessors
`func (o *ComputeRackUnitAllOf) UnsetProcessors()`

UnsetProcessors ensures that no value is present for Processors, not even an explicit nil
### GetPsus

`func (o *ComputeRackUnitAllOf) GetPsus() []EquipmentPsuRelationship`

GetPsus returns the Psus field if non-nil, zero value otherwise.

### GetPsusOk

`func (o *ComputeRackUnitAllOf) GetPsusOk() (*[]EquipmentPsuRelationship, bool)`

GetPsusOk returns a tuple with the Psus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPsus

`func (o *ComputeRackUnitAllOf) SetPsus(v []EquipmentPsuRelationship)`

SetPsus sets Psus field to given value.

### HasPsus

`func (o *ComputeRackUnitAllOf) HasPsus() bool`

HasPsus returns a boolean if a field has been set.

### SetPsusNil

`func (o *ComputeRackUnitAllOf) SetPsusNil(b bool)`

 SetPsusNil sets the value for Psus to be an explicit nil

### UnsetPsus
`func (o *ComputeRackUnitAllOf) UnsetPsus()`

UnsetPsus ensures that no value is present for Psus, not even an explicit nil
### GetRackEnclosureSlot

`func (o *ComputeRackUnitAllOf) GetRackEnclosureSlot() EquipmentRackEnclosureSlotRelationship`

GetRackEnclosureSlot returns the RackEnclosureSlot field if non-nil, zero value otherwise.

### GetRackEnclosureSlotOk

`func (o *ComputeRackUnitAllOf) GetRackEnclosureSlotOk() (*EquipmentRackEnclosureSlotRelationship, bool)`

GetRackEnclosureSlotOk returns a tuple with the RackEnclosureSlot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRackEnclosureSlot

`func (o *ComputeRackUnitAllOf) SetRackEnclosureSlot(v EquipmentRackEnclosureSlotRelationship)`

SetRackEnclosureSlot sets RackEnclosureSlot field to given value.

### HasRackEnclosureSlot

`func (o *ComputeRackUnitAllOf) HasRackEnclosureSlot() bool`

HasRackEnclosureSlot returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *ComputeRackUnitAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *ComputeRackUnitAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *ComputeRackUnitAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *ComputeRackUnitAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### GetSasExpanders

`func (o *ComputeRackUnitAllOf) GetSasExpanders() []StorageSasExpanderRelationship`

GetSasExpanders returns the SasExpanders field if non-nil, zero value otherwise.

### GetSasExpandersOk

`func (o *ComputeRackUnitAllOf) GetSasExpandersOk() (*[]StorageSasExpanderRelationship, bool)`

GetSasExpandersOk returns a tuple with the SasExpanders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSasExpanders

`func (o *ComputeRackUnitAllOf) SetSasExpanders(v []StorageSasExpanderRelationship)`

SetSasExpanders sets SasExpanders field to given value.

### HasSasExpanders

`func (o *ComputeRackUnitAllOf) HasSasExpanders() bool`

HasSasExpanders returns a boolean if a field has been set.

### SetSasExpandersNil

`func (o *ComputeRackUnitAllOf) SetSasExpandersNil(b bool)`

 SetSasExpandersNil sets the value for SasExpanders to be an explicit nil

### UnsetSasExpanders
`func (o *ComputeRackUnitAllOf) UnsetSasExpanders()`

UnsetSasExpanders ensures that no value is present for SasExpanders, not even an explicit nil
### GetStorageControllers

`func (o *ComputeRackUnitAllOf) GetStorageControllers() []StorageControllerRelationship`

GetStorageControllers returns the StorageControllers field if non-nil, zero value otherwise.

### GetStorageControllersOk

`func (o *ComputeRackUnitAllOf) GetStorageControllersOk() (*[]StorageControllerRelationship, bool)`

GetStorageControllersOk returns a tuple with the StorageControllers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageControllers

`func (o *ComputeRackUnitAllOf) SetStorageControllers(v []StorageControllerRelationship)`

SetStorageControllers sets StorageControllers field to given value.

### HasStorageControllers

`func (o *ComputeRackUnitAllOf) HasStorageControllers() bool`

HasStorageControllers returns a boolean if a field has been set.

### SetStorageControllersNil

`func (o *ComputeRackUnitAllOf) SetStorageControllersNil(b bool)`

 SetStorageControllersNil sets the value for StorageControllers to be an explicit nil

### UnsetStorageControllers
`func (o *ComputeRackUnitAllOf) UnsetStorageControllers()`

UnsetStorageControllers ensures that no value is present for StorageControllers, not even an explicit nil
### GetStorageEnclosures

`func (o *ComputeRackUnitAllOf) GetStorageEnclosures() []StorageEnclosureRelationship`

GetStorageEnclosures returns the StorageEnclosures field if non-nil, zero value otherwise.

### GetStorageEnclosuresOk

`func (o *ComputeRackUnitAllOf) GetStorageEnclosuresOk() (*[]StorageEnclosureRelationship, bool)`

GetStorageEnclosuresOk returns a tuple with the StorageEnclosures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageEnclosures

`func (o *ComputeRackUnitAllOf) SetStorageEnclosures(v []StorageEnclosureRelationship)`

SetStorageEnclosures sets StorageEnclosures field to given value.

### HasStorageEnclosures

`func (o *ComputeRackUnitAllOf) HasStorageEnclosures() bool`

HasStorageEnclosures returns a boolean if a field has been set.

### SetStorageEnclosuresNil

`func (o *ComputeRackUnitAllOf) SetStorageEnclosuresNil(b bool)`

 SetStorageEnclosuresNil sets the value for StorageEnclosures to be an explicit nil

### UnsetStorageEnclosures
`func (o *ComputeRackUnitAllOf) UnsetStorageEnclosures()`

UnsetStorageEnclosures ensures that no value is present for StorageEnclosures, not even an explicit nil
### GetTopSystem

`func (o *ComputeRackUnitAllOf) GetTopSystem() TopSystemRelationship`

GetTopSystem returns the TopSystem field if non-nil, zero value otherwise.

### GetTopSystemOk

`func (o *ComputeRackUnitAllOf) GetTopSystemOk() (*TopSystemRelationship, bool)`

GetTopSystemOk returns a tuple with the TopSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopSystem

`func (o *ComputeRackUnitAllOf) SetTopSystem(v TopSystemRelationship)`

SetTopSystem sets TopSystem field to given value.

### HasTopSystem

`func (o *ComputeRackUnitAllOf) HasTopSystem() bool`

HasTopSystem returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



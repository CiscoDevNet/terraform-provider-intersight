# ComputeRackUnit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "compute.RackUnit"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "compute.RackUnit"]
**ConnectionStatus** | Pointer to **string** | Connectivity Status of RackUnit to Switch - A or B or AB. | [optional] [readonly] 
**ServerId** | Pointer to **int64** | RackUnit ID that uniquely identifies the server. | [optional] [readonly] 
**TopologyScanStatus** | Pointer to **string** | To maintain the Topology workflow run status. | [optional] 
**Adapters** | Pointer to [**[]AdapterUnitRelationship**](AdapterUnitRelationship.md) | An array of relationships to adapterUnit resources. | [optional] [readonly] 
**BiosBootmode** | Pointer to [**NullableBiosBootModeRelationship**](BiosBootModeRelationship.md) |  | [optional] 
**BiosTokenSettings** | Pointer to [**NullableBiosTokenSettingsRelationship**](BiosTokenSettingsRelationship.md) |  | [optional] 
**BiosVfSelectMemoryRasConfiguration** | Pointer to [**NullableBiosVfSelectMemoryRasConfigurationRelationship**](BiosVfSelectMemoryRasConfigurationRelationship.md) |  | [optional] 
**Biosunits** | Pointer to [**[]BiosUnitRelationship**](BiosUnitRelationship.md) | An array of relationships to biosUnit resources. | [optional] [readonly] 
**Bmc** | Pointer to [**NullableManagementControllerRelationship**](ManagementControllerRelationship.md) |  | [optional] 
**Board** | Pointer to [**NullableComputeBoardRelationship**](ComputeBoardRelationship.md) |  | [optional] 
**BootDeviceBootmode** | Pointer to [**NullableBootDeviceBootModeRelationship**](BootDeviceBootModeRelationship.md) |  | [optional] 
**ComputePersonality** | Pointer to [**[]ComputePersonalityRelationship**](ComputePersonalityRelationship.md) | An array of relationships to computePersonality resources. | [optional] [readonly] 
**Fanmodules** | Pointer to [**[]EquipmentFanModuleRelationship**](EquipmentFanModuleRelationship.md) | An array of relationships to equipmentFanModule resources. | [optional] [readonly] 
**GenericInventoryHolders** | Pointer to [**[]InventoryGenericInventoryHolderRelationship**](InventoryGenericInventoryHolderRelationship.md) | An array of relationships to inventoryGenericInventoryHolder resources. | [optional] [readonly] 
**GraphicsCards** | Pointer to [**[]GraphicsCardRelationship**](GraphicsCardRelationship.md) | An array of relationships to graphicsCard resources. | [optional] 
**HybridDriveSlots** | Pointer to [**[]EquipmentHybridDriveSlotRelationship**](EquipmentHybridDriveSlotRelationship.md) | An array of relationships to equipmentHybridDriveSlot resources. | [optional] [readonly] 
**InventoryDeviceInfo** | Pointer to [**NullableInventoryDeviceInfoRelationship**](InventoryDeviceInfoRelationship.md) |  | [optional] 
**LocatorLed** | Pointer to [**NullableEquipmentLocatorLedRelationship**](EquipmentLocatorLedRelationship.md) |  | [optional] 
**MemoryArrays** | Pointer to [**[]MemoryArrayRelationship**](MemoryArrayRelationship.md) | An array of relationships to memoryArray resources. | [optional] 
**PciDevices** | Pointer to [**[]PciDeviceRelationship**](PciDeviceRelationship.md) | An array of relationships to pciDevice resources. | [optional] [readonly] 
**Processors** | Pointer to [**[]ProcessorUnitRelationship**](ProcessorUnitRelationship.md) | An array of relationships to processorUnit resources. | [optional] 
**Psus** | Pointer to [**[]EquipmentPsuRelationship**](EquipmentPsuRelationship.md) | An array of relationships to equipmentPsu resources. | [optional] [readonly] 
**RackEnclosureSlot** | Pointer to [**NullableEquipmentRackEnclosureSlotRelationship**](EquipmentRackEnclosureSlotRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 
**SasExpanders** | Pointer to [**[]StorageSasExpanderRelationship**](StorageSasExpanderRelationship.md) | An array of relationships to storageSasExpander resources. | [optional] 
**StorageControllers** | Pointer to [**[]StorageControllerRelationship**](StorageControllerRelationship.md) | An array of relationships to storageController resources. | [optional] 
**StorageEnclosures** | Pointer to [**[]StorageEnclosureRelationship**](StorageEnclosureRelationship.md) | An array of relationships to storageEnclosure resources. | [optional] [readonly] 
**TopSystem** | Pointer to [**NullableTopSystemRelationship**](TopSystemRelationship.md) |  | [optional] 
**UnitPersonality** | Pointer to [**[]RackUnitPersonalityRelationship**](RackUnitPersonalityRelationship.md) | An array of relationships to rackUnitPersonality resources. | [optional] [readonly] 

## Methods

### NewComputeRackUnit

`func NewComputeRackUnit(classId string, objectType string, ) *ComputeRackUnit`

NewComputeRackUnit instantiates a new ComputeRackUnit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputeRackUnitWithDefaults

`func NewComputeRackUnitWithDefaults() *ComputeRackUnit`

NewComputeRackUnitWithDefaults instantiates a new ComputeRackUnit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ComputeRackUnit) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ComputeRackUnit) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ComputeRackUnit) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ComputeRackUnit) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ComputeRackUnit) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ComputeRackUnit) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetConnectionStatus

`func (o *ComputeRackUnit) GetConnectionStatus() string`

GetConnectionStatus returns the ConnectionStatus field if non-nil, zero value otherwise.

### GetConnectionStatusOk

`func (o *ComputeRackUnit) GetConnectionStatusOk() (*string, bool)`

GetConnectionStatusOk returns a tuple with the ConnectionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionStatus

`func (o *ComputeRackUnit) SetConnectionStatus(v string)`

SetConnectionStatus sets ConnectionStatus field to given value.

### HasConnectionStatus

`func (o *ComputeRackUnit) HasConnectionStatus() bool`

HasConnectionStatus returns a boolean if a field has been set.

### GetServerId

`func (o *ComputeRackUnit) GetServerId() int64`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *ComputeRackUnit) GetServerIdOk() (*int64, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *ComputeRackUnit) SetServerId(v int64)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *ComputeRackUnit) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetTopologyScanStatus

`func (o *ComputeRackUnit) GetTopologyScanStatus() string`

GetTopologyScanStatus returns the TopologyScanStatus field if non-nil, zero value otherwise.

### GetTopologyScanStatusOk

`func (o *ComputeRackUnit) GetTopologyScanStatusOk() (*string, bool)`

GetTopologyScanStatusOk returns a tuple with the TopologyScanStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopologyScanStatus

`func (o *ComputeRackUnit) SetTopologyScanStatus(v string)`

SetTopologyScanStatus sets TopologyScanStatus field to given value.

### HasTopologyScanStatus

`func (o *ComputeRackUnit) HasTopologyScanStatus() bool`

HasTopologyScanStatus returns a boolean if a field has been set.

### GetAdapters

`func (o *ComputeRackUnit) GetAdapters() []AdapterUnitRelationship`

GetAdapters returns the Adapters field if non-nil, zero value otherwise.

### GetAdaptersOk

`func (o *ComputeRackUnit) GetAdaptersOk() (*[]AdapterUnitRelationship, bool)`

GetAdaptersOk returns a tuple with the Adapters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapters

`func (o *ComputeRackUnit) SetAdapters(v []AdapterUnitRelationship)`

SetAdapters sets Adapters field to given value.

### HasAdapters

`func (o *ComputeRackUnit) HasAdapters() bool`

HasAdapters returns a boolean if a field has been set.

### SetAdaptersNil

`func (o *ComputeRackUnit) SetAdaptersNil(b bool)`

 SetAdaptersNil sets the value for Adapters to be an explicit nil

### UnsetAdapters
`func (o *ComputeRackUnit) UnsetAdapters()`

UnsetAdapters ensures that no value is present for Adapters, not even an explicit nil
### GetBiosBootmode

`func (o *ComputeRackUnit) GetBiosBootmode() BiosBootModeRelationship`

GetBiosBootmode returns the BiosBootmode field if non-nil, zero value otherwise.

### GetBiosBootmodeOk

`func (o *ComputeRackUnit) GetBiosBootmodeOk() (*BiosBootModeRelationship, bool)`

GetBiosBootmodeOk returns a tuple with the BiosBootmode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiosBootmode

`func (o *ComputeRackUnit) SetBiosBootmode(v BiosBootModeRelationship)`

SetBiosBootmode sets BiosBootmode field to given value.

### HasBiosBootmode

`func (o *ComputeRackUnit) HasBiosBootmode() bool`

HasBiosBootmode returns a boolean if a field has been set.

### SetBiosBootmodeNil

`func (o *ComputeRackUnit) SetBiosBootmodeNil(b bool)`

 SetBiosBootmodeNil sets the value for BiosBootmode to be an explicit nil

### UnsetBiosBootmode
`func (o *ComputeRackUnit) UnsetBiosBootmode()`

UnsetBiosBootmode ensures that no value is present for BiosBootmode, not even an explicit nil
### GetBiosTokenSettings

`func (o *ComputeRackUnit) GetBiosTokenSettings() BiosTokenSettingsRelationship`

GetBiosTokenSettings returns the BiosTokenSettings field if non-nil, zero value otherwise.

### GetBiosTokenSettingsOk

`func (o *ComputeRackUnit) GetBiosTokenSettingsOk() (*BiosTokenSettingsRelationship, bool)`

GetBiosTokenSettingsOk returns a tuple with the BiosTokenSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiosTokenSettings

`func (o *ComputeRackUnit) SetBiosTokenSettings(v BiosTokenSettingsRelationship)`

SetBiosTokenSettings sets BiosTokenSettings field to given value.

### HasBiosTokenSettings

`func (o *ComputeRackUnit) HasBiosTokenSettings() bool`

HasBiosTokenSettings returns a boolean if a field has been set.

### SetBiosTokenSettingsNil

`func (o *ComputeRackUnit) SetBiosTokenSettingsNil(b bool)`

 SetBiosTokenSettingsNil sets the value for BiosTokenSettings to be an explicit nil

### UnsetBiosTokenSettings
`func (o *ComputeRackUnit) UnsetBiosTokenSettings()`

UnsetBiosTokenSettings ensures that no value is present for BiosTokenSettings, not even an explicit nil
### GetBiosVfSelectMemoryRasConfiguration

`func (o *ComputeRackUnit) GetBiosVfSelectMemoryRasConfiguration() BiosVfSelectMemoryRasConfigurationRelationship`

GetBiosVfSelectMemoryRasConfiguration returns the BiosVfSelectMemoryRasConfiguration field if non-nil, zero value otherwise.

### GetBiosVfSelectMemoryRasConfigurationOk

`func (o *ComputeRackUnit) GetBiosVfSelectMemoryRasConfigurationOk() (*BiosVfSelectMemoryRasConfigurationRelationship, bool)`

GetBiosVfSelectMemoryRasConfigurationOk returns a tuple with the BiosVfSelectMemoryRasConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiosVfSelectMemoryRasConfiguration

`func (o *ComputeRackUnit) SetBiosVfSelectMemoryRasConfiguration(v BiosVfSelectMemoryRasConfigurationRelationship)`

SetBiosVfSelectMemoryRasConfiguration sets BiosVfSelectMemoryRasConfiguration field to given value.

### HasBiosVfSelectMemoryRasConfiguration

`func (o *ComputeRackUnit) HasBiosVfSelectMemoryRasConfiguration() bool`

HasBiosVfSelectMemoryRasConfiguration returns a boolean if a field has been set.

### SetBiosVfSelectMemoryRasConfigurationNil

`func (o *ComputeRackUnit) SetBiosVfSelectMemoryRasConfigurationNil(b bool)`

 SetBiosVfSelectMemoryRasConfigurationNil sets the value for BiosVfSelectMemoryRasConfiguration to be an explicit nil

### UnsetBiosVfSelectMemoryRasConfiguration
`func (o *ComputeRackUnit) UnsetBiosVfSelectMemoryRasConfiguration()`

UnsetBiosVfSelectMemoryRasConfiguration ensures that no value is present for BiosVfSelectMemoryRasConfiguration, not even an explicit nil
### GetBiosunits

`func (o *ComputeRackUnit) GetBiosunits() []BiosUnitRelationship`

GetBiosunits returns the Biosunits field if non-nil, zero value otherwise.

### GetBiosunitsOk

`func (o *ComputeRackUnit) GetBiosunitsOk() (*[]BiosUnitRelationship, bool)`

GetBiosunitsOk returns a tuple with the Biosunits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiosunits

`func (o *ComputeRackUnit) SetBiosunits(v []BiosUnitRelationship)`

SetBiosunits sets Biosunits field to given value.

### HasBiosunits

`func (o *ComputeRackUnit) HasBiosunits() bool`

HasBiosunits returns a boolean if a field has been set.

### SetBiosunitsNil

`func (o *ComputeRackUnit) SetBiosunitsNil(b bool)`

 SetBiosunitsNil sets the value for Biosunits to be an explicit nil

### UnsetBiosunits
`func (o *ComputeRackUnit) UnsetBiosunits()`

UnsetBiosunits ensures that no value is present for Biosunits, not even an explicit nil
### GetBmc

`func (o *ComputeRackUnit) GetBmc() ManagementControllerRelationship`

GetBmc returns the Bmc field if non-nil, zero value otherwise.

### GetBmcOk

`func (o *ComputeRackUnit) GetBmcOk() (*ManagementControllerRelationship, bool)`

GetBmcOk returns a tuple with the Bmc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBmc

`func (o *ComputeRackUnit) SetBmc(v ManagementControllerRelationship)`

SetBmc sets Bmc field to given value.

### HasBmc

`func (o *ComputeRackUnit) HasBmc() bool`

HasBmc returns a boolean if a field has been set.

### SetBmcNil

`func (o *ComputeRackUnit) SetBmcNil(b bool)`

 SetBmcNil sets the value for Bmc to be an explicit nil

### UnsetBmc
`func (o *ComputeRackUnit) UnsetBmc()`

UnsetBmc ensures that no value is present for Bmc, not even an explicit nil
### GetBoard

`func (o *ComputeRackUnit) GetBoard() ComputeBoardRelationship`

GetBoard returns the Board field if non-nil, zero value otherwise.

### GetBoardOk

`func (o *ComputeRackUnit) GetBoardOk() (*ComputeBoardRelationship, bool)`

GetBoardOk returns a tuple with the Board field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoard

`func (o *ComputeRackUnit) SetBoard(v ComputeBoardRelationship)`

SetBoard sets Board field to given value.

### HasBoard

`func (o *ComputeRackUnit) HasBoard() bool`

HasBoard returns a boolean if a field has been set.

### SetBoardNil

`func (o *ComputeRackUnit) SetBoardNil(b bool)`

 SetBoardNil sets the value for Board to be an explicit nil

### UnsetBoard
`func (o *ComputeRackUnit) UnsetBoard()`

UnsetBoard ensures that no value is present for Board, not even an explicit nil
### GetBootDeviceBootmode

`func (o *ComputeRackUnit) GetBootDeviceBootmode() BootDeviceBootModeRelationship`

GetBootDeviceBootmode returns the BootDeviceBootmode field if non-nil, zero value otherwise.

### GetBootDeviceBootmodeOk

`func (o *ComputeRackUnit) GetBootDeviceBootmodeOk() (*BootDeviceBootModeRelationship, bool)`

GetBootDeviceBootmodeOk returns a tuple with the BootDeviceBootmode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootDeviceBootmode

`func (o *ComputeRackUnit) SetBootDeviceBootmode(v BootDeviceBootModeRelationship)`

SetBootDeviceBootmode sets BootDeviceBootmode field to given value.

### HasBootDeviceBootmode

`func (o *ComputeRackUnit) HasBootDeviceBootmode() bool`

HasBootDeviceBootmode returns a boolean if a field has been set.

### SetBootDeviceBootmodeNil

`func (o *ComputeRackUnit) SetBootDeviceBootmodeNil(b bool)`

 SetBootDeviceBootmodeNil sets the value for BootDeviceBootmode to be an explicit nil

### UnsetBootDeviceBootmode
`func (o *ComputeRackUnit) UnsetBootDeviceBootmode()`

UnsetBootDeviceBootmode ensures that no value is present for BootDeviceBootmode, not even an explicit nil
### GetComputePersonality

`func (o *ComputeRackUnit) GetComputePersonality() []ComputePersonalityRelationship`

GetComputePersonality returns the ComputePersonality field if non-nil, zero value otherwise.

### GetComputePersonalityOk

`func (o *ComputeRackUnit) GetComputePersonalityOk() (*[]ComputePersonalityRelationship, bool)`

GetComputePersonalityOk returns a tuple with the ComputePersonality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputePersonality

`func (o *ComputeRackUnit) SetComputePersonality(v []ComputePersonalityRelationship)`

SetComputePersonality sets ComputePersonality field to given value.

### HasComputePersonality

`func (o *ComputeRackUnit) HasComputePersonality() bool`

HasComputePersonality returns a boolean if a field has been set.

### SetComputePersonalityNil

`func (o *ComputeRackUnit) SetComputePersonalityNil(b bool)`

 SetComputePersonalityNil sets the value for ComputePersonality to be an explicit nil

### UnsetComputePersonality
`func (o *ComputeRackUnit) UnsetComputePersonality()`

UnsetComputePersonality ensures that no value is present for ComputePersonality, not even an explicit nil
### GetFanmodules

`func (o *ComputeRackUnit) GetFanmodules() []EquipmentFanModuleRelationship`

GetFanmodules returns the Fanmodules field if non-nil, zero value otherwise.

### GetFanmodulesOk

`func (o *ComputeRackUnit) GetFanmodulesOk() (*[]EquipmentFanModuleRelationship, bool)`

GetFanmodulesOk returns a tuple with the Fanmodules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFanmodules

`func (o *ComputeRackUnit) SetFanmodules(v []EquipmentFanModuleRelationship)`

SetFanmodules sets Fanmodules field to given value.

### HasFanmodules

`func (o *ComputeRackUnit) HasFanmodules() bool`

HasFanmodules returns a boolean if a field has been set.

### SetFanmodulesNil

`func (o *ComputeRackUnit) SetFanmodulesNil(b bool)`

 SetFanmodulesNil sets the value for Fanmodules to be an explicit nil

### UnsetFanmodules
`func (o *ComputeRackUnit) UnsetFanmodules()`

UnsetFanmodules ensures that no value is present for Fanmodules, not even an explicit nil
### GetGenericInventoryHolders

`func (o *ComputeRackUnit) GetGenericInventoryHolders() []InventoryGenericInventoryHolderRelationship`

GetGenericInventoryHolders returns the GenericInventoryHolders field if non-nil, zero value otherwise.

### GetGenericInventoryHoldersOk

`func (o *ComputeRackUnit) GetGenericInventoryHoldersOk() (*[]InventoryGenericInventoryHolderRelationship, bool)`

GetGenericInventoryHoldersOk returns a tuple with the GenericInventoryHolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenericInventoryHolders

`func (o *ComputeRackUnit) SetGenericInventoryHolders(v []InventoryGenericInventoryHolderRelationship)`

SetGenericInventoryHolders sets GenericInventoryHolders field to given value.

### HasGenericInventoryHolders

`func (o *ComputeRackUnit) HasGenericInventoryHolders() bool`

HasGenericInventoryHolders returns a boolean if a field has been set.

### SetGenericInventoryHoldersNil

`func (o *ComputeRackUnit) SetGenericInventoryHoldersNil(b bool)`

 SetGenericInventoryHoldersNil sets the value for GenericInventoryHolders to be an explicit nil

### UnsetGenericInventoryHolders
`func (o *ComputeRackUnit) UnsetGenericInventoryHolders()`

UnsetGenericInventoryHolders ensures that no value is present for GenericInventoryHolders, not even an explicit nil
### GetGraphicsCards

`func (o *ComputeRackUnit) GetGraphicsCards() []GraphicsCardRelationship`

GetGraphicsCards returns the GraphicsCards field if non-nil, zero value otherwise.

### GetGraphicsCardsOk

`func (o *ComputeRackUnit) GetGraphicsCardsOk() (*[]GraphicsCardRelationship, bool)`

GetGraphicsCardsOk returns a tuple with the GraphicsCards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraphicsCards

`func (o *ComputeRackUnit) SetGraphicsCards(v []GraphicsCardRelationship)`

SetGraphicsCards sets GraphicsCards field to given value.

### HasGraphicsCards

`func (o *ComputeRackUnit) HasGraphicsCards() bool`

HasGraphicsCards returns a boolean if a field has been set.

### SetGraphicsCardsNil

`func (o *ComputeRackUnit) SetGraphicsCardsNil(b bool)`

 SetGraphicsCardsNil sets the value for GraphicsCards to be an explicit nil

### UnsetGraphicsCards
`func (o *ComputeRackUnit) UnsetGraphicsCards()`

UnsetGraphicsCards ensures that no value is present for GraphicsCards, not even an explicit nil
### GetHybridDriveSlots

`func (o *ComputeRackUnit) GetHybridDriveSlots() []EquipmentHybridDriveSlotRelationship`

GetHybridDriveSlots returns the HybridDriveSlots field if non-nil, zero value otherwise.

### GetHybridDriveSlotsOk

`func (o *ComputeRackUnit) GetHybridDriveSlotsOk() (*[]EquipmentHybridDriveSlotRelationship, bool)`

GetHybridDriveSlotsOk returns a tuple with the HybridDriveSlots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHybridDriveSlots

`func (o *ComputeRackUnit) SetHybridDriveSlots(v []EquipmentHybridDriveSlotRelationship)`

SetHybridDriveSlots sets HybridDriveSlots field to given value.

### HasHybridDriveSlots

`func (o *ComputeRackUnit) HasHybridDriveSlots() bool`

HasHybridDriveSlots returns a boolean if a field has been set.

### SetHybridDriveSlotsNil

`func (o *ComputeRackUnit) SetHybridDriveSlotsNil(b bool)`

 SetHybridDriveSlotsNil sets the value for HybridDriveSlots to be an explicit nil

### UnsetHybridDriveSlots
`func (o *ComputeRackUnit) UnsetHybridDriveSlots()`

UnsetHybridDriveSlots ensures that no value is present for HybridDriveSlots, not even an explicit nil
### GetInventoryDeviceInfo

`func (o *ComputeRackUnit) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *ComputeRackUnit) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *ComputeRackUnit) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *ComputeRackUnit) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### SetInventoryDeviceInfoNil

`func (o *ComputeRackUnit) SetInventoryDeviceInfoNil(b bool)`

 SetInventoryDeviceInfoNil sets the value for InventoryDeviceInfo to be an explicit nil

### UnsetInventoryDeviceInfo
`func (o *ComputeRackUnit) UnsetInventoryDeviceInfo()`

UnsetInventoryDeviceInfo ensures that no value is present for InventoryDeviceInfo, not even an explicit nil
### GetLocatorLed

`func (o *ComputeRackUnit) GetLocatorLed() EquipmentLocatorLedRelationship`

GetLocatorLed returns the LocatorLed field if non-nil, zero value otherwise.

### GetLocatorLedOk

`func (o *ComputeRackUnit) GetLocatorLedOk() (*EquipmentLocatorLedRelationship, bool)`

GetLocatorLedOk returns a tuple with the LocatorLed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocatorLed

`func (o *ComputeRackUnit) SetLocatorLed(v EquipmentLocatorLedRelationship)`

SetLocatorLed sets LocatorLed field to given value.

### HasLocatorLed

`func (o *ComputeRackUnit) HasLocatorLed() bool`

HasLocatorLed returns a boolean if a field has been set.

### SetLocatorLedNil

`func (o *ComputeRackUnit) SetLocatorLedNil(b bool)`

 SetLocatorLedNil sets the value for LocatorLed to be an explicit nil

### UnsetLocatorLed
`func (o *ComputeRackUnit) UnsetLocatorLed()`

UnsetLocatorLed ensures that no value is present for LocatorLed, not even an explicit nil
### GetMemoryArrays

`func (o *ComputeRackUnit) GetMemoryArrays() []MemoryArrayRelationship`

GetMemoryArrays returns the MemoryArrays field if non-nil, zero value otherwise.

### GetMemoryArraysOk

`func (o *ComputeRackUnit) GetMemoryArraysOk() (*[]MemoryArrayRelationship, bool)`

GetMemoryArraysOk returns a tuple with the MemoryArrays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryArrays

`func (o *ComputeRackUnit) SetMemoryArrays(v []MemoryArrayRelationship)`

SetMemoryArrays sets MemoryArrays field to given value.

### HasMemoryArrays

`func (o *ComputeRackUnit) HasMemoryArrays() bool`

HasMemoryArrays returns a boolean if a field has been set.

### SetMemoryArraysNil

`func (o *ComputeRackUnit) SetMemoryArraysNil(b bool)`

 SetMemoryArraysNil sets the value for MemoryArrays to be an explicit nil

### UnsetMemoryArrays
`func (o *ComputeRackUnit) UnsetMemoryArrays()`

UnsetMemoryArrays ensures that no value is present for MemoryArrays, not even an explicit nil
### GetPciDevices

`func (o *ComputeRackUnit) GetPciDevices() []PciDeviceRelationship`

GetPciDevices returns the PciDevices field if non-nil, zero value otherwise.

### GetPciDevicesOk

`func (o *ComputeRackUnit) GetPciDevicesOk() (*[]PciDeviceRelationship, bool)`

GetPciDevicesOk returns a tuple with the PciDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciDevices

`func (o *ComputeRackUnit) SetPciDevices(v []PciDeviceRelationship)`

SetPciDevices sets PciDevices field to given value.

### HasPciDevices

`func (o *ComputeRackUnit) HasPciDevices() bool`

HasPciDevices returns a boolean if a field has been set.

### SetPciDevicesNil

`func (o *ComputeRackUnit) SetPciDevicesNil(b bool)`

 SetPciDevicesNil sets the value for PciDevices to be an explicit nil

### UnsetPciDevices
`func (o *ComputeRackUnit) UnsetPciDevices()`

UnsetPciDevices ensures that no value is present for PciDevices, not even an explicit nil
### GetProcessors

`func (o *ComputeRackUnit) GetProcessors() []ProcessorUnitRelationship`

GetProcessors returns the Processors field if non-nil, zero value otherwise.

### GetProcessorsOk

`func (o *ComputeRackUnit) GetProcessorsOk() (*[]ProcessorUnitRelationship, bool)`

GetProcessorsOk returns a tuple with the Processors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessors

`func (o *ComputeRackUnit) SetProcessors(v []ProcessorUnitRelationship)`

SetProcessors sets Processors field to given value.

### HasProcessors

`func (o *ComputeRackUnit) HasProcessors() bool`

HasProcessors returns a boolean if a field has been set.

### SetProcessorsNil

`func (o *ComputeRackUnit) SetProcessorsNil(b bool)`

 SetProcessorsNil sets the value for Processors to be an explicit nil

### UnsetProcessors
`func (o *ComputeRackUnit) UnsetProcessors()`

UnsetProcessors ensures that no value is present for Processors, not even an explicit nil
### GetPsus

`func (o *ComputeRackUnit) GetPsus() []EquipmentPsuRelationship`

GetPsus returns the Psus field if non-nil, zero value otherwise.

### GetPsusOk

`func (o *ComputeRackUnit) GetPsusOk() (*[]EquipmentPsuRelationship, bool)`

GetPsusOk returns a tuple with the Psus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPsus

`func (o *ComputeRackUnit) SetPsus(v []EquipmentPsuRelationship)`

SetPsus sets Psus field to given value.

### HasPsus

`func (o *ComputeRackUnit) HasPsus() bool`

HasPsus returns a boolean if a field has been set.

### SetPsusNil

`func (o *ComputeRackUnit) SetPsusNil(b bool)`

 SetPsusNil sets the value for Psus to be an explicit nil

### UnsetPsus
`func (o *ComputeRackUnit) UnsetPsus()`

UnsetPsus ensures that no value is present for Psus, not even an explicit nil
### GetRackEnclosureSlot

`func (o *ComputeRackUnit) GetRackEnclosureSlot() EquipmentRackEnclosureSlotRelationship`

GetRackEnclosureSlot returns the RackEnclosureSlot field if non-nil, zero value otherwise.

### GetRackEnclosureSlotOk

`func (o *ComputeRackUnit) GetRackEnclosureSlotOk() (*EquipmentRackEnclosureSlotRelationship, bool)`

GetRackEnclosureSlotOk returns a tuple with the RackEnclosureSlot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRackEnclosureSlot

`func (o *ComputeRackUnit) SetRackEnclosureSlot(v EquipmentRackEnclosureSlotRelationship)`

SetRackEnclosureSlot sets RackEnclosureSlot field to given value.

### HasRackEnclosureSlot

`func (o *ComputeRackUnit) HasRackEnclosureSlot() bool`

HasRackEnclosureSlot returns a boolean if a field has been set.

### SetRackEnclosureSlotNil

`func (o *ComputeRackUnit) SetRackEnclosureSlotNil(b bool)`

 SetRackEnclosureSlotNil sets the value for RackEnclosureSlot to be an explicit nil

### UnsetRackEnclosureSlot
`func (o *ComputeRackUnit) UnsetRackEnclosureSlot()`

UnsetRackEnclosureSlot ensures that no value is present for RackEnclosureSlot, not even an explicit nil
### GetRegisteredDevice

`func (o *ComputeRackUnit) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *ComputeRackUnit) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *ComputeRackUnit) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *ComputeRackUnit) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *ComputeRackUnit) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *ComputeRackUnit) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
### GetSasExpanders

`func (o *ComputeRackUnit) GetSasExpanders() []StorageSasExpanderRelationship`

GetSasExpanders returns the SasExpanders field if non-nil, zero value otherwise.

### GetSasExpandersOk

`func (o *ComputeRackUnit) GetSasExpandersOk() (*[]StorageSasExpanderRelationship, bool)`

GetSasExpandersOk returns a tuple with the SasExpanders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSasExpanders

`func (o *ComputeRackUnit) SetSasExpanders(v []StorageSasExpanderRelationship)`

SetSasExpanders sets SasExpanders field to given value.

### HasSasExpanders

`func (o *ComputeRackUnit) HasSasExpanders() bool`

HasSasExpanders returns a boolean if a field has been set.

### SetSasExpandersNil

`func (o *ComputeRackUnit) SetSasExpandersNil(b bool)`

 SetSasExpandersNil sets the value for SasExpanders to be an explicit nil

### UnsetSasExpanders
`func (o *ComputeRackUnit) UnsetSasExpanders()`

UnsetSasExpanders ensures that no value is present for SasExpanders, not even an explicit nil
### GetStorageControllers

`func (o *ComputeRackUnit) GetStorageControllers() []StorageControllerRelationship`

GetStorageControllers returns the StorageControllers field if non-nil, zero value otherwise.

### GetStorageControllersOk

`func (o *ComputeRackUnit) GetStorageControllersOk() (*[]StorageControllerRelationship, bool)`

GetStorageControllersOk returns a tuple with the StorageControllers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageControllers

`func (o *ComputeRackUnit) SetStorageControllers(v []StorageControllerRelationship)`

SetStorageControllers sets StorageControllers field to given value.

### HasStorageControllers

`func (o *ComputeRackUnit) HasStorageControllers() bool`

HasStorageControllers returns a boolean if a field has been set.

### SetStorageControllersNil

`func (o *ComputeRackUnit) SetStorageControllersNil(b bool)`

 SetStorageControllersNil sets the value for StorageControllers to be an explicit nil

### UnsetStorageControllers
`func (o *ComputeRackUnit) UnsetStorageControllers()`

UnsetStorageControllers ensures that no value is present for StorageControllers, not even an explicit nil
### GetStorageEnclosures

`func (o *ComputeRackUnit) GetStorageEnclosures() []StorageEnclosureRelationship`

GetStorageEnclosures returns the StorageEnclosures field if non-nil, zero value otherwise.

### GetStorageEnclosuresOk

`func (o *ComputeRackUnit) GetStorageEnclosuresOk() (*[]StorageEnclosureRelationship, bool)`

GetStorageEnclosuresOk returns a tuple with the StorageEnclosures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageEnclosures

`func (o *ComputeRackUnit) SetStorageEnclosures(v []StorageEnclosureRelationship)`

SetStorageEnclosures sets StorageEnclosures field to given value.

### HasStorageEnclosures

`func (o *ComputeRackUnit) HasStorageEnclosures() bool`

HasStorageEnclosures returns a boolean if a field has been set.

### SetStorageEnclosuresNil

`func (o *ComputeRackUnit) SetStorageEnclosuresNil(b bool)`

 SetStorageEnclosuresNil sets the value for StorageEnclosures to be an explicit nil

### UnsetStorageEnclosures
`func (o *ComputeRackUnit) UnsetStorageEnclosures()`

UnsetStorageEnclosures ensures that no value is present for StorageEnclosures, not even an explicit nil
### GetTopSystem

`func (o *ComputeRackUnit) GetTopSystem() TopSystemRelationship`

GetTopSystem returns the TopSystem field if non-nil, zero value otherwise.

### GetTopSystemOk

`func (o *ComputeRackUnit) GetTopSystemOk() (*TopSystemRelationship, bool)`

GetTopSystemOk returns a tuple with the TopSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopSystem

`func (o *ComputeRackUnit) SetTopSystem(v TopSystemRelationship)`

SetTopSystem sets TopSystem field to given value.

### HasTopSystem

`func (o *ComputeRackUnit) HasTopSystem() bool`

HasTopSystem returns a boolean if a field has been set.

### SetTopSystemNil

`func (o *ComputeRackUnit) SetTopSystemNil(b bool)`

 SetTopSystemNil sets the value for TopSystem to be an explicit nil

### UnsetTopSystem
`func (o *ComputeRackUnit) UnsetTopSystem()`

UnsetTopSystem ensures that no value is present for TopSystem, not even an explicit nil
### GetUnitPersonality

`func (o *ComputeRackUnit) GetUnitPersonality() []RackUnitPersonalityRelationship`

GetUnitPersonality returns the UnitPersonality field if non-nil, zero value otherwise.

### GetUnitPersonalityOk

`func (o *ComputeRackUnit) GetUnitPersonalityOk() (*[]RackUnitPersonalityRelationship, bool)`

GetUnitPersonalityOk returns a tuple with the UnitPersonality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPersonality

`func (o *ComputeRackUnit) SetUnitPersonality(v []RackUnitPersonalityRelationship)`

SetUnitPersonality sets UnitPersonality field to given value.

### HasUnitPersonality

`func (o *ComputeRackUnit) HasUnitPersonality() bool`

HasUnitPersonality returns a boolean if a field has been set.

### SetUnitPersonalityNil

`func (o *ComputeRackUnit) SetUnitPersonalityNil(b bool)`

 SetUnitPersonalityNil sets the value for UnitPersonality to be an explicit nil

### UnsetUnitPersonality
`func (o *ComputeRackUnit) UnsetUnitPersonality()`

UnsetUnitPersonality ensures that no value is present for UnitPersonality, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



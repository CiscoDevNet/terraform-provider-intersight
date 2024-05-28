# ComputeBlade

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "compute.Blade"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "compute.Blade"]
**ChassisId** | Pointer to **string** | The id of the chassis that the blade is discovered in. | [optional] [readonly] 
**ScaledMode** | Pointer to **string** | The mode of the server that determines it is scaled. | [optional] [readonly] 
**SlotId** | Pointer to **int64** | The slot number in the chassis that the blade is discovered in. | [optional] [readonly] 
**Adapters** | Pointer to [**[]AdapterUnitRelationship**](AdapterUnitRelationship.md) | An array of relationships to adapterUnit resources. | [optional] [readonly] 
**BiosBootmode** | Pointer to [**NullableBiosBootModeRelationship**](BiosBootModeRelationship.md) |  | [optional] 
**BiosTokenSettings** | Pointer to [**NullableBiosTokenSettingsRelationship**](BiosTokenSettingsRelationship.md) |  | [optional] 
**BiosUnits** | Pointer to [**[]BiosUnitRelationship**](BiosUnitRelationship.md) | An array of relationships to biosUnit resources. | [optional] [readonly] 
**BiosVfSelectMemoryRasConfiguration** | Pointer to [**NullableBiosVfSelectMemoryRasConfigurationRelationship**](BiosVfSelectMemoryRasConfigurationRelationship.md) |  | [optional] 
**Bmc** | Pointer to [**NullableManagementControllerRelationship**](ManagementControllerRelationship.md) |  | [optional] 
**Board** | Pointer to [**NullableComputeBoardRelationship**](ComputeBoardRelationship.md) |  | [optional] 
**BootDeviceBootmode** | Pointer to [**NullableBootDeviceBootModeRelationship**](BootDeviceBootModeRelationship.md) |  | [optional] 
**ComputePersonality** | Pointer to [**[]ComputePersonalityRelationship**](ComputePersonalityRelationship.md) | An array of relationships to computePersonality resources. | [optional] [readonly] 
**EquipmentChassis** | Pointer to [**NullableEquipmentChassisRelationship**](EquipmentChassisRelationship.md) |  | [optional] 
**EquipmentIoExpanders** | Pointer to [**[]EquipmentIoExpanderRelationship**](EquipmentIoExpanderRelationship.md) | An array of relationships to equipmentIoExpander resources. | [optional] [readonly] 
**GenericInventoryHolders** | Pointer to [**[]InventoryGenericInventoryHolderRelationship**](InventoryGenericInventoryHolderRelationship.md) | An array of relationships to inventoryGenericInventoryHolder resources. | [optional] [readonly] 
**GraphicsCards** | Pointer to [**[]GraphicsCardRelationship**](GraphicsCardRelationship.md) | An array of relationships to graphicsCard resources. | [optional] 
**HybridDriveSlots** | Pointer to [**[]EquipmentHybridDriveSlotRelationship**](EquipmentHybridDriveSlotRelationship.md) | An array of relationships to equipmentHybridDriveSlot resources. | [optional] [readonly] 
**InventoryDeviceInfo** | Pointer to [**NullableInventoryDeviceInfoRelationship**](InventoryDeviceInfoRelationship.md) |  | [optional] 
**LocatorLed** | Pointer to [**NullableEquipmentLocatorLedRelationship**](EquipmentLocatorLedRelationship.md) |  | [optional] 
**MemoryArrays** | Pointer to [**[]MemoryArrayRelationship**](MemoryArrayRelationship.md) | An array of relationships to memoryArray resources. | [optional] 
**PciDevices** | Pointer to [**[]PciDeviceRelationship**](PciDeviceRelationship.md) | An array of relationships to pciDevice resources. | [optional] [readonly] 
**PciNodes** | Pointer to [**[]PciNodeRelationship**](PciNodeRelationship.md) | An array of relationships to pciNode resources. | [optional] [readonly] 
**Processors** | Pointer to [**[]ProcessorUnitRelationship**](ProcessorUnitRelationship.md) | An array of relationships to processorUnit resources. | [optional] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 
**StorageControllers** | Pointer to [**[]StorageControllerRelationship**](StorageControllerRelationship.md) | An array of relationships to storageController resources. | [optional] 
**StorageEnclosures** | Pointer to [**[]StorageEnclosureRelationship**](StorageEnclosureRelationship.md) | An array of relationships to storageEnclosure resources. | [optional] [readonly] 
**TopSystem** | Pointer to [**NullableTopSystemRelationship**](TopSystemRelationship.md) |  | [optional] 

## Methods

### NewComputeBlade

`func NewComputeBlade(classId string, objectType string, ) *ComputeBlade`

NewComputeBlade instantiates a new ComputeBlade object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputeBladeWithDefaults

`func NewComputeBladeWithDefaults() *ComputeBlade`

NewComputeBladeWithDefaults instantiates a new ComputeBlade object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ComputeBlade) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ComputeBlade) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ComputeBlade) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ComputeBlade) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ComputeBlade) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ComputeBlade) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetChassisId

`func (o *ComputeBlade) GetChassisId() string`

GetChassisId returns the ChassisId field if non-nil, zero value otherwise.

### GetChassisIdOk

`func (o *ComputeBlade) GetChassisIdOk() (*string, bool)`

GetChassisIdOk returns a tuple with the ChassisId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassisId

`func (o *ComputeBlade) SetChassisId(v string)`

SetChassisId sets ChassisId field to given value.

### HasChassisId

`func (o *ComputeBlade) HasChassisId() bool`

HasChassisId returns a boolean if a field has been set.

### GetScaledMode

`func (o *ComputeBlade) GetScaledMode() string`

GetScaledMode returns the ScaledMode field if non-nil, zero value otherwise.

### GetScaledModeOk

`func (o *ComputeBlade) GetScaledModeOk() (*string, bool)`

GetScaledModeOk returns a tuple with the ScaledMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaledMode

`func (o *ComputeBlade) SetScaledMode(v string)`

SetScaledMode sets ScaledMode field to given value.

### HasScaledMode

`func (o *ComputeBlade) HasScaledMode() bool`

HasScaledMode returns a boolean if a field has been set.

### GetSlotId

`func (o *ComputeBlade) GetSlotId() int64`

GetSlotId returns the SlotId field if non-nil, zero value otherwise.

### GetSlotIdOk

`func (o *ComputeBlade) GetSlotIdOk() (*int64, bool)`

GetSlotIdOk returns a tuple with the SlotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotId

`func (o *ComputeBlade) SetSlotId(v int64)`

SetSlotId sets SlotId field to given value.

### HasSlotId

`func (o *ComputeBlade) HasSlotId() bool`

HasSlotId returns a boolean if a field has been set.

### GetAdapters

`func (o *ComputeBlade) GetAdapters() []AdapterUnitRelationship`

GetAdapters returns the Adapters field if non-nil, zero value otherwise.

### GetAdaptersOk

`func (o *ComputeBlade) GetAdaptersOk() (*[]AdapterUnitRelationship, bool)`

GetAdaptersOk returns a tuple with the Adapters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapters

`func (o *ComputeBlade) SetAdapters(v []AdapterUnitRelationship)`

SetAdapters sets Adapters field to given value.

### HasAdapters

`func (o *ComputeBlade) HasAdapters() bool`

HasAdapters returns a boolean if a field has been set.

### SetAdaptersNil

`func (o *ComputeBlade) SetAdaptersNil(b bool)`

 SetAdaptersNil sets the value for Adapters to be an explicit nil

### UnsetAdapters
`func (o *ComputeBlade) UnsetAdapters()`

UnsetAdapters ensures that no value is present for Adapters, not even an explicit nil
### GetBiosBootmode

`func (o *ComputeBlade) GetBiosBootmode() BiosBootModeRelationship`

GetBiosBootmode returns the BiosBootmode field if non-nil, zero value otherwise.

### GetBiosBootmodeOk

`func (o *ComputeBlade) GetBiosBootmodeOk() (*BiosBootModeRelationship, bool)`

GetBiosBootmodeOk returns a tuple with the BiosBootmode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiosBootmode

`func (o *ComputeBlade) SetBiosBootmode(v BiosBootModeRelationship)`

SetBiosBootmode sets BiosBootmode field to given value.

### HasBiosBootmode

`func (o *ComputeBlade) HasBiosBootmode() bool`

HasBiosBootmode returns a boolean if a field has been set.

### SetBiosBootmodeNil

`func (o *ComputeBlade) SetBiosBootmodeNil(b bool)`

 SetBiosBootmodeNil sets the value for BiosBootmode to be an explicit nil

### UnsetBiosBootmode
`func (o *ComputeBlade) UnsetBiosBootmode()`

UnsetBiosBootmode ensures that no value is present for BiosBootmode, not even an explicit nil
### GetBiosTokenSettings

`func (o *ComputeBlade) GetBiosTokenSettings() BiosTokenSettingsRelationship`

GetBiosTokenSettings returns the BiosTokenSettings field if non-nil, zero value otherwise.

### GetBiosTokenSettingsOk

`func (o *ComputeBlade) GetBiosTokenSettingsOk() (*BiosTokenSettingsRelationship, bool)`

GetBiosTokenSettingsOk returns a tuple with the BiosTokenSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiosTokenSettings

`func (o *ComputeBlade) SetBiosTokenSettings(v BiosTokenSettingsRelationship)`

SetBiosTokenSettings sets BiosTokenSettings field to given value.

### HasBiosTokenSettings

`func (o *ComputeBlade) HasBiosTokenSettings() bool`

HasBiosTokenSettings returns a boolean if a field has been set.

### SetBiosTokenSettingsNil

`func (o *ComputeBlade) SetBiosTokenSettingsNil(b bool)`

 SetBiosTokenSettingsNil sets the value for BiosTokenSettings to be an explicit nil

### UnsetBiosTokenSettings
`func (o *ComputeBlade) UnsetBiosTokenSettings()`

UnsetBiosTokenSettings ensures that no value is present for BiosTokenSettings, not even an explicit nil
### GetBiosUnits

`func (o *ComputeBlade) GetBiosUnits() []BiosUnitRelationship`

GetBiosUnits returns the BiosUnits field if non-nil, zero value otherwise.

### GetBiosUnitsOk

`func (o *ComputeBlade) GetBiosUnitsOk() (*[]BiosUnitRelationship, bool)`

GetBiosUnitsOk returns a tuple with the BiosUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiosUnits

`func (o *ComputeBlade) SetBiosUnits(v []BiosUnitRelationship)`

SetBiosUnits sets BiosUnits field to given value.

### HasBiosUnits

`func (o *ComputeBlade) HasBiosUnits() bool`

HasBiosUnits returns a boolean if a field has been set.

### SetBiosUnitsNil

`func (o *ComputeBlade) SetBiosUnitsNil(b bool)`

 SetBiosUnitsNil sets the value for BiosUnits to be an explicit nil

### UnsetBiosUnits
`func (o *ComputeBlade) UnsetBiosUnits()`

UnsetBiosUnits ensures that no value is present for BiosUnits, not even an explicit nil
### GetBiosVfSelectMemoryRasConfiguration

`func (o *ComputeBlade) GetBiosVfSelectMemoryRasConfiguration() BiosVfSelectMemoryRasConfigurationRelationship`

GetBiosVfSelectMemoryRasConfiguration returns the BiosVfSelectMemoryRasConfiguration field if non-nil, zero value otherwise.

### GetBiosVfSelectMemoryRasConfigurationOk

`func (o *ComputeBlade) GetBiosVfSelectMemoryRasConfigurationOk() (*BiosVfSelectMemoryRasConfigurationRelationship, bool)`

GetBiosVfSelectMemoryRasConfigurationOk returns a tuple with the BiosVfSelectMemoryRasConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiosVfSelectMemoryRasConfiguration

`func (o *ComputeBlade) SetBiosVfSelectMemoryRasConfiguration(v BiosVfSelectMemoryRasConfigurationRelationship)`

SetBiosVfSelectMemoryRasConfiguration sets BiosVfSelectMemoryRasConfiguration field to given value.

### HasBiosVfSelectMemoryRasConfiguration

`func (o *ComputeBlade) HasBiosVfSelectMemoryRasConfiguration() bool`

HasBiosVfSelectMemoryRasConfiguration returns a boolean if a field has been set.

### SetBiosVfSelectMemoryRasConfigurationNil

`func (o *ComputeBlade) SetBiosVfSelectMemoryRasConfigurationNil(b bool)`

 SetBiosVfSelectMemoryRasConfigurationNil sets the value for BiosVfSelectMemoryRasConfiguration to be an explicit nil

### UnsetBiosVfSelectMemoryRasConfiguration
`func (o *ComputeBlade) UnsetBiosVfSelectMemoryRasConfiguration()`

UnsetBiosVfSelectMemoryRasConfiguration ensures that no value is present for BiosVfSelectMemoryRasConfiguration, not even an explicit nil
### GetBmc

`func (o *ComputeBlade) GetBmc() ManagementControllerRelationship`

GetBmc returns the Bmc field if non-nil, zero value otherwise.

### GetBmcOk

`func (o *ComputeBlade) GetBmcOk() (*ManagementControllerRelationship, bool)`

GetBmcOk returns a tuple with the Bmc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBmc

`func (o *ComputeBlade) SetBmc(v ManagementControllerRelationship)`

SetBmc sets Bmc field to given value.

### HasBmc

`func (o *ComputeBlade) HasBmc() bool`

HasBmc returns a boolean if a field has been set.

### SetBmcNil

`func (o *ComputeBlade) SetBmcNil(b bool)`

 SetBmcNil sets the value for Bmc to be an explicit nil

### UnsetBmc
`func (o *ComputeBlade) UnsetBmc()`

UnsetBmc ensures that no value is present for Bmc, not even an explicit nil
### GetBoard

`func (o *ComputeBlade) GetBoard() ComputeBoardRelationship`

GetBoard returns the Board field if non-nil, zero value otherwise.

### GetBoardOk

`func (o *ComputeBlade) GetBoardOk() (*ComputeBoardRelationship, bool)`

GetBoardOk returns a tuple with the Board field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoard

`func (o *ComputeBlade) SetBoard(v ComputeBoardRelationship)`

SetBoard sets Board field to given value.

### HasBoard

`func (o *ComputeBlade) HasBoard() bool`

HasBoard returns a boolean if a field has been set.

### SetBoardNil

`func (o *ComputeBlade) SetBoardNil(b bool)`

 SetBoardNil sets the value for Board to be an explicit nil

### UnsetBoard
`func (o *ComputeBlade) UnsetBoard()`

UnsetBoard ensures that no value is present for Board, not even an explicit nil
### GetBootDeviceBootmode

`func (o *ComputeBlade) GetBootDeviceBootmode() BootDeviceBootModeRelationship`

GetBootDeviceBootmode returns the BootDeviceBootmode field if non-nil, zero value otherwise.

### GetBootDeviceBootmodeOk

`func (o *ComputeBlade) GetBootDeviceBootmodeOk() (*BootDeviceBootModeRelationship, bool)`

GetBootDeviceBootmodeOk returns a tuple with the BootDeviceBootmode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootDeviceBootmode

`func (o *ComputeBlade) SetBootDeviceBootmode(v BootDeviceBootModeRelationship)`

SetBootDeviceBootmode sets BootDeviceBootmode field to given value.

### HasBootDeviceBootmode

`func (o *ComputeBlade) HasBootDeviceBootmode() bool`

HasBootDeviceBootmode returns a boolean if a field has been set.

### SetBootDeviceBootmodeNil

`func (o *ComputeBlade) SetBootDeviceBootmodeNil(b bool)`

 SetBootDeviceBootmodeNil sets the value for BootDeviceBootmode to be an explicit nil

### UnsetBootDeviceBootmode
`func (o *ComputeBlade) UnsetBootDeviceBootmode()`

UnsetBootDeviceBootmode ensures that no value is present for BootDeviceBootmode, not even an explicit nil
### GetComputePersonality

`func (o *ComputeBlade) GetComputePersonality() []ComputePersonalityRelationship`

GetComputePersonality returns the ComputePersonality field if non-nil, zero value otherwise.

### GetComputePersonalityOk

`func (o *ComputeBlade) GetComputePersonalityOk() (*[]ComputePersonalityRelationship, bool)`

GetComputePersonalityOk returns a tuple with the ComputePersonality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputePersonality

`func (o *ComputeBlade) SetComputePersonality(v []ComputePersonalityRelationship)`

SetComputePersonality sets ComputePersonality field to given value.

### HasComputePersonality

`func (o *ComputeBlade) HasComputePersonality() bool`

HasComputePersonality returns a boolean if a field has been set.

### SetComputePersonalityNil

`func (o *ComputeBlade) SetComputePersonalityNil(b bool)`

 SetComputePersonalityNil sets the value for ComputePersonality to be an explicit nil

### UnsetComputePersonality
`func (o *ComputeBlade) UnsetComputePersonality()`

UnsetComputePersonality ensures that no value is present for ComputePersonality, not even an explicit nil
### GetEquipmentChassis

`func (o *ComputeBlade) GetEquipmentChassis() EquipmentChassisRelationship`

GetEquipmentChassis returns the EquipmentChassis field if non-nil, zero value otherwise.

### GetEquipmentChassisOk

`func (o *ComputeBlade) GetEquipmentChassisOk() (*EquipmentChassisRelationship, bool)`

GetEquipmentChassisOk returns a tuple with the EquipmentChassis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentChassis

`func (o *ComputeBlade) SetEquipmentChassis(v EquipmentChassisRelationship)`

SetEquipmentChassis sets EquipmentChassis field to given value.

### HasEquipmentChassis

`func (o *ComputeBlade) HasEquipmentChassis() bool`

HasEquipmentChassis returns a boolean if a field has been set.

### SetEquipmentChassisNil

`func (o *ComputeBlade) SetEquipmentChassisNil(b bool)`

 SetEquipmentChassisNil sets the value for EquipmentChassis to be an explicit nil

### UnsetEquipmentChassis
`func (o *ComputeBlade) UnsetEquipmentChassis()`

UnsetEquipmentChassis ensures that no value is present for EquipmentChassis, not even an explicit nil
### GetEquipmentIoExpanders

`func (o *ComputeBlade) GetEquipmentIoExpanders() []EquipmentIoExpanderRelationship`

GetEquipmentIoExpanders returns the EquipmentIoExpanders field if non-nil, zero value otherwise.

### GetEquipmentIoExpandersOk

`func (o *ComputeBlade) GetEquipmentIoExpandersOk() (*[]EquipmentIoExpanderRelationship, bool)`

GetEquipmentIoExpandersOk returns a tuple with the EquipmentIoExpanders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentIoExpanders

`func (o *ComputeBlade) SetEquipmentIoExpanders(v []EquipmentIoExpanderRelationship)`

SetEquipmentIoExpanders sets EquipmentIoExpanders field to given value.

### HasEquipmentIoExpanders

`func (o *ComputeBlade) HasEquipmentIoExpanders() bool`

HasEquipmentIoExpanders returns a boolean if a field has been set.

### SetEquipmentIoExpandersNil

`func (o *ComputeBlade) SetEquipmentIoExpandersNil(b bool)`

 SetEquipmentIoExpandersNil sets the value for EquipmentIoExpanders to be an explicit nil

### UnsetEquipmentIoExpanders
`func (o *ComputeBlade) UnsetEquipmentIoExpanders()`

UnsetEquipmentIoExpanders ensures that no value is present for EquipmentIoExpanders, not even an explicit nil
### GetGenericInventoryHolders

`func (o *ComputeBlade) GetGenericInventoryHolders() []InventoryGenericInventoryHolderRelationship`

GetGenericInventoryHolders returns the GenericInventoryHolders field if non-nil, zero value otherwise.

### GetGenericInventoryHoldersOk

`func (o *ComputeBlade) GetGenericInventoryHoldersOk() (*[]InventoryGenericInventoryHolderRelationship, bool)`

GetGenericInventoryHoldersOk returns a tuple with the GenericInventoryHolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenericInventoryHolders

`func (o *ComputeBlade) SetGenericInventoryHolders(v []InventoryGenericInventoryHolderRelationship)`

SetGenericInventoryHolders sets GenericInventoryHolders field to given value.

### HasGenericInventoryHolders

`func (o *ComputeBlade) HasGenericInventoryHolders() bool`

HasGenericInventoryHolders returns a boolean if a field has been set.

### SetGenericInventoryHoldersNil

`func (o *ComputeBlade) SetGenericInventoryHoldersNil(b bool)`

 SetGenericInventoryHoldersNil sets the value for GenericInventoryHolders to be an explicit nil

### UnsetGenericInventoryHolders
`func (o *ComputeBlade) UnsetGenericInventoryHolders()`

UnsetGenericInventoryHolders ensures that no value is present for GenericInventoryHolders, not even an explicit nil
### GetGraphicsCards

`func (o *ComputeBlade) GetGraphicsCards() []GraphicsCardRelationship`

GetGraphicsCards returns the GraphicsCards field if non-nil, zero value otherwise.

### GetGraphicsCardsOk

`func (o *ComputeBlade) GetGraphicsCardsOk() (*[]GraphicsCardRelationship, bool)`

GetGraphicsCardsOk returns a tuple with the GraphicsCards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraphicsCards

`func (o *ComputeBlade) SetGraphicsCards(v []GraphicsCardRelationship)`

SetGraphicsCards sets GraphicsCards field to given value.

### HasGraphicsCards

`func (o *ComputeBlade) HasGraphicsCards() bool`

HasGraphicsCards returns a boolean if a field has been set.

### SetGraphicsCardsNil

`func (o *ComputeBlade) SetGraphicsCardsNil(b bool)`

 SetGraphicsCardsNil sets the value for GraphicsCards to be an explicit nil

### UnsetGraphicsCards
`func (o *ComputeBlade) UnsetGraphicsCards()`

UnsetGraphicsCards ensures that no value is present for GraphicsCards, not even an explicit nil
### GetHybridDriveSlots

`func (o *ComputeBlade) GetHybridDriveSlots() []EquipmentHybridDriveSlotRelationship`

GetHybridDriveSlots returns the HybridDriveSlots field if non-nil, zero value otherwise.

### GetHybridDriveSlotsOk

`func (o *ComputeBlade) GetHybridDriveSlotsOk() (*[]EquipmentHybridDriveSlotRelationship, bool)`

GetHybridDriveSlotsOk returns a tuple with the HybridDriveSlots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHybridDriveSlots

`func (o *ComputeBlade) SetHybridDriveSlots(v []EquipmentHybridDriveSlotRelationship)`

SetHybridDriveSlots sets HybridDriveSlots field to given value.

### HasHybridDriveSlots

`func (o *ComputeBlade) HasHybridDriveSlots() bool`

HasHybridDriveSlots returns a boolean if a field has been set.

### SetHybridDriveSlotsNil

`func (o *ComputeBlade) SetHybridDriveSlotsNil(b bool)`

 SetHybridDriveSlotsNil sets the value for HybridDriveSlots to be an explicit nil

### UnsetHybridDriveSlots
`func (o *ComputeBlade) UnsetHybridDriveSlots()`

UnsetHybridDriveSlots ensures that no value is present for HybridDriveSlots, not even an explicit nil
### GetInventoryDeviceInfo

`func (o *ComputeBlade) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *ComputeBlade) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *ComputeBlade) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *ComputeBlade) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### SetInventoryDeviceInfoNil

`func (o *ComputeBlade) SetInventoryDeviceInfoNil(b bool)`

 SetInventoryDeviceInfoNil sets the value for InventoryDeviceInfo to be an explicit nil

### UnsetInventoryDeviceInfo
`func (o *ComputeBlade) UnsetInventoryDeviceInfo()`

UnsetInventoryDeviceInfo ensures that no value is present for InventoryDeviceInfo, not even an explicit nil
### GetLocatorLed

`func (o *ComputeBlade) GetLocatorLed() EquipmentLocatorLedRelationship`

GetLocatorLed returns the LocatorLed field if non-nil, zero value otherwise.

### GetLocatorLedOk

`func (o *ComputeBlade) GetLocatorLedOk() (*EquipmentLocatorLedRelationship, bool)`

GetLocatorLedOk returns a tuple with the LocatorLed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocatorLed

`func (o *ComputeBlade) SetLocatorLed(v EquipmentLocatorLedRelationship)`

SetLocatorLed sets LocatorLed field to given value.

### HasLocatorLed

`func (o *ComputeBlade) HasLocatorLed() bool`

HasLocatorLed returns a boolean if a field has been set.

### SetLocatorLedNil

`func (o *ComputeBlade) SetLocatorLedNil(b bool)`

 SetLocatorLedNil sets the value for LocatorLed to be an explicit nil

### UnsetLocatorLed
`func (o *ComputeBlade) UnsetLocatorLed()`

UnsetLocatorLed ensures that no value is present for LocatorLed, not even an explicit nil
### GetMemoryArrays

`func (o *ComputeBlade) GetMemoryArrays() []MemoryArrayRelationship`

GetMemoryArrays returns the MemoryArrays field if non-nil, zero value otherwise.

### GetMemoryArraysOk

`func (o *ComputeBlade) GetMemoryArraysOk() (*[]MemoryArrayRelationship, bool)`

GetMemoryArraysOk returns a tuple with the MemoryArrays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryArrays

`func (o *ComputeBlade) SetMemoryArrays(v []MemoryArrayRelationship)`

SetMemoryArrays sets MemoryArrays field to given value.

### HasMemoryArrays

`func (o *ComputeBlade) HasMemoryArrays() bool`

HasMemoryArrays returns a boolean if a field has been set.

### SetMemoryArraysNil

`func (o *ComputeBlade) SetMemoryArraysNil(b bool)`

 SetMemoryArraysNil sets the value for MemoryArrays to be an explicit nil

### UnsetMemoryArrays
`func (o *ComputeBlade) UnsetMemoryArrays()`

UnsetMemoryArrays ensures that no value is present for MemoryArrays, not even an explicit nil
### GetPciDevices

`func (o *ComputeBlade) GetPciDevices() []PciDeviceRelationship`

GetPciDevices returns the PciDevices field if non-nil, zero value otherwise.

### GetPciDevicesOk

`func (o *ComputeBlade) GetPciDevicesOk() (*[]PciDeviceRelationship, bool)`

GetPciDevicesOk returns a tuple with the PciDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciDevices

`func (o *ComputeBlade) SetPciDevices(v []PciDeviceRelationship)`

SetPciDevices sets PciDevices field to given value.

### HasPciDevices

`func (o *ComputeBlade) HasPciDevices() bool`

HasPciDevices returns a boolean if a field has been set.

### SetPciDevicesNil

`func (o *ComputeBlade) SetPciDevicesNil(b bool)`

 SetPciDevicesNil sets the value for PciDevices to be an explicit nil

### UnsetPciDevices
`func (o *ComputeBlade) UnsetPciDevices()`

UnsetPciDevices ensures that no value is present for PciDevices, not even an explicit nil
### GetPciNodes

`func (o *ComputeBlade) GetPciNodes() []PciNodeRelationship`

GetPciNodes returns the PciNodes field if non-nil, zero value otherwise.

### GetPciNodesOk

`func (o *ComputeBlade) GetPciNodesOk() (*[]PciNodeRelationship, bool)`

GetPciNodesOk returns a tuple with the PciNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciNodes

`func (o *ComputeBlade) SetPciNodes(v []PciNodeRelationship)`

SetPciNodes sets PciNodes field to given value.

### HasPciNodes

`func (o *ComputeBlade) HasPciNodes() bool`

HasPciNodes returns a boolean if a field has been set.

### SetPciNodesNil

`func (o *ComputeBlade) SetPciNodesNil(b bool)`

 SetPciNodesNil sets the value for PciNodes to be an explicit nil

### UnsetPciNodes
`func (o *ComputeBlade) UnsetPciNodes()`

UnsetPciNodes ensures that no value is present for PciNodes, not even an explicit nil
### GetProcessors

`func (o *ComputeBlade) GetProcessors() []ProcessorUnitRelationship`

GetProcessors returns the Processors field if non-nil, zero value otherwise.

### GetProcessorsOk

`func (o *ComputeBlade) GetProcessorsOk() (*[]ProcessorUnitRelationship, bool)`

GetProcessorsOk returns a tuple with the Processors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessors

`func (o *ComputeBlade) SetProcessors(v []ProcessorUnitRelationship)`

SetProcessors sets Processors field to given value.

### HasProcessors

`func (o *ComputeBlade) HasProcessors() bool`

HasProcessors returns a boolean if a field has been set.

### SetProcessorsNil

`func (o *ComputeBlade) SetProcessorsNil(b bool)`

 SetProcessorsNil sets the value for Processors to be an explicit nil

### UnsetProcessors
`func (o *ComputeBlade) UnsetProcessors()`

UnsetProcessors ensures that no value is present for Processors, not even an explicit nil
### GetRegisteredDevice

`func (o *ComputeBlade) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *ComputeBlade) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *ComputeBlade) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *ComputeBlade) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *ComputeBlade) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *ComputeBlade) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
### GetStorageControllers

`func (o *ComputeBlade) GetStorageControllers() []StorageControllerRelationship`

GetStorageControllers returns the StorageControllers field if non-nil, zero value otherwise.

### GetStorageControllersOk

`func (o *ComputeBlade) GetStorageControllersOk() (*[]StorageControllerRelationship, bool)`

GetStorageControllersOk returns a tuple with the StorageControllers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageControllers

`func (o *ComputeBlade) SetStorageControllers(v []StorageControllerRelationship)`

SetStorageControllers sets StorageControllers field to given value.

### HasStorageControllers

`func (o *ComputeBlade) HasStorageControllers() bool`

HasStorageControllers returns a boolean if a field has been set.

### SetStorageControllersNil

`func (o *ComputeBlade) SetStorageControllersNil(b bool)`

 SetStorageControllersNil sets the value for StorageControllers to be an explicit nil

### UnsetStorageControllers
`func (o *ComputeBlade) UnsetStorageControllers()`

UnsetStorageControllers ensures that no value is present for StorageControllers, not even an explicit nil
### GetStorageEnclosures

`func (o *ComputeBlade) GetStorageEnclosures() []StorageEnclosureRelationship`

GetStorageEnclosures returns the StorageEnclosures field if non-nil, zero value otherwise.

### GetStorageEnclosuresOk

`func (o *ComputeBlade) GetStorageEnclosuresOk() (*[]StorageEnclosureRelationship, bool)`

GetStorageEnclosuresOk returns a tuple with the StorageEnclosures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageEnclosures

`func (o *ComputeBlade) SetStorageEnclosures(v []StorageEnclosureRelationship)`

SetStorageEnclosures sets StorageEnclosures field to given value.

### HasStorageEnclosures

`func (o *ComputeBlade) HasStorageEnclosures() bool`

HasStorageEnclosures returns a boolean if a field has been set.

### SetStorageEnclosuresNil

`func (o *ComputeBlade) SetStorageEnclosuresNil(b bool)`

 SetStorageEnclosuresNil sets the value for StorageEnclosures to be an explicit nil

### UnsetStorageEnclosures
`func (o *ComputeBlade) UnsetStorageEnclosures()`

UnsetStorageEnclosures ensures that no value is present for StorageEnclosures, not even an explicit nil
### GetTopSystem

`func (o *ComputeBlade) GetTopSystem() TopSystemRelationship`

GetTopSystem returns the TopSystem field if non-nil, zero value otherwise.

### GetTopSystemOk

`func (o *ComputeBlade) GetTopSystemOk() (*TopSystemRelationship, bool)`

GetTopSystemOk returns a tuple with the TopSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopSystem

`func (o *ComputeBlade) SetTopSystem(v TopSystemRelationship)`

SetTopSystem sets TopSystem field to given value.

### HasTopSystem

`func (o *ComputeBlade) HasTopSystem() bool`

HasTopSystem returns a boolean if a field has been set.

### SetTopSystemNil

`func (o *ComputeBlade) SetTopSystemNil(b bool)`

 SetTopSystemNil sets the value for TopSystem to be an explicit nil

### UnsetTopSystem
`func (o *ComputeBlade) UnsetTopSystem()`

UnsetTopSystem ensures that no value is present for TopSystem, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



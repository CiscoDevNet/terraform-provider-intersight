# FirmwareRunningFirmware

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "firmware.RunningFirmware"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "firmware.RunningFirmware"]
**Component** | Pointer to **string** | Kind of the firmware - boot-booloader/system/kernel. | [optional] [readonly] 
**PackageVersion** | Pointer to **string** | Bundle version which the firmware belongs to. | [optional] [readonly] 
**Type** | Pointer to **string** | The type of the firmware. | [optional] [readonly] 
**Version** | Pointer to **string** | The version of the firmware. | [optional] [readonly] 
**BiosUnit** | Pointer to [**NullableBiosUnitRelationship**](BiosUnitRelationship.md) |  | [optional] 
**EquipmentEnclosureElement** | Pointer to [**NullableEquipmentEnclosureElementRelationship**](EquipmentEnclosureElementRelationship.md) |  | [optional] 
**GraphicsCard** | Pointer to [**NullableGraphicsCardRelationship**](GraphicsCardRelationship.md) |  | [optional] 
**InventoryDeviceInfo** | Pointer to [**NullableInventoryDeviceInfoRelationship**](InventoryDeviceInfoRelationship.md) |  | [optional] 
**ManagementController** | Pointer to [**NullableManagementControllerRelationship**](ManagementControllerRelationship.md) |  | [optional] 
**NetworkElements** | Pointer to [**[]NetworkElementRelationship**](NetworkElementRelationship.md) | An array of relationships to networkElement resources. | [optional] 
**PciSwitch** | Pointer to [**NullablePciSwitchRelationship**](PciSwitchRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 
**StorageController** | Pointer to [**NullableStorageControllerRelationship**](StorageControllerRelationship.md) |  | [optional] 
**StorageFlexFlashController** | Pointer to [**NullableStorageFlexFlashControllerRelationship**](StorageFlexFlashControllerRelationship.md) |  | [optional] 
**StoragePhysicalDisk** | Pointer to [**NullableStoragePhysicalDiskRelationship**](StoragePhysicalDiskRelationship.md) |  | [optional] 

## Methods

### NewFirmwareRunningFirmware

`func NewFirmwareRunningFirmware(classId string, objectType string, ) *FirmwareRunningFirmware`

NewFirmwareRunningFirmware instantiates a new FirmwareRunningFirmware object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirmwareRunningFirmwareWithDefaults

`func NewFirmwareRunningFirmwareWithDefaults() *FirmwareRunningFirmware`

NewFirmwareRunningFirmwareWithDefaults instantiates a new FirmwareRunningFirmware object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FirmwareRunningFirmware) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FirmwareRunningFirmware) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FirmwareRunningFirmware) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FirmwareRunningFirmware) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FirmwareRunningFirmware) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FirmwareRunningFirmware) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetComponent

`func (o *FirmwareRunningFirmware) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *FirmwareRunningFirmware) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *FirmwareRunningFirmware) SetComponent(v string)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *FirmwareRunningFirmware) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetPackageVersion

`func (o *FirmwareRunningFirmware) GetPackageVersion() string`

GetPackageVersion returns the PackageVersion field if non-nil, zero value otherwise.

### GetPackageVersionOk

`func (o *FirmwareRunningFirmware) GetPackageVersionOk() (*string, bool)`

GetPackageVersionOk returns a tuple with the PackageVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageVersion

`func (o *FirmwareRunningFirmware) SetPackageVersion(v string)`

SetPackageVersion sets PackageVersion field to given value.

### HasPackageVersion

`func (o *FirmwareRunningFirmware) HasPackageVersion() bool`

HasPackageVersion returns a boolean if a field has been set.

### GetType

`func (o *FirmwareRunningFirmware) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FirmwareRunningFirmware) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FirmwareRunningFirmware) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FirmwareRunningFirmware) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVersion

`func (o *FirmwareRunningFirmware) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *FirmwareRunningFirmware) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *FirmwareRunningFirmware) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *FirmwareRunningFirmware) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetBiosUnit

`func (o *FirmwareRunningFirmware) GetBiosUnit() BiosUnitRelationship`

GetBiosUnit returns the BiosUnit field if non-nil, zero value otherwise.

### GetBiosUnitOk

`func (o *FirmwareRunningFirmware) GetBiosUnitOk() (*BiosUnitRelationship, bool)`

GetBiosUnitOk returns a tuple with the BiosUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiosUnit

`func (o *FirmwareRunningFirmware) SetBiosUnit(v BiosUnitRelationship)`

SetBiosUnit sets BiosUnit field to given value.

### HasBiosUnit

`func (o *FirmwareRunningFirmware) HasBiosUnit() bool`

HasBiosUnit returns a boolean if a field has been set.

### SetBiosUnitNil

`func (o *FirmwareRunningFirmware) SetBiosUnitNil(b bool)`

 SetBiosUnitNil sets the value for BiosUnit to be an explicit nil

### UnsetBiosUnit
`func (o *FirmwareRunningFirmware) UnsetBiosUnit()`

UnsetBiosUnit ensures that no value is present for BiosUnit, not even an explicit nil
### GetEquipmentEnclosureElement

`func (o *FirmwareRunningFirmware) GetEquipmentEnclosureElement() EquipmentEnclosureElementRelationship`

GetEquipmentEnclosureElement returns the EquipmentEnclosureElement field if non-nil, zero value otherwise.

### GetEquipmentEnclosureElementOk

`func (o *FirmwareRunningFirmware) GetEquipmentEnclosureElementOk() (*EquipmentEnclosureElementRelationship, bool)`

GetEquipmentEnclosureElementOk returns a tuple with the EquipmentEnclosureElement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentEnclosureElement

`func (o *FirmwareRunningFirmware) SetEquipmentEnclosureElement(v EquipmentEnclosureElementRelationship)`

SetEquipmentEnclosureElement sets EquipmentEnclosureElement field to given value.

### HasEquipmentEnclosureElement

`func (o *FirmwareRunningFirmware) HasEquipmentEnclosureElement() bool`

HasEquipmentEnclosureElement returns a boolean if a field has been set.

### SetEquipmentEnclosureElementNil

`func (o *FirmwareRunningFirmware) SetEquipmentEnclosureElementNil(b bool)`

 SetEquipmentEnclosureElementNil sets the value for EquipmentEnclosureElement to be an explicit nil

### UnsetEquipmentEnclosureElement
`func (o *FirmwareRunningFirmware) UnsetEquipmentEnclosureElement()`

UnsetEquipmentEnclosureElement ensures that no value is present for EquipmentEnclosureElement, not even an explicit nil
### GetGraphicsCard

`func (o *FirmwareRunningFirmware) GetGraphicsCard() GraphicsCardRelationship`

GetGraphicsCard returns the GraphicsCard field if non-nil, zero value otherwise.

### GetGraphicsCardOk

`func (o *FirmwareRunningFirmware) GetGraphicsCardOk() (*GraphicsCardRelationship, bool)`

GetGraphicsCardOk returns a tuple with the GraphicsCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraphicsCard

`func (o *FirmwareRunningFirmware) SetGraphicsCard(v GraphicsCardRelationship)`

SetGraphicsCard sets GraphicsCard field to given value.

### HasGraphicsCard

`func (o *FirmwareRunningFirmware) HasGraphicsCard() bool`

HasGraphicsCard returns a boolean if a field has been set.

### SetGraphicsCardNil

`func (o *FirmwareRunningFirmware) SetGraphicsCardNil(b bool)`

 SetGraphicsCardNil sets the value for GraphicsCard to be an explicit nil

### UnsetGraphicsCard
`func (o *FirmwareRunningFirmware) UnsetGraphicsCard()`

UnsetGraphicsCard ensures that no value is present for GraphicsCard, not even an explicit nil
### GetInventoryDeviceInfo

`func (o *FirmwareRunningFirmware) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *FirmwareRunningFirmware) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *FirmwareRunningFirmware) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *FirmwareRunningFirmware) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### SetInventoryDeviceInfoNil

`func (o *FirmwareRunningFirmware) SetInventoryDeviceInfoNil(b bool)`

 SetInventoryDeviceInfoNil sets the value for InventoryDeviceInfo to be an explicit nil

### UnsetInventoryDeviceInfo
`func (o *FirmwareRunningFirmware) UnsetInventoryDeviceInfo()`

UnsetInventoryDeviceInfo ensures that no value is present for InventoryDeviceInfo, not even an explicit nil
### GetManagementController

`func (o *FirmwareRunningFirmware) GetManagementController() ManagementControllerRelationship`

GetManagementController returns the ManagementController field if non-nil, zero value otherwise.

### GetManagementControllerOk

`func (o *FirmwareRunningFirmware) GetManagementControllerOk() (*ManagementControllerRelationship, bool)`

GetManagementControllerOk returns a tuple with the ManagementController field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementController

`func (o *FirmwareRunningFirmware) SetManagementController(v ManagementControllerRelationship)`

SetManagementController sets ManagementController field to given value.

### HasManagementController

`func (o *FirmwareRunningFirmware) HasManagementController() bool`

HasManagementController returns a boolean if a field has been set.

### SetManagementControllerNil

`func (o *FirmwareRunningFirmware) SetManagementControllerNil(b bool)`

 SetManagementControllerNil sets the value for ManagementController to be an explicit nil

### UnsetManagementController
`func (o *FirmwareRunningFirmware) UnsetManagementController()`

UnsetManagementController ensures that no value is present for ManagementController, not even an explicit nil
### GetNetworkElements

`func (o *FirmwareRunningFirmware) GetNetworkElements() []NetworkElementRelationship`

GetNetworkElements returns the NetworkElements field if non-nil, zero value otherwise.

### GetNetworkElementsOk

`func (o *FirmwareRunningFirmware) GetNetworkElementsOk() (*[]NetworkElementRelationship, bool)`

GetNetworkElementsOk returns a tuple with the NetworkElements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkElements

`func (o *FirmwareRunningFirmware) SetNetworkElements(v []NetworkElementRelationship)`

SetNetworkElements sets NetworkElements field to given value.

### HasNetworkElements

`func (o *FirmwareRunningFirmware) HasNetworkElements() bool`

HasNetworkElements returns a boolean if a field has been set.

### SetNetworkElementsNil

`func (o *FirmwareRunningFirmware) SetNetworkElementsNil(b bool)`

 SetNetworkElementsNil sets the value for NetworkElements to be an explicit nil

### UnsetNetworkElements
`func (o *FirmwareRunningFirmware) UnsetNetworkElements()`

UnsetNetworkElements ensures that no value is present for NetworkElements, not even an explicit nil
### GetPciSwitch

`func (o *FirmwareRunningFirmware) GetPciSwitch() PciSwitchRelationship`

GetPciSwitch returns the PciSwitch field if non-nil, zero value otherwise.

### GetPciSwitchOk

`func (o *FirmwareRunningFirmware) GetPciSwitchOk() (*PciSwitchRelationship, bool)`

GetPciSwitchOk returns a tuple with the PciSwitch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciSwitch

`func (o *FirmwareRunningFirmware) SetPciSwitch(v PciSwitchRelationship)`

SetPciSwitch sets PciSwitch field to given value.

### HasPciSwitch

`func (o *FirmwareRunningFirmware) HasPciSwitch() bool`

HasPciSwitch returns a boolean if a field has been set.

### SetPciSwitchNil

`func (o *FirmwareRunningFirmware) SetPciSwitchNil(b bool)`

 SetPciSwitchNil sets the value for PciSwitch to be an explicit nil

### UnsetPciSwitch
`func (o *FirmwareRunningFirmware) UnsetPciSwitch()`

UnsetPciSwitch ensures that no value is present for PciSwitch, not even an explicit nil
### GetRegisteredDevice

`func (o *FirmwareRunningFirmware) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *FirmwareRunningFirmware) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *FirmwareRunningFirmware) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *FirmwareRunningFirmware) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *FirmwareRunningFirmware) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *FirmwareRunningFirmware) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
### GetStorageController

`func (o *FirmwareRunningFirmware) GetStorageController() StorageControllerRelationship`

GetStorageController returns the StorageController field if non-nil, zero value otherwise.

### GetStorageControllerOk

`func (o *FirmwareRunningFirmware) GetStorageControllerOk() (*StorageControllerRelationship, bool)`

GetStorageControllerOk returns a tuple with the StorageController field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageController

`func (o *FirmwareRunningFirmware) SetStorageController(v StorageControllerRelationship)`

SetStorageController sets StorageController field to given value.

### HasStorageController

`func (o *FirmwareRunningFirmware) HasStorageController() bool`

HasStorageController returns a boolean if a field has been set.

### SetStorageControllerNil

`func (o *FirmwareRunningFirmware) SetStorageControllerNil(b bool)`

 SetStorageControllerNil sets the value for StorageController to be an explicit nil

### UnsetStorageController
`func (o *FirmwareRunningFirmware) UnsetStorageController()`

UnsetStorageController ensures that no value is present for StorageController, not even an explicit nil
### GetStorageFlexFlashController

`func (o *FirmwareRunningFirmware) GetStorageFlexFlashController() StorageFlexFlashControllerRelationship`

GetStorageFlexFlashController returns the StorageFlexFlashController field if non-nil, zero value otherwise.

### GetStorageFlexFlashControllerOk

`func (o *FirmwareRunningFirmware) GetStorageFlexFlashControllerOk() (*StorageFlexFlashControllerRelationship, bool)`

GetStorageFlexFlashControllerOk returns a tuple with the StorageFlexFlashController field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageFlexFlashController

`func (o *FirmwareRunningFirmware) SetStorageFlexFlashController(v StorageFlexFlashControllerRelationship)`

SetStorageFlexFlashController sets StorageFlexFlashController field to given value.

### HasStorageFlexFlashController

`func (o *FirmwareRunningFirmware) HasStorageFlexFlashController() bool`

HasStorageFlexFlashController returns a boolean if a field has been set.

### SetStorageFlexFlashControllerNil

`func (o *FirmwareRunningFirmware) SetStorageFlexFlashControllerNil(b bool)`

 SetStorageFlexFlashControllerNil sets the value for StorageFlexFlashController to be an explicit nil

### UnsetStorageFlexFlashController
`func (o *FirmwareRunningFirmware) UnsetStorageFlexFlashController()`

UnsetStorageFlexFlashController ensures that no value is present for StorageFlexFlashController, not even an explicit nil
### GetStoragePhysicalDisk

`func (o *FirmwareRunningFirmware) GetStoragePhysicalDisk() StoragePhysicalDiskRelationship`

GetStoragePhysicalDisk returns the StoragePhysicalDisk field if non-nil, zero value otherwise.

### GetStoragePhysicalDiskOk

`func (o *FirmwareRunningFirmware) GetStoragePhysicalDiskOk() (*StoragePhysicalDiskRelationship, bool)`

GetStoragePhysicalDiskOk returns a tuple with the StoragePhysicalDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoragePhysicalDisk

`func (o *FirmwareRunningFirmware) SetStoragePhysicalDisk(v StoragePhysicalDiskRelationship)`

SetStoragePhysicalDisk sets StoragePhysicalDisk field to given value.

### HasStoragePhysicalDisk

`func (o *FirmwareRunningFirmware) HasStoragePhysicalDisk() bool`

HasStoragePhysicalDisk returns a boolean if a field has been set.

### SetStoragePhysicalDiskNil

`func (o *FirmwareRunningFirmware) SetStoragePhysicalDiskNil(b bool)`

 SetStoragePhysicalDiskNil sets the value for StoragePhysicalDisk to be an explicit nil

### UnsetStoragePhysicalDisk
`func (o *FirmwareRunningFirmware) UnsetStoragePhysicalDisk()`

UnsetStoragePhysicalDisk ensures that no value is present for StoragePhysicalDisk, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



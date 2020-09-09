# FirmwareRunningFirmwareAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Component** | Pointer to **string** | Kind of the firmware - boot-booloader/system/kernel. | [optional] [readonly] 
**PackageVersion** | Pointer to **string** | Package version which the firmware belongs to. | [optional] [readonly] 
**Type** | Pointer to **string** | The type of the firmware. | [optional] [readonly] 
**Version** | Pointer to **string** | The version of the firmware. | [optional] [readonly] 
**BiosUnit** | Pointer to [**BiosUnitRelationship**](bios.Unit.Relationship.md) |  | [optional] 
**GraphicsCard** | Pointer to [**GraphicsCardRelationship**](graphics.Card.Relationship.md) |  | [optional] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](inventory.DeviceInfo.Relationship.md) |  | [optional] 
**ManagementController** | Pointer to [**ManagementControllerRelationship**](management.Controller.Relationship.md) |  | [optional] 
**NetworkElements** | Pointer to [**[]NetworkElementRelationship**](network.Element.Relationship.md) | An array of relationships to networkElement resources. | [optional] 
**PciSwitch** | Pointer to [**PciSwitchRelationship**](pci.Switch.Relationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 
**StorageController** | Pointer to [**StorageControllerRelationship**](storage.Controller.Relationship.md) |  | [optional] 
**StorageFlexFlashController** | Pointer to [**StorageFlexFlashControllerRelationship**](storage.FlexFlashController.Relationship.md) |  | [optional] 
**StoragePhysicalDisk** | Pointer to [**StoragePhysicalDiskRelationship**](storage.PhysicalDisk.Relationship.md) |  | [optional] 

## Methods

### NewFirmwareRunningFirmwareAllOf

`func NewFirmwareRunningFirmwareAllOf() *FirmwareRunningFirmwareAllOf`

NewFirmwareRunningFirmwareAllOf instantiates a new FirmwareRunningFirmwareAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirmwareRunningFirmwareAllOfWithDefaults

`func NewFirmwareRunningFirmwareAllOfWithDefaults() *FirmwareRunningFirmwareAllOf`

NewFirmwareRunningFirmwareAllOfWithDefaults instantiates a new FirmwareRunningFirmwareAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponent

`func (o *FirmwareRunningFirmwareAllOf) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *FirmwareRunningFirmwareAllOf) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *FirmwareRunningFirmwareAllOf) SetComponent(v string)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *FirmwareRunningFirmwareAllOf) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetPackageVersion

`func (o *FirmwareRunningFirmwareAllOf) GetPackageVersion() string`

GetPackageVersion returns the PackageVersion field if non-nil, zero value otherwise.

### GetPackageVersionOk

`func (o *FirmwareRunningFirmwareAllOf) GetPackageVersionOk() (*string, bool)`

GetPackageVersionOk returns a tuple with the PackageVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageVersion

`func (o *FirmwareRunningFirmwareAllOf) SetPackageVersion(v string)`

SetPackageVersion sets PackageVersion field to given value.

### HasPackageVersion

`func (o *FirmwareRunningFirmwareAllOf) HasPackageVersion() bool`

HasPackageVersion returns a boolean if a field has been set.

### GetType

`func (o *FirmwareRunningFirmwareAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FirmwareRunningFirmwareAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FirmwareRunningFirmwareAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FirmwareRunningFirmwareAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVersion

`func (o *FirmwareRunningFirmwareAllOf) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *FirmwareRunningFirmwareAllOf) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *FirmwareRunningFirmwareAllOf) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *FirmwareRunningFirmwareAllOf) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetBiosUnit

`func (o *FirmwareRunningFirmwareAllOf) GetBiosUnit() BiosUnitRelationship`

GetBiosUnit returns the BiosUnit field if non-nil, zero value otherwise.

### GetBiosUnitOk

`func (o *FirmwareRunningFirmwareAllOf) GetBiosUnitOk() (*BiosUnitRelationship, bool)`

GetBiosUnitOk returns a tuple with the BiosUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiosUnit

`func (o *FirmwareRunningFirmwareAllOf) SetBiosUnit(v BiosUnitRelationship)`

SetBiosUnit sets BiosUnit field to given value.

### HasBiosUnit

`func (o *FirmwareRunningFirmwareAllOf) HasBiosUnit() bool`

HasBiosUnit returns a boolean if a field has been set.

### GetGraphicsCard

`func (o *FirmwareRunningFirmwareAllOf) GetGraphicsCard() GraphicsCardRelationship`

GetGraphicsCard returns the GraphicsCard field if non-nil, zero value otherwise.

### GetGraphicsCardOk

`func (o *FirmwareRunningFirmwareAllOf) GetGraphicsCardOk() (*GraphicsCardRelationship, bool)`

GetGraphicsCardOk returns a tuple with the GraphicsCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraphicsCard

`func (o *FirmwareRunningFirmwareAllOf) SetGraphicsCard(v GraphicsCardRelationship)`

SetGraphicsCard sets GraphicsCard field to given value.

### HasGraphicsCard

`func (o *FirmwareRunningFirmwareAllOf) HasGraphicsCard() bool`

HasGraphicsCard returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *FirmwareRunningFirmwareAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *FirmwareRunningFirmwareAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *FirmwareRunningFirmwareAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *FirmwareRunningFirmwareAllOf) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetManagementController

`func (o *FirmwareRunningFirmwareAllOf) GetManagementController() ManagementControllerRelationship`

GetManagementController returns the ManagementController field if non-nil, zero value otherwise.

### GetManagementControllerOk

`func (o *FirmwareRunningFirmwareAllOf) GetManagementControllerOk() (*ManagementControllerRelationship, bool)`

GetManagementControllerOk returns a tuple with the ManagementController field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementController

`func (o *FirmwareRunningFirmwareAllOf) SetManagementController(v ManagementControllerRelationship)`

SetManagementController sets ManagementController field to given value.

### HasManagementController

`func (o *FirmwareRunningFirmwareAllOf) HasManagementController() bool`

HasManagementController returns a boolean if a field has been set.

### GetNetworkElements

`func (o *FirmwareRunningFirmwareAllOf) GetNetworkElements() []NetworkElementRelationship`

GetNetworkElements returns the NetworkElements field if non-nil, zero value otherwise.

### GetNetworkElementsOk

`func (o *FirmwareRunningFirmwareAllOf) GetNetworkElementsOk() (*[]NetworkElementRelationship, bool)`

GetNetworkElementsOk returns a tuple with the NetworkElements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkElements

`func (o *FirmwareRunningFirmwareAllOf) SetNetworkElements(v []NetworkElementRelationship)`

SetNetworkElements sets NetworkElements field to given value.

### HasNetworkElements

`func (o *FirmwareRunningFirmwareAllOf) HasNetworkElements() bool`

HasNetworkElements returns a boolean if a field has been set.

### SetNetworkElementsNil

`func (o *FirmwareRunningFirmwareAllOf) SetNetworkElementsNil(b bool)`

 SetNetworkElementsNil sets the value for NetworkElements to be an explicit nil

### UnsetNetworkElements
`func (o *FirmwareRunningFirmwareAllOf) UnsetNetworkElements()`

UnsetNetworkElements ensures that no value is present for NetworkElements, not even an explicit nil
### GetPciSwitch

`func (o *FirmwareRunningFirmwareAllOf) GetPciSwitch() PciSwitchRelationship`

GetPciSwitch returns the PciSwitch field if non-nil, zero value otherwise.

### GetPciSwitchOk

`func (o *FirmwareRunningFirmwareAllOf) GetPciSwitchOk() (*PciSwitchRelationship, bool)`

GetPciSwitchOk returns a tuple with the PciSwitch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciSwitch

`func (o *FirmwareRunningFirmwareAllOf) SetPciSwitch(v PciSwitchRelationship)`

SetPciSwitch sets PciSwitch field to given value.

### HasPciSwitch

`func (o *FirmwareRunningFirmwareAllOf) HasPciSwitch() bool`

HasPciSwitch returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *FirmwareRunningFirmwareAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *FirmwareRunningFirmwareAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *FirmwareRunningFirmwareAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *FirmwareRunningFirmwareAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### GetStorageController

`func (o *FirmwareRunningFirmwareAllOf) GetStorageController() StorageControllerRelationship`

GetStorageController returns the StorageController field if non-nil, zero value otherwise.

### GetStorageControllerOk

`func (o *FirmwareRunningFirmwareAllOf) GetStorageControllerOk() (*StorageControllerRelationship, bool)`

GetStorageControllerOk returns a tuple with the StorageController field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageController

`func (o *FirmwareRunningFirmwareAllOf) SetStorageController(v StorageControllerRelationship)`

SetStorageController sets StorageController field to given value.

### HasStorageController

`func (o *FirmwareRunningFirmwareAllOf) HasStorageController() bool`

HasStorageController returns a boolean if a field has been set.

### GetStorageFlexFlashController

`func (o *FirmwareRunningFirmwareAllOf) GetStorageFlexFlashController() StorageFlexFlashControllerRelationship`

GetStorageFlexFlashController returns the StorageFlexFlashController field if non-nil, zero value otherwise.

### GetStorageFlexFlashControllerOk

`func (o *FirmwareRunningFirmwareAllOf) GetStorageFlexFlashControllerOk() (*StorageFlexFlashControllerRelationship, bool)`

GetStorageFlexFlashControllerOk returns a tuple with the StorageFlexFlashController field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageFlexFlashController

`func (o *FirmwareRunningFirmwareAllOf) SetStorageFlexFlashController(v StorageFlexFlashControllerRelationship)`

SetStorageFlexFlashController sets StorageFlexFlashController field to given value.

### HasStorageFlexFlashController

`func (o *FirmwareRunningFirmwareAllOf) HasStorageFlexFlashController() bool`

HasStorageFlexFlashController returns a boolean if a field has been set.

### GetStoragePhysicalDisk

`func (o *FirmwareRunningFirmwareAllOf) GetStoragePhysicalDisk() StoragePhysicalDiskRelationship`

GetStoragePhysicalDisk returns the StoragePhysicalDisk field if non-nil, zero value otherwise.

### GetStoragePhysicalDiskOk

`func (o *FirmwareRunningFirmwareAllOf) GetStoragePhysicalDiskOk() (*StoragePhysicalDiskRelationship, bool)`

GetStoragePhysicalDiskOk returns a tuple with the StoragePhysicalDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoragePhysicalDisk

`func (o *FirmwareRunningFirmwareAllOf) SetStoragePhysicalDisk(v StoragePhysicalDiskRelationship)`

SetStoragePhysicalDisk sets StoragePhysicalDisk field to given value.

### HasStoragePhysicalDisk

`func (o *FirmwareRunningFirmwareAllOf) HasStoragePhysicalDisk() bool`

HasStoragePhysicalDisk returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



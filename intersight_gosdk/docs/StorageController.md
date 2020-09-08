# StorageController

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ControllerFlags** | Pointer to **string** | The flags for the storage controller. | [optional] [readonly] 
**ControllerId** | Pointer to **string** | The Id of the storage controller. | [optional] [readonly] 
**ControllerStatus** | Pointer to **string** | The current status of controller. | [optional] [readonly] 
**ForeignConfigPresent** | Pointer to **bool** | Storage controller has detected disks in foreign config. | [optional] 
**HwRevision** | Pointer to **string** | The hardware revision of controller. | [optional] [readonly] 
**InterfaceType** | Pointer to **string** | Interface types are Sas, Sata, PCH. | [optional] 
**MaxVolumesSupported** | Pointer to **int64** | Maximum virtual drives that can be created on this Storage Controller. | [optional] 
**OobInterfaceSupported** | Pointer to **string** | The CIMC support for out-of-band configuration of controller. | [optional] [readonly] 
**OperState** | Pointer to **string** | The current operational state of controller. | [optional] [readonly] 
**Operability** | Pointer to **string** | Operability state of the storage controller. | [optional] [readonly] 
**PciAddr** | Pointer to **string** | The current pci address of controller. | [optional] [readonly] 
**PciSlot** | Pointer to **string** | The pci slot name for the controller. | [optional] [readonly] 
**Presence** | Pointer to **string** | Physical Presence State for the Storage Controller. | [optional] [readonly] 
**RaidSupport** | Pointer to **string** | The RAID levels supported by controller. | [optional] [readonly] 
**RebuildRate** | Pointer to **string** | Logical volume or RAID rebuild rate of Storage Controller. | [optional] [readonly] 
**SelfEncryptEnabled** | Pointer to **string** | Storage controller disk self encryption state. | [optional] 
**Type** | Pointer to **string** | Controller types are Raid, FlexFlash. | [optional] [readonly] 
**ComputeBlade** | Pointer to [**ComputeBladeRelationship**](compute.Blade.Relationship.md) |  | [optional] 
**ComputeBoard** | Pointer to [**ComputeBoardRelationship**](compute.Board.Relationship.md) |  | [optional] 
**ComputeRackUnit** | Pointer to [**ComputeRackUnitRelationship**](compute.RackUnit.Relationship.md) |  | [optional] 
**DiskGroup** | Pointer to [**[]StorageDiskGroupRelationship**](storage.DiskGroup.Relationship.md) | An array of relationships to storageDiskGroup resources. | [optional] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](inventory.DeviceInfo.Relationship.md) |  | [optional] 
**PhysicalDiskExtensions** | Pointer to [**[]StoragePhysicalDiskExtensionRelationship**](storage.PhysicalDiskExtension.Relationship.md) | An array of relationships to storagePhysicalDiskExtension resources. | [optional] [readonly] 
**PhysicalDisks** | Pointer to [**[]StoragePhysicalDiskRelationship**](storage.PhysicalDisk.Relationship.md) | An array of relationships to storagePhysicalDisk resources. | [optional] [readonly] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 
**RunningFirmware** | Pointer to [**[]FirmwareRunningFirmwareRelationship**](firmware.RunningFirmware.Relationship.md) | An array of relationships to firmwareRunningFirmware resources. | [optional] [readonly] 
**VirtualDriveExtensions** | Pointer to [**[]StorageVirtualDriveExtensionRelationship**](storage.VirtualDriveExtension.Relationship.md) | An array of relationships to storageVirtualDriveExtension resources. | [optional] [readonly] 
**VirtualDrives** | Pointer to [**[]StorageVirtualDriveRelationship**](storage.VirtualDrive.Relationship.md) | An array of relationships to storageVirtualDrive resources. | [optional] [readonly] 

## Methods

### NewStorageController

`func NewStorageController() *StorageController`

NewStorageController instantiates a new StorageController object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageControllerWithDefaults

`func NewStorageControllerWithDefaults() *StorageController`

NewStorageControllerWithDefaults instantiates a new StorageController object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetControllerFlags

`func (o *StorageController) GetControllerFlags() string`

GetControllerFlags returns the ControllerFlags field if non-nil, zero value otherwise.

### GetControllerFlagsOk

`func (o *StorageController) GetControllerFlagsOk() (*string, bool)`

GetControllerFlagsOk returns a tuple with the ControllerFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerFlags

`func (o *StorageController) SetControllerFlags(v string)`

SetControllerFlags sets ControllerFlags field to given value.

### HasControllerFlags

`func (o *StorageController) HasControllerFlags() bool`

HasControllerFlags returns a boolean if a field has been set.

### GetControllerId

`func (o *StorageController) GetControllerId() string`

GetControllerId returns the ControllerId field if non-nil, zero value otherwise.

### GetControllerIdOk

`func (o *StorageController) GetControllerIdOk() (*string, bool)`

GetControllerIdOk returns a tuple with the ControllerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerId

`func (o *StorageController) SetControllerId(v string)`

SetControllerId sets ControllerId field to given value.

### HasControllerId

`func (o *StorageController) HasControllerId() bool`

HasControllerId returns a boolean if a field has been set.

### GetControllerStatus

`func (o *StorageController) GetControllerStatus() string`

GetControllerStatus returns the ControllerStatus field if non-nil, zero value otherwise.

### GetControllerStatusOk

`func (o *StorageController) GetControllerStatusOk() (*string, bool)`

GetControllerStatusOk returns a tuple with the ControllerStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerStatus

`func (o *StorageController) SetControllerStatus(v string)`

SetControllerStatus sets ControllerStatus field to given value.

### HasControllerStatus

`func (o *StorageController) HasControllerStatus() bool`

HasControllerStatus returns a boolean if a field has been set.

### GetForeignConfigPresent

`func (o *StorageController) GetForeignConfigPresent() bool`

GetForeignConfigPresent returns the ForeignConfigPresent field if non-nil, zero value otherwise.

### GetForeignConfigPresentOk

`func (o *StorageController) GetForeignConfigPresentOk() (*bool, bool)`

GetForeignConfigPresentOk returns a tuple with the ForeignConfigPresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignConfigPresent

`func (o *StorageController) SetForeignConfigPresent(v bool)`

SetForeignConfigPresent sets ForeignConfigPresent field to given value.

### HasForeignConfigPresent

`func (o *StorageController) HasForeignConfigPresent() bool`

HasForeignConfigPresent returns a boolean if a field has been set.

### GetHwRevision

`func (o *StorageController) GetHwRevision() string`

GetHwRevision returns the HwRevision field if non-nil, zero value otherwise.

### GetHwRevisionOk

`func (o *StorageController) GetHwRevisionOk() (*string, bool)`

GetHwRevisionOk returns a tuple with the HwRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHwRevision

`func (o *StorageController) SetHwRevision(v string)`

SetHwRevision sets HwRevision field to given value.

### HasHwRevision

`func (o *StorageController) HasHwRevision() bool`

HasHwRevision returns a boolean if a field has been set.

### GetInterfaceType

`func (o *StorageController) GetInterfaceType() string`

GetInterfaceType returns the InterfaceType field if non-nil, zero value otherwise.

### GetInterfaceTypeOk

`func (o *StorageController) GetInterfaceTypeOk() (*string, bool)`

GetInterfaceTypeOk returns a tuple with the InterfaceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceType

`func (o *StorageController) SetInterfaceType(v string)`

SetInterfaceType sets InterfaceType field to given value.

### HasInterfaceType

`func (o *StorageController) HasInterfaceType() bool`

HasInterfaceType returns a boolean if a field has been set.

### GetMaxVolumesSupported

`func (o *StorageController) GetMaxVolumesSupported() int64`

GetMaxVolumesSupported returns the MaxVolumesSupported field if non-nil, zero value otherwise.

### GetMaxVolumesSupportedOk

`func (o *StorageController) GetMaxVolumesSupportedOk() (*int64, bool)`

GetMaxVolumesSupportedOk returns a tuple with the MaxVolumesSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVolumesSupported

`func (o *StorageController) SetMaxVolumesSupported(v int64)`

SetMaxVolumesSupported sets MaxVolumesSupported field to given value.

### HasMaxVolumesSupported

`func (o *StorageController) HasMaxVolumesSupported() bool`

HasMaxVolumesSupported returns a boolean if a field has been set.

### GetOobInterfaceSupported

`func (o *StorageController) GetOobInterfaceSupported() string`

GetOobInterfaceSupported returns the OobInterfaceSupported field if non-nil, zero value otherwise.

### GetOobInterfaceSupportedOk

`func (o *StorageController) GetOobInterfaceSupportedOk() (*string, bool)`

GetOobInterfaceSupportedOk returns a tuple with the OobInterfaceSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOobInterfaceSupported

`func (o *StorageController) SetOobInterfaceSupported(v string)`

SetOobInterfaceSupported sets OobInterfaceSupported field to given value.

### HasOobInterfaceSupported

`func (o *StorageController) HasOobInterfaceSupported() bool`

HasOobInterfaceSupported returns a boolean if a field has been set.

### GetOperState

`func (o *StorageController) GetOperState() string`

GetOperState returns the OperState field if non-nil, zero value otherwise.

### GetOperStateOk

`func (o *StorageController) GetOperStateOk() (*string, bool)`

GetOperStateOk returns a tuple with the OperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperState

`func (o *StorageController) SetOperState(v string)`

SetOperState sets OperState field to given value.

### HasOperState

`func (o *StorageController) HasOperState() bool`

HasOperState returns a boolean if a field has been set.

### GetOperability

`func (o *StorageController) GetOperability() string`

GetOperability returns the Operability field if non-nil, zero value otherwise.

### GetOperabilityOk

`func (o *StorageController) GetOperabilityOk() (*string, bool)`

GetOperabilityOk returns a tuple with the Operability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperability

`func (o *StorageController) SetOperability(v string)`

SetOperability sets Operability field to given value.

### HasOperability

`func (o *StorageController) HasOperability() bool`

HasOperability returns a boolean if a field has been set.

### GetPciAddr

`func (o *StorageController) GetPciAddr() string`

GetPciAddr returns the PciAddr field if non-nil, zero value otherwise.

### GetPciAddrOk

`func (o *StorageController) GetPciAddrOk() (*string, bool)`

GetPciAddrOk returns a tuple with the PciAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciAddr

`func (o *StorageController) SetPciAddr(v string)`

SetPciAddr sets PciAddr field to given value.

### HasPciAddr

`func (o *StorageController) HasPciAddr() bool`

HasPciAddr returns a boolean if a field has been set.

### GetPciSlot

`func (o *StorageController) GetPciSlot() string`

GetPciSlot returns the PciSlot field if non-nil, zero value otherwise.

### GetPciSlotOk

`func (o *StorageController) GetPciSlotOk() (*string, bool)`

GetPciSlotOk returns a tuple with the PciSlot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciSlot

`func (o *StorageController) SetPciSlot(v string)`

SetPciSlot sets PciSlot field to given value.

### HasPciSlot

`func (o *StorageController) HasPciSlot() bool`

HasPciSlot returns a boolean if a field has been set.

### GetPresence

`func (o *StorageController) GetPresence() string`

GetPresence returns the Presence field if non-nil, zero value otherwise.

### GetPresenceOk

`func (o *StorageController) GetPresenceOk() (*string, bool)`

GetPresenceOk returns a tuple with the Presence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresence

`func (o *StorageController) SetPresence(v string)`

SetPresence sets Presence field to given value.

### HasPresence

`func (o *StorageController) HasPresence() bool`

HasPresence returns a boolean if a field has been set.

### GetRaidSupport

`func (o *StorageController) GetRaidSupport() string`

GetRaidSupport returns the RaidSupport field if non-nil, zero value otherwise.

### GetRaidSupportOk

`func (o *StorageController) GetRaidSupportOk() (*string, bool)`

GetRaidSupportOk returns a tuple with the RaidSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaidSupport

`func (o *StorageController) SetRaidSupport(v string)`

SetRaidSupport sets RaidSupport field to given value.

### HasRaidSupport

`func (o *StorageController) HasRaidSupport() bool`

HasRaidSupport returns a boolean if a field has been set.

### GetRebuildRate

`func (o *StorageController) GetRebuildRate() string`

GetRebuildRate returns the RebuildRate field if non-nil, zero value otherwise.

### GetRebuildRateOk

`func (o *StorageController) GetRebuildRateOk() (*string, bool)`

GetRebuildRateOk returns a tuple with the RebuildRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRebuildRate

`func (o *StorageController) SetRebuildRate(v string)`

SetRebuildRate sets RebuildRate field to given value.

### HasRebuildRate

`func (o *StorageController) HasRebuildRate() bool`

HasRebuildRate returns a boolean if a field has been set.

### GetSelfEncryptEnabled

`func (o *StorageController) GetSelfEncryptEnabled() string`

GetSelfEncryptEnabled returns the SelfEncryptEnabled field if non-nil, zero value otherwise.

### GetSelfEncryptEnabledOk

`func (o *StorageController) GetSelfEncryptEnabledOk() (*string, bool)`

GetSelfEncryptEnabledOk returns a tuple with the SelfEncryptEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfEncryptEnabled

`func (o *StorageController) SetSelfEncryptEnabled(v string)`

SetSelfEncryptEnabled sets SelfEncryptEnabled field to given value.

### HasSelfEncryptEnabled

`func (o *StorageController) HasSelfEncryptEnabled() bool`

HasSelfEncryptEnabled returns a boolean if a field has been set.

### GetType

`func (o *StorageController) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StorageController) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StorageController) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *StorageController) HasType() bool`

HasType returns a boolean if a field has been set.

### GetComputeBlade

`func (o *StorageController) GetComputeBlade() ComputeBladeRelationship`

GetComputeBlade returns the ComputeBlade field if non-nil, zero value otherwise.

### GetComputeBladeOk

`func (o *StorageController) GetComputeBladeOk() (*ComputeBladeRelationship, bool)`

GetComputeBladeOk returns a tuple with the ComputeBlade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeBlade

`func (o *StorageController) SetComputeBlade(v ComputeBladeRelationship)`

SetComputeBlade sets ComputeBlade field to given value.

### HasComputeBlade

`func (o *StorageController) HasComputeBlade() bool`

HasComputeBlade returns a boolean if a field has been set.

### GetComputeBoard

`func (o *StorageController) GetComputeBoard() ComputeBoardRelationship`

GetComputeBoard returns the ComputeBoard field if non-nil, zero value otherwise.

### GetComputeBoardOk

`func (o *StorageController) GetComputeBoardOk() (*ComputeBoardRelationship, bool)`

GetComputeBoardOk returns a tuple with the ComputeBoard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeBoard

`func (o *StorageController) SetComputeBoard(v ComputeBoardRelationship)`

SetComputeBoard sets ComputeBoard field to given value.

### HasComputeBoard

`func (o *StorageController) HasComputeBoard() bool`

HasComputeBoard returns a boolean if a field has been set.

### GetComputeRackUnit

`func (o *StorageController) GetComputeRackUnit() ComputeRackUnitRelationship`

GetComputeRackUnit returns the ComputeRackUnit field if non-nil, zero value otherwise.

### GetComputeRackUnitOk

`func (o *StorageController) GetComputeRackUnitOk() (*ComputeRackUnitRelationship, bool)`

GetComputeRackUnitOk returns a tuple with the ComputeRackUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRackUnit

`func (o *StorageController) SetComputeRackUnit(v ComputeRackUnitRelationship)`

SetComputeRackUnit sets ComputeRackUnit field to given value.

### HasComputeRackUnit

`func (o *StorageController) HasComputeRackUnit() bool`

HasComputeRackUnit returns a boolean if a field has been set.

### GetDiskGroup

`func (o *StorageController) GetDiskGroup() []StorageDiskGroupRelationship`

GetDiskGroup returns the DiskGroup field if non-nil, zero value otherwise.

### GetDiskGroupOk

`func (o *StorageController) GetDiskGroupOk() (*[]StorageDiskGroupRelationship, bool)`

GetDiskGroupOk returns a tuple with the DiskGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskGroup

`func (o *StorageController) SetDiskGroup(v []StorageDiskGroupRelationship)`

SetDiskGroup sets DiskGroup field to given value.

### HasDiskGroup

`func (o *StorageController) HasDiskGroup() bool`

HasDiskGroup returns a boolean if a field has been set.

### SetDiskGroupNil

`func (o *StorageController) SetDiskGroupNil(b bool)`

 SetDiskGroupNil sets the value for DiskGroup to be an explicit nil

### UnsetDiskGroup
`func (o *StorageController) UnsetDiskGroup()`

UnsetDiskGroup ensures that no value is present for DiskGroup, not even an explicit nil
### GetInventoryDeviceInfo

`func (o *StorageController) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *StorageController) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *StorageController) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *StorageController) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetPhysicalDiskExtensions

`func (o *StorageController) GetPhysicalDiskExtensions() []StoragePhysicalDiskExtensionRelationship`

GetPhysicalDiskExtensions returns the PhysicalDiskExtensions field if non-nil, zero value otherwise.

### GetPhysicalDiskExtensionsOk

`func (o *StorageController) GetPhysicalDiskExtensionsOk() (*[]StoragePhysicalDiskExtensionRelationship, bool)`

GetPhysicalDiskExtensionsOk returns a tuple with the PhysicalDiskExtensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalDiskExtensions

`func (o *StorageController) SetPhysicalDiskExtensions(v []StoragePhysicalDiskExtensionRelationship)`

SetPhysicalDiskExtensions sets PhysicalDiskExtensions field to given value.

### HasPhysicalDiskExtensions

`func (o *StorageController) HasPhysicalDiskExtensions() bool`

HasPhysicalDiskExtensions returns a boolean if a field has been set.

### SetPhysicalDiskExtensionsNil

`func (o *StorageController) SetPhysicalDiskExtensionsNil(b bool)`

 SetPhysicalDiskExtensionsNil sets the value for PhysicalDiskExtensions to be an explicit nil

### UnsetPhysicalDiskExtensions
`func (o *StorageController) UnsetPhysicalDiskExtensions()`

UnsetPhysicalDiskExtensions ensures that no value is present for PhysicalDiskExtensions, not even an explicit nil
### GetPhysicalDisks

`func (o *StorageController) GetPhysicalDisks() []StoragePhysicalDiskRelationship`

GetPhysicalDisks returns the PhysicalDisks field if non-nil, zero value otherwise.

### GetPhysicalDisksOk

`func (o *StorageController) GetPhysicalDisksOk() (*[]StoragePhysicalDiskRelationship, bool)`

GetPhysicalDisksOk returns a tuple with the PhysicalDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalDisks

`func (o *StorageController) SetPhysicalDisks(v []StoragePhysicalDiskRelationship)`

SetPhysicalDisks sets PhysicalDisks field to given value.

### HasPhysicalDisks

`func (o *StorageController) HasPhysicalDisks() bool`

HasPhysicalDisks returns a boolean if a field has been set.

### SetPhysicalDisksNil

`func (o *StorageController) SetPhysicalDisksNil(b bool)`

 SetPhysicalDisksNil sets the value for PhysicalDisks to be an explicit nil

### UnsetPhysicalDisks
`func (o *StorageController) UnsetPhysicalDisks()`

UnsetPhysicalDisks ensures that no value is present for PhysicalDisks, not even an explicit nil
### GetRegisteredDevice

`func (o *StorageController) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StorageController) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StorageController) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StorageController) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### GetRunningFirmware

`func (o *StorageController) GetRunningFirmware() []FirmwareRunningFirmwareRelationship`

GetRunningFirmware returns the RunningFirmware field if non-nil, zero value otherwise.

### GetRunningFirmwareOk

`func (o *StorageController) GetRunningFirmwareOk() (*[]FirmwareRunningFirmwareRelationship, bool)`

GetRunningFirmwareOk returns a tuple with the RunningFirmware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningFirmware

`func (o *StorageController) SetRunningFirmware(v []FirmwareRunningFirmwareRelationship)`

SetRunningFirmware sets RunningFirmware field to given value.

### HasRunningFirmware

`func (o *StorageController) HasRunningFirmware() bool`

HasRunningFirmware returns a boolean if a field has been set.

### SetRunningFirmwareNil

`func (o *StorageController) SetRunningFirmwareNil(b bool)`

 SetRunningFirmwareNil sets the value for RunningFirmware to be an explicit nil

### UnsetRunningFirmware
`func (o *StorageController) UnsetRunningFirmware()`

UnsetRunningFirmware ensures that no value is present for RunningFirmware, not even an explicit nil
### GetVirtualDriveExtensions

`func (o *StorageController) GetVirtualDriveExtensions() []StorageVirtualDriveExtensionRelationship`

GetVirtualDriveExtensions returns the VirtualDriveExtensions field if non-nil, zero value otherwise.

### GetVirtualDriveExtensionsOk

`func (o *StorageController) GetVirtualDriveExtensionsOk() (*[]StorageVirtualDriveExtensionRelationship, bool)`

GetVirtualDriveExtensionsOk returns a tuple with the VirtualDriveExtensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualDriveExtensions

`func (o *StorageController) SetVirtualDriveExtensions(v []StorageVirtualDriveExtensionRelationship)`

SetVirtualDriveExtensions sets VirtualDriveExtensions field to given value.

### HasVirtualDriveExtensions

`func (o *StorageController) HasVirtualDriveExtensions() bool`

HasVirtualDriveExtensions returns a boolean if a field has been set.

### SetVirtualDriveExtensionsNil

`func (o *StorageController) SetVirtualDriveExtensionsNil(b bool)`

 SetVirtualDriveExtensionsNil sets the value for VirtualDriveExtensions to be an explicit nil

### UnsetVirtualDriveExtensions
`func (o *StorageController) UnsetVirtualDriveExtensions()`

UnsetVirtualDriveExtensions ensures that no value is present for VirtualDriveExtensions, not even an explicit nil
### GetVirtualDrives

`func (o *StorageController) GetVirtualDrives() []StorageVirtualDriveRelationship`

GetVirtualDrives returns the VirtualDrives field if non-nil, zero value otherwise.

### GetVirtualDrivesOk

`func (o *StorageController) GetVirtualDrivesOk() (*[]StorageVirtualDriveRelationship, bool)`

GetVirtualDrivesOk returns a tuple with the VirtualDrives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualDrives

`func (o *StorageController) SetVirtualDrives(v []StorageVirtualDriveRelationship)`

SetVirtualDrives sets VirtualDrives field to given value.

### HasVirtualDrives

`func (o *StorageController) HasVirtualDrives() bool`

HasVirtualDrives returns a boolean if a field has been set.

### SetVirtualDrivesNil

`func (o *StorageController) SetVirtualDrivesNil(b bool)`

 SetVirtualDrivesNil sets the value for VirtualDrives to be an explicit nil

### UnsetVirtualDrives
`func (o *StorageController) UnsetVirtualDrives()`

UnsetVirtualDrives ensures that no value is present for VirtualDrives, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



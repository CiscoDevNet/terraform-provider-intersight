# StorageControllerAllOf

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

### NewStorageControllerAllOf

`func NewStorageControllerAllOf() *StorageControllerAllOf`

NewStorageControllerAllOf instantiates a new StorageControllerAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageControllerAllOfWithDefaults

`func NewStorageControllerAllOfWithDefaults() *StorageControllerAllOf`

NewStorageControllerAllOfWithDefaults instantiates a new StorageControllerAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetControllerFlags

`func (o *StorageControllerAllOf) GetControllerFlags() string`

GetControllerFlags returns the ControllerFlags field if non-nil, zero value otherwise.

### GetControllerFlagsOk

`func (o *StorageControllerAllOf) GetControllerFlagsOk() (*string, bool)`

GetControllerFlagsOk returns a tuple with the ControllerFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerFlags

`func (o *StorageControllerAllOf) SetControllerFlags(v string)`

SetControllerFlags sets ControllerFlags field to given value.

### HasControllerFlags

`func (o *StorageControllerAllOf) HasControllerFlags() bool`

HasControllerFlags returns a boolean if a field has been set.

### GetControllerId

`func (o *StorageControllerAllOf) GetControllerId() string`

GetControllerId returns the ControllerId field if non-nil, zero value otherwise.

### GetControllerIdOk

`func (o *StorageControllerAllOf) GetControllerIdOk() (*string, bool)`

GetControllerIdOk returns a tuple with the ControllerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerId

`func (o *StorageControllerAllOf) SetControllerId(v string)`

SetControllerId sets ControllerId field to given value.

### HasControllerId

`func (o *StorageControllerAllOf) HasControllerId() bool`

HasControllerId returns a boolean if a field has been set.

### GetControllerStatus

`func (o *StorageControllerAllOf) GetControllerStatus() string`

GetControllerStatus returns the ControllerStatus field if non-nil, zero value otherwise.

### GetControllerStatusOk

`func (o *StorageControllerAllOf) GetControllerStatusOk() (*string, bool)`

GetControllerStatusOk returns a tuple with the ControllerStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerStatus

`func (o *StorageControllerAllOf) SetControllerStatus(v string)`

SetControllerStatus sets ControllerStatus field to given value.

### HasControllerStatus

`func (o *StorageControllerAllOf) HasControllerStatus() bool`

HasControllerStatus returns a boolean if a field has been set.

### GetForeignConfigPresent

`func (o *StorageControllerAllOf) GetForeignConfigPresent() bool`

GetForeignConfigPresent returns the ForeignConfigPresent field if non-nil, zero value otherwise.

### GetForeignConfigPresentOk

`func (o *StorageControllerAllOf) GetForeignConfigPresentOk() (*bool, bool)`

GetForeignConfigPresentOk returns a tuple with the ForeignConfigPresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignConfigPresent

`func (o *StorageControllerAllOf) SetForeignConfigPresent(v bool)`

SetForeignConfigPresent sets ForeignConfigPresent field to given value.

### HasForeignConfigPresent

`func (o *StorageControllerAllOf) HasForeignConfigPresent() bool`

HasForeignConfigPresent returns a boolean if a field has been set.

### GetHwRevision

`func (o *StorageControllerAllOf) GetHwRevision() string`

GetHwRevision returns the HwRevision field if non-nil, zero value otherwise.

### GetHwRevisionOk

`func (o *StorageControllerAllOf) GetHwRevisionOk() (*string, bool)`

GetHwRevisionOk returns a tuple with the HwRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHwRevision

`func (o *StorageControllerAllOf) SetHwRevision(v string)`

SetHwRevision sets HwRevision field to given value.

### HasHwRevision

`func (o *StorageControllerAllOf) HasHwRevision() bool`

HasHwRevision returns a boolean if a field has been set.

### GetInterfaceType

`func (o *StorageControllerAllOf) GetInterfaceType() string`

GetInterfaceType returns the InterfaceType field if non-nil, zero value otherwise.

### GetInterfaceTypeOk

`func (o *StorageControllerAllOf) GetInterfaceTypeOk() (*string, bool)`

GetInterfaceTypeOk returns a tuple with the InterfaceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceType

`func (o *StorageControllerAllOf) SetInterfaceType(v string)`

SetInterfaceType sets InterfaceType field to given value.

### HasInterfaceType

`func (o *StorageControllerAllOf) HasInterfaceType() bool`

HasInterfaceType returns a boolean if a field has been set.

### GetMaxVolumesSupported

`func (o *StorageControllerAllOf) GetMaxVolumesSupported() int64`

GetMaxVolumesSupported returns the MaxVolumesSupported field if non-nil, zero value otherwise.

### GetMaxVolumesSupportedOk

`func (o *StorageControllerAllOf) GetMaxVolumesSupportedOk() (*int64, bool)`

GetMaxVolumesSupportedOk returns a tuple with the MaxVolumesSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVolumesSupported

`func (o *StorageControllerAllOf) SetMaxVolumesSupported(v int64)`

SetMaxVolumesSupported sets MaxVolumesSupported field to given value.

### HasMaxVolumesSupported

`func (o *StorageControllerAllOf) HasMaxVolumesSupported() bool`

HasMaxVolumesSupported returns a boolean if a field has been set.

### GetOobInterfaceSupported

`func (o *StorageControllerAllOf) GetOobInterfaceSupported() string`

GetOobInterfaceSupported returns the OobInterfaceSupported field if non-nil, zero value otherwise.

### GetOobInterfaceSupportedOk

`func (o *StorageControllerAllOf) GetOobInterfaceSupportedOk() (*string, bool)`

GetOobInterfaceSupportedOk returns a tuple with the OobInterfaceSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOobInterfaceSupported

`func (o *StorageControllerAllOf) SetOobInterfaceSupported(v string)`

SetOobInterfaceSupported sets OobInterfaceSupported field to given value.

### HasOobInterfaceSupported

`func (o *StorageControllerAllOf) HasOobInterfaceSupported() bool`

HasOobInterfaceSupported returns a boolean if a field has been set.

### GetOperState

`func (o *StorageControllerAllOf) GetOperState() string`

GetOperState returns the OperState field if non-nil, zero value otherwise.

### GetOperStateOk

`func (o *StorageControllerAllOf) GetOperStateOk() (*string, bool)`

GetOperStateOk returns a tuple with the OperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperState

`func (o *StorageControllerAllOf) SetOperState(v string)`

SetOperState sets OperState field to given value.

### HasOperState

`func (o *StorageControllerAllOf) HasOperState() bool`

HasOperState returns a boolean if a field has been set.

### GetOperability

`func (o *StorageControllerAllOf) GetOperability() string`

GetOperability returns the Operability field if non-nil, zero value otherwise.

### GetOperabilityOk

`func (o *StorageControllerAllOf) GetOperabilityOk() (*string, bool)`

GetOperabilityOk returns a tuple with the Operability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperability

`func (o *StorageControllerAllOf) SetOperability(v string)`

SetOperability sets Operability field to given value.

### HasOperability

`func (o *StorageControllerAllOf) HasOperability() bool`

HasOperability returns a boolean if a field has been set.

### GetPciAddr

`func (o *StorageControllerAllOf) GetPciAddr() string`

GetPciAddr returns the PciAddr field if non-nil, zero value otherwise.

### GetPciAddrOk

`func (o *StorageControllerAllOf) GetPciAddrOk() (*string, bool)`

GetPciAddrOk returns a tuple with the PciAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciAddr

`func (o *StorageControllerAllOf) SetPciAddr(v string)`

SetPciAddr sets PciAddr field to given value.

### HasPciAddr

`func (o *StorageControllerAllOf) HasPciAddr() bool`

HasPciAddr returns a boolean if a field has been set.

### GetPciSlot

`func (o *StorageControllerAllOf) GetPciSlot() string`

GetPciSlot returns the PciSlot field if non-nil, zero value otherwise.

### GetPciSlotOk

`func (o *StorageControllerAllOf) GetPciSlotOk() (*string, bool)`

GetPciSlotOk returns a tuple with the PciSlot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciSlot

`func (o *StorageControllerAllOf) SetPciSlot(v string)`

SetPciSlot sets PciSlot field to given value.

### HasPciSlot

`func (o *StorageControllerAllOf) HasPciSlot() bool`

HasPciSlot returns a boolean if a field has been set.

### GetPresence

`func (o *StorageControllerAllOf) GetPresence() string`

GetPresence returns the Presence field if non-nil, zero value otherwise.

### GetPresenceOk

`func (o *StorageControllerAllOf) GetPresenceOk() (*string, bool)`

GetPresenceOk returns a tuple with the Presence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresence

`func (o *StorageControllerAllOf) SetPresence(v string)`

SetPresence sets Presence field to given value.

### HasPresence

`func (o *StorageControllerAllOf) HasPresence() bool`

HasPresence returns a boolean if a field has been set.

### GetRaidSupport

`func (o *StorageControllerAllOf) GetRaidSupport() string`

GetRaidSupport returns the RaidSupport field if non-nil, zero value otherwise.

### GetRaidSupportOk

`func (o *StorageControllerAllOf) GetRaidSupportOk() (*string, bool)`

GetRaidSupportOk returns a tuple with the RaidSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaidSupport

`func (o *StorageControllerAllOf) SetRaidSupport(v string)`

SetRaidSupport sets RaidSupport field to given value.

### HasRaidSupport

`func (o *StorageControllerAllOf) HasRaidSupport() bool`

HasRaidSupport returns a boolean if a field has been set.

### GetRebuildRate

`func (o *StorageControllerAllOf) GetRebuildRate() string`

GetRebuildRate returns the RebuildRate field if non-nil, zero value otherwise.

### GetRebuildRateOk

`func (o *StorageControllerAllOf) GetRebuildRateOk() (*string, bool)`

GetRebuildRateOk returns a tuple with the RebuildRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRebuildRate

`func (o *StorageControllerAllOf) SetRebuildRate(v string)`

SetRebuildRate sets RebuildRate field to given value.

### HasRebuildRate

`func (o *StorageControllerAllOf) HasRebuildRate() bool`

HasRebuildRate returns a boolean if a field has been set.

### GetSelfEncryptEnabled

`func (o *StorageControllerAllOf) GetSelfEncryptEnabled() string`

GetSelfEncryptEnabled returns the SelfEncryptEnabled field if non-nil, zero value otherwise.

### GetSelfEncryptEnabledOk

`func (o *StorageControllerAllOf) GetSelfEncryptEnabledOk() (*string, bool)`

GetSelfEncryptEnabledOk returns a tuple with the SelfEncryptEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfEncryptEnabled

`func (o *StorageControllerAllOf) SetSelfEncryptEnabled(v string)`

SetSelfEncryptEnabled sets SelfEncryptEnabled field to given value.

### HasSelfEncryptEnabled

`func (o *StorageControllerAllOf) HasSelfEncryptEnabled() bool`

HasSelfEncryptEnabled returns a boolean if a field has been set.

### GetType

`func (o *StorageControllerAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StorageControllerAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StorageControllerAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *StorageControllerAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetComputeBlade

`func (o *StorageControllerAllOf) GetComputeBlade() ComputeBladeRelationship`

GetComputeBlade returns the ComputeBlade field if non-nil, zero value otherwise.

### GetComputeBladeOk

`func (o *StorageControllerAllOf) GetComputeBladeOk() (*ComputeBladeRelationship, bool)`

GetComputeBladeOk returns a tuple with the ComputeBlade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeBlade

`func (o *StorageControllerAllOf) SetComputeBlade(v ComputeBladeRelationship)`

SetComputeBlade sets ComputeBlade field to given value.

### HasComputeBlade

`func (o *StorageControllerAllOf) HasComputeBlade() bool`

HasComputeBlade returns a boolean if a field has been set.

### GetComputeBoard

`func (o *StorageControllerAllOf) GetComputeBoard() ComputeBoardRelationship`

GetComputeBoard returns the ComputeBoard field if non-nil, zero value otherwise.

### GetComputeBoardOk

`func (o *StorageControllerAllOf) GetComputeBoardOk() (*ComputeBoardRelationship, bool)`

GetComputeBoardOk returns a tuple with the ComputeBoard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeBoard

`func (o *StorageControllerAllOf) SetComputeBoard(v ComputeBoardRelationship)`

SetComputeBoard sets ComputeBoard field to given value.

### HasComputeBoard

`func (o *StorageControllerAllOf) HasComputeBoard() bool`

HasComputeBoard returns a boolean if a field has been set.

### GetComputeRackUnit

`func (o *StorageControllerAllOf) GetComputeRackUnit() ComputeRackUnitRelationship`

GetComputeRackUnit returns the ComputeRackUnit field if non-nil, zero value otherwise.

### GetComputeRackUnitOk

`func (o *StorageControllerAllOf) GetComputeRackUnitOk() (*ComputeRackUnitRelationship, bool)`

GetComputeRackUnitOk returns a tuple with the ComputeRackUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRackUnit

`func (o *StorageControllerAllOf) SetComputeRackUnit(v ComputeRackUnitRelationship)`

SetComputeRackUnit sets ComputeRackUnit field to given value.

### HasComputeRackUnit

`func (o *StorageControllerAllOf) HasComputeRackUnit() bool`

HasComputeRackUnit returns a boolean if a field has been set.

### GetDiskGroup

`func (o *StorageControllerAllOf) GetDiskGroup() []StorageDiskGroupRelationship`

GetDiskGroup returns the DiskGroup field if non-nil, zero value otherwise.

### GetDiskGroupOk

`func (o *StorageControllerAllOf) GetDiskGroupOk() (*[]StorageDiskGroupRelationship, bool)`

GetDiskGroupOk returns a tuple with the DiskGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskGroup

`func (o *StorageControllerAllOf) SetDiskGroup(v []StorageDiskGroupRelationship)`

SetDiskGroup sets DiskGroup field to given value.

### HasDiskGroup

`func (o *StorageControllerAllOf) HasDiskGroup() bool`

HasDiskGroup returns a boolean if a field has been set.

### SetDiskGroupNil

`func (o *StorageControllerAllOf) SetDiskGroupNil(b bool)`

 SetDiskGroupNil sets the value for DiskGroup to be an explicit nil

### UnsetDiskGroup
`func (o *StorageControllerAllOf) UnsetDiskGroup()`

UnsetDiskGroup ensures that no value is present for DiskGroup, not even an explicit nil
### GetInventoryDeviceInfo

`func (o *StorageControllerAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *StorageControllerAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *StorageControllerAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *StorageControllerAllOf) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetPhysicalDiskExtensions

`func (o *StorageControllerAllOf) GetPhysicalDiskExtensions() []StoragePhysicalDiskExtensionRelationship`

GetPhysicalDiskExtensions returns the PhysicalDiskExtensions field if non-nil, zero value otherwise.

### GetPhysicalDiskExtensionsOk

`func (o *StorageControllerAllOf) GetPhysicalDiskExtensionsOk() (*[]StoragePhysicalDiskExtensionRelationship, bool)`

GetPhysicalDiskExtensionsOk returns a tuple with the PhysicalDiskExtensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalDiskExtensions

`func (o *StorageControllerAllOf) SetPhysicalDiskExtensions(v []StoragePhysicalDiskExtensionRelationship)`

SetPhysicalDiskExtensions sets PhysicalDiskExtensions field to given value.

### HasPhysicalDiskExtensions

`func (o *StorageControllerAllOf) HasPhysicalDiskExtensions() bool`

HasPhysicalDiskExtensions returns a boolean if a field has been set.

### SetPhysicalDiskExtensionsNil

`func (o *StorageControllerAllOf) SetPhysicalDiskExtensionsNil(b bool)`

 SetPhysicalDiskExtensionsNil sets the value for PhysicalDiskExtensions to be an explicit nil

### UnsetPhysicalDiskExtensions
`func (o *StorageControllerAllOf) UnsetPhysicalDiskExtensions()`

UnsetPhysicalDiskExtensions ensures that no value is present for PhysicalDiskExtensions, not even an explicit nil
### GetPhysicalDisks

`func (o *StorageControllerAllOf) GetPhysicalDisks() []StoragePhysicalDiskRelationship`

GetPhysicalDisks returns the PhysicalDisks field if non-nil, zero value otherwise.

### GetPhysicalDisksOk

`func (o *StorageControllerAllOf) GetPhysicalDisksOk() (*[]StoragePhysicalDiskRelationship, bool)`

GetPhysicalDisksOk returns a tuple with the PhysicalDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalDisks

`func (o *StorageControllerAllOf) SetPhysicalDisks(v []StoragePhysicalDiskRelationship)`

SetPhysicalDisks sets PhysicalDisks field to given value.

### HasPhysicalDisks

`func (o *StorageControllerAllOf) HasPhysicalDisks() bool`

HasPhysicalDisks returns a boolean if a field has been set.

### SetPhysicalDisksNil

`func (o *StorageControllerAllOf) SetPhysicalDisksNil(b bool)`

 SetPhysicalDisksNil sets the value for PhysicalDisks to be an explicit nil

### UnsetPhysicalDisks
`func (o *StorageControllerAllOf) UnsetPhysicalDisks()`

UnsetPhysicalDisks ensures that no value is present for PhysicalDisks, not even an explicit nil
### GetRegisteredDevice

`func (o *StorageControllerAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StorageControllerAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StorageControllerAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StorageControllerAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### GetRunningFirmware

`func (o *StorageControllerAllOf) GetRunningFirmware() []FirmwareRunningFirmwareRelationship`

GetRunningFirmware returns the RunningFirmware field if non-nil, zero value otherwise.

### GetRunningFirmwareOk

`func (o *StorageControllerAllOf) GetRunningFirmwareOk() (*[]FirmwareRunningFirmwareRelationship, bool)`

GetRunningFirmwareOk returns a tuple with the RunningFirmware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningFirmware

`func (o *StorageControllerAllOf) SetRunningFirmware(v []FirmwareRunningFirmwareRelationship)`

SetRunningFirmware sets RunningFirmware field to given value.

### HasRunningFirmware

`func (o *StorageControllerAllOf) HasRunningFirmware() bool`

HasRunningFirmware returns a boolean if a field has been set.

### SetRunningFirmwareNil

`func (o *StorageControllerAllOf) SetRunningFirmwareNil(b bool)`

 SetRunningFirmwareNil sets the value for RunningFirmware to be an explicit nil

### UnsetRunningFirmware
`func (o *StorageControllerAllOf) UnsetRunningFirmware()`

UnsetRunningFirmware ensures that no value is present for RunningFirmware, not even an explicit nil
### GetVirtualDriveExtensions

`func (o *StorageControllerAllOf) GetVirtualDriveExtensions() []StorageVirtualDriveExtensionRelationship`

GetVirtualDriveExtensions returns the VirtualDriveExtensions field if non-nil, zero value otherwise.

### GetVirtualDriveExtensionsOk

`func (o *StorageControllerAllOf) GetVirtualDriveExtensionsOk() (*[]StorageVirtualDriveExtensionRelationship, bool)`

GetVirtualDriveExtensionsOk returns a tuple with the VirtualDriveExtensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualDriveExtensions

`func (o *StorageControllerAllOf) SetVirtualDriveExtensions(v []StorageVirtualDriveExtensionRelationship)`

SetVirtualDriveExtensions sets VirtualDriveExtensions field to given value.

### HasVirtualDriveExtensions

`func (o *StorageControllerAllOf) HasVirtualDriveExtensions() bool`

HasVirtualDriveExtensions returns a boolean if a field has been set.

### SetVirtualDriveExtensionsNil

`func (o *StorageControllerAllOf) SetVirtualDriveExtensionsNil(b bool)`

 SetVirtualDriveExtensionsNil sets the value for VirtualDriveExtensions to be an explicit nil

### UnsetVirtualDriveExtensions
`func (o *StorageControllerAllOf) UnsetVirtualDriveExtensions()`

UnsetVirtualDriveExtensions ensures that no value is present for VirtualDriveExtensions, not even an explicit nil
### GetVirtualDrives

`func (o *StorageControllerAllOf) GetVirtualDrives() []StorageVirtualDriveRelationship`

GetVirtualDrives returns the VirtualDrives field if non-nil, zero value otherwise.

### GetVirtualDrivesOk

`func (o *StorageControllerAllOf) GetVirtualDrivesOk() (*[]StorageVirtualDriveRelationship, bool)`

GetVirtualDrivesOk returns a tuple with the VirtualDrives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualDrives

`func (o *StorageControllerAllOf) SetVirtualDrives(v []StorageVirtualDriveRelationship)`

SetVirtualDrives sets VirtualDrives field to given value.

### HasVirtualDrives

`func (o *StorageControllerAllOf) HasVirtualDrives() bool`

HasVirtualDrives returns a boolean if a field has been set.

### SetVirtualDrivesNil

`func (o *StorageControllerAllOf) SetVirtualDrivesNil(b bool)`

 SetVirtualDrivesNil sets the value for VirtualDrives to be an explicit nil

### UnsetVirtualDrives
`func (o *StorageControllerAllOf) UnsetVirtualDrives()`

UnsetVirtualDrives ensures that no value is present for VirtualDrives, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



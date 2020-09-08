# StorageVirtualDrive

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessPolicy** | Pointer to **string** | The access policy of the virtual drive. | [optional] [readonly] 
**ActualWriteCachePolicy** | Pointer to **string** | The current write cache policy of the virtual drive. | [optional] [readonly] 
**AvailableSize** | Pointer to **string** | Available storage capacity of the virtual drive. | [optional] [readonly] 
**BlockSize** | Pointer to **string** | Block size of the virtual drive. | [optional] [readonly] 
**Bootable** | Pointer to **string** | The virtual drive bootable state. | [optional] [readonly] 
**ConfigState** | Pointer to **string** | The configuration state of the virtual drive. | [optional] [readonly] 
**ConfiguredWriteCachePolicy** | Pointer to **string** | The requested write cache policy of the virtual drive. | [optional] [readonly] 
**ConnectionProtocol** | Pointer to **string** | The connection protocol of the virtual drive. | [optional] [readonly] 
**DriveCache** | Pointer to **string** | The state of the drive cache of the virtual drive. | [optional] [readonly] 
**DriveSecurity** | Pointer to **string** | The driveSecurity state of the Virtual drive. | [optional] [readonly] 
**DriveState** | Pointer to **string** | The state of the Virtual drive. | [optional] [readonly] 
**IoPolicy** | Pointer to **string** | The Input/Output Policy defined on the Virtual drive. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the Virtual drive. | [optional] [readonly] 
**NumBlocks** | Pointer to **string** | Number of Blocks on the Physical Disk. | [optional] [readonly] 
**OperState** | Pointer to **string** | The current operational state of Virtual drive. | [optional] [readonly] 
**Operability** | Pointer to **string** | The current operability state of Virtual drive. | [optional] [readonly] 
**PhysicalBlockSize** | Pointer to **string** | The block size of the the virtual drive. | [optional] [readonly] 
**Presence** | Pointer to **string** | The presence status of the virtual drive. | [optional] [readonly] 
**ReadPolicy** | Pointer to **string** | The read-ahead cache mode of the virtual drive. | [optional] [readonly] 
**SecurityFlags** | Pointer to **string** | The security flags set for this virtual drive. | [optional] [readonly] 
**Size** | Pointer to **string** | The size of the virtual drive in MB. | [optional] [readonly] 
**StripSize** | Pointer to **string** | The strip size is the portion of a stripe that resides on a single drive in the drive group, this is measured in KB. | [optional] [readonly] 
**Type** | Pointer to **string** | The raid type of the virtual drive. | [optional] [readonly] 
**Uuid** | Pointer to **string** | The uuid of the virtual drive. | [optional] [readonly] 
**VendorUuid** | Pointer to **string** | The UUID value of the vendor. | [optional] [readonly] 
**VirtualDriveId** | Pointer to **string** | The identifier for this Virtual drive. | [optional] [readonly] 
**DiskGroup** | Pointer to [**StorageDiskGroupRelationship**](storage.DiskGroup.Relationship.md) |  | [optional] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](inventory.DeviceInfo.Relationship.md) |  | [optional] 
**PhysicalDiskUsages** | Pointer to [**[]StoragePhysicalDiskUsageRelationship**](storage.PhysicalDiskUsage.Relationship.md) | An array of relationships to storagePhysicalDiskUsage resources. | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 
**StorageController** | Pointer to [**StorageControllerRelationship**](storage.Controller.Relationship.md) |  | [optional] 
**VdMemberEps** | Pointer to [**[]StorageVdMemberEpRelationship**](storage.VdMemberEp.Relationship.md) | An array of relationships to storageVdMemberEp resources. | [optional] [readonly] 
**VirtualDriveExtension** | Pointer to [**StorageVirtualDriveExtensionRelationship**](storage.VirtualDriveExtension.Relationship.md) |  | [optional] 

## Methods

### NewStorageVirtualDrive

`func NewStorageVirtualDrive() *StorageVirtualDrive`

NewStorageVirtualDrive instantiates a new StorageVirtualDrive object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageVirtualDriveWithDefaults

`func NewStorageVirtualDriveWithDefaults() *StorageVirtualDrive`

NewStorageVirtualDriveWithDefaults instantiates a new StorageVirtualDrive object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessPolicy

`func (o *StorageVirtualDrive) GetAccessPolicy() string`

GetAccessPolicy returns the AccessPolicy field if non-nil, zero value otherwise.

### GetAccessPolicyOk

`func (o *StorageVirtualDrive) GetAccessPolicyOk() (*string, bool)`

GetAccessPolicyOk returns a tuple with the AccessPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessPolicy

`func (o *StorageVirtualDrive) SetAccessPolicy(v string)`

SetAccessPolicy sets AccessPolicy field to given value.

### HasAccessPolicy

`func (o *StorageVirtualDrive) HasAccessPolicy() bool`

HasAccessPolicy returns a boolean if a field has been set.

### GetActualWriteCachePolicy

`func (o *StorageVirtualDrive) GetActualWriteCachePolicy() string`

GetActualWriteCachePolicy returns the ActualWriteCachePolicy field if non-nil, zero value otherwise.

### GetActualWriteCachePolicyOk

`func (o *StorageVirtualDrive) GetActualWriteCachePolicyOk() (*string, bool)`

GetActualWriteCachePolicyOk returns a tuple with the ActualWriteCachePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualWriteCachePolicy

`func (o *StorageVirtualDrive) SetActualWriteCachePolicy(v string)`

SetActualWriteCachePolicy sets ActualWriteCachePolicy field to given value.

### HasActualWriteCachePolicy

`func (o *StorageVirtualDrive) HasActualWriteCachePolicy() bool`

HasActualWriteCachePolicy returns a boolean if a field has been set.

### GetAvailableSize

`func (o *StorageVirtualDrive) GetAvailableSize() string`

GetAvailableSize returns the AvailableSize field if non-nil, zero value otherwise.

### GetAvailableSizeOk

`func (o *StorageVirtualDrive) GetAvailableSizeOk() (*string, bool)`

GetAvailableSizeOk returns a tuple with the AvailableSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableSize

`func (o *StorageVirtualDrive) SetAvailableSize(v string)`

SetAvailableSize sets AvailableSize field to given value.

### HasAvailableSize

`func (o *StorageVirtualDrive) HasAvailableSize() bool`

HasAvailableSize returns a boolean if a field has been set.

### GetBlockSize

`func (o *StorageVirtualDrive) GetBlockSize() string`

GetBlockSize returns the BlockSize field if non-nil, zero value otherwise.

### GetBlockSizeOk

`func (o *StorageVirtualDrive) GetBlockSizeOk() (*string, bool)`

GetBlockSizeOk returns a tuple with the BlockSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockSize

`func (o *StorageVirtualDrive) SetBlockSize(v string)`

SetBlockSize sets BlockSize field to given value.

### HasBlockSize

`func (o *StorageVirtualDrive) HasBlockSize() bool`

HasBlockSize returns a boolean if a field has been set.

### GetBootable

`func (o *StorageVirtualDrive) GetBootable() string`

GetBootable returns the Bootable field if non-nil, zero value otherwise.

### GetBootableOk

`func (o *StorageVirtualDrive) GetBootableOk() (*string, bool)`

GetBootableOk returns a tuple with the Bootable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootable

`func (o *StorageVirtualDrive) SetBootable(v string)`

SetBootable sets Bootable field to given value.

### HasBootable

`func (o *StorageVirtualDrive) HasBootable() bool`

HasBootable returns a boolean if a field has been set.

### GetConfigState

`func (o *StorageVirtualDrive) GetConfigState() string`

GetConfigState returns the ConfigState field if non-nil, zero value otherwise.

### GetConfigStateOk

`func (o *StorageVirtualDrive) GetConfigStateOk() (*string, bool)`

GetConfigStateOk returns a tuple with the ConfigState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigState

`func (o *StorageVirtualDrive) SetConfigState(v string)`

SetConfigState sets ConfigState field to given value.

### HasConfigState

`func (o *StorageVirtualDrive) HasConfigState() bool`

HasConfigState returns a boolean if a field has been set.

### GetConfiguredWriteCachePolicy

`func (o *StorageVirtualDrive) GetConfiguredWriteCachePolicy() string`

GetConfiguredWriteCachePolicy returns the ConfiguredWriteCachePolicy field if non-nil, zero value otherwise.

### GetConfiguredWriteCachePolicyOk

`func (o *StorageVirtualDrive) GetConfiguredWriteCachePolicyOk() (*string, bool)`

GetConfiguredWriteCachePolicyOk returns a tuple with the ConfiguredWriteCachePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguredWriteCachePolicy

`func (o *StorageVirtualDrive) SetConfiguredWriteCachePolicy(v string)`

SetConfiguredWriteCachePolicy sets ConfiguredWriteCachePolicy field to given value.

### HasConfiguredWriteCachePolicy

`func (o *StorageVirtualDrive) HasConfiguredWriteCachePolicy() bool`

HasConfiguredWriteCachePolicy returns a boolean if a field has been set.

### GetConnectionProtocol

`func (o *StorageVirtualDrive) GetConnectionProtocol() string`

GetConnectionProtocol returns the ConnectionProtocol field if non-nil, zero value otherwise.

### GetConnectionProtocolOk

`func (o *StorageVirtualDrive) GetConnectionProtocolOk() (*string, bool)`

GetConnectionProtocolOk returns a tuple with the ConnectionProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionProtocol

`func (o *StorageVirtualDrive) SetConnectionProtocol(v string)`

SetConnectionProtocol sets ConnectionProtocol field to given value.

### HasConnectionProtocol

`func (o *StorageVirtualDrive) HasConnectionProtocol() bool`

HasConnectionProtocol returns a boolean if a field has been set.

### GetDriveCache

`func (o *StorageVirtualDrive) GetDriveCache() string`

GetDriveCache returns the DriveCache field if non-nil, zero value otherwise.

### GetDriveCacheOk

`func (o *StorageVirtualDrive) GetDriveCacheOk() (*string, bool)`

GetDriveCacheOk returns a tuple with the DriveCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriveCache

`func (o *StorageVirtualDrive) SetDriveCache(v string)`

SetDriveCache sets DriveCache field to given value.

### HasDriveCache

`func (o *StorageVirtualDrive) HasDriveCache() bool`

HasDriveCache returns a boolean if a field has been set.

### GetDriveSecurity

`func (o *StorageVirtualDrive) GetDriveSecurity() string`

GetDriveSecurity returns the DriveSecurity field if non-nil, zero value otherwise.

### GetDriveSecurityOk

`func (o *StorageVirtualDrive) GetDriveSecurityOk() (*string, bool)`

GetDriveSecurityOk returns a tuple with the DriveSecurity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriveSecurity

`func (o *StorageVirtualDrive) SetDriveSecurity(v string)`

SetDriveSecurity sets DriveSecurity field to given value.

### HasDriveSecurity

`func (o *StorageVirtualDrive) HasDriveSecurity() bool`

HasDriveSecurity returns a boolean if a field has been set.

### GetDriveState

`func (o *StorageVirtualDrive) GetDriveState() string`

GetDriveState returns the DriveState field if non-nil, zero value otherwise.

### GetDriveStateOk

`func (o *StorageVirtualDrive) GetDriveStateOk() (*string, bool)`

GetDriveStateOk returns a tuple with the DriveState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriveState

`func (o *StorageVirtualDrive) SetDriveState(v string)`

SetDriveState sets DriveState field to given value.

### HasDriveState

`func (o *StorageVirtualDrive) HasDriveState() bool`

HasDriveState returns a boolean if a field has been set.

### GetIoPolicy

`func (o *StorageVirtualDrive) GetIoPolicy() string`

GetIoPolicy returns the IoPolicy field if non-nil, zero value otherwise.

### GetIoPolicyOk

`func (o *StorageVirtualDrive) GetIoPolicyOk() (*string, bool)`

GetIoPolicyOk returns a tuple with the IoPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIoPolicy

`func (o *StorageVirtualDrive) SetIoPolicy(v string)`

SetIoPolicy sets IoPolicy field to given value.

### HasIoPolicy

`func (o *StorageVirtualDrive) HasIoPolicy() bool`

HasIoPolicy returns a boolean if a field has been set.

### GetName

`func (o *StorageVirtualDrive) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageVirtualDrive) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageVirtualDrive) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StorageVirtualDrive) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumBlocks

`func (o *StorageVirtualDrive) GetNumBlocks() string`

GetNumBlocks returns the NumBlocks field if non-nil, zero value otherwise.

### GetNumBlocksOk

`func (o *StorageVirtualDrive) GetNumBlocksOk() (*string, bool)`

GetNumBlocksOk returns a tuple with the NumBlocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumBlocks

`func (o *StorageVirtualDrive) SetNumBlocks(v string)`

SetNumBlocks sets NumBlocks field to given value.

### HasNumBlocks

`func (o *StorageVirtualDrive) HasNumBlocks() bool`

HasNumBlocks returns a boolean if a field has been set.

### GetOperState

`func (o *StorageVirtualDrive) GetOperState() string`

GetOperState returns the OperState field if non-nil, zero value otherwise.

### GetOperStateOk

`func (o *StorageVirtualDrive) GetOperStateOk() (*string, bool)`

GetOperStateOk returns a tuple with the OperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperState

`func (o *StorageVirtualDrive) SetOperState(v string)`

SetOperState sets OperState field to given value.

### HasOperState

`func (o *StorageVirtualDrive) HasOperState() bool`

HasOperState returns a boolean if a field has been set.

### GetOperability

`func (o *StorageVirtualDrive) GetOperability() string`

GetOperability returns the Operability field if non-nil, zero value otherwise.

### GetOperabilityOk

`func (o *StorageVirtualDrive) GetOperabilityOk() (*string, bool)`

GetOperabilityOk returns a tuple with the Operability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperability

`func (o *StorageVirtualDrive) SetOperability(v string)`

SetOperability sets Operability field to given value.

### HasOperability

`func (o *StorageVirtualDrive) HasOperability() bool`

HasOperability returns a boolean if a field has been set.

### GetPhysicalBlockSize

`func (o *StorageVirtualDrive) GetPhysicalBlockSize() string`

GetPhysicalBlockSize returns the PhysicalBlockSize field if non-nil, zero value otherwise.

### GetPhysicalBlockSizeOk

`func (o *StorageVirtualDrive) GetPhysicalBlockSizeOk() (*string, bool)`

GetPhysicalBlockSizeOk returns a tuple with the PhysicalBlockSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalBlockSize

`func (o *StorageVirtualDrive) SetPhysicalBlockSize(v string)`

SetPhysicalBlockSize sets PhysicalBlockSize field to given value.

### HasPhysicalBlockSize

`func (o *StorageVirtualDrive) HasPhysicalBlockSize() bool`

HasPhysicalBlockSize returns a boolean if a field has been set.

### GetPresence

`func (o *StorageVirtualDrive) GetPresence() string`

GetPresence returns the Presence field if non-nil, zero value otherwise.

### GetPresenceOk

`func (o *StorageVirtualDrive) GetPresenceOk() (*string, bool)`

GetPresenceOk returns a tuple with the Presence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresence

`func (o *StorageVirtualDrive) SetPresence(v string)`

SetPresence sets Presence field to given value.

### HasPresence

`func (o *StorageVirtualDrive) HasPresence() bool`

HasPresence returns a boolean if a field has been set.

### GetReadPolicy

`func (o *StorageVirtualDrive) GetReadPolicy() string`

GetReadPolicy returns the ReadPolicy field if non-nil, zero value otherwise.

### GetReadPolicyOk

`func (o *StorageVirtualDrive) GetReadPolicyOk() (*string, bool)`

GetReadPolicyOk returns a tuple with the ReadPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadPolicy

`func (o *StorageVirtualDrive) SetReadPolicy(v string)`

SetReadPolicy sets ReadPolicy field to given value.

### HasReadPolicy

`func (o *StorageVirtualDrive) HasReadPolicy() bool`

HasReadPolicy returns a boolean if a field has been set.

### GetSecurityFlags

`func (o *StorageVirtualDrive) GetSecurityFlags() string`

GetSecurityFlags returns the SecurityFlags field if non-nil, zero value otherwise.

### GetSecurityFlagsOk

`func (o *StorageVirtualDrive) GetSecurityFlagsOk() (*string, bool)`

GetSecurityFlagsOk returns a tuple with the SecurityFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityFlags

`func (o *StorageVirtualDrive) SetSecurityFlags(v string)`

SetSecurityFlags sets SecurityFlags field to given value.

### HasSecurityFlags

`func (o *StorageVirtualDrive) HasSecurityFlags() bool`

HasSecurityFlags returns a boolean if a field has been set.

### GetSize

`func (o *StorageVirtualDrive) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *StorageVirtualDrive) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *StorageVirtualDrive) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *StorageVirtualDrive) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetStripSize

`func (o *StorageVirtualDrive) GetStripSize() string`

GetStripSize returns the StripSize field if non-nil, zero value otherwise.

### GetStripSizeOk

`func (o *StorageVirtualDrive) GetStripSizeOk() (*string, bool)`

GetStripSizeOk returns a tuple with the StripSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripSize

`func (o *StorageVirtualDrive) SetStripSize(v string)`

SetStripSize sets StripSize field to given value.

### HasStripSize

`func (o *StorageVirtualDrive) HasStripSize() bool`

HasStripSize returns a boolean if a field has been set.

### GetType

`func (o *StorageVirtualDrive) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StorageVirtualDrive) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StorageVirtualDrive) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *StorageVirtualDrive) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUuid

`func (o *StorageVirtualDrive) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *StorageVirtualDrive) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *StorageVirtualDrive) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *StorageVirtualDrive) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetVendorUuid

`func (o *StorageVirtualDrive) GetVendorUuid() string`

GetVendorUuid returns the VendorUuid field if non-nil, zero value otherwise.

### GetVendorUuidOk

`func (o *StorageVirtualDrive) GetVendorUuidOk() (*string, bool)`

GetVendorUuidOk returns a tuple with the VendorUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorUuid

`func (o *StorageVirtualDrive) SetVendorUuid(v string)`

SetVendorUuid sets VendorUuid field to given value.

### HasVendorUuid

`func (o *StorageVirtualDrive) HasVendorUuid() bool`

HasVendorUuid returns a boolean if a field has been set.

### GetVirtualDriveId

`func (o *StorageVirtualDrive) GetVirtualDriveId() string`

GetVirtualDriveId returns the VirtualDriveId field if non-nil, zero value otherwise.

### GetVirtualDriveIdOk

`func (o *StorageVirtualDrive) GetVirtualDriveIdOk() (*string, bool)`

GetVirtualDriveIdOk returns a tuple with the VirtualDriveId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualDriveId

`func (o *StorageVirtualDrive) SetVirtualDriveId(v string)`

SetVirtualDriveId sets VirtualDriveId field to given value.

### HasVirtualDriveId

`func (o *StorageVirtualDrive) HasVirtualDriveId() bool`

HasVirtualDriveId returns a boolean if a field has been set.

### GetDiskGroup

`func (o *StorageVirtualDrive) GetDiskGroup() StorageDiskGroupRelationship`

GetDiskGroup returns the DiskGroup field if non-nil, zero value otherwise.

### GetDiskGroupOk

`func (o *StorageVirtualDrive) GetDiskGroupOk() (*StorageDiskGroupRelationship, bool)`

GetDiskGroupOk returns a tuple with the DiskGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskGroup

`func (o *StorageVirtualDrive) SetDiskGroup(v StorageDiskGroupRelationship)`

SetDiskGroup sets DiskGroup field to given value.

### HasDiskGroup

`func (o *StorageVirtualDrive) HasDiskGroup() bool`

HasDiskGroup returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *StorageVirtualDrive) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *StorageVirtualDrive) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *StorageVirtualDrive) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *StorageVirtualDrive) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetPhysicalDiskUsages

`func (o *StorageVirtualDrive) GetPhysicalDiskUsages() []StoragePhysicalDiskUsageRelationship`

GetPhysicalDiskUsages returns the PhysicalDiskUsages field if non-nil, zero value otherwise.

### GetPhysicalDiskUsagesOk

`func (o *StorageVirtualDrive) GetPhysicalDiskUsagesOk() (*[]StoragePhysicalDiskUsageRelationship, bool)`

GetPhysicalDiskUsagesOk returns a tuple with the PhysicalDiskUsages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalDiskUsages

`func (o *StorageVirtualDrive) SetPhysicalDiskUsages(v []StoragePhysicalDiskUsageRelationship)`

SetPhysicalDiskUsages sets PhysicalDiskUsages field to given value.

### HasPhysicalDiskUsages

`func (o *StorageVirtualDrive) HasPhysicalDiskUsages() bool`

HasPhysicalDiskUsages returns a boolean if a field has been set.

### SetPhysicalDiskUsagesNil

`func (o *StorageVirtualDrive) SetPhysicalDiskUsagesNil(b bool)`

 SetPhysicalDiskUsagesNil sets the value for PhysicalDiskUsages to be an explicit nil

### UnsetPhysicalDiskUsages
`func (o *StorageVirtualDrive) UnsetPhysicalDiskUsages()`

UnsetPhysicalDiskUsages ensures that no value is present for PhysicalDiskUsages, not even an explicit nil
### GetRegisteredDevice

`func (o *StorageVirtualDrive) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StorageVirtualDrive) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StorageVirtualDrive) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StorageVirtualDrive) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### GetStorageController

`func (o *StorageVirtualDrive) GetStorageController() StorageControllerRelationship`

GetStorageController returns the StorageController field if non-nil, zero value otherwise.

### GetStorageControllerOk

`func (o *StorageVirtualDrive) GetStorageControllerOk() (*StorageControllerRelationship, bool)`

GetStorageControllerOk returns a tuple with the StorageController field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageController

`func (o *StorageVirtualDrive) SetStorageController(v StorageControllerRelationship)`

SetStorageController sets StorageController field to given value.

### HasStorageController

`func (o *StorageVirtualDrive) HasStorageController() bool`

HasStorageController returns a boolean if a field has been set.

### GetVdMemberEps

`func (o *StorageVirtualDrive) GetVdMemberEps() []StorageVdMemberEpRelationship`

GetVdMemberEps returns the VdMemberEps field if non-nil, zero value otherwise.

### GetVdMemberEpsOk

`func (o *StorageVirtualDrive) GetVdMemberEpsOk() (*[]StorageVdMemberEpRelationship, bool)`

GetVdMemberEpsOk returns a tuple with the VdMemberEps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdMemberEps

`func (o *StorageVirtualDrive) SetVdMemberEps(v []StorageVdMemberEpRelationship)`

SetVdMemberEps sets VdMemberEps field to given value.

### HasVdMemberEps

`func (o *StorageVirtualDrive) HasVdMemberEps() bool`

HasVdMemberEps returns a boolean if a field has been set.

### SetVdMemberEpsNil

`func (o *StorageVirtualDrive) SetVdMemberEpsNil(b bool)`

 SetVdMemberEpsNil sets the value for VdMemberEps to be an explicit nil

### UnsetVdMemberEps
`func (o *StorageVirtualDrive) UnsetVdMemberEps()`

UnsetVdMemberEps ensures that no value is present for VdMemberEps, not even an explicit nil
### GetVirtualDriveExtension

`func (o *StorageVirtualDrive) GetVirtualDriveExtension() StorageVirtualDriveExtensionRelationship`

GetVirtualDriveExtension returns the VirtualDriveExtension field if non-nil, zero value otherwise.

### GetVirtualDriveExtensionOk

`func (o *StorageVirtualDrive) GetVirtualDriveExtensionOk() (*StorageVirtualDriveExtensionRelationship, bool)`

GetVirtualDriveExtensionOk returns a tuple with the VirtualDriveExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualDriveExtension

`func (o *StorageVirtualDrive) SetVirtualDriveExtension(v StorageVirtualDriveExtensionRelationship)`

SetVirtualDriveExtension sets VirtualDriveExtension field to given value.

### HasVirtualDriveExtension

`func (o *StorageVirtualDrive) HasVirtualDriveExtension() bool`

HasVirtualDriveExtension returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



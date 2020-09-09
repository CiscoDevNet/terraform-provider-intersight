# StorageVirtualDriveAllOf

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

### NewStorageVirtualDriveAllOf

`func NewStorageVirtualDriveAllOf() *StorageVirtualDriveAllOf`

NewStorageVirtualDriveAllOf instantiates a new StorageVirtualDriveAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageVirtualDriveAllOfWithDefaults

`func NewStorageVirtualDriveAllOfWithDefaults() *StorageVirtualDriveAllOf`

NewStorageVirtualDriveAllOfWithDefaults instantiates a new StorageVirtualDriveAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessPolicy

`func (o *StorageVirtualDriveAllOf) GetAccessPolicy() string`

GetAccessPolicy returns the AccessPolicy field if non-nil, zero value otherwise.

### GetAccessPolicyOk

`func (o *StorageVirtualDriveAllOf) GetAccessPolicyOk() (*string, bool)`

GetAccessPolicyOk returns a tuple with the AccessPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessPolicy

`func (o *StorageVirtualDriveAllOf) SetAccessPolicy(v string)`

SetAccessPolicy sets AccessPolicy field to given value.

### HasAccessPolicy

`func (o *StorageVirtualDriveAllOf) HasAccessPolicy() bool`

HasAccessPolicy returns a boolean if a field has been set.

### GetActualWriteCachePolicy

`func (o *StorageVirtualDriveAllOf) GetActualWriteCachePolicy() string`

GetActualWriteCachePolicy returns the ActualWriteCachePolicy field if non-nil, zero value otherwise.

### GetActualWriteCachePolicyOk

`func (o *StorageVirtualDriveAllOf) GetActualWriteCachePolicyOk() (*string, bool)`

GetActualWriteCachePolicyOk returns a tuple with the ActualWriteCachePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualWriteCachePolicy

`func (o *StorageVirtualDriveAllOf) SetActualWriteCachePolicy(v string)`

SetActualWriteCachePolicy sets ActualWriteCachePolicy field to given value.

### HasActualWriteCachePolicy

`func (o *StorageVirtualDriveAllOf) HasActualWriteCachePolicy() bool`

HasActualWriteCachePolicy returns a boolean if a field has been set.

### GetAvailableSize

`func (o *StorageVirtualDriveAllOf) GetAvailableSize() string`

GetAvailableSize returns the AvailableSize field if non-nil, zero value otherwise.

### GetAvailableSizeOk

`func (o *StorageVirtualDriveAllOf) GetAvailableSizeOk() (*string, bool)`

GetAvailableSizeOk returns a tuple with the AvailableSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableSize

`func (o *StorageVirtualDriveAllOf) SetAvailableSize(v string)`

SetAvailableSize sets AvailableSize field to given value.

### HasAvailableSize

`func (o *StorageVirtualDriveAllOf) HasAvailableSize() bool`

HasAvailableSize returns a boolean if a field has been set.

### GetBlockSize

`func (o *StorageVirtualDriveAllOf) GetBlockSize() string`

GetBlockSize returns the BlockSize field if non-nil, zero value otherwise.

### GetBlockSizeOk

`func (o *StorageVirtualDriveAllOf) GetBlockSizeOk() (*string, bool)`

GetBlockSizeOk returns a tuple with the BlockSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockSize

`func (o *StorageVirtualDriveAllOf) SetBlockSize(v string)`

SetBlockSize sets BlockSize field to given value.

### HasBlockSize

`func (o *StorageVirtualDriveAllOf) HasBlockSize() bool`

HasBlockSize returns a boolean if a field has been set.

### GetBootable

`func (o *StorageVirtualDriveAllOf) GetBootable() string`

GetBootable returns the Bootable field if non-nil, zero value otherwise.

### GetBootableOk

`func (o *StorageVirtualDriveAllOf) GetBootableOk() (*string, bool)`

GetBootableOk returns a tuple with the Bootable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootable

`func (o *StorageVirtualDriveAllOf) SetBootable(v string)`

SetBootable sets Bootable field to given value.

### HasBootable

`func (o *StorageVirtualDriveAllOf) HasBootable() bool`

HasBootable returns a boolean if a field has been set.

### GetConfigState

`func (o *StorageVirtualDriveAllOf) GetConfigState() string`

GetConfigState returns the ConfigState field if non-nil, zero value otherwise.

### GetConfigStateOk

`func (o *StorageVirtualDriveAllOf) GetConfigStateOk() (*string, bool)`

GetConfigStateOk returns a tuple with the ConfigState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigState

`func (o *StorageVirtualDriveAllOf) SetConfigState(v string)`

SetConfigState sets ConfigState field to given value.

### HasConfigState

`func (o *StorageVirtualDriveAllOf) HasConfigState() bool`

HasConfigState returns a boolean if a field has been set.

### GetConfiguredWriteCachePolicy

`func (o *StorageVirtualDriveAllOf) GetConfiguredWriteCachePolicy() string`

GetConfiguredWriteCachePolicy returns the ConfiguredWriteCachePolicy field if non-nil, zero value otherwise.

### GetConfiguredWriteCachePolicyOk

`func (o *StorageVirtualDriveAllOf) GetConfiguredWriteCachePolicyOk() (*string, bool)`

GetConfiguredWriteCachePolicyOk returns a tuple with the ConfiguredWriteCachePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguredWriteCachePolicy

`func (o *StorageVirtualDriveAllOf) SetConfiguredWriteCachePolicy(v string)`

SetConfiguredWriteCachePolicy sets ConfiguredWriteCachePolicy field to given value.

### HasConfiguredWriteCachePolicy

`func (o *StorageVirtualDriveAllOf) HasConfiguredWriteCachePolicy() bool`

HasConfiguredWriteCachePolicy returns a boolean if a field has been set.

### GetConnectionProtocol

`func (o *StorageVirtualDriveAllOf) GetConnectionProtocol() string`

GetConnectionProtocol returns the ConnectionProtocol field if non-nil, zero value otherwise.

### GetConnectionProtocolOk

`func (o *StorageVirtualDriveAllOf) GetConnectionProtocolOk() (*string, bool)`

GetConnectionProtocolOk returns a tuple with the ConnectionProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionProtocol

`func (o *StorageVirtualDriveAllOf) SetConnectionProtocol(v string)`

SetConnectionProtocol sets ConnectionProtocol field to given value.

### HasConnectionProtocol

`func (o *StorageVirtualDriveAllOf) HasConnectionProtocol() bool`

HasConnectionProtocol returns a boolean if a field has been set.

### GetDriveCache

`func (o *StorageVirtualDriveAllOf) GetDriveCache() string`

GetDriveCache returns the DriveCache field if non-nil, zero value otherwise.

### GetDriveCacheOk

`func (o *StorageVirtualDriveAllOf) GetDriveCacheOk() (*string, bool)`

GetDriveCacheOk returns a tuple with the DriveCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriveCache

`func (o *StorageVirtualDriveAllOf) SetDriveCache(v string)`

SetDriveCache sets DriveCache field to given value.

### HasDriveCache

`func (o *StorageVirtualDriveAllOf) HasDriveCache() bool`

HasDriveCache returns a boolean if a field has been set.

### GetDriveSecurity

`func (o *StorageVirtualDriveAllOf) GetDriveSecurity() string`

GetDriveSecurity returns the DriveSecurity field if non-nil, zero value otherwise.

### GetDriveSecurityOk

`func (o *StorageVirtualDriveAllOf) GetDriveSecurityOk() (*string, bool)`

GetDriveSecurityOk returns a tuple with the DriveSecurity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriveSecurity

`func (o *StorageVirtualDriveAllOf) SetDriveSecurity(v string)`

SetDriveSecurity sets DriveSecurity field to given value.

### HasDriveSecurity

`func (o *StorageVirtualDriveAllOf) HasDriveSecurity() bool`

HasDriveSecurity returns a boolean if a field has been set.

### GetDriveState

`func (o *StorageVirtualDriveAllOf) GetDriveState() string`

GetDriveState returns the DriveState field if non-nil, zero value otherwise.

### GetDriveStateOk

`func (o *StorageVirtualDriveAllOf) GetDriveStateOk() (*string, bool)`

GetDriveStateOk returns a tuple with the DriveState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriveState

`func (o *StorageVirtualDriveAllOf) SetDriveState(v string)`

SetDriveState sets DriveState field to given value.

### HasDriveState

`func (o *StorageVirtualDriveAllOf) HasDriveState() bool`

HasDriveState returns a boolean if a field has been set.

### GetIoPolicy

`func (o *StorageVirtualDriveAllOf) GetIoPolicy() string`

GetIoPolicy returns the IoPolicy field if non-nil, zero value otherwise.

### GetIoPolicyOk

`func (o *StorageVirtualDriveAllOf) GetIoPolicyOk() (*string, bool)`

GetIoPolicyOk returns a tuple with the IoPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIoPolicy

`func (o *StorageVirtualDriveAllOf) SetIoPolicy(v string)`

SetIoPolicy sets IoPolicy field to given value.

### HasIoPolicy

`func (o *StorageVirtualDriveAllOf) HasIoPolicy() bool`

HasIoPolicy returns a boolean if a field has been set.

### GetName

`func (o *StorageVirtualDriveAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageVirtualDriveAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageVirtualDriveAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StorageVirtualDriveAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumBlocks

`func (o *StorageVirtualDriveAllOf) GetNumBlocks() string`

GetNumBlocks returns the NumBlocks field if non-nil, zero value otherwise.

### GetNumBlocksOk

`func (o *StorageVirtualDriveAllOf) GetNumBlocksOk() (*string, bool)`

GetNumBlocksOk returns a tuple with the NumBlocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumBlocks

`func (o *StorageVirtualDriveAllOf) SetNumBlocks(v string)`

SetNumBlocks sets NumBlocks field to given value.

### HasNumBlocks

`func (o *StorageVirtualDriveAllOf) HasNumBlocks() bool`

HasNumBlocks returns a boolean if a field has been set.

### GetOperState

`func (o *StorageVirtualDriveAllOf) GetOperState() string`

GetOperState returns the OperState field if non-nil, zero value otherwise.

### GetOperStateOk

`func (o *StorageVirtualDriveAllOf) GetOperStateOk() (*string, bool)`

GetOperStateOk returns a tuple with the OperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperState

`func (o *StorageVirtualDriveAllOf) SetOperState(v string)`

SetOperState sets OperState field to given value.

### HasOperState

`func (o *StorageVirtualDriveAllOf) HasOperState() bool`

HasOperState returns a boolean if a field has been set.

### GetOperability

`func (o *StorageVirtualDriveAllOf) GetOperability() string`

GetOperability returns the Operability field if non-nil, zero value otherwise.

### GetOperabilityOk

`func (o *StorageVirtualDriveAllOf) GetOperabilityOk() (*string, bool)`

GetOperabilityOk returns a tuple with the Operability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperability

`func (o *StorageVirtualDriveAllOf) SetOperability(v string)`

SetOperability sets Operability field to given value.

### HasOperability

`func (o *StorageVirtualDriveAllOf) HasOperability() bool`

HasOperability returns a boolean if a field has been set.

### GetPhysicalBlockSize

`func (o *StorageVirtualDriveAllOf) GetPhysicalBlockSize() string`

GetPhysicalBlockSize returns the PhysicalBlockSize field if non-nil, zero value otherwise.

### GetPhysicalBlockSizeOk

`func (o *StorageVirtualDriveAllOf) GetPhysicalBlockSizeOk() (*string, bool)`

GetPhysicalBlockSizeOk returns a tuple with the PhysicalBlockSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalBlockSize

`func (o *StorageVirtualDriveAllOf) SetPhysicalBlockSize(v string)`

SetPhysicalBlockSize sets PhysicalBlockSize field to given value.

### HasPhysicalBlockSize

`func (o *StorageVirtualDriveAllOf) HasPhysicalBlockSize() bool`

HasPhysicalBlockSize returns a boolean if a field has been set.

### GetPresence

`func (o *StorageVirtualDriveAllOf) GetPresence() string`

GetPresence returns the Presence field if non-nil, zero value otherwise.

### GetPresenceOk

`func (o *StorageVirtualDriveAllOf) GetPresenceOk() (*string, bool)`

GetPresenceOk returns a tuple with the Presence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresence

`func (o *StorageVirtualDriveAllOf) SetPresence(v string)`

SetPresence sets Presence field to given value.

### HasPresence

`func (o *StorageVirtualDriveAllOf) HasPresence() bool`

HasPresence returns a boolean if a field has been set.

### GetReadPolicy

`func (o *StorageVirtualDriveAllOf) GetReadPolicy() string`

GetReadPolicy returns the ReadPolicy field if non-nil, zero value otherwise.

### GetReadPolicyOk

`func (o *StorageVirtualDriveAllOf) GetReadPolicyOk() (*string, bool)`

GetReadPolicyOk returns a tuple with the ReadPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadPolicy

`func (o *StorageVirtualDriveAllOf) SetReadPolicy(v string)`

SetReadPolicy sets ReadPolicy field to given value.

### HasReadPolicy

`func (o *StorageVirtualDriveAllOf) HasReadPolicy() bool`

HasReadPolicy returns a boolean if a field has been set.

### GetSecurityFlags

`func (o *StorageVirtualDriveAllOf) GetSecurityFlags() string`

GetSecurityFlags returns the SecurityFlags field if non-nil, zero value otherwise.

### GetSecurityFlagsOk

`func (o *StorageVirtualDriveAllOf) GetSecurityFlagsOk() (*string, bool)`

GetSecurityFlagsOk returns a tuple with the SecurityFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityFlags

`func (o *StorageVirtualDriveAllOf) SetSecurityFlags(v string)`

SetSecurityFlags sets SecurityFlags field to given value.

### HasSecurityFlags

`func (o *StorageVirtualDriveAllOf) HasSecurityFlags() bool`

HasSecurityFlags returns a boolean if a field has been set.

### GetSize

`func (o *StorageVirtualDriveAllOf) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *StorageVirtualDriveAllOf) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *StorageVirtualDriveAllOf) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *StorageVirtualDriveAllOf) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetStripSize

`func (o *StorageVirtualDriveAllOf) GetStripSize() string`

GetStripSize returns the StripSize field if non-nil, zero value otherwise.

### GetStripSizeOk

`func (o *StorageVirtualDriveAllOf) GetStripSizeOk() (*string, bool)`

GetStripSizeOk returns a tuple with the StripSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripSize

`func (o *StorageVirtualDriveAllOf) SetStripSize(v string)`

SetStripSize sets StripSize field to given value.

### HasStripSize

`func (o *StorageVirtualDriveAllOf) HasStripSize() bool`

HasStripSize returns a boolean if a field has been set.

### GetType

`func (o *StorageVirtualDriveAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StorageVirtualDriveAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StorageVirtualDriveAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *StorageVirtualDriveAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUuid

`func (o *StorageVirtualDriveAllOf) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *StorageVirtualDriveAllOf) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *StorageVirtualDriveAllOf) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *StorageVirtualDriveAllOf) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetVendorUuid

`func (o *StorageVirtualDriveAllOf) GetVendorUuid() string`

GetVendorUuid returns the VendorUuid field if non-nil, zero value otherwise.

### GetVendorUuidOk

`func (o *StorageVirtualDriveAllOf) GetVendorUuidOk() (*string, bool)`

GetVendorUuidOk returns a tuple with the VendorUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorUuid

`func (o *StorageVirtualDriveAllOf) SetVendorUuid(v string)`

SetVendorUuid sets VendorUuid field to given value.

### HasVendorUuid

`func (o *StorageVirtualDriveAllOf) HasVendorUuid() bool`

HasVendorUuid returns a boolean if a field has been set.

### GetVirtualDriveId

`func (o *StorageVirtualDriveAllOf) GetVirtualDriveId() string`

GetVirtualDriveId returns the VirtualDriveId field if non-nil, zero value otherwise.

### GetVirtualDriveIdOk

`func (o *StorageVirtualDriveAllOf) GetVirtualDriveIdOk() (*string, bool)`

GetVirtualDriveIdOk returns a tuple with the VirtualDriveId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualDriveId

`func (o *StorageVirtualDriveAllOf) SetVirtualDriveId(v string)`

SetVirtualDriveId sets VirtualDriveId field to given value.

### HasVirtualDriveId

`func (o *StorageVirtualDriveAllOf) HasVirtualDriveId() bool`

HasVirtualDriveId returns a boolean if a field has been set.

### GetDiskGroup

`func (o *StorageVirtualDriveAllOf) GetDiskGroup() StorageDiskGroupRelationship`

GetDiskGroup returns the DiskGroup field if non-nil, zero value otherwise.

### GetDiskGroupOk

`func (o *StorageVirtualDriveAllOf) GetDiskGroupOk() (*StorageDiskGroupRelationship, bool)`

GetDiskGroupOk returns a tuple with the DiskGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskGroup

`func (o *StorageVirtualDriveAllOf) SetDiskGroup(v StorageDiskGroupRelationship)`

SetDiskGroup sets DiskGroup field to given value.

### HasDiskGroup

`func (o *StorageVirtualDriveAllOf) HasDiskGroup() bool`

HasDiskGroup returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *StorageVirtualDriveAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *StorageVirtualDriveAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *StorageVirtualDriveAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *StorageVirtualDriveAllOf) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetPhysicalDiskUsages

`func (o *StorageVirtualDriveAllOf) GetPhysicalDiskUsages() []StoragePhysicalDiskUsageRelationship`

GetPhysicalDiskUsages returns the PhysicalDiskUsages field if non-nil, zero value otherwise.

### GetPhysicalDiskUsagesOk

`func (o *StorageVirtualDriveAllOf) GetPhysicalDiskUsagesOk() (*[]StoragePhysicalDiskUsageRelationship, bool)`

GetPhysicalDiskUsagesOk returns a tuple with the PhysicalDiskUsages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalDiskUsages

`func (o *StorageVirtualDriveAllOf) SetPhysicalDiskUsages(v []StoragePhysicalDiskUsageRelationship)`

SetPhysicalDiskUsages sets PhysicalDiskUsages field to given value.

### HasPhysicalDiskUsages

`func (o *StorageVirtualDriveAllOf) HasPhysicalDiskUsages() bool`

HasPhysicalDiskUsages returns a boolean if a field has been set.

### SetPhysicalDiskUsagesNil

`func (o *StorageVirtualDriveAllOf) SetPhysicalDiskUsagesNil(b bool)`

 SetPhysicalDiskUsagesNil sets the value for PhysicalDiskUsages to be an explicit nil

### UnsetPhysicalDiskUsages
`func (o *StorageVirtualDriveAllOf) UnsetPhysicalDiskUsages()`

UnsetPhysicalDiskUsages ensures that no value is present for PhysicalDiskUsages, not even an explicit nil
### GetRegisteredDevice

`func (o *StorageVirtualDriveAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StorageVirtualDriveAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StorageVirtualDriveAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StorageVirtualDriveAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### GetStorageController

`func (o *StorageVirtualDriveAllOf) GetStorageController() StorageControllerRelationship`

GetStorageController returns the StorageController field if non-nil, zero value otherwise.

### GetStorageControllerOk

`func (o *StorageVirtualDriveAllOf) GetStorageControllerOk() (*StorageControllerRelationship, bool)`

GetStorageControllerOk returns a tuple with the StorageController field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageController

`func (o *StorageVirtualDriveAllOf) SetStorageController(v StorageControllerRelationship)`

SetStorageController sets StorageController field to given value.

### HasStorageController

`func (o *StorageVirtualDriveAllOf) HasStorageController() bool`

HasStorageController returns a boolean if a field has been set.

### GetVdMemberEps

`func (o *StorageVirtualDriveAllOf) GetVdMemberEps() []StorageVdMemberEpRelationship`

GetVdMemberEps returns the VdMemberEps field if non-nil, zero value otherwise.

### GetVdMemberEpsOk

`func (o *StorageVirtualDriveAllOf) GetVdMemberEpsOk() (*[]StorageVdMemberEpRelationship, bool)`

GetVdMemberEpsOk returns a tuple with the VdMemberEps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdMemberEps

`func (o *StorageVirtualDriveAllOf) SetVdMemberEps(v []StorageVdMemberEpRelationship)`

SetVdMemberEps sets VdMemberEps field to given value.

### HasVdMemberEps

`func (o *StorageVirtualDriveAllOf) HasVdMemberEps() bool`

HasVdMemberEps returns a boolean if a field has been set.

### SetVdMemberEpsNil

`func (o *StorageVirtualDriveAllOf) SetVdMemberEpsNil(b bool)`

 SetVdMemberEpsNil sets the value for VdMemberEps to be an explicit nil

### UnsetVdMemberEps
`func (o *StorageVirtualDriveAllOf) UnsetVdMemberEps()`

UnsetVdMemberEps ensures that no value is present for VdMemberEps, not even an explicit nil
### GetVirtualDriveExtension

`func (o *StorageVirtualDriveAllOf) GetVirtualDriveExtension() StorageVirtualDriveExtensionRelationship`

GetVirtualDriveExtension returns the VirtualDriveExtension field if non-nil, zero value otherwise.

### GetVirtualDriveExtensionOk

`func (o *StorageVirtualDriveAllOf) GetVirtualDriveExtensionOk() (*StorageVirtualDriveExtensionRelationship, bool)`

GetVirtualDriveExtensionOk returns a tuple with the VirtualDriveExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualDriveExtension

`func (o *StorageVirtualDriveAllOf) SetVirtualDriveExtension(v StorageVirtualDriveExtensionRelationship)`

SetVirtualDriveExtension sets VirtualDriveExtension field to given value.

### HasVirtualDriveExtension

`func (o *StorageVirtualDriveAllOf) HasVirtualDriveExtension() bool`

HasVirtualDriveExtension returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



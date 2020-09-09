# StorageEnclosureDiskAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockSize** | Pointer to **string** | The block size of the physical disk in bytes. | [optional] 
**DiskId** | Pointer to **string** | This field represents the disk Id in the storage enclosure. | [optional] 
**DiskState** | Pointer to **string** | This field identifies the current disk configuration applied in the physical disk. | [optional] 
**Health** | Pointer to **string** | The current health state of the enclosure disk. | [optional] 
**NumBlocks** | Pointer to **string** | The number of blocks present on the physical disk. | [optional] 
**Pid** | Pointer to **string** | This field identifies the Product ID for physicalDisk. | [optional] [readonly] 
**SasAddress1** | Pointer to **string** | This field identifies the SAS address assigned to the disk SAS port-1. | [optional] 
**SasAddress2** | Pointer to **string** | This field identifies the SAS address assigned to the disk SAS port-2. | [optional] 
**Size** | Pointer to **string** | The size of the physical disk in MB. | [optional] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](inventory.DeviceInfo.Relationship.md) |  | [optional] 
**PhysicalDisk** | Pointer to [**StoragePhysicalDiskRelationship**](storage.PhysicalDisk.Relationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 
**StorageEnclosure** | Pointer to [**StorageEnclosureRelationship**](storage.Enclosure.Relationship.md) |  | [optional] 

## Methods

### NewStorageEnclosureDiskAllOf

`func NewStorageEnclosureDiskAllOf() *StorageEnclosureDiskAllOf`

NewStorageEnclosureDiskAllOf instantiates a new StorageEnclosureDiskAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageEnclosureDiskAllOfWithDefaults

`func NewStorageEnclosureDiskAllOfWithDefaults() *StorageEnclosureDiskAllOf`

NewStorageEnclosureDiskAllOfWithDefaults instantiates a new StorageEnclosureDiskAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockSize

`func (o *StorageEnclosureDiskAllOf) GetBlockSize() string`

GetBlockSize returns the BlockSize field if non-nil, zero value otherwise.

### GetBlockSizeOk

`func (o *StorageEnclosureDiskAllOf) GetBlockSizeOk() (*string, bool)`

GetBlockSizeOk returns a tuple with the BlockSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockSize

`func (o *StorageEnclosureDiskAllOf) SetBlockSize(v string)`

SetBlockSize sets BlockSize field to given value.

### HasBlockSize

`func (o *StorageEnclosureDiskAllOf) HasBlockSize() bool`

HasBlockSize returns a boolean if a field has been set.

### GetDiskId

`func (o *StorageEnclosureDiskAllOf) GetDiskId() string`

GetDiskId returns the DiskId field if non-nil, zero value otherwise.

### GetDiskIdOk

`func (o *StorageEnclosureDiskAllOf) GetDiskIdOk() (*string, bool)`

GetDiskIdOk returns a tuple with the DiskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskId

`func (o *StorageEnclosureDiskAllOf) SetDiskId(v string)`

SetDiskId sets DiskId field to given value.

### HasDiskId

`func (o *StorageEnclosureDiskAllOf) HasDiskId() bool`

HasDiskId returns a boolean if a field has been set.

### GetDiskState

`func (o *StorageEnclosureDiskAllOf) GetDiskState() string`

GetDiskState returns the DiskState field if non-nil, zero value otherwise.

### GetDiskStateOk

`func (o *StorageEnclosureDiskAllOf) GetDiskStateOk() (*string, bool)`

GetDiskStateOk returns a tuple with the DiskState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskState

`func (o *StorageEnclosureDiskAllOf) SetDiskState(v string)`

SetDiskState sets DiskState field to given value.

### HasDiskState

`func (o *StorageEnclosureDiskAllOf) HasDiskState() bool`

HasDiskState returns a boolean if a field has been set.

### GetHealth

`func (o *StorageEnclosureDiskAllOf) GetHealth() string`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *StorageEnclosureDiskAllOf) GetHealthOk() (*string, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *StorageEnclosureDiskAllOf) SetHealth(v string)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *StorageEnclosureDiskAllOf) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetNumBlocks

`func (o *StorageEnclosureDiskAllOf) GetNumBlocks() string`

GetNumBlocks returns the NumBlocks field if non-nil, zero value otherwise.

### GetNumBlocksOk

`func (o *StorageEnclosureDiskAllOf) GetNumBlocksOk() (*string, bool)`

GetNumBlocksOk returns a tuple with the NumBlocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumBlocks

`func (o *StorageEnclosureDiskAllOf) SetNumBlocks(v string)`

SetNumBlocks sets NumBlocks field to given value.

### HasNumBlocks

`func (o *StorageEnclosureDiskAllOf) HasNumBlocks() bool`

HasNumBlocks returns a boolean if a field has been set.

### GetPid

`func (o *StorageEnclosureDiskAllOf) GetPid() string`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *StorageEnclosureDiskAllOf) GetPidOk() (*string, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *StorageEnclosureDiskAllOf) SetPid(v string)`

SetPid sets Pid field to given value.

### HasPid

`func (o *StorageEnclosureDiskAllOf) HasPid() bool`

HasPid returns a boolean if a field has been set.

### GetSasAddress1

`func (o *StorageEnclosureDiskAllOf) GetSasAddress1() string`

GetSasAddress1 returns the SasAddress1 field if non-nil, zero value otherwise.

### GetSasAddress1Ok

`func (o *StorageEnclosureDiskAllOf) GetSasAddress1Ok() (*string, bool)`

GetSasAddress1Ok returns a tuple with the SasAddress1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSasAddress1

`func (o *StorageEnclosureDiskAllOf) SetSasAddress1(v string)`

SetSasAddress1 sets SasAddress1 field to given value.

### HasSasAddress1

`func (o *StorageEnclosureDiskAllOf) HasSasAddress1() bool`

HasSasAddress1 returns a boolean if a field has been set.

### GetSasAddress2

`func (o *StorageEnclosureDiskAllOf) GetSasAddress2() string`

GetSasAddress2 returns the SasAddress2 field if non-nil, zero value otherwise.

### GetSasAddress2Ok

`func (o *StorageEnclosureDiskAllOf) GetSasAddress2Ok() (*string, bool)`

GetSasAddress2Ok returns a tuple with the SasAddress2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSasAddress2

`func (o *StorageEnclosureDiskAllOf) SetSasAddress2(v string)`

SetSasAddress2 sets SasAddress2 field to given value.

### HasSasAddress2

`func (o *StorageEnclosureDiskAllOf) HasSasAddress2() bool`

HasSasAddress2 returns a boolean if a field has been set.

### GetSize

`func (o *StorageEnclosureDiskAllOf) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *StorageEnclosureDiskAllOf) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *StorageEnclosureDiskAllOf) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *StorageEnclosureDiskAllOf) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *StorageEnclosureDiskAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *StorageEnclosureDiskAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *StorageEnclosureDiskAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *StorageEnclosureDiskAllOf) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetPhysicalDisk

`func (o *StorageEnclosureDiskAllOf) GetPhysicalDisk() StoragePhysicalDiskRelationship`

GetPhysicalDisk returns the PhysicalDisk field if non-nil, zero value otherwise.

### GetPhysicalDiskOk

`func (o *StorageEnclosureDiskAllOf) GetPhysicalDiskOk() (*StoragePhysicalDiskRelationship, bool)`

GetPhysicalDiskOk returns a tuple with the PhysicalDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalDisk

`func (o *StorageEnclosureDiskAllOf) SetPhysicalDisk(v StoragePhysicalDiskRelationship)`

SetPhysicalDisk sets PhysicalDisk field to given value.

### HasPhysicalDisk

`func (o *StorageEnclosureDiskAllOf) HasPhysicalDisk() bool`

HasPhysicalDisk returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *StorageEnclosureDiskAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StorageEnclosureDiskAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StorageEnclosureDiskAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StorageEnclosureDiskAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### GetStorageEnclosure

`func (o *StorageEnclosureDiskAllOf) GetStorageEnclosure() StorageEnclosureRelationship`

GetStorageEnclosure returns the StorageEnclosure field if non-nil, zero value otherwise.

### GetStorageEnclosureOk

`func (o *StorageEnclosureDiskAllOf) GetStorageEnclosureOk() (*StorageEnclosureRelationship, bool)`

GetStorageEnclosureOk returns a tuple with the StorageEnclosure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageEnclosure

`func (o *StorageEnclosureDiskAllOf) SetStorageEnclosure(v StorageEnclosureRelationship)`

SetStorageEnclosure sets StorageEnclosure field to given value.

### HasStorageEnclosure

`func (o *StorageEnclosureDiskAllOf) HasStorageEnclosure() bool`

HasStorageEnclosure returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



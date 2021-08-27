# StorageEnclosureDisk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.EnclosureDisk"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.EnclosureDisk"]
**BlockSize** | Pointer to **string** | The block size of the physical disk in bytes. | [optional] 
**DiskId** | Pointer to **string** | This field represents the disk Id in the storage enclosure. | [optional] 
**DiskState** | Pointer to **string** | This field identifies the current disk configuration applied in the physical disk. | [optional] 
**Health** | Pointer to **string** | The current health state of the enclosure disk. | [optional] 
**NumBlocks** | Pointer to **string** | The number of blocks present on the physical disk. | [optional] 
**Pid** | Pointer to **string** | This field identifies the Product ID for physicalDisk. | [optional] [readonly] 
**SasAddress1** | Pointer to **string** | This field identifies the SAS address assigned to the disk SAS port-1. | [optional] 
**SasAddress2** | Pointer to **string** | This field identifies the SAS address assigned to the disk SAS port-2. | [optional] 
**Size** | Pointer to **string** | The size of the physical disk in MB. | [optional] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](InventoryDeviceInfoRelationship.md) |  | [optional] 
**PhysicalDisk** | Pointer to [**StoragePhysicalDiskRelationship**](StoragePhysicalDiskRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 
**StorageEnclosure** | Pointer to [**StorageEnclosureRelationship**](StorageEnclosureRelationship.md) |  | [optional] 

## Methods

### NewStorageEnclosureDisk

`func NewStorageEnclosureDisk(classId string, objectType string, ) *StorageEnclosureDisk`

NewStorageEnclosureDisk instantiates a new StorageEnclosureDisk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageEnclosureDiskWithDefaults

`func NewStorageEnclosureDiskWithDefaults() *StorageEnclosureDisk`

NewStorageEnclosureDiskWithDefaults instantiates a new StorageEnclosureDisk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageEnclosureDisk) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageEnclosureDisk) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageEnclosureDisk) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageEnclosureDisk) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageEnclosureDisk) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageEnclosureDisk) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBlockSize

`func (o *StorageEnclosureDisk) GetBlockSize() string`

GetBlockSize returns the BlockSize field if non-nil, zero value otherwise.

### GetBlockSizeOk

`func (o *StorageEnclosureDisk) GetBlockSizeOk() (*string, bool)`

GetBlockSizeOk returns a tuple with the BlockSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockSize

`func (o *StorageEnclosureDisk) SetBlockSize(v string)`

SetBlockSize sets BlockSize field to given value.

### HasBlockSize

`func (o *StorageEnclosureDisk) HasBlockSize() bool`

HasBlockSize returns a boolean if a field has been set.

### GetDiskId

`func (o *StorageEnclosureDisk) GetDiskId() string`

GetDiskId returns the DiskId field if non-nil, zero value otherwise.

### GetDiskIdOk

`func (o *StorageEnclosureDisk) GetDiskIdOk() (*string, bool)`

GetDiskIdOk returns a tuple with the DiskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskId

`func (o *StorageEnclosureDisk) SetDiskId(v string)`

SetDiskId sets DiskId field to given value.

### HasDiskId

`func (o *StorageEnclosureDisk) HasDiskId() bool`

HasDiskId returns a boolean if a field has been set.

### GetDiskState

`func (o *StorageEnclosureDisk) GetDiskState() string`

GetDiskState returns the DiskState field if non-nil, zero value otherwise.

### GetDiskStateOk

`func (o *StorageEnclosureDisk) GetDiskStateOk() (*string, bool)`

GetDiskStateOk returns a tuple with the DiskState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskState

`func (o *StorageEnclosureDisk) SetDiskState(v string)`

SetDiskState sets DiskState field to given value.

### HasDiskState

`func (o *StorageEnclosureDisk) HasDiskState() bool`

HasDiskState returns a boolean if a field has been set.

### GetHealth

`func (o *StorageEnclosureDisk) GetHealth() string`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *StorageEnclosureDisk) GetHealthOk() (*string, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *StorageEnclosureDisk) SetHealth(v string)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *StorageEnclosureDisk) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetNumBlocks

`func (o *StorageEnclosureDisk) GetNumBlocks() string`

GetNumBlocks returns the NumBlocks field if non-nil, zero value otherwise.

### GetNumBlocksOk

`func (o *StorageEnclosureDisk) GetNumBlocksOk() (*string, bool)`

GetNumBlocksOk returns a tuple with the NumBlocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumBlocks

`func (o *StorageEnclosureDisk) SetNumBlocks(v string)`

SetNumBlocks sets NumBlocks field to given value.

### HasNumBlocks

`func (o *StorageEnclosureDisk) HasNumBlocks() bool`

HasNumBlocks returns a boolean if a field has been set.

### GetPid

`func (o *StorageEnclosureDisk) GetPid() string`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *StorageEnclosureDisk) GetPidOk() (*string, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *StorageEnclosureDisk) SetPid(v string)`

SetPid sets Pid field to given value.

### HasPid

`func (o *StorageEnclosureDisk) HasPid() bool`

HasPid returns a boolean if a field has been set.

### GetSasAddress1

`func (o *StorageEnclosureDisk) GetSasAddress1() string`

GetSasAddress1 returns the SasAddress1 field if non-nil, zero value otherwise.

### GetSasAddress1Ok

`func (o *StorageEnclosureDisk) GetSasAddress1Ok() (*string, bool)`

GetSasAddress1Ok returns a tuple with the SasAddress1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSasAddress1

`func (o *StorageEnclosureDisk) SetSasAddress1(v string)`

SetSasAddress1 sets SasAddress1 field to given value.

### HasSasAddress1

`func (o *StorageEnclosureDisk) HasSasAddress1() bool`

HasSasAddress1 returns a boolean if a field has been set.

### GetSasAddress2

`func (o *StorageEnclosureDisk) GetSasAddress2() string`

GetSasAddress2 returns the SasAddress2 field if non-nil, zero value otherwise.

### GetSasAddress2Ok

`func (o *StorageEnclosureDisk) GetSasAddress2Ok() (*string, bool)`

GetSasAddress2Ok returns a tuple with the SasAddress2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSasAddress2

`func (o *StorageEnclosureDisk) SetSasAddress2(v string)`

SetSasAddress2 sets SasAddress2 field to given value.

### HasSasAddress2

`func (o *StorageEnclosureDisk) HasSasAddress2() bool`

HasSasAddress2 returns a boolean if a field has been set.

### GetSize

`func (o *StorageEnclosureDisk) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *StorageEnclosureDisk) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *StorageEnclosureDisk) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *StorageEnclosureDisk) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *StorageEnclosureDisk) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *StorageEnclosureDisk) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *StorageEnclosureDisk) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *StorageEnclosureDisk) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetPhysicalDisk

`func (o *StorageEnclosureDisk) GetPhysicalDisk() StoragePhysicalDiskRelationship`

GetPhysicalDisk returns the PhysicalDisk field if non-nil, zero value otherwise.

### GetPhysicalDiskOk

`func (o *StorageEnclosureDisk) GetPhysicalDiskOk() (*StoragePhysicalDiskRelationship, bool)`

GetPhysicalDiskOk returns a tuple with the PhysicalDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalDisk

`func (o *StorageEnclosureDisk) SetPhysicalDisk(v StoragePhysicalDiskRelationship)`

SetPhysicalDisk sets PhysicalDisk field to given value.

### HasPhysicalDisk

`func (o *StorageEnclosureDisk) HasPhysicalDisk() bool`

HasPhysicalDisk returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *StorageEnclosureDisk) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StorageEnclosureDisk) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StorageEnclosureDisk) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StorageEnclosureDisk) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### GetStorageEnclosure

`func (o *StorageEnclosureDisk) GetStorageEnclosure() StorageEnclosureRelationship`

GetStorageEnclosure returns the StorageEnclosure field if non-nil, zero value otherwise.

### GetStorageEnclosureOk

`func (o *StorageEnclosureDisk) GetStorageEnclosureOk() (*StorageEnclosureRelationship, bool)`

GetStorageEnclosureOk returns a tuple with the StorageEnclosure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageEnclosure

`func (o *StorageEnclosureDisk) SetStorageEnclosure(v StorageEnclosureRelationship)`

SetStorageEnclosure sets StorageEnclosure field to given value.

### HasStorageEnclosure

`func (o *StorageEnclosureDisk) HasStorageEnclosure() bool`

HasStorageEnclosure returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



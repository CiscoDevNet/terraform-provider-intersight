# StoragePhysicalDiskExtension

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bootable** | Pointer to **string** | The whether disk is bootable or not. | [optional] [readonly] 
**DiskDn** | Pointer to **string** | The distinguished name of the Physical drive. | [optional] [readonly] 
**DiskId** | Pointer to **int64** | The storage Enclosure slotId. | [optional] [readonly] 
**DiskState** | Pointer to **string** | The current drive state of disk. | [optional] [readonly] 
**Health** | Pointer to **string** | The current drive state of disk. | [optional] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](inventory.DeviceInfo.Relationship.md) |  | [optional] 
**PhysicalDisk** | Pointer to [**StoragePhysicalDiskRelationship**](storage.PhysicalDisk.Relationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 
**StorageController** | Pointer to [**StorageControllerRelationship**](storage.Controller.Relationship.md) |  | [optional] 

## Methods

### NewStoragePhysicalDiskExtension

`func NewStoragePhysicalDiskExtension() *StoragePhysicalDiskExtension`

NewStoragePhysicalDiskExtension instantiates a new StoragePhysicalDiskExtension object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoragePhysicalDiskExtensionWithDefaults

`func NewStoragePhysicalDiskExtensionWithDefaults() *StoragePhysicalDiskExtension`

NewStoragePhysicalDiskExtensionWithDefaults instantiates a new StoragePhysicalDiskExtension object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBootable

`func (o *StoragePhysicalDiskExtension) GetBootable() string`

GetBootable returns the Bootable field if non-nil, zero value otherwise.

### GetBootableOk

`func (o *StoragePhysicalDiskExtension) GetBootableOk() (*string, bool)`

GetBootableOk returns a tuple with the Bootable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootable

`func (o *StoragePhysicalDiskExtension) SetBootable(v string)`

SetBootable sets Bootable field to given value.

### HasBootable

`func (o *StoragePhysicalDiskExtension) HasBootable() bool`

HasBootable returns a boolean if a field has been set.

### GetDiskDn

`func (o *StoragePhysicalDiskExtension) GetDiskDn() string`

GetDiskDn returns the DiskDn field if non-nil, zero value otherwise.

### GetDiskDnOk

`func (o *StoragePhysicalDiskExtension) GetDiskDnOk() (*string, bool)`

GetDiskDnOk returns a tuple with the DiskDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskDn

`func (o *StoragePhysicalDiskExtension) SetDiskDn(v string)`

SetDiskDn sets DiskDn field to given value.

### HasDiskDn

`func (o *StoragePhysicalDiskExtension) HasDiskDn() bool`

HasDiskDn returns a boolean if a field has been set.

### GetDiskId

`func (o *StoragePhysicalDiskExtension) GetDiskId() int64`

GetDiskId returns the DiskId field if non-nil, zero value otherwise.

### GetDiskIdOk

`func (o *StoragePhysicalDiskExtension) GetDiskIdOk() (*int64, bool)`

GetDiskIdOk returns a tuple with the DiskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskId

`func (o *StoragePhysicalDiskExtension) SetDiskId(v int64)`

SetDiskId sets DiskId field to given value.

### HasDiskId

`func (o *StoragePhysicalDiskExtension) HasDiskId() bool`

HasDiskId returns a boolean if a field has been set.

### GetDiskState

`func (o *StoragePhysicalDiskExtension) GetDiskState() string`

GetDiskState returns the DiskState field if non-nil, zero value otherwise.

### GetDiskStateOk

`func (o *StoragePhysicalDiskExtension) GetDiskStateOk() (*string, bool)`

GetDiskStateOk returns a tuple with the DiskState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskState

`func (o *StoragePhysicalDiskExtension) SetDiskState(v string)`

SetDiskState sets DiskState field to given value.

### HasDiskState

`func (o *StoragePhysicalDiskExtension) HasDiskState() bool`

HasDiskState returns a boolean if a field has been set.

### GetHealth

`func (o *StoragePhysicalDiskExtension) GetHealth() string`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *StoragePhysicalDiskExtension) GetHealthOk() (*string, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *StoragePhysicalDiskExtension) SetHealth(v string)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *StoragePhysicalDiskExtension) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *StoragePhysicalDiskExtension) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *StoragePhysicalDiskExtension) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *StoragePhysicalDiskExtension) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *StoragePhysicalDiskExtension) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetPhysicalDisk

`func (o *StoragePhysicalDiskExtension) GetPhysicalDisk() StoragePhysicalDiskRelationship`

GetPhysicalDisk returns the PhysicalDisk field if non-nil, zero value otherwise.

### GetPhysicalDiskOk

`func (o *StoragePhysicalDiskExtension) GetPhysicalDiskOk() (*StoragePhysicalDiskRelationship, bool)`

GetPhysicalDiskOk returns a tuple with the PhysicalDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalDisk

`func (o *StoragePhysicalDiskExtension) SetPhysicalDisk(v StoragePhysicalDiskRelationship)`

SetPhysicalDisk sets PhysicalDisk field to given value.

### HasPhysicalDisk

`func (o *StoragePhysicalDiskExtension) HasPhysicalDisk() bool`

HasPhysicalDisk returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *StoragePhysicalDiskExtension) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StoragePhysicalDiskExtension) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StoragePhysicalDiskExtension) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StoragePhysicalDiskExtension) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### GetStorageController

`func (o *StoragePhysicalDiskExtension) GetStorageController() StorageControllerRelationship`

GetStorageController returns the StorageController field if non-nil, zero value otherwise.

### GetStorageControllerOk

`func (o *StoragePhysicalDiskExtension) GetStorageControllerOk() (*StorageControllerRelationship, bool)`

GetStorageControllerOk returns a tuple with the StorageController field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageController

`func (o *StoragePhysicalDiskExtension) SetStorageController(v StorageControllerRelationship)`

SetStorageController sets StorageController field to given value.

### HasStorageController

`func (o *StoragePhysicalDiskExtension) HasStorageController() bool`

HasStorageController returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



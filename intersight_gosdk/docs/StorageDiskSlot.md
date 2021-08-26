# StorageDiskSlot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.DiskSlot"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.DiskSlot"]
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](InventoryDeviceInfoRelationship.md) |  | [optional] 
**PhysicalDisk** | Pointer to [**StoragePhysicalDiskRelationship**](StoragePhysicalDiskRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 
**StorageController** | Pointer to [**StorageControllerRelationship**](StorageControllerRelationship.md) |  | [optional] 

## Methods

### NewStorageDiskSlot

`func NewStorageDiskSlot(classId string, objectType string, ) *StorageDiskSlot`

NewStorageDiskSlot instantiates a new StorageDiskSlot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageDiskSlotWithDefaults

`func NewStorageDiskSlotWithDefaults() *StorageDiskSlot`

NewStorageDiskSlotWithDefaults instantiates a new StorageDiskSlot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageDiskSlot) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageDiskSlot) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageDiskSlot) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageDiskSlot) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageDiskSlot) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageDiskSlot) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetInventoryDeviceInfo

`func (o *StorageDiskSlot) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *StorageDiskSlot) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *StorageDiskSlot) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *StorageDiskSlot) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetPhysicalDisk

`func (o *StorageDiskSlot) GetPhysicalDisk() StoragePhysicalDiskRelationship`

GetPhysicalDisk returns the PhysicalDisk field if non-nil, zero value otherwise.

### GetPhysicalDiskOk

`func (o *StorageDiskSlot) GetPhysicalDiskOk() (*StoragePhysicalDiskRelationship, bool)`

GetPhysicalDiskOk returns a tuple with the PhysicalDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalDisk

`func (o *StorageDiskSlot) SetPhysicalDisk(v StoragePhysicalDiskRelationship)`

SetPhysicalDisk sets PhysicalDisk field to given value.

### HasPhysicalDisk

`func (o *StorageDiskSlot) HasPhysicalDisk() bool`

HasPhysicalDisk returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *StorageDiskSlot) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StorageDiskSlot) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StorageDiskSlot) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StorageDiskSlot) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### GetStorageController

`func (o *StorageDiskSlot) GetStorageController() StorageControllerRelationship`

GetStorageController returns the StorageController field if non-nil, zero value otherwise.

### GetStorageControllerOk

`func (o *StorageDiskSlot) GetStorageControllerOk() (*StorageControllerRelationship, bool)`

GetStorageControllerOk returns a tuple with the StorageController field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageController

`func (o *StorageDiskSlot) SetStorageController(v StorageControllerRelationship)`

SetStorageController sets StorageController field to given value.

### HasStorageController

`func (o *StorageDiskSlot) HasStorageController() bool`

HasStorageController returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



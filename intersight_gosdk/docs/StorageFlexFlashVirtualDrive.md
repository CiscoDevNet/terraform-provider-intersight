# StorageFlexFlashVirtualDrive

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DriveScope** | Pointer to **string** | The drive scope of the flex flash virtual drive. | [optional] 
**DriveStatus** | Pointer to **string** | Status of virtual drive on the flex controller. | [optional] 
**PartitionId** | Pointer to **string** | The partition Id of the flex flash virtual Drive. | [optional] 
**ResidentImage** | Pointer to **string** | The resident image on the flex flash virtual Drive. | [optional] 
**Size** | Pointer to **string** | Size of virtual drive on the flex controller. | [optional] 
**VirtualDrive** | Pointer to **string** | Virtual drive on the flex flash controller. | [optional] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](inventory.DeviceInfo.Relationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 
**StorageFlexFlashController** | Pointer to [**StorageFlexFlashControllerRelationship**](storage.FlexFlashController.Relationship.md) |  | [optional] 

## Methods

### NewStorageFlexFlashVirtualDrive

`func NewStorageFlexFlashVirtualDrive() *StorageFlexFlashVirtualDrive`

NewStorageFlexFlashVirtualDrive instantiates a new StorageFlexFlashVirtualDrive object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageFlexFlashVirtualDriveWithDefaults

`func NewStorageFlexFlashVirtualDriveWithDefaults() *StorageFlexFlashVirtualDrive`

NewStorageFlexFlashVirtualDriveWithDefaults instantiates a new StorageFlexFlashVirtualDrive object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDriveScope

`func (o *StorageFlexFlashVirtualDrive) GetDriveScope() string`

GetDriveScope returns the DriveScope field if non-nil, zero value otherwise.

### GetDriveScopeOk

`func (o *StorageFlexFlashVirtualDrive) GetDriveScopeOk() (*string, bool)`

GetDriveScopeOk returns a tuple with the DriveScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriveScope

`func (o *StorageFlexFlashVirtualDrive) SetDriveScope(v string)`

SetDriveScope sets DriveScope field to given value.

### HasDriveScope

`func (o *StorageFlexFlashVirtualDrive) HasDriveScope() bool`

HasDriveScope returns a boolean if a field has been set.

### GetDriveStatus

`func (o *StorageFlexFlashVirtualDrive) GetDriveStatus() string`

GetDriveStatus returns the DriveStatus field if non-nil, zero value otherwise.

### GetDriveStatusOk

`func (o *StorageFlexFlashVirtualDrive) GetDriveStatusOk() (*string, bool)`

GetDriveStatusOk returns a tuple with the DriveStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriveStatus

`func (o *StorageFlexFlashVirtualDrive) SetDriveStatus(v string)`

SetDriveStatus sets DriveStatus field to given value.

### HasDriveStatus

`func (o *StorageFlexFlashVirtualDrive) HasDriveStatus() bool`

HasDriveStatus returns a boolean if a field has been set.

### GetPartitionId

`func (o *StorageFlexFlashVirtualDrive) GetPartitionId() string`

GetPartitionId returns the PartitionId field if non-nil, zero value otherwise.

### GetPartitionIdOk

`func (o *StorageFlexFlashVirtualDrive) GetPartitionIdOk() (*string, bool)`

GetPartitionIdOk returns a tuple with the PartitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionId

`func (o *StorageFlexFlashVirtualDrive) SetPartitionId(v string)`

SetPartitionId sets PartitionId field to given value.

### HasPartitionId

`func (o *StorageFlexFlashVirtualDrive) HasPartitionId() bool`

HasPartitionId returns a boolean if a field has been set.

### GetResidentImage

`func (o *StorageFlexFlashVirtualDrive) GetResidentImage() string`

GetResidentImage returns the ResidentImage field if non-nil, zero value otherwise.

### GetResidentImageOk

`func (o *StorageFlexFlashVirtualDrive) GetResidentImageOk() (*string, bool)`

GetResidentImageOk returns a tuple with the ResidentImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResidentImage

`func (o *StorageFlexFlashVirtualDrive) SetResidentImage(v string)`

SetResidentImage sets ResidentImage field to given value.

### HasResidentImage

`func (o *StorageFlexFlashVirtualDrive) HasResidentImage() bool`

HasResidentImage returns a boolean if a field has been set.

### GetSize

`func (o *StorageFlexFlashVirtualDrive) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *StorageFlexFlashVirtualDrive) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *StorageFlexFlashVirtualDrive) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *StorageFlexFlashVirtualDrive) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetVirtualDrive

`func (o *StorageFlexFlashVirtualDrive) GetVirtualDrive() string`

GetVirtualDrive returns the VirtualDrive field if non-nil, zero value otherwise.

### GetVirtualDriveOk

`func (o *StorageFlexFlashVirtualDrive) GetVirtualDriveOk() (*string, bool)`

GetVirtualDriveOk returns a tuple with the VirtualDrive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualDrive

`func (o *StorageFlexFlashVirtualDrive) SetVirtualDrive(v string)`

SetVirtualDrive sets VirtualDrive field to given value.

### HasVirtualDrive

`func (o *StorageFlexFlashVirtualDrive) HasVirtualDrive() bool`

HasVirtualDrive returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *StorageFlexFlashVirtualDrive) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *StorageFlexFlashVirtualDrive) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *StorageFlexFlashVirtualDrive) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *StorageFlexFlashVirtualDrive) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *StorageFlexFlashVirtualDrive) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StorageFlexFlashVirtualDrive) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StorageFlexFlashVirtualDrive) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StorageFlexFlashVirtualDrive) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### GetStorageFlexFlashController

`func (o *StorageFlexFlashVirtualDrive) GetStorageFlexFlashController() StorageFlexFlashControllerRelationship`

GetStorageFlexFlashController returns the StorageFlexFlashController field if non-nil, zero value otherwise.

### GetStorageFlexFlashControllerOk

`func (o *StorageFlexFlashVirtualDrive) GetStorageFlexFlashControllerOk() (*StorageFlexFlashControllerRelationship, bool)`

GetStorageFlexFlashControllerOk returns a tuple with the StorageFlexFlashController field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageFlexFlashController

`func (o *StorageFlexFlashVirtualDrive) SetStorageFlexFlashController(v StorageFlexFlashControllerRelationship)`

SetStorageFlexFlashController sets StorageFlexFlashController field to given value.

### HasStorageFlexFlashController

`func (o *StorageFlexFlashVirtualDrive) HasStorageFlexFlashController() bool`

HasStorageFlexFlashController returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



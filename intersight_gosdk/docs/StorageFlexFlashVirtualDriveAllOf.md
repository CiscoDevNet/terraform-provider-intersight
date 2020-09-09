# StorageFlexFlashVirtualDriveAllOf

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

### NewStorageFlexFlashVirtualDriveAllOf

`func NewStorageFlexFlashVirtualDriveAllOf() *StorageFlexFlashVirtualDriveAllOf`

NewStorageFlexFlashVirtualDriveAllOf instantiates a new StorageFlexFlashVirtualDriveAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageFlexFlashVirtualDriveAllOfWithDefaults

`func NewStorageFlexFlashVirtualDriveAllOfWithDefaults() *StorageFlexFlashVirtualDriveAllOf`

NewStorageFlexFlashVirtualDriveAllOfWithDefaults instantiates a new StorageFlexFlashVirtualDriveAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDriveScope

`func (o *StorageFlexFlashVirtualDriveAllOf) GetDriveScope() string`

GetDriveScope returns the DriveScope field if non-nil, zero value otherwise.

### GetDriveScopeOk

`func (o *StorageFlexFlashVirtualDriveAllOf) GetDriveScopeOk() (*string, bool)`

GetDriveScopeOk returns a tuple with the DriveScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriveScope

`func (o *StorageFlexFlashVirtualDriveAllOf) SetDriveScope(v string)`

SetDriveScope sets DriveScope field to given value.

### HasDriveScope

`func (o *StorageFlexFlashVirtualDriveAllOf) HasDriveScope() bool`

HasDriveScope returns a boolean if a field has been set.

### GetDriveStatus

`func (o *StorageFlexFlashVirtualDriveAllOf) GetDriveStatus() string`

GetDriveStatus returns the DriveStatus field if non-nil, zero value otherwise.

### GetDriveStatusOk

`func (o *StorageFlexFlashVirtualDriveAllOf) GetDriveStatusOk() (*string, bool)`

GetDriveStatusOk returns a tuple with the DriveStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriveStatus

`func (o *StorageFlexFlashVirtualDriveAllOf) SetDriveStatus(v string)`

SetDriveStatus sets DriveStatus field to given value.

### HasDriveStatus

`func (o *StorageFlexFlashVirtualDriveAllOf) HasDriveStatus() bool`

HasDriveStatus returns a boolean if a field has been set.

### GetPartitionId

`func (o *StorageFlexFlashVirtualDriveAllOf) GetPartitionId() string`

GetPartitionId returns the PartitionId field if non-nil, zero value otherwise.

### GetPartitionIdOk

`func (o *StorageFlexFlashVirtualDriveAllOf) GetPartitionIdOk() (*string, bool)`

GetPartitionIdOk returns a tuple with the PartitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionId

`func (o *StorageFlexFlashVirtualDriveAllOf) SetPartitionId(v string)`

SetPartitionId sets PartitionId field to given value.

### HasPartitionId

`func (o *StorageFlexFlashVirtualDriveAllOf) HasPartitionId() bool`

HasPartitionId returns a boolean if a field has been set.

### GetResidentImage

`func (o *StorageFlexFlashVirtualDriveAllOf) GetResidentImage() string`

GetResidentImage returns the ResidentImage field if non-nil, zero value otherwise.

### GetResidentImageOk

`func (o *StorageFlexFlashVirtualDriveAllOf) GetResidentImageOk() (*string, bool)`

GetResidentImageOk returns a tuple with the ResidentImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResidentImage

`func (o *StorageFlexFlashVirtualDriveAllOf) SetResidentImage(v string)`

SetResidentImage sets ResidentImage field to given value.

### HasResidentImage

`func (o *StorageFlexFlashVirtualDriveAllOf) HasResidentImage() bool`

HasResidentImage returns a boolean if a field has been set.

### GetSize

`func (o *StorageFlexFlashVirtualDriveAllOf) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *StorageFlexFlashVirtualDriveAllOf) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *StorageFlexFlashVirtualDriveAllOf) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *StorageFlexFlashVirtualDriveAllOf) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetVirtualDrive

`func (o *StorageFlexFlashVirtualDriveAllOf) GetVirtualDrive() string`

GetVirtualDrive returns the VirtualDrive field if non-nil, zero value otherwise.

### GetVirtualDriveOk

`func (o *StorageFlexFlashVirtualDriveAllOf) GetVirtualDriveOk() (*string, bool)`

GetVirtualDriveOk returns a tuple with the VirtualDrive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualDrive

`func (o *StorageFlexFlashVirtualDriveAllOf) SetVirtualDrive(v string)`

SetVirtualDrive sets VirtualDrive field to given value.

### HasVirtualDrive

`func (o *StorageFlexFlashVirtualDriveAllOf) HasVirtualDrive() bool`

HasVirtualDrive returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *StorageFlexFlashVirtualDriveAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *StorageFlexFlashVirtualDriveAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *StorageFlexFlashVirtualDriveAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *StorageFlexFlashVirtualDriveAllOf) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *StorageFlexFlashVirtualDriveAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StorageFlexFlashVirtualDriveAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StorageFlexFlashVirtualDriveAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StorageFlexFlashVirtualDriveAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### GetStorageFlexFlashController

`func (o *StorageFlexFlashVirtualDriveAllOf) GetStorageFlexFlashController() StorageFlexFlashControllerRelationship`

GetStorageFlexFlashController returns the StorageFlexFlashController field if non-nil, zero value otherwise.

### GetStorageFlexFlashControllerOk

`func (o *StorageFlexFlashVirtualDriveAllOf) GetStorageFlexFlashControllerOk() (*StorageFlexFlashControllerRelationship, bool)`

GetStorageFlexFlashControllerOk returns a tuple with the StorageFlexFlashController field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageFlexFlashController

`func (o *StorageFlexFlashVirtualDriveAllOf) SetStorageFlexFlashController(v StorageFlexFlashControllerRelationship)`

SetStorageFlexFlashController sets StorageFlexFlashController field to given value.

### HasStorageFlexFlashController

`func (o *StorageFlexFlashVirtualDriveAllOf) HasStorageFlexFlashController() bool`

HasStorageFlexFlashController returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# StorageFlexUtilVirtualDriveAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DriveStatus** | Pointer to **string** | Status of the Flex Util virtual drive. | [optional] 
**DriveType** | Pointer to **string** | Type of virtual drive managed by flex util controller. | [optional] 
**PartitionId** | Pointer to **string** | Disk Partition Id of virtual drive managed by flex util controller. | [optional] 
**PartitionName** | Pointer to **string** | Partition name of the Flex Util virtual drive. | [optional] 
**ResidentImage** | Pointer to **string** | The resident image on the flex util virtual Drive. | [optional] 
**Size** | Pointer to **string** | Size of the Flex Util virtual drive. | [optional] 
**VirtualDrive** | Pointer to **string** | Virtual drive on the Flex Util controller. | [optional] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](inventory.DeviceInfo.Relationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 
**StorageFlexUtilController** | Pointer to [**StorageFlexUtilControllerRelationship**](storage.FlexUtilController.Relationship.md) |  | [optional] 

## Methods

### NewStorageFlexUtilVirtualDriveAllOf

`func NewStorageFlexUtilVirtualDriveAllOf() *StorageFlexUtilVirtualDriveAllOf`

NewStorageFlexUtilVirtualDriveAllOf instantiates a new StorageFlexUtilVirtualDriveAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageFlexUtilVirtualDriveAllOfWithDefaults

`func NewStorageFlexUtilVirtualDriveAllOfWithDefaults() *StorageFlexUtilVirtualDriveAllOf`

NewStorageFlexUtilVirtualDriveAllOfWithDefaults instantiates a new StorageFlexUtilVirtualDriveAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDriveStatus

`func (o *StorageFlexUtilVirtualDriveAllOf) GetDriveStatus() string`

GetDriveStatus returns the DriveStatus field if non-nil, zero value otherwise.

### GetDriveStatusOk

`func (o *StorageFlexUtilVirtualDriveAllOf) GetDriveStatusOk() (*string, bool)`

GetDriveStatusOk returns a tuple with the DriveStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriveStatus

`func (o *StorageFlexUtilVirtualDriveAllOf) SetDriveStatus(v string)`

SetDriveStatus sets DriveStatus field to given value.

### HasDriveStatus

`func (o *StorageFlexUtilVirtualDriveAllOf) HasDriveStatus() bool`

HasDriveStatus returns a boolean if a field has been set.

### GetDriveType

`func (o *StorageFlexUtilVirtualDriveAllOf) GetDriveType() string`

GetDriveType returns the DriveType field if non-nil, zero value otherwise.

### GetDriveTypeOk

`func (o *StorageFlexUtilVirtualDriveAllOf) GetDriveTypeOk() (*string, bool)`

GetDriveTypeOk returns a tuple with the DriveType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriveType

`func (o *StorageFlexUtilVirtualDriveAllOf) SetDriveType(v string)`

SetDriveType sets DriveType field to given value.

### HasDriveType

`func (o *StorageFlexUtilVirtualDriveAllOf) HasDriveType() bool`

HasDriveType returns a boolean if a field has been set.

### GetPartitionId

`func (o *StorageFlexUtilVirtualDriveAllOf) GetPartitionId() string`

GetPartitionId returns the PartitionId field if non-nil, zero value otherwise.

### GetPartitionIdOk

`func (o *StorageFlexUtilVirtualDriveAllOf) GetPartitionIdOk() (*string, bool)`

GetPartitionIdOk returns a tuple with the PartitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionId

`func (o *StorageFlexUtilVirtualDriveAllOf) SetPartitionId(v string)`

SetPartitionId sets PartitionId field to given value.

### HasPartitionId

`func (o *StorageFlexUtilVirtualDriveAllOf) HasPartitionId() bool`

HasPartitionId returns a boolean if a field has been set.

### GetPartitionName

`func (o *StorageFlexUtilVirtualDriveAllOf) GetPartitionName() string`

GetPartitionName returns the PartitionName field if non-nil, zero value otherwise.

### GetPartitionNameOk

`func (o *StorageFlexUtilVirtualDriveAllOf) GetPartitionNameOk() (*string, bool)`

GetPartitionNameOk returns a tuple with the PartitionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionName

`func (o *StorageFlexUtilVirtualDriveAllOf) SetPartitionName(v string)`

SetPartitionName sets PartitionName field to given value.

### HasPartitionName

`func (o *StorageFlexUtilVirtualDriveAllOf) HasPartitionName() bool`

HasPartitionName returns a boolean if a field has been set.

### GetResidentImage

`func (o *StorageFlexUtilVirtualDriveAllOf) GetResidentImage() string`

GetResidentImage returns the ResidentImage field if non-nil, zero value otherwise.

### GetResidentImageOk

`func (o *StorageFlexUtilVirtualDriveAllOf) GetResidentImageOk() (*string, bool)`

GetResidentImageOk returns a tuple with the ResidentImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResidentImage

`func (o *StorageFlexUtilVirtualDriveAllOf) SetResidentImage(v string)`

SetResidentImage sets ResidentImage field to given value.

### HasResidentImage

`func (o *StorageFlexUtilVirtualDriveAllOf) HasResidentImage() bool`

HasResidentImage returns a boolean if a field has been set.

### GetSize

`func (o *StorageFlexUtilVirtualDriveAllOf) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *StorageFlexUtilVirtualDriveAllOf) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *StorageFlexUtilVirtualDriveAllOf) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *StorageFlexUtilVirtualDriveAllOf) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetVirtualDrive

`func (o *StorageFlexUtilVirtualDriveAllOf) GetVirtualDrive() string`

GetVirtualDrive returns the VirtualDrive field if non-nil, zero value otherwise.

### GetVirtualDriveOk

`func (o *StorageFlexUtilVirtualDriveAllOf) GetVirtualDriveOk() (*string, bool)`

GetVirtualDriveOk returns a tuple with the VirtualDrive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualDrive

`func (o *StorageFlexUtilVirtualDriveAllOf) SetVirtualDrive(v string)`

SetVirtualDrive sets VirtualDrive field to given value.

### HasVirtualDrive

`func (o *StorageFlexUtilVirtualDriveAllOf) HasVirtualDrive() bool`

HasVirtualDrive returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *StorageFlexUtilVirtualDriveAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *StorageFlexUtilVirtualDriveAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *StorageFlexUtilVirtualDriveAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *StorageFlexUtilVirtualDriveAllOf) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *StorageFlexUtilVirtualDriveAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StorageFlexUtilVirtualDriveAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StorageFlexUtilVirtualDriveAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StorageFlexUtilVirtualDriveAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### GetStorageFlexUtilController

`func (o *StorageFlexUtilVirtualDriveAllOf) GetStorageFlexUtilController() StorageFlexUtilControllerRelationship`

GetStorageFlexUtilController returns the StorageFlexUtilController field if non-nil, zero value otherwise.

### GetStorageFlexUtilControllerOk

`func (o *StorageFlexUtilVirtualDriveAllOf) GetStorageFlexUtilControllerOk() (*StorageFlexUtilControllerRelationship, bool)`

GetStorageFlexUtilControllerOk returns a tuple with the StorageFlexUtilController field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageFlexUtilController

`func (o *StorageFlexUtilVirtualDriveAllOf) SetStorageFlexUtilController(v StorageFlexUtilControllerRelationship)`

SetStorageFlexUtilController sets StorageFlexUtilController field to given value.

### HasStorageFlexUtilController

`func (o *StorageFlexUtilVirtualDriveAllOf) HasStorageFlexUtilController() bool`

HasStorageFlexUtilController returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



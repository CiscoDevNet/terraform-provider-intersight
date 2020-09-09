# StorageFlexUtilPhysicalDrive

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockSize** | Pointer to **string** | Block size of the FlexUtil Physical drive. | [optional] 
**Capacity** | Pointer to **string** | Capacity of the FlexUtil Physical drive. | [optional] 
**Controller** | Pointer to **string** | Type of the Physical Drive Controller. | [optional] 
**DrivesEnabled** | Pointer to **string** | The number of drives enabled in the FlexUtil Physical Drive. | [optional] 
**Health** | Pointer to **string** | Health of the FlexUtil Physical drive. | [optional] 
**ManufacturerDate** | Pointer to **string** | Manufacturing date of the FlexUtil Physical Drive. | [optional] 
**ManufacturerId** | Pointer to **string** | Manufacturer identity of the FlexUtil Physical Drive. | [optional] 
**OemId** | Pointer to **string** | The OEM Identifier of the FlexUtil physical drive. | [optional] 
**PartitionCount** | Pointer to **string** | The number of partitions present on the FlexUtil Physical Drive. | [optional] 
**PdStatus** | Pointer to **string** | Status of the FlexUtil Physical Drive. | [optional] 
**PhysicalDrive** | Pointer to **string** | The type of physical drive. Example - microSD. | [optional] 
**ProductName** | Pointer to **string** | Product name of the FlexUtil Physical Drive. | [optional] 
**ProductRevision** | Pointer to **string** | Product revision of the FlexUtil Physical Drive. | [optional] 
**ReadErrorCount** | Pointer to **string** | Read error count of the FlexUtil Physical Drive. | [optional] 
**ReadErrorThreshold** | Pointer to **string** | Read error threshold for FlexUtil Physical Drive. | [optional] 
**WriteEnabled** | Pointer to **string** | Write access state of the FlexUtil Physical Drive. | [optional] 
**WriteErrorCount** | Pointer to **string** | Write error count of the FlexUtil Physical Drive. | [optional] 
**WriteErrorThreshold** | Pointer to **string** | Write error threshold for FlexUtil Physical Drive. | [optional] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](inventory.DeviceInfo.Relationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 
**StorageFlexUtilController** | Pointer to [**StorageFlexUtilControllerRelationship**](storage.FlexUtilController.Relationship.md) |  | [optional] 

## Methods

### NewStorageFlexUtilPhysicalDrive

`func NewStorageFlexUtilPhysicalDrive() *StorageFlexUtilPhysicalDrive`

NewStorageFlexUtilPhysicalDrive instantiates a new StorageFlexUtilPhysicalDrive object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageFlexUtilPhysicalDriveWithDefaults

`func NewStorageFlexUtilPhysicalDriveWithDefaults() *StorageFlexUtilPhysicalDrive`

NewStorageFlexUtilPhysicalDriveWithDefaults instantiates a new StorageFlexUtilPhysicalDrive object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockSize

`func (o *StorageFlexUtilPhysicalDrive) GetBlockSize() string`

GetBlockSize returns the BlockSize field if non-nil, zero value otherwise.

### GetBlockSizeOk

`func (o *StorageFlexUtilPhysicalDrive) GetBlockSizeOk() (*string, bool)`

GetBlockSizeOk returns a tuple with the BlockSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockSize

`func (o *StorageFlexUtilPhysicalDrive) SetBlockSize(v string)`

SetBlockSize sets BlockSize field to given value.

### HasBlockSize

`func (o *StorageFlexUtilPhysicalDrive) HasBlockSize() bool`

HasBlockSize returns a boolean if a field has been set.

### GetCapacity

`func (o *StorageFlexUtilPhysicalDrive) GetCapacity() string`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *StorageFlexUtilPhysicalDrive) GetCapacityOk() (*string, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *StorageFlexUtilPhysicalDrive) SetCapacity(v string)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *StorageFlexUtilPhysicalDrive) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### GetController

`func (o *StorageFlexUtilPhysicalDrive) GetController() string`

GetController returns the Controller field if non-nil, zero value otherwise.

### GetControllerOk

`func (o *StorageFlexUtilPhysicalDrive) GetControllerOk() (*string, bool)`

GetControllerOk returns a tuple with the Controller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetController

`func (o *StorageFlexUtilPhysicalDrive) SetController(v string)`

SetController sets Controller field to given value.

### HasController

`func (o *StorageFlexUtilPhysicalDrive) HasController() bool`

HasController returns a boolean if a field has been set.

### GetDrivesEnabled

`func (o *StorageFlexUtilPhysicalDrive) GetDrivesEnabled() string`

GetDrivesEnabled returns the DrivesEnabled field if non-nil, zero value otherwise.

### GetDrivesEnabledOk

`func (o *StorageFlexUtilPhysicalDrive) GetDrivesEnabledOk() (*string, bool)`

GetDrivesEnabledOk returns a tuple with the DrivesEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrivesEnabled

`func (o *StorageFlexUtilPhysicalDrive) SetDrivesEnabled(v string)`

SetDrivesEnabled sets DrivesEnabled field to given value.

### HasDrivesEnabled

`func (o *StorageFlexUtilPhysicalDrive) HasDrivesEnabled() bool`

HasDrivesEnabled returns a boolean if a field has been set.

### GetHealth

`func (o *StorageFlexUtilPhysicalDrive) GetHealth() string`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *StorageFlexUtilPhysicalDrive) GetHealthOk() (*string, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *StorageFlexUtilPhysicalDrive) SetHealth(v string)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *StorageFlexUtilPhysicalDrive) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetManufacturerDate

`func (o *StorageFlexUtilPhysicalDrive) GetManufacturerDate() string`

GetManufacturerDate returns the ManufacturerDate field if non-nil, zero value otherwise.

### GetManufacturerDateOk

`func (o *StorageFlexUtilPhysicalDrive) GetManufacturerDateOk() (*string, bool)`

GetManufacturerDateOk returns a tuple with the ManufacturerDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturerDate

`func (o *StorageFlexUtilPhysicalDrive) SetManufacturerDate(v string)`

SetManufacturerDate sets ManufacturerDate field to given value.

### HasManufacturerDate

`func (o *StorageFlexUtilPhysicalDrive) HasManufacturerDate() bool`

HasManufacturerDate returns a boolean if a field has been set.

### GetManufacturerId

`func (o *StorageFlexUtilPhysicalDrive) GetManufacturerId() string`

GetManufacturerId returns the ManufacturerId field if non-nil, zero value otherwise.

### GetManufacturerIdOk

`func (o *StorageFlexUtilPhysicalDrive) GetManufacturerIdOk() (*string, bool)`

GetManufacturerIdOk returns a tuple with the ManufacturerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturerId

`func (o *StorageFlexUtilPhysicalDrive) SetManufacturerId(v string)`

SetManufacturerId sets ManufacturerId field to given value.

### HasManufacturerId

`func (o *StorageFlexUtilPhysicalDrive) HasManufacturerId() bool`

HasManufacturerId returns a boolean if a field has been set.

### GetOemId

`func (o *StorageFlexUtilPhysicalDrive) GetOemId() string`

GetOemId returns the OemId field if non-nil, zero value otherwise.

### GetOemIdOk

`func (o *StorageFlexUtilPhysicalDrive) GetOemIdOk() (*string, bool)`

GetOemIdOk returns a tuple with the OemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOemId

`func (o *StorageFlexUtilPhysicalDrive) SetOemId(v string)`

SetOemId sets OemId field to given value.

### HasOemId

`func (o *StorageFlexUtilPhysicalDrive) HasOemId() bool`

HasOemId returns a boolean if a field has been set.

### GetPartitionCount

`func (o *StorageFlexUtilPhysicalDrive) GetPartitionCount() string`

GetPartitionCount returns the PartitionCount field if non-nil, zero value otherwise.

### GetPartitionCountOk

`func (o *StorageFlexUtilPhysicalDrive) GetPartitionCountOk() (*string, bool)`

GetPartitionCountOk returns a tuple with the PartitionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionCount

`func (o *StorageFlexUtilPhysicalDrive) SetPartitionCount(v string)`

SetPartitionCount sets PartitionCount field to given value.

### HasPartitionCount

`func (o *StorageFlexUtilPhysicalDrive) HasPartitionCount() bool`

HasPartitionCount returns a boolean if a field has been set.

### GetPdStatus

`func (o *StorageFlexUtilPhysicalDrive) GetPdStatus() string`

GetPdStatus returns the PdStatus field if non-nil, zero value otherwise.

### GetPdStatusOk

`func (o *StorageFlexUtilPhysicalDrive) GetPdStatusOk() (*string, bool)`

GetPdStatusOk returns a tuple with the PdStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdStatus

`func (o *StorageFlexUtilPhysicalDrive) SetPdStatus(v string)`

SetPdStatus sets PdStatus field to given value.

### HasPdStatus

`func (o *StorageFlexUtilPhysicalDrive) HasPdStatus() bool`

HasPdStatus returns a boolean if a field has been set.

### GetPhysicalDrive

`func (o *StorageFlexUtilPhysicalDrive) GetPhysicalDrive() string`

GetPhysicalDrive returns the PhysicalDrive field if non-nil, zero value otherwise.

### GetPhysicalDriveOk

`func (o *StorageFlexUtilPhysicalDrive) GetPhysicalDriveOk() (*string, bool)`

GetPhysicalDriveOk returns a tuple with the PhysicalDrive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalDrive

`func (o *StorageFlexUtilPhysicalDrive) SetPhysicalDrive(v string)`

SetPhysicalDrive sets PhysicalDrive field to given value.

### HasPhysicalDrive

`func (o *StorageFlexUtilPhysicalDrive) HasPhysicalDrive() bool`

HasPhysicalDrive returns a boolean if a field has been set.

### GetProductName

`func (o *StorageFlexUtilPhysicalDrive) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *StorageFlexUtilPhysicalDrive) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *StorageFlexUtilPhysicalDrive) SetProductName(v string)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *StorageFlexUtilPhysicalDrive) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### GetProductRevision

`func (o *StorageFlexUtilPhysicalDrive) GetProductRevision() string`

GetProductRevision returns the ProductRevision field if non-nil, zero value otherwise.

### GetProductRevisionOk

`func (o *StorageFlexUtilPhysicalDrive) GetProductRevisionOk() (*string, bool)`

GetProductRevisionOk returns a tuple with the ProductRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductRevision

`func (o *StorageFlexUtilPhysicalDrive) SetProductRevision(v string)`

SetProductRevision sets ProductRevision field to given value.

### HasProductRevision

`func (o *StorageFlexUtilPhysicalDrive) HasProductRevision() bool`

HasProductRevision returns a boolean if a field has been set.

### GetReadErrorCount

`func (o *StorageFlexUtilPhysicalDrive) GetReadErrorCount() string`

GetReadErrorCount returns the ReadErrorCount field if non-nil, zero value otherwise.

### GetReadErrorCountOk

`func (o *StorageFlexUtilPhysicalDrive) GetReadErrorCountOk() (*string, bool)`

GetReadErrorCountOk returns a tuple with the ReadErrorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadErrorCount

`func (o *StorageFlexUtilPhysicalDrive) SetReadErrorCount(v string)`

SetReadErrorCount sets ReadErrorCount field to given value.

### HasReadErrorCount

`func (o *StorageFlexUtilPhysicalDrive) HasReadErrorCount() bool`

HasReadErrorCount returns a boolean if a field has been set.

### GetReadErrorThreshold

`func (o *StorageFlexUtilPhysicalDrive) GetReadErrorThreshold() string`

GetReadErrorThreshold returns the ReadErrorThreshold field if non-nil, zero value otherwise.

### GetReadErrorThresholdOk

`func (o *StorageFlexUtilPhysicalDrive) GetReadErrorThresholdOk() (*string, bool)`

GetReadErrorThresholdOk returns a tuple with the ReadErrorThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadErrorThreshold

`func (o *StorageFlexUtilPhysicalDrive) SetReadErrorThreshold(v string)`

SetReadErrorThreshold sets ReadErrorThreshold field to given value.

### HasReadErrorThreshold

`func (o *StorageFlexUtilPhysicalDrive) HasReadErrorThreshold() bool`

HasReadErrorThreshold returns a boolean if a field has been set.

### GetWriteEnabled

`func (o *StorageFlexUtilPhysicalDrive) GetWriteEnabled() string`

GetWriteEnabled returns the WriteEnabled field if non-nil, zero value otherwise.

### GetWriteEnabledOk

`func (o *StorageFlexUtilPhysicalDrive) GetWriteEnabledOk() (*string, bool)`

GetWriteEnabledOk returns a tuple with the WriteEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteEnabled

`func (o *StorageFlexUtilPhysicalDrive) SetWriteEnabled(v string)`

SetWriteEnabled sets WriteEnabled field to given value.

### HasWriteEnabled

`func (o *StorageFlexUtilPhysicalDrive) HasWriteEnabled() bool`

HasWriteEnabled returns a boolean if a field has been set.

### GetWriteErrorCount

`func (o *StorageFlexUtilPhysicalDrive) GetWriteErrorCount() string`

GetWriteErrorCount returns the WriteErrorCount field if non-nil, zero value otherwise.

### GetWriteErrorCountOk

`func (o *StorageFlexUtilPhysicalDrive) GetWriteErrorCountOk() (*string, bool)`

GetWriteErrorCountOk returns a tuple with the WriteErrorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteErrorCount

`func (o *StorageFlexUtilPhysicalDrive) SetWriteErrorCount(v string)`

SetWriteErrorCount sets WriteErrorCount field to given value.

### HasWriteErrorCount

`func (o *StorageFlexUtilPhysicalDrive) HasWriteErrorCount() bool`

HasWriteErrorCount returns a boolean if a field has been set.

### GetWriteErrorThreshold

`func (o *StorageFlexUtilPhysicalDrive) GetWriteErrorThreshold() string`

GetWriteErrorThreshold returns the WriteErrorThreshold field if non-nil, zero value otherwise.

### GetWriteErrorThresholdOk

`func (o *StorageFlexUtilPhysicalDrive) GetWriteErrorThresholdOk() (*string, bool)`

GetWriteErrorThresholdOk returns a tuple with the WriteErrorThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteErrorThreshold

`func (o *StorageFlexUtilPhysicalDrive) SetWriteErrorThreshold(v string)`

SetWriteErrorThreshold sets WriteErrorThreshold field to given value.

### HasWriteErrorThreshold

`func (o *StorageFlexUtilPhysicalDrive) HasWriteErrorThreshold() bool`

HasWriteErrorThreshold returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *StorageFlexUtilPhysicalDrive) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *StorageFlexUtilPhysicalDrive) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *StorageFlexUtilPhysicalDrive) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *StorageFlexUtilPhysicalDrive) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *StorageFlexUtilPhysicalDrive) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StorageFlexUtilPhysicalDrive) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StorageFlexUtilPhysicalDrive) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StorageFlexUtilPhysicalDrive) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### GetStorageFlexUtilController

`func (o *StorageFlexUtilPhysicalDrive) GetStorageFlexUtilController() StorageFlexUtilControllerRelationship`

GetStorageFlexUtilController returns the StorageFlexUtilController field if non-nil, zero value otherwise.

### GetStorageFlexUtilControllerOk

`func (o *StorageFlexUtilPhysicalDrive) GetStorageFlexUtilControllerOk() (*StorageFlexUtilControllerRelationship, bool)`

GetStorageFlexUtilControllerOk returns a tuple with the StorageFlexUtilController field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageFlexUtilController

`func (o *StorageFlexUtilPhysicalDrive) SetStorageFlexUtilController(v StorageFlexUtilControllerRelationship)`

SetStorageFlexUtilController sets StorageFlexUtilController field to given value.

### HasStorageFlexUtilController

`func (o *StorageFlexUtilPhysicalDrive) HasStorageFlexUtilController() bool`

HasStorageFlexUtilController returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



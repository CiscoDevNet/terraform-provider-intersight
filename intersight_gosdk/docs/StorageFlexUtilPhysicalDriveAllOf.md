# StorageFlexUtilPhysicalDriveAllOf

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

### NewStorageFlexUtilPhysicalDriveAllOf

`func NewStorageFlexUtilPhysicalDriveAllOf() *StorageFlexUtilPhysicalDriveAllOf`

NewStorageFlexUtilPhysicalDriveAllOf instantiates a new StorageFlexUtilPhysicalDriveAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageFlexUtilPhysicalDriveAllOfWithDefaults

`func NewStorageFlexUtilPhysicalDriveAllOfWithDefaults() *StorageFlexUtilPhysicalDriveAllOf`

NewStorageFlexUtilPhysicalDriveAllOfWithDefaults instantiates a new StorageFlexUtilPhysicalDriveAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockSize

`func (o *StorageFlexUtilPhysicalDriveAllOf) GetBlockSize() string`

GetBlockSize returns the BlockSize field if non-nil, zero value otherwise.

### GetBlockSizeOk

`func (o *StorageFlexUtilPhysicalDriveAllOf) GetBlockSizeOk() (*string, bool)`

GetBlockSizeOk returns a tuple with the BlockSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockSize

`func (o *StorageFlexUtilPhysicalDriveAllOf) SetBlockSize(v string)`

SetBlockSize sets BlockSize field to given value.

### HasBlockSize

`func (o *StorageFlexUtilPhysicalDriveAllOf) HasBlockSize() bool`

HasBlockSize returns a boolean if a field has been set.

### GetCapacity

`func (o *StorageFlexUtilPhysicalDriveAllOf) GetCapacity() string`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *StorageFlexUtilPhysicalDriveAllOf) GetCapacityOk() (*string, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *StorageFlexUtilPhysicalDriveAllOf) SetCapacity(v string)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *StorageFlexUtilPhysicalDriveAllOf) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### GetController

`func (o *StorageFlexUtilPhysicalDriveAllOf) GetController() string`

GetController returns the Controller field if non-nil, zero value otherwise.

### GetControllerOk

`func (o *StorageFlexUtilPhysicalDriveAllOf) GetControllerOk() (*string, bool)`

GetControllerOk returns a tuple with the Controller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetController

`func (o *StorageFlexUtilPhysicalDriveAllOf) SetController(v string)`

SetController sets Controller field to given value.

### HasController

`func (o *StorageFlexUtilPhysicalDriveAllOf) HasController() bool`

HasController returns a boolean if a field has been set.

### GetDrivesEnabled

`func (o *StorageFlexUtilPhysicalDriveAllOf) GetDrivesEnabled() string`

GetDrivesEnabled returns the DrivesEnabled field if non-nil, zero value otherwise.

### GetDrivesEnabledOk

`func (o *StorageFlexUtilPhysicalDriveAllOf) GetDrivesEnabledOk() (*string, bool)`

GetDrivesEnabledOk returns a tuple with the DrivesEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrivesEnabled

`func (o *StorageFlexUtilPhysicalDriveAllOf) SetDrivesEnabled(v string)`

SetDrivesEnabled sets DrivesEnabled field to given value.

### HasDrivesEnabled

`func (o *StorageFlexUtilPhysicalDriveAllOf) HasDrivesEnabled() bool`

HasDrivesEnabled returns a boolean if a field has been set.

### GetHealth

`func (o *StorageFlexUtilPhysicalDriveAllOf) GetHealth() string`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *StorageFlexUtilPhysicalDriveAllOf) GetHealthOk() (*string, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *StorageFlexUtilPhysicalDriveAllOf) SetHealth(v string)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *StorageFlexUtilPhysicalDriveAllOf) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetManufacturerDate

`func (o *StorageFlexUtilPhysicalDriveAllOf) GetManufacturerDate() string`

GetManufacturerDate returns the ManufacturerDate field if non-nil, zero value otherwise.

### GetManufacturerDateOk

`func (o *StorageFlexUtilPhysicalDriveAllOf) GetManufacturerDateOk() (*string, bool)`

GetManufacturerDateOk returns a tuple with the ManufacturerDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturerDate

`func (o *StorageFlexUtilPhysicalDriveAllOf) SetManufacturerDate(v string)`

SetManufacturerDate sets ManufacturerDate field to given value.

### HasManufacturerDate

`func (o *StorageFlexUtilPhysicalDriveAllOf) HasManufacturerDate() bool`

HasManufacturerDate returns a boolean if a field has been set.

### GetManufacturerId

`func (o *StorageFlexUtilPhysicalDriveAllOf) GetManufacturerId() string`

GetManufacturerId returns the ManufacturerId field if non-nil, zero value otherwise.

### GetManufacturerIdOk

`func (o *StorageFlexUtilPhysicalDriveAllOf) GetManufacturerIdOk() (*string, bool)`

GetManufacturerIdOk returns a tuple with the ManufacturerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturerId

`func (o *StorageFlexUtilPhysicalDriveAllOf) SetManufacturerId(v string)`

SetManufacturerId sets ManufacturerId field to given value.

### HasManufacturerId

`func (o *StorageFlexUtilPhysicalDriveAllOf) HasManufacturerId() bool`

HasManufacturerId returns a boolean if a field has been set.

### GetOemId

`func (o *StorageFlexUtilPhysicalDriveAllOf) GetOemId() string`

GetOemId returns the OemId field if non-nil, zero value otherwise.

### GetOemIdOk

`func (o *StorageFlexUtilPhysicalDriveAllOf) GetOemIdOk() (*string, bool)`

GetOemIdOk returns a tuple with the OemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOemId

`func (o *StorageFlexUtilPhysicalDriveAllOf) SetOemId(v string)`

SetOemId sets OemId field to given value.

### HasOemId

`func (o *StorageFlexUtilPhysicalDriveAllOf) HasOemId() bool`

HasOemId returns a boolean if a field has been set.

### GetPartitionCount

`func (o *StorageFlexUtilPhysicalDriveAllOf) GetPartitionCount() string`

GetPartitionCount returns the PartitionCount field if non-nil, zero value otherwise.

### GetPartitionCountOk

`func (o *StorageFlexUtilPhysicalDriveAllOf) GetPartitionCountOk() (*string, bool)`

GetPartitionCountOk returns a tuple with the PartitionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionCount

`func (o *StorageFlexUtilPhysicalDriveAllOf) SetPartitionCount(v string)`

SetPartitionCount sets PartitionCount field to given value.

### HasPartitionCount

`func (o *StorageFlexUtilPhysicalDriveAllOf) HasPartitionCount() bool`

HasPartitionCount returns a boolean if a field has been set.

### GetPdStatus

`func (o *StorageFlexUtilPhysicalDriveAllOf) GetPdStatus() string`

GetPdStatus returns the PdStatus field if non-nil, zero value otherwise.

### GetPdStatusOk

`func (o *StorageFlexUtilPhysicalDriveAllOf) GetPdStatusOk() (*string, bool)`

GetPdStatusOk returns a tuple with the PdStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdStatus

`func (o *StorageFlexUtilPhysicalDriveAllOf) SetPdStatus(v string)`

SetPdStatus sets PdStatus field to given value.

### HasPdStatus

`func (o *StorageFlexUtilPhysicalDriveAllOf) HasPdStatus() bool`

HasPdStatus returns a boolean if a field has been set.

### GetPhysicalDrive

`func (o *StorageFlexUtilPhysicalDriveAllOf) GetPhysicalDrive() string`

GetPhysicalDrive returns the PhysicalDrive field if non-nil, zero value otherwise.

### GetPhysicalDriveOk

`func (o *StorageFlexUtilPhysicalDriveAllOf) GetPhysicalDriveOk() (*string, bool)`

GetPhysicalDriveOk returns a tuple with the PhysicalDrive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalDrive

`func (o *StorageFlexUtilPhysicalDriveAllOf) SetPhysicalDrive(v string)`

SetPhysicalDrive sets PhysicalDrive field to given value.

### HasPhysicalDrive

`func (o *StorageFlexUtilPhysicalDriveAllOf) HasPhysicalDrive() bool`

HasPhysicalDrive returns a boolean if a field has been set.

### GetProductName

`func (o *StorageFlexUtilPhysicalDriveAllOf) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *StorageFlexUtilPhysicalDriveAllOf) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *StorageFlexUtilPhysicalDriveAllOf) SetProductName(v string)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *StorageFlexUtilPhysicalDriveAllOf) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### GetProductRevision

`func (o *StorageFlexUtilPhysicalDriveAllOf) GetProductRevision() string`

GetProductRevision returns the ProductRevision field if non-nil, zero value otherwise.

### GetProductRevisionOk

`func (o *StorageFlexUtilPhysicalDriveAllOf) GetProductRevisionOk() (*string, bool)`

GetProductRevisionOk returns a tuple with the ProductRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductRevision

`func (o *StorageFlexUtilPhysicalDriveAllOf) SetProductRevision(v string)`

SetProductRevision sets ProductRevision field to given value.

### HasProductRevision

`func (o *StorageFlexUtilPhysicalDriveAllOf) HasProductRevision() bool`

HasProductRevision returns a boolean if a field has been set.

### GetReadErrorCount

`func (o *StorageFlexUtilPhysicalDriveAllOf) GetReadErrorCount() string`

GetReadErrorCount returns the ReadErrorCount field if non-nil, zero value otherwise.

### GetReadErrorCountOk

`func (o *StorageFlexUtilPhysicalDriveAllOf) GetReadErrorCountOk() (*string, bool)`

GetReadErrorCountOk returns a tuple with the ReadErrorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadErrorCount

`func (o *StorageFlexUtilPhysicalDriveAllOf) SetReadErrorCount(v string)`

SetReadErrorCount sets ReadErrorCount field to given value.

### HasReadErrorCount

`func (o *StorageFlexUtilPhysicalDriveAllOf) HasReadErrorCount() bool`

HasReadErrorCount returns a boolean if a field has been set.

### GetReadErrorThreshold

`func (o *StorageFlexUtilPhysicalDriveAllOf) GetReadErrorThreshold() string`

GetReadErrorThreshold returns the ReadErrorThreshold field if non-nil, zero value otherwise.

### GetReadErrorThresholdOk

`func (o *StorageFlexUtilPhysicalDriveAllOf) GetReadErrorThresholdOk() (*string, bool)`

GetReadErrorThresholdOk returns a tuple with the ReadErrorThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadErrorThreshold

`func (o *StorageFlexUtilPhysicalDriveAllOf) SetReadErrorThreshold(v string)`

SetReadErrorThreshold sets ReadErrorThreshold field to given value.

### HasReadErrorThreshold

`func (o *StorageFlexUtilPhysicalDriveAllOf) HasReadErrorThreshold() bool`

HasReadErrorThreshold returns a boolean if a field has been set.

### GetWriteEnabled

`func (o *StorageFlexUtilPhysicalDriveAllOf) GetWriteEnabled() string`

GetWriteEnabled returns the WriteEnabled field if non-nil, zero value otherwise.

### GetWriteEnabledOk

`func (o *StorageFlexUtilPhysicalDriveAllOf) GetWriteEnabledOk() (*string, bool)`

GetWriteEnabledOk returns a tuple with the WriteEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteEnabled

`func (o *StorageFlexUtilPhysicalDriveAllOf) SetWriteEnabled(v string)`

SetWriteEnabled sets WriteEnabled field to given value.

### HasWriteEnabled

`func (o *StorageFlexUtilPhysicalDriveAllOf) HasWriteEnabled() bool`

HasWriteEnabled returns a boolean if a field has been set.

### GetWriteErrorCount

`func (o *StorageFlexUtilPhysicalDriveAllOf) GetWriteErrorCount() string`

GetWriteErrorCount returns the WriteErrorCount field if non-nil, zero value otherwise.

### GetWriteErrorCountOk

`func (o *StorageFlexUtilPhysicalDriveAllOf) GetWriteErrorCountOk() (*string, bool)`

GetWriteErrorCountOk returns a tuple with the WriteErrorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteErrorCount

`func (o *StorageFlexUtilPhysicalDriveAllOf) SetWriteErrorCount(v string)`

SetWriteErrorCount sets WriteErrorCount field to given value.

### HasWriteErrorCount

`func (o *StorageFlexUtilPhysicalDriveAllOf) HasWriteErrorCount() bool`

HasWriteErrorCount returns a boolean if a field has been set.

### GetWriteErrorThreshold

`func (o *StorageFlexUtilPhysicalDriveAllOf) GetWriteErrorThreshold() string`

GetWriteErrorThreshold returns the WriteErrorThreshold field if non-nil, zero value otherwise.

### GetWriteErrorThresholdOk

`func (o *StorageFlexUtilPhysicalDriveAllOf) GetWriteErrorThresholdOk() (*string, bool)`

GetWriteErrorThresholdOk returns a tuple with the WriteErrorThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteErrorThreshold

`func (o *StorageFlexUtilPhysicalDriveAllOf) SetWriteErrorThreshold(v string)`

SetWriteErrorThreshold sets WriteErrorThreshold field to given value.

### HasWriteErrorThreshold

`func (o *StorageFlexUtilPhysicalDriveAllOf) HasWriteErrorThreshold() bool`

HasWriteErrorThreshold returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *StorageFlexUtilPhysicalDriveAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *StorageFlexUtilPhysicalDriveAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *StorageFlexUtilPhysicalDriveAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *StorageFlexUtilPhysicalDriveAllOf) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *StorageFlexUtilPhysicalDriveAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StorageFlexUtilPhysicalDriveAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StorageFlexUtilPhysicalDriveAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StorageFlexUtilPhysicalDriveAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### GetStorageFlexUtilController

`func (o *StorageFlexUtilPhysicalDriveAllOf) GetStorageFlexUtilController() StorageFlexUtilControllerRelationship`

GetStorageFlexUtilController returns the StorageFlexUtilController field if non-nil, zero value otherwise.

### GetStorageFlexUtilControllerOk

`func (o *StorageFlexUtilPhysicalDriveAllOf) GetStorageFlexUtilControllerOk() (*StorageFlexUtilControllerRelationship, bool)`

GetStorageFlexUtilControllerOk returns a tuple with the StorageFlexUtilController field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageFlexUtilController

`func (o *StorageFlexUtilPhysicalDriveAllOf) SetStorageFlexUtilController(v StorageFlexUtilControllerRelationship)`

SetStorageFlexUtilController sets StorageFlexUtilController field to given value.

### HasStorageFlexUtilController

`func (o *StorageFlexUtilPhysicalDriveAllOf) HasStorageFlexUtilController() bool`

HasStorageFlexUtilController returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



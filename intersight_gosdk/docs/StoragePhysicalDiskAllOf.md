# StoragePhysicalDiskAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.PhysicalDisk"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.PhysicalDisk"]
**BackgroundOperations** | Pointer to **string** | List of background operations underway. | [optional] 
**BlockSize** | Pointer to **string** | The block size of the physical disk in bytes. | [optional] [readonly] 
**Bootable** | Pointer to **string** | This field identifies the disk drive as bootable if set to true. | [optional] [readonly] 
**ConfigurationCheckpoint** | Pointer to **string** | The current configuration checkpoint of the physical disk. | [optional] [readonly] 
**ConfigurationState** | Pointer to **string** | The current configuration state of the physical disk. | [optional] [readonly] 
**DisabledForRemoval** | Pointer to **bool** | The physical disk is disabled for removal. | [optional] 
**DiscoveredPath** | Pointer to **string** | The discovered path of the physical disk. | [optional] [readonly] 
**DiskId** | Pointer to **string** | This field identifies the ID assigned to physical disks. | [optional] [readonly] 
**DiskState** | Pointer to **string** | This field identifies the health of the disk. | [optional] [readonly] 
**DriveFirmware** | Pointer to **string** | This field identifies the disk firmware running in the disk. | [optional] 
**DriveState** | Pointer to **string** | The drive state as reported by the controller. | [optional] [readonly] 
**EncryptionStatus** | Pointer to **string** | Encryption status of the physical disk. | [optional] 
**FailurePredicted** | Pointer to **bool** | Possibility of physical disk failure. | [optional] 
**FdeCapable** | Pointer to **string** | Full-Disk Encryption capability parameter of the physical disk. | [optional] 
**HotSpareType** | Pointer to **string** | Type of hotspare configured on the physical disk. | [optional] 
**IndicatorLed** | Pointer to **string** | Status of the locator LED corresponding to the physical disk. | [optional] 
**LinkSpeed** | Pointer to **string** | The speed of the link between the drive and the controller. | [optional] [readonly] 
**LinkState** | Pointer to **string** | The current link state of the physical disk. | [optional] [readonly] 
**MaximumOperatingTemperature** | Pointer to **int64** | Maximum operating temperature of drive in Celsius. | [optional] 
**MediaErrorCount** | Pointer to **int64** | Media error count on the physical disk. | [optional] 
**Name** | Pointer to **string** | Detailed name of the physical disk. | [optional] 
**NonCoercedSizeBytes** | Pointer to **int64** | Physical disk non-coerced size in bytes. | [optional] 
**NumBlocks** | Pointer to **string** | The number of blocks present on the physical disk. | [optional] [readonly] 
**OperPowerState** | Pointer to **string** | Operational power of the physical disk. | [optional] [readonly] 
**OperQualifierReason** | Pointer to **string** | For certain states, indicates the reason why the operState is in that state. | [optional] [readonly] 
**Operability** | Pointer to **string** | This field identifies the disk operability of the disk. | [optional] [readonly] 
**OperatingTemperature** | Pointer to **int64** | Operating temperature of drive in Celsius. | [optional] 
**PercentLifeLeft** | Pointer to **int64** | Percentage of write cycles remaining in a solid state drive (SSD). | [optional] 
**PercentReservedCapacityConsumed** | Pointer to **int64** | Percentage of reserve capacity consumed. | [optional] 
**PerformancePercent** | Pointer to **int64** | Performance at which the device operating expressed in percentage. | [optional] 
**PhysicalBlockSize** | Pointer to **string** | The block size of the installed physical disk. | [optional] [readonly] 
**Pid** | Pointer to **string** | This field identifies the Product ID for physicalDisk. | [optional] [readonly] 
**PowerCycleCount** | Pointer to **int64** | Number of powercycles the drive has undergone. | [optional] 
**PowerOnHours** | Pointer to **int64** | Number of hours the drive has been powered on. | [optional] 
**PredictedMediaLifeLeftPercent** | Pointer to **int64** | Predicted physical disk life left in percentage. | [optional] 
**PredictiveFailureCount** | Pointer to **int64** | Error count on the physical disk. | [optional] 
**Protocol** | Pointer to **string** | This field identifies the disk protocol used for communication. | [optional] [readonly] 
**RawSize** | Pointer to **string** | The raw size of the physical disk in MB. | [optional] [readonly] 
**ReadErrorCountThreshold** | Pointer to **int64** | The number of read errors that are permitted while accessing the drive/card. | [optional] 
**ReadIoErrorCount** | Pointer to **int64** | Number of IO Errors that occured while reading data from the disk. | [optional] 
**Secured** | Pointer to **string** | This field identifies whether the disk is encrypted. | [optional] 
**Size** | Pointer to **string** | The size of the physical disk in MB. | [optional] [readonly] 
**Thermal** | Pointer to **string** | Thermal state of the physical disk. | [optional] [readonly] 
**ThresholdOperatingTemperature** | Pointer to **int64** | Rated threshold operating temperature in Celsius. | [optional] 
**Type** | Pointer to **string** | This field identifies the type of the physical disk. | [optional] [readonly] 
**VariantType** | Pointer to **string** | The variant type of the physical disk. | [optional] [readonly] 
**WearStatusInDays** | Pointer to **int64** | The number of days an SSD has gone through with the write cycles. | [optional] 
**WriteErrorCountThreshold** | Pointer to **int64** | The number of write errors that are permitted while accessing the drive/card. | [optional] 
**WriteIoErrorCount** | Pointer to **int64** | Number of IO Errors that occured while writing data to the disk. | [optional] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](InventoryDeviceInfoRelationship.md) |  | [optional] 
**LocatorLed** | Pointer to [**EquipmentLocatorLedRelationship**](EquipmentLocatorLedRelationship.md) |  | [optional] 
**PhysicalDiskExtensions** | Pointer to [**[]StoragePhysicalDiskExtensionRelationship**](StoragePhysicalDiskExtensionRelationship.md) | An array of relationships to storagePhysicalDiskExtension resources. | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 
**RunningFirmware** | Pointer to [**[]FirmwareRunningFirmwareRelationship**](FirmwareRunningFirmwareRelationship.md) | An array of relationships to firmwareRunningFirmware resources. | [optional] [readonly] 
**SasPorts** | Pointer to [**[]StorageSasPortRelationship**](StorageSasPortRelationship.md) | An array of relationships to storageSasPort resources. | [optional] [readonly] 
**StorageController** | Pointer to [**StorageControllerRelationship**](StorageControllerRelationship.md) |  | [optional] 
**StorageEnclosure** | Pointer to [**StorageEnclosureRelationship**](StorageEnclosureRelationship.md) |  | [optional] 

## Methods

### NewStoragePhysicalDiskAllOf

`func NewStoragePhysicalDiskAllOf(classId string, objectType string, ) *StoragePhysicalDiskAllOf`

NewStoragePhysicalDiskAllOf instantiates a new StoragePhysicalDiskAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoragePhysicalDiskAllOfWithDefaults

`func NewStoragePhysicalDiskAllOfWithDefaults() *StoragePhysicalDiskAllOf`

NewStoragePhysicalDiskAllOfWithDefaults instantiates a new StoragePhysicalDiskAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StoragePhysicalDiskAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StoragePhysicalDiskAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StoragePhysicalDiskAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StoragePhysicalDiskAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StoragePhysicalDiskAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StoragePhysicalDiskAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBackgroundOperations

`func (o *StoragePhysicalDiskAllOf) GetBackgroundOperations() string`

GetBackgroundOperations returns the BackgroundOperations field if non-nil, zero value otherwise.

### GetBackgroundOperationsOk

`func (o *StoragePhysicalDiskAllOf) GetBackgroundOperationsOk() (*string, bool)`

GetBackgroundOperationsOk returns a tuple with the BackgroundOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundOperations

`func (o *StoragePhysicalDiskAllOf) SetBackgroundOperations(v string)`

SetBackgroundOperations sets BackgroundOperations field to given value.

### HasBackgroundOperations

`func (o *StoragePhysicalDiskAllOf) HasBackgroundOperations() bool`

HasBackgroundOperations returns a boolean if a field has been set.

### GetBlockSize

`func (o *StoragePhysicalDiskAllOf) GetBlockSize() string`

GetBlockSize returns the BlockSize field if non-nil, zero value otherwise.

### GetBlockSizeOk

`func (o *StoragePhysicalDiskAllOf) GetBlockSizeOk() (*string, bool)`

GetBlockSizeOk returns a tuple with the BlockSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockSize

`func (o *StoragePhysicalDiskAllOf) SetBlockSize(v string)`

SetBlockSize sets BlockSize field to given value.

### HasBlockSize

`func (o *StoragePhysicalDiskAllOf) HasBlockSize() bool`

HasBlockSize returns a boolean if a field has been set.

### GetBootable

`func (o *StoragePhysicalDiskAllOf) GetBootable() string`

GetBootable returns the Bootable field if non-nil, zero value otherwise.

### GetBootableOk

`func (o *StoragePhysicalDiskAllOf) GetBootableOk() (*string, bool)`

GetBootableOk returns a tuple with the Bootable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootable

`func (o *StoragePhysicalDiskAllOf) SetBootable(v string)`

SetBootable sets Bootable field to given value.

### HasBootable

`func (o *StoragePhysicalDiskAllOf) HasBootable() bool`

HasBootable returns a boolean if a field has been set.

### GetConfigurationCheckpoint

`func (o *StoragePhysicalDiskAllOf) GetConfigurationCheckpoint() string`

GetConfigurationCheckpoint returns the ConfigurationCheckpoint field if non-nil, zero value otherwise.

### GetConfigurationCheckpointOk

`func (o *StoragePhysicalDiskAllOf) GetConfigurationCheckpointOk() (*string, bool)`

GetConfigurationCheckpointOk returns a tuple with the ConfigurationCheckpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationCheckpoint

`func (o *StoragePhysicalDiskAllOf) SetConfigurationCheckpoint(v string)`

SetConfigurationCheckpoint sets ConfigurationCheckpoint field to given value.

### HasConfigurationCheckpoint

`func (o *StoragePhysicalDiskAllOf) HasConfigurationCheckpoint() bool`

HasConfigurationCheckpoint returns a boolean if a field has been set.

### GetConfigurationState

`func (o *StoragePhysicalDiskAllOf) GetConfigurationState() string`

GetConfigurationState returns the ConfigurationState field if non-nil, zero value otherwise.

### GetConfigurationStateOk

`func (o *StoragePhysicalDiskAllOf) GetConfigurationStateOk() (*string, bool)`

GetConfigurationStateOk returns a tuple with the ConfigurationState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationState

`func (o *StoragePhysicalDiskAllOf) SetConfigurationState(v string)`

SetConfigurationState sets ConfigurationState field to given value.

### HasConfigurationState

`func (o *StoragePhysicalDiskAllOf) HasConfigurationState() bool`

HasConfigurationState returns a boolean if a field has been set.

### GetDisabledForRemoval

`func (o *StoragePhysicalDiskAllOf) GetDisabledForRemoval() bool`

GetDisabledForRemoval returns the DisabledForRemoval field if non-nil, zero value otherwise.

### GetDisabledForRemovalOk

`func (o *StoragePhysicalDiskAllOf) GetDisabledForRemovalOk() (*bool, bool)`

GetDisabledForRemovalOk returns a tuple with the DisabledForRemoval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledForRemoval

`func (o *StoragePhysicalDiskAllOf) SetDisabledForRemoval(v bool)`

SetDisabledForRemoval sets DisabledForRemoval field to given value.

### HasDisabledForRemoval

`func (o *StoragePhysicalDiskAllOf) HasDisabledForRemoval() bool`

HasDisabledForRemoval returns a boolean if a field has been set.

### GetDiscoveredPath

`func (o *StoragePhysicalDiskAllOf) GetDiscoveredPath() string`

GetDiscoveredPath returns the DiscoveredPath field if non-nil, zero value otherwise.

### GetDiscoveredPathOk

`func (o *StoragePhysicalDiskAllOf) GetDiscoveredPathOk() (*string, bool)`

GetDiscoveredPathOk returns a tuple with the DiscoveredPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveredPath

`func (o *StoragePhysicalDiskAllOf) SetDiscoveredPath(v string)`

SetDiscoveredPath sets DiscoveredPath field to given value.

### HasDiscoveredPath

`func (o *StoragePhysicalDiskAllOf) HasDiscoveredPath() bool`

HasDiscoveredPath returns a boolean if a field has been set.

### GetDiskId

`func (o *StoragePhysicalDiskAllOf) GetDiskId() string`

GetDiskId returns the DiskId field if non-nil, zero value otherwise.

### GetDiskIdOk

`func (o *StoragePhysicalDiskAllOf) GetDiskIdOk() (*string, bool)`

GetDiskIdOk returns a tuple with the DiskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskId

`func (o *StoragePhysicalDiskAllOf) SetDiskId(v string)`

SetDiskId sets DiskId field to given value.

### HasDiskId

`func (o *StoragePhysicalDiskAllOf) HasDiskId() bool`

HasDiskId returns a boolean if a field has been set.

### GetDiskState

`func (o *StoragePhysicalDiskAllOf) GetDiskState() string`

GetDiskState returns the DiskState field if non-nil, zero value otherwise.

### GetDiskStateOk

`func (o *StoragePhysicalDiskAllOf) GetDiskStateOk() (*string, bool)`

GetDiskStateOk returns a tuple with the DiskState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskState

`func (o *StoragePhysicalDiskAllOf) SetDiskState(v string)`

SetDiskState sets DiskState field to given value.

### HasDiskState

`func (o *StoragePhysicalDiskAllOf) HasDiskState() bool`

HasDiskState returns a boolean if a field has been set.

### GetDriveFirmware

`func (o *StoragePhysicalDiskAllOf) GetDriveFirmware() string`

GetDriveFirmware returns the DriveFirmware field if non-nil, zero value otherwise.

### GetDriveFirmwareOk

`func (o *StoragePhysicalDiskAllOf) GetDriveFirmwareOk() (*string, bool)`

GetDriveFirmwareOk returns a tuple with the DriveFirmware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriveFirmware

`func (o *StoragePhysicalDiskAllOf) SetDriveFirmware(v string)`

SetDriveFirmware sets DriveFirmware field to given value.

### HasDriveFirmware

`func (o *StoragePhysicalDiskAllOf) HasDriveFirmware() bool`

HasDriveFirmware returns a boolean if a field has been set.

### GetDriveState

`func (o *StoragePhysicalDiskAllOf) GetDriveState() string`

GetDriveState returns the DriveState field if non-nil, zero value otherwise.

### GetDriveStateOk

`func (o *StoragePhysicalDiskAllOf) GetDriveStateOk() (*string, bool)`

GetDriveStateOk returns a tuple with the DriveState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriveState

`func (o *StoragePhysicalDiskAllOf) SetDriveState(v string)`

SetDriveState sets DriveState field to given value.

### HasDriveState

`func (o *StoragePhysicalDiskAllOf) HasDriveState() bool`

HasDriveState returns a boolean if a field has been set.

### GetEncryptionStatus

`func (o *StoragePhysicalDiskAllOf) GetEncryptionStatus() string`

GetEncryptionStatus returns the EncryptionStatus field if non-nil, zero value otherwise.

### GetEncryptionStatusOk

`func (o *StoragePhysicalDiskAllOf) GetEncryptionStatusOk() (*string, bool)`

GetEncryptionStatusOk returns a tuple with the EncryptionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionStatus

`func (o *StoragePhysicalDiskAllOf) SetEncryptionStatus(v string)`

SetEncryptionStatus sets EncryptionStatus field to given value.

### HasEncryptionStatus

`func (o *StoragePhysicalDiskAllOf) HasEncryptionStatus() bool`

HasEncryptionStatus returns a boolean if a field has been set.

### GetFailurePredicted

`func (o *StoragePhysicalDiskAllOf) GetFailurePredicted() bool`

GetFailurePredicted returns the FailurePredicted field if non-nil, zero value otherwise.

### GetFailurePredictedOk

`func (o *StoragePhysicalDiskAllOf) GetFailurePredictedOk() (*bool, bool)`

GetFailurePredictedOk returns a tuple with the FailurePredicted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailurePredicted

`func (o *StoragePhysicalDiskAllOf) SetFailurePredicted(v bool)`

SetFailurePredicted sets FailurePredicted field to given value.

### HasFailurePredicted

`func (o *StoragePhysicalDiskAllOf) HasFailurePredicted() bool`

HasFailurePredicted returns a boolean if a field has been set.

### GetFdeCapable

`func (o *StoragePhysicalDiskAllOf) GetFdeCapable() string`

GetFdeCapable returns the FdeCapable field if non-nil, zero value otherwise.

### GetFdeCapableOk

`func (o *StoragePhysicalDiskAllOf) GetFdeCapableOk() (*string, bool)`

GetFdeCapableOk returns a tuple with the FdeCapable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFdeCapable

`func (o *StoragePhysicalDiskAllOf) SetFdeCapable(v string)`

SetFdeCapable sets FdeCapable field to given value.

### HasFdeCapable

`func (o *StoragePhysicalDiskAllOf) HasFdeCapable() bool`

HasFdeCapable returns a boolean if a field has been set.

### GetHotSpareType

`func (o *StoragePhysicalDiskAllOf) GetHotSpareType() string`

GetHotSpareType returns the HotSpareType field if non-nil, zero value otherwise.

### GetHotSpareTypeOk

`func (o *StoragePhysicalDiskAllOf) GetHotSpareTypeOk() (*string, bool)`

GetHotSpareTypeOk returns a tuple with the HotSpareType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHotSpareType

`func (o *StoragePhysicalDiskAllOf) SetHotSpareType(v string)`

SetHotSpareType sets HotSpareType field to given value.

### HasHotSpareType

`func (o *StoragePhysicalDiskAllOf) HasHotSpareType() bool`

HasHotSpareType returns a boolean if a field has been set.

### GetIndicatorLed

`func (o *StoragePhysicalDiskAllOf) GetIndicatorLed() string`

GetIndicatorLed returns the IndicatorLed field if non-nil, zero value otherwise.

### GetIndicatorLedOk

`func (o *StoragePhysicalDiskAllOf) GetIndicatorLedOk() (*string, bool)`

GetIndicatorLedOk returns a tuple with the IndicatorLed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndicatorLed

`func (o *StoragePhysicalDiskAllOf) SetIndicatorLed(v string)`

SetIndicatorLed sets IndicatorLed field to given value.

### HasIndicatorLed

`func (o *StoragePhysicalDiskAllOf) HasIndicatorLed() bool`

HasIndicatorLed returns a boolean if a field has been set.

### GetLinkSpeed

`func (o *StoragePhysicalDiskAllOf) GetLinkSpeed() string`

GetLinkSpeed returns the LinkSpeed field if non-nil, zero value otherwise.

### GetLinkSpeedOk

`func (o *StoragePhysicalDiskAllOf) GetLinkSpeedOk() (*string, bool)`

GetLinkSpeedOk returns a tuple with the LinkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkSpeed

`func (o *StoragePhysicalDiskAllOf) SetLinkSpeed(v string)`

SetLinkSpeed sets LinkSpeed field to given value.

### HasLinkSpeed

`func (o *StoragePhysicalDiskAllOf) HasLinkSpeed() bool`

HasLinkSpeed returns a boolean if a field has been set.

### GetLinkState

`func (o *StoragePhysicalDiskAllOf) GetLinkState() string`

GetLinkState returns the LinkState field if non-nil, zero value otherwise.

### GetLinkStateOk

`func (o *StoragePhysicalDiskAllOf) GetLinkStateOk() (*string, bool)`

GetLinkStateOk returns a tuple with the LinkState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkState

`func (o *StoragePhysicalDiskAllOf) SetLinkState(v string)`

SetLinkState sets LinkState field to given value.

### HasLinkState

`func (o *StoragePhysicalDiskAllOf) HasLinkState() bool`

HasLinkState returns a boolean if a field has been set.

### GetMaximumOperatingTemperature

`func (o *StoragePhysicalDiskAllOf) GetMaximumOperatingTemperature() int64`

GetMaximumOperatingTemperature returns the MaximumOperatingTemperature field if non-nil, zero value otherwise.

### GetMaximumOperatingTemperatureOk

`func (o *StoragePhysicalDiskAllOf) GetMaximumOperatingTemperatureOk() (*int64, bool)`

GetMaximumOperatingTemperatureOk returns a tuple with the MaximumOperatingTemperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumOperatingTemperature

`func (o *StoragePhysicalDiskAllOf) SetMaximumOperatingTemperature(v int64)`

SetMaximumOperatingTemperature sets MaximumOperatingTemperature field to given value.

### HasMaximumOperatingTemperature

`func (o *StoragePhysicalDiskAllOf) HasMaximumOperatingTemperature() bool`

HasMaximumOperatingTemperature returns a boolean if a field has been set.

### GetMediaErrorCount

`func (o *StoragePhysicalDiskAllOf) GetMediaErrorCount() int64`

GetMediaErrorCount returns the MediaErrorCount field if non-nil, zero value otherwise.

### GetMediaErrorCountOk

`func (o *StoragePhysicalDiskAllOf) GetMediaErrorCountOk() (*int64, bool)`

GetMediaErrorCountOk returns a tuple with the MediaErrorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaErrorCount

`func (o *StoragePhysicalDiskAllOf) SetMediaErrorCount(v int64)`

SetMediaErrorCount sets MediaErrorCount field to given value.

### HasMediaErrorCount

`func (o *StoragePhysicalDiskAllOf) HasMediaErrorCount() bool`

HasMediaErrorCount returns a boolean if a field has been set.

### GetName

`func (o *StoragePhysicalDiskAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StoragePhysicalDiskAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StoragePhysicalDiskAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StoragePhysicalDiskAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNonCoercedSizeBytes

`func (o *StoragePhysicalDiskAllOf) GetNonCoercedSizeBytes() int64`

GetNonCoercedSizeBytes returns the NonCoercedSizeBytes field if non-nil, zero value otherwise.

### GetNonCoercedSizeBytesOk

`func (o *StoragePhysicalDiskAllOf) GetNonCoercedSizeBytesOk() (*int64, bool)`

GetNonCoercedSizeBytesOk returns a tuple with the NonCoercedSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonCoercedSizeBytes

`func (o *StoragePhysicalDiskAllOf) SetNonCoercedSizeBytes(v int64)`

SetNonCoercedSizeBytes sets NonCoercedSizeBytes field to given value.

### HasNonCoercedSizeBytes

`func (o *StoragePhysicalDiskAllOf) HasNonCoercedSizeBytes() bool`

HasNonCoercedSizeBytes returns a boolean if a field has been set.

### GetNumBlocks

`func (o *StoragePhysicalDiskAllOf) GetNumBlocks() string`

GetNumBlocks returns the NumBlocks field if non-nil, zero value otherwise.

### GetNumBlocksOk

`func (o *StoragePhysicalDiskAllOf) GetNumBlocksOk() (*string, bool)`

GetNumBlocksOk returns a tuple with the NumBlocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumBlocks

`func (o *StoragePhysicalDiskAllOf) SetNumBlocks(v string)`

SetNumBlocks sets NumBlocks field to given value.

### HasNumBlocks

`func (o *StoragePhysicalDiskAllOf) HasNumBlocks() bool`

HasNumBlocks returns a boolean if a field has been set.

### GetOperPowerState

`func (o *StoragePhysicalDiskAllOf) GetOperPowerState() string`

GetOperPowerState returns the OperPowerState field if non-nil, zero value otherwise.

### GetOperPowerStateOk

`func (o *StoragePhysicalDiskAllOf) GetOperPowerStateOk() (*string, bool)`

GetOperPowerStateOk returns a tuple with the OperPowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperPowerState

`func (o *StoragePhysicalDiskAllOf) SetOperPowerState(v string)`

SetOperPowerState sets OperPowerState field to given value.

### HasOperPowerState

`func (o *StoragePhysicalDiskAllOf) HasOperPowerState() bool`

HasOperPowerState returns a boolean if a field has been set.

### GetOperQualifierReason

`func (o *StoragePhysicalDiskAllOf) GetOperQualifierReason() string`

GetOperQualifierReason returns the OperQualifierReason field if non-nil, zero value otherwise.

### GetOperQualifierReasonOk

`func (o *StoragePhysicalDiskAllOf) GetOperQualifierReasonOk() (*string, bool)`

GetOperQualifierReasonOk returns a tuple with the OperQualifierReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperQualifierReason

`func (o *StoragePhysicalDiskAllOf) SetOperQualifierReason(v string)`

SetOperQualifierReason sets OperQualifierReason field to given value.

### HasOperQualifierReason

`func (o *StoragePhysicalDiskAllOf) HasOperQualifierReason() bool`

HasOperQualifierReason returns a boolean if a field has been set.

### GetOperability

`func (o *StoragePhysicalDiskAllOf) GetOperability() string`

GetOperability returns the Operability field if non-nil, zero value otherwise.

### GetOperabilityOk

`func (o *StoragePhysicalDiskAllOf) GetOperabilityOk() (*string, bool)`

GetOperabilityOk returns a tuple with the Operability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperability

`func (o *StoragePhysicalDiskAllOf) SetOperability(v string)`

SetOperability sets Operability field to given value.

### HasOperability

`func (o *StoragePhysicalDiskAllOf) HasOperability() bool`

HasOperability returns a boolean if a field has been set.

### GetOperatingTemperature

`func (o *StoragePhysicalDiskAllOf) GetOperatingTemperature() int64`

GetOperatingTemperature returns the OperatingTemperature field if non-nil, zero value otherwise.

### GetOperatingTemperatureOk

`func (o *StoragePhysicalDiskAllOf) GetOperatingTemperatureOk() (*int64, bool)`

GetOperatingTemperatureOk returns a tuple with the OperatingTemperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingTemperature

`func (o *StoragePhysicalDiskAllOf) SetOperatingTemperature(v int64)`

SetOperatingTemperature sets OperatingTemperature field to given value.

### HasOperatingTemperature

`func (o *StoragePhysicalDiskAllOf) HasOperatingTemperature() bool`

HasOperatingTemperature returns a boolean if a field has been set.

### GetPercentLifeLeft

`func (o *StoragePhysicalDiskAllOf) GetPercentLifeLeft() int64`

GetPercentLifeLeft returns the PercentLifeLeft field if non-nil, zero value otherwise.

### GetPercentLifeLeftOk

`func (o *StoragePhysicalDiskAllOf) GetPercentLifeLeftOk() (*int64, bool)`

GetPercentLifeLeftOk returns a tuple with the PercentLifeLeft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentLifeLeft

`func (o *StoragePhysicalDiskAllOf) SetPercentLifeLeft(v int64)`

SetPercentLifeLeft sets PercentLifeLeft field to given value.

### HasPercentLifeLeft

`func (o *StoragePhysicalDiskAllOf) HasPercentLifeLeft() bool`

HasPercentLifeLeft returns a boolean if a field has been set.

### GetPercentReservedCapacityConsumed

`func (o *StoragePhysicalDiskAllOf) GetPercentReservedCapacityConsumed() int64`

GetPercentReservedCapacityConsumed returns the PercentReservedCapacityConsumed field if non-nil, zero value otherwise.

### GetPercentReservedCapacityConsumedOk

`func (o *StoragePhysicalDiskAllOf) GetPercentReservedCapacityConsumedOk() (*int64, bool)`

GetPercentReservedCapacityConsumedOk returns a tuple with the PercentReservedCapacityConsumed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentReservedCapacityConsumed

`func (o *StoragePhysicalDiskAllOf) SetPercentReservedCapacityConsumed(v int64)`

SetPercentReservedCapacityConsumed sets PercentReservedCapacityConsumed field to given value.

### HasPercentReservedCapacityConsumed

`func (o *StoragePhysicalDiskAllOf) HasPercentReservedCapacityConsumed() bool`

HasPercentReservedCapacityConsumed returns a boolean if a field has been set.

### GetPerformancePercent

`func (o *StoragePhysicalDiskAllOf) GetPerformancePercent() int64`

GetPerformancePercent returns the PerformancePercent field if non-nil, zero value otherwise.

### GetPerformancePercentOk

`func (o *StoragePhysicalDiskAllOf) GetPerformancePercentOk() (*int64, bool)`

GetPerformancePercentOk returns a tuple with the PerformancePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformancePercent

`func (o *StoragePhysicalDiskAllOf) SetPerformancePercent(v int64)`

SetPerformancePercent sets PerformancePercent field to given value.

### HasPerformancePercent

`func (o *StoragePhysicalDiskAllOf) HasPerformancePercent() bool`

HasPerformancePercent returns a boolean if a field has been set.

### GetPhysicalBlockSize

`func (o *StoragePhysicalDiskAllOf) GetPhysicalBlockSize() string`

GetPhysicalBlockSize returns the PhysicalBlockSize field if non-nil, zero value otherwise.

### GetPhysicalBlockSizeOk

`func (o *StoragePhysicalDiskAllOf) GetPhysicalBlockSizeOk() (*string, bool)`

GetPhysicalBlockSizeOk returns a tuple with the PhysicalBlockSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalBlockSize

`func (o *StoragePhysicalDiskAllOf) SetPhysicalBlockSize(v string)`

SetPhysicalBlockSize sets PhysicalBlockSize field to given value.

### HasPhysicalBlockSize

`func (o *StoragePhysicalDiskAllOf) HasPhysicalBlockSize() bool`

HasPhysicalBlockSize returns a boolean if a field has been set.

### GetPid

`func (o *StoragePhysicalDiskAllOf) GetPid() string`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *StoragePhysicalDiskAllOf) GetPidOk() (*string, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *StoragePhysicalDiskAllOf) SetPid(v string)`

SetPid sets Pid field to given value.

### HasPid

`func (o *StoragePhysicalDiskAllOf) HasPid() bool`

HasPid returns a boolean if a field has been set.

### GetPowerCycleCount

`func (o *StoragePhysicalDiskAllOf) GetPowerCycleCount() int64`

GetPowerCycleCount returns the PowerCycleCount field if non-nil, zero value otherwise.

### GetPowerCycleCountOk

`func (o *StoragePhysicalDiskAllOf) GetPowerCycleCountOk() (*int64, bool)`

GetPowerCycleCountOk returns a tuple with the PowerCycleCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerCycleCount

`func (o *StoragePhysicalDiskAllOf) SetPowerCycleCount(v int64)`

SetPowerCycleCount sets PowerCycleCount field to given value.

### HasPowerCycleCount

`func (o *StoragePhysicalDiskAllOf) HasPowerCycleCount() bool`

HasPowerCycleCount returns a boolean if a field has been set.

### GetPowerOnHours

`func (o *StoragePhysicalDiskAllOf) GetPowerOnHours() int64`

GetPowerOnHours returns the PowerOnHours field if non-nil, zero value otherwise.

### GetPowerOnHoursOk

`func (o *StoragePhysicalDiskAllOf) GetPowerOnHoursOk() (*int64, bool)`

GetPowerOnHoursOk returns a tuple with the PowerOnHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerOnHours

`func (o *StoragePhysicalDiskAllOf) SetPowerOnHours(v int64)`

SetPowerOnHours sets PowerOnHours field to given value.

### HasPowerOnHours

`func (o *StoragePhysicalDiskAllOf) HasPowerOnHours() bool`

HasPowerOnHours returns a boolean if a field has been set.

### GetPredictedMediaLifeLeftPercent

`func (o *StoragePhysicalDiskAllOf) GetPredictedMediaLifeLeftPercent() int64`

GetPredictedMediaLifeLeftPercent returns the PredictedMediaLifeLeftPercent field if non-nil, zero value otherwise.

### GetPredictedMediaLifeLeftPercentOk

`func (o *StoragePhysicalDiskAllOf) GetPredictedMediaLifeLeftPercentOk() (*int64, bool)`

GetPredictedMediaLifeLeftPercentOk returns a tuple with the PredictedMediaLifeLeftPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredictedMediaLifeLeftPercent

`func (o *StoragePhysicalDiskAllOf) SetPredictedMediaLifeLeftPercent(v int64)`

SetPredictedMediaLifeLeftPercent sets PredictedMediaLifeLeftPercent field to given value.

### HasPredictedMediaLifeLeftPercent

`func (o *StoragePhysicalDiskAllOf) HasPredictedMediaLifeLeftPercent() bool`

HasPredictedMediaLifeLeftPercent returns a boolean if a field has been set.

### GetPredictiveFailureCount

`func (o *StoragePhysicalDiskAllOf) GetPredictiveFailureCount() int64`

GetPredictiveFailureCount returns the PredictiveFailureCount field if non-nil, zero value otherwise.

### GetPredictiveFailureCountOk

`func (o *StoragePhysicalDiskAllOf) GetPredictiveFailureCountOk() (*int64, bool)`

GetPredictiveFailureCountOk returns a tuple with the PredictiveFailureCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredictiveFailureCount

`func (o *StoragePhysicalDiskAllOf) SetPredictiveFailureCount(v int64)`

SetPredictiveFailureCount sets PredictiveFailureCount field to given value.

### HasPredictiveFailureCount

`func (o *StoragePhysicalDiskAllOf) HasPredictiveFailureCount() bool`

HasPredictiveFailureCount returns a boolean if a field has been set.

### GetProtocol

`func (o *StoragePhysicalDiskAllOf) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *StoragePhysicalDiskAllOf) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *StoragePhysicalDiskAllOf) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *StoragePhysicalDiskAllOf) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetRawSize

`func (o *StoragePhysicalDiskAllOf) GetRawSize() string`

GetRawSize returns the RawSize field if non-nil, zero value otherwise.

### GetRawSizeOk

`func (o *StoragePhysicalDiskAllOf) GetRawSizeOk() (*string, bool)`

GetRawSizeOk returns a tuple with the RawSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawSize

`func (o *StoragePhysicalDiskAllOf) SetRawSize(v string)`

SetRawSize sets RawSize field to given value.

### HasRawSize

`func (o *StoragePhysicalDiskAllOf) HasRawSize() bool`

HasRawSize returns a boolean if a field has been set.

### GetReadErrorCountThreshold

`func (o *StoragePhysicalDiskAllOf) GetReadErrorCountThreshold() int64`

GetReadErrorCountThreshold returns the ReadErrorCountThreshold field if non-nil, zero value otherwise.

### GetReadErrorCountThresholdOk

`func (o *StoragePhysicalDiskAllOf) GetReadErrorCountThresholdOk() (*int64, bool)`

GetReadErrorCountThresholdOk returns a tuple with the ReadErrorCountThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadErrorCountThreshold

`func (o *StoragePhysicalDiskAllOf) SetReadErrorCountThreshold(v int64)`

SetReadErrorCountThreshold sets ReadErrorCountThreshold field to given value.

### HasReadErrorCountThreshold

`func (o *StoragePhysicalDiskAllOf) HasReadErrorCountThreshold() bool`

HasReadErrorCountThreshold returns a boolean if a field has been set.

### GetReadIoErrorCount

`func (o *StoragePhysicalDiskAllOf) GetReadIoErrorCount() int64`

GetReadIoErrorCount returns the ReadIoErrorCount field if non-nil, zero value otherwise.

### GetReadIoErrorCountOk

`func (o *StoragePhysicalDiskAllOf) GetReadIoErrorCountOk() (*int64, bool)`

GetReadIoErrorCountOk returns a tuple with the ReadIoErrorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadIoErrorCount

`func (o *StoragePhysicalDiskAllOf) SetReadIoErrorCount(v int64)`

SetReadIoErrorCount sets ReadIoErrorCount field to given value.

### HasReadIoErrorCount

`func (o *StoragePhysicalDiskAllOf) HasReadIoErrorCount() bool`

HasReadIoErrorCount returns a boolean if a field has been set.

### GetSecured

`func (o *StoragePhysicalDiskAllOf) GetSecured() string`

GetSecured returns the Secured field if non-nil, zero value otherwise.

### GetSecuredOk

`func (o *StoragePhysicalDiskAllOf) GetSecuredOk() (*string, bool)`

GetSecuredOk returns a tuple with the Secured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecured

`func (o *StoragePhysicalDiskAllOf) SetSecured(v string)`

SetSecured sets Secured field to given value.

### HasSecured

`func (o *StoragePhysicalDiskAllOf) HasSecured() bool`

HasSecured returns a boolean if a field has been set.

### GetSize

`func (o *StoragePhysicalDiskAllOf) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *StoragePhysicalDiskAllOf) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *StoragePhysicalDiskAllOf) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *StoragePhysicalDiskAllOf) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetThermal

`func (o *StoragePhysicalDiskAllOf) GetThermal() string`

GetThermal returns the Thermal field if non-nil, zero value otherwise.

### GetThermalOk

`func (o *StoragePhysicalDiskAllOf) GetThermalOk() (*string, bool)`

GetThermalOk returns a tuple with the Thermal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThermal

`func (o *StoragePhysicalDiskAllOf) SetThermal(v string)`

SetThermal sets Thermal field to given value.

### HasThermal

`func (o *StoragePhysicalDiskAllOf) HasThermal() bool`

HasThermal returns a boolean if a field has been set.

### GetThresholdOperatingTemperature

`func (o *StoragePhysicalDiskAllOf) GetThresholdOperatingTemperature() int64`

GetThresholdOperatingTemperature returns the ThresholdOperatingTemperature field if non-nil, zero value otherwise.

### GetThresholdOperatingTemperatureOk

`func (o *StoragePhysicalDiskAllOf) GetThresholdOperatingTemperatureOk() (*int64, bool)`

GetThresholdOperatingTemperatureOk returns a tuple with the ThresholdOperatingTemperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdOperatingTemperature

`func (o *StoragePhysicalDiskAllOf) SetThresholdOperatingTemperature(v int64)`

SetThresholdOperatingTemperature sets ThresholdOperatingTemperature field to given value.

### HasThresholdOperatingTemperature

`func (o *StoragePhysicalDiskAllOf) HasThresholdOperatingTemperature() bool`

HasThresholdOperatingTemperature returns a boolean if a field has been set.

### GetType

`func (o *StoragePhysicalDiskAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StoragePhysicalDiskAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StoragePhysicalDiskAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *StoragePhysicalDiskAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVariantType

`func (o *StoragePhysicalDiskAllOf) GetVariantType() string`

GetVariantType returns the VariantType field if non-nil, zero value otherwise.

### GetVariantTypeOk

`func (o *StoragePhysicalDiskAllOf) GetVariantTypeOk() (*string, bool)`

GetVariantTypeOk returns a tuple with the VariantType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariantType

`func (o *StoragePhysicalDiskAllOf) SetVariantType(v string)`

SetVariantType sets VariantType field to given value.

### HasVariantType

`func (o *StoragePhysicalDiskAllOf) HasVariantType() bool`

HasVariantType returns a boolean if a field has been set.

### GetWearStatusInDays

`func (o *StoragePhysicalDiskAllOf) GetWearStatusInDays() int64`

GetWearStatusInDays returns the WearStatusInDays field if non-nil, zero value otherwise.

### GetWearStatusInDaysOk

`func (o *StoragePhysicalDiskAllOf) GetWearStatusInDaysOk() (*int64, bool)`

GetWearStatusInDaysOk returns a tuple with the WearStatusInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWearStatusInDays

`func (o *StoragePhysicalDiskAllOf) SetWearStatusInDays(v int64)`

SetWearStatusInDays sets WearStatusInDays field to given value.

### HasWearStatusInDays

`func (o *StoragePhysicalDiskAllOf) HasWearStatusInDays() bool`

HasWearStatusInDays returns a boolean if a field has been set.

### GetWriteErrorCountThreshold

`func (o *StoragePhysicalDiskAllOf) GetWriteErrorCountThreshold() int64`

GetWriteErrorCountThreshold returns the WriteErrorCountThreshold field if non-nil, zero value otherwise.

### GetWriteErrorCountThresholdOk

`func (o *StoragePhysicalDiskAllOf) GetWriteErrorCountThresholdOk() (*int64, bool)`

GetWriteErrorCountThresholdOk returns a tuple with the WriteErrorCountThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteErrorCountThreshold

`func (o *StoragePhysicalDiskAllOf) SetWriteErrorCountThreshold(v int64)`

SetWriteErrorCountThreshold sets WriteErrorCountThreshold field to given value.

### HasWriteErrorCountThreshold

`func (o *StoragePhysicalDiskAllOf) HasWriteErrorCountThreshold() bool`

HasWriteErrorCountThreshold returns a boolean if a field has been set.

### GetWriteIoErrorCount

`func (o *StoragePhysicalDiskAllOf) GetWriteIoErrorCount() int64`

GetWriteIoErrorCount returns the WriteIoErrorCount field if non-nil, zero value otherwise.

### GetWriteIoErrorCountOk

`func (o *StoragePhysicalDiskAllOf) GetWriteIoErrorCountOk() (*int64, bool)`

GetWriteIoErrorCountOk returns a tuple with the WriteIoErrorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteIoErrorCount

`func (o *StoragePhysicalDiskAllOf) SetWriteIoErrorCount(v int64)`

SetWriteIoErrorCount sets WriteIoErrorCount field to given value.

### HasWriteIoErrorCount

`func (o *StoragePhysicalDiskAllOf) HasWriteIoErrorCount() bool`

HasWriteIoErrorCount returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *StoragePhysicalDiskAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *StoragePhysicalDiskAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *StoragePhysicalDiskAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *StoragePhysicalDiskAllOf) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetLocatorLed

`func (o *StoragePhysicalDiskAllOf) GetLocatorLed() EquipmentLocatorLedRelationship`

GetLocatorLed returns the LocatorLed field if non-nil, zero value otherwise.

### GetLocatorLedOk

`func (o *StoragePhysicalDiskAllOf) GetLocatorLedOk() (*EquipmentLocatorLedRelationship, bool)`

GetLocatorLedOk returns a tuple with the LocatorLed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocatorLed

`func (o *StoragePhysicalDiskAllOf) SetLocatorLed(v EquipmentLocatorLedRelationship)`

SetLocatorLed sets LocatorLed field to given value.

### HasLocatorLed

`func (o *StoragePhysicalDiskAllOf) HasLocatorLed() bool`

HasLocatorLed returns a boolean if a field has been set.

### GetPhysicalDiskExtensions

`func (o *StoragePhysicalDiskAllOf) GetPhysicalDiskExtensions() []StoragePhysicalDiskExtensionRelationship`

GetPhysicalDiskExtensions returns the PhysicalDiskExtensions field if non-nil, zero value otherwise.

### GetPhysicalDiskExtensionsOk

`func (o *StoragePhysicalDiskAllOf) GetPhysicalDiskExtensionsOk() (*[]StoragePhysicalDiskExtensionRelationship, bool)`

GetPhysicalDiskExtensionsOk returns a tuple with the PhysicalDiskExtensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalDiskExtensions

`func (o *StoragePhysicalDiskAllOf) SetPhysicalDiskExtensions(v []StoragePhysicalDiskExtensionRelationship)`

SetPhysicalDiskExtensions sets PhysicalDiskExtensions field to given value.

### HasPhysicalDiskExtensions

`func (o *StoragePhysicalDiskAllOf) HasPhysicalDiskExtensions() bool`

HasPhysicalDiskExtensions returns a boolean if a field has been set.

### SetPhysicalDiskExtensionsNil

`func (o *StoragePhysicalDiskAllOf) SetPhysicalDiskExtensionsNil(b bool)`

 SetPhysicalDiskExtensionsNil sets the value for PhysicalDiskExtensions to be an explicit nil

### UnsetPhysicalDiskExtensions
`func (o *StoragePhysicalDiskAllOf) UnsetPhysicalDiskExtensions()`

UnsetPhysicalDiskExtensions ensures that no value is present for PhysicalDiskExtensions, not even an explicit nil
### GetRegisteredDevice

`func (o *StoragePhysicalDiskAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StoragePhysicalDiskAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StoragePhysicalDiskAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StoragePhysicalDiskAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### GetRunningFirmware

`func (o *StoragePhysicalDiskAllOf) GetRunningFirmware() []FirmwareRunningFirmwareRelationship`

GetRunningFirmware returns the RunningFirmware field if non-nil, zero value otherwise.

### GetRunningFirmwareOk

`func (o *StoragePhysicalDiskAllOf) GetRunningFirmwareOk() (*[]FirmwareRunningFirmwareRelationship, bool)`

GetRunningFirmwareOk returns a tuple with the RunningFirmware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningFirmware

`func (o *StoragePhysicalDiskAllOf) SetRunningFirmware(v []FirmwareRunningFirmwareRelationship)`

SetRunningFirmware sets RunningFirmware field to given value.

### HasRunningFirmware

`func (o *StoragePhysicalDiskAllOf) HasRunningFirmware() bool`

HasRunningFirmware returns a boolean if a field has been set.

### SetRunningFirmwareNil

`func (o *StoragePhysicalDiskAllOf) SetRunningFirmwareNil(b bool)`

 SetRunningFirmwareNil sets the value for RunningFirmware to be an explicit nil

### UnsetRunningFirmware
`func (o *StoragePhysicalDiskAllOf) UnsetRunningFirmware()`

UnsetRunningFirmware ensures that no value is present for RunningFirmware, not even an explicit nil
### GetSasPorts

`func (o *StoragePhysicalDiskAllOf) GetSasPorts() []StorageSasPortRelationship`

GetSasPorts returns the SasPorts field if non-nil, zero value otherwise.

### GetSasPortsOk

`func (o *StoragePhysicalDiskAllOf) GetSasPortsOk() (*[]StorageSasPortRelationship, bool)`

GetSasPortsOk returns a tuple with the SasPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSasPorts

`func (o *StoragePhysicalDiskAllOf) SetSasPorts(v []StorageSasPortRelationship)`

SetSasPorts sets SasPorts field to given value.

### HasSasPorts

`func (o *StoragePhysicalDiskAllOf) HasSasPorts() bool`

HasSasPorts returns a boolean if a field has been set.

### SetSasPortsNil

`func (o *StoragePhysicalDiskAllOf) SetSasPortsNil(b bool)`

 SetSasPortsNil sets the value for SasPorts to be an explicit nil

### UnsetSasPorts
`func (o *StoragePhysicalDiskAllOf) UnsetSasPorts()`

UnsetSasPorts ensures that no value is present for SasPorts, not even an explicit nil
### GetStorageController

`func (o *StoragePhysicalDiskAllOf) GetStorageController() StorageControllerRelationship`

GetStorageController returns the StorageController field if non-nil, zero value otherwise.

### GetStorageControllerOk

`func (o *StoragePhysicalDiskAllOf) GetStorageControllerOk() (*StorageControllerRelationship, bool)`

GetStorageControllerOk returns a tuple with the StorageController field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageController

`func (o *StoragePhysicalDiskAllOf) SetStorageController(v StorageControllerRelationship)`

SetStorageController sets StorageController field to given value.

### HasStorageController

`func (o *StoragePhysicalDiskAllOf) HasStorageController() bool`

HasStorageController returns a boolean if a field has been set.

### GetStorageEnclosure

`func (o *StoragePhysicalDiskAllOf) GetStorageEnclosure() StorageEnclosureRelationship`

GetStorageEnclosure returns the StorageEnclosure field if non-nil, zero value otherwise.

### GetStorageEnclosureOk

`func (o *StoragePhysicalDiskAllOf) GetStorageEnclosureOk() (*StorageEnclosureRelationship, bool)`

GetStorageEnclosureOk returns a tuple with the StorageEnclosure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageEnclosure

`func (o *StoragePhysicalDiskAllOf) SetStorageEnclosure(v StorageEnclosureRelationship)`

SetStorageEnclosure sets StorageEnclosure field to given value.

### HasStorageEnclosure

`func (o *StoragePhysicalDiskAllOf) HasStorageEnclosure() bool`

HasStorageEnclosure returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# StoragePhysicalDisk

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
**Description** | Pointer to **string** | This field displays the description of the physical disk. | [optional] [readonly] 
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
**IsPlatformSupported** | Pointer to **bool** | This field indicates whether the physical disk is supported on the server or not. | [optional] [readonly] [default to true]
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
**PartNumber** | Pointer to **string** | This field displays the part number of the physical disk. | [optional] [readonly] 
**PercentLifeLeft** | Pointer to **int64** | Percentage of write cycles remaining in a solid state drive (SSD). | [optional] 
**PercentReservedCapacityConsumed** | Pointer to **int64** | Percentage of reserve capacity consumed. | [optional] 
**PerformancePercent** | Pointer to **int64** | Performance at which the device operating expressed in percentage. | [optional] 
**PhysicalBlockSize** | Pointer to **string** | The block size of the installed physical disk. | [optional] [readonly] 
**Pid** | Pointer to **string** | This field displays the product ID of the physical disk. | [optional] [readonly] 
**PowerCycleCount** | Pointer to **int64** | Number of powercycles the drive has undergone. | [optional] 
**PowerOnHours** | Pointer to **int64** | Number of hours the drive has been powered on. | [optional] 
**PowerOnHoursPercentage** | Pointer to **int64** | Percentage of life used based on five year life span of Cisco supported drives. | [optional] [readonly] 
**PredictedMediaLifeLeftPercent** | Pointer to **int64** | Predicted physical disk life left in percentage. | [optional] 
**PredictiveFailureCount** | Pointer to **int64** | Error count on the physical disk. | [optional] 
**Protocol** | Pointer to **string** | This field identifies the disk protocol used for communication. | [optional] [readonly] 
**RawSize** | Pointer to **string** | The raw size of the physical disk in MB. | [optional] [readonly] 
**ReadErrorCountThreshold** | Pointer to **int64** | The number of read errors that are permitted while accessing the drive/card. | [optional] [readonly] 
**ReadIoErrorCount** | Pointer to **int64** | Number of IO Errors that occured while reading data from the disk. | [optional] 
**Secured** | Pointer to **string** | This field identifies whether the disk is encrypted. | [optional] 
**Size** | Pointer to **string** | The size of the physical disk in MB. | [optional] [readonly] 
**Thermal** | Pointer to **string** | Thermal state of the physical disk. | [optional] [readonly] 
**ThresholdOperatingTemperature** | Pointer to **int64** | Rated threshold operating temperature in Celsius. | [optional] 
**Type** | Pointer to **string** | This field identifies the type of the physical disk. | [optional] [readonly] 
**VariantType** | Pointer to **string** | The variant type of the physical disk. | [optional] [readonly] 
**WearStatusInDays** | Pointer to **int64** | The number of days an SSD has gone through with the write cycles. | [optional] 
**WriteErrorCountThreshold** | Pointer to **int64** | The number of write errors that are permitted while accessing the drive/card. | [optional] [readonly] 
**WriteIoErrorCount** | Pointer to **int64** | Number of IO Errors that occured while writing data to the disk. | [optional] [readonly] 
**InventoryDeviceInfo** | Pointer to [**NullableInventoryDeviceInfoRelationship**](InventoryDeviceInfoRelationship.md) |  | [optional] 
**LocatorLed** | Pointer to [**NullableEquipmentLocatorLedRelationship**](EquipmentLocatorLedRelationship.md) |  | [optional] 
**PhysicalDiskExtensions** | Pointer to [**[]StoragePhysicalDiskExtensionRelationship**](StoragePhysicalDiskExtensionRelationship.md) | An array of relationships to storagePhysicalDiskExtension resources. | [optional] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 
**RunningFirmware** | Pointer to [**[]FirmwareRunningFirmwareRelationship**](FirmwareRunningFirmwareRelationship.md) | An array of relationships to firmwareRunningFirmware resources. | [optional] [readonly] 
**SasPorts** | Pointer to [**[]StorageSasPortRelationship**](StorageSasPortRelationship.md) | An array of relationships to storageSasPort resources. | [optional] [readonly] 
**StorageController** | Pointer to [**NullableStorageControllerRelationship**](StorageControllerRelationship.md) |  | [optional] 
**StorageEnclosure** | Pointer to [**NullableStorageEnclosureRelationship**](StorageEnclosureRelationship.md) |  | [optional] 

## Methods

### NewStoragePhysicalDisk

`func NewStoragePhysicalDisk(classId string, objectType string, ) *StoragePhysicalDisk`

NewStoragePhysicalDisk instantiates a new StoragePhysicalDisk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoragePhysicalDiskWithDefaults

`func NewStoragePhysicalDiskWithDefaults() *StoragePhysicalDisk`

NewStoragePhysicalDiskWithDefaults instantiates a new StoragePhysicalDisk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StoragePhysicalDisk) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StoragePhysicalDisk) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StoragePhysicalDisk) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StoragePhysicalDisk) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StoragePhysicalDisk) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StoragePhysicalDisk) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBackgroundOperations

`func (o *StoragePhysicalDisk) GetBackgroundOperations() string`

GetBackgroundOperations returns the BackgroundOperations field if non-nil, zero value otherwise.

### GetBackgroundOperationsOk

`func (o *StoragePhysicalDisk) GetBackgroundOperationsOk() (*string, bool)`

GetBackgroundOperationsOk returns a tuple with the BackgroundOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundOperations

`func (o *StoragePhysicalDisk) SetBackgroundOperations(v string)`

SetBackgroundOperations sets BackgroundOperations field to given value.

### HasBackgroundOperations

`func (o *StoragePhysicalDisk) HasBackgroundOperations() bool`

HasBackgroundOperations returns a boolean if a field has been set.

### GetBlockSize

`func (o *StoragePhysicalDisk) GetBlockSize() string`

GetBlockSize returns the BlockSize field if non-nil, zero value otherwise.

### GetBlockSizeOk

`func (o *StoragePhysicalDisk) GetBlockSizeOk() (*string, bool)`

GetBlockSizeOk returns a tuple with the BlockSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockSize

`func (o *StoragePhysicalDisk) SetBlockSize(v string)`

SetBlockSize sets BlockSize field to given value.

### HasBlockSize

`func (o *StoragePhysicalDisk) HasBlockSize() bool`

HasBlockSize returns a boolean if a field has been set.

### GetBootable

`func (o *StoragePhysicalDisk) GetBootable() string`

GetBootable returns the Bootable field if non-nil, zero value otherwise.

### GetBootableOk

`func (o *StoragePhysicalDisk) GetBootableOk() (*string, bool)`

GetBootableOk returns a tuple with the Bootable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootable

`func (o *StoragePhysicalDisk) SetBootable(v string)`

SetBootable sets Bootable field to given value.

### HasBootable

`func (o *StoragePhysicalDisk) HasBootable() bool`

HasBootable returns a boolean if a field has been set.

### GetConfigurationCheckpoint

`func (o *StoragePhysicalDisk) GetConfigurationCheckpoint() string`

GetConfigurationCheckpoint returns the ConfigurationCheckpoint field if non-nil, zero value otherwise.

### GetConfigurationCheckpointOk

`func (o *StoragePhysicalDisk) GetConfigurationCheckpointOk() (*string, bool)`

GetConfigurationCheckpointOk returns a tuple with the ConfigurationCheckpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationCheckpoint

`func (o *StoragePhysicalDisk) SetConfigurationCheckpoint(v string)`

SetConfigurationCheckpoint sets ConfigurationCheckpoint field to given value.

### HasConfigurationCheckpoint

`func (o *StoragePhysicalDisk) HasConfigurationCheckpoint() bool`

HasConfigurationCheckpoint returns a boolean if a field has been set.

### GetConfigurationState

`func (o *StoragePhysicalDisk) GetConfigurationState() string`

GetConfigurationState returns the ConfigurationState field if non-nil, zero value otherwise.

### GetConfigurationStateOk

`func (o *StoragePhysicalDisk) GetConfigurationStateOk() (*string, bool)`

GetConfigurationStateOk returns a tuple with the ConfigurationState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationState

`func (o *StoragePhysicalDisk) SetConfigurationState(v string)`

SetConfigurationState sets ConfigurationState field to given value.

### HasConfigurationState

`func (o *StoragePhysicalDisk) HasConfigurationState() bool`

HasConfigurationState returns a boolean if a field has been set.

### GetDescription

`func (o *StoragePhysicalDisk) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StoragePhysicalDisk) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StoragePhysicalDisk) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *StoragePhysicalDisk) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisabledForRemoval

`func (o *StoragePhysicalDisk) GetDisabledForRemoval() bool`

GetDisabledForRemoval returns the DisabledForRemoval field if non-nil, zero value otherwise.

### GetDisabledForRemovalOk

`func (o *StoragePhysicalDisk) GetDisabledForRemovalOk() (*bool, bool)`

GetDisabledForRemovalOk returns a tuple with the DisabledForRemoval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledForRemoval

`func (o *StoragePhysicalDisk) SetDisabledForRemoval(v bool)`

SetDisabledForRemoval sets DisabledForRemoval field to given value.

### HasDisabledForRemoval

`func (o *StoragePhysicalDisk) HasDisabledForRemoval() bool`

HasDisabledForRemoval returns a boolean if a field has been set.

### GetDiscoveredPath

`func (o *StoragePhysicalDisk) GetDiscoveredPath() string`

GetDiscoveredPath returns the DiscoveredPath field if non-nil, zero value otherwise.

### GetDiscoveredPathOk

`func (o *StoragePhysicalDisk) GetDiscoveredPathOk() (*string, bool)`

GetDiscoveredPathOk returns a tuple with the DiscoveredPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveredPath

`func (o *StoragePhysicalDisk) SetDiscoveredPath(v string)`

SetDiscoveredPath sets DiscoveredPath field to given value.

### HasDiscoveredPath

`func (o *StoragePhysicalDisk) HasDiscoveredPath() bool`

HasDiscoveredPath returns a boolean if a field has been set.

### GetDiskId

`func (o *StoragePhysicalDisk) GetDiskId() string`

GetDiskId returns the DiskId field if non-nil, zero value otherwise.

### GetDiskIdOk

`func (o *StoragePhysicalDisk) GetDiskIdOk() (*string, bool)`

GetDiskIdOk returns a tuple with the DiskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskId

`func (o *StoragePhysicalDisk) SetDiskId(v string)`

SetDiskId sets DiskId field to given value.

### HasDiskId

`func (o *StoragePhysicalDisk) HasDiskId() bool`

HasDiskId returns a boolean if a field has been set.

### GetDiskState

`func (o *StoragePhysicalDisk) GetDiskState() string`

GetDiskState returns the DiskState field if non-nil, zero value otherwise.

### GetDiskStateOk

`func (o *StoragePhysicalDisk) GetDiskStateOk() (*string, bool)`

GetDiskStateOk returns a tuple with the DiskState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskState

`func (o *StoragePhysicalDisk) SetDiskState(v string)`

SetDiskState sets DiskState field to given value.

### HasDiskState

`func (o *StoragePhysicalDisk) HasDiskState() bool`

HasDiskState returns a boolean if a field has been set.

### GetDriveFirmware

`func (o *StoragePhysicalDisk) GetDriveFirmware() string`

GetDriveFirmware returns the DriveFirmware field if non-nil, zero value otherwise.

### GetDriveFirmwareOk

`func (o *StoragePhysicalDisk) GetDriveFirmwareOk() (*string, bool)`

GetDriveFirmwareOk returns a tuple with the DriveFirmware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriveFirmware

`func (o *StoragePhysicalDisk) SetDriveFirmware(v string)`

SetDriveFirmware sets DriveFirmware field to given value.

### HasDriveFirmware

`func (o *StoragePhysicalDisk) HasDriveFirmware() bool`

HasDriveFirmware returns a boolean if a field has been set.

### GetDriveState

`func (o *StoragePhysicalDisk) GetDriveState() string`

GetDriveState returns the DriveState field if non-nil, zero value otherwise.

### GetDriveStateOk

`func (o *StoragePhysicalDisk) GetDriveStateOk() (*string, bool)`

GetDriveStateOk returns a tuple with the DriveState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriveState

`func (o *StoragePhysicalDisk) SetDriveState(v string)`

SetDriveState sets DriveState field to given value.

### HasDriveState

`func (o *StoragePhysicalDisk) HasDriveState() bool`

HasDriveState returns a boolean if a field has been set.

### GetEncryptionStatus

`func (o *StoragePhysicalDisk) GetEncryptionStatus() string`

GetEncryptionStatus returns the EncryptionStatus field if non-nil, zero value otherwise.

### GetEncryptionStatusOk

`func (o *StoragePhysicalDisk) GetEncryptionStatusOk() (*string, bool)`

GetEncryptionStatusOk returns a tuple with the EncryptionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionStatus

`func (o *StoragePhysicalDisk) SetEncryptionStatus(v string)`

SetEncryptionStatus sets EncryptionStatus field to given value.

### HasEncryptionStatus

`func (o *StoragePhysicalDisk) HasEncryptionStatus() bool`

HasEncryptionStatus returns a boolean if a field has been set.

### GetFailurePredicted

`func (o *StoragePhysicalDisk) GetFailurePredicted() bool`

GetFailurePredicted returns the FailurePredicted field if non-nil, zero value otherwise.

### GetFailurePredictedOk

`func (o *StoragePhysicalDisk) GetFailurePredictedOk() (*bool, bool)`

GetFailurePredictedOk returns a tuple with the FailurePredicted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailurePredicted

`func (o *StoragePhysicalDisk) SetFailurePredicted(v bool)`

SetFailurePredicted sets FailurePredicted field to given value.

### HasFailurePredicted

`func (o *StoragePhysicalDisk) HasFailurePredicted() bool`

HasFailurePredicted returns a boolean if a field has been set.

### GetFdeCapable

`func (o *StoragePhysicalDisk) GetFdeCapable() string`

GetFdeCapable returns the FdeCapable field if non-nil, zero value otherwise.

### GetFdeCapableOk

`func (o *StoragePhysicalDisk) GetFdeCapableOk() (*string, bool)`

GetFdeCapableOk returns a tuple with the FdeCapable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFdeCapable

`func (o *StoragePhysicalDisk) SetFdeCapable(v string)`

SetFdeCapable sets FdeCapable field to given value.

### HasFdeCapable

`func (o *StoragePhysicalDisk) HasFdeCapable() bool`

HasFdeCapable returns a boolean if a field has been set.

### GetHotSpareType

`func (o *StoragePhysicalDisk) GetHotSpareType() string`

GetHotSpareType returns the HotSpareType field if non-nil, zero value otherwise.

### GetHotSpareTypeOk

`func (o *StoragePhysicalDisk) GetHotSpareTypeOk() (*string, bool)`

GetHotSpareTypeOk returns a tuple with the HotSpareType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHotSpareType

`func (o *StoragePhysicalDisk) SetHotSpareType(v string)`

SetHotSpareType sets HotSpareType field to given value.

### HasHotSpareType

`func (o *StoragePhysicalDisk) HasHotSpareType() bool`

HasHotSpareType returns a boolean if a field has been set.

### GetIndicatorLed

`func (o *StoragePhysicalDisk) GetIndicatorLed() string`

GetIndicatorLed returns the IndicatorLed field if non-nil, zero value otherwise.

### GetIndicatorLedOk

`func (o *StoragePhysicalDisk) GetIndicatorLedOk() (*string, bool)`

GetIndicatorLedOk returns a tuple with the IndicatorLed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndicatorLed

`func (o *StoragePhysicalDisk) SetIndicatorLed(v string)`

SetIndicatorLed sets IndicatorLed field to given value.

### HasIndicatorLed

`func (o *StoragePhysicalDisk) HasIndicatorLed() bool`

HasIndicatorLed returns a boolean if a field has been set.

### GetIsPlatformSupported

`func (o *StoragePhysicalDisk) GetIsPlatformSupported() bool`

GetIsPlatformSupported returns the IsPlatformSupported field if non-nil, zero value otherwise.

### GetIsPlatformSupportedOk

`func (o *StoragePhysicalDisk) GetIsPlatformSupportedOk() (*bool, bool)`

GetIsPlatformSupportedOk returns a tuple with the IsPlatformSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPlatformSupported

`func (o *StoragePhysicalDisk) SetIsPlatformSupported(v bool)`

SetIsPlatformSupported sets IsPlatformSupported field to given value.

### HasIsPlatformSupported

`func (o *StoragePhysicalDisk) HasIsPlatformSupported() bool`

HasIsPlatformSupported returns a boolean if a field has been set.

### GetLinkSpeed

`func (o *StoragePhysicalDisk) GetLinkSpeed() string`

GetLinkSpeed returns the LinkSpeed field if non-nil, zero value otherwise.

### GetLinkSpeedOk

`func (o *StoragePhysicalDisk) GetLinkSpeedOk() (*string, bool)`

GetLinkSpeedOk returns a tuple with the LinkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkSpeed

`func (o *StoragePhysicalDisk) SetLinkSpeed(v string)`

SetLinkSpeed sets LinkSpeed field to given value.

### HasLinkSpeed

`func (o *StoragePhysicalDisk) HasLinkSpeed() bool`

HasLinkSpeed returns a boolean if a field has been set.

### GetLinkState

`func (o *StoragePhysicalDisk) GetLinkState() string`

GetLinkState returns the LinkState field if non-nil, zero value otherwise.

### GetLinkStateOk

`func (o *StoragePhysicalDisk) GetLinkStateOk() (*string, bool)`

GetLinkStateOk returns a tuple with the LinkState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkState

`func (o *StoragePhysicalDisk) SetLinkState(v string)`

SetLinkState sets LinkState field to given value.

### HasLinkState

`func (o *StoragePhysicalDisk) HasLinkState() bool`

HasLinkState returns a boolean if a field has been set.

### GetMaximumOperatingTemperature

`func (o *StoragePhysicalDisk) GetMaximumOperatingTemperature() int64`

GetMaximumOperatingTemperature returns the MaximumOperatingTemperature field if non-nil, zero value otherwise.

### GetMaximumOperatingTemperatureOk

`func (o *StoragePhysicalDisk) GetMaximumOperatingTemperatureOk() (*int64, bool)`

GetMaximumOperatingTemperatureOk returns a tuple with the MaximumOperatingTemperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumOperatingTemperature

`func (o *StoragePhysicalDisk) SetMaximumOperatingTemperature(v int64)`

SetMaximumOperatingTemperature sets MaximumOperatingTemperature field to given value.

### HasMaximumOperatingTemperature

`func (o *StoragePhysicalDisk) HasMaximumOperatingTemperature() bool`

HasMaximumOperatingTemperature returns a boolean if a field has been set.

### GetMediaErrorCount

`func (o *StoragePhysicalDisk) GetMediaErrorCount() int64`

GetMediaErrorCount returns the MediaErrorCount field if non-nil, zero value otherwise.

### GetMediaErrorCountOk

`func (o *StoragePhysicalDisk) GetMediaErrorCountOk() (*int64, bool)`

GetMediaErrorCountOk returns a tuple with the MediaErrorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaErrorCount

`func (o *StoragePhysicalDisk) SetMediaErrorCount(v int64)`

SetMediaErrorCount sets MediaErrorCount field to given value.

### HasMediaErrorCount

`func (o *StoragePhysicalDisk) HasMediaErrorCount() bool`

HasMediaErrorCount returns a boolean if a field has been set.

### GetName

`func (o *StoragePhysicalDisk) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StoragePhysicalDisk) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StoragePhysicalDisk) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StoragePhysicalDisk) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNonCoercedSizeBytes

`func (o *StoragePhysicalDisk) GetNonCoercedSizeBytes() int64`

GetNonCoercedSizeBytes returns the NonCoercedSizeBytes field if non-nil, zero value otherwise.

### GetNonCoercedSizeBytesOk

`func (o *StoragePhysicalDisk) GetNonCoercedSizeBytesOk() (*int64, bool)`

GetNonCoercedSizeBytesOk returns a tuple with the NonCoercedSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonCoercedSizeBytes

`func (o *StoragePhysicalDisk) SetNonCoercedSizeBytes(v int64)`

SetNonCoercedSizeBytes sets NonCoercedSizeBytes field to given value.

### HasNonCoercedSizeBytes

`func (o *StoragePhysicalDisk) HasNonCoercedSizeBytes() bool`

HasNonCoercedSizeBytes returns a boolean if a field has been set.

### GetNumBlocks

`func (o *StoragePhysicalDisk) GetNumBlocks() string`

GetNumBlocks returns the NumBlocks field if non-nil, zero value otherwise.

### GetNumBlocksOk

`func (o *StoragePhysicalDisk) GetNumBlocksOk() (*string, bool)`

GetNumBlocksOk returns a tuple with the NumBlocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumBlocks

`func (o *StoragePhysicalDisk) SetNumBlocks(v string)`

SetNumBlocks sets NumBlocks field to given value.

### HasNumBlocks

`func (o *StoragePhysicalDisk) HasNumBlocks() bool`

HasNumBlocks returns a boolean if a field has been set.

### GetOperPowerState

`func (o *StoragePhysicalDisk) GetOperPowerState() string`

GetOperPowerState returns the OperPowerState field if non-nil, zero value otherwise.

### GetOperPowerStateOk

`func (o *StoragePhysicalDisk) GetOperPowerStateOk() (*string, bool)`

GetOperPowerStateOk returns a tuple with the OperPowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperPowerState

`func (o *StoragePhysicalDisk) SetOperPowerState(v string)`

SetOperPowerState sets OperPowerState field to given value.

### HasOperPowerState

`func (o *StoragePhysicalDisk) HasOperPowerState() bool`

HasOperPowerState returns a boolean if a field has been set.

### GetOperQualifierReason

`func (o *StoragePhysicalDisk) GetOperQualifierReason() string`

GetOperQualifierReason returns the OperQualifierReason field if non-nil, zero value otherwise.

### GetOperQualifierReasonOk

`func (o *StoragePhysicalDisk) GetOperQualifierReasonOk() (*string, bool)`

GetOperQualifierReasonOk returns a tuple with the OperQualifierReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperQualifierReason

`func (o *StoragePhysicalDisk) SetOperQualifierReason(v string)`

SetOperQualifierReason sets OperQualifierReason field to given value.

### HasOperQualifierReason

`func (o *StoragePhysicalDisk) HasOperQualifierReason() bool`

HasOperQualifierReason returns a boolean if a field has been set.

### GetOperability

`func (o *StoragePhysicalDisk) GetOperability() string`

GetOperability returns the Operability field if non-nil, zero value otherwise.

### GetOperabilityOk

`func (o *StoragePhysicalDisk) GetOperabilityOk() (*string, bool)`

GetOperabilityOk returns a tuple with the Operability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperability

`func (o *StoragePhysicalDisk) SetOperability(v string)`

SetOperability sets Operability field to given value.

### HasOperability

`func (o *StoragePhysicalDisk) HasOperability() bool`

HasOperability returns a boolean if a field has been set.

### GetOperatingTemperature

`func (o *StoragePhysicalDisk) GetOperatingTemperature() int64`

GetOperatingTemperature returns the OperatingTemperature field if non-nil, zero value otherwise.

### GetOperatingTemperatureOk

`func (o *StoragePhysicalDisk) GetOperatingTemperatureOk() (*int64, bool)`

GetOperatingTemperatureOk returns a tuple with the OperatingTemperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingTemperature

`func (o *StoragePhysicalDisk) SetOperatingTemperature(v int64)`

SetOperatingTemperature sets OperatingTemperature field to given value.

### HasOperatingTemperature

`func (o *StoragePhysicalDisk) HasOperatingTemperature() bool`

HasOperatingTemperature returns a boolean if a field has been set.

### GetPartNumber

`func (o *StoragePhysicalDisk) GetPartNumber() string`

GetPartNumber returns the PartNumber field if non-nil, zero value otherwise.

### GetPartNumberOk

`func (o *StoragePhysicalDisk) GetPartNumberOk() (*string, bool)`

GetPartNumberOk returns a tuple with the PartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartNumber

`func (o *StoragePhysicalDisk) SetPartNumber(v string)`

SetPartNumber sets PartNumber field to given value.

### HasPartNumber

`func (o *StoragePhysicalDisk) HasPartNumber() bool`

HasPartNumber returns a boolean if a field has been set.

### GetPercentLifeLeft

`func (o *StoragePhysicalDisk) GetPercentLifeLeft() int64`

GetPercentLifeLeft returns the PercentLifeLeft field if non-nil, zero value otherwise.

### GetPercentLifeLeftOk

`func (o *StoragePhysicalDisk) GetPercentLifeLeftOk() (*int64, bool)`

GetPercentLifeLeftOk returns a tuple with the PercentLifeLeft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentLifeLeft

`func (o *StoragePhysicalDisk) SetPercentLifeLeft(v int64)`

SetPercentLifeLeft sets PercentLifeLeft field to given value.

### HasPercentLifeLeft

`func (o *StoragePhysicalDisk) HasPercentLifeLeft() bool`

HasPercentLifeLeft returns a boolean if a field has been set.

### GetPercentReservedCapacityConsumed

`func (o *StoragePhysicalDisk) GetPercentReservedCapacityConsumed() int64`

GetPercentReservedCapacityConsumed returns the PercentReservedCapacityConsumed field if non-nil, zero value otherwise.

### GetPercentReservedCapacityConsumedOk

`func (o *StoragePhysicalDisk) GetPercentReservedCapacityConsumedOk() (*int64, bool)`

GetPercentReservedCapacityConsumedOk returns a tuple with the PercentReservedCapacityConsumed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentReservedCapacityConsumed

`func (o *StoragePhysicalDisk) SetPercentReservedCapacityConsumed(v int64)`

SetPercentReservedCapacityConsumed sets PercentReservedCapacityConsumed field to given value.

### HasPercentReservedCapacityConsumed

`func (o *StoragePhysicalDisk) HasPercentReservedCapacityConsumed() bool`

HasPercentReservedCapacityConsumed returns a boolean if a field has been set.

### GetPerformancePercent

`func (o *StoragePhysicalDisk) GetPerformancePercent() int64`

GetPerformancePercent returns the PerformancePercent field if non-nil, zero value otherwise.

### GetPerformancePercentOk

`func (o *StoragePhysicalDisk) GetPerformancePercentOk() (*int64, bool)`

GetPerformancePercentOk returns a tuple with the PerformancePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformancePercent

`func (o *StoragePhysicalDisk) SetPerformancePercent(v int64)`

SetPerformancePercent sets PerformancePercent field to given value.

### HasPerformancePercent

`func (o *StoragePhysicalDisk) HasPerformancePercent() bool`

HasPerformancePercent returns a boolean if a field has been set.

### GetPhysicalBlockSize

`func (o *StoragePhysicalDisk) GetPhysicalBlockSize() string`

GetPhysicalBlockSize returns the PhysicalBlockSize field if non-nil, zero value otherwise.

### GetPhysicalBlockSizeOk

`func (o *StoragePhysicalDisk) GetPhysicalBlockSizeOk() (*string, bool)`

GetPhysicalBlockSizeOk returns a tuple with the PhysicalBlockSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalBlockSize

`func (o *StoragePhysicalDisk) SetPhysicalBlockSize(v string)`

SetPhysicalBlockSize sets PhysicalBlockSize field to given value.

### HasPhysicalBlockSize

`func (o *StoragePhysicalDisk) HasPhysicalBlockSize() bool`

HasPhysicalBlockSize returns a boolean if a field has been set.

### GetPid

`func (o *StoragePhysicalDisk) GetPid() string`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *StoragePhysicalDisk) GetPidOk() (*string, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *StoragePhysicalDisk) SetPid(v string)`

SetPid sets Pid field to given value.

### HasPid

`func (o *StoragePhysicalDisk) HasPid() bool`

HasPid returns a boolean if a field has been set.

### GetPowerCycleCount

`func (o *StoragePhysicalDisk) GetPowerCycleCount() int64`

GetPowerCycleCount returns the PowerCycleCount field if non-nil, zero value otherwise.

### GetPowerCycleCountOk

`func (o *StoragePhysicalDisk) GetPowerCycleCountOk() (*int64, bool)`

GetPowerCycleCountOk returns a tuple with the PowerCycleCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerCycleCount

`func (o *StoragePhysicalDisk) SetPowerCycleCount(v int64)`

SetPowerCycleCount sets PowerCycleCount field to given value.

### HasPowerCycleCount

`func (o *StoragePhysicalDisk) HasPowerCycleCount() bool`

HasPowerCycleCount returns a boolean if a field has been set.

### GetPowerOnHours

`func (o *StoragePhysicalDisk) GetPowerOnHours() int64`

GetPowerOnHours returns the PowerOnHours field if non-nil, zero value otherwise.

### GetPowerOnHoursOk

`func (o *StoragePhysicalDisk) GetPowerOnHoursOk() (*int64, bool)`

GetPowerOnHoursOk returns a tuple with the PowerOnHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerOnHours

`func (o *StoragePhysicalDisk) SetPowerOnHours(v int64)`

SetPowerOnHours sets PowerOnHours field to given value.

### HasPowerOnHours

`func (o *StoragePhysicalDisk) HasPowerOnHours() bool`

HasPowerOnHours returns a boolean if a field has been set.

### GetPowerOnHoursPercentage

`func (o *StoragePhysicalDisk) GetPowerOnHoursPercentage() int64`

GetPowerOnHoursPercentage returns the PowerOnHoursPercentage field if non-nil, zero value otherwise.

### GetPowerOnHoursPercentageOk

`func (o *StoragePhysicalDisk) GetPowerOnHoursPercentageOk() (*int64, bool)`

GetPowerOnHoursPercentageOk returns a tuple with the PowerOnHoursPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerOnHoursPercentage

`func (o *StoragePhysicalDisk) SetPowerOnHoursPercentage(v int64)`

SetPowerOnHoursPercentage sets PowerOnHoursPercentage field to given value.

### HasPowerOnHoursPercentage

`func (o *StoragePhysicalDisk) HasPowerOnHoursPercentage() bool`

HasPowerOnHoursPercentage returns a boolean if a field has been set.

### GetPredictedMediaLifeLeftPercent

`func (o *StoragePhysicalDisk) GetPredictedMediaLifeLeftPercent() int64`

GetPredictedMediaLifeLeftPercent returns the PredictedMediaLifeLeftPercent field if non-nil, zero value otherwise.

### GetPredictedMediaLifeLeftPercentOk

`func (o *StoragePhysicalDisk) GetPredictedMediaLifeLeftPercentOk() (*int64, bool)`

GetPredictedMediaLifeLeftPercentOk returns a tuple with the PredictedMediaLifeLeftPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredictedMediaLifeLeftPercent

`func (o *StoragePhysicalDisk) SetPredictedMediaLifeLeftPercent(v int64)`

SetPredictedMediaLifeLeftPercent sets PredictedMediaLifeLeftPercent field to given value.

### HasPredictedMediaLifeLeftPercent

`func (o *StoragePhysicalDisk) HasPredictedMediaLifeLeftPercent() bool`

HasPredictedMediaLifeLeftPercent returns a boolean if a field has been set.

### GetPredictiveFailureCount

`func (o *StoragePhysicalDisk) GetPredictiveFailureCount() int64`

GetPredictiveFailureCount returns the PredictiveFailureCount field if non-nil, zero value otherwise.

### GetPredictiveFailureCountOk

`func (o *StoragePhysicalDisk) GetPredictiveFailureCountOk() (*int64, bool)`

GetPredictiveFailureCountOk returns a tuple with the PredictiveFailureCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredictiveFailureCount

`func (o *StoragePhysicalDisk) SetPredictiveFailureCount(v int64)`

SetPredictiveFailureCount sets PredictiveFailureCount field to given value.

### HasPredictiveFailureCount

`func (o *StoragePhysicalDisk) HasPredictiveFailureCount() bool`

HasPredictiveFailureCount returns a boolean if a field has been set.

### GetProtocol

`func (o *StoragePhysicalDisk) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *StoragePhysicalDisk) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *StoragePhysicalDisk) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *StoragePhysicalDisk) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetRawSize

`func (o *StoragePhysicalDisk) GetRawSize() string`

GetRawSize returns the RawSize field if non-nil, zero value otherwise.

### GetRawSizeOk

`func (o *StoragePhysicalDisk) GetRawSizeOk() (*string, bool)`

GetRawSizeOk returns a tuple with the RawSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawSize

`func (o *StoragePhysicalDisk) SetRawSize(v string)`

SetRawSize sets RawSize field to given value.

### HasRawSize

`func (o *StoragePhysicalDisk) HasRawSize() bool`

HasRawSize returns a boolean if a field has been set.

### GetReadErrorCountThreshold

`func (o *StoragePhysicalDisk) GetReadErrorCountThreshold() int64`

GetReadErrorCountThreshold returns the ReadErrorCountThreshold field if non-nil, zero value otherwise.

### GetReadErrorCountThresholdOk

`func (o *StoragePhysicalDisk) GetReadErrorCountThresholdOk() (*int64, bool)`

GetReadErrorCountThresholdOk returns a tuple with the ReadErrorCountThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadErrorCountThreshold

`func (o *StoragePhysicalDisk) SetReadErrorCountThreshold(v int64)`

SetReadErrorCountThreshold sets ReadErrorCountThreshold field to given value.

### HasReadErrorCountThreshold

`func (o *StoragePhysicalDisk) HasReadErrorCountThreshold() bool`

HasReadErrorCountThreshold returns a boolean if a field has been set.

### GetReadIoErrorCount

`func (o *StoragePhysicalDisk) GetReadIoErrorCount() int64`

GetReadIoErrorCount returns the ReadIoErrorCount field if non-nil, zero value otherwise.

### GetReadIoErrorCountOk

`func (o *StoragePhysicalDisk) GetReadIoErrorCountOk() (*int64, bool)`

GetReadIoErrorCountOk returns a tuple with the ReadIoErrorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadIoErrorCount

`func (o *StoragePhysicalDisk) SetReadIoErrorCount(v int64)`

SetReadIoErrorCount sets ReadIoErrorCount field to given value.

### HasReadIoErrorCount

`func (o *StoragePhysicalDisk) HasReadIoErrorCount() bool`

HasReadIoErrorCount returns a boolean if a field has been set.

### GetSecured

`func (o *StoragePhysicalDisk) GetSecured() string`

GetSecured returns the Secured field if non-nil, zero value otherwise.

### GetSecuredOk

`func (o *StoragePhysicalDisk) GetSecuredOk() (*string, bool)`

GetSecuredOk returns a tuple with the Secured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecured

`func (o *StoragePhysicalDisk) SetSecured(v string)`

SetSecured sets Secured field to given value.

### HasSecured

`func (o *StoragePhysicalDisk) HasSecured() bool`

HasSecured returns a boolean if a field has been set.

### GetSize

`func (o *StoragePhysicalDisk) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *StoragePhysicalDisk) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *StoragePhysicalDisk) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *StoragePhysicalDisk) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetThermal

`func (o *StoragePhysicalDisk) GetThermal() string`

GetThermal returns the Thermal field if non-nil, zero value otherwise.

### GetThermalOk

`func (o *StoragePhysicalDisk) GetThermalOk() (*string, bool)`

GetThermalOk returns a tuple with the Thermal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThermal

`func (o *StoragePhysicalDisk) SetThermal(v string)`

SetThermal sets Thermal field to given value.

### HasThermal

`func (o *StoragePhysicalDisk) HasThermal() bool`

HasThermal returns a boolean if a field has been set.

### GetThresholdOperatingTemperature

`func (o *StoragePhysicalDisk) GetThresholdOperatingTemperature() int64`

GetThresholdOperatingTemperature returns the ThresholdOperatingTemperature field if non-nil, zero value otherwise.

### GetThresholdOperatingTemperatureOk

`func (o *StoragePhysicalDisk) GetThresholdOperatingTemperatureOk() (*int64, bool)`

GetThresholdOperatingTemperatureOk returns a tuple with the ThresholdOperatingTemperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdOperatingTemperature

`func (o *StoragePhysicalDisk) SetThresholdOperatingTemperature(v int64)`

SetThresholdOperatingTemperature sets ThresholdOperatingTemperature field to given value.

### HasThresholdOperatingTemperature

`func (o *StoragePhysicalDisk) HasThresholdOperatingTemperature() bool`

HasThresholdOperatingTemperature returns a boolean if a field has been set.

### GetType

`func (o *StoragePhysicalDisk) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StoragePhysicalDisk) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StoragePhysicalDisk) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *StoragePhysicalDisk) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVariantType

`func (o *StoragePhysicalDisk) GetVariantType() string`

GetVariantType returns the VariantType field if non-nil, zero value otherwise.

### GetVariantTypeOk

`func (o *StoragePhysicalDisk) GetVariantTypeOk() (*string, bool)`

GetVariantTypeOk returns a tuple with the VariantType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariantType

`func (o *StoragePhysicalDisk) SetVariantType(v string)`

SetVariantType sets VariantType field to given value.

### HasVariantType

`func (o *StoragePhysicalDisk) HasVariantType() bool`

HasVariantType returns a boolean if a field has been set.

### GetWearStatusInDays

`func (o *StoragePhysicalDisk) GetWearStatusInDays() int64`

GetWearStatusInDays returns the WearStatusInDays field if non-nil, zero value otherwise.

### GetWearStatusInDaysOk

`func (o *StoragePhysicalDisk) GetWearStatusInDaysOk() (*int64, bool)`

GetWearStatusInDaysOk returns a tuple with the WearStatusInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWearStatusInDays

`func (o *StoragePhysicalDisk) SetWearStatusInDays(v int64)`

SetWearStatusInDays sets WearStatusInDays field to given value.

### HasWearStatusInDays

`func (o *StoragePhysicalDisk) HasWearStatusInDays() bool`

HasWearStatusInDays returns a boolean if a field has been set.

### GetWriteErrorCountThreshold

`func (o *StoragePhysicalDisk) GetWriteErrorCountThreshold() int64`

GetWriteErrorCountThreshold returns the WriteErrorCountThreshold field if non-nil, zero value otherwise.

### GetWriteErrorCountThresholdOk

`func (o *StoragePhysicalDisk) GetWriteErrorCountThresholdOk() (*int64, bool)`

GetWriteErrorCountThresholdOk returns a tuple with the WriteErrorCountThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteErrorCountThreshold

`func (o *StoragePhysicalDisk) SetWriteErrorCountThreshold(v int64)`

SetWriteErrorCountThreshold sets WriteErrorCountThreshold field to given value.

### HasWriteErrorCountThreshold

`func (o *StoragePhysicalDisk) HasWriteErrorCountThreshold() bool`

HasWriteErrorCountThreshold returns a boolean if a field has been set.

### GetWriteIoErrorCount

`func (o *StoragePhysicalDisk) GetWriteIoErrorCount() int64`

GetWriteIoErrorCount returns the WriteIoErrorCount field if non-nil, zero value otherwise.

### GetWriteIoErrorCountOk

`func (o *StoragePhysicalDisk) GetWriteIoErrorCountOk() (*int64, bool)`

GetWriteIoErrorCountOk returns a tuple with the WriteIoErrorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteIoErrorCount

`func (o *StoragePhysicalDisk) SetWriteIoErrorCount(v int64)`

SetWriteIoErrorCount sets WriteIoErrorCount field to given value.

### HasWriteIoErrorCount

`func (o *StoragePhysicalDisk) HasWriteIoErrorCount() bool`

HasWriteIoErrorCount returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *StoragePhysicalDisk) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *StoragePhysicalDisk) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *StoragePhysicalDisk) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *StoragePhysicalDisk) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### SetInventoryDeviceInfoNil

`func (o *StoragePhysicalDisk) SetInventoryDeviceInfoNil(b bool)`

 SetInventoryDeviceInfoNil sets the value for InventoryDeviceInfo to be an explicit nil

### UnsetInventoryDeviceInfo
`func (o *StoragePhysicalDisk) UnsetInventoryDeviceInfo()`

UnsetInventoryDeviceInfo ensures that no value is present for InventoryDeviceInfo, not even an explicit nil
### GetLocatorLed

`func (o *StoragePhysicalDisk) GetLocatorLed() EquipmentLocatorLedRelationship`

GetLocatorLed returns the LocatorLed field if non-nil, zero value otherwise.

### GetLocatorLedOk

`func (o *StoragePhysicalDisk) GetLocatorLedOk() (*EquipmentLocatorLedRelationship, bool)`

GetLocatorLedOk returns a tuple with the LocatorLed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocatorLed

`func (o *StoragePhysicalDisk) SetLocatorLed(v EquipmentLocatorLedRelationship)`

SetLocatorLed sets LocatorLed field to given value.

### HasLocatorLed

`func (o *StoragePhysicalDisk) HasLocatorLed() bool`

HasLocatorLed returns a boolean if a field has been set.

### SetLocatorLedNil

`func (o *StoragePhysicalDisk) SetLocatorLedNil(b bool)`

 SetLocatorLedNil sets the value for LocatorLed to be an explicit nil

### UnsetLocatorLed
`func (o *StoragePhysicalDisk) UnsetLocatorLed()`

UnsetLocatorLed ensures that no value is present for LocatorLed, not even an explicit nil
### GetPhysicalDiskExtensions

`func (o *StoragePhysicalDisk) GetPhysicalDiskExtensions() []StoragePhysicalDiskExtensionRelationship`

GetPhysicalDiskExtensions returns the PhysicalDiskExtensions field if non-nil, zero value otherwise.

### GetPhysicalDiskExtensionsOk

`func (o *StoragePhysicalDisk) GetPhysicalDiskExtensionsOk() (*[]StoragePhysicalDiskExtensionRelationship, bool)`

GetPhysicalDiskExtensionsOk returns a tuple with the PhysicalDiskExtensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalDiskExtensions

`func (o *StoragePhysicalDisk) SetPhysicalDiskExtensions(v []StoragePhysicalDiskExtensionRelationship)`

SetPhysicalDiskExtensions sets PhysicalDiskExtensions field to given value.

### HasPhysicalDiskExtensions

`func (o *StoragePhysicalDisk) HasPhysicalDiskExtensions() bool`

HasPhysicalDiskExtensions returns a boolean if a field has been set.

### SetPhysicalDiskExtensionsNil

`func (o *StoragePhysicalDisk) SetPhysicalDiskExtensionsNil(b bool)`

 SetPhysicalDiskExtensionsNil sets the value for PhysicalDiskExtensions to be an explicit nil

### UnsetPhysicalDiskExtensions
`func (o *StoragePhysicalDisk) UnsetPhysicalDiskExtensions()`

UnsetPhysicalDiskExtensions ensures that no value is present for PhysicalDiskExtensions, not even an explicit nil
### GetRegisteredDevice

`func (o *StoragePhysicalDisk) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StoragePhysicalDisk) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StoragePhysicalDisk) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StoragePhysicalDisk) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *StoragePhysicalDisk) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *StoragePhysicalDisk) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
### GetRunningFirmware

`func (o *StoragePhysicalDisk) GetRunningFirmware() []FirmwareRunningFirmwareRelationship`

GetRunningFirmware returns the RunningFirmware field if non-nil, zero value otherwise.

### GetRunningFirmwareOk

`func (o *StoragePhysicalDisk) GetRunningFirmwareOk() (*[]FirmwareRunningFirmwareRelationship, bool)`

GetRunningFirmwareOk returns a tuple with the RunningFirmware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningFirmware

`func (o *StoragePhysicalDisk) SetRunningFirmware(v []FirmwareRunningFirmwareRelationship)`

SetRunningFirmware sets RunningFirmware field to given value.

### HasRunningFirmware

`func (o *StoragePhysicalDisk) HasRunningFirmware() bool`

HasRunningFirmware returns a boolean if a field has been set.

### SetRunningFirmwareNil

`func (o *StoragePhysicalDisk) SetRunningFirmwareNil(b bool)`

 SetRunningFirmwareNil sets the value for RunningFirmware to be an explicit nil

### UnsetRunningFirmware
`func (o *StoragePhysicalDisk) UnsetRunningFirmware()`

UnsetRunningFirmware ensures that no value is present for RunningFirmware, not even an explicit nil
### GetSasPorts

`func (o *StoragePhysicalDisk) GetSasPorts() []StorageSasPortRelationship`

GetSasPorts returns the SasPorts field if non-nil, zero value otherwise.

### GetSasPortsOk

`func (o *StoragePhysicalDisk) GetSasPortsOk() (*[]StorageSasPortRelationship, bool)`

GetSasPortsOk returns a tuple with the SasPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSasPorts

`func (o *StoragePhysicalDisk) SetSasPorts(v []StorageSasPortRelationship)`

SetSasPorts sets SasPorts field to given value.

### HasSasPorts

`func (o *StoragePhysicalDisk) HasSasPorts() bool`

HasSasPorts returns a boolean if a field has been set.

### SetSasPortsNil

`func (o *StoragePhysicalDisk) SetSasPortsNil(b bool)`

 SetSasPortsNil sets the value for SasPorts to be an explicit nil

### UnsetSasPorts
`func (o *StoragePhysicalDisk) UnsetSasPorts()`

UnsetSasPorts ensures that no value is present for SasPorts, not even an explicit nil
### GetStorageController

`func (o *StoragePhysicalDisk) GetStorageController() StorageControllerRelationship`

GetStorageController returns the StorageController field if non-nil, zero value otherwise.

### GetStorageControllerOk

`func (o *StoragePhysicalDisk) GetStorageControllerOk() (*StorageControllerRelationship, bool)`

GetStorageControllerOk returns a tuple with the StorageController field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageController

`func (o *StoragePhysicalDisk) SetStorageController(v StorageControllerRelationship)`

SetStorageController sets StorageController field to given value.

### HasStorageController

`func (o *StoragePhysicalDisk) HasStorageController() bool`

HasStorageController returns a boolean if a field has been set.

### SetStorageControllerNil

`func (o *StoragePhysicalDisk) SetStorageControllerNil(b bool)`

 SetStorageControllerNil sets the value for StorageController to be an explicit nil

### UnsetStorageController
`func (o *StoragePhysicalDisk) UnsetStorageController()`

UnsetStorageController ensures that no value is present for StorageController, not even an explicit nil
### GetStorageEnclosure

`func (o *StoragePhysicalDisk) GetStorageEnclosure() StorageEnclosureRelationship`

GetStorageEnclosure returns the StorageEnclosure field if non-nil, zero value otherwise.

### GetStorageEnclosureOk

`func (o *StoragePhysicalDisk) GetStorageEnclosureOk() (*StorageEnclosureRelationship, bool)`

GetStorageEnclosureOk returns a tuple with the StorageEnclosure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageEnclosure

`func (o *StoragePhysicalDisk) SetStorageEnclosure(v StorageEnclosureRelationship)`

SetStorageEnclosure sets StorageEnclosure field to given value.

### HasStorageEnclosure

`func (o *StoragePhysicalDisk) HasStorageEnclosure() bool`

HasStorageEnclosure returns a boolean if a field has been set.

### SetStorageEnclosureNil

`func (o *StoragePhysicalDisk) SetStorageEnclosureNil(b bool)`

 SetStorageEnclosureNil sets the value for StorageEnclosure to be an explicit nil

### UnsetStorageEnclosure
`func (o *StoragePhysicalDisk) UnsetStorageEnclosure()`

UnsetStorageEnclosure ensures that no value is present for StorageEnclosure, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



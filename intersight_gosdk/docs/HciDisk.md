# HciDisk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hci.Disk"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hci.Disk"]
**ClusterExtId** | Pointer to **string** | The unique identifier of the cluster. | [optional] [readonly] 
**ClusterName** | Pointer to **string** | The name of the cluster this disk belongs to. | [optional] [readonly] 
**CvmIpAddress** | Pointer to [**NullableHciIpAddress**](HciIpAddress.md) |  | [optional] 
**DiskExtId** | Pointer to **string** | The unique identifier of the disk. | [optional] [readonly] 
**DiskSizeBytes** | Pointer to **int64** | The size of the disk in bytes. | [optional] [readonly] 
**FirmwareVersion** | Pointer to **string** | The current firmware version of the disk. | [optional] [readonly] 
**HasBootPartitionsOnly** | Pointer to **bool** | Indicates if the disk is boot only and no disk operation to be run on it. | [optional] [readonly] 
**HostName** | Pointer to **string** | The name of the host the disk is running on. | [optional] [readonly] 
**IsBootDisk** | Pointer to **bool** | Indicate if the disk is a boot disk. | [optional] [readonly] 
**IsDataMigrated** | Pointer to **bool** | Indicates whether the disk data is migrated. | [optional] [readonly] 
**IsDiagnosticInfoAvailable** | Pointer to **bool** | Indicates whether the diagnostic information is available. | [optional] [readonly] 
**IsErrorFoundInLogs** | Pointer to **bool** | Indicates whether the error is found in kernel logs. | [optional] [readonly] 
**IsMarkedForRemoval** | Pointer to **bool** | Indicates whether the disk is marked for removal. | [optional] [readonly] 
**IsMounted** | Pointer to **bool** | Indicates whether the disk is mounted. | [optional] [readonly] 
**IsOnline** | Pointer to **bool** | Indicates whether the disk is online or offline. | [optional] [readonly] 
**IsPasswordProtected** | Pointer to **bool** | The password protection status of the disk. | [optional] [readonly] 
**IsPlannedOutage** | Pointer to **bool** | Indicates if diagnostics are running on the Disk. | [optional] [readonly] 
**IsSelfEncryptingDrive** | Pointer to **bool** | The self-encrypting drive status of the disk. | [optional] [readonly] 
**IsSelfManagedNvme** | Pointer to **bool** | Indicates if the NVMe Disk is self managed and no host/CVM reboot is required. | [optional] [readonly] 
**IsSpdkManaged** | Pointer to **bool** | Indicates if NVMe device is managed by storage performance development  kit (SPDK). | [optional] [readonly] 
**IsSuspectedUnhealthy** | Pointer to **bool** | Indicates whether the disk is suspected unhealthy. | [optional] [readonly] 
**IsUnderDiagnosis** | Pointer to **bool** | Indicates whether the disk is under diagnosis. | [optional] [readonly] 
**IsUnhealthy** | Pointer to **bool** | Indicates whether the disk is unhealthy. | [optional] [readonly] 
**Location** | Pointer to **int64** | The location of the disk in the node. | [optional] [readonly] 
**Model** | Pointer to **string** | The model of the reported disk. | [optional] [readonly] 
**MountPath** | Pointer to **string** | The mount path of the disk. | [optional] [readonly] 
**NodeExtId** | Pointer to **string** | The unique identifier of the node. | [optional] [readonly] 
**NodeIpAddress** | Pointer to [**NullableHciIpAddress**](HciIpAddress.md) |  | [optional] 
**NvmePciePath** | Pointer to **string** | The PCIe path of the NVMe disk. | [optional] [readonly] 
**PhysicalCapacityBytes** | Pointer to **int64** | The physical capacity of the disk in bytes. | [optional] [readonly] 
**SerialNumber** | Pointer to **string** | The serial number of the disk. | [optional] [readonly] 
**ServiceVmId** | Pointer to **string** | The unique identifier of the service VM on the node. | [optional] [readonly] 
**Status** | Pointer to **string** | The status of the disk such as NORMAL, MARKED_FOR_REMOVAL_BUT_NOT_DETACHABLE, DETACHABLE, DATA_MIGRATION_INITIATED. | [optional] [readonly] 
**StoragePoolExtId** | Pointer to **string** | The unique identifier of the storage pool. | [optional] [readonly] 
**StorageTier** | Pointer to **string** | The storage tier of the disk such as SSD_PCIE, SSD_SATA, DAS_SATA, CLOUD, SSD_MEM_NVME. | [optional] [readonly] 
**TargetFirmwareVersion** | Pointer to **string** | The target firmware version of the disk. | [optional] [readonly] 
**Vendor** | Pointer to **string** | The vendor of the reported disk. | [optional] [readonly] 
**Node** | Pointer to [**NullableHciNodeRelationship**](HciNodeRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewHciDisk

`func NewHciDisk(classId string, objectType string, ) *HciDisk`

NewHciDisk instantiates a new HciDisk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHciDiskWithDefaults

`func NewHciDiskWithDefaults() *HciDisk`

NewHciDiskWithDefaults instantiates a new HciDisk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HciDisk) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HciDisk) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HciDisk) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HciDisk) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HciDisk) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HciDisk) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetClusterExtId

`func (o *HciDisk) GetClusterExtId() string`

GetClusterExtId returns the ClusterExtId field if non-nil, zero value otherwise.

### GetClusterExtIdOk

`func (o *HciDisk) GetClusterExtIdOk() (*string, bool)`

GetClusterExtIdOk returns a tuple with the ClusterExtId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterExtId

`func (o *HciDisk) SetClusterExtId(v string)`

SetClusterExtId sets ClusterExtId field to given value.

### HasClusterExtId

`func (o *HciDisk) HasClusterExtId() bool`

HasClusterExtId returns a boolean if a field has been set.

### GetClusterName

`func (o *HciDisk) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *HciDisk) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *HciDisk) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *HciDisk) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.

### GetCvmIpAddress

`func (o *HciDisk) GetCvmIpAddress() HciIpAddress`

GetCvmIpAddress returns the CvmIpAddress field if non-nil, zero value otherwise.

### GetCvmIpAddressOk

`func (o *HciDisk) GetCvmIpAddressOk() (*HciIpAddress, bool)`

GetCvmIpAddressOk returns a tuple with the CvmIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvmIpAddress

`func (o *HciDisk) SetCvmIpAddress(v HciIpAddress)`

SetCvmIpAddress sets CvmIpAddress field to given value.

### HasCvmIpAddress

`func (o *HciDisk) HasCvmIpAddress() bool`

HasCvmIpAddress returns a boolean if a field has been set.

### SetCvmIpAddressNil

`func (o *HciDisk) SetCvmIpAddressNil(b bool)`

 SetCvmIpAddressNil sets the value for CvmIpAddress to be an explicit nil

### UnsetCvmIpAddress
`func (o *HciDisk) UnsetCvmIpAddress()`

UnsetCvmIpAddress ensures that no value is present for CvmIpAddress, not even an explicit nil
### GetDiskExtId

`func (o *HciDisk) GetDiskExtId() string`

GetDiskExtId returns the DiskExtId field if non-nil, zero value otherwise.

### GetDiskExtIdOk

`func (o *HciDisk) GetDiskExtIdOk() (*string, bool)`

GetDiskExtIdOk returns a tuple with the DiskExtId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskExtId

`func (o *HciDisk) SetDiskExtId(v string)`

SetDiskExtId sets DiskExtId field to given value.

### HasDiskExtId

`func (o *HciDisk) HasDiskExtId() bool`

HasDiskExtId returns a boolean if a field has been set.

### GetDiskSizeBytes

`func (o *HciDisk) GetDiskSizeBytes() int64`

GetDiskSizeBytes returns the DiskSizeBytes field if non-nil, zero value otherwise.

### GetDiskSizeBytesOk

`func (o *HciDisk) GetDiskSizeBytesOk() (*int64, bool)`

GetDiskSizeBytesOk returns a tuple with the DiskSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSizeBytes

`func (o *HciDisk) SetDiskSizeBytes(v int64)`

SetDiskSizeBytes sets DiskSizeBytes field to given value.

### HasDiskSizeBytes

`func (o *HciDisk) HasDiskSizeBytes() bool`

HasDiskSizeBytes returns a boolean if a field has been set.

### GetFirmwareVersion

`func (o *HciDisk) GetFirmwareVersion() string`

GetFirmwareVersion returns the FirmwareVersion field if non-nil, zero value otherwise.

### GetFirmwareVersionOk

`func (o *HciDisk) GetFirmwareVersionOk() (*string, bool)`

GetFirmwareVersionOk returns a tuple with the FirmwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmwareVersion

`func (o *HciDisk) SetFirmwareVersion(v string)`

SetFirmwareVersion sets FirmwareVersion field to given value.

### HasFirmwareVersion

`func (o *HciDisk) HasFirmwareVersion() bool`

HasFirmwareVersion returns a boolean if a field has been set.

### GetHasBootPartitionsOnly

`func (o *HciDisk) GetHasBootPartitionsOnly() bool`

GetHasBootPartitionsOnly returns the HasBootPartitionsOnly field if non-nil, zero value otherwise.

### GetHasBootPartitionsOnlyOk

`func (o *HciDisk) GetHasBootPartitionsOnlyOk() (*bool, bool)`

GetHasBootPartitionsOnlyOk returns a tuple with the HasBootPartitionsOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasBootPartitionsOnly

`func (o *HciDisk) SetHasBootPartitionsOnly(v bool)`

SetHasBootPartitionsOnly sets HasBootPartitionsOnly field to given value.

### HasHasBootPartitionsOnly

`func (o *HciDisk) HasHasBootPartitionsOnly() bool`

HasHasBootPartitionsOnly returns a boolean if a field has been set.

### GetHostName

`func (o *HciDisk) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *HciDisk) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *HciDisk) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *HciDisk) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetIsBootDisk

`func (o *HciDisk) GetIsBootDisk() bool`

GetIsBootDisk returns the IsBootDisk field if non-nil, zero value otherwise.

### GetIsBootDiskOk

`func (o *HciDisk) GetIsBootDiskOk() (*bool, bool)`

GetIsBootDiskOk returns a tuple with the IsBootDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBootDisk

`func (o *HciDisk) SetIsBootDisk(v bool)`

SetIsBootDisk sets IsBootDisk field to given value.

### HasIsBootDisk

`func (o *HciDisk) HasIsBootDisk() bool`

HasIsBootDisk returns a boolean if a field has been set.

### GetIsDataMigrated

`func (o *HciDisk) GetIsDataMigrated() bool`

GetIsDataMigrated returns the IsDataMigrated field if non-nil, zero value otherwise.

### GetIsDataMigratedOk

`func (o *HciDisk) GetIsDataMigratedOk() (*bool, bool)`

GetIsDataMigratedOk returns a tuple with the IsDataMigrated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDataMigrated

`func (o *HciDisk) SetIsDataMigrated(v bool)`

SetIsDataMigrated sets IsDataMigrated field to given value.

### HasIsDataMigrated

`func (o *HciDisk) HasIsDataMigrated() bool`

HasIsDataMigrated returns a boolean if a field has been set.

### GetIsDiagnosticInfoAvailable

`func (o *HciDisk) GetIsDiagnosticInfoAvailable() bool`

GetIsDiagnosticInfoAvailable returns the IsDiagnosticInfoAvailable field if non-nil, zero value otherwise.

### GetIsDiagnosticInfoAvailableOk

`func (o *HciDisk) GetIsDiagnosticInfoAvailableOk() (*bool, bool)`

GetIsDiagnosticInfoAvailableOk returns a tuple with the IsDiagnosticInfoAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDiagnosticInfoAvailable

`func (o *HciDisk) SetIsDiagnosticInfoAvailable(v bool)`

SetIsDiagnosticInfoAvailable sets IsDiagnosticInfoAvailable field to given value.

### HasIsDiagnosticInfoAvailable

`func (o *HciDisk) HasIsDiagnosticInfoAvailable() bool`

HasIsDiagnosticInfoAvailable returns a boolean if a field has been set.

### GetIsErrorFoundInLogs

`func (o *HciDisk) GetIsErrorFoundInLogs() bool`

GetIsErrorFoundInLogs returns the IsErrorFoundInLogs field if non-nil, zero value otherwise.

### GetIsErrorFoundInLogsOk

`func (o *HciDisk) GetIsErrorFoundInLogsOk() (*bool, bool)`

GetIsErrorFoundInLogsOk returns a tuple with the IsErrorFoundInLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsErrorFoundInLogs

`func (o *HciDisk) SetIsErrorFoundInLogs(v bool)`

SetIsErrorFoundInLogs sets IsErrorFoundInLogs field to given value.

### HasIsErrorFoundInLogs

`func (o *HciDisk) HasIsErrorFoundInLogs() bool`

HasIsErrorFoundInLogs returns a boolean if a field has been set.

### GetIsMarkedForRemoval

`func (o *HciDisk) GetIsMarkedForRemoval() bool`

GetIsMarkedForRemoval returns the IsMarkedForRemoval field if non-nil, zero value otherwise.

### GetIsMarkedForRemovalOk

`func (o *HciDisk) GetIsMarkedForRemovalOk() (*bool, bool)`

GetIsMarkedForRemovalOk returns a tuple with the IsMarkedForRemoval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMarkedForRemoval

`func (o *HciDisk) SetIsMarkedForRemoval(v bool)`

SetIsMarkedForRemoval sets IsMarkedForRemoval field to given value.

### HasIsMarkedForRemoval

`func (o *HciDisk) HasIsMarkedForRemoval() bool`

HasIsMarkedForRemoval returns a boolean if a field has been set.

### GetIsMounted

`func (o *HciDisk) GetIsMounted() bool`

GetIsMounted returns the IsMounted field if non-nil, zero value otherwise.

### GetIsMountedOk

`func (o *HciDisk) GetIsMountedOk() (*bool, bool)`

GetIsMountedOk returns a tuple with the IsMounted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMounted

`func (o *HciDisk) SetIsMounted(v bool)`

SetIsMounted sets IsMounted field to given value.

### HasIsMounted

`func (o *HciDisk) HasIsMounted() bool`

HasIsMounted returns a boolean if a field has been set.

### GetIsOnline

`func (o *HciDisk) GetIsOnline() bool`

GetIsOnline returns the IsOnline field if non-nil, zero value otherwise.

### GetIsOnlineOk

`func (o *HciDisk) GetIsOnlineOk() (*bool, bool)`

GetIsOnlineOk returns a tuple with the IsOnline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOnline

`func (o *HciDisk) SetIsOnline(v bool)`

SetIsOnline sets IsOnline field to given value.

### HasIsOnline

`func (o *HciDisk) HasIsOnline() bool`

HasIsOnline returns a boolean if a field has been set.

### GetIsPasswordProtected

`func (o *HciDisk) GetIsPasswordProtected() bool`

GetIsPasswordProtected returns the IsPasswordProtected field if non-nil, zero value otherwise.

### GetIsPasswordProtectedOk

`func (o *HciDisk) GetIsPasswordProtectedOk() (*bool, bool)`

GetIsPasswordProtectedOk returns a tuple with the IsPasswordProtected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPasswordProtected

`func (o *HciDisk) SetIsPasswordProtected(v bool)`

SetIsPasswordProtected sets IsPasswordProtected field to given value.

### HasIsPasswordProtected

`func (o *HciDisk) HasIsPasswordProtected() bool`

HasIsPasswordProtected returns a boolean if a field has been set.

### GetIsPlannedOutage

`func (o *HciDisk) GetIsPlannedOutage() bool`

GetIsPlannedOutage returns the IsPlannedOutage field if non-nil, zero value otherwise.

### GetIsPlannedOutageOk

`func (o *HciDisk) GetIsPlannedOutageOk() (*bool, bool)`

GetIsPlannedOutageOk returns a tuple with the IsPlannedOutage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPlannedOutage

`func (o *HciDisk) SetIsPlannedOutage(v bool)`

SetIsPlannedOutage sets IsPlannedOutage field to given value.

### HasIsPlannedOutage

`func (o *HciDisk) HasIsPlannedOutage() bool`

HasIsPlannedOutage returns a boolean if a field has been set.

### GetIsSelfEncryptingDrive

`func (o *HciDisk) GetIsSelfEncryptingDrive() bool`

GetIsSelfEncryptingDrive returns the IsSelfEncryptingDrive field if non-nil, zero value otherwise.

### GetIsSelfEncryptingDriveOk

`func (o *HciDisk) GetIsSelfEncryptingDriveOk() (*bool, bool)`

GetIsSelfEncryptingDriveOk returns a tuple with the IsSelfEncryptingDrive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSelfEncryptingDrive

`func (o *HciDisk) SetIsSelfEncryptingDrive(v bool)`

SetIsSelfEncryptingDrive sets IsSelfEncryptingDrive field to given value.

### HasIsSelfEncryptingDrive

`func (o *HciDisk) HasIsSelfEncryptingDrive() bool`

HasIsSelfEncryptingDrive returns a boolean if a field has been set.

### GetIsSelfManagedNvme

`func (o *HciDisk) GetIsSelfManagedNvme() bool`

GetIsSelfManagedNvme returns the IsSelfManagedNvme field if non-nil, zero value otherwise.

### GetIsSelfManagedNvmeOk

`func (o *HciDisk) GetIsSelfManagedNvmeOk() (*bool, bool)`

GetIsSelfManagedNvmeOk returns a tuple with the IsSelfManagedNvme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSelfManagedNvme

`func (o *HciDisk) SetIsSelfManagedNvme(v bool)`

SetIsSelfManagedNvme sets IsSelfManagedNvme field to given value.

### HasIsSelfManagedNvme

`func (o *HciDisk) HasIsSelfManagedNvme() bool`

HasIsSelfManagedNvme returns a boolean if a field has been set.

### GetIsSpdkManaged

`func (o *HciDisk) GetIsSpdkManaged() bool`

GetIsSpdkManaged returns the IsSpdkManaged field if non-nil, zero value otherwise.

### GetIsSpdkManagedOk

`func (o *HciDisk) GetIsSpdkManagedOk() (*bool, bool)`

GetIsSpdkManagedOk returns a tuple with the IsSpdkManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSpdkManaged

`func (o *HciDisk) SetIsSpdkManaged(v bool)`

SetIsSpdkManaged sets IsSpdkManaged field to given value.

### HasIsSpdkManaged

`func (o *HciDisk) HasIsSpdkManaged() bool`

HasIsSpdkManaged returns a boolean if a field has been set.

### GetIsSuspectedUnhealthy

`func (o *HciDisk) GetIsSuspectedUnhealthy() bool`

GetIsSuspectedUnhealthy returns the IsSuspectedUnhealthy field if non-nil, zero value otherwise.

### GetIsSuspectedUnhealthyOk

`func (o *HciDisk) GetIsSuspectedUnhealthyOk() (*bool, bool)`

GetIsSuspectedUnhealthyOk returns a tuple with the IsSuspectedUnhealthy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSuspectedUnhealthy

`func (o *HciDisk) SetIsSuspectedUnhealthy(v bool)`

SetIsSuspectedUnhealthy sets IsSuspectedUnhealthy field to given value.

### HasIsSuspectedUnhealthy

`func (o *HciDisk) HasIsSuspectedUnhealthy() bool`

HasIsSuspectedUnhealthy returns a boolean if a field has been set.

### GetIsUnderDiagnosis

`func (o *HciDisk) GetIsUnderDiagnosis() bool`

GetIsUnderDiagnosis returns the IsUnderDiagnosis field if non-nil, zero value otherwise.

### GetIsUnderDiagnosisOk

`func (o *HciDisk) GetIsUnderDiagnosisOk() (*bool, bool)`

GetIsUnderDiagnosisOk returns a tuple with the IsUnderDiagnosis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUnderDiagnosis

`func (o *HciDisk) SetIsUnderDiagnosis(v bool)`

SetIsUnderDiagnosis sets IsUnderDiagnosis field to given value.

### HasIsUnderDiagnosis

`func (o *HciDisk) HasIsUnderDiagnosis() bool`

HasIsUnderDiagnosis returns a boolean if a field has been set.

### GetIsUnhealthy

`func (o *HciDisk) GetIsUnhealthy() bool`

GetIsUnhealthy returns the IsUnhealthy field if non-nil, zero value otherwise.

### GetIsUnhealthyOk

`func (o *HciDisk) GetIsUnhealthyOk() (*bool, bool)`

GetIsUnhealthyOk returns a tuple with the IsUnhealthy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUnhealthy

`func (o *HciDisk) SetIsUnhealthy(v bool)`

SetIsUnhealthy sets IsUnhealthy field to given value.

### HasIsUnhealthy

`func (o *HciDisk) HasIsUnhealthy() bool`

HasIsUnhealthy returns a boolean if a field has been set.

### GetLocation

`func (o *HciDisk) GetLocation() int64`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *HciDisk) GetLocationOk() (*int64, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *HciDisk) SetLocation(v int64)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *HciDisk) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetModel

`func (o *HciDisk) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *HciDisk) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *HciDisk) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *HciDisk) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetMountPath

`func (o *HciDisk) GetMountPath() string`

GetMountPath returns the MountPath field if non-nil, zero value otherwise.

### GetMountPathOk

`func (o *HciDisk) GetMountPathOk() (*string, bool)`

GetMountPathOk returns a tuple with the MountPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountPath

`func (o *HciDisk) SetMountPath(v string)`

SetMountPath sets MountPath field to given value.

### HasMountPath

`func (o *HciDisk) HasMountPath() bool`

HasMountPath returns a boolean if a field has been set.

### GetNodeExtId

`func (o *HciDisk) GetNodeExtId() string`

GetNodeExtId returns the NodeExtId field if non-nil, zero value otherwise.

### GetNodeExtIdOk

`func (o *HciDisk) GetNodeExtIdOk() (*string, bool)`

GetNodeExtIdOk returns a tuple with the NodeExtId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeExtId

`func (o *HciDisk) SetNodeExtId(v string)`

SetNodeExtId sets NodeExtId field to given value.

### HasNodeExtId

`func (o *HciDisk) HasNodeExtId() bool`

HasNodeExtId returns a boolean if a field has been set.

### GetNodeIpAddress

`func (o *HciDisk) GetNodeIpAddress() HciIpAddress`

GetNodeIpAddress returns the NodeIpAddress field if non-nil, zero value otherwise.

### GetNodeIpAddressOk

`func (o *HciDisk) GetNodeIpAddressOk() (*HciIpAddress, bool)`

GetNodeIpAddressOk returns a tuple with the NodeIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeIpAddress

`func (o *HciDisk) SetNodeIpAddress(v HciIpAddress)`

SetNodeIpAddress sets NodeIpAddress field to given value.

### HasNodeIpAddress

`func (o *HciDisk) HasNodeIpAddress() bool`

HasNodeIpAddress returns a boolean if a field has been set.

### SetNodeIpAddressNil

`func (o *HciDisk) SetNodeIpAddressNil(b bool)`

 SetNodeIpAddressNil sets the value for NodeIpAddress to be an explicit nil

### UnsetNodeIpAddress
`func (o *HciDisk) UnsetNodeIpAddress()`

UnsetNodeIpAddress ensures that no value is present for NodeIpAddress, not even an explicit nil
### GetNvmePciePath

`func (o *HciDisk) GetNvmePciePath() string`

GetNvmePciePath returns the NvmePciePath field if non-nil, zero value otherwise.

### GetNvmePciePathOk

`func (o *HciDisk) GetNvmePciePathOk() (*string, bool)`

GetNvmePciePathOk returns a tuple with the NvmePciePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNvmePciePath

`func (o *HciDisk) SetNvmePciePath(v string)`

SetNvmePciePath sets NvmePciePath field to given value.

### HasNvmePciePath

`func (o *HciDisk) HasNvmePciePath() bool`

HasNvmePciePath returns a boolean if a field has been set.

### GetPhysicalCapacityBytes

`func (o *HciDisk) GetPhysicalCapacityBytes() int64`

GetPhysicalCapacityBytes returns the PhysicalCapacityBytes field if non-nil, zero value otherwise.

### GetPhysicalCapacityBytesOk

`func (o *HciDisk) GetPhysicalCapacityBytesOk() (*int64, bool)`

GetPhysicalCapacityBytesOk returns a tuple with the PhysicalCapacityBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalCapacityBytes

`func (o *HciDisk) SetPhysicalCapacityBytes(v int64)`

SetPhysicalCapacityBytes sets PhysicalCapacityBytes field to given value.

### HasPhysicalCapacityBytes

`func (o *HciDisk) HasPhysicalCapacityBytes() bool`

HasPhysicalCapacityBytes returns a boolean if a field has been set.

### GetSerialNumber

`func (o *HciDisk) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *HciDisk) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *HciDisk) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *HciDisk) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetServiceVmId

`func (o *HciDisk) GetServiceVmId() string`

GetServiceVmId returns the ServiceVmId field if non-nil, zero value otherwise.

### GetServiceVmIdOk

`func (o *HciDisk) GetServiceVmIdOk() (*string, bool)`

GetServiceVmIdOk returns a tuple with the ServiceVmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceVmId

`func (o *HciDisk) SetServiceVmId(v string)`

SetServiceVmId sets ServiceVmId field to given value.

### HasServiceVmId

`func (o *HciDisk) HasServiceVmId() bool`

HasServiceVmId returns a boolean if a field has been set.

### GetStatus

`func (o *HciDisk) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HciDisk) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HciDisk) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HciDisk) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStoragePoolExtId

`func (o *HciDisk) GetStoragePoolExtId() string`

GetStoragePoolExtId returns the StoragePoolExtId field if non-nil, zero value otherwise.

### GetStoragePoolExtIdOk

`func (o *HciDisk) GetStoragePoolExtIdOk() (*string, bool)`

GetStoragePoolExtIdOk returns a tuple with the StoragePoolExtId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoragePoolExtId

`func (o *HciDisk) SetStoragePoolExtId(v string)`

SetStoragePoolExtId sets StoragePoolExtId field to given value.

### HasStoragePoolExtId

`func (o *HciDisk) HasStoragePoolExtId() bool`

HasStoragePoolExtId returns a boolean if a field has been set.

### GetStorageTier

`func (o *HciDisk) GetStorageTier() string`

GetStorageTier returns the StorageTier field if non-nil, zero value otherwise.

### GetStorageTierOk

`func (o *HciDisk) GetStorageTierOk() (*string, bool)`

GetStorageTierOk returns a tuple with the StorageTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageTier

`func (o *HciDisk) SetStorageTier(v string)`

SetStorageTier sets StorageTier field to given value.

### HasStorageTier

`func (o *HciDisk) HasStorageTier() bool`

HasStorageTier returns a boolean if a field has been set.

### GetTargetFirmwareVersion

`func (o *HciDisk) GetTargetFirmwareVersion() string`

GetTargetFirmwareVersion returns the TargetFirmwareVersion field if non-nil, zero value otherwise.

### GetTargetFirmwareVersionOk

`func (o *HciDisk) GetTargetFirmwareVersionOk() (*string, bool)`

GetTargetFirmwareVersionOk returns a tuple with the TargetFirmwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetFirmwareVersion

`func (o *HciDisk) SetTargetFirmwareVersion(v string)`

SetTargetFirmwareVersion sets TargetFirmwareVersion field to given value.

### HasTargetFirmwareVersion

`func (o *HciDisk) HasTargetFirmwareVersion() bool`

HasTargetFirmwareVersion returns a boolean if a field has been set.

### GetVendor

`func (o *HciDisk) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *HciDisk) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *HciDisk) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *HciDisk) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetNode

`func (o *HciDisk) GetNode() HciNodeRelationship`

GetNode returns the Node field if non-nil, zero value otherwise.

### GetNodeOk

`func (o *HciDisk) GetNodeOk() (*HciNodeRelationship, bool)`

GetNodeOk returns a tuple with the Node field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNode

`func (o *HciDisk) SetNode(v HciNodeRelationship)`

SetNode sets Node field to given value.

### HasNode

`func (o *HciDisk) HasNode() bool`

HasNode returns a boolean if a field has been set.

### SetNodeNil

`func (o *HciDisk) SetNodeNil(b bool)`

 SetNodeNil sets the value for Node to be an explicit nil

### UnsetNode
`func (o *HciDisk) UnsetNode()`

UnsetNode ensures that no value is present for Node, not even an explicit nil
### GetRegisteredDevice

`func (o *HciDisk) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *HciDisk) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *HciDisk) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *HciDisk) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *HciDisk) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *HciDisk) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



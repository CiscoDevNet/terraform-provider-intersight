# ComputePhysicalAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**AdminPowerState** | Pointer to **string** | The desired power state of the server. | [optional] [readonly] 
**AlarmSummary** | Pointer to [**NullableComputeAlarmSummary**](ComputeAlarmSummary.md) |  | [optional] 
**AssetTag** | Pointer to **string** | The user defined asset tag assigned to the server. | [optional] [readonly] 
**AvailableMemory** | Pointer to **int64** | The amount of memory available on the server. | [optional] [readonly] 
**BiosPostComplete** | Pointer to **bool** | The BIOS POST completion status of the server. | [optional] 
**FaultSummary** | Pointer to **int64** | The fault summary for the server. | [optional] 
**FrontPanelLockState** | Pointer to **string** | The actual front panel state of the server. * &#x60;None&#x60; - Front Panel of the server is set to None state. It is required so that the next frontPanelLockState operation can be triggered. * &#x60;Lock&#x60; - Front Panel of the server is set to Locked state. * &#x60;Unlock&#x60; - Front Panel of the server is set to Unlocked state. | [optional] [default to "None"]
**HardwareUuid** | Pointer to **string** | The universally unique hardware identity of the server provided by the manufacturer. | [optional] 
**KvmIpAddresses** | Pointer to [**[]ComputeIpAddress**](ComputeIpAddress.md) |  | [optional] 
**KvmServerStateEnabled** | Pointer to **bool** | The KVM server state of the server. | [optional] [readonly] 
**KvmVendor** | Pointer to **string** | The KVM Vendor for the server. | [optional] [readonly] 
**Lifecycle** | Pointer to **string** | The lifecycle state of the server. This will map to the discovery lifecycle as represented in the server Identity object. * &#x60;None&#x60; - Default state of an equipment. This should be an initial state when no state is defined for an equipment. * &#x60;Active&#x60; - Default Lifecycle State for a physical entity. * &#x60;Decommissioned&#x60; - Decommission Lifecycle state. * &#x60;DiscoveryInProgress&#x60; - DiscoveryInProgress Lifecycle state. * &#x60;DiscoveryFailed&#x60; - DiscoveryFailed Lifecycle state. * &#x60;FirmwareUpgradeInProgress&#x60; - Firmware upgrade is in progress on given physical entity. * &#x60;BladeMigrationInProgress&#x60; - Server slot migration is in progress on given physical entity. * &#x60;SlotMismatch&#x60; - The blade server is detected in a different chassis/slot than it was previously. | [optional] [readonly] [default to "None"]
**ManagementMode** | Pointer to **string** | The management mode of the server. * &#x60;IntersightStandalone&#x60; - Intersight Standalone mode of operation. * &#x60;UCSM&#x60; - Unified Computing System Manager mode of operation. * &#x60;Intersight&#x60; - Intersight managed mode of operation. | [optional] [default to "IntersightStandalone"]
**MemorySpeed** | Pointer to **string** | The maximum memory speed in MHz available on the server. | [optional] [readonly] 
**MgmtIpAddress** | Pointer to **string** | Management address of the server. | [optional] 
**Name** | Pointer to **string** | The name of the UCS Fabric Interconnect cluster or Cisco Integrated Management Controller (CIMC). When this server is attached to a UCS Fabric Interconnect, the value of this property is the name of the UCS Fabric Interconnect along with chassis/server Id. When this server configured in standalone mode, the value of this property is the name of the Cisco Integrated Management Controller. when this server is configired in IMM mode, the value of this property contains model and chassis/server Id. | [optional] 
**NumAdaptors** | Pointer to **int64** | The total number of network adapters present on the server. | [optional] [readonly] 
**NumCpuCores** | Pointer to **int64** | The total number of CPU cores present on the server. | [optional] [readonly] 
**NumCpuCoresEnabled** | Pointer to **int64** | The total number of CPU cores enabled on the server. | [optional] [readonly] 
**NumCpus** | Pointer to **int64** | The total number of CPUs present on the server. | [optional] [readonly] 
**NumEthHostInterfaces** | Pointer to **int64** | The total number of vNICs which are visible to a host on the server. | [optional] [readonly] 
**NumFcHostInterfaces** | Pointer to **int64** | The total number of vHBAs which are visible to a host on the server. | [optional] [readonly] 
**NumThreads** | Pointer to **int64** | The total number of threads the server is capable of handling. | [optional] [readonly] 
**OperPowerState** | Pointer to **string** | The actual power state of the server. | [optional] [readonly] 
**OperReason** | Pointer to **[]string** |  | [optional] 
**OperState** | Pointer to **string** | The operational state of the server. | [optional] [readonly] 
**Operability** | Pointer to **string** | The operability of the server. | [optional] [readonly] 
**PlatformType** | Pointer to **string** | The platform type of the registered device - whether managed by UCSM or operating in standalone mode. | [optional] 
**ServiceProfile** | Pointer to **string** | The distinguished name of the service profile to which the server is associated to. It is applicable only for servers which are managed via UCSM. | [optional] [readonly] 
**TotalMemory** | Pointer to **int64** | The total memory available on the server. | [optional] [readonly] 
**TunneledKvm** | Pointer to **bool** | The Tunneled vKVM status of the server. | [optional] 
**UserLabel** | Pointer to **string** | The user defined label assigned to the server. | [optional] [readonly] 
**Uuid** | Pointer to **string** | The universally unique identity of the server. | [optional] [readonly] 
**BootCddDevices** | Pointer to [**[]BootCddDeviceRelationship**](BootCddDeviceRelationship.md) | An array of relationships to bootCddDevice resources. | [optional] 
**BootDeviceBootSecurity** | Pointer to [**BootDeviceBootSecurityRelationship**](BootDeviceBootSecurityRelationship.md) |  | [optional] 
**BootHddDevices** | Pointer to [**[]BootHddDeviceRelationship**](BootHddDeviceRelationship.md) | An array of relationships to bootHddDevice resources. | [optional] 
**BootIscsiDevices** | Pointer to [**[]BootIscsiDeviceRelationship**](BootIscsiDeviceRelationship.md) | An array of relationships to bootIscsiDevice resources. | [optional] 
**BootNvmeDevices** | Pointer to [**[]BootNvmeDeviceRelationship**](BootNvmeDeviceRelationship.md) | An array of relationships to bootNvmeDevice resources. | [optional] 
**BootPchStorageDevices** | Pointer to [**[]BootPchStorageDeviceRelationship**](BootPchStorageDeviceRelationship.md) | An array of relationships to bootPchStorageDevice resources. | [optional] 
**BootPxeDevices** | Pointer to [**[]BootPxeDeviceRelationship**](BootPxeDeviceRelationship.md) | An array of relationships to bootPxeDevice resources. | [optional] 
**BootSanDevices** | Pointer to [**[]BootSanDeviceRelationship**](BootSanDeviceRelationship.md) | An array of relationships to bootSanDevice resources. | [optional] 
**BootSdDevices** | Pointer to [**[]BootSdDeviceRelationship**](BootSdDeviceRelationship.md) | An array of relationships to bootSdDevice resources. | [optional] 
**BootUefiShellDevices** | Pointer to [**[]BootUefiShellDeviceRelationship**](BootUefiShellDeviceRelationship.md) | An array of relationships to bootUefiShellDevice resources. | [optional] 
**BootUsbDevices** | Pointer to [**[]BootUsbDeviceRelationship**](BootUsbDeviceRelationship.md) | An array of relationships to bootUsbDevice resources. | [optional] 
**BootVmediaDevices** | Pointer to [**[]BootVmediaDeviceRelationship**](BootVmediaDeviceRelationship.md) | An array of relationships to bootVmediaDevice resources. | [optional] 
**MgmtIdentity** | Pointer to [**EquipmentPhysicalIdentityRelationship**](EquipmentPhysicalIdentityRelationship.md) |  | [optional] 
**Vmedia** | Pointer to [**ComputeVmediaRelationship**](ComputeVmediaRelationship.md) |  | [optional] 

## Methods

### NewComputePhysicalAllOf

`func NewComputePhysicalAllOf(classId string, objectType string, ) *ComputePhysicalAllOf`

NewComputePhysicalAllOf instantiates a new ComputePhysicalAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputePhysicalAllOfWithDefaults

`func NewComputePhysicalAllOfWithDefaults() *ComputePhysicalAllOf`

NewComputePhysicalAllOfWithDefaults instantiates a new ComputePhysicalAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ComputePhysicalAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ComputePhysicalAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ComputePhysicalAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ComputePhysicalAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ComputePhysicalAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ComputePhysicalAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdminPowerState

`func (o *ComputePhysicalAllOf) GetAdminPowerState() string`

GetAdminPowerState returns the AdminPowerState field if non-nil, zero value otherwise.

### GetAdminPowerStateOk

`func (o *ComputePhysicalAllOf) GetAdminPowerStateOk() (*string, bool)`

GetAdminPowerStateOk returns a tuple with the AdminPowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminPowerState

`func (o *ComputePhysicalAllOf) SetAdminPowerState(v string)`

SetAdminPowerState sets AdminPowerState field to given value.

### HasAdminPowerState

`func (o *ComputePhysicalAllOf) HasAdminPowerState() bool`

HasAdminPowerState returns a boolean if a field has been set.

### GetAlarmSummary

`func (o *ComputePhysicalAllOf) GetAlarmSummary() ComputeAlarmSummary`

GetAlarmSummary returns the AlarmSummary field if non-nil, zero value otherwise.

### GetAlarmSummaryOk

`func (o *ComputePhysicalAllOf) GetAlarmSummaryOk() (*ComputeAlarmSummary, bool)`

GetAlarmSummaryOk returns a tuple with the AlarmSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmSummary

`func (o *ComputePhysicalAllOf) SetAlarmSummary(v ComputeAlarmSummary)`

SetAlarmSummary sets AlarmSummary field to given value.

### HasAlarmSummary

`func (o *ComputePhysicalAllOf) HasAlarmSummary() bool`

HasAlarmSummary returns a boolean if a field has been set.

### SetAlarmSummaryNil

`func (o *ComputePhysicalAllOf) SetAlarmSummaryNil(b bool)`

 SetAlarmSummaryNil sets the value for AlarmSummary to be an explicit nil

### UnsetAlarmSummary
`func (o *ComputePhysicalAllOf) UnsetAlarmSummary()`

UnsetAlarmSummary ensures that no value is present for AlarmSummary, not even an explicit nil
### GetAssetTag

`func (o *ComputePhysicalAllOf) GetAssetTag() string`

GetAssetTag returns the AssetTag field if non-nil, zero value otherwise.

### GetAssetTagOk

`func (o *ComputePhysicalAllOf) GetAssetTagOk() (*string, bool)`

GetAssetTagOk returns a tuple with the AssetTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetTag

`func (o *ComputePhysicalAllOf) SetAssetTag(v string)`

SetAssetTag sets AssetTag field to given value.

### HasAssetTag

`func (o *ComputePhysicalAllOf) HasAssetTag() bool`

HasAssetTag returns a boolean if a field has been set.

### GetAvailableMemory

`func (o *ComputePhysicalAllOf) GetAvailableMemory() int64`

GetAvailableMemory returns the AvailableMemory field if non-nil, zero value otherwise.

### GetAvailableMemoryOk

`func (o *ComputePhysicalAllOf) GetAvailableMemoryOk() (*int64, bool)`

GetAvailableMemoryOk returns a tuple with the AvailableMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableMemory

`func (o *ComputePhysicalAllOf) SetAvailableMemory(v int64)`

SetAvailableMemory sets AvailableMemory field to given value.

### HasAvailableMemory

`func (o *ComputePhysicalAllOf) HasAvailableMemory() bool`

HasAvailableMemory returns a boolean if a field has been set.

### GetBiosPostComplete

`func (o *ComputePhysicalAllOf) GetBiosPostComplete() bool`

GetBiosPostComplete returns the BiosPostComplete field if non-nil, zero value otherwise.

### GetBiosPostCompleteOk

`func (o *ComputePhysicalAllOf) GetBiosPostCompleteOk() (*bool, bool)`

GetBiosPostCompleteOk returns a tuple with the BiosPostComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiosPostComplete

`func (o *ComputePhysicalAllOf) SetBiosPostComplete(v bool)`

SetBiosPostComplete sets BiosPostComplete field to given value.

### HasBiosPostComplete

`func (o *ComputePhysicalAllOf) HasBiosPostComplete() bool`

HasBiosPostComplete returns a boolean if a field has been set.

### GetFaultSummary

`func (o *ComputePhysicalAllOf) GetFaultSummary() int64`

GetFaultSummary returns the FaultSummary field if non-nil, zero value otherwise.

### GetFaultSummaryOk

`func (o *ComputePhysicalAllOf) GetFaultSummaryOk() (*int64, bool)`

GetFaultSummaryOk returns a tuple with the FaultSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaultSummary

`func (o *ComputePhysicalAllOf) SetFaultSummary(v int64)`

SetFaultSummary sets FaultSummary field to given value.

### HasFaultSummary

`func (o *ComputePhysicalAllOf) HasFaultSummary() bool`

HasFaultSummary returns a boolean if a field has been set.

### GetFrontPanelLockState

`func (o *ComputePhysicalAllOf) GetFrontPanelLockState() string`

GetFrontPanelLockState returns the FrontPanelLockState field if non-nil, zero value otherwise.

### GetFrontPanelLockStateOk

`func (o *ComputePhysicalAllOf) GetFrontPanelLockStateOk() (*string, bool)`

GetFrontPanelLockStateOk returns a tuple with the FrontPanelLockState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrontPanelLockState

`func (o *ComputePhysicalAllOf) SetFrontPanelLockState(v string)`

SetFrontPanelLockState sets FrontPanelLockState field to given value.

### HasFrontPanelLockState

`func (o *ComputePhysicalAllOf) HasFrontPanelLockState() bool`

HasFrontPanelLockState returns a boolean if a field has been set.

### GetHardwareUuid

`func (o *ComputePhysicalAllOf) GetHardwareUuid() string`

GetHardwareUuid returns the HardwareUuid field if non-nil, zero value otherwise.

### GetHardwareUuidOk

`func (o *ComputePhysicalAllOf) GetHardwareUuidOk() (*string, bool)`

GetHardwareUuidOk returns a tuple with the HardwareUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareUuid

`func (o *ComputePhysicalAllOf) SetHardwareUuid(v string)`

SetHardwareUuid sets HardwareUuid field to given value.

### HasHardwareUuid

`func (o *ComputePhysicalAllOf) HasHardwareUuid() bool`

HasHardwareUuid returns a boolean if a field has been set.

### GetKvmIpAddresses

`func (o *ComputePhysicalAllOf) GetKvmIpAddresses() []ComputeIpAddress`

GetKvmIpAddresses returns the KvmIpAddresses field if non-nil, zero value otherwise.

### GetKvmIpAddressesOk

`func (o *ComputePhysicalAllOf) GetKvmIpAddressesOk() (*[]ComputeIpAddress, bool)`

GetKvmIpAddressesOk returns a tuple with the KvmIpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKvmIpAddresses

`func (o *ComputePhysicalAllOf) SetKvmIpAddresses(v []ComputeIpAddress)`

SetKvmIpAddresses sets KvmIpAddresses field to given value.

### HasKvmIpAddresses

`func (o *ComputePhysicalAllOf) HasKvmIpAddresses() bool`

HasKvmIpAddresses returns a boolean if a field has been set.

### SetKvmIpAddressesNil

`func (o *ComputePhysicalAllOf) SetKvmIpAddressesNil(b bool)`

 SetKvmIpAddressesNil sets the value for KvmIpAddresses to be an explicit nil

### UnsetKvmIpAddresses
`func (o *ComputePhysicalAllOf) UnsetKvmIpAddresses()`

UnsetKvmIpAddresses ensures that no value is present for KvmIpAddresses, not even an explicit nil
### GetKvmServerStateEnabled

`func (o *ComputePhysicalAllOf) GetKvmServerStateEnabled() bool`

GetKvmServerStateEnabled returns the KvmServerStateEnabled field if non-nil, zero value otherwise.

### GetKvmServerStateEnabledOk

`func (o *ComputePhysicalAllOf) GetKvmServerStateEnabledOk() (*bool, bool)`

GetKvmServerStateEnabledOk returns a tuple with the KvmServerStateEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKvmServerStateEnabled

`func (o *ComputePhysicalAllOf) SetKvmServerStateEnabled(v bool)`

SetKvmServerStateEnabled sets KvmServerStateEnabled field to given value.

### HasKvmServerStateEnabled

`func (o *ComputePhysicalAllOf) HasKvmServerStateEnabled() bool`

HasKvmServerStateEnabled returns a boolean if a field has been set.

### GetKvmVendor

`func (o *ComputePhysicalAllOf) GetKvmVendor() string`

GetKvmVendor returns the KvmVendor field if non-nil, zero value otherwise.

### GetKvmVendorOk

`func (o *ComputePhysicalAllOf) GetKvmVendorOk() (*string, bool)`

GetKvmVendorOk returns a tuple with the KvmVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKvmVendor

`func (o *ComputePhysicalAllOf) SetKvmVendor(v string)`

SetKvmVendor sets KvmVendor field to given value.

### HasKvmVendor

`func (o *ComputePhysicalAllOf) HasKvmVendor() bool`

HasKvmVendor returns a boolean if a field has been set.

### GetLifecycle

`func (o *ComputePhysicalAllOf) GetLifecycle() string`

GetLifecycle returns the Lifecycle field if non-nil, zero value otherwise.

### GetLifecycleOk

`func (o *ComputePhysicalAllOf) GetLifecycleOk() (*string, bool)`

GetLifecycleOk returns a tuple with the Lifecycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycle

`func (o *ComputePhysicalAllOf) SetLifecycle(v string)`

SetLifecycle sets Lifecycle field to given value.

### HasLifecycle

`func (o *ComputePhysicalAllOf) HasLifecycle() bool`

HasLifecycle returns a boolean if a field has been set.

### GetManagementMode

`func (o *ComputePhysicalAllOf) GetManagementMode() string`

GetManagementMode returns the ManagementMode field if non-nil, zero value otherwise.

### GetManagementModeOk

`func (o *ComputePhysicalAllOf) GetManagementModeOk() (*string, bool)`

GetManagementModeOk returns a tuple with the ManagementMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementMode

`func (o *ComputePhysicalAllOf) SetManagementMode(v string)`

SetManagementMode sets ManagementMode field to given value.

### HasManagementMode

`func (o *ComputePhysicalAllOf) HasManagementMode() bool`

HasManagementMode returns a boolean if a field has been set.

### GetMemorySpeed

`func (o *ComputePhysicalAllOf) GetMemorySpeed() string`

GetMemorySpeed returns the MemorySpeed field if non-nil, zero value otherwise.

### GetMemorySpeedOk

`func (o *ComputePhysicalAllOf) GetMemorySpeedOk() (*string, bool)`

GetMemorySpeedOk returns a tuple with the MemorySpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemorySpeed

`func (o *ComputePhysicalAllOf) SetMemorySpeed(v string)`

SetMemorySpeed sets MemorySpeed field to given value.

### HasMemorySpeed

`func (o *ComputePhysicalAllOf) HasMemorySpeed() bool`

HasMemorySpeed returns a boolean if a field has been set.

### GetMgmtIpAddress

`func (o *ComputePhysicalAllOf) GetMgmtIpAddress() string`

GetMgmtIpAddress returns the MgmtIpAddress field if non-nil, zero value otherwise.

### GetMgmtIpAddressOk

`func (o *ComputePhysicalAllOf) GetMgmtIpAddressOk() (*string, bool)`

GetMgmtIpAddressOk returns a tuple with the MgmtIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgmtIpAddress

`func (o *ComputePhysicalAllOf) SetMgmtIpAddress(v string)`

SetMgmtIpAddress sets MgmtIpAddress field to given value.

### HasMgmtIpAddress

`func (o *ComputePhysicalAllOf) HasMgmtIpAddress() bool`

HasMgmtIpAddress returns a boolean if a field has been set.

### GetName

`func (o *ComputePhysicalAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ComputePhysicalAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ComputePhysicalAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ComputePhysicalAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumAdaptors

`func (o *ComputePhysicalAllOf) GetNumAdaptors() int64`

GetNumAdaptors returns the NumAdaptors field if non-nil, zero value otherwise.

### GetNumAdaptorsOk

`func (o *ComputePhysicalAllOf) GetNumAdaptorsOk() (*int64, bool)`

GetNumAdaptorsOk returns a tuple with the NumAdaptors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumAdaptors

`func (o *ComputePhysicalAllOf) SetNumAdaptors(v int64)`

SetNumAdaptors sets NumAdaptors field to given value.

### HasNumAdaptors

`func (o *ComputePhysicalAllOf) HasNumAdaptors() bool`

HasNumAdaptors returns a boolean if a field has been set.

### GetNumCpuCores

`func (o *ComputePhysicalAllOf) GetNumCpuCores() int64`

GetNumCpuCores returns the NumCpuCores field if non-nil, zero value otherwise.

### GetNumCpuCoresOk

`func (o *ComputePhysicalAllOf) GetNumCpuCoresOk() (*int64, bool)`

GetNumCpuCoresOk returns a tuple with the NumCpuCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumCpuCores

`func (o *ComputePhysicalAllOf) SetNumCpuCores(v int64)`

SetNumCpuCores sets NumCpuCores field to given value.

### HasNumCpuCores

`func (o *ComputePhysicalAllOf) HasNumCpuCores() bool`

HasNumCpuCores returns a boolean if a field has been set.

### GetNumCpuCoresEnabled

`func (o *ComputePhysicalAllOf) GetNumCpuCoresEnabled() int64`

GetNumCpuCoresEnabled returns the NumCpuCoresEnabled field if non-nil, zero value otherwise.

### GetNumCpuCoresEnabledOk

`func (o *ComputePhysicalAllOf) GetNumCpuCoresEnabledOk() (*int64, bool)`

GetNumCpuCoresEnabledOk returns a tuple with the NumCpuCoresEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumCpuCoresEnabled

`func (o *ComputePhysicalAllOf) SetNumCpuCoresEnabled(v int64)`

SetNumCpuCoresEnabled sets NumCpuCoresEnabled field to given value.

### HasNumCpuCoresEnabled

`func (o *ComputePhysicalAllOf) HasNumCpuCoresEnabled() bool`

HasNumCpuCoresEnabled returns a boolean if a field has been set.

### GetNumCpus

`func (o *ComputePhysicalAllOf) GetNumCpus() int64`

GetNumCpus returns the NumCpus field if non-nil, zero value otherwise.

### GetNumCpusOk

`func (o *ComputePhysicalAllOf) GetNumCpusOk() (*int64, bool)`

GetNumCpusOk returns a tuple with the NumCpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumCpus

`func (o *ComputePhysicalAllOf) SetNumCpus(v int64)`

SetNumCpus sets NumCpus field to given value.

### HasNumCpus

`func (o *ComputePhysicalAllOf) HasNumCpus() bool`

HasNumCpus returns a boolean if a field has been set.

### GetNumEthHostInterfaces

`func (o *ComputePhysicalAllOf) GetNumEthHostInterfaces() int64`

GetNumEthHostInterfaces returns the NumEthHostInterfaces field if non-nil, zero value otherwise.

### GetNumEthHostInterfacesOk

`func (o *ComputePhysicalAllOf) GetNumEthHostInterfacesOk() (*int64, bool)`

GetNumEthHostInterfacesOk returns a tuple with the NumEthHostInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumEthHostInterfaces

`func (o *ComputePhysicalAllOf) SetNumEthHostInterfaces(v int64)`

SetNumEthHostInterfaces sets NumEthHostInterfaces field to given value.

### HasNumEthHostInterfaces

`func (o *ComputePhysicalAllOf) HasNumEthHostInterfaces() bool`

HasNumEthHostInterfaces returns a boolean if a field has been set.

### GetNumFcHostInterfaces

`func (o *ComputePhysicalAllOf) GetNumFcHostInterfaces() int64`

GetNumFcHostInterfaces returns the NumFcHostInterfaces field if non-nil, zero value otherwise.

### GetNumFcHostInterfacesOk

`func (o *ComputePhysicalAllOf) GetNumFcHostInterfacesOk() (*int64, bool)`

GetNumFcHostInterfacesOk returns a tuple with the NumFcHostInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumFcHostInterfaces

`func (o *ComputePhysicalAllOf) SetNumFcHostInterfaces(v int64)`

SetNumFcHostInterfaces sets NumFcHostInterfaces field to given value.

### HasNumFcHostInterfaces

`func (o *ComputePhysicalAllOf) HasNumFcHostInterfaces() bool`

HasNumFcHostInterfaces returns a boolean if a field has been set.

### GetNumThreads

`func (o *ComputePhysicalAllOf) GetNumThreads() int64`

GetNumThreads returns the NumThreads field if non-nil, zero value otherwise.

### GetNumThreadsOk

`func (o *ComputePhysicalAllOf) GetNumThreadsOk() (*int64, bool)`

GetNumThreadsOk returns a tuple with the NumThreads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumThreads

`func (o *ComputePhysicalAllOf) SetNumThreads(v int64)`

SetNumThreads sets NumThreads field to given value.

### HasNumThreads

`func (o *ComputePhysicalAllOf) HasNumThreads() bool`

HasNumThreads returns a boolean if a field has been set.

### GetOperPowerState

`func (o *ComputePhysicalAllOf) GetOperPowerState() string`

GetOperPowerState returns the OperPowerState field if non-nil, zero value otherwise.

### GetOperPowerStateOk

`func (o *ComputePhysicalAllOf) GetOperPowerStateOk() (*string, bool)`

GetOperPowerStateOk returns a tuple with the OperPowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperPowerState

`func (o *ComputePhysicalAllOf) SetOperPowerState(v string)`

SetOperPowerState sets OperPowerState field to given value.

### HasOperPowerState

`func (o *ComputePhysicalAllOf) HasOperPowerState() bool`

HasOperPowerState returns a boolean if a field has been set.

### GetOperReason

`func (o *ComputePhysicalAllOf) GetOperReason() []string`

GetOperReason returns the OperReason field if non-nil, zero value otherwise.

### GetOperReasonOk

`func (o *ComputePhysicalAllOf) GetOperReasonOk() (*[]string, bool)`

GetOperReasonOk returns a tuple with the OperReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperReason

`func (o *ComputePhysicalAllOf) SetOperReason(v []string)`

SetOperReason sets OperReason field to given value.

### HasOperReason

`func (o *ComputePhysicalAllOf) HasOperReason() bool`

HasOperReason returns a boolean if a field has been set.

### SetOperReasonNil

`func (o *ComputePhysicalAllOf) SetOperReasonNil(b bool)`

 SetOperReasonNil sets the value for OperReason to be an explicit nil

### UnsetOperReason
`func (o *ComputePhysicalAllOf) UnsetOperReason()`

UnsetOperReason ensures that no value is present for OperReason, not even an explicit nil
### GetOperState

`func (o *ComputePhysicalAllOf) GetOperState() string`

GetOperState returns the OperState field if non-nil, zero value otherwise.

### GetOperStateOk

`func (o *ComputePhysicalAllOf) GetOperStateOk() (*string, bool)`

GetOperStateOk returns a tuple with the OperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperState

`func (o *ComputePhysicalAllOf) SetOperState(v string)`

SetOperState sets OperState field to given value.

### HasOperState

`func (o *ComputePhysicalAllOf) HasOperState() bool`

HasOperState returns a boolean if a field has been set.

### GetOperability

`func (o *ComputePhysicalAllOf) GetOperability() string`

GetOperability returns the Operability field if non-nil, zero value otherwise.

### GetOperabilityOk

`func (o *ComputePhysicalAllOf) GetOperabilityOk() (*string, bool)`

GetOperabilityOk returns a tuple with the Operability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperability

`func (o *ComputePhysicalAllOf) SetOperability(v string)`

SetOperability sets Operability field to given value.

### HasOperability

`func (o *ComputePhysicalAllOf) HasOperability() bool`

HasOperability returns a boolean if a field has been set.

### GetPlatformType

`func (o *ComputePhysicalAllOf) GetPlatformType() string`

GetPlatformType returns the PlatformType field if non-nil, zero value otherwise.

### GetPlatformTypeOk

`func (o *ComputePhysicalAllOf) GetPlatformTypeOk() (*string, bool)`

GetPlatformTypeOk returns a tuple with the PlatformType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformType

`func (o *ComputePhysicalAllOf) SetPlatformType(v string)`

SetPlatformType sets PlatformType field to given value.

### HasPlatformType

`func (o *ComputePhysicalAllOf) HasPlatformType() bool`

HasPlatformType returns a boolean if a field has been set.

### GetServiceProfile

`func (o *ComputePhysicalAllOf) GetServiceProfile() string`

GetServiceProfile returns the ServiceProfile field if non-nil, zero value otherwise.

### GetServiceProfileOk

`func (o *ComputePhysicalAllOf) GetServiceProfileOk() (*string, bool)`

GetServiceProfileOk returns a tuple with the ServiceProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceProfile

`func (o *ComputePhysicalAllOf) SetServiceProfile(v string)`

SetServiceProfile sets ServiceProfile field to given value.

### HasServiceProfile

`func (o *ComputePhysicalAllOf) HasServiceProfile() bool`

HasServiceProfile returns a boolean if a field has been set.

### GetTotalMemory

`func (o *ComputePhysicalAllOf) GetTotalMemory() int64`

GetTotalMemory returns the TotalMemory field if non-nil, zero value otherwise.

### GetTotalMemoryOk

`func (o *ComputePhysicalAllOf) GetTotalMemoryOk() (*int64, bool)`

GetTotalMemoryOk returns a tuple with the TotalMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMemory

`func (o *ComputePhysicalAllOf) SetTotalMemory(v int64)`

SetTotalMemory sets TotalMemory field to given value.

### HasTotalMemory

`func (o *ComputePhysicalAllOf) HasTotalMemory() bool`

HasTotalMemory returns a boolean if a field has been set.

### GetTunneledKvm

`func (o *ComputePhysicalAllOf) GetTunneledKvm() bool`

GetTunneledKvm returns the TunneledKvm field if non-nil, zero value otherwise.

### GetTunneledKvmOk

`func (o *ComputePhysicalAllOf) GetTunneledKvmOk() (*bool, bool)`

GetTunneledKvmOk returns a tuple with the TunneledKvm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunneledKvm

`func (o *ComputePhysicalAllOf) SetTunneledKvm(v bool)`

SetTunneledKvm sets TunneledKvm field to given value.

### HasTunneledKvm

`func (o *ComputePhysicalAllOf) HasTunneledKvm() bool`

HasTunneledKvm returns a boolean if a field has been set.

### GetUserLabel

`func (o *ComputePhysicalAllOf) GetUserLabel() string`

GetUserLabel returns the UserLabel field if non-nil, zero value otherwise.

### GetUserLabelOk

`func (o *ComputePhysicalAllOf) GetUserLabelOk() (*string, bool)`

GetUserLabelOk returns a tuple with the UserLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLabel

`func (o *ComputePhysicalAllOf) SetUserLabel(v string)`

SetUserLabel sets UserLabel field to given value.

### HasUserLabel

`func (o *ComputePhysicalAllOf) HasUserLabel() bool`

HasUserLabel returns a boolean if a field has been set.

### GetUuid

`func (o *ComputePhysicalAllOf) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ComputePhysicalAllOf) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ComputePhysicalAllOf) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ComputePhysicalAllOf) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetBootCddDevices

`func (o *ComputePhysicalAllOf) GetBootCddDevices() []BootCddDeviceRelationship`

GetBootCddDevices returns the BootCddDevices field if non-nil, zero value otherwise.

### GetBootCddDevicesOk

`func (o *ComputePhysicalAllOf) GetBootCddDevicesOk() (*[]BootCddDeviceRelationship, bool)`

GetBootCddDevicesOk returns a tuple with the BootCddDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootCddDevices

`func (o *ComputePhysicalAllOf) SetBootCddDevices(v []BootCddDeviceRelationship)`

SetBootCddDevices sets BootCddDevices field to given value.

### HasBootCddDevices

`func (o *ComputePhysicalAllOf) HasBootCddDevices() bool`

HasBootCddDevices returns a boolean if a field has been set.

### SetBootCddDevicesNil

`func (o *ComputePhysicalAllOf) SetBootCddDevicesNil(b bool)`

 SetBootCddDevicesNil sets the value for BootCddDevices to be an explicit nil

### UnsetBootCddDevices
`func (o *ComputePhysicalAllOf) UnsetBootCddDevices()`

UnsetBootCddDevices ensures that no value is present for BootCddDevices, not even an explicit nil
### GetBootDeviceBootSecurity

`func (o *ComputePhysicalAllOf) GetBootDeviceBootSecurity() BootDeviceBootSecurityRelationship`

GetBootDeviceBootSecurity returns the BootDeviceBootSecurity field if non-nil, zero value otherwise.

### GetBootDeviceBootSecurityOk

`func (o *ComputePhysicalAllOf) GetBootDeviceBootSecurityOk() (*BootDeviceBootSecurityRelationship, bool)`

GetBootDeviceBootSecurityOk returns a tuple with the BootDeviceBootSecurity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootDeviceBootSecurity

`func (o *ComputePhysicalAllOf) SetBootDeviceBootSecurity(v BootDeviceBootSecurityRelationship)`

SetBootDeviceBootSecurity sets BootDeviceBootSecurity field to given value.

### HasBootDeviceBootSecurity

`func (o *ComputePhysicalAllOf) HasBootDeviceBootSecurity() bool`

HasBootDeviceBootSecurity returns a boolean if a field has been set.

### GetBootHddDevices

`func (o *ComputePhysicalAllOf) GetBootHddDevices() []BootHddDeviceRelationship`

GetBootHddDevices returns the BootHddDevices field if non-nil, zero value otherwise.

### GetBootHddDevicesOk

`func (o *ComputePhysicalAllOf) GetBootHddDevicesOk() (*[]BootHddDeviceRelationship, bool)`

GetBootHddDevicesOk returns a tuple with the BootHddDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootHddDevices

`func (o *ComputePhysicalAllOf) SetBootHddDevices(v []BootHddDeviceRelationship)`

SetBootHddDevices sets BootHddDevices field to given value.

### HasBootHddDevices

`func (o *ComputePhysicalAllOf) HasBootHddDevices() bool`

HasBootHddDevices returns a boolean if a field has been set.

### SetBootHddDevicesNil

`func (o *ComputePhysicalAllOf) SetBootHddDevicesNil(b bool)`

 SetBootHddDevicesNil sets the value for BootHddDevices to be an explicit nil

### UnsetBootHddDevices
`func (o *ComputePhysicalAllOf) UnsetBootHddDevices()`

UnsetBootHddDevices ensures that no value is present for BootHddDevices, not even an explicit nil
### GetBootIscsiDevices

`func (o *ComputePhysicalAllOf) GetBootIscsiDevices() []BootIscsiDeviceRelationship`

GetBootIscsiDevices returns the BootIscsiDevices field if non-nil, zero value otherwise.

### GetBootIscsiDevicesOk

`func (o *ComputePhysicalAllOf) GetBootIscsiDevicesOk() (*[]BootIscsiDeviceRelationship, bool)`

GetBootIscsiDevicesOk returns a tuple with the BootIscsiDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootIscsiDevices

`func (o *ComputePhysicalAllOf) SetBootIscsiDevices(v []BootIscsiDeviceRelationship)`

SetBootIscsiDevices sets BootIscsiDevices field to given value.

### HasBootIscsiDevices

`func (o *ComputePhysicalAllOf) HasBootIscsiDevices() bool`

HasBootIscsiDevices returns a boolean if a field has been set.

### SetBootIscsiDevicesNil

`func (o *ComputePhysicalAllOf) SetBootIscsiDevicesNil(b bool)`

 SetBootIscsiDevicesNil sets the value for BootIscsiDevices to be an explicit nil

### UnsetBootIscsiDevices
`func (o *ComputePhysicalAllOf) UnsetBootIscsiDevices()`

UnsetBootIscsiDevices ensures that no value is present for BootIscsiDevices, not even an explicit nil
### GetBootNvmeDevices

`func (o *ComputePhysicalAllOf) GetBootNvmeDevices() []BootNvmeDeviceRelationship`

GetBootNvmeDevices returns the BootNvmeDevices field if non-nil, zero value otherwise.

### GetBootNvmeDevicesOk

`func (o *ComputePhysicalAllOf) GetBootNvmeDevicesOk() (*[]BootNvmeDeviceRelationship, bool)`

GetBootNvmeDevicesOk returns a tuple with the BootNvmeDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootNvmeDevices

`func (o *ComputePhysicalAllOf) SetBootNvmeDevices(v []BootNvmeDeviceRelationship)`

SetBootNvmeDevices sets BootNvmeDevices field to given value.

### HasBootNvmeDevices

`func (o *ComputePhysicalAllOf) HasBootNvmeDevices() bool`

HasBootNvmeDevices returns a boolean if a field has been set.

### SetBootNvmeDevicesNil

`func (o *ComputePhysicalAllOf) SetBootNvmeDevicesNil(b bool)`

 SetBootNvmeDevicesNil sets the value for BootNvmeDevices to be an explicit nil

### UnsetBootNvmeDevices
`func (o *ComputePhysicalAllOf) UnsetBootNvmeDevices()`

UnsetBootNvmeDevices ensures that no value is present for BootNvmeDevices, not even an explicit nil
### GetBootPchStorageDevices

`func (o *ComputePhysicalAllOf) GetBootPchStorageDevices() []BootPchStorageDeviceRelationship`

GetBootPchStorageDevices returns the BootPchStorageDevices field if non-nil, zero value otherwise.

### GetBootPchStorageDevicesOk

`func (o *ComputePhysicalAllOf) GetBootPchStorageDevicesOk() (*[]BootPchStorageDeviceRelationship, bool)`

GetBootPchStorageDevicesOk returns a tuple with the BootPchStorageDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootPchStorageDevices

`func (o *ComputePhysicalAllOf) SetBootPchStorageDevices(v []BootPchStorageDeviceRelationship)`

SetBootPchStorageDevices sets BootPchStorageDevices field to given value.

### HasBootPchStorageDevices

`func (o *ComputePhysicalAllOf) HasBootPchStorageDevices() bool`

HasBootPchStorageDevices returns a boolean if a field has been set.

### SetBootPchStorageDevicesNil

`func (o *ComputePhysicalAllOf) SetBootPchStorageDevicesNil(b bool)`

 SetBootPchStorageDevicesNil sets the value for BootPchStorageDevices to be an explicit nil

### UnsetBootPchStorageDevices
`func (o *ComputePhysicalAllOf) UnsetBootPchStorageDevices()`

UnsetBootPchStorageDevices ensures that no value is present for BootPchStorageDevices, not even an explicit nil
### GetBootPxeDevices

`func (o *ComputePhysicalAllOf) GetBootPxeDevices() []BootPxeDeviceRelationship`

GetBootPxeDevices returns the BootPxeDevices field if non-nil, zero value otherwise.

### GetBootPxeDevicesOk

`func (o *ComputePhysicalAllOf) GetBootPxeDevicesOk() (*[]BootPxeDeviceRelationship, bool)`

GetBootPxeDevicesOk returns a tuple with the BootPxeDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootPxeDevices

`func (o *ComputePhysicalAllOf) SetBootPxeDevices(v []BootPxeDeviceRelationship)`

SetBootPxeDevices sets BootPxeDevices field to given value.

### HasBootPxeDevices

`func (o *ComputePhysicalAllOf) HasBootPxeDevices() bool`

HasBootPxeDevices returns a boolean if a field has been set.

### SetBootPxeDevicesNil

`func (o *ComputePhysicalAllOf) SetBootPxeDevicesNil(b bool)`

 SetBootPxeDevicesNil sets the value for BootPxeDevices to be an explicit nil

### UnsetBootPxeDevices
`func (o *ComputePhysicalAllOf) UnsetBootPxeDevices()`

UnsetBootPxeDevices ensures that no value is present for BootPxeDevices, not even an explicit nil
### GetBootSanDevices

`func (o *ComputePhysicalAllOf) GetBootSanDevices() []BootSanDeviceRelationship`

GetBootSanDevices returns the BootSanDevices field if non-nil, zero value otherwise.

### GetBootSanDevicesOk

`func (o *ComputePhysicalAllOf) GetBootSanDevicesOk() (*[]BootSanDeviceRelationship, bool)`

GetBootSanDevicesOk returns a tuple with the BootSanDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootSanDevices

`func (o *ComputePhysicalAllOf) SetBootSanDevices(v []BootSanDeviceRelationship)`

SetBootSanDevices sets BootSanDevices field to given value.

### HasBootSanDevices

`func (o *ComputePhysicalAllOf) HasBootSanDevices() bool`

HasBootSanDevices returns a boolean if a field has been set.

### SetBootSanDevicesNil

`func (o *ComputePhysicalAllOf) SetBootSanDevicesNil(b bool)`

 SetBootSanDevicesNil sets the value for BootSanDevices to be an explicit nil

### UnsetBootSanDevices
`func (o *ComputePhysicalAllOf) UnsetBootSanDevices()`

UnsetBootSanDevices ensures that no value is present for BootSanDevices, not even an explicit nil
### GetBootSdDevices

`func (o *ComputePhysicalAllOf) GetBootSdDevices() []BootSdDeviceRelationship`

GetBootSdDevices returns the BootSdDevices field if non-nil, zero value otherwise.

### GetBootSdDevicesOk

`func (o *ComputePhysicalAllOf) GetBootSdDevicesOk() (*[]BootSdDeviceRelationship, bool)`

GetBootSdDevicesOk returns a tuple with the BootSdDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootSdDevices

`func (o *ComputePhysicalAllOf) SetBootSdDevices(v []BootSdDeviceRelationship)`

SetBootSdDevices sets BootSdDevices field to given value.

### HasBootSdDevices

`func (o *ComputePhysicalAllOf) HasBootSdDevices() bool`

HasBootSdDevices returns a boolean if a field has been set.

### SetBootSdDevicesNil

`func (o *ComputePhysicalAllOf) SetBootSdDevicesNil(b bool)`

 SetBootSdDevicesNil sets the value for BootSdDevices to be an explicit nil

### UnsetBootSdDevices
`func (o *ComputePhysicalAllOf) UnsetBootSdDevices()`

UnsetBootSdDevices ensures that no value is present for BootSdDevices, not even an explicit nil
### GetBootUefiShellDevices

`func (o *ComputePhysicalAllOf) GetBootUefiShellDevices() []BootUefiShellDeviceRelationship`

GetBootUefiShellDevices returns the BootUefiShellDevices field if non-nil, zero value otherwise.

### GetBootUefiShellDevicesOk

`func (o *ComputePhysicalAllOf) GetBootUefiShellDevicesOk() (*[]BootUefiShellDeviceRelationship, bool)`

GetBootUefiShellDevicesOk returns a tuple with the BootUefiShellDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootUefiShellDevices

`func (o *ComputePhysicalAllOf) SetBootUefiShellDevices(v []BootUefiShellDeviceRelationship)`

SetBootUefiShellDevices sets BootUefiShellDevices field to given value.

### HasBootUefiShellDevices

`func (o *ComputePhysicalAllOf) HasBootUefiShellDevices() bool`

HasBootUefiShellDevices returns a boolean if a field has been set.

### SetBootUefiShellDevicesNil

`func (o *ComputePhysicalAllOf) SetBootUefiShellDevicesNil(b bool)`

 SetBootUefiShellDevicesNil sets the value for BootUefiShellDevices to be an explicit nil

### UnsetBootUefiShellDevices
`func (o *ComputePhysicalAllOf) UnsetBootUefiShellDevices()`

UnsetBootUefiShellDevices ensures that no value is present for BootUefiShellDevices, not even an explicit nil
### GetBootUsbDevices

`func (o *ComputePhysicalAllOf) GetBootUsbDevices() []BootUsbDeviceRelationship`

GetBootUsbDevices returns the BootUsbDevices field if non-nil, zero value otherwise.

### GetBootUsbDevicesOk

`func (o *ComputePhysicalAllOf) GetBootUsbDevicesOk() (*[]BootUsbDeviceRelationship, bool)`

GetBootUsbDevicesOk returns a tuple with the BootUsbDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootUsbDevices

`func (o *ComputePhysicalAllOf) SetBootUsbDevices(v []BootUsbDeviceRelationship)`

SetBootUsbDevices sets BootUsbDevices field to given value.

### HasBootUsbDevices

`func (o *ComputePhysicalAllOf) HasBootUsbDevices() bool`

HasBootUsbDevices returns a boolean if a field has been set.

### SetBootUsbDevicesNil

`func (o *ComputePhysicalAllOf) SetBootUsbDevicesNil(b bool)`

 SetBootUsbDevicesNil sets the value for BootUsbDevices to be an explicit nil

### UnsetBootUsbDevices
`func (o *ComputePhysicalAllOf) UnsetBootUsbDevices()`

UnsetBootUsbDevices ensures that no value is present for BootUsbDevices, not even an explicit nil
### GetBootVmediaDevices

`func (o *ComputePhysicalAllOf) GetBootVmediaDevices() []BootVmediaDeviceRelationship`

GetBootVmediaDevices returns the BootVmediaDevices field if non-nil, zero value otherwise.

### GetBootVmediaDevicesOk

`func (o *ComputePhysicalAllOf) GetBootVmediaDevicesOk() (*[]BootVmediaDeviceRelationship, bool)`

GetBootVmediaDevicesOk returns a tuple with the BootVmediaDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootVmediaDevices

`func (o *ComputePhysicalAllOf) SetBootVmediaDevices(v []BootVmediaDeviceRelationship)`

SetBootVmediaDevices sets BootVmediaDevices field to given value.

### HasBootVmediaDevices

`func (o *ComputePhysicalAllOf) HasBootVmediaDevices() bool`

HasBootVmediaDevices returns a boolean if a field has been set.

### SetBootVmediaDevicesNil

`func (o *ComputePhysicalAllOf) SetBootVmediaDevicesNil(b bool)`

 SetBootVmediaDevicesNil sets the value for BootVmediaDevices to be an explicit nil

### UnsetBootVmediaDevices
`func (o *ComputePhysicalAllOf) UnsetBootVmediaDevices()`

UnsetBootVmediaDevices ensures that no value is present for BootVmediaDevices, not even an explicit nil
### GetMgmtIdentity

`func (o *ComputePhysicalAllOf) GetMgmtIdentity() EquipmentPhysicalIdentityRelationship`

GetMgmtIdentity returns the MgmtIdentity field if non-nil, zero value otherwise.

### GetMgmtIdentityOk

`func (o *ComputePhysicalAllOf) GetMgmtIdentityOk() (*EquipmentPhysicalIdentityRelationship, bool)`

GetMgmtIdentityOk returns a tuple with the MgmtIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgmtIdentity

`func (o *ComputePhysicalAllOf) SetMgmtIdentity(v EquipmentPhysicalIdentityRelationship)`

SetMgmtIdentity sets MgmtIdentity field to given value.

### HasMgmtIdentity

`func (o *ComputePhysicalAllOf) HasMgmtIdentity() bool`

HasMgmtIdentity returns a boolean if a field has been set.

### GetVmedia

`func (o *ComputePhysicalAllOf) GetVmedia() ComputeVmediaRelationship`

GetVmedia returns the Vmedia field if non-nil, zero value otherwise.

### GetVmediaOk

`func (o *ComputePhysicalAllOf) GetVmediaOk() (*ComputeVmediaRelationship, bool)`

GetVmediaOk returns a tuple with the Vmedia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmedia

`func (o *ComputePhysicalAllOf) SetVmedia(v ComputeVmediaRelationship)`

SetVmedia sets Vmedia field to given value.

### HasVmedia

`func (o *ComputePhysicalAllOf) HasVmedia() bool`

HasVmedia returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



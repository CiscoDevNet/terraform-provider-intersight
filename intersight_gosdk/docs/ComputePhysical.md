# ComputePhysical

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
**BmcInventoryReady** | Pointer to **bool** | The BMC inventory readiness status of the server. | [optional] 
**BootLastState** | Pointer to **string** | The boot progress state of a rack or blade server. * &#x60;None&#x60; - The server is powered off. * &#x60;OSBootStarted&#x60; - The operating system boot process has started. * &#x60;OSRunning&#x60; - The operating system boot process has started and running. * &#x60;OEM&#x60; - The server is in an OEM-defined startup state. | [optional] [readonly] [default to "None"]
**BootOemLastState** | Pointer to **string** | The last known OEM-defined startup state of a rack or blade server. * &#x60;None&#x60; - The server has not yet completed OEM initialization. * &#x60;PlatformInitializing&#x60; - The server is initializing after being powered on. | [optional] [readonly] [default to "None"]
**CoolingMode** | Pointer to **string** | Cooling mode representation of the server. * &#x60;Air&#x60; - Cooling mode of the device is set to Air. * &#x60;Immersion&#x60; - Cooling mode of the device is set to Immersion. | [optional] [readonly] [default to "Air"]
**CpuCapacity** | Pointer to **float32** | Total processing capacity of the server. | [optional] [readonly] 
**DimmBlocklisting** | Pointer to **string** | The DIMM Blocklisting setting configured on the server. * &#x60;Unknown&#x60; - The configured state is unknown. * &#x60;Enabled&#x60; - The configured state is enabled. * &#x60;Disabled&#x60; - The configured state is disabled. | [optional] [readonly] [default to "Unknown"]
**FaultSummary** | Pointer to **int64** | The fault summary for the server. | [optional] 
**FrontPanelLockState** | Pointer to **string** | The actual front panel state of the server. * &#x60;None&#x60; - Front Panel of the server is set to None state. It is required so that the next frontPanelLockState operation can be triggered. * &#x60;Lock&#x60; - Front Panel of the server is set to Locked state. * &#x60;Unlock&#x60; - Front Panel of the server is set to Unlocked state. | [optional] [default to "None"]
**HardwareUuid** | Pointer to **string** | The universally unique hardware identity of the server provided by the manufacturer. | [optional] 
**KvmIpAddresses** | Pointer to [**[]ComputeIpAddress**](ComputeIpAddress.md) |  | [optional] 
**KvmServerStateEnabled** | Pointer to **bool** | The KVM server state of the server. | [optional] [readonly] 
**KvmVendor** | Pointer to **string** | The KVM Vendor for the server. | [optional] [readonly] 
**LastPowerStateChangedTime** | Pointer to **string** | The Last host power state changed time of the server. | [optional] [readonly] 
**Lifecycle** | Pointer to **string** | The lifecycle state of the server. This will map to the discovery lifecycle as represented in the server Identity object. * &#x60;None&#x60; - Default state of an equipment. This should be an initial state when no state is defined for an equipment. * &#x60;Active&#x60; - Default Lifecycle State for a physical entity. * &#x60;Decommissioned&#x60; - Decommission Lifecycle state. * &#x60;DiscoveryInProgress&#x60; - DiscoveryInProgress Lifecycle state. * &#x60;DiscoveryFailed&#x60; - DiscoveryFailed Lifecycle state. * &#x60;FirmwareUpgradeInProgress&#x60; - Firmware upgrade is in progress on given physical entity. * &#x60;SecureEraseInProgress&#x60; - Secure Erase is in progress on given physical entity. * &#x60;ScrubInProgress&#x60; - Scrub is in progress on given physical entity. * &#x60;BladeMigrationInProgress&#x60; - Server slot migration is in progress on given physical entity. * &#x60;SlotMismatch&#x60; - The blade server is detected in a different chassis/slot than it was previously. * &#x60;Removed&#x60; - The blade server has been removed from its discovered slot, and not detected anywhere else. Blade inventory can be cleaned up by performing a software remove operation on the physically removed blade. * &#x60;Moved&#x60; - The blade server has been moved from its discovered location to a new location. Blade inventory can be updated by performing a rediscover operation on the moved blade. * &#x60;Replaced&#x60; - The blade server has been removed from its discovered location and another blade has been inserted in that location. Blade inventory can be cleaned up and updated by doing a software remove operation on the physically removed blade. * &#x60;MovedAndReplaced&#x60; - The blade server has been moved from its discovered location to a new location and another blade has been inserted into the old discovered location. Blade inventory can be updated by performing a rediscover operation on the moved blade. | [optional] [readonly] [default to "None"]
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
**BootDeviceBootSecurity** | Pointer to [**NullableBootDeviceBootSecurityRelationship**](BootDeviceBootSecurityRelationship.md) |  | [optional] 
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
**MgmtIdentity** | Pointer to [**NullableEquipmentPhysicalIdentityRelationship**](EquipmentPhysicalIdentityRelationship.md) |  | [optional] 
**Vmedia** | Pointer to [**NullableComputeVmediaRelationship**](ComputeVmediaRelationship.md) |  | [optional] 

## Methods

### NewComputePhysical

`func NewComputePhysical(classId string, objectType string, ) *ComputePhysical`

NewComputePhysical instantiates a new ComputePhysical object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputePhysicalWithDefaults

`func NewComputePhysicalWithDefaults() *ComputePhysical`

NewComputePhysicalWithDefaults instantiates a new ComputePhysical object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ComputePhysical) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ComputePhysical) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ComputePhysical) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ComputePhysical) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ComputePhysical) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ComputePhysical) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdminPowerState

`func (o *ComputePhysical) GetAdminPowerState() string`

GetAdminPowerState returns the AdminPowerState field if non-nil, zero value otherwise.

### GetAdminPowerStateOk

`func (o *ComputePhysical) GetAdminPowerStateOk() (*string, bool)`

GetAdminPowerStateOk returns a tuple with the AdminPowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminPowerState

`func (o *ComputePhysical) SetAdminPowerState(v string)`

SetAdminPowerState sets AdminPowerState field to given value.

### HasAdminPowerState

`func (o *ComputePhysical) HasAdminPowerState() bool`

HasAdminPowerState returns a boolean if a field has been set.

### GetAlarmSummary

`func (o *ComputePhysical) GetAlarmSummary() ComputeAlarmSummary`

GetAlarmSummary returns the AlarmSummary field if non-nil, zero value otherwise.

### GetAlarmSummaryOk

`func (o *ComputePhysical) GetAlarmSummaryOk() (*ComputeAlarmSummary, bool)`

GetAlarmSummaryOk returns a tuple with the AlarmSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmSummary

`func (o *ComputePhysical) SetAlarmSummary(v ComputeAlarmSummary)`

SetAlarmSummary sets AlarmSummary field to given value.

### HasAlarmSummary

`func (o *ComputePhysical) HasAlarmSummary() bool`

HasAlarmSummary returns a boolean if a field has been set.

### SetAlarmSummaryNil

`func (o *ComputePhysical) SetAlarmSummaryNil(b bool)`

 SetAlarmSummaryNil sets the value for AlarmSummary to be an explicit nil

### UnsetAlarmSummary
`func (o *ComputePhysical) UnsetAlarmSummary()`

UnsetAlarmSummary ensures that no value is present for AlarmSummary, not even an explicit nil
### GetAssetTag

`func (o *ComputePhysical) GetAssetTag() string`

GetAssetTag returns the AssetTag field if non-nil, zero value otherwise.

### GetAssetTagOk

`func (o *ComputePhysical) GetAssetTagOk() (*string, bool)`

GetAssetTagOk returns a tuple with the AssetTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetTag

`func (o *ComputePhysical) SetAssetTag(v string)`

SetAssetTag sets AssetTag field to given value.

### HasAssetTag

`func (o *ComputePhysical) HasAssetTag() bool`

HasAssetTag returns a boolean if a field has been set.

### GetAvailableMemory

`func (o *ComputePhysical) GetAvailableMemory() int64`

GetAvailableMemory returns the AvailableMemory field if non-nil, zero value otherwise.

### GetAvailableMemoryOk

`func (o *ComputePhysical) GetAvailableMemoryOk() (*int64, bool)`

GetAvailableMemoryOk returns a tuple with the AvailableMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableMemory

`func (o *ComputePhysical) SetAvailableMemory(v int64)`

SetAvailableMemory sets AvailableMemory field to given value.

### HasAvailableMemory

`func (o *ComputePhysical) HasAvailableMemory() bool`

HasAvailableMemory returns a boolean if a field has been set.

### GetBiosPostComplete

`func (o *ComputePhysical) GetBiosPostComplete() bool`

GetBiosPostComplete returns the BiosPostComplete field if non-nil, zero value otherwise.

### GetBiosPostCompleteOk

`func (o *ComputePhysical) GetBiosPostCompleteOk() (*bool, bool)`

GetBiosPostCompleteOk returns a tuple with the BiosPostComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiosPostComplete

`func (o *ComputePhysical) SetBiosPostComplete(v bool)`

SetBiosPostComplete sets BiosPostComplete field to given value.

### HasBiosPostComplete

`func (o *ComputePhysical) HasBiosPostComplete() bool`

HasBiosPostComplete returns a boolean if a field has been set.

### GetBmcInventoryReady

`func (o *ComputePhysical) GetBmcInventoryReady() bool`

GetBmcInventoryReady returns the BmcInventoryReady field if non-nil, zero value otherwise.

### GetBmcInventoryReadyOk

`func (o *ComputePhysical) GetBmcInventoryReadyOk() (*bool, bool)`

GetBmcInventoryReadyOk returns a tuple with the BmcInventoryReady field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBmcInventoryReady

`func (o *ComputePhysical) SetBmcInventoryReady(v bool)`

SetBmcInventoryReady sets BmcInventoryReady field to given value.

### HasBmcInventoryReady

`func (o *ComputePhysical) HasBmcInventoryReady() bool`

HasBmcInventoryReady returns a boolean if a field has been set.

### GetBootLastState

`func (o *ComputePhysical) GetBootLastState() string`

GetBootLastState returns the BootLastState field if non-nil, zero value otherwise.

### GetBootLastStateOk

`func (o *ComputePhysical) GetBootLastStateOk() (*string, bool)`

GetBootLastStateOk returns a tuple with the BootLastState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootLastState

`func (o *ComputePhysical) SetBootLastState(v string)`

SetBootLastState sets BootLastState field to given value.

### HasBootLastState

`func (o *ComputePhysical) HasBootLastState() bool`

HasBootLastState returns a boolean if a field has been set.

### GetBootOemLastState

`func (o *ComputePhysical) GetBootOemLastState() string`

GetBootOemLastState returns the BootOemLastState field if non-nil, zero value otherwise.

### GetBootOemLastStateOk

`func (o *ComputePhysical) GetBootOemLastStateOk() (*string, bool)`

GetBootOemLastStateOk returns a tuple with the BootOemLastState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootOemLastState

`func (o *ComputePhysical) SetBootOemLastState(v string)`

SetBootOemLastState sets BootOemLastState field to given value.

### HasBootOemLastState

`func (o *ComputePhysical) HasBootOemLastState() bool`

HasBootOemLastState returns a boolean if a field has been set.

### GetCoolingMode

`func (o *ComputePhysical) GetCoolingMode() string`

GetCoolingMode returns the CoolingMode field if non-nil, zero value otherwise.

### GetCoolingModeOk

`func (o *ComputePhysical) GetCoolingModeOk() (*string, bool)`

GetCoolingModeOk returns a tuple with the CoolingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoolingMode

`func (o *ComputePhysical) SetCoolingMode(v string)`

SetCoolingMode sets CoolingMode field to given value.

### HasCoolingMode

`func (o *ComputePhysical) HasCoolingMode() bool`

HasCoolingMode returns a boolean if a field has been set.

### GetCpuCapacity

`func (o *ComputePhysical) GetCpuCapacity() float32`

GetCpuCapacity returns the CpuCapacity field if non-nil, zero value otherwise.

### GetCpuCapacityOk

`func (o *ComputePhysical) GetCpuCapacityOk() (*float32, bool)`

GetCpuCapacityOk returns a tuple with the CpuCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCapacity

`func (o *ComputePhysical) SetCpuCapacity(v float32)`

SetCpuCapacity sets CpuCapacity field to given value.

### HasCpuCapacity

`func (o *ComputePhysical) HasCpuCapacity() bool`

HasCpuCapacity returns a boolean if a field has been set.

### GetDimmBlocklisting

`func (o *ComputePhysical) GetDimmBlocklisting() string`

GetDimmBlocklisting returns the DimmBlocklisting field if non-nil, zero value otherwise.

### GetDimmBlocklistingOk

`func (o *ComputePhysical) GetDimmBlocklistingOk() (*string, bool)`

GetDimmBlocklistingOk returns a tuple with the DimmBlocklisting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimmBlocklisting

`func (o *ComputePhysical) SetDimmBlocklisting(v string)`

SetDimmBlocklisting sets DimmBlocklisting field to given value.

### HasDimmBlocklisting

`func (o *ComputePhysical) HasDimmBlocklisting() bool`

HasDimmBlocklisting returns a boolean if a field has been set.

### GetFaultSummary

`func (o *ComputePhysical) GetFaultSummary() int64`

GetFaultSummary returns the FaultSummary field if non-nil, zero value otherwise.

### GetFaultSummaryOk

`func (o *ComputePhysical) GetFaultSummaryOk() (*int64, bool)`

GetFaultSummaryOk returns a tuple with the FaultSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaultSummary

`func (o *ComputePhysical) SetFaultSummary(v int64)`

SetFaultSummary sets FaultSummary field to given value.

### HasFaultSummary

`func (o *ComputePhysical) HasFaultSummary() bool`

HasFaultSummary returns a boolean if a field has been set.

### GetFrontPanelLockState

`func (o *ComputePhysical) GetFrontPanelLockState() string`

GetFrontPanelLockState returns the FrontPanelLockState field if non-nil, zero value otherwise.

### GetFrontPanelLockStateOk

`func (o *ComputePhysical) GetFrontPanelLockStateOk() (*string, bool)`

GetFrontPanelLockStateOk returns a tuple with the FrontPanelLockState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrontPanelLockState

`func (o *ComputePhysical) SetFrontPanelLockState(v string)`

SetFrontPanelLockState sets FrontPanelLockState field to given value.

### HasFrontPanelLockState

`func (o *ComputePhysical) HasFrontPanelLockState() bool`

HasFrontPanelLockState returns a boolean if a field has been set.

### GetHardwareUuid

`func (o *ComputePhysical) GetHardwareUuid() string`

GetHardwareUuid returns the HardwareUuid field if non-nil, zero value otherwise.

### GetHardwareUuidOk

`func (o *ComputePhysical) GetHardwareUuidOk() (*string, bool)`

GetHardwareUuidOk returns a tuple with the HardwareUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareUuid

`func (o *ComputePhysical) SetHardwareUuid(v string)`

SetHardwareUuid sets HardwareUuid field to given value.

### HasHardwareUuid

`func (o *ComputePhysical) HasHardwareUuid() bool`

HasHardwareUuid returns a boolean if a field has been set.

### GetKvmIpAddresses

`func (o *ComputePhysical) GetKvmIpAddresses() []ComputeIpAddress`

GetKvmIpAddresses returns the KvmIpAddresses field if non-nil, zero value otherwise.

### GetKvmIpAddressesOk

`func (o *ComputePhysical) GetKvmIpAddressesOk() (*[]ComputeIpAddress, bool)`

GetKvmIpAddressesOk returns a tuple with the KvmIpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKvmIpAddresses

`func (o *ComputePhysical) SetKvmIpAddresses(v []ComputeIpAddress)`

SetKvmIpAddresses sets KvmIpAddresses field to given value.

### HasKvmIpAddresses

`func (o *ComputePhysical) HasKvmIpAddresses() bool`

HasKvmIpAddresses returns a boolean if a field has been set.

### SetKvmIpAddressesNil

`func (o *ComputePhysical) SetKvmIpAddressesNil(b bool)`

 SetKvmIpAddressesNil sets the value for KvmIpAddresses to be an explicit nil

### UnsetKvmIpAddresses
`func (o *ComputePhysical) UnsetKvmIpAddresses()`

UnsetKvmIpAddresses ensures that no value is present for KvmIpAddresses, not even an explicit nil
### GetKvmServerStateEnabled

`func (o *ComputePhysical) GetKvmServerStateEnabled() bool`

GetKvmServerStateEnabled returns the KvmServerStateEnabled field if non-nil, zero value otherwise.

### GetKvmServerStateEnabledOk

`func (o *ComputePhysical) GetKvmServerStateEnabledOk() (*bool, bool)`

GetKvmServerStateEnabledOk returns a tuple with the KvmServerStateEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKvmServerStateEnabled

`func (o *ComputePhysical) SetKvmServerStateEnabled(v bool)`

SetKvmServerStateEnabled sets KvmServerStateEnabled field to given value.

### HasKvmServerStateEnabled

`func (o *ComputePhysical) HasKvmServerStateEnabled() bool`

HasKvmServerStateEnabled returns a boolean if a field has been set.

### GetKvmVendor

`func (o *ComputePhysical) GetKvmVendor() string`

GetKvmVendor returns the KvmVendor field if non-nil, zero value otherwise.

### GetKvmVendorOk

`func (o *ComputePhysical) GetKvmVendorOk() (*string, bool)`

GetKvmVendorOk returns a tuple with the KvmVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKvmVendor

`func (o *ComputePhysical) SetKvmVendor(v string)`

SetKvmVendor sets KvmVendor field to given value.

### HasKvmVendor

`func (o *ComputePhysical) HasKvmVendor() bool`

HasKvmVendor returns a boolean if a field has been set.

### GetLastPowerStateChangedTime

`func (o *ComputePhysical) GetLastPowerStateChangedTime() string`

GetLastPowerStateChangedTime returns the LastPowerStateChangedTime field if non-nil, zero value otherwise.

### GetLastPowerStateChangedTimeOk

`func (o *ComputePhysical) GetLastPowerStateChangedTimeOk() (*string, bool)`

GetLastPowerStateChangedTimeOk returns a tuple with the LastPowerStateChangedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPowerStateChangedTime

`func (o *ComputePhysical) SetLastPowerStateChangedTime(v string)`

SetLastPowerStateChangedTime sets LastPowerStateChangedTime field to given value.

### HasLastPowerStateChangedTime

`func (o *ComputePhysical) HasLastPowerStateChangedTime() bool`

HasLastPowerStateChangedTime returns a boolean if a field has been set.

### GetLifecycle

`func (o *ComputePhysical) GetLifecycle() string`

GetLifecycle returns the Lifecycle field if non-nil, zero value otherwise.

### GetLifecycleOk

`func (o *ComputePhysical) GetLifecycleOk() (*string, bool)`

GetLifecycleOk returns a tuple with the Lifecycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycle

`func (o *ComputePhysical) SetLifecycle(v string)`

SetLifecycle sets Lifecycle field to given value.

### HasLifecycle

`func (o *ComputePhysical) HasLifecycle() bool`

HasLifecycle returns a boolean if a field has been set.

### GetManagementMode

`func (o *ComputePhysical) GetManagementMode() string`

GetManagementMode returns the ManagementMode field if non-nil, zero value otherwise.

### GetManagementModeOk

`func (o *ComputePhysical) GetManagementModeOk() (*string, bool)`

GetManagementModeOk returns a tuple with the ManagementMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementMode

`func (o *ComputePhysical) SetManagementMode(v string)`

SetManagementMode sets ManagementMode field to given value.

### HasManagementMode

`func (o *ComputePhysical) HasManagementMode() bool`

HasManagementMode returns a boolean if a field has been set.

### GetMemorySpeed

`func (o *ComputePhysical) GetMemorySpeed() string`

GetMemorySpeed returns the MemorySpeed field if non-nil, zero value otherwise.

### GetMemorySpeedOk

`func (o *ComputePhysical) GetMemorySpeedOk() (*string, bool)`

GetMemorySpeedOk returns a tuple with the MemorySpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemorySpeed

`func (o *ComputePhysical) SetMemorySpeed(v string)`

SetMemorySpeed sets MemorySpeed field to given value.

### HasMemorySpeed

`func (o *ComputePhysical) HasMemorySpeed() bool`

HasMemorySpeed returns a boolean if a field has been set.

### GetMgmtIpAddress

`func (o *ComputePhysical) GetMgmtIpAddress() string`

GetMgmtIpAddress returns the MgmtIpAddress field if non-nil, zero value otherwise.

### GetMgmtIpAddressOk

`func (o *ComputePhysical) GetMgmtIpAddressOk() (*string, bool)`

GetMgmtIpAddressOk returns a tuple with the MgmtIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgmtIpAddress

`func (o *ComputePhysical) SetMgmtIpAddress(v string)`

SetMgmtIpAddress sets MgmtIpAddress field to given value.

### HasMgmtIpAddress

`func (o *ComputePhysical) HasMgmtIpAddress() bool`

HasMgmtIpAddress returns a boolean if a field has been set.

### GetName

`func (o *ComputePhysical) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ComputePhysical) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ComputePhysical) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ComputePhysical) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumAdaptors

`func (o *ComputePhysical) GetNumAdaptors() int64`

GetNumAdaptors returns the NumAdaptors field if non-nil, zero value otherwise.

### GetNumAdaptorsOk

`func (o *ComputePhysical) GetNumAdaptorsOk() (*int64, bool)`

GetNumAdaptorsOk returns a tuple with the NumAdaptors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumAdaptors

`func (o *ComputePhysical) SetNumAdaptors(v int64)`

SetNumAdaptors sets NumAdaptors field to given value.

### HasNumAdaptors

`func (o *ComputePhysical) HasNumAdaptors() bool`

HasNumAdaptors returns a boolean if a field has been set.

### GetNumCpuCores

`func (o *ComputePhysical) GetNumCpuCores() int64`

GetNumCpuCores returns the NumCpuCores field if non-nil, zero value otherwise.

### GetNumCpuCoresOk

`func (o *ComputePhysical) GetNumCpuCoresOk() (*int64, bool)`

GetNumCpuCoresOk returns a tuple with the NumCpuCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumCpuCores

`func (o *ComputePhysical) SetNumCpuCores(v int64)`

SetNumCpuCores sets NumCpuCores field to given value.

### HasNumCpuCores

`func (o *ComputePhysical) HasNumCpuCores() bool`

HasNumCpuCores returns a boolean if a field has been set.

### GetNumCpuCoresEnabled

`func (o *ComputePhysical) GetNumCpuCoresEnabled() int64`

GetNumCpuCoresEnabled returns the NumCpuCoresEnabled field if non-nil, zero value otherwise.

### GetNumCpuCoresEnabledOk

`func (o *ComputePhysical) GetNumCpuCoresEnabledOk() (*int64, bool)`

GetNumCpuCoresEnabledOk returns a tuple with the NumCpuCoresEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumCpuCoresEnabled

`func (o *ComputePhysical) SetNumCpuCoresEnabled(v int64)`

SetNumCpuCoresEnabled sets NumCpuCoresEnabled field to given value.

### HasNumCpuCoresEnabled

`func (o *ComputePhysical) HasNumCpuCoresEnabled() bool`

HasNumCpuCoresEnabled returns a boolean if a field has been set.

### GetNumCpus

`func (o *ComputePhysical) GetNumCpus() int64`

GetNumCpus returns the NumCpus field if non-nil, zero value otherwise.

### GetNumCpusOk

`func (o *ComputePhysical) GetNumCpusOk() (*int64, bool)`

GetNumCpusOk returns a tuple with the NumCpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumCpus

`func (o *ComputePhysical) SetNumCpus(v int64)`

SetNumCpus sets NumCpus field to given value.

### HasNumCpus

`func (o *ComputePhysical) HasNumCpus() bool`

HasNumCpus returns a boolean if a field has been set.

### GetNumEthHostInterfaces

`func (o *ComputePhysical) GetNumEthHostInterfaces() int64`

GetNumEthHostInterfaces returns the NumEthHostInterfaces field if non-nil, zero value otherwise.

### GetNumEthHostInterfacesOk

`func (o *ComputePhysical) GetNumEthHostInterfacesOk() (*int64, bool)`

GetNumEthHostInterfacesOk returns a tuple with the NumEthHostInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumEthHostInterfaces

`func (o *ComputePhysical) SetNumEthHostInterfaces(v int64)`

SetNumEthHostInterfaces sets NumEthHostInterfaces field to given value.

### HasNumEthHostInterfaces

`func (o *ComputePhysical) HasNumEthHostInterfaces() bool`

HasNumEthHostInterfaces returns a boolean if a field has been set.

### GetNumFcHostInterfaces

`func (o *ComputePhysical) GetNumFcHostInterfaces() int64`

GetNumFcHostInterfaces returns the NumFcHostInterfaces field if non-nil, zero value otherwise.

### GetNumFcHostInterfacesOk

`func (o *ComputePhysical) GetNumFcHostInterfacesOk() (*int64, bool)`

GetNumFcHostInterfacesOk returns a tuple with the NumFcHostInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumFcHostInterfaces

`func (o *ComputePhysical) SetNumFcHostInterfaces(v int64)`

SetNumFcHostInterfaces sets NumFcHostInterfaces field to given value.

### HasNumFcHostInterfaces

`func (o *ComputePhysical) HasNumFcHostInterfaces() bool`

HasNumFcHostInterfaces returns a boolean if a field has been set.

### GetNumThreads

`func (o *ComputePhysical) GetNumThreads() int64`

GetNumThreads returns the NumThreads field if non-nil, zero value otherwise.

### GetNumThreadsOk

`func (o *ComputePhysical) GetNumThreadsOk() (*int64, bool)`

GetNumThreadsOk returns a tuple with the NumThreads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumThreads

`func (o *ComputePhysical) SetNumThreads(v int64)`

SetNumThreads sets NumThreads field to given value.

### HasNumThreads

`func (o *ComputePhysical) HasNumThreads() bool`

HasNumThreads returns a boolean if a field has been set.

### GetOperPowerState

`func (o *ComputePhysical) GetOperPowerState() string`

GetOperPowerState returns the OperPowerState field if non-nil, zero value otherwise.

### GetOperPowerStateOk

`func (o *ComputePhysical) GetOperPowerStateOk() (*string, bool)`

GetOperPowerStateOk returns a tuple with the OperPowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperPowerState

`func (o *ComputePhysical) SetOperPowerState(v string)`

SetOperPowerState sets OperPowerState field to given value.

### HasOperPowerState

`func (o *ComputePhysical) HasOperPowerState() bool`

HasOperPowerState returns a boolean if a field has been set.

### GetOperReason

`func (o *ComputePhysical) GetOperReason() []string`

GetOperReason returns the OperReason field if non-nil, zero value otherwise.

### GetOperReasonOk

`func (o *ComputePhysical) GetOperReasonOk() (*[]string, bool)`

GetOperReasonOk returns a tuple with the OperReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperReason

`func (o *ComputePhysical) SetOperReason(v []string)`

SetOperReason sets OperReason field to given value.

### HasOperReason

`func (o *ComputePhysical) HasOperReason() bool`

HasOperReason returns a boolean if a field has been set.

### SetOperReasonNil

`func (o *ComputePhysical) SetOperReasonNil(b bool)`

 SetOperReasonNil sets the value for OperReason to be an explicit nil

### UnsetOperReason
`func (o *ComputePhysical) UnsetOperReason()`

UnsetOperReason ensures that no value is present for OperReason, not even an explicit nil
### GetOperState

`func (o *ComputePhysical) GetOperState() string`

GetOperState returns the OperState field if non-nil, zero value otherwise.

### GetOperStateOk

`func (o *ComputePhysical) GetOperStateOk() (*string, bool)`

GetOperStateOk returns a tuple with the OperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperState

`func (o *ComputePhysical) SetOperState(v string)`

SetOperState sets OperState field to given value.

### HasOperState

`func (o *ComputePhysical) HasOperState() bool`

HasOperState returns a boolean if a field has been set.

### GetOperability

`func (o *ComputePhysical) GetOperability() string`

GetOperability returns the Operability field if non-nil, zero value otherwise.

### GetOperabilityOk

`func (o *ComputePhysical) GetOperabilityOk() (*string, bool)`

GetOperabilityOk returns a tuple with the Operability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperability

`func (o *ComputePhysical) SetOperability(v string)`

SetOperability sets Operability field to given value.

### HasOperability

`func (o *ComputePhysical) HasOperability() bool`

HasOperability returns a boolean if a field has been set.

### GetPlatformType

`func (o *ComputePhysical) GetPlatformType() string`

GetPlatformType returns the PlatformType field if non-nil, zero value otherwise.

### GetPlatformTypeOk

`func (o *ComputePhysical) GetPlatformTypeOk() (*string, bool)`

GetPlatformTypeOk returns a tuple with the PlatformType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformType

`func (o *ComputePhysical) SetPlatformType(v string)`

SetPlatformType sets PlatformType field to given value.

### HasPlatformType

`func (o *ComputePhysical) HasPlatformType() bool`

HasPlatformType returns a boolean if a field has been set.

### GetServiceProfile

`func (o *ComputePhysical) GetServiceProfile() string`

GetServiceProfile returns the ServiceProfile field if non-nil, zero value otherwise.

### GetServiceProfileOk

`func (o *ComputePhysical) GetServiceProfileOk() (*string, bool)`

GetServiceProfileOk returns a tuple with the ServiceProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceProfile

`func (o *ComputePhysical) SetServiceProfile(v string)`

SetServiceProfile sets ServiceProfile field to given value.

### HasServiceProfile

`func (o *ComputePhysical) HasServiceProfile() bool`

HasServiceProfile returns a boolean if a field has been set.

### GetTotalMemory

`func (o *ComputePhysical) GetTotalMemory() int64`

GetTotalMemory returns the TotalMemory field if non-nil, zero value otherwise.

### GetTotalMemoryOk

`func (o *ComputePhysical) GetTotalMemoryOk() (*int64, bool)`

GetTotalMemoryOk returns a tuple with the TotalMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMemory

`func (o *ComputePhysical) SetTotalMemory(v int64)`

SetTotalMemory sets TotalMemory field to given value.

### HasTotalMemory

`func (o *ComputePhysical) HasTotalMemory() bool`

HasTotalMemory returns a boolean if a field has been set.

### GetTunneledKvm

`func (o *ComputePhysical) GetTunneledKvm() bool`

GetTunneledKvm returns the TunneledKvm field if non-nil, zero value otherwise.

### GetTunneledKvmOk

`func (o *ComputePhysical) GetTunneledKvmOk() (*bool, bool)`

GetTunneledKvmOk returns a tuple with the TunneledKvm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunneledKvm

`func (o *ComputePhysical) SetTunneledKvm(v bool)`

SetTunneledKvm sets TunneledKvm field to given value.

### HasTunneledKvm

`func (o *ComputePhysical) HasTunneledKvm() bool`

HasTunneledKvm returns a boolean if a field has been set.

### GetUserLabel

`func (o *ComputePhysical) GetUserLabel() string`

GetUserLabel returns the UserLabel field if non-nil, zero value otherwise.

### GetUserLabelOk

`func (o *ComputePhysical) GetUserLabelOk() (*string, bool)`

GetUserLabelOk returns a tuple with the UserLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLabel

`func (o *ComputePhysical) SetUserLabel(v string)`

SetUserLabel sets UserLabel field to given value.

### HasUserLabel

`func (o *ComputePhysical) HasUserLabel() bool`

HasUserLabel returns a boolean if a field has been set.

### GetUuid

`func (o *ComputePhysical) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ComputePhysical) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ComputePhysical) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ComputePhysical) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetBootCddDevices

`func (o *ComputePhysical) GetBootCddDevices() []BootCddDeviceRelationship`

GetBootCddDevices returns the BootCddDevices field if non-nil, zero value otherwise.

### GetBootCddDevicesOk

`func (o *ComputePhysical) GetBootCddDevicesOk() (*[]BootCddDeviceRelationship, bool)`

GetBootCddDevicesOk returns a tuple with the BootCddDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootCddDevices

`func (o *ComputePhysical) SetBootCddDevices(v []BootCddDeviceRelationship)`

SetBootCddDevices sets BootCddDevices field to given value.

### HasBootCddDevices

`func (o *ComputePhysical) HasBootCddDevices() bool`

HasBootCddDevices returns a boolean if a field has been set.

### SetBootCddDevicesNil

`func (o *ComputePhysical) SetBootCddDevicesNil(b bool)`

 SetBootCddDevicesNil sets the value for BootCddDevices to be an explicit nil

### UnsetBootCddDevices
`func (o *ComputePhysical) UnsetBootCddDevices()`

UnsetBootCddDevices ensures that no value is present for BootCddDevices, not even an explicit nil
### GetBootDeviceBootSecurity

`func (o *ComputePhysical) GetBootDeviceBootSecurity() BootDeviceBootSecurityRelationship`

GetBootDeviceBootSecurity returns the BootDeviceBootSecurity field if non-nil, zero value otherwise.

### GetBootDeviceBootSecurityOk

`func (o *ComputePhysical) GetBootDeviceBootSecurityOk() (*BootDeviceBootSecurityRelationship, bool)`

GetBootDeviceBootSecurityOk returns a tuple with the BootDeviceBootSecurity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootDeviceBootSecurity

`func (o *ComputePhysical) SetBootDeviceBootSecurity(v BootDeviceBootSecurityRelationship)`

SetBootDeviceBootSecurity sets BootDeviceBootSecurity field to given value.

### HasBootDeviceBootSecurity

`func (o *ComputePhysical) HasBootDeviceBootSecurity() bool`

HasBootDeviceBootSecurity returns a boolean if a field has been set.

### SetBootDeviceBootSecurityNil

`func (o *ComputePhysical) SetBootDeviceBootSecurityNil(b bool)`

 SetBootDeviceBootSecurityNil sets the value for BootDeviceBootSecurity to be an explicit nil

### UnsetBootDeviceBootSecurity
`func (o *ComputePhysical) UnsetBootDeviceBootSecurity()`

UnsetBootDeviceBootSecurity ensures that no value is present for BootDeviceBootSecurity, not even an explicit nil
### GetBootHddDevices

`func (o *ComputePhysical) GetBootHddDevices() []BootHddDeviceRelationship`

GetBootHddDevices returns the BootHddDevices field if non-nil, zero value otherwise.

### GetBootHddDevicesOk

`func (o *ComputePhysical) GetBootHddDevicesOk() (*[]BootHddDeviceRelationship, bool)`

GetBootHddDevicesOk returns a tuple with the BootHddDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootHddDevices

`func (o *ComputePhysical) SetBootHddDevices(v []BootHddDeviceRelationship)`

SetBootHddDevices sets BootHddDevices field to given value.

### HasBootHddDevices

`func (o *ComputePhysical) HasBootHddDevices() bool`

HasBootHddDevices returns a boolean if a field has been set.

### SetBootHddDevicesNil

`func (o *ComputePhysical) SetBootHddDevicesNil(b bool)`

 SetBootHddDevicesNil sets the value for BootHddDevices to be an explicit nil

### UnsetBootHddDevices
`func (o *ComputePhysical) UnsetBootHddDevices()`

UnsetBootHddDevices ensures that no value is present for BootHddDevices, not even an explicit nil
### GetBootIscsiDevices

`func (o *ComputePhysical) GetBootIscsiDevices() []BootIscsiDeviceRelationship`

GetBootIscsiDevices returns the BootIscsiDevices field if non-nil, zero value otherwise.

### GetBootIscsiDevicesOk

`func (o *ComputePhysical) GetBootIscsiDevicesOk() (*[]BootIscsiDeviceRelationship, bool)`

GetBootIscsiDevicesOk returns a tuple with the BootIscsiDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootIscsiDevices

`func (o *ComputePhysical) SetBootIscsiDevices(v []BootIscsiDeviceRelationship)`

SetBootIscsiDevices sets BootIscsiDevices field to given value.

### HasBootIscsiDevices

`func (o *ComputePhysical) HasBootIscsiDevices() bool`

HasBootIscsiDevices returns a boolean if a field has been set.

### SetBootIscsiDevicesNil

`func (o *ComputePhysical) SetBootIscsiDevicesNil(b bool)`

 SetBootIscsiDevicesNil sets the value for BootIscsiDevices to be an explicit nil

### UnsetBootIscsiDevices
`func (o *ComputePhysical) UnsetBootIscsiDevices()`

UnsetBootIscsiDevices ensures that no value is present for BootIscsiDevices, not even an explicit nil
### GetBootNvmeDevices

`func (o *ComputePhysical) GetBootNvmeDevices() []BootNvmeDeviceRelationship`

GetBootNvmeDevices returns the BootNvmeDevices field if non-nil, zero value otherwise.

### GetBootNvmeDevicesOk

`func (o *ComputePhysical) GetBootNvmeDevicesOk() (*[]BootNvmeDeviceRelationship, bool)`

GetBootNvmeDevicesOk returns a tuple with the BootNvmeDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootNvmeDevices

`func (o *ComputePhysical) SetBootNvmeDevices(v []BootNvmeDeviceRelationship)`

SetBootNvmeDevices sets BootNvmeDevices field to given value.

### HasBootNvmeDevices

`func (o *ComputePhysical) HasBootNvmeDevices() bool`

HasBootNvmeDevices returns a boolean if a field has been set.

### SetBootNvmeDevicesNil

`func (o *ComputePhysical) SetBootNvmeDevicesNil(b bool)`

 SetBootNvmeDevicesNil sets the value for BootNvmeDevices to be an explicit nil

### UnsetBootNvmeDevices
`func (o *ComputePhysical) UnsetBootNvmeDevices()`

UnsetBootNvmeDevices ensures that no value is present for BootNvmeDevices, not even an explicit nil
### GetBootPchStorageDevices

`func (o *ComputePhysical) GetBootPchStorageDevices() []BootPchStorageDeviceRelationship`

GetBootPchStorageDevices returns the BootPchStorageDevices field if non-nil, zero value otherwise.

### GetBootPchStorageDevicesOk

`func (o *ComputePhysical) GetBootPchStorageDevicesOk() (*[]BootPchStorageDeviceRelationship, bool)`

GetBootPchStorageDevicesOk returns a tuple with the BootPchStorageDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootPchStorageDevices

`func (o *ComputePhysical) SetBootPchStorageDevices(v []BootPchStorageDeviceRelationship)`

SetBootPchStorageDevices sets BootPchStorageDevices field to given value.

### HasBootPchStorageDevices

`func (o *ComputePhysical) HasBootPchStorageDevices() bool`

HasBootPchStorageDevices returns a boolean if a field has been set.

### SetBootPchStorageDevicesNil

`func (o *ComputePhysical) SetBootPchStorageDevicesNil(b bool)`

 SetBootPchStorageDevicesNil sets the value for BootPchStorageDevices to be an explicit nil

### UnsetBootPchStorageDevices
`func (o *ComputePhysical) UnsetBootPchStorageDevices()`

UnsetBootPchStorageDevices ensures that no value is present for BootPchStorageDevices, not even an explicit nil
### GetBootPxeDevices

`func (o *ComputePhysical) GetBootPxeDevices() []BootPxeDeviceRelationship`

GetBootPxeDevices returns the BootPxeDevices field if non-nil, zero value otherwise.

### GetBootPxeDevicesOk

`func (o *ComputePhysical) GetBootPxeDevicesOk() (*[]BootPxeDeviceRelationship, bool)`

GetBootPxeDevicesOk returns a tuple with the BootPxeDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootPxeDevices

`func (o *ComputePhysical) SetBootPxeDevices(v []BootPxeDeviceRelationship)`

SetBootPxeDevices sets BootPxeDevices field to given value.

### HasBootPxeDevices

`func (o *ComputePhysical) HasBootPxeDevices() bool`

HasBootPxeDevices returns a boolean if a field has been set.

### SetBootPxeDevicesNil

`func (o *ComputePhysical) SetBootPxeDevicesNil(b bool)`

 SetBootPxeDevicesNil sets the value for BootPxeDevices to be an explicit nil

### UnsetBootPxeDevices
`func (o *ComputePhysical) UnsetBootPxeDevices()`

UnsetBootPxeDevices ensures that no value is present for BootPxeDevices, not even an explicit nil
### GetBootSanDevices

`func (o *ComputePhysical) GetBootSanDevices() []BootSanDeviceRelationship`

GetBootSanDevices returns the BootSanDevices field if non-nil, zero value otherwise.

### GetBootSanDevicesOk

`func (o *ComputePhysical) GetBootSanDevicesOk() (*[]BootSanDeviceRelationship, bool)`

GetBootSanDevicesOk returns a tuple with the BootSanDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootSanDevices

`func (o *ComputePhysical) SetBootSanDevices(v []BootSanDeviceRelationship)`

SetBootSanDevices sets BootSanDevices field to given value.

### HasBootSanDevices

`func (o *ComputePhysical) HasBootSanDevices() bool`

HasBootSanDevices returns a boolean if a field has been set.

### SetBootSanDevicesNil

`func (o *ComputePhysical) SetBootSanDevicesNil(b bool)`

 SetBootSanDevicesNil sets the value for BootSanDevices to be an explicit nil

### UnsetBootSanDevices
`func (o *ComputePhysical) UnsetBootSanDevices()`

UnsetBootSanDevices ensures that no value is present for BootSanDevices, not even an explicit nil
### GetBootSdDevices

`func (o *ComputePhysical) GetBootSdDevices() []BootSdDeviceRelationship`

GetBootSdDevices returns the BootSdDevices field if non-nil, zero value otherwise.

### GetBootSdDevicesOk

`func (o *ComputePhysical) GetBootSdDevicesOk() (*[]BootSdDeviceRelationship, bool)`

GetBootSdDevicesOk returns a tuple with the BootSdDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootSdDevices

`func (o *ComputePhysical) SetBootSdDevices(v []BootSdDeviceRelationship)`

SetBootSdDevices sets BootSdDevices field to given value.

### HasBootSdDevices

`func (o *ComputePhysical) HasBootSdDevices() bool`

HasBootSdDevices returns a boolean if a field has been set.

### SetBootSdDevicesNil

`func (o *ComputePhysical) SetBootSdDevicesNil(b bool)`

 SetBootSdDevicesNil sets the value for BootSdDevices to be an explicit nil

### UnsetBootSdDevices
`func (o *ComputePhysical) UnsetBootSdDevices()`

UnsetBootSdDevices ensures that no value is present for BootSdDevices, not even an explicit nil
### GetBootUefiShellDevices

`func (o *ComputePhysical) GetBootUefiShellDevices() []BootUefiShellDeviceRelationship`

GetBootUefiShellDevices returns the BootUefiShellDevices field if non-nil, zero value otherwise.

### GetBootUefiShellDevicesOk

`func (o *ComputePhysical) GetBootUefiShellDevicesOk() (*[]BootUefiShellDeviceRelationship, bool)`

GetBootUefiShellDevicesOk returns a tuple with the BootUefiShellDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootUefiShellDevices

`func (o *ComputePhysical) SetBootUefiShellDevices(v []BootUefiShellDeviceRelationship)`

SetBootUefiShellDevices sets BootUefiShellDevices field to given value.

### HasBootUefiShellDevices

`func (o *ComputePhysical) HasBootUefiShellDevices() bool`

HasBootUefiShellDevices returns a boolean if a field has been set.

### SetBootUefiShellDevicesNil

`func (o *ComputePhysical) SetBootUefiShellDevicesNil(b bool)`

 SetBootUefiShellDevicesNil sets the value for BootUefiShellDevices to be an explicit nil

### UnsetBootUefiShellDevices
`func (o *ComputePhysical) UnsetBootUefiShellDevices()`

UnsetBootUefiShellDevices ensures that no value is present for BootUefiShellDevices, not even an explicit nil
### GetBootUsbDevices

`func (o *ComputePhysical) GetBootUsbDevices() []BootUsbDeviceRelationship`

GetBootUsbDevices returns the BootUsbDevices field if non-nil, zero value otherwise.

### GetBootUsbDevicesOk

`func (o *ComputePhysical) GetBootUsbDevicesOk() (*[]BootUsbDeviceRelationship, bool)`

GetBootUsbDevicesOk returns a tuple with the BootUsbDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootUsbDevices

`func (o *ComputePhysical) SetBootUsbDevices(v []BootUsbDeviceRelationship)`

SetBootUsbDevices sets BootUsbDevices field to given value.

### HasBootUsbDevices

`func (o *ComputePhysical) HasBootUsbDevices() bool`

HasBootUsbDevices returns a boolean if a field has been set.

### SetBootUsbDevicesNil

`func (o *ComputePhysical) SetBootUsbDevicesNil(b bool)`

 SetBootUsbDevicesNil sets the value for BootUsbDevices to be an explicit nil

### UnsetBootUsbDevices
`func (o *ComputePhysical) UnsetBootUsbDevices()`

UnsetBootUsbDevices ensures that no value is present for BootUsbDevices, not even an explicit nil
### GetBootVmediaDevices

`func (o *ComputePhysical) GetBootVmediaDevices() []BootVmediaDeviceRelationship`

GetBootVmediaDevices returns the BootVmediaDevices field if non-nil, zero value otherwise.

### GetBootVmediaDevicesOk

`func (o *ComputePhysical) GetBootVmediaDevicesOk() (*[]BootVmediaDeviceRelationship, bool)`

GetBootVmediaDevicesOk returns a tuple with the BootVmediaDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootVmediaDevices

`func (o *ComputePhysical) SetBootVmediaDevices(v []BootVmediaDeviceRelationship)`

SetBootVmediaDevices sets BootVmediaDevices field to given value.

### HasBootVmediaDevices

`func (o *ComputePhysical) HasBootVmediaDevices() bool`

HasBootVmediaDevices returns a boolean if a field has been set.

### SetBootVmediaDevicesNil

`func (o *ComputePhysical) SetBootVmediaDevicesNil(b bool)`

 SetBootVmediaDevicesNil sets the value for BootVmediaDevices to be an explicit nil

### UnsetBootVmediaDevices
`func (o *ComputePhysical) UnsetBootVmediaDevices()`

UnsetBootVmediaDevices ensures that no value is present for BootVmediaDevices, not even an explicit nil
### GetMgmtIdentity

`func (o *ComputePhysical) GetMgmtIdentity() EquipmentPhysicalIdentityRelationship`

GetMgmtIdentity returns the MgmtIdentity field if non-nil, zero value otherwise.

### GetMgmtIdentityOk

`func (o *ComputePhysical) GetMgmtIdentityOk() (*EquipmentPhysicalIdentityRelationship, bool)`

GetMgmtIdentityOk returns a tuple with the MgmtIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgmtIdentity

`func (o *ComputePhysical) SetMgmtIdentity(v EquipmentPhysicalIdentityRelationship)`

SetMgmtIdentity sets MgmtIdentity field to given value.

### HasMgmtIdentity

`func (o *ComputePhysical) HasMgmtIdentity() bool`

HasMgmtIdentity returns a boolean if a field has been set.

### SetMgmtIdentityNil

`func (o *ComputePhysical) SetMgmtIdentityNil(b bool)`

 SetMgmtIdentityNil sets the value for MgmtIdentity to be an explicit nil

### UnsetMgmtIdentity
`func (o *ComputePhysical) UnsetMgmtIdentity()`

UnsetMgmtIdentity ensures that no value is present for MgmtIdentity, not even an explicit nil
### GetVmedia

`func (o *ComputePhysical) GetVmedia() ComputeVmediaRelationship`

GetVmedia returns the Vmedia field if non-nil, zero value otherwise.

### GetVmediaOk

`func (o *ComputePhysical) GetVmediaOk() (*ComputeVmediaRelationship, bool)`

GetVmediaOk returns a tuple with the Vmedia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmedia

`func (o *ComputePhysical) SetVmedia(v ComputeVmediaRelationship)`

SetVmedia sets Vmedia field to given value.

### HasVmedia

`func (o *ComputePhysical) HasVmedia() bool`

HasVmedia returns a boolean if a field has been set.

### SetVmediaNil

`func (o *ComputePhysical) SetVmediaNil(b bool)`

 SetVmediaNil sets the value for Vmedia to be an explicit nil

### UnsetVmedia
`func (o *ComputePhysical) UnsetVmedia()`

UnsetVmedia ensures that no value is present for Vmedia, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



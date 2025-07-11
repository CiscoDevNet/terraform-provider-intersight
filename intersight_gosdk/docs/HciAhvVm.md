# HciAhvVm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hci.AhvVm"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hci.AhvVm"]
**BiosUuid** | Pointer to **string** | The BIOS UUID of the VM, similar to physical server&#39;s serial number. | [optional] [readonly] 
**CreationTime** | Pointer to **time.Time** | The time the VM was created. | [optional] [readonly] 
**EnabledCpuFeatures** | Pointer to **[]string** |  | [optional] 
**GenerationUuid** | Pointer to **string** | The generation UUID of the VM. | [optional] [readonly] 
**GuestTools** | Pointer to [**NullableHciAhvGuestTools**](HciAhvGuestTools.md) |  | [optional] 
**HardwareClockTimezone** | Pointer to **string** | VM hardware clock timezone in IANA TZDB format (America/Los_Angeles).  Default: UTC. | [optional] [readonly] 
**IsAgentVm** | Pointer to **bool** | Indicates whether the VM is an agent VM or not. When their host enters maintenance mode, once the normal VMs are evacuated, the  agent VMs are powered off. When the host is restored, agent VMs  are powered on before the normal VMs are restored. In other words,  agent VMs cannot be HA-protected or live migrated. | [optional] [readonly] 
**IsBrandingEnabled** | Pointer to **bool** | Indicates whether to remove AHV branding from VM firmware tables or not. | [optional] [readonly] 
**IsCpuHotplugEnabled** | Pointer to **bool** | The CPU hotplug status of the VM. It indicates whether the CPU  hotplug feature should be enabled for the VM or not. If enabled,  the VM can add or remove vCPUs while the VM is running. | [optional] [readonly] 
**IsCpuPassthroughEnabled** | Pointer to **bool** | The CPU passthrough status of the VM. It Indicates whether to passthrough the host CPU features to the guest or not.  Enabling this will make VM incapable of live migration. | [optional] [readonly] 
**IsCrossClusterMigrationInProgress** | Pointer to **bool** | Indicates whether the VM is in the process of cross cluster migration or not. | [optional] [readonly] 
**IsGpuConsoleEnabled** | Pointer to **bool** | The GPU console status of the VM. It indicates whether the GPU  console should be enabled for the VM or not. If enabled, the VM  will have access to the GPU console. | [optional] [readonly] 
**IsLiveMigrateCapable** | Pointer to **bool** | Indicates whether the VM is live migrate capable or not.  If the VM is not live migrate capable, it cannot be live migrated. | [optional] [readonly] 
**IsMemoryOvercommitEnabled** | Pointer to **bool** | The memory overcommit status of the VM. It indicates whether the memory overcommit feature should be enabled for the VM or not. If enabled, parts of the VM memory may reside outside of the  hypervisor physical memory. Once enabled, it should be expected  that the VM may suffer performance degradation. | [optional] [readonly] 
**IsScsiControllerEnabled** | Pointer to **bool** | The SCSI controller status of the VM. It indicates whether the  SCSI controller should be enabled for the VM or not. If enabled,  the VM will have access to the SCSI controller. | [optional] [readonly] 
**IsVcpuHardPinningEnabled** | Pointer to **bool** | The hard pinning status of the vCPU. It indicates whether the vCPUs  should be hard pinned to specific pCPUs or not. | [optional] [readonly] 
**IsVgaConsoleEnabled** | Pointer to **bool** | Indicates whether to enable VGA console for the VM or not. | [optional] [readonly] 
**IsVtpmEnabled** | Pointer to **bool** | Indicates whether the VM has a virtual TPM enabled or not. | [optional] [readonly] 
**MachineType** | Pointer to **string** | The machine type of the VM. Possible values are PC, PSERIES, Q35. | [optional] [readonly] 
**NumNumaNodes** | Pointer to **int32** | Number of NUMA nodes. 0 means NUMA is disabled. | [optional] [readonly] 
**NumSockets** | Pointer to **int32** | The number of sockets of the VM. | [optional] [readonly] 
**NumThreadsPerCore** | Pointer to **int32** | The number of threads per core of the VM. | [optional] [readonly] 
**ProtectionType** | Pointer to **string** | The type of protection applied on a VM.  Possible values are UNPROTECTED, PD_PROTECTED, and RULE_PROTECTED. PD_PROTECTED indicates a VM is protected using the Prism Element.  RULE_PROTECTED indicates a VM protection using the Prism Central. | [optional] [readonly] 
**SourceUuid** | Pointer to **string** | The source UUID of the VM. | [optional] [readonly] 
**UpdateTime** | Pointer to **time.Time** | The time the VM was last updated. | [optional] [readonly] 
**VtpmModuleVersion** | Pointer to **string** | The version of the vTPM module. | [optional] [readonly] 
**Disks** | Pointer to [**[]HciAhvVmDiskRelationship**](HciAhvVmDiskRelationship.md) | An array of relationships to hciAhvVmDisk resources. | [optional] [readonly] 
**Gpus** | Pointer to [**[]HciAhvVmGpuRelationship**](HciAhvVmGpuRelationship.md) | An array of relationships to hciAhvVmGpu resources. | [optional] [readonly] 
**Nics** | Pointer to [**[]HciAhvVmNicRelationship**](HciAhvVmNicRelationship.md) | An array of relationships to hciAhvVmNic resources. | [optional] [readonly] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewHciAhvVm

`func NewHciAhvVm(classId string, objectType string, ) *HciAhvVm`

NewHciAhvVm instantiates a new HciAhvVm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHciAhvVmWithDefaults

`func NewHciAhvVmWithDefaults() *HciAhvVm`

NewHciAhvVmWithDefaults instantiates a new HciAhvVm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HciAhvVm) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HciAhvVm) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HciAhvVm) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HciAhvVm) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HciAhvVm) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HciAhvVm) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBiosUuid

`func (o *HciAhvVm) GetBiosUuid() string`

GetBiosUuid returns the BiosUuid field if non-nil, zero value otherwise.

### GetBiosUuidOk

`func (o *HciAhvVm) GetBiosUuidOk() (*string, bool)`

GetBiosUuidOk returns a tuple with the BiosUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiosUuid

`func (o *HciAhvVm) SetBiosUuid(v string)`

SetBiosUuid sets BiosUuid field to given value.

### HasBiosUuid

`func (o *HciAhvVm) HasBiosUuid() bool`

HasBiosUuid returns a boolean if a field has been set.

### GetCreationTime

`func (o *HciAhvVm) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *HciAhvVm) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *HciAhvVm) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *HciAhvVm) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetEnabledCpuFeatures

`func (o *HciAhvVm) GetEnabledCpuFeatures() []string`

GetEnabledCpuFeatures returns the EnabledCpuFeatures field if non-nil, zero value otherwise.

### GetEnabledCpuFeaturesOk

`func (o *HciAhvVm) GetEnabledCpuFeaturesOk() (*[]string, bool)`

GetEnabledCpuFeaturesOk returns a tuple with the EnabledCpuFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledCpuFeatures

`func (o *HciAhvVm) SetEnabledCpuFeatures(v []string)`

SetEnabledCpuFeatures sets EnabledCpuFeatures field to given value.

### HasEnabledCpuFeatures

`func (o *HciAhvVm) HasEnabledCpuFeatures() bool`

HasEnabledCpuFeatures returns a boolean if a field has been set.

### SetEnabledCpuFeaturesNil

`func (o *HciAhvVm) SetEnabledCpuFeaturesNil(b bool)`

 SetEnabledCpuFeaturesNil sets the value for EnabledCpuFeatures to be an explicit nil

### UnsetEnabledCpuFeatures
`func (o *HciAhvVm) UnsetEnabledCpuFeatures()`

UnsetEnabledCpuFeatures ensures that no value is present for EnabledCpuFeatures, not even an explicit nil
### GetGenerationUuid

`func (o *HciAhvVm) GetGenerationUuid() string`

GetGenerationUuid returns the GenerationUuid field if non-nil, zero value otherwise.

### GetGenerationUuidOk

`func (o *HciAhvVm) GetGenerationUuidOk() (*string, bool)`

GetGenerationUuidOk returns a tuple with the GenerationUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerationUuid

`func (o *HciAhvVm) SetGenerationUuid(v string)`

SetGenerationUuid sets GenerationUuid field to given value.

### HasGenerationUuid

`func (o *HciAhvVm) HasGenerationUuid() bool`

HasGenerationUuid returns a boolean if a field has been set.

### GetGuestTools

`func (o *HciAhvVm) GetGuestTools() HciAhvGuestTools`

GetGuestTools returns the GuestTools field if non-nil, zero value otherwise.

### GetGuestToolsOk

`func (o *HciAhvVm) GetGuestToolsOk() (*HciAhvGuestTools, bool)`

GetGuestToolsOk returns a tuple with the GuestTools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestTools

`func (o *HciAhvVm) SetGuestTools(v HciAhvGuestTools)`

SetGuestTools sets GuestTools field to given value.

### HasGuestTools

`func (o *HciAhvVm) HasGuestTools() bool`

HasGuestTools returns a boolean if a field has been set.

### SetGuestToolsNil

`func (o *HciAhvVm) SetGuestToolsNil(b bool)`

 SetGuestToolsNil sets the value for GuestTools to be an explicit nil

### UnsetGuestTools
`func (o *HciAhvVm) UnsetGuestTools()`

UnsetGuestTools ensures that no value is present for GuestTools, not even an explicit nil
### GetHardwareClockTimezone

`func (o *HciAhvVm) GetHardwareClockTimezone() string`

GetHardwareClockTimezone returns the HardwareClockTimezone field if non-nil, zero value otherwise.

### GetHardwareClockTimezoneOk

`func (o *HciAhvVm) GetHardwareClockTimezoneOk() (*string, bool)`

GetHardwareClockTimezoneOk returns a tuple with the HardwareClockTimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareClockTimezone

`func (o *HciAhvVm) SetHardwareClockTimezone(v string)`

SetHardwareClockTimezone sets HardwareClockTimezone field to given value.

### HasHardwareClockTimezone

`func (o *HciAhvVm) HasHardwareClockTimezone() bool`

HasHardwareClockTimezone returns a boolean if a field has been set.

### GetIsAgentVm

`func (o *HciAhvVm) GetIsAgentVm() bool`

GetIsAgentVm returns the IsAgentVm field if non-nil, zero value otherwise.

### GetIsAgentVmOk

`func (o *HciAhvVm) GetIsAgentVmOk() (*bool, bool)`

GetIsAgentVmOk returns a tuple with the IsAgentVm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAgentVm

`func (o *HciAhvVm) SetIsAgentVm(v bool)`

SetIsAgentVm sets IsAgentVm field to given value.

### HasIsAgentVm

`func (o *HciAhvVm) HasIsAgentVm() bool`

HasIsAgentVm returns a boolean if a field has been set.

### GetIsBrandingEnabled

`func (o *HciAhvVm) GetIsBrandingEnabled() bool`

GetIsBrandingEnabled returns the IsBrandingEnabled field if non-nil, zero value otherwise.

### GetIsBrandingEnabledOk

`func (o *HciAhvVm) GetIsBrandingEnabledOk() (*bool, bool)`

GetIsBrandingEnabledOk returns a tuple with the IsBrandingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBrandingEnabled

`func (o *HciAhvVm) SetIsBrandingEnabled(v bool)`

SetIsBrandingEnabled sets IsBrandingEnabled field to given value.

### HasIsBrandingEnabled

`func (o *HciAhvVm) HasIsBrandingEnabled() bool`

HasIsBrandingEnabled returns a boolean if a field has been set.

### GetIsCpuHotplugEnabled

`func (o *HciAhvVm) GetIsCpuHotplugEnabled() bool`

GetIsCpuHotplugEnabled returns the IsCpuHotplugEnabled field if non-nil, zero value otherwise.

### GetIsCpuHotplugEnabledOk

`func (o *HciAhvVm) GetIsCpuHotplugEnabledOk() (*bool, bool)`

GetIsCpuHotplugEnabledOk returns a tuple with the IsCpuHotplugEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCpuHotplugEnabled

`func (o *HciAhvVm) SetIsCpuHotplugEnabled(v bool)`

SetIsCpuHotplugEnabled sets IsCpuHotplugEnabled field to given value.

### HasIsCpuHotplugEnabled

`func (o *HciAhvVm) HasIsCpuHotplugEnabled() bool`

HasIsCpuHotplugEnabled returns a boolean if a field has been set.

### GetIsCpuPassthroughEnabled

`func (o *HciAhvVm) GetIsCpuPassthroughEnabled() bool`

GetIsCpuPassthroughEnabled returns the IsCpuPassthroughEnabled field if non-nil, zero value otherwise.

### GetIsCpuPassthroughEnabledOk

`func (o *HciAhvVm) GetIsCpuPassthroughEnabledOk() (*bool, bool)`

GetIsCpuPassthroughEnabledOk returns a tuple with the IsCpuPassthroughEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCpuPassthroughEnabled

`func (o *HciAhvVm) SetIsCpuPassthroughEnabled(v bool)`

SetIsCpuPassthroughEnabled sets IsCpuPassthroughEnabled field to given value.

### HasIsCpuPassthroughEnabled

`func (o *HciAhvVm) HasIsCpuPassthroughEnabled() bool`

HasIsCpuPassthroughEnabled returns a boolean if a field has been set.

### GetIsCrossClusterMigrationInProgress

`func (o *HciAhvVm) GetIsCrossClusterMigrationInProgress() bool`

GetIsCrossClusterMigrationInProgress returns the IsCrossClusterMigrationInProgress field if non-nil, zero value otherwise.

### GetIsCrossClusterMigrationInProgressOk

`func (o *HciAhvVm) GetIsCrossClusterMigrationInProgressOk() (*bool, bool)`

GetIsCrossClusterMigrationInProgressOk returns a tuple with the IsCrossClusterMigrationInProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCrossClusterMigrationInProgress

`func (o *HciAhvVm) SetIsCrossClusterMigrationInProgress(v bool)`

SetIsCrossClusterMigrationInProgress sets IsCrossClusterMigrationInProgress field to given value.

### HasIsCrossClusterMigrationInProgress

`func (o *HciAhvVm) HasIsCrossClusterMigrationInProgress() bool`

HasIsCrossClusterMigrationInProgress returns a boolean if a field has been set.

### GetIsGpuConsoleEnabled

`func (o *HciAhvVm) GetIsGpuConsoleEnabled() bool`

GetIsGpuConsoleEnabled returns the IsGpuConsoleEnabled field if non-nil, zero value otherwise.

### GetIsGpuConsoleEnabledOk

`func (o *HciAhvVm) GetIsGpuConsoleEnabledOk() (*bool, bool)`

GetIsGpuConsoleEnabledOk returns a tuple with the IsGpuConsoleEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGpuConsoleEnabled

`func (o *HciAhvVm) SetIsGpuConsoleEnabled(v bool)`

SetIsGpuConsoleEnabled sets IsGpuConsoleEnabled field to given value.

### HasIsGpuConsoleEnabled

`func (o *HciAhvVm) HasIsGpuConsoleEnabled() bool`

HasIsGpuConsoleEnabled returns a boolean if a field has been set.

### GetIsLiveMigrateCapable

`func (o *HciAhvVm) GetIsLiveMigrateCapable() bool`

GetIsLiveMigrateCapable returns the IsLiveMigrateCapable field if non-nil, zero value otherwise.

### GetIsLiveMigrateCapableOk

`func (o *HciAhvVm) GetIsLiveMigrateCapableOk() (*bool, bool)`

GetIsLiveMigrateCapableOk returns a tuple with the IsLiveMigrateCapable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLiveMigrateCapable

`func (o *HciAhvVm) SetIsLiveMigrateCapable(v bool)`

SetIsLiveMigrateCapable sets IsLiveMigrateCapable field to given value.

### HasIsLiveMigrateCapable

`func (o *HciAhvVm) HasIsLiveMigrateCapable() bool`

HasIsLiveMigrateCapable returns a boolean if a field has been set.

### GetIsMemoryOvercommitEnabled

`func (o *HciAhvVm) GetIsMemoryOvercommitEnabled() bool`

GetIsMemoryOvercommitEnabled returns the IsMemoryOvercommitEnabled field if non-nil, zero value otherwise.

### GetIsMemoryOvercommitEnabledOk

`func (o *HciAhvVm) GetIsMemoryOvercommitEnabledOk() (*bool, bool)`

GetIsMemoryOvercommitEnabledOk returns a tuple with the IsMemoryOvercommitEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMemoryOvercommitEnabled

`func (o *HciAhvVm) SetIsMemoryOvercommitEnabled(v bool)`

SetIsMemoryOvercommitEnabled sets IsMemoryOvercommitEnabled field to given value.

### HasIsMemoryOvercommitEnabled

`func (o *HciAhvVm) HasIsMemoryOvercommitEnabled() bool`

HasIsMemoryOvercommitEnabled returns a boolean if a field has been set.

### GetIsScsiControllerEnabled

`func (o *HciAhvVm) GetIsScsiControllerEnabled() bool`

GetIsScsiControllerEnabled returns the IsScsiControllerEnabled field if non-nil, zero value otherwise.

### GetIsScsiControllerEnabledOk

`func (o *HciAhvVm) GetIsScsiControllerEnabledOk() (*bool, bool)`

GetIsScsiControllerEnabledOk returns a tuple with the IsScsiControllerEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsScsiControllerEnabled

`func (o *HciAhvVm) SetIsScsiControllerEnabled(v bool)`

SetIsScsiControllerEnabled sets IsScsiControllerEnabled field to given value.

### HasIsScsiControllerEnabled

`func (o *HciAhvVm) HasIsScsiControllerEnabled() bool`

HasIsScsiControllerEnabled returns a boolean if a field has been set.

### GetIsVcpuHardPinningEnabled

`func (o *HciAhvVm) GetIsVcpuHardPinningEnabled() bool`

GetIsVcpuHardPinningEnabled returns the IsVcpuHardPinningEnabled field if non-nil, zero value otherwise.

### GetIsVcpuHardPinningEnabledOk

`func (o *HciAhvVm) GetIsVcpuHardPinningEnabledOk() (*bool, bool)`

GetIsVcpuHardPinningEnabledOk returns a tuple with the IsVcpuHardPinningEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVcpuHardPinningEnabled

`func (o *HciAhvVm) SetIsVcpuHardPinningEnabled(v bool)`

SetIsVcpuHardPinningEnabled sets IsVcpuHardPinningEnabled field to given value.

### HasIsVcpuHardPinningEnabled

`func (o *HciAhvVm) HasIsVcpuHardPinningEnabled() bool`

HasIsVcpuHardPinningEnabled returns a boolean if a field has been set.

### GetIsVgaConsoleEnabled

`func (o *HciAhvVm) GetIsVgaConsoleEnabled() bool`

GetIsVgaConsoleEnabled returns the IsVgaConsoleEnabled field if non-nil, zero value otherwise.

### GetIsVgaConsoleEnabledOk

`func (o *HciAhvVm) GetIsVgaConsoleEnabledOk() (*bool, bool)`

GetIsVgaConsoleEnabledOk returns a tuple with the IsVgaConsoleEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVgaConsoleEnabled

`func (o *HciAhvVm) SetIsVgaConsoleEnabled(v bool)`

SetIsVgaConsoleEnabled sets IsVgaConsoleEnabled field to given value.

### HasIsVgaConsoleEnabled

`func (o *HciAhvVm) HasIsVgaConsoleEnabled() bool`

HasIsVgaConsoleEnabled returns a boolean if a field has been set.

### GetIsVtpmEnabled

`func (o *HciAhvVm) GetIsVtpmEnabled() bool`

GetIsVtpmEnabled returns the IsVtpmEnabled field if non-nil, zero value otherwise.

### GetIsVtpmEnabledOk

`func (o *HciAhvVm) GetIsVtpmEnabledOk() (*bool, bool)`

GetIsVtpmEnabledOk returns a tuple with the IsVtpmEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVtpmEnabled

`func (o *HciAhvVm) SetIsVtpmEnabled(v bool)`

SetIsVtpmEnabled sets IsVtpmEnabled field to given value.

### HasIsVtpmEnabled

`func (o *HciAhvVm) HasIsVtpmEnabled() bool`

HasIsVtpmEnabled returns a boolean if a field has been set.

### GetMachineType

`func (o *HciAhvVm) GetMachineType() string`

GetMachineType returns the MachineType field if non-nil, zero value otherwise.

### GetMachineTypeOk

`func (o *HciAhvVm) GetMachineTypeOk() (*string, bool)`

GetMachineTypeOk returns a tuple with the MachineType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineType

`func (o *HciAhvVm) SetMachineType(v string)`

SetMachineType sets MachineType field to given value.

### HasMachineType

`func (o *HciAhvVm) HasMachineType() bool`

HasMachineType returns a boolean if a field has been set.

### GetNumNumaNodes

`func (o *HciAhvVm) GetNumNumaNodes() int32`

GetNumNumaNodes returns the NumNumaNodes field if non-nil, zero value otherwise.

### GetNumNumaNodesOk

`func (o *HciAhvVm) GetNumNumaNodesOk() (*int32, bool)`

GetNumNumaNodesOk returns a tuple with the NumNumaNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumNumaNodes

`func (o *HciAhvVm) SetNumNumaNodes(v int32)`

SetNumNumaNodes sets NumNumaNodes field to given value.

### HasNumNumaNodes

`func (o *HciAhvVm) HasNumNumaNodes() bool`

HasNumNumaNodes returns a boolean if a field has been set.

### GetNumSockets

`func (o *HciAhvVm) GetNumSockets() int32`

GetNumSockets returns the NumSockets field if non-nil, zero value otherwise.

### GetNumSocketsOk

`func (o *HciAhvVm) GetNumSocketsOk() (*int32, bool)`

GetNumSocketsOk returns a tuple with the NumSockets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumSockets

`func (o *HciAhvVm) SetNumSockets(v int32)`

SetNumSockets sets NumSockets field to given value.

### HasNumSockets

`func (o *HciAhvVm) HasNumSockets() bool`

HasNumSockets returns a boolean if a field has been set.

### GetNumThreadsPerCore

`func (o *HciAhvVm) GetNumThreadsPerCore() int32`

GetNumThreadsPerCore returns the NumThreadsPerCore field if non-nil, zero value otherwise.

### GetNumThreadsPerCoreOk

`func (o *HciAhvVm) GetNumThreadsPerCoreOk() (*int32, bool)`

GetNumThreadsPerCoreOk returns a tuple with the NumThreadsPerCore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumThreadsPerCore

`func (o *HciAhvVm) SetNumThreadsPerCore(v int32)`

SetNumThreadsPerCore sets NumThreadsPerCore field to given value.

### HasNumThreadsPerCore

`func (o *HciAhvVm) HasNumThreadsPerCore() bool`

HasNumThreadsPerCore returns a boolean if a field has been set.

### GetProtectionType

`func (o *HciAhvVm) GetProtectionType() string`

GetProtectionType returns the ProtectionType field if non-nil, zero value otherwise.

### GetProtectionTypeOk

`func (o *HciAhvVm) GetProtectionTypeOk() (*string, bool)`

GetProtectionTypeOk returns a tuple with the ProtectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionType

`func (o *HciAhvVm) SetProtectionType(v string)`

SetProtectionType sets ProtectionType field to given value.

### HasProtectionType

`func (o *HciAhvVm) HasProtectionType() bool`

HasProtectionType returns a boolean if a field has been set.

### GetSourceUuid

`func (o *HciAhvVm) GetSourceUuid() string`

GetSourceUuid returns the SourceUuid field if non-nil, zero value otherwise.

### GetSourceUuidOk

`func (o *HciAhvVm) GetSourceUuidOk() (*string, bool)`

GetSourceUuidOk returns a tuple with the SourceUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUuid

`func (o *HciAhvVm) SetSourceUuid(v string)`

SetSourceUuid sets SourceUuid field to given value.

### HasSourceUuid

`func (o *HciAhvVm) HasSourceUuid() bool`

HasSourceUuid returns a boolean if a field has been set.

### GetUpdateTime

`func (o *HciAhvVm) GetUpdateTime() time.Time`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *HciAhvVm) GetUpdateTimeOk() (*time.Time, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *HciAhvVm) SetUpdateTime(v time.Time)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *HciAhvVm) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetVtpmModuleVersion

`func (o *HciAhvVm) GetVtpmModuleVersion() string`

GetVtpmModuleVersion returns the VtpmModuleVersion field if non-nil, zero value otherwise.

### GetVtpmModuleVersionOk

`func (o *HciAhvVm) GetVtpmModuleVersionOk() (*string, bool)`

GetVtpmModuleVersionOk returns a tuple with the VtpmModuleVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVtpmModuleVersion

`func (o *HciAhvVm) SetVtpmModuleVersion(v string)`

SetVtpmModuleVersion sets VtpmModuleVersion field to given value.

### HasVtpmModuleVersion

`func (o *HciAhvVm) HasVtpmModuleVersion() bool`

HasVtpmModuleVersion returns a boolean if a field has been set.

### GetDisks

`func (o *HciAhvVm) GetDisks() []HciAhvVmDiskRelationship`

GetDisks returns the Disks field if non-nil, zero value otherwise.

### GetDisksOk

`func (o *HciAhvVm) GetDisksOk() (*[]HciAhvVmDiskRelationship, bool)`

GetDisksOk returns a tuple with the Disks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisks

`func (o *HciAhvVm) SetDisks(v []HciAhvVmDiskRelationship)`

SetDisks sets Disks field to given value.

### HasDisks

`func (o *HciAhvVm) HasDisks() bool`

HasDisks returns a boolean if a field has been set.

### SetDisksNil

`func (o *HciAhvVm) SetDisksNil(b bool)`

 SetDisksNil sets the value for Disks to be an explicit nil

### UnsetDisks
`func (o *HciAhvVm) UnsetDisks()`

UnsetDisks ensures that no value is present for Disks, not even an explicit nil
### GetGpus

`func (o *HciAhvVm) GetGpus() []HciAhvVmGpuRelationship`

GetGpus returns the Gpus field if non-nil, zero value otherwise.

### GetGpusOk

`func (o *HciAhvVm) GetGpusOk() (*[]HciAhvVmGpuRelationship, bool)`

GetGpusOk returns a tuple with the Gpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpus

`func (o *HciAhvVm) SetGpus(v []HciAhvVmGpuRelationship)`

SetGpus sets Gpus field to given value.

### HasGpus

`func (o *HciAhvVm) HasGpus() bool`

HasGpus returns a boolean if a field has been set.

### SetGpusNil

`func (o *HciAhvVm) SetGpusNil(b bool)`

 SetGpusNil sets the value for Gpus to be an explicit nil

### UnsetGpus
`func (o *HciAhvVm) UnsetGpus()`

UnsetGpus ensures that no value is present for Gpus, not even an explicit nil
### GetNics

`func (o *HciAhvVm) GetNics() []HciAhvVmNicRelationship`

GetNics returns the Nics field if non-nil, zero value otherwise.

### GetNicsOk

`func (o *HciAhvVm) GetNicsOk() (*[]HciAhvVmNicRelationship, bool)`

GetNicsOk returns a tuple with the Nics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNics

`func (o *HciAhvVm) SetNics(v []HciAhvVmNicRelationship)`

SetNics sets Nics field to given value.

### HasNics

`func (o *HciAhvVm) HasNics() bool`

HasNics returns a boolean if a field has been set.

### SetNicsNil

`func (o *HciAhvVm) SetNicsNil(b bool)`

 SetNicsNil sets the value for Nics to be an explicit nil

### UnsetNics
`func (o *HciAhvVm) UnsetNics()`

UnsetNics ensures that no value is present for Nics, not even an explicit nil
### GetRegisteredDevice

`func (o *HciAhvVm) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *HciAhvVm) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *HciAhvVm) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *HciAhvVm) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *HciAhvVm) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *HciAhvVm) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



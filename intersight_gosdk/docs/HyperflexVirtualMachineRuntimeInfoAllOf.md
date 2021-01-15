# HyperflexVirtualMachineRuntimeInfoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.VirtualMachineRuntimeInfo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.VirtualMachineRuntimeInfo"]
**BiosUuid** | Pointer to **string** | BiosUuid of the Protected Virtual Machine. | [optional] [readonly] 
**ConnectionState** | Pointer to **string** | Connection state of the VM. | [optional] [readonly] 
**CpuUsage** | Pointer to **int64** | CPU Usage of Virtual Machine. | [optional] [readonly] 
**Folder** | Pointer to **string** | Folder which VM belongs to. | [optional] [readonly] 
**GuestFamily** | Pointer to **string** | Guest operating system family, if known. | [optional] [readonly] 
**GuestFullName** | Pointer to **string** | Guest operating system full name, if known. | [optional] [readonly] 
**GuestId** | Pointer to **string** | Guest operating system identifier (short name), if known. | [optional] [readonly] 
**GuestState** | Pointer to **string** | VMware guest reset, powercycle, or boot action states. | [optional] [readonly] 
**HostName** | Pointer to **string** | Hostname of Virtual Machine. | [optional] [readonly] 
**InstanceUuid** | Pointer to **string** | InstanceUuid of the Protected Virtual Machine. | [optional] [readonly] 
**MemoryMb** | Pointer to **int64** | CPU Memory in MB of Virtual Machine. | [optional] [readonly] 
**MemoryUsage** | Pointer to **int64** | Memory usage of Virtual Machine. | [optional] [readonly] 
**Moid** | Pointer to **string** | Virtual Machine unique MOID. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the Virtual Machine. | [optional] [readonly] 
**Networks** | Pointer to **[]string** |  | [optional] 
**NumCpu** | Pointer to **int64** | Number of CPUs for the VM. | [optional] [readonly] 
**PowerState** | Pointer to **string** | Power state of the Virtual Machine. | [optional] [readonly] 
**ProvisionedSize** | Pointer to **int64** | Provisioned Size of Virtual Machine. | [optional] [readonly] 
**ResourcePool** | Pointer to **string** | Resource pool which VM belongs to. | [optional] [readonly] 
**UsedSize** | Pointer to **int64** | Used Size of Virtual Machine. | [optional] [readonly] 
**Version** | Pointer to **string** | Version of the Virtual Machine. | [optional] [readonly] 
**VmxPath** | Pointer to **string** | Vmx Path in VC datastore format. | [optional] [readonly] 

## Methods

### NewHyperflexVirtualMachineRuntimeInfoAllOf

`func NewHyperflexVirtualMachineRuntimeInfoAllOf(classId string, objectType string, ) *HyperflexVirtualMachineRuntimeInfoAllOf`

NewHyperflexVirtualMachineRuntimeInfoAllOf instantiates a new HyperflexVirtualMachineRuntimeInfoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexVirtualMachineRuntimeInfoAllOfWithDefaults

`func NewHyperflexVirtualMachineRuntimeInfoAllOfWithDefaults() *HyperflexVirtualMachineRuntimeInfoAllOf`

NewHyperflexVirtualMachineRuntimeInfoAllOfWithDefaults instantiates a new HyperflexVirtualMachineRuntimeInfoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexVirtualMachineRuntimeInfoAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexVirtualMachineRuntimeInfoAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexVirtualMachineRuntimeInfoAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexVirtualMachineRuntimeInfoAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexVirtualMachineRuntimeInfoAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexVirtualMachineRuntimeInfoAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBiosUuid

`func (o *HyperflexVirtualMachineRuntimeInfoAllOf) GetBiosUuid() string`

GetBiosUuid returns the BiosUuid field if non-nil, zero value otherwise.

### GetBiosUuidOk

`func (o *HyperflexVirtualMachineRuntimeInfoAllOf) GetBiosUuidOk() (*string, bool)`

GetBiosUuidOk returns a tuple with the BiosUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiosUuid

`func (o *HyperflexVirtualMachineRuntimeInfoAllOf) SetBiosUuid(v string)`

SetBiosUuid sets BiosUuid field to given value.

### HasBiosUuid

`func (o *HyperflexVirtualMachineRuntimeInfoAllOf) HasBiosUuid() bool`

HasBiosUuid returns a boolean if a field has been set.

### GetConnectionState

`func (o *HyperflexVirtualMachineRuntimeInfoAllOf) GetConnectionState() string`

GetConnectionState returns the ConnectionState field if non-nil, zero value otherwise.

### GetConnectionStateOk

`func (o *HyperflexVirtualMachineRuntimeInfoAllOf) GetConnectionStateOk() (*string, bool)`

GetConnectionStateOk returns a tuple with the ConnectionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionState

`func (o *HyperflexVirtualMachineRuntimeInfoAllOf) SetConnectionState(v string)`

SetConnectionState sets ConnectionState field to given value.

### HasConnectionState

`func (o *HyperflexVirtualMachineRuntimeInfoAllOf) HasConnectionState() bool`

HasConnectionState returns a boolean if a field has been set.

### GetCpuUsage

`func (o *HyperflexVirtualMachineRuntimeInfoAllOf) GetCpuUsage() int64`

GetCpuUsage returns the CpuUsage field if non-nil, zero value otherwise.

### GetCpuUsageOk

`func (o *HyperflexVirtualMachineRuntimeInfoAllOf) GetCpuUsageOk() (*int64, bool)`

GetCpuUsageOk returns a tuple with the CpuUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUsage

`func (o *HyperflexVirtualMachineRuntimeInfoAllOf) SetCpuUsage(v int64)`

SetCpuUsage sets CpuUsage field to given value.

### HasCpuUsage

`func (o *HyperflexVirtualMachineRuntimeInfoAllOf) HasCpuUsage() bool`

HasCpuUsage returns a boolean if a field has been set.

### GetFolder

`func (o *HyperflexVirtualMachineRuntimeInfoAllOf) GetFolder() string`

GetFolder returns the Folder field if non-nil, zero value otherwise.

### GetFolderOk

`func (o *HyperflexVirtualMachineRuntimeInfoAllOf) GetFolderOk() (*string, bool)`

GetFolderOk returns a tuple with the Folder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolder

`func (o *HyperflexVirtualMachineRuntimeInfoAllOf) SetFolder(v string)`

SetFolder sets Folder field to given value.

### HasFolder

`func (o *HyperflexVirtualMachineRuntimeInfoAllOf) HasFolder() bool`

HasFolder returns a boolean if a field has been set.

### GetGuestFamily

`func (o *HyperflexVirtualMachineRuntimeInfoAllOf) GetGuestFamily() string`

GetGuestFamily returns the GuestFamily field if non-nil, zero value otherwise.

### GetGuestFamilyOk

`func (o *HyperflexVirtualMachineRuntimeInfoAllOf) GetGuestFamilyOk() (*string, bool)`

GetGuestFamilyOk returns a tuple with the GuestFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestFamily

`func (o *HyperflexVirtualMachineRuntimeInfoAllOf) SetGuestFamily(v string)`

SetGuestFamily sets GuestFamily field to given value.

### HasGuestFamily

`func (o *HyperflexVirtualMachineRuntimeInfoAllOf) HasGuestFamily() bool`

HasGuestFamily returns a boolean if a field has been set.

### GetGuestFullName

`func (o *HyperflexVirtualMachineRuntimeInfoAllOf) GetGuestFullName() string`

GetGuestFullName returns the GuestFullName field if non-nil, zero value otherwise.

### GetGuestFullNameOk

`func (o *HyperflexVirtualMachineRuntimeInfoAllOf) GetGuestFullNameOk() (*string, bool)`

GetGuestFullNameOk returns a tuple with the GuestFullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestFullName

`func (o *HyperflexVirtualMachineRuntimeInfoAllOf) SetGuestFullName(v string)`

SetGuestFullName sets GuestFullName field to given value.

### HasGuestFullName

`func (o *HyperflexVirtualMachineRuntimeInfoAllOf) HasGuestFullName() bool`

HasGuestFullName returns a boolean if a field has been set.

### GetGuestId

`func (o *HyperflexVirtualMachineRuntimeInfoAllOf) GetGuestId() string`

GetGuestId returns the GuestId field if non-nil, zero value otherwise.

### GetGuestIdOk

`func (o *HyperflexVirtualMachineRuntimeInfoAllOf) GetGuestIdOk() (*string, bool)`

GetGuestIdOk returns a tuple with the GuestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestId

`func (o *HyperflexVirtualMachineRuntimeInfoAllOf) SetGuestId(v string)`

SetGuestId sets GuestId field to given value.

### HasGuestId

`func (o *HyperflexVirtualMachineRuntimeInfoAllOf) HasGuestId() bool`

HasGuestId returns a boolean if a field has been set.

### GetGuestState

`func (o *HyperflexVirtualMachineRuntimeInfoAllOf) GetGuestState() string`

GetGuestState returns the GuestState field if non-nil, zero value otherwise.

### GetGuestStateOk

`func (o *HyperflexVirtualMachineRuntimeInfoAllOf) GetGuestStateOk() (*string, bool)`

GetGuestStateOk returns a tuple with the GuestState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestState

`func (o *HyperflexVirtualMachineRuntimeInfoAllOf) SetGuestState(v string)`

SetGuestState sets GuestState field to given value.

### HasGuestState

`func (o *HyperflexVirtualMachineRuntimeInfoAllOf) HasGuestState() bool`

HasGuestState returns a boolean if a field has been set.

### GetHostName

`func (o *HyperflexVirtualMachineRuntimeInfoAllOf) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *HyperflexVirtualMachineRuntimeInfoAllOf) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *HyperflexVirtualMachineRuntimeInfoAllOf) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *HyperflexVirtualMachineRuntimeInfoAllOf) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetInstanceUuid

`func (o *HyperflexVirtualMachineRuntimeInfoAllOf) GetInstanceUuid() string`

GetInstanceUuid returns the InstanceUuid field if non-nil, zero value otherwise.

### GetInstanceUuidOk

`func (o *HyperflexVirtualMachineRuntimeInfoAllOf) GetInstanceUuidOk() (*string, bool)`

GetInstanceUuidOk returns a tuple with the InstanceUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceUuid

`func (o *HyperflexVirtualMachineRuntimeInfoAllOf) SetInstanceUuid(v string)`

SetInstanceUuid sets InstanceUuid field to given value.

### HasInstanceUuid

`func (o *HyperflexVirtualMachineRuntimeInfoAllOf) HasInstanceUuid() bool`

HasInstanceUuid returns a boolean if a field has been set.

### GetMemoryMb

`func (o *HyperflexVirtualMachineRuntimeInfoAllOf) GetMemoryMb() int64`

GetMemoryMb returns the MemoryMb field if non-nil, zero value otherwise.

### GetMemoryMbOk

`func (o *HyperflexVirtualMachineRuntimeInfoAllOf) GetMemoryMbOk() (*int64, bool)`

GetMemoryMbOk returns a tuple with the MemoryMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryMb

`func (o *HyperflexVirtualMachineRuntimeInfoAllOf) SetMemoryMb(v int64)`

SetMemoryMb sets MemoryMb field to given value.

### HasMemoryMb

`func (o *HyperflexVirtualMachineRuntimeInfoAllOf) HasMemoryMb() bool`

HasMemoryMb returns a boolean if a field has been set.

### GetMemoryUsage

`func (o *HyperflexVirtualMachineRuntimeInfoAllOf) GetMemoryUsage() int64`

GetMemoryUsage returns the MemoryUsage field if non-nil, zero value otherwise.

### GetMemoryUsageOk

`func (o *HyperflexVirtualMachineRuntimeInfoAllOf) GetMemoryUsageOk() (*int64, bool)`

GetMemoryUsageOk returns a tuple with the MemoryUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryUsage

`func (o *HyperflexVirtualMachineRuntimeInfoAllOf) SetMemoryUsage(v int64)`

SetMemoryUsage sets MemoryUsage field to given value.

### HasMemoryUsage

`func (o *HyperflexVirtualMachineRuntimeInfoAllOf) HasMemoryUsage() bool`

HasMemoryUsage returns a boolean if a field has been set.

### GetMoid

`func (o *HyperflexVirtualMachineRuntimeInfoAllOf) GetMoid() string`

GetMoid returns the Moid field if non-nil, zero value otherwise.

### GetMoidOk

`func (o *HyperflexVirtualMachineRuntimeInfoAllOf) GetMoidOk() (*string, bool)`

GetMoidOk returns a tuple with the Moid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoid

`func (o *HyperflexVirtualMachineRuntimeInfoAllOf) SetMoid(v string)`

SetMoid sets Moid field to given value.

### HasMoid

`func (o *HyperflexVirtualMachineRuntimeInfoAllOf) HasMoid() bool`

HasMoid returns a boolean if a field has been set.

### GetName

`func (o *HyperflexVirtualMachineRuntimeInfoAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HyperflexVirtualMachineRuntimeInfoAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HyperflexVirtualMachineRuntimeInfoAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HyperflexVirtualMachineRuntimeInfoAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetworks

`func (o *HyperflexVirtualMachineRuntimeInfoAllOf) GetNetworks() []string`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *HyperflexVirtualMachineRuntimeInfoAllOf) GetNetworksOk() (*[]string, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *HyperflexVirtualMachineRuntimeInfoAllOf) SetNetworks(v []string)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *HyperflexVirtualMachineRuntimeInfoAllOf) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.

### SetNetworksNil

`func (o *HyperflexVirtualMachineRuntimeInfoAllOf) SetNetworksNil(b bool)`

 SetNetworksNil sets the value for Networks to be an explicit nil

### UnsetNetworks
`func (o *HyperflexVirtualMachineRuntimeInfoAllOf) UnsetNetworks()`

UnsetNetworks ensures that no value is present for Networks, not even an explicit nil
### GetNumCpu

`func (o *HyperflexVirtualMachineRuntimeInfoAllOf) GetNumCpu() int64`

GetNumCpu returns the NumCpu field if non-nil, zero value otherwise.

### GetNumCpuOk

`func (o *HyperflexVirtualMachineRuntimeInfoAllOf) GetNumCpuOk() (*int64, bool)`

GetNumCpuOk returns a tuple with the NumCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumCpu

`func (o *HyperflexVirtualMachineRuntimeInfoAllOf) SetNumCpu(v int64)`

SetNumCpu sets NumCpu field to given value.

### HasNumCpu

`func (o *HyperflexVirtualMachineRuntimeInfoAllOf) HasNumCpu() bool`

HasNumCpu returns a boolean if a field has been set.

### GetPowerState

`func (o *HyperflexVirtualMachineRuntimeInfoAllOf) GetPowerState() string`

GetPowerState returns the PowerState field if non-nil, zero value otherwise.

### GetPowerStateOk

`func (o *HyperflexVirtualMachineRuntimeInfoAllOf) GetPowerStateOk() (*string, bool)`

GetPowerStateOk returns a tuple with the PowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerState

`func (o *HyperflexVirtualMachineRuntimeInfoAllOf) SetPowerState(v string)`

SetPowerState sets PowerState field to given value.

### HasPowerState

`func (o *HyperflexVirtualMachineRuntimeInfoAllOf) HasPowerState() bool`

HasPowerState returns a boolean if a field has been set.

### GetProvisionedSize

`func (o *HyperflexVirtualMachineRuntimeInfoAllOf) GetProvisionedSize() int64`

GetProvisionedSize returns the ProvisionedSize field if non-nil, zero value otherwise.

### GetProvisionedSizeOk

`func (o *HyperflexVirtualMachineRuntimeInfoAllOf) GetProvisionedSizeOk() (*int64, bool)`

GetProvisionedSizeOk returns a tuple with the ProvisionedSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionedSize

`func (o *HyperflexVirtualMachineRuntimeInfoAllOf) SetProvisionedSize(v int64)`

SetProvisionedSize sets ProvisionedSize field to given value.

### HasProvisionedSize

`func (o *HyperflexVirtualMachineRuntimeInfoAllOf) HasProvisionedSize() bool`

HasProvisionedSize returns a boolean if a field has been set.

### GetResourcePool

`func (o *HyperflexVirtualMachineRuntimeInfoAllOf) GetResourcePool() string`

GetResourcePool returns the ResourcePool field if non-nil, zero value otherwise.

### GetResourcePoolOk

`func (o *HyperflexVirtualMachineRuntimeInfoAllOf) GetResourcePoolOk() (*string, bool)`

GetResourcePoolOk returns a tuple with the ResourcePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePool

`func (o *HyperflexVirtualMachineRuntimeInfoAllOf) SetResourcePool(v string)`

SetResourcePool sets ResourcePool field to given value.

### HasResourcePool

`func (o *HyperflexVirtualMachineRuntimeInfoAllOf) HasResourcePool() bool`

HasResourcePool returns a boolean if a field has been set.

### GetUsedSize

`func (o *HyperflexVirtualMachineRuntimeInfoAllOf) GetUsedSize() int64`

GetUsedSize returns the UsedSize field if non-nil, zero value otherwise.

### GetUsedSizeOk

`func (o *HyperflexVirtualMachineRuntimeInfoAllOf) GetUsedSizeOk() (*int64, bool)`

GetUsedSizeOk returns a tuple with the UsedSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedSize

`func (o *HyperflexVirtualMachineRuntimeInfoAllOf) SetUsedSize(v int64)`

SetUsedSize sets UsedSize field to given value.

### HasUsedSize

`func (o *HyperflexVirtualMachineRuntimeInfoAllOf) HasUsedSize() bool`

HasUsedSize returns a boolean if a field has been set.

### GetVersion

`func (o *HyperflexVirtualMachineRuntimeInfoAllOf) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *HyperflexVirtualMachineRuntimeInfoAllOf) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *HyperflexVirtualMachineRuntimeInfoAllOf) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *HyperflexVirtualMachineRuntimeInfoAllOf) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetVmxPath

`func (o *HyperflexVirtualMachineRuntimeInfoAllOf) GetVmxPath() string`

GetVmxPath returns the VmxPath field if non-nil, zero value otherwise.

### GetVmxPathOk

`func (o *HyperflexVirtualMachineRuntimeInfoAllOf) GetVmxPathOk() (*string, bool)`

GetVmxPathOk returns a tuple with the VmxPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmxPath

`func (o *HyperflexVirtualMachineRuntimeInfoAllOf) SetVmxPath(v string)`

SetVmxPath sets VmxPath field to given value.

### HasVmxPath

`func (o *HyperflexVirtualMachineRuntimeInfoAllOf) HasVmxPath() bool`

HasVmxPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



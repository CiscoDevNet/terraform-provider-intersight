# VirtualizationVmwareVirtualMachine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Annotation** | Pointer to **string** | List of annotations provided to this VM by user. Can be long. | [optional] 
**BootTime** | Pointer to [**time.Time**](time.Time.md) | Time when this VM booted up. | [optional] 
**ConfigName** | Pointer to **string** | The configuration name for this VM. This maybe the same as the guest hostname. | [optional] 
**ConnectionState** | Pointer to **string** | Shows if virtual machine is connected to vCenter. Values are Connected, Disconnected, Orphaned, Inaccessible, and Invalid. | [optional] 
**CpuHotAddEnabled** | Pointer to **bool** | Indicates if the capability to add CPUs to a running VM is enabled. | [optional] 
**CpuShares** | Pointer to [**VirtualizationVmwareVmCpuShareInfo**](virtualization.VmwareVmCpuShareInfo.md) |  | [optional] 
**CpuSocketInfo** | Pointer to [**VirtualizationVmwareVmCpuSocketInfo**](virtualization.VmwareVmCpuSocketInfo.md) |  | [optional] 
**CustomAttributes** | Pointer to **[]string** |  | [optional] 
**DefaultPowerOffType** | Pointer to **string** | Indicates how the VM will be powered off (soft, hard etc.). | [optional] 
**DhcpEnabled** | Pointer to **bool** | Shows if DHCP is used for IP/DNS on this VM. | [optional] 
**DiskCommitInfo** | Pointer to [**VirtualizationVmwareVmDiskCommitInfo**](virtualization.VmwareVmDiskCommitInfo.md) |  | [optional] 
**DnsServerList** | Pointer to **[]string** |  | [optional] 
**DnsSuffixList** | Pointer to **[]string** |  | [optional] 
**Folder** | Pointer to **string** | The folder name associated with this VM. | [optional] 
**GuestState** | Pointer to **string** | The state of the guest OS running on this VM. Could be running, not running etc. * &#x60;Unknown&#x60; - Indicates that the guest OS state cannot be determined. * &#x60;NotRunning&#x60; - Indicates that the guest OS is not running. * &#x60;Resetting&#x60; - Indicates that the guest OS is resetting. * &#x60;Running&#x60; - Indicates that the guest OS is running normally. * &#x60;ShuttingDown&#x60; - Indicates that the guest OS is shutting down. * &#x60;Standby&#x60; - Indicates that the guest OS is in standby mode. | [optional] [default to "Unknown"]
**InstanceUuid** | Pointer to **string** | UUID assigned by vCenter to every VM. | [optional] 
**IsTemplate** | Pointer to **bool** | If true, indicates that the entity refers to a template of a virtual machine and not a real virtual machine. | [optional] 
**MacAddress** | Pointer to **[]string** |  | [optional] 
**MemShares** | Pointer to [**VirtualizationVmwareVmMemoryShareInfo**](virtualization.VmwareVmMemoryShareInfo.md) |  | [optional] 
**MemoryHotAddEnabled** | Pointer to **bool** | Adding memory to a running VM. | [optional] 
**NetworkCount** | Pointer to **int64** | Indicates how many networks are used by this VM. | [optional] 
**PortGroups** | Pointer to **[]string** |  | [optional] 
**ProtectedVm** | Pointer to **bool** | Shows if this is a protected VM. VMs can be in protection groups. | [optional] 
**RemoteDisplayInfo** | Pointer to [**VirtualizationVmwareRemoteDisplayInfo**](virtualization.VmwareRemoteDisplayInfo.md) |  | [optional] 
**RemoteDisplayVncEnabled** | Pointer to **bool** | Shows if support for a remote VNC access is enabled. | [optional] 
**ResourcePool** | Pointer to **string** | Name of the resource pool to which this VM belongs (optional). | [optional] 
**ResourcePoolOwner** | Pointer to **string** | Who owns the resource pool. | [optional] 
**ResourcePoolParent** | Pointer to **string** | The parent of the current resource pool to which this VM belongs. | [optional] 
**ToolRunningStatus** | Pointer to **string** | Indicates if guest tools are running on this VM. Could be set to guestToolNotRunning or guestToolsRunning. | [optional] 
**ToolsVersion** | Pointer to **string** | The version of the guest tools, usually not specified. | [optional] 
**VmDiskCount** | Pointer to **int64** | Shows the number of disks assigned to this VM. | [optional] 
**VmOverallStatus** | Pointer to **string** | The operational state of the VM. Could be Available, Provisioned, Maintenance mode, Deleting, etc. | [optional] 
**VmPath** | Pointer to **string** | Example - [datastore3] VCSA-134/VCSA-134.vmx. | [optional] 
**VmVersion** | Pointer to **string** | Information about the version of this VM (vmx-09, vmx-11 etc.). | [optional] 
**VmVnicCount** | Pointer to **int64** | How many vnics are present. | [optional] 
**VnicDeviceConfigId** | Pointer to **string** | Information related to the guest info&#39;s VNIC virtual device. It is a comma-separated list. | [optional] 
**Cluster** | Pointer to [**VirtualizationVmwareClusterRelationship**](virtualization.VmwareCluster.Relationship.md) |  | [optional] 
**Datacenter** | Pointer to [**VirtualizationVmwareDatacenterRelationship**](virtualization.VmwareDatacenter.Relationship.md) |  | [optional] 
**Datastores** | Pointer to [**[]VirtualizationVmwareDatastoreRelationship**](virtualization.VmwareDatastore.Relationship.md) | An array of relationships to virtualizationVmwareDatastore resources. | [optional] [readonly] 
**Host** | Pointer to [**VirtualizationVmwareHostRelationship**](virtualization.VmwareHost.Relationship.md) |  | [optional] 

## Methods

### NewVirtualizationVmwareVirtualMachine

`func NewVirtualizationVmwareVirtualMachine() *VirtualizationVmwareVirtualMachine`

NewVirtualizationVmwareVirtualMachine instantiates a new VirtualizationVmwareVirtualMachine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationVmwareVirtualMachineWithDefaults

`func NewVirtualizationVmwareVirtualMachineWithDefaults() *VirtualizationVmwareVirtualMachine`

NewVirtualizationVmwareVirtualMachineWithDefaults instantiates a new VirtualizationVmwareVirtualMachine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnnotation

`func (o *VirtualizationVmwareVirtualMachine) GetAnnotation() string`

GetAnnotation returns the Annotation field if non-nil, zero value otherwise.

### GetAnnotationOk

`func (o *VirtualizationVmwareVirtualMachine) GetAnnotationOk() (*string, bool)`

GetAnnotationOk returns a tuple with the Annotation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotation

`func (o *VirtualizationVmwareVirtualMachine) SetAnnotation(v string)`

SetAnnotation sets Annotation field to given value.

### HasAnnotation

`func (o *VirtualizationVmwareVirtualMachine) HasAnnotation() bool`

HasAnnotation returns a boolean if a field has been set.

### GetBootTime

`func (o *VirtualizationVmwareVirtualMachine) GetBootTime() time.Time`

GetBootTime returns the BootTime field if non-nil, zero value otherwise.

### GetBootTimeOk

`func (o *VirtualizationVmwareVirtualMachine) GetBootTimeOk() (*time.Time, bool)`

GetBootTimeOk returns a tuple with the BootTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootTime

`func (o *VirtualizationVmwareVirtualMachine) SetBootTime(v time.Time)`

SetBootTime sets BootTime field to given value.

### HasBootTime

`func (o *VirtualizationVmwareVirtualMachine) HasBootTime() bool`

HasBootTime returns a boolean if a field has been set.

### GetConfigName

`func (o *VirtualizationVmwareVirtualMachine) GetConfigName() string`

GetConfigName returns the ConfigName field if non-nil, zero value otherwise.

### GetConfigNameOk

`func (o *VirtualizationVmwareVirtualMachine) GetConfigNameOk() (*string, bool)`

GetConfigNameOk returns a tuple with the ConfigName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigName

`func (o *VirtualizationVmwareVirtualMachine) SetConfigName(v string)`

SetConfigName sets ConfigName field to given value.

### HasConfigName

`func (o *VirtualizationVmwareVirtualMachine) HasConfigName() bool`

HasConfigName returns a boolean if a field has been set.

### GetConnectionState

`func (o *VirtualizationVmwareVirtualMachine) GetConnectionState() string`

GetConnectionState returns the ConnectionState field if non-nil, zero value otherwise.

### GetConnectionStateOk

`func (o *VirtualizationVmwareVirtualMachine) GetConnectionStateOk() (*string, bool)`

GetConnectionStateOk returns a tuple with the ConnectionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionState

`func (o *VirtualizationVmwareVirtualMachine) SetConnectionState(v string)`

SetConnectionState sets ConnectionState field to given value.

### HasConnectionState

`func (o *VirtualizationVmwareVirtualMachine) HasConnectionState() bool`

HasConnectionState returns a boolean if a field has been set.

### GetCpuHotAddEnabled

`func (o *VirtualizationVmwareVirtualMachine) GetCpuHotAddEnabled() bool`

GetCpuHotAddEnabled returns the CpuHotAddEnabled field if non-nil, zero value otherwise.

### GetCpuHotAddEnabledOk

`func (o *VirtualizationVmwareVirtualMachine) GetCpuHotAddEnabledOk() (*bool, bool)`

GetCpuHotAddEnabledOk returns a tuple with the CpuHotAddEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuHotAddEnabled

`func (o *VirtualizationVmwareVirtualMachine) SetCpuHotAddEnabled(v bool)`

SetCpuHotAddEnabled sets CpuHotAddEnabled field to given value.

### HasCpuHotAddEnabled

`func (o *VirtualizationVmwareVirtualMachine) HasCpuHotAddEnabled() bool`

HasCpuHotAddEnabled returns a boolean if a field has been set.

### GetCpuShares

`func (o *VirtualizationVmwareVirtualMachine) GetCpuShares() VirtualizationVmwareVmCpuShareInfo`

GetCpuShares returns the CpuShares field if non-nil, zero value otherwise.

### GetCpuSharesOk

`func (o *VirtualizationVmwareVirtualMachine) GetCpuSharesOk() (*VirtualizationVmwareVmCpuShareInfo, bool)`

GetCpuSharesOk returns a tuple with the CpuShares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuShares

`func (o *VirtualizationVmwareVirtualMachine) SetCpuShares(v VirtualizationVmwareVmCpuShareInfo)`

SetCpuShares sets CpuShares field to given value.

### HasCpuShares

`func (o *VirtualizationVmwareVirtualMachine) HasCpuShares() bool`

HasCpuShares returns a boolean if a field has been set.

### GetCpuSocketInfo

`func (o *VirtualizationVmwareVirtualMachine) GetCpuSocketInfo() VirtualizationVmwareVmCpuSocketInfo`

GetCpuSocketInfo returns the CpuSocketInfo field if non-nil, zero value otherwise.

### GetCpuSocketInfoOk

`func (o *VirtualizationVmwareVirtualMachine) GetCpuSocketInfoOk() (*VirtualizationVmwareVmCpuSocketInfo, bool)`

GetCpuSocketInfoOk returns a tuple with the CpuSocketInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuSocketInfo

`func (o *VirtualizationVmwareVirtualMachine) SetCpuSocketInfo(v VirtualizationVmwareVmCpuSocketInfo)`

SetCpuSocketInfo sets CpuSocketInfo field to given value.

### HasCpuSocketInfo

`func (o *VirtualizationVmwareVirtualMachine) HasCpuSocketInfo() bool`

HasCpuSocketInfo returns a boolean if a field has been set.

### GetCustomAttributes

`func (o *VirtualizationVmwareVirtualMachine) GetCustomAttributes() []string`

GetCustomAttributes returns the CustomAttributes field if non-nil, zero value otherwise.

### GetCustomAttributesOk

`func (o *VirtualizationVmwareVirtualMachine) GetCustomAttributesOk() (*[]string, bool)`

GetCustomAttributesOk returns a tuple with the CustomAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomAttributes

`func (o *VirtualizationVmwareVirtualMachine) SetCustomAttributes(v []string)`

SetCustomAttributes sets CustomAttributes field to given value.

### HasCustomAttributes

`func (o *VirtualizationVmwareVirtualMachine) HasCustomAttributes() bool`

HasCustomAttributes returns a boolean if a field has been set.

### GetDefaultPowerOffType

`func (o *VirtualizationVmwareVirtualMachine) GetDefaultPowerOffType() string`

GetDefaultPowerOffType returns the DefaultPowerOffType field if non-nil, zero value otherwise.

### GetDefaultPowerOffTypeOk

`func (o *VirtualizationVmwareVirtualMachine) GetDefaultPowerOffTypeOk() (*string, bool)`

GetDefaultPowerOffTypeOk returns a tuple with the DefaultPowerOffType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPowerOffType

`func (o *VirtualizationVmwareVirtualMachine) SetDefaultPowerOffType(v string)`

SetDefaultPowerOffType sets DefaultPowerOffType field to given value.

### HasDefaultPowerOffType

`func (o *VirtualizationVmwareVirtualMachine) HasDefaultPowerOffType() bool`

HasDefaultPowerOffType returns a boolean if a field has been set.

### GetDhcpEnabled

`func (o *VirtualizationVmwareVirtualMachine) GetDhcpEnabled() bool`

GetDhcpEnabled returns the DhcpEnabled field if non-nil, zero value otherwise.

### GetDhcpEnabledOk

`func (o *VirtualizationVmwareVirtualMachine) GetDhcpEnabledOk() (*bool, bool)`

GetDhcpEnabledOk returns a tuple with the DhcpEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpEnabled

`func (o *VirtualizationVmwareVirtualMachine) SetDhcpEnabled(v bool)`

SetDhcpEnabled sets DhcpEnabled field to given value.

### HasDhcpEnabled

`func (o *VirtualizationVmwareVirtualMachine) HasDhcpEnabled() bool`

HasDhcpEnabled returns a boolean if a field has been set.

### GetDiskCommitInfo

`func (o *VirtualizationVmwareVirtualMachine) GetDiskCommitInfo() VirtualizationVmwareVmDiskCommitInfo`

GetDiskCommitInfo returns the DiskCommitInfo field if non-nil, zero value otherwise.

### GetDiskCommitInfoOk

`func (o *VirtualizationVmwareVirtualMachine) GetDiskCommitInfoOk() (*VirtualizationVmwareVmDiskCommitInfo, bool)`

GetDiskCommitInfoOk returns a tuple with the DiskCommitInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskCommitInfo

`func (o *VirtualizationVmwareVirtualMachine) SetDiskCommitInfo(v VirtualizationVmwareVmDiskCommitInfo)`

SetDiskCommitInfo sets DiskCommitInfo field to given value.

### HasDiskCommitInfo

`func (o *VirtualizationVmwareVirtualMachine) HasDiskCommitInfo() bool`

HasDiskCommitInfo returns a boolean if a field has been set.

### GetDnsServerList

`func (o *VirtualizationVmwareVirtualMachine) GetDnsServerList() []string`

GetDnsServerList returns the DnsServerList field if non-nil, zero value otherwise.

### GetDnsServerListOk

`func (o *VirtualizationVmwareVirtualMachine) GetDnsServerListOk() (*[]string, bool)`

GetDnsServerListOk returns a tuple with the DnsServerList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsServerList

`func (o *VirtualizationVmwareVirtualMachine) SetDnsServerList(v []string)`

SetDnsServerList sets DnsServerList field to given value.

### HasDnsServerList

`func (o *VirtualizationVmwareVirtualMachine) HasDnsServerList() bool`

HasDnsServerList returns a boolean if a field has been set.

### GetDnsSuffixList

`func (o *VirtualizationVmwareVirtualMachine) GetDnsSuffixList() []string`

GetDnsSuffixList returns the DnsSuffixList field if non-nil, zero value otherwise.

### GetDnsSuffixListOk

`func (o *VirtualizationVmwareVirtualMachine) GetDnsSuffixListOk() (*[]string, bool)`

GetDnsSuffixListOk returns a tuple with the DnsSuffixList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSuffixList

`func (o *VirtualizationVmwareVirtualMachine) SetDnsSuffixList(v []string)`

SetDnsSuffixList sets DnsSuffixList field to given value.

### HasDnsSuffixList

`func (o *VirtualizationVmwareVirtualMachine) HasDnsSuffixList() bool`

HasDnsSuffixList returns a boolean if a field has been set.

### GetFolder

`func (o *VirtualizationVmwareVirtualMachine) GetFolder() string`

GetFolder returns the Folder field if non-nil, zero value otherwise.

### GetFolderOk

`func (o *VirtualizationVmwareVirtualMachine) GetFolderOk() (*string, bool)`

GetFolderOk returns a tuple with the Folder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolder

`func (o *VirtualizationVmwareVirtualMachine) SetFolder(v string)`

SetFolder sets Folder field to given value.

### HasFolder

`func (o *VirtualizationVmwareVirtualMachine) HasFolder() bool`

HasFolder returns a boolean if a field has been set.

### GetGuestState

`func (o *VirtualizationVmwareVirtualMachine) GetGuestState() string`

GetGuestState returns the GuestState field if non-nil, zero value otherwise.

### GetGuestStateOk

`func (o *VirtualizationVmwareVirtualMachine) GetGuestStateOk() (*string, bool)`

GetGuestStateOk returns a tuple with the GuestState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestState

`func (o *VirtualizationVmwareVirtualMachine) SetGuestState(v string)`

SetGuestState sets GuestState field to given value.

### HasGuestState

`func (o *VirtualizationVmwareVirtualMachine) HasGuestState() bool`

HasGuestState returns a boolean if a field has been set.

### GetInstanceUuid

`func (o *VirtualizationVmwareVirtualMachine) GetInstanceUuid() string`

GetInstanceUuid returns the InstanceUuid field if non-nil, zero value otherwise.

### GetInstanceUuidOk

`func (o *VirtualizationVmwareVirtualMachine) GetInstanceUuidOk() (*string, bool)`

GetInstanceUuidOk returns a tuple with the InstanceUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceUuid

`func (o *VirtualizationVmwareVirtualMachine) SetInstanceUuid(v string)`

SetInstanceUuid sets InstanceUuid field to given value.

### HasInstanceUuid

`func (o *VirtualizationVmwareVirtualMachine) HasInstanceUuid() bool`

HasInstanceUuid returns a boolean if a field has been set.

### GetIsTemplate

`func (o *VirtualizationVmwareVirtualMachine) GetIsTemplate() bool`

GetIsTemplate returns the IsTemplate field if non-nil, zero value otherwise.

### GetIsTemplateOk

`func (o *VirtualizationVmwareVirtualMachine) GetIsTemplateOk() (*bool, bool)`

GetIsTemplateOk returns a tuple with the IsTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTemplate

`func (o *VirtualizationVmwareVirtualMachine) SetIsTemplate(v bool)`

SetIsTemplate sets IsTemplate field to given value.

### HasIsTemplate

`func (o *VirtualizationVmwareVirtualMachine) HasIsTemplate() bool`

HasIsTemplate returns a boolean if a field has been set.

### GetMacAddress

`func (o *VirtualizationVmwareVirtualMachine) GetMacAddress() []string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *VirtualizationVmwareVirtualMachine) GetMacAddressOk() (*[]string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *VirtualizationVmwareVirtualMachine) SetMacAddress(v []string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *VirtualizationVmwareVirtualMachine) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetMemShares

`func (o *VirtualizationVmwareVirtualMachine) GetMemShares() VirtualizationVmwareVmMemoryShareInfo`

GetMemShares returns the MemShares field if non-nil, zero value otherwise.

### GetMemSharesOk

`func (o *VirtualizationVmwareVirtualMachine) GetMemSharesOk() (*VirtualizationVmwareVmMemoryShareInfo, bool)`

GetMemSharesOk returns a tuple with the MemShares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemShares

`func (o *VirtualizationVmwareVirtualMachine) SetMemShares(v VirtualizationVmwareVmMemoryShareInfo)`

SetMemShares sets MemShares field to given value.

### HasMemShares

`func (o *VirtualizationVmwareVirtualMachine) HasMemShares() bool`

HasMemShares returns a boolean if a field has been set.

### GetMemoryHotAddEnabled

`func (o *VirtualizationVmwareVirtualMachine) GetMemoryHotAddEnabled() bool`

GetMemoryHotAddEnabled returns the MemoryHotAddEnabled field if non-nil, zero value otherwise.

### GetMemoryHotAddEnabledOk

`func (o *VirtualizationVmwareVirtualMachine) GetMemoryHotAddEnabledOk() (*bool, bool)`

GetMemoryHotAddEnabledOk returns a tuple with the MemoryHotAddEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryHotAddEnabled

`func (o *VirtualizationVmwareVirtualMachine) SetMemoryHotAddEnabled(v bool)`

SetMemoryHotAddEnabled sets MemoryHotAddEnabled field to given value.

### HasMemoryHotAddEnabled

`func (o *VirtualizationVmwareVirtualMachine) HasMemoryHotAddEnabled() bool`

HasMemoryHotAddEnabled returns a boolean if a field has been set.

### GetNetworkCount

`func (o *VirtualizationVmwareVirtualMachine) GetNetworkCount() int64`

GetNetworkCount returns the NetworkCount field if non-nil, zero value otherwise.

### GetNetworkCountOk

`func (o *VirtualizationVmwareVirtualMachine) GetNetworkCountOk() (*int64, bool)`

GetNetworkCountOk returns a tuple with the NetworkCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkCount

`func (o *VirtualizationVmwareVirtualMachine) SetNetworkCount(v int64)`

SetNetworkCount sets NetworkCount field to given value.

### HasNetworkCount

`func (o *VirtualizationVmwareVirtualMachine) HasNetworkCount() bool`

HasNetworkCount returns a boolean if a field has been set.

### GetPortGroups

`func (o *VirtualizationVmwareVirtualMachine) GetPortGroups() []string`

GetPortGroups returns the PortGroups field if non-nil, zero value otherwise.

### GetPortGroupsOk

`func (o *VirtualizationVmwareVirtualMachine) GetPortGroupsOk() (*[]string, bool)`

GetPortGroupsOk returns a tuple with the PortGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortGroups

`func (o *VirtualizationVmwareVirtualMachine) SetPortGroups(v []string)`

SetPortGroups sets PortGroups field to given value.

### HasPortGroups

`func (o *VirtualizationVmwareVirtualMachine) HasPortGroups() bool`

HasPortGroups returns a boolean if a field has been set.

### GetProtectedVm

`func (o *VirtualizationVmwareVirtualMachine) GetProtectedVm() bool`

GetProtectedVm returns the ProtectedVm field if non-nil, zero value otherwise.

### GetProtectedVmOk

`func (o *VirtualizationVmwareVirtualMachine) GetProtectedVmOk() (*bool, bool)`

GetProtectedVmOk returns a tuple with the ProtectedVm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectedVm

`func (o *VirtualizationVmwareVirtualMachine) SetProtectedVm(v bool)`

SetProtectedVm sets ProtectedVm field to given value.

### HasProtectedVm

`func (o *VirtualizationVmwareVirtualMachine) HasProtectedVm() bool`

HasProtectedVm returns a boolean if a field has been set.

### GetRemoteDisplayInfo

`func (o *VirtualizationVmwareVirtualMachine) GetRemoteDisplayInfo() VirtualizationVmwareRemoteDisplayInfo`

GetRemoteDisplayInfo returns the RemoteDisplayInfo field if non-nil, zero value otherwise.

### GetRemoteDisplayInfoOk

`func (o *VirtualizationVmwareVirtualMachine) GetRemoteDisplayInfoOk() (*VirtualizationVmwareRemoteDisplayInfo, bool)`

GetRemoteDisplayInfoOk returns a tuple with the RemoteDisplayInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteDisplayInfo

`func (o *VirtualizationVmwareVirtualMachine) SetRemoteDisplayInfo(v VirtualizationVmwareRemoteDisplayInfo)`

SetRemoteDisplayInfo sets RemoteDisplayInfo field to given value.

### HasRemoteDisplayInfo

`func (o *VirtualizationVmwareVirtualMachine) HasRemoteDisplayInfo() bool`

HasRemoteDisplayInfo returns a boolean if a field has been set.

### GetRemoteDisplayVncEnabled

`func (o *VirtualizationVmwareVirtualMachine) GetRemoteDisplayVncEnabled() bool`

GetRemoteDisplayVncEnabled returns the RemoteDisplayVncEnabled field if non-nil, zero value otherwise.

### GetRemoteDisplayVncEnabledOk

`func (o *VirtualizationVmwareVirtualMachine) GetRemoteDisplayVncEnabledOk() (*bool, bool)`

GetRemoteDisplayVncEnabledOk returns a tuple with the RemoteDisplayVncEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteDisplayVncEnabled

`func (o *VirtualizationVmwareVirtualMachine) SetRemoteDisplayVncEnabled(v bool)`

SetRemoteDisplayVncEnabled sets RemoteDisplayVncEnabled field to given value.

### HasRemoteDisplayVncEnabled

`func (o *VirtualizationVmwareVirtualMachine) HasRemoteDisplayVncEnabled() bool`

HasRemoteDisplayVncEnabled returns a boolean if a field has been set.

### GetResourcePool

`func (o *VirtualizationVmwareVirtualMachine) GetResourcePool() string`

GetResourcePool returns the ResourcePool field if non-nil, zero value otherwise.

### GetResourcePoolOk

`func (o *VirtualizationVmwareVirtualMachine) GetResourcePoolOk() (*string, bool)`

GetResourcePoolOk returns a tuple with the ResourcePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePool

`func (o *VirtualizationVmwareVirtualMachine) SetResourcePool(v string)`

SetResourcePool sets ResourcePool field to given value.

### HasResourcePool

`func (o *VirtualizationVmwareVirtualMachine) HasResourcePool() bool`

HasResourcePool returns a boolean if a field has been set.

### GetResourcePoolOwner

`func (o *VirtualizationVmwareVirtualMachine) GetResourcePoolOwner() string`

GetResourcePoolOwner returns the ResourcePoolOwner field if non-nil, zero value otherwise.

### GetResourcePoolOwnerOk

`func (o *VirtualizationVmwareVirtualMachine) GetResourcePoolOwnerOk() (*string, bool)`

GetResourcePoolOwnerOk returns a tuple with the ResourcePoolOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolOwner

`func (o *VirtualizationVmwareVirtualMachine) SetResourcePoolOwner(v string)`

SetResourcePoolOwner sets ResourcePoolOwner field to given value.

### HasResourcePoolOwner

`func (o *VirtualizationVmwareVirtualMachine) HasResourcePoolOwner() bool`

HasResourcePoolOwner returns a boolean if a field has been set.

### GetResourcePoolParent

`func (o *VirtualizationVmwareVirtualMachine) GetResourcePoolParent() string`

GetResourcePoolParent returns the ResourcePoolParent field if non-nil, zero value otherwise.

### GetResourcePoolParentOk

`func (o *VirtualizationVmwareVirtualMachine) GetResourcePoolParentOk() (*string, bool)`

GetResourcePoolParentOk returns a tuple with the ResourcePoolParent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolParent

`func (o *VirtualizationVmwareVirtualMachine) SetResourcePoolParent(v string)`

SetResourcePoolParent sets ResourcePoolParent field to given value.

### HasResourcePoolParent

`func (o *VirtualizationVmwareVirtualMachine) HasResourcePoolParent() bool`

HasResourcePoolParent returns a boolean if a field has been set.

### GetToolRunningStatus

`func (o *VirtualizationVmwareVirtualMachine) GetToolRunningStatus() string`

GetToolRunningStatus returns the ToolRunningStatus field if non-nil, zero value otherwise.

### GetToolRunningStatusOk

`func (o *VirtualizationVmwareVirtualMachine) GetToolRunningStatusOk() (*string, bool)`

GetToolRunningStatusOk returns a tuple with the ToolRunningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolRunningStatus

`func (o *VirtualizationVmwareVirtualMachine) SetToolRunningStatus(v string)`

SetToolRunningStatus sets ToolRunningStatus field to given value.

### HasToolRunningStatus

`func (o *VirtualizationVmwareVirtualMachine) HasToolRunningStatus() bool`

HasToolRunningStatus returns a boolean if a field has been set.

### GetToolsVersion

`func (o *VirtualizationVmwareVirtualMachine) GetToolsVersion() string`

GetToolsVersion returns the ToolsVersion field if non-nil, zero value otherwise.

### GetToolsVersionOk

`func (o *VirtualizationVmwareVirtualMachine) GetToolsVersionOk() (*string, bool)`

GetToolsVersionOk returns a tuple with the ToolsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolsVersion

`func (o *VirtualizationVmwareVirtualMachine) SetToolsVersion(v string)`

SetToolsVersion sets ToolsVersion field to given value.

### HasToolsVersion

`func (o *VirtualizationVmwareVirtualMachine) HasToolsVersion() bool`

HasToolsVersion returns a boolean if a field has been set.

### GetVmDiskCount

`func (o *VirtualizationVmwareVirtualMachine) GetVmDiskCount() int64`

GetVmDiskCount returns the VmDiskCount field if non-nil, zero value otherwise.

### GetVmDiskCountOk

`func (o *VirtualizationVmwareVirtualMachine) GetVmDiskCountOk() (*int64, bool)`

GetVmDiskCountOk returns a tuple with the VmDiskCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmDiskCount

`func (o *VirtualizationVmwareVirtualMachine) SetVmDiskCount(v int64)`

SetVmDiskCount sets VmDiskCount field to given value.

### HasVmDiskCount

`func (o *VirtualizationVmwareVirtualMachine) HasVmDiskCount() bool`

HasVmDiskCount returns a boolean if a field has been set.

### GetVmOverallStatus

`func (o *VirtualizationVmwareVirtualMachine) GetVmOverallStatus() string`

GetVmOverallStatus returns the VmOverallStatus field if non-nil, zero value otherwise.

### GetVmOverallStatusOk

`func (o *VirtualizationVmwareVirtualMachine) GetVmOverallStatusOk() (*string, bool)`

GetVmOverallStatusOk returns a tuple with the VmOverallStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmOverallStatus

`func (o *VirtualizationVmwareVirtualMachine) SetVmOverallStatus(v string)`

SetVmOverallStatus sets VmOverallStatus field to given value.

### HasVmOverallStatus

`func (o *VirtualizationVmwareVirtualMachine) HasVmOverallStatus() bool`

HasVmOverallStatus returns a boolean if a field has been set.

### GetVmPath

`func (o *VirtualizationVmwareVirtualMachine) GetVmPath() string`

GetVmPath returns the VmPath field if non-nil, zero value otherwise.

### GetVmPathOk

`func (o *VirtualizationVmwareVirtualMachine) GetVmPathOk() (*string, bool)`

GetVmPathOk returns a tuple with the VmPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmPath

`func (o *VirtualizationVmwareVirtualMachine) SetVmPath(v string)`

SetVmPath sets VmPath field to given value.

### HasVmPath

`func (o *VirtualizationVmwareVirtualMachine) HasVmPath() bool`

HasVmPath returns a boolean if a field has been set.

### GetVmVersion

`func (o *VirtualizationVmwareVirtualMachine) GetVmVersion() string`

GetVmVersion returns the VmVersion field if non-nil, zero value otherwise.

### GetVmVersionOk

`func (o *VirtualizationVmwareVirtualMachine) GetVmVersionOk() (*string, bool)`

GetVmVersionOk returns a tuple with the VmVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmVersion

`func (o *VirtualizationVmwareVirtualMachine) SetVmVersion(v string)`

SetVmVersion sets VmVersion field to given value.

### HasVmVersion

`func (o *VirtualizationVmwareVirtualMachine) HasVmVersion() bool`

HasVmVersion returns a boolean if a field has been set.

### GetVmVnicCount

`func (o *VirtualizationVmwareVirtualMachine) GetVmVnicCount() int64`

GetVmVnicCount returns the VmVnicCount field if non-nil, zero value otherwise.

### GetVmVnicCountOk

`func (o *VirtualizationVmwareVirtualMachine) GetVmVnicCountOk() (*int64, bool)`

GetVmVnicCountOk returns a tuple with the VmVnicCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmVnicCount

`func (o *VirtualizationVmwareVirtualMachine) SetVmVnicCount(v int64)`

SetVmVnicCount sets VmVnicCount field to given value.

### HasVmVnicCount

`func (o *VirtualizationVmwareVirtualMachine) HasVmVnicCount() bool`

HasVmVnicCount returns a boolean if a field has been set.

### GetVnicDeviceConfigId

`func (o *VirtualizationVmwareVirtualMachine) GetVnicDeviceConfigId() string`

GetVnicDeviceConfigId returns the VnicDeviceConfigId field if non-nil, zero value otherwise.

### GetVnicDeviceConfigIdOk

`func (o *VirtualizationVmwareVirtualMachine) GetVnicDeviceConfigIdOk() (*string, bool)`

GetVnicDeviceConfigIdOk returns a tuple with the VnicDeviceConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVnicDeviceConfigId

`func (o *VirtualizationVmwareVirtualMachine) SetVnicDeviceConfigId(v string)`

SetVnicDeviceConfigId sets VnicDeviceConfigId field to given value.

### HasVnicDeviceConfigId

`func (o *VirtualizationVmwareVirtualMachine) HasVnicDeviceConfigId() bool`

HasVnicDeviceConfigId returns a boolean if a field has been set.

### GetCluster

`func (o *VirtualizationVmwareVirtualMachine) GetCluster() VirtualizationVmwareClusterRelationship`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *VirtualizationVmwareVirtualMachine) GetClusterOk() (*VirtualizationVmwareClusterRelationship, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *VirtualizationVmwareVirtualMachine) SetCluster(v VirtualizationVmwareClusterRelationship)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *VirtualizationVmwareVirtualMachine) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetDatacenter

`func (o *VirtualizationVmwareVirtualMachine) GetDatacenter() VirtualizationVmwareDatacenterRelationship`

GetDatacenter returns the Datacenter field if non-nil, zero value otherwise.

### GetDatacenterOk

`func (o *VirtualizationVmwareVirtualMachine) GetDatacenterOk() (*VirtualizationVmwareDatacenterRelationship, bool)`

GetDatacenterOk returns a tuple with the Datacenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenter

`func (o *VirtualizationVmwareVirtualMachine) SetDatacenter(v VirtualizationVmwareDatacenterRelationship)`

SetDatacenter sets Datacenter field to given value.

### HasDatacenter

`func (o *VirtualizationVmwareVirtualMachine) HasDatacenter() bool`

HasDatacenter returns a boolean if a field has been set.

### GetDatastores

`func (o *VirtualizationVmwareVirtualMachine) GetDatastores() []VirtualizationVmwareDatastoreRelationship`

GetDatastores returns the Datastores field if non-nil, zero value otherwise.

### GetDatastoresOk

`func (o *VirtualizationVmwareVirtualMachine) GetDatastoresOk() (*[]VirtualizationVmwareDatastoreRelationship, bool)`

GetDatastoresOk returns a tuple with the Datastores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastores

`func (o *VirtualizationVmwareVirtualMachine) SetDatastores(v []VirtualizationVmwareDatastoreRelationship)`

SetDatastores sets Datastores field to given value.

### HasDatastores

`func (o *VirtualizationVmwareVirtualMachine) HasDatastores() bool`

HasDatastores returns a boolean if a field has been set.

### SetDatastoresNil

`func (o *VirtualizationVmwareVirtualMachine) SetDatastoresNil(b bool)`

 SetDatastoresNil sets the value for Datastores to be an explicit nil

### UnsetDatastores
`func (o *VirtualizationVmwareVirtualMachine) UnsetDatastores()`

UnsetDatastores ensures that no value is present for Datastores, not even an explicit nil
### GetHost

`func (o *VirtualizationVmwareVirtualMachine) GetHost() VirtualizationVmwareHostRelationship`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *VirtualizationVmwareVirtualMachine) GetHostOk() (*VirtualizationVmwareHostRelationship, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *VirtualizationVmwareVirtualMachine) SetHost(v VirtualizationVmwareHostRelationship)`

SetHost sets Host field to given value.

### HasHost

`func (o *VirtualizationVmwareVirtualMachine) HasHost() bool`

HasHost returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



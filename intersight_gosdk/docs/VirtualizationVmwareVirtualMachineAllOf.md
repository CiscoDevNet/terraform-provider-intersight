# VirtualizationVmwareVirtualMachineAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "virtualization.VmwareVirtualMachine"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "virtualization.VmwareVirtualMachine"]
**Annotation** | Pointer to **string** | List of annotations provided to this VM by user. Can be long. | [optional] 
**ConfigName** | Pointer to **string** | The configuration name for this VM. This maybe the same as the guest hostname. | [optional] 
**ConnectionState** | Pointer to **string** | Shows if virtual machine is connected to vCenter. Values are Connected, Disconnected, Orphaned, Inaccessible, and Invalid. | [optional] 
**CpuHotAddEnabled** | Pointer to **bool** | Indicates if the capability to add CPUs to a running VM is enabled. | [optional] 
**CpuShares** | Pointer to [**NullableVirtualizationVmwareVmCpuShareInfo**](virtualization.VmwareVmCpuShareInfo.md) |  | [optional] 
**CpuSocketInfo** | Pointer to [**NullableVirtualizationVmwareVmCpuSocketInfo**](virtualization.VmwareVmCpuSocketInfo.md) |  | [optional] 
**CustomAttributes** | Pointer to **[]string** |  | [optional] 
**DefaultPowerOffType** | Pointer to **string** | Indicates how the VM will be powered off (soft, hard etc.). | [optional] 
**DhcpEnabled** | Pointer to **bool** | Shows if DHCP is used for IP/DNS on this VM. | [optional] 
**DiskCommitInfo** | Pointer to [**NullableVirtualizationVmwareVmDiskCommitInfo**](virtualization.VmwareVmDiskCommitInfo.md) |  | [optional] 
**DnsServerList** | Pointer to **[]string** |  | [optional] 
**DnsSuffixList** | Pointer to **[]string** |  | [optional] 
**ExtraConfig** | Pointer to **interface{}** | Additional custom configuration settings applied to this VM. It is a set of name-value pairs stored as json. | [optional] 
**Folder** | Pointer to **string** | The folder name associated with this VM. | [optional] 
**GuestState** | Pointer to **string** | The state of the guest OS running on this VM. Could be running, not running etc. * &#x60;Unknown&#x60; - Indicates that the guest OS state cannot be determined. * &#x60;NotRunning&#x60; - Indicates that the guest OS is not running. * &#x60;Resetting&#x60; - Indicates that the guest OS is resetting. * &#x60;Running&#x60; - Indicates that the guest OS is running normally. * &#x60;ShuttingDown&#x60; - Indicates that the guest OS is shutting down. * &#x60;Standby&#x60; - Indicates that the guest OS is in standby mode. | [optional] [default to "Unknown"]
**InstanceUuid** | Pointer to **string** | UUID assigned by vCenter to every VM. | [optional] 
**InventoryPath** | Pointer to **string** | Inventory path to the VM. Example - /DC/vm/folder/VMName. | [optional] 
**IsTemplate** | Pointer to **bool** | If true, indicates that the entity refers to a template of a virtual machine and not a real virtual machine. | [optional] 
**MacAddress** | Pointer to **[]string** |  | [optional] 
**MemShares** | Pointer to [**NullableVirtualizationVmwareVmMemoryShareInfo**](virtualization.VmwareVmMemoryShareInfo.md) |  | [optional] 
**MemoryHotAddEnabled** | Pointer to **bool** | Adding memory to a running VM. | [optional] 
**NetworkCount** | Pointer to **int64** | Indicates how many networks are used by this VM. | [optional] 
**PortGroups** | Pointer to **[]string** |  | [optional] 
**ProtectedVm** | Pointer to **bool** | Shows if this is a protected VM. VMs can be in protection groups. | [optional] 
**RemoteDisplayInfo** | Pointer to [**NullableVirtualizationVmwareRemoteDisplayInfo**](virtualization.VmwareRemoteDisplayInfo.md) |  | [optional] 
**RemoteDisplayVncEnabled** | Pointer to **bool** | Shows if support for a remote VNC access is enabled. | [optional] 
**ResourcePool** | Pointer to **string** | Name of the resource pool to which this VM belongs (optional). | [optional] 
**ResourcePoolOwner** | Pointer to **string** | Who owns the resource pool. | [optional] 
**ResourcePoolParent** | Pointer to **string** | The parent of the current resource pool to which this VM belongs. | [optional] 
**ToolRunningStatus** | Pointer to **string** | Indicates if guest tools are running on this VM. Could be set to guestToolNotRunning or guestToolsRunning. | [optional] 
**ToolsVersion** | Pointer to **string** | The version of the guest tools, usually not specified. | [optional] 
**VirtualDisks** | Pointer to **[]int64** |  | [optional] 
**VirtualNetworkInterfaces** | Pointer to **[]int64** |  | [optional] 
**VmDiskCount** | Pointer to **int64** | Shows the number of disks assigned to this VM. | [optional] 
**VmOverallStatus** | Pointer to **string** | The operational state of the VM. Could be Available, Provisioned, Maintenance mode, Deleting, etc. | [optional] 
**VmPath** | Pointer to **string** | Path to the vmx file of the VM. Example - [datastore3] VCSA-134/VCSA-134.vmx. | [optional] 
**VmVersion** | Pointer to **string** | Information about the version of this VM (vmx-09, vmx-11 etc.). | [optional] 
**VmVnicCount** | Pointer to **int64** | How many vnics are present. | [optional] 
**VnicDeviceConfigId** | Pointer to **string** | Information related to the guest info&#39;s VNIC virtual device. It is a comma-separated list. | [optional] 
**Cluster** | Pointer to [**VirtualizationVmwareClusterRelationship**](virtualization.VmwareCluster.Relationship.md) |  | [optional] 
**Datacenter** | Pointer to [**VirtualizationVmwareDatacenterRelationship**](virtualization.VmwareDatacenter.Relationship.md) |  | [optional] 
**Datastores** | Pointer to [**[]VirtualizationVmwareDatastoreRelationship**](VirtualizationVmwareDatastoreRelationship.md) | An array of relationships to virtualizationVmwareDatastore resources. | [optional] [readonly] 
**Host** | Pointer to [**VirtualizationVmwareHostRelationship**](virtualization.VmwareHost.Relationship.md) |  | [optional] 
**Networks** | Pointer to [**[]VirtualizationBaseNetworkRelationship**](VirtualizationBaseNetworkRelationship.md) | An array of relationships to virtualizationBaseNetwork resources. | [optional] [readonly] 
**ParentFolder** | Pointer to [**VirtualizationVmwareFolderRelationship**](virtualization.VmwareFolder.Relationship.md) |  | [optional] 

## Methods

### NewVirtualizationVmwareVirtualMachineAllOf

`func NewVirtualizationVmwareVirtualMachineAllOf(classId string, objectType string, ) *VirtualizationVmwareVirtualMachineAllOf`

NewVirtualizationVmwareVirtualMachineAllOf instantiates a new VirtualizationVmwareVirtualMachineAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationVmwareVirtualMachineAllOfWithDefaults

`func NewVirtualizationVmwareVirtualMachineAllOfWithDefaults() *VirtualizationVmwareVirtualMachineAllOf`

NewVirtualizationVmwareVirtualMachineAllOfWithDefaults instantiates a new VirtualizationVmwareVirtualMachineAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationVmwareVirtualMachineAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationVmwareVirtualMachineAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationVmwareVirtualMachineAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationVmwareVirtualMachineAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationVmwareVirtualMachineAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationVmwareVirtualMachineAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAnnotation

`func (o *VirtualizationVmwareVirtualMachineAllOf) GetAnnotation() string`

GetAnnotation returns the Annotation field if non-nil, zero value otherwise.

### GetAnnotationOk

`func (o *VirtualizationVmwareVirtualMachineAllOf) GetAnnotationOk() (*string, bool)`

GetAnnotationOk returns a tuple with the Annotation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotation

`func (o *VirtualizationVmwareVirtualMachineAllOf) SetAnnotation(v string)`

SetAnnotation sets Annotation field to given value.

### HasAnnotation

`func (o *VirtualizationVmwareVirtualMachineAllOf) HasAnnotation() bool`

HasAnnotation returns a boolean if a field has been set.

### GetConfigName

`func (o *VirtualizationVmwareVirtualMachineAllOf) GetConfigName() string`

GetConfigName returns the ConfigName field if non-nil, zero value otherwise.

### GetConfigNameOk

`func (o *VirtualizationVmwareVirtualMachineAllOf) GetConfigNameOk() (*string, bool)`

GetConfigNameOk returns a tuple with the ConfigName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigName

`func (o *VirtualizationVmwareVirtualMachineAllOf) SetConfigName(v string)`

SetConfigName sets ConfigName field to given value.

### HasConfigName

`func (o *VirtualizationVmwareVirtualMachineAllOf) HasConfigName() bool`

HasConfigName returns a boolean if a field has been set.

### GetConnectionState

`func (o *VirtualizationVmwareVirtualMachineAllOf) GetConnectionState() string`

GetConnectionState returns the ConnectionState field if non-nil, zero value otherwise.

### GetConnectionStateOk

`func (o *VirtualizationVmwareVirtualMachineAllOf) GetConnectionStateOk() (*string, bool)`

GetConnectionStateOk returns a tuple with the ConnectionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionState

`func (o *VirtualizationVmwareVirtualMachineAllOf) SetConnectionState(v string)`

SetConnectionState sets ConnectionState field to given value.

### HasConnectionState

`func (o *VirtualizationVmwareVirtualMachineAllOf) HasConnectionState() bool`

HasConnectionState returns a boolean if a field has been set.

### GetCpuHotAddEnabled

`func (o *VirtualizationVmwareVirtualMachineAllOf) GetCpuHotAddEnabled() bool`

GetCpuHotAddEnabled returns the CpuHotAddEnabled field if non-nil, zero value otherwise.

### GetCpuHotAddEnabledOk

`func (o *VirtualizationVmwareVirtualMachineAllOf) GetCpuHotAddEnabledOk() (*bool, bool)`

GetCpuHotAddEnabledOk returns a tuple with the CpuHotAddEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuHotAddEnabled

`func (o *VirtualizationVmwareVirtualMachineAllOf) SetCpuHotAddEnabled(v bool)`

SetCpuHotAddEnabled sets CpuHotAddEnabled field to given value.

### HasCpuHotAddEnabled

`func (o *VirtualizationVmwareVirtualMachineAllOf) HasCpuHotAddEnabled() bool`

HasCpuHotAddEnabled returns a boolean if a field has been set.

### GetCpuShares

`func (o *VirtualizationVmwareVirtualMachineAllOf) GetCpuShares() VirtualizationVmwareVmCpuShareInfo`

GetCpuShares returns the CpuShares field if non-nil, zero value otherwise.

### GetCpuSharesOk

`func (o *VirtualizationVmwareVirtualMachineAllOf) GetCpuSharesOk() (*VirtualizationVmwareVmCpuShareInfo, bool)`

GetCpuSharesOk returns a tuple with the CpuShares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuShares

`func (o *VirtualizationVmwareVirtualMachineAllOf) SetCpuShares(v VirtualizationVmwareVmCpuShareInfo)`

SetCpuShares sets CpuShares field to given value.

### HasCpuShares

`func (o *VirtualizationVmwareVirtualMachineAllOf) HasCpuShares() bool`

HasCpuShares returns a boolean if a field has been set.

### SetCpuSharesNil

`func (o *VirtualizationVmwareVirtualMachineAllOf) SetCpuSharesNil(b bool)`

 SetCpuSharesNil sets the value for CpuShares to be an explicit nil

### UnsetCpuShares
`func (o *VirtualizationVmwareVirtualMachineAllOf) UnsetCpuShares()`

UnsetCpuShares ensures that no value is present for CpuShares, not even an explicit nil
### GetCpuSocketInfo

`func (o *VirtualizationVmwareVirtualMachineAllOf) GetCpuSocketInfo() VirtualizationVmwareVmCpuSocketInfo`

GetCpuSocketInfo returns the CpuSocketInfo field if non-nil, zero value otherwise.

### GetCpuSocketInfoOk

`func (o *VirtualizationVmwareVirtualMachineAllOf) GetCpuSocketInfoOk() (*VirtualizationVmwareVmCpuSocketInfo, bool)`

GetCpuSocketInfoOk returns a tuple with the CpuSocketInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuSocketInfo

`func (o *VirtualizationVmwareVirtualMachineAllOf) SetCpuSocketInfo(v VirtualizationVmwareVmCpuSocketInfo)`

SetCpuSocketInfo sets CpuSocketInfo field to given value.

### HasCpuSocketInfo

`func (o *VirtualizationVmwareVirtualMachineAllOf) HasCpuSocketInfo() bool`

HasCpuSocketInfo returns a boolean if a field has been set.

### SetCpuSocketInfoNil

`func (o *VirtualizationVmwareVirtualMachineAllOf) SetCpuSocketInfoNil(b bool)`

 SetCpuSocketInfoNil sets the value for CpuSocketInfo to be an explicit nil

### UnsetCpuSocketInfo
`func (o *VirtualizationVmwareVirtualMachineAllOf) UnsetCpuSocketInfo()`

UnsetCpuSocketInfo ensures that no value is present for CpuSocketInfo, not even an explicit nil
### GetCustomAttributes

`func (o *VirtualizationVmwareVirtualMachineAllOf) GetCustomAttributes() []string`

GetCustomAttributes returns the CustomAttributes field if non-nil, zero value otherwise.

### GetCustomAttributesOk

`func (o *VirtualizationVmwareVirtualMachineAllOf) GetCustomAttributesOk() (*[]string, bool)`

GetCustomAttributesOk returns a tuple with the CustomAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomAttributes

`func (o *VirtualizationVmwareVirtualMachineAllOf) SetCustomAttributes(v []string)`

SetCustomAttributes sets CustomAttributes field to given value.

### HasCustomAttributes

`func (o *VirtualizationVmwareVirtualMachineAllOf) HasCustomAttributes() bool`

HasCustomAttributes returns a boolean if a field has been set.

### SetCustomAttributesNil

`func (o *VirtualizationVmwareVirtualMachineAllOf) SetCustomAttributesNil(b bool)`

 SetCustomAttributesNil sets the value for CustomAttributes to be an explicit nil

### UnsetCustomAttributes
`func (o *VirtualizationVmwareVirtualMachineAllOf) UnsetCustomAttributes()`

UnsetCustomAttributes ensures that no value is present for CustomAttributes, not even an explicit nil
### GetDefaultPowerOffType

`func (o *VirtualizationVmwareVirtualMachineAllOf) GetDefaultPowerOffType() string`

GetDefaultPowerOffType returns the DefaultPowerOffType field if non-nil, zero value otherwise.

### GetDefaultPowerOffTypeOk

`func (o *VirtualizationVmwareVirtualMachineAllOf) GetDefaultPowerOffTypeOk() (*string, bool)`

GetDefaultPowerOffTypeOk returns a tuple with the DefaultPowerOffType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPowerOffType

`func (o *VirtualizationVmwareVirtualMachineAllOf) SetDefaultPowerOffType(v string)`

SetDefaultPowerOffType sets DefaultPowerOffType field to given value.

### HasDefaultPowerOffType

`func (o *VirtualizationVmwareVirtualMachineAllOf) HasDefaultPowerOffType() bool`

HasDefaultPowerOffType returns a boolean if a field has been set.

### GetDhcpEnabled

`func (o *VirtualizationVmwareVirtualMachineAllOf) GetDhcpEnabled() bool`

GetDhcpEnabled returns the DhcpEnabled field if non-nil, zero value otherwise.

### GetDhcpEnabledOk

`func (o *VirtualizationVmwareVirtualMachineAllOf) GetDhcpEnabledOk() (*bool, bool)`

GetDhcpEnabledOk returns a tuple with the DhcpEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpEnabled

`func (o *VirtualizationVmwareVirtualMachineAllOf) SetDhcpEnabled(v bool)`

SetDhcpEnabled sets DhcpEnabled field to given value.

### HasDhcpEnabled

`func (o *VirtualizationVmwareVirtualMachineAllOf) HasDhcpEnabled() bool`

HasDhcpEnabled returns a boolean if a field has been set.

### GetDiskCommitInfo

`func (o *VirtualizationVmwareVirtualMachineAllOf) GetDiskCommitInfo() VirtualizationVmwareVmDiskCommitInfo`

GetDiskCommitInfo returns the DiskCommitInfo field if non-nil, zero value otherwise.

### GetDiskCommitInfoOk

`func (o *VirtualizationVmwareVirtualMachineAllOf) GetDiskCommitInfoOk() (*VirtualizationVmwareVmDiskCommitInfo, bool)`

GetDiskCommitInfoOk returns a tuple with the DiskCommitInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskCommitInfo

`func (o *VirtualizationVmwareVirtualMachineAllOf) SetDiskCommitInfo(v VirtualizationVmwareVmDiskCommitInfo)`

SetDiskCommitInfo sets DiskCommitInfo field to given value.

### HasDiskCommitInfo

`func (o *VirtualizationVmwareVirtualMachineAllOf) HasDiskCommitInfo() bool`

HasDiskCommitInfo returns a boolean if a field has been set.

### SetDiskCommitInfoNil

`func (o *VirtualizationVmwareVirtualMachineAllOf) SetDiskCommitInfoNil(b bool)`

 SetDiskCommitInfoNil sets the value for DiskCommitInfo to be an explicit nil

### UnsetDiskCommitInfo
`func (o *VirtualizationVmwareVirtualMachineAllOf) UnsetDiskCommitInfo()`

UnsetDiskCommitInfo ensures that no value is present for DiskCommitInfo, not even an explicit nil
### GetDnsServerList

`func (o *VirtualizationVmwareVirtualMachineAllOf) GetDnsServerList() []string`

GetDnsServerList returns the DnsServerList field if non-nil, zero value otherwise.

### GetDnsServerListOk

`func (o *VirtualizationVmwareVirtualMachineAllOf) GetDnsServerListOk() (*[]string, bool)`

GetDnsServerListOk returns a tuple with the DnsServerList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsServerList

`func (o *VirtualizationVmwareVirtualMachineAllOf) SetDnsServerList(v []string)`

SetDnsServerList sets DnsServerList field to given value.

### HasDnsServerList

`func (o *VirtualizationVmwareVirtualMachineAllOf) HasDnsServerList() bool`

HasDnsServerList returns a boolean if a field has been set.

### SetDnsServerListNil

`func (o *VirtualizationVmwareVirtualMachineAllOf) SetDnsServerListNil(b bool)`

 SetDnsServerListNil sets the value for DnsServerList to be an explicit nil

### UnsetDnsServerList
`func (o *VirtualizationVmwareVirtualMachineAllOf) UnsetDnsServerList()`

UnsetDnsServerList ensures that no value is present for DnsServerList, not even an explicit nil
### GetDnsSuffixList

`func (o *VirtualizationVmwareVirtualMachineAllOf) GetDnsSuffixList() []string`

GetDnsSuffixList returns the DnsSuffixList field if non-nil, zero value otherwise.

### GetDnsSuffixListOk

`func (o *VirtualizationVmwareVirtualMachineAllOf) GetDnsSuffixListOk() (*[]string, bool)`

GetDnsSuffixListOk returns a tuple with the DnsSuffixList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSuffixList

`func (o *VirtualizationVmwareVirtualMachineAllOf) SetDnsSuffixList(v []string)`

SetDnsSuffixList sets DnsSuffixList field to given value.

### HasDnsSuffixList

`func (o *VirtualizationVmwareVirtualMachineAllOf) HasDnsSuffixList() bool`

HasDnsSuffixList returns a boolean if a field has been set.

### SetDnsSuffixListNil

`func (o *VirtualizationVmwareVirtualMachineAllOf) SetDnsSuffixListNil(b bool)`

 SetDnsSuffixListNil sets the value for DnsSuffixList to be an explicit nil

### UnsetDnsSuffixList
`func (o *VirtualizationVmwareVirtualMachineAllOf) UnsetDnsSuffixList()`

UnsetDnsSuffixList ensures that no value is present for DnsSuffixList, not even an explicit nil
### GetExtraConfig

`func (o *VirtualizationVmwareVirtualMachineAllOf) GetExtraConfig() interface{}`

GetExtraConfig returns the ExtraConfig field if non-nil, zero value otherwise.

### GetExtraConfigOk

`func (o *VirtualizationVmwareVirtualMachineAllOf) GetExtraConfigOk() (*interface{}, bool)`

GetExtraConfigOk returns a tuple with the ExtraConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraConfig

`func (o *VirtualizationVmwareVirtualMachineAllOf) SetExtraConfig(v interface{})`

SetExtraConfig sets ExtraConfig field to given value.

### HasExtraConfig

`func (o *VirtualizationVmwareVirtualMachineAllOf) HasExtraConfig() bool`

HasExtraConfig returns a boolean if a field has been set.

### SetExtraConfigNil

`func (o *VirtualizationVmwareVirtualMachineAllOf) SetExtraConfigNil(b bool)`

 SetExtraConfigNil sets the value for ExtraConfig to be an explicit nil

### UnsetExtraConfig
`func (o *VirtualizationVmwareVirtualMachineAllOf) UnsetExtraConfig()`

UnsetExtraConfig ensures that no value is present for ExtraConfig, not even an explicit nil
### GetFolder

`func (o *VirtualizationVmwareVirtualMachineAllOf) GetFolder() string`

GetFolder returns the Folder field if non-nil, zero value otherwise.

### GetFolderOk

`func (o *VirtualizationVmwareVirtualMachineAllOf) GetFolderOk() (*string, bool)`

GetFolderOk returns a tuple with the Folder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolder

`func (o *VirtualizationVmwareVirtualMachineAllOf) SetFolder(v string)`

SetFolder sets Folder field to given value.

### HasFolder

`func (o *VirtualizationVmwareVirtualMachineAllOf) HasFolder() bool`

HasFolder returns a boolean if a field has been set.

### GetGuestState

`func (o *VirtualizationVmwareVirtualMachineAllOf) GetGuestState() string`

GetGuestState returns the GuestState field if non-nil, zero value otherwise.

### GetGuestStateOk

`func (o *VirtualizationVmwareVirtualMachineAllOf) GetGuestStateOk() (*string, bool)`

GetGuestStateOk returns a tuple with the GuestState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestState

`func (o *VirtualizationVmwareVirtualMachineAllOf) SetGuestState(v string)`

SetGuestState sets GuestState field to given value.

### HasGuestState

`func (o *VirtualizationVmwareVirtualMachineAllOf) HasGuestState() bool`

HasGuestState returns a boolean if a field has been set.

### GetInstanceUuid

`func (o *VirtualizationVmwareVirtualMachineAllOf) GetInstanceUuid() string`

GetInstanceUuid returns the InstanceUuid field if non-nil, zero value otherwise.

### GetInstanceUuidOk

`func (o *VirtualizationVmwareVirtualMachineAllOf) GetInstanceUuidOk() (*string, bool)`

GetInstanceUuidOk returns a tuple with the InstanceUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceUuid

`func (o *VirtualizationVmwareVirtualMachineAllOf) SetInstanceUuid(v string)`

SetInstanceUuid sets InstanceUuid field to given value.

### HasInstanceUuid

`func (o *VirtualizationVmwareVirtualMachineAllOf) HasInstanceUuid() bool`

HasInstanceUuid returns a boolean if a field has been set.

### GetInventoryPath

`func (o *VirtualizationVmwareVirtualMachineAllOf) GetInventoryPath() string`

GetInventoryPath returns the InventoryPath field if non-nil, zero value otherwise.

### GetInventoryPathOk

`func (o *VirtualizationVmwareVirtualMachineAllOf) GetInventoryPathOk() (*string, bool)`

GetInventoryPathOk returns a tuple with the InventoryPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryPath

`func (o *VirtualizationVmwareVirtualMachineAllOf) SetInventoryPath(v string)`

SetInventoryPath sets InventoryPath field to given value.

### HasInventoryPath

`func (o *VirtualizationVmwareVirtualMachineAllOf) HasInventoryPath() bool`

HasInventoryPath returns a boolean if a field has been set.

### GetIsTemplate

`func (o *VirtualizationVmwareVirtualMachineAllOf) GetIsTemplate() bool`

GetIsTemplate returns the IsTemplate field if non-nil, zero value otherwise.

### GetIsTemplateOk

`func (o *VirtualizationVmwareVirtualMachineAllOf) GetIsTemplateOk() (*bool, bool)`

GetIsTemplateOk returns a tuple with the IsTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTemplate

`func (o *VirtualizationVmwareVirtualMachineAllOf) SetIsTemplate(v bool)`

SetIsTemplate sets IsTemplate field to given value.

### HasIsTemplate

`func (o *VirtualizationVmwareVirtualMachineAllOf) HasIsTemplate() bool`

HasIsTemplate returns a boolean if a field has been set.

### GetMacAddress

`func (o *VirtualizationVmwareVirtualMachineAllOf) GetMacAddress() []string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *VirtualizationVmwareVirtualMachineAllOf) GetMacAddressOk() (*[]string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *VirtualizationVmwareVirtualMachineAllOf) SetMacAddress(v []string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *VirtualizationVmwareVirtualMachineAllOf) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### SetMacAddressNil

`func (o *VirtualizationVmwareVirtualMachineAllOf) SetMacAddressNil(b bool)`

 SetMacAddressNil sets the value for MacAddress to be an explicit nil

### UnsetMacAddress
`func (o *VirtualizationVmwareVirtualMachineAllOf) UnsetMacAddress()`

UnsetMacAddress ensures that no value is present for MacAddress, not even an explicit nil
### GetMemShares

`func (o *VirtualizationVmwareVirtualMachineAllOf) GetMemShares() VirtualizationVmwareVmMemoryShareInfo`

GetMemShares returns the MemShares field if non-nil, zero value otherwise.

### GetMemSharesOk

`func (o *VirtualizationVmwareVirtualMachineAllOf) GetMemSharesOk() (*VirtualizationVmwareVmMemoryShareInfo, bool)`

GetMemSharesOk returns a tuple with the MemShares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemShares

`func (o *VirtualizationVmwareVirtualMachineAllOf) SetMemShares(v VirtualizationVmwareVmMemoryShareInfo)`

SetMemShares sets MemShares field to given value.

### HasMemShares

`func (o *VirtualizationVmwareVirtualMachineAllOf) HasMemShares() bool`

HasMemShares returns a boolean if a field has been set.

### SetMemSharesNil

`func (o *VirtualizationVmwareVirtualMachineAllOf) SetMemSharesNil(b bool)`

 SetMemSharesNil sets the value for MemShares to be an explicit nil

### UnsetMemShares
`func (o *VirtualizationVmwareVirtualMachineAllOf) UnsetMemShares()`

UnsetMemShares ensures that no value is present for MemShares, not even an explicit nil
### GetMemoryHotAddEnabled

`func (o *VirtualizationVmwareVirtualMachineAllOf) GetMemoryHotAddEnabled() bool`

GetMemoryHotAddEnabled returns the MemoryHotAddEnabled field if non-nil, zero value otherwise.

### GetMemoryHotAddEnabledOk

`func (o *VirtualizationVmwareVirtualMachineAllOf) GetMemoryHotAddEnabledOk() (*bool, bool)`

GetMemoryHotAddEnabledOk returns a tuple with the MemoryHotAddEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryHotAddEnabled

`func (o *VirtualizationVmwareVirtualMachineAllOf) SetMemoryHotAddEnabled(v bool)`

SetMemoryHotAddEnabled sets MemoryHotAddEnabled field to given value.

### HasMemoryHotAddEnabled

`func (o *VirtualizationVmwareVirtualMachineAllOf) HasMemoryHotAddEnabled() bool`

HasMemoryHotAddEnabled returns a boolean if a field has been set.

### GetNetworkCount

`func (o *VirtualizationVmwareVirtualMachineAllOf) GetNetworkCount() int64`

GetNetworkCount returns the NetworkCount field if non-nil, zero value otherwise.

### GetNetworkCountOk

`func (o *VirtualizationVmwareVirtualMachineAllOf) GetNetworkCountOk() (*int64, bool)`

GetNetworkCountOk returns a tuple with the NetworkCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkCount

`func (o *VirtualizationVmwareVirtualMachineAllOf) SetNetworkCount(v int64)`

SetNetworkCount sets NetworkCount field to given value.

### HasNetworkCount

`func (o *VirtualizationVmwareVirtualMachineAllOf) HasNetworkCount() bool`

HasNetworkCount returns a boolean if a field has been set.

### GetPortGroups

`func (o *VirtualizationVmwareVirtualMachineAllOf) GetPortGroups() []string`

GetPortGroups returns the PortGroups field if non-nil, zero value otherwise.

### GetPortGroupsOk

`func (o *VirtualizationVmwareVirtualMachineAllOf) GetPortGroupsOk() (*[]string, bool)`

GetPortGroupsOk returns a tuple with the PortGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortGroups

`func (o *VirtualizationVmwareVirtualMachineAllOf) SetPortGroups(v []string)`

SetPortGroups sets PortGroups field to given value.

### HasPortGroups

`func (o *VirtualizationVmwareVirtualMachineAllOf) HasPortGroups() bool`

HasPortGroups returns a boolean if a field has been set.

### SetPortGroupsNil

`func (o *VirtualizationVmwareVirtualMachineAllOf) SetPortGroupsNil(b bool)`

 SetPortGroupsNil sets the value for PortGroups to be an explicit nil

### UnsetPortGroups
`func (o *VirtualizationVmwareVirtualMachineAllOf) UnsetPortGroups()`

UnsetPortGroups ensures that no value is present for PortGroups, not even an explicit nil
### GetProtectedVm

`func (o *VirtualizationVmwareVirtualMachineAllOf) GetProtectedVm() bool`

GetProtectedVm returns the ProtectedVm field if non-nil, zero value otherwise.

### GetProtectedVmOk

`func (o *VirtualizationVmwareVirtualMachineAllOf) GetProtectedVmOk() (*bool, bool)`

GetProtectedVmOk returns a tuple with the ProtectedVm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectedVm

`func (o *VirtualizationVmwareVirtualMachineAllOf) SetProtectedVm(v bool)`

SetProtectedVm sets ProtectedVm field to given value.

### HasProtectedVm

`func (o *VirtualizationVmwareVirtualMachineAllOf) HasProtectedVm() bool`

HasProtectedVm returns a boolean if a field has been set.

### GetRemoteDisplayInfo

`func (o *VirtualizationVmwareVirtualMachineAllOf) GetRemoteDisplayInfo() VirtualizationVmwareRemoteDisplayInfo`

GetRemoteDisplayInfo returns the RemoteDisplayInfo field if non-nil, zero value otherwise.

### GetRemoteDisplayInfoOk

`func (o *VirtualizationVmwareVirtualMachineAllOf) GetRemoteDisplayInfoOk() (*VirtualizationVmwareRemoteDisplayInfo, bool)`

GetRemoteDisplayInfoOk returns a tuple with the RemoteDisplayInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteDisplayInfo

`func (o *VirtualizationVmwareVirtualMachineAllOf) SetRemoteDisplayInfo(v VirtualizationVmwareRemoteDisplayInfo)`

SetRemoteDisplayInfo sets RemoteDisplayInfo field to given value.

### HasRemoteDisplayInfo

`func (o *VirtualizationVmwareVirtualMachineAllOf) HasRemoteDisplayInfo() bool`

HasRemoteDisplayInfo returns a boolean if a field has been set.

### SetRemoteDisplayInfoNil

`func (o *VirtualizationVmwareVirtualMachineAllOf) SetRemoteDisplayInfoNil(b bool)`

 SetRemoteDisplayInfoNil sets the value for RemoteDisplayInfo to be an explicit nil

### UnsetRemoteDisplayInfo
`func (o *VirtualizationVmwareVirtualMachineAllOf) UnsetRemoteDisplayInfo()`

UnsetRemoteDisplayInfo ensures that no value is present for RemoteDisplayInfo, not even an explicit nil
### GetRemoteDisplayVncEnabled

`func (o *VirtualizationVmwareVirtualMachineAllOf) GetRemoteDisplayVncEnabled() bool`

GetRemoteDisplayVncEnabled returns the RemoteDisplayVncEnabled field if non-nil, zero value otherwise.

### GetRemoteDisplayVncEnabledOk

`func (o *VirtualizationVmwareVirtualMachineAllOf) GetRemoteDisplayVncEnabledOk() (*bool, bool)`

GetRemoteDisplayVncEnabledOk returns a tuple with the RemoteDisplayVncEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteDisplayVncEnabled

`func (o *VirtualizationVmwareVirtualMachineAllOf) SetRemoteDisplayVncEnabled(v bool)`

SetRemoteDisplayVncEnabled sets RemoteDisplayVncEnabled field to given value.

### HasRemoteDisplayVncEnabled

`func (o *VirtualizationVmwareVirtualMachineAllOf) HasRemoteDisplayVncEnabled() bool`

HasRemoteDisplayVncEnabled returns a boolean if a field has been set.

### GetResourcePool

`func (o *VirtualizationVmwareVirtualMachineAllOf) GetResourcePool() string`

GetResourcePool returns the ResourcePool field if non-nil, zero value otherwise.

### GetResourcePoolOk

`func (o *VirtualizationVmwareVirtualMachineAllOf) GetResourcePoolOk() (*string, bool)`

GetResourcePoolOk returns a tuple with the ResourcePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePool

`func (o *VirtualizationVmwareVirtualMachineAllOf) SetResourcePool(v string)`

SetResourcePool sets ResourcePool field to given value.

### HasResourcePool

`func (o *VirtualizationVmwareVirtualMachineAllOf) HasResourcePool() bool`

HasResourcePool returns a boolean if a field has been set.

### GetResourcePoolOwner

`func (o *VirtualizationVmwareVirtualMachineAllOf) GetResourcePoolOwner() string`

GetResourcePoolOwner returns the ResourcePoolOwner field if non-nil, zero value otherwise.

### GetResourcePoolOwnerOk

`func (o *VirtualizationVmwareVirtualMachineAllOf) GetResourcePoolOwnerOk() (*string, bool)`

GetResourcePoolOwnerOk returns a tuple with the ResourcePoolOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolOwner

`func (o *VirtualizationVmwareVirtualMachineAllOf) SetResourcePoolOwner(v string)`

SetResourcePoolOwner sets ResourcePoolOwner field to given value.

### HasResourcePoolOwner

`func (o *VirtualizationVmwareVirtualMachineAllOf) HasResourcePoolOwner() bool`

HasResourcePoolOwner returns a boolean if a field has been set.

### GetResourcePoolParent

`func (o *VirtualizationVmwareVirtualMachineAllOf) GetResourcePoolParent() string`

GetResourcePoolParent returns the ResourcePoolParent field if non-nil, zero value otherwise.

### GetResourcePoolParentOk

`func (o *VirtualizationVmwareVirtualMachineAllOf) GetResourcePoolParentOk() (*string, bool)`

GetResourcePoolParentOk returns a tuple with the ResourcePoolParent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolParent

`func (o *VirtualizationVmwareVirtualMachineAllOf) SetResourcePoolParent(v string)`

SetResourcePoolParent sets ResourcePoolParent field to given value.

### HasResourcePoolParent

`func (o *VirtualizationVmwareVirtualMachineAllOf) HasResourcePoolParent() bool`

HasResourcePoolParent returns a boolean if a field has been set.

### GetToolRunningStatus

`func (o *VirtualizationVmwareVirtualMachineAllOf) GetToolRunningStatus() string`

GetToolRunningStatus returns the ToolRunningStatus field if non-nil, zero value otherwise.

### GetToolRunningStatusOk

`func (o *VirtualizationVmwareVirtualMachineAllOf) GetToolRunningStatusOk() (*string, bool)`

GetToolRunningStatusOk returns a tuple with the ToolRunningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolRunningStatus

`func (o *VirtualizationVmwareVirtualMachineAllOf) SetToolRunningStatus(v string)`

SetToolRunningStatus sets ToolRunningStatus field to given value.

### HasToolRunningStatus

`func (o *VirtualizationVmwareVirtualMachineAllOf) HasToolRunningStatus() bool`

HasToolRunningStatus returns a boolean if a field has been set.

### GetToolsVersion

`func (o *VirtualizationVmwareVirtualMachineAllOf) GetToolsVersion() string`

GetToolsVersion returns the ToolsVersion field if non-nil, zero value otherwise.

### GetToolsVersionOk

`func (o *VirtualizationVmwareVirtualMachineAllOf) GetToolsVersionOk() (*string, bool)`

GetToolsVersionOk returns a tuple with the ToolsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolsVersion

`func (o *VirtualizationVmwareVirtualMachineAllOf) SetToolsVersion(v string)`

SetToolsVersion sets ToolsVersion field to given value.

### HasToolsVersion

`func (o *VirtualizationVmwareVirtualMachineAllOf) HasToolsVersion() bool`

HasToolsVersion returns a boolean if a field has been set.

### GetVirtualDisks

`func (o *VirtualizationVmwareVirtualMachineAllOf) GetVirtualDisks() []int64`

GetVirtualDisks returns the VirtualDisks field if non-nil, zero value otherwise.

### GetVirtualDisksOk

`func (o *VirtualizationVmwareVirtualMachineAllOf) GetVirtualDisksOk() (*[]int64, bool)`

GetVirtualDisksOk returns a tuple with the VirtualDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualDisks

`func (o *VirtualizationVmwareVirtualMachineAllOf) SetVirtualDisks(v []int64)`

SetVirtualDisks sets VirtualDisks field to given value.

### HasVirtualDisks

`func (o *VirtualizationVmwareVirtualMachineAllOf) HasVirtualDisks() bool`

HasVirtualDisks returns a boolean if a field has been set.

### SetVirtualDisksNil

`func (o *VirtualizationVmwareVirtualMachineAllOf) SetVirtualDisksNil(b bool)`

 SetVirtualDisksNil sets the value for VirtualDisks to be an explicit nil

### UnsetVirtualDisks
`func (o *VirtualizationVmwareVirtualMachineAllOf) UnsetVirtualDisks()`

UnsetVirtualDisks ensures that no value is present for VirtualDisks, not even an explicit nil
### GetVirtualNetworkInterfaces

`func (o *VirtualizationVmwareVirtualMachineAllOf) GetVirtualNetworkInterfaces() []int64`

GetVirtualNetworkInterfaces returns the VirtualNetworkInterfaces field if non-nil, zero value otherwise.

### GetVirtualNetworkInterfacesOk

`func (o *VirtualizationVmwareVirtualMachineAllOf) GetVirtualNetworkInterfacesOk() (*[]int64, bool)`

GetVirtualNetworkInterfacesOk returns a tuple with the VirtualNetworkInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualNetworkInterfaces

`func (o *VirtualizationVmwareVirtualMachineAllOf) SetVirtualNetworkInterfaces(v []int64)`

SetVirtualNetworkInterfaces sets VirtualNetworkInterfaces field to given value.

### HasVirtualNetworkInterfaces

`func (o *VirtualizationVmwareVirtualMachineAllOf) HasVirtualNetworkInterfaces() bool`

HasVirtualNetworkInterfaces returns a boolean if a field has been set.

### SetVirtualNetworkInterfacesNil

`func (o *VirtualizationVmwareVirtualMachineAllOf) SetVirtualNetworkInterfacesNil(b bool)`

 SetVirtualNetworkInterfacesNil sets the value for VirtualNetworkInterfaces to be an explicit nil

### UnsetVirtualNetworkInterfaces
`func (o *VirtualizationVmwareVirtualMachineAllOf) UnsetVirtualNetworkInterfaces()`

UnsetVirtualNetworkInterfaces ensures that no value is present for VirtualNetworkInterfaces, not even an explicit nil
### GetVmDiskCount

`func (o *VirtualizationVmwareVirtualMachineAllOf) GetVmDiskCount() int64`

GetVmDiskCount returns the VmDiskCount field if non-nil, zero value otherwise.

### GetVmDiskCountOk

`func (o *VirtualizationVmwareVirtualMachineAllOf) GetVmDiskCountOk() (*int64, bool)`

GetVmDiskCountOk returns a tuple with the VmDiskCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmDiskCount

`func (o *VirtualizationVmwareVirtualMachineAllOf) SetVmDiskCount(v int64)`

SetVmDiskCount sets VmDiskCount field to given value.

### HasVmDiskCount

`func (o *VirtualizationVmwareVirtualMachineAllOf) HasVmDiskCount() bool`

HasVmDiskCount returns a boolean if a field has been set.

### GetVmOverallStatus

`func (o *VirtualizationVmwareVirtualMachineAllOf) GetVmOverallStatus() string`

GetVmOverallStatus returns the VmOverallStatus field if non-nil, zero value otherwise.

### GetVmOverallStatusOk

`func (o *VirtualizationVmwareVirtualMachineAllOf) GetVmOverallStatusOk() (*string, bool)`

GetVmOverallStatusOk returns a tuple with the VmOverallStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmOverallStatus

`func (o *VirtualizationVmwareVirtualMachineAllOf) SetVmOverallStatus(v string)`

SetVmOverallStatus sets VmOverallStatus field to given value.

### HasVmOverallStatus

`func (o *VirtualizationVmwareVirtualMachineAllOf) HasVmOverallStatus() bool`

HasVmOverallStatus returns a boolean if a field has been set.

### GetVmPath

`func (o *VirtualizationVmwareVirtualMachineAllOf) GetVmPath() string`

GetVmPath returns the VmPath field if non-nil, zero value otherwise.

### GetVmPathOk

`func (o *VirtualizationVmwareVirtualMachineAllOf) GetVmPathOk() (*string, bool)`

GetVmPathOk returns a tuple with the VmPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmPath

`func (o *VirtualizationVmwareVirtualMachineAllOf) SetVmPath(v string)`

SetVmPath sets VmPath field to given value.

### HasVmPath

`func (o *VirtualizationVmwareVirtualMachineAllOf) HasVmPath() bool`

HasVmPath returns a boolean if a field has been set.

### GetVmVersion

`func (o *VirtualizationVmwareVirtualMachineAllOf) GetVmVersion() string`

GetVmVersion returns the VmVersion field if non-nil, zero value otherwise.

### GetVmVersionOk

`func (o *VirtualizationVmwareVirtualMachineAllOf) GetVmVersionOk() (*string, bool)`

GetVmVersionOk returns a tuple with the VmVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmVersion

`func (o *VirtualizationVmwareVirtualMachineAllOf) SetVmVersion(v string)`

SetVmVersion sets VmVersion field to given value.

### HasVmVersion

`func (o *VirtualizationVmwareVirtualMachineAllOf) HasVmVersion() bool`

HasVmVersion returns a boolean if a field has been set.

### GetVmVnicCount

`func (o *VirtualizationVmwareVirtualMachineAllOf) GetVmVnicCount() int64`

GetVmVnicCount returns the VmVnicCount field if non-nil, zero value otherwise.

### GetVmVnicCountOk

`func (o *VirtualizationVmwareVirtualMachineAllOf) GetVmVnicCountOk() (*int64, bool)`

GetVmVnicCountOk returns a tuple with the VmVnicCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmVnicCount

`func (o *VirtualizationVmwareVirtualMachineAllOf) SetVmVnicCount(v int64)`

SetVmVnicCount sets VmVnicCount field to given value.

### HasVmVnicCount

`func (o *VirtualizationVmwareVirtualMachineAllOf) HasVmVnicCount() bool`

HasVmVnicCount returns a boolean if a field has been set.

### GetVnicDeviceConfigId

`func (o *VirtualizationVmwareVirtualMachineAllOf) GetVnicDeviceConfigId() string`

GetVnicDeviceConfigId returns the VnicDeviceConfigId field if non-nil, zero value otherwise.

### GetVnicDeviceConfigIdOk

`func (o *VirtualizationVmwareVirtualMachineAllOf) GetVnicDeviceConfigIdOk() (*string, bool)`

GetVnicDeviceConfigIdOk returns a tuple with the VnicDeviceConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVnicDeviceConfigId

`func (o *VirtualizationVmwareVirtualMachineAllOf) SetVnicDeviceConfigId(v string)`

SetVnicDeviceConfigId sets VnicDeviceConfigId field to given value.

### HasVnicDeviceConfigId

`func (o *VirtualizationVmwareVirtualMachineAllOf) HasVnicDeviceConfigId() bool`

HasVnicDeviceConfigId returns a boolean if a field has been set.

### GetCluster

`func (o *VirtualizationVmwareVirtualMachineAllOf) GetCluster() VirtualizationVmwareClusterRelationship`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *VirtualizationVmwareVirtualMachineAllOf) GetClusterOk() (*VirtualizationVmwareClusterRelationship, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *VirtualizationVmwareVirtualMachineAllOf) SetCluster(v VirtualizationVmwareClusterRelationship)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *VirtualizationVmwareVirtualMachineAllOf) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetDatacenter

`func (o *VirtualizationVmwareVirtualMachineAllOf) GetDatacenter() VirtualizationVmwareDatacenterRelationship`

GetDatacenter returns the Datacenter field if non-nil, zero value otherwise.

### GetDatacenterOk

`func (o *VirtualizationVmwareVirtualMachineAllOf) GetDatacenterOk() (*VirtualizationVmwareDatacenterRelationship, bool)`

GetDatacenterOk returns a tuple with the Datacenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenter

`func (o *VirtualizationVmwareVirtualMachineAllOf) SetDatacenter(v VirtualizationVmwareDatacenterRelationship)`

SetDatacenter sets Datacenter field to given value.

### HasDatacenter

`func (o *VirtualizationVmwareVirtualMachineAllOf) HasDatacenter() bool`

HasDatacenter returns a boolean if a field has been set.

### GetDatastores

`func (o *VirtualizationVmwareVirtualMachineAllOf) GetDatastores() []VirtualizationVmwareDatastoreRelationship`

GetDatastores returns the Datastores field if non-nil, zero value otherwise.

### GetDatastoresOk

`func (o *VirtualizationVmwareVirtualMachineAllOf) GetDatastoresOk() (*[]VirtualizationVmwareDatastoreRelationship, bool)`

GetDatastoresOk returns a tuple with the Datastores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastores

`func (o *VirtualizationVmwareVirtualMachineAllOf) SetDatastores(v []VirtualizationVmwareDatastoreRelationship)`

SetDatastores sets Datastores field to given value.

### HasDatastores

`func (o *VirtualizationVmwareVirtualMachineAllOf) HasDatastores() bool`

HasDatastores returns a boolean if a field has been set.

### SetDatastoresNil

`func (o *VirtualizationVmwareVirtualMachineAllOf) SetDatastoresNil(b bool)`

 SetDatastoresNil sets the value for Datastores to be an explicit nil

### UnsetDatastores
`func (o *VirtualizationVmwareVirtualMachineAllOf) UnsetDatastores()`

UnsetDatastores ensures that no value is present for Datastores, not even an explicit nil
### GetHost

`func (o *VirtualizationVmwareVirtualMachineAllOf) GetHost() VirtualizationVmwareHostRelationship`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *VirtualizationVmwareVirtualMachineAllOf) GetHostOk() (*VirtualizationVmwareHostRelationship, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *VirtualizationVmwareVirtualMachineAllOf) SetHost(v VirtualizationVmwareHostRelationship)`

SetHost sets Host field to given value.

### HasHost

`func (o *VirtualizationVmwareVirtualMachineAllOf) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetNetworks

`func (o *VirtualizationVmwareVirtualMachineAllOf) GetNetworks() []VirtualizationBaseNetworkRelationship`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *VirtualizationVmwareVirtualMachineAllOf) GetNetworksOk() (*[]VirtualizationBaseNetworkRelationship, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *VirtualizationVmwareVirtualMachineAllOf) SetNetworks(v []VirtualizationBaseNetworkRelationship)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *VirtualizationVmwareVirtualMachineAllOf) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.

### SetNetworksNil

`func (o *VirtualizationVmwareVirtualMachineAllOf) SetNetworksNil(b bool)`

 SetNetworksNil sets the value for Networks to be an explicit nil

### UnsetNetworks
`func (o *VirtualizationVmwareVirtualMachineAllOf) UnsetNetworks()`

UnsetNetworks ensures that no value is present for Networks, not even an explicit nil
### GetParentFolder

`func (o *VirtualizationVmwareVirtualMachineAllOf) GetParentFolder() VirtualizationVmwareFolderRelationship`

GetParentFolder returns the ParentFolder field if non-nil, zero value otherwise.

### GetParentFolderOk

`func (o *VirtualizationVmwareVirtualMachineAllOf) GetParentFolderOk() (*VirtualizationVmwareFolderRelationship, bool)`

GetParentFolderOk returns a tuple with the ParentFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentFolder

`func (o *VirtualizationVmwareVirtualMachineAllOf) SetParentFolder(v VirtualizationVmwareFolderRelationship)`

SetParentFolder sets ParentFolder field to given value.

### HasParentFolder

`func (o *VirtualizationVmwareVirtualMachineAllOf) HasParentFolder() bool`

HasParentFolder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



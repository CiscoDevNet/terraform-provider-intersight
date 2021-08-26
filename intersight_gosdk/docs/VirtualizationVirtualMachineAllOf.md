# VirtualizationVirtualMachineAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "virtualization.VirtualMachine"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "virtualization.VirtualMachine"]
**Action** | Pointer to **string** | Action to be performed on a virtual machine (Create, PowerState, Migrate, Clone etc). * &#x60;None&#x60; - A place holder for the default value. * &#x60;PowerState&#x60; - Power action is performed on the virtual machine. * &#x60;Migrate&#x60; - The virtual machine will be migrated from existing node to a different node in cluster. The behavior depends on the underlying hypervisor. * &#x60;Create&#x60; - The virtual machine will be created on the specified hypervisor. This action is also useful if the virtual machine creation failed during first POST operation on VirtualMachine managed object. User can set this action to retry the virtual machine creation. * &#x60;Delete&#x60; - The virtual machine will be deleted from the specified hypervisor. User can either set this action or can do a DELETE operation on the VirtualMachine managed object. | [optional] [default to "None"]
**ActionInfo** | Pointer to [**NullableVirtualizationActionInfo**](VirtualizationActionInfo.md) |  | [optional] 
**AffinitySelectors** | Pointer to [**[]InfraMetaData**](InfraMetaData.md) |  | [optional] 
**AntiAffinitySelectors** | Pointer to [**[]InfraMetaData**](InfraMetaData.md) |  | [optional] 
**CloudInitConfig** | Pointer to [**NullableVirtualizationCloudInitConfig**](VirtualizationCloudInitConfig.md) |  | [optional] 
**ClusterEsxi** | Pointer to **string** | Cluster where virtual machine is deployed. | [optional] 
**Cpu** | Pointer to **int64** | Number of vCPUs allocated to virtual machine. | [optional] 
**Discovered** | Pointer to **bool** | Flag to indicate whether the configuration is created from inventory object. | [optional] [readonly] 
**Disk** | Pointer to [**[]VirtualizationVirtualMachineDisk**](VirtualizationVirtualMachineDisk.md) |  | [optional] 
**ForceDelete** | Pointer to **bool** | Normally any virtual machine that is still powered on cannot be deleted. The expected sequence from a user is to first power off the virtual machine and then invoke the delete operation. However, in special circumstances, the owner of the virtual machine may know very well that the virtual machine is no longer needed and just wants to dispose it off. In such situations a delete operation of a virtual machine object is accepted only when this forceDelete attribute is set to true. Under normal circumstances (forceDelete is false), delete operation first confirms that the virtual machine is powered off and then proceeds to delete the virtual machine. | [optional] 
**GuestOs** | Pointer to **string** | Guest operating system running on virtual machine. * &#x60;linux&#x60; - A Linux operating system. * &#x60;windows&#x60; - A Windows operating system. | [optional] [default to "linux"]
**HostEsxi** | Pointer to **string** | Host where virtual machine is deployed. | [optional] 
**HypervisorType** | Pointer to **string** | Identifies the broad product type of the hypervisor but without any version information. It is here to easily identify the type of the virtual machine. There are other entities (Host, Cluster, etc.) that can be indirectly used to determine the hypervisor but a direct attribute makes it easier to work with. * &#x60;ESXi&#x60; - The hypervisor running on the HyperFlex cluster is a Vmware ESXi hypervisor of any version. * &#x60;HyperFlexAp&#x60; - The hypervisor running on the HyperFlex cluster is Cisco HyperFlex Application Platform. * &#x60;Hyper-V&#x60; - The hypervisor running on the HyperFlex cluster is Microsoft Hyper-V. * &#x60;Unknown&#x60; - The hypervisor running on the HyperFlex cluster is not known. | [optional] [readonly] [default to "ESXi"]
**Interfaces** | Pointer to [**[]VirtualizationNetworkInterface**](VirtualizationNetworkInterface.md) |  | [optional] 
**Labels** | Pointer to [**[]InfraMetaData**](InfraMetaData.md) |  | [optional] 
**Memory** | Pointer to **int64** | Virtual machine memory in mebi bytes (one mebibyte, 1MiB, is 1048576 bytes, and 1KiB is 1024 bytes). Input must be a whole number and scientific notation is not acceptable. For example, enter 1730 and not 1.73e03. | [optional] 
**Name** | Pointer to **string** | Virtual machine name that is unique. Hypervisors enforce platform specific limits and character sets. The name length limit, both min and max, vary among hypervisors. Therefore, the basic limits are set here and proper enforcement is done elsewhere. | [optional] 
**PowerState** | Pointer to **string** | Expected power state of virtual machine (PowerOn, PowerOff, Restart). * &#x60;PowerOff&#x60; - The virtual machine will be powered off if it is already not in powered off state. If it is already powered off, no side-effects are expected. * &#x60;PowerOn&#x60; - The virtual machine will be powered on if it is already not in powered on state. If it is already powered on, no side-effects are expected. * &#x60;Suspend&#x60; - The virtual machine will be put into  a suspended state. * &#x60;ShutDownGuestOS&#x60; - The guest operating system is shut down gracefully. * &#x60;RestartGuestOS&#x60; - It can either act as a reset switch and abruptly reset the guest operating system, or it can send a restart signal to the guest operating system so that it shuts down gracefully and restarts. * &#x60;Reset&#x60; - Resets the virtual machine abruptly, with no consideration for work in progress. * &#x60;Restart&#x60; - The virtual machine will be restarted only if it is in powered on state. If it is powered off, it will not be started up. * &#x60;Unknown&#x60; - Power state of the entity is unknown. | [optional] [default to "PowerOff"]
**ProvisionType** | Pointer to **string** | Identifies the provision type to create a new virtual machine. * &#x60;OVA&#x60; - Deploy virtual machine using OVA/F file. * &#x60;Template&#x60; - Provision virtual machine using a template file. * &#x60;Discovered&#x60; - A virtual machine was &#39;discovered&#39; and not created from Intersight. No provisioning information is available. | [optional] [default to "OVA"]
**VmConfig** | Pointer to [**NullableVirtualizationBaseVmConfiguration**](VirtualizationBaseVmConfiguration.md) |  | [optional] 
**Cluster** | Pointer to [**VirtualizationBaseClusterRelationship**](VirtualizationBaseClusterRelationship.md) |  | [optional] 
**Host** | Pointer to [**VirtualizationBaseHostRelationship**](VirtualizationBaseHostRelationship.md) |  | [optional] 
**Inventory** | Pointer to [**VirtualizationBaseVirtualMachineRelationship**](VirtualizationBaseVirtualMachineRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 
**WorkflowInfo** | Pointer to [**WorkflowWorkflowInfoRelationship**](WorkflowWorkflowInfoRelationship.md) |  | [optional] 

## Methods

### NewVirtualizationVirtualMachineAllOf

`func NewVirtualizationVirtualMachineAllOf(classId string, objectType string, ) *VirtualizationVirtualMachineAllOf`

NewVirtualizationVirtualMachineAllOf instantiates a new VirtualizationVirtualMachineAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationVirtualMachineAllOfWithDefaults

`func NewVirtualizationVirtualMachineAllOfWithDefaults() *VirtualizationVirtualMachineAllOf`

NewVirtualizationVirtualMachineAllOfWithDefaults instantiates a new VirtualizationVirtualMachineAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationVirtualMachineAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationVirtualMachineAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationVirtualMachineAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationVirtualMachineAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationVirtualMachineAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationVirtualMachineAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAction

`func (o *VirtualizationVirtualMachineAllOf) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *VirtualizationVirtualMachineAllOf) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *VirtualizationVirtualMachineAllOf) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *VirtualizationVirtualMachineAllOf) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetActionInfo

`func (o *VirtualizationVirtualMachineAllOf) GetActionInfo() VirtualizationActionInfo`

GetActionInfo returns the ActionInfo field if non-nil, zero value otherwise.

### GetActionInfoOk

`func (o *VirtualizationVirtualMachineAllOf) GetActionInfoOk() (*VirtualizationActionInfo, bool)`

GetActionInfoOk returns a tuple with the ActionInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionInfo

`func (o *VirtualizationVirtualMachineAllOf) SetActionInfo(v VirtualizationActionInfo)`

SetActionInfo sets ActionInfo field to given value.

### HasActionInfo

`func (o *VirtualizationVirtualMachineAllOf) HasActionInfo() bool`

HasActionInfo returns a boolean if a field has been set.

### SetActionInfoNil

`func (o *VirtualizationVirtualMachineAllOf) SetActionInfoNil(b bool)`

 SetActionInfoNil sets the value for ActionInfo to be an explicit nil

### UnsetActionInfo
`func (o *VirtualizationVirtualMachineAllOf) UnsetActionInfo()`

UnsetActionInfo ensures that no value is present for ActionInfo, not even an explicit nil
### GetAffinitySelectors

`func (o *VirtualizationVirtualMachineAllOf) GetAffinitySelectors() []InfraMetaData`

GetAffinitySelectors returns the AffinitySelectors field if non-nil, zero value otherwise.

### GetAffinitySelectorsOk

`func (o *VirtualizationVirtualMachineAllOf) GetAffinitySelectorsOk() (*[]InfraMetaData, bool)`

GetAffinitySelectorsOk returns a tuple with the AffinitySelectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffinitySelectors

`func (o *VirtualizationVirtualMachineAllOf) SetAffinitySelectors(v []InfraMetaData)`

SetAffinitySelectors sets AffinitySelectors field to given value.

### HasAffinitySelectors

`func (o *VirtualizationVirtualMachineAllOf) HasAffinitySelectors() bool`

HasAffinitySelectors returns a boolean if a field has been set.

### SetAffinitySelectorsNil

`func (o *VirtualizationVirtualMachineAllOf) SetAffinitySelectorsNil(b bool)`

 SetAffinitySelectorsNil sets the value for AffinitySelectors to be an explicit nil

### UnsetAffinitySelectors
`func (o *VirtualizationVirtualMachineAllOf) UnsetAffinitySelectors()`

UnsetAffinitySelectors ensures that no value is present for AffinitySelectors, not even an explicit nil
### GetAntiAffinitySelectors

`func (o *VirtualizationVirtualMachineAllOf) GetAntiAffinitySelectors() []InfraMetaData`

GetAntiAffinitySelectors returns the AntiAffinitySelectors field if non-nil, zero value otherwise.

### GetAntiAffinitySelectorsOk

`func (o *VirtualizationVirtualMachineAllOf) GetAntiAffinitySelectorsOk() (*[]InfraMetaData, bool)`

GetAntiAffinitySelectorsOk returns a tuple with the AntiAffinitySelectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntiAffinitySelectors

`func (o *VirtualizationVirtualMachineAllOf) SetAntiAffinitySelectors(v []InfraMetaData)`

SetAntiAffinitySelectors sets AntiAffinitySelectors field to given value.

### HasAntiAffinitySelectors

`func (o *VirtualizationVirtualMachineAllOf) HasAntiAffinitySelectors() bool`

HasAntiAffinitySelectors returns a boolean if a field has been set.

### SetAntiAffinitySelectorsNil

`func (o *VirtualizationVirtualMachineAllOf) SetAntiAffinitySelectorsNil(b bool)`

 SetAntiAffinitySelectorsNil sets the value for AntiAffinitySelectors to be an explicit nil

### UnsetAntiAffinitySelectors
`func (o *VirtualizationVirtualMachineAllOf) UnsetAntiAffinitySelectors()`

UnsetAntiAffinitySelectors ensures that no value is present for AntiAffinitySelectors, not even an explicit nil
### GetCloudInitConfig

`func (o *VirtualizationVirtualMachineAllOf) GetCloudInitConfig() VirtualizationCloudInitConfig`

GetCloudInitConfig returns the CloudInitConfig field if non-nil, zero value otherwise.

### GetCloudInitConfigOk

`func (o *VirtualizationVirtualMachineAllOf) GetCloudInitConfigOk() (*VirtualizationCloudInitConfig, bool)`

GetCloudInitConfigOk returns a tuple with the CloudInitConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInitConfig

`func (o *VirtualizationVirtualMachineAllOf) SetCloudInitConfig(v VirtualizationCloudInitConfig)`

SetCloudInitConfig sets CloudInitConfig field to given value.

### HasCloudInitConfig

`func (o *VirtualizationVirtualMachineAllOf) HasCloudInitConfig() bool`

HasCloudInitConfig returns a boolean if a field has been set.

### SetCloudInitConfigNil

`func (o *VirtualizationVirtualMachineAllOf) SetCloudInitConfigNil(b bool)`

 SetCloudInitConfigNil sets the value for CloudInitConfig to be an explicit nil

### UnsetCloudInitConfig
`func (o *VirtualizationVirtualMachineAllOf) UnsetCloudInitConfig()`

UnsetCloudInitConfig ensures that no value is present for CloudInitConfig, not even an explicit nil
### GetClusterEsxi

`func (o *VirtualizationVirtualMachineAllOf) GetClusterEsxi() string`

GetClusterEsxi returns the ClusterEsxi field if non-nil, zero value otherwise.

### GetClusterEsxiOk

`func (o *VirtualizationVirtualMachineAllOf) GetClusterEsxiOk() (*string, bool)`

GetClusterEsxiOk returns a tuple with the ClusterEsxi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterEsxi

`func (o *VirtualizationVirtualMachineAllOf) SetClusterEsxi(v string)`

SetClusterEsxi sets ClusterEsxi field to given value.

### HasClusterEsxi

`func (o *VirtualizationVirtualMachineAllOf) HasClusterEsxi() bool`

HasClusterEsxi returns a boolean if a field has been set.

### GetCpu

`func (o *VirtualizationVirtualMachineAllOf) GetCpu() int64`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *VirtualizationVirtualMachineAllOf) GetCpuOk() (*int64, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *VirtualizationVirtualMachineAllOf) SetCpu(v int64)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *VirtualizationVirtualMachineAllOf) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetDiscovered

`func (o *VirtualizationVirtualMachineAllOf) GetDiscovered() bool`

GetDiscovered returns the Discovered field if non-nil, zero value otherwise.

### GetDiscoveredOk

`func (o *VirtualizationVirtualMachineAllOf) GetDiscoveredOk() (*bool, bool)`

GetDiscoveredOk returns a tuple with the Discovered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscovered

`func (o *VirtualizationVirtualMachineAllOf) SetDiscovered(v bool)`

SetDiscovered sets Discovered field to given value.

### HasDiscovered

`func (o *VirtualizationVirtualMachineAllOf) HasDiscovered() bool`

HasDiscovered returns a boolean if a field has been set.

### GetDisk

`func (o *VirtualizationVirtualMachineAllOf) GetDisk() []VirtualizationVirtualMachineDisk`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *VirtualizationVirtualMachineAllOf) GetDiskOk() (*[]VirtualizationVirtualMachineDisk, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *VirtualizationVirtualMachineAllOf) SetDisk(v []VirtualizationVirtualMachineDisk)`

SetDisk sets Disk field to given value.

### HasDisk

`func (o *VirtualizationVirtualMachineAllOf) HasDisk() bool`

HasDisk returns a boolean if a field has been set.

### SetDiskNil

`func (o *VirtualizationVirtualMachineAllOf) SetDiskNil(b bool)`

 SetDiskNil sets the value for Disk to be an explicit nil

### UnsetDisk
`func (o *VirtualizationVirtualMachineAllOf) UnsetDisk()`

UnsetDisk ensures that no value is present for Disk, not even an explicit nil
### GetForceDelete

`func (o *VirtualizationVirtualMachineAllOf) GetForceDelete() bool`

GetForceDelete returns the ForceDelete field if non-nil, zero value otherwise.

### GetForceDeleteOk

`func (o *VirtualizationVirtualMachineAllOf) GetForceDeleteOk() (*bool, bool)`

GetForceDeleteOk returns a tuple with the ForceDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceDelete

`func (o *VirtualizationVirtualMachineAllOf) SetForceDelete(v bool)`

SetForceDelete sets ForceDelete field to given value.

### HasForceDelete

`func (o *VirtualizationVirtualMachineAllOf) HasForceDelete() bool`

HasForceDelete returns a boolean if a field has been set.

### GetGuestOs

`func (o *VirtualizationVirtualMachineAllOf) GetGuestOs() string`

GetGuestOs returns the GuestOs field if non-nil, zero value otherwise.

### GetGuestOsOk

`func (o *VirtualizationVirtualMachineAllOf) GetGuestOsOk() (*string, bool)`

GetGuestOsOk returns a tuple with the GuestOs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestOs

`func (o *VirtualizationVirtualMachineAllOf) SetGuestOs(v string)`

SetGuestOs sets GuestOs field to given value.

### HasGuestOs

`func (o *VirtualizationVirtualMachineAllOf) HasGuestOs() bool`

HasGuestOs returns a boolean if a field has been set.

### GetHostEsxi

`func (o *VirtualizationVirtualMachineAllOf) GetHostEsxi() string`

GetHostEsxi returns the HostEsxi field if non-nil, zero value otherwise.

### GetHostEsxiOk

`func (o *VirtualizationVirtualMachineAllOf) GetHostEsxiOk() (*string, bool)`

GetHostEsxiOk returns a tuple with the HostEsxi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostEsxi

`func (o *VirtualizationVirtualMachineAllOf) SetHostEsxi(v string)`

SetHostEsxi sets HostEsxi field to given value.

### HasHostEsxi

`func (o *VirtualizationVirtualMachineAllOf) HasHostEsxi() bool`

HasHostEsxi returns a boolean if a field has been set.

### GetHypervisorType

`func (o *VirtualizationVirtualMachineAllOf) GetHypervisorType() string`

GetHypervisorType returns the HypervisorType field if non-nil, zero value otherwise.

### GetHypervisorTypeOk

`func (o *VirtualizationVirtualMachineAllOf) GetHypervisorTypeOk() (*string, bool)`

GetHypervisorTypeOk returns a tuple with the HypervisorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisorType

`func (o *VirtualizationVirtualMachineAllOf) SetHypervisorType(v string)`

SetHypervisorType sets HypervisorType field to given value.

### HasHypervisorType

`func (o *VirtualizationVirtualMachineAllOf) HasHypervisorType() bool`

HasHypervisorType returns a boolean if a field has been set.

### GetInterfaces

`func (o *VirtualizationVirtualMachineAllOf) GetInterfaces() []VirtualizationNetworkInterface`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *VirtualizationVirtualMachineAllOf) GetInterfacesOk() (*[]VirtualizationNetworkInterface, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *VirtualizationVirtualMachineAllOf) SetInterfaces(v []VirtualizationNetworkInterface)`

SetInterfaces sets Interfaces field to given value.

### HasInterfaces

`func (o *VirtualizationVirtualMachineAllOf) HasInterfaces() bool`

HasInterfaces returns a boolean if a field has been set.

### SetInterfacesNil

`func (o *VirtualizationVirtualMachineAllOf) SetInterfacesNil(b bool)`

 SetInterfacesNil sets the value for Interfaces to be an explicit nil

### UnsetInterfaces
`func (o *VirtualizationVirtualMachineAllOf) UnsetInterfaces()`

UnsetInterfaces ensures that no value is present for Interfaces, not even an explicit nil
### GetLabels

`func (o *VirtualizationVirtualMachineAllOf) GetLabels() []InfraMetaData`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *VirtualizationVirtualMachineAllOf) GetLabelsOk() (*[]InfraMetaData, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *VirtualizationVirtualMachineAllOf) SetLabels(v []InfraMetaData)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *VirtualizationVirtualMachineAllOf) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *VirtualizationVirtualMachineAllOf) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *VirtualizationVirtualMachineAllOf) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetMemory

`func (o *VirtualizationVirtualMachineAllOf) GetMemory() int64`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *VirtualizationVirtualMachineAllOf) GetMemoryOk() (*int64, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *VirtualizationVirtualMachineAllOf) SetMemory(v int64)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *VirtualizationVirtualMachineAllOf) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetName

`func (o *VirtualizationVirtualMachineAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VirtualizationVirtualMachineAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VirtualizationVirtualMachineAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VirtualizationVirtualMachineAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPowerState

`func (o *VirtualizationVirtualMachineAllOf) GetPowerState() string`

GetPowerState returns the PowerState field if non-nil, zero value otherwise.

### GetPowerStateOk

`func (o *VirtualizationVirtualMachineAllOf) GetPowerStateOk() (*string, bool)`

GetPowerStateOk returns a tuple with the PowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerState

`func (o *VirtualizationVirtualMachineAllOf) SetPowerState(v string)`

SetPowerState sets PowerState field to given value.

### HasPowerState

`func (o *VirtualizationVirtualMachineAllOf) HasPowerState() bool`

HasPowerState returns a boolean if a field has been set.

### GetProvisionType

`func (o *VirtualizationVirtualMachineAllOf) GetProvisionType() string`

GetProvisionType returns the ProvisionType field if non-nil, zero value otherwise.

### GetProvisionTypeOk

`func (o *VirtualizationVirtualMachineAllOf) GetProvisionTypeOk() (*string, bool)`

GetProvisionTypeOk returns a tuple with the ProvisionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionType

`func (o *VirtualizationVirtualMachineAllOf) SetProvisionType(v string)`

SetProvisionType sets ProvisionType field to given value.

### HasProvisionType

`func (o *VirtualizationVirtualMachineAllOf) HasProvisionType() bool`

HasProvisionType returns a boolean if a field has been set.

### GetVmConfig

`func (o *VirtualizationVirtualMachineAllOf) GetVmConfig() VirtualizationBaseVmConfiguration`

GetVmConfig returns the VmConfig field if non-nil, zero value otherwise.

### GetVmConfigOk

`func (o *VirtualizationVirtualMachineAllOf) GetVmConfigOk() (*VirtualizationBaseVmConfiguration, bool)`

GetVmConfigOk returns a tuple with the VmConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmConfig

`func (o *VirtualizationVirtualMachineAllOf) SetVmConfig(v VirtualizationBaseVmConfiguration)`

SetVmConfig sets VmConfig field to given value.

### HasVmConfig

`func (o *VirtualizationVirtualMachineAllOf) HasVmConfig() bool`

HasVmConfig returns a boolean if a field has been set.

### SetVmConfigNil

`func (o *VirtualizationVirtualMachineAllOf) SetVmConfigNil(b bool)`

 SetVmConfigNil sets the value for VmConfig to be an explicit nil

### UnsetVmConfig
`func (o *VirtualizationVirtualMachineAllOf) UnsetVmConfig()`

UnsetVmConfig ensures that no value is present for VmConfig, not even an explicit nil
### GetCluster

`func (o *VirtualizationVirtualMachineAllOf) GetCluster() VirtualizationBaseClusterRelationship`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *VirtualizationVirtualMachineAllOf) GetClusterOk() (*VirtualizationBaseClusterRelationship, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *VirtualizationVirtualMachineAllOf) SetCluster(v VirtualizationBaseClusterRelationship)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *VirtualizationVirtualMachineAllOf) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetHost

`func (o *VirtualizationVirtualMachineAllOf) GetHost() VirtualizationBaseHostRelationship`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *VirtualizationVirtualMachineAllOf) GetHostOk() (*VirtualizationBaseHostRelationship, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *VirtualizationVirtualMachineAllOf) SetHost(v VirtualizationBaseHostRelationship)`

SetHost sets Host field to given value.

### HasHost

`func (o *VirtualizationVirtualMachineAllOf) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetInventory

`func (o *VirtualizationVirtualMachineAllOf) GetInventory() VirtualizationBaseVirtualMachineRelationship`

GetInventory returns the Inventory field if non-nil, zero value otherwise.

### GetInventoryOk

`func (o *VirtualizationVirtualMachineAllOf) GetInventoryOk() (*VirtualizationBaseVirtualMachineRelationship, bool)`

GetInventoryOk returns a tuple with the Inventory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventory

`func (o *VirtualizationVirtualMachineAllOf) SetInventory(v VirtualizationBaseVirtualMachineRelationship)`

SetInventory sets Inventory field to given value.

### HasInventory

`func (o *VirtualizationVirtualMachineAllOf) HasInventory() bool`

HasInventory returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *VirtualizationVirtualMachineAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *VirtualizationVirtualMachineAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *VirtualizationVirtualMachineAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *VirtualizationVirtualMachineAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### GetWorkflowInfo

`func (o *VirtualizationVirtualMachineAllOf) GetWorkflowInfo() WorkflowWorkflowInfoRelationship`

GetWorkflowInfo returns the WorkflowInfo field if non-nil, zero value otherwise.

### GetWorkflowInfoOk

`func (o *VirtualizationVirtualMachineAllOf) GetWorkflowInfoOk() (*WorkflowWorkflowInfoRelationship, bool)`

GetWorkflowInfoOk returns a tuple with the WorkflowInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowInfo

`func (o *VirtualizationVirtualMachineAllOf) SetWorkflowInfo(v WorkflowWorkflowInfoRelationship)`

SetWorkflowInfo sets WorkflowInfo field to given value.

### HasWorkflowInfo

`func (o *VirtualizationVirtualMachineAllOf) HasWorkflowInfo() bool`

HasWorkflowInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



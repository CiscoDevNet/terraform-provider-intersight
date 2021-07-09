# VirtualizationBaseVirtualMachine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**BootTime** | Pointer to **time.Time** | Time when this VM booted up. | [optional] 
**Capacity** | Pointer to [**NullableInfraHardwareInfo**](infra.HardwareInfo.md) |  | [optional] 
**GuestInfo** | Pointer to [**NullableVirtualizationGuestInfo**](virtualization.GuestInfo.md) |  | [optional] 
**HypervisorType** | Pointer to **string** | Type of hypervisor where the virtual machine is hosted for example ESXi. * &#x60;ESXi&#x60; - The hypervisor running on the HyperFlex cluster is a Vmware ESXi hypervisor of any version. * &#x60;HyperFlexAp&#x60; - The hypervisor running on the HyperFlex cluster is Cisco HyperFlex Application Platform. * &#x60;Hyper-V&#x60; - The hypervisor running on the HyperFlex cluster is Microsoft Hyper-V. * &#x60;Unknown&#x60; - The hypervisor running on the HyperFlex cluster is not known. | [optional] [default to "ESXi"]
**Identity** | Pointer to **string** | The internally generated identity of this VM. This entity is not manipulated by users. It aids in uniquely identifying the virtual machine object. For VMware, this is MOR (managed object reference). | [optional] 
**IpAddress** | Pointer to **[]string** |  | [optional] 
**MemoryCapacity** | Pointer to [**NullableVirtualizationMemoryCapacity**](virtualization.MemoryCapacity.md) |  | [optional] 
**Name** | Pointer to **string** | User-provided name to identify the virtual machine. | [optional] 
**PowerState** | Pointer to **string** | Power state of the virtual machine. * &#x60;Unknown&#x60; - The entity&#39;s power state is unknown. * &#x60;PoweringOn&#x60; - The entity is powering on. * &#x60;PoweredOn&#x60; - The entity is powered on. * &#x60;PoweringOff&#x60; - The entity is powering off. * &#x60;PoweredOff&#x60; - The entity is powered down. * &#x60;StandBy&#x60; - The entity is in standby mode. * &#x60;Paused&#x60; - The entity is in pause state. * &#x60;Rebooting&#x60; - The entity reboot is in progress. * &#x60;&#x60; - The entity&#39;s power state is not available. | [optional] [default to "Unknown"]
**ProcessorCapacity** | Pointer to [**NullableVirtualizationComputeCapacity**](virtualization.ComputeCapacity.md) |  | [optional] 
**Provider** | Pointer to **string** | Cloud platform, where the virtual machine is launched. * &#x60;Unknown&#x60; - Cloud provider is not known. * &#x60;VMwarevSphere&#x60; - Cloud provider named VMware vSphere. * &#x60;AmazonWebServices&#x60; - Cloud provider named Amazon Web Services. * &#x60;MicrosoftAzure&#x60; - Cloud provider named Microsoft Azure. * &#x60;GoogleCloudPlatform&#x60; - Cloud provider named Google Cloud Platform. | [optional] [default to "Unknown"]
**State** | Pointer to **string** | The current state of the virtual machine. For example, starting, stopped, etc. * &#x60;None&#x60; - A place holder for the default value. * &#x60;Creating&#x60; - Virtual machine creation is in progress. * &#x60;Pending&#x60; - The virtual machine is preparing to enter the started state. * &#x60;Starting&#x60; - The virtual machine is starting. * &#x60;Started&#x60; - The virtual machine is running and ready for use. * &#x60;Stopping&#x60; - The virtual machine is preparing to be stopped. * &#x60;Stopped&#x60; - The virtual machine is shut down and cannot be used. The virtual machine can be started again at any time. * &#x60;Pausing&#x60; - The virtual machine is preparing to be paused. * &#x60;Paused&#x60; - The virtual machine enters into paused state due to low free disk space. * &#x60;Suspending&#x60; - The virtual machine is preparing to be suspended. * &#x60;Suspended&#x60; - Virtual machine is in sleep mode.When a virtual machine is suspended, the current state of theoperating system, and applications is saved, and the virtual machine put into a suspended mode. * &#x60;Deleting&#x60; - The virtual machine is preparing to be terminated. * &#x60;Terminated&#x60; - The virtual machine has been permanently deleted and cannot be started. * &#x60;Rebooting&#x60; - The virtual machine reboot is in progress. * &#x60;Error&#x60; - The deployment of virtual machine is failed. | [optional] [default to "None"]
**Uuid** | Pointer to **string** | The uuid of this virtual machine. The uuid is internally generated and not user specified. | [optional] 
**VmCreationTime** | Pointer to **time.Time** | Time when this virtualmachine is created. | [optional] 

## Methods

### NewVirtualizationBaseVirtualMachine

`func NewVirtualizationBaseVirtualMachine(classId string, objectType string, ) *VirtualizationBaseVirtualMachine`

NewVirtualizationBaseVirtualMachine instantiates a new VirtualizationBaseVirtualMachine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationBaseVirtualMachineWithDefaults

`func NewVirtualizationBaseVirtualMachineWithDefaults() *VirtualizationBaseVirtualMachine`

NewVirtualizationBaseVirtualMachineWithDefaults instantiates a new VirtualizationBaseVirtualMachine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationBaseVirtualMachine) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationBaseVirtualMachine) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationBaseVirtualMachine) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationBaseVirtualMachine) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationBaseVirtualMachine) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationBaseVirtualMachine) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBootTime

`func (o *VirtualizationBaseVirtualMachine) GetBootTime() time.Time`

GetBootTime returns the BootTime field if non-nil, zero value otherwise.

### GetBootTimeOk

`func (o *VirtualizationBaseVirtualMachine) GetBootTimeOk() (*time.Time, bool)`

GetBootTimeOk returns a tuple with the BootTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootTime

`func (o *VirtualizationBaseVirtualMachine) SetBootTime(v time.Time)`

SetBootTime sets BootTime field to given value.

### HasBootTime

`func (o *VirtualizationBaseVirtualMachine) HasBootTime() bool`

HasBootTime returns a boolean if a field has been set.

### GetCapacity

`func (o *VirtualizationBaseVirtualMachine) GetCapacity() InfraHardwareInfo`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *VirtualizationBaseVirtualMachine) GetCapacityOk() (*InfraHardwareInfo, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *VirtualizationBaseVirtualMachine) SetCapacity(v InfraHardwareInfo)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *VirtualizationBaseVirtualMachine) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### SetCapacityNil

`func (o *VirtualizationBaseVirtualMachine) SetCapacityNil(b bool)`

 SetCapacityNil sets the value for Capacity to be an explicit nil

### UnsetCapacity
`func (o *VirtualizationBaseVirtualMachine) UnsetCapacity()`

UnsetCapacity ensures that no value is present for Capacity, not even an explicit nil
### GetGuestInfo

`func (o *VirtualizationBaseVirtualMachine) GetGuestInfo() VirtualizationGuestInfo`

GetGuestInfo returns the GuestInfo field if non-nil, zero value otherwise.

### GetGuestInfoOk

`func (o *VirtualizationBaseVirtualMachine) GetGuestInfoOk() (*VirtualizationGuestInfo, bool)`

GetGuestInfoOk returns a tuple with the GuestInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestInfo

`func (o *VirtualizationBaseVirtualMachine) SetGuestInfo(v VirtualizationGuestInfo)`

SetGuestInfo sets GuestInfo field to given value.

### HasGuestInfo

`func (o *VirtualizationBaseVirtualMachine) HasGuestInfo() bool`

HasGuestInfo returns a boolean if a field has been set.

### SetGuestInfoNil

`func (o *VirtualizationBaseVirtualMachine) SetGuestInfoNil(b bool)`

 SetGuestInfoNil sets the value for GuestInfo to be an explicit nil

### UnsetGuestInfo
`func (o *VirtualizationBaseVirtualMachine) UnsetGuestInfo()`

UnsetGuestInfo ensures that no value is present for GuestInfo, not even an explicit nil
### GetHypervisorType

`func (o *VirtualizationBaseVirtualMachine) GetHypervisorType() string`

GetHypervisorType returns the HypervisorType field if non-nil, zero value otherwise.

### GetHypervisorTypeOk

`func (o *VirtualizationBaseVirtualMachine) GetHypervisorTypeOk() (*string, bool)`

GetHypervisorTypeOk returns a tuple with the HypervisorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisorType

`func (o *VirtualizationBaseVirtualMachine) SetHypervisorType(v string)`

SetHypervisorType sets HypervisorType field to given value.

### HasHypervisorType

`func (o *VirtualizationBaseVirtualMachine) HasHypervisorType() bool`

HasHypervisorType returns a boolean if a field has been set.

### GetIdentity

`func (o *VirtualizationBaseVirtualMachine) GetIdentity() string`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *VirtualizationBaseVirtualMachine) GetIdentityOk() (*string, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *VirtualizationBaseVirtualMachine) SetIdentity(v string)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *VirtualizationBaseVirtualMachine) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetIpAddress

`func (o *VirtualizationBaseVirtualMachine) GetIpAddress() []string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *VirtualizationBaseVirtualMachine) GetIpAddressOk() (*[]string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *VirtualizationBaseVirtualMachine) SetIpAddress(v []string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *VirtualizationBaseVirtualMachine) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### SetIpAddressNil

`func (o *VirtualizationBaseVirtualMachine) SetIpAddressNil(b bool)`

 SetIpAddressNil sets the value for IpAddress to be an explicit nil

### UnsetIpAddress
`func (o *VirtualizationBaseVirtualMachine) UnsetIpAddress()`

UnsetIpAddress ensures that no value is present for IpAddress, not even an explicit nil
### GetMemoryCapacity

`func (o *VirtualizationBaseVirtualMachine) GetMemoryCapacity() VirtualizationMemoryCapacity`

GetMemoryCapacity returns the MemoryCapacity field if non-nil, zero value otherwise.

### GetMemoryCapacityOk

`func (o *VirtualizationBaseVirtualMachine) GetMemoryCapacityOk() (*VirtualizationMemoryCapacity, bool)`

GetMemoryCapacityOk returns a tuple with the MemoryCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryCapacity

`func (o *VirtualizationBaseVirtualMachine) SetMemoryCapacity(v VirtualizationMemoryCapacity)`

SetMemoryCapacity sets MemoryCapacity field to given value.

### HasMemoryCapacity

`func (o *VirtualizationBaseVirtualMachine) HasMemoryCapacity() bool`

HasMemoryCapacity returns a boolean if a field has been set.

### SetMemoryCapacityNil

`func (o *VirtualizationBaseVirtualMachine) SetMemoryCapacityNil(b bool)`

 SetMemoryCapacityNil sets the value for MemoryCapacity to be an explicit nil

### UnsetMemoryCapacity
`func (o *VirtualizationBaseVirtualMachine) UnsetMemoryCapacity()`

UnsetMemoryCapacity ensures that no value is present for MemoryCapacity, not even an explicit nil
### GetName

`func (o *VirtualizationBaseVirtualMachine) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VirtualizationBaseVirtualMachine) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VirtualizationBaseVirtualMachine) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VirtualizationBaseVirtualMachine) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPowerState

`func (o *VirtualizationBaseVirtualMachine) GetPowerState() string`

GetPowerState returns the PowerState field if non-nil, zero value otherwise.

### GetPowerStateOk

`func (o *VirtualizationBaseVirtualMachine) GetPowerStateOk() (*string, bool)`

GetPowerStateOk returns a tuple with the PowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerState

`func (o *VirtualizationBaseVirtualMachine) SetPowerState(v string)`

SetPowerState sets PowerState field to given value.

### HasPowerState

`func (o *VirtualizationBaseVirtualMachine) HasPowerState() bool`

HasPowerState returns a boolean if a field has been set.

### GetProcessorCapacity

`func (o *VirtualizationBaseVirtualMachine) GetProcessorCapacity() VirtualizationComputeCapacity`

GetProcessorCapacity returns the ProcessorCapacity field if non-nil, zero value otherwise.

### GetProcessorCapacityOk

`func (o *VirtualizationBaseVirtualMachine) GetProcessorCapacityOk() (*VirtualizationComputeCapacity, bool)`

GetProcessorCapacityOk returns a tuple with the ProcessorCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorCapacity

`func (o *VirtualizationBaseVirtualMachine) SetProcessorCapacity(v VirtualizationComputeCapacity)`

SetProcessorCapacity sets ProcessorCapacity field to given value.

### HasProcessorCapacity

`func (o *VirtualizationBaseVirtualMachine) HasProcessorCapacity() bool`

HasProcessorCapacity returns a boolean if a field has been set.

### SetProcessorCapacityNil

`func (o *VirtualizationBaseVirtualMachine) SetProcessorCapacityNil(b bool)`

 SetProcessorCapacityNil sets the value for ProcessorCapacity to be an explicit nil

### UnsetProcessorCapacity
`func (o *VirtualizationBaseVirtualMachine) UnsetProcessorCapacity()`

UnsetProcessorCapacity ensures that no value is present for ProcessorCapacity, not even an explicit nil
### GetProvider

`func (o *VirtualizationBaseVirtualMachine) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *VirtualizationBaseVirtualMachine) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *VirtualizationBaseVirtualMachine) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *VirtualizationBaseVirtualMachine) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetState

`func (o *VirtualizationBaseVirtualMachine) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *VirtualizationBaseVirtualMachine) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *VirtualizationBaseVirtualMachine) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *VirtualizationBaseVirtualMachine) HasState() bool`

HasState returns a boolean if a field has been set.

### GetUuid

`func (o *VirtualizationBaseVirtualMachine) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *VirtualizationBaseVirtualMachine) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *VirtualizationBaseVirtualMachine) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *VirtualizationBaseVirtualMachine) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetVmCreationTime

`func (o *VirtualizationBaseVirtualMachine) GetVmCreationTime() time.Time`

GetVmCreationTime returns the VmCreationTime field if non-nil, zero value otherwise.

### GetVmCreationTimeOk

`func (o *VirtualizationBaseVirtualMachine) GetVmCreationTimeOk() (*time.Time, bool)`

GetVmCreationTimeOk returns a tuple with the VmCreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmCreationTime

`func (o *VirtualizationBaseVirtualMachine) SetVmCreationTime(v time.Time)`

SetVmCreationTime sets VmCreationTime field to given value.

### HasVmCreationTime

`func (o *VirtualizationBaseVirtualMachine) HasVmCreationTime() bool`

HasVmCreationTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



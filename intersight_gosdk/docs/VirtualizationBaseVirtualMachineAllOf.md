# VirtualizationBaseVirtualMachineAllOf

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

### NewVirtualizationBaseVirtualMachineAllOf

`func NewVirtualizationBaseVirtualMachineAllOf(classId string, objectType string, ) *VirtualizationBaseVirtualMachineAllOf`

NewVirtualizationBaseVirtualMachineAllOf instantiates a new VirtualizationBaseVirtualMachineAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationBaseVirtualMachineAllOfWithDefaults

`func NewVirtualizationBaseVirtualMachineAllOfWithDefaults() *VirtualizationBaseVirtualMachineAllOf`

NewVirtualizationBaseVirtualMachineAllOfWithDefaults instantiates a new VirtualizationBaseVirtualMachineAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationBaseVirtualMachineAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationBaseVirtualMachineAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationBaseVirtualMachineAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationBaseVirtualMachineAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationBaseVirtualMachineAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationBaseVirtualMachineAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBootTime

`func (o *VirtualizationBaseVirtualMachineAllOf) GetBootTime() time.Time`

GetBootTime returns the BootTime field if non-nil, zero value otherwise.

### GetBootTimeOk

`func (o *VirtualizationBaseVirtualMachineAllOf) GetBootTimeOk() (*time.Time, bool)`

GetBootTimeOk returns a tuple with the BootTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootTime

`func (o *VirtualizationBaseVirtualMachineAllOf) SetBootTime(v time.Time)`

SetBootTime sets BootTime field to given value.

### HasBootTime

`func (o *VirtualizationBaseVirtualMachineAllOf) HasBootTime() bool`

HasBootTime returns a boolean if a field has been set.

### GetCapacity

`func (o *VirtualizationBaseVirtualMachineAllOf) GetCapacity() InfraHardwareInfo`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *VirtualizationBaseVirtualMachineAllOf) GetCapacityOk() (*InfraHardwareInfo, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *VirtualizationBaseVirtualMachineAllOf) SetCapacity(v InfraHardwareInfo)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *VirtualizationBaseVirtualMachineAllOf) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### SetCapacityNil

`func (o *VirtualizationBaseVirtualMachineAllOf) SetCapacityNil(b bool)`

 SetCapacityNil sets the value for Capacity to be an explicit nil

### UnsetCapacity
`func (o *VirtualizationBaseVirtualMachineAllOf) UnsetCapacity()`

UnsetCapacity ensures that no value is present for Capacity, not even an explicit nil
### GetGuestInfo

`func (o *VirtualizationBaseVirtualMachineAllOf) GetGuestInfo() VirtualizationGuestInfo`

GetGuestInfo returns the GuestInfo field if non-nil, zero value otherwise.

### GetGuestInfoOk

`func (o *VirtualizationBaseVirtualMachineAllOf) GetGuestInfoOk() (*VirtualizationGuestInfo, bool)`

GetGuestInfoOk returns a tuple with the GuestInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestInfo

`func (o *VirtualizationBaseVirtualMachineAllOf) SetGuestInfo(v VirtualizationGuestInfo)`

SetGuestInfo sets GuestInfo field to given value.

### HasGuestInfo

`func (o *VirtualizationBaseVirtualMachineAllOf) HasGuestInfo() bool`

HasGuestInfo returns a boolean if a field has been set.

### SetGuestInfoNil

`func (o *VirtualizationBaseVirtualMachineAllOf) SetGuestInfoNil(b bool)`

 SetGuestInfoNil sets the value for GuestInfo to be an explicit nil

### UnsetGuestInfo
`func (o *VirtualizationBaseVirtualMachineAllOf) UnsetGuestInfo()`

UnsetGuestInfo ensures that no value is present for GuestInfo, not even an explicit nil
### GetHypervisorType

`func (o *VirtualizationBaseVirtualMachineAllOf) GetHypervisorType() string`

GetHypervisorType returns the HypervisorType field if non-nil, zero value otherwise.

### GetHypervisorTypeOk

`func (o *VirtualizationBaseVirtualMachineAllOf) GetHypervisorTypeOk() (*string, bool)`

GetHypervisorTypeOk returns a tuple with the HypervisorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisorType

`func (o *VirtualizationBaseVirtualMachineAllOf) SetHypervisorType(v string)`

SetHypervisorType sets HypervisorType field to given value.

### HasHypervisorType

`func (o *VirtualizationBaseVirtualMachineAllOf) HasHypervisorType() bool`

HasHypervisorType returns a boolean if a field has been set.

### GetIdentity

`func (o *VirtualizationBaseVirtualMachineAllOf) GetIdentity() string`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *VirtualizationBaseVirtualMachineAllOf) GetIdentityOk() (*string, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *VirtualizationBaseVirtualMachineAllOf) SetIdentity(v string)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *VirtualizationBaseVirtualMachineAllOf) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetIpAddress

`func (o *VirtualizationBaseVirtualMachineAllOf) GetIpAddress() []string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *VirtualizationBaseVirtualMachineAllOf) GetIpAddressOk() (*[]string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *VirtualizationBaseVirtualMachineAllOf) SetIpAddress(v []string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *VirtualizationBaseVirtualMachineAllOf) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### SetIpAddressNil

`func (o *VirtualizationBaseVirtualMachineAllOf) SetIpAddressNil(b bool)`

 SetIpAddressNil sets the value for IpAddress to be an explicit nil

### UnsetIpAddress
`func (o *VirtualizationBaseVirtualMachineAllOf) UnsetIpAddress()`

UnsetIpAddress ensures that no value is present for IpAddress, not even an explicit nil
### GetMemoryCapacity

`func (o *VirtualizationBaseVirtualMachineAllOf) GetMemoryCapacity() VirtualizationMemoryCapacity`

GetMemoryCapacity returns the MemoryCapacity field if non-nil, zero value otherwise.

### GetMemoryCapacityOk

`func (o *VirtualizationBaseVirtualMachineAllOf) GetMemoryCapacityOk() (*VirtualizationMemoryCapacity, bool)`

GetMemoryCapacityOk returns a tuple with the MemoryCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryCapacity

`func (o *VirtualizationBaseVirtualMachineAllOf) SetMemoryCapacity(v VirtualizationMemoryCapacity)`

SetMemoryCapacity sets MemoryCapacity field to given value.

### HasMemoryCapacity

`func (o *VirtualizationBaseVirtualMachineAllOf) HasMemoryCapacity() bool`

HasMemoryCapacity returns a boolean if a field has been set.

### SetMemoryCapacityNil

`func (o *VirtualizationBaseVirtualMachineAllOf) SetMemoryCapacityNil(b bool)`

 SetMemoryCapacityNil sets the value for MemoryCapacity to be an explicit nil

### UnsetMemoryCapacity
`func (o *VirtualizationBaseVirtualMachineAllOf) UnsetMemoryCapacity()`

UnsetMemoryCapacity ensures that no value is present for MemoryCapacity, not even an explicit nil
### GetName

`func (o *VirtualizationBaseVirtualMachineAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VirtualizationBaseVirtualMachineAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VirtualizationBaseVirtualMachineAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VirtualizationBaseVirtualMachineAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPowerState

`func (o *VirtualizationBaseVirtualMachineAllOf) GetPowerState() string`

GetPowerState returns the PowerState field if non-nil, zero value otherwise.

### GetPowerStateOk

`func (o *VirtualizationBaseVirtualMachineAllOf) GetPowerStateOk() (*string, bool)`

GetPowerStateOk returns a tuple with the PowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerState

`func (o *VirtualizationBaseVirtualMachineAllOf) SetPowerState(v string)`

SetPowerState sets PowerState field to given value.

### HasPowerState

`func (o *VirtualizationBaseVirtualMachineAllOf) HasPowerState() bool`

HasPowerState returns a boolean if a field has been set.

### GetProcessorCapacity

`func (o *VirtualizationBaseVirtualMachineAllOf) GetProcessorCapacity() VirtualizationComputeCapacity`

GetProcessorCapacity returns the ProcessorCapacity field if non-nil, zero value otherwise.

### GetProcessorCapacityOk

`func (o *VirtualizationBaseVirtualMachineAllOf) GetProcessorCapacityOk() (*VirtualizationComputeCapacity, bool)`

GetProcessorCapacityOk returns a tuple with the ProcessorCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorCapacity

`func (o *VirtualizationBaseVirtualMachineAllOf) SetProcessorCapacity(v VirtualizationComputeCapacity)`

SetProcessorCapacity sets ProcessorCapacity field to given value.

### HasProcessorCapacity

`func (o *VirtualizationBaseVirtualMachineAllOf) HasProcessorCapacity() bool`

HasProcessorCapacity returns a boolean if a field has been set.

### SetProcessorCapacityNil

`func (o *VirtualizationBaseVirtualMachineAllOf) SetProcessorCapacityNil(b bool)`

 SetProcessorCapacityNil sets the value for ProcessorCapacity to be an explicit nil

### UnsetProcessorCapacity
`func (o *VirtualizationBaseVirtualMachineAllOf) UnsetProcessorCapacity()`

UnsetProcessorCapacity ensures that no value is present for ProcessorCapacity, not even an explicit nil
### GetProvider

`func (o *VirtualizationBaseVirtualMachineAllOf) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *VirtualizationBaseVirtualMachineAllOf) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *VirtualizationBaseVirtualMachineAllOf) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *VirtualizationBaseVirtualMachineAllOf) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetState

`func (o *VirtualizationBaseVirtualMachineAllOf) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *VirtualizationBaseVirtualMachineAllOf) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *VirtualizationBaseVirtualMachineAllOf) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *VirtualizationBaseVirtualMachineAllOf) HasState() bool`

HasState returns a boolean if a field has been set.

### GetUuid

`func (o *VirtualizationBaseVirtualMachineAllOf) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *VirtualizationBaseVirtualMachineAllOf) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *VirtualizationBaseVirtualMachineAllOf) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *VirtualizationBaseVirtualMachineAllOf) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetVmCreationTime

`func (o *VirtualizationBaseVirtualMachineAllOf) GetVmCreationTime() time.Time`

GetVmCreationTime returns the VmCreationTime field if non-nil, zero value otherwise.

### GetVmCreationTimeOk

`func (o *VirtualizationBaseVirtualMachineAllOf) GetVmCreationTimeOk() (*time.Time, bool)`

GetVmCreationTimeOk returns a tuple with the VmCreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmCreationTime

`func (o *VirtualizationBaseVirtualMachineAllOf) SetVmCreationTime(v time.Time)`

SetVmCreationTime sets VmCreationTime field to given value.

### HasVmCreationTime

`func (o *VirtualizationBaseVirtualMachineAllOf) HasVmCreationTime() bool`

HasVmCreationTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



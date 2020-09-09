# VirtualizationBaseVirtualMachineAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capacity** | Pointer to [**InfraHardwareInfo**](infra.HardwareInfo.md) |  | [optional] 
**GuestInfo** | Pointer to [**VirtualizationGuestInfo**](virtualization.GuestInfo.md) |  | [optional] 
**HypervisorType** | Pointer to **string** | Type of hypervisor where the virtual machine is hosted for example ESXi. * &#x60;Unknown&#x60; - The type of the hypervisor is unknown. * &#x60;ESXi&#x60; - A Vmware ESXi hypervisor of any version. * &#x60;HXAP&#x60; - A Cisco HyperFlex Application Platform hypervisor. | [optional] [default to "Unknown"]
**Identity** | Pointer to **string** | The internally generated identity of this VM. This entity is not manipulated by users. It aids in uniquely identifying the virtual machine object. For VMware, this is MOR (managed object reference). | [optional] 
**IpAddress** | Pointer to **[]string** |  | [optional] 
**MemoryCapacity** | Pointer to [**VirtualizationMemoryCapacity**](virtualization.MemoryCapacity.md) |  | [optional] 
**Name** | Pointer to **string** | User-provided name to identify the virtual machine. | [optional] 
**PowerState** | Pointer to **string** | Power state of the virtual machine. * &#x60;Unknown&#x60; - The entity&#39;s power state is unknown. * &#x60;PoweredOn&#x60; - The entity is powered on. * &#x60;PoweredOff&#x60; - The entity is powered down. * &#x60;StandBy&#x60; - The entity is in standby mode. * &#x60;Paused&#x60; - The entity is in pause state. | [optional] [default to "Unknown"]
**ProcessorCapacity** | Pointer to [**VirtualizationComputeCapacity**](virtualization.ComputeCapacity.md) |  | [optional] 
**Uuid** | Pointer to **string** | The uuid of this virtual machine. The uuid is internally generated and not user specified. | [optional] 

## Methods

### NewVirtualizationBaseVirtualMachineAllOf

`func NewVirtualizationBaseVirtualMachineAllOf() *VirtualizationBaseVirtualMachineAllOf`

NewVirtualizationBaseVirtualMachineAllOf instantiates a new VirtualizationBaseVirtualMachineAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationBaseVirtualMachineAllOfWithDefaults

`func NewVirtualizationBaseVirtualMachineAllOfWithDefaults() *VirtualizationBaseVirtualMachineAllOf`

NewVirtualizationBaseVirtualMachineAllOfWithDefaults instantiates a new VirtualizationBaseVirtualMachineAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



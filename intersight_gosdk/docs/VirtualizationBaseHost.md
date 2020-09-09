# VirtualizationBaseHost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CpuInfo** | Pointer to [**VirtualizationCpuInfo**](virtualization.CpuInfo.md) |  | [optional] 
**HardwareInfo** | Pointer to [**InfraHardwareInfo**](infra.HardwareInfo.md) |  | [optional] 
**HypervisorType** | Pointer to **string** | Identifies the broad type of the underlying hypervisor. * &#x60;Unknown&#x60; - The type of the hypervisor is unknown. * &#x60;ESXi&#x60; - A Vmware ESXi hypervisor of any version. * &#x60;HXAP&#x60; - A Cisco HyperFlex Application Platform hypervisor. | [optional] [default to "Unknown"]
**Identity** | Pointer to **string** | The internally generated identity of this host. This entity is not manipulated by users. It aids in uniquely identifying the datacenter object. For VMware, this is an MOR (managed object reference). | [optional] 
**MaintenanceMode** | Pointer to **bool** | Is this host in maintenance mode. Set to true or false. | [optional] 
**MemoryCapacity** | Pointer to [**VirtualizationMemoryCapacity**](virtualization.MemoryCapacity.md) |  | [optional] 
**Model** | Pointer to **string** | Commercial model information about this hardware. | [optional] 
**Name** | Pointer to **string** | Name of this host supplied by user. It is not the identity of the host. The name is subject to user manipulations. | [optional] 
**ProcessorCapacity** | Pointer to [**VirtualizationComputeCapacity**](virtualization.ComputeCapacity.md) |  | [optional] 
**ProductInfo** | Pointer to [**VirtualizationProductInfo**](virtualization.ProductInfo.md) |  | [optional] 
**Serial** | Pointer to **string** | Serial number of this host (internally generated). | [optional] 
**Status** | Pointer to **string** | Host health status, as reported by the hypervisor platform. * &#x60;Unknown&#x60; - Entity status is unknown. * &#x60;Degraded&#x60; - State is degraded, and might impact normal operation of the entity. * &#x60;Critical&#x60; - Entity is in a critical state, impacting operations. * &#x60;Ok&#x60; - Entity status is in a stable state, operating normally. | [optional] [default to "Unknown"]
**UpTime** | Pointer to **string** | The uptime of the host, stored as Duration (from w3c). | [optional] 
**Uuid** | Pointer to **string** | Universally unique identity of this host (example b3d4483b-5560-9342-8309-b486c9236610). Internally generated. | [optional] 
**Vendor** | Pointer to **string** | Commercial vendor details of this hardware. | [optional] 

## Methods

### NewVirtualizationBaseHost

`func NewVirtualizationBaseHost() *VirtualizationBaseHost`

NewVirtualizationBaseHost instantiates a new VirtualizationBaseHost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationBaseHostWithDefaults

`func NewVirtualizationBaseHostWithDefaults() *VirtualizationBaseHost`

NewVirtualizationBaseHostWithDefaults instantiates a new VirtualizationBaseHost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpuInfo

`func (o *VirtualizationBaseHost) GetCpuInfo() VirtualizationCpuInfo`

GetCpuInfo returns the CpuInfo field if non-nil, zero value otherwise.

### GetCpuInfoOk

`func (o *VirtualizationBaseHost) GetCpuInfoOk() (*VirtualizationCpuInfo, bool)`

GetCpuInfoOk returns a tuple with the CpuInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuInfo

`func (o *VirtualizationBaseHost) SetCpuInfo(v VirtualizationCpuInfo)`

SetCpuInfo sets CpuInfo field to given value.

### HasCpuInfo

`func (o *VirtualizationBaseHost) HasCpuInfo() bool`

HasCpuInfo returns a boolean if a field has been set.

### GetHardwareInfo

`func (o *VirtualizationBaseHost) GetHardwareInfo() InfraHardwareInfo`

GetHardwareInfo returns the HardwareInfo field if non-nil, zero value otherwise.

### GetHardwareInfoOk

`func (o *VirtualizationBaseHost) GetHardwareInfoOk() (*InfraHardwareInfo, bool)`

GetHardwareInfoOk returns a tuple with the HardwareInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareInfo

`func (o *VirtualizationBaseHost) SetHardwareInfo(v InfraHardwareInfo)`

SetHardwareInfo sets HardwareInfo field to given value.

### HasHardwareInfo

`func (o *VirtualizationBaseHost) HasHardwareInfo() bool`

HasHardwareInfo returns a boolean if a field has been set.

### GetHypervisorType

`func (o *VirtualizationBaseHost) GetHypervisorType() string`

GetHypervisorType returns the HypervisorType field if non-nil, zero value otherwise.

### GetHypervisorTypeOk

`func (o *VirtualizationBaseHost) GetHypervisorTypeOk() (*string, bool)`

GetHypervisorTypeOk returns a tuple with the HypervisorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisorType

`func (o *VirtualizationBaseHost) SetHypervisorType(v string)`

SetHypervisorType sets HypervisorType field to given value.

### HasHypervisorType

`func (o *VirtualizationBaseHost) HasHypervisorType() bool`

HasHypervisorType returns a boolean if a field has been set.

### GetIdentity

`func (o *VirtualizationBaseHost) GetIdentity() string`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *VirtualizationBaseHost) GetIdentityOk() (*string, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *VirtualizationBaseHost) SetIdentity(v string)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *VirtualizationBaseHost) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetMaintenanceMode

`func (o *VirtualizationBaseHost) GetMaintenanceMode() bool`

GetMaintenanceMode returns the MaintenanceMode field if non-nil, zero value otherwise.

### GetMaintenanceModeOk

`func (o *VirtualizationBaseHost) GetMaintenanceModeOk() (*bool, bool)`

GetMaintenanceModeOk returns a tuple with the MaintenanceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceMode

`func (o *VirtualizationBaseHost) SetMaintenanceMode(v bool)`

SetMaintenanceMode sets MaintenanceMode field to given value.

### HasMaintenanceMode

`func (o *VirtualizationBaseHost) HasMaintenanceMode() bool`

HasMaintenanceMode returns a boolean if a field has been set.

### GetMemoryCapacity

`func (o *VirtualizationBaseHost) GetMemoryCapacity() VirtualizationMemoryCapacity`

GetMemoryCapacity returns the MemoryCapacity field if non-nil, zero value otherwise.

### GetMemoryCapacityOk

`func (o *VirtualizationBaseHost) GetMemoryCapacityOk() (*VirtualizationMemoryCapacity, bool)`

GetMemoryCapacityOk returns a tuple with the MemoryCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryCapacity

`func (o *VirtualizationBaseHost) SetMemoryCapacity(v VirtualizationMemoryCapacity)`

SetMemoryCapacity sets MemoryCapacity field to given value.

### HasMemoryCapacity

`func (o *VirtualizationBaseHost) HasMemoryCapacity() bool`

HasMemoryCapacity returns a boolean if a field has been set.

### GetModel

`func (o *VirtualizationBaseHost) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *VirtualizationBaseHost) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *VirtualizationBaseHost) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *VirtualizationBaseHost) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetName

`func (o *VirtualizationBaseHost) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VirtualizationBaseHost) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VirtualizationBaseHost) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VirtualizationBaseHost) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProcessorCapacity

`func (o *VirtualizationBaseHost) GetProcessorCapacity() VirtualizationComputeCapacity`

GetProcessorCapacity returns the ProcessorCapacity field if non-nil, zero value otherwise.

### GetProcessorCapacityOk

`func (o *VirtualizationBaseHost) GetProcessorCapacityOk() (*VirtualizationComputeCapacity, bool)`

GetProcessorCapacityOk returns a tuple with the ProcessorCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorCapacity

`func (o *VirtualizationBaseHost) SetProcessorCapacity(v VirtualizationComputeCapacity)`

SetProcessorCapacity sets ProcessorCapacity field to given value.

### HasProcessorCapacity

`func (o *VirtualizationBaseHost) HasProcessorCapacity() bool`

HasProcessorCapacity returns a boolean if a field has been set.

### GetProductInfo

`func (o *VirtualizationBaseHost) GetProductInfo() VirtualizationProductInfo`

GetProductInfo returns the ProductInfo field if non-nil, zero value otherwise.

### GetProductInfoOk

`func (o *VirtualizationBaseHost) GetProductInfoOk() (*VirtualizationProductInfo, bool)`

GetProductInfoOk returns a tuple with the ProductInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductInfo

`func (o *VirtualizationBaseHost) SetProductInfo(v VirtualizationProductInfo)`

SetProductInfo sets ProductInfo field to given value.

### HasProductInfo

`func (o *VirtualizationBaseHost) HasProductInfo() bool`

HasProductInfo returns a boolean if a field has been set.

### GetSerial

`func (o *VirtualizationBaseHost) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *VirtualizationBaseHost) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *VirtualizationBaseHost) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *VirtualizationBaseHost) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetStatus

`func (o *VirtualizationBaseHost) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VirtualizationBaseHost) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VirtualizationBaseHost) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VirtualizationBaseHost) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpTime

`func (o *VirtualizationBaseHost) GetUpTime() string`

GetUpTime returns the UpTime field if non-nil, zero value otherwise.

### GetUpTimeOk

`func (o *VirtualizationBaseHost) GetUpTimeOk() (*string, bool)`

GetUpTimeOk returns a tuple with the UpTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpTime

`func (o *VirtualizationBaseHost) SetUpTime(v string)`

SetUpTime sets UpTime field to given value.

### HasUpTime

`func (o *VirtualizationBaseHost) HasUpTime() bool`

HasUpTime returns a boolean if a field has been set.

### GetUuid

`func (o *VirtualizationBaseHost) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *VirtualizationBaseHost) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *VirtualizationBaseHost) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *VirtualizationBaseHost) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetVendor

`func (o *VirtualizationBaseHost) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *VirtualizationBaseHost) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *VirtualizationBaseHost) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *VirtualizationBaseHost) HasVendor() bool`

HasVendor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



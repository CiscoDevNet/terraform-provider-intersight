# VirtualizationBaseHostAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**CpuInfo** | Pointer to [**NullableVirtualizationCpuInfo**](VirtualizationCpuInfo.md) |  | [optional] 
**HardwareInfo** | Pointer to [**NullableInfraHardwareInfo**](InfraHardwareInfo.md) |  | [optional] 
**HypervisorType** | Pointer to **string** | Identifies the broad type of the underlying hypervisor. * &#x60;ESXi&#x60; - The hypervisor running on the HyperFlex cluster is a Vmware ESXi hypervisor of any version. * &#x60;HyperFlexAp&#x60; - The hypervisor running on the HyperFlex cluster is Cisco HyperFlex Application Platform. * &#x60;Hyper-V&#x60; - The hypervisor running on the HyperFlex cluster is Microsoft Hyper-V. * &#x60;Unknown&#x60; - The hypervisor running on the HyperFlex cluster is not known. | [optional] [default to "ESXi"]
**Identity** | Pointer to **string** | The internally generated identity of this host. This entity is not manipulated by users. It aids in uniquely identifying the datacenter object. For VMware, this is an MOR (managed object reference). | [optional] 
**MaintenanceMode** | Pointer to **bool** | Is this host in maintenance mode. Set to true or false. | [optional] 
**MemoryCapacity** | Pointer to [**NullableVirtualizationMemoryCapacity**](VirtualizationMemoryCapacity.md) |  | [optional] 
**Model** | Pointer to **string** | Commercial model information about this hardware. | [optional] 
**Name** | Pointer to **string** | Name of this host supplied by user. It is not the identity of the host. The name is subject to user manipulations. | [optional] 
**ProcessorCapacity** | Pointer to [**NullableVirtualizationComputeCapacity**](VirtualizationComputeCapacity.md) |  | [optional] 
**ProductInfo** | Pointer to [**NullableVirtualizationProductInfo**](VirtualizationProductInfo.md) |  | [optional] 
**Serial** | Pointer to **string** | Serial number of this host (internally generated). | [optional] 
**Status** | Pointer to **string** | Host health status, as reported by the hypervisor platform. * &#x60;Unknown&#x60; - Entity status is unknown. * &#x60;Degraded&#x60; - State is degraded, and might impact normal operation of the entity. * &#x60;Critical&#x60; - Entity is in a critical state, impacting operations. * &#x60;Ok&#x60; - Entity status is in a stable state, operating normally. | [optional] [default to "Unknown"]
**UpTime** | Pointer to **string** | The uptime of the host, stored as Duration (from w3c). | [optional] 
**Uuid** | Pointer to **string** | Universally unique identity of this host (example b3d4483b-5560-9342-8309-b486c9236610). Internally generated. | [optional] 
**Vendor** | Pointer to **string** | Commercial vendor details of this hardware. | [optional] 

## Methods

### NewVirtualizationBaseHostAllOf

`func NewVirtualizationBaseHostAllOf(classId string, objectType string, ) *VirtualizationBaseHostAllOf`

NewVirtualizationBaseHostAllOf instantiates a new VirtualizationBaseHostAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationBaseHostAllOfWithDefaults

`func NewVirtualizationBaseHostAllOfWithDefaults() *VirtualizationBaseHostAllOf`

NewVirtualizationBaseHostAllOfWithDefaults instantiates a new VirtualizationBaseHostAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationBaseHostAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationBaseHostAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationBaseHostAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationBaseHostAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationBaseHostAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationBaseHostAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCpuInfo

`func (o *VirtualizationBaseHostAllOf) GetCpuInfo() VirtualizationCpuInfo`

GetCpuInfo returns the CpuInfo field if non-nil, zero value otherwise.

### GetCpuInfoOk

`func (o *VirtualizationBaseHostAllOf) GetCpuInfoOk() (*VirtualizationCpuInfo, bool)`

GetCpuInfoOk returns a tuple with the CpuInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuInfo

`func (o *VirtualizationBaseHostAllOf) SetCpuInfo(v VirtualizationCpuInfo)`

SetCpuInfo sets CpuInfo field to given value.

### HasCpuInfo

`func (o *VirtualizationBaseHostAllOf) HasCpuInfo() bool`

HasCpuInfo returns a boolean if a field has been set.

### SetCpuInfoNil

`func (o *VirtualizationBaseHostAllOf) SetCpuInfoNil(b bool)`

 SetCpuInfoNil sets the value for CpuInfo to be an explicit nil

### UnsetCpuInfo
`func (o *VirtualizationBaseHostAllOf) UnsetCpuInfo()`

UnsetCpuInfo ensures that no value is present for CpuInfo, not even an explicit nil
### GetHardwareInfo

`func (o *VirtualizationBaseHostAllOf) GetHardwareInfo() InfraHardwareInfo`

GetHardwareInfo returns the HardwareInfo field if non-nil, zero value otherwise.

### GetHardwareInfoOk

`func (o *VirtualizationBaseHostAllOf) GetHardwareInfoOk() (*InfraHardwareInfo, bool)`

GetHardwareInfoOk returns a tuple with the HardwareInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareInfo

`func (o *VirtualizationBaseHostAllOf) SetHardwareInfo(v InfraHardwareInfo)`

SetHardwareInfo sets HardwareInfo field to given value.

### HasHardwareInfo

`func (o *VirtualizationBaseHostAllOf) HasHardwareInfo() bool`

HasHardwareInfo returns a boolean if a field has been set.

### SetHardwareInfoNil

`func (o *VirtualizationBaseHostAllOf) SetHardwareInfoNil(b bool)`

 SetHardwareInfoNil sets the value for HardwareInfo to be an explicit nil

### UnsetHardwareInfo
`func (o *VirtualizationBaseHostAllOf) UnsetHardwareInfo()`

UnsetHardwareInfo ensures that no value is present for HardwareInfo, not even an explicit nil
### GetHypervisorType

`func (o *VirtualizationBaseHostAllOf) GetHypervisorType() string`

GetHypervisorType returns the HypervisorType field if non-nil, zero value otherwise.

### GetHypervisorTypeOk

`func (o *VirtualizationBaseHostAllOf) GetHypervisorTypeOk() (*string, bool)`

GetHypervisorTypeOk returns a tuple with the HypervisorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisorType

`func (o *VirtualizationBaseHostAllOf) SetHypervisorType(v string)`

SetHypervisorType sets HypervisorType field to given value.

### HasHypervisorType

`func (o *VirtualizationBaseHostAllOf) HasHypervisorType() bool`

HasHypervisorType returns a boolean if a field has been set.

### GetIdentity

`func (o *VirtualizationBaseHostAllOf) GetIdentity() string`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *VirtualizationBaseHostAllOf) GetIdentityOk() (*string, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *VirtualizationBaseHostAllOf) SetIdentity(v string)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *VirtualizationBaseHostAllOf) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetMaintenanceMode

`func (o *VirtualizationBaseHostAllOf) GetMaintenanceMode() bool`

GetMaintenanceMode returns the MaintenanceMode field if non-nil, zero value otherwise.

### GetMaintenanceModeOk

`func (o *VirtualizationBaseHostAllOf) GetMaintenanceModeOk() (*bool, bool)`

GetMaintenanceModeOk returns a tuple with the MaintenanceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceMode

`func (o *VirtualizationBaseHostAllOf) SetMaintenanceMode(v bool)`

SetMaintenanceMode sets MaintenanceMode field to given value.

### HasMaintenanceMode

`func (o *VirtualizationBaseHostAllOf) HasMaintenanceMode() bool`

HasMaintenanceMode returns a boolean if a field has been set.

### GetMemoryCapacity

`func (o *VirtualizationBaseHostAllOf) GetMemoryCapacity() VirtualizationMemoryCapacity`

GetMemoryCapacity returns the MemoryCapacity field if non-nil, zero value otherwise.

### GetMemoryCapacityOk

`func (o *VirtualizationBaseHostAllOf) GetMemoryCapacityOk() (*VirtualizationMemoryCapacity, bool)`

GetMemoryCapacityOk returns a tuple with the MemoryCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryCapacity

`func (o *VirtualizationBaseHostAllOf) SetMemoryCapacity(v VirtualizationMemoryCapacity)`

SetMemoryCapacity sets MemoryCapacity field to given value.

### HasMemoryCapacity

`func (o *VirtualizationBaseHostAllOf) HasMemoryCapacity() bool`

HasMemoryCapacity returns a boolean if a field has been set.

### SetMemoryCapacityNil

`func (o *VirtualizationBaseHostAllOf) SetMemoryCapacityNil(b bool)`

 SetMemoryCapacityNil sets the value for MemoryCapacity to be an explicit nil

### UnsetMemoryCapacity
`func (o *VirtualizationBaseHostAllOf) UnsetMemoryCapacity()`

UnsetMemoryCapacity ensures that no value is present for MemoryCapacity, not even an explicit nil
### GetModel

`func (o *VirtualizationBaseHostAllOf) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *VirtualizationBaseHostAllOf) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *VirtualizationBaseHostAllOf) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *VirtualizationBaseHostAllOf) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetName

`func (o *VirtualizationBaseHostAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VirtualizationBaseHostAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VirtualizationBaseHostAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VirtualizationBaseHostAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProcessorCapacity

`func (o *VirtualizationBaseHostAllOf) GetProcessorCapacity() VirtualizationComputeCapacity`

GetProcessorCapacity returns the ProcessorCapacity field if non-nil, zero value otherwise.

### GetProcessorCapacityOk

`func (o *VirtualizationBaseHostAllOf) GetProcessorCapacityOk() (*VirtualizationComputeCapacity, bool)`

GetProcessorCapacityOk returns a tuple with the ProcessorCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorCapacity

`func (o *VirtualizationBaseHostAllOf) SetProcessorCapacity(v VirtualizationComputeCapacity)`

SetProcessorCapacity sets ProcessorCapacity field to given value.

### HasProcessorCapacity

`func (o *VirtualizationBaseHostAllOf) HasProcessorCapacity() bool`

HasProcessorCapacity returns a boolean if a field has been set.

### SetProcessorCapacityNil

`func (o *VirtualizationBaseHostAllOf) SetProcessorCapacityNil(b bool)`

 SetProcessorCapacityNil sets the value for ProcessorCapacity to be an explicit nil

### UnsetProcessorCapacity
`func (o *VirtualizationBaseHostAllOf) UnsetProcessorCapacity()`

UnsetProcessorCapacity ensures that no value is present for ProcessorCapacity, not even an explicit nil
### GetProductInfo

`func (o *VirtualizationBaseHostAllOf) GetProductInfo() VirtualizationProductInfo`

GetProductInfo returns the ProductInfo field if non-nil, zero value otherwise.

### GetProductInfoOk

`func (o *VirtualizationBaseHostAllOf) GetProductInfoOk() (*VirtualizationProductInfo, bool)`

GetProductInfoOk returns a tuple with the ProductInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductInfo

`func (o *VirtualizationBaseHostAllOf) SetProductInfo(v VirtualizationProductInfo)`

SetProductInfo sets ProductInfo field to given value.

### HasProductInfo

`func (o *VirtualizationBaseHostAllOf) HasProductInfo() bool`

HasProductInfo returns a boolean if a field has been set.

### SetProductInfoNil

`func (o *VirtualizationBaseHostAllOf) SetProductInfoNil(b bool)`

 SetProductInfoNil sets the value for ProductInfo to be an explicit nil

### UnsetProductInfo
`func (o *VirtualizationBaseHostAllOf) UnsetProductInfo()`

UnsetProductInfo ensures that no value is present for ProductInfo, not even an explicit nil
### GetSerial

`func (o *VirtualizationBaseHostAllOf) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *VirtualizationBaseHostAllOf) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *VirtualizationBaseHostAllOf) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *VirtualizationBaseHostAllOf) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetStatus

`func (o *VirtualizationBaseHostAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VirtualizationBaseHostAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VirtualizationBaseHostAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VirtualizationBaseHostAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpTime

`func (o *VirtualizationBaseHostAllOf) GetUpTime() string`

GetUpTime returns the UpTime field if non-nil, zero value otherwise.

### GetUpTimeOk

`func (o *VirtualizationBaseHostAllOf) GetUpTimeOk() (*string, bool)`

GetUpTimeOk returns a tuple with the UpTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpTime

`func (o *VirtualizationBaseHostAllOf) SetUpTime(v string)`

SetUpTime sets UpTime field to given value.

### HasUpTime

`func (o *VirtualizationBaseHostAllOf) HasUpTime() bool`

HasUpTime returns a boolean if a field has been set.

### GetUuid

`func (o *VirtualizationBaseHostAllOf) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *VirtualizationBaseHostAllOf) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *VirtualizationBaseHostAllOf) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *VirtualizationBaseHostAllOf) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetVendor

`func (o *VirtualizationBaseHostAllOf) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *VirtualizationBaseHostAllOf) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *VirtualizationBaseHostAllOf) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *VirtualizationBaseHostAllOf) HasVendor() bool`

HasVendor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



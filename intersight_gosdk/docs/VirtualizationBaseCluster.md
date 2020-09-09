# VirtualizationBaseCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HypervisorType** | Pointer to **string** | Identifies the broad type of the underlying hypervisor. * &#x60;Unknown&#x60; - The type of the hypervisor is unknown. * &#x60;ESXi&#x60; - A Vmware ESXi hypervisor of any version. * &#x60;HXAP&#x60; - A Cisco HyperFlex Application Platform hypervisor. | [optional] [default to "Unknown"]
**Identity** | Pointer to **string** | The internally generated identity of this cluster. This entity is not manipulated by users. It aids in uniquely identifying the cluster object. In case of VMware, this is a MOR (managed object reference). | [optional] [readonly] 
**MemoryCapacity** | Pointer to [**VirtualizationMemoryCapacity**](virtualization.MemoryCapacity.md) |  | [optional] 
**Name** | Pointer to **string** | The user-provided name for this cluster to facilitate identification. | [optional] [readonly] 
**ProcessorCapacity** | Pointer to [**VirtualizationComputeCapacity**](virtualization.ComputeCapacity.md) |  | [optional] 
**Status** | Pointer to **string** | Cluster health status as reported by the hypervisor platform. * &#x60;Unknown&#x60; - Entity status is unknown. * &#x60;Degraded&#x60; - State is degraded, and might impact normal operation of the entity. * &#x60;Critical&#x60; - Entity is in a critical state, impacting operations. * &#x60;Ok&#x60; - Entity status is in a stable state, operating normally. | [optional] [readonly] [default to "Unknown"]
**TotalCores** | Pointer to **int64** | Total number of CPU cores in this cluster. It is a cumulative number across all hosts in the cluster. | [optional] 

## Methods

### NewVirtualizationBaseCluster

`func NewVirtualizationBaseCluster() *VirtualizationBaseCluster`

NewVirtualizationBaseCluster instantiates a new VirtualizationBaseCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationBaseClusterWithDefaults

`func NewVirtualizationBaseClusterWithDefaults() *VirtualizationBaseCluster`

NewVirtualizationBaseClusterWithDefaults instantiates a new VirtualizationBaseCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHypervisorType

`func (o *VirtualizationBaseCluster) GetHypervisorType() string`

GetHypervisorType returns the HypervisorType field if non-nil, zero value otherwise.

### GetHypervisorTypeOk

`func (o *VirtualizationBaseCluster) GetHypervisorTypeOk() (*string, bool)`

GetHypervisorTypeOk returns a tuple with the HypervisorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisorType

`func (o *VirtualizationBaseCluster) SetHypervisorType(v string)`

SetHypervisorType sets HypervisorType field to given value.

### HasHypervisorType

`func (o *VirtualizationBaseCluster) HasHypervisorType() bool`

HasHypervisorType returns a boolean if a field has been set.

### GetIdentity

`func (o *VirtualizationBaseCluster) GetIdentity() string`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *VirtualizationBaseCluster) GetIdentityOk() (*string, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *VirtualizationBaseCluster) SetIdentity(v string)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *VirtualizationBaseCluster) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetMemoryCapacity

`func (o *VirtualizationBaseCluster) GetMemoryCapacity() VirtualizationMemoryCapacity`

GetMemoryCapacity returns the MemoryCapacity field if non-nil, zero value otherwise.

### GetMemoryCapacityOk

`func (o *VirtualizationBaseCluster) GetMemoryCapacityOk() (*VirtualizationMemoryCapacity, bool)`

GetMemoryCapacityOk returns a tuple with the MemoryCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryCapacity

`func (o *VirtualizationBaseCluster) SetMemoryCapacity(v VirtualizationMemoryCapacity)`

SetMemoryCapacity sets MemoryCapacity field to given value.

### HasMemoryCapacity

`func (o *VirtualizationBaseCluster) HasMemoryCapacity() bool`

HasMemoryCapacity returns a boolean if a field has been set.

### GetName

`func (o *VirtualizationBaseCluster) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VirtualizationBaseCluster) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VirtualizationBaseCluster) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VirtualizationBaseCluster) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProcessorCapacity

`func (o *VirtualizationBaseCluster) GetProcessorCapacity() VirtualizationComputeCapacity`

GetProcessorCapacity returns the ProcessorCapacity field if non-nil, zero value otherwise.

### GetProcessorCapacityOk

`func (o *VirtualizationBaseCluster) GetProcessorCapacityOk() (*VirtualizationComputeCapacity, bool)`

GetProcessorCapacityOk returns a tuple with the ProcessorCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorCapacity

`func (o *VirtualizationBaseCluster) SetProcessorCapacity(v VirtualizationComputeCapacity)`

SetProcessorCapacity sets ProcessorCapacity field to given value.

### HasProcessorCapacity

`func (o *VirtualizationBaseCluster) HasProcessorCapacity() bool`

HasProcessorCapacity returns a boolean if a field has been set.

### GetStatus

`func (o *VirtualizationBaseCluster) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VirtualizationBaseCluster) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VirtualizationBaseCluster) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VirtualizationBaseCluster) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTotalCores

`func (o *VirtualizationBaseCluster) GetTotalCores() int64`

GetTotalCores returns the TotalCores field if non-nil, zero value otherwise.

### GetTotalCoresOk

`func (o *VirtualizationBaseCluster) GetTotalCoresOk() (*int64, bool)`

GetTotalCoresOk returns a tuple with the TotalCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCores

`func (o *VirtualizationBaseCluster) SetTotalCores(v int64)`

SetTotalCores sets TotalCores field to given value.

### HasTotalCores

`func (o *VirtualizationBaseCluster) HasTotalCores() bool`

HasTotalCores returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



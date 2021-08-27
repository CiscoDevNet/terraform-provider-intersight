# VirtualizationBaseClusterAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**HypervisorType** | Pointer to **string** | Identifies the broad type of the underlying hypervisor. * &#x60;ESXi&#x60; - The hypervisor running on the HyperFlex cluster is a Vmware ESXi hypervisor of any version. * &#x60;HyperFlexAp&#x60; - The hypervisor running on the HyperFlex cluster is Cisco HyperFlex Application Platform. * &#x60;Hyper-V&#x60; - The hypervisor running on the HyperFlex cluster is Microsoft Hyper-V. * &#x60;Unknown&#x60; - The hypervisor running on the HyperFlex cluster is not known. | [optional] [default to "ESXi"]
**Identity** | Pointer to **string** | The internally generated identity of this cluster. This entity is not manipulated by users. It aids in uniquely identifying the cluster object. In case of VMware, this is a MOR (managed object reference). | [optional] [readonly] 
**MemoryCapacity** | Pointer to [**NullableVirtualizationMemoryCapacity**](VirtualizationMemoryCapacity.md) |  | [optional] 
**Name** | Pointer to **string** | The user-provided name for this cluster to facilitate identification. | [optional] [readonly] 
**ProcessorCapacity** | Pointer to [**NullableVirtualizationComputeCapacity**](VirtualizationComputeCapacity.md) |  | [optional] 
**Status** | Pointer to **string** | Cluster health status as reported by the hypervisor platform. * &#x60;Unknown&#x60; - Entity status is unknown. * &#x60;Degraded&#x60; - State is degraded, and might impact normal operation of the entity. * &#x60;Critical&#x60; - Entity is in a critical state, impacting operations. * &#x60;Ok&#x60; - Entity status is in a stable state, operating normally. | [optional] [readonly] [default to "Unknown"]
**TotalCores** | Pointer to **int64** | Total number of CPU cores in this cluster. It is a cumulative number across all hosts in the cluster. | [optional] 

## Methods

### NewVirtualizationBaseClusterAllOf

`func NewVirtualizationBaseClusterAllOf(classId string, objectType string, ) *VirtualizationBaseClusterAllOf`

NewVirtualizationBaseClusterAllOf instantiates a new VirtualizationBaseClusterAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationBaseClusterAllOfWithDefaults

`func NewVirtualizationBaseClusterAllOfWithDefaults() *VirtualizationBaseClusterAllOf`

NewVirtualizationBaseClusterAllOfWithDefaults instantiates a new VirtualizationBaseClusterAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationBaseClusterAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationBaseClusterAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationBaseClusterAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationBaseClusterAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationBaseClusterAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationBaseClusterAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetHypervisorType

`func (o *VirtualizationBaseClusterAllOf) GetHypervisorType() string`

GetHypervisorType returns the HypervisorType field if non-nil, zero value otherwise.

### GetHypervisorTypeOk

`func (o *VirtualizationBaseClusterAllOf) GetHypervisorTypeOk() (*string, bool)`

GetHypervisorTypeOk returns a tuple with the HypervisorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisorType

`func (o *VirtualizationBaseClusterAllOf) SetHypervisorType(v string)`

SetHypervisorType sets HypervisorType field to given value.

### HasHypervisorType

`func (o *VirtualizationBaseClusterAllOf) HasHypervisorType() bool`

HasHypervisorType returns a boolean if a field has been set.

### GetIdentity

`func (o *VirtualizationBaseClusterAllOf) GetIdentity() string`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *VirtualizationBaseClusterAllOf) GetIdentityOk() (*string, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *VirtualizationBaseClusterAllOf) SetIdentity(v string)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *VirtualizationBaseClusterAllOf) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetMemoryCapacity

`func (o *VirtualizationBaseClusterAllOf) GetMemoryCapacity() VirtualizationMemoryCapacity`

GetMemoryCapacity returns the MemoryCapacity field if non-nil, zero value otherwise.

### GetMemoryCapacityOk

`func (o *VirtualizationBaseClusterAllOf) GetMemoryCapacityOk() (*VirtualizationMemoryCapacity, bool)`

GetMemoryCapacityOk returns a tuple with the MemoryCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryCapacity

`func (o *VirtualizationBaseClusterAllOf) SetMemoryCapacity(v VirtualizationMemoryCapacity)`

SetMemoryCapacity sets MemoryCapacity field to given value.

### HasMemoryCapacity

`func (o *VirtualizationBaseClusterAllOf) HasMemoryCapacity() bool`

HasMemoryCapacity returns a boolean if a field has been set.

### SetMemoryCapacityNil

`func (o *VirtualizationBaseClusterAllOf) SetMemoryCapacityNil(b bool)`

 SetMemoryCapacityNil sets the value for MemoryCapacity to be an explicit nil

### UnsetMemoryCapacity
`func (o *VirtualizationBaseClusterAllOf) UnsetMemoryCapacity()`

UnsetMemoryCapacity ensures that no value is present for MemoryCapacity, not even an explicit nil
### GetName

`func (o *VirtualizationBaseClusterAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VirtualizationBaseClusterAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VirtualizationBaseClusterAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VirtualizationBaseClusterAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProcessorCapacity

`func (o *VirtualizationBaseClusterAllOf) GetProcessorCapacity() VirtualizationComputeCapacity`

GetProcessorCapacity returns the ProcessorCapacity field if non-nil, zero value otherwise.

### GetProcessorCapacityOk

`func (o *VirtualizationBaseClusterAllOf) GetProcessorCapacityOk() (*VirtualizationComputeCapacity, bool)`

GetProcessorCapacityOk returns a tuple with the ProcessorCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorCapacity

`func (o *VirtualizationBaseClusterAllOf) SetProcessorCapacity(v VirtualizationComputeCapacity)`

SetProcessorCapacity sets ProcessorCapacity field to given value.

### HasProcessorCapacity

`func (o *VirtualizationBaseClusterAllOf) HasProcessorCapacity() bool`

HasProcessorCapacity returns a boolean if a field has been set.

### SetProcessorCapacityNil

`func (o *VirtualizationBaseClusterAllOf) SetProcessorCapacityNil(b bool)`

 SetProcessorCapacityNil sets the value for ProcessorCapacity to be an explicit nil

### UnsetProcessorCapacity
`func (o *VirtualizationBaseClusterAllOf) UnsetProcessorCapacity()`

UnsetProcessorCapacity ensures that no value is present for ProcessorCapacity, not even an explicit nil
### GetStatus

`func (o *VirtualizationBaseClusterAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VirtualizationBaseClusterAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VirtualizationBaseClusterAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VirtualizationBaseClusterAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTotalCores

`func (o *VirtualizationBaseClusterAllOf) GetTotalCores() int64`

GetTotalCores returns the TotalCores field if non-nil, zero value otherwise.

### GetTotalCoresOk

`func (o *VirtualizationBaseClusterAllOf) GetTotalCoresOk() (*int64, bool)`

GetTotalCoresOk returns a tuple with the TotalCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCores

`func (o *VirtualizationBaseClusterAllOf) SetTotalCores(v int64)`

SetTotalCores sets TotalCores field to given value.

### HasTotalCores

`func (o *VirtualizationBaseClusterAllOf) HasTotalCores() bool`

HasTotalCores returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



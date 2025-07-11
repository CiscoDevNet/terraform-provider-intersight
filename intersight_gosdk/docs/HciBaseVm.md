# HciBaseVm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ClusterExtId** | Pointer to **string** | The unique identifier of the cluster which owns this VM. | [optional] [readonly] 
**Description** | Pointer to **string** | The description of the VM. | [optional] [readonly] 
**GuestOsName** | Pointer to **string** | The guest OS name of the VM. | [optional] [readonly] 
**HypervisorType** | Pointer to **string** | The hypervisor type of the given VM. It could be AHV, ESX etc. | [optional] [readonly] 
**MemorySizeBytes** | Pointer to **int64** | The memory size in bytes of the VM. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the VM reported by the Prism Central. | [optional] [readonly] 
**NodeExtId** | Pointer to **string** | The unique identifier of the node. | [optional] [readonly] 
**NumCoresPerSocket** | Pointer to **int32** | The number of cores per socket of the VM. | [optional] [readonly] 
**NumCpuCores** | Pointer to **int64** | The number of CPU cores of the VM. | [optional] [readonly] 
**PowerState** | Pointer to **string** | The power state of the VM. The possible values are ON, OFF, SUSPENDED (ESXi), PAUSED (AHV), UNDETERMINED. * &#x60;UNDETERMINED&#x60; - The VM power state is currently unknown. * &#x60;OFF&#x60; - The VM&#39;s power state is powered-off. * &#x60;ON&#x60; - The VM&#39;s power state is powered-on. * &#x60;PAUSED&#x60; - The VM&#39;s power state is paused, applicable only to AHV VM. * &#x60;SUSPENDED&#x60; - The VM&#39;s power state is suspended, applicable only to ESXi VM. | [optional] [readonly] [default to "UNDETERMINED"]
**VmExtId** | Pointer to **string** | The unique identifier of the VM. | [optional] [readonly] 
**Cluster** | Pointer to [**NullableHciClusterRelationship**](HciClusterRelationship.md) |  | [optional] 
**Node** | Pointer to [**NullableHciNodeRelationship**](HciNodeRelationship.md) |  | [optional] 

## Methods

### NewHciBaseVm

`func NewHciBaseVm(classId string, objectType string, ) *HciBaseVm`

NewHciBaseVm instantiates a new HciBaseVm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHciBaseVmWithDefaults

`func NewHciBaseVmWithDefaults() *HciBaseVm`

NewHciBaseVmWithDefaults instantiates a new HciBaseVm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HciBaseVm) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HciBaseVm) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HciBaseVm) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HciBaseVm) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HciBaseVm) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HciBaseVm) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetClusterExtId

`func (o *HciBaseVm) GetClusterExtId() string`

GetClusterExtId returns the ClusterExtId field if non-nil, zero value otherwise.

### GetClusterExtIdOk

`func (o *HciBaseVm) GetClusterExtIdOk() (*string, bool)`

GetClusterExtIdOk returns a tuple with the ClusterExtId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterExtId

`func (o *HciBaseVm) SetClusterExtId(v string)`

SetClusterExtId sets ClusterExtId field to given value.

### HasClusterExtId

`func (o *HciBaseVm) HasClusterExtId() bool`

HasClusterExtId returns a boolean if a field has been set.

### GetDescription

`func (o *HciBaseVm) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HciBaseVm) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HciBaseVm) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HciBaseVm) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGuestOsName

`func (o *HciBaseVm) GetGuestOsName() string`

GetGuestOsName returns the GuestOsName field if non-nil, zero value otherwise.

### GetGuestOsNameOk

`func (o *HciBaseVm) GetGuestOsNameOk() (*string, bool)`

GetGuestOsNameOk returns a tuple with the GuestOsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestOsName

`func (o *HciBaseVm) SetGuestOsName(v string)`

SetGuestOsName sets GuestOsName field to given value.

### HasGuestOsName

`func (o *HciBaseVm) HasGuestOsName() bool`

HasGuestOsName returns a boolean if a field has been set.

### GetHypervisorType

`func (o *HciBaseVm) GetHypervisorType() string`

GetHypervisorType returns the HypervisorType field if non-nil, zero value otherwise.

### GetHypervisorTypeOk

`func (o *HciBaseVm) GetHypervisorTypeOk() (*string, bool)`

GetHypervisorTypeOk returns a tuple with the HypervisorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisorType

`func (o *HciBaseVm) SetHypervisorType(v string)`

SetHypervisorType sets HypervisorType field to given value.

### HasHypervisorType

`func (o *HciBaseVm) HasHypervisorType() bool`

HasHypervisorType returns a boolean if a field has been set.

### GetMemorySizeBytes

`func (o *HciBaseVm) GetMemorySizeBytes() int64`

GetMemorySizeBytes returns the MemorySizeBytes field if non-nil, zero value otherwise.

### GetMemorySizeBytesOk

`func (o *HciBaseVm) GetMemorySizeBytesOk() (*int64, bool)`

GetMemorySizeBytesOk returns a tuple with the MemorySizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemorySizeBytes

`func (o *HciBaseVm) SetMemorySizeBytes(v int64)`

SetMemorySizeBytes sets MemorySizeBytes field to given value.

### HasMemorySizeBytes

`func (o *HciBaseVm) HasMemorySizeBytes() bool`

HasMemorySizeBytes returns a boolean if a field has been set.

### GetName

`func (o *HciBaseVm) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HciBaseVm) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HciBaseVm) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HciBaseVm) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNodeExtId

`func (o *HciBaseVm) GetNodeExtId() string`

GetNodeExtId returns the NodeExtId field if non-nil, zero value otherwise.

### GetNodeExtIdOk

`func (o *HciBaseVm) GetNodeExtIdOk() (*string, bool)`

GetNodeExtIdOk returns a tuple with the NodeExtId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeExtId

`func (o *HciBaseVm) SetNodeExtId(v string)`

SetNodeExtId sets NodeExtId field to given value.

### HasNodeExtId

`func (o *HciBaseVm) HasNodeExtId() bool`

HasNodeExtId returns a boolean if a field has been set.

### GetNumCoresPerSocket

`func (o *HciBaseVm) GetNumCoresPerSocket() int32`

GetNumCoresPerSocket returns the NumCoresPerSocket field if non-nil, zero value otherwise.

### GetNumCoresPerSocketOk

`func (o *HciBaseVm) GetNumCoresPerSocketOk() (*int32, bool)`

GetNumCoresPerSocketOk returns a tuple with the NumCoresPerSocket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumCoresPerSocket

`func (o *HciBaseVm) SetNumCoresPerSocket(v int32)`

SetNumCoresPerSocket sets NumCoresPerSocket field to given value.

### HasNumCoresPerSocket

`func (o *HciBaseVm) HasNumCoresPerSocket() bool`

HasNumCoresPerSocket returns a boolean if a field has been set.

### GetNumCpuCores

`func (o *HciBaseVm) GetNumCpuCores() int64`

GetNumCpuCores returns the NumCpuCores field if non-nil, zero value otherwise.

### GetNumCpuCoresOk

`func (o *HciBaseVm) GetNumCpuCoresOk() (*int64, bool)`

GetNumCpuCoresOk returns a tuple with the NumCpuCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumCpuCores

`func (o *HciBaseVm) SetNumCpuCores(v int64)`

SetNumCpuCores sets NumCpuCores field to given value.

### HasNumCpuCores

`func (o *HciBaseVm) HasNumCpuCores() bool`

HasNumCpuCores returns a boolean if a field has been set.

### GetPowerState

`func (o *HciBaseVm) GetPowerState() string`

GetPowerState returns the PowerState field if non-nil, zero value otherwise.

### GetPowerStateOk

`func (o *HciBaseVm) GetPowerStateOk() (*string, bool)`

GetPowerStateOk returns a tuple with the PowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerState

`func (o *HciBaseVm) SetPowerState(v string)`

SetPowerState sets PowerState field to given value.

### HasPowerState

`func (o *HciBaseVm) HasPowerState() bool`

HasPowerState returns a boolean if a field has been set.

### GetVmExtId

`func (o *HciBaseVm) GetVmExtId() string`

GetVmExtId returns the VmExtId field if non-nil, zero value otherwise.

### GetVmExtIdOk

`func (o *HciBaseVm) GetVmExtIdOk() (*string, bool)`

GetVmExtIdOk returns a tuple with the VmExtId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmExtId

`func (o *HciBaseVm) SetVmExtId(v string)`

SetVmExtId sets VmExtId field to given value.

### HasVmExtId

`func (o *HciBaseVm) HasVmExtId() bool`

HasVmExtId returns a boolean if a field has been set.

### GetCluster

`func (o *HciBaseVm) GetCluster() HciClusterRelationship`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *HciBaseVm) GetClusterOk() (*HciClusterRelationship, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *HciBaseVm) SetCluster(v HciClusterRelationship)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *HciBaseVm) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### SetClusterNil

`func (o *HciBaseVm) SetClusterNil(b bool)`

 SetClusterNil sets the value for Cluster to be an explicit nil

### UnsetCluster
`func (o *HciBaseVm) UnsetCluster()`

UnsetCluster ensures that no value is present for Cluster, not even an explicit nil
### GetNode

`func (o *HciBaseVm) GetNode() HciNodeRelationship`

GetNode returns the Node field if non-nil, zero value otherwise.

### GetNodeOk

`func (o *HciBaseVm) GetNodeOk() (*HciNodeRelationship, bool)`

GetNodeOk returns a tuple with the Node field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNode

`func (o *HciBaseVm) SetNode(v HciNodeRelationship)`

SetNode sets Node field to given value.

### HasNode

`func (o *HciBaseVm) HasNode() bool`

HasNode returns a boolean if a field has been set.

### SetNodeNil

`func (o *HciBaseVm) SetNodeNil(b bool)`

 SetNodeNil sets the value for Node to be an explicit nil

### UnsetNode
`func (o *HciBaseVm) UnsetNode()`

UnsetNode ensures that no value is present for Node, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



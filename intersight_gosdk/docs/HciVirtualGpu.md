# HciVirtualGpu

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hci.VirtualGpu"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hci.VirtualGpu"]
**AllocatedVmExtIds** | Pointer to **[]string** |  | [optional] 
**Assignable** | Pointer to **bool** | If the GPU resources is available to be allocated to virtual machines (VMs)  within this cluster. | [optional] [readonly] 
**ClusterExtId** | Pointer to **string** | The unique identifier of the cluster which owns this virtual GPU. | [optional] [readonly] 
**DeviceId** | Pointer to **int64** | The GPU type of the virtual GPU in an integer format. It is similar to DeviceName  which shows the GPU type in string format. | [optional] [readonly] 
**DeviceName** | Pointer to **string** | The GPU type of the virtual GPU in string format. It is similar to DeviceId which shows the GPU type in integer format. | [optional] [readonly] 
**Fraction** | Pointer to **int64** | The fraction of the physical GPU assigned. | [optional] [readonly] 
**FrameBufferSizeBytes** | Pointer to **int64** | The frame buffer size in bytes of the virtual GPU. | [optional] [readonly] 
**GuestDriverVersion** | Pointer to **string** | The guest driver version of the virtual GPU. | [optional] [readonly] 
**IsInUse** | Pointer to **bool** | The in use status of the virtual GPU. | [optional] [readonly] 
**Licenses** | Pointer to **[]string** |  | [optional] 
**MaxInstancesPerVm** | Pointer to **int64** | The maximum instances per VM of the virtual GPU. | [optional] [readonly] 
**MaxResolution** | Pointer to **string** | The maximum resolution of the virtual GPU. | [optional] [readonly] 
**NumaNode** | Pointer to **string** | Each GPU in a system may be physically connected to a specific CPU socket or NUMA node.  The numaNode value specifies which node the GPU is associated with.  In a NUMA system, a computer&#39;s memory is divided into multiple nodes. Each node is a  combination of processors and a portion of the system&#39;s memory. While processors can  access memory from all nodes, they have faster access to the memory in their own node  compared to memory in other nodes. | [optional] [readonly] 
**NumberOfVirtualDisplayHeads** | Pointer to **int64** | The number of virtual display heads of the virtual GPU. | [optional] [readonly] 
**Sbdf** | Pointer to **string** | The SBDF address of the virtual GPU. | [optional] [readonly] 
**Type** | Pointer to **string** | The type of the irtual GPU. Possible values are PASSTHROUGH_GRAPHICS, PASSTHROUGH_COMPUTE, VIRTUAL. | [optional] [readonly] 
**Vendor** | Pointer to **string** | The vendor name of the virtual GPU. | [optional] [readonly] 
**VirtualGpuExtId** | Pointer to **string** | The unique identifier of the virtual GPU. | [optional] [readonly] 
**AllocatedVms** | Pointer to [**[]HciBaseVmRelationship**](HciBaseVmRelationship.md) | An array of relationships to hciBaseVm resources. | [optional] [readonly] 
**Cluster** | Pointer to [**NullableHciClusterRelationship**](HciClusterRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewHciVirtualGpu

`func NewHciVirtualGpu(classId string, objectType string, ) *HciVirtualGpu`

NewHciVirtualGpu instantiates a new HciVirtualGpu object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHciVirtualGpuWithDefaults

`func NewHciVirtualGpuWithDefaults() *HciVirtualGpu`

NewHciVirtualGpuWithDefaults instantiates a new HciVirtualGpu object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HciVirtualGpu) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HciVirtualGpu) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HciVirtualGpu) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HciVirtualGpu) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HciVirtualGpu) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HciVirtualGpu) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAllocatedVmExtIds

`func (o *HciVirtualGpu) GetAllocatedVmExtIds() []string`

GetAllocatedVmExtIds returns the AllocatedVmExtIds field if non-nil, zero value otherwise.

### GetAllocatedVmExtIdsOk

`func (o *HciVirtualGpu) GetAllocatedVmExtIdsOk() (*[]string, bool)`

GetAllocatedVmExtIdsOk returns a tuple with the AllocatedVmExtIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedVmExtIds

`func (o *HciVirtualGpu) SetAllocatedVmExtIds(v []string)`

SetAllocatedVmExtIds sets AllocatedVmExtIds field to given value.

### HasAllocatedVmExtIds

`func (o *HciVirtualGpu) HasAllocatedVmExtIds() bool`

HasAllocatedVmExtIds returns a boolean if a field has been set.

### SetAllocatedVmExtIdsNil

`func (o *HciVirtualGpu) SetAllocatedVmExtIdsNil(b bool)`

 SetAllocatedVmExtIdsNil sets the value for AllocatedVmExtIds to be an explicit nil

### UnsetAllocatedVmExtIds
`func (o *HciVirtualGpu) UnsetAllocatedVmExtIds()`

UnsetAllocatedVmExtIds ensures that no value is present for AllocatedVmExtIds, not even an explicit nil
### GetAssignable

`func (o *HciVirtualGpu) GetAssignable() bool`

GetAssignable returns the Assignable field if non-nil, zero value otherwise.

### GetAssignableOk

`func (o *HciVirtualGpu) GetAssignableOk() (*bool, bool)`

GetAssignableOk returns a tuple with the Assignable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignable

`func (o *HciVirtualGpu) SetAssignable(v bool)`

SetAssignable sets Assignable field to given value.

### HasAssignable

`func (o *HciVirtualGpu) HasAssignable() bool`

HasAssignable returns a boolean if a field has been set.

### GetClusterExtId

`func (o *HciVirtualGpu) GetClusterExtId() string`

GetClusterExtId returns the ClusterExtId field if non-nil, zero value otherwise.

### GetClusterExtIdOk

`func (o *HciVirtualGpu) GetClusterExtIdOk() (*string, bool)`

GetClusterExtIdOk returns a tuple with the ClusterExtId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterExtId

`func (o *HciVirtualGpu) SetClusterExtId(v string)`

SetClusterExtId sets ClusterExtId field to given value.

### HasClusterExtId

`func (o *HciVirtualGpu) HasClusterExtId() bool`

HasClusterExtId returns a boolean if a field has been set.

### GetDeviceId

`func (o *HciVirtualGpu) GetDeviceId() int64`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *HciVirtualGpu) GetDeviceIdOk() (*int64, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *HciVirtualGpu) SetDeviceId(v int64)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *HciVirtualGpu) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetDeviceName

`func (o *HciVirtualGpu) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *HciVirtualGpu) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *HciVirtualGpu) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *HciVirtualGpu) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetFraction

`func (o *HciVirtualGpu) GetFraction() int64`

GetFraction returns the Fraction field if non-nil, zero value otherwise.

### GetFractionOk

`func (o *HciVirtualGpu) GetFractionOk() (*int64, bool)`

GetFractionOk returns a tuple with the Fraction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFraction

`func (o *HciVirtualGpu) SetFraction(v int64)`

SetFraction sets Fraction field to given value.

### HasFraction

`func (o *HciVirtualGpu) HasFraction() bool`

HasFraction returns a boolean if a field has been set.

### GetFrameBufferSizeBytes

`func (o *HciVirtualGpu) GetFrameBufferSizeBytes() int64`

GetFrameBufferSizeBytes returns the FrameBufferSizeBytes field if non-nil, zero value otherwise.

### GetFrameBufferSizeBytesOk

`func (o *HciVirtualGpu) GetFrameBufferSizeBytesOk() (*int64, bool)`

GetFrameBufferSizeBytesOk returns a tuple with the FrameBufferSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrameBufferSizeBytes

`func (o *HciVirtualGpu) SetFrameBufferSizeBytes(v int64)`

SetFrameBufferSizeBytes sets FrameBufferSizeBytes field to given value.

### HasFrameBufferSizeBytes

`func (o *HciVirtualGpu) HasFrameBufferSizeBytes() bool`

HasFrameBufferSizeBytes returns a boolean if a field has been set.

### GetGuestDriverVersion

`func (o *HciVirtualGpu) GetGuestDriverVersion() string`

GetGuestDriverVersion returns the GuestDriverVersion field if non-nil, zero value otherwise.

### GetGuestDriverVersionOk

`func (o *HciVirtualGpu) GetGuestDriverVersionOk() (*string, bool)`

GetGuestDriverVersionOk returns a tuple with the GuestDriverVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestDriverVersion

`func (o *HciVirtualGpu) SetGuestDriverVersion(v string)`

SetGuestDriverVersion sets GuestDriverVersion field to given value.

### HasGuestDriverVersion

`func (o *HciVirtualGpu) HasGuestDriverVersion() bool`

HasGuestDriverVersion returns a boolean if a field has been set.

### GetIsInUse

`func (o *HciVirtualGpu) GetIsInUse() bool`

GetIsInUse returns the IsInUse field if non-nil, zero value otherwise.

### GetIsInUseOk

`func (o *HciVirtualGpu) GetIsInUseOk() (*bool, bool)`

GetIsInUseOk returns a tuple with the IsInUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInUse

`func (o *HciVirtualGpu) SetIsInUse(v bool)`

SetIsInUse sets IsInUse field to given value.

### HasIsInUse

`func (o *HciVirtualGpu) HasIsInUse() bool`

HasIsInUse returns a boolean if a field has been set.

### GetLicenses

`func (o *HciVirtualGpu) GetLicenses() []string`

GetLicenses returns the Licenses field if non-nil, zero value otherwise.

### GetLicensesOk

`func (o *HciVirtualGpu) GetLicensesOk() (*[]string, bool)`

GetLicensesOk returns a tuple with the Licenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenses

`func (o *HciVirtualGpu) SetLicenses(v []string)`

SetLicenses sets Licenses field to given value.

### HasLicenses

`func (o *HciVirtualGpu) HasLicenses() bool`

HasLicenses returns a boolean if a field has been set.

### SetLicensesNil

`func (o *HciVirtualGpu) SetLicensesNil(b bool)`

 SetLicensesNil sets the value for Licenses to be an explicit nil

### UnsetLicenses
`func (o *HciVirtualGpu) UnsetLicenses()`

UnsetLicenses ensures that no value is present for Licenses, not even an explicit nil
### GetMaxInstancesPerVm

`func (o *HciVirtualGpu) GetMaxInstancesPerVm() int64`

GetMaxInstancesPerVm returns the MaxInstancesPerVm field if non-nil, zero value otherwise.

### GetMaxInstancesPerVmOk

`func (o *HciVirtualGpu) GetMaxInstancesPerVmOk() (*int64, bool)`

GetMaxInstancesPerVmOk returns a tuple with the MaxInstancesPerVm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxInstancesPerVm

`func (o *HciVirtualGpu) SetMaxInstancesPerVm(v int64)`

SetMaxInstancesPerVm sets MaxInstancesPerVm field to given value.

### HasMaxInstancesPerVm

`func (o *HciVirtualGpu) HasMaxInstancesPerVm() bool`

HasMaxInstancesPerVm returns a boolean if a field has been set.

### GetMaxResolution

`func (o *HciVirtualGpu) GetMaxResolution() string`

GetMaxResolution returns the MaxResolution field if non-nil, zero value otherwise.

### GetMaxResolutionOk

`func (o *HciVirtualGpu) GetMaxResolutionOk() (*string, bool)`

GetMaxResolutionOk returns a tuple with the MaxResolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxResolution

`func (o *HciVirtualGpu) SetMaxResolution(v string)`

SetMaxResolution sets MaxResolution field to given value.

### HasMaxResolution

`func (o *HciVirtualGpu) HasMaxResolution() bool`

HasMaxResolution returns a boolean if a field has been set.

### GetNumaNode

`func (o *HciVirtualGpu) GetNumaNode() string`

GetNumaNode returns the NumaNode field if non-nil, zero value otherwise.

### GetNumaNodeOk

`func (o *HciVirtualGpu) GetNumaNodeOk() (*string, bool)`

GetNumaNodeOk returns a tuple with the NumaNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumaNode

`func (o *HciVirtualGpu) SetNumaNode(v string)`

SetNumaNode sets NumaNode field to given value.

### HasNumaNode

`func (o *HciVirtualGpu) HasNumaNode() bool`

HasNumaNode returns a boolean if a field has been set.

### GetNumberOfVirtualDisplayHeads

`func (o *HciVirtualGpu) GetNumberOfVirtualDisplayHeads() int64`

GetNumberOfVirtualDisplayHeads returns the NumberOfVirtualDisplayHeads field if non-nil, zero value otherwise.

### GetNumberOfVirtualDisplayHeadsOk

`func (o *HciVirtualGpu) GetNumberOfVirtualDisplayHeadsOk() (*int64, bool)`

GetNumberOfVirtualDisplayHeadsOk returns a tuple with the NumberOfVirtualDisplayHeads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfVirtualDisplayHeads

`func (o *HciVirtualGpu) SetNumberOfVirtualDisplayHeads(v int64)`

SetNumberOfVirtualDisplayHeads sets NumberOfVirtualDisplayHeads field to given value.

### HasNumberOfVirtualDisplayHeads

`func (o *HciVirtualGpu) HasNumberOfVirtualDisplayHeads() bool`

HasNumberOfVirtualDisplayHeads returns a boolean if a field has been set.

### GetSbdf

`func (o *HciVirtualGpu) GetSbdf() string`

GetSbdf returns the Sbdf field if non-nil, zero value otherwise.

### GetSbdfOk

`func (o *HciVirtualGpu) GetSbdfOk() (*string, bool)`

GetSbdfOk returns a tuple with the Sbdf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSbdf

`func (o *HciVirtualGpu) SetSbdf(v string)`

SetSbdf sets Sbdf field to given value.

### HasSbdf

`func (o *HciVirtualGpu) HasSbdf() bool`

HasSbdf returns a boolean if a field has been set.

### GetType

`func (o *HciVirtualGpu) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HciVirtualGpu) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HciVirtualGpu) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *HciVirtualGpu) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVendor

`func (o *HciVirtualGpu) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *HciVirtualGpu) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *HciVirtualGpu) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *HciVirtualGpu) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetVirtualGpuExtId

`func (o *HciVirtualGpu) GetVirtualGpuExtId() string`

GetVirtualGpuExtId returns the VirtualGpuExtId field if non-nil, zero value otherwise.

### GetVirtualGpuExtIdOk

`func (o *HciVirtualGpu) GetVirtualGpuExtIdOk() (*string, bool)`

GetVirtualGpuExtIdOk returns a tuple with the VirtualGpuExtId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualGpuExtId

`func (o *HciVirtualGpu) SetVirtualGpuExtId(v string)`

SetVirtualGpuExtId sets VirtualGpuExtId field to given value.

### HasVirtualGpuExtId

`func (o *HciVirtualGpu) HasVirtualGpuExtId() bool`

HasVirtualGpuExtId returns a boolean if a field has been set.

### GetAllocatedVms

`func (o *HciVirtualGpu) GetAllocatedVms() []HciBaseVmRelationship`

GetAllocatedVms returns the AllocatedVms field if non-nil, zero value otherwise.

### GetAllocatedVmsOk

`func (o *HciVirtualGpu) GetAllocatedVmsOk() (*[]HciBaseVmRelationship, bool)`

GetAllocatedVmsOk returns a tuple with the AllocatedVms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedVms

`func (o *HciVirtualGpu) SetAllocatedVms(v []HciBaseVmRelationship)`

SetAllocatedVms sets AllocatedVms field to given value.

### HasAllocatedVms

`func (o *HciVirtualGpu) HasAllocatedVms() bool`

HasAllocatedVms returns a boolean if a field has been set.

### SetAllocatedVmsNil

`func (o *HciVirtualGpu) SetAllocatedVmsNil(b bool)`

 SetAllocatedVmsNil sets the value for AllocatedVms to be an explicit nil

### UnsetAllocatedVms
`func (o *HciVirtualGpu) UnsetAllocatedVms()`

UnsetAllocatedVms ensures that no value is present for AllocatedVms, not even an explicit nil
### GetCluster

`func (o *HciVirtualGpu) GetCluster() HciClusterRelationship`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *HciVirtualGpu) GetClusterOk() (*HciClusterRelationship, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *HciVirtualGpu) SetCluster(v HciClusterRelationship)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *HciVirtualGpu) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### SetClusterNil

`func (o *HciVirtualGpu) SetClusterNil(b bool)`

 SetClusterNil sets the value for Cluster to be an explicit nil

### UnsetCluster
`func (o *HciVirtualGpu) UnsetCluster()`

UnsetCluster ensures that no value is present for Cluster, not even an explicit nil
### GetRegisteredDevice

`func (o *HciVirtualGpu) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *HciVirtualGpu) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *HciVirtualGpu) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *HciVirtualGpu) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *HciVirtualGpu) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *HciVirtualGpu) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



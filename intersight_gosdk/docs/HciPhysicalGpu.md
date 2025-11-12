# HciPhysicalGpu

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hci.PhysicalGpu"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hci.PhysicalGpu"]
**AllocatedVmExtIds** | Pointer to **[]string** |  | [optional] 
**Assignable** | Pointer to **bool** | If the GPU resources is available to be allocated to virtual machines (VMs)  within this cluster. | [optional] [readonly] 
**ClusterExtId** | Pointer to **string** | The unique identifier of the cluster which owns this physical GPU. | [optional] [readonly] 
**DeviceId** | Pointer to **int32** | The GPU type of the physical GPU in an integer format. It is similar to DeviceName which shows the GPU type in string format. | [optional] [readonly] 
**DeviceName** | Pointer to **string** | The GPU type of the physical GPU in string format. It is similar to DeviceId which shows the GPU type in integer format. | [optional] [readonly] 
**FrameBufferSizeBytes** | Pointer to **int64** | The frame buffer size in bytes of the physical GPU. | [optional] [readonly] 
**IsInUse** | Pointer to **bool** | The in use status of the physical GPU. | [optional] [readonly] 
**Mode** | Pointer to **string** | The mode of the physical GPU. Possible values are - UNUSED, USED_FOR_PASSTHROUGH, USED_FOR_VIRTUAL. | [optional] [readonly] 
**NumaNode** | Pointer to **string** | Each GPU in a system may be physically connected to a specific CPU socket or NUMA node.  The numaNode value specifies which node the GPU is associated with.  In a NUMA system, a computer&#39;s memory is divided into multiple nodes. Each node is a  combination of processors and a portion of the system&#39;s memory. While processors can  access memory from all nodes, they have faster access to the memory in their own node  compared to memory in other nodes. | [optional] [readonly] 
**PhysicalGpuExtId** | Pointer to **string** | The unique identifier of the physical GPU. | [optional] [readonly] 
**Sbdf** | Pointer to **string** | The SBDF address of the physical GPU. | [optional] [readonly] 
**Type** | Pointer to **string** | The type of the physical GPU. Possible values are PASSTHROUGH_GRAPHICS, PASSTHROUGH_COMPUTE, VIRTUAL. | [optional] [readonly] 
**Vendor** | Pointer to **string** | The vendor name of the physical GPU. | [optional] [readonly] 
**AllocatedVms** | Pointer to [**[]HciBaseVmRelationship**](HciBaseVmRelationship.md) | An array of relationships to hciBaseVm resources. | [optional] [readonly] 
**Cluster** | Pointer to [**NullableHciClusterRelationship**](HciClusterRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewHciPhysicalGpu

`func NewHciPhysicalGpu(classId string, objectType string, ) *HciPhysicalGpu`

NewHciPhysicalGpu instantiates a new HciPhysicalGpu object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHciPhysicalGpuWithDefaults

`func NewHciPhysicalGpuWithDefaults() *HciPhysicalGpu`

NewHciPhysicalGpuWithDefaults instantiates a new HciPhysicalGpu object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HciPhysicalGpu) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HciPhysicalGpu) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HciPhysicalGpu) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HciPhysicalGpu) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HciPhysicalGpu) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HciPhysicalGpu) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAllocatedVmExtIds

`func (o *HciPhysicalGpu) GetAllocatedVmExtIds() []string`

GetAllocatedVmExtIds returns the AllocatedVmExtIds field if non-nil, zero value otherwise.

### GetAllocatedVmExtIdsOk

`func (o *HciPhysicalGpu) GetAllocatedVmExtIdsOk() (*[]string, bool)`

GetAllocatedVmExtIdsOk returns a tuple with the AllocatedVmExtIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedVmExtIds

`func (o *HciPhysicalGpu) SetAllocatedVmExtIds(v []string)`

SetAllocatedVmExtIds sets AllocatedVmExtIds field to given value.

### HasAllocatedVmExtIds

`func (o *HciPhysicalGpu) HasAllocatedVmExtIds() bool`

HasAllocatedVmExtIds returns a boolean if a field has been set.

### SetAllocatedVmExtIdsNil

`func (o *HciPhysicalGpu) SetAllocatedVmExtIdsNil(b bool)`

 SetAllocatedVmExtIdsNil sets the value for AllocatedVmExtIds to be an explicit nil

### UnsetAllocatedVmExtIds
`func (o *HciPhysicalGpu) UnsetAllocatedVmExtIds()`

UnsetAllocatedVmExtIds ensures that no value is present for AllocatedVmExtIds, not even an explicit nil
### GetAssignable

`func (o *HciPhysicalGpu) GetAssignable() bool`

GetAssignable returns the Assignable field if non-nil, zero value otherwise.

### GetAssignableOk

`func (o *HciPhysicalGpu) GetAssignableOk() (*bool, bool)`

GetAssignableOk returns a tuple with the Assignable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignable

`func (o *HciPhysicalGpu) SetAssignable(v bool)`

SetAssignable sets Assignable field to given value.

### HasAssignable

`func (o *HciPhysicalGpu) HasAssignable() bool`

HasAssignable returns a boolean if a field has been set.

### GetClusterExtId

`func (o *HciPhysicalGpu) GetClusterExtId() string`

GetClusterExtId returns the ClusterExtId field if non-nil, zero value otherwise.

### GetClusterExtIdOk

`func (o *HciPhysicalGpu) GetClusterExtIdOk() (*string, bool)`

GetClusterExtIdOk returns a tuple with the ClusterExtId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterExtId

`func (o *HciPhysicalGpu) SetClusterExtId(v string)`

SetClusterExtId sets ClusterExtId field to given value.

### HasClusterExtId

`func (o *HciPhysicalGpu) HasClusterExtId() bool`

HasClusterExtId returns a boolean if a field has been set.

### GetDeviceId

`func (o *HciPhysicalGpu) GetDeviceId() int32`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *HciPhysicalGpu) GetDeviceIdOk() (*int32, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *HciPhysicalGpu) SetDeviceId(v int32)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *HciPhysicalGpu) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetDeviceName

`func (o *HciPhysicalGpu) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *HciPhysicalGpu) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *HciPhysicalGpu) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *HciPhysicalGpu) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetFrameBufferSizeBytes

`func (o *HciPhysicalGpu) GetFrameBufferSizeBytes() int64`

GetFrameBufferSizeBytes returns the FrameBufferSizeBytes field if non-nil, zero value otherwise.

### GetFrameBufferSizeBytesOk

`func (o *HciPhysicalGpu) GetFrameBufferSizeBytesOk() (*int64, bool)`

GetFrameBufferSizeBytesOk returns a tuple with the FrameBufferSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrameBufferSizeBytes

`func (o *HciPhysicalGpu) SetFrameBufferSizeBytes(v int64)`

SetFrameBufferSizeBytes sets FrameBufferSizeBytes field to given value.

### HasFrameBufferSizeBytes

`func (o *HciPhysicalGpu) HasFrameBufferSizeBytes() bool`

HasFrameBufferSizeBytes returns a boolean if a field has been set.

### GetIsInUse

`func (o *HciPhysicalGpu) GetIsInUse() bool`

GetIsInUse returns the IsInUse field if non-nil, zero value otherwise.

### GetIsInUseOk

`func (o *HciPhysicalGpu) GetIsInUseOk() (*bool, bool)`

GetIsInUseOk returns a tuple with the IsInUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInUse

`func (o *HciPhysicalGpu) SetIsInUse(v bool)`

SetIsInUse sets IsInUse field to given value.

### HasIsInUse

`func (o *HciPhysicalGpu) HasIsInUse() bool`

HasIsInUse returns a boolean if a field has been set.

### GetMode

`func (o *HciPhysicalGpu) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *HciPhysicalGpu) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *HciPhysicalGpu) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *HciPhysicalGpu) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetNumaNode

`func (o *HciPhysicalGpu) GetNumaNode() string`

GetNumaNode returns the NumaNode field if non-nil, zero value otherwise.

### GetNumaNodeOk

`func (o *HciPhysicalGpu) GetNumaNodeOk() (*string, bool)`

GetNumaNodeOk returns a tuple with the NumaNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumaNode

`func (o *HciPhysicalGpu) SetNumaNode(v string)`

SetNumaNode sets NumaNode field to given value.

### HasNumaNode

`func (o *HciPhysicalGpu) HasNumaNode() bool`

HasNumaNode returns a boolean if a field has been set.

### GetPhysicalGpuExtId

`func (o *HciPhysicalGpu) GetPhysicalGpuExtId() string`

GetPhysicalGpuExtId returns the PhysicalGpuExtId field if non-nil, zero value otherwise.

### GetPhysicalGpuExtIdOk

`func (o *HciPhysicalGpu) GetPhysicalGpuExtIdOk() (*string, bool)`

GetPhysicalGpuExtIdOk returns a tuple with the PhysicalGpuExtId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalGpuExtId

`func (o *HciPhysicalGpu) SetPhysicalGpuExtId(v string)`

SetPhysicalGpuExtId sets PhysicalGpuExtId field to given value.

### HasPhysicalGpuExtId

`func (o *HciPhysicalGpu) HasPhysicalGpuExtId() bool`

HasPhysicalGpuExtId returns a boolean if a field has been set.

### GetSbdf

`func (o *HciPhysicalGpu) GetSbdf() string`

GetSbdf returns the Sbdf field if non-nil, zero value otherwise.

### GetSbdfOk

`func (o *HciPhysicalGpu) GetSbdfOk() (*string, bool)`

GetSbdfOk returns a tuple with the Sbdf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSbdf

`func (o *HciPhysicalGpu) SetSbdf(v string)`

SetSbdf sets Sbdf field to given value.

### HasSbdf

`func (o *HciPhysicalGpu) HasSbdf() bool`

HasSbdf returns a boolean if a field has been set.

### GetType

`func (o *HciPhysicalGpu) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HciPhysicalGpu) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HciPhysicalGpu) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *HciPhysicalGpu) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVendor

`func (o *HciPhysicalGpu) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *HciPhysicalGpu) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *HciPhysicalGpu) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *HciPhysicalGpu) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetAllocatedVms

`func (o *HciPhysicalGpu) GetAllocatedVms() []HciBaseVmRelationship`

GetAllocatedVms returns the AllocatedVms field if non-nil, zero value otherwise.

### GetAllocatedVmsOk

`func (o *HciPhysicalGpu) GetAllocatedVmsOk() (*[]HciBaseVmRelationship, bool)`

GetAllocatedVmsOk returns a tuple with the AllocatedVms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedVms

`func (o *HciPhysicalGpu) SetAllocatedVms(v []HciBaseVmRelationship)`

SetAllocatedVms sets AllocatedVms field to given value.

### HasAllocatedVms

`func (o *HciPhysicalGpu) HasAllocatedVms() bool`

HasAllocatedVms returns a boolean if a field has been set.

### SetAllocatedVmsNil

`func (o *HciPhysicalGpu) SetAllocatedVmsNil(b bool)`

 SetAllocatedVmsNil sets the value for AllocatedVms to be an explicit nil

### UnsetAllocatedVms
`func (o *HciPhysicalGpu) UnsetAllocatedVms()`

UnsetAllocatedVms ensures that no value is present for AllocatedVms, not even an explicit nil
### GetCluster

`func (o *HciPhysicalGpu) GetCluster() HciClusterRelationship`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *HciPhysicalGpu) GetClusterOk() (*HciClusterRelationship, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *HciPhysicalGpu) SetCluster(v HciClusterRelationship)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *HciPhysicalGpu) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### SetClusterNil

`func (o *HciPhysicalGpu) SetClusterNil(b bool)`

 SetClusterNil sets the value for Cluster to be an explicit nil

### UnsetCluster
`func (o *HciPhysicalGpu) UnsetCluster()`

UnsetCluster ensures that no value is present for Cluster, not even an explicit nil
### GetRegisteredDevice

`func (o *HciPhysicalGpu) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *HciPhysicalGpu) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *HciPhysicalGpu) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *HciPhysicalGpu) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *HciPhysicalGpu) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *HciPhysicalGpu) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



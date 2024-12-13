# HciGpu

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hci.Gpu"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hci.Gpu"]
**ClusterExtId** | Pointer to **string** | The unique identifier of the cluster. | [optional] [readonly] 
**ClusterName** | Pointer to **string** | The name of the cluster this GPU belongs to. | [optional] [readonly] 
**ControllerVmId** | Pointer to **string** | The unique identifier of the controller VM. | [optional] [readonly] 
**GpuExtId** | Pointer to **string** | The unique identifier of the gpu. | [optional] [readonly] 
**NodeExtId** | Pointer to **string** | The unique identifier of the node. | [optional] [readonly] 
**NumberOfVgpusAllocated** | Pointer to **int64** | The number of vGPUs allocated. | [optional] [readonly] 
**VirtualGpuConfig** | Pointer to [**NullableHciVirtualGpuConfig**](HciVirtualGpuConfig.md) |  | [optional] 
**Node** | Pointer to [**NullableHciNodeRelationship**](HciNodeRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewHciGpu

`func NewHciGpu(classId string, objectType string, ) *HciGpu`

NewHciGpu instantiates a new HciGpu object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHciGpuWithDefaults

`func NewHciGpuWithDefaults() *HciGpu`

NewHciGpuWithDefaults instantiates a new HciGpu object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HciGpu) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HciGpu) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HciGpu) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HciGpu) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HciGpu) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HciGpu) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetClusterExtId

`func (o *HciGpu) GetClusterExtId() string`

GetClusterExtId returns the ClusterExtId field if non-nil, zero value otherwise.

### GetClusterExtIdOk

`func (o *HciGpu) GetClusterExtIdOk() (*string, bool)`

GetClusterExtIdOk returns a tuple with the ClusterExtId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterExtId

`func (o *HciGpu) SetClusterExtId(v string)`

SetClusterExtId sets ClusterExtId field to given value.

### HasClusterExtId

`func (o *HciGpu) HasClusterExtId() bool`

HasClusterExtId returns a boolean if a field has been set.

### GetClusterName

`func (o *HciGpu) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *HciGpu) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *HciGpu) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *HciGpu) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.

### GetControllerVmId

`func (o *HciGpu) GetControllerVmId() string`

GetControllerVmId returns the ControllerVmId field if non-nil, zero value otherwise.

### GetControllerVmIdOk

`func (o *HciGpu) GetControllerVmIdOk() (*string, bool)`

GetControllerVmIdOk returns a tuple with the ControllerVmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerVmId

`func (o *HciGpu) SetControllerVmId(v string)`

SetControllerVmId sets ControllerVmId field to given value.

### HasControllerVmId

`func (o *HciGpu) HasControllerVmId() bool`

HasControllerVmId returns a boolean if a field has been set.

### GetGpuExtId

`func (o *HciGpu) GetGpuExtId() string`

GetGpuExtId returns the GpuExtId field if non-nil, zero value otherwise.

### GetGpuExtIdOk

`func (o *HciGpu) GetGpuExtIdOk() (*string, bool)`

GetGpuExtIdOk returns a tuple with the GpuExtId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuExtId

`func (o *HciGpu) SetGpuExtId(v string)`

SetGpuExtId sets GpuExtId field to given value.

### HasGpuExtId

`func (o *HciGpu) HasGpuExtId() bool`

HasGpuExtId returns a boolean if a field has been set.

### GetNodeExtId

`func (o *HciGpu) GetNodeExtId() string`

GetNodeExtId returns the NodeExtId field if non-nil, zero value otherwise.

### GetNodeExtIdOk

`func (o *HciGpu) GetNodeExtIdOk() (*string, bool)`

GetNodeExtIdOk returns a tuple with the NodeExtId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeExtId

`func (o *HciGpu) SetNodeExtId(v string)`

SetNodeExtId sets NodeExtId field to given value.

### HasNodeExtId

`func (o *HciGpu) HasNodeExtId() bool`

HasNodeExtId returns a boolean if a field has been set.

### GetNumberOfVgpusAllocated

`func (o *HciGpu) GetNumberOfVgpusAllocated() int64`

GetNumberOfVgpusAllocated returns the NumberOfVgpusAllocated field if non-nil, zero value otherwise.

### GetNumberOfVgpusAllocatedOk

`func (o *HciGpu) GetNumberOfVgpusAllocatedOk() (*int64, bool)`

GetNumberOfVgpusAllocatedOk returns a tuple with the NumberOfVgpusAllocated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfVgpusAllocated

`func (o *HciGpu) SetNumberOfVgpusAllocated(v int64)`

SetNumberOfVgpusAllocated sets NumberOfVgpusAllocated field to given value.

### HasNumberOfVgpusAllocated

`func (o *HciGpu) HasNumberOfVgpusAllocated() bool`

HasNumberOfVgpusAllocated returns a boolean if a field has been set.

### GetVirtualGpuConfig

`func (o *HciGpu) GetVirtualGpuConfig() HciVirtualGpuConfig`

GetVirtualGpuConfig returns the VirtualGpuConfig field if non-nil, zero value otherwise.

### GetVirtualGpuConfigOk

`func (o *HciGpu) GetVirtualGpuConfigOk() (*HciVirtualGpuConfig, bool)`

GetVirtualGpuConfigOk returns a tuple with the VirtualGpuConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualGpuConfig

`func (o *HciGpu) SetVirtualGpuConfig(v HciVirtualGpuConfig)`

SetVirtualGpuConfig sets VirtualGpuConfig field to given value.

### HasVirtualGpuConfig

`func (o *HciGpu) HasVirtualGpuConfig() bool`

HasVirtualGpuConfig returns a boolean if a field has been set.

### SetVirtualGpuConfigNil

`func (o *HciGpu) SetVirtualGpuConfigNil(b bool)`

 SetVirtualGpuConfigNil sets the value for VirtualGpuConfig to be an explicit nil

### UnsetVirtualGpuConfig
`func (o *HciGpu) UnsetVirtualGpuConfig()`

UnsetVirtualGpuConfig ensures that no value is present for VirtualGpuConfig, not even an explicit nil
### GetNode

`func (o *HciGpu) GetNode() HciNodeRelationship`

GetNode returns the Node field if non-nil, zero value otherwise.

### GetNodeOk

`func (o *HciGpu) GetNodeOk() (*HciNodeRelationship, bool)`

GetNodeOk returns a tuple with the Node field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNode

`func (o *HciGpu) SetNode(v HciNodeRelationship)`

SetNode sets Node field to given value.

### HasNode

`func (o *HciGpu) HasNode() bool`

HasNode returns a boolean if a field has been set.

### SetNodeNil

`func (o *HciGpu) SetNodeNil(b bool)`

 SetNodeNil sets the value for Node to be an explicit nil

### UnsetNode
`func (o *HciGpu) UnsetNode()`

UnsetNode ensures that no value is present for Node, not even an explicit nil
### GetRegisteredDevice

`func (o *HciGpu) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *HciGpu) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *HciGpu) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *HciGpu) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *HciGpu) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *HciGpu) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



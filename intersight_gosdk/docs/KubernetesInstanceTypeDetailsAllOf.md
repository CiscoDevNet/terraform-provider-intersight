# KubernetesInstanceTypeDetailsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "kubernetes.InstanceTypeDetails"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "kubernetes.InstanceTypeDetails"]
**Cpu** | Pointer to **int64** | Number of CPUs in the instancetype. | [optional] 
**DiskSize** | Pointer to **int64** | Ephemeral disk capacity to be provided with units example - 10Gi. | [optional] 
**Memory** | Pointer to **int64** | Virtual machine memory defined in mebibytes (MiB). | [optional] 

## Methods

### NewKubernetesInstanceTypeDetailsAllOf

`func NewKubernetesInstanceTypeDetailsAllOf(classId string, objectType string, ) *KubernetesInstanceTypeDetailsAllOf`

NewKubernetesInstanceTypeDetailsAllOf instantiates a new KubernetesInstanceTypeDetailsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesInstanceTypeDetailsAllOfWithDefaults

`func NewKubernetesInstanceTypeDetailsAllOfWithDefaults() *KubernetesInstanceTypeDetailsAllOf`

NewKubernetesInstanceTypeDetailsAllOfWithDefaults instantiates a new KubernetesInstanceTypeDetailsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *KubernetesInstanceTypeDetailsAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *KubernetesInstanceTypeDetailsAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *KubernetesInstanceTypeDetailsAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *KubernetesInstanceTypeDetailsAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KubernetesInstanceTypeDetailsAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KubernetesInstanceTypeDetailsAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCpu

`func (o *KubernetesInstanceTypeDetailsAllOf) GetCpu() int64`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *KubernetesInstanceTypeDetailsAllOf) GetCpuOk() (*int64, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *KubernetesInstanceTypeDetailsAllOf) SetCpu(v int64)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *KubernetesInstanceTypeDetailsAllOf) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetDiskSize

`func (o *KubernetesInstanceTypeDetailsAllOf) GetDiskSize() int64`

GetDiskSize returns the DiskSize field if non-nil, zero value otherwise.

### GetDiskSizeOk

`func (o *KubernetesInstanceTypeDetailsAllOf) GetDiskSizeOk() (*int64, bool)`

GetDiskSizeOk returns a tuple with the DiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSize

`func (o *KubernetesInstanceTypeDetailsAllOf) SetDiskSize(v int64)`

SetDiskSize sets DiskSize field to given value.

### HasDiskSize

`func (o *KubernetesInstanceTypeDetailsAllOf) HasDiskSize() bool`

HasDiskSize returns a boolean if a field has been set.

### GetMemory

`func (o *KubernetesInstanceTypeDetailsAllOf) GetMemory() int64`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *KubernetesInstanceTypeDetailsAllOf) GetMemoryOk() (*int64, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *KubernetesInstanceTypeDetailsAllOf) SetMemory(v int64)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *KubernetesInstanceTypeDetailsAllOf) HasMemory() bool`

HasMemory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



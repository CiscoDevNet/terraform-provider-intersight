# VirtualizationVmwareResourceConsumption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "virtualization.VmwareResourceConsumption"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "virtualization.VmwareResourceConsumption"]
**CpuConsumed** | Pointer to **int64** | The amount of CPU consumed in Hz. | [optional] 
**MemoryConsumed** | Pointer to **int64** | Memory consumed by this host in bytes. | [optional] 

## Methods

### NewVirtualizationVmwareResourceConsumption

`func NewVirtualizationVmwareResourceConsumption(classId string, objectType string, ) *VirtualizationVmwareResourceConsumption`

NewVirtualizationVmwareResourceConsumption instantiates a new VirtualizationVmwareResourceConsumption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationVmwareResourceConsumptionWithDefaults

`func NewVirtualizationVmwareResourceConsumptionWithDefaults() *VirtualizationVmwareResourceConsumption`

NewVirtualizationVmwareResourceConsumptionWithDefaults instantiates a new VirtualizationVmwareResourceConsumption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationVmwareResourceConsumption) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationVmwareResourceConsumption) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationVmwareResourceConsumption) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationVmwareResourceConsumption) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationVmwareResourceConsumption) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationVmwareResourceConsumption) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCpuConsumed

`func (o *VirtualizationVmwareResourceConsumption) GetCpuConsumed() int64`

GetCpuConsumed returns the CpuConsumed field if non-nil, zero value otherwise.

### GetCpuConsumedOk

`func (o *VirtualizationVmwareResourceConsumption) GetCpuConsumedOk() (*int64, bool)`

GetCpuConsumedOk returns a tuple with the CpuConsumed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuConsumed

`func (o *VirtualizationVmwareResourceConsumption) SetCpuConsumed(v int64)`

SetCpuConsumed sets CpuConsumed field to given value.

### HasCpuConsumed

`func (o *VirtualizationVmwareResourceConsumption) HasCpuConsumed() bool`

HasCpuConsumed returns a boolean if a field has been set.

### GetMemoryConsumed

`func (o *VirtualizationVmwareResourceConsumption) GetMemoryConsumed() int64`

GetMemoryConsumed returns the MemoryConsumed field if non-nil, zero value otherwise.

### GetMemoryConsumedOk

`func (o *VirtualizationVmwareResourceConsumption) GetMemoryConsumedOk() (*int64, bool)`

GetMemoryConsumedOk returns a tuple with the MemoryConsumed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryConsumed

`func (o *VirtualizationVmwareResourceConsumption) SetMemoryConsumed(v int64)`

SetMemoryConsumed sets MemoryConsumed field to given value.

### HasMemoryConsumed

`func (o *VirtualizationVmwareResourceConsumption) HasMemoryConsumed() bool`

HasMemoryConsumed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



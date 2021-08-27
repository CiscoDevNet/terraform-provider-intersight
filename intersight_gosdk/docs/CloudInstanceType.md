# CloudInstanceType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "cloud.InstanceType"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "cloud.InstanceType"]
**Architecture** | Pointer to **string** | Cpu architecture, for example, x86. * &#x60;x86_64&#x60; - The virtual machine cpu architecture type x86_64. * &#x60;x86&#x60; - The virtual machine cpu architecture type x86. | [optional] [readonly] [default to "x86_64"]
**CpuSpeed** | Pointer to **int64** | The speed of the processor, in GHz. | [optional] [readonly] 
**Cpus** | Pointer to **int64** | The number of CPUs for the instance type. | [optional] [readonly] 
**InstanceTypeId** | Pointer to **string** | The ID of the instance type. | [optional] [readonly] 
**Memory** | Pointer to **int64** | The size of the memory, in bytes. | [optional] [readonly] 
**Name** | Pointer to **string** | Name to identity the instance type. | [optional] [readonly] 
**Platform** | Pointer to **string** | Operation System, for example, Linux. | [optional] [readonly] 

## Methods

### NewCloudInstanceType

`func NewCloudInstanceType(classId string, objectType string, ) *CloudInstanceType`

NewCloudInstanceType instantiates a new CloudInstanceType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudInstanceTypeWithDefaults

`func NewCloudInstanceTypeWithDefaults() *CloudInstanceType`

NewCloudInstanceTypeWithDefaults instantiates a new CloudInstanceType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CloudInstanceType) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CloudInstanceType) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CloudInstanceType) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CloudInstanceType) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CloudInstanceType) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CloudInstanceType) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetArchitecture

`func (o *CloudInstanceType) GetArchitecture() string`

GetArchitecture returns the Architecture field if non-nil, zero value otherwise.

### GetArchitectureOk

`func (o *CloudInstanceType) GetArchitectureOk() (*string, bool)`

GetArchitectureOk returns a tuple with the Architecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitecture

`func (o *CloudInstanceType) SetArchitecture(v string)`

SetArchitecture sets Architecture field to given value.

### HasArchitecture

`func (o *CloudInstanceType) HasArchitecture() bool`

HasArchitecture returns a boolean if a field has been set.

### GetCpuSpeed

`func (o *CloudInstanceType) GetCpuSpeed() int64`

GetCpuSpeed returns the CpuSpeed field if non-nil, zero value otherwise.

### GetCpuSpeedOk

`func (o *CloudInstanceType) GetCpuSpeedOk() (*int64, bool)`

GetCpuSpeedOk returns a tuple with the CpuSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuSpeed

`func (o *CloudInstanceType) SetCpuSpeed(v int64)`

SetCpuSpeed sets CpuSpeed field to given value.

### HasCpuSpeed

`func (o *CloudInstanceType) HasCpuSpeed() bool`

HasCpuSpeed returns a boolean if a field has been set.

### GetCpus

`func (o *CloudInstanceType) GetCpus() int64`

GetCpus returns the Cpus field if non-nil, zero value otherwise.

### GetCpusOk

`func (o *CloudInstanceType) GetCpusOk() (*int64, bool)`

GetCpusOk returns a tuple with the Cpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpus

`func (o *CloudInstanceType) SetCpus(v int64)`

SetCpus sets Cpus field to given value.

### HasCpus

`func (o *CloudInstanceType) HasCpus() bool`

HasCpus returns a boolean if a field has been set.

### GetInstanceTypeId

`func (o *CloudInstanceType) GetInstanceTypeId() string`

GetInstanceTypeId returns the InstanceTypeId field if non-nil, zero value otherwise.

### GetInstanceTypeIdOk

`func (o *CloudInstanceType) GetInstanceTypeIdOk() (*string, bool)`

GetInstanceTypeIdOk returns a tuple with the InstanceTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceTypeId

`func (o *CloudInstanceType) SetInstanceTypeId(v string)`

SetInstanceTypeId sets InstanceTypeId field to given value.

### HasInstanceTypeId

`func (o *CloudInstanceType) HasInstanceTypeId() bool`

HasInstanceTypeId returns a boolean if a field has been set.

### GetMemory

`func (o *CloudInstanceType) GetMemory() int64`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *CloudInstanceType) GetMemoryOk() (*int64, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *CloudInstanceType) SetMemory(v int64)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *CloudInstanceType) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetName

`func (o *CloudInstanceType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudInstanceType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudInstanceType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudInstanceType) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPlatform

`func (o *CloudInstanceType) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *CloudInstanceType) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *CloudInstanceType) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *CloudInstanceType) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



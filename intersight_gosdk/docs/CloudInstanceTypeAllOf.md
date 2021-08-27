# CloudInstanceTypeAllOf

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

### NewCloudInstanceTypeAllOf

`func NewCloudInstanceTypeAllOf(classId string, objectType string, ) *CloudInstanceTypeAllOf`

NewCloudInstanceTypeAllOf instantiates a new CloudInstanceTypeAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudInstanceTypeAllOfWithDefaults

`func NewCloudInstanceTypeAllOfWithDefaults() *CloudInstanceTypeAllOf`

NewCloudInstanceTypeAllOfWithDefaults instantiates a new CloudInstanceTypeAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CloudInstanceTypeAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CloudInstanceTypeAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CloudInstanceTypeAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CloudInstanceTypeAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CloudInstanceTypeAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CloudInstanceTypeAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetArchitecture

`func (o *CloudInstanceTypeAllOf) GetArchitecture() string`

GetArchitecture returns the Architecture field if non-nil, zero value otherwise.

### GetArchitectureOk

`func (o *CloudInstanceTypeAllOf) GetArchitectureOk() (*string, bool)`

GetArchitectureOk returns a tuple with the Architecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitecture

`func (o *CloudInstanceTypeAllOf) SetArchitecture(v string)`

SetArchitecture sets Architecture field to given value.

### HasArchitecture

`func (o *CloudInstanceTypeAllOf) HasArchitecture() bool`

HasArchitecture returns a boolean if a field has been set.

### GetCpuSpeed

`func (o *CloudInstanceTypeAllOf) GetCpuSpeed() int64`

GetCpuSpeed returns the CpuSpeed field if non-nil, zero value otherwise.

### GetCpuSpeedOk

`func (o *CloudInstanceTypeAllOf) GetCpuSpeedOk() (*int64, bool)`

GetCpuSpeedOk returns a tuple with the CpuSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuSpeed

`func (o *CloudInstanceTypeAllOf) SetCpuSpeed(v int64)`

SetCpuSpeed sets CpuSpeed field to given value.

### HasCpuSpeed

`func (o *CloudInstanceTypeAllOf) HasCpuSpeed() bool`

HasCpuSpeed returns a boolean if a field has been set.

### GetCpus

`func (o *CloudInstanceTypeAllOf) GetCpus() int64`

GetCpus returns the Cpus field if non-nil, zero value otherwise.

### GetCpusOk

`func (o *CloudInstanceTypeAllOf) GetCpusOk() (*int64, bool)`

GetCpusOk returns a tuple with the Cpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpus

`func (o *CloudInstanceTypeAllOf) SetCpus(v int64)`

SetCpus sets Cpus field to given value.

### HasCpus

`func (o *CloudInstanceTypeAllOf) HasCpus() bool`

HasCpus returns a boolean if a field has been set.

### GetInstanceTypeId

`func (o *CloudInstanceTypeAllOf) GetInstanceTypeId() string`

GetInstanceTypeId returns the InstanceTypeId field if non-nil, zero value otherwise.

### GetInstanceTypeIdOk

`func (o *CloudInstanceTypeAllOf) GetInstanceTypeIdOk() (*string, bool)`

GetInstanceTypeIdOk returns a tuple with the InstanceTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceTypeId

`func (o *CloudInstanceTypeAllOf) SetInstanceTypeId(v string)`

SetInstanceTypeId sets InstanceTypeId field to given value.

### HasInstanceTypeId

`func (o *CloudInstanceTypeAllOf) HasInstanceTypeId() bool`

HasInstanceTypeId returns a boolean if a field has been set.

### GetMemory

`func (o *CloudInstanceTypeAllOf) GetMemory() int64`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *CloudInstanceTypeAllOf) GetMemoryOk() (*int64, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *CloudInstanceTypeAllOf) SetMemory(v int64)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *CloudInstanceTypeAllOf) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetName

`func (o *CloudInstanceTypeAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudInstanceTypeAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudInstanceTypeAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudInstanceTypeAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPlatform

`func (o *CloudInstanceTypeAllOf) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *CloudInstanceTypeAllOf) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *CloudInstanceTypeAllOf) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *CloudInstanceTypeAllOf) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



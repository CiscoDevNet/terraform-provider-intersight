# CloudSkuInstanceType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "cloud.SkuInstanceType"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "cloud.SkuInstanceType"]
**ArchitectureType** | Pointer to **string** | Indicates if the instance type supports 32 or 64 bit or both. * &#x60;64Bit&#x60; - Enum to indicate that the instance type suppports only 64 bit architecture. * &#x60;32Bit&#x60; - Enum to indicate that the instance type supports only 32 bit architecture. * &#x60;both&#x60; - Enum to indicate that the instance type supports both 32 and 64 bit architecture. | [optional] [default to "64Bit"]
**CpuUnit** | Pointer to **string** | The cpu unit for this instance type. * &#x60;VIRTUAL_CPU&#x60; - The CPU unit used for virtual machines. * &#x60;MILLI_CPU&#x60; - The CPU unit used by containers. | [optional] [default to "VIRTUAL_CPU"]
**CudaSupport** | Pointer to **bool** | Does the instanceType support CUDA architecture. | [optional] 
**LocalStorageSize** | Pointer to **float32** | Total size of local storage for this instance type. | [optional] 
**LocalStorageSizeUnit** | Pointer to **string** | The units for this storage. For e.g. MB or GB or TB. * &#x60;MB&#x60; - Enum to represent mega byte storage unit. * &#x60;GB&#x60; - Enum to represent Giga byte storage unit. * &#x60;TB&#x60; - Enum to represent Tera byte storage unit. | [optional] [default to "MB"]
**MemorySize** | Pointer to **float32** | The RAM size for this instance type. | [optional] 
**MemorySizeUnit** | Pointer to **string** | Units to represent memory, for e.g. MB or GB. * &#x60;MB&#x60; - Enum to represent mega byte storage unit. * &#x60;GB&#x60; - Enum to represent Giga byte storage unit. * &#x60;TB&#x60; - Enum to represent Tera byte storage unit. | [optional] [default to "MB"]
**NetworkPerformance** | Pointer to **string** | Network performance of this instance type. | [optional] 
**NumOfCpus** | Pointer to **int64** | The number of CPUs in this instance type. | [optional] 
**NumOfNics** | Pointer to **int64** | Maximum number of NICs supported by this instance type. | [optional] 

## Methods

### NewCloudSkuInstanceType

`func NewCloudSkuInstanceType(classId string, objectType string, ) *CloudSkuInstanceType`

NewCloudSkuInstanceType instantiates a new CloudSkuInstanceType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudSkuInstanceTypeWithDefaults

`func NewCloudSkuInstanceTypeWithDefaults() *CloudSkuInstanceType`

NewCloudSkuInstanceTypeWithDefaults instantiates a new CloudSkuInstanceType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CloudSkuInstanceType) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CloudSkuInstanceType) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CloudSkuInstanceType) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CloudSkuInstanceType) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CloudSkuInstanceType) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CloudSkuInstanceType) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetArchitectureType

`func (o *CloudSkuInstanceType) GetArchitectureType() string`

GetArchitectureType returns the ArchitectureType field if non-nil, zero value otherwise.

### GetArchitectureTypeOk

`func (o *CloudSkuInstanceType) GetArchitectureTypeOk() (*string, bool)`

GetArchitectureTypeOk returns a tuple with the ArchitectureType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitectureType

`func (o *CloudSkuInstanceType) SetArchitectureType(v string)`

SetArchitectureType sets ArchitectureType field to given value.

### HasArchitectureType

`func (o *CloudSkuInstanceType) HasArchitectureType() bool`

HasArchitectureType returns a boolean if a field has been set.

### GetCpuUnit

`func (o *CloudSkuInstanceType) GetCpuUnit() string`

GetCpuUnit returns the CpuUnit field if non-nil, zero value otherwise.

### GetCpuUnitOk

`func (o *CloudSkuInstanceType) GetCpuUnitOk() (*string, bool)`

GetCpuUnitOk returns a tuple with the CpuUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUnit

`func (o *CloudSkuInstanceType) SetCpuUnit(v string)`

SetCpuUnit sets CpuUnit field to given value.

### HasCpuUnit

`func (o *CloudSkuInstanceType) HasCpuUnit() bool`

HasCpuUnit returns a boolean if a field has been set.

### GetCudaSupport

`func (o *CloudSkuInstanceType) GetCudaSupport() bool`

GetCudaSupport returns the CudaSupport field if non-nil, zero value otherwise.

### GetCudaSupportOk

`func (o *CloudSkuInstanceType) GetCudaSupportOk() (*bool, bool)`

GetCudaSupportOk returns a tuple with the CudaSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCudaSupport

`func (o *CloudSkuInstanceType) SetCudaSupport(v bool)`

SetCudaSupport sets CudaSupport field to given value.

### HasCudaSupport

`func (o *CloudSkuInstanceType) HasCudaSupport() bool`

HasCudaSupport returns a boolean if a field has been set.

### GetLocalStorageSize

`func (o *CloudSkuInstanceType) GetLocalStorageSize() float32`

GetLocalStorageSize returns the LocalStorageSize field if non-nil, zero value otherwise.

### GetLocalStorageSizeOk

`func (o *CloudSkuInstanceType) GetLocalStorageSizeOk() (*float32, bool)`

GetLocalStorageSizeOk returns a tuple with the LocalStorageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalStorageSize

`func (o *CloudSkuInstanceType) SetLocalStorageSize(v float32)`

SetLocalStorageSize sets LocalStorageSize field to given value.

### HasLocalStorageSize

`func (o *CloudSkuInstanceType) HasLocalStorageSize() bool`

HasLocalStorageSize returns a boolean if a field has been set.

### GetLocalStorageSizeUnit

`func (o *CloudSkuInstanceType) GetLocalStorageSizeUnit() string`

GetLocalStorageSizeUnit returns the LocalStorageSizeUnit field if non-nil, zero value otherwise.

### GetLocalStorageSizeUnitOk

`func (o *CloudSkuInstanceType) GetLocalStorageSizeUnitOk() (*string, bool)`

GetLocalStorageSizeUnitOk returns a tuple with the LocalStorageSizeUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalStorageSizeUnit

`func (o *CloudSkuInstanceType) SetLocalStorageSizeUnit(v string)`

SetLocalStorageSizeUnit sets LocalStorageSizeUnit field to given value.

### HasLocalStorageSizeUnit

`func (o *CloudSkuInstanceType) HasLocalStorageSizeUnit() bool`

HasLocalStorageSizeUnit returns a boolean if a field has been set.

### GetMemorySize

`func (o *CloudSkuInstanceType) GetMemorySize() float32`

GetMemorySize returns the MemorySize field if non-nil, zero value otherwise.

### GetMemorySizeOk

`func (o *CloudSkuInstanceType) GetMemorySizeOk() (*float32, bool)`

GetMemorySizeOk returns a tuple with the MemorySize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemorySize

`func (o *CloudSkuInstanceType) SetMemorySize(v float32)`

SetMemorySize sets MemorySize field to given value.

### HasMemorySize

`func (o *CloudSkuInstanceType) HasMemorySize() bool`

HasMemorySize returns a boolean if a field has been set.

### GetMemorySizeUnit

`func (o *CloudSkuInstanceType) GetMemorySizeUnit() string`

GetMemorySizeUnit returns the MemorySizeUnit field if non-nil, zero value otherwise.

### GetMemorySizeUnitOk

`func (o *CloudSkuInstanceType) GetMemorySizeUnitOk() (*string, bool)`

GetMemorySizeUnitOk returns a tuple with the MemorySizeUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemorySizeUnit

`func (o *CloudSkuInstanceType) SetMemorySizeUnit(v string)`

SetMemorySizeUnit sets MemorySizeUnit field to given value.

### HasMemorySizeUnit

`func (o *CloudSkuInstanceType) HasMemorySizeUnit() bool`

HasMemorySizeUnit returns a boolean if a field has been set.

### GetNetworkPerformance

`func (o *CloudSkuInstanceType) GetNetworkPerformance() string`

GetNetworkPerformance returns the NetworkPerformance field if non-nil, zero value otherwise.

### GetNetworkPerformanceOk

`func (o *CloudSkuInstanceType) GetNetworkPerformanceOk() (*string, bool)`

GetNetworkPerformanceOk returns a tuple with the NetworkPerformance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkPerformance

`func (o *CloudSkuInstanceType) SetNetworkPerformance(v string)`

SetNetworkPerformance sets NetworkPerformance field to given value.

### HasNetworkPerformance

`func (o *CloudSkuInstanceType) HasNetworkPerformance() bool`

HasNetworkPerformance returns a boolean if a field has been set.

### GetNumOfCpus

`func (o *CloudSkuInstanceType) GetNumOfCpus() int64`

GetNumOfCpus returns the NumOfCpus field if non-nil, zero value otherwise.

### GetNumOfCpusOk

`func (o *CloudSkuInstanceType) GetNumOfCpusOk() (*int64, bool)`

GetNumOfCpusOk returns a tuple with the NumOfCpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumOfCpus

`func (o *CloudSkuInstanceType) SetNumOfCpus(v int64)`

SetNumOfCpus sets NumOfCpus field to given value.

### HasNumOfCpus

`func (o *CloudSkuInstanceType) HasNumOfCpus() bool`

HasNumOfCpus returns a boolean if a field has been set.

### GetNumOfNics

`func (o *CloudSkuInstanceType) GetNumOfNics() int64`

GetNumOfNics returns the NumOfNics field if non-nil, zero value otherwise.

### GetNumOfNicsOk

`func (o *CloudSkuInstanceType) GetNumOfNicsOk() (*int64, bool)`

GetNumOfNicsOk returns a tuple with the NumOfNics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumOfNics

`func (o *CloudSkuInstanceType) SetNumOfNics(v int64)`

SetNumOfNics sets NumOfNics field to given value.

### HasNumOfNics

`func (o *CloudSkuInstanceType) HasNumOfNics() bool`

HasNumOfNics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



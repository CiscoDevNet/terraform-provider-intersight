# InfraHardwareInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "infra.HardwareInfo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "infra.HardwareInfo"]
**CpuCores** | Pointer to **int64** | The number of cpu cores on this hardware platform. | [optional] 
**CpuSpeed** | Pointer to **int64** | Speed of cpu in MHz. Usually cpu speeds are reported for modern cpus in GHz but MHz makes it more precise. | [optional] 
**MemorySize** | Pointer to **int64** | The amount of memory allocated (bytes) to this hardware platform. | [optional] 

## Methods

### NewInfraHardwareInfo

`func NewInfraHardwareInfo(classId string, objectType string, ) *InfraHardwareInfo`

NewInfraHardwareInfo instantiates a new InfraHardwareInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInfraHardwareInfoWithDefaults

`func NewInfraHardwareInfoWithDefaults() *InfraHardwareInfo`

NewInfraHardwareInfoWithDefaults instantiates a new InfraHardwareInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *InfraHardwareInfo) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *InfraHardwareInfo) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *InfraHardwareInfo) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *InfraHardwareInfo) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *InfraHardwareInfo) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *InfraHardwareInfo) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCpuCores

`func (o *InfraHardwareInfo) GetCpuCores() int64`

GetCpuCores returns the CpuCores field if non-nil, zero value otherwise.

### GetCpuCoresOk

`func (o *InfraHardwareInfo) GetCpuCoresOk() (*int64, bool)`

GetCpuCoresOk returns a tuple with the CpuCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCores

`func (o *InfraHardwareInfo) SetCpuCores(v int64)`

SetCpuCores sets CpuCores field to given value.

### HasCpuCores

`func (o *InfraHardwareInfo) HasCpuCores() bool`

HasCpuCores returns a boolean if a field has been set.

### GetCpuSpeed

`func (o *InfraHardwareInfo) GetCpuSpeed() int64`

GetCpuSpeed returns the CpuSpeed field if non-nil, zero value otherwise.

### GetCpuSpeedOk

`func (o *InfraHardwareInfo) GetCpuSpeedOk() (*int64, bool)`

GetCpuSpeedOk returns a tuple with the CpuSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuSpeed

`func (o *InfraHardwareInfo) SetCpuSpeed(v int64)`

SetCpuSpeed sets CpuSpeed field to given value.

### HasCpuSpeed

`func (o *InfraHardwareInfo) HasCpuSpeed() bool`

HasCpuSpeed returns a boolean if a field has been set.

### GetMemorySize

`func (o *InfraHardwareInfo) GetMemorySize() int64`

GetMemorySize returns the MemorySize field if non-nil, zero value otherwise.

### GetMemorySizeOk

`func (o *InfraHardwareInfo) GetMemorySizeOk() (*int64, bool)`

GetMemorySizeOk returns a tuple with the MemorySize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemorySize

`func (o *InfraHardwareInfo) SetMemorySize(v int64)`

SetMemorySize sets MemorySize field to given value.

### HasMemorySize

`func (o *InfraHardwareInfo) HasMemorySize() bool`

HasMemorySize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



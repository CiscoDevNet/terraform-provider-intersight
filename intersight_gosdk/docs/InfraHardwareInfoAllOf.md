# InfraHardwareInfoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "infra.HardwareInfo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "infra.HardwareInfo"]
**CpuCores** | Pointer to **int64** | The number of cpu cores on this hardware platform. | [optional] 
**CpuSpeed** | Pointer to **int64** | Speed of cpu in MHz. Usually cpu speeds are reported for modern cpus in GHz but MHz makes it more precise. | [optional] 
**MemorySize** | Pointer to **int64** | The amount of memory allocated (bytes) to this hardware platform. | [optional] 

## Methods

### NewInfraHardwareInfoAllOf

`func NewInfraHardwareInfoAllOf(classId string, objectType string, ) *InfraHardwareInfoAllOf`

NewInfraHardwareInfoAllOf instantiates a new InfraHardwareInfoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInfraHardwareInfoAllOfWithDefaults

`func NewInfraHardwareInfoAllOfWithDefaults() *InfraHardwareInfoAllOf`

NewInfraHardwareInfoAllOfWithDefaults instantiates a new InfraHardwareInfoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *InfraHardwareInfoAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *InfraHardwareInfoAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *InfraHardwareInfoAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *InfraHardwareInfoAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *InfraHardwareInfoAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *InfraHardwareInfoAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCpuCores

`func (o *InfraHardwareInfoAllOf) GetCpuCores() int64`

GetCpuCores returns the CpuCores field if non-nil, zero value otherwise.

### GetCpuCoresOk

`func (o *InfraHardwareInfoAllOf) GetCpuCoresOk() (*int64, bool)`

GetCpuCoresOk returns a tuple with the CpuCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCores

`func (o *InfraHardwareInfoAllOf) SetCpuCores(v int64)`

SetCpuCores sets CpuCores field to given value.

### HasCpuCores

`func (o *InfraHardwareInfoAllOf) HasCpuCores() bool`

HasCpuCores returns a boolean if a field has been set.

### GetCpuSpeed

`func (o *InfraHardwareInfoAllOf) GetCpuSpeed() int64`

GetCpuSpeed returns the CpuSpeed field if non-nil, zero value otherwise.

### GetCpuSpeedOk

`func (o *InfraHardwareInfoAllOf) GetCpuSpeedOk() (*int64, bool)`

GetCpuSpeedOk returns a tuple with the CpuSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuSpeed

`func (o *InfraHardwareInfoAllOf) SetCpuSpeed(v int64)`

SetCpuSpeed sets CpuSpeed field to given value.

### HasCpuSpeed

`func (o *InfraHardwareInfoAllOf) HasCpuSpeed() bool`

HasCpuSpeed returns a boolean if a field has been set.

### GetMemorySize

`func (o *InfraHardwareInfoAllOf) GetMemorySize() int64`

GetMemorySize returns the MemorySize field if non-nil, zero value otherwise.

### GetMemorySizeOk

`func (o *InfraHardwareInfoAllOf) GetMemorySizeOk() (*int64, bool)`

GetMemorySizeOk returns a tuple with the MemorySize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemorySize

`func (o *InfraHardwareInfoAllOf) SetMemorySize(v int64)`

SetMemorySize sets MemorySize field to given value.

### HasMemorySize

`func (o *InfraHardwareInfoAllOf) HasMemorySize() bool`

HasMemorySize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



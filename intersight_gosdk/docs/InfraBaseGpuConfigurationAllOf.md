# InfraBaseGpuConfigurationAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**MemorySize** | Pointer to **int64** | The amount of memory on the GPU (GBs). | [optional] 

## Methods

### NewInfraBaseGpuConfigurationAllOf

`func NewInfraBaseGpuConfigurationAllOf(classId string, objectType string, ) *InfraBaseGpuConfigurationAllOf`

NewInfraBaseGpuConfigurationAllOf instantiates a new InfraBaseGpuConfigurationAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInfraBaseGpuConfigurationAllOfWithDefaults

`func NewInfraBaseGpuConfigurationAllOfWithDefaults() *InfraBaseGpuConfigurationAllOf`

NewInfraBaseGpuConfigurationAllOfWithDefaults instantiates a new InfraBaseGpuConfigurationAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *InfraBaseGpuConfigurationAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *InfraBaseGpuConfigurationAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *InfraBaseGpuConfigurationAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *InfraBaseGpuConfigurationAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *InfraBaseGpuConfigurationAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *InfraBaseGpuConfigurationAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetMemorySize

`func (o *InfraBaseGpuConfigurationAllOf) GetMemorySize() int64`

GetMemorySize returns the MemorySize field if non-nil, zero value otherwise.

### GetMemorySizeOk

`func (o *InfraBaseGpuConfigurationAllOf) GetMemorySizeOk() (*int64, bool)`

GetMemorySizeOk returns a tuple with the MemorySize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemorySize

`func (o *InfraBaseGpuConfigurationAllOf) SetMemorySize(v int64)`

SetMemorySize sets MemorySize field to given value.

### HasMemorySize

`func (o *InfraBaseGpuConfigurationAllOf) HasMemorySize() bool`

HasMemorySize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



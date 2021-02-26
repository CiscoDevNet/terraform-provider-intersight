# MetadataSkuContainerTypeAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "metadata.SkuContainerType"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "metadata.SkuContainerType"]
**CpuUnit** | Pointer to **string** | The CpuUnit. Will be MILLI_CPU for containers. * &#x60;VIRTUAL_CPU&#x60; - The CPU unit used for virtual machines. * &#x60;MILLI_CPU&#x60; - The CPU unit used by containers. | [optional] [default to "VIRTUAL_CPU"]
**NumOfCpus** | Pointer to **int64** | The number of CPUs in this instance type. | [optional] 

## Methods

### NewMetadataSkuContainerTypeAllOf

`func NewMetadataSkuContainerTypeAllOf(classId string, objectType string, ) *MetadataSkuContainerTypeAllOf`

NewMetadataSkuContainerTypeAllOf instantiates a new MetadataSkuContainerTypeAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataSkuContainerTypeAllOfWithDefaults

`func NewMetadataSkuContainerTypeAllOfWithDefaults() *MetadataSkuContainerTypeAllOf`

NewMetadataSkuContainerTypeAllOfWithDefaults instantiates a new MetadataSkuContainerTypeAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *MetadataSkuContainerTypeAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *MetadataSkuContainerTypeAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *MetadataSkuContainerTypeAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *MetadataSkuContainerTypeAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *MetadataSkuContainerTypeAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *MetadataSkuContainerTypeAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCpuUnit

`func (o *MetadataSkuContainerTypeAllOf) GetCpuUnit() string`

GetCpuUnit returns the CpuUnit field if non-nil, zero value otherwise.

### GetCpuUnitOk

`func (o *MetadataSkuContainerTypeAllOf) GetCpuUnitOk() (*string, bool)`

GetCpuUnitOk returns a tuple with the CpuUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUnit

`func (o *MetadataSkuContainerTypeAllOf) SetCpuUnit(v string)`

SetCpuUnit sets CpuUnit field to given value.

### HasCpuUnit

`func (o *MetadataSkuContainerTypeAllOf) HasCpuUnit() bool`

HasCpuUnit returns a boolean if a field has been set.

### GetNumOfCpus

`func (o *MetadataSkuContainerTypeAllOf) GetNumOfCpus() int64`

GetNumOfCpus returns the NumOfCpus field if non-nil, zero value otherwise.

### GetNumOfCpusOk

`func (o *MetadataSkuContainerTypeAllOf) GetNumOfCpusOk() (*int64, bool)`

GetNumOfCpusOk returns a tuple with the NumOfCpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumOfCpus

`func (o *MetadataSkuContainerTypeAllOf) SetNumOfCpus(v int64)`

SetNumOfCpus sets NumOfCpus field to given value.

### HasNumOfCpus

`func (o *MetadataSkuContainerTypeAllOf) HasNumOfCpus() bool`

HasNumOfCpus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



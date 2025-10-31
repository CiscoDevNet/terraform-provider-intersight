# ComputeCpuConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "compute.CpuConfig"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "compute.CpuConfig"]
**CpuId** | Pointer to **string** | Unique identifier for a CPU in the server. * &#x60;Any&#x60; - To map PCIe devices to any CPU. * &#x60;1&#x60; - Identifier used to map PCIe device to CPU-1. * &#x60;2&#x60; - Identifier used to map PCIe device to CPU-2. | [optional] [default to "Any"]

## Methods

### NewComputeCpuConfig

`func NewComputeCpuConfig(classId string, objectType string, ) *ComputeCpuConfig`

NewComputeCpuConfig instantiates a new ComputeCpuConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputeCpuConfigWithDefaults

`func NewComputeCpuConfigWithDefaults() *ComputeCpuConfig`

NewComputeCpuConfigWithDefaults instantiates a new ComputeCpuConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ComputeCpuConfig) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ComputeCpuConfig) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ComputeCpuConfig) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ComputeCpuConfig) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ComputeCpuConfig) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ComputeCpuConfig) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCpuId

`func (o *ComputeCpuConfig) GetCpuId() string`

GetCpuId returns the CpuId field if non-nil, zero value otherwise.

### GetCpuIdOk

`func (o *ComputeCpuConfig) GetCpuIdOk() (*string, bool)`

GetCpuIdOk returns a tuple with the CpuId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuId

`func (o *ComputeCpuConfig) SetCpuId(v string)`

SetCpuId sets CpuId field to given value.

### HasCpuId

`func (o *ComputeCpuConfig) HasCpuId() bool`

HasCpuId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



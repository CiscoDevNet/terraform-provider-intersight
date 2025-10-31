# ComputeGpuPropFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "compute.GpuPropFilter"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "compute.GpuPropFilter"]
**Count** | Pointer to **int64** | Number of GPUs to be selected. | [optional] [default to 1]
**Model** | Pointer to **string** | Specific GPU model to select. * &#x60;Any&#x60; - To select any GPU model available in GPU Node. * &#x60;UCSC-GPU-H100-NVL&#x60; - NVIDIA H100 GPU with NVL interface. * &#x60;UCSC-GPU-H200-NVL&#x60; - NVIDIA H200 GPU with NVLink interface. * &#x60;UCSC-GPU-L40S&#x60; - NVIDIA L40S GPU with PCIe interface. * &#x60;UCSC-GPU-RTXP6000&#x60; - NVIDIA RTXP6000 GPU with PCIe interface. | [optional] [default to "Any"]

## Methods

### NewComputeGpuPropFilter

`func NewComputeGpuPropFilter(classId string, objectType string, ) *ComputeGpuPropFilter`

NewComputeGpuPropFilter instantiates a new ComputeGpuPropFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputeGpuPropFilterWithDefaults

`func NewComputeGpuPropFilterWithDefaults() *ComputeGpuPropFilter`

NewComputeGpuPropFilterWithDefaults instantiates a new ComputeGpuPropFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ComputeGpuPropFilter) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ComputeGpuPropFilter) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ComputeGpuPropFilter) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ComputeGpuPropFilter) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ComputeGpuPropFilter) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ComputeGpuPropFilter) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCount

`func (o *ComputeGpuPropFilter) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ComputeGpuPropFilter) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ComputeGpuPropFilter) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *ComputeGpuPropFilter) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetModel

`func (o *ComputeGpuPropFilter) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ComputeGpuPropFilter) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ComputeGpuPropFilter) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ComputeGpuPropFilter) HasModel() bool`

HasModel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



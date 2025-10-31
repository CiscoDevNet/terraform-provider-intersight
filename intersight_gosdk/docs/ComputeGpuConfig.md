# ComputeGpuConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "compute.GpuConfig"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "compute.GpuConfig"]
**GpuPropFilters** | Pointer to [**[]ComputeGpuPropFilter**](ComputeGpuPropFilter.md) |  | [optional] 

## Methods

### NewComputeGpuConfig

`func NewComputeGpuConfig(classId string, objectType string, ) *ComputeGpuConfig`

NewComputeGpuConfig instantiates a new ComputeGpuConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputeGpuConfigWithDefaults

`func NewComputeGpuConfigWithDefaults() *ComputeGpuConfig`

NewComputeGpuConfigWithDefaults instantiates a new ComputeGpuConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ComputeGpuConfig) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ComputeGpuConfig) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ComputeGpuConfig) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ComputeGpuConfig) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ComputeGpuConfig) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ComputeGpuConfig) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetGpuPropFilters

`func (o *ComputeGpuConfig) GetGpuPropFilters() []ComputeGpuPropFilter`

GetGpuPropFilters returns the GpuPropFilters field if non-nil, zero value otherwise.

### GetGpuPropFiltersOk

`func (o *ComputeGpuConfig) GetGpuPropFiltersOk() (*[]ComputeGpuPropFilter, bool)`

GetGpuPropFiltersOk returns a tuple with the GpuPropFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuPropFilters

`func (o *ComputeGpuConfig) SetGpuPropFilters(v []ComputeGpuPropFilter)`

SetGpuPropFilters sets GpuPropFilters field to given value.

### HasGpuPropFilters

`func (o *ComputeGpuConfig) HasGpuPropFilters() bool`

HasGpuPropFilters returns a boolean if a field has been set.

### SetGpuPropFiltersNil

`func (o *ComputeGpuConfig) SetGpuPropFiltersNil(b bool)`

 SetGpuPropFiltersNil sets the value for GpuPropFilters to be an explicit nil

### UnsetGpuPropFilters
`func (o *ComputeGpuConfig) UnsetGpuPropFilters()`

UnsetGpuPropFilters ensures that no value is present for GpuPropFilters, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# ResourceGpuQualifier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "resource.GpuQualifier"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "resource.GpuQualifier"]
**GpuControllersRange** | Pointer to [**ResourceGpuControllersRangeFilter**](ResourceGpuControllersRangeFilter.md) |  | [optional] 
**GpuEvaluationType** | Pointer to **string** | The GPU qualifier used for resource qualification is based on gpuEvaluationType. * &#x60;ServerWithoutGpu&#x60; - It is used to specify the evaluation of a resource without a GPU. * &#x60;ServerWithGpu&#x60; - It is used to specify the evaluation of a resource with a GPU. * &#x60;Unspecified&#x60; - It is used to specify that the GPU qualifier should be ignored for the evaluation of a resource. | [optional] [default to "ServerWithoutGpu"]
**Pids** | Pointer to **[]string** |  | [optional] 
**Vendor** | Pointer to **string** | The qualification of resources is based on vendor of GPU. | [optional] 

## Methods

### NewResourceGpuQualifier

`func NewResourceGpuQualifier(classId string, objectType string, ) *ResourceGpuQualifier`

NewResourceGpuQualifier instantiates a new ResourceGpuQualifier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceGpuQualifierWithDefaults

`func NewResourceGpuQualifierWithDefaults() *ResourceGpuQualifier`

NewResourceGpuQualifierWithDefaults instantiates a new ResourceGpuQualifier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ResourceGpuQualifier) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ResourceGpuQualifier) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ResourceGpuQualifier) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ResourceGpuQualifier) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ResourceGpuQualifier) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ResourceGpuQualifier) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetGpuControllersRange

`func (o *ResourceGpuQualifier) GetGpuControllersRange() ResourceGpuControllersRangeFilter`

GetGpuControllersRange returns the GpuControllersRange field if non-nil, zero value otherwise.

### GetGpuControllersRangeOk

`func (o *ResourceGpuQualifier) GetGpuControllersRangeOk() (*ResourceGpuControllersRangeFilter, bool)`

GetGpuControllersRangeOk returns a tuple with the GpuControllersRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuControllersRange

`func (o *ResourceGpuQualifier) SetGpuControllersRange(v ResourceGpuControllersRangeFilter)`

SetGpuControllersRange sets GpuControllersRange field to given value.

### HasGpuControllersRange

`func (o *ResourceGpuQualifier) HasGpuControllersRange() bool`

HasGpuControllersRange returns a boolean if a field has been set.

### GetGpuEvaluationType

`func (o *ResourceGpuQualifier) GetGpuEvaluationType() string`

GetGpuEvaluationType returns the GpuEvaluationType field if non-nil, zero value otherwise.

### GetGpuEvaluationTypeOk

`func (o *ResourceGpuQualifier) GetGpuEvaluationTypeOk() (*string, bool)`

GetGpuEvaluationTypeOk returns a tuple with the GpuEvaluationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuEvaluationType

`func (o *ResourceGpuQualifier) SetGpuEvaluationType(v string)`

SetGpuEvaluationType sets GpuEvaluationType field to given value.

### HasGpuEvaluationType

`func (o *ResourceGpuQualifier) HasGpuEvaluationType() bool`

HasGpuEvaluationType returns a boolean if a field has been set.

### GetPids

`func (o *ResourceGpuQualifier) GetPids() []string`

GetPids returns the Pids field if non-nil, zero value otherwise.

### GetPidsOk

`func (o *ResourceGpuQualifier) GetPidsOk() (*[]string, bool)`

GetPidsOk returns a tuple with the Pids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPids

`func (o *ResourceGpuQualifier) SetPids(v []string)`

SetPids sets Pids field to given value.

### HasPids

`func (o *ResourceGpuQualifier) HasPids() bool`

HasPids returns a boolean if a field has been set.

### SetPidsNil

`func (o *ResourceGpuQualifier) SetPidsNil(b bool)`

 SetPidsNil sets the value for Pids to be an explicit nil

### UnsetPids
`func (o *ResourceGpuQualifier) UnsetPids()`

UnsetPids ensures that no value is present for Pids, not even an explicit nil
### GetVendor

`func (o *ResourceGpuQualifier) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *ResourceGpuQualifier) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *ResourceGpuQualifier) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *ResourceGpuQualifier) HasVendor() bool`

HasVendor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



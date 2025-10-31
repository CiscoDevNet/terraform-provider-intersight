# ComputeAdapterUnitConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "compute.AdapterUnitConfig"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "compute.AdapterUnitConfig"]
**AdapterUnitPropFilters** | Pointer to [**[]ComputeAdapterUnitPropFilter**](ComputeAdapterUnitPropFilter.md) |  | [optional] 

## Methods

### NewComputeAdapterUnitConfig

`func NewComputeAdapterUnitConfig(classId string, objectType string, ) *ComputeAdapterUnitConfig`

NewComputeAdapterUnitConfig instantiates a new ComputeAdapterUnitConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputeAdapterUnitConfigWithDefaults

`func NewComputeAdapterUnitConfigWithDefaults() *ComputeAdapterUnitConfig`

NewComputeAdapterUnitConfigWithDefaults instantiates a new ComputeAdapterUnitConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ComputeAdapterUnitConfig) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ComputeAdapterUnitConfig) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ComputeAdapterUnitConfig) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ComputeAdapterUnitConfig) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ComputeAdapterUnitConfig) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ComputeAdapterUnitConfig) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdapterUnitPropFilters

`func (o *ComputeAdapterUnitConfig) GetAdapterUnitPropFilters() []ComputeAdapterUnitPropFilter`

GetAdapterUnitPropFilters returns the AdapterUnitPropFilters field if non-nil, zero value otherwise.

### GetAdapterUnitPropFiltersOk

`func (o *ComputeAdapterUnitConfig) GetAdapterUnitPropFiltersOk() (*[]ComputeAdapterUnitPropFilter, bool)`

GetAdapterUnitPropFiltersOk returns a tuple with the AdapterUnitPropFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapterUnitPropFilters

`func (o *ComputeAdapterUnitConfig) SetAdapterUnitPropFilters(v []ComputeAdapterUnitPropFilter)`

SetAdapterUnitPropFilters sets AdapterUnitPropFilters field to given value.

### HasAdapterUnitPropFilters

`func (o *ComputeAdapterUnitConfig) HasAdapterUnitPropFilters() bool`

HasAdapterUnitPropFilters returns a boolean if a field has been set.

### SetAdapterUnitPropFiltersNil

`func (o *ComputeAdapterUnitConfig) SetAdapterUnitPropFiltersNil(b bool)`

 SetAdapterUnitPropFiltersNil sets the value for AdapterUnitPropFilters to be an explicit nil

### UnsetAdapterUnitPropFilters
`func (o *ComputeAdapterUnitConfig) UnsetAdapterUnitPropFilters()`

UnsetAdapterUnitPropFilters ensures that no value is present for AdapterUnitPropFilters, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



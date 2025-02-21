# CapabilityFexSupportMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "capability.FexSupportMeta"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "capability.FexSupportMeta"]
**Description** | Pointer to **string** | Information related to the list of FEXs. | [optional] [readonly] 
**SeriesId** | Pointer to **string** | Series names of FEXs which will be unsupported in the firmware operation. | [optional] [readonly] 
**SupportedModels** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCapabilityFexSupportMeta

`func NewCapabilityFexSupportMeta(classId string, objectType string, ) *CapabilityFexSupportMeta`

NewCapabilityFexSupportMeta instantiates a new CapabilityFexSupportMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilityFexSupportMetaWithDefaults

`func NewCapabilityFexSupportMetaWithDefaults() *CapabilityFexSupportMeta`

NewCapabilityFexSupportMetaWithDefaults instantiates a new CapabilityFexSupportMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CapabilityFexSupportMeta) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CapabilityFexSupportMeta) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CapabilityFexSupportMeta) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CapabilityFexSupportMeta) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CapabilityFexSupportMeta) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CapabilityFexSupportMeta) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDescription

`func (o *CapabilityFexSupportMeta) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CapabilityFexSupportMeta) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CapabilityFexSupportMeta) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CapabilityFexSupportMeta) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSeriesId

`func (o *CapabilityFexSupportMeta) GetSeriesId() string`

GetSeriesId returns the SeriesId field if non-nil, zero value otherwise.

### GetSeriesIdOk

`func (o *CapabilityFexSupportMeta) GetSeriesIdOk() (*string, bool)`

GetSeriesIdOk returns a tuple with the SeriesId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesId

`func (o *CapabilityFexSupportMeta) SetSeriesId(v string)`

SetSeriesId sets SeriesId field to given value.

### HasSeriesId

`func (o *CapabilityFexSupportMeta) HasSeriesId() bool`

HasSeriesId returns a boolean if a field has been set.

### GetSupportedModels

`func (o *CapabilityFexSupportMeta) GetSupportedModels() []string`

GetSupportedModels returns the SupportedModels field if non-nil, zero value otherwise.

### GetSupportedModelsOk

`func (o *CapabilityFexSupportMeta) GetSupportedModelsOk() (*[]string, bool)`

GetSupportedModelsOk returns a tuple with the SupportedModels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedModels

`func (o *CapabilityFexSupportMeta) SetSupportedModels(v []string)`

SetSupportedModels sets SupportedModels field to given value.

### HasSupportedModels

`func (o *CapabilityFexSupportMeta) HasSupportedModels() bool`

HasSupportedModels returns a boolean if a field has been set.

### SetSupportedModelsNil

`func (o *CapabilityFexSupportMeta) SetSupportedModelsNil(b bool)`

 SetSupportedModelsNil sets the value for SupportedModels to be an explicit nil

### UnsetSupportedModels
`func (o *CapabilityFexSupportMeta) UnsetSupportedModels()`

UnsetSupportedModels ensures that no value is present for SupportedModels, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



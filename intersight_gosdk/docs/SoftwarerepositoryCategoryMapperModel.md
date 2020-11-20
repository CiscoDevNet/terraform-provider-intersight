# SoftwarerepositoryCategoryMapperModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "softwarerepository.CategoryMapperModel"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "softwarerepository.CategoryMapperModel"]
**Category** | Pointer to **string** | The category of the model series. | [optional] 
**DistTag** | Pointer to **string** | The distributable tag value of the model series. | [optional] 
**RegexPattern** | Pointer to **string** | The regex that all images of this model follow. | [optional] 
**SeriesId** | Pointer to **string** | Cisco hardware model series. | [optional] 
**SupportedModels** | Pointer to **[]string** |  | [optional] 

## Methods

### NewSoftwarerepositoryCategoryMapperModel

`func NewSoftwarerepositoryCategoryMapperModel(classId string, objectType string, ) *SoftwarerepositoryCategoryMapperModel`

NewSoftwarerepositoryCategoryMapperModel instantiates a new SoftwarerepositoryCategoryMapperModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoftwarerepositoryCategoryMapperModelWithDefaults

`func NewSoftwarerepositoryCategoryMapperModelWithDefaults() *SoftwarerepositoryCategoryMapperModel`

NewSoftwarerepositoryCategoryMapperModelWithDefaults instantiates a new SoftwarerepositoryCategoryMapperModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *SoftwarerepositoryCategoryMapperModel) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *SoftwarerepositoryCategoryMapperModel) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *SoftwarerepositoryCategoryMapperModel) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *SoftwarerepositoryCategoryMapperModel) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *SoftwarerepositoryCategoryMapperModel) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *SoftwarerepositoryCategoryMapperModel) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCategory

`func (o *SoftwarerepositoryCategoryMapperModel) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *SoftwarerepositoryCategoryMapperModel) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *SoftwarerepositoryCategoryMapperModel) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *SoftwarerepositoryCategoryMapperModel) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetDistTag

`func (o *SoftwarerepositoryCategoryMapperModel) GetDistTag() string`

GetDistTag returns the DistTag field if non-nil, zero value otherwise.

### GetDistTagOk

`func (o *SoftwarerepositoryCategoryMapperModel) GetDistTagOk() (*string, bool)`

GetDistTagOk returns a tuple with the DistTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistTag

`func (o *SoftwarerepositoryCategoryMapperModel) SetDistTag(v string)`

SetDistTag sets DistTag field to given value.

### HasDistTag

`func (o *SoftwarerepositoryCategoryMapperModel) HasDistTag() bool`

HasDistTag returns a boolean if a field has been set.

### GetRegexPattern

`func (o *SoftwarerepositoryCategoryMapperModel) GetRegexPattern() string`

GetRegexPattern returns the RegexPattern field if non-nil, zero value otherwise.

### GetRegexPatternOk

`func (o *SoftwarerepositoryCategoryMapperModel) GetRegexPatternOk() (*string, bool)`

GetRegexPatternOk returns a tuple with the RegexPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegexPattern

`func (o *SoftwarerepositoryCategoryMapperModel) SetRegexPattern(v string)`

SetRegexPattern sets RegexPattern field to given value.

### HasRegexPattern

`func (o *SoftwarerepositoryCategoryMapperModel) HasRegexPattern() bool`

HasRegexPattern returns a boolean if a field has been set.

### GetSeriesId

`func (o *SoftwarerepositoryCategoryMapperModel) GetSeriesId() string`

GetSeriesId returns the SeriesId field if non-nil, zero value otherwise.

### GetSeriesIdOk

`func (o *SoftwarerepositoryCategoryMapperModel) GetSeriesIdOk() (*string, bool)`

GetSeriesIdOk returns a tuple with the SeriesId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesId

`func (o *SoftwarerepositoryCategoryMapperModel) SetSeriesId(v string)`

SetSeriesId sets SeriesId field to given value.

### HasSeriesId

`func (o *SoftwarerepositoryCategoryMapperModel) HasSeriesId() bool`

HasSeriesId returns a boolean if a field has been set.

### GetSupportedModels

`func (o *SoftwarerepositoryCategoryMapperModel) GetSupportedModels() []string`

GetSupportedModels returns the SupportedModels field if non-nil, zero value otherwise.

### GetSupportedModelsOk

`func (o *SoftwarerepositoryCategoryMapperModel) GetSupportedModelsOk() (*[]string, bool)`

GetSupportedModelsOk returns a tuple with the SupportedModels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedModels

`func (o *SoftwarerepositoryCategoryMapperModel) SetSupportedModels(v []string)`

SetSupportedModels sets SupportedModels field to given value.

### HasSupportedModels

`func (o *SoftwarerepositoryCategoryMapperModel) HasSupportedModels() bool`

HasSupportedModels returns a boolean if a field has been set.

### SetSupportedModelsNil

`func (o *SoftwarerepositoryCategoryMapperModel) SetSupportedModelsNil(b bool)`

 SetSupportedModelsNil sets the value for SupportedModels to be an explicit nil

### UnsetSupportedModels
`func (o *SoftwarerepositoryCategoryMapperModel) UnsetSupportedModels()`

UnsetSupportedModels ensures that no value is present for SupportedModels, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# SoftwarerepositoryCategoryMapperModelAllOf

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

### NewSoftwarerepositoryCategoryMapperModelAllOf

`func NewSoftwarerepositoryCategoryMapperModelAllOf(classId string, objectType string, ) *SoftwarerepositoryCategoryMapperModelAllOf`

NewSoftwarerepositoryCategoryMapperModelAllOf instantiates a new SoftwarerepositoryCategoryMapperModelAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoftwarerepositoryCategoryMapperModelAllOfWithDefaults

`func NewSoftwarerepositoryCategoryMapperModelAllOfWithDefaults() *SoftwarerepositoryCategoryMapperModelAllOf`

NewSoftwarerepositoryCategoryMapperModelAllOfWithDefaults instantiates a new SoftwarerepositoryCategoryMapperModelAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *SoftwarerepositoryCategoryMapperModelAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *SoftwarerepositoryCategoryMapperModelAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *SoftwarerepositoryCategoryMapperModelAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *SoftwarerepositoryCategoryMapperModelAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *SoftwarerepositoryCategoryMapperModelAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *SoftwarerepositoryCategoryMapperModelAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCategory

`func (o *SoftwarerepositoryCategoryMapperModelAllOf) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *SoftwarerepositoryCategoryMapperModelAllOf) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *SoftwarerepositoryCategoryMapperModelAllOf) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *SoftwarerepositoryCategoryMapperModelAllOf) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetDistTag

`func (o *SoftwarerepositoryCategoryMapperModelAllOf) GetDistTag() string`

GetDistTag returns the DistTag field if non-nil, zero value otherwise.

### GetDistTagOk

`func (o *SoftwarerepositoryCategoryMapperModelAllOf) GetDistTagOk() (*string, bool)`

GetDistTagOk returns a tuple with the DistTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistTag

`func (o *SoftwarerepositoryCategoryMapperModelAllOf) SetDistTag(v string)`

SetDistTag sets DistTag field to given value.

### HasDistTag

`func (o *SoftwarerepositoryCategoryMapperModelAllOf) HasDistTag() bool`

HasDistTag returns a boolean if a field has been set.

### GetRegexPattern

`func (o *SoftwarerepositoryCategoryMapperModelAllOf) GetRegexPattern() string`

GetRegexPattern returns the RegexPattern field if non-nil, zero value otherwise.

### GetRegexPatternOk

`func (o *SoftwarerepositoryCategoryMapperModelAllOf) GetRegexPatternOk() (*string, bool)`

GetRegexPatternOk returns a tuple with the RegexPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegexPattern

`func (o *SoftwarerepositoryCategoryMapperModelAllOf) SetRegexPattern(v string)`

SetRegexPattern sets RegexPattern field to given value.

### HasRegexPattern

`func (o *SoftwarerepositoryCategoryMapperModelAllOf) HasRegexPattern() bool`

HasRegexPattern returns a boolean if a field has been set.

### GetSeriesId

`func (o *SoftwarerepositoryCategoryMapperModelAllOf) GetSeriesId() string`

GetSeriesId returns the SeriesId field if non-nil, zero value otherwise.

### GetSeriesIdOk

`func (o *SoftwarerepositoryCategoryMapperModelAllOf) GetSeriesIdOk() (*string, bool)`

GetSeriesIdOk returns a tuple with the SeriesId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesId

`func (o *SoftwarerepositoryCategoryMapperModelAllOf) SetSeriesId(v string)`

SetSeriesId sets SeriesId field to given value.

### HasSeriesId

`func (o *SoftwarerepositoryCategoryMapperModelAllOf) HasSeriesId() bool`

HasSeriesId returns a boolean if a field has been set.

### GetSupportedModels

`func (o *SoftwarerepositoryCategoryMapperModelAllOf) GetSupportedModels() []string`

GetSupportedModels returns the SupportedModels field if non-nil, zero value otherwise.

### GetSupportedModelsOk

`func (o *SoftwarerepositoryCategoryMapperModelAllOf) GetSupportedModelsOk() (*[]string, bool)`

GetSupportedModelsOk returns a tuple with the SupportedModels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedModels

`func (o *SoftwarerepositoryCategoryMapperModelAllOf) SetSupportedModels(v []string)`

SetSupportedModels sets SupportedModels field to given value.

### HasSupportedModels

`func (o *SoftwarerepositoryCategoryMapperModelAllOf) HasSupportedModels() bool`

HasSupportedModels returns a boolean if a field has been set.

### SetSupportedModelsNil

`func (o *SoftwarerepositoryCategoryMapperModelAllOf) SetSupportedModelsNil(b bool)`

 SetSupportedModelsNil sets the value for SupportedModels to be an explicit nil

### UnsetSupportedModels
`func (o *SoftwarerepositoryCategoryMapperModelAllOf) UnsetSupportedModels()`

UnsetSupportedModels ensures that no value is present for SupportedModels, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



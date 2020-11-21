# SoftwarerepositoryCategoryMapper

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "softwarerepository.CategoryMapper"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "softwarerepository.CategoryMapper"]
**Category** | Pointer to **string** | The category of the model series. | [optional] 
**FileType** | Pointer to **string** | The type of distributable image, example huu, scu, driver, os. * &#x60;Distributable&#x60; - Stores firmware host utility images and fabric images. * &#x60;DriverDistributable&#x60; - Stores driver distributable images. * &#x60;ServerConfigurationUtilityDistributable&#x60; - Stores server configuration utility images. * &#x60;OperatingSystemFile&#x60; - Stores operating system iso images. * &#x60;HyperflexDistributable&#x60; - It stores HyperFlex images. | [optional] [default to "Distributable"]
**MdfId** | Pointer to **string** | Cisco software repository image category identifier. | [optional] 
**RegexPattern** | Pointer to **string** | The regex that all images of this category follow. | [optional] 
**Source** | Pointer to **string** | The image can be downloaded from cisco.com or external cloud store. * &#x60;Cisco&#x60; - External repository hosted on cisco.com. * &#x60;IntersightCloud&#x60; - Repository hosted by the Intersight Cloud. * &#x60;LocalMachine&#x60; - The file is available on the local client machine. Used as an upload source type. * &#x60;NetworkShare&#x60; - External repository in the customer datacenter. This will typically be a file server. | [optional] [default to "Cisco"]
**SupportedModels** | Pointer to **[]string** |  | [optional] 
**SwId** | Pointer to **string** | The software type id provided by cisco.com. | [optional] 
**TagTypes** | Pointer to **[]string** |  | [optional] 
**Version** | Pointer to **string** | The version from which user can download images from amazon store, if source is external cloud store. | [optional] 

## Methods

### NewSoftwarerepositoryCategoryMapper

`func NewSoftwarerepositoryCategoryMapper(classId string, objectType string, ) *SoftwarerepositoryCategoryMapper`

NewSoftwarerepositoryCategoryMapper instantiates a new SoftwarerepositoryCategoryMapper object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoftwarerepositoryCategoryMapperWithDefaults

`func NewSoftwarerepositoryCategoryMapperWithDefaults() *SoftwarerepositoryCategoryMapper`

NewSoftwarerepositoryCategoryMapperWithDefaults instantiates a new SoftwarerepositoryCategoryMapper object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *SoftwarerepositoryCategoryMapper) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *SoftwarerepositoryCategoryMapper) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *SoftwarerepositoryCategoryMapper) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *SoftwarerepositoryCategoryMapper) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *SoftwarerepositoryCategoryMapper) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *SoftwarerepositoryCategoryMapper) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCategory

`func (o *SoftwarerepositoryCategoryMapper) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *SoftwarerepositoryCategoryMapper) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *SoftwarerepositoryCategoryMapper) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *SoftwarerepositoryCategoryMapper) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetFileType

`func (o *SoftwarerepositoryCategoryMapper) GetFileType() string`

GetFileType returns the FileType field if non-nil, zero value otherwise.

### GetFileTypeOk

`func (o *SoftwarerepositoryCategoryMapper) GetFileTypeOk() (*string, bool)`

GetFileTypeOk returns a tuple with the FileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileType

`func (o *SoftwarerepositoryCategoryMapper) SetFileType(v string)`

SetFileType sets FileType field to given value.

### HasFileType

`func (o *SoftwarerepositoryCategoryMapper) HasFileType() bool`

HasFileType returns a boolean if a field has been set.

### GetMdfId

`func (o *SoftwarerepositoryCategoryMapper) GetMdfId() string`

GetMdfId returns the MdfId field if non-nil, zero value otherwise.

### GetMdfIdOk

`func (o *SoftwarerepositoryCategoryMapper) GetMdfIdOk() (*string, bool)`

GetMdfIdOk returns a tuple with the MdfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdfId

`func (o *SoftwarerepositoryCategoryMapper) SetMdfId(v string)`

SetMdfId sets MdfId field to given value.

### HasMdfId

`func (o *SoftwarerepositoryCategoryMapper) HasMdfId() bool`

HasMdfId returns a boolean if a field has been set.

### GetRegexPattern

`func (o *SoftwarerepositoryCategoryMapper) GetRegexPattern() string`

GetRegexPattern returns the RegexPattern field if non-nil, zero value otherwise.

### GetRegexPatternOk

`func (o *SoftwarerepositoryCategoryMapper) GetRegexPatternOk() (*string, bool)`

GetRegexPatternOk returns a tuple with the RegexPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegexPattern

`func (o *SoftwarerepositoryCategoryMapper) SetRegexPattern(v string)`

SetRegexPattern sets RegexPattern field to given value.

### HasRegexPattern

`func (o *SoftwarerepositoryCategoryMapper) HasRegexPattern() bool`

HasRegexPattern returns a boolean if a field has been set.

### GetSource

`func (o *SoftwarerepositoryCategoryMapper) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *SoftwarerepositoryCategoryMapper) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *SoftwarerepositoryCategoryMapper) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *SoftwarerepositoryCategoryMapper) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSupportedModels

`func (o *SoftwarerepositoryCategoryMapper) GetSupportedModels() []string`

GetSupportedModels returns the SupportedModels field if non-nil, zero value otherwise.

### GetSupportedModelsOk

`func (o *SoftwarerepositoryCategoryMapper) GetSupportedModelsOk() (*[]string, bool)`

GetSupportedModelsOk returns a tuple with the SupportedModels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedModels

`func (o *SoftwarerepositoryCategoryMapper) SetSupportedModels(v []string)`

SetSupportedModels sets SupportedModels field to given value.

### HasSupportedModels

`func (o *SoftwarerepositoryCategoryMapper) HasSupportedModels() bool`

HasSupportedModels returns a boolean if a field has been set.

### SetSupportedModelsNil

`func (o *SoftwarerepositoryCategoryMapper) SetSupportedModelsNil(b bool)`

 SetSupportedModelsNil sets the value for SupportedModels to be an explicit nil

### UnsetSupportedModels
`func (o *SoftwarerepositoryCategoryMapper) UnsetSupportedModels()`

UnsetSupportedModels ensures that no value is present for SupportedModels, not even an explicit nil
### GetSwId

`func (o *SoftwarerepositoryCategoryMapper) GetSwId() string`

GetSwId returns the SwId field if non-nil, zero value otherwise.

### GetSwIdOk

`func (o *SoftwarerepositoryCategoryMapper) GetSwIdOk() (*string, bool)`

GetSwIdOk returns a tuple with the SwId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwId

`func (o *SoftwarerepositoryCategoryMapper) SetSwId(v string)`

SetSwId sets SwId field to given value.

### HasSwId

`func (o *SoftwarerepositoryCategoryMapper) HasSwId() bool`

HasSwId returns a boolean if a field has been set.

### GetTagTypes

`func (o *SoftwarerepositoryCategoryMapper) GetTagTypes() []string`

GetTagTypes returns the TagTypes field if non-nil, zero value otherwise.

### GetTagTypesOk

`func (o *SoftwarerepositoryCategoryMapper) GetTagTypesOk() (*[]string, bool)`

GetTagTypesOk returns a tuple with the TagTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagTypes

`func (o *SoftwarerepositoryCategoryMapper) SetTagTypes(v []string)`

SetTagTypes sets TagTypes field to given value.

### HasTagTypes

`func (o *SoftwarerepositoryCategoryMapper) HasTagTypes() bool`

HasTagTypes returns a boolean if a field has been set.

### SetTagTypesNil

`func (o *SoftwarerepositoryCategoryMapper) SetTagTypesNil(b bool)`

 SetTagTypesNil sets the value for TagTypes to be an explicit nil

### UnsetTagTypes
`func (o *SoftwarerepositoryCategoryMapper) UnsetTagTypes()`

UnsetTagTypes ensures that no value is present for TagTypes, not even an explicit nil
### GetVersion

`func (o *SoftwarerepositoryCategoryMapper) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SoftwarerepositoryCategoryMapper) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SoftwarerepositoryCategoryMapper) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SoftwarerepositoryCategoryMapper) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



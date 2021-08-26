# HyperflexTrackedFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.TrackedFile"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.TrackedFile"]
**FilePath** | Pointer to [**NullableHyperflexFilePath**](HyperflexFilePath.md) |  | [optional] 
**FileType** | Pointer to **string** | File type for the tracked file. | [optional] [readonly] 

## Methods

### NewHyperflexTrackedFile

`func NewHyperflexTrackedFile(classId string, objectType string, ) *HyperflexTrackedFile`

NewHyperflexTrackedFile instantiates a new HyperflexTrackedFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexTrackedFileWithDefaults

`func NewHyperflexTrackedFileWithDefaults() *HyperflexTrackedFile`

NewHyperflexTrackedFileWithDefaults instantiates a new HyperflexTrackedFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexTrackedFile) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexTrackedFile) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexTrackedFile) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexTrackedFile) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexTrackedFile) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexTrackedFile) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetFilePath

`func (o *HyperflexTrackedFile) GetFilePath() HyperflexFilePath`

GetFilePath returns the FilePath field if non-nil, zero value otherwise.

### GetFilePathOk

`func (o *HyperflexTrackedFile) GetFilePathOk() (*HyperflexFilePath, bool)`

GetFilePathOk returns a tuple with the FilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePath

`func (o *HyperflexTrackedFile) SetFilePath(v HyperflexFilePath)`

SetFilePath sets FilePath field to given value.

### HasFilePath

`func (o *HyperflexTrackedFile) HasFilePath() bool`

HasFilePath returns a boolean if a field has been set.

### SetFilePathNil

`func (o *HyperflexTrackedFile) SetFilePathNil(b bool)`

 SetFilePathNil sets the value for FilePath to be an explicit nil

### UnsetFilePath
`func (o *HyperflexTrackedFile) UnsetFilePath()`

UnsetFilePath ensures that no value is present for FilePath, not even an explicit nil
### GetFileType

`func (o *HyperflexTrackedFile) GetFileType() string`

GetFileType returns the FileType field if non-nil, zero value otherwise.

### GetFileTypeOk

`func (o *HyperflexTrackedFile) GetFileTypeOk() (*string, bool)`

GetFileTypeOk returns a tuple with the FileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileType

`func (o *HyperflexTrackedFile) SetFileType(v string)`

SetFileType sets FileType field to given value.

### HasFileType

`func (o *HyperflexTrackedFile) HasFileType() bool`

HasFileType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



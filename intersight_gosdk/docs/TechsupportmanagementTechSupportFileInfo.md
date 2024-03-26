# TechsupportmanagementTechSupportFileInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "techsupportmanagement.TechSupportFileInfo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "techsupportmanagement.TechSupportFileInfo"]
**FileName** | Pointer to **string** | The name of the techsupport file. | [optional] [readonly] 
**FileSize** | Pointer to **int64** | Techsupport file size in bytes. | [optional] [readonly] 
**TechsupportDownloadUrl** | Pointer to **string** | The Url to download the techsupport file. | [optional] [readonly] 
**UploadStatus** | Pointer to **string** | The upload status of the techsupport file. | [optional] [readonly] [default to "UploadQueued"]

## Methods

### NewTechsupportmanagementTechSupportFileInfo

`func NewTechsupportmanagementTechSupportFileInfo(classId string, objectType string, ) *TechsupportmanagementTechSupportFileInfo`

NewTechsupportmanagementTechSupportFileInfo instantiates a new TechsupportmanagementTechSupportFileInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTechsupportmanagementTechSupportFileInfoWithDefaults

`func NewTechsupportmanagementTechSupportFileInfoWithDefaults() *TechsupportmanagementTechSupportFileInfo`

NewTechsupportmanagementTechSupportFileInfoWithDefaults instantiates a new TechsupportmanagementTechSupportFileInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *TechsupportmanagementTechSupportFileInfo) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *TechsupportmanagementTechSupportFileInfo) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *TechsupportmanagementTechSupportFileInfo) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *TechsupportmanagementTechSupportFileInfo) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *TechsupportmanagementTechSupportFileInfo) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *TechsupportmanagementTechSupportFileInfo) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetFileName

`func (o *TechsupportmanagementTechSupportFileInfo) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *TechsupportmanagementTechSupportFileInfo) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *TechsupportmanagementTechSupportFileInfo) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *TechsupportmanagementTechSupportFileInfo) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetFileSize

`func (o *TechsupportmanagementTechSupportFileInfo) GetFileSize() int64`

GetFileSize returns the FileSize field if non-nil, zero value otherwise.

### GetFileSizeOk

`func (o *TechsupportmanagementTechSupportFileInfo) GetFileSizeOk() (*int64, bool)`

GetFileSizeOk returns a tuple with the FileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSize

`func (o *TechsupportmanagementTechSupportFileInfo) SetFileSize(v int64)`

SetFileSize sets FileSize field to given value.

### HasFileSize

`func (o *TechsupportmanagementTechSupportFileInfo) HasFileSize() bool`

HasFileSize returns a boolean if a field has been set.

### GetTechsupportDownloadUrl

`func (o *TechsupportmanagementTechSupportFileInfo) GetTechsupportDownloadUrl() string`

GetTechsupportDownloadUrl returns the TechsupportDownloadUrl field if non-nil, zero value otherwise.

### GetTechsupportDownloadUrlOk

`func (o *TechsupportmanagementTechSupportFileInfo) GetTechsupportDownloadUrlOk() (*string, bool)`

GetTechsupportDownloadUrlOk returns a tuple with the TechsupportDownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechsupportDownloadUrl

`func (o *TechsupportmanagementTechSupportFileInfo) SetTechsupportDownloadUrl(v string)`

SetTechsupportDownloadUrl sets TechsupportDownloadUrl field to given value.

### HasTechsupportDownloadUrl

`func (o *TechsupportmanagementTechSupportFileInfo) HasTechsupportDownloadUrl() bool`

HasTechsupportDownloadUrl returns a boolean if a field has been set.

### GetUploadStatus

`func (o *TechsupportmanagementTechSupportFileInfo) GetUploadStatus() string`

GetUploadStatus returns the UploadStatus field if non-nil, zero value otherwise.

### GetUploadStatusOk

`func (o *TechsupportmanagementTechSupportFileInfo) GetUploadStatusOk() (*string, bool)`

GetUploadStatusOk returns a tuple with the UploadStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadStatus

`func (o *TechsupportmanagementTechSupportFileInfo) SetUploadStatus(v string)`

SetUploadStatus sets UploadStatus field to given value.

### HasUploadStatus

`func (o *TechsupportmanagementTechSupportFileInfo) HasUploadStatus() bool`

HasUploadStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



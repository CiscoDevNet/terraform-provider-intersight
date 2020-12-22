# OnpremImagePackageAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "onprem.ImagePackage"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "onprem.ImagePackage"]
**FilePath** | Pointer to **string** | Optional file path of the image package. | [optional] [readonly] 
**FileSha** | Pointer to **string** | Image file&#39;s fingerprint. Fingerprint is calculated using SHA256 algorithm. | [optional] [readonly] 
**FileSize** | Pointer to **int64** | Image file size in bytes. | [optional] [readonly] 
**FileTime** | Pointer to **time.Time** | Image file&#39;s last modified date and time. | [optional] [readonly] 
**Filename** | Pointer to **string** | Filename of the image package. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the software image package. | [optional] [readonly] 
**PackageType** | Pointer to **string** | Image package type (e.g. service, system etc.). | [optional] [readonly] 
**Version** | Pointer to **string** | Image package version string. | [optional] [readonly] 

## Methods

### NewOnpremImagePackageAllOf

`func NewOnpremImagePackageAllOf(classId string, objectType string, ) *OnpremImagePackageAllOf`

NewOnpremImagePackageAllOf instantiates a new OnpremImagePackageAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnpremImagePackageAllOfWithDefaults

`func NewOnpremImagePackageAllOfWithDefaults() *OnpremImagePackageAllOf`

NewOnpremImagePackageAllOfWithDefaults instantiates a new OnpremImagePackageAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *OnpremImagePackageAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *OnpremImagePackageAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *OnpremImagePackageAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *OnpremImagePackageAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *OnpremImagePackageAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *OnpremImagePackageAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetFilePath

`func (o *OnpremImagePackageAllOf) GetFilePath() string`

GetFilePath returns the FilePath field if non-nil, zero value otherwise.

### GetFilePathOk

`func (o *OnpremImagePackageAllOf) GetFilePathOk() (*string, bool)`

GetFilePathOk returns a tuple with the FilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePath

`func (o *OnpremImagePackageAllOf) SetFilePath(v string)`

SetFilePath sets FilePath field to given value.

### HasFilePath

`func (o *OnpremImagePackageAllOf) HasFilePath() bool`

HasFilePath returns a boolean if a field has been set.

### GetFileSha

`func (o *OnpremImagePackageAllOf) GetFileSha() string`

GetFileSha returns the FileSha field if non-nil, zero value otherwise.

### GetFileShaOk

`func (o *OnpremImagePackageAllOf) GetFileShaOk() (*string, bool)`

GetFileShaOk returns a tuple with the FileSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSha

`func (o *OnpremImagePackageAllOf) SetFileSha(v string)`

SetFileSha sets FileSha field to given value.

### HasFileSha

`func (o *OnpremImagePackageAllOf) HasFileSha() bool`

HasFileSha returns a boolean if a field has been set.

### GetFileSize

`func (o *OnpremImagePackageAllOf) GetFileSize() int64`

GetFileSize returns the FileSize field if non-nil, zero value otherwise.

### GetFileSizeOk

`func (o *OnpremImagePackageAllOf) GetFileSizeOk() (*int64, bool)`

GetFileSizeOk returns a tuple with the FileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSize

`func (o *OnpremImagePackageAllOf) SetFileSize(v int64)`

SetFileSize sets FileSize field to given value.

### HasFileSize

`func (o *OnpremImagePackageAllOf) HasFileSize() bool`

HasFileSize returns a boolean if a field has been set.

### GetFileTime

`func (o *OnpremImagePackageAllOf) GetFileTime() time.Time`

GetFileTime returns the FileTime field if non-nil, zero value otherwise.

### GetFileTimeOk

`func (o *OnpremImagePackageAllOf) GetFileTimeOk() (*time.Time, bool)`

GetFileTimeOk returns a tuple with the FileTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileTime

`func (o *OnpremImagePackageAllOf) SetFileTime(v time.Time)`

SetFileTime sets FileTime field to given value.

### HasFileTime

`func (o *OnpremImagePackageAllOf) HasFileTime() bool`

HasFileTime returns a boolean if a field has been set.

### GetFilename

`func (o *OnpremImagePackageAllOf) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *OnpremImagePackageAllOf) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *OnpremImagePackageAllOf) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *OnpremImagePackageAllOf) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetName

`func (o *OnpremImagePackageAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OnpremImagePackageAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OnpremImagePackageAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OnpremImagePackageAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPackageType

`func (o *OnpremImagePackageAllOf) GetPackageType() string`

GetPackageType returns the PackageType field if non-nil, zero value otherwise.

### GetPackageTypeOk

`func (o *OnpremImagePackageAllOf) GetPackageTypeOk() (*string, bool)`

GetPackageTypeOk returns a tuple with the PackageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageType

`func (o *OnpremImagePackageAllOf) SetPackageType(v string)`

SetPackageType sets PackageType field to given value.

### HasPackageType

`func (o *OnpremImagePackageAllOf) HasPackageType() bool`

HasPackageType returns a boolean if a field has been set.

### GetVersion

`func (o *OnpremImagePackageAllOf) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *OnpremImagePackageAllOf) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *OnpremImagePackageAllOf) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *OnpremImagePackageAllOf) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



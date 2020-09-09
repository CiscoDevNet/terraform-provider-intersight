# OnpremImagePackage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilePath** | Pointer to **string** | Optional file path of the image package. | [optional] [readonly] 
**FileSha** | Pointer to **string** | Image file&#39;s fingerprint. Fingerprint is calculated using SHA256 algorithm. | [optional] [readonly] 
**FileSize** | Pointer to **int64** | Image file size in bytes. | [optional] [readonly] 
**FileTime** | Pointer to [**time.Time**](time.Time.md) | Image file&#39;s last modified date and time. | [optional] [readonly] 
**Filename** | Pointer to **string** | Filename of the image package. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the software image package. | [optional] [readonly] 
**PackageType** | Pointer to **string** | Image package type (e.g. service, system etc.). | [optional] [readonly] 
**Version** | Pointer to **string** | Image package version string. | [optional] [readonly] 

## Methods

### NewOnpremImagePackage

`func NewOnpremImagePackage() *OnpremImagePackage`

NewOnpremImagePackage instantiates a new OnpremImagePackage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnpremImagePackageWithDefaults

`func NewOnpremImagePackageWithDefaults() *OnpremImagePackage`

NewOnpremImagePackageWithDefaults instantiates a new OnpremImagePackage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilePath

`func (o *OnpremImagePackage) GetFilePath() string`

GetFilePath returns the FilePath field if non-nil, zero value otherwise.

### GetFilePathOk

`func (o *OnpremImagePackage) GetFilePathOk() (*string, bool)`

GetFilePathOk returns a tuple with the FilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePath

`func (o *OnpremImagePackage) SetFilePath(v string)`

SetFilePath sets FilePath field to given value.

### HasFilePath

`func (o *OnpremImagePackage) HasFilePath() bool`

HasFilePath returns a boolean if a field has been set.

### GetFileSha

`func (o *OnpremImagePackage) GetFileSha() string`

GetFileSha returns the FileSha field if non-nil, zero value otherwise.

### GetFileShaOk

`func (o *OnpremImagePackage) GetFileShaOk() (*string, bool)`

GetFileShaOk returns a tuple with the FileSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSha

`func (o *OnpremImagePackage) SetFileSha(v string)`

SetFileSha sets FileSha field to given value.

### HasFileSha

`func (o *OnpremImagePackage) HasFileSha() bool`

HasFileSha returns a boolean if a field has been set.

### GetFileSize

`func (o *OnpremImagePackage) GetFileSize() int64`

GetFileSize returns the FileSize field if non-nil, zero value otherwise.

### GetFileSizeOk

`func (o *OnpremImagePackage) GetFileSizeOk() (*int64, bool)`

GetFileSizeOk returns a tuple with the FileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSize

`func (o *OnpremImagePackage) SetFileSize(v int64)`

SetFileSize sets FileSize field to given value.

### HasFileSize

`func (o *OnpremImagePackage) HasFileSize() bool`

HasFileSize returns a boolean if a field has been set.

### GetFileTime

`func (o *OnpremImagePackage) GetFileTime() time.Time`

GetFileTime returns the FileTime field if non-nil, zero value otherwise.

### GetFileTimeOk

`func (o *OnpremImagePackage) GetFileTimeOk() (*time.Time, bool)`

GetFileTimeOk returns a tuple with the FileTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileTime

`func (o *OnpremImagePackage) SetFileTime(v time.Time)`

SetFileTime sets FileTime field to given value.

### HasFileTime

`func (o *OnpremImagePackage) HasFileTime() bool`

HasFileTime returns a boolean if a field has been set.

### GetFilename

`func (o *OnpremImagePackage) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *OnpremImagePackage) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *OnpremImagePackage) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *OnpremImagePackage) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetName

`func (o *OnpremImagePackage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OnpremImagePackage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OnpremImagePackage) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OnpremImagePackage) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPackageType

`func (o *OnpremImagePackage) GetPackageType() string`

GetPackageType returns the PackageType field if non-nil, zero value otherwise.

### GetPackageTypeOk

`func (o *OnpremImagePackage) GetPackageTypeOk() (*string, bool)`

GetPackageTypeOk returns a tuple with the PackageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageType

`func (o *OnpremImagePackage) SetPackageType(v string)`

SetPackageType sets PackageType field to given value.

### HasPackageType

`func (o *OnpremImagePackage) HasPackageType() bool`

HasPackageType returns a boolean if a field has been set.

### GetVersion

`func (o *OnpremImagePackage) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *OnpremImagePackage) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *OnpremImagePackage) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *OnpremImagePackage) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



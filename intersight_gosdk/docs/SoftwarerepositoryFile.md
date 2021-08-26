# SoftwarerepositoryFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**Description** | Pointer to **string** | User provided description about the file. Cisco provided description for image inventoried from a Cisco repository. | [optional] 
**DownloadCount** | Pointer to **int64** | The number of times this file has been downloaded from the local repository. It is used by the repository monitoring process to determine the files that are to be evicted from the cache. | [optional] [readonly] 
**ImportAction** | Pointer to **string** | The action to be performed on the imported file. If &#39;PreCache&#39; is set, the image will be cached in Appliance. Applicable in Intersight appliance deployment. If &#39;Evict&#39; is set, the cached file will be removed. Applicable in Intersight appliance deployment. If &#39;GeneratePreSignedUploadUrl&#39; is set, generates pre signed URL (s) for the file to be imported into the repository. Applicable for local machine source. The URL (s) will be populated under LocalMachine file server. If &#39;CompleteImportProcess&#39; is set, the ImportState is marked as &#39;Imported&#39;. Applicable for local machine source. If &#39;Cancel&#39; is set, the ImportState is marked as &#39;Failed&#39;. Applicable for local machine source. * &#x60;None&#x60; - No action should be taken on the imported file. * &#x60;GeneratePreSignedUploadUrl&#x60; - Generate pre signed URL of file for importing into the repository. * &#x60;GeneratePreSignedDownloadUrl&#x60; - Generate pre signed URL of file in the repository to download. * &#x60;CompleteImportProcess&#x60; - Mark that the import process of the file into the repository is complete. * &#x60;MarkImportFailed&#x60; - Mark to indicate that the import process of the file into the repository failed. * &#x60;PreCache&#x60; - Cache the file into the Intersight Appliance. * &#x60;Cancel&#x60; - The cancel import process for the file into the repository. * &#x60;Extract&#x60; - The action to extract the file in the external repository. * &#x60;Evict&#x60; - Evict the cached file from the Intersight Appliance. | [optional] [default to "None"]
**ImportState** | Pointer to **string** | The state  of this file in the repository or Appliance. The importState is updated during the import operation and as part of the repository monitoring process. * &#x60;ReadyForImport&#x60; - The image is ready to be imported into the repository. * &#x60;Importing&#x60; - The image is being imported into the repository. * &#x60;Imported&#x60; - The image has been extracted and imported into the repository. * &#x60;PendingExtraction&#x60; - Indicates that the image has been imported but not extracted in the repository. * &#x60;Extracting&#x60; - Indicates that the image is being extracted into the repository. * &#x60;Extracted&#x60; - Indicates that the image has been extracted into the repository. * &#x60;Failed&#x60; - The image import from an external source to the repository has failed. * &#x60;MetaOnly&#x60; - The image is present in an external repository. * &#x60;ReadyForCache&#x60; - The image is ready to be cached into the Intersight Appliance. * &#x60;Caching&#x60; - Indicates that the image is being cached into the Intersight Appliance or endpoint cache. * &#x60;Cached&#x60; - Indicates that the image has been cached into the Intersight Appliance or endpoint cache. * &#x60;CachingFailed&#x60; - Indicates that the image caching into the Intersight Appliance failed or endpoint cache. * &#x60;Corrupted&#x60; - Indicates that the image in the local repository (or endpoint cache) has been corrupted after it was cached. * &#x60;Evicted&#x60; - Indicates that the image has been evicted from the Intersight Appliance (or endpoint cache) to reclaim storage space. * &#x60;Invalid&#x60; - Indicates that the corresponding distributable MO has been removed from the backend. This can be due to unpublishing of an image. | [optional] [readonly] [default to "ReadyForImport"]
**ImportedTime** | Pointer to **time.Time** | The time at which this image or file was imported/cached into the repositry. if the &#39;ImportState&#39; is &#39;Imported&#39;, the time at which this image or file was imported. if the &#39;ImportState&#39; is &#39;Cached&#39;, the time at which this image or file was cached. | [optional] [readonly] 
**LastAccessTime** | Pointer to **time.Time** | The time at which this file was last downloaded from the local repository. It is used by the repository monitoring process to determine the files that are to be evicted from the cache. | [optional] [readonly] 
**Md5eTag** | Pointer to **string** | The MD5 ETag for a file that is stored in Intersight repository or in the appliance cache. Warning - MD5 is currently broken and this will be migrated to SHA shortly. | [optional] 
**Md5sum** | Pointer to **string** | The md5sum checksum of the file. This information is available for all Cisco distributed images and files imported to the local repository. | [optional] 
**Name** | Pointer to **string** | The name of the file. It is populated as part of the image import operation. | [optional] 
**ReleaseDate** | Pointer to **time.Time** | The date on which the file was released or distributed by its vendor. | [optional] 
**Sha512sum** | Pointer to **string** | The sha512sum of the file. This information is available for all Cisco distributed images and files imported to the local repository. | [optional] 
**Size** | Pointer to **int64** | The size (in bytes) of the file. This information is available for all Cisco distributed images and files imported to the local repository. | [optional] 
**SoftwareAdvisoryUrl** | Pointer to **string** | The software advisory, if any, provided by the vendor for this file. | [optional] 
**Source** | Pointer to [**NullableSoftwarerepositoryFileServer**](SoftwarerepositoryFileServer.md) |  | [optional] 
**Version** | Pointer to **string** | Vendor provided version for the file. | [optional] 

## Methods

### NewSoftwarerepositoryFile

`func NewSoftwarerepositoryFile(classId string, objectType string, ) *SoftwarerepositoryFile`

NewSoftwarerepositoryFile instantiates a new SoftwarerepositoryFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoftwarerepositoryFileWithDefaults

`func NewSoftwarerepositoryFileWithDefaults() *SoftwarerepositoryFile`

NewSoftwarerepositoryFileWithDefaults instantiates a new SoftwarerepositoryFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *SoftwarerepositoryFile) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *SoftwarerepositoryFile) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *SoftwarerepositoryFile) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *SoftwarerepositoryFile) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *SoftwarerepositoryFile) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *SoftwarerepositoryFile) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDescription

`func (o *SoftwarerepositoryFile) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SoftwarerepositoryFile) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SoftwarerepositoryFile) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SoftwarerepositoryFile) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDownloadCount

`func (o *SoftwarerepositoryFile) GetDownloadCount() int64`

GetDownloadCount returns the DownloadCount field if non-nil, zero value otherwise.

### GetDownloadCountOk

`func (o *SoftwarerepositoryFile) GetDownloadCountOk() (*int64, bool)`

GetDownloadCountOk returns a tuple with the DownloadCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadCount

`func (o *SoftwarerepositoryFile) SetDownloadCount(v int64)`

SetDownloadCount sets DownloadCount field to given value.

### HasDownloadCount

`func (o *SoftwarerepositoryFile) HasDownloadCount() bool`

HasDownloadCount returns a boolean if a field has been set.

### GetImportAction

`func (o *SoftwarerepositoryFile) GetImportAction() string`

GetImportAction returns the ImportAction field if non-nil, zero value otherwise.

### GetImportActionOk

`func (o *SoftwarerepositoryFile) GetImportActionOk() (*string, bool)`

GetImportActionOk returns a tuple with the ImportAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportAction

`func (o *SoftwarerepositoryFile) SetImportAction(v string)`

SetImportAction sets ImportAction field to given value.

### HasImportAction

`func (o *SoftwarerepositoryFile) HasImportAction() bool`

HasImportAction returns a boolean if a field has been set.

### GetImportState

`func (o *SoftwarerepositoryFile) GetImportState() string`

GetImportState returns the ImportState field if non-nil, zero value otherwise.

### GetImportStateOk

`func (o *SoftwarerepositoryFile) GetImportStateOk() (*string, bool)`

GetImportStateOk returns a tuple with the ImportState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportState

`func (o *SoftwarerepositoryFile) SetImportState(v string)`

SetImportState sets ImportState field to given value.

### HasImportState

`func (o *SoftwarerepositoryFile) HasImportState() bool`

HasImportState returns a boolean if a field has been set.

### GetImportedTime

`func (o *SoftwarerepositoryFile) GetImportedTime() time.Time`

GetImportedTime returns the ImportedTime field if non-nil, zero value otherwise.

### GetImportedTimeOk

`func (o *SoftwarerepositoryFile) GetImportedTimeOk() (*time.Time, bool)`

GetImportedTimeOk returns a tuple with the ImportedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportedTime

`func (o *SoftwarerepositoryFile) SetImportedTime(v time.Time)`

SetImportedTime sets ImportedTime field to given value.

### HasImportedTime

`func (o *SoftwarerepositoryFile) HasImportedTime() bool`

HasImportedTime returns a boolean if a field has been set.

### GetLastAccessTime

`func (o *SoftwarerepositoryFile) GetLastAccessTime() time.Time`

GetLastAccessTime returns the LastAccessTime field if non-nil, zero value otherwise.

### GetLastAccessTimeOk

`func (o *SoftwarerepositoryFile) GetLastAccessTimeOk() (*time.Time, bool)`

GetLastAccessTimeOk returns a tuple with the LastAccessTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAccessTime

`func (o *SoftwarerepositoryFile) SetLastAccessTime(v time.Time)`

SetLastAccessTime sets LastAccessTime field to given value.

### HasLastAccessTime

`func (o *SoftwarerepositoryFile) HasLastAccessTime() bool`

HasLastAccessTime returns a boolean if a field has been set.

### GetMd5eTag

`func (o *SoftwarerepositoryFile) GetMd5eTag() string`

GetMd5eTag returns the Md5eTag field if non-nil, zero value otherwise.

### GetMd5eTagOk

`func (o *SoftwarerepositoryFile) GetMd5eTagOk() (*string, bool)`

GetMd5eTagOk returns a tuple with the Md5eTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5eTag

`func (o *SoftwarerepositoryFile) SetMd5eTag(v string)`

SetMd5eTag sets Md5eTag field to given value.

### HasMd5eTag

`func (o *SoftwarerepositoryFile) HasMd5eTag() bool`

HasMd5eTag returns a boolean if a field has been set.

### GetMd5sum

`func (o *SoftwarerepositoryFile) GetMd5sum() string`

GetMd5sum returns the Md5sum field if non-nil, zero value otherwise.

### GetMd5sumOk

`func (o *SoftwarerepositoryFile) GetMd5sumOk() (*string, bool)`

GetMd5sumOk returns a tuple with the Md5sum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5sum

`func (o *SoftwarerepositoryFile) SetMd5sum(v string)`

SetMd5sum sets Md5sum field to given value.

### HasMd5sum

`func (o *SoftwarerepositoryFile) HasMd5sum() bool`

HasMd5sum returns a boolean if a field has been set.

### GetName

`func (o *SoftwarerepositoryFile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SoftwarerepositoryFile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SoftwarerepositoryFile) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SoftwarerepositoryFile) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReleaseDate

`func (o *SoftwarerepositoryFile) GetReleaseDate() time.Time`

GetReleaseDate returns the ReleaseDate field if non-nil, zero value otherwise.

### GetReleaseDateOk

`func (o *SoftwarerepositoryFile) GetReleaseDateOk() (*time.Time, bool)`

GetReleaseDateOk returns a tuple with the ReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDate

`func (o *SoftwarerepositoryFile) SetReleaseDate(v time.Time)`

SetReleaseDate sets ReleaseDate field to given value.

### HasReleaseDate

`func (o *SoftwarerepositoryFile) HasReleaseDate() bool`

HasReleaseDate returns a boolean if a field has been set.

### GetSha512sum

`func (o *SoftwarerepositoryFile) GetSha512sum() string`

GetSha512sum returns the Sha512sum field if non-nil, zero value otherwise.

### GetSha512sumOk

`func (o *SoftwarerepositoryFile) GetSha512sumOk() (*string, bool)`

GetSha512sumOk returns a tuple with the Sha512sum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha512sum

`func (o *SoftwarerepositoryFile) SetSha512sum(v string)`

SetSha512sum sets Sha512sum field to given value.

### HasSha512sum

`func (o *SoftwarerepositoryFile) HasSha512sum() bool`

HasSha512sum returns a boolean if a field has been set.

### GetSize

`func (o *SoftwarerepositoryFile) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *SoftwarerepositoryFile) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *SoftwarerepositoryFile) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *SoftwarerepositoryFile) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSoftwareAdvisoryUrl

`func (o *SoftwarerepositoryFile) GetSoftwareAdvisoryUrl() string`

GetSoftwareAdvisoryUrl returns the SoftwareAdvisoryUrl field if non-nil, zero value otherwise.

### GetSoftwareAdvisoryUrlOk

`func (o *SoftwarerepositoryFile) GetSoftwareAdvisoryUrlOk() (*string, bool)`

GetSoftwareAdvisoryUrlOk returns a tuple with the SoftwareAdvisoryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareAdvisoryUrl

`func (o *SoftwarerepositoryFile) SetSoftwareAdvisoryUrl(v string)`

SetSoftwareAdvisoryUrl sets SoftwareAdvisoryUrl field to given value.

### HasSoftwareAdvisoryUrl

`func (o *SoftwarerepositoryFile) HasSoftwareAdvisoryUrl() bool`

HasSoftwareAdvisoryUrl returns a boolean if a field has been set.

### GetSource

`func (o *SoftwarerepositoryFile) GetSource() SoftwarerepositoryFileServer`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *SoftwarerepositoryFile) GetSourceOk() (*SoftwarerepositoryFileServer, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *SoftwarerepositoryFile) SetSource(v SoftwarerepositoryFileServer)`

SetSource sets Source field to given value.

### HasSource

`func (o *SoftwarerepositoryFile) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *SoftwarerepositoryFile) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *SoftwarerepositoryFile) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetVersion

`func (o *SoftwarerepositoryFile) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SoftwarerepositoryFile) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SoftwarerepositoryFile) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SoftwarerepositoryFile) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



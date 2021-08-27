# SoftwarerepositoryCachedImage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "softwarerepository.CachedImage"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "softwarerepository.CachedImage"]
**Action** | Pointer to **string** | The action to be performed on the imported file. If &#39;PreCache&#39; is set, the image will be cached in endpoint. If &#39;Evict&#39; is set, the cached file will be removed. Applicable in Intersight appliance deployment. If &#39;Cancel&#39; is set, the ImportState is marked as &#39;Failed&#39;. Applicable for local machine source. * &#x60;None&#x60; - No action should be taken on the imported file. * &#x60;GeneratePreSignedUploadUrl&#x60; - Generate pre signed URL of file for importing into the repository. * &#x60;GeneratePreSignedDownloadUrl&#x60; - Generate pre signed URL of file in the repository to download. * &#x60;CompleteImportProcess&#x60; - Mark that the import process of the file into the repository is complete. * &#x60;MarkImportFailed&#x60; - Mark to indicate that the import process of the file into the repository failed. * &#x60;PreCache&#x60; - Cache the file into the Intersight Appliance. * &#x60;Cancel&#x60; - The cancel import process for the file into the repository. * &#x60;Extract&#x60; - The action to extract the file in the external repository. * &#x60;Evict&#x60; - Evict the cached file from the Intersight Appliance. | [optional] [default to "None"]
**CacheState** | Pointer to **string** | The state  of this file in the endpoint The importState is updated during the cache operation and as part of the cache monitoring process. * &#x60;ReadyForImport&#x60; - The image is ready to be imported into the repository. * &#x60;Importing&#x60; - The image is being imported into the repository. * &#x60;Imported&#x60; - The image has been extracted and imported into the repository. * &#x60;PendingExtraction&#x60; - Indicates that the image has been imported but not extracted in the repository. * &#x60;Extracting&#x60; - Indicates that the image is being extracted into the repository. * &#x60;Extracted&#x60; - Indicates that the image has been extracted into the repository. * &#x60;Failed&#x60; - The image import from an external source to the repository has failed. * &#x60;MetaOnly&#x60; - The image is present in an external repository. * &#x60;ReadyForCache&#x60; - The image is ready to be cached into the Intersight Appliance. * &#x60;Caching&#x60; - Indicates that the image is being cached into the Intersight Appliance or endpoint cache. * &#x60;Cached&#x60; - Indicates that the image has been cached into the Intersight Appliance or endpoint cache. * &#x60;CachingFailed&#x60; - Indicates that the image caching into the Intersight Appliance failed or endpoint cache. * &#x60;Corrupted&#x60; - Indicates that the image in the local repository (or endpoint cache) has been corrupted after it was cached. * &#x60;Evicted&#x60; - Indicates that the image has been evicted from the Intersight Appliance (or endpoint cache) to reclaim storage space. * &#x60;Invalid&#x60; - Indicates that the corresponding distributable MO has been removed from the backend. This can be due to unpublishing of an image. | [optional] [readonly] [default to "ReadyForImport"]
**CachedTime** | Pointer to **time.Time** | The time when the image or file was cached into the FI storage. | [optional] [readonly] 
**LastAccessTime** | Pointer to **time.Time** | Used by the cache monitoring process to determine the files that are to be evicted from the cache. | [optional] [readonly] 
**Md5sum** | Pointer to **string** | The MD5 sum of the firmware image that will be used by the endpoint to validate the integrity of the image. | [optional] [readonly] 
**OriginalSha512sum** | Pointer to **string** | The actual sha512sum of the cached image. | [optional] [readonly] 
**Path** | Pointer to **string** | The absolute path of the imported file in the endpoint. | [optional] [readonly] 
**RegisteredWorkflows** | Pointer to **[]string** |  | [optional] 
**UsedCount** | Pointer to **int64** | The number of times this file has been used to copy or upgrade or install actions. Used by the cache monitoring process to determine the files to be evicted from the cache. | [optional] [readonly] 
**File** | Pointer to [**SoftwarerepositoryFileRelationship**](SoftwarerepositoryFileRelationship.md) |  | [optional] 
**NetworkElement** | Pointer to [**NetworkElementRelationship**](NetworkElementRelationship.md) |  | [optional] 

## Methods

### NewSoftwarerepositoryCachedImage

`func NewSoftwarerepositoryCachedImage(classId string, objectType string, ) *SoftwarerepositoryCachedImage`

NewSoftwarerepositoryCachedImage instantiates a new SoftwarerepositoryCachedImage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoftwarerepositoryCachedImageWithDefaults

`func NewSoftwarerepositoryCachedImageWithDefaults() *SoftwarerepositoryCachedImage`

NewSoftwarerepositoryCachedImageWithDefaults instantiates a new SoftwarerepositoryCachedImage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *SoftwarerepositoryCachedImage) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *SoftwarerepositoryCachedImage) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *SoftwarerepositoryCachedImage) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *SoftwarerepositoryCachedImage) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *SoftwarerepositoryCachedImage) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *SoftwarerepositoryCachedImage) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAction

`func (o *SoftwarerepositoryCachedImage) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *SoftwarerepositoryCachedImage) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *SoftwarerepositoryCachedImage) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *SoftwarerepositoryCachedImage) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetCacheState

`func (o *SoftwarerepositoryCachedImage) GetCacheState() string`

GetCacheState returns the CacheState field if non-nil, zero value otherwise.

### GetCacheStateOk

`func (o *SoftwarerepositoryCachedImage) GetCacheStateOk() (*string, bool)`

GetCacheStateOk returns a tuple with the CacheState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheState

`func (o *SoftwarerepositoryCachedImage) SetCacheState(v string)`

SetCacheState sets CacheState field to given value.

### HasCacheState

`func (o *SoftwarerepositoryCachedImage) HasCacheState() bool`

HasCacheState returns a boolean if a field has been set.

### GetCachedTime

`func (o *SoftwarerepositoryCachedImage) GetCachedTime() time.Time`

GetCachedTime returns the CachedTime field if non-nil, zero value otherwise.

### GetCachedTimeOk

`func (o *SoftwarerepositoryCachedImage) GetCachedTimeOk() (*time.Time, bool)`

GetCachedTimeOk returns a tuple with the CachedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCachedTime

`func (o *SoftwarerepositoryCachedImage) SetCachedTime(v time.Time)`

SetCachedTime sets CachedTime field to given value.

### HasCachedTime

`func (o *SoftwarerepositoryCachedImage) HasCachedTime() bool`

HasCachedTime returns a boolean if a field has been set.

### GetLastAccessTime

`func (o *SoftwarerepositoryCachedImage) GetLastAccessTime() time.Time`

GetLastAccessTime returns the LastAccessTime field if non-nil, zero value otherwise.

### GetLastAccessTimeOk

`func (o *SoftwarerepositoryCachedImage) GetLastAccessTimeOk() (*time.Time, bool)`

GetLastAccessTimeOk returns a tuple with the LastAccessTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAccessTime

`func (o *SoftwarerepositoryCachedImage) SetLastAccessTime(v time.Time)`

SetLastAccessTime sets LastAccessTime field to given value.

### HasLastAccessTime

`func (o *SoftwarerepositoryCachedImage) HasLastAccessTime() bool`

HasLastAccessTime returns a boolean if a field has been set.

### GetMd5sum

`func (o *SoftwarerepositoryCachedImage) GetMd5sum() string`

GetMd5sum returns the Md5sum field if non-nil, zero value otherwise.

### GetMd5sumOk

`func (o *SoftwarerepositoryCachedImage) GetMd5sumOk() (*string, bool)`

GetMd5sumOk returns a tuple with the Md5sum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5sum

`func (o *SoftwarerepositoryCachedImage) SetMd5sum(v string)`

SetMd5sum sets Md5sum field to given value.

### HasMd5sum

`func (o *SoftwarerepositoryCachedImage) HasMd5sum() bool`

HasMd5sum returns a boolean if a field has been set.

### GetOriginalSha512sum

`func (o *SoftwarerepositoryCachedImage) GetOriginalSha512sum() string`

GetOriginalSha512sum returns the OriginalSha512sum field if non-nil, zero value otherwise.

### GetOriginalSha512sumOk

`func (o *SoftwarerepositoryCachedImage) GetOriginalSha512sumOk() (*string, bool)`

GetOriginalSha512sumOk returns a tuple with the OriginalSha512sum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalSha512sum

`func (o *SoftwarerepositoryCachedImage) SetOriginalSha512sum(v string)`

SetOriginalSha512sum sets OriginalSha512sum field to given value.

### HasOriginalSha512sum

`func (o *SoftwarerepositoryCachedImage) HasOriginalSha512sum() bool`

HasOriginalSha512sum returns a boolean if a field has been set.

### GetPath

`func (o *SoftwarerepositoryCachedImage) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *SoftwarerepositoryCachedImage) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *SoftwarerepositoryCachedImage) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *SoftwarerepositoryCachedImage) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetRegisteredWorkflows

`func (o *SoftwarerepositoryCachedImage) GetRegisteredWorkflows() []string`

GetRegisteredWorkflows returns the RegisteredWorkflows field if non-nil, zero value otherwise.

### GetRegisteredWorkflowsOk

`func (o *SoftwarerepositoryCachedImage) GetRegisteredWorkflowsOk() (*[]string, bool)`

GetRegisteredWorkflowsOk returns a tuple with the RegisteredWorkflows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredWorkflows

`func (o *SoftwarerepositoryCachedImage) SetRegisteredWorkflows(v []string)`

SetRegisteredWorkflows sets RegisteredWorkflows field to given value.

### HasRegisteredWorkflows

`func (o *SoftwarerepositoryCachedImage) HasRegisteredWorkflows() bool`

HasRegisteredWorkflows returns a boolean if a field has been set.

### SetRegisteredWorkflowsNil

`func (o *SoftwarerepositoryCachedImage) SetRegisteredWorkflowsNil(b bool)`

 SetRegisteredWorkflowsNil sets the value for RegisteredWorkflows to be an explicit nil

### UnsetRegisteredWorkflows
`func (o *SoftwarerepositoryCachedImage) UnsetRegisteredWorkflows()`

UnsetRegisteredWorkflows ensures that no value is present for RegisteredWorkflows, not even an explicit nil
### GetUsedCount

`func (o *SoftwarerepositoryCachedImage) GetUsedCount() int64`

GetUsedCount returns the UsedCount field if non-nil, zero value otherwise.

### GetUsedCountOk

`func (o *SoftwarerepositoryCachedImage) GetUsedCountOk() (*int64, bool)`

GetUsedCountOk returns a tuple with the UsedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedCount

`func (o *SoftwarerepositoryCachedImage) SetUsedCount(v int64)`

SetUsedCount sets UsedCount field to given value.

### HasUsedCount

`func (o *SoftwarerepositoryCachedImage) HasUsedCount() bool`

HasUsedCount returns a boolean if a field has been set.

### GetFile

`func (o *SoftwarerepositoryCachedImage) GetFile() SoftwarerepositoryFileRelationship`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *SoftwarerepositoryCachedImage) GetFileOk() (*SoftwarerepositoryFileRelationship, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *SoftwarerepositoryCachedImage) SetFile(v SoftwarerepositoryFileRelationship)`

SetFile sets File field to given value.

### HasFile

`func (o *SoftwarerepositoryCachedImage) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetNetworkElement

`func (o *SoftwarerepositoryCachedImage) GetNetworkElement() NetworkElementRelationship`

GetNetworkElement returns the NetworkElement field if non-nil, zero value otherwise.

### GetNetworkElementOk

`func (o *SoftwarerepositoryCachedImage) GetNetworkElementOk() (*NetworkElementRelationship, bool)`

GetNetworkElementOk returns a tuple with the NetworkElement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkElement

`func (o *SoftwarerepositoryCachedImage) SetNetworkElement(v NetworkElementRelationship)`

SetNetworkElement sets NetworkElement field to given value.

### HasNetworkElement

`func (o *SoftwarerepositoryCachedImage) HasNetworkElement() bool`

HasNetworkElement returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



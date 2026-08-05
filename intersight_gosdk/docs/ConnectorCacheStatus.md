# ConnectorCacheStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "firmware.UpgradeStatus"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "firmware.UpgradeStatus"]
**CacheError** | Pointer to **string** | Any error encountered in caching. Example, Space unavailability due to too many active workflows running. | [optional] 
**CacheMessage** | Pointer to **string** | Message to notify caching operation status. | [optional] 
**CacheState** | Pointer to **string** | The current cache status of the file. * &#x60;ReadyForImport&#x60; - The image is ready to be imported into the repository. * &#x60;Importing&#x60; - The image is being imported into the repository. * &#x60;Imported&#x60; - The image has been extracted and imported into the repository. * &#x60;ComputingMetadata&#x60; - Indicates that the image has been imported but its metadata computation has not been done. * &#x60;PendingExtraction&#x60; - Indicates that the image has been imported but not extracted in the repository. * &#x60;Extracting&#x60; - Indicates that the image is being extracted into the repository. * &#x60;Extracted&#x60; - Indicates that the image has been extracted into the repository. * &#x60;Failed&#x60; - The image import from an external source to the repository has failed. * &#x60;MetaOnly&#x60; - The image is present in an external repository. * &#x60;ReadyForCache&#x60; - The image is ready to be cached into the Intersight Appliance. * &#x60;Caching&#x60; - Indicates that the image is being cached into the Intersight Appliance or endpoint cache. * &#x60;Cached&#x60; - Indicates that the image has been cached into the Intersight Appliance or endpoint cache. * &#x60;CachingFailed&#x60; - Indicates that the image caching into the Intersight Appliance failed or endpoint cache. * &#x60;Corrupted&#x60; - Indicates that the image in the local repository (or endpoint cache) has been corrupted after it was cached. * &#x60;Evicted&#x60; - Indicates that the image has been evicted from the Intersight Appliance (or endpoint cache) to reclaim storage space. * &#x60;Invalid&#x60; - Indicates that the corresponding distributable MO has been removed from the backend. This can be due to unpublishing of an image. | [optional] [default to "ReadyForImport"]
**DownloadError** | Pointer to **string** | Any error encountered. Set to empty when download is in progress or completed. | [optional] 
**DownloadProgress** | Pointer to **int64** | The download progress of the file represented as a percentage between 0 and 100. If progress reporting is not possible, a value of -1 is sent. | [optional] 
**DownloadRetries** | Pointer to **int64** | The number of retries the plugin attempted before succeeding or failing the download. | [optional] 

## Methods

### NewConnectorCacheStatus

`func NewConnectorCacheStatus(classId string, objectType string, ) *ConnectorCacheStatus`

NewConnectorCacheStatus instantiates a new ConnectorCacheStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorCacheStatusWithDefaults

`func NewConnectorCacheStatusWithDefaults() *ConnectorCacheStatus`

NewConnectorCacheStatusWithDefaults instantiates a new ConnectorCacheStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ConnectorCacheStatus) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ConnectorCacheStatus) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ConnectorCacheStatus) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ConnectorCacheStatus) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ConnectorCacheStatus) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ConnectorCacheStatus) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCacheError

`func (o *ConnectorCacheStatus) GetCacheError() string`

GetCacheError returns the CacheError field if non-nil, zero value otherwise.

### GetCacheErrorOk

`func (o *ConnectorCacheStatus) GetCacheErrorOk() (*string, bool)`

GetCacheErrorOk returns a tuple with the CacheError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheError

`func (o *ConnectorCacheStatus) SetCacheError(v string)`

SetCacheError sets CacheError field to given value.

### HasCacheError

`func (o *ConnectorCacheStatus) HasCacheError() bool`

HasCacheError returns a boolean if a field has been set.

### GetCacheMessage

`func (o *ConnectorCacheStatus) GetCacheMessage() string`

GetCacheMessage returns the CacheMessage field if non-nil, zero value otherwise.

### GetCacheMessageOk

`func (o *ConnectorCacheStatus) GetCacheMessageOk() (*string, bool)`

GetCacheMessageOk returns a tuple with the CacheMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheMessage

`func (o *ConnectorCacheStatus) SetCacheMessage(v string)`

SetCacheMessage sets CacheMessage field to given value.

### HasCacheMessage

`func (o *ConnectorCacheStatus) HasCacheMessage() bool`

HasCacheMessage returns a boolean if a field has been set.

### GetCacheState

`func (o *ConnectorCacheStatus) GetCacheState() string`

GetCacheState returns the CacheState field if non-nil, zero value otherwise.

### GetCacheStateOk

`func (o *ConnectorCacheStatus) GetCacheStateOk() (*string, bool)`

GetCacheStateOk returns a tuple with the CacheState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheState

`func (o *ConnectorCacheStatus) SetCacheState(v string)`

SetCacheState sets CacheState field to given value.

### HasCacheState

`func (o *ConnectorCacheStatus) HasCacheState() bool`

HasCacheState returns a boolean if a field has been set.

### GetDownloadError

`func (o *ConnectorCacheStatus) GetDownloadError() string`

GetDownloadError returns the DownloadError field if non-nil, zero value otherwise.

### GetDownloadErrorOk

`func (o *ConnectorCacheStatus) GetDownloadErrorOk() (*string, bool)`

GetDownloadErrorOk returns a tuple with the DownloadError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadError

`func (o *ConnectorCacheStatus) SetDownloadError(v string)`

SetDownloadError sets DownloadError field to given value.

### HasDownloadError

`func (o *ConnectorCacheStatus) HasDownloadError() bool`

HasDownloadError returns a boolean if a field has been set.

### GetDownloadProgress

`func (o *ConnectorCacheStatus) GetDownloadProgress() int64`

GetDownloadProgress returns the DownloadProgress field if non-nil, zero value otherwise.

### GetDownloadProgressOk

`func (o *ConnectorCacheStatus) GetDownloadProgressOk() (*int64, bool)`

GetDownloadProgressOk returns a tuple with the DownloadProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadProgress

`func (o *ConnectorCacheStatus) SetDownloadProgress(v int64)`

SetDownloadProgress sets DownloadProgress field to given value.

### HasDownloadProgress

`func (o *ConnectorCacheStatus) HasDownloadProgress() bool`

HasDownloadProgress returns a boolean if a field has been set.

### GetDownloadRetries

`func (o *ConnectorCacheStatus) GetDownloadRetries() int64`

GetDownloadRetries returns the DownloadRetries field if non-nil, zero value otherwise.

### GetDownloadRetriesOk

`func (o *ConnectorCacheStatus) GetDownloadRetriesOk() (*int64, bool)`

GetDownloadRetriesOk returns a tuple with the DownloadRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadRetries

`func (o *ConnectorCacheStatus) SetDownloadRetries(v int64)`

SetDownloadRetries sets DownloadRetries field to given value.

### HasDownloadRetries

`func (o *ConnectorCacheStatus) HasDownloadRetries() bool`

HasDownloadRetries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



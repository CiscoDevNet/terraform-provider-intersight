# TaskFileDownloadInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "task.FileDownloadInfo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "task.FileDownloadInfo"]
**Contents** | Pointer to **string** | When the type of the download is inline, it will read the file from the contents here. | [optional] 
**Path** | Pointer to **string** | The path of the download from the specified source location for type S3, then this is the object key. | [optional] 
**Source** | Pointer to **string** | The source of the download location and if type is S3, then this is the bucket name. In case of MoRef download type  the source will have the Moid of the workflow definition. | [optional] 
**Type** | Pointer to **string** | The type of download location is captured in type. * &#x60;S3&#x60; - Download workflow from S3. * &#x60;Local&#x60; - Download workflow from local path. * &#x60;Inline&#x60; - Workflow is provided inline. * &#x60;MoRef&#x60; - Start an existing workflow registered with given Moid. | [optional] [default to "S3"]

## Methods

### NewTaskFileDownloadInfo

`func NewTaskFileDownloadInfo(classId string, objectType string, ) *TaskFileDownloadInfo`

NewTaskFileDownloadInfo instantiates a new TaskFileDownloadInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskFileDownloadInfoWithDefaults

`func NewTaskFileDownloadInfoWithDefaults() *TaskFileDownloadInfo`

NewTaskFileDownloadInfoWithDefaults instantiates a new TaskFileDownloadInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *TaskFileDownloadInfo) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *TaskFileDownloadInfo) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *TaskFileDownloadInfo) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *TaskFileDownloadInfo) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *TaskFileDownloadInfo) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *TaskFileDownloadInfo) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetContents

`func (o *TaskFileDownloadInfo) GetContents() string`

GetContents returns the Contents field if non-nil, zero value otherwise.

### GetContentsOk

`func (o *TaskFileDownloadInfo) GetContentsOk() (*string, bool)`

GetContentsOk returns a tuple with the Contents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContents

`func (o *TaskFileDownloadInfo) SetContents(v string)`

SetContents sets Contents field to given value.

### HasContents

`func (o *TaskFileDownloadInfo) HasContents() bool`

HasContents returns a boolean if a field has been set.

### GetPath

`func (o *TaskFileDownloadInfo) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *TaskFileDownloadInfo) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *TaskFileDownloadInfo) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *TaskFileDownloadInfo) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetSource

`func (o *TaskFileDownloadInfo) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *TaskFileDownloadInfo) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *TaskFileDownloadInfo) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *TaskFileDownloadInfo) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetType

`func (o *TaskFileDownloadInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TaskFileDownloadInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TaskFileDownloadInfo) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TaskFileDownloadInfo) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



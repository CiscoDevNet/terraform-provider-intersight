# WorkflowFileDownloadOp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.FileDownloadOp"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.FileDownloadOp"]
**DestinationPath** | Pointer to **string** | Path on the Intersight connected device to which file needs to be downloaded. | [optional] 
**SourceBucket** | Pointer to **string** | Source bucket name hosting the file. | [optional] 
**SourceFile** | Pointer to **string** | Name of the file to be downloaded from bucket to endpoint devices. | [optional] 

## Methods

### NewWorkflowFileDownloadOp

`func NewWorkflowFileDownloadOp(classId string, objectType string, ) *WorkflowFileDownloadOp`

NewWorkflowFileDownloadOp instantiates a new WorkflowFileDownloadOp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowFileDownloadOpWithDefaults

`func NewWorkflowFileDownloadOpWithDefaults() *WorkflowFileDownloadOp`

NewWorkflowFileDownloadOpWithDefaults instantiates a new WorkflowFileDownloadOp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowFileDownloadOp) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowFileDownloadOp) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowFileDownloadOp) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowFileDownloadOp) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowFileDownloadOp) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowFileDownloadOp) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDestinationPath

`func (o *WorkflowFileDownloadOp) GetDestinationPath() string`

GetDestinationPath returns the DestinationPath field if non-nil, zero value otherwise.

### GetDestinationPathOk

`func (o *WorkflowFileDownloadOp) GetDestinationPathOk() (*string, bool)`

GetDestinationPathOk returns a tuple with the DestinationPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationPath

`func (o *WorkflowFileDownloadOp) SetDestinationPath(v string)`

SetDestinationPath sets DestinationPath field to given value.

### HasDestinationPath

`func (o *WorkflowFileDownloadOp) HasDestinationPath() bool`

HasDestinationPath returns a boolean if a field has been set.

### GetSourceBucket

`func (o *WorkflowFileDownloadOp) GetSourceBucket() string`

GetSourceBucket returns the SourceBucket field if non-nil, zero value otherwise.

### GetSourceBucketOk

`func (o *WorkflowFileDownloadOp) GetSourceBucketOk() (*string, bool)`

GetSourceBucketOk returns a tuple with the SourceBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceBucket

`func (o *WorkflowFileDownloadOp) SetSourceBucket(v string)`

SetSourceBucket sets SourceBucket field to given value.

### HasSourceBucket

`func (o *WorkflowFileDownloadOp) HasSourceBucket() bool`

HasSourceBucket returns a boolean if a field has been set.

### GetSourceFile

`func (o *WorkflowFileDownloadOp) GetSourceFile() string`

GetSourceFile returns the SourceFile field if non-nil, zero value otherwise.

### GetSourceFileOk

`func (o *WorkflowFileDownloadOp) GetSourceFileOk() (*string, bool)`

GetSourceFileOk returns a tuple with the SourceFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceFile

`func (o *WorkflowFileDownloadOp) SetSourceFile(v string)`

SetSourceFile sets SourceFile field to given value.

### HasSourceFile

`func (o *WorkflowFileDownloadOp) HasSourceFile() bool`

HasSourceFile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



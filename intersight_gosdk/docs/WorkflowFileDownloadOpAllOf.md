# WorkflowFileDownloadOpAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.FileDownloadOp"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.FileDownloadOp"]
**DestinationPath** | Pointer to **string** | Path on the Intersight connected device to which file needs to be downloaded. | [optional] 
**SourceBucket** | Pointer to **string** | Source bucket name hosting the file. | [optional] 
**SourceFile** | Pointer to **string** | Name of the file to be downloaded from bucket to endpoint devices. | [optional] 

## Methods

### NewWorkflowFileDownloadOpAllOf

`func NewWorkflowFileDownloadOpAllOf(classId string, objectType string, ) *WorkflowFileDownloadOpAllOf`

NewWorkflowFileDownloadOpAllOf instantiates a new WorkflowFileDownloadOpAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowFileDownloadOpAllOfWithDefaults

`func NewWorkflowFileDownloadOpAllOfWithDefaults() *WorkflowFileDownloadOpAllOf`

NewWorkflowFileDownloadOpAllOfWithDefaults instantiates a new WorkflowFileDownloadOpAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowFileDownloadOpAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowFileDownloadOpAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowFileDownloadOpAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowFileDownloadOpAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowFileDownloadOpAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowFileDownloadOpAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDestinationPath

`func (o *WorkflowFileDownloadOpAllOf) GetDestinationPath() string`

GetDestinationPath returns the DestinationPath field if non-nil, zero value otherwise.

### GetDestinationPathOk

`func (o *WorkflowFileDownloadOpAllOf) GetDestinationPathOk() (*string, bool)`

GetDestinationPathOk returns a tuple with the DestinationPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationPath

`func (o *WorkflowFileDownloadOpAllOf) SetDestinationPath(v string)`

SetDestinationPath sets DestinationPath field to given value.

### HasDestinationPath

`func (o *WorkflowFileDownloadOpAllOf) HasDestinationPath() bool`

HasDestinationPath returns a boolean if a field has been set.

### GetSourceBucket

`func (o *WorkflowFileDownloadOpAllOf) GetSourceBucket() string`

GetSourceBucket returns the SourceBucket field if non-nil, zero value otherwise.

### GetSourceBucketOk

`func (o *WorkflowFileDownloadOpAllOf) GetSourceBucketOk() (*string, bool)`

GetSourceBucketOk returns a tuple with the SourceBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceBucket

`func (o *WorkflowFileDownloadOpAllOf) SetSourceBucket(v string)`

SetSourceBucket sets SourceBucket field to given value.

### HasSourceBucket

`func (o *WorkflowFileDownloadOpAllOf) HasSourceBucket() bool`

HasSourceBucket returns a boolean if a field has been set.

### GetSourceFile

`func (o *WorkflowFileDownloadOpAllOf) GetSourceFile() string`

GetSourceFile returns the SourceFile field if non-nil, zero value otherwise.

### GetSourceFileOk

`func (o *WorkflowFileDownloadOpAllOf) GetSourceFileOk() (*string, bool)`

GetSourceFileOk returns a tuple with the SourceFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceFile

`func (o *WorkflowFileDownloadOpAllOf) SetSourceFile(v string)`

SetSourceFile sets SourceFile field to given value.

### HasSourceFile

`func (o *WorkflowFileDownloadOpAllOf) HasSourceFile() bool`

HasSourceFile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



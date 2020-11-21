# WorkflowFileTransfer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.FileTransfer"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.FileTransfer"]
**DestinationFilePath** | Pointer to **string** | Destination file path on the target server. | [optional] 
**FileMode** | Pointer to **int64** | File permission to set on the transferred file. | [optional] 
**SourceFilePath** | Pointer to **string** | Source file path on the Intersight connected device. | [optional] 

## Methods

### NewWorkflowFileTransfer

`func NewWorkflowFileTransfer(classId string, objectType string, ) *WorkflowFileTransfer`

NewWorkflowFileTransfer instantiates a new WorkflowFileTransfer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowFileTransferWithDefaults

`func NewWorkflowFileTransferWithDefaults() *WorkflowFileTransfer`

NewWorkflowFileTransferWithDefaults instantiates a new WorkflowFileTransfer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowFileTransfer) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowFileTransfer) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowFileTransfer) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowFileTransfer) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowFileTransfer) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowFileTransfer) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDestinationFilePath

`func (o *WorkflowFileTransfer) GetDestinationFilePath() string`

GetDestinationFilePath returns the DestinationFilePath field if non-nil, zero value otherwise.

### GetDestinationFilePathOk

`func (o *WorkflowFileTransfer) GetDestinationFilePathOk() (*string, bool)`

GetDestinationFilePathOk returns a tuple with the DestinationFilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationFilePath

`func (o *WorkflowFileTransfer) SetDestinationFilePath(v string)`

SetDestinationFilePath sets DestinationFilePath field to given value.

### HasDestinationFilePath

`func (o *WorkflowFileTransfer) HasDestinationFilePath() bool`

HasDestinationFilePath returns a boolean if a field has been set.

### GetFileMode

`func (o *WorkflowFileTransfer) GetFileMode() int64`

GetFileMode returns the FileMode field if non-nil, zero value otherwise.

### GetFileModeOk

`func (o *WorkflowFileTransfer) GetFileModeOk() (*int64, bool)`

GetFileModeOk returns a tuple with the FileMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileMode

`func (o *WorkflowFileTransfer) SetFileMode(v int64)`

SetFileMode sets FileMode field to given value.

### HasFileMode

`func (o *WorkflowFileTransfer) HasFileMode() bool`

HasFileMode returns a boolean if a field has been set.

### GetSourceFilePath

`func (o *WorkflowFileTransfer) GetSourceFilePath() string`

GetSourceFilePath returns the SourceFilePath field if non-nil, zero value otherwise.

### GetSourceFilePathOk

`func (o *WorkflowFileTransfer) GetSourceFilePathOk() (*string, bool)`

GetSourceFilePathOk returns a tuple with the SourceFilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceFilePath

`func (o *WorkflowFileTransfer) SetSourceFilePath(v string)`

SetSourceFilePath sets SourceFilePath field to given value.

### HasSourceFilePath

`func (o *WorkflowFileTransfer) HasSourceFilePath() bool`

HasSourceFilePath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



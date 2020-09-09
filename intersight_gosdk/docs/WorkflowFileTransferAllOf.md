# WorkflowFileTransferAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationFilePath** | Pointer to **string** | Destination file path on the target server. | [optional] 
**FileMode** | Pointer to **int64** | File permission to set on the transferred file. | [optional] 
**SourceFilePath** | Pointer to **string** | Source file path on the Intersight connected device. | [optional] 

## Methods

### NewWorkflowFileTransferAllOf

`func NewWorkflowFileTransferAllOf() *WorkflowFileTransferAllOf`

NewWorkflowFileTransferAllOf instantiates a new WorkflowFileTransferAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowFileTransferAllOfWithDefaults

`func NewWorkflowFileTransferAllOfWithDefaults() *WorkflowFileTransferAllOf`

NewWorkflowFileTransferAllOfWithDefaults instantiates a new WorkflowFileTransferAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationFilePath

`func (o *WorkflowFileTransferAllOf) GetDestinationFilePath() string`

GetDestinationFilePath returns the DestinationFilePath field if non-nil, zero value otherwise.

### GetDestinationFilePathOk

`func (o *WorkflowFileTransferAllOf) GetDestinationFilePathOk() (*string, bool)`

GetDestinationFilePathOk returns a tuple with the DestinationFilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationFilePath

`func (o *WorkflowFileTransferAllOf) SetDestinationFilePath(v string)`

SetDestinationFilePath sets DestinationFilePath field to given value.

### HasDestinationFilePath

`func (o *WorkflowFileTransferAllOf) HasDestinationFilePath() bool`

HasDestinationFilePath returns a boolean if a field has been set.

### GetFileMode

`func (o *WorkflowFileTransferAllOf) GetFileMode() int64`

GetFileMode returns the FileMode field if non-nil, zero value otherwise.

### GetFileModeOk

`func (o *WorkflowFileTransferAllOf) GetFileModeOk() (*int64, bool)`

GetFileModeOk returns a tuple with the FileMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileMode

`func (o *WorkflowFileTransferAllOf) SetFileMode(v int64)`

SetFileMode sets FileMode field to given value.

### HasFileMode

`func (o *WorkflowFileTransferAllOf) HasFileMode() bool`

HasFileMode returns a boolean if a field has been set.

### GetSourceFilePath

`func (o *WorkflowFileTransferAllOf) GetSourceFilePath() string`

GetSourceFilePath returns the SourceFilePath field if non-nil, zero value otherwise.

### GetSourceFilePathOk

`func (o *WorkflowFileTransferAllOf) GetSourceFilePathOk() (*string, bool)`

GetSourceFilePathOk returns a tuple with the SourceFilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceFilePath

`func (o *WorkflowFileTransferAllOf) SetSourceFilePath(v string)`

SetSourceFilePath sets SourceFilePath field to given value.

### HasSourceFilePath

`func (o *WorkflowFileTransferAllOf) HasSourceFilePath() bool`

HasSourceFilePath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



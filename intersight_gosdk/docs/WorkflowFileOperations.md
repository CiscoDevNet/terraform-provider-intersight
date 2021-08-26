# WorkflowFileOperations

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.FileOperations"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.FileOperations"]
**FileDownload** | Pointer to [**NullableWorkflowFileDownloadOp**](WorkflowFileDownloadOp.md) |  | [optional] 
**FileTemplate** | Pointer to [**NullableWorkflowFileTemplateOp**](WorkflowFileTemplateOp.md) |  | [optional] 
**OperationType** | Pointer to **string** | File operation type to be executed on the connected endpoint. * &#x60;FileDownload&#x60; - The API is executed in a remote device connected to the Intersightthrough its device connector. This operation is to download the filefrom specified storage bucket to the specific path on the device. * &#x60;FileTemplatize&#x60; - Populates data driven template file with input values to generate textual output.Inputs - the path of the template file on the device and json values to populate.An error will be returned if the file does not exists or if there is an error whileexecuting the template. | [optional] [default to "FileDownload"]

## Methods

### NewWorkflowFileOperations

`func NewWorkflowFileOperations(classId string, objectType string, ) *WorkflowFileOperations`

NewWorkflowFileOperations instantiates a new WorkflowFileOperations object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowFileOperationsWithDefaults

`func NewWorkflowFileOperationsWithDefaults() *WorkflowFileOperations`

NewWorkflowFileOperationsWithDefaults instantiates a new WorkflowFileOperations object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowFileOperations) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowFileOperations) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowFileOperations) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowFileOperations) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowFileOperations) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowFileOperations) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetFileDownload

`func (o *WorkflowFileOperations) GetFileDownload() WorkflowFileDownloadOp`

GetFileDownload returns the FileDownload field if non-nil, zero value otherwise.

### GetFileDownloadOk

`func (o *WorkflowFileOperations) GetFileDownloadOk() (*WorkflowFileDownloadOp, bool)`

GetFileDownloadOk returns a tuple with the FileDownload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileDownload

`func (o *WorkflowFileOperations) SetFileDownload(v WorkflowFileDownloadOp)`

SetFileDownload sets FileDownload field to given value.

### HasFileDownload

`func (o *WorkflowFileOperations) HasFileDownload() bool`

HasFileDownload returns a boolean if a field has been set.

### SetFileDownloadNil

`func (o *WorkflowFileOperations) SetFileDownloadNil(b bool)`

 SetFileDownloadNil sets the value for FileDownload to be an explicit nil

### UnsetFileDownload
`func (o *WorkflowFileOperations) UnsetFileDownload()`

UnsetFileDownload ensures that no value is present for FileDownload, not even an explicit nil
### GetFileTemplate

`func (o *WorkflowFileOperations) GetFileTemplate() WorkflowFileTemplateOp`

GetFileTemplate returns the FileTemplate field if non-nil, zero value otherwise.

### GetFileTemplateOk

`func (o *WorkflowFileOperations) GetFileTemplateOk() (*WorkflowFileTemplateOp, bool)`

GetFileTemplateOk returns a tuple with the FileTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileTemplate

`func (o *WorkflowFileOperations) SetFileTemplate(v WorkflowFileTemplateOp)`

SetFileTemplate sets FileTemplate field to given value.

### HasFileTemplate

`func (o *WorkflowFileOperations) HasFileTemplate() bool`

HasFileTemplate returns a boolean if a field has been set.

### SetFileTemplateNil

`func (o *WorkflowFileOperations) SetFileTemplateNil(b bool)`

 SetFileTemplateNil sets the value for FileTemplate to be an explicit nil

### UnsetFileTemplate
`func (o *WorkflowFileOperations) UnsetFileTemplate()`

UnsetFileTemplate ensures that no value is present for FileTemplate, not even an explicit nil
### GetOperationType

`func (o *WorkflowFileOperations) GetOperationType() string`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *WorkflowFileOperations) GetOperationTypeOk() (*string, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *WorkflowFileOperations) SetOperationType(v string)`

SetOperationType sets OperationType field to given value.

### HasOperationType

`func (o *WorkflowFileOperations) HasOperationType() bool`

HasOperationType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



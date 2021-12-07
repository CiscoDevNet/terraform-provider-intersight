# WorkflowSshSessionAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.SshSession"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.SshSession"]
**CaptureCompleteResponse** | Pointer to **string** | Flag to allow capturing entire command response as batch API output. | [optional] 
**ExpectedExitCodes** | Pointer to **interface{}** | Optional array of integer values to specify the expected exit codes of a SSH command execution. SSH command execution is marked success upon receiving any of the expected exit code from command execution. If not set, success exit code of 0 is expected from command execution. | [optional] 
**FileTransferToRemote** | Pointer to [**WorkflowFileTransfer**](WorkflowFileTransfer.md) |  | [optional] 
**MessageType** | Pointer to **string** | The type of SSH message to be sent to the remote server. * &#x60;ExecuteCommand&#x60; - Execute a SSH command on the remote server. * &#x60;NewSession&#x60; - Open a new SSH connection to the remote server. * &#x60;FileTransfer&#x60; - Transfer a file from Intersight connected device to the remote server. * &#x60;CloseSession&#x60; - Close the SSH connection to the remote server. | [optional] [default to "ExecuteCommand"]
**SshCommand** | Pointer to **interface{}** | SSH command to execute on the remote server. | [optional] 
**SshConfiguration** | Pointer to [**WorkflowSshConfig**](WorkflowSshConfig.md) |  | [optional] 
**SshOpTimeout** | Pointer to **string** | SSH operation timeout value in seconds. The provided string value should be able to convert to respective integer value. | [optional] 

## Methods

### NewWorkflowSshSessionAllOf

`func NewWorkflowSshSessionAllOf(classId string, objectType string, ) *WorkflowSshSessionAllOf`

NewWorkflowSshSessionAllOf instantiates a new WorkflowSshSessionAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowSshSessionAllOfWithDefaults

`func NewWorkflowSshSessionAllOfWithDefaults() *WorkflowSshSessionAllOf`

NewWorkflowSshSessionAllOfWithDefaults instantiates a new WorkflowSshSessionAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowSshSessionAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowSshSessionAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowSshSessionAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowSshSessionAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowSshSessionAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowSshSessionAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCaptureCompleteResponse

`func (o *WorkflowSshSessionAllOf) GetCaptureCompleteResponse() string`

GetCaptureCompleteResponse returns the CaptureCompleteResponse field if non-nil, zero value otherwise.

### GetCaptureCompleteResponseOk

`func (o *WorkflowSshSessionAllOf) GetCaptureCompleteResponseOk() (*string, bool)`

GetCaptureCompleteResponseOk returns a tuple with the CaptureCompleteResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptureCompleteResponse

`func (o *WorkflowSshSessionAllOf) SetCaptureCompleteResponse(v string)`

SetCaptureCompleteResponse sets CaptureCompleteResponse field to given value.

### HasCaptureCompleteResponse

`func (o *WorkflowSshSessionAllOf) HasCaptureCompleteResponse() bool`

HasCaptureCompleteResponse returns a boolean if a field has been set.

### GetExpectedExitCodes

`func (o *WorkflowSshSessionAllOf) GetExpectedExitCodes() interface{}`

GetExpectedExitCodes returns the ExpectedExitCodes field if non-nil, zero value otherwise.

### GetExpectedExitCodesOk

`func (o *WorkflowSshSessionAllOf) GetExpectedExitCodesOk() (*interface{}, bool)`

GetExpectedExitCodesOk returns a tuple with the ExpectedExitCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedExitCodes

`func (o *WorkflowSshSessionAllOf) SetExpectedExitCodes(v interface{})`

SetExpectedExitCodes sets ExpectedExitCodes field to given value.

### HasExpectedExitCodes

`func (o *WorkflowSshSessionAllOf) HasExpectedExitCodes() bool`

HasExpectedExitCodes returns a boolean if a field has been set.

### SetExpectedExitCodesNil

`func (o *WorkflowSshSessionAllOf) SetExpectedExitCodesNil(b bool)`

 SetExpectedExitCodesNil sets the value for ExpectedExitCodes to be an explicit nil

### UnsetExpectedExitCodes
`func (o *WorkflowSshSessionAllOf) UnsetExpectedExitCodes()`

UnsetExpectedExitCodes ensures that no value is present for ExpectedExitCodes, not even an explicit nil
### GetFileTransferToRemote

`func (o *WorkflowSshSessionAllOf) GetFileTransferToRemote() WorkflowFileTransfer`

GetFileTransferToRemote returns the FileTransferToRemote field if non-nil, zero value otherwise.

### GetFileTransferToRemoteOk

`func (o *WorkflowSshSessionAllOf) GetFileTransferToRemoteOk() (*WorkflowFileTransfer, bool)`

GetFileTransferToRemoteOk returns a tuple with the FileTransferToRemote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileTransferToRemote

`func (o *WorkflowSshSessionAllOf) SetFileTransferToRemote(v WorkflowFileTransfer)`

SetFileTransferToRemote sets FileTransferToRemote field to given value.

### HasFileTransferToRemote

`func (o *WorkflowSshSessionAllOf) HasFileTransferToRemote() bool`

HasFileTransferToRemote returns a boolean if a field has been set.

### GetMessageType

`func (o *WorkflowSshSessionAllOf) GetMessageType() string`

GetMessageType returns the MessageType field if non-nil, zero value otherwise.

### GetMessageTypeOk

`func (o *WorkflowSshSessionAllOf) GetMessageTypeOk() (*string, bool)`

GetMessageTypeOk returns a tuple with the MessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageType

`func (o *WorkflowSshSessionAllOf) SetMessageType(v string)`

SetMessageType sets MessageType field to given value.

### HasMessageType

`func (o *WorkflowSshSessionAllOf) HasMessageType() bool`

HasMessageType returns a boolean if a field has been set.

### GetSshCommand

`func (o *WorkflowSshSessionAllOf) GetSshCommand() interface{}`

GetSshCommand returns the SshCommand field if non-nil, zero value otherwise.

### GetSshCommandOk

`func (o *WorkflowSshSessionAllOf) GetSshCommandOk() (*interface{}, bool)`

GetSshCommandOk returns a tuple with the SshCommand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshCommand

`func (o *WorkflowSshSessionAllOf) SetSshCommand(v interface{})`

SetSshCommand sets SshCommand field to given value.

### HasSshCommand

`func (o *WorkflowSshSessionAllOf) HasSshCommand() bool`

HasSshCommand returns a boolean if a field has been set.

### SetSshCommandNil

`func (o *WorkflowSshSessionAllOf) SetSshCommandNil(b bool)`

 SetSshCommandNil sets the value for SshCommand to be an explicit nil

### UnsetSshCommand
`func (o *WorkflowSshSessionAllOf) UnsetSshCommand()`

UnsetSshCommand ensures that no value is present for SshCommand, not even an explicit nil
### GetSshConfiguration

`func (o *WorkflowSshSessionAllOf) GetSshConfiguration() WorkflowSshConfig`

GetSshConfiguration returns the SshConfiguration field if non-nil, zero value otherwise.

### GetSshConfigurationOk

`func (o *WorkflowSshSessionAllOf) GetSshConfigurationOk() (*WorkflowSshConfig, bool)`

GetSshConfigurationOk returns a tuple with the SshConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshConfiguration

`func (o *WorkflowSshSessionAllOf) SetSshConfiguration(v WorkflowSshConfig)`

SetSshConfiguration sets SshConfiguration field to given value.

### HasSshConfiguration

`func (o *WorkflowSshSessionAllOf) HasSshConfiguration() bool`

HasSshConfiguration returns a boolean if a field has been set.

### GetSshOpTimeout

`func (o *WorkflowSshSessionAllOf) GetSshOpTimeout() string`

GetSshOpTimeout returns the SshOpTimeout field if non-nil, zero value otherwise.

### GetSshOpTimeoutOk

`func (o *WorkflowSshSessionAllOf) GetSshOpTimeoutOk() (*string, bool)`

GetSshOpTimeoutOk returns a tuple with the SshOpTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshOpTimeout

`func (o *WorkflowSshSessionAllOf) SetSshOpTimeout(v string)`

SetSshOpTimeout sets SshOpTimeout field to given value.

### HasSshOpTimeout

`func (o *WorkflowSshSessionAllOf) HasSshOpTimeout() bool`

HasSshOpTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



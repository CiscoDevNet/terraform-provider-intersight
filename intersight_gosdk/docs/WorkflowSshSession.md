# WorkflowSshSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileTransferToRemote** | Pointer to [**WorkflowFileTransfer**](workflow.FileTransfer.md) |  | [optional] 
**MessageType** | Pointer to **string** | The type of SSH message to send to the remote server. * &#x60;ExecuteCommand&#x60; - Execute a SSH command on the remote server. * &#x60;NewSession&#x60; - Open a new SSH connection to the remote server. * &#x60;FileTransfer&#x60; - Transfer a file from Intersight connected device to the remote server. * &#x60;CloseSession&#x60; - Close the SSH connection to the remote server. | [optional] [default to "ExecuteCommand"]
**SshCommand** | Pointer to [**WorkflowSshCmd**](workflow.SshCmd.md) |  | [optional] 
**SshConfiguration** | Pointer to [**WorkflowSshConfig**](workflow.SshConfig.md) |  | [optional] 

## Methods

### NewWorkflowSshSession

`func NewWorkflowSshSession() *WorkflowSshSession`

NewWorkflowSshSession instantiates a new WorkflowSshSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowSshSessionWithDefaults

`func NewWorkflowSshSessionWithDefaults() *WorkflowSshSession`

NewWorkflowSshSessionWithDefaults instantiates a new WorkflowSshSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileTransferToRemote

`func (o *WorkflowSshSession) GetFileTransferToRemote() WorkflowFileTransfer`

GetFileTransferToRemote returns the FileTransferToRemote field if non-nil, zero value otherwise.

### GetFileTransferToRemoteOk

`func (o *WorkflowSshSession) GetFileTransferToRemoteOk() (*WorkflowFileTransfer, bool)`

GetFileTransferToRemoteOk returns a tuple with the FileTransferToRemote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileTransferToRemote

`func (o *WorkflowSshSession) SetFileTransferToRemote(v WorkflowFileTransfer)`

SetFileTransferToRemote sets FileTransferToRemote field to given value.

### HasFileTransferToRemote

`func (o *WorkflowSshSession) HasFileTransferToRemote() bool`

HasFileTransferToRemote returns a boolean if a field has been set.

### GetMessageType

`func (o *WorkflowSshSession) GetMessageType() string`

GetMessageType returns the MessageType field if non-nil, zero value otherwise.

### GetMessageTypeOk

`func (o *WorkflowSshSession) GetMessageTypeOk() (*string, bool)`

GetMessageTypeOk returns a tuple with the MessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageType

`func (o *WorkflowSshSession) SetMessageType(v string)`

SetMessageType sets MessageType field to given value.

### HasMessageType

`func (o *WorkflowSshSession) HasMessageType() bool`

HasMessageType returns a boolean if a field has been set.

### GetSshCommand

`func (o *WorkflowSshSession) GetSshCommand() WorkflowSshCmd`

GetSshCommand returns the SshCommand field if non-nil, zero value otherwise.

### GetSshCommandOk

`func (o *WorkflowSshSession) GetSshCommandOk() (*WorkflowSshCmd, bool)`

GetSshCommandOk returns a tuple with the SshCommand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshCommand

`func (o *WorkflowSshSession) SetSshCommand(v WorkflowSshCmd)`

SetSshCommand sets SshCommand field to given value.

### HasSshCommand

`func (o *WorkflowSshSession) HasSshCommand() bool`

HasSshCommand returns a boolean if a field has been set.

### GetSshConfiguration

`func (o *WorkflowSshSession) GetSshConfiguration() WorkflowSshConfig`

GetSshConfiguration returns the SshConfiguration field if non-nil, zero value otherwise.

### GetSshConfigurationOk

`func (o *WorkflowSshSession) GetSshConfigurationOk() (*WorkflowSshConfig, bool)`

GetSshConfigurationOk returns a tuple with the SshConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshConfiguration

`func (o *WorkflowSshSession) SetSshConfiguration(v WorkflowSshConfig)`

SetSshConfiguration sets SshConfiguration field to given value.

### HasSshConfiguration

`func (o *WorkflowSshSession) HasSshConfiguration() bool`

HasSshConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# WorkflowSshSessionAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.SshSession"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.SshSession"]
**FileTransferToRemote** | Pointer to [**WorkflowFileTransfer**](WorkflowFileTransfer.md) |  | [optional] 
**MessageType** | Pointer to **string** | The type of SSH message to send to the remote server. * &#x60;ExecuteCommand&#x60; - Execute a SSH command on the remote server. * &#x60;NewSession&#x60; - Open a new SSH connection to the remote server. * &#x60;FileTransfer&#x60; - Transfer a file from Intersight connected device to the remote server. * &#x60;CloseSession&#x60; - Close the SSH connection to the remote server. | [optional] [default to "ExecuteCommand"]
**SshCommand** | Pointer to [**WorkflowSshCmd**](WorkflowSshCmd.md) |  | [optional] 
**SshConfiguration** | Pointer to [**WorkflowSshConfig**](WorkflowSshConfig.md) |  | [optional] 

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

`func (o *WorkflowSshSessionAllOf) GetSshCommand() WorkflowSshCmd`

GetSshCommand returns the SshCommand field if non-nil, zero value otherwise.

### GetSshCommandOk

`func (o *WorkflowSshSessionAllOf) GetSshCommandOk() (*WorkflowSshCmd, bool)`

GetSshCommandOk returns a tuple with the SshCommand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshCommand

`func (o *WorkflowSshSessionAllOf) SetSshCommand(v WorkflowSshCmd)`

SetSshCommand sets SshCommand field to given value.

### HasSshCommand

`func (o *WorkflowSshSessionAllOf) HasSshCommand() bool`

HasSshCommand returns a boolean if a field has been set.

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



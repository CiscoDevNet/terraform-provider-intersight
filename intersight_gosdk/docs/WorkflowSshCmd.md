# WorkflowSshCmd

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Command** | Pointer to **string** | SSH command to execute on the remote server. | [optional] 
**CommandType** | Pointer to **string** | SSH command type to execute on the remote server. * &#x60;NonInteractiveCmd&#x60; - Execute a non-interactive SSH command on the remote server. * &#x60;InteractiveCmd&#x60; - Execute an interactive SSH command on the remote server. | [optional] [default to "NonInteractiveCmd"]
**ExpectPrompts** | Pointer to [**[]ConnectorExpectPrompt**](connector.ExpectPrompt.md) |  | [optional] 
**ShellPrompt** | Pointer to **string** | Regex of the remote server&#39;s shell prompt. | [optional] 
**ShellPromptTimeout** | Pointer to **int64** | Expect timeout value in seconds for the shell prompt. | [optional] 

## Methods

### NewWorkflowSshCmd

`func NewWorkflowSshCmd() *WorkflowSshCmd`

NewWorkflowSshCmd instantiates a new WorkflowSshCmd object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowSshCmdWithDefaults

`func NewWorkflowSshCmdWithDefaults() *WorkflowSshCmd`

NewWorkflowSshCmdWithDefaults instantiates a new WorkflowSshCmd object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommand

`func (o *WorkflowSshCmd) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *WorkflowSshCmd) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *WorkflowSshCmd) SetCommand(v string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *WorkflowSshCmd) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### GetCommandType

`func (o *WorkflowSshCmd) GetCommandType() string`

GetCommandType returns the CommandType field if non-nil, zero value otherwise.

### GetCommandTypeOk

`func (o *WorkflowSshCmd) GetCommandTypeOk() (*string, bool)`

GetCommandTypeOk returns a tuple with the CommandType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandType

`func (o *WorkflowSshCmd) SetCommandType(v string)`

SetCommandType sets CommandType field to given value.

### HasCommandType

`func (o *WorkflowSshCmd) HasCommandType() bool`

HasCommandType returns a boolean if a field has been set.

### GetExpectPrompts

`func (o *WorkflowSshCmd) GetExpectPrompts() []ConnectorExpectPrompt`

GetExpectPrompts returns the ExpectPrompts field if non-nil, zero value otherwise.

### GetExpectPromptsOk

`func (o *WorkflowSshCmd) GetExpectPromptsOk() (*[]ConnectorExpectPrompt, bool)`

GetExpectPromptsOk returns a tuple with the ExpectPrompts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectPrompts

`func (o *WorkflowSshCmd) SetExpectPrompts(v []ConnectorExpectPrompt)`

SetExpectPrompts sets ExpectPrompts field to given value.

### HasExpectPrompts

`func (o *WorkflowSshCmd) HasExpectPrompts() bool`

HasExpectPrompts returns a boolean if a field has been set.

### GetShellPrompt

`func (o *WorkflowSshCmd) GetShellPrompt() string`

GetShellPrompt returns the ShellPrompt field if non-nil, zero value otherwise.

### GetShellPromptOk

`func (o *WorkflowSshCmd) GetShellPromptOk() (*string, bool)`

GetShellPromptOk returns a tuple with the ShellPrompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShellPrompt

`func (o *WorkflowSshCmd) SetShellPrompt(v string)`

SetShellPrompt sets ShellPrompt field to given value.

### HasShellPrompt

`func (o *WorkflowSshCmd) HasShellPrompt() bool`

HasShellPrompt returns a boolean if a field has been set.

### GetShellPromptTimeout

`func (o *WorkflowSshCmd) GetShellPromptTimeout() int64`

GetShellPromptTimeout returns the ShellPromptTimeout field if non-nil, zero value otherwise.

### GetShellPromptTimeoutOk

`func (o *WorkflowSshCmd) GetShellPromptTimeoutOk() (*int64, bool)`

GetShellPromptTimeoutOk returns a tuple with the ShellPromptTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShellPromptTimeout

`func (o *WorkflowSshCmd) SetShellPromptTimeout(v int64)`

SetShellPromptTimeout sets ShellPromptTimeout field to given value.

### HasShellPromptTimeout

`func (o *WorkflowSshCmd) HasShellPromptTimeout() bool`

HasShellPromptTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



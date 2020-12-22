# WorkflowSshCmdAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.SshCmd"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.SshCmd"]
**Command** | Pointer to **string** | SSH command to execute on the remote server. | [optional] 
**CommandType** | Pointer to **string** | SSH command type to execute on the remote server. * &#x60;NonInteractiveCmd&#x60; - Execute a non-interactive SSH command on the remote server. * &#x60;InteractiveCmd&#x60; - Execute an interactive SSH command on the remote server. | [optional] [default to "NonInteractiveCmd"]
**ExpectPrompts** | Pointer to [**[]ConnectorExpectPrompt**](ConnectorExpectPrompt.md) |  | [optional] 
**ShellPrompt** | Pointer to **string** | Regex of the remote server&#39;s shell prompt. | [optional] 
**ShellPromptTimeout** | Pointer to **int64** | Expect timeout value in seconds for the shell prompt. | [optional] 

## Methods

### NewWorkflowSshCmdAllOf

`func NewWorkflowSshCmdAllOf(classId string, objectType string, ) *WorkflowSshCmdAllOf`

NewWorkflowSshCmdAllOf instantiates a new WorkflowSshCmdAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowSshCmdAllOfWithDefaults

`func NewWorkflowSshCmdAllOfWithDefaults() *WorkflowSshCmdAllOf`

NewWorkflowSshCmdAllOfWithDefaults instantiates a new WorkflowSshCmdAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowSshCmdAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowSshCmdAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowSshCmdAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowSshCmdAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowSshCmdAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowSshCmdAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCommand

`func (o *WorkflowSshCmdAllOf) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *WorkflowSshCmdAllOf) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *WorkflowSshCmdAllOf) SetCommand(v string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *WorkflowSshCmdAllOf) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### GetCommandType

`func (o *WorkflowSshCmdAllOf) GetCommandType() string`

GetCommandType returns the CommandType field if non-nil, zero value otherwise.

### GetCommandTypeOk

`func (o *WorkflowSshCmdAllOf) GetCommandTypeOk() (*string, bool)`

GetCommandTypeOk returns a tuple with the CommandType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandType

`func (o *WorkflowSshCmdAllOf) SetCommandType(v string)`

SetCommandType sets CommandType field to given value.

### HasCommandType

`func (o *WorkflowSshCmdAllOf) HasCommandType() bool`

HasCommandType returns a boolean if a field has been set.

### GetExpectPrompts

`func (o *WorkflowSshCmdAllOf) GetExpectPrompts() []ConnectorExpectPrompt`

GetExpectPrompts returns the ExpectPrompts field if non-nil, zero value otherwise.

### GetExpectPromptsOk

`func (o *WorkflowSshCmdAllOf) GetExpectPromptsOk() (*[]ConnectorExpectPrompt, bool)`

GetExpectPromptsOk returns a tuple with the ExpectPrompts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectPrompts

`func (o *WorkflowSshCmdAllOf) SetExpectPrompts(v []ConnectorExpectPrompt)`

SetExpectPrompts sets ExpectPrompts field to given value.

### HasExpectPrompts

`func (o *WorkflowSshCmdAllOf) HasExpectPrompts() bool`

HasExpectPrompts returns a boolean if a field has been set.

### SetExpectPromptsNil

`func (o *WorkflowSshCmdAllOf) SetExpectPromptsNil(b bool)`

 SetExpectPromptsNil sets the value for ExpectPrompts to be an explicit nil

### UnsetExpectPrompts
`func (o *WorkflowSshCmdAllOf) UnsetExpectPrompts()`

UnsetExpectPrompts ensures that no value is present for ExpectPrompts, not even an explicit nil
### GetShellPrompt

`func (o *WorkflowSshCmdAllOf) GetShellPrompt() string`

GetShellPrompt returns the ShellPrompt field if non-nil, zero value otherwise.

### GetShellPromptOk

`func (o *WorkflowSshCmdAllOf) GetShellPromptOk() (*string, bool)`

GetShellPromptOk returns a tuple with the ShellPrompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShellPrompt

`func (o *WorkflowSshCmdAllOf) SetShellPrompt(v string)`

SetShellPrompt sets ShellPrompt field to given value.

### HasShellPrompt

`func (o *WorkflowSshCmdAllOf) HasShellPrompt() bool`

HasShellPrompt returns a boolean if a field has been set.

### GetShellPromptTimeout

`func (o *WorkflowSshCmdAllOf) GetShellPromptTimeout() int64`

GetShellPromptTimeout returns the ShellPromptTimeout field if non-nil, zero value otherwise.

### GetShellPromptTimeoutOk

`func (o *WorkflowSshCmdAllOf) GetShellPromptTimeoutOk() (*int64, bool)`

GetShellPromptTimeoutOk returns a tuple with the ShellPromptTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShellPromptTimeout

`func (o *WorkflowSshCmdAllOf) SetShellPromptTimeout(v int64)`

SetShellPromptTimeout sets ShellPromptTimeout field to given value.

### HasShellPromptTimeout

`func (o *WorkflowSshCmdAllOf) HasShellPromptTimeout() bool`

HasShellPromptTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



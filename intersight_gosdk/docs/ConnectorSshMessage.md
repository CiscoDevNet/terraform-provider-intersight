# ConnectorSshMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpectPrompts** | Pointer to [**[]ConnectorExpectPrompt**](connector.ExpectPrompt.md) |  | [optional] 
**MsgType** | Pointer to **int64** | The operation to execute on a new or existing session. | [optional] 
**SessionId** | Pointer to **string** | Unique id of session to route messages to. | [optional] 
**ShellPrompt** | Pointer to **string** | The regex of the secure shell prompt. | [optional] 
**Stream** | Pointer to **string** | Input to the SSH operation to be executed. e.g. file contents to write. | [optional] 
**Timeout** | Pointer to **int64** | The timeout for the ssh command to complete and exit after starting or receiving input. If timeout is not set a default of 10 minutes will be used. | [optional] 

## Methods

### NewConnectorSshMessage

`func NewConnectorSshMessage() *ConnectorSshMessage`

NewConnectorSshMessage instantiates a new ConnectorSshMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorSshMessageWithDefaults

`func NewConnectorSshMessageWithDefaults() *ConnectorSshMessage`

NewConnectorSshMessageWithDefaults instantiates a new ConnectorSshMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpectPrompts

`func (o *ConnectorSshMessage) GetExpectPrompts() []ConnectorExpectPrompt`

GetExpectPrompts returns the ExpectPrompts field if non-nil, zero value otherwise.

### GetExpectPromptsOk

`func (o *ConnectorSshMessage) GetExpectPromptsOk() (*[]ConnectorExpectPrompt, bool)`

GetExpectPromptsOk returns a tuple with the ExpectPrompts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectPrompts

`func (o *ConnectorSshMessage) SetExpectPrompts(v []ConnectorExpectPrompt)`

SetExpectPrompts sets ExpectPrompts field to given value.

### HasExpectPrompts

`func (o *ConnectorSshMessage) HasExpectPrompts() bool`

HasExpectPrompts returns a boolean if a field has been set.

### GetMsgType

`func (o *ConnectorSshMessage) GetMsgType() int64`

GetMsgType returns the MsgType field if non-nil, zero value otherwise.

### GetMsgTypeOk

`func (o *ConnectorSshMessage) GetMsgTypeOk() (*int64, bool)`

GetMsgTypeOk returns a tuple with the MsgType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgType

`func (o *ConnectorSshMessage) SetMsgType(v int64)`

SetMsgType sets MsgType field to given value.

### HasMsgType

`func (o *ConnectorSshMessage) HasMsgType() bool`

HasMsgType returns a boolean if a field has been set.

### GetSessionId

`func (o *ConnectorSshMessage) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *ConnectorSshMessage) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *ConnectorSshMessage) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *ConnectorSshMessage) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetShellPrompt

`func (o *ConnectorSshMessage) GetShellPrompt() string`

GetShellPrompt returns the ShellPrompt field if non-nil, zero value otherwise.

### GetShellPromptOk

`func (o *ConnectorSshMessage) GetShellPromptOk() (*string, bool)`

GetShellPromptOk returns a tuple with the ShellPrompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShellPrompt

`func (o *ConnectorSshMessage) SetShellPrompt(v string)`

SetShellPrompt sets ShellPrompt field to given value.

### HasShellPrompt

`func (o *ConnectorSshMessage) HasShellPrompt() bool`

HasShellPrompt returns a boolean if a field has been set.

### GetStream

`func (o *ConnectorSshMessage) GetStream() string`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *ConnectorSshMessage) GetStreamOk() (*string, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *ConnectorSshMessage) SetStream(v string)`

SetStream sets Stream field to given value.

### HasStream

`func (o *ConnectorSshMessage) HasStream() bool`

HasStream returns a boolean if a field has been set.

### GetTimeout

`func (o *ConnectorSshMessage) GetTimeout() int64`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *ConnectorSshMessage) GetTimeoutOk() (*int64, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *ConnectorSshMessage) SetTimeout(v int64)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *ConnectorSshMessage) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



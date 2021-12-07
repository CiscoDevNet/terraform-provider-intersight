# ConnectorSshMessageAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "connector.SshMessage"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "connector.SshMessage"]
**ExpectPrompts** | Pointer to [**[]ConnectorExpectPrompt**](ConnectorExpectPrompt.md) |  | [optional] 
**MsgType** | Pointer to **int64** | The operation to execute on a new or existing session. | [optional] 
**SessionId** | Pointer to **string** | Unique id of session to route messages to. | [optional] 
**ShellPrompt** | Pointer to **string** | The regex of the secure shell prompt. | [optional] 
**Stream** | Pointer to **string** | Input to the SSH operation to be executed. e.g. file contents to write. | [optional] 
**Timeout** | Pointer to **int64** | The timeout for the ssh command to complete and exit after starting or receiving input. If timeout is not set a default of 10 minutes will be used. | [optional] 

## Methods

### NewConnectorSshMessageAllOf

`func NewConnectorSshMessageAllOf(classId string, objectType string, ) *ConnectorSshMessageAllOf`

NewConnectorSshMessageAllOf instantiates a new ConnectorSshMessageAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorSshMessageAllOfWithDefaults

`func NewConnectorSshMessageAllOfWithDefaults() *ConnectorSshMessageAllOf`

NewConnectorSshMessageAllOfWithDefaults instantiates a new ConnectorSshMessageAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ConnectorSshMessageAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ConnectorSshMessageAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ConnectorSshMessageAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ConnectorSshMessageAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ConnectorSshMessageAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ConnectorSshMessageAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetExpectPrompts

`func (o *ConnectorSshMessageAllOf) GetExpectPrompts() []ConnectorExpectPrompt`

GetExpectPrompts returns the ExpectPrompts field if non-nil, zero value otherwise.

### GetExpectPromptsOk

`func (o *ConnectorSshMessageAllOf) GetExpectPromptsOk() (*[]ConnectorExpectPrompt, bool)`

GetExpectPromptsOk returns a tuple with the ExpectPrompts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectPrompts

`func (o *ConnectorSshMessageAllOf) SetExpectPrompts(v []ConnectorExpectPrompt)`

SetExpectPrompts sets ExpectPrompts field to given value.

### HasExpectPrompts

`func (o *ConnectorSshMessageAllOf) HasExpectPrompts() bool`

HasExpectPrompts returns a boolean if a field has been set.

### SetExpectPromptsNil

`func (o *ConnectorSshMessageAllOf) SetExpectPromptsNil(b bool)`

 SetExpectPromptsNil sets the value for ExpectPrompts to be an explicit nil

### UnsetExpectPrompts
`func (o *ConnectorSshMessageAllOf) UnsetExpectPrompts()`

UnsetExpectPrompts ensures that no value is present for ExpectPrompts, not even an explicit nil
### GetMsgType

`func (o *ConnectorSshMessageAllOf) GetMsgType() int64`

GetMsgType returns the MsgType field if non-nil, zero value otherwise.

### GetMsgTypeOk

`func (o *ConnectorSshMessageAllOf) GetMsgTypeOk() (*int64, bool)`

GetMsgTypeOk returns a tuple with the MsgType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgType

`func (o *ConnectorSshMessageAllOf) SetMsgType(v int64)`

SetMsgType sets MsgType field to given value.

### HasMsgType

`func (o *ConnectorSshMessageAllOf) HasMsgType() bool`

HasMsgType returns a boolean if a field has been set.

### GetSessionId

`func (o *ConnectorSshMessageAllOf) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *ConnectorSshMessageAllOf) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *ConnectorSshMessageAllOf) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *ConnectorSshMessageAllOf) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetShellPrompt

`func (o *ConnectorSshMessageAllOf) GetShellPrompt() string`

GetShellPrompt returns the ShellPrompt field if non-nil, zero value otherwise.

### GetShellPromptOk

`func (o *ConnectorSshMessageAllOf) GetShellPromptOk() (*string, bool)`

GetShellPromptOk returns a tuple with the ShellPrompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShellPrompt

`func (o *ConnectorSshMessageAllOf) SetShellPrompt(v string)`

SetShellPrompt sets ShellPrompt field to given value.

### HasShellPrompt

`func (o *ConnectorSshMessageAllOf) HasShellPrompt() bool`

HasShellPrompt returns a boolean if a field has been set.

### GetStream

`func (o *ConnectorSshMessageAllOf) GetStream() string`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *ConnectorSshMessageAllOf) GetStreamOk() (*string, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *ConnectorSshMessageAllOf) SetStream(v string)`

SetStream sets Stream field to given value.

### HasStream

`func (o *ConnectorSshMessageAllOf) HasStream() bool`

HasStream returns a boolean if a field has been set.

### GetTimeout

`func (o *ConnectorSshMessageAllOf) GetTimeout() int64`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *ConnectorSshMessageAllOf) GetTimeoutOk() (*int64, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *ConnectorSshMessageAllOf) SetTimeout(v int64)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *ConnectorSshMessageAllOf) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# ConnectorCommandControlMessageAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dir** | Pointer to **string** | The working directory of the command. If empty command is executed in the same directory the device connector process was called. | [optional] 
**MsgType** | Pointer to **string** | Message carrying the operation to perform. | [optional] 
**Stream** | Pointer to **string** | The command to execute. Commands must be whitelisted by platform implementation, if a command does not match any whitelisted command patterns an error will be returned to the requesting service on command start. | [optional] 
**Terminal** | Pointer to **bool** | Indicates that a pseudo terminal should be attached to the command. Used for interactive commands. e.g A cross launch cli. | [optional] 
**Timeout** | Pointer to **int64** | The timeout for the command to complete and exit after starting or receiving input. If timeout is not set a default of 10 minutes will be used. If there is input to the command stream the timeout is extended. | [optional] 

## Methods

### NewConnectorCommandControlMessageAllOf

`func NewConnectorCommandControlMessageAllOf() *ConnectorCommandControlMessageAllOf`

NewConnectorCommandControlMessageAllOf instantiates a new ConnectorCommandControlMessageAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorCommandControlMessageAllOfWithDefaults

`func NewConnectorCommandControlMessageAllOfWithDefaults() *ConnectorCommandControlMessageAllOf`

NewConnectorCommandControlMessageAllOfWithDefaults instantiates a new ConnectorCommandControlMessageAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDir

`func (o *ConnectorCommandControlMessageAllOf) GetDir() string`

GetDir returns the Dir field if non-nil, zero value otherwise.

### GetDirOk

`func (o *ConnectorCommandControlMessageAllOf) GetDirOk() (*string, bool)`

GetDirOk returns a tuple with the Dir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDir

`func (o *ConnectorCommandControlMessageAllOf) SetDir(v string)`

SetDir sets Dir field to given value.

### HasDir

`func (o *ConnectorCommandControlMessageAllOf) HasDir() bool`

HasDir returns a boolean if a field has been set.

### GetMsgType

`func (o *ConnectorCommandControlMessageAllOf) GetMsgType() string`

GetMsgType returns the MsgType field if non-nil, zero value otherwise.

### GetMsgTypeOk

`func (o *ConnectorCommandControlMessageAllOf) GetMsgTypeOk() (*string, bool)`

GetMsgTypeOk returns a tuple with the MsgType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgType

`func (o *ConnectorCommandControlMessageAllOf) SetMsgType(v string)`

SetMsgType sets MsgType field to given value.

### HasMsgType

`func (o *ConnectorCommandControlMessageAllOf) HasMsgType() bool`

HasMsgType returns a boolean if a field has been set.

### GetStream

`func (o *ConnectorCommandControlMessageAllOf) GetStream() string`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *ConnectorCommandControlMessageAllOf) GetStreamOk() (*string, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *ConnectorCommandControlMessageAllOf) SetStream(v string)`

SetStream sets Stream field to given value.

### HasStream

`func (o *ConnectorCommandControlMessageAllOf) HasStream() bool`

HasStream returns a boolean if a field has been set.

### GetTerminal

`func (o *ConnectorCommandControlMessageAllOf) GetTerminal() bool`

GetTerminal returns the Terminal field if non-nil, zero value otherwise.

### GetTerminalOk

`func (o *ConnectorCommandControlMessageAllOf) GetTerminalOk() (*bool, bool)`

GetTerminalOk returns a tuple with the Terminal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminal

`func (o *ConnectorCommandControlMessageAllOf) SetTerminal(v bool)`

SetTerminal sets Terminal field to given value.

### HasTerminal

`func (o *ConnectorCommandControlMessageAllOf) HasTerminal() bool`

HasTerminal returns a boolean if a field has been set.

### GetTimeout

`func (o *ConnectorCommandControlMessageAllOf) GetTimeout() int64`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *ConnectorCommandControlMessageAllOf) GetTimeoutOk() (*int64, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *ConnectorCommandControlMessageAllOf) SetTimeout(v int64)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *ConnectorCommandControlMessageAllOf) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



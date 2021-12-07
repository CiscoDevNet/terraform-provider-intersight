# ConnectorCommandControlMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "connector.CommandControlMessage"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "connector.CommandControlMessage"]
**Dir** | Pointer to **string** | The working directory of the command. If empty command is executed in the same directory the device connector process was called. | [optional] 
**MsgType** | Pointer to **string** | Message carrying the operation to perform. | [optional] 
**Stream** | Pointer to **string** | The command to execute. Commands must be whitelisted by platform implementation, if a command does not match any whitelisted command patterns an error will be returned to the requesting service on command start. | [optional] 
**Terminal** | Pointer to **bool** | Indicates that a pseudo terminal should be attached to the command. Used for interactive commands. e.g A cross launch cli. | [optional] 
**Timeout** | Pointer to **int64** | The timeout for the command to complete and exit after starting or receiving input. If timeout is not set a default of 10 minutes will be used. If there is input to the command stream the timeout is extended. | [optional] 

## Methods

### NewConnectorCommandControlMessage

`func NewConnectorCommandControlMessage(classId string, objectType string, ) *ConnectorCommandControlMessage`

NewConnectorCommandControlMessage instantiates a new ConnectorCommandControlMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorCommandControlMessageWithDefaults

`func NewConnectorCommandControlMessageWithDefaults() *ConnectorCommandControlMessage`

NewConnectorCommandControlMessageWithDefaults instantiates a new ConnectorCommandControlMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ConnectorCommandControlMessage) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ConnectorCommandControlMessage) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ConnectorCommandControlMessage) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ConnectorCommandControlMessage) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ConnectorCommandControlMessage) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ConnectorCommandControlMessage) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDir

`func (o *ConnectorCommandControlMessage) GetDir() string`

GetDir returns the Dir field if non-nil, zero value otherwise.

### GetDirOk

`func (o *ConnectorCommandControlMessage) GetDirOk() (*string, bool)`

GetDirOk returns a tuple with the Dir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDir

`func (o *ConnectorCommandControlMessage) SetDir(v string)`

SetDir sets Dir field to given value.

### HasDir

`func (o *ConnectorCommandControlMessage) HasDir() bool`

HasDir returns a boolean if a field has been set.

### GetMsgType

`func (o *ConnectorCommandControlMessage) GetMsgType() string`

GetMsgType returns the MsgType field if non-nil, zero value otherwise.

### GetMsgTypeOk

`func (o *ConnectorCommandControlMessage) GetMsgTypeOk() (*string, bool)`

GetMsgTypeOk returns a tuple with the MsgType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgType

`func (o *ConnectorCommandControlMessage) SetMsgType(v string)`

SetMsgType sets MsgType field to given value.

### HasMsgType

`func (o *ConnectorCommandControlMessage) HasMsgType() bool`

HasMsgType returns a boolean if a field has been set.

### GetStream

`func (o *ConnectorCommandControlMessage) GetStream() string`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *ConnectorCommandControlMessage) GetStreamOk() (*string, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *ConnectorCommandControlMessage) SetStream(v string)`

SetStream sets Stream field to given value.

### HasStream

`func (o *ConnectorCommandControlMessage) HasStream() bool`

HasStream returns a boolean if a field has been set.

### GetTerminal

`func (o *ConnectorCommandControlMessage) GetTerminal() bool`

GetTerminal returns the Terminal field if non-nil, zero value otherwise.

### GetTerminalOk

`func (o *ConnectorCommandControlMessage) GetTerminalOk() (*bool, bool)`

GetTerminalOk returns a tuple with the Terminal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminal

`func (o *ConnectorCommandControlMessage) SetTerminal(v bool)`

SetTerminal sets Terminal field to given value.

### HasTerminal

`func (o *ConnectorCommandControlMessage) HasTerminal() bool`

HasTerminal returns a boolean if a field has been set.

### GetTimeout

`func (o *ConnectorCommandControlMessage) GetTimeout() int64`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *ConnectorCommandControlMessage) GetTimeoutOk() (*int64, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *ConnectorCommandControlMessage) SetTimeout(v int64)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *ConnectorCommandControlMessage) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



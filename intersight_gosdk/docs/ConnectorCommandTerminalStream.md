# ConnectorCommandTerminalStream

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "connector.CommandTerminalStream"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "connector.CommandTerminalStream"]
**MsgType** | Pointer to **string** | The type of data this message contains. | [optional] 
**Sequence** | Pointer to **int64** | Sequence of the message within a session to handle out-of-order delivery. | [optional] 
**Stream** | Pointer to **string** | The input/output payload to/from the pseudo terminal session. When sent from the cloud service if the msgType is CommandInput stream is piped to stdin of the command or a resize message if msgType is CommandResize. From the device connector value is always the combined output of stdout &amp; stderr. | [optional] 

## Methods

### NewConnectorCommandTerminalStream

`func NewConnectorCommandTerminalStream(classId string, objectType string, ) *ConnectorCommandTerminalStream`

NewConnectorCommandTerminalStream instantiates a new ConnectorCommandTerminalStream object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorCommandTerminalStreamWithDefaults

`func NewConnectorCommandTerminalStreamWithDefaults() *ConnectorCommandTerminalStream`

NewConnectorCommandTerminalStreamWithDefaults instantiates a new ConnectorCommandTerminalStream object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ConnectorCommandTerminalStream) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ConnectorCommandTerminalStream) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ConnectorCommandTerminalStream) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ConnectorCommandTerminalStream) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ConnectorCommandTerminalStream) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ConnectorCommandTerminalStream) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetMsgType

`func (o *ConnectorCommandTerminalStream) GetMsgType() string`

GetMsgType returns the MsgType field if non-nil, zero value otherwise.

### GetMsgTypeOk

`func (o *ConnectorCommandTerminalStream) GetMsgTypeOk() (*string, bool)`

GetMsgTypeOk returns a tuple with the MsgType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgType

`func (o *ConnectorCommandTerminalStream) SetMsgType(v string)`

SetMsgType sets MsgType field to given value.

### HasMsgType

`func (o *ConnectorCommandTerminalStream) HasMsgType() bool`

HasMsgType returns a boolean if a field has been set.

### GetSequence

`func (o *ConnectorCommandTerminalStream) GetSequence() int64`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *ConnectorCommandTerminalStream) GetSequenceOk() (*int64, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *ConnectorCommandTerminalStream) SetSequence(v int64)`

SetSequence sets Sequence field to given value.

### HasSequence

`func (o *ConnectorCommandTerminalStream) HasSequence() bool`

HasSequence returns a boolean if a field has been set.

### GetStream

`func (o *ConnectorCommandTerminalStream) GetStream() string`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *ConnectorCommandTerminalStream) GetStreamOk() (*string, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *ConnectorCommandTerminalStream) SetStream(v string)`

SetStream sets Stream field to given value.

### HasStream

`func (o *ConnectorCommandTerminalStream) HasStream() bool`

HasStream returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



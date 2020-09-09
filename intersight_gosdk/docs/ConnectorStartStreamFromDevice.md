# ConnectorStartStreamFromDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MemberId** | Pointer to **string** | The asset.ClusterMember member identity that is opening this stream. | [optional] 
**MemberStream** | Pointer to **bool** | The stream is to be started against the cluster member. | [optional] 
**StreamConfig** | Pointer to **interface{}** | Any extra configuration needed to open/identify a stream. | [optional] 
**StreamType** | Pointer to **string** | Identifies the type of stream to open to the device. The Intersight service will validate that the device should open a stream of this type and if so build a stream configuration and send it down to the device. The streamType should identify a unique stream to open to a device, that is if the device sends a stream open message and a stream of that type is already open in the cloud the existing stream should be re-used. | [optional] 
**Topic** | Pointer to **string** | The topic the device should send the stream open message to. | [optional] 

## Methods

### NewConnectorStartStreamFromDevice

`func NewConnectorStartStreamFromDevice() *ConnectorStartStreamFromDevice`

NewConnectorStartStreamFromDevice instantiates a new ConnectorStartStreamFromDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorStartStreamFromDeviceWithDefaults

`func NewConnectorStartStreamFromDeviceWithDefaults() *ConnectorStartStreamFromDevice`

NewConnectorStartStreamFromDeviceWithDefaults instantiates a new ConnectorStartStreamFromDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMemberId

`func (o *ConnectorStartStreamFromDevice) GetMemberId() string`

GetMemberId returns the MemberId field if non-nil, zero value otherwise.

### GetMemberIdOk

`func (o *ConnectorStartStreamFromDevice) GetMemberIdOk() (*string, bool)`

GetMemberIdOk returns a tuple with the MemberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberId

`func (o *ConnectorStartStreamFromDevice) SetMemberId(v string)`

SetMemberId sets MemberId field to given value.

### HasMemberId

`func (o *ConnectorStartStreamFromDevice) HasMemberId() bool`

HasMemberId returns a boolean if a field has been set.

### GetMemberStream

`func (o *ConnectorStartStreamFromDevice) GetMemberStream() bool`

GetMemberStream returns the MemberStream field if non-nil, zero value otherwise.

### GetMemberStreamOk

`func (o *ConnectorStartStreamFromDevice) GetMemberStreamOk() (*bool, bool)`

GetMemberStreamOk returns a tuple with the MemberStream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberStream

`func (o *ConnectorStartStreamFromDevice) SetMemberStream(v bool)`

SetMemberStream sets MemberStream field to given value.

### HasMemberStream

`func (o *ConnectorStartStreamFromDevice) HasMemberStream() bool`

HasMemberStream returns a boolean if a field has been set.

### GetStreamConfig

`func (o *ConnectorStartStreamFromDevice) GetStreamConfig() interface{}`

GetStreamConfig returns the StreamConfig field if non-nil, zero value otherwise.

### GetStreamConfigOk

`func (o *ConnectorStartStreamFromDevice) GetStreamConfigOk() (*interface{}, bool)`

GetStreamConfigOk returns a tuple with the StreamConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamConfig

`func (o *ConnectorStartStreamFromDevice) SetStreamConfig(v interface{})`

SetStreamConfig sets StreamConfig field to given value.

### HasStreamConfig

`func (o *ConnectorStartStreamFromDevice) HasStreamConfig() bool`

HasStreamConfig returns a boolean if a field has been set.

### SetStreamConfigNil

`func (o *ConnectorStartStreamFromDevice) SetStreamConfigNil(b bool)`

 SetStreamConfigNil sets the value for StreamConfig to be an explicit nil

### UnsetStreamConfig
`func (o *ConnectorStartStreamFromDevice) UnsetStreamConfig()`

UnsetStreamConfig ensures that no value is present for StreamConfig, not even an explicit nil
### GetStreamType

`func (o *ConnectorStartStreamFromDevice) GetStreamType() string`

GetStreamType returns the StreamType field if non-nil, zero value otherwise.

### GetStreamTypeOk

`func (o *ConnectorStartStreamFromDevice) GetStreamTypeOk() (*string, bool)`

GetStreamTypeOk returns a tuple with the StreamType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamType

`func (o *ConnectorStartStreamFromDevice) SetStreamType(v string)`

SetStreamType sets StreamType field to given value.

### HasStreamType

`func (o *ConnectorStartStreamFromDevice) HasStreamType() bool`

HasStreamType returns a boolean if a field has been set.

### GetTopic

`func (o *ConnectorStartStreamFromDevice) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *ConnectorStartStreamFromDevice) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *ConnectorStartStreamFromDevice) SetTopic(v string)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *ConnectorStartStreamFromDevice) HasTopic() bool`

HasTopic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



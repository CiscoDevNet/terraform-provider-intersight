# ConnectorStartStreamFromDeviceAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MemberId** | Pointer to **string** | The asset.ClusterMember member identity that is opening this stream. | [optional] 
**MemberStream** | Pointer to **bool** | The stream is to be started against the cluster member. | [optional] 
**StreamConfig** | Pointer to **interface{}** | Any extra configuration needed to open/identify a stream. | [optional] 
**StreamType** | Pointer to **string** | Identifies the type of stream to open to the device. The Intersight service will validate that the device should open a stream of this type and if so build a stream configuration and send it down to the device. The streamType should identify a unique stream to open to a device, that is if the device sends a stream open message and a stream of that type is already open in the cloud the existing stream should be re-used. | [optional] 
**Topic** | Pointer to **string** | The topic the device should send the stream open message to. | [optional] 

## Methods

### NewConnectorStartStreamFromDeviceAllOf

`func NewConnectorStartStreamFromDeviceAllOf() *ConnectorStartStreamFromDeviceAllOf`

NewConnectorStartStreamFromDeviceAllOf instantiates a new ConnectorStartStreamFromDeviceAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorStartStreamFromDeviceAllOfWithDefaults

`func NewConnectorStartStreamFromDeviceAllOfWithDefaults() *ConnectorStartStreamFromDeviceAllOf`

NewConnectorStartStreamFromDeviceAllOfWithDefaults instantiates a new ConnectorStartStreamFromDeviceAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMemberId

`func (o *ConnectorStartStreamFromDeviceAllOf) GetMemberId() string`

GetMemberId returns the MemberId field if non-nil, zero value otherwise.

### GetMemberIdOk

`func (o *ConnectorStartStreamFromDeviceAllOf) GetMemberIdOk() (*string, bool)`

GetMemberIdOk returns a tuple with the MemberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberId

`func (o *ConnectorStartStreamFromDeviceAllOf) SetMemberId(v string)`

SetMemberId sets MemberId field to given value.

### HasMemberId

`func (o *ConnectorStartStreamFromDeviceAllOf) HasMemberId() bool`

HasMemberId returns a boolean if a field has been set.

### GetMemberStream

`func (o *ConnectorStartStreamFromDeviceAllOf) GetMemberStream() bool`

GetMemberStream returns the MemberStream field if non-nil, zero value otherwise.

### GetMemberStreamOk

`func (o *ConnectorStartStreamFromDeviceAllOf) GetMemberStreamOk() (*bool, bool)`

GetMemberStreamOk returns a tuple with the MemberStream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberStream

`func (o *ConnectorStartStreamFromDeviceAllOf) SetMemberStream(v bool)`

SetMemberStream sets MemberStream field to given value.

### HasMemberStream

`func (o *ConnectorStartStreamFromDeviceAllOf) HasMemberStream() bool`

HasMemberStream returns a boolean if a field has been set.

### GetStreamConfig

`func (o *ConnectorStartStreamFromDeviceAllOf) GetStreamConfig() interface{}`

GetStreamConfig returns the StreamConfig field if non-nil, zero value otherwise.

### GetStreamConfigOk

`func (o *ConnectorStartStreamFromDeviceAllOf) GetStreamConfigOk() (*interface{}, bool)`

GetStreamConfigOk returns a tuple with the StreamConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamConfig

`func (o *ConnectorStartStreamFromDeviceAllOf) SetStreamConfig(v interface{})`

SetStreamConfig sets StreamConfig field to given value.

### HasStreamConfig

`func (o *ConnectorStartStreamFromDeviceAllOf) HasStreamConfig() bool`

HasStreamConfig returns a boolean if a field has been set.

### SetStreamConfigNil

`func (o *ConnectorStartStreamFromDeviceAllOf) SetStreamConfigNil(b bool)`

 SetStreamConfigNil sets the value for StreamConfig to be an explicit nil

### UnsetStreamConfig
`func (o *ConnectorStartStreamFromDeviceAllOf) UnsetStreamConfig()`

UnsetStreamConfig ensures that no value is present for StreamConfig, not even an explicit nil
### GetStreamType

`func (o *ConnectorStartStreamFromDeviceAllOf) GetStreamType() string`

GetStreamType returns the StreamType field if non-nil, zero value otherwise.

### GetStreamTypeOk

`func (o *ConnectorStartStreamFromDeviceAllOf) GetStreamTypeOk() (*string, bool)`

GetStreamTypeOk returns a tuple with the StreamType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamType

`func (o *ConnectorStartStreamFromDeviceAllOf) SetStreamType(v string)`

SetStreamType sets StreamType field to given value.

### HasStreamType

`func (o *ConnectorStartStreamFromDeviceAllOf) HasStreamType() bool`

HasStreamType returns a boolean if a field has been set.

### GetTopic

`func (o *ConnectorStartStreamFromDeviceAllOf) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *ConnectorStartStreamFromDeviceAllOf) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *ConnectorStartStreamFromDeviceAllOf) SetTopic(v string)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *ConnectorStartStreamFromDeviceAllOf) HasTopic() bool`

HasTopic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



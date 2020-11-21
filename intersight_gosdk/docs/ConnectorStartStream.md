# ConnectorStartStream

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "connector.StartStream"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "connector.StartStream"]
**BatchSize** | Pointer to **int64** | The number of outputs from a plugin to collect into a single message. Applicable only to streams that involve polling plugins and plugins which support emitting batchable data. Default value of zero indicates no batching. | [optional] 
**ForceRebuild** | Pointer to **bool** | Flag to force a rebuild of an existing stream. To be used if a stream is unable to recover itself in response to dropped messages. | [optional] 
**Input** | Pointer to **string** | Input to the plugin to start the start the stream or collect stream messages. | [optional] 
**KeepAliveInterval** | Pointer to **int64** | Interval at which device should emit a keepalive message for this stream. Device will also expect a keepalive response from the cloud within the interval. If zero, no keepalive is required and stream should not timeout. | [optional] 
**PluginName** | Pointer to **string** | The plugin to run the stream on. | [optional] 
**PollInterval** | Pointer to **int64** | The desired interval to emit messages from this stream. The stream plugin will poll plugins at this interval to create a stream event. | [optional] 
**Priority** | Pointer to **int64** | The priority level to apply to messages emitted by this stream. | [optional] 
**ProtocolVersion** | Pointer to **int64** | The version of the device connector stream protocol. Used to change behavior of the device connector stream plugin based on the version of the Intersight service. Allows for multiple versions of Intersight services to interact with the stream plugin of devices. | [optional] 
**ResponseTopic** | Pointer to **string** | The topic for the device connector to publish messages to. | [optional] 

## Methods

### NewConnectorStartStream

`func NewConnectorStartStream(classId string, objectType string, ) *ConnectorStartStream`

NewConnectorStartStream instantiates a new ConnectorStartStream object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorStartStreamWithDefaults

`func NewConnectorStartStreamWithDefaults() *ConnectorStartStream`

NewConnectorStartStreamWithDefaults instantiates a new ConnectorStartStream object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ConnectorStartStream) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ConnectorStartStream) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ConnectorStartStream) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ConnectorStartStream) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ConnectorStartStream) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ConnectorStartStream) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBatchSize

`func (o *ConnectorStartStream) GetBatchSize() int64`

GetBatchSize returns the BatchSize field if non-nil, zero value otherwise.

### GetBatchSizeOk

`func (o *ConnectorStartStream) GetBatchSizeOk() (*int64, bool)`

GetBatchSizeOk returns a tuple with the BatchSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchSize

`func (o *ConnectorStartStream) SetBatchSize(v int64)`

SetBatchSize sets BatchSize field to given value.

### HasBatchSize

`func (o *ConnectorStartStream) HasBatchSize() bool`

HasBatchSize returns a boolean if a field has been set.

### GetForceRebuild

`func (o *ConnectorStartStream) GetForceRebuild() bool`

GetForceRebuild returns the ForceRebuild field if non-nil, zero value otherwise.

### GetForceRebuildOk

`func (o *ConnectorStartStream) GetForceRebuildOk() (*bool, bool)`

GetForceRebuildOk returns a tuple with the ForceRebuild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceRebuild

`func (o *ConnectorStartStream) SetForceRebuild(v bool)`

SetForceRebuild sets ForceRebuild field to given value.

### HasForceRebuild

`func (o *ConnectorStartStream) HasForceRebuild() bool`

HasForceRebuild returns a boolean if a field has been set.

### GetInput

`func (o *ConnectorStartStream) GetInput() string`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *ConnectorStartStream) GetInputOk() (*string, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *ConnectorStartStream) SetInput(v string)`

SetInput sets Input field to given value.

### HasInput

`func (o *ConnectorStartStream) HasInput() bool`

HasInput returns a boolean if a field has been set.

### GetKeepAliveInterval

`func (o *ConnectorStartStream) GetKeepAliveInterval() int64`

GetKeepAliveInterval returns the KeepAliveInterval field if non-nil, zero value otherwise.

### GetKeepAliveIntervalOk

`func (o *ConnectorStartStream) GetKeepAliveIntervalOk() (*int64, bool)`

GetKeepAliveIntervalOk returns a tuple with the KeepAliveInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepAliveInterval

`func (o *ConnectorStartStream) SetKeepAliveInterval(v int64)`

SetKeepAliveInterval sets KeepAliveInterval field to given value.

### HasKeepAliveInterval

`func (o *ConnectorStartStream) HasKeepAliveInterval() bool`

HasKeepAliveInterval returns a boolean if a field has been set.

### GetPluginName

`func (o *ConnectorStartStream) GetPluginName() string`

GetPluginName returns the PluginName field if non-nil, zero value otherwise.

### GetPluginNameOk

`func (o *ConnectorStartStream) GetPluginNameOk() (*string, bool)`

GetPluginNameOk returns a tuple with the PluginName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginName

`func (o *ConnectorStartStream) SetPluginName(v string)`

SetPluginName sets PluginName field to given value.

### HasPluginName

`func (o *ConnectorStartStream) HasPluginName() bool`

HasPluginName returns a boolean if a field has been set.

### GetPollInterval

`func (o *ConnectorStartStream) GetPollInterval() int64`

GetPollInterval returns the PollInterval field if non-nil, zero value otherwise.

### GetPollIntervalOk

`func (o *ConnectorStartStream) GetPollIntervalOk() (*int64, bool)`

GetPollIntervalOk returns a tuple with the PollInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPollInterval

`func (o *ConnectorStartStream) SetPollInterval(v int64)`

SetPollInterval sets PollInterval field to given value.

### HasPollInterval

`func (o *ConnectorStartStream) HasPollInterval() bool`

HasPollInterval returns a boolean if a field has been set.

### GetPriority

`func (o *ConnectorStartStream) GetPriority() int64`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *ConnectorStartStream) GetPriorityOk() (*int64, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *ConnectorStartStream) SetPriority(v int64)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *ConnectorStartStream) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetProtocolVersion

`func (o *ConnectorStartStream) GetProtocolVersion() int64`

GetProtocolVersion returns the ProtocolVersion field if non-nil, zero value otherwise.

### GetProtocolVersionOk

`func (o *ConnectorStartStream) GetProtocolVersionOk() (*int64, bool)`

GetProtocolVersionOk returns a tuple with the ProtocolVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolVersion

`func (o *ConnectorStartStream) SetProtocolVersion(v int64)`

SetProtocolVersion sets ProtocolVersion field to given value.

### HasProtocolVersion

`func (o *ConnectorStartStream) HasProtocolVersion() bool`

HasProtocolVersion returns a boolean if a field has been set.

### GetResponseTopic

`func (o *ConnectorStartStream) GetResponseTopic() string`

GetResponseTopic returns the ResponseTopic field if non-nil, zero value otherwise.

### GetResponseTopicOk

`func (o *ConnectorStartStream) GetResponseTopicOk() (*string, bool)`

GetResponseTopicOk returns a tuple with the ResponseTopic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseTopic

`func (o *ConnectorStartStream) SetResponseTopic(v string)`

SetResponseTopic sets ResponseTopic field to given value.

### HasResponseTopic

`func (o *ConnectorStartStream) HasResponseTopic() bool`

HasResponseTopic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



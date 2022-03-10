# ConnectorStreamKeepalive

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "connector.StreamKeepalive"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "connector.StreamKeepalive"]
**Paused** | Pointer to **bool** | Stream has been paused and is not actively sending outputs to cloud. Device will pause streams when it does not receive an acknowledgment of messages after sending number messages cloud has configured in the stream acknowledge rate (set in StartStream.ackRate field on stream start). If cloud receives a keepalive messages with paused set it will attempt to acknowledge the latest messages processed from the device in order to unpause the stream. | [optional] 

## Methods

### NewConnectorStreamKeepalive

`func NewConnectorStreamKeepalive(classId string, objectType string, ) *ConnectorStreamKeepalive`

NewConnectorStreamKeepalive instantiates a new ConnectorStreamKeepalive object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorStreamKeepaliveWithDefaults

`func NewConnectorStreamKeepaliveWithDefaults() *ConnectorStreamKeepalive`

NewConnectorStreamKeepaliveWithDefaults instantiates a new ConnectorStreamKeepalive object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ConnectorStreamKeepalive) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ConnectorStreamKeepalive) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ConnectorStreamKeepalive) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ConnectorStreamKeepalive) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ConnectorStreamKeepalive) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ConnectorStreamKeepalive) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetPaused

`func (o *ConnectorStreamKeepalive) GetPaused() bool`

GetPaused returns the Paused field if non-nil, zero value otherwise.

### GetPausedOk

`func (o *ConnectorStreamKeepalive) GetPausedOk() (*bool, bool)`

GetPausedOk returns a tuple with the Paused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaused

`func (o *ConnectorStreamKeepalive) SetPaused(v bool)`

SetPaused sets Paused field to given value.

### HasPaused

`func (o *ConnectorStreamKeepalive) HasPaused() bool`

HasPaused returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



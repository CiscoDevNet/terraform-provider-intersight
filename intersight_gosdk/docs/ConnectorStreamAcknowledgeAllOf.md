# ConnectorStreamAcknowledgeAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "connector.StreamAcknowledge"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "connector.StreamAcknowledge"]
**AckSequence** | Pointer to **int64** | The latest message sequence processed in the cloud. Device connector will drop all messages up to this sequence from its cache. | [optional] 

## Methods

### NewConnectorStreamAcknowledgeAllOf

`func NewConnectorStreamAcknowledgeAllOf(classId string, objectType string, ) *ConnectorStreamAcknowledgeAllOf`

NewConnectorStreamAcknowledgeAllOf instantiates a new ConnectorStreamAcknowledgeAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorStreamAcknowledgeAllOfWithDefaults

`func NewConnectorStreamAcknowledgeAllOfWithDefaults() *ConnectorStreamAcknowledgeAllOf`

NewConnectorStreamAcknowledgeAllOfWithDefaults instantiates a new ConnectorStreamAcknowledgeAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ConnectorStreamAcknowledgeAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ConnectorStreamAcknowledgeAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ConnectorStreamAcknowledgeAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ConnectorStreamAcknowledgeAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ConnectorStreamAcknowledgeAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ConnectorStreamAcknowledgeAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAckSequence

`func (o *ConnectorStreamAcknowledgeAllOf) GetAckSequence() int64`

GetAckSequence returns the AckSequence field if non-nil, zero value otherwise.

### GetAckSequenceOk

`func (o *ConnectorStreamAcknowledgeAllOf) GetAckSequenceOk() (*int64, bool)`

GetAckSequenceOk returns a tuple with the AckSequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckSequence

`func (o *ConnectorStreamAcknowledgeAllOf) SetAckSequence(v int64)`

SetAckSequence sets AckSequence field to given value.

### HasAckSequence

`func (o *ConnectorStreamAcknowledgeAllOf) HasAckSequence() bool`

HasAckSequence returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



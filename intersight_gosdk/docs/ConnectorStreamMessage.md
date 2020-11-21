# ConnectorStreamMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**StreamName** | Pointer to **string** | The requested stream name. Stream names are unique per device endpoint. | [optional] 

## Methods

### NewConnectorStreamMessage

`func NewConnectorStreamMessage(classId string, objectType string, ) *ConnectorStreamMessage`

NewConnectorStreamMessage instantiates a new ConnectorStreamMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorStreamMessageWithDefaults

`func NewConnectorStreamMessageWithDefaults() *ConnectorStreamMessage`

NewConnectorStreamMessageWithDefaults instantiates a new ConnectorStreamMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ConnectorStreamMessage) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ConnectorStreamMessage) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ConnectorStreamMessage) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ConnectorStreamMessage) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ConnectorStreamMessage) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ConnectorStreamMessage) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetStreamName

`func (o *ConnectorStreamMessage) GetStreamName() string`

GetStreamName returns the StreamName field if non-nil, zero value otherwise.

### GetStreamNameOk

`func (o *ConnectorStreamMessage) GetStreamNameOk() (*string, bool)`

GetStreamNameOk returns a tuple with the StreamName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamName

`func (o *ConnectorStreamMessage) SetStreamName(v string)`

SetStreamName sets StreamName field to given value.

### HasStreamName

`func (o *ConnectorStreamMessage) HasStreamName() bool`

HasStreamName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# ConnectorStreamMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StreamName** | Pointer to **string** | The requested stream name. Stream names are unique per device endpoint. | [optional] 

## Methods

### NewConnectorStreamMessage

`func NewConnectorStreamMessage() *ConnectorStreamMessage`

NewConnectorStreamMessage instantiates a new ConnectorStreamMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorStreamMessageWithDefaults

`func NewConnectorStreamMessageWithDefaults() *ConnectorStreamMessage`

NewConnectorStreamMessageWithDefaults instantiates a new ConnectorStreamMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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



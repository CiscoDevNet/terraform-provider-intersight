# ConnectorFetchStreamMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sequences** | Pointer to **[]int64** |  | [optional] 

## Methods

### NewConnectorFetchStreamMessage

`func NewConnectorFetchStreamMessage() *ConnectorFetchStreamMessage`

NewConnectorFetchStreamMessage instantiates a new ConnectorFetchStreamMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorFetchStreamMessageWithDefaults

`func NewConnectorFetchStreamMessageWithDefaults() *ConnectorFetchStreamMessage`

NewConnectorFetchStreamMessageWithDefaults instantiates a new ConnectorFetchStreamMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSequences

`func (o *ConnectorFetchStreamMessage) GetSequences() []int64`

GetSequences returns the Sequences field if non-nil, zero value otherwise.

### GetSequencesOk

`func (o *ConnectorFetchStreamMessage) GetSequencesOk() (*[]int64, bool)`

GetSequencesOk returns a tuple with the Sequences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequences

`func (o *ConnectorFetchStreamMessage) SetSequences(v []int64)`

SetSequences sets Sequences field to given value.

### HasSequences

`func (o *ConnectorFetchStreamMessage) HasSequences() bool`

HasSequences returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



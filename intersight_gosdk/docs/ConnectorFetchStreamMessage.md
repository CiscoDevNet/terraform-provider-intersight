# ConnectorFetchStreamMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "connector.FetchStreamMessage"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "connector.FetchStreamMessage"]
**Sequences** | Pointer to **[]int64** |  | [optional] 

## Methods

### NewConnectorFetchStreamMessage

`func NewConnectorFetchStreamMessage(classId string, objectType string, ) *ConnectorFetchStreamMessage`

NewConnectorFetchStreamMessage instantiates a new ConnectorFetchStreamMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorFetchStreamMessageWithDefaults

`func NewConnectorFetchStreamMessageWithDefaults() *ConnectorFetchStreamMessage`

NewConnectorFetchStreamMessageWithDefaults instantiates a new ConnectorFetchStreamMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ConnectorFetchStreamMessage) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ConnectorFetchStreamMessage) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ConnectorFetchStreamMessage) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ConnectorFetchStreamMessage) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ConnectorFetchStreamMessage) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ConnectorFetchStreamMessage) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


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

### SetSequencesNil

`func (o *ConnectorFetchStreamMessage) SetSequencesNil(b bool)`

 SetSequencesNil sets the value for Sequences to be an explicit nil

### UnsetSequences
`func (o *ConnectorFetchStreamMessage) UnsetSequences()`

UnsetSequences ensures that no value is present for Sequences, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



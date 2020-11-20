# ConnectorFetchStreamMessageAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "connector.FetchStreamMessage"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "connector.FetchStreamMessage"]
**Sequences** | Pointer to **[]int64** |  | [optional] 

## Methods

### NewConnectorFetchStreamMessageAllOf

`func NewConnectorFetchStreamMessageAllOf(classId string, objectType string, ) *ConnectorFetchStreamMessageAllOf`

NewConnectorFetchStreamMessageAllOf instantiates a new ConnectorFetchStreamMessageAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorFetchStreamMessageAllOfWithDefaults

`func NewConnectorFetchStreamMessageAllOfWithDefaults() *ConnectorFetchStreamMessageAllOf`

NewConnectorFetchStreamMessageAllOfWithDefaults instantiates a new ConnectorFetchStreamMessageAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ConnectorFetchStreamMessageAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ConnectorFetchStreamMessageAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ConnectorFetchStreamMessageAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ConnectorFetchStreamMessageAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ConnectorFetchStreamMessageAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ConnectorFetchStreamMessageAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetSequences

`func (o *ConnectorFetchStreamMessageAllOf) GetSequences() []int64`

GetSequences returns the Sequences field if non-nil, zero value otherwise.

### GetSequencesOk

`func (o *ConnectorFetchStreamMessageAllOf) GetSequencesOk() (*[]int64, bool)`

GetSequencesOk returns a tuple with the Sequences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequences

`func (o *ConnectorFetchStreamMessageAllOf) SetSequences(v []int64)`

SetSequences sets Sequences field to given value.

### HasSequences

`func (o *ConnectorFetchStreamMessageAllOf) HasSequences() bool`

HasSequences returns a boolean if a field has been set.

### SetSequencesNil

`func (o *ConnectorFetchStreamMessageAllOf) SetSequencesNil(b bool)`

 SetSequencesNil sets the value for Sequences to be an explicit nil

### UnsetSequences
`func (o *ConnectorFetchStreamMessageAllOf) UnsetSequences()`

UnsetSequences ensures that no value is present for Sequences, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



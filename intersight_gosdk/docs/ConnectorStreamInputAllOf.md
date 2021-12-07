# ConnectorStreamInputAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "connector.StreamInput"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "connector.StreamInput"]
**Input** | Pointer to **string** | The input to write to the stream plugin. | [optional] 

## Methods

### NewConnectorStreamInputAllOf

`func NewConnectorStreamInputAllOf(classId string, objectType string, ) *ConnectorStreamInputAllOf`

NewConnectorStreamInputAllOf instantiates a new ConnectorStreamInputAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorStreamInputAllOfWithDefaults

`func NewConnectorStreamInputAllOfWithDefaults() *ConnectorStreamInputAllOf`

NewConnectorStreamInputAllOfWithDefaults instantiates a new ConnectorStreamInputAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ConnectorStreamInputAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ConnectorStreamInputAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ConnectorStreamInputAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ConnectorStreamInputAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ConnectorStreamInputAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ConnectorStreamInputAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetInput

`func (o *ConnectorStreamInputAllOf) GetInput() string`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *ConnectorStreamInputAllOf) GetInputOk() (*string, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *ConnectorStreamInputAllOf) SetInput(v string)`

SetInput sets Input field to given value.

### HasInput

`func (o *ConnectorStreamInputAllOf) HasInput() bool`

HasInput returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# ConnectorEventLogAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "connector.EventLog"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "connector.EventLog"]
**Contents** | Pointer to **string** | Change contents of this event. | [optional] [readonly] 
**Type** | Pointer to **string** | Type of the change event, e.g. connected, disconnected, claimed, unclaimed, etc. | [optional] [readonly] 

## Methods

### NewConnectorEventLogAllOf

`func NewConnectorEventLogAllOf(classId string, objectType string, ) *ConnectorEventLogAllOf`

NewConnectorEventLogAllOf instantiates a new ConnectorEventLogAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorEventLogAllOfWithDefaults

`func NewConnectorEventLogAllOfWithDefaults() *ConnectorEventLogAllOf`

NewConnectorEventLogAllOfWithDefaults instantiates a new ConnectorEventLogAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ConnectorEventLogAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ConnectorEventLogAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ConnectorEventLogAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ConnectorEventLogAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ConnectorEventLogAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ConnectorEventLogAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetContents

`func (o *ConnectorEventLogAllOf) GetContents() string`

GetContents returns the Contents field if non-nil, zero value otherwise.

### GetContentsOk

`func (o *ConnectorEventLogAllOf) GetContentsOk() (*string, bool)`

GetContentsOk returns a tuple with the Contents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContents

`func (o *ConnectorEventLogAllOf) SetContents(v string)`

SetContents sets Contents field to given value.

### HasContents

`func (o *ConnectorEventLogAllOf) HasContents() bool`

HasContents returns a boolean if a field has been set.

### GetType

`func (o *ConnectorEventLogAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConnectorEventLogAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConnectorEventLogAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ConnectorEventLogAllOf) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



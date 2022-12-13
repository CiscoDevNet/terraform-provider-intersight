# ConnectorEventLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "connector.EventLog"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "connector.EventLog"]
**Contents** | Pointer to **string** | Change contents of this event. | [optional] [readonly] 
**Type** | Pointer to **string** | Type of the change event, e.g. connected, disconnected, claimed, unclaimed, etc. | [optional] [readonly] 

## Methods

### NewConnectorEventLog

`func NewConnectorEventLog(classId string, objectType string, ) *ConnectorEventLog`

NewConnectorEventLog instantiates a new ConnectorEventLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorEventLogWithDefaults

`func NewConnectorEventLogWithDefaults() *ConnectorEventLog`

NewConnectorEventLogWithDefaults instantiates a new ConnectorEventLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ConnectorEventLog) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ConnectorEventLog) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ConnectorEventLog) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ConnectorEventLog) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ConnectorEventLog) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ConnectorEventLog) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetContents

`func (o *ConnectorEventLog) GetContents() string`

GetContents returns the Contents field if non-nil, zero value otherwise.

### GetContentsOk

`func (o *ConnectorEventLog) GetContentsOk() (*string, bool)`

GetContentsOk returns a tuple with the Contents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContents

`func (o *ConnectorEventLog) SetContents(v string)`

SetContents sets Contents field to given value.

### HasContents

`func (o *ConnectorEventLog) HasContents() bool`

HasContents returns a boolean if a field has been set.

### GetType

`func (o *ConnectorEventLog) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConnectorEventLog) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConnectorEventLog) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ConnectorEventLog) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# ConnectorScopedInventoryAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**NamingProperty** | Pointer to **string** | A property that uniquely identifies the object to be inventoried as a part of the scoped inventory. | [optional] 
**Queries** | Pointer to **interface{}** | Set of queries to identify objects to be inventoried as part of this scoped inventory action. | [optional] 
**Type** | Pointer to **string** | Type of the object for which scoped inventory needs to be run. | [optional] 
**Values** | Pointer to **[]string** |  | [optional] 

## Methods

### NewConnectorScopedInventoryAllOf

`func NewConnectorScopedInventoryAllOf(classId string, objectType string, ) *ConnectorScopedInventoryAllOf`

NewConnectorScopedInventoryAllOf instantiates a new ConnectorScopedInventoryAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorScopedInventoryAllOfWithDefaults

`func NewConnectorScopedInventoryAllOfWithDefaults() *ConnectorScopedInventoryAllOf`

NewConnectorScopedInventoryAllOfWithDefaults instantiates a new ConnectorScopedInventoryAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ConnectorScopedInventoryAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ConnectorScopedInventoryAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ConnectorScopedInventoryAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ConnectorScopedInventoryAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ConnectorScopedInventoryAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ConnectorScopedInventoryAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetNamingProperty

`func (o *ConnectorScopedInventoryAllOf) GetNamingProperty() string`

GetNamingProperty returns the NamingProperty field if non-nil, zero value otherwise.

### GetNamingPropertyOk

`func (o *ConnectorScopedInventoryAllOf) GetNamingPropertyOk() (*string, bool)`

GetNamingPropertyOk returns a tuple with the NamingProperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamingProperty

`func (o *ConnectorScopedInventoryAllOf) SetNamingProperty(v string)`

SetNamingProperty sets NamingProperty field to given value.

### HasNamingProperty

`func (o *ConnectorScopedInventoryAllOf) HasNamingProperty() bool`

HasNamingProperty returns a boolean if a field has been set.

### GetQueries

`func (o *ConnectorScopedInventoryAllOf) GetQueries() interface{}`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *ConnectorScopedInventoryAllOf) GetQueriesOk() (*interface{}, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *ConnectorScopedInventoryAllOf) SetQueries(v interface{})`

SetQueries sets Queries field to given value.

### HasQueries

`func (o *ConnectorScopedInventoryAllOf) HasQueries() bool`

HasQueries returns a boolean if a field has been set.

### SetQueriesNil

`func (o *ConnectorScopedInventoryAllOf) SetQueriesNil(b bool)`

 SetQueriesNil sets the value for Queries to be an explicit nil

### UnsetQueries
`func (o *ConnectorScopedInventoryAllOf) UnsetQueries()`

UnsetQueries ensures that no value is present for Queries, not even an explicit nil
### GetType

`func (o *ConnectorScopedInventoryAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConnectorScopedInventoryAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConnectorScopedInventoryAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ConnectorScopedInventoryAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValues

`func (o *ConnectorScopedInventoryAllOf) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *ConnectorScopedInventoryAllOf) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *ConnectorScopedInventoryAllOf) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *ConnectorScopedInventoryAllOf) HasValues() bool`

HasValues returns a boolean if a field has been set.

### SetValuesNil

`func (o *ConnectorScopedInventoryAllOf) SetValuesNil(b bool)`

 SetValuesNil sets the value for Values to be an explicit nil

### UnsetValues
`func (o *ConnectorScopedInventoryAllOf) UnsetValues()`

UnsetValues ensures that no value is present for Values, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



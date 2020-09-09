# ConnectorScopedInventoryAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NamingProperty** | Pointer to **string** | A property that uniquely identifies the object to be inventoried as a part of the scoped inventory. | [optional] 
**Type** | Pointer to **string** | Type of the object for which scoped inventory needs to be run. | [optional] 
**Values** | Pointer to **[]string** |  | [optional] 

## Methods

### NewConnectorScopedInventoryAllOf

`func NewConnectorScopedInventoryAllOf() *ConnectorScopedInventoryAllOf`

NewConnectorScopedInventoryAllOf instantiates a new ConnectorScopedInventoryAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorScopedInventoryAllOfWithDefaults

`func NewConnectorScopedInventoryAllOfWithDefaults() *ConnectorScopedInventoryAllOf`

NewConnectorScopedInventoryAllOfWithDefaults instantiates a new ConnectorScopedInventoryAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



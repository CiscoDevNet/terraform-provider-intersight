# ConnectorScopedInventory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**NamingProperty** | Pointer to **string** | A property that uniquely identifies the object to be inventoried as a part of the scoped inventory. | [optional] 
**Type** | Pointer to **string** | Type of the object for which scoped inventory needs to be run. | [optional] 
**Values** | Pointer to **[]string** |  | [optional] 

## Methods

### NewConnectorScopedInventory

`func NewConnectorScopedInventory(classId string, objectType string, ) *ConnectorScopedInventory`

NewConnectorScopedInventory instantiates a new ConnectorScopedInventory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorScopedInventoryWithDefaults

`func NewConnectorScopedInventoryWithDefaults() *ConnectorScopedInventory`

NewConnectorScopedInventoryWithDefaults instantiates a new ConnectorScopedInventory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ConnectorScopedInventory) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ConnectorScopedInventory) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ConnectorScopedInventory) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ConnectorScopedInventory) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ConnectorScopedInventory) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ConnectorScopedInventory) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetNamingProperty

`func (o *ConnectorScopedInventory) GetNamingProperty() string`

GetNamingProperty returns the NamingProperty field if non-nil, zero value otherwise.

### GetNamingPropertyOk

`func (o *ConnectorScopedInventory) GetNamingPropertyOk() (*string, bool)`

GetNamingPropertyOk returns a tuple with the NamingProperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamingProperty

`func (o *ConnectorScopedInventory) SetNamingProperty(v string)`

SetNamingProperty sets NamingProperty field to given value.

### HasNamingProperty

`func (o *ConnectorScopedInventory) HasNamingProperty() bool`

HasNamingProperty returns a boolean if a field has been set.

### GetType

`func (o *ConnectorScopedInventory) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConnectorScopedInventory) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConnectorScopedInventory) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ConnectorScopedInventory) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValues

`func (o *ConnectorScopedInventory) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *ConnectorScopedInventory) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *ConnectorScopedInventory) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *ConnectorScopedInventory) HasValues() bool`

HasValues returns a boolean if a field has been set.

### SetValuesNil

`func (o *ConnectorScopedInventory) SetValuesNil(b bool)`

 SetValuesNil sets the value for Values to be an explicit nil

### UnsetValues
`func (o *ConnectorScopedInventory) UnsetValues()`

UnsetValues ensures that no value is present for Values, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



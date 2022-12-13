# HclServerHwCatalogInfoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hcl.ServerHwCatalogInfo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hcl.ServerHwCatalogInfo"]
**Constraints** | Pointer to [**[]HclConstraint**](HclConstraint.md) |  | [optional] 
**DisplayModel** | Pointer to **string** | Display model of the server hardware. In many cases, the model string used in the catalog will be the hardware Cisco PID and the Intersight model string is a more user-friendly string with vendor information in it. This will be the user-friendly modal string to be used in Intersight. | [optional] 
**ExtId** | Pointer to **string** | Identifier of specific tag. | [optional] 
**Model** | Pointer to **string** | Model of the server hardware from the catalog file. | [optional] 
**Type** | Pointer to **string** | Type of specific tag, required to choose the correct datatype while reading the value. | [optional] 
**Value** | Pointer to **interface{}** | Value of specific tag, having type as json. | [optional] 
**ServerHwInfo** | Pointer to [**[]HclServerHwCatalogInfoRelationship**](HclServerHwCatalogInfoRelationship.md) | An array of relationships to hclServerHwCatalogInfo resources. | [optional] [readonly] 

## Methods

### NewHclServerHwCatalogInfoAllOf

`func NewHclServerHwCatalogInfoAllOf(classId string, objectType string, ) *HclServerHwCatalogInfoAllOf`

NewHclServerHwCatalogInfoAllOf instantiates a new HclServerHwCatalogInfoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHclServerHwCatalogInfoAllOfWithDefaults

`func NewHclServerHwCatalogInfoAllOfWithDefaults() *HclServerHwCatalogInfoAllOf`

NewHclServerHwCatalogInfoAllOfWithDefaults instantiates a new HclServerHwCatalogInfoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HclServerHwCatalogInfoAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HclServerHwCatalogInfoAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HclServerHwCatalogInfoAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HclServerHwCatalogInfoAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HclServerHwCatalogInfoAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HclServerHwCatalogInfoAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetConstraints

`func (o *HclServerHwCatalogInfoAllOf) GetConstraints() []HclConstraint`

GetConstraints returns the Constraints field if non-nil, zero value otherwise.

### GetConstraintsOk

`func (o *HclServerHwCatalogInfoAllOf) GetConstraintsOk() (*[]HclConstraint, bool)`

GetConstraintsOk returns a tuple with the Constraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraints

`func (o *HclServerHwCatalogInfoAllOf) SetConstraints(v []HclConstraint)`

SetConstraints sets Constraints field to given value.

### HasConstraints

`func (o *HclServerHwCatalogInfoAllOf) HasConstraints() bool`

HasConstraints returns a boolean if a field has been set.

### SetConstraintsNil

`func (o *HclServerHwCatalogInfoAllOf) SetConstraintsNil(b bool)`

 SetConstraintsNil sets the value for Constraints to be an explicit nil

### UnsetConstraints
`func (o *HclServerHwCatalogInfoAllOf) UnsetConstraints()`

UnsetConstraints ensures that no value is present for Constraints, not even an explicit nil
### GetDisplayModel

`func (o *HclServerHwCatalogInfoAllOf) GetDisplayModel() string`

GetDisplayModel returns the DisplayModel field if non-nil, zero value otherwise.

### GetDisplayModelOk

`func (o *HclServerHwCatalogInfoAllOf) GetDisplayModelOk() (*string, bool)`

GetDisplayModelOk returns a tuple with the DisplayModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayModel

`func (o *HclServerHwCatalogInfoAllOf) SetDisplayModel(v string)`

SetDisplayModel sets DisplayModel field to given value.

### HasDisplayModel

`func (o *HclServerHwCatalogInfoAllOf) HasDisplayModel() bool`

HasDisplayModel returns a boolean if a field has been set.

### GetExtId

`func (o *HclServerHwCatalogInfoAllOf) GetExtId() string`

GetExtId returns the ExtId field if non-nil, zero value otherwise.

### GetExtIdOk

`func (o *HclServerHwCatalogInfoAllOf) GetExtIdOk() (*string, bool)`

GetExtIdOk returns a tuple with the ExtId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtId

`func (o *HclServerHwCatalogInfoAllOf) SetExtId(v string)`

SetExtId sets ExtId field to given value.

### HasExtId

`func (o *HclServerHwCatalogInfoAllOf) HasExtId() bool`

HasExtId returns a boolean if a field has been set.

### GetModel

`func (o *HclServerHwCatalogInfoAllOf) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *HclServerHwCatalogInfoAllOf) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *HclServerHwCatalogInfoAllOf) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *HclServerHwCatalogInfoAllOf) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetType

`func (o *HclServerHwCatalogInfoAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HclServerHwCatalogInfoAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HclServerHwCatalogInfoAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *HclServerHwCatalogInfoAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *HclServerHwCatalogInfoAllOf) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *HclServerHwCatalogInfoAllOf) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *HclServerHwCatalogInfoAllOf) SetValue(v interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *HclServerHwCatalogInfoAllOf) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *HclServerHwCatalogInfoAllOf) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *HclServerHwCatalogInfoAllOf) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetServerHwInfo

`func (o *HclServerHwCatalogInfoAllOf) GetServerHwInfo() []HclServerHwCatalogInfoRelationship`

GetServerHwInfo returns the ServerHwInfo field if non-nil, zero value otherwise.

### GetServerHwInfoOk

`func (o *HclServerHwCatalogInfoAllOf) GetServerHwInfoOk() (*[]HclServerHwCatalogInfoRelationship, bool)`

GetServerHwInfoOk returns a tuple with the ServerHwInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerHwInfo

`func (o *HclServerHwCatalogInfoAllOf) SetServerHwInfo(v []HclServerHwCatalogInfoRelationship)`

SetServerHwInfo sets ServerHwInfo field to given value.

### HasServerHwInfo

`func (o *HclServerHwCatalogInfoAllOf) HasServerHwInfo() bool`

HasServerHwInfo returns a boolean if a field has been set.

### SetServerHwInfoNil

`func (o *HclServerHwCatalogInfoAllOf) SetServerHwInfoNil(b bool)`

 SetServerHwInfoNil sets the value for ServerHwInfo to be an explicit nil

### UnsetServerHwInfo
`func (o *HclServerHwCatalogInfoAllOf) UnsetServerHwInfo()`

UnsetServerHwInfo ensures that no value is present for ServerHwInfo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



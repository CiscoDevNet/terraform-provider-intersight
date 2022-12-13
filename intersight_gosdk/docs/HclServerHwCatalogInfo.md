# HclServerHwCatalogInfo

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

### NewHclServerHwCatalogInfo

`func NewHclServerHwCatalogInfo(classId string, objectType string, ) *HclServerHwCatalogInfo`

NewHclServerHwCatalogInfo instantiates a new HclServerHwCatalogInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHclServerHwCatalogInfoWithDefaults

`func NewHclServerHwCatalogInfoWithDefaults() *HclServerHwCatalogInfo`

NewHclServerHwCatalogInfoWithDefaults instantiates a new HclServerHwCatalogInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HclServerHwCatalogInfo) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HclServerHwCatalogInfo) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HclServerHwCatalogInfo) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HclServerHwCatalogInfo) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HclServerHwCatalogInfo) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HclServerHwCatalogInfo) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetConstraints

`func (o *HclServerHwCatalogInfo) GetConstraints() []HclConstraint`

GetConstraints returns the Constraints field if non-nil, zero value otherwise.

### GetConstraintsOk

`func (o *HclServerHwCatalogInfo) GetConstraintsOk() (*[]HclConstraint, bool)`

GetConstraintsOk returns a tuple with the Constraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraints

`func (o *HclServerHwCatalogInfo) SetConstraints(v []HclConstraint)`

SetConstraints sets Constraints field to given value.

### HasConstraints

`func (o *HclServerHwCatalogInfo) HasConstraints() bool`

HasConstraints returns a boolean if a field has been set.

### SetConstraintsNil

`func (o *HclServerHwCatalogInfo) SetConstraintsNil(b bool)`

 SetConstraintsNil sets the value for Constraints to be an explicit nil

### UnsetConstraints
`func (o *HclServerHwCatalogInfo) UnsetConstraints()`

UnsetConstraints ensures that no value is present for Constraints, not even an explicit nil
### GetDisplayModel

`func (o *HclServerHwCatalogInfo) GetDisplayModel() string`

GetDisplayModel returns the DisplayModel field if non-nil, zero value otherwise.

### GetDisplayModelOk

`func (o *HclServerHwCatalogInfo) GetDisplayModelOk() (*string, bool)`

GetDisplayModelOk returns a tuple with the DisplayModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayModel

`func (o *HclServerHwCatalogInfo) SetDisplayModel(v string)`

SetDisplayModel sets DisplayModel field to given value.

### HasDisplayModel

`func (o *HclServerHwCatalogInfo) HasDisplayModel() bool`

HasDisplayModel returns a boolean if a field has been set.

### GetExtId

`func (o *HclServerHwCatalogInfo) GetExtId() string`

GetExtId returns the ExtId field if non-nil, zero value otherwise.

### GetExtIdOk

`func (o *HclServerHwCatalogInfo) GetExtIdOk() (*string, bool)`

GetExtIdOk returns a tuple with the ExtId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtId

`func (o *HclServerHwCatalogInfo) SetExtId(v string)`

SetExtId sets ExtId field to given value.

### HasExtId

`func (o *HclServerHwCatalogInfo) HasExtId() bool`

HasExtId returns a boolean if a field has been set.

### GetModel

`func (o *HclServerHwCatalogInfo) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *HclServerHwCatalogInfo) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *HclServerHwCatalogInfo) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *HclServerHwCatalogInfo) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetType

`func (o *HclServerHwCatalogInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HclServerHwCatalogInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HclServerHwCatalogInfo) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *HclServerHwCatalogInfo) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *HclServerHwCatalogInfo) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *HclServerHwCatalogInfo) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *HclServerHwCatalogInfo) SetValue(v interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *HclServerHwCatalogInfo) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *HclServerHwCatalogInfo) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *HclServerHwCatalogInfo) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetServerHwInfo

`func (o *HclServerHwCatalogInfo) GetServerHwInfo() []HclServerHwCatalogInfoRelationship`

GetServerHwInfo returns the ServerHwInfo field if non-nil, zero value otherwise.

### GetServerHwInfoOk

`func (o *HclServerHwCatalogInfo) GetServerHwInfoOk() (*[]HclServerHwCatalogInfoRelationship, bool)`

GetServerHwInfoOk returns a tuple with the ServerHwInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerHwInfo

`func (o *HclServerHwCatalogInfo) SetServerHwInfo(v []HclServerHwCatalogInfoRelationship)`

SetServerHwInfo sets ServerHwInfo field to given value.

### HasServerHwInfo

`func (o *HclServerHwCatalogInfo) HasServerHwInfo() bool`

HasServerHwInfo returns a boolean if a field has been set.

### SetServerHwInfoNil

`func (o *HclServerHwCatalogInfo) SetServerHwInfoNil(b bool)`

 SetServerHwInfoNil sets the value for ServerHwInfo to be an explicit nil

### UnsetServerHwInfo
`func (o *HclServerHwCatalogInfo) UnsetServerHwInfo()`

UnsetServerHwInfo ensures that no value is present for ServerHwInfo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



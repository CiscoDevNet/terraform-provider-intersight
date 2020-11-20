# InventoryInventoryMo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "inventory.InventoryMo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "inventory.InventoryMo"]
**MoDn** | Pointer to **string** | The UCS DN of the MO for which the latest inventory to be fetched. If this property is empty and moId property has the Moid of the MO to be updated, the Moid will be used. If this property is empty and moId is also empty, all the MOs of the given moType will be updated. | [optional] 
**MoId** | Pointer to **string** | The MO id of an MO for which the latest inventory to be fetched. If this property is empty and moDn property has the UCS DN of the MO to be updated, the DN will be used. If this property is empty and moDn is also empty, all the MOs of the given moType will be updated. | [optional] 
**MoType** | Pointer to **string** | The type of the MO for which the latest inventory to be fetched. | [optional] 

## Methods

### NewInventoryInventoryMo

`func NewInventoryInventoryMo(classId string, objectType string, ) *InventoryInventoryMo`

NewInventoryInventoryMo instantiates a new InventoryInventoryMo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInventoryInventoryMoWithDefaults

`func NewInventoryInventoryMoWithDefaults() *InventoryInventoryMo`

NewInventoryInventoryMoWithDefaults instantiates a new InventoryInventoryMo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *InventoryInventoryMo) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *InventoryInventoryMo) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *InventoryInventoryMo) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *InventoryInventoryMo) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *InventoryInventoryMo) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *InventoryInventoryMo) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetMoDn

`func (o *InventoryInventoryMo) GetMoDn() string`

GetMoDn returns the MoDn field if non-nil, zero value otherwise.

### GetMoDnOk

`func (o *InventoryInventoryMo) GetMoDnOk() (*string, bool)`

GetMoDnOk returns a tuple with the MoDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoDn

`func (o *InventoryInventoryMo) SetMoDn(v string)`

SetMoDn sets MoDn field to given value.

### HasMoDn

`func (o *InventoryInventoryMo) HasMoDn() bool`

HasMoDn returns a boolean if a field has been set.

### GetMoId

`func (o *InventoryInventoryMo) GetMoId() string`

GetMoId returns the MoId field if non-nil, zero value otherwise.

### GetMoIdOk

`func (o *InventoryInventoryMo) GetMoIdOk() (*string, bool)`

GetMoIdOk returns a tuple with the MoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoId

`func (o *InventoryInventoryMo) SetMoId(v string)`

SetMoId sets MoId field to given value.

### HasMoId

`func (o *InventoryInventoryMo) HasMoId() bool`

HasMoId returns a boolean if a field has been set.

### GetMoType

`func (o *InventoryInventoryMo) GetMoType() string`

GetMoType returns the MoType field if non-nil, zero value otherwise.

### GetMoTypeOk

`func (o *InventoryInventoryMo) GetMoTypeOk() (*string, bool)`

GetMoTypeOk returns a tuple with the MoType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoType

`func (o *InventoryInventoryMo) SetMoType(v string)`

SetMoType sets MoType field to given value.

### HasMoType

`func (o *InventoryInventoryMo) HasMoType() bool`

HasMoType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# InventoryBaseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**DeviceMoId** | Pointer to **string** | The database identifier of the registered device of an object. | [optional] [readonly] 
**Dn** | Pointer to **string** | The Distinguished Name unambiguously identifies an object in the system. | [optional] [readonly] 
**Rn** | Pointer to **string** | The Relative Name uniquely identifies an object within a given context. | [optional] [readonly] 

## Methods

### NewInventoryBaseAllOf

`func NewInventoryBaseAllOf(classId string, objectType string, ) *InventoryBaseAllOf`

NewInventoryBaseAllOf instantiates a new InventoryBaseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInventoryBaseAllOfWithDefaults

`func NewInventoryBaseAllOfWithDefaults() *InventoryBaseAllOf`

NewInventoryBaseAllOfWithDefaults instantiates a new InventoryBaseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *InventoryBaseAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *InventoryBaseAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *InventoryBaseAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *InventoryBaseAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *InventoryBaseAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *InventoryBaseAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDeviceMoId

`func (o *InventoryBaseAllOf) GetDeviceMoId() string`

GetDeviceMoId returns the DeviceMoId field if non-nil, zero value otherwise.

### GetDeviceMoIdOk

`func (o *InventoryBaseAllOf) GetDeviceMoIdOk() (*string, bool)`

GetDeviceMoIdOk returns a tuple with the DeviceMoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMoId

`func (o *InventoryBaseAllOf) SetDeviceMoId(v string)`

SetDeviceMoId sets DeviceMoId field to given value.

### HasDeviceMoId

`func (o *InventoryBaseAllOf) HasDeviceMoId() bool`

HasDeviceMoId returns a boolean if a field has been set.

### GetDn

`func (o *InventoryBaseAllOf) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *InventoryBaseAllOf) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *InventoryBaseAllOf) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *InventoryBaseAllOf) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetRn

`func (o *InventoryBaseAllOf) GetRn() string`

GetRn returns the Rn field if non-nil, zero value otherwise.

### GetRnOk

`func (o *InventoryBaseAllOf) GetRnOk() (*string, bool)`

GetRnOk returns a tuple with the Rn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRn

`func (o *InventoryBaseAllOf) SetRn(v string)`

SetRn sets Rn field to given value.

### HasRn

`func (o *InventoryBaseAllOf) HasRn() bool`

HasRn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



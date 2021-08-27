# InventoryDnMoBinding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "inventory.DnMoBinding"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "inventory.DnMoBinding"]
**Dn** | Pointer to **string** | The Distinguished Name for this object, used to uniquely identify this object. | [optional] [readonly] 
**TargetMoId** | Pointer to **string** | The MO ID of the target MO for this particular Distinguished Name (dn). | [optional] [readonly] 
**TargetMoType** | Pointer to **string** | The type of the target MO for this particular Distinguished Name (dn). | [optional] [readonly] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewInventoryDnMoBinding

`func NewInventoryDnMoBinding(classId string, objectType string, ) *InventoryDnMoBinding`

NewInventoryDnMoBinding instantiates a new InventoryDnMoBinding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInventoryDnMoBindingWithDefaults

`func NewInventoryDnMoBindingWithDefaults() *InventoryDnMoBinding`

NewInventoryDnMoBindingWithDefaults instantiates a new InventoryDnMoBinding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *InventoryDnMoBinding) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *InventoryDnMoBinding) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *InventoryDnMoBinding) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *InventoryDnMoBinding) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *InventoryDnMoBinding) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *InventoryDnMoBinding) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDn

`func (o *InventoryDnMoBinding) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *InventoryDnMoBinding) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *InventoryDnMoBinding) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *InventoryDnMoBinding) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetTargetMoId

`func (o *InventoryDnMoBinding) GetTargetMoId() string`

GetTargetMoId returns the TargetMoId field if non-nil, zero value otherwise.

### GetTargetMoIdOk

`func (o *InventoryDnMoBinding) GetTargetMoIdOk() (*string, bool)`

GetTargetMoIdOk returns a tuple with the TargetMoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetMoId

`func (o *InventoryDnMoBinding) SetTargetMoId(v string)`

SetTargetMoId sets TargetMoId field to given value.

### HasTargetMoId

`func (o *InventoryDnMoBinding) HasTargetMoId() bool`

HasTargetMoId returns a boolean if a field has been set.

### GetTargetMoType

`func (o *InventoryDnMoBinding) GetTargetMoType() string`

GetTargetMoType returns the TargetMoType field if non-nil, zero value otherwise.

### GetTargetMoTypeOk

`func (o *InventoryDnMoBinding) GetTargetMoTypeOk() (*string, bool)`

GetTargetMoTypeOk returns a tuple with the TargetMoType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetMoType

`func (o *InventoryDnMoBinding) SetTargetMoType(v string)`

SetTargetMoType sets TargetMoType field to given value.

### HasTargetMoType

`func (o *InventoryDnMoBinding) HasTargetMoType() bool`

HasTargetMoType returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *InventoryDnMoBinding) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *InventoryDnMoBinding) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *InventoryDnMoBinding) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *InventoryDnMoBinding) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



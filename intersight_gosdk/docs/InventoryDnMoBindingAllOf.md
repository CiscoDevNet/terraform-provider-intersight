# InventoryDnMoBindingAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dn** | Pointer to **string** | The Distinguished Name for this object, used to uniquely identify this object. | [optional] [readonly] 
**TargetMoId** | Pointer to **string** | The MO ID of the target MO for this particular Distinguished Name (dn). | [optional] [readonly] 
**TargetMoType** | Pointer to **string** | The type of the target MO for this particular Distinguished Name (dn). | [optional] [readonly] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewInventoryDnMoBindingAllOf

`func NewInventoryDnMoBindingAllOf() *InventoryDnMoBindingAllOf`

NewInventoryDnMoBindingAllOf instantiates a new InventoryDnMoBindingAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInventoryDnMoBindingAllOfWithDefaults

`func NewInventoryDnMoBindingAllOfWithDefaults() *InventoryDnMoBindingAllOf`

NewInventoryDnMoBindingAllOfWithDefaults instantiates a new InventoryDnMoBindingAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDn

`func (o *InventoryDnMoBindingAllOf) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *InventoryDnMoBindingAllOf) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *InventoryDnMoBindingAllOf) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *InventoryDnMoBindingAllOf) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetTargetMoId

`func (o *InventoryDnMoBindingAllOf) GetTargetMoId() string`

GetTargetMoId returns the TargetMoId field if non-nil, zero value otherwise.

### GetTargetMoIdOk

`func (o *InventoryDnMoBindingAllOf) GetTargetMoIdOk() (*string, bool)`

GetTargetMoIdOk returns a tuple with the TargetMoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetMoId

`func (o *InventoryDnMoBindingAllOf) SetTargetMoId(v string)`

SetTargetMoId sets TargetMoId field to given value.

### HasTargetMoId

`func (o *InventoryDnMoBindingAllOf) HasTargetMoId() bool`

HasTargetMoId returns a boolean if a field has been set.

### GetTargetMoType

`func (o *InventoryDnMoBindingAllOf) GetTargetMoType() string`

GetTargetMoType returns the TargetMoType field if non-nil, zero value otherwise.

### GetTargetMoTypeOk

`func (o *InventoryDnMoBindingAllOf) GetTargetMoTypeOk() (*string, bool)`

GetTargetMoTypeOk returns a tuple with the TargetMoType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetMoType

`func (o *InventoryDnMoBindingAllOf) SetTargetMoType(v string)`

SetTargetMoType sets TargetMoType field to given value.

### HasTargetMoType

`func (o *InventoryDnMoBindingAllOf) HasTargetMoType() bool`

HasTargetMoType returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *InventoryDnMoBindingAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *InventoryDnMoBindingAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *InventoryDnMoBindingAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *InventoryDnMoBindingAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



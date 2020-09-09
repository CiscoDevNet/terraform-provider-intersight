# InventoryRequestAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mos** | Pointer to [**[]InventoryInventoryMo**](inventory.InventoryMo.md) |  | [optional] 
**Device** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewInventoryRequestAllOf

`func NewInventoryRequestAllOf() *InventoryRequestAllOf`

NewInventoryRequestAllOf instantiates a new InventoryRequestAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInventoryRequestAllOfWithDefaults

`func NewInventoryRequestAllOfWithDefaults() *InventoryRequestAllOf`

NewInventoryRequestAllOfWithDefaults instantiates a new InventoryRequestAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMos

`func (o *InventoryRequestAllOf) GetMos() []InventoryInventoryMo`

GetMos returns the Mos field if non-nil, zero value otherwise.

### GetMosOk

`func (o *InventoryRequestAllOf) GetMosOk() (*[]InventoryInventoryMo, bool)`

GetMosOk returns a tuple with the Mos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMos

`func (o *InventoryRequestAllOf) SetMos(v []InventoryInventoryMo)`

SetMos sets Mos field to given value.

### HasMos

`func (o *InventoryRequestAllOf) HasMos() bool`

HasMos returns a boolean if a field has been set.

### GetDevice

`func (o *InventoryRequestAllOf) GetDevice() AssetDeviceRegistrationRelationship`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *InventoryRequestAllOf) GetDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *InventoryRequestAllOf) SetDevice(v AssetDeviceRegistrationRelationship)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *InventoryRequestAllOf) HasDevice() bool`

HasDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



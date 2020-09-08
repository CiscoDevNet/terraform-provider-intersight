# InventoryGenericInventoryAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | Key of inventory data for Generic Inventory data set. | [optional] [readonly] 
**Type** | Pointer to **string** | Type of inventory data for Generic Inventory data set. | [optional] [readonly] 
**Value** | Pointer to **string** | Value of inventory data for Generic Inventory data set. | [optional] [readonly] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](inventory.DeviceInfo.Relationship.md) |  | [optional] 
**InventoryGenericInventoryHolder** | Pointer to [**InventoryGenericInventoryHolderRelationship**](inventory.GenericInventoryHolder.Relationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewInventoryGenericInventoryAllOf

`func NewInventoryGenericInventoryAllOf() *InventoryGenericInventoryAllOf`

NewInventoryGenericInventoryAllOf instantiates a new InventoryGenericInventoryAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInventoryGenericInventoryAllOfWithDefaults

`func NewInventoryGenericInventoryAllOfWithDefaults() *InventoryGenericInventoryAllOf`

NewInventoryGenericInventoryAllOfWithDefaults instantiates a new InventoryGenericInventoryAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *InventoryGenericInventoryAllOf) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *InventoryGenericInventoryAllOf) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *InventoryGenericInventoryAllOf) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *InventoryGenericInventoryAllOf) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetType

`func (o *InventoryGenericInventoryAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InventoryGenericInventoryAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InventoryGenericInventoryAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InventoryGenericInventoryAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *InventoryGenericInventoryAllOf) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *InventoryGenericInventoryAllOf) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *InventoryGenericInventoryAllOf) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *InventoryGenericInventoryAllOf) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *InventoryGenericInventoryAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *InventoryGenericInventoryAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *InventoryGenericInventoryAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *InventoryGenericInventoryAllOf) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetInventoryGenericInventoryHolder

`func (o *InventoryGenericInventoryAllOf) GetInventoryGenericInventoryHolder() InventoryGenericInventoryHolderRelationship`

GetInventoryGenericInventoryHolder returns the InventoryGenericInventoryHolder field if non-nil, zero value otherwise.

### GetInventoryGenericInventoryHolderOk

`func (o *InventoryGenericInventoryAllOf) GetInventoryGenericInventoryHolderOk() (*InventoryGenericInventoryHolderRelationship, bool)`

GetInventoryGenericInventoryHolderOk returns a tuple with the InventoryGenericInventoryHolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryGenericInventoryHolder

`func (o *InventoryGenericInventoryAllOf) SetInventoryGenericInventoryHolder(v InventoryGenericInventoryHolderRelationship)`

SetInventoryGenericInventoryHolder sets InventoryGenericInventoryHolder field to given value.

### HasInventoryGenericInventoryHolder

`func (o *InventoryGenericInventoryAllOf) HasInventoryGenericInventoryHolder() bool`

HasInventoryGenericInventoryHolder returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *InventoryGenericInventoryAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *InventoryGenericInventoryAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *InventoryGenericInventoryAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *InventoryGenericInventoryAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



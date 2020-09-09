# InventoryGenericInventory

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

### NewInventoryGenericInventory

`func NewInventoryGenericInventory() *InventoryGenericInventory`

NewInventoryGenericInventory instantiates a new InventoryGenericInventory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInventoryGenericInventoryWithDefaults

`func NewInventoryGenericInventoryWithDefaults() *InventoryGenericInventory`

NewInventoryGenericInventoryWithDefaults instantiates a new InventoryGenericInventory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *InventoryGenericInventory) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *InventoryGenericInventory) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *InventoryGenericInventory) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *InventoryGenericInventory) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetType

`func (o *InventoryGenericInventory) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InventoryGenericInventory) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InventoryGenericInventory) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InventoryGenericInventory) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *InventoryGenericInventory) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *InventoryGenericInventory) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *InventoryGenericInventory) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *InventoryGenericInventory) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *InventoryGenericInventory) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *InventoryGenericInventory) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *InventoryGenericInventory) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *InventoryGenericInventory) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetInventoryGenericInventoryHolder

`func (o *InventoryGenericInventory) GetInventoryGenericInventoryHolder() InventoryGenericInventoryHolderRelationship`

GetInventoryGenericInventoryHolder returns the InventoryGenericInventoryHolder field if non-nil, zero value otherwise.

### GetInventoryGenericInventoryHolderOk

`func (o *InventoryGenericInventory) GetInventoryGenericInventoryHolderOk() (*InventoryGenericInventoryHolderRelationship, bool)`

GetInventoryGenericInventoryHolderOk returns a tuple with the InventoryGenericInventoryHolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryGenericInventoryHolder

`func (o *InventoryGenericInventory) SetInventoryGenericInventoryHolder(v InventoryGenericInventoryHolderRelationship)`

SetInventoryGenericInventoryHolder sets InventoryGenericInventoryHolder field to given value.

### HasInventoryGenericInventoryHolder

`func (o *InventoryGenericInventory) HasInventoryGenericInventoryHolder() bool`

HasInventoryGenericInventoryHolder returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *InventoryGenericInventory) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *InventoryGenericInventory) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *InventoryGenericInventory) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *InventoryGenericInventory) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



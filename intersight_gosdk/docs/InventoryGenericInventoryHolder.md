# InventoryGenericInventoryHolder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Endpoint** | Pointer to **string** | The endpoint represented by this holder. | [optional] [readonly] 
**ComputeBlade** | Pointer to [**ComputeBladeRelationship**](compute.Blade.Relationship.md) |  | [optional] 
**ComputeRackUnit** | Pointer to [**ComputeRackUnitRelationship**](compute.RackUnit.Relationship.md) |  | [optional] 
**GenericInventory** | Pointer to [**[]InventoryGenericInventoryRelationship**](inventory.GenericInventory.Relationship.md) | An array of relationships to inventoryGenericInventory resources. | [optional] [readonly] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](inventory.DeviceInfo.Relationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewInventoryGenericInventoryHolder

`func NewInventoryGenericInventoryHolder() *InventoryGenericInventoryHolder`

NewInventoryGenericInventoryHolder instantiates a new InventoryGenericInventoryHolder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInventoryGenericInventoryHolderWithDefaults

`func NewInventoryGenericInventoryHolderWithDefaults() *InventoryGenericInventoryHolder`

NewInventoryGenericInventoryHolderWithDefaults instantiates a new InventoryGenericInventoryHolder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpoint

`func (o *InventoryGenericInventoryHolder) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *InventoryGenericInventoryHolder) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *InventoryGenericInventoryHolder) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *InventoryGenericInventoryHolder) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetComputeBlade

`func (o *InventoryGenericInventoryHolder) GetComputeBlade() ComputeBladeRelationship`

GetComputeBlade returns the ComputeBlade field if non-nil, zero value otherwise.

### GetComputeBladeOk

`func (o *InventoryGenericInventoryHolder) GetComputeBladeOk() (*ComputeBladeRelationship, bool)`

GetComputeBladeOk returns a tuple with the ComputeBlade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeBlade

`func (o *InventoryGenericInventoryHolder) SetComputeBlade(v ComputeBladeRelationship)`

SetComputeBlade sets ComputeBlade field to given value.

### HasComputeBlade

`func (o *InventoryGenericInventoryHolder) HasComputeBlade() bool`

HasComputeBlade returns a boolean if a field has been set.

### GetComputeRackUnit

`func (o *InventoryGenericInventoryHolder) GetComputeRackUnit() ComputeRackUnitRelationship`

GetComputeRackUnit returns the ComputeRackUnit field if non-nil, zero value otherwise.

### GetComputeRackUnitOk

`func (o *InventoryGenericInventoryHolder) GetComputeRackUnitOk() (*ComputeRackUnitRelationship, bool)`

GetComputeRackUnitOk returns a tuple with the ComputeRackUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRackUnit

`func (o *InventoryGenericInventoryHolder) SetComputeRackUnit(v ComputeRackUnitRelationship)`

SetComputeRackUnit sets ComputeRackUnit field to given value.

### HasComputeRackUnit

`func (o *InventoryGenericInventoryHolder) HasComputeRackUnit() bool`

HasComputeRackUnit returns a boolean if a field has been set.

### GetGenericInventory

`func (o *InventoryGenericInventoryHolder) GetGenericInventory() []InventoryGenericInventoryRelationship`

GetGenericInventory returns the GenericInventory field if non-nil, zero value otherwise.

### GetGenericInventoryOk

`func (o *InventoryGenericInventoryHolder) GetGenericInventoryOk() (*[]InventoryGenericInventoryRelationship, bool)`

GetGenericInventoryOk returns a tuple with the GenericInventory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenericInventory

`func (o *InventoryGenericInventoryHolder) SetGenericInventory(v []InventoryGenericInventoryRelationship)`

SetGenericInventory sets GenericInventory field to given value.

### HasGenericInventory

`func (o *InventoryGenericInventoryHolder) HasGenericInventory() bool`

HasGenericInventory returns a boolean if a field has been set.

### SetGenericInventoryNil

`func (o *InventoryGenericInventoryHolder) SetGenericInventoryNil(b bool)`

 SetGenericInventoryNil sets the value for GenericInventory to be an explicit nil

### UnsetGenericInventory
`func (o *InventoryGenericInventoryHolder) UnsetGenericInventory()`

UnsetGenericInventory ensures that no value is present for GenericInventory, not even an explicit nil
### GetInventoryDeviceInfo

`func (o *InventoryGenericInventoryHolder) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *InventoryGenericInventoryHolder) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *InventoryGenericInventoryHolder) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *InventoryGenericInventoryHolder) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *InventoryGenericInventoryHolder) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *InventoryGenericInventoryHolder) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *InventoryGenericInventoryHolder) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *InventoryGenericInventoryHolder) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



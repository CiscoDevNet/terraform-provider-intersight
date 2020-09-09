# BootDeviceBootModeAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfiguredBootMode** | Pointer to **string** | The user desired BIOS boot mode as configured in the boot policy. | [optional] 
**ComputeRackUnit** | Pointer to [**ComputeRackUnitRelationship**](compute.RackUnit.Relationship.md) |  | [optional] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](inventory.DeviceInfo.Relationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewBootDeviceBootModeAllOf

`func NewBootDeviceBootModeAllOf() *BootDeviceBootModeAllOf`

NewBootDeviceBootModeAllOf instantiates a new BootDeviceBootModeAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBootDeviceBootModeAllOfWithDefaults

`func NewBootDeviceBootModeAllOfWithDefaults() *BootDeviceBootModeAllOf`

NewBootDeviceBootModeAllOfWithDefaults instantiates a new BootDeviceBootModeAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfiguredBootMode

`func (o *BootDeviceBootModeAllOf) GetConfiguredBootMode() string`

GetConfiguredBootMode returns the ConfiguredBootMode field if non-nil, zero value otherwise.

### GetConfiguredBootModeOk

`func (o *BootDeviceBootModeAllOf) GetConfiguredBootModeOk() (*string, bool)`

GetConfiguredBootModeOk returns a tuple with the ConfiguredBootMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguredBootMode

`func (o *BootDeviceBootModeAllOf) SetConfiguredBootMode(v string)`

SetConfiguredBootMode sets ConfiguredBootMode field to given value.

### HasConfiguredBootMode

`func (o *BootDeviceBootModeAllOf) HasConfiguredBootMode() bool`

HasConfiguredBootMode returns a boolean if a field has been set.

### GetComputeRackUnit

`func (o *BootDeviceBootModeAllOf) GetComputeRackUnit() ComputeRackUnitRelationship`

GetComputeRackUnit returns the ComputeRackUnit field if non-nil, zero value otherwise.

### GetComputeRackUnitOk

`func (o *BootDeviceBootModeAllOf) GetComputeRackUnitOk() (*ComputeRackUnitRelationship, bool)`

GetComputeRackUnitOk returns a tuple with the ComputeRackUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRackUnit

`func (o *BootDeviceBootModeAllOf) SetComputeRackUnit(v ComputeRackUnitRelationship)`

SetComputeRackUnit sets ComputeRackUnit field to given value.

### HasComputeRackUnit

`func (o *BootDeviceBootModeAllOf) HasComputeRackUnit() bool`

HasComputeRackUnit returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *BootDeviceBootModeAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *BootDeviceBootModeAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *BootDeviceBootModeAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *BootDeviceBootModeAllOf) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *BootDeviceBootModeAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *BootDeviceBootModeAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *BootDeviceBootModeAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *BootDeviceBootModeAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



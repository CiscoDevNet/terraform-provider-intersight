# EquipmentRackEnclosureSlot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RackId** | Pointer to **int64** | Server ID which is part of Rack Enclosure Slot. | [optional] [readonly] 
**RackUnitDn** | Pointer to **string** | Server DN which is part of Rack Enclosure Slot. | [optional] [readonly] 
**EquipmentRackEnclosure** | Pointer to [**EquipmentRackEnclosureRelationship**](equipment.RackEnclosure.Relationship.md) |  | [optional] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](inventory.DeviceInfo.Relationship.md) |  | [optional] 
**RackUnit** | Pointer to [**ComputeRackUnitRelationship**](compute.RackUnit.Relationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewEquipmentRackEnclosureSlot

`func NewEquipmentRackEnclosureSlot() *EquipmentRackEnclosureSlot`

NewEquipmentRackEnclosureSlot instantiates a new EquipmentRackEnclosureSlot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquipmentRackEnclosureSlotWithDefaults

`func NewEquipmentRackEnclosureSlotWithDefaults() *EquipmentRackEnclosureSlot`

NewEquipmentRackEnclosureSlotWithDefaults instantiates a new EquipmentRackEnclosureSlot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRackId

`func (o *EquipmentRackEnclosureSlot) GetRackId() int64`

GetRackId returns the RackId field if non-nil, zero value otherwise.

### GetRackIdOk

`func (o *EquipmentRackEnclosureSlot) GetRackIdOk() (*int64, bool)`

GetRackIdOk returns a tuple with the RackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRackId

`func (o *EquipmentRackEnclosureSlot) SetRackId(v int64)`

SetRackId sets RackId field to given value.

### HasRackId

`func (o *EquipmentRackEnclosureSlot) HasRackId() bool`

HasRackId returns a boolean if a field has been set.

### GetRackUnitDn

`func (o *EquipmentRackEnclosureSlot) GetRackUnitDn() string`

GetRackUnitDn returns the RackUnitDn field if non-nil, zero value otherwise.

### GetRackUnitDnOk

`func (o *EquipmentRackEnclosureSlot) GetRackUnitDnOk() (*string, bool)`

GetRackUnitDnOk returns a tuple with the RackUnitDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRackUnitDn

`func (o *EquipmentRackEnclosureSlot) SetRackUnitDn(v string)`

SetRackUnitDn sets RackUnitDn field to given value.

### HasRackUnitDn

`func (o *EquipmentRackEnclosureSlot) HasRackUnitDn() bool`

HasRackUnitDn returns a boolean if a field has been set.

### GetEquipmentRackEnclosure

`func (o *EquipmentRackEnclosureSlot) GetEquipmentRackEnclosure() EquipmentRackEnclosureRelationship`

GetEquipmentRackEnclosure returns the EquipmentRackEnclosure field if non-nil, zero value otherwise.

### GetEquipmentRackEnclosureOk

`func (o *EquipmentRackEnclosureSlot) GetEquipmentRackEnclosureOk() (*EquipmentRackEnclosureRelationship, bool)`

GetEquipmentRackEnclosureOk returns a tuple with the EquipmentRackEnclosure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentRackEnclosure

`func (o *EquipmentRackEnclosureSlot) SetEquipmentRackEnclosure(v EquipmentRackEnclosureRelationship)`

SetEquipmentRackEnclosure sets EquipmentRackEnclosure field to given value.

### HasEquipmentRackEnclosure

`func (o *EquipmentRackEnclosureSlot) HasEquipmentRackEnclosure() bool`

HasEquipmentRackEnclosure returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *EquipmentRackEnclosureSlot) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *EquipmentRackEnclosureSlot) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *EquipmentRackEnclosureSlot) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *EquipmentRackEnclosureSlot) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetRackUnit

`func (o *EquipmentRackEnclosureSlot) GetRackUnit() ComputeRackUnitRelationship`

GetRackUnit returns the RackUnit field if non-nil, zero value otherwise.

### GetRackUnitOk

`func (o *EquipmentRackEnclosureSlot) GetRackUnitOk() (*ComputeRackUnitRelationship, bool)`

GetRackUnitOk returns a tuple with the RackUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRackUnit

`func (o *EquipmentRackEnclosureSlot) SetRackUnit(v ComputeRackUnitRelationship)`

SetRackUnit sets RackUnit field to given value.

### HasRackUnit

`func (o *EquipmentRackEnclosureSlot) HasRackUnit() bool`

HasRackUnit returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *EquipmentRackEnclosureSlot) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *EquipmentRackEnclosureSlot) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *EquipmentRackEnclosureSlot) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *EquipmentRackEnclosureSlot) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



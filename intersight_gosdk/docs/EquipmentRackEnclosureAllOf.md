# EquipmentRackEnclosureAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnclosureId** | Pointer to **int64** | This represents the Enclosure Identifier for Rack servers. | [optional] [readonly] 
**Fanmodules** | Pointer to [**[]EquipmentFanModuleRelationship**](equipment.FanModule.Relationship.md) | An array of relationships to equipmentFanModule resources. | [optional] [readonly] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](inventory.DeviceInfo.Relationship.md) |  | [optional] 
**Psus** | Pointer to [**[]EquipmentPsuRelationship**](equipment.Psu.Relationship.md) | An array of relationships to equipmentPsu resources. | [optional] [readonly] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 
**Slots** | Pointer to [**[]EquipmentRackEnclosureSlotRelationship**](equipment.RackEnclosureSlot.Relationship.md) | An array of relationships to equipmentRackEnclosureSlot resources. | [optional] [readonly] 

## Methods

### NewEquipmentRackEnclosureAllOf

`func NewEquipmentRackEnclosureAllOf() *EquipmentRackEnclosureAllOf`

NewEquipmentRackEnclosureAllOf instantiates a new EquipmentRackEnclosureAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquipmentRackEnclosureAllOfWithDefaults

`func NewEquipmentRackEnclosureAllOfWithDefaults() *EquipmentRackEnclosureAllOf`

NewEquipmentRackEnclosureAllOfWithDefaults instantiates a new EquipmentRackEnclosureAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnclosureId

`func (o *EquipmentRackEnclosureAllOf) GetEnclosureId() int64`

GetEnclosureId returns the EnclosureId field if non-nil, zero value otherwise.

### GetEnclosureIdOk

`func (o *EquipmentRackEnclosureAllOf) GetEnclosureIdOk() (*int64, bool)`

GetEnclosureIdOk returns a tuple with the EnclosureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnclosureId

`func (o *EquipmentRackEnclosureAllOf) SetEnclosureId(v int64)`

SetEnclosureId sets EnclosureId field to given value.

### HasEnclosureId

`func (o *EquipmentRackEnclosureAllOf) HasEnclosureId() bool`

HasEnclosureId returns a boolean if a field has been set.

### GetFanmodules

`func (o *EquipmentRackEnclosureAllOf) GetFanmodules() []EquipmentFanModuleRelationship`

GetFanmodules returns the Fanmodules field if non-nil, zero value otherwise.

### GetFanmodulesOk

`func (o *EquipmentRackEnclosureAllOf) GetFanmodulesOk() (*[]EquipmentFanModuleRelationship, bool)`

GetFanmodulesOk returns a tuple with the Fanmodules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFanmodules

`func (o *EquipmentRackEnclosureAllOf) SetFanmodules(v []EquipmentFanModuleRelationship)`

SetFanmodules sets Fanmodules field to given value.

### HasFanmodules

`func (o *EquipmentRackEnclosureAllOf) HasFanmodules() bool`

HasFanmodules returns a boolean if a field has been set.

### SetFanmodulesNil

`func (o *EquipmentRackEnclosureAllOf) SetFanmodulesNil(b bool)`

 SetFanmodulesNil sets the value for Fanmodules to be an explicit nil

### UnsetFanmodules
`func (o *EquipmentRackEnclosureAllOf) UnsetFanmodules()`

UnsetFanmodules ensures that no value is present for Fanmodules, not even an explicit nil
### GetInventoryDeviceInfo

`func (o *EquipmentRackEnclosureAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *EquipmentRackEnclosureAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *EquipmentRackEnclosureAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *EquipmentRackEnclosureAllOf) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetPsus

`func (o *EquipmentRackEnclosureAllOf) GetPsus() []EquipmentPsuRelationship`

GetPsus returns the Psus field if non-nil, zero value otherwise.

### GetPsusOk

`func (o *EquipmentRackEnclosureAllOf) GetPsusOk() (*[]EquipmentPsuRelationship, bool)`

GetPsusOk returns a tuple with the Psus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPsus

`func (o *EquipmentRackEnclosureAllOf) SetPsus(v []EquipmentPsuRelationship)`

SetPsus sets Psus field to given value.

### HasPsus

`func (o *EquipmentRackEnclosureAllOf) HasPsus() bool`

HasPsus returns a boolean if a field has been set.

### SetPsusNil

`func (o *EquipmentRackEnclosureAllOf) SetPsusNil(b bool)`

 SetPsusNil sets the value for Psus to be an explicit nil

### UnsetPsus
`func (o *EquipmentRackEnclosureAllOf) UnsetPsus()`

UnsetPsus ensures that no value is present for Psus, not even an explicit nil
### GetRegisteredDevice

`func (o *EquipmentRackEnclosureAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *EquipmentRackEnclosureAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *EquipmentRackEnclosureAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *EquipmentRackEnclosureAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### GetSlots

`func (o *EquipmentRackEnclosureAllOf) GetSlots() []EquipmentRackEnclosureSlotRelationship`

GetSlots returns the Slots field if non-nil, zero value otherwise.

### GetSlotsOk

`func (o *EquipmentRackEnclosureAllOf) GetSlotsOk() (*[]EquipmentRackEnclosureSlotRelationship, bool)`

GetSlotsOk returns a tuple with the Slots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlots

`func (o *EquipmentRackEnclosureAllOf) SetSlots(v []EquipmentRackEnclosureSlotRelationship)`

SetSlots sets Slots field to given value.

### HasSlots

`func (o *EquipmentRackEnclosureAllOf) HasSlots() bool`

HasSlots returns a boolean if a field has been set.

### SetSlotsNil

`func (o *EquipmentRackEnclosureAllOf) SetSlotsNil(b bool)`

 SetSlotsNil sets the value for Slots to be an explicit nil

### UnsetSlots
`func (o *EquipmentRackEnclosureAllOf) UnsetSlots()`

UnsetSlots ensures that no value is present for Slots, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



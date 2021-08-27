# EquipmentRackEnclosureSlot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "equipment.RackEnclosureSlot"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "equipment.RackEnclosureSlot"]
**RackId** | Pointer to **int64** | Server ID which is part of Rack Enclosure Slot. | [optional] [readonly] 
**RackUnitDn** | Pointer to **string** | Server DN which is part of Rack Enclosure Slot. | [optional] [readonly] 
**EquipmentRackEnclosure** | Pointer to [**EquipmentRackEnclosureRelationship**](EquipmentRackEnclosureRelationship.md) |  | [optional] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](InventoryDeviceInfoRelationship.md) |  | [optional] 
**RackUnit** | Pointer to [**ComputeRackUnitRelationship**](ComputeRackUnitRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewEquipmentRackEnclosureSlot

`func NewEquipmentRackEnclosureSlot(classId string, objectType string, ) *EquipmentRackEnclosureSlot`

NewEquipmentRackEnclosureSlot instantiates a new EquipmentRackEnclosureSlot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquipmentRackEnclosureSlotWithDefaults

`func NewEquipmentRackEnclosureSlotWithDefaults() *EquipmentRackEnclosureSlot`

NewEquipmentRackEnclosureSlotWithDefaults instantiates a new EquipmentRackEnclosureSlot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *EquipmentRackEnclosureSlot) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *EquipmentRackEnclosureSlot) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *EquipmentRackEnclosureSlot) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *EquipmentRackEnclosureSlot) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *EquipmentRackEnclosureSlot) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *EquipmentRackEnclosureSlot) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


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



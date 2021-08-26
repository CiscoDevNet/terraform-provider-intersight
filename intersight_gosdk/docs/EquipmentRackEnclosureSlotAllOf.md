# EquipmentRackEnclosureSlotAllOf

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

### NewEquipmentRackEnclosureSlotAllOf

`func NewEquipmentRackEnclosureSlotAllOf(classId string, objectType string, ) *EquipmentRackEnclosureSlotAllOf`

NewEquipmentRackEnclosureSlotAllOf instantiates a new EquipmentRackEnclosureSlotAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquipmentRackEnclosureSlotAllOfWithDefaults

`func NewEquipmentRackEnclosureSlotAllOfWithDefaults() *EquipmentRackEnclosureSlotAllOf`

NewEquipmentRackEnclosureSlotAllOfWithDefaults instantiates a new EquipmentRackEnclosureSlotAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *EquipmentRackEnclosureSlotAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *EquipmentRackEnclosureSlotAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *EquipmentRackEnclosureSlotAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *EquipmentRackEnclosureSlotAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *EquipmentRackEnclosureSlotAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *EquipmentRackEnclosureSlotAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetRackId

`func (o *EquipmentRackEnclosureSlotAllOf) GetRackId() int64`

GetRackId returns the RackId field if non-nil, zero value otherwise.

### GetRackIdOk

`func (o *EquipmentRackEnclosureSlotAllOf) GetRackIdOk() (*int64, bool)`

GetRackIdOk returns a tuple with the RackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRackId

`func (o *EquipmentRackEnclosureSlotAllOf) SetRackId(v int64)`

SetRackId sets RackId field to given value.

### HasRackId

`func (o *EquipmentRackEnclosureSlotAllOf) HasRackId() bool`

HasRackId returns a boolean if a field has been set.

### GetRackUnitDn

`func (o *EquipmentRackEnclosureSlotAllOf) GetRackUnitDn() string`

GetRackUnitDn returns the RackUnitDn field if non-nil, zero value otherwise.

### GetRackUnitDnOk

`func (o *EquipmentRackEnclosureSlotAllOf) GetRackUnitDnOk() (*string, bool)`

GetRackUnitDnOk returns a tuple with the RackUnitDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRackUnitDn

`func (o *EquipmentRackEnclosureSlotAllOf) SetRackUnitDn(v string)`

SetRackUnitDn sets RackUnitDn field to given value.

### HasRackUnitDn

`func (o *EquipmentRackEnclosureSlotAllOf) HasRackUnitDn() bool`

HasRackUnitDn returns a boolean if a field has been set.

### GetEquipmentRackEnclosure

`func (o *EquipmentRackEnclosureSlotAllOf) GetEquipmentRackEnclosure() EquipmentRackEnclosureRelationship`

GetEquipmentRackEnclosure returns the EquipmentRackEnclosure field if non-nil, zero value otherwise.

### GetEquipmentRackEnclosureOk

`func (o *EquipmentRackEnclosureSlotAllOf) GetEquipmentRackEnclosureOk() (*EquipmentRackEnclosureRelationship, bool)`

GetEquipmentRackEnclosureOk returns a tuple with the EquipmentRackEnclosure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentRackEnclosure

`func (o *EquipmentRackEnclosureSlotAllOf) SetEquipmentRackEnclosure(v EquipmentRackEnclosureRelationship)`

SetEquipmentRackEnclosure sets EquipmentRackEnclosure field to given value.

### HasEquipmentRackEnclosure

`func (o *EquipmentRackEnclosureSlotAllOf) HasEquipmentRackEnclosure() bool`

HasEquipmentRackEnclosure returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *EquipmentRackEnclosureSlotAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *EquipmentRackEnclosureSlotAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *EquipmentRackEnclosureSlotAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *EquipmentRackEnclosureSlotAllOf) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetRackUnit

`func (o *EquipmentRackEnclosureSlotAllOf) GetRackUnit() ComputeRackUnitRelationship`

GetRackUnit returns the RackUnit field if non-nil, zero value otherwise.

### GetRackUnitOk

`func (o *EquipmentRackEnclosureSlotAllOf) GetRackUnitOk() (*ComputeRackUnitRelationship, bool)`

GetRackUnitOk returns a tuple with the RackUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRackUnit

`func (o *EquipmentRackEnclosureSlotAllOf) SetRackUnit(v ComputeRackUnitRelationship)`

SetRackUnit sets RackUnit field to given value.

### HasRackUnit

`func (o *EquipmentRackEnclosureSlotAllOf) HasRackUnit() bool`

HasRackUnit returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *EquipmentRackEnclosureSlotAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *EquipmentRackEnclosureSlotAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *EquipmentRackEnclosureSlotAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *EquipmentRackEnclosureSlotAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



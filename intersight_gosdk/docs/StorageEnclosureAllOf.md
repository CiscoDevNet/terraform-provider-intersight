# StorageEnclosureAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChassisId** | Pointer to **int64** | This represent the chassis-ID that houses the storage enclosure. | [optional] [readonly] 
**Description** | Pointer to **string** | This represnets the description for the storage enclosure. | [optional] [readonly] 
**EnclosureId** | Pointer to **int64** | This represnets the Identifier for the storage enclosure. | [optional] [readonly] 
**NumSlots** | Pointer to **int64** | This represent the number of slots present in storage enclosure. | [optional] [readonly] 
**Presence** | Pointer to **string** | This represent the availability of storage enclosure. | [optional] [readonly] 
**ServerId** | Pointer to **int64** | This represent the server-ID that houses the storage enclosure. | [optional] [readonly] 
**Type** | Pointer to **string** | This represent the type of storage enclosure. | [optional] [readonly] 
**ComputeBlade** | Pointer to [**ComputeBladeRelationship**](compute.Blade.Relationship.md) |  | [optional] 
**ComputeRackUnit** | Pointer to [**ComputeRackUnitRelationship**](compute.RackUnit.Relationship.md) |  | [optional] 
**EnclosureDiskSlots** | Pointer to [**[]StorageEnclosureDiskSlotEpRelationship**](storage.EnclosureDiskSlotEp.Relationship.md) | An array of relationships to storageEnclosureDiskSlotEp resources. | [optional] [readonly] 
**EnclosureDisks** | Pointer to [**[]StorageEnclosureDiskRelationship**](storage.EnclosureDisk.Relationship.md) | An array of relationships to storageEnclosureDisk resources. | [optional] [readonly] 
**EquipmentChassis** | Pointer to [**EquipmentChassisRelationship**](equipment.Chassis.Relationship.md) |  | [optional] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](inventory.DeviceInfo.Relationship.md) |  | [optional] 
**PhysicalDisks** | Pointer to [**[]StoragePhysicalDiskRelationship**](storage.PhysicalDisk.Relationship.md) | An array of relationships to storagePhysicalDisk resources. | [optional] [readonly] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewStorageEnclosureAllOf

`func NewStorageEnclosureAllOf() *StorageEnclosureAllOf`

NewStorageEnclosureAllOf instantiates a new StorageEnclosureAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageEnclosureAllOfWithDefaults

`func NewStorageEnclosureAllOfWithDefaults() *StorageEnclosureAllOf`

NewStorageEnclosureAllOfWithDefaults instantiates a new StorageEnclosureAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChassisId

`func (o *StorageEnclosureAllOf) GetChassisId() int64`

GetChassisId returns the ChassisId field if non-nil, zero value otherwise.

### GetChassisIdOk

`func (o *StorageEnclosureAllOf) GetChassisIdOk() (*int64, bool)`

GetChassisIdOk returns a tuple with the ChassisId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassisId

`func (o *StorageEnclosureAllOf) SetChassisId(v int64)`

SetChassisId sets ChassisId field to given value.

### HasChassisId

`func (o *StorageEnclosureAllOf) HasChassisId() bool`

HasChassisId returns a boolean if a field has been set.

### GetDescription

`func (o *StorageEnclosureAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StorageEnclosureAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StorageEnclosureAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *StorageEnclosureAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnclosureId

`func (o *StorageEnclosureAllOf) GetEnclosureId() int64`

GetEnclosureId returns the EnclosureId field if non-nil, zero value otherwise.

### GetEnclosureIdOk

`func (o *StorageEnclosureAllOf) GetEnclosureIdOk() (*int64, bool)`

GetEnclosureIdOk returns a tuple with the EnclosureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnclosureId

`func (o *StorageEnclosureAllOf) SetEnclosureId(v int64)`

SetEnclosureId sets EnclosureId field to given value.

### HasEnclosureId

`func (o *StorageEnclosureAllOf) HasEnclosureId() bool`

HasEnclosureId returns a boolean if a field has been set.

### GetNumSlots

`func (o *StorageEnclosureAllOf) GetNumSlots() int64`

GetNumSlots returns the NumSlots field if non-nil, zero value otherwise.

### GetNumSlotsOk

`func (o *StorageEnclosureAllOf) GetNumSlotsOk() (*int64, bool)`

GetNumSlotsOk returns a tuple with the NumSlots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumSlots

`func (o *StorageEnclosureAllOf) SetNumSlots(v int64)`

SetNumSlots sets NumSlots field to given value.

### HasNumSlots

`func (o *StorageEnclosureAllOf) HasNumSlots() bool`

HasNumSlots returns a boolean if a field has been set.

### GetPresence

`func (o *StorageEnclosureAllOf) GetPresence() string`

GetPresence returns the Presence field if non-nil, zero value otherwise.

### GetPresenceOk

`func (o *StorageEnclosureAllOf) GetPresenceOk() (*string, bool)`

GetPresenceOk returns a tuple with the Presence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresence

`func (o *StorageEnclosureAllOf) SetPresence(v string)`

SetPresence sets Presence field to given value.

### HasPresence

`func (o *StorageEnclosureAllOf) HasPresence() bool`

HasPresence returns a boolean if a field has been set.

### GetServerId

`func (o *StorageEnclosureAllOf) GetServerId() int64`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *StorageEnclosureAllOf) GetServerIdOk() (*int64, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *StorageEnclosureAllOf) SetServerId(v int64)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *StorageEnclosureAllOf) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetType

`func (o *StorageEnclosureAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StorageEnclosureAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StorageEnclosureAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *StorageEnclosureAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetComputeBlade

`func (o *StorageEnclosureAllOf) GetComputeBlade() ComputeBladeRelationship`

GetComputeBlade returns the ComputeBlade field if non-nil, zero value otherwise.

### GetComputeBladeOk

`func (o *StorageEnclosureAllOf) GetComputeBladeOk() (*ComputeBladeRelationship, bool)`

GetComputeBladeOk returns a tuple with the ComputeBlade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeBlade

`func (o *StorageEnclosureAllOf) SetComputeBlade(v ComputeBladeRelationship)`

SetComputeBlade sets ComputeBlade field to given value.

### HasComputeBlade

`func (o *StorageEnclosureAllOf) HasComputeBlade() bool`

HasComputeBlade returns a boolean if a field has been set.

### GetComputeRackUnit

`func (o *StorageEnclosureAllOf) GetComputeRackUnit() ComputeRackUnitRelationship`

GetComputeRackUnit returns the ComputeRackUnit field if non-nil, zero value otherwise.

### GetComputeRackUnitOk

`func (o *StorageEnclosureAllOf) GetComputeRackUnitOk() (*ComputeRackUnitRelationship, bool)`

GetComputeRackUnitOk returns a tuple with the ComputeRackUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRackUnit

`func (o *StorageEnclosureAllOf) SetComputeRackUnit(v ComputeRackUnitRelationship)`

SetComputeRackUnit sets ComputeRackUnit field to given value.

### HasComputeRackUnit

`func (o *StorageEnclosureAllOf) HasComputeRackUnit() bool`

HasComputeRackUnit returns a boolean if a field has been set.

### GetEnclosureDiskSlots

`func (o *StorageEnclosureAllOf) GetEnclosureDiskSlots() []StorageEnclosureDiskSlotEpRelationship`

GetEnclosureDiskSlots returns the EnclosureDiskSlots field if non-nil, zero value otherwise.

### GetEnclosureDiskSlotsOk

`func (o *StorageEnclosureAllOf) GetEnclosureDiskSlotsOk() (*[]StorageEnclosureDiskSlotEpRelationship, bool)`

GetEnclosureDiskSlotsOk returns a tuple with the EnclosureDiskSlots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnclosureDiskSlots

`func (o *StorageEnclosureAllOf) SetEnclosureDiskSlots(v []StorageEnclosureDiskSlotEpRelationship)`

SetEnclosureDiskSlots sets EnclosureDiskSlots field to given value.

### HasEnclosureDiskSlots

`func (o *StorageEnclosureAllOf) HasEnclosureDiskSlots() bool`

HasEnclosureDiskSlots returns a boolean if a field has been set.

### SetEnclosureDiskSlotsNil

`func (o *StorageEnclosureAllOf) SetEnclosureDiskSlotsNil(b bool)`

 SetEnclosureDiskSlotsNil sets the value for EnclosureDiskSlots to be an explicit nil

### UnsetEnclosureDiskSlots
`func (o *StorageEnclosureAllOf) UnsetEnclosureDiskSlots()`

UnsetEnclosureDiskSlots ensures that no value is present for EnclosureDiskSlots, not even an explicit nil
### GetEnclosureDisks

`func (o *StorageEnclosureAllOf) GetEnclosureDisks() []StorageEnclosureDiskRelationship`

GetEnclosureDisks returns the EnclosureDisks field if non-nil, zero value otherwise.

### GetEnclosureDisksOk

`func (o *StorageEnclosureAllOf) GetEnclosureDisksOk() (*[]StorageEnclosureDiskRelationship, bool)`

GetEnclosureDisksOk returns a tuple with the EnclosureDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnclosureDisks

`func (o *StorageEnclosureAllOf) SetEnclosureDisks(v []StorageEnclosureDiskRelationship)`

SetEnclosureDisks sets EnclosureDisks field to given value.

### HasEnclosureDisks

`func (o *StorageEnclosureAllOf) HasEnclosureDisks() bool`

HasEnclosureDisks returns a boolean if a field has been set.

### SetEnclosureDisksNil

`func (o *StorageEnclosureAllOf) SetEnclosureDisksNil(b bool)`

 SetEnclosureDisksNil sets the value for EnclosureDisks to be an explicit nil

### UnsetEnclosureDisks
`func (o *StorageEnclosureAllOf) UnsetEnclosureDisks()`

UnsetEnclosureDisks ensures that no value is present for EnclosureDisks, not even an explicit nil
### GetEquipmentChassis

`func (o *StorageEnclosureAllOf) GetEquipmentChassis() EquipmentChassisRelationship`

GetEquipmentChassis returns the EquipmentChassis field if non-nil, zero value otherwise.

### GetEquipmentChassisOk

`func (o *StorageEnclosureAllOf) GetEquipmentChassisOk() (*EquipmentChassisRelationship, bool)`

GetEquipmentChassisOk returns a tuple with the EquipmentChassis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentChassis

`func (o *StorageEnclosureAllOf) SetEquipmentChassis(v EquipmentChassisRelationship)`

SetEquipmentChassis sets EquipmentChassis field to given value.

### HasEquipmentChassis

`func (o *StorageEnclosureAllOf) HasEquipmentChassis() bool`

HasEquipmentChassis returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *StorageEnclosureAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *StorageEnclosureAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *StorageEnclosureAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *StorageEnclosureAllOf) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetPhysicalDisks

`func (o *StorageEnclosureAllOf) GetPhysicalDisks() []StoragePhysicalDiskRelationship`

GetPhysicalDisks returns the PhysicalDisks field if non-nil, zero value otherwise.

### GetPhysicalDisksOk

`func (o *StorageEnclosureAllOf) GetPhysicalDisksOk() (*[]StoragePhysicalDiskRelationship, bool)`

GetPhysicalDisksOk returns a tuple with the PhysicalDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalDisks

`func (o *StorageEnclosureAllOf) SetPhysicalDisks(v []StoragePhysicalDiskRelationship)`

SetPhysicalDisks sets PhysicalDisks field to given value.

### HasPhysicalDisks

`func (o *StorageEnclosureAllOf) HasPhysicalDisks() bool`

HasPhysicalDisks returns a boolean if a field has been set.

### SetPhysicalDisksNil

`func (o *StorageEnclosureAllOf) SetPhysicalDisksNil(b bool)`

 SetPhysicalDisksNil sets the value for PhysicalDisks to be an explicit nil

### UnsetPhysicalDisks
`func (o *StorageEnclosureAllOf) UnsetPhysicalDisks()`

UnsetPhysicalDisks ensures that no value is present for PhysicalDisks, not even an explicit nil
### GetRegisteredDevice

`func (o *StorageEnclosureAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StorageEnclosureAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StorageEnclosureAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StorageEnclosureAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



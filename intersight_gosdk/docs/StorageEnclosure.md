# StorageEnclosure

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

### NewStorageEnclosure

`func NewStorageEnclosure() *StorageEnclosure`

NewStorageEnclosure instantiates a new StorageEnclosure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageEnclosureWithDefaults

`func NewStorageEnclosureWithDefaults() *StorageEnclosure`

NewStorageEnclosureWithDefaults instantiates a new StorageEnclosure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChassisId

`func (o *StorageEnclosure) GetChassisId() int64`

GetChassisId returns the ChassisId field if non-nil, zero value otherwise.

### GetChassisIdOk

`func (o *StorageEnclosure) GetChassisIdOk() (*int64, bool)`

GetChassisIdOk returns a tuple with the ChassisId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassisId

`func (o *StorageEnclosure) SetChassisId(v int64)`

SetChassisId sets ChassisId field to given value.

### HasChassisId

`func (o *StorageEnclosure) HasChassisId() bool`

HasChassisId returns a boolean if a field has been set.

### GetDescription

`func (o *StorageEnclosure) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StorageEnclosure) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StorageEnclosure) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *StorageEnclosure) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnclosureId

`func (o *StorageEnclosure) GetEnclosureId() int64`

GetEnclosureId returns the EnclosureId field if non-nil, zero value otherwise.

### GetEnclosureIdOk

`func (o *StorageEnclosure) GetEnclosureIdOk() (*int64, bool)`

GetEnclosureIdOk returns a tuple with the EnclosureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnclosureId

`func (o *StorageEnclosure) SetEnclosureId(v int64)`

SetEnclosureId sets EnclosureId field to given value.

### HasEnclosureId

`func (o *StorageEnclosure) HasEnclosureId() bool`

HasEnclosureId returns a boolean if a field has been set.

### GetNumSlots

`func (o *StorageEnclosure) GetNumSlots() int64`

GetNumSlots returns the NumSlots field if non-nil, zero value otherwise.

### GetNumSlotsOk

`func (o *StorageEnclosure) GetNumSlotsOk() (*int64, bool)`

GetNumSlotsOk returns a tuple with the NumSlots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumSlots

`func (o *StorageEnclosure) SetNumSlots(v int64)`

SetNumSlots sets NumSlots field to given value.

### HasNumSlots

`func (o *StorageEnclosure) HasNumSlots() bool`

HasNumSlots returns a boolean if a field has been set.

### GetPresence

`func (o *StorageEnclosure) GetPresence() string`

GetPresence returns the Presence field if non-nil, zero value otherwise.

### GetPresenceOk

`func (o *StorageEnclosure) GetPresenceOk() (*string, bool)`

GetPresenceOk returns a tuple with the Presence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresence

`func (o *StorageEnclosure) SetPresence(v string)`

SetPresence sets Presence field to given value.

### HasPresence

`func (o *StorageEnclosure) HasPresence() bool`

HasPresence returns a boolean if a field has been set.

### GetServerId

`func (o *StorageEnclosure) GetServerId() int64`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *StorageEnclosure) GetServerIdOk() (*int64, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *StorageEnclosure) SetServerId(v int64)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *StorageEnclosure) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetType

`func (o *StorageEnclosure) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StorageEnclosure) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StorageEnclosure) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *StorageEnclosure) HasType() bool`

HasType returns a boolean if a field has been set.

### GetComputeBlade

`func (o *StorageEnclosure) GetComputeBlade() ComputeBladeRelationship`

GetComputeBlade returns the ComputeBlade field if non-nil, zero value otherwise.

### GetComputeBladeOk

`func (o *StorageEnclosure) GetComputeBladeOk() (*ComputeBladeRelationship, bool)`

GetComputeBladeOk returns a tuple with the ComputeBlade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeBlade

`func (o *StorageEnclosure) SetComputeBlade(v ComputeBladeRelationship)`

SetComputeBlade sets ComputeBlade field to given value.

### HasComputeBlade

`func (o *StorageEnclosure) HasComputeBlade() bool`

HasComputeBlade returns a boolean if a field has been set.

### GetComputeRackUnit

`func (o *StorageEnclosure) GetComputeRackUnit() ComputeRackUnitRelationship`

GetComputeRackUnit returns the ComputeRackUnit field if non-nil, zero value otherwise.

### GetComputeRackUnitOk

`func (o *StorageEnclosure) GetComputeRackUnitOk() (*ComputeRackUnitRelationship, bool)`

GetComputeRackUnitOk returns a tuple with the ComputeRackUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRackUnit

`func (o *StorageEnclosure) SetComputeRackUnit(v ComputeRackUnitRelationship)`

SetComputeRackUnit sets ComputeRackUnit field to given value.

### HasComputeRackUnit

`func (o *StorageEnclosure) HasComputeRackUnit() bool`

HasComputeRackUnit returns a boolean if a field has been set.

### GetEnclosureDiskSlots

`func (o *StorageEnclosure) GetEnclosureDiskSlots() []StorageEnclosureDiskSlotEpRelationship`

GetEnclosureDiskSlots returns the EnclosureDiskSlots field if non-nil, zero value otherwise.

### GetEnclosureDiskSlotsOk

`func (o *StorageEnclosure) GetEnclosureDiskSlotsOk() (*[]StorageEnclosureDiskSlotEpRelationship, bool)`

GetEnclosureDiskSlotsOk returns a tuple with the EnclosureDiskSlots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnclosureDiskSlots

`func (o *StorageEnclosure) SetEnclosureDiskSlots(v []StorageEnclosureDiskSlotEpRelationship)`

SetEnclosureDiskSlots sets EnclosureDiskSlots field to given value.

### HasEnclosureDiskSlots

`func (o *StorageEnclosure) HasEnclosureDiskSlots() bool`

HasEnclosureDiskSlots returns a boolean if a field has been set.

### SetEnclosureDiskSlotsNil

`func (o *StorageEnclosure) SetEnclosureDiskSlotsNil(b bool)`

 SetEnclosureDiskSlotsNil sets the value for EnclosureDiskSlots to be an explicit nil

### UnsetEnclosureDiskSlots
`func (o *StorageEnclosure) UnsetEnclosureDiskSlots()`

UnsetEnclosureDiskSlots ensures that no value is present for EnclosureDiskSlots, not even an explicit nil
### GetEnclosureDisks

`func (o *StorageEnclosure) GetEnclosureDisks() []StorageEnclosureDiskRelationship`

GetEnclosureDisks returns the EnclosureDisks field if non-nil, zero value otherwise.

### GetEnclosureDisksOk

`func (o *StorageEnclosure) GetEnclosureDisksOk() (*[]StorageEnclosureDiskRelationship, bool)`

GetEnclosureDisksOk returns a tuple with the EnclosureDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnclosureDisks

`func (o *StorageEnclosure) SetEnclosureDisks(v []StorageEnclosureDiskRelationship)`

SetEnclosureDisks sets EnclosureDisks field to given value.

### HasEnclosureDisks

`func (o *StorageEnclosure) HasEnclosureDisks() bool`

HasEnclosureDisks returns a boolean if a field has been set.

### SetEnclosureDisksNil

`func (o *StorageEnclosure) SetEnclosureDisksNil(b bool)`

 SetEnclosureDisksNil sets the value for EnclosureDisks to be an explicit nil

### UnsetEnclosureDisks
`func (o *StorageEnclosure) UnsetEnclosureDisks()`

UnsetEnclosureDisks ensures that no value is present for EnclosureDisks, not even an explicit nil
### GetEquipmentChassis

`func (o *StorageEnclosure) GetEquipmentChassis() EquipmentChassisRelationship`

GetEquipmentChassis returns the EquipmentChassis field if non-nil, zero value otherwise.

### GetEquipmentChassisOk

`func (o *StorageEnclosure) GetEquipmentChassisOk() (*EquipmentChassisRelationship, bool)`

GetEquipmentChassisOk returns a tuple with the EquipmentChassis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentChassis

`func (o *StorageEnclosure) SetEquipmentChassis(v EquipmentChassisRelationship)`

SetEquipmentChassis sets EquipmentChassis field to given value.

### HasEquipmentChassis

`func (o *StorageEnclosure) HasEquipmentChassis() bool`

HasEquipmentChassis returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *StorageEnclosure) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *StorageEnclosure) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *StorageEnclosure) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *StorageEnclosure) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetPhysicalDisks

`func (o *StorageEnclosure) GetPhysicalDisks() []StoragePhysicalDiskRelationship`

GetPhysicalDisks returns the PhysicalDisks field if non-nil, zero value otherwise.

### GetPhysicalDisksOk

`func (o *StorageEnclosure) GetPhysicalDisksOk() (*[]StoragePhysicalDiskRelationship, bool)`

GetPhysicalDisksOk returns a tuple with the PhysicalDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalDisks

`func (o *StorageEnclosure) SetPhysicalDisks(v []StoragePhysicalDiskRelationship)`

SetPhysicalDisks sets PhysicalDisks field to given value.

### HasPhysicalDisks

`func (o *StorageEnclosure) HasPhysicalDisks() bool`

HasPhysicalDisks returns a boolean if a field has been set.

### SetPhysicalDisksNil

`func (o *StorageEnclosure) SetPhysicalDisksNil(b bool)`

 SetPhysicalDisksNil sets the value for PhysicalDisks to be an explicit nil

### UnsetPhysicalDisks
`func (o *StorageEnclosure) UnsetPhysicalDisks()`

UnsetPhysicalDisks ensures that no value is present for PhysicalDisks, not even an explicit nil
### GetRegisteredDevice

`func (o *StorageEnclosure) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StorageEnclosure) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StorageEnclosure) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StorageEnclosure) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



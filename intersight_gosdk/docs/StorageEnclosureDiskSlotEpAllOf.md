# StorageEnclosureDiskSlotEpAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.EnclosureDiskSlotEp"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.EnclosureDiskSlotEp"]
**DrivePath** | Pointer to **string** | This field identifies the zoning configuration applied to  this enclosure slot. | [optional] 
**Health** | Pointer to **string** | This field identifies the health of the disk inserted in the slot. | [optional] 
**Presence** | Pointer to **string** | This field identifies the disk is present in the enclosure slot. | [optional] 
**Slot** | Pointer to **string** | This field represents the slot Id in the storage enclosure. | [optional] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](InventoryDeviceInfoRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 
**StorageEnclosure** | Pointer to [**StorageEnclosureRelationship**](StorageEnclosureRelationship.md) |  | [optional] 

## Methods

### NewStorageEnclosureDiskSlotEpAllOf

`func NewStorageEnclosureDiskSlotEpAllOf(classId string, objectType string, ) *StorageEnclosureDiskSlotEpAllOf`

NewStorageEnclosureDiskSlotEpAllOf instantiates a new StorageEnclosureDiskSlotEpAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageEnclosureDiskSlotEpAllOfWithDefaults

`func NewStorageEnclosureDiskSlotEpAllOfWithDefaults() *StorageEnclosureDiskSlotEpAllOf`

NewStorageEnclosureDiskSlotEpAllOfWithDefaults instantiates a new StorageEnclosureDiskSlotEpAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageEnclosureDiskSlotEpAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageEnclosureDiskSlotEpAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageEnclosureDiskSlotEpAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageEnclosureDiskSlotEpAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageEnclosureDiskSlotEpAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageEnclosureDiskSlotEpAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDrivePath

`func (o *StorageEnclosureDiskSlotEpAllOf) GetDrivePath() string`

GetDrivePath returns the DrivePath field if non-nil, zero value otherwise.

### GetDrivePathOk

`func (o *StorageEnclosureDiskSlotEpAllOf) GetDrivePathOk() (*string, bool)`

GetDrivePathOk returns a tuple with the DrivePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrivePath

`func (o *StorageEnclosureDiskSlotEpAllOf) SetDrivePath(v string)`

SetDrivePath sets DrivePath field to given value.

### HasDrivePath

`func (o *StorageEnclosureDiskSlotEpAllOf) HasDrivePath() bool`

HasDrivePath returns a boolean if a field has been set.

### GetHealth

`func (o *StorageEnclosureDiskSlotEpAllOf) GetHealth() string`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *StorageEnclosureDiskSlotEpAllOf) GetHealthOk() (*string, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *StorageEnclosureDiskSlotEpAllOf) SetHealth(v string)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *StorageEnclosureDiskSlotEpAllOf) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetPresence

`func (o *StorageEnclosureDiskSlotEpAllOf) GetPresence() string`

GetPresence returns the Presence field if non-nil, zero value otherwise.

### GetPresenceOk

`func (o *StorageEnclosureDiskSlotEpAllOf) GetPresenceOk() (*string, bool)`

GetPresenceOk returns a tuple with the Presence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresence

`func (o *StorageEnclosureDiskSlotEpAllOf) SetPresence(v string)`

SetPresence sets Presence field to given value.

### HasPresence

`func (o *StorageEnclosureDiskSlotEpAllOf) HasPresence() bool`

HasPresence returns a boolean if a field has been set.

### GetSlot

`func (o *StorageEnclosureDiskSlotEpAllOf) GetSlot() string`

GetSlot returns the Slot field if non-nil, zero value otherwise.

### GetSlotOk

`func (o *StorageEnclosureDiskSlotEpAllOf) GetSlotOk() (*string, bool)`

GetSlotOk returns a tuple with the Slot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlot

`func (o *StorageEnclosureDiskSlotEpAllOf) SetSlot(v string)`

SetSlot sets Slot field to given value.

### HasSlot

`func (o *StorageEnclosureDiskSlotEpAllOf) HasSlot() bool`

HasSlot returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *StorageEnclosureDiskSlotEpAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *StorageEnclosureDiskSlotEpAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *StorageEnclosureDiskSlotEpAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *StorageEnclosureDiskSlotEpAllOf) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *StorageEnclosureDiskSlotEpAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StorageEnclosureDiskSlotEpAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StorageEnclosureDiskSlotEpAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StorageEnclosureDiskSlotEpAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### GetStorageEnclosure

`func (o *StorageEnclosureDiskSlotEpAllOf) GetStorageEnclosure() StorageEnclosureRelationship`

GetStorageEnclosure returns the StorageEnclosure field if non-nil, zero value otherwise.

### GetStorageEnclosureOk

`func (o *StorageEnclosureDiskSlotEpAllOf) GetStorageEnclosureOk() (*StorageEnclosureRelationship, bool)`

GetStorageEnclosureOk returns a tuple with the StorageEnclosure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageEnclosure

`func (o *StorageEnclosureDiskSlotEpAllOf) SetStorageEnclosure(v StorageEnclosureRelationship)`

SetStorageEnclosure sets StorageEnclosure field to given value.

### HasStorageEnclosure

`func (o *StorageEnclosureDiskSlotEpAllOf) HasStorageEnclosure() bool`

HasStorageEnclosure returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



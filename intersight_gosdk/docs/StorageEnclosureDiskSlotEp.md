# StorageEnclosureDiskSlotEp

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

### NewStorageEnclosureDiskSlotEp

`func NewStorageEnclosureDiskSlotEp(classId string, objectType string, ) *StorageEnclosureDiskSlotEp`

NewStorageEnclosureDiskSlotEp instantiates a new StorageEnclosureDiskSlotEp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageEnclosureDiskSlotEpWithDefaults

`func NewStorageEnclosureDiskSlotEpWithDefaults() *StorageEnclosureDiskSlotEp`

NewStorageEnclosureDiskSlotEpWithDefaults instantiates a new StorageEnclosureDiskSlotEp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageEnclosureDiskSlotEp) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageEnclosureDiskSlotEp) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageEnclosureDiskSlotEp) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageEnclosureDiskSlotEp) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageEnclosureDiskSlotEp) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageEnclosureDiskSlotEp) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDrivePath

`func (o *StorageEnclosureDiskSlotEp) GetDrivePath() string`

GetDrivePath returns the DrivePath field if non-nil, zero value otherwise.

### GetDrivePathOk

`func (o *StorageEnclosureDiskSlotEp) GetDrivePathOk() (*string, bool)`

GetDrivePathOk returns a tuple with the DrivePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrivePath

`func (o *StorageEnclosureDiskSlotEp) SetDrivePath(v string)`

SetDrivePath sets DrivePath field to given value.

### HasDrivePath

`func (o *StorageEnclosureDiskSlotEp) HasDrivePath() bool`

HasDrivePath returns a boolean if a field has been set.

### GetHealth

`func (o *StorageEnclosureDiskSlotEp) GetHealth() string`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *StorageEnclosureDiskSlotEp) GetHealthOk() (*string, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *StorageEnclosureDiskSlotEp) SetHealth(v string)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *StorageEnclosureDiskSlotEp) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetPresence

`func (o *StorageEnclosureDiskSlotEp) GetPresence() string`

GetPresence returns the Presence field if non-nil, zero value otherwise.

### GetPresenceOk

`func (o *StorageEnclosureDiskSlotEp) GetPresenceOk() (*string, bool)`

GetPresenceOk returns a tuple with the Presence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresence

`func (o *StorageEnclosureDiskSlotEp) SetPresence(v string)`

SetPresence sets Presence field to given value.

### HasPresence

`func (o *StorageEnclosureDiskSlotEp) HasPresence() bool`

HasPresence returns a boolean if a field has been set.

### GetSlot

`func (o *StorageEnclosureDiskSlotEp) GetSlot() string`

GetSlot returns the Slot field if non-nil, zero value otherwise.

### GetSlotOk

`func (o *StorageEnclosureDiskSlotEp) GetSlotOk() (*string, bool)`

GetSlotOk returns a tuple with the Slot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlot

`func (o *StorageEnclosureDiskSlotEp) SetSlot(v string)`

SetSlot sets Slot field to given value.

### HasSlot

`func (o *StorageEnclosureDiskSlotEp) HasSlot() bool`

HasSlot returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *StorageEnclosureDiskSlotEp) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *StorageEnclosureDiskSlotEp) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *StorageEnclosureDiskSlotEp) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *StorageEnclosureDiskSlotEp) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *StorageEnclosureDiskSlotEp) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StorageEnclosureDiskSlotEp) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StorageEnclosureDiskSlotEp) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StorageEnclosureDiskSlotEp) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### GetStorageEnclosure

`func (o *StorageEnclosureDiskSlotEp) GetStorageEnclosure() StorageEnclosureRelationship`

GetStorageEnclosure returns the StorageEnclosure field if non-nil, zero value otherwise.

### GetStorageEnclosureOk

`func (o *StorageEnclosureDiskSlotEp) GetStorageEnclosureOk() (*StorageEnclosureRelationship, bool)`

GetStorageEnclosureOk returns a tuple with the StorageEnclosure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageEnclosure

`func (o *StorageEnclosureDiskSlotEp) SetStorageEnclosure(v StorageEnclosureRelationship)`

SetStorageEnclosure sets StorageEnclosure field to given value.

### HasStorageEnclosure

`func (o *StorageEnclosureDiskSlotEp) HasStorageEnclosure() bool`

HasStorageEnclosure returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



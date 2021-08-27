# StorageDiskGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.DiskGroup"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.DiskGroup"]
**Name** | Pointer to **string** | Name to identity this disk group in the controller. | [optional] 
**RaidType** | Pointer to **string** | Raid level of the virtual drives in this diskgroup. | [optional] 
**DedicatedHotSpares** | Pointer to [**[]StoragePhysicalDiskRelationship**](StoragePhysicalDiskRelationship.md) | An array of relationships to storagePhysicalDisk resources. | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 
**Spans** | Pointer to [**[]StorageSpanRelationship**](StorageSpanRelationship.md) | An array of relationships to storageSpan resources. | [optional] 
**StorageController** | Pointer to [**StorageControllerRelationship**](StorageControllerRelationship.md) |  | [optional] 
**VirtualDrives** | Pointer to [**[]StorageVirtualDriveRelationship**](StorageVirtualDriveRelationship.md) | An array of relationships to storageVirtualDrive resources. | [optional] 

## Methods

### NewStorageDiskGroup

`func NewStorageDiskGroup(classId string, objectType string, ) *StorageDiskGroup`

NewStorageDiskGroup instantiates a new StorageDiskGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageDiskGroupWithDefaults

`func NewStorageDiskGroupWithDefaults() *StorageDiskGroup`

NewStorageDiskGroupWithDefaults instantiates a new StorageDiskGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageDiskGroup) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageDiskGroup) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageDiskGroup) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageDiskGroup) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageDiskGroup) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageDiskGroup) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetName

`func (o *StorageDiskGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageDiskGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageDiskGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StorageDiskGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRaidType

`func (o *StorageDiskGroup) GetRaidType() string`

GetRaidType returns the RaidType field if non-nil, zero value otherwise.

### GetRaidTypeOk

`func (o *StorageDiskGroup) GetRaidTypeOk() (*string, bool)`

GetRaidTypeOk returns a tuple with the RaidType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaidType

`func (o *StorageDiskGroup) SetRaidType(v string)`

SetRaidType sets RaidType field to given value.

### HasRaidType

`func (o *StorageDiskGroup) HasRaidType() bool`

HasRaidType returns a boolean if a field has been set.

### GetDedicatedHotSpares

`func (o *StorageDiskGroup) GetDedicatedHotSpares() []StoragePhysicalDiskRelationship`

GetDedicatedHotSpares returns the DedicatedHotSpares field if non-nil, zero value otherwise.

### GetDedicatedHotSparesOk

`func (o *StorageDiskGroup) GetDedicatedHotSparesOk() (*[]StoragePhysicalDiskRelationship, bool)`

GetDedicatedHotSparesOk returns a tuple with the DedicatedHotSpares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDedicatedHotSpares

`func (o *StorageDiskGroup) SetDedicatedHotSpares(v []StoragePhysicalDiskRelationship)`

SetDedicatedHotSpares sets DedicatedHotSpares field to given value.

### HasDedicatedHotSpares

`func (o *StorageDiskGroup) HasDedicatedHotSpares() bool`

HasDedicatedHotSpares returns a boolean if a field has been set.

### SetDedicatedHotSparesNil

`func (o *StorageDiskGroup) SetDedicatedHotSparesNil(b bool)`

 SetDedicatedHotSparesNil sets the value for DedicatedHotSpares to be an explicit nil

### UnsetDedicatedHotSpares
`func (o *StorageDiskGroup) UnsetDedicatedHotSpares()`

UnsetDedicatedHotSpares ensures that no value is present for DedicatedHotSpares, not even an explicit nil
### GetRegisteredDevice

`func (o *StorageDiskGroup) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StorageDiskGroup) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StorageDiskGroup) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StorageDiskGroup) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### GetSpans

`func (o *StorageDiskGroup) GetSpans() []StorageSpanRelationship`

GetSpans returns the Spans field if non-nil, zero value otherwise.

### GetSpansOk

`func (o *StorageDiskGroup) GetSpansOk() (*[]StorageSpanRelationship, bool)`

GetSpansOk returns a tuple with the Spans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpans

`func (o *StorageDiskGroup) SetSpans(v []StorageSpanRelationship)`

SetSpans sets Spans field to given value.

### HasSpans

`func (o *StorageDiskGroup) HasSpans() bool`

HasSpans returns a boolean if a field has been set.

### SetSpansNil

`func (o *StorageDiskGroup) SetSpansNil(b bool)`

 SetSpansNil sets the value for Spans to be an explicit nil

### UnsetSpans
`func (o *StorageDiskGroup) UnsetSpans()`

UnsetSpans ensures that no value is present for Spans, not even an explicit nil
### GetStorageController

`func (o *StorageDiskGroup) GetStorageController() StorageControllerRelationship`

GetStorageController returns the StorageController field if non-nil, zero value otherwise.

### GetStorageControllerOk

`func (o *StorageDiskGroup) GetStorageControllerOk() (*StorageControllerRelationship, bool)`

GetStorageControllerOk returns a tuple with the StorageController field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageController

`func (o *StorageDiskGroup) SetStorageController(v StorageControllerRelationship)`

SetStorageController sets StorageController field to given value.

### HasStorageController

`func (o *StorageDiskGroup) HasStorageController() bool`

HasStorageController returns a boolean if a field has been set.

### GetVirtualDrives

`func (o *StorageDiskGroup) GetVirtualDrives() []StorageVirtualDriveRelationship`

GetVirtualDrives returns the VirtualDrives field if non-nil, zero value otherwise.

### GetVirtualDrivesOk

`func (o *StorageDiskGroup) GetVirtualDrivesOk() (*[]StorageVirtualDriveRelationship, bool)`

GetVirtualDrivesOk returns a tuple with the VirtualDrives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualDrives

`func (o *StorageDiskGroup) SetVirtualDrives(v []StorageVirtualDriveRelationship)`

SetVirtualDrives sets VirtualDrives field to given value.

### HasVirtualDrives

`func (o *StorageDiskGroup) HasVirtualDrives() bool`

HasVirtualDrives returns a boolean if a field has been set.

### SetVirtualDrivesNil

`func (o *StorageDiskGroup) SetVirtualDrivesNil(b bool)`

 SetVirtualDrivesNil sets the value for VirtualDrives to be an explicit nil

### UnsetVirtualDrives
`func (o *StorageDiskGroup) UnsetVirtualDrives()`

UnsetVirtualDrives ensures that no value is present for VirtualDrives, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



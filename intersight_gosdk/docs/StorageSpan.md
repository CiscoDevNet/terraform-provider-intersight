# StorageSpan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Slots** | Pointer to **[]int64** |  | [optional] 
**SpanId** | Pointer to **int64** | Unique identifier value of this span. | [optional] 
**DiskGroup** | Pointer to [**StorageDiskGroupRelationship**](storage.DiskGroup.Relationship.md) |  | [optional] 
**PhysicalDisks** | Pointer to [**[]StoragePhysicalDiskRelationship**](storage.PhysicalDisk.Relationship.md) | An array of relationships to storagePhysicalDisk resources. | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewStorageSpan

`func NewStorageSpan() *StorageSpan`

NewStorageSpan instantiates a new StorageSpan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageSpanWithDefaults

`func NewStorageSpanWithDefaults() *StorageSpan`

NewStorageSpanWithDefaults instantiates a new StorageSpan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSlots

`func (o *StorageSpan) GetSlots() []int64`

GetSlots returns the Slots field if non-nil, zero value otherwise.

### GetSlotsOk

`func (o *StorageSpan) GetSlotsOk() (*[]int64, bool)`

GetSlotsOk returns a tuple with the Slots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlots

`func (o *StorageSpan) SetSlots(v []int64)`

SetSlots sets Slots field to given value.

### HasSlots

`func (o *StorageSpan) HasSlots() bool`

HasSlots returns a boolean if a field has been set.

### GetSpanId

`func (o *StorageSpan) GetSpanId() int64`

GetSpanId returns the SpanId field if non-nil, zero value otherwise.

### GetSpanIdOk

`func (o *StorageSpan) GetSpanIdOk() (*int64, bool)`

GetSpanIdOk returns a tuple with the SpanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanId

`func (o *StorageSpan) SetSpanId(v int64)`

SetSpanId sets SpanId field to given value.

### HasSpanId

`func (o *StorageSpan) HasSpanId() bool`

HasSpanId returns a boolean if a field has been set.

### GetDiskGroup

`func (o *StorageSpan) GetDiskGroup() StorageDiskGroupRelationship`

GetDiskGroup returns the DiskGroup field if non-nil, zero value otherwise.

### GetDiskGroupOk

`func (o *StorageSpan) GetDiskGroupOk() (*StorageDiskGroupRelationship, bool)`

GetDiskGroupOk returns a tuple with the DiskGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskGroup

`func (o *StorageSpan) SetDiskGroup(v StorageDiskGroupRelationship)`

SetDiskGroup sets DiskGroup field to given value.

### HasDiskGroup

`func (o *StorageSpan) HasDiskGroup() bool`

HasDiskGroup returns a boolean if a field has been set.

### GetPhysicalDisks

`func (o *StorageSpan) GetPhysicalDisks() []StoragePhysicalDiskRelationship`

GetPhysicalDisks returns the PhysicalDisks field if non-nil, zero value otherwise.

### GetPhysicalDisksOk

`func (o *StorageSpan) GetPhysicalDisksOk() (*[]StoragePhysicalDiskRelationship, bool)`

GetPhysicalDisksOk returns a tuple with the PhysicalDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalDisks

`func (o *StorageSpan) SetPhysicalDisks(v []StoragePhysicalDiskRelationship)`

SetPhysicalDisks sets PhysicalDisks field to given value.

### HasPhysicalDisks

`func (o *StorageSpan) HasPhysicalDisks() bool`

HasPhysicalDisks returns a boolean if a field has been set.

### SetPhysicalDisksNil

`func (o *StorageSpan) SetPhysicalDisksNil(b bool)`

 SetPhysicalDisksNil sets the value for PhysicalDisks to be an explicit nil

### UnsetPhysicalDisks
`func (o *StorageSpan) UnsetPhysicalDisks()`

UnsetPhysicalDisks ensures that no value is present for PhysicalDisks, not even an explicit nil
### GetRegisteredDevice

`func (o *StorageSpan) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StorageSpan) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StorageSpan) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StorageSpan) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



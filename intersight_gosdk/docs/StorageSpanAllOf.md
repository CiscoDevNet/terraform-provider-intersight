# StorageSpanAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.Span"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.Span"]
**Slots** | Pointer to **[]int64** |  | [optional] 
**SpanId** | Pointer to **int64** | Unique identifier value of this span. | [optional] 
**DiskGroup** | Pointer to [**StorageDiskGroupRelationship**](StorageDiskGroupRelationship.md) |  | [optional] 
**PhysicalDisks** | Pointer to [**[]StoragePhysicalDiskRelationship**](StoragePhysicalDiskRelationship.md) | An array of relationships to storagePhysicalDisk resources. | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewStorageSpanAllOf

`func NewStorageSpanAllOf(classId string, objectType string, ) *StorageSpanAllOf`

NewStorageSpanAllOf instantiates a new StorageSpanAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageSpanAllOfWithDefaults

`func NewStorageSpanAllOfWithDefaults() *StorageSpanAllOf`

NewStorageSpanAllOfWithDefaults instantiates a new StorageSpanAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageSpanAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageSpanAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageSpanAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageSpanAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageSpanAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageSpanAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetSlots

`func (o *StorageSpanAllOf) GetSlots() []int64`

GetSlots returns the Slots field if non-nil, zero value otherwise.

### GetSlotsOk

`func (o *StorageSpanAllOf) GetSlotsOk() (*[]int64, bool)`

GetSlotsOk returns a tuple with the Slots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlots

`func (o *StorageSpanAllOf) SetSlots(v []int64)`

SetSlots sets Slots field to given value.

### HasSlots

`func (o *StorageSpanAllOf) HasSlots() bool`

HasSlots returns a boolean if a field has been set.

### SetSlotsNil

`func (o *StorageSpanAllOf) SetSlotsNil(b bool)`

 SetSlotsNil sets the value for Slots to be an explicit nil

### UnsetSlots
`func (o *StorageSpanAllOf) UnsetSlots()`

UnsetSlots ensures that no value is present for Slots, not even an explicit nil
### GetSpanId

`func (o *StorageSpanAllOf) GetSpanId() int64`

GetSpanId returns the SpanId field if non-nil, zero value otherwise.

### GetSpanIdOk

`func (o *StorageSpanAllOf) GetSpanIdOk() (*int64, bool)`

GetSpanIdOk returns a tuple with the SpanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanId

`func (o *StorageSpanAllOf) SetSpanId(v int64)`

SetSpanId sets SpanId field to given value.

### HasSpanId

`func (o *StorageSpanAllOf) HasSpanId() bool`

HasSpanId returns a boolean if a field has been set.

### GetDiskGroup

`func (o *StorageSpanAllOf) GetDiskGroup() StorageDiskGroupRelationship`

GetDiskGroup returns the DiskGroup field if non-nil, zero value otherwise.

### GetDiskGroupOk

`func (o *StorageSpanAllOf) GetDiskGroupOk() (*StorageDiskGroupRelationship, bool)`

GetDiskGroupOk returns a tuple with the DiskGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskGroup

`func (o *StorageSpanAllOf) SetDiskGroup(v StorageDiskGroupRelationship)`

SetDiskGroup sets DiskGroup field to given value.

### HasDiskGroup

`func (o *StorageSpanAllOf) HasDiskGroup() bool`

HasDiskGroup returns a boolean if a field has been set.

### GetPhysicalDisks

`func (o *StorageSpanAllOf) GetPhysicalDisks() []StoragePhysicalDiskRelationship`

GetPhysicalDisks returns the PhysicalDisks field if non-nil, zero value otherwise.

### GetPhysicalDisksOk

`func (o *StorageSpanAllOf) GetPhysicalDisksOk() (*[]StoragePhysicalDiskRelationship, bool)`

GetPhysicalDisksOk returns a tuple with the PhysicalDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalDisks

`func (o *StorageSpanAllOf) SetPhysicalDisks(v []StoragePhysicalDiskRelationship)`

SetPhysicalDisks sets PhysicalDisks field to given value.

### HasPhysicalDisks

`func (o *StorageSpanAllOf) HasPhysicalDisks() bool`

HasPhysicalDisks returns a boolean if a field has been set.

### SetPhysicalDisksNil

`func (o *StorageSpanAllOf) SetPhysicalDisksNil(b bool)`

 SetPhysicalDisksNil sets the value for PhysicalDisks to be an explicit nil

### UnsetPhysicalDisks
`func (o *StorageSpanAllOf) UnsetPhysicalDisks()`

UnsetPhysicalDisks ensures that no value is present for PhysicalDisks, not even an explicit nil
### GetRegisteredDevice

`func (o *StorageSpanAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StorageSpanAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StorageSpanAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StorageSpanAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



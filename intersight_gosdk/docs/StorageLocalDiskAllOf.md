# StorageLocalDiskAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SlotNumber** | Pointer to **int64** | The slot number of the disk to be referenced. As this is a policy, this slot number may or may not be valid depending on the number of disks in the associated server. | [optional] 

## Methods

### NewStorageLocalDiskAllOf

`func NewStorageLocalDiskAllOf() *StorageLocalDiskAllOf`

NewStorageLocalDiskAllOf instantiates a new StorageLocalDiskAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageLocalDiskAllOfWithDefaults

`func NewStorageLocalDiskAllOfWithDefaults() *StorageLocalDiskAllOf`

NewStorageLocalDiskAllOfWithDefaults instantiates a new StorageLocalDiskAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSlotNumber

`func (o *StorageLocalDiskAllOf) GetSlotNumber() int64`

GetSlotNumber returns the SlotNumber field if non-nil, zero value otherwise.

### GetSlotNumberOk

`func (o *StorageLocalDiskAllOf) GetSlotNumberOk() (*int64, bool)`

GetSlotNumberOk returns a tuple with the SlotNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotNumber

`func (o *StorageLocalDiskAllOf) SetSlotNumber(v int64)`

SetSlotNumber sets SlotNumber field to given value.

### HasSlotNumber

`func (o *StorageLocalDiskAllOf) HasSlotNumber() bool`

HasSlotNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



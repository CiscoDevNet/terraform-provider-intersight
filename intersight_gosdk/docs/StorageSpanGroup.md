# StorageSpanGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Disks** | Pointer to [**[]StorageLocalDisk**](storage.LocalDisk.md) |  | [optional] 

## Methods

### NewStorageSpanGroup

`func NewStorageSpanGroup() *StorageSpanGroup`

NewStorageSpanGroup instantiates a new StorageSpanGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageSpanGroupWithDefaults

`func NewStorageSpanGroupWithDefaults() *StorageSpanGroup`

NewStorageSpanGroupWithDefaults instantiates a new StorageSpanGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisks

`func (o *StorageSpanGroup) GetDisks() []StorageLocalDisk`

GetDisks returns the Disks field if non-nil, zero value otherwise.

### GetDisksOk

`func (o *StorageSpanGroup) GetDisksOk() (*[]StorageLocalDisk, bool)`

GetDisksOk returns a tuple with the Disks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisks

`func (o *StorageSpanGroup) SetDisks(v []StorageLocalDisk)`

SetDisks sets Disks field to given value.

### HasDisks

`func (o *StorageSpanGroup) HasDisks() bool`

HasDisks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



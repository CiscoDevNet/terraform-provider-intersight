# StorageBaseProtectionGroupSnapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedTime** | Pointer to [**time.Time**](time.Time.md) | Protection group snapshot creation time. | [optional] [readonly] 
**Name** | Pointer to **string** | Protection group snapshot name which represents point-in-time copy of all members in protection group. | [optional] [readonly] 
**Size** | Pointer to **int64** | Snapshot size represented in bytes. It is a cumulative size of all snapshots in a set. | [optional] [readonly] 
**Source** | Pointer to **string** | Source protection group name on which the snapshot is created. | [optional] [readonly] 

## Methods

### NewStorageBaseProtectionGroupSnapshot

`func NewStorageBaseProtectionGroupSnapshot() *StorageBaseProtectionGroupSnapshot`

NewStorageBaseProtectionGroupSnapshot instantiates a new StorageBaseProtectionGroupSnapshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageBaseProtectionGroupSnapshotWithDefaults

`func NewStorageBaseProtectionGroupSnapshotWithDefaults() *StorageBaseProtectionGroupSnapshot`

NewStorageBaseProtectionGroupSnapshotWithDefaults instantiates a new StorageBaseProtectionGroupSnapshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedTime

`func (o *StorageBaseProtectionGroupSnapshot) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *StorageBaseProtectionGroupSnapshot) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *StorageBaseProtectionGroupSnapshot) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *StorageBaseProtectionGroupSnapshot) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetName

`func (o *StorageBaseProtectionGroupSnapshot) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageBaseProtectionGroupSnapshot) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageBaseProtectionGroupSnapshot) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StorageBaseProtectionGroupSnapshot) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSize

`func (o *StorageBaseProtectionGroupSnapshot) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *StorageBaseProtectionGroupSnapshot) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *StorageBaseProtectionGroupSnapshot) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *StorageBaseProtectionGroupSnapshot) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSource

`func (o *StorageBaseProtectionGroupSnapshot) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *StorageBaseProtectionGroupSnapshot) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *StorageBaseProtectionGroupSnapshot) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *StorageBaseProtectionGroupSnapshot) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



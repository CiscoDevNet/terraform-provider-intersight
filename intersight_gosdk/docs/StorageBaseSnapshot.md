# StorageBaseSnapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedTime** | Pointer to [**time.Time**](time.Time.md) | Exact date and time at which snapshot was created. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the snapshot which represents point-in-time copy of volume. | [optional] [readonly] 
**ProtectionGroupName** | Pointer to **string** | Name of the protection group to which the snapshot belongs. Value is empty, if the snapshot is created directly on volume. | [optional] [readonly] 
**Size** | Pointer to **int64** | Snapshot size represented in bytes. | [optional] [readonly] 
**Source** | Pointer to **string** | Source object on which the snapshot is created. It is the name of the originating volume. | [optional] [readonly] 

## Methods

### NewStorageBaseSnapshot

`func NewStorageBaseSnapshot() *StorageBaseSnapshot`

NewStorageBaseSnapshot instantiates a new StorageBaseSnapshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageBaseSnapshotWithDefaults

`func NewStorageBaseSnapshotWithDefaults() *StorageBaseSnapshot`

NewStorageBaseSnapshotWithDefaults instantiates a new StorageBaseSnapshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedTime

`func (o *StorageBaseSnapshot) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *StorageBaseSnapshot) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *StorageBaseSnapshot) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *StorageBaseSnapshot) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetName

`func (o *StorageBaseSnapshot) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageBaseSnapshot) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageBaseSnapshot) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StorageBaseSnapshot) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProtectionGroupName

`func (o *StorageBaseSnapshot) GetProtectionGroupName() string`

GetProtectionGroupName returns the ProtectionGroupName field if non-nil, zero value otherwise.

### GetProtectionGroupNameOk

`func (o *StorageBaseSnapshot) GetProtectionGroupNameOk() (*string, bool)`

GetProtectionGroupNameOk returns a tuple with the ProtectionGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroupName

`func (o *StorageBaseSnapshot) SetProtectionGroupName(v string)`

SetProtectionGroupName sets ProtectionGroupName field to given value.

### HasProtectionGroupName

`func (o *StorageBaseSnapshot) HasProtectionGroupName() bool`

HasProtectionGroupName returns a boolean if a field has been set.

### GetSize

`func (o *StorageBaseSnapshot) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *StorageBaseSnapshot) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *StorageBaseSnapshot) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *StorageBaseSnapshot) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSource

`func (o *StorageBaseSnapshot) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *StorageBaseSnapshot) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *StorageBaseSnapshot) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *StorageBaseSnapshot) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



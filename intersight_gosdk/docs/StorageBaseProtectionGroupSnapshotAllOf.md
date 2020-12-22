# StorageBaseProtectionGroupSnapshotAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "storage.PureProtectionGroupSnapshot"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "storage.PureProtectionGroupSnapshot"]
**CreatedTime** | Pointer to **time.Time** | Protection group snapshot creation time. | [optional] [readonly] 
**Name** | Pointer to **string** | Protection group snapshot name which represents point-in-time copy of all members in protection group. | [optional] [readonly] 
**Size** | Pointer to **int64** | Snapshot size represented in bytes. It is a cumulative size of all snapshots in a set. | [optional] [readonly] 
**Source** | Pointer to **string** | Source protection group name on which the snapshot is created. | [optional] [readonly] 

## Methods

### NewStorageBaseProtectionGroupSnapshotAllOf

`func NewStorageBaseProtectionGroupSnapshotAllOf(classId string, objectType string, ) *StorageBaseProtectionGroupSnapshotAllOf`

NewStorageBaseProtectionGroupSnapshotAllOf instantiates a new StorageBaseProtectionGroupSnapshotAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageBaseProtectionGroupSnapshotAllOfWithDefaults

`func NewStorageBaseProtectionGroupSnapshotAllOfWithDefaults() *StorageBaseProtectionGroupSnapshotAllOf`

NewStorageBaseProtectionGroupSnapshotAllOfWithDefaults instantiates a new StorageBaseProtectionGroupSnapshotAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageBaseProtectionGroupSnapshotAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageBaseProtectionGroupSnapshotAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageBaseProtectionGroupSnapshotAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageBaseProtectionGroupSnapshotAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageBaseProtectionGroupSnapshotAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageBaseProtectionGroupSnapshotAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCreatedTime

`func (o *StorageBaseProtectionGroupSnapshotAllOf) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *StorageBaseProtectionGroupSnapshotAllOf) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *StorageBaseProtectionGroupSnapshotAllOf) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *StorageBaseProtectionGroupSnapshotAllOf) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetName

`func (o *StorageBaseProtectionGroupSnapshotAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageBaseProtectionGroupSnapshotAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageBaseProtectionGroupSnapshotAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StorageBaseProtectionGroupSnapshotAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSize

`func (o *StorageBaseProtectionGroupSnapshotAllOf) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *StorageBaseProtectionGroupSnapshotAllOf) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *StorageBaseProtectionGroupSnapshotAllOf) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *StorageBaseProtectionGroupSnapshotAllOf) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSource

`func (o *StorageBaseProtectionGroupSnapshotAllOf) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *StorageBaseProtectionGroupSnapshotAllOf) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *StorageBaseProtectionGroupSnapshotAllOf) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *StorageBaseProtectionGroupSnapshotAllOf) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# StorageBaseSnapshotAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**CreatedTime** | Pointer to **time.Time** | Exact date and time at which snapshot was created. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the snapshot which represents point-in-time copy of volume. | [optional] [readonly] 
**ProtectionGroupName** | Pointer to **string** | Name of the protection group to which the snapshot belongs. Value is empty, if the snapshot is created directly on volume. | [optional] [readonly] 
**Size** | Pointer to **int64** | Snapshot size represented in bytes. | [optional] [readonly] 
**Source** | Pointer to **string** | Source object on which the snapshot is created. It is the name of the originating volume. | [optional] [readonly] 

## Methods

### NewStorageBaseSnapshotAllOf

`func NewStorageBaseSnapshotAllOf(classId string, objectType string, ) *StorageBaseSnapshotAllOf`

NewStorageBaseSnapshotAllOf instantiates a new StorageBaseSnapshotAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageBaseSnapshotAllOfWithDefaults

`func NewStorageBaseSnapshotAllOfWithDefaults() *StorageBaseSnapshotAllOf`

NewStorageBaseSnapshotAllOfWithDefaults instantiates a new StorageBaseSnapshotAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageBaseSnapshotAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageBaseSnapshotAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageBaseSnapshotAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageBaseSnapshotAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageBaseSnapshotAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageBaseSnapshotAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCreatedTime

`func (o *StorageBaseSnapshotAllOf) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *StorageBaseSnapshotAllOf) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *StorageBaseSnapshotAllOf) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *StorageBaseSnapshotAllOf) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetName

`func (o *StorageBaseSnapshotAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageBaseSnapshotAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageBaseSnapshotAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StorageBaseSnapshotAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProtectionGroupName

`func (o *StorageBaseSnapshotAllOf) GetProtectionGroupName() string`

GetProtectionGroupName returns the ProtectionGroupName field if non-nil, zero value otherwise.

### GetProtectionGroupNameOk

`func (o *StorageBaseSnapshotAllOf) GetProtectionGroupNameOk() (*string, bool)`

GetProtectionGroupNameOk returns a tuple with the ProtectionGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroupName

`func (o *StorageBaseSnapshotAllOf) SetProtectionGroupName(v string)`

SetProtectionGroupName sets ProtectionGroupName field to given value.

### HasProtectionGroupName

`func (o *StorageBaseSnapshotAllOf) HasProtectionGroupName() bool`

HasProtectionGroupName returns a boolean if a field has been set.

### GetSize

`func (o *StorageBaseSnapshotAllOf) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *StorageBaseSnapshotAllOf) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *StorageBaseSnapshotAllOf) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *StorageBaseSnapshotAllOf) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSource

`func (o *StorageBaseSnapshotAllOf) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *StorageBaseSnapshotAllOf) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *StorageBaseSnapshotAllOf) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *StorageBaseSnapshotAllOf) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



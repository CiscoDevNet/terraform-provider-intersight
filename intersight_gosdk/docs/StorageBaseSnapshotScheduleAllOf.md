# StorageBaseSnapshotScheduleAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "storage.PureSnapshotSchedule"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "storage.PureSnapshotSchedule"]
**Frequency** | Pointer to **string** | Snapshot frequency. It is an interval at which snapshot is set to trigger on source array. Examples:     PT2H Snapshot is generated every 2 hours.     P4D, Snapshot is scheduled for every 4 days.     PT2H34M56.123S is 2 hours, 34 minutes, 56 seconds and 123 milliseconds. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the snapshot schedule. | [optional] 
**RetentionTime** | Pointer to **string** | Duration to keep the snapshots on the source array. Once this period expires, system deletes the snapshot automatically from the source array. Examples: P200D,  200 days. PT2H34M56.123S, 2 hours, 34 minutes, 56 seconds and 123 milliseconds. | [optional] [readonly] 
**SnapshotTime** | Pointer to **string** | Preferred time of the day to capture the snapshot. It is applicable only if the frequency is set for a day or more. Format: hh:mm:ss Example: 08:30:00, Snapshot is set for 08:30 AM. | [optional] [readonly] 

## Methods

### NewStorageBaseSnapshotScheduleAllOf

`func NewStorageBaseSnapshotScheduleAllOf(classId string, objectType string, ) *StorageBaseSnapshotScheduleAllOf`

NewStorageBaseSnapshotScheduleAllOf instantiates a new StorageBaseSnapshotScheduleAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageBaseSnapshotScheduleAllOfWithDefaults

`func NewStorageBaseSnapshotScheduleAllOfWithDefaults() *StorageBaseSnapshotScheduleAllOf`

NewStorageBaseSnapshotScheduleAllOfWithDefaults instantiates a new StorageBaseSnapshotScheduleAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageBaseSnapshotScheduleAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageBaseSnapshotScheduleAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageBaseSnapshotScheduleAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageBaseSnapshotScheduleAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageBaseSnapshotScheduleAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageBaseSnapshotScheduleAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetFrequency

`func (o *StorageBaseSnapshotScheduleAllOf) GetFrequency() string`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *StorageBaseSnapshotScheduleAllOf) GetFrequencyOk() (*string, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *StorageBaseSnapshotScheduleAllOf) SetFrequency(v string)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *StorageBaseSnapshotScheduleAllOf) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### GetName

`func (o *StorageBaseSnapshotScheduleAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageBaseSnapshotScheduleAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageBaseSnapshotScheduleAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StorageBaseSnapshotScheduleAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRetentionTime

`func (o *StorageBaseSnapshotScheduleAllOf) GetRetentionTime() string`

GetRetentionTime returns the RetentionTime field if non-nil, zero value otherwise.

### GetRetentionTimeOk

`func (o *StorageBaseSnapshotScheduleAllOf) GetRetentionTimeOk() (*string, bool)`

GetRetentionTimeOk returns a tuple with the RetentionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionTime

`func (o *StorageBaseSnapshotScheduleAllOf) SetRetentionTime(v string)`

SetRetentionTime sets RetentionTime field to given value.

### HasRetentionTime

`func (o *StorageBaseSnapshotScheduleAllOf) HasRetentionTime() bool`

HasRetentionTime returns a boolean if a field has been set.

### GetSnapshotTime

`func (o *StorageBaseSnapshotScheduleAllOf) GetSnapshotTime() string`

GetSnapshotTime returns the SnapshotTime field if non-nil, zero value otherwise.

### GetSnapshotTimeOk

`func (o *StorageBaseSnapshotScheduleAllOf) GetSnapshotTimeOk() (*string, bool)`

GetSnapshotTimeOk returns a tuple with the SnapshotTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotTime

`func (o *StorageBaseSnapshotScheduleAllOf) SetSnapshotTime(v string)`

SetSnapshotTime sets SnapshotTime field to given value.

### HasSnapshotTime

`func (o *StorageBaseSnapshotScheduleAllOf) HasSnapshotTime() bool`

HasSnapshotTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



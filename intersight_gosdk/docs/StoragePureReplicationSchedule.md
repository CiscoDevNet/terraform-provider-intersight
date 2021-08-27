# StoragePureReplicationSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.PureReplicationSchedule"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.PureReplicationSchedule"]
**DailyLimit** | Pointer to **int64** | Total number of snapshots per day to be available on target above and over the specified retention time. PureStorage FlashArray maintains all created snapshot until retention period. Daily limit is applied only on the snapshots once retention time is expired. In case of, daily limit is less than the number of snapshot available on source, system select snapshots evenly spaced out throughout the day. | [optional] [readonly] 
**ReplicationBlackoutIntervals** | Pointer to [**[]StoragePureReplicationBlackout**](StoragePureReplicationBlackout.md) |  | [optional] 
**SnapshotExpiryTime** | Pointer to **string** | Duration to keep the daily limit snapshots on target array. StorageArray deletes the snapshots permanently from the targets beyond this period. | [optional] [readonly] 
**Array** | Pointer to [**StoragePureArrayRelationship**](StoragePureArrayRelationship.md) |  | [optional] 
**ProtectionGroup** | Pointer to [**StoragePureProtectionGroupRelationship**](StoragePureProtectionGroupRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewStoragePureReplicationSchedule

`func NewStoragePureReplicationSchedule(classId string, objectType string, ) *StoragePureReplicationSchedule`

NewStoragePureReplicationSchedule instantiates a new StoragePureReplicationSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoragePureReplicationScheduleWithDefaults

`func NewStoragePureReplicationScheduleWithDefaults() *StoragePureReplicationSchedule`

NewStoragePureReplicationScheduleWithDefaults instantiates a new StoragePureReplicationSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StoragePureReplicationSchedule) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StoragePureReplicationSchedule) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StoragePureReplicationSchedule) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StoragePureReplicationSchedule) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StoragePureReplicationSchedule) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StoragePureReplicationSchedule) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDailyLimit

`func (o *StoragePureReplicationSchedule) GetDailyLimit() int64`

GetDailyLimit returns the DailyLimit field if non-nil, zero value otherwise.

### GetDailyLimitOk

`func (o *StoragePureReplicationSchedule) GetDailyLimitOk() (*int64, bool)`

GetDailyLimitOk returns a tuple with the DailyLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyLimit

`func (o *StoragePureReplicationSchedule) SetDailyLimit(v int64)`

SetDailyLimit sets DailyLimit field to given value.

### HasDailyLimit

`func (o *StoragePureReplicationSchedule) HasDailyLimit() bool`

HasDailyLimit returns a boolean if a field has been set.

### GetReplicationBlackoutIntervals

`func (o *StoragePureReplicationSchedule) GetReplicationBlackoutIntervals() []StoragePureReplicationBlackout`

GetReplicationBlackoutIntervals returns the ReplicationBlackoutIntervals field if non-nil, zero value otherwise.

### GetReplicationBlackoutIntervalsOk

`func (o *StoragePureReplicationSchedule) GetReplicationBlackoutIntervalsOk() (*[]StoragePureReplicationBlackout, bool)`

GetReplicationBlackoutIntervalsOk returns a tuple with the ReplicationBlackoutIntervals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationBlackoutIntervals

`func (o *StoragePureReplicationSchedule) SetReplicationBlackoutIntervals(v []StoragePureReplicationBlackout)`

SetReplicationBlackoutIntervals sets ReplicationBlackoutIntervals field to given value.

### HasReplicationBlackoutIntervals

`func (o *StoragePureReplicationSchedule) HasReplicationBlackoutIntervals() bool`

HasReplicationBlackoutIntervals returns a boolean if a field has been set.

### SetReplicationBlackoutIntervalsNil

`func (o *StoragePureReplicationSchedule) SetReplicationBlackoutIntervalsNil(b bool)`

 SetReplicationBlackoutIntervalsNil sets the value for ReplicationBlackoutIntervals to be an explicit nil

### UnsetReplicationBlackoutIntervals
`func (o *StoragePureReplicationSchedule) UnsetReplicationBlackoutIntervals()`

UnsetReplicationBlackoutIntervals ensures that no value is present for ReplicationBlackoutIntervals, not even an explicit nil
### GetSnapshotExpiryTime

`func (o *StoragePureReplicationSchedule) GetSnapshotExpiryTime() string`

GetSnapshotExpiryTime returns the SnapshotExpiryTime field if non-nil, zero value otherwise.

### GetSnapshotExpiryTimeOk

`func (o *StoragePureReplicationSchedule) GetSnapshotExpiryTimeOk() (*string, bool)`

GetSnapshotExpiryTimeOk returns a tuple with the SnapshotExpiryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotExpiryTime

`func (o *StoragePureReplicationSchedule) SetSnapshotExpiryTime(v string)`

SetSnapshotExpiryTime sets SnapshotExpiryTime field to given value.

### HasSnapshotExpiryTime

`func (o *StoragePureReplicationSchedule) HasSnapshotExpiryTime() bool`

HasSnapshotExpiryTime returns a boolean if a field has been set.

### GetArray

`func (o *StoragePureReplicationSchedule) GetArray() StoragePureArrayRelationship`

GetArray returns the Array field if non-nil, zero value otherwise.

### GetArrayOk

`func (o *StoragePureReplicationSchedule) GetArrayOk() (*StoragePureArrayRelationship, bool)`

GetArrayOk returns a tuple with the Array field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArray

`func (o *StoragePureReplicationSchedule) SetArray(v StoragePureArrayRelationship)`

SetArray sets Array field to given value.

### HasArray

`func (o *StoragePureReplicationSchedule) HasArray() bool`

HasArray returns a boolean if a field has been set.

### GetProtectionGroup

`func (o *StoragePureReplicationSchedule) GetProtectionGroup() StoragePureProtectionGroupRelationship`

GetProtectionGroup returns the ProtectionGroup field if non-nil, zero value otherwise.

### GetProtectionGroupOk

`func (o *StoragePureReplicationSchedule) GetProtectionGroupOk() (*StoragePureProtectionGroupRelationship, bool)`

GetProtectionGroupOk returns a tuple with the ProtectionGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroup

`func (o *StoragePureReplicationSchedule) SetProtectionGroup(v StoragePureProtectionGroupRelationship)`

SetProtectionGroup sets ProtectionGroup field to given value.

### HasProtectionGroup

`func (o *StoragePureReplicationSchedule) HasProtectionGroup() bool`

HasProtectionGroup returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *StoragePureReplicationSchedule) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StoragePureReplicationSchedule) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StoragePureReplicationSchedule) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StoragePureReplicationSchedule) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



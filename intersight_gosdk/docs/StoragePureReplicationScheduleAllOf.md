# StoragePureReplicationScheduleAllOf

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

### NewStoragePureReplicationScheduleAllOf

`func NewStoragePureReplicationScheduleAllOf(classId string, objectType string, ) *StoragePureReplicationScheduleAllOf`

NewStoragePureReplicationScheduleAllOf instantiates a new StoragePureReplicationScheduleAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoragePureReplicationScheduleAllOfWithDefaults

`func NewStoragePureReplicationScheduleAllOfWithDefaults() *StoragePureReplicationScheduleAllOf`

NewStoragePureReplicationScheduleAllOfWithDefaults instantiates a new StoragePureReplicationScheduleAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StoragePureReplicationScheduleAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StoragePureReplicationScheduleAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StoragePureReplicationScheduleAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StoragePureReplicationScheduleAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StoragePureReplicationScheduleAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StoragePureReplicationScheduleAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDailyLimit

`func (o *StoragePureReplicationScheduleAllOf) GetDailyLimit() int64`

GetDailyLimit returns the DailyLimit field if non-nil, zero value otherwise.

### GetDailyLimitOk

`func (o *StoragePureReplicationScheduleAllOf) GetDailyLimitOk() (*int64, bool)`

GetDailyLimitOk returns a tuple with the DailyLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyLimit

`func (o *StoragePureReplicationScheduleAllOf) SetDailyLimit(v int64)`

SetDailyLimit sets DailyLimit field to given value.

### HasDailyLimit

`func (o *StoragePureReplicationScheduleAllOf) HasDailyLimit() bool`

HasDailyLimit returns a boolean if a field has been set.

### GetReplicationBlackoutIntervals

`func (o *StoragePureReplicationScheduleAllOf) GetReplicationBlackoutIntervals() []StoragePureReplicationBlackout`

GetReplicationBlackoutIntervals returns the ReplicationBlackoutIntervals field if non-nil, zero value otherwise.

### GetReplicationBlackoutIntervalsOk

`func (o *StoragePureReplicationScheduleAllOf) GetReplicationBlackoutIntervalsOk() (*[]StoragePureReplicationBlackout, bool)`

GetReplicationBlackoutIntervalsOk returns a tuple with the ReplicationBlackoutIntervals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationBlackoutIntervals

`func (o *StoragePureReplicationScheduleAllOf) SetReplicationBlackoutIntervals(v []StoragePureReplicationBlackout)`

SetReplicationBlackoutIntervals sets ReplicationBlackoutIntervals field to given value.

### HasReplicationBlackoutIntervals

`func (o *StoragePureReplicationScheduleAllOf) HasReplicationBlackoutIntervals() bool`

HasReplicationBlackoutIntervals returns a boolean if a field has been set.

### SetReplicationBlackoutIntervalsNil

`func (o *StoragePureReplicationScheduleAllOf) SetReplicationBlackoutIntervalsNil(b bool)`

 SetReplicationBlackoutIntervalsNil sets the value for ReplicationBlackoutIntervals to be an explicit nil

### UnsetReplicationBlackoutIntervals
`func (o *StoragePureReplicationScheduleAllOf) UnsetReplicationBlackoutIntervals()`

UnsetReplicationBlackoutIntervals ensures that no value is present for ReplicationBlackoutIntervals, not even an explicit nil
### GetSnapshotExpiryTime

`func (o *StoragePureReplicationScheduleAllOf) GetSnapshotExpiryTime() string`

GetSnapshotExpiryTime returns the SnapshotExpiryTime field if non-nil, zero value otherwise.

### GetSnapshotExpiryTimeOk

`func (o *StoragePureReplicationScheduleAllOf) GetSnapshotExpiryTimeOk() (*string, bool)`

GetSnapshotExpiryTimeOk returns a tuple with the SnapshotExpiryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotExpiryTime

`func (o *StoragePureReplicationScheduleAllOf) SetSnapshotExpiryTime(v string)`

SetSnapshotExpiryTime sets SnapshotExpiryTime field to given value.

### HasSnapshotExpiryTime

`func (o *StoragePureReplicationScheduleAllOf) HasSnapshotExpiryTime() bool`

HasSnapshotExpiryTime returns a boolean if a field has been set.

### GetArray

`func (o *StoragePureReplicationScheduleAllOf) GetArray() StoragePureArrayRelationship`

GetArray returns the Array field if non-nil, zero value otherwise.

### GetArrayOk

`func (o *StoragePureReplicationScheduleAllOf) GetArrayOk() (*StoragePureArrayRelationship, bool)`

GetArrayOk returns a tuple with the Array field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArray

`func (o *StoragePureReplicationScheduleAllOf) SetArray(v StoragePureArrayRelationship)`

SetArray sets Array field to given value.

### HasArray

`func (o *StoragePureReplicationScheduleAllOf) HasArray() bool`

HasArray returns a boolean if a field has been set.

### GetProtectionGroup

`func (o *StoragePureReplicationScheduleAllOf) GetProtectionGroup() StoragePureProtectionGroupRelationship`

GetProtectionGroup returns the ProtectionGroup field if non-nil, zero value otherwise.

### GetProtectionGroupOk

`func (o *StoragePureReplicationScheduleAllOf) GetProtectionGroupOk() (*StoragePureProtectionGroupRelationship, bool)`

GetProtectionGroupOk returns a tuple with the ProtectionGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroup

`func (o *StoragePureReplicationScheduleAllOf) SetProtectionGroup(v StoragePureProtectionGroupRelationship)`

SetProtectionGroup sets ProtectionGroup field to given value.

### HasProtectionGroup

`func (o *StoragePureReplicationScheduleAllOf) HasProtectionGroup() bool`

HasProtectionGroup returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *StoragePureReplicationScheduleAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StoragePureReplicationScheduleAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StoragePureReplicationScheduleAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StoragePureReplicationScheduleAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



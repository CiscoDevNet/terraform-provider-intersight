# StoragePureSnapshotSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DailyLimit** | Pointer to **int64** | Total number of snapshots per day to be available on source above and over the specified retention time. PureStorage FlashArray maintains all created snapshot until retention period. Daily limit is applied only on the snapshots once retention time is expired. In case of, daily limit is less than the number of snapshot available on source, system select snapshots evenly spaced out throughout the day. | [optional] [readonly] 
**SnapshotExpiryTime** | Pointer to **string** | Duration to keep the daily limit snapshots on source array. StorageArray deletes the snapshots permanently from source beyond this period. | [optional] [readonly] 
**Array** | Pointer to [**StoragePureArrayRelationship**](storage.PureArray.Relationship.md) |  | [optional] 
**ProtectionGroup** | Pointer to [**StoragePureProtectionGroupRelationship**](storage.PureProtectionGroup.Relationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewStoragePureSnapshotSchedule

`func NewStoragePureSnapshotSchedule() *StoragePureSnapshotSchedule`

NewStoragePureSnapshotSchedule instantiates a new StoragePureSnapshotSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoragePureSnapshotScheduleWithDefaults

`func NewStoragePureSnapshotScheduleWithDefaults() *StoragePureSnapshotSchedule`

NewStoragePureSnapshotScheduleWithDefaults instantiates a new StoragePureSnapshotSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDailyLimit

`func (o *StoragePureSnapshotSchedule) GetDailyLimit() int64`

GetDailyLimit returns the DailyLimit field if non-nil, zero value otherwise.

### GetDailyLimitOk

`func (o *StoragePureSnapshotSchedule) GetDailyLimitOk() (*int64, bool)`

GetDailyLimitOk returns a tuple with the DailyLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyLimit

`func (o *StoragePureSnapshotSchedule) SetDailyLimit(v int64)`

SetDailyLimit sets DailyLimit field to given value.

### HasDailyLimit

`func (o *StoragePureSnapshotSchedule) HasDailyLimit() bool`

HasDailyLimit returns a boolean if a field has been set.

### GetSnapshotExpiryTime

`func (o *StoragePureSnapshotSchedule) GetSnapshotExpiryTime() string`

GetSnapshotExpiryTime returns the SnapshotExpiryTime field if non-nil, zero value otherwise.

### GetSnapshotExpiryTimeOk

`func (o *StoragePureSnapshotSchedule) GetSnapshotExpiryTimeOk() (*string, bool)`

GetSnapshotExpiryTimeOk returns a tuple with the SnapshotExpiryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotExpiryTime

`func (o *StoragePureSnapshotSchedule) SetSnapshotExpiryTime(v string)`

SetSnapshotExpiryTime sets SnapshotExpiryTime field to given value.

### HasSnapshotExpiryTime

`func (o *StoragePureSnapshotSchedule) HasSnapshotExpiryTime() bool`

HasSnapshotExpiryTime returns a boolean if a field has been set.

### GetArray

`func (o *StoragePureSnapshotSchedule) GetArray() StoragePureArrayRelationship`

GetArray returns the Array field if non-nil, zero value otherwise.

### GetArrayOk

`func (o *StoragePureSnapshotSchedule) GetArrayOk() (*StoragePureArrayRelationship, bool)`

GetArrayOk returns a tuple with the Array field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArray

`func (o *StoragePureSnapshotSchedule) SetArray(v StoragePureArrayRelationship)`

SetArray sets Array field to given value.

### HasArray

`func (o *StoragePureSnapshotSchedule) HasArray() bool`

HasArray returns a boolean if a field has been set.

### GetProtectionGroup

`func (o *StoragePureSnapshotSchedule) GetProtectionGroup() StoragePureProtectionGroupRelationship`

GetProtectionGroup returns the ProtectionGroup field if non-nil, zero value otherwise.

### GetProtectionGroupOk

`func (o *StoragePureSnapshotSchedule) GetProtectionGroupOk() (*StoragePureProtectionGroupRelationship, bool)`

GetProtectionGroupOk returns a tuple with the ProtectionGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroup

`func (o *StoragePureSnapshotSchedule) SetProtectionGroup(v StoragePureProtectionGroupRelationship)`

SetProtectionGroup sets ProtectionGroup field to given value.

### HasProtectionGroup

`func (o *StoragePureSnapshotSchedule) HasProtectionGroup() bool`

HasProtectionGroup returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *StoragePureSnapshotSchedule) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StoragePureSnapshotSchedule) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StoragePureSnapshotSchedule) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StoragePureSnapshotSchedule) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



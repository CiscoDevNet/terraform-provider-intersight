# StorageBaseReplicationScheduleAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Frequency** | Pointer to **string** | Replication frequency. It is an interval at which replication is set to trigger. Examples:     PT2H, Snapshot is generated every 2 hours.     P30D, Snapshot is scheduled for every 30 days.     PT2H34M56.123S is 2 hours, 34 minutes, 56 seconds and 123 milliseconds. | [optional] [readonly] 
**Name** | Pointer to **string** | Replication schedule name. | [optional] [readonly] 
**ReplicationTime** | Pointer to **string** | Preferred time of day at which to replicate the snapshots on target array. It is applicable only if the replication frequency is set for a day or more. Format: hh:mm:ss Example: 15:00:00, Replication is set for 3:00 PM. | [optional] [readonly] 
**RetentionTime** | Pointer to **string** | Duration to keep the replicated snapshots on the targets. Replicated snapshots are deleted from target array once the retention period has elapsed. Examples: P30D, Snapshots are available for 30 days. PT2H34M56.123S, 2 hours, 34 minutes, 56 seconds and 123 milliseconds. | [optional] [readonly] 

## Methods

### NewStorageBaseReplicationScheduleAllOf

`func NewStorageBaseReplicationScheduleAllOf() *StorageBaseReplicationScheduleAllOf`

NewStorageBaseReplicationScheduleAllOf instantiates a new StorageBaseReplicationScheduleAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageBaseReplicationScheduleAllOfWithDefaults

`func NewStorageBaseReplicationScheduleAllOfWithDefaults() *StorageBaseReplicationScheduleAllOf`

NewStorageBaseReplicationScheduleAllOfWithDefaults instantiates a new StorageBaseReplicationScheduleAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrequency

`func (o *StorageBaseReplicationScheduleAllOf) GetFrequency() string`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *StorageBaseReplicationScheduleAllOf) GetFrequencyOk() (*string, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *StorageBaseReplicationScheduleAllOf) SetFrequency(v string)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *StorageBaseReplicationScheduleAllOf) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### GetName

`func (o *StorageBaseReplicationScheduleAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageBaseReplicationScheduleAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageBaseReplicationScheduleAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StorageBaseReplicationScheduleAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReplicationTime

`func (o *StorageBaseReplicationScheduleAllOf) GetReplicationTime() string`

GetReplicationTime returns the ReplicationTime field if non-nil, zero value otherwise.

### GetReplicationTimeOk

`func (o *StorageBaseReplicationScheduleAllOf) GetReplicationTimeOk() (*string, bool)`

GetReplicationTimeOk returns a tuple with the ReplicationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationTime

`func (o *StorageBaseReplicationScheduleAllOf) SetReplicationTime(v string)`

SetReplicationTime sets ReplicationTime field to given value.

### HasReplicationTime

`func (o *StorageBaseReplicationScheduleAllOf) HasReplicationTime() bool`

HasReplicationTime returns a boolean if a field has been set.

### GetRetentionTime

`func (o *StorageBaseReplicationScheduleAllOf) GetRetentionTime() string`

GetRetentionTime returns the RetentionTime field if non-nil, zero value otherwise.

### GetRetentionTimeOk

`func (o *StorageBaseReplicationScheduleAllOf) GetRetentionTimeOk() (*string, bool)`

GetRetentionTimeOk returns a tuple with the RetentionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionTime

`func (o *StorageBaseReplicationScheduleAllOf) SetRetentionTime(v string)`

SetRetentionTime sets RetentionTime field to given value.

### HasRetentionTime

`func (o *StorageBaseReplicationScheduleAllOf) HasRetentionTime() bool`

HasRetentionTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



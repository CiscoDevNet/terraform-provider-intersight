# HyperflexBackupPolicySettingsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.BackupPolicySettings"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.BackupPolicySettings"]
**BackupDataStoreName** | Pointer to **string** | Backup datastore name used during auto creation of the datastore. All Virtual Machines created in this datastore will be automatically backed up. | [optional] [readonly] 
**BackupDataStoreSize** | Pointer to **int64** | Capacity of Backup source datastore. | [optional] [readonly] 
**BackupDataStoreSizeUnit** | Pointer to **string** | Unit of backupDataStoreSize.  Allowable values are \&quot;GB\&quot; (Gigabytes) or \&quot;TB\&quot; (Terabytes). * &#x60;GB&#x60; - BackupDataStoreSize is specified in Gigabytes. * &#x60;TB&#x60; - BackupDataStoreSize is specified in Terabytes. | [optional] [readonly] [default to "GB"]
**DataStoreEncryptionEnabled** | Pointer to **bool** | Whether the datastore is encrypted or not. | [optional] [readonly] 
**LocalSnapshotRetentionCount** | Pointer to **int64** | Number of snapshots that will be retained. | [optional] [readonly] 
**ReplicationIntervalInMinutes** | Pointer to **int64** | Snapshot replication interval. | [optional] [readonly] 
**ReplicationPairNamePrefix** | Pointer to **string** | Replication cluster pairing name prefix. | [optional] [readonly] 
**SnapshotRetentionCount** | Pointer to **int64** | Number of snapshots that will be retained. | [optional] [readonly] 

## Methods

### NewHyperflexBackupPolicySettingsAllOf

`func NewHyperflexBackupPolicySettingsAllOf(classId string, objectType string, ) *HyperflexBackupPolicySettingsAllOf`

NewHyperflexBackupPolicySettingsAllOf instantiates a new HyperflexBackupPolicySettingsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexBackupPolicySettingsAllOfWithDefaults

`func NewHyperflexBackupPolicySettingsAllOfWithDefaults() *HyperflexBackupPolicySettingsAllOf`

NewHyperflexBackupPolicySettingsAllOfWithDefaults instantiates a new HyperflexBackupPolicySettingsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexBackupPolicySettingsAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexBackupPolicySettingsAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexBackupPolicySettingsAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexBackupPolicySettingsAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexBackupPolicySettingsAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexBackupPolicySettingsAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBackupDataStoreName

`func (o *HyperflexBackupPolicySettingsAllOf) GetBackupDataStoreName() string`

GetBackupDataStoreName returns the BackupDataStoreName field if non-nil, zero value otherwise.

### GetBackupDataStoreNameOk

`func (o *HyperflexBackupPolicySettingsAllOf) GetBackupDataStoreNameOk() (*string, bool)`

GetBackupDataStoreNameOk returns a tuple with the BackupDataStoreName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupDataStoreName

`func (o *HyperflexBackupPolicySettingsAllOf) SetBackupDataStoreName(v string)`

SetBackupDataStoreName sets BackupDataStoreName field to given value.

### HasBackupDataStoreName

`func (o *HyperflexBackupPolicySettingsAllOf) HasBackupDataStoreName() bool`

HasBackupDataStoreName returns a boolean if a field has been set.

### GetBackupDataStoreSize

`func (o *HyperflexBackupPolicySettingsAllOf) GetBackupDataStoreSize() int64`

GetBackupDataStoreSize returns the BackupDataStoreSize field if non-nil, zero value otherwise.

### GetBackupDataStoreSizeOk

`func (o *HyperflexBackupPolicySettingsAllOf) GetBackupDataStoreSizeOk() (*int64, bool)`

GetBackupDataStoreSizeOk returns a tuple with the BackupDataStoreSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupDataStoreSize

`func (o *HyperflexBackupPolicySettingsAllOf) SetBackupDataStoreSize(v int64)`

SetBackupDataStoreSize sets BackupDataStoreSize field to given value.

### HasBackupDataStoreSize

`func (o *HyperflexBackupPolicySettingsAllOf) HasBackupDataStoreSize() bool`

HasBackupDataStoreSize returns a boolean if a field has been set.

### GetBackupDataStoreSizeUnit

`func (o *HyperflexBackupPolicySettingsAllOf) GetBackupDataStoreSizeUnit() string`

GetBackupDataStoreSizeUnit returns the BackupDataStoreSizeUnit field if non-nil, zero value otherwise.

### GetBackupDataStoreSizeUnitOk

`func (o *HyperflexBackupPolicySettingsAllOf) GetBackupDataStoreSizeUnitOk() (*string, bool)`

GetBackupDataStoreSizeUnitOk returns a tuple with the BackupDataStoreSizeUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupDataStoreSizeUnit

`func (o *HyperflexBackupPolicySettingsAllOf) SetBackupDataStoreSizeUnit(v string)`

SetBackupDataStoreSizeUnit sets BackupDataStoreSizeUnit field to given value.

### HasBackupDataStoreSizeUnit

`func (o *HyperflexBackupPolicySettingsAllOf) HasBackupDataStoreSizeUnit() bool`

HasBackupDataStoreSizeUnit returns a boolean if a field has been set.

### GetDataStoreEncryptionEnabled

`func (o *HyperflexBackupPolicySettingsAllOf) GetDataStoreEncryptionEnabled() bool`

GetDataStoreEncryptionEnabled returns the DataStoreEncryptionEnabled field if non-nil, zero value otherwise.

### GetDataStoreEncryptionEnabledOk

`func (o *HyperflexBackupPolicySettingsAllOf) GetDataStoreEncryptionEnabledOk() (*bool, bool)`

GetDataStoreEncryptionEnabledOk returns a tuple with the DataStoreEncryptionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataStoreEncryptionEnabled

`func (o *HyperflexBackupPolicySettingsAllOf) SetDataStoreEncryptionEnabled(v bool)`

SetDataStoreEncryptionEnabled sets DataStoreEncryptionEnabled field to given value.

### HasDataStoreEncryptionEnabled

`func (o *HyperflexBackupPolicySettingsAllOf) HasDataStoreEncryptionEnabled() bool`

HasDataStoreEncryptionEnabled returns a boolean if a field has been set.

### GetLocalSnapshotRetentionCount

`func (o *HyperflexBackupPolicySettingsAllOf) GetLocalSnapshotRetentionCount() int64`

GetLocalSnapshotRetentionCount returns the LocalSnapshotRetentionCount field if non-nil, zero value otherwise.

### GetLocalSnapshotRetentionCountOk

`func (o *HyperflexBackupPolicySettingsAllOf) GetLocalSnapshotRetentionCountOk() (*int64, bool)`

GetLocalSnapshotRetentionCountOk returns a tuple with the LocalSnapshotRetentionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalSnapshotRetentionCount

`func (o *HyperflexBackupPolicySettingsAllOf) SetLocalSnapshotRetentionCount(v int64)`

SetLocalSnapshotRetentionCount sets LocalSnapshotRetentionCount field to given value.

### HasLocalSnapshotRetentionCount

`func (o *HyperflexBackupPolicySettingsAllOf) HasLocalSnapshotRetentionCount() bool`

HasLocalSnapshotRetentionCount returns a boolean if a field has been set.

### GetReplicationIntervalInMinutes

`func (o *HyperflexBackupPolicySettingsAllOf) GetReplicationIntervalInMinutes() int64`

GetReplicationIntervalInMinutes returns the ReplicationIntervalInMinutes field if non-nil, zero value otherwise.

### GetReplicationIntervalInMinutesOk

`func (o *HyperflexBackupPolicySettingsAllOf) GetReplicationIntervalInMinutesOk() (*int64, bool)`

GetReplicationIntervalInMinutesOk returns a tuple with the ReplicationIntervalInMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationIntervalInMinutes

`func (o *HyperflexBackupPolicySettingsAllOf) SetReplicationIntervalInMinutes(v int64)`

SetReplicationIntervalInMinutes sets ReplicationIntervalInMinutes field to given value.

### HasReplicationIntervalInMinutes

`func (o *HyperflexBackupPolicySettingsAllOf) HasReplicationIntervalInMinutes() bool`

HasReplicationIntervalInMinutes returns a boolean if a field has been set.

### GetReplicationPairNamePrefix

`func (o *HyperflexBackupPolicySettingsAllOf) GetReplicationPairNamePrefix() string`

GetReplicationPairNamePrefix returns the ReplicationPairNamePrefix field if non-nil, zero value otherwise.

### GetReplicationPairNamePrefixOk

`func (o *HyperflexBackupPolicySettingsAllOf) GetReplicationPairNamePrefixOk() (*string, bool)`

GetReplicationPairNamePrefixOk returns a tuple with the ReplicationPairNamePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationPairNamePrefix

`func (o *HyperflexBackupPolicySettingsAllOf) SetReplicationPairNamePrefix(v string)`

SetReplicationPairNamePrefix sets ReplicationPairNamePrefix field to given value.

### HasReplicationPairNamePrefix

`func (o *HyperflexBackupPolicySettingsAllOf) HasReplicationPairNamePrefix() bool`

HasReplicationPairNamePrefix returns a boolean if a field has been set.

### GetSnapshotRetentionCount

`func (o *HyperflexBackupPolicySettingsAllOf) GetSnapshotRetentionCount() int64`

GetSnapshotRetentionCount returns the SnapshotRetentionCount field if non-nil, zero value otherwise.

### GetSnapshotRetentionCountOk

`func (o *HyperflexBackupPolicySettingsAllOf) GetSnapshotRetentionCountOk() (*int64, bool)`

GetSnapshotRetentionCountOk returns a tuple with the SnapshotRetentionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotRetentionCount

`func (o *HyperflexBackupPolicySettingsAllOf) SetSnapshotRetentionCount(v int64)`

SetSnapshotRetentionCount sets SnapshotRetentionCount field to given value.

### HasSnapshotRetentionCount

`func (o *HyperflexBackupPolicySettingsAllOf) HasSnapshotRetentionCount() bool`

HasSnapshotRetentionCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



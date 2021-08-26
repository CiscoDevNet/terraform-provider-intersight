# HyperflexClusterBackupPolicyAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.ClusterBackupPolicy"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.ClusterBackupPolicy"]
**BackupDataStoreName** | Pointer to **string** | Backup datastore name prefix used during the auto creation of the datastore. All VMs created in this datastore will be automatically backed up. | [optional] [default to "backup-source-ds"]
**BackupDataStoreSize** | Pointer to **int64** | Replication data store size in backupDataStoreSizeUnit. | [optional] [default to 2]
**BackupDataStoreSizeUnit** | Pointer to **string** | Replication data store size. | [optional] [default to "TB"]
**ReplicationPairNamePrefix** | Pointer to **string** | Replication cluster pairing name prefix. | [optional] [default to "backup"]
**ReplicationSchedule** | Pointer to [**NullableHyperflexReplicationSchedule**](HyperflexReplicationSchedule.md) |  | [optional] 
**SnapshotRetentionCount** | Pointer to **int64** | Number of snapshots that will be retained as part of the Multi Point in Time support. | [optional] [default to 4]
**BackupTarget** | Pointer to [**HyperflexClusterRelationship**](HyperflexClusterRelationship.md) |  | [optional] 
**ClusterProfiles** | Pointer to [**[]HyperflexClusterProfileRelationship**](HyperflexClusterProfileRelationship.md) | An array of relationships to hyperflexClusterProfile resources. | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 

## Methods

### NewHyperflexClusterBackupPolicyAllOf

`func NewHyperflexClusterBackupPolicyAllOf(classId string, objectType string, ) *HyperflexClusterBackupPolicyAllOf`

NewHyperflexClusterBackupPolicyAllOf instantiates a new HyperflexClusterBackupPolicyAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexClusterBackupPolicyAllOfWithDefaults

`func NewHyperflexClusterBackupPolicyAllOfWithDefaults() *HyperflexClusterBackupPolicyAllOf`

NewHyperflexClusterBackupPolicyAllOfWithDefaults instantiates a new HyperflexClusterBackupPolicyAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexClusterBackupPolicyAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexClusterBackupPolicyAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexClusterBackupPolicyAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexClusterBackupPolicyAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexClusterBackupPolicyAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexClusterBackupPolicyAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBackupDataStoreName

`func (o *HyperflexClusterBackupPolicyAllOf) GetBackupDataStoreName() string`

GetBackupDataStoreName returns the BackupDataStoreName field if non-nil, zero value otherwise.

### GetBackupDataStoreNameOk

`func (o *HyperflexClusterBackupPolicyAllOf) GetBackupDataStoreNameOk() (*string, bool)`

GetBackupDataStoreNameOk returns a tuple with the BackupDataStoreName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupDataStoreName

`func (o *HyperflexClusterBackupPolicyAllOf) SetBackupDataStoreName(v string)`

SetBackupDataStoreName sets BackupDataStoreName field to given value.

### HasBackupDataStoreName

`func (o *HyperflexClusterBackupPolicyAllOf) HasBackupDataStoreName() bool`

HasBackupDataStoreName returns a boolean if a field has been set.

### GetBackupDataStoreSize

`func (o *HyperflexClusterBackupPolicyAllOf) GetBackupDataStoreSize() int64`

GetBackupDataStoreSize returns the BackupDataStoreSize field if non-nil, zero value otherwise.

### GetBackupDataStoreSizeOk

`func (o *HyperflexClusterBackupPolicyAllOf) GetBackupDataStoreSizeOk() (*int64, bool)`

GetBackupDataStoreSizeOk returns a tuple with the BackupDataStoreSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupDataStoreSize

`func (o *HyperflexClusterBackupPolicyAllOf) SetBackupDataStoreSize(v int64)`

SetBackupDataStoreSize sets BackupDataStoreSize field to given value.

### HasBackupDataStoreSize

`func (o *HyperflexClusterBackupPolicyAllOf) HasBackupDataStoreSize() bool`

HasBackupDataStoreSize returns a boolean if a field has been set.

### GetBackupDataStoreSizeUnit

`func (o *HyperflexClusterBackupPolicyAllOf) GetBackupDataStoreSizeUnit() string`

GetBackupDataStoreSizeUnit returns the BackupDataStoreSizeUnit field if non-nil, zero value otherwise.

### GetBackupDataStoreSizeUnitOk

`func (o *HyperflexClusterBackupPolicyAllOf) GetBackupDataStoreSizeUnitOk() (*string, bool)`

GetBackupDataStoreSizeUnitOk returns a tuple with the BackupDataStoreSizeUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupDataStoreSizeUnit

`func (o *HyperflexClusterBackupPolicyAllOf) SetBackupDataStoreSizeUnit(v string)`

SetBackupDataStoreSizeUnit sets BackupDataStoreSizeUnit field to given value.

### HasBackupDataStoreSizeUnit

`func (o *HyperflexClusterBackupPolicyAllOf) HasBackupDataStoreSizeUnit() bool`

HasBackupDataStoreSizeUnit returns a boolean if a field has been set.

### GetReplicationPairNamePrefix

`func (o *HyperflexClusterBackupPolicyAllOf) GetReplicationPairNamePrefix() string`

GetReplicationPairNamePrefix returns the ReplicationPairNamePrefix field if non-nil, zero value otherwise.

### GetReplicationPairNamePrefixOk

`func (o *HyperflexClusterBackupPolicyAllOf) GetReplicationPairNamePrefixOk() (*string, bool)`

GetReplicationPairNamePrefixOk returns a tuple with the ReplicationPairNamePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationPairNamePrefix

`func (o *HyperflexClusterBackupPolicyAllOf) SetReplicationPairNamePrefix(v string)`

SetReplicationPairNamePrefix sets ReplicationPairNamePrefix field to given value.

### HasReplicationPairNamePrefix

`func (o *HyperflexClusterBackupPolicyAllOf) HasReplicationPairNamePrefix() bool`

HasReplicationPairNamePrefix returns a boolean if a field has been set.

### GetReplicationSchedule

`func (o *HyperflexClusterBackupPolicyAllOf) GetReplicationSchedule() HyperflexReplicationSchedule`

GetReplicationSchedule returns the ReplicationSchedule field if non-nil, zero value otherwise.

### GetReplicationScheduleOk

`func (o *HyperflexClusterBackupPolicyAllOf) GetReplicationScheduleOk() (*HyperflexReplicationSchedule, bool)`

GetReplicationScheduleOk returns a tuple with the ReplicationSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationSchedule

`func (o *HyperflexClusterBackupPolicyAllOf) SetReplicationSchedule(v HyperflexReplicationSchedule)`

SetReplicationSchedule sets ReplicationSchedule field to given value.

### HasReplicationSchedule

`func (o *HyperflexClusterBackupPolicyAllOf) HasReplicationSchedule() bool`

HasReplicationSchedule returns a boolean if a field has been set.

### SetReplicationScheduleNil

`func (o *HyperflexClusterBackupPolicyAllOf) SetReplicationScheduleNil(b bool)`

 SetReplicationScheduleNil sets the value for ReplicationSchedule to be an explicit nil

### UnsetReplicationSchedule
`func (o *HyperflexClusterBackupPolicyAllOf) UnsetReplicationSchedule()`

UnsetReplicationSchedule ensures that no value is present for ReplicationSchedule, not even an explicit nil
### GetSnapshotRetentionCount

`func (o *HyperflexClusterBackupPolicyAllOf) GetSnapshotRetentionCount() int64`

GetSnapshotRetentionCount returns the SnapshotRetentionCount field if non-nil, zero value otherwise.

### GetSnapshotRetentionCountOk

`func (o *HyperflexClusterBackupPolicyAllOf) GetSnapshotRetentionCountOk() (*int64, bool)`

GetSnapshotRetentionCountOk returns a tuple with the SnapshotRetentionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotRetentionCount

`func (o *HyperflexClusterBackupPolicyAllOf) SetSnapshotRetentionCount(v int64)`

SetSnapshotRetentionCount sets SnapshotRetentionCount field to given value.

### HasSnapshotRetentionCount

`func (o *HyperflexClusterBackupPolicyAllOf) HasSnapshotRetentionCount() bool`

HasSnapshotRetentionCount returns a boolean if a field has been set.

### GetBackupTarget

`func (o *HyperflexClusterBackupPolicyAllOf) GetBackupTarget() HyperflexClusterRelationship`

GetBackupTarget returns the BackupTarget field if non-nil, zero value otherwise.

### GetBackupTargetOk

`func (o *HyperflexClusterBackupPolicyAllOf) GetBackupTargetOk() (*HyperflexClusterRelationship, bool)`

GetBackupTargetOk returns a tuple with the BackupTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupTarget

`func (o *HyperflexClusterBackupPolicyAllOf) SetBackupTarget(v HyperflexClusterRelationship)`

SetBackupTarget sets BackupTarget field to given value.

### HasBackupTarget

`func (o *HyperflexClusterBackupPolicyAllOf) HasBackupTarget() bool`

HasBackupTarget returns a boolean if a field has been set.

### GetClusterProfiles

`func (o *HyperflexClusterBackupPolicyAllOf) GetClusterProfiles() []HyperflexClusterProfileRelationship`

GetClusterProfiles returns the ClusterProfiles field if non-nil, zero value otherwise.

### GetClusterProfilesOk

`func (o *HyperflexClusterBackupPolicyAllOf) GetClusterProfilesOk() (*[]HyperflexClusterProfileRelationship, bool)`

GetClusterProfilesOk returns a tuple with the ClusterProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterProfiles

`func (o *HyperflexClusterBackupPolicyAllOf) SetClusterProfiles(v []HyperflexClusterProfileRelationship)`

SetClusterProfiles sets ClusterProfiles field to given value.

### HasClusterProfiles

`func (o *HyperflexClusterBackupPolicyAllOf) HasClusterProfiles() bool`

HasClusterProfiles returns a boolean if a field has been set.

### SetClusterProfilesNil

`func (o *HyperflexClusterBackupPolicyAllOf) SetClusterProfilesNil(b bool)`

 SetClusterProfilesNil sets the value for ClusterProfiles to be an explicit nil

### UnsetClusterProfiles
`func (o *HyperflexClusterBackupPolicyAllOf) UnsetClusterProfiles()`

UnsetClusterProfiles ensures that no value is present for ClusterProfiles, not even an explicit nil
### GetOrganization

`func (o *HyperflexClusterBackupPolicyAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *HyperflexClusterBackupPolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *HyperflexClusterBackupPolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *HyperflexClusterBackupPolicyAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



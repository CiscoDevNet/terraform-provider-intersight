# HyperflexClusterBackupPolicyDeployment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.ClusterBackupPolicyDeployment"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.ClusterBackupPolicyDeployment"]
**BackupDataStoreName** | Pointer to **string** | Backup data store name used during the auto creation of the datastore. All VMs created in this data store will be automatically backed up. | [optional] [readonly] [default to "backup-source-ds"]
**BackupDataStoreSize** | Pointer to **int64** | Replication data store size in backupDataStoreSizeUnit. | [optional] [readonly] [default to 2]
**BackupDataStoreSizeUnit** | Pointer to **string** | Replication data store size. | [optional] [readonly] [default to "TB"]
**Description** | Pointer to **string** | Description from corresponding ClusterBackupPolicy. | [optional] [readonly] 
**Discovered** | Pointer to **bool** | True if record created by discovery on HyperFlex cluster. | [optional] 
**Name** | Pointer to **string** | Name from corresponding ClusterBackupPolicy. | [optional] [readonly] 
**PolicyMoid** | Pointer to **string** | Deployed cluster policy moid. | [optional] [readonly] 
**ProfileMoid** | Pointer to **string** | Deployed cluster profile moid. | [optional] [readonly] 
**ReplicationPairNamePrefix** | Pointer to **string** | Replication cluster pairing name prefix. | [optional] [readonly] [default to "backup"]
**ReplicationSchedule** | Pointer to [**NullableHyperflexReplicationSchedule**](HyperflexReplicationSchedule.md) |  | [optional] 
**SnapshotRetentionCount** | Pointer to **int64** | Number of snapshots that will be retained as part of the Multi Point in Time support. | [optional] [readonly] [default to 4]
**SourceDetached** | Pointer to **bool** | True if policy was detached from source Hyperflex Cluster. | [optional] 
**SourceRequestId** | Pointer to **string** | Unique source cluster request ID allowing retry of the same logical request following a transient communication failure. | [optional] [readonly] 
**SourceUuid** | Pointer to **string** | Uuid of the source Hyperflex Cluster. | [optional] [readonly] 
**TargetDetached** | Pointer to **bool** | True if policy was detached from target Hyperflex Cluster. | [optional] 
**TargetRequestId** | Pointer to **string** | Unique target cluster request ID allowing retry of the same logical request following a transient communication failure. | [optional] [readonly] 
**TargetUuid** | Pointer to **string** | Uuid of the target Hyperflex Cluster. | [optional] [readonly] 
**BackupTarget** | Pointer to [**HyperflexClusterRelationship**](HyperflexClusterRelationship.md) |  | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 
**SourceCluster** | Pointer to [**HyperflexClusterRelationship**](HyperflexClusterRelationship.md) |  | [optional] 

## Methods

### NewHyperflexClusterBackupPolicyDeployment

`func NewHyperflexClusterBackupPolicyDeployment(classId string, objectType string, ) *HyperflexClusterBackupPolicyDeployment`

NewHyperflexClusterBackupPolicyDeployment instantiates a new HyperflexClusterBackupPolicyDeployment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexClusterBackupPolicyDeploymentWithDefaults

`func NewHyperflexClusterBackupPolicyDeploymentWithDefaults() *HyperflexClusterBackupPolicyDeployment`

NewHyperflexClusterBackupPolicyDeploymentWithDefaults instantiates a new HyperflexClusterBackupPolicyDeployment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexClusterBackupPolicyDeployment) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexClusterBackupPolicyDeployment) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexClusterBackupPolicyDeployment) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexClusterBackupPolicyDeployment) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexClusterBackupPolicyDeployment) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexClusterBackupPolicyDeployment) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBackupDataStoreName

`func (o *HyperflexClusterBackupPolicyDeployment) GetBackupDataStoreName() string`

GetBackupDataStoreName returns the BackupDataStoreName field if non-nil, zero value otherwise.

### GetBackupDataStoreNameOk

`func (o *HyperflexClusterBackupPolicyDeployment) GetBackupDataStoreNameOk() (*string, bool)`

GetBackupDataStoreNameOk returns a tuple with the BackupDataStoreName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupDataStoreName

`func (o *HyperflexClusterBackupPolicyDeployment) SetBackupDataStoreName(v string)`

SetBackupDataStoreName sets BackupDataStoreName field to given value.

### HasBackupDataStoreName

`func (o *HyperflexClusterBackupPolicyDeployment) HasBackupDataStoreName() bool`

HasBackupDataStoreName returns a boolean if a field has been set.

### GetBackupDataStoreSize

`func (o *HyperflexClusterBackupPolicyDeployment) GetBackupDataStoreSize() int64`

GetBackupDataStoreSize returns the BackupDataStoreSize field if non-nil, zero value otherwise.

### GetBackupDataStoreSizeOk

`func (o *HyperflexClusterBackupPolicyDeployment) GetBackupDataStoreSizeOk() (*int64, bool)`

GetBackupDataStoreSizeOk returns a tuple with the BackupDataStoreSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupDataStoreSize

`func (o *HyperflexClusterBackupPolicyDeployment) SetBackupDataStoreSize(v int64)`

SetBackupDataStoreSize sets BackupDataStoreSize field to given value.

### HasBackupDataStoreSize

`func (o *HyperflexClusterBackupPolicyDeployment) HasBackupDataStoreSize() bool`

HasBackupDataStoreSize returns a boolean if a field has been set.

### GetBackupDataStoreSizeUnit

`func (o *HyperflexClusterBackupPolicyDeployment) GetBackupDataStoreSizeUnit() string`

GetBackupDataStoreSizeUnit returns the BackupDataStoreSizeUnit field if non-nil, zero value otherwise.

### GetBackupDataStoreSizeUnitOk

`func (o *HyperflexClusterBackupPolicyDeployment) GetBackupDataStoreSizeUnitOk() (*string, bool)`

GetBackupDataStoreSizeUnitOk returns a tuple with the BackupDataStoreSizeUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupDataStoreSizeUnit

`func (o *HyperflexClusterBackupPolicyDeployment) SetBackupDataStoreSizeUnit(v string)`

SetBackupDataStoreSizeUnit sets BackupDataStoreSizeUnit field to given value.

### HasBackupDataStoreSizeUnit

`func (o *HyperflexClusterBackupPolicyDeployment) HasBackupDataStoreSizeUnit() bool`

HasBackupDataStoreSizeUnit returns a boolean if a field has been set.

### GetDescription

`func (o *HyperflexClusterBackupPolicyDeployment) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HyperflexClusterBackupPolicyDeployment) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HyperflexClusterBackupPolicyDeployment) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HyperflexClusterBackupPolicyDeployment) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDiscovered

`func (o *HyperflexClusterBackupPolicyDeployment) GetDiscovered() bool`

GetDiscovered returns the Discovered field if non-nil, zero value otherwise.

### GetDiscoveredOk

`func (o *HyperflexClusterBackupPolicyDeployment) GetDiscoveredOk() (*bool, bool)`

GetDiscoveredOk returns a tuple with the Discovered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscovered

`func (o *HyperflexClusterBackupPolicyDeployment) SetDiscovered(v bool)`

SetDiscovered sets Discovered field to given value.

### HasDiscovered

`func (o *HyperflexClusterBackupPolicyDeployment) HasDiscovered() bool`

HasDiscovered returns a boolean if a field has been set.

### GetName

`func (o *HyperflexClusterBackupPolicyDeployment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HyperflexClusterBackupPolicyDeployment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HyperflexClusterBackupPolicyDeployment) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HyperflexClusterBackupPolicyDeployment) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPolicyMoid

`func (o *HyperflexClusterBackupPolicyDeployment) GetPolicyMoid() string`

GetPolicyMoid returns the PolicyMoid field if non-nil, zero value otherwise.

### GetPolicyMoidOk

`func (o *HyperflexClusterBackupPolicyDeployment) GetPolicyMoidOk() (*string, bool)`

GetPolicyMoidOk returns a tuple with the PolicyMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyMoid

`func (o *HyperflexClusterBackupPolicyDeployment) SetPolicyMoid(v string)`

SetPolicyMoid sets PolicyMoid field to given value.

### HasPolicyMoid

`func (o *HyperflexClusterBackupPolicyDeployment) HasPolicyMoid() bool`

HasPolicyMoid returns a boolean if a field has been set.

### GetProfileMoid

`func (o *HyperflexClusterBackupPolicyDeployment) GetProfileMoid() string`

GetProfileMoid returns the ProfileMoid field if non-nil, zero value otherwise.

### GetProfileMoidOk

`func (o *HyperflexClusterBackupPolicyDeployment) GetProfileMoidOk() (*string, bool)`

GetProfileMoidOk returns a tuple with the ProfileMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileMoid

`func (o *HyperflexClusterBackupPolicyDeployment) SetProfileMoid(v string)`

SetProfileMoid sets ProfileMoid field to given value.

### HasProfileMoid

`func (o *HyperflexClusterBackupPolicyDeployment) HasProfileMoid() bool`

HasProfileMoid returns a boolean if a field has been set.

### GetReplicationPairNamePrefix

`func (o *HyperflexClusterBackupPolicyDeployment) GetReplicationPairNamePrefix() string`

GetReplicationPairNamePrefix returns the ReplicationPairNamePrefix field if non-nil, zero value otherwise.

### GetReplicationPairNamePrefixOk

`func (o *HyperflexClusterBackupPolicyDeployment) GetReplicationPairNamePrefixOk() (*string, bool)`

GetReplicationPairNamePrefixOk returns a tuple with the ReplicationPairNamePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationPairNamePrefix

`func (o *HyperflexClusterBackupPolicyDeployment) SetReplicationPairNamePrefix(v string)`

SetReplicationPairNamePrefix sets ReplicationPairNamePrefix field to given value.

### HasReplicationPairNamePrefix

`func (o *HyperflexClusterBackupPolicyDeployment) HasReplicationPairNamePrefix() bool`

HasReplicationPairNamePrefix returns a boolean if a field has been set.

### GetReplicationSchedule

`func (o *HyperflexClusterBackupPolicyDeployment) GetReplicationSchedule() HyperflexReplicationSchedule`

GetReplicationSchedule returns the ReplicationSchedule field if non-nil, zero value otherwise.

### GetReplicationScheduleOk

`func (o *HyperflexClusterBackupPolicyDeployment) GetReplicationScheduleOk() (*HyperflexReplicationSchedule, bool)`

GetReplicationScheduleOk returns a tuple with the ReplicationSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationSchedule

`func (o *HyperflexClusterBackupPolicyDeployment) SetReplicationSchedule(v HyperflexReplicationSchedule)`

SetReplicationSchedule sets ReplicationSchedule field to given value.

### HasReplicationSchedule

`func (o *HyperflexClusterBackupPolicyDeployment) HasReplicationSchedule() bool`

HasReplicationSchedule returns a boolean if a field has been set.

### SetReplicationScheduleNil

`func (o *HyperflexClusterBackupPolicyDeployment) SetReplicationScheduleNil(b bool)`

 SetReplicationScheduleNil sets the value for ReplicationSchedule to be an explicit nil

### UnsetReplicationSchedule
`func (o *HyperflexClusterBackupPolicyDeployment) UnsetReplicationSchedule()`

UnsetReplicationSchedule ensures that no value is present for ReplicationSchedule, not even an explicit nil
### GetSnapshotRetentionCount

`func (o *HyperflexClusterBackupPolicyDeployment) GetSnapshotRetentionCount() int64`

GetSnapshotRetentionCount returns the SnapshotRetentionCount field if non-nil, zero value otherwise.

### GetSnapshotRetentionCountOk

`func (o *HyperflexClusterBackupPolicyDeployment) GetSnapshotRetentionCountOk() (*int64, bool)`

GetSnapshotRetentionCountOk returns a tuple with the SnapshotRetentionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotRetentionCount

`func (o *HyperflexClusterBackupPolicyDeployment) SetSnapshotRetentionCount(v int64)`

SetSnapshotRetentionCount sets SnapshotRetentionCount field to given value.

### HasSnapshotRetentionCount

`func (o *HyperflexClusterBackupPolicyDeployment) HasSnapshotRetentionCount() bool`

HasSnapshotRetentionCount returns a boolean if a field has been set.

### GetSourceDetached

`func (o *HyperflexClusterBackupPolicyDeployment) GetSourceDetached() bool`

GetSourceDetached returns the SourceDetached field if non-nil, zero value otherwise.

### GetSourceDetachedOk

`func (o *HyperflexClusterBackupPolicyDeployment) GetSourceDetachedOk() (*bool, bool)`

GetSourceDetachedOk returns a tuple with the SourceDetached field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDetached

`func (o *HyperflexClusterBackupPolicyDeployment) SetSourceDetached(v bool)`

SetSourceDetached sets SourceDetached field to given value.

### HasSourceDetached

`func (o *HyperflexClusterBackupPolicyDeployment) HasSourceDetached() bool`

HasSourceDetached returns a boolean if a field has been set.

### GetSourceRequestId

`func (o *HyperflexClusterBackupPolicyDeployment) GetSourceRequestId() string`

GetSourceRequestId returns the SourceRequestId field if non-nil, zero value otherwise.

### GetSourceRequestIdOk

`func (o *HyperflexClusterBackupPolicyDeployment) GetSourceRequestIdOk() (*string, bool)`

GetSourceRequestIdOk returns a tuple with the SourceRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceRequestId

`func (o *HyperflexClusterBackupPolicyDeployment) SetSourceRequestId(v string)`

SetSourceRequestId sets SourceRequestId field to given value.

### HasSourceRequestId

`func (o *HyperflexClusterBackupPolicyDeployment) HasSourceRequestId() bool`

HasSourceRequestId returns a boolean if a field has been set.

### GetSourceUuid

`func (o *HyperflexClusterBackupPolicyDeployment) GetSourceUuid() string`

GetSourceUuid returns the SourceUuid field if non-nil, zero value otherwise.

### GetSourceUuidOk

`func (o *HyperflexClusterBackupPolicyDeployment) GetSourceUuidOk() (*string, bool)`

GetSourceUuidOk returns a tuple with the SourceUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUuid

`func (o *HyperflexClusterBackupPolicyDeployment) SetSourceUuid(v string)`

SetSourceUuid sets SourceUuid field to given value.

### HasSourceUuid

`func (o *HyperflexClusterBackupPolicyDeployment) HasSourceUuid() bool`

HasSourceUuid returns a boolean if a field has been set.

### GetTargetDetached

`func (o *HyperflexClusterBackupPolicyDeployment) GetTargetDetached() bool`

GetTargetDetached returns the TargetDetached field if non-nil, zero value otherwise.

### GetTargetDetachedOk

`func (o *HyperflexClusterBackupPolicyDeployment) GetTargetDetachedOk() (*bool, bool)`

GetTargetDetachedOk returns a tuple with the TargetDetached field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetDetached

`func (o *HyperflexClusterBackupPolicyDeployment) SetTargetDetached(v bool)`

SetTargetDetached sets TargetDetached field to given value.

### HasTargetDetached

`func (o *HyperflexClusterBackupPolicyDeployment) HasTargetDetached() bool`

HasTargetDetached returns a boolean if a field has been set.

### GetTargetRequestId

`func (o *HyperflexClusterBackupPolicyDeployment) GetTargetRequestId() string`

GetTargetRequestId returns the TargetRequestId field if non-nil, zero value otherwise.

### GetTargetRequestIdOk

`func (o *HyperflexClusterBackupPolicyDeployment) GetTargetRequestIdOk() (*string, bool)`

GetTargetRequestIdOk returns a tuple with the TargetRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetRequestId

`func (o *HyperflexClusterBackupPolicyDeployment) SetTargetRequestId(v string)`

SetTargetRequestId sets TargetRequestId field to given value.

### HasTargetRequestId

`func (o *HyperflexClusterBackupPolicyDeployment) HasTargetRequestId() bool`

HasTargetRequestId returns a boolean if a field has been set.

### GetTargetUuid

`func (o *HyperflexClusterBackupPolicyDeployment) GetTargetUuid() string`

GetTargetUuid returns the TargetUuid field if non-nil, zero value otherwise.

### GetTargetUuidOk

`func (o *HyperflexClusterBackupPolicyDeployment) GetTargetUuidOk() (*string, bool)`

GetTargetUuidOk returns a tuple with the TargetUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUuid

`func (o *HyperflexClusterBackupPolicyDeployment) SetTargetUuid(v string)`

SetTargetUuid sets TargetUuid field to given value.

### HasTargetUuid

`func (o *HyperflexClusterBackupPolicyDeployment) HasTargetUuid() bool`

HasTargetUuid returns a boolean if a field has been set.

### GetBackupTarget

`func (o *HyperflexClusterBackupPolicyDeployment) GetBackupTarget() HyperflexClusterRelationship`

GetBackupTarget returns the BackupTarget field if non-nil, zero value otherwise.

### GetBackupTargetOk

`func (o *HyperflexClusterBackupPolicyDeployment) GetBackupTargetOk() (*HyperflexClusterRelationship, bool)`

GetBackupTargetOk returns a tuple with the BackupTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupTarget

`func (o *HyperflexClusterBackupPolicyDeployment) SetBackupTarget(v HyperflexClusterRelationship)`

SetBackupTarget sets BackupTarget field to given value.

### HasBackupTarget

`func (o *HyperflexClusterBackupPolicyDeployment) HasBackupTarget() bool`

HasBackupTarget returns a boolean if a field has been set.

### GetOrganization

`func (o *HyperflexClusterBackupPolicyDeployment) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *HyperflexClusterBackupPolicyDeployment) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *HyperflexClusterBackupPolicyDeployment) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *HyperflexClusterBackupPolicyDeployment) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetSourceCluster

`func (o *HyperflexClusterBackupPolicyDeployment) GetSourceCluster() HyperflexClusterRelationship`

GetSourceCluster returns the SourceCluster field if non-nil, zero value otherwise.

### GetSourceClusterOk

`func (o *HyperflexClusterBackupPolicyDeployment) GetSourceClusterOk() (*HyperflexClusterRelationship, bool)`

GetSourceClusterOk returns a tuple with the SourceCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceCluster

`func (o *HyperflexClusterBackupPolicyDeployment) SetSourceCluster(v HyperflexClusterRelationship)`

SetSourceCluster sets SourceCluster field to given value.

### HasSourceCluster

`func (o *HyperflexClusterBackupPolicyDeployment) HasSourceCluster() bool`

HasSourceCluster returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



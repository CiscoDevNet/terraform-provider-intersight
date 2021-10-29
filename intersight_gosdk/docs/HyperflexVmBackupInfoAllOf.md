# HyperflexVmBackupInfoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.VmBackupInfo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.VmBackupInfo"]
**BackupStatus** | Pointer to **string** | Description of the backup status of this VmBackupInfo. * &#x60;InitializingProtection&#x60; - Protection has started, but not completed. * &#x60;Protected&#x60; - Protection has completed successfully. * &#x60;ExceedsInterval&#x60; - Protection has not completed successfully in over two times the backup interval. | [optional] [readonly] [default to "InitializingProtection"]
**ClusterEntityReference** | Pointer to [**NullableHyperflexEntityReference**](HyperflexEntityReference.md) |  | [optional] 
**ClusterIdProtectionInfoMap** | Pointer to [**[]HyperflexMapClusterIdToProtectionInfo**](HyperflexMapClusterIdToProtectionInfo.md) |  | [optional] 
**Error** | Pointer to [**NullableHyperflexErrorStack**](HyperflexErrorStack.md) |  | [optional] 
**LocalSnapshotRetentionCount** | Pointer to **int64** | Retention count from backup policy for local snapshots. | [optional] [readonly] 
**PowerOn** | Pointer to **bool** | The power state of the Virtual Machine. | [optional] [readonly] 
**ProtectionStatus** | Pointer to **string** | Description of the protection status of this VmBackupInfo. * &#x60;PREPARE_FAILOVER_STARTED&#x60; - The protection status is prepare failover started. * &#x60;PREPARE_FAILOVER_FAILED&#x60; - The protection status is prepare failover failed. * &#x60;PREPARE_FAILOVER_COMPLETED&#x60; - The protection status is prepaire failover completed. * &#x60;FAILOVER_STARTED&#x60; - The protection status is failover started. * &#x60;FAILOVER_FAILED&#x60; - The protection status is failover failed. * &#x60;FAILOVER_COMPLETED&#x60; - The protection status is failover completed. * &#x60;PREPARE_REVERSEPROTECT_STARTED&#x60; - The protection status is prepare reverse protect started. * &#x60;PREPARE_REVERSEPROTECT_FAILED&#x60; - The protection status is prepare reverse protect failed. * &#x60;PREPARE_REVERSEPROTECT_COMPLETED&#x60; - The protection status is prepair reverse protect completed. * &#x60;REVERSEPROTECT_STARTED&#x60; - The protection status is reverse protect started. * &#x60;REVERSEPROTECT_FAILED&#x60; - The protection status is reverse protect failed. * &#x60;ACTIVE&#x60; - The protection status is active. * &#x60;CREATION_IN_PROGRESS&#x60; - The protection status is failover in progress. * &#x60;CREATION_FAILED&#x60; - The protection status is creation failed. * &#x60;LOCAL_RESTORE_STARTED&#x60; - The protection status is local restore started. * &#x60;LOCAL_RESTORE_FAILED&#x60; - The protection status is local restore failed. | [optional] [readonly] [default to "PREPARE_FAILOVER_STARTED"]
**Schedule** | Pointer to [**[]HyperflexReplicationClusterReferenceToSchedule**](HyperflexReplicationClusterReferenceToSchedule.md) |  | [optional] 
**SnapshotRetentionCount** | Pointer to **int64** | Retention count from backup policy for remote snapshots. | [optional] [readonly] 
**SrcClusterName** | Pointer to **string** | Name for the source cluster this Virtual Machine is residing on. | [optional] [readonly] 
**TgtClusterName** | Pointer to **string** | Name for the target cluster this Virtual Machine is residing on. | [optional] [readonly] 
**VmEntityReference** | Pointer to [**NullableHyperflexEntityReference**](HyperflexEntityReference.md) |  | [optional] 
**VmInfo** | Pointer to [**NullableHyperflexVirtualMachine**](HyperflexVirtualMachine.md) |  | [optional] 
**SrcBackupCluster** | Pointer to [**HyperflexBackupClusterRelationship**](HyperflexBackupClusterRelationship.md) |  | [optional] 
**SrcCluster** | Pointer to [**HyperflexClusterRelationship**](HyperflexClusterRelationship.md) |  | [optional] 
**TgtCluster** | Pointer to [**HyperflexClusterRelationship**](HyperflexClusterRelationship.md) |  | [optional] 

## Methods

### NewHyperflexVmBackupInfoAllOf

`func NewHyperflexVmBackupInfoAllOf(classId string, objectType string, ) *HyperflexVmBackupInfoAllOf`

NewHyperflexVmBackupInfoAllOf instantiates a new HyperflexVmBackupInfoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexVmBackupInfoAllOfWithDefaults

`func NewHyperflexVmBackupInfoAllOfWithDefaults() *HyperflexVmBackupInfoAllOf`

NewHyperflexVmBackupInfoAllOfWithDefaults instantiates a new HyperflexVmBackupInfoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexVmBackupInfoAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexVmBackupInfoAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexVmBackupInfoAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexVmBackupInfoAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexVmBackupInfoAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexVmBackupInfoAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBackupStatus

`func (o *HyperflexVmBackupInfoAllOf) GetBackupStatus() string`

GetBackupStatus returns the BackupStatus field if non-nil, zero value otherwise.

### GetBackupStatusOk

`func (o *HyperflexVmBackupInfoAllOf) GetBackupStatusOk() (*string, bool)`

GetBackupStatusOk returns a tuple with the BackupStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupStatus

`func (o *HyperflexVmBackupInfoAllOf) SetBackupStatus(v string)`

SetBackupStatus sets BackupStatus field to given value.

### HasBackupStatus

`func (o *HyperflexVmBackupInfoAllOf) HasBackupStatus() bool`

HasBackupStatus returns a boolean if a field has been set.

### GetClusterEntityReference

`func (o *HyperflexVmBackupInfoAllOf) GetClusterEntityReference() HyperflexEntityReference`

GetClusterEntityReference returns the ClusterEntityReference field if non-nil, zero value otherwise.

### GetClusterEntityReferenceOk

`func (o *HyperflexVmBackupInfoAllOf) GetClusterEntityReferenceOk() (*HyperflexEntityReference, bool)`

GetClusterEntityReferenceOk returns a tuple with the ClusterEntityReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterEntityReference

`func (o *HyperflexVmBackupInfoAllOf) SetClusterEntityReference(v HyperflexEntityReference)`

SetClusterEntityReference sets ClusterEntityReference field to given value.

### HasClusterEntityReference

`func (o *HyperflexVmBackupInfoAllOf) HasClusterEntityReference() bool`

HasClusterEntityReference returns a boolean if a field has been set.

### SetClusterEntityReferenceNil

`func (o *HyperflexVmBackupInfoAllOf) SetClusterEntityReferenceNil(b bool)`

 SetClusterEntityReferenceNil sets the value for ClusterEntityReference to be an explicit nil

### UnsetClusterEntityReference
`func (o *HyperflexVmBackupInfoAllOf) UnsetClusterEntityReference()`

UnsetClusterEntityReference ensures that no value is present for ClusterEntityReference, not even an explicit nil
### GetClusterIdProtectionInfoMap

`func (o *HyperflexVmBackupInfoAllOf) GetClusterIdProtectionInfoMap() []HyperflexMapClusterIdToProtectionInfo`

GetClusterIdProtectionInfoMap returns the ClusterIdProtectionInfoMap field if non-nil, zero value otherwise.

### GetClusterIdProtectionInfoMapOk

`func (o *HyperflexVmBackupInfoAllOf) GetClusterIdProtectionInfoMapOk() (*[]HyperflexMapClusterIdToProtectionInfo, bool)`

GetClusterIdProtectionInfoMapOk returns a tuple with the ClusterIdProtectionInfoMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterIdProtectionInfoMap

`func (o *HyperflexVmBackupInfoAllOf) SetClusterIdProtectionInfoMap(v []HyperflexMapClusterIdToProtectionInfo)`

SetClusterIdProtectionInfoMap sets ClusterIdProtectionInfoMap field to given value.

### HasClusterIdProtectionInfoMap

`func (o *HyperflexVmBackupInfoAllOf) HasClusterIdProtectionInfoMap() bool`

HasClusterIdProtectionInfoMap returns a boolean if a field has been set.

### SetClusterIdProtectionInfoMapNil

`func (o *HyperflexVmBackupInfoAllOf) SetClusterIdProtectionInfoMapNil(b bool)`

 SetClusterIdProtectionInfoMapNil sets the value for ClusterIdProtectionInfoMap to be an explicit nil

### UnsetClusterIdProtectionInfoMap
`func (o *HyperflexVmBackupInfoAllOf) UnsetClusterIdProtectionInfoMap()`

UnsetClusterIdProtectionInfoMap ensures that no value is present for ClusterIdProtectionInfoMap, not even an explicit nil
### GetError

`func (o *HyperflexVmBackupInfoAllOf) GetError() HyperflexErrorStack`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *HyperflexVmBackupInfoAllOf) GetErrorOk() (*HyperflexErrorStack, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *HyperflexVmBackupInfoAllOf) SetError(v HyperflexErrorStack)`

SetError sets Error field to given value.

### HasError

`func (o *HyperflexVmBackupInfoAllOf) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *HyperflexVmBackupInfoAllOf) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *HyperflexVmBackupInfoAllOf) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetLocalSnapshotRetentionCount

`func (o *HyperflexVmBackupInfoAllOf) GetLocalSnapshotRetentionCount() int64`

GetLocalSnapshotRetentionCount returns the LocalSnapshotRetentionCount field if non-nil, zero value otherwise.

### GetLocalSnapshotRetentionCountOk

`func (o *HyperflexVmBackupInfoAllOf) GetLocalSnapshotRetentionCountOk() (*int64, bool)`

GetLocalSnapshotRetentionCountOk returns a tuple with the LocalSnapshotRetentionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalSnapshotRetentionCount

`func (o *HyperflexVmBackupInfoAllOf) SetLocalSnapshotRetentionCount(v int64)`

SetLocalSnapshotRetentionCount sets LocalSnapshotRetentionCount field to given value.

### HasLocalSnapshotRetentionCount

`func (o *HyperflexVmBackupInfoAllOf) HasLocalSnapshotRetentionCount() bool`

HasLocalSnapshotRetentionCount returns a boolean if a field has been set.

### GetPowerOn

`func (o *HyperflexVmBackupInfoAllOf) GetPowerOn() bool`

GetPowerOn returns the PowerOn field if non-nil, zero value otherwise.

### GetPowerOnOk

`func (o *HyperflexVmBackupInfoAllOf) GetPowerOnOk() (*bool, bool)`

GetPowerOnOk returns a tuple with the PowerOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerOn

`func (o *HyperflexVmBackupInfoAllOf) SetPowerOn(v bool)`

SetPowerOn sets PowerOn field to given value.

### HasPowerOn

`func (o *HyperflexVmBackupInfoAllOf) HasPowerOn() bool`

HasPowerOn returns a boolean if a field has been set.

### GetProtectionStatus

`func (o *HyperflexVmBackupInfoAllOf) GetProtectionStatus() string`

GetProtectionStatus returns the ProtectionStatus field if non-nil, zero value otherwise.

### GetProtectionStatusOk

`func (o *HyperflexVmBackupInfoAllOf) GetProtectionStatusOk() (*string, bool)`

GetProtectionStatusOk returns a tuple with the ProtectionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionStatus

`func (o *HyperflexVmBackupInfoAllOf) SetProtectionStatus(v string)`

SetProtectionStatus sets ProtectionStatus field to given value.

### HasProtectionStatus

`func (o *HyperflexVmBackupInfoAllOf) HasProtectionStatus() bool`

HasProtectionStatus returns a boolean if a field has been set.

### GetSchedule

`func (o *HyperflexVmBackupInfoAllOf) GetSchedule() []HyperflexReplicationClusterReferenceToSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *HyperflexVmBackupInfoAllOf) GetScheduleOk() (*[]HyperflexReplicationClusterReferenceToSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *HyperflexVmBackupInfoAllOf) SetSchedule(v []HyperflexReplicationClusterReferenceToSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *HyperflexVmBackupInfoAllOf) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### SetScheduleNil

`func (o *HyperflexVmBackupInfoAllOf) SetScheduleNil(b bool)`

 SetScheduleNil sets the value for Schedule to be an explicit nil

### UnsetSchedule
`func (o *HyperflexVmBackupInfoAllOf) UnsetSchedule()`

UnsetSchedule ensures that no value is present for Schedule, not even an explicit nil
### GetSnapshotRetentionCount

`func (o *HyperflexVmBackupInfoAllOf) GetSnapshotRetentionCount() int64`

GetSnapshotRetentionCount returns the SnapshotRetentionCount field if non-nil, zero value otherwise.

### GetSnapshotRetentionCountOk

`func (o *HyperflexVmBackupInfoAllOf) GetSnapshotRetentionCountOk() (*int64, bool)`

GetSnapshotRetentionCountOk returns a tuple with the SnapshotRetentionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotRetentionCount

`func (o *HyperflexVmBackupInfoAllOf) SetSnapshotRetentionCount(v int64)`

SetSnapshotRetentionCount sets SnapshotRetentionCount field to given value.

### HasSnapshotRetentionCount

`func (o *HyperflexVmBackupInfoAllOf) HasSnapshotRetentionCount() bool`

HasSnapshotRetentionCount returns a boolean if a field has been set.

### GetSrcClusterName

`func (o *HyperflexVmBackupInfoAllOf) GetSrcClusterName() string`

GetSrcClusterName returns the SrcClusterName field if non-nil, zero value otherwise.

### GetSrcClusterNameOk

`func (o *HyperflexVmBackupInfoAllOf) GetSrcClusterNameOk() (*string, bool)`

GetSrcClusterNameOk returns a tuple with the SrcClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcClusterName

`func (o *HyperflexVmBackupInfoAllOf) SetSrcClusterName(v string)`

SetSrcClusterName sets SrcClusterName field to given value.

### HasSrcClusterName

`func (o *HyperflexVmBackupInfoAllOf) HasSrcClusterName() bool`

HasSrcClusterName returns a boolean if a field has been set.

### GetTgtClusterName

`func (o *HyperflexVmBackupInfoAllOf) GetTgtClusterName() string`

GetTgtClusterName returns the TgtClusterName field if non-nil, zero value otherwise.

### GetTgtClusterNameOk

`func (o *HyperflexVmBackupInfoAllOf) GetTgtClusterNameOk() (*string, bool)`

GetTgtClusterNameOk returns a tuple with the TgtClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTgtClusterName

`func (o *HyperflexVmBackupInfoAllOf) SetTgtClusterName(v string)`

SetTgtClusterName sets TgtClusterName field to given value.

### HasTgtClusterName

`func (o *HyperflexVmBackupInfoAllOf) HasTgtClusterName() bool`

HasTgtClusterName returns a boolean if a field has been set.

### GetVmEntityReference

`func (o *HyperflexVmBackupInfoAllOf) GetVmEntityReference() HyperflexEntityReference`

GetVmEntityReference returns the VmEntityReference field if non-nil, zero value otherwise.

### GetVmEntityReferenceOk

`func (o *HyperflexVmBackupInfoAllOf) GetVmEntityReferenceOk() (*HyperflexEntityReference, bool)`

GetVmEntityReferenceOk returns a tuple with the VmEntityReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmEntityReference

`func (o *HyperflexVmBackupInfoAllOf) SetVmEntityReference(v HyperflexEntityReference)`

SetVmEntityReference sets VmEntityReference field to given value.

### HasVmEntityReference

`func (o *HyperflexVmBackupInfoAllOf) HasVmEntityReference() bool`

HasVmEntityReference returns a boolean if a field has been set.

### SetVmEntityReferenceNil

`func (o *HyperflexVmBackupInfoAllOf) SetVmEntityReferenceNil(b bool)`

 SetVmEntityReferenceNil sets the value for VmEntityReference to be an explicit nil

### UnsetVmEntityReference
`func (o *HyperflexVmBackupInfoAllOf) UnsetVmEntityReference()`

UnsetVmEntityReference ensures that no value is present for VmEntityReference, not even an explicit nil
### GetVmInfo

`func (o *HyperflexVmBackupInfoAllOf) GetVmInfo() HyperflexVirtualMachine`

GetVmInfo returns the VmInfo field if non-nil, zero value otherwise.

### GetVmInfoOk

`func (o *HyperflexVmBackupInfoAllOf) GetVmInfoOk() (*HyperflexVirtualMachine, bool)`

GetVmInfoOk returns a tuple with the VmInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmInfo

`func (o *HyperflexVmBackupInfoAllOf) SetVmInfo(v HyperflexVirtualMachine)`

SetVmInfo sets VmInfo field to given value.

### HasVmInfo

`func (o *HyperflexVmBackupInfoAllOf) HasVmInfo() bool`

HasVmInfo returns a boolean if a field has been set.

### SetVmInfoNil

`func (o *HyperflexVmBackupInfoAllOf) SetVmInfoNil(b bool)`

 SetVmInfoNil sets the value for VmInfo to be an explicit nil

### UnsetVmInfo
`func (o *HyperflexVmBackupInfoAllOf) UnsetVmInfo()`

UnsetVmInfo ensures that no value is present for VmInfo, not even an explicit nil
### GetSrcBackupCluster

`func (o *HyperflexVmBackupInfoAllOf) GetSrcBackupCluster() HyperflexBackupClusterRelationship`

GetSrcBackupCluster returns the SrcBackupCluster field if non-nil, zero value otherwise.

### GetSrcBackupClusterOk

`func (o *HyperflexVmBackupInfoAllOf) GetSrcBackupClusterOk() (*HyperflexBackupClusterRelationship, bool)`

GetSrcBackupClusterOk returns a tuple with the SrcBackupCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcBackupCluster

`func (o *HyperflexVmBackupInfoAllOf) SetSrcBackupCluster(v HyperflexBackupClusterRelationship)`

SetSrcBackupCluster sets SrcBackupCluster field to given value.

### HasSrcBackupCluster

`func (o *HyperflexVmBackupInfoAllOf) HasSrcBackupCluster() bool`

HasSrcBackupCluster returns a boolean if a field has been set.

### GetSrcCluster

`func (o *HyperflexVmBackupInfoAllOf) GetSrcCluster() HyperflexClusterRelationship`

GetSrcCluster returns the SrcCluster field if non-nil, zero value otherwise.

### GetSrcClusterOk

`func (o *HyperflexVmBackupInfoAllOf) GetSrcClusterOk() (*HyperflexClusterRelationship, bool)`

GetSrcClusterOk returns a tuple with the SrcCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcCluster

`func (o *HyperflexVmBackupInfoAllOf) SetSrcCluster(v HyperflexClusterRelationship)`

SetSrcCluster sets SrcCluster field to given value.

### HasSrcCluster

`func (o *HyperflexVmBackupInfoAllOf) HasSrcCluster() bool`

HasSrcCluster returns a boolean if a field has been set.

### GetTgtCluster

`func (o *HyperflexVmBackupInfoAllOf) GetTgtCluster() HyperflexClusterRelationship`

GetTgtCluster returns the TgtCluster field if non-nil, zero value otherwise.

### GetTgtClusterOk

`func (o *HyperflexVmBackupInfoAllOf) GetTgtClusterOk() (*HyperflexClusterRelationship, bool)`

GetTgtClusterOk returns a tuple with the TgtCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTgtCluster

`func (o *HyperflexVmBackupInfoAllOf) SetTgtCluster(v HyperflexClusterRelationship)`

SetTgtCluster sets TgtCluster field to given value.

### HasTgtCluster

`func (o *HyperflexVmBackupInfoAllOf) HasTgtCluster() bool`

HasTgtCluster returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



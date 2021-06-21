# HyperflexVmBackupInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.VmBackupInfo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.VmBackupInfo"]
**BackupStatus** | Pointer to **string** | Description of the backup status of this VmBackupInfo. * &#x60;InitializingProtection&#x60; - Protection has started, but not completed. * &#x60;Protected&#x60; - Protection has completed successfully. * &#x60;ExceedsInterval&#x60; - Protection has not completed successfully in over two times the backup interval. | [optional] [readonly] [default to "InitializingProtection"]
**ClusterEntityReference** | Pointer to [**NullableHyperflexEntityReference**](hyperflex.EntityReference.md) |  | [optional] 
**ClusterIdProtectionInfoMap** | Pointer to [**[]HyperflexMapClusterIdToProtectionInfo**](HyperflexMapClusterIdToProtectionInfo.md) |  | [optional] 
**Error** | Pointer to [**NullableHyperflexErrorStack**](hyperflex.ErrorStack.md) |  | [optional] 
**PowerOn** | Pointer to **bool** | The power state of the Virtual Machine. | [optional] [readonly] 
**ProtectionStatus** | Pointer to **string** | Description of the protection status of this VmBackupInfo. * &#x60;PREPARE_FAILOVER_STARTED&#x60; - The protection status is prepare failover started. * &#x60;PREPARE_FAILOVER_FAILED&#x60; - The protection status is prepare failover failed. * &#x60;PREPARE_FAILOVER_COMPLETED&#x60; - The protection status is prepaire failover completed. * &#x60;FAILOVER_STARTED&#x60; - The protection status is failover started. * &#x60;FAILOVER_FAILED&#x60; - The protection status is failover failed. * &#x60;FAILOVER_COMPLETED&#x60; - The protection status is failover completed. * &#x60;PREPARE_REVERSEPROTECT_STARTED&#x60; - The protection status is prepare reverse protect started. * &#x60;PREPARE_REVERSEPROTECT_FAILED&#x60; - The protection status is prepare reverse protect failed. * &#x60;PREPARE_REVERSEPROTECT_COMPLETED&#x60; - The protection status is prepair reverse protect completed. * &#x60;REVERSEPROTECT_STARTED&#x60; - The protection status is reverse protect started. * &#x60;REVERSEPROTECT_FAILED&#x60; - The protection status is reverse protect failed. * &#x60;ACTIVE&#x60; - The protection status is active. * &#x60;CREATION_IN_PROGRESS&#x60; - The protection status is failover in progress. * &#x60;CREATION_FAILED&#x60; - The protection status is creation failed. * &#x60;LOCAL_RESTORE_STARTED&#x60; - The protection status is local restore started. * &#x60;LOCAL_RESTORE_FAILED&#x60; - The protection status is local restore failed. | [optional] [readonly] [default to "PREPARE_FAILOVER_STARTED"]
**Schedule** | Pointer to [**[]HyperflexReplicationClusterReferenceToSchedule**](HyperflexReplicationClusterReferenceToSchedule.md) |  | [optional] 
**VmEntityReference** | Pointer to [**NullableHyperflexEntityReference**](hyperflex.EntityReference.md) |  | [optional] 
**VmInfo** | Pointer to [**NullableHyperflexVirtualMachine**](hyperflex.VirtualMachine.md) |  | [optional] 
**SrcBackupCluster** | Pointer to [**HyperflexBackupClusterRelationship**](hyperflex.BackupCluster.Relationship.md) |  | [optional] 
**SrcCluster** | Pointer to [**HyperflexClusterRelationship**](hyperflex.Cluster.Relationship.md) |  | [optional] 
**TgtCluster** | Pointer to [**HyperflexClusterRelationship**](hyperflex.Cluster.Relationship.md) |  | [optional] 

## Methods

### NewHyperflexVmBackupInfo

`func NewHyperflexVmBackupInfo(classId string, objectType string, ) *HyperflexVmBackupInfo`

NewHyperflexVmBackupInfo instantiates a new HyperflexVmBackupInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexVmBackupInfoWithDefaults

`func NewHyperflexVmBackupInfoWithDefaults() *HyperflexVmBackupInfo`

NewHyperflexVmBackupInfoWithDefaults instantiates a new HyperflexVmBackupInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexVmBackupInfo) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexVmBackupInfo) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexVmBackupInfo) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexVmBackupInfo) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexVmBackupInfo) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexVmBackupInfo) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBackupStatus

`func (o *HyperflexVmBackupInfo) GetBackupStatus() string`

GetBackupStatus returns the BackupStatus field if non-nil, zero value otherwise.

### GetBackupStatusOk

`func (o *HyperflexVmBackupInfo) GetBackupStatusOk() (*string, bool)`

GetBackupStatusOk returns a tuple with the BackupStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupStatus

`func (o *HyperflexVmBackupInfo) SetBackupStatus(v string)`

SetBackupStatus sets BackupStatus field to given value.

### HasBackupStatus

`func (o *HyperflexVmBackupInfo) HasBackupStatus() bool`

HasBackupStatus returns a boolean if a field has been set.

### GetClusterEntityReference

`func (o *HyperflexVmBackupInfo) GetClusterEntityReference() HyperflexEntityReference`

GetClusterEntityReference returns the ClusterEntityReference field if non-nil, zero value otherwise.

### GetClusterEntityReferenceOk

`func (o *HyperflexVmBackupInfo) GetClusterEntityReferenceOk() (*HyperflexEntityReference, bool)`

GetClusterEntityReferenceOk returns a tuple with the ClusterEntityReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterEntityReference

`func (o *HyperflexVmBackupInfo) SetClusterEntityReference(v HyperflexEntityReference)`

SetClusterEntityReference sets ClusterEntityReference field to given value.

### HasClusterEntityReference

`func (o *HyperflexVmBackupInfo) HasClusterEntityReference() bool`

HasClusterEntityReference returns a boolean if a field has been set.

### SetClusterEntityReferenceNil

`func (o *HyperflexVmBackupInfo) SetClusterEntityReferenceNil(b bool)`

 SetClusterEntityReferenceNil sets the value for ClusterEntityReference to be an explicit nil

### UnsetClusterEntityReference
`func (o *HyperflexVmBackupInfo) UnsetClusterEntityReference()`

UnsetClusterEntityReference ensures that no value is present for ClusterEntityReference, not even an explicit nil
### GetClusterIdProtectionInfoMap

`func (o *HyperflexVmBackupInfo) GetClusterIdProtectionInfoMap() []HyperflexMapClusterIdToProtectionInfo`

GetClusterIdProtectionInfoMap returns the ClusterIdProtectionInfoMap field if non-nil, zero value otherwise.

### GetClusterIdProtectionInfoMapOk

`func (o *HyperflexVmBackupInfo) GetClusterIdProtectionInfoMapOk() (*[]HyperflexMapClusterIdToProtectionInfo, bool)`

GetClusterIdProtectionInfoMapOk returns a tuple with the ClusterIdProtectionInfoMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterIdProtectionInfoMap

`func (o *HyperflexVmBackupInfo) SetClusterIdProtectionInfoMap(v []HyperflexMapClusterIdToProtectionInfo)`

SetClusterIdProtectionInfoMap sets ClusterIdProtectionInfoMap field to given value.

### HasClusterIdProtectionInfoMap

`func (o *HyperflexVmBackupInfo) HasClusterIdProtectionInfoMap() bool`

HasClusterIdProtectionInfoMap returns a boolean if a field has been set.

### SetClusterIdProtectionInfoMapNil

`func (o *HyperflexVmBackupInfo) SetClusterIdProtectionInfoMapNil(b bool)`

 SetClusterIdProtectionInfoMapNil sets the value for ClusterIdProtectionInfoMap to be an explicit nil

### UnsetClusterIdProtectionInfoMap
`func (o *HyperflexVmBackupInfo) UnsetClusterIdProtectionInfoMap()`

UnsetClusterIdProtectionInfoMap ensures that no value is present for ClusterIdProtectionInfoMap, not even an explicit nil
### GetError

`func (o *HyperflexVmBackupInfo) GetError() HyperflexErrorStack`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *HyperflexVmBackupInfo) GetErrorOk() (*HyperflexErrorStack, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *HyperflexVmBackupInfo) SetError(v HyperflexErrorStack)`

SetError sets Error field to given value.

### HasError

`func (o *HyperflexVmBackupInfo) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *HyperflexVmBackupInfo) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *HyperflexVmBackupInfo) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetPowerOn

`func (o *HyperflexVmBackupInfo) GetPowerOn() bool`

GetPowerOn returns the PowerOn field if non-nil, zero value otherwise.

### GetPowerOnOk

`func (o *HyperflexVmBackupInfo) GetPowerOnOk() (*bool, bool)`

GetPowerOnOk returns a tuple with the PowerOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerOn

`func (o *HyperflexVmBackupInfo) SetPowerOn(v bool)`

SetPowerOn sets PowerOn field to given value.

### HasPowerOn

`func (o *HyperflexVmBackupInfo) HasPowerOn() bool`

HasPowerOn returns a boolean if a field has been set.

### GetProtectionStatus

`func (o *HyperflexVmBackupInfo) GetProtectionStatus() string`

GetProtectionStatus returns the ProtectionStatus field if non-nil, zero value otherwise.

### GetProtectionStatusOk

`func (o *HyperflexVmBackupInfo) GetProtectionStatusOk() (*string, bool)`

GetProtectionStatusOk returns a tuple with the ProtectionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionStatus

`func (o *HyperflexVmBackupInfo) SetProtectionStatus(v string)`

SetProtectionStatus sets ProtectionStatus field to given value.

### HasProtectionStatus

`func (o *HyperflexVmBackupInfo) HasProtectionStatus() bool`

HasProtectionStatus returns a boolean if a field has been set.

### GetSchedule

`func (o *HyperflexVmBackupInfo) GetSchedule() []HyperflexReplicationClusterReferenceToSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *HyperflexVmBackupInfo) GetScheduleOk() (*[]HyperflexReplicationClusterReferenceToSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *HyperflexVmBackupInfo) SetSchedule(v []HyperflexReplicationClusterReferenceToSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *HyperflexVmBackupInfo) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### SetScheduleNil

`func (o *HyperflexVmBackupInfo) SetScheduleNil(b bool)`

 SetScheduleNil sets the value for Schedule to be an explicit nil

### UnsetSchedule
`func (o *HyperflexVmBackupInfo) UnsetSchedule()`

UnsetSchedule ensures that no value is present for Schedule, not even an explicit nil
### GetVmEntityReference

`func (o *HyperflexVmBackupInfo) GetVmEntityReference() HyperflexEntityReference`

GetVmEntityReference returns the VmEntityReference field if non-nil, zero value otherwise.

### GetVmEntityReferenceOk

`func (o *HyperflexVmBackupInfo) GetVmEntityReferenceOk() (*HyperflexEntityReference, bool)`

GetVmEntityReferenceOk returns a tuple with the VmEntityReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmEntityReference

`func (o *HyperflexVmBackupInfo) SetVmEntityReference(v HyperflexEntityReference)`

SetVmEntityReference sets VmEntityReference field to given value.

### HasVmEntityReference

`func (o *HyperflexVmBackupInfo) HasVmEntityReference() bool`

HasVmEntityReference returns a boolean if a field has been set.

### SetVmEntityReferenceNil

`func (o *HyperflexVmBackupInfo) SetVmEntityReferenceNil(b bool)`

 SetVmEntityReferenceNil sets the value for VmEntityReference to be an explicit nil

### UnsetVmEntityReference
`func (o *HyperflexVmBackupInfo) UnsetVmEntityReference()`

UnsetVmEntityReference ensures that no value is present for VmEntityReference, not even an explicit nil
### GetVmInfo

`func (o *HyperflexVmBackupInfo) GetVmInfo() HyperflexVirtualMachine`

GetVmInfo returns the VmInfo field if non-nil, zero value otherwise.

### GetVmInfoOk

`func (o *HyperflexVmBackupInfo) GetVmInfoOk() (*HyperflexVirtualMachine, bool)`

GetVmInfoOk returns a tuple with the VmInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmInfo

`func (o *HyperflexVmBackupInfo) SetVmInfo(v HyperflexVirtualMachine)`

SetVmInfo sets VmInfo field to given value.

### HasVmInfo

`func (o *HyperflexVmBackupInfo) HasVmInfo() bool`

HasVmInfo returns a boolean if a field has been set.

### SetVmInfoNil

`func (o *HyperflexVmBackupInfo) SetVmInfoNil(b bool)`

 SetVmInfoNil sets the value for VmInfo to be an explicit nil

### UnsetVmInfo
`func (o *HyperflexVmBackupInfo) UnsetVmInfo()`

UnsetVmInfo ensures that no value is present for VmInfo, not even an explicit nil
### GetSrcBackupCluster

`func (o *HyperflexVmBackupInfo) GetSrcBackupCluster() HyperflexBackupClusterRelationship`

GetSrcBackupCluster returns the SrcBackupCluster field if non-nil, zero value otherwise.

### GetSrcBackupClusterOk

`func (o *HyperflexVmBackupInfo) GetSrcBackupClusterOk() (*HyperflexBackupClusterRelationship, bool)`

GetSrcBackupClusterOk returns a tuple with the SrcBackupCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcBackupCluster

`func (o *HyperflexVmBackupInfo) SetSrcBackupCluster(v HyperflexBackupClusterRelationship)`

SetSrcBackupCluster sets SrcBackupCluster field to given value.

### HasSrcBackupCluster

`func (o *HyperflexVmBackupInfo) HasSrcBackupCluster() bool`

HasSrcBackupCluster returns a boolean if a field has been set.

### GetSrcCluster

`func (o *HyperflexVmBackupInfo) GetSrcCluster() HyperflexClusterRelationship`

GetSrcCluster returns the SrcCluster field if non-nil, zero value otherwise.

### GetSrcClusterOk

`func (o *HyperflexVmBackupInfo) GetSrcClusterOk() (*HyperflexClusterRelationship, bool)`

GetSrcClusterOk returns a tuple with the SrcCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcCluster

`func (o *HyperflexVmBackupInfo) SetSrcCluster(v HyperflexClusterRelationship)`

SetSrcCluster sets SrcCluster field to given value.

### HasSrcCluster

`func (o *HyperflexVmBackupInfo) HasSrcCluster() bool`

HasSrcCluster returns a boolean if a field has been set.

### GetTgtCluster

`func (o *HyperflexVmBackupInfo) GetTgtCluster() HyperflexClusterRelationship`

GetTgtCluster returns the TgtCluster field if non-nil, zero value otherwise.

### GetTgtClusterOk

`func (o *HyperflexVmBackupInfo) GetTgtClusterOk() (*HyperflexClusterRelationship, bool)`

GetTgtClusterOk returns a tuple with the TgtCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTgtCluster

`func (o *HyperflexVmBackupInfo) SetTgtCluster(v HyperflexClusterRelationship)`

SetTgtCluster sets TgtCluster field to given value.

### HasTgtCluster

`func (o *HyperflexVmBackupInfo) HasTgtCluster() bool`

HasTgtCluster returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



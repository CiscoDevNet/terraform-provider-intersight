# HyperflexVmSnapshotInfoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.VmSnapshotInfo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.VmSnapshotInfo"]
**ClusterIdSnapMap** | Pointer to [**[]HyperflexMapClusterIdToStSnapshotPoint**](HyperflexMapClusterIdToStSnapshotPoint.md) |  | [optional] 
**DisplayStatus** | Pointer to **string** | Combined status code from replication and snapshot status to display in the Intersight UI. * &#x60;NONE&#x60; - Snapshot replication state is none. * &#x60;SUCCESS&#x60; - Snapshot completed successfully. * &#x60;FAILED&#x60; - Snapshot failed replication status code. * &#x60;PAUSED&#x60; - Snapshot replication paused status code. * &#x60;IN_USE&#x60; - Snapshot replica in use status code. * &#x60;STARTING&#x60; - Snapshot replication starting. * &#x60;REPLICATING&#x60; - Snapshot replication in progress. | [optional] [readonly] [default to "NONE"]
**Error** | Pointer to [**NullableHyperflexErrorStack**](HyperflexErrorStack.md) |  | [optional] 
**Label** | Pointer to **string** | The name of the Virtual Machine and the time stamp of the snapshot. | [optional] 
**Mode** | Pointer to **string** | Quiesce Mode for snapshot taken on Hyperflex cluster. * &#x60;NONE&#x60; - The snapshot quiesce mode is none. * &#x60;CRASH&#x60; - The snapshot quiesce mode is crash. * &#x60;VMTOOLS&#x60; - The snapshot quiesce mode is VMTOOLS. * &#x60;APP_CONSISTENT&#x60; - The snapshot quiesce mode is app consistent. | [optional] [readonly] [default to "NONE"]
**ParentSnapshot** | Pointer to [**NullableHyperflexEntityReference**](HyperflexEntityReference.md) |  | [optional] 
**ReplicationStatus** | Pointer to **string** | Replication status of the least successful target cluster. * &#x60;NONE&#x60; - Snapshot replication state is none. * &#x60;SUCCESS&#x60; - Snapshot completed successfully. * &#x60;FAILED&#x60; - Snapshot failed replication status code. * &#x60;PAUSED&#x60; - Snapshot replication paused status code. * &#x60;IN_USE&#x60; - Snapshot replica in use status code. * &#x60;STARTING&#x60; - Snapshot replication starting. * &#x60;REPLICATING&#x60; - Snapshot replication in progress. | [optional] [readonly] [default to "NONE"]
**SnapshotErrorMsg** | Pointer to **string** | Error message from snapshot creation or replcation if any exist. | [optional] [readonly] 
**SnapshotStatus** | Pointer to **string** | Snapshot status of the source cluster. * &#x60;SUCCESS&#x60; - This snapshot status code is success. * &#x60;FAILED&#x60; - This snapshot status code is failed. * &#x60;IN_PROGRESS&#x60; - This snapshot status code is in progress. * &#x60;DELETING&#x60; - This snapshot status code is deleting. * &#x60;DELETED&#x60; - This snapshot status code is deleted. * &#x60;NONE&#x60; - This snapshot status code is none. * &#x60;INIT&#x60; - This snapshot status code is initializing. | [optional] [readonly] [default to "SUCCESS"]
**SourceTimestamp** | Pointer to **int64** | Timestamp the snapshot was created on the source cluster. | [optional] [readonly] 
**SrcClusterName** | Pointer to **string** | Name of the cluster this snapshot resides on. | [optional] [readonly] 
**TargetCompletionTimestamp** | Pointer to **int64** | Timestamp the snapshot was finished replicating on the target cluster. | [optional] [readonly] 
**TgtClusterName** | Pointer to **string** | Name of the cluster this snapshot is replicated to. | [optional] [readonly] 
**VmEntityReference** | Pointer to [**NullableHyperflexEntityReference**](HyperflexEntityReference.md) |  | [optional] 
**VmSnapshotEntityReference** | Pointer to [**NullableHyperflexEntityReference**](HyperflexEntityReference.md) |  | [optional] 
**SrcCluster** | Pointer to [**HyperflexClusterRelationship**](HyperflexClusterRelationship.md) |  | [optional] 
**TgtCluster** | Pointer to [**HyperflexClusterRelationship**](HyperflexClusterRelationship.md) |  | [optional] 
**VmBackupInfo** | Pointer to [**HyperflexVmBackupInfoRelationship**](HyperflexVmBackupInfoRelationship.md) |  | [optional] 

## Methods

### NewHyperflexVmSnapshotInfoAllOf

`func NewHyperflexVmSnapshotInfoAllOf(classId string, objectType string, ) *HyperflexVmSnapshotInfoAllOf`

NewHyperflexVmSnapshotInfoAllOf instantiates a new HyperflexVmSnapshotInfoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexVmSnapshotInfoAllOfWithDefaults

`func NewHyperflexVmSnapshotInfoAllOfWithDefaults() *HyperflexVmSnapshotInfoAllOf`

NewHyperflexVmSnapshotInfoAllOfWithDefaults instantiates a new HyperflexVmSnapshotInfoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexVmSnapshotInfoAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexVmSnapshotInfoAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexVmSnapshotInfoAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexVmSnapshotInfoAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexVmSnapshotInfoAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexVmSnapshotInfoAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetClusterIdSnapMap

`func (o *HyperflexVmSnapshotInfoAllOf) GetClusterIdSnapMap() []HyperflexMapClusterIdToStSnapshotPoint`

GetClusterIdSnapMap returns the ClusterIdSnapMap field if non-nil, zero value otherwise.

### GetClusterIdSnapMapOk

`func (o *HyperflexVmSnapshotInfoAllOf) GetClusterIdSnapMapOk() (*[]HyperflexMapClusterIdToStSnapshotPoint, bool)`

GetClusterIdSnapMapOk returns a tuple with the ClusterIdSnapMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterIdSnapMap

`func (o *HyperflexVmSnapshotInfoAllOf) SetClusterIdSnapMap(v []HyperflexMapClusterIdToStSnapshotPoint)`

SetClusterIdSnapMap sets ClusterIdSnapMap field to given value.

### HasClusterIdSnapMap

`func (o *HyperflexVmSnapshotInfoAllOf) HasClusterIdSnapMap() bool`

HasClusterIdSnapMap returns a boolean if a field has been set.

### SetClusterIdSnapMapNil

`func (o *HyperflexVmSnapshotInfoAllOf) SetClusterIdSnapMapNil(b bool)`

 SetClusterIdSnapMapNil sets the value for ClusterIdSnapMap to be an explicit nil

### UnsetClusterIdSnapMap
`func (o *HyperflexVmSnapshotInfoAllOf) UnsetClusterIdSnapMap()`

UnsetClusterIdSnapMap ensures that no value is present for ClusterIdSnapMap, not even an explicit nil
### GetDisplayStatus

`func (o *HyperflexVmSnapshotInfoAllOf) GetDisplayStatus() string`

GetDisplayStatus returns the DisplayStatus field if non-nil, zero value otherwise.

### GetDisplayStatusOk

`func (o *HyperflexVmSnapshotInfoAllOf) GetDisplayStatusOk() (*string, bool)`

GetDisplayStatusOk returns a tuple with the DisplayStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayStatus

`func (o *HyperflexVmSnapshotInfoAllOf) SetDisplayStatus(v string)`

SetDisplayStatus sets DisplayStatus field to given value.

### HasDisplayStatus

`func (o *HyperflexVmSnapshotInfoAllOf) HasDisplayStatus() bool`

HasDisplayStatus returns a boolean if a field has been set.

### GetError

`func (o *HyperflexVmSnapshotInfoAllOf) GetError() HyperflexErrorStack`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *HyperflexVmSnapshotInfoAllOf) GetErrorOk() (*HyperflexErrorStack, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *HyperflexVmSnapshotInfoAllOf) SetError(v HyperflexErrorStack)`

SetError sets Error field to given value.

### HasError

`func (o *HyperflexVmSnapshotInfoAllOf) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *HyperflexVmSnapshotInfoAllOf) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *HyperflexVmSnapshotInfoAllOf) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetLabel

`func (o *HyperflexVmSnapshotInfoAllOf) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *HyperflexVmSnapshotInfoAllOf) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *HyperflexVmSnapshotInfoAllOf) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *HyperflexVmSnapshotInfoAllOf) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetMode

`func (o *HyperflexVmSnapshotInfoAllOf) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *HyperflexVmSnapshotInfoAllOf) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *HyperflexVmSnapshotInfoAllOf) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *HyperflexVmSnapshotInfoAllOf) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetParentSnapshot

`func (o *HyperflexVmSnapshotInfoAllOf) GetParentSnapshot() HyperflexEntityReference`

GetParentSnapshot returns the ParentSnapshot field if non-nil, zero value otherwise.

### GetParentSnapshotOk

`func (o *HyperflexVmSnapshotInfoAllOf) GetParentSnapshotOk() (*HyperflexEntityReference, bool)`

GetParentSnapshotOk returns a tuple with the ParentSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSnapshot

`func (o *HyperflexVmSnapshotInfoAllOf) SetParentSnapshot(v HyperflexEntityReference)`

SetParentSnapshot sets ParentSnapshot field to given value.

### HasParentSnapshot

`func (o *HyperflexVmSnapshotInfoAllOf) HasParentSnapshot() bool`

HasParentSnapshot returns a boolean if a field has been set.

### SetParentSnapshotNil

`func (o *HyperflexVmSnapshotInfoAllOf) SetParentSnapshotNil(b bool)`

 SetParentSnapshotNil sets the value for ParentSnapshot to be an explicit nil

### UnsetParentSnapshot
`func (o *HyperflexVmSnapshotInfoAllOf) UnsetParentSnapshot()`

UnsetParentSnapshot ensures that no value is present for ParentSnapshot, not even an explicit nil
### GetReplicationStatus

`func (o *HyperflexVmSnapshotInfoAllOf) GetReplicationStatus() string`

GetReplicationStatus returns the ReplicationStatus field if non-nil, zero value otherwise.

### GetReplicationStatusOk

`func (o *HyperflexVmSnapshotInfoAllOf) GetReplicationStatusOk() (*string, bool)`

GetReplicationStatusOk returns a tuple with the ReplicationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationStatus

`func (o *HyperflexVmSnapshotInfoAllOf) SetReplicationStatus(v string)`

SetReplicationStatus sets ReplicationStatus field to given value.

### HasReplicationStatus

`func (o *HyperflexVmSnapshotInfoAllOf) HasReplicationStatus() bool`

HasReplicationStatus returns a boolean if a field has been set.

### GetSnapshotErrorMsg

`func (o *HyperflexVmSnapshotInfoAllOf) GetSnapshotErrorMsg() string`

GetSnapshotErrorMsg returns the SnapshotErrorMsg field if non-nil, zero value otherwise.

### GetSnapshotErrorMsgOk

`func (o *HyperflexVmSnapshotInfoAllOf) GetSnapshotErrorMsgOk() (*string, bool)`

GetSnapshotErrorMsgOk returns a tuple with the SnapshotErrorMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotErrorMsg

`func (o *HyperflexVmSnapshotInfoAllOf) SetSnapshotErrorMsg(v string)`

SetSnapshotErrorMsg sets SnapshotErrorMsg field to given value.

### HasSnapshotErrorMsg

`func (o *HyperflexVmSnapshotInfoAllOf) HasSnapshotErrorMsg() bool`

HasSnapshotErrorMsg returns a boolean if a field has been set.

### GetSnapshotStatus

`func (o *HyperflexVmSnapshotInfoAllOf) GetSnapshotStatus() string`

GetSnapshotStatus returns the SnapshotStatus field if non-nil, zero value otherwise.

### GetSnapshotStatusOk

`func (o *HyperflexVmSnapshotInfoAllOf) GetSnapshotStatusOk() (*string, bool)`

GetSnapshotStatusOk returns a tuple with the SnapshotStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotStatus

`func (o *HyperflexVmSnapshotInfoAllOf) SetSnapshotStatus(v string)`

SetSnapshotStatus sets SnapshotStatus field to given value.

### HasSnapshotStatus

`func (o *HyperflexVmSnapshotInfoAllOf) HasSnapshotStatus() bool`

HasSnapshotStatus returns a boolean if a field has been set.

### GetSourceTimestamp

`func (o *HyperflexVmSnapshotInfoAllOf) GetSourceTimestamp() int64`

GetSourceTimestamp returns the SourceTimestamp field if non-nil, zero value otherwise.

### GetSourceTimestampOk

`func (o *HyperflexVmSnapshotInfoAllOf) GetSourceTimestampOk() (*int64, bool)`

GetSourceTimestampOk returns a tuple with the SourceTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceTimestamp

`func (o *HyperflexVmSnapshotInfoAllOf) SetSourceTimestamp(v int64)`

SetSourceTimestamp sets SourceTimestamp field to given value.

### HasSourceTimestamp

`func (o *HyperflexVmSnapshotInfoAllOf) HasSourceTimestamp() bool`

HasSourceTimestamp returns a boolean if a field has been set.

### GetSrcClusterName

`func (o *HyperflexVmSnapshotInfoAllOf) GetSrcClusterName() string`

GetSrcClusterName returns the SrcClusterName field if non-nil, zero value otherwise.

### GetSrcClusterNameOk

`func (o *HyperflexVmSnapshotInfoAllOf) GetSrcClusterNameOk() (*string, bool)`

GetSrcClusterNameOk returns a tuple with the SrcClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcClusterName

`func (o *HyperflexVmSnapshotInfoAllOf) SetSrcClusterName(v string)`

SetSrcClusterName sets SrcClusterName field to given value.

### HasSrcClusterName

`func (o *HyperflexVmSnapshotInfoAllOf) HasSrcClusterName() bool`

HasSrcClusterName returns a boolean if a field has been set.

### GetTargetCompletionTimestamp

`func (o *HyperflexVmSnapshotInfoAllOf) GetTargetCompletionTimestamp() int64`

GetTargetCompletionTimestamp returns the TargetCompletionTimestamp field if non-nil, zero value otherwise.

### GetTargetCompletionTimestampOk

`func (o *HyperflexVmSnapshotInfoAllOf) GetTargetCompletionTimestampOk() (*int64, bool)`

GetTargetCompletionTimestampOk returns a tuple with the TargetCompletionTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetCompletionTimestamp

`func (o *HyperflexVmSnapshotInfoAllOf) SetTargetCompletionTimestamp(v int64)`

SetTargetCompletionTimestamp sets TargetCompletionTimestamp field to given value.

### HasTargetCompletionTimestamp

`func (o *HyperflexVmSnapshotInfoAllOf) HasTargetCompletionTimestamp() bool`

HasTargetCompletionTimestamp returns a boolean if a field has been set.

### GetTgtClusterName

`func (o *HyperflexVmSnapshotInfoAllOf) GetTgtClusterName() string`

GetTgtClusterName returns the TgtClusterName field if non-nil, zero value otherwise.

### GetTgtClusterNameOk

`func (o *HyperflexVmSnapshotInfoAllOf) GetTgtClusterNameOk() (*string, bool)`

GetTgtClusterNameOk returns a tuple with the TgtClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTgtClusterName

`func (o *HyperflexVmSnapshotInfoAllOf) SetTgtClusterName(v string)`

SetTgtClusterName sets TgtClusterName field to given value.

### HasTgtClusterName

`func (o *HyperflexVmSnapshotInfoAllOf) HasTgtClusterName() bool`

HasTgtClusterName returns a boolean if a field has been set.

### GetVmEntityReference

`func (o *HyperflexVmSnapshotInfoAllOf) GetVmEntityReference() HyperflexEntityReference`

GetVmEntityReference returns the VmEntityReference field if non-nil, zero value otherwise.

### GetVmEntityReferenceOk

`func (o *HyperflexVmSnapshotInfoAllOf) GetVmEntityReferenceOk() (*HyperflexEntityReference, bool)`

GetVmEntityReferenceOk returns a tuple with the VmEntityReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmEntityReference

`func (o *HyperflexVmSnapshotInfoAllOf) SetVmEntityReference(v HyperflexEntityReference)`

SetVmEntityReference sets VmEntityReference field to given value.

### HasVmEntityReference

`func (o *HyperflexVmSnapshotInfoAllOf) HasVmEntityReference() bool`

HasVmEntityReference returns a boolean if a field has been set.

### SetVmEntityReferenceNil

`func (o *HyperflexVmSnapshotInfoAllOf) SetVmEntityReferenceNil(b bool)`

 SetVmEntityReferenceNil sets the value for VmEntityReference to be an explicit nil

### UnsetVmEntityReference
`func (o *HyperflexVmSnapshotInfoAllOf) UnsetVmEntityReference()`

UnsetVmEntityReference ensures that no value is present for VmEntityReference, not even an explicit nil
### GetVmSnapshotEntityReference

`func (o *HyperflexVmSnapshotInfoAllOf) GetVmSnapshotEntityReference() HyperflexEntityReference`

GetVmSnapshotEntityReference returns the VmSnapshotEntityReference field if non-nil, zero value otherwise.

### GetVmSnapshotEntityReferenceOk

`func (o *HyperflexVmSnapshotInfoAllOf) GetVmSnapshotEntityReferenceOk() (*HyperflexEntityReference, bool)`

GetVmSnapshotEntityReferenceOk returns a tuple with the VmSnapshotEntityReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmSnapshotEntityReference

`func (o *HyperflexVmSnapshotInfoAllOf) SetVmSnapshotEntityReference(v HyperflexEntityReference)`

SetVmSnapshotEntityReference sets VmSnapshotEntityReference field to given value.

### HasVmSnapshotEntityReference

`func (o *HyperflexVmSnapshotInfoAllOf) HasVmSnapshotEntityReference() bool`

HasVmSnapshotEntityReference returns a boolean if a field has been set.

### SetVmSnapshotEntityReferenceNil

`func (o *HyperflexVmSnapshotInfoAllOf) SetVmSnapshotEntityReferenceNil(b bool)`

 SetVmSnapshotEntityReferenceNil sets the value for VmSnapshotEntityReference to be an explicit nil

### UnsetVmSnapshotEntityReference
`func (o *HyperflexVmSnapshotInfoAllOf) UnsetVmSnapshotEntityReference()`

UnsetVmSnapshotEntityReference ensures that no value is present for VmSnapshotEntityReference, not even an explicit nil
### GetSrcCluster

`func (o *HyperflexVmSnapshotInfoAllOf) GetSrcCluster() HyperflexClusterRelationship`

GetSrcCluster returns the SrcCluster field if non-nil, zero value otherwise.

### GetSrcClusterOk

`func (o *HyperflexVmSnapshotInfoAllOf) GetSrcClusterOk() (*HyperflexClusterRelationship, bool)`

GetSrcClusterOk returns a tuple with the SrcCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcCluster

`func (o *HyperflexVmSnapshotInfoAllOf) SetSrcCluster(v HyperflexClusterRelationship)`

SetSrcCluster sets SrcCluster field to given value.

### HasSrcCluster

`func (o *HyperflexVmSnapshotInfoAllOf) HasSrcCluster() bool`

HasSrcCluster returns a boolean if a field has been set.

### GetTgtCluster

`func (o *HyperflexVmSnapshotInfoAllOf) GetTgtCluster() HyperflexClusterRelationship`

GetTgtCluster returns the TgtCluster field if non-nil, zero value otherwise.

### GetTgtClusterOk

`func (o *HyperflexVmSnapshotInfoAllOf) GetTgtClusterOk() (*HyperflexClusterRelationship, bool)`

GetTgtClusterOk returns a tuple with the TgtCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTgtCluster

`func (o *HyperflexVmSnapshotInfoAllOf) SetTgtCluster(v HyperflexClusterRelationship)`

SetTgtCluster sets TgtCluster field to given value.

### HasTgtCluster

`func (o *HyperflexVmSnapshotInfoAllOf) HasTgtCluster() bool`

HasTgtCluster returns a boolean if a field has been set.

### GetVmBackupInfo

`func (o *HyperflexVmSnapshotInfoAllOf) GetVmBackupInfo() HyperflexVmBackupInfoRelationship`

GetVmBackupInfo returns the VmBackupInfo field if non-nil, zero value otherwise.

### GetVmBackupInfoOk

`func (o *HyperflexVmSnapshotInfoAllOf) GetVmBackupInfoOk() (*HyperflexVmBackupInfoRelationship, bool)`

GetVmBackupInfoOk returns a tuple with the VmBackupInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmBackupInfo

`func (o *HyperflexVmSnapshotInfoAllOf) SetVmBackupInfo(v HyperflexVmBackupInfoRelationship)`

SetVmBackupInfo sets VmBackupInfo field to given value.

### HasVmBackupInfo

`func (o *HyperflexVmSnapshotInfoAllOf) HasVmBackupInfo() bool`

HasVmBackupInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



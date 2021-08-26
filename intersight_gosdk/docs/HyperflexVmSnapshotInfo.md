# HyperflexVmSnapshotInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.VmSnapshotInfo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.VmSnapshotInfo"]
**ClusterIdSnapMap** | Pointer to [**[]HyperflexMapClusterIdToStSnapshotPoint**](HyperflexMapClusterIdToStSnapshotPoint.md) |  | [optional] 
**ErrorStack** | Pointer to [**NullableHyperflexErrorStack**](HyperflexErrorStack.md) |  | [optional] 
**Label** | Pointer to **string** | The name of the Virtual Machine and the time stamp of the snapshot. | [optional] 
**Mode** | Pointer to **string** | Quiesce Mode for snapshot taken on Hyperflex cluster. * &#x60;NONE&#x60; - The snapshot quiesce mode is none. * &#x60;CRASH&#x60; - The snapshot quiesce mode is crash. * &#x60;VMTOOLS&#x60; - The snapshot quiesce mode is VMTOOLS. * &#x60;APP_CONSISTENT&#x60; - The snapshot quiesce mode is app consistent. | [optional] [readonly] [default to "NONE"]
**ParentSnapshot** | Pointer to [**NullableHyperflexEntityReference**](HyperflexEntityReference.md) |  | [optional] 
**ReplicationStatus** | Pointer to **string** | Replication status of the least successful target cluster. * &#x60;NONE&#x60; - Snapshot replication state is none. * &#x60;SUCCESS&#x60; - Snapshot completed successfully. * &#x60;FAILED&#x60; - Snapshot failed replication status code. * &#x60;PAUSED&#x60; - Snapshot replication paused status code. * &#x60;IN_USE&#x60; - Snapshot replica in use status code. * &#x60;STARTING&#x60; - Snapshot replication starting. * &#x60;REPLICATING&#x60; - Snapshot replication in progress. | [optional] [readonly] [default to "NONE"]
**SnapshotStatus** | Pointer to **string** | Snapshot status of the source cluster. * &#x60;SUCCESS&#x60; - This snapshot status code is success. * &#x60;FAILED&#x60; - This snapshot status code is failed. * &#x60;IN_PROGRESS&#x60; - This snapshot status code is in progress. * &#x60;DELETING&#x60; - This snapshot status code is deleting. * &#x60;DELETED&#x60; - This snapshot status code is deleted. * &#x60;NONE&#x60; - This snapshot status code is none. | [optional] [readonly] [default to "SUCCESS"]
**SourceTimestamp** | Pointer to **int64** | Timestamp the snapshot was created on the source cluster. | [optional] [readonly] 
**VmEntityReference** | Pointer to [**NullableHyperflexEntityReference**](HyperflexEntityReference.md) |  | [optional] 
**VmSnapshotEntityReference** | Pointer to [**NullableHyperflexEntityReference**](HyperflexEntityReference.md) |  | [optional] 
**SrcCluster** | Pointer to [**HyperflexClusterRelationship**](HyperflexClusterRelationship.md) |  | [optional] 
**TgtCluster** | Pointer to [**HyperflexClusterRelationship**](HyperflexClusterRelationship.md) |  | [optional] 
**VmBackupInfo** | Pointer to [**HyperflexVmBackupInfoRelationship**](HyperflexVmBackupInfoRelationship.md) |  | [optional] 

## Methods

### NewHyperflexVmSnapshotInfo

`func NewHyperflexVmSnapshotInfo(classId string, objectType string, ) *HyperflexVmSnapshotInfo`

NewHyperflexVmSnapshotInfo instantiates a new HyperflexVmSnapshotInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexVmSnapshotInfoWithDefaults

`func NewHyperflexVmSnapshotInfoWithDefaults() *HyperflexVmSnapshotInfo`

NewHyperflexVmSnapshotInfoWithDefaults instantiates a new HyperflexVmSnapshotInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexVmSnapshotInfo) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexVmSnapshotInfo) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexVmSnapshotInfo) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexVmSnapshotInfo) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexVmSnapshotInfo) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexVmSnapshotInfo) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetClusterIdSnapMap

`func (o *HyperflexVmSnapshotInfo) GetClusterIdSnapMap() []HyperflexMapClusterIdToStSnapshotPoint`

GetClusterIdSnapMap returns the ClusterIdSnapMap field if non-nil, zero value otherwise.

### GetClusterIdSnapMapOk

`func (o *HyperflexVmSnapshotInfo) GetClusterIdSnapMapOk() (*[]HyperflexMapClusterIdToStSnapshotPoint, bool)`

GetClusterIdSnapMapOk returns a tuple with the ClusterIdSnapMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterIdSnapMap

`func (o *HyperflexVmSnapshotInfo) SetClusterIdSnapMap(v []HyperflexMapClusterIdToStSnapshotPoint)`

SetClusterIdSnapMap sets ClusterIdSnapMap field to given value.

### HasClusterIdSnapMap

`func (o *HyperflexVmSnapshotInfo) HasClusterIdSnapMap() bool`

HasClusterIdSnapMap returns a boolean if a field has been set.

### SetClusterIdSnapMapNil

`func (o *HyperflexVmSnapshotInfo) SetClusterIdSnapMapNil(b bool)`

 SetClusterIdSnapMapNil sets the value for ClusterIdSnapMap to be an explicit nil

### UnsetClusterIdSnapMap
`func (o *HyperflexVmSnapshotInfo) UnsetClusterIdSnapMap()`

UnsetClusterIdSnapMap ensures that no value is present for ClusterIdSnapMap, not even an explicit nil
### GetErrorStack

`func (o *HyperflexVmSnapshotInfo) GetErrorStack() HyperflexErrorStack`

GetErrorStack returns the ErrorStack field if non-nil, zero value otherwise.

### GetErrorStackOk

`func (o *HyperflexVmSnapshotInfo) GetErrorStackOk() (*HyperflexErrorStack, bool)`

GetErrorStackOk returns a tuple with the ErrorStack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorStack

`func (o *HyperflexVmSnapshotInfo) SetErrorStack(v HyperflexErrorStack)`

SetErrorStack sets ErrorStack field to given value.

### HasErrorStack

`func (o *HyperflexVmSnapshotInfo) HasErrorStack() bool`

HasErrorStack returns a boolean if a field has been set.

### SetErrorStackNil

`func (o *HyperflexVmSnapshotInfo) SetErrorStackNil(b bool)`

 SetErrorStackNil sets the value for ErrorStack to be an explicit nil

### UnsetErrorStack
`func (o *HyperflexVmSnapshotInfo) UnsetErrorStack()`

UnsetErrorStack ensures that no value is present for ErrorStack, not even an explicit nil
### GetLabel

`func (o *HyperflexVmSnapshotInfo) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *HyperflexVmSnapshotInfo) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *HyperflexVmSnapshotInfo) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *HyperflexVmSnapshotInfo) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetMode

`func (o *HyperflexVmSnapshotInfo) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *HyperflexVmSnapshotInfo) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *HyperflexVmSnapshotInfo) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *HyperflexVmSnapshotInfo) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetParentSnapshot

`func (o *HyperflexVmSnapshotInfo) GetParentSnapshot() HyperflexEntityReference`

GetParentSnapshot returns the ParentSnapshot field if non-nil, zero value otherwise.

### GetParentSnapshotOk

`func (o *HyperflexVmSnapshotInfo) GetParentSnapshotOk() (*HyperflexEntityReference, bool)`

GetParentSnapshotOk returns a tuple with the ParentSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSnapshot

`func (o *HyperflexVmSnapshotInfo) SetParentSnapshot(v HyperflexEntityReference)`

SetParentSnapshot sets ParentSnapshot field to given value.

### HasParentSnapshot

`func (o *HyperflexVmSnapshotInfo) HasParentSnapshot() bool`

HasParentSnapshot returns a boolean if a field has been set.

### SetParentSnapshotNil

`func (o *HyperflexVmSnapshotInfo) SetParentSnapshotNil(b bool)`

 SetParentSnapshotNil sets the value for ParentSnapshot to be an explicit nil

### UnsetParentSnapshot
`func (o *HyperflexVmSnapshotInfo) UnsetParentSnapshot()`

UnsetParentSnapshot ensures that no value is present for ParentSnapshot, not even an explicit nil
### GetReplicationStatus

`func (o *HyperflexVmSnapshotInfo) GetReplicationStatus() string`

GetReplicationStatus returns the ReplicationStatus field if non-nil, zero value otherwise.

### GetReplicationStatusOk

`func (o *HyperflexVmSnapshotInfo) GetReplicationStatusOk() (*string, bool)`

GetReplicationStatusOk returns a tuple with the ReplicationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationStatus

`func (o *HyperflexVmSnapshotInfo) SetReplicationStatus(v string)`

SetReplicationStatus sets ReplicationStatus field to given value.

### HasReplicationStatus

`func (o *HyperflexVmSnapshotInfo) HasReplicationStatus() bool`

HasReplicationStatus returns a boolean if a field has been set.

### GetSnapshotStatus

`func (o *HyperflexVmSnapshotInfo) GetSnapshotStatus() string`

GetSnapshotStatus returns the SnapshotStatus field if non-nil, zero value otherwise.

### GetSnapshotStatusOk

`func (o *HyperflexVmSnapshotInfo) GetSnapshotStatusOk() (*string, bool)`

GetSnapshotStatusOk returns a tuple with the SnapshotStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotStatus

`func (o *HyperflexVmSnapshotInfo) SetSnapshotStatus(v string)`

SetSnapshotStatus sets SnapshotStatus field to given value.

### HasSnapshotStatus

`func (o *HyperflexVmSnapshotInfo) HasSnapshotStatus() bool`

HasSnapshotStatus returns a boolean if a field has been set.

### GetSourceTimestamp

`func (o *HyperflexVmSnapshotInfo) GetSourceTimestamp() int64`

GetSourceTimestamp returns the SourceTimestamp field if non-nil, zero value otherwise.

### GetSourceTimestampOk

`func (o *HyperflexVmSnapshotInfo) GetSourceTimestampOk() (*int64, bool)`

GetSourceTimestampOk returns a tuple with the SourceTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceTimestamp

`func (o *HyperflexVmSnapshotInfo) SetSourceTimestamp(v int64)`

SetSourceTimestamp sets SourceTimestamp field to given value.

### HasSourceTimestamp

`func (o *HyperflexVmSnapshotInfo) HasSourceTimestamp() bool`

HasSourceTimestamp returns a boolean if a field has been set.

### GetVmEntityReference

`func (o *HyperflexVmSnapshotInfo) GetVmEntityReference() HyperflexEntityReference`

GetVmEntityReference returns the VmEntityReference field if non-nil, zero value otherwise.

### GetVmEntityReferenceOk

`func (o *HyperflexVmSnapshotInfo) GetVmEntityReferenceOk() (*HyperflexEntityReference, bool)`

GetVmEntityReferenceOk returns a tuple with the VmEntityReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmEntityReference

`func (o *HyperflexVmSnapshotInfo) SetVmEntityReference(v HyperflexEntityReference)`

SetVmEntityReference sets VmEntityReference field to given value.

### HasVmEntityReference

`func (o *HyperflexVmSnapshotInfo) HasVmEntityReference() bool`

HasVmEntityReference returns a boolean if a field has been set.

### SetVmEntityReferenceNil

`func (o *HyperflexVmSnapshotInfo) SetVmEntityReferenceNil(b bool)`

 SetVmEntityReferenceNil sets the value for VmEntityReference to be an explicit nil

### UnsetVmEntityReference
`func (o *HyperflexVmSnapshotInfo) UnsetVmEntityReference()`

UnsetVmEntityReference ensures that no value is present for VmEntityReference, not even an explicit nil
### GetVmSnapshotEntityReference

`func (o *HyperflexVmSnapshotInfo) GetVmSnapshotEntityReference() HyperflexEntityReference`

GetVmSnapshotEntityReference returns the VmSnapshotEntityReference field if non-nil, zero value otherwise.

### GetVmSnapshotEntityReferenceOk

`func (o *HyperflexVmSnapshotInfo) GetVmSnapshotEntityReferenceOk() (*HyperflexEntityReference, bool)`

GetVmSnapshotEntityReferenceOk returns a tuple with the VmSnapshotEntityReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmSnapshotEntityReference

`func (o *HyperflexVmSnapshotInfo) SetVmSnapshotEntityReference(v HyperflexEntityReference)`

SetVmSnapshotEntityReference sets VmSnapshotEntityReference field to given value.

### HasVmSnapshotEntityReference

`func (o *HyperflexVmSnapshotInfo) HasVmSnapshotEntityReference() bool`

HasVmSnapshotEntityReference returns a boolean if a field has been set.

### SetVmSnapshotEntityReferenceNil

`func (o *HyperflexVmSnapshotInfo) SetVmSnapshotEntityReferenceNil(b bool)`

 SetVmSnapshotEntityReferenceNil sets the value for VmSnapshotEntityReference to be an explicit nil

### UnsetVmSnapshotEntityReference
`func (o *HyperflexVmSnapshotInfo) UnsetVmSnapshotEntityReference()`

UnsetVmSnapshotEntityReference ensures that no value is present for VmSnapshotEntityReference, not even an explicit nil
### GetSrcCluster

`func (o *HyperflexVmSnapshotInfo) GetSrcCluster() HyperflexClusterRelationship`

GetSrcCluster returns the SrcCluster field if non-nil, zero value otherwise.

### GetSrcClusterOk

`func (o *HyperflexVmSnapshotInfo) GetSrcClusterOk() (*HyperflexClusterRelationship, bool)`

GetSrcClusterOk returns a tuple with the SrcCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcCluster

`func (o *HyperflexVmSnapshotInfo) SetSrcCluster(v HyperflexClusterRelationship)`

SetSrcCluster sets SrcCluster field to given value.

### HasSrcCluster

`func (o *HyperflexVmSnapshotInfo) HasSrcCluster() bool`

HasSrcCluster returns a boolean if a field has been set.

### GetTgtCluster

`func (o *HyperflexVmSnapshotInfo) GetTgtCluster() HyperflexClusterRelationship`

GetTgtCluster returns the TgtCluster field if non-nil, zero value otherwise.

### GetTgtClusterOk

`func (o *HyperflexVmSnapshotInfo) GetTgtClusterOk() (*HyperflexClusterRelationship, bool)`

GetTgtClusterOk returns a tuple with the TgtCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTgtCluster

`func (o *HyperflexVmSnapshotInfo) SetTgtCluster(v HyperflexClusterRelationship)`

SetTgtCluster sets TgtCluster field to given value.

### HasTgtCluster

`func (o *HyperflexVmSnapshotInfo) HasTgtCluster() bool`

HasTgtCluster returns a boolean if a field has been set.

### GetVmBackupInfo

`func (o *HyperflexVmSnapshotInfo) GetVmBackupInfo() HyperflexVmBackupInfoRelationship`

GetVmBackupInfo returns the VmBackupInfo field if non-nil, zero value otherwise.

### GetVmBackupInfoOk

`func (o *HyperflexVmSnapshotInfo) GetVmBackupInfoOk() (*HyperflexVmBackupInfoRelationship, bool)`

GetVmBackupInfoOk returns a tuple with the VmBackupInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmBackupInfo

`func (o *HyperflexVmSnapshotInfo) SetVmBackupInfo(v HyperflexVmBackupInfoRelationship)`

SetVmBackupInfo sets VmBackupInfo field to given value.

### HasVmBackupInfo

`func (o *HyperflexVmSnapshotInfo) HasVmBackupInfo() bool`

HasVmBackupInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



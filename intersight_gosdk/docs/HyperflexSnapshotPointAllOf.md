# HyperflexSnapshotPointAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.SnapshotPoint"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.SnapshotPoint"]
**ClusterEntityReference** | Pointer to [**NullableHyperflexEntityReference**](HyperflexEntityReference.md) |  | [optional] 
**ReplicationStatus** | Pointer to [**NullableHyperflexReplicationStatus**](HyperflexReplicationStatus.md) |  | [optional] 
**SnapshotFiles** | Pointer to [**NullableHyperflexSnapshotFiles**](HyperflexSnapshotFiles.md) |  | [optional] 
**SnapshotPointEntityReference** | Pointer to [**NullableHyperflexEntityReference**](HyperflexEntityReference.md) |  | [optional] 
**SnapshotStatus** | Pointer to [**NullableHyperflexSnapshotStatus**](HyperflexSnapshotStatus.md) |  | [optional] 
**VmRuntimeInfo** | Pointer to [**NullableHyperflexVirtualMachine**](HyperflexVirtualMachine.md) |  | [optional] 

## Methods

### NewHyperflexSnapshotPointAllOf

`func NewHyperflexSnapshotPointAllOf(classId string, objectType string, ) *HyperflexSnapshotPointAllOf`

NewHyperflexSnapshotPointAllOf instantiates a new HyperflexSnapshotPointAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexSnapshotPointAllOfWithDefaults

`func NewHyperflexSnapshotPointAllOfWithDefaults() *HyperflexSnapshotPointAllOf`

NewHyperflexSnapshotPointAllOfWithDefaults instantiates a new HyperflexSnapshotPointAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexSnapshotPointAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexSnapshotPointAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexSnapshotPointAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexSnapshotPointAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexSnapshotPointAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexSnapshotPointAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetClusterEntityReference

`func (o *HyperflexSnapshotPointAllOf) GetClusterEntityReference() HyperflexEntityReference`

GetClusterEntityReference returns the ClusterEntityReference field if non-nil, zero value otherwise.

### GetClusterEntityReferenceOk

`func (o *HyperflexSnapshotPointAllOf) GetClusterEntityReferenceOk() (*HyperflexEntityReference, bool)`

GetClusterEntityReferenceOk returns a tuple with the ClusterEntityReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterEntityReference

`func (o *HyperflexSnapshotPointAllOf) SetClusterEntityReference(v HyperflexEntityReference)`

SetClusterEntityReference sets ClusterEntityReference field to given value.

### HasClusterEntityReference

`func (o *HyperflexSnapshotPointAllOf) HasClusterEntityReference() bool`

HasClusterEntityReference returns a boolean if a field has been set.

### SetClusterEntityReferenceNil

`func (o *HyperflexSnapshotPointAllOf) SetClusterEntityReferenceNil(b bool)`

 SetClusterEntityReferenceNil sets the value for ClusterEntityReference to be an explicit nil

### UnsetClusterEntityReference
`func (o *HyperflexSnapshotPointAllOf) UnsetClusterEntityReference()`

UnsetClusterEntityReference ensures that no value is present for ClusterEntityReference, not even an explicit nil
### GetReplicationStatus

`func (o *HyperflexSnapshotPointAllOf) GetReplicationStatus() HyperflexReplicationStatus`

GetReplicationStatus returns the ReplicationStatus field if non-nil, zero value otherwise.

### GetReplicationStatusOk

`func (o *HyperflexSnapshotPointAllOf) GetReplicationStatusOk() (*HyperflexReplicationStatus, bool)`

GetReplicationStatusOk returns a tuple with the ReplicationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationStatus

`func (o *HyperflexSnapshotPointAllOf) SetReplicationStatus(v HyperflexReplicationStatus)`

SetReplicationStatus sets ReplicationStatus field to given value.

### HasReplicationStatus

`func (o *HyperflexSnapshotPointAllOf) HasReplicationStatus() bool`

HasReplicationStatus returns a boolean if a field has been set.

### SetReplicationStatusNil

`func (o *HyperflexSnapshotPointAllOf) SetReplicationStatusNil(b bool)`

 SetReplicationStatusNil sets the value for ReplicationStatus to be an explicit nil

### UnsetReplicationStatus
`func (o *HyperflexSnapshotPointAllOf) UnsetReplicationStatus()`

UnsetReplicationStatus ensures that no value is present for ReplicationStatus, not even an explicit nil
### GetSnapshotFiles

`func (o *HyperflexSnapshotPointAllOf) GetSnapshotFiles() HyperflexSnapshotFiles`

GetSnapshotFiles returns the SnapshotFiles field if non-nil, zero value otherwise.

### GetSnapshotFilesOk

`func (o *HyperflexSnapshotPointAllOf) GetSnapshotFilesOk() (*HyperflexSnapshotFiles, bool)`

GetSnapshotFilesOk returns a tuple with the SnapshotFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotFiles

`func (o *HyperflexSnapshotPointAllOf) SetSnapshotFiles(v HyperflexSnapshotFiles)`

SetSnapshotFiles sets SnapshotFiles field to given value.

### HasSnapshotFiles

`func (o *HyperflexSnapshotPointAllOf) HasSnapshotFiles() bool`

HasSnapshotFiles returns a boolean if a field has been set.

### SetSnapshotFilesNil

`func (o *HyperflexSnapshotPointAllOf) SetSnapshotFilesNil(b bool)`

 SetSnapshotFilesNil sets the value for SnapshotFiles to be an explicit nil

### UnsetSnapshotFiles
`func (o *HyperflexSnapshotPointAllOf) UnsetSnapshotFiles()`

UnsetSnapshotFiles ensures that no value is present for SnapshotFiles, not even an explicit nil
### GetSnapshotPointEntityReference

`func (o *HyperflexSnapshotPointAllOf) GetSnapshotPointEntityReference() HyperflexEntityReference`

GetSnapshotPointEntityReference returns the SnapshotPointEntityReference field if non-nil, zero value otherwise.

### GetSnapshotPointEntityReferenceOk

`func (o *HyperflexSnapshotPointAllOf) GetSnapshotPointEntityReferenceOk() (*HyperflexEntityReference, bool)`

GetSnapshotPointEntityReferenceOk returns a tuple with the SnapshotPointEntityReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotPointEntityReference

`func (o *HyperflexSnapshotPointAllOf) SetSnapshotPointEntityReference(v HyperflexEntityReference)`

SetSnapshotPointEntityReference sets SnapshotPointEntityReference field to given value.

### HasSnapshotPointEntityReference

`func (o *HyperflexSnapshotPointAllOf) HasSnapshotPointEntityReference() bool`

HasSnapshotPointEntityReference returns a boolean if a field has been set.

### SetSnapshotPointEntityReferenceNil

`func (o *HyperflexSnapshotPointAllOf) SetSnapshotPointEntityReferenceNil(b bool)`

 SetSnapshotPointEntityReferenceNil sets the value for SnapshotPointEntityReference to be an explicit nil

### UnsetSnapshotPointEntityReference
`func (o *HyperflexSnapshotPointAllOf) UnsetSnapshotPointEntityReference()`

UnsetSnapshotPointEntityReference ensures that no value is present for SnapshotPointEntityReference, not even an explicit nil
### GetSnapshotStatus

`func (o *HyperflexSnapshotPointAllOf) GetSnapshotStatus() HyperflexSnapshotStatus`

GetSnapshotStatus returns the SnapshotStatus field if non-nil, zero value otherwise.

### GetSnapshotStatusOk

`func (o *HyperflexSnapshotPointAllOf) GetSnapshotStatusOk() (*HyperflexSnapshotStatus, bool)`

GetSnapshotStatusOk returns a tuple with the SnapshotStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotStatus

`func (o *HyperflexSnapshotPointAllOf) SetSnapshotStatus(v HyperflexSnapshotStatus)`

SetSnapshotStatus sets SnapshotStatus field to given value.

### HasSnapshotStatus

`func (o *HyperflexSnapshotPointAllOf) HasSnapshotStatus() bool`

HasSnapshotStatus returns a boolean if a field has been set.

### SetSnapshotStatusNil

`func (o *HyperflexSnapshotPointAllOf) SetSnapshotStatusNil(b bool)`

 SetSnapshotStatusNil sets the value for SnapshotStatus to be an explicit nil

### UnsetSnapshotStatus
`func (o *HyperflexSnapshotPointAllOf) UnsetSnapshotStatus()`

UnsetSnapshotStatus ensures that no value is present for SnapshotStatus, not even an explicit nil
### GetVmRuntimeInfo

`func (o *HyperflexSnapshotPointAllOf) GetVmRuntimeInfo() HyperflexVirtualMachine`

GetVmRuntimeInfo returns the VmRuntimeInfo field if non-nil, zero value otherwise.

### GetVmRuntimeInfoOk

`func (o *HyperflexSnapshotPointAllOf) GetVmRuntimeInfoOk() (*HyperflexVirtualMachine, bool)`

GetVmRuntimeInfoOk returns a tuple with the VmRuntimeInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmRuntimeInfo

`func (o *HyperflexSnapshotPointAllOf) SetVmRuntimeInfo(v HyperflexVirtualMachine)`

SetVmRuntimeInfo sets VmRuntimeInfo field to given value.

### HasVmRuntimeInfo

`func (o *HyperflexSnapshotPointAllOf) HasVmRuntimeInfo() bool`

HasVmRuntimeInfo returns a boolean if a field has been set.

### SetVmRuntimeInfoNil

`func (o *HyperflexSnapshotPointAllOf) SetVmRuntimeInfoNil(b bool)`

 SetVmRuntimeInfoNil sets the value for VmRuntimeInfo to be an explicit nil

### UnsetVmRuntimeInfo
`func (o *HyperflexSnapshotPointAllOf) UnsetVmRuntimeInfo()`

UnsetVmRuntimeInfo ensures that no value is present for VmRuntimeInfo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



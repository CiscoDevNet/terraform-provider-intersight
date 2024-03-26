# StorageNetAppBaseSnapMirrorPolicyAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**Comment** | Pointer to **string** | Comment associated with the policy. | [optional] [readonly] 
**CopyAllSourceSnapshots** | Pointer to **bool** | Specifies whether all source Snapshot copies should be copied to the destination on a transfer rather than specifying specific retentions. It is applicable only to async policies. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the NetApp SnapMirror policy. | [optional] [readonly] 
**Scope** | Pointer to **string** | Identifies whether the SnapMirror policy is owned by the storage virtual machine or the cluster. | [optional] [readonly] 
**SyncType** | Pointer to **string** | SnapMirror policy sync_type is either sync, strict_sync, or automated_failover. Property is applicable only to the policies of type \&quot;sync\&quot;. | [optional] [readonly] 
**TransferScheduleName** | Pointer to **string** | Name of the schedule used to update asynchronous relationships. | [optional] [readonly] 
**TransferScheduleUuid** | Pointer to **string** | Uuid of the schedule used to update asynchronous relationships. | [optional] [readonly] 
**Type** | Pointer to **string** | SnapMirror policy type can be async, sync, or continuous. | [optional] [readonly] 
**Uuid** | Pointer to **string** | Uuid of the NetApp SnapMirror policy. | [optional] [readonly] 

## Methods

### NewStorageNetAppBaseSnapMirrorPolicyAllOf

`func NewStorageNetAppBaseSnapMirrorPolicyAllOf(classId string, objectType string, ) *StorageNetAppBaseSnapMirrorPolicyAllOf`

NewStorageNetAppBaseSnapMirrorPolicyAllOf instantiates a new StorageNetAppBaseSnapMirrorPolicyAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNetAppBaseSnapMirrorPolicyAllOfWithDefaults

`func NewStorageNetAppBaseSnapMirrorPolicyAllOfWithDefaults() *StorageNetAppBaseSnapMirrorPolicyAllOf`

NewStorageNetAppBaseSnapMirrorPolicyAllOfWithDefaults instantiates a new StorageNetAppBaseSnapMirrorPolicyAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageNetAppBaseSnapMirrorPolicyAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageNetAppBaseSnapMirrorPolicyAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageNetAppBaseSnapMirrorPolicyAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageNetAppBaseSnapMirrorPolicyAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageNetAppBaseSnapMirrorPolicyAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageNetAppBaseSnapMirrorPolicyAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetComment

`func (o *StorageNetAppBaseSnapMirrorPolicyAllOf) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *StorageNetAppBaseSnapMirrorPolicyAllOf) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *StorageNetAppBaseSnapMirrorPolicyAllOf) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *StorageNetAppBaseSnapMirrorPolicyAllOf) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCopyAllSourceSnapshots

`func (o *StorageNetAppBaseSnapMirrorPolicyAllOf) GetCopyAllSourceSnapshots() bool`

GetCopyAllSourceSnapshots returns the CopyAllSourceSnapshots field if non-nil, zero value otherwise.

### GetCopyAllSourceSnapshotsOk

`func (o *StorageNetAppBaseSnapMirrorPolicyAllOf) GetCopyAllSourceSnapshotsOk() (*bool, bool)`

GetCopyAllSourceSnapshotsOk returns a tuple with the CopyAllSourceSnapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyAllSourceSnapshots

`func (o *StorageNetAppBaseSnapMirrorPolicyAllOf) SetCopyAllSourceSnapshots(v bool)`

SetCopyAllSourceSnapshots sets CopyAllSourceSnapshots field to given value.

### HasCopyAllSourceSnapshots

`func (o *StorageNetAppBaseSnapMirrorPolicyAllOf) HasCopyAllSourceSnapshots() bool`

HasCopyAllSourceSnapshots returns a boolean if a field has been set.

### GetName

`func (o *StorageNetAppBaseSnapMirrorPolicyAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageNetAppBaseSnapMirrorPolicyAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageNetAppBaseSnapMirrorPolicyAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StorageNetAppBaseSnapMirrorPolicyAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetScope

`func (o *StorageNetAppBaseSnapMirrorPolicyAllOf) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *StorageNetAppBaseSnapMirrorPolicyAllOf) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *StorageNetAppBaseSnapMirrorPolicyAllOf) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *StorageNetAppBaseSnapMirrorPolicyAllOf) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetSyncType

`func (o *StorageNetAppBaseSnapMirrorPolicyAllOf) GetSyncType() string`

GetSyncType returns the SyncType field if non-nil, zero value otherwise.

### GetSyncTypeOk

`func (o *StorageNetAppBaseSnapMirrorPolicyAllOf) GetSyncTypeOk() (*string, bool)`

GetSyncTypeOk returns a tuple with the SyncType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncType

`func (o *StorageNetAppBaseSnapMirrorPolicyAllOf) SetSyncType(v string)`

SetSyncType sets SyncType field to given value.

### HasSyncType

`func (o *StorageNetAppBaseSnapMirrorPolicyAllOf) HasSyncType() bool`

HasSyncType returns a boolean if a field has been set.

### GetTransferScheduleName

`func (o *StorageNetAppBaseSnapMirrorPolicyAllOf) GetTransferScheduleName() string`

GetTransferScheduleName returns the TransferScheduleName field if non-nil, zero value otherwise.

### GetTransferScheduleNameOk

`func (o *StorageNetAppBaseSnapMirrorPolicyAllOf) GetTransferScheduleNameOk() (*string, bool)`

GetTransferScheduleNameOk returns a tuple with the TransferScheduleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferScheduleName

`func (o *StorageNetAppBaseSnapMirrorPolicyAllOf) SetTransferScheduleName(v string)`

SetTransferScheduleName sets TransferScheduleName field to given value.

### HasTransferScheduleName

`func (o *StorageNetAppBaseSnapMirrorPolicyAllOf) HasTransferScheduleName() bool`

HasTransferScheduleName returns a boolean if a field has been set.

### GetTransferScheduleUuid

`func (o *StorageNetAppBaseSnapMirrorPolicyAllOf) GetTransferScheduleUuid() string`

GetTransferScheduleUuid returns the TransferScheduleUuid field if non-nil, zero value otherwise.

### GetTransferScheduleUuidOk

`func (o *StorageNetAppBaseSnapMirrorPolicyAllOf) GetTransferScheduleUuidOk() (*string, bool)`

GetTransferScheduleUuidOk returns a tuple with the TransferScheduleUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferScheduleUuid

`func (o *StorageNetAppBaseSnapMirrorPolicyAllOf) SetTransferScheduleUuid(v string)`

SetTransferScheduleUuid sets TransferScheduleUuid field to given value.

### HasTransferScheduleUuid

`func (o *StorageNetAppBaseSnapMirrorPolicyAllOf) HasTransferScheduleUuid() bool`

HasTransferScheduleUuid returns a boolean if a field has been set.

### GetType

`func (o *StorageNetAppBaseSnapMirrorPolicyAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StorageNetAppBaseSnapMirrorPolicyAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StorageNetAppBaseSnapMirrorPolicyAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *StorageNetAppBaseSnapMirrorPolicyAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUuid

`func (o *StorageNetAppBaseSnapMirrorPolicyAllOf) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *StorageNetAppBaseSnapMirrorPolicyAllOf) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *StorageNetAppBaseSnapMirrorPolicyAllOf) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *StorageNetAppBaseSnapMirrorPolicyAllOf) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



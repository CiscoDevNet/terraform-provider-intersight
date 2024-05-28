# StorageHitachiRemoteReplication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.HitachiRemoteReplication"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.HitachiRemoteReplication"]
**ConsistencyGroupId** | Pointer to **string** | Consistency group ID. If no consistency group consists, information is not input. | [optional] [readonly] 
**CopyPace** | Pointer to **string** | Copy speed. Number for the size of tracks to be copied. | [optional] [readonly] 
**DeltaStatus** | Pointer to **string** | Status of the 3DC multi-target configuration that uses delta resync. | [optional] [readonly] 
**FenceLevel** | Pointer to **string** | Fence level. Whether the P-VOL can be written to when the pair is split due to error. | [optional] [readonly] 
**MuNumber** | Pointer to **string** | MU (mirror unit) number of the volume. | [optional] [readonly] 
**Name** | Pointer to **string** | Object ID of the remote copy pair. | [optional] [readonly] 
**PathGroupId** | Pointer to **string** | Path group ID of the RCU. | [optional] [readonly] 
**PvolIoMode** | Pointer to **string** | I-O mode of the P-VOL. Information is input only in the case of global-active device. | [optional] [readonly] 
**PvolJournalId** | Pointer to **string** | Journal ID of the P-VOL. A value is input only in the case of Universal Replicator. | [optional] [readonly] 
**PvolLdevId** | Pointer to **int64** | LDEV number of primary volume. | [optional] [readonly] 
**PvolStorageSerial** | Pointer to **string** | Serial number of the storage system on the P-VOL. | [optional] [readonly] 
**QuorumDiskId** | Pointer to **string** | ID of the Quorum disk. A value is input only in the case of global-active device. | [optional] [readonly] 
**ReplicationType** | Pointer to **string** | Pair type of the remote copy pair. | [optional] [readonly] 
**Status** | Pointer to **string** | Status of the remote copy pair. | [optional] [readonly] 
**SvolIoMode** | Pointer to **string** | I-O mode of the S-VOL. Information is input only in the case of global-active device. | [optional] [readonly] 
**SvolJournalId** | Pointer to **string** | Journal ID of the S-VOL. A value is input only in the case of Universal Replicator. | [optional] [readonly] 
**SvolLdevId** | Pointer to **int64** | LDEV number of secondary volume. | [optional] [readonly] 
**SvolStorageSerial** | Pointer to **string** | Serial number of the storage system on the S-VOL. | [optional] [readonly] 
**Array** | Pointer to [**NullableStorageHitachiArrayRelationship**](StorageHitachiArrayRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewStorageHitachiRemoteReplication

`func NewStorageHitachiRemoteReplication(classId string, objectType string, ) *StorageHitachiRemoteReplication`

NewStorageHitachiRemoteReplication instantiates a new StorageHitachiRemoteReplication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageHitachiRemoteReplicationWithDefaults

`func NewStorageHitachiRemoteReplicationWithDefaults() *StorageHitachiRemoteReplication`

NewStorageHitachiRemoteReplicationWithDefaults instantiates a new StorageHitachiRemoteReplication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageHitachiRemoteReplication) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageHitachiRemoteReplication) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageHitachiRemoteReplication) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageHitachiRemoteReplication) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageHitachiRemoteReplication) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageHitachiRemoteReplication) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetConsistencyGroupId

`func (o *StorageHitachiRemoteReplication) GetConsistencyGroupId() string`

GetConsistencyGroupId returns the ConsistencyGroupId field if non-nil, zero value otherwise.

### GetConsistencyGroupIdOk

`func (o *StorageHitachiRemoteReplication) GetConsistencyGroupIdOk() (*string, bool)`

GetConsistencyGroupIdOk returns a tuple with the ConsistencyGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsistencyGroupId

`func (o *StorageHitachiRemoteReplication) SetConsistencyGroupId(v string)`

SetConsistencyGroupId sets ConsistencyGroupId field to given value.

### HasConsistencyGroupId

`func (o *StorageHitachiRemoteReplication) HasConsistencyGroupId() bool`

HasConsistencyGroupId returns a boolean if a field has been set.

### GetCopyPace

`func (o *StorageHitachiRemoteReplication) GetCopyPace() string`

GetCopyPace returns the CopyPace field if non-nil, zero value otherwise.

### GetCopyPaceOk

`func (o *StorageHitachiRemoteReplication) GetCopyPaceOk() (*string, bool)`

GetCopyPaceOk returns a tuple with the CopyPace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyPace

`func (o *StorageHitachiRemoteReplication) SetCopyPace(v string)`

SetCopyPace sets CopyPace field to given value.

### HasCopyPace

`func (o *StorageHitachiRemoteReplication) HasCopyPace() bool`

HasCopyPace returns a boolean if a field has been set.

### GetDeltaStatus

`func (o *StorageHitachiRemoteReplication) GetDeltaStatus() string`

GetDeltaStatus returns the DeltaStatus field if non-nil, zero value otherwise.

### GetDeltaStatusOk

`func (o *StorageHitachiRemoteReplication) GetDeltaStatusOk() (*string, bool)`

GetDeltaStatusOk returns a tuple with the DeltaStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeltaStatus

`func (o *StorageHitachiRemoteReplication) SetDeltaStatus(v string)`

SetDeltaStatus sets DeltaStatus field to given value.

### HasDeltaStatus

`func (o *StorageHitachiRemoteReplication) HasDeltaStatus() bool`

HasDeltaStatus returns a boolean if a field has been set.

### GetFenceLevel

`func (o *StorageHitachiRemoteReplication) GetFenceLevel() string`

GetFenceLevel returns the FenceLevel field if non-nil, zero value otherwise.

### GetFenceLevelOk

`func (o *StorageHitachiRemoteReplication) GetFenceLevelOk() (*string, bool)`

GetFenceLevelOk returns a tuple with the FenceLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFenceLevel

`func (o *StorageHitachiRemoteReplication) SetFenceLevel(v string)`

SetFenceLevel sets FenceLevel field to given value.

### HasFenceLevel

`func (o *StorageHitachiRemoteReplication) HasFenceLevel() bool`

HasFenceLevel returns a boolean if a field has been set.

### GetMuNumber

`func (o *StorageHitachiRemoteReplication) GetMuNumber() string`

GetMuNumber returns the MuNumber field if non-nil, zero value otherwise.

### GetMuNumberOk

`func (o *StorageHitachiRemoteReplication) GetMuNumberOk() (*string, bool)`

GetMuNumberOk returns a tuple with the MuNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuNumber

`func (o *StorageHitachiRemoteReplication) SetMuNumber(v string)`

SetMuNumber sets MuNumber field to given value.

### HasMuNumber

`func (o *StorageHitachiRemoteReplication) HasMuNumber() bool`

HasMuNumber returns a boolean if a field has been set.

### GetName

`func (o *StorageHitachiRemoteReplication) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageHitachiRemoteReplication) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageHitachiRemoteReplication) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StorageHitachiRemoteReplication) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPathGroupId

`func (o *StorageHitachiRemoteReplication) GetPathGroupId() string`

GetPathGroupId returns the PathGroupId field if non-nil, zero value otherwise.

### GetPathGroupIdOk

`func (o *StorageHitachiRemoteReplication) GetPathGroupIdOk() (*string, bool)`

GetPathGroupIdOk returns a tuple with the PathGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathGroupId

`func (o *StorageHitachiRemoteReplication) SetPathGroupId(v string)`

SetPathGroupId sets PathGroupId field to given value.

### HasPathGroupId

`func (o *StorageHitachiRemoteReplication) HasPathGroupId() bool`

HasPathGroupId returns a boolean if a field has been set.

### GetPvolIoMode

`func (o *StorageHitachiRemoteReplication) GetPvolIoMode() string`

GetPvolIoMode returns the PvolIoMode field if non-nil, zero value otherwise.

### GetPvolIoModeOk

`func (o *StorageHitachiRemoteReplication) GetPvolIoModeOk() (*string, bool)`

GetPvolIoModeOk returns a tuple with the PvolIoMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPvolIoMode

`func (o *StorageHitachiRemoteReplication) SetPvolIoMode(v string)`

SetPvolIoMode sets PvolIoMode field to given value.

### HasPvolIoMode

`func (o *StorageHitachiRemoteReplication) HasPvolIoMode() bool`

HasPvolIoMode returns a boolean if a field has been set.

### GetPvolJournalId

`func (o *StorageHitachiRemoteReplication) GetPvolJournalId() string`

GetPvolJournalId returns the PvolJournalId field if non-nil, zero value otherwise.

### GetPvolJournalIdOk

`func (o *StorageHitachiRemoteReplication) GetPvolJournalIdOk() (*string, bool)`

GetPvolJournalIdOk returns a tuple with the PvolJournalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPvolJournalId

`func (o *StorageHitachiRemoteReplication) SetPvolJournalId(v string)`

SetPvolJournalId sets PvolJournalId field to given value.

### HasPvolJournalId

`func (o *StorageHitachiRemoteReplication) HasPvolJournalId() bool`

HasPvolJournalId returns a boolean if a field has been set.

### GetPvolLdevId

`func (o *StorageHitachiRemoteReplication) GetPvolLdevId() int64`

GetPvolLdevId returns the PvolLdevId field if non-nil, zero value otherwise.

### GetPvolLdevIdOk

`func (o *StorageHitachiRemoteReplication) GetPvolLdevIdOk() (*int64, bool)`

GetPvolLdevIdOk returns a tuple with the PvolLdevId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPvolLdevId

`func (o *StorageHitachiRemoteReplication) SetPvolLdevId(v int64)`

SetPvolLdevId sets PvolLdevId field to given value.

### HasPvolLdevId

`func (o *StorageHitachiRemoteReplication) HasPvolLdevId() bool`

HasPvolLdevId returns a boolean if a field has been set.

### GetPvolStorageSerial

`func (o *StorageHitachiRemoteReplication) GetPvolStorageSerial() string`

GetPvolStorageSerial returns the PvolStorageSerial field if non-nil, zero value otherwise.

### GetPvolStorageSerialOk

`func (o *StorageHitachiRemoteReplication) GetPvolStorageSerialOk() (*string, bool)`

GetPvolStorageSerialOk returns a tuple with the PvolStorageSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPvolStorageSerial

`func (o *StorageHitachiRemoteReplication) SetPvolStorageSerial(v string)`

SetPvolStorageSerial sets PvolStorageSerial field to given value.

### HasPvolStorageSerial

`func (o *StorageHitachiRemoteReplication) HasPvolStorageSerial() bool`

HasPvolStorageSerial returns a boolean if a field has been set.

### GetQuorumDiskId

`func (o *StorageHitachiRemoteReplication) GetQuorumDiskId() string`

GetQuorumDiskId returns the QuorumDiskId field if non-nil, zero value otherwise.

### GetQuorumDiskIdOk

`func (o *StorageHitachiRemoteReplication) GetQuorumDiskIdOk() (*string, bool)`

GetQuorumDiskIdOk returns a tuple with the QuorumDiskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuorumDiskId

`func (o *StorageHitachiRemoteReplication) SetQuorumDiskId(v string)`

SetQuorumDiskId sets QuorumDiskId field to given value.

### HasQuorumDiskId

`func (o *StorageHitachiRemoteReplication) HasQuorumDiskId() bool`

HasQuorumDiskId returns a boolean if a field has been set.

### GetReplicationType

`func (o *StorageHitachiRemoteReplication) GetReplicationType() string`

GetReplicationType returns the ReplicationType field if non-nil, zero value otherwise.

### GetReplicationTypeOk

`func (o *StorageHitachiRemoteReplication) GetReplicationTypeOk() (*string, bool)`

GetReplicationTypeOk returns a tuple with the ReplicationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationType

`func (o *StorageHitachiRemoteReplication) SetReplicationType(v string)`

SetReplicationType sets ReplicationType field to given value.

### HasReplicationType

`func (o *StorageHitachiRemoteReplication) HasReplicationType() bool`

HasReplicationType returns a boolean if a field has been set.

### GetStatus

`func (o *StorageHitachiRemoteReplication) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StorageHitachiRemoteReplication) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StorageHitachiRemoteReplication) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *StorageHitachiRemoteReplication) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSvolIoMode

`func (o *StorageHitachiRemoteReplication) GetSvolIoMode() string`

GetSvolIoMode returns the SvolIoMode field if non-nil, zero value otherwise.

### GetSvolIoModeOk

`func (o *StorageHitachiRemoteReplication) GetSvolIoModeOk() (*string, bool)`

GetSvolIoModeOk returns a tuple with the SvolIoMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSvolIoMode

`func (o *StorageHitachiRemoteReplication) SetSvolIoMode(v string)`

SetSvolIoMode sets SvolIoMode field to given value.

### HasSvolIoMode

`func (o *StorageHitachiRemoteReplication) HasSvolIoMode() bool`

HasSvolIoMode returns a boolean if a field has been set.

### GetSvolJournalId

`func (o *StorageHitachiRemoteReplication) GetSvolJournalId() string`

GetSvolJournalId returns the SvolJournalId field if non-nil, zero value otherwise.

### GetSvolJournalIdOk

`func (o *StorageHitachiRemoteReplication) GetSvolJournalIdOk() (*string, bool)`

GetSvolJournalIdOk returns a tuple with the SvolJournalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSvolJournalId

`func (o *StorageHitachiRemoteReplication) SetSvolJournalId(v string)`

SetSvolJournalId sets SvolJournalId field to given value.

### HasSvolJournalId

`func (o *StorageHitachiRemoteReplication) HasSvolJournalId() bool`

HasSvolJournalId returns a boolean if a field has been set.

### GetSvolLdevId

`func (o *StorageHitachiRemoteReplication) GetSvolLdevId() int64`

GetSvolLdevId returns the SvolLdevId field if non-nil, zero value otherwise.

### GetSvolLdevIdOk

`func (o *StorageHitachiRemoteReplication) GetSvolLdevIdOk() (*int64, bool)`

GetSvolLdevIdOk returns a tuple with the SvolLdevId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSvolLdevId

`func (o *StorageHitachiRemoteReplication) SetSvolLdevId(v int64)`

SetSvolLdevId sets SvolLdevId field to given value.

### HasSvolLdevId

`func (o *StorageHitachiRemoteReplication) HasSvolLdevId() bool`

HasSvolLdevId returns a boolean if a field has been set.

### GetSvolStorageSerial

`func (o *StorageHitachiRemoteReplication) GetSvolStorageSerial() string`

GetSvolStorageSerial returns the SvolStorageSerial field if non-nil, zero value otherwise.

### GetSvolStorageSerialOk

`func (o *StorageHitachiRemoteReplication) GetSvolStorageSerialOk() (*string, bool)`

GetSvolStorageSerialOk returns a tuple with the SvolStorageSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSvolStorageSerial

`func (o *StorageHitachiRemoteReplication) SetSvolStorageSerial(v string)`

SetSvolStorageSerial sets SvolStorageSerial field to given value.

### HasSvolStorageSerial

`func (o *StorageHitachiRemoteReplication) HasSvolStorageSerial() bool`

HasSvolStorageSerial returns a boolean if a field has been set.

### GetArray

`func (o *StorageHitachiRemoteReplication) GetArray() StorageHitachiArrayRelationship`

GetArray returns the Array field if non-nil, zero value otherwise.

### GetArrayOk

`func (o *StorageHitachiRemoteReplication) GetArrayOk() (*StorageHitachiArrayRelationship, bool)`

GetArrayOk returns a tuple with the Array field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArray

`func (o *StorageHitachiRemoteReplication) SetArray(v StorageHitachiArrayRelationship)`

SetArray sets Array field to given value.

### HasArray

`func (o *StorageHitachiRemoteReplication) HasArray() bool`

HasArray returns a boolean if a field has been set.

### SetArrayNil

`func (o *StorageHitachiRemoteReplication) SetArrayNil(b bool)`

 SetArrayNil sets the value for Array to be an explicit nil

### UnsetArray
`func (o *StorageHitachiRemoteReplication) UnsetArray()`

UnsetArray ensures that no value is present for Array, not even an explicit nil
### GetRegisteredDevice

`func (o *StorageHitachiRemoteReplication) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StorageHitachiRemoteReplication) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StorageHitachiRemoteReplication) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StorageHitachiRemoteReplication) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *StorageHitachiRemoteReplication) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *StorageHitachiRemoteReplication) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



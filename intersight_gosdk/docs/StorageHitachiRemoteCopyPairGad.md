# StorageHitachiRemoteCopyPairGad

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.HitachiRemoteCopyPairGad"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.HitachiRemoteCopyPairGad"]
**CreatedLocalTime** | Pointer to **string** | Time at which the pair was created. The local time of the storage system is displayed. | [optional] [readonly] 
**IsSsws** | Pointer to **bool** | Indicates whether the state of the volume on the local storage system is SSWS. | [optional] [readonly] 
**MuNumber** | Pointer to **string** | MU (mirror unit) number of the volume. | [optional] [readonly] 
**Name** | Pointer to **string** | Object ID of the remote copy pair. | [optional] [readonly] 
**PvolLdevId** | Pointer to **int64** | LDEV number of primary volume. | [optional] [readonly] 
**PvolStorageSerial** | Pointer to **string** | Serial number of the storage system on the P-VOL. | [optional] [readonly] 
**QuorumDiskId** | Pointer to **string** | ID of the Quorum disk. A value is input only in the case of global-active device. | [optional] [readonly] 
**ReplicationType** | Pointer to **string** | Pair type of the remote copy pair. | [optional] [readonly] 
**Status** | Pointer to **string** | Status of the remote copy pair. | [optional] [readonly] 
**SuspendedMode** | Pointer to **string** | Block or Remote instructions when a pair is suspended. | [optional] [readonly] 
**SvolLdevId** | Pointer to **int64** | LDEV number of secondary volume. | [optional] [readonly] 
**SvolStorageSerial** | Pointer to **string** | Serial number of the storage system on the S-VOL. | [optional] [readonly] 
**Array** | Pointer to [**NullableStorageHitachiArrayRelationship**](StorageHitachiArrayRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewStorageHitachiRemoteCopyPairGad

`func NewStorageHitachiRemoteCopyPairGad(classId string, objectType string, ) *StorageHitachiRemoteCopyPairGad`

NewStorageHitachiRemoteCopyPairGad instantiates a new StorageHitachiRemoteCopyPairGad object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageHitachiRemoteCopyPairGadWithDefaults

`func NewStorageHitachiRemoteCopyPairGadWithDefaults() *StorageHitachiRemoteCopyPairGad`

NewStorageHitachiRemoteCopyPairGadWithDefaults instantiates a new StorageHitachiRemoteCopyPairGad object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageHitachiRemoteCopyPairGad) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageHitachiRemoteCopyPairGad) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageHitachiRemoteCopyPairGad) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageHitachiRemoteCopyPairGad) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageHitachiRemoteCopyPairGad) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageHitachiRemoteCopyPairGad) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCreatedLocalTime

`func (o *StorageHitachiRemoteCopyPairGad) GetCreatedLocalTime() string`

GetCreatedLocalTime returns the CreatedLocalTime field if non-nil, zero value otherwise.

### GetCreatedLocalTimeOk

`func (o *StorageHitachiRemoteCopyPairGad) GetCreatedLocalTimeOk() (*string, bool)`

GetCreatedLocalTimeOk returns a tuple with the CreatedLocalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedLocalTime

`func (o *StorageHitachiRemoteCopyPairGad) SetCreatedLocalTime(v string)`

SetCreatedLocalTime sets CreatedLocalTime field to given value.

### HasCreatedLocalTime

`func (o *StorageHitachiRemoteCopyPairGad) HasCreatedLocalTime() bool`

HasCreatedLocalTime returns a boolean if a field has been set.

### GetIsSsws

`func (o *StorageHitachiRemoteCopyPairGad) GetIsSsws() bool`

GetIsSsws returns the IsSsws field if non-nil, zero value otherwise.

### GetIsSswsOk

`func (o *StorageHitachiRemoteCopyPairGad) GetIsSswsOk() (*bool, bool)`

GetIsSswsOk returns a tuple with the IsSsws field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSsws

`func (o *StorageHitachiRemoteCopyPairGad) SetIsSsws(v bool)`

SetIsSsws sets IsSsws field to given value.

### HasIsSsws

`func (o *StorageHitachiRemoteCopyPairGad) HasIsSsws() bool`

HasIsSsws returns a boolean if a field has been set.

### GetMuNumber

`func (o *StorageHitachiRemoteCopyPairGad) GetMuNumber() string`

GetMuNumber returns the MuNumber field if non-nil, zero value otherwise.

### GetMuNumberOk

`func (o *StorageHitachiRemoteCopyPairGad) GetMuNumberOk() (*string, bool)`

GetMuNumberOk returns a tuple with the MuNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuNumber

`func (o *StorageHitachiRemoteCopyPairGad) SetMuNumber(v string)`

SetMuNumber sets MuNumber field to given value.

### HasMuNumber

`func (o *StorageHitachiRemoteCopyPairGad) HasMuNumber() bool`

HasMuNumber returns a boolean if a field has been set.

### GetName

`func (o *StorageHitachiRemoteCopyPairGad) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageHitachiRemoteCopyPairGad) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageHitachiRemoteCopyPairGad) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StorageHitachiRemoteCopyPairGad) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPvolLdevId

`func (o *StorageHitachiRemoteCopyPairGad) GetPvolLdevId() int64`

GetPvolLdevId returns the PvolLdevId field if non-nil, zero value otherwise.

### GetPvolLdevIdOk

`func (o *StorageHitachiRemoteCopyPairGad) GetPvolLdevIdOk() (*int64, bool)`

GetPvolLdevIdOk returns a tuple with the PvolLdevId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPvolLdevId

`func (o *StorageHitachiRemoteCopyPairGad) SetPvolLdevId(v int64)`

SetPvolLdevId sets PvolLdevId field to given value.

### HasPvolLdevId

`func (o *StorageHitachiRemoteCopyPairGad) HasPvolLdevId() bool`

HasPvolLdevId returns a boolean if a field has been set.

### GetPvolStorageSerial

`func (o *StorageHitachiRemoteCopyPairGad) GetPvolStorageSerial() string`

GetPvolStorageSerial returns the PvolStorageSerial field if non-nil, zero value otherwise.

### GetPvolStorageSerialOk

`func (o *StorageHitachiRemoteCopyPairGad) GetPvolStorageSerialOk() (*string, bool)`

GetPvolStorageSerialOk returns a tuple with the PvolStorageSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPvolStorageSerial

`func (o *StorageHitachiRemoteCopyPairGad) SetPvolStorageSerial(v string)`

SetPvolStorageSerial sets PvolStorageSerial field to given value.

### HasPvolStorageSerial

`func (o *StorageHitachiRemoteCopyPairGad) HasPvolStorageSerial() bool`

HasPvolStorageSerial returns a boolean if a field has been set.

### GetQuorumDiskId

`func (o *StorageHitachiRemoteCopyPairGad) GetQuorumDiskId() string`

GetQuorumDiskId returns the QuorumDiskId field if non-nil, zero value otherwise.

### GetQuorumDiskIdOk

`func (o *StorageHitachiRemoteCopyPairGad) GetQuorumDiskIdOk() (*string, bool)`

GetQuorumDiskIdOk returns a tuple with the QuorumDiskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuorumDiskId

`func (o *StorageHitachiRemoteCopyPairGad) SetQuorumDiskId(v string)`

SetQuorumDiskId sets QuorumDiskId field to given value.

### HasQuorumDiskId

`func (o *StorageHitachiRemoteCopyPairGad) HasQuorumDiskId() bool`

HasQuorumDiskId returns a boolean if a field has been set.

### GetReplicationType

`func (o *StorageHitachiRemoteCopyPairGad) GetReplicationType() string`

GetReplicationType returns the ReplicationType field if non-nil, zero value otherwise.

### GetReplicationTypeOk

`func (o *StorageHitachiRemoteCopyPairGad) GetReplicationTypeOk() (*string, bool)`

GetReplicationTypeOk returns a tuple with the ReplicationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationType

`func (o *StorageHitachiRemoteCopyPairGad) SetReplicationType(v string)`

SetReplicationType sets ReplicationType field to given value.

### HasReplicationType

`func (o *StorageHitachiRemoteCopyPairGad) HasReplicationType() bool`

HasReplicationType returns a boolean if a field has been set.

### GetStatus

`func (o *StorageHitachiRemoteCopyPairGad) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StorageHitachiRemoteCopyPairGad) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StorageHitachiRemoteCopyPairGad) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *StorageHitachiRemoteCopyPairGad) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSuspendedMode

`func (o *StorageHitachiRemoteCopyPairGad) GetSuspendedMode() string`

GetSuspendedMode returns the SuspendedMode field if non-nil, zero value otherwise.

### GetSuspendedModeOk

`func (o *StorageHitachiRemoteCopyPairGad) GetSuspendedModeOk() (*string, bool)`

GetSuspendedModeOk returns a tuple with the SuspendedMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspendedMode

`func (o *StorageHitachiRemoteCopyPairGad) SetSuspendedMode(v string)`

SetSuspendedMode sets SuspendedMode field to given value.

### HasSuspendedMode

`func (o *StorageHitachiRemoteCopyPairGad) HasSuspendedMode() bool`

HasSuspendedMode returns a boolean if a field has been set.

### GetSvolLdevId

`func (o *StorageHitachiRemoteCopyPairGad) GetSvolLdevId() int64`

GetSvolLdevId returns the SvolLdevId field if non-nil, zero value otherwise.

### GetSvolLdevIdOk

`func (o *StorageHitachiRemoteCopyPairGad) GetSvolLdevIdOk() (*int64, bool)`

GetSvolLdevIdOk returns a tuple with the SvolLdevId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSvolLdevId

`func (o *StorageHitachiRemoteCopyPairGad) SetSvolLdevId(v int64)`

SetSvolLdevId sets SvolLdevId field to given value.

### HasSvolLdevId

`func (o *StorageHitachiRemoteCopyPairGad) HasSvolLdevId() bool`

HasSvolLdevId returns a boolean if a field has been set.

### GetSvolStorageSerial

`func (o *StorageHitachiRemoteCopyPairGad) GetSvolStorageSerial() string`

GetSvolStorageSerial returns the SvolStorageSerial field if non-nil, zero value otherwise.

### GetSvolStorageSerialOk

`func (o *StorageHitachiRemoteCopyPairGad) GetSvolStorageSerialOk() (*string, bool)`

GetSvolStorageSerialOk returns a tuple with the SvolStorageSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSvolStorageSerial

`func (o *StorageHitachiRemoteCopyPairGad) SetSvolStorageSerial(v string)`

SetSvolStorageSerial sets SvolStorageSerial field to given value.

### HasSvolStorageSerial

`func (o *StorageHitachiRemoteCopyPairGad) HasSvolStorageSerial() bool`

HasSvolStorageSerial returns a boolean if a field has been set.

### GetArray

`func (o *StorageHitachiRemoteCopyPairGad) GetArray() StorageHitachiArrayRelationship`

GetArray returns the Array field if non-nil, zero value otherwise.

### GetArrayOk

`func (o *StorageHitachiRemoteCopyPairGad) GetArrayOk() (*StorageHitachiArrayRelationship, bool)`

GetArrayOk returns a tuple with the Array field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArray

`func (o *StorageHitachiRemoteCopyPairGad) SetArray(v StorageHitachiArrayRelationship)`

SetArray sets Array field to given value.

### HasArray

`func (o *StorageHitachiRemoteCopyPairGad) HasArray() bool`

HasArray returns a boolean if a field has been set.

### SetArrayNil

`func (o *StorageHitachiRemoteCopyPairGad) SetArrayNil(b bool)`

 SetArrayNil sets the value for Array to be an explicit nil

### UnsetArray
`func (o *StorageHitachiRemoteCopyPairGad) UnsetArray()`

UnsetArray ensures that no value is present for Array, not even an explicit nil
### GetRegisteredDevice

`func (o *StorageHitachiRemoteCopyPairGad) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StorageHitachiRemoteCopyPairGad) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StorageHitachiRemoteCopyPairGad) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StorageHitachiRemoteCopyPairGad) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *StorageHitachiRemoteCopyPairGad) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *StorageHitachiRemoteCopyPairGad) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



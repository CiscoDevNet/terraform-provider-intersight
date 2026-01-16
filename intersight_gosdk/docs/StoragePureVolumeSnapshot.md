# StoragePureVolumeSnapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.PureVolumeSnapshot"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.PureVolumeSnapshot"]
**Pod** | Pointer to **string** | A pod representing a collection of protection groups and volumes is created on one array and stretched to another array, resulting in fully synchronized writes between the two arrays. | [optional] [readonly] 
**RealmName** | Pointer to **string** | A realm is the core multi-tenancy component on a Pure Flash Array, providing a self-contained, virtual storage environment with dedicated policies and quotas for secure data isolation and predictable performance. | [optional] [readonly] 
**Serial** | Pointer to **string** | Unique serial number of the snapshot allocated by the storage array. | [optional] [readonly] 
**SnapshotSize** | Pointer to **int64** | The size of the snapshot created. | [optional] [readonly] 
**TotalProvisioned** | Pointer to **int64** | The overall size of the snapshot allocated by the storage array. | [optional] [readonly] 
**UsedProvisioned** | Pointer to **int64** | The used size of the snapshot allocated by the storage array. | [optional] [readonly] 
**VolumeGroup** | Pointer to **string** | Volume groups organize volumes into logical groupings. If virtual volumes are configured, each volume group on the FlashArray array represents its associated virtual machine, and inside each of those volumes groups are the FlashArray volumes that are assigned to the virtual machine. | [optional] [readonly] 
**Array** | Pointer to [**NullableStoragePureArrayRelationship**](StoragePureArrayRelationship.md) |  | [optional] 
**ProtectionGroupSnapshot** | Pointer to [**NullableStoragePureProtectionGroupSnapshotRelationship**](StoragePureProtectionGroupSnapshotRelationship.md) |  | [optional] 
**Realm** | Pointer to [**NullableStoragePureRealmRelationship**](StoragePureRealmRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 
**Volume** | Pointer to [**NullableStoragePureVolumeRelationship**](StoragePureVolumeRelationship.md) |  | [optional] 

## Methods

### NewStoragePureVolumeSnapshot

`func NewStoragePureVolumeSnapshot(classId string, objectType string, ) *StoragePureVolumeSnapshot`

NewStoragePureVolumeSnapshot instantiates a new StoragePureVolumeSnapshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoragePureVolumeSnapshotWithDefaults

`func NewStoragePureVolumeSnapshotWithDefaults() *StoragePureVolumeSnapshot`

NewStoragePureVolumeSnapshotWithDefaults instantiates a new StoragePureVolumeSnapshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StoragePureVolumeSnapshot) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StoragePureVolumeSnapshot) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StoragePureVolumeSnapshot) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StoragePureVolumeSnapshot) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StoragePureVolumeSnapshot) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StoragePureVolumeSnapshot) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetPod

`func (o *StoragePureVolumeSnapshot) GetPod() string`

GetPod returns the Pod field if non-nil, zero value otherwise.

### GetPodOk

`func (o *StoragePureVolumeSnapshot) GetPodOk() (*string, bool)`

GetPodOk returns a tuple with the Pod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPod

`func (o *StoragePureVolumeSnapshot) SetPod(v string)`

SetPod sets Pod field to given value.

### HasPod

`func (o *StoragePureVolumeSnapshot) HasPod() bool`

HasPod returns a boolean if a field has been set.

### GetRealmName

`func (o *StoragePureVolumeSnapshot) GetRealmName() string`

GetRealmName returns the RealmName field if non-nil, zero value otherwise.

### GetRealmNameOk

`func (o *StoragePureVolumeSnapshot) GetRealmNameOk() (*string, bool)`

GetRealmNameOk returns a tuple with the RealmName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmName

`func (o *StoragePureVolumeSnapshot) SetRealmName(v string)`

SetRealmName sets RealmName field to given value.

### HasRealmName

`func (o *StoragePureVolumeSnapshot) HasRealmName() bool`

HasRealmName returns a boolean if a field has been set.

### GetSerial

`func (o *StoragePureVolumeSnapshot) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *StoragePureVolumeSnapshot) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *StoragePureVolumeSnapshot) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *StoragePureVolumeSnapshot) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetSnapshotSize

`func (o *StoragePureVolumeSnapshot) GetSnapshotSize() int64`

GetSnapshotSize returns the SnapshotSize field if non-nil, zero value otherwise.

### GetSnapshotSizeOk

`func (o *StoragePureVolumeSnapshot) GetSnapshotSizeOk() (*int64, bool)`

GetSnapshotSizeOk returns a tuple with the SnapshotSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotSize

`func (o *StoragePureVolumeSnapshot) SetSnapshotSize(v int64)`

SetSnapshotSize sets SnapshotSize field to given value.

### HasSnapshotSize

`func (o *StoragePureVolumeSnapshot) HasSnapshotSize() bool`

HasSnapshotSize returns a boolean if a field has been set.

### GetTotalProvisioned

`func (o *StoragePureVolumeSnapshot) GetTotalProvisioned() int64`

GetTotalProvisioned returns the TotalProvisioned field if non-nil, zero value otherwise.

### GetTotalProvisionedOk

`func (o *StoragePureVolumeSnapshot) GetTotalProvisionedOk() (*int64, bool)`

GetTotalProvisionedOk returns a tuple with the TotalProvisioned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalProvisioned

`func (o *StoragePureVolumeSnapshot) SetTotalProvisioned(v int64)`

SetTotalProvisioned sets TotalProvisioned field to given value.

### HasTotalProvisioned

`func (o *StoragePureVolumeSnapshot) HasTotalProvisioned() bool`

HasTotalProvisioned returns a boolean if a field has been set.

### GetUsedProvisioned

`func (o *StoragePureVolumeSnapshot) GetUsedProvisioned() int64`

GetUsedProvisioned returns the UsedProvisioned field if non-nil, zero value otherwise.

### GetUsedProvisionedOk

`func (o *StoragePureVolumeSnapshot) GetUsedProvisionedOk() (*int64, bool)`

GetUsedProvisionedOk returns a tuple with the UsedProvisioned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedProvisioned

`func (o *StoragePureVolumeSnapshot) SetUsedProvisioned(v int64)`

SetUsedProvisioned sets UsedProvisioned field to given value.

### HasUsedProvisioned

`func (o *StoragePureVolumeSnapshot) HasUsedProvisioned() bool`

HasUsedProvisioned returns a boolean if a field has been set.

### GetVolumeGroup

`func (o *StoragePureVolumeSnapshot) GetVolumeGroup() string`

GetVolumeGroup returns the VolumeGroup field if non-nil, zero value otherwise.

### GetVolumeGroupOk

`func (o *StoragePureVolumeSnapshot) GetVolumeGroupOk() (*string, bool)`

GetVolumeGroupOk returns a tuple with the VolumeGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeGroup

`func (o *StoragePureVolumeSnapshot) SetVolumeGroup(v string)`

SetVolumeGroup sets VolumeGroup field to given value.

### HasVolumeGroup

`func (o *StoragePureVolumeSnapshot) HasVolumeGroup() bool`

HasVolumeGroup returns a boolean if a field has been set.

### GetArray

`func (o *StoragePureVolumeSnapshot) GetArray() StoragePureArrayRelationship`

GetArray returns the Array field if non-nil, zero value otherwise.

### GetArrayOk

`func (o *StoragePureVolumeSnapshot) GetArrayOk() (*StoragePureArrayRelationship, bool)`

GetArrayOk returns a tuple with the Array field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArray

`func (o *StoragePureVolumeSnapshot) SetArray(v StoragePureArrayRelationship)`

SetArray sets Array field to given value.

### HasArray

`func (o *StoragePureVolumeSnapshot) HasArray() bool`

HasArray returns a boolean if a field has been set.

### SetArrayNil

`func (o *StoragePureVolumeSnapshot) SetArrayNil(b bool)`

 SetArrayNil sets the value for Array to be an explicit nil

### UnsetArray
`func (o *StoragePureVolumeSnapshot) UnsetArray()`

UnsetArray ensures that no value is present for Array, not even an explicit nil
### GetProtectionGroupSnapshot

`func (o *StoragePureVolumeSnapshot) GetProtectionGroupSnapshot() StoragePureProtectionGroupSnapshotRelationship`

GetProtectionGroupSnapshot returns the ProtectionGroupSnapshot field if non-nil, zero value otherwise.

### GetProtectionGroupSnapshotOk

`func (o *StoragePureVolumeSnapshot) GetProtectionGroupSnapshotOk() (*StoragePureProtectionGroupSnapshotRelationship, bool)`

GetProtectionGroupSnapshotOk returns a tuple with the ProtectionGroupSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroupSnapshot

`func (o *StoragePureVolumeSnapshot) SetProtectionGroupSnapshot(v StoragePureProtectionGroupSnapshotRelationship)`

SetProtectionGroupSnapshot sets ProtectionGroupSnapshot field to given value.

### HasProtectionGroupSnapshot

`func (o *StoragePureVolumeSnapshot) HasProtectionGroupSnapshot() bool`

HasProtectionGroupSnapshot returns a boolean if a field has been set.

### SetProtectionGroupSnapshotNil

`func (o *StoragePureVolumeSnapshot) SetProtectionGroupSnapshotNil(b bool)`

 SetProtectionGroupSnapshotNil sets the value for ProtectionGroupSnapshot to be an explicit nil

### UnsetProtectionGroupSnapshot
`func (o *StoragePureVolumeSnapshot) UnsetProtectionGroupSnapshot()`

UnsetProtectionGroupSnapshot ensures that no value is present for ProtectionGroupSnapshot, not even an explicit nil
### GetRealm

`func (o *StoragePureVolumeSnapshot) GetRealm() StoragePureRealmRelationship`

GetRealm returns the Realm field if non-nil, zero value otherwise.

### GetRealmOk

`func (o *StoragePureVolumeSnapshot) GetRealmOk() (*StoragePureRealmRelationship, bool)`

GetRealmOk returns a tuple with the Realm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealm

`func (o *StoragePureVolumeSnapshot) SetRealm(v StoragePureRealmRelationship)`

SetRealm sets Realm field to given value.

### HasRealm

`func (o *StoragePureVolumeSnapshot) HasRealm() bool`

HasRealm returns a boolean if a field has been set.

### SetRealmNil

`func (o *StoragePureVolumeSnapshot) SetRealmNil(b bool)`

 SetRealmNil sets the value for Realm to be an explicit nil

### UnsetRealm
`func (o *StoragePureVolumeSnapshot) UnsetRealm()`

UnsetRealm ensures that no value is present for Realm, not even an explicit nil
### GetRegisteredDevice

`func (o *StoragePureVolumeSnapshot) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StoragePureVolumeSnapshot) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StoragePureVolumeSnapshot) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StoragePureVolumeSnapshot) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *StoragePureVolumeSnapshot) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *StoragePureVolumeSnapshot) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
### GetVolume

`func (o *StoragePureVolumeSnapshot) GetVolume() StoragePureVolumeRelationship`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *StoragePureVolumeSnapshot) GetVolumeOk() (*StoragePureVolumeRelationship, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *StoragePureVolumeSnapshot) SetVolume(v StoragePureVolumeRelationship)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *StoragePureVolumeSnapshot) HasVolume() bool`

HasVolume returns a boolean if a field has been set.

### SetVolumeNil

`func (o *StoragePureVolumeSnapshot) SetVolumeNil(b bool)`

 SetVolumeNil sets the value for Volume to be an explicit nil

### UnsetVolume
`func (o *StoragePureVolumeSnapshot) UnsetVolume()`

UnsetVolume ensures that no value is present for Volume, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



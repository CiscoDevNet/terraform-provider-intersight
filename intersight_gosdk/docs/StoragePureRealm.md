# StoragePureRealm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.PureRealm"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.PureRealm"]
**BandwidthLimit** | Pointer to **int64** | Maximum allowed data throughput for this object. | [optional] [readonly] 
**Created** | Pointer to **time.Time** | Creation time of the realm. | [optional] [readonly] 
**IopsLimit** | Pointer to **int64** | Maximum allowed IO operations per second for this object. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the realm as configured on the Pure Storage array. | [optional] [readonly] 
**QuotaLimit** | Pointer to **int64** | Maximum logical space the object is allowed to consume. | [optional] [readonly] 
**StorageUtilization** | Pointer to [**NullableStoragePureRealmUtilization**](StoragePureRealmUtilization.md) |  | [optional] 
**Array** | Pointer to [**NullableStoragePureArrayRelationship**](StoragePureArrayRelationship.md) |  | [optional] 
**Hosts** | Pointer to [**[]StoragePureHostRelationship**](StoragePureHostRelationship.md) | An array of relationships to storagePureHost resources. | [optional] [readonly] 
**Pods** | Pointer to [**[]StoragePurePodRelationship**](StoragePurePodRelationship.md) | An array of relationships to storagePurePod resources. | [optional] [readonly] 
**ProtectionGroupSnapshots** | Pointer to [**[]StoragePureProtectionGroupSnapshotRelationship**](StoragePureProtectionGroupSnapshotRelationship.md) | An array of relationships to storagePureProtectionGroupSnapshot resources. | [optional] [readonly] 
**ProtectionGroups** | Pointer to [**[]StoragePureProtectionGroupRelationship**](StoragePureProtectionGroupRelationship.md) | An array of relationships to storagePureProtectionGroup resources. | [optional] [readonly] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 
**VolumeSnapshots** | Pointer to [**[]StoragePureVolumeSnapshotRelationship**](StoragePureVolumeSnapshotRelationship.md) | An array of relationships to storagePureVolumeSnapshot resources. | [optional] [readonly] 
**Volumes** | Pointer to [**[]StoragePureVolumeRelationship**](StoragePureVolumeRelationship.md) | An array of relationships to storagePureVolume resources. | [optional] [readonly] 

## Methods

### NewStoragePureRealm

`func NewStoragePureRealm(classId string, objectType string, ) *StoragePureRealm`

NewStoragePureRealm instantiates a new StoragePureRealm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoragePureRealmWithDefaults

`func NewStoragePureRealmWithDefaults() *StoragePureRealm`

NewStoragePureRealmWithDefaults instantiates a new StoragePureRealm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StoragePureRealm) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StoragePureRealm) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StoragePureRealm) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StoragePureRealm) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StoragePureRealm) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StoragePureRealm) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBandwidthLimit

`func (o *StoragePureRealm) GetBandwidthLimit() int64`

GetBandwidthLimit returns the BandwidthLimit field if non-nil, zero value otherwise.

### GetBandwidthLimitOk

`func (o *StoragePureRealm) GetBandwidthLimitOk() (*int64, bool)`

GetBandwidthLimitOk returns a tuple with the BandwidthLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidthLimit

`func (o *StoragePureRealm) SetBandwidthLimit(v int64)`

SetBandwidthLimit sets BandwidthLimit field to given value.

### HasBandwidthLimit

`func (o *StoragePureRealm) HasBandwidthLimit() bool`

HasBandwidthLimit returns a boolean if a field has been set.

### GetCreated

`func (o *StoragePureRealm) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *StoragePureRealm) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *StoragePureRealm) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *StoragePureRealm) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetIopsLimit

`func (o *StoragePureRealm) GetIopsLimit() int64`

GetIopsLimit returns the IopsLimit field if non-nil, zero value otherwise.

### GetIopsLimitOk

`func (o *StoragePureRealm) GetIopsLimitOk() (*int64, bool)`

GetIopsLimitOk returns a tuple with the IopsLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIopsLimit

`func (o *StoragePureRealm) SetIopsLimit(v int64)`

SetIopsLimit sets IopsLimit field to given value.

### HasIopsLimit

`func (o *StoragePureRealm) HasIopsLimit() bool`

HasIopsLimit returns a boolean if a field has been set.

### GetName

`func (o *StoragePureRealm) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StoragePureRealm) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StoragePureRealm) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StoragePureRealm) HasName() bool`

HasName returns a boolean if a field has been set.

### GetQuotaLimit

`func (o *StoragePureRealm) GetQuotaLimit() int64`

GetQuotaLimit returns the QuotaLimit field if non-nil, zero value otherwise.

### GetQuotaLimitOk

`func (o *StoragePureRealm) GetQuotaLimitOk() (*int64, bool)`

GetQuotaLimitOk returns a tuple with the QuotaLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaLimit

`func (o *StoragePureRealm) SetQuotaLimit(v int64)`

SetQuotaLimit sets QuotaLimit field to given value.

### HasQuotaLimit

`func (o *StoragePureRealm) HasQuotaLimit() bool`

HasQuotaLimit returns a boolean if a field has been set.

### GetStorageUtilization

`func (o *StoragePureRealm) GetStorageUtilization() StoragePureRealmUtilization`

GetStorageUtilization returns the StorageUtilization field if non-nil, zero value otherwise.

### GetStorageUtilizationOk

`func (o *StoragePureRealm) GetStorageUtilizationOk() (*StoragePureRealmUtilization, bool)`

GetStorageUtilizationOk returns a tuple with the StorageUtilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageUtilization

`func (o *StoragePureRealm) SetStorageUtilization(v StoragePureRealmUtilization)`

SetStorageUtilization sets StorageUtilization field to given value.

### HasStorageUtilization

`func (o *StoragePureRealm) HasStorageUtilization() bool`

HasStorageUtilization returns a boolean if a field has been set.

### SetStorageUtilizationNil

`func (o *StoragePureRealm) SetStorageUtilizationNil(b bool)`

 SetStorageUtilizationNil sets the value for StorageUtilization to be an explicit nil

### UnsetStorageUtilization
`func (o *StoragePureRealm) UnsetStorageUtilization()`

UnsetStorageUtilization ensures that no value is present for StorageUtilization, not even an explicit nil
### GetArray

`func (o *StoragePureRealm) GetArray() StoragePureArrayRelationship`

GetArray returns the Array field if non-nil, zero value otherwise.

### GetArrayOk

`func (o *StoragePureRealm) GetArrayOk() (*StoragePureArrayRelationship, bool)`

GetArrayOk returns a tuple with the Array field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArray

`func (o *StoragePureRealm) SetArray(v StoragePureArrayRelationship)`

SetArray sets Array field to given value.

### HasArray

`func (o *StoragePureRealm) HasArray() bool`

HasArray returns a boolean if a field has been set.

### SetArrayNil

`func (o *StoragePureRealm) SetArrayNil(b bool)`

 SetArrayNil sets the value for Array to be an explicit nil

### UnsetArray
`func (o *StoragePureRealm) UnsetArray()`

UnsetArray ensures that no value is present for Array, not even an explicit nil
### GetHosts

`func (o *StoragePureRealm) GetHosts() []StoragePureHostRelationship`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *StoragePureRealm) GetHostsOk() (*[]StoragePureHostRelationship, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *StoragePureRealm) SetHosts(v []StoragePureHostRelationship)`

SetHosts sets Hosts field to given value.

### HasHosts

`func (o *StoragePureRealm) HasHosts() bool`

HasHosts returns a boolean if a field has been set.

### SetHostsNil

`func (o *StoragePureRealm) SetHostsNil(b bool)`

 SetHostsNil sets the value for Hosts to be an explicit nil

### UnsetHosts
`func (o *StoragePureRealm) UnsetHosts()`

UnsetHosts ensures that no value is present for Hosts, not even an explicit nil
### GetPods

`func (o *StoragePureRealm) GetPods() []StoragePurePodRelationship`

GetPods returns the Pods field if non-nil, zero value otherwise.

### GetPodsOk

`func (o *StoragePureRealm) GetPodsOk() (*[]StoragePurePodRelationship, bool)`

GetPodsOk returns a tuple with the Pods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPods

`func (o *StoragePureRealm) SetPods(v []StoragePurePodRelationship)`

SetPods sets Pods field to given value.

### HasPods

`func (o *StoragePureRealm) HasPods() bool`

HasPods returns a boolean if a field has been set.

### SetPodsNil

`func (o *StoragePureRealm) SetPodsNil(b bool)`

 SetPodsNil sets the value for Pods to be an explicit nil

### UnsetPods
`func (o *StoragePureRealm) UnsetPods()`

UnsetPods ensures that no value is present for Pods, not even an explicit nil
### GetProtectionGroupSnapshots

`func (o *StoragePureRealm) GetProtectionGroupSnapshots() []StoragePureProtectionGroupSnapshotRelationship`

GetProtectionGroupSnapshots returns the ProtectionGroupSnapshots field if non-nil, zero value otherwise.

### GetProtectionGroupSnapshotsOk

`func (o *StoragePureRealm) GetProtectionGroupSnapshotsOk() (*[]StoragePureProtectionGroupSnapshotRelationship, bool)`

GetProtectionGroupSnapshotsOk returns a tuple with the ProtectionGroupSnapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroupSnapshots

`func (o *StoragePureRealm) SetProtectionGroupSnapshots(v []StoragePureProtectionGroupSnapshotRelationship)`

SetProtectionGroupSnapshots sets ProtectionGroupSnapshots field to given value.

### HasProtectionGroupSnapshots

`func (o *StoragePureRealm) HasProtectionGroupSnapshots() bool`

HasProtectionGroupSnapshots returns a boolean if a field has been set.

### SetProtectionGroupSnapshotsNil

`func (o *StoragePureRealm) SetProtectionGroupSnapshotsNil(b bool)`

 SetProtectionGroupSnapshotsNil sets the value for ProtectionGroupSnapshots to be an explicit nil

### UnsetProtectionGroupSnapshots
`func (o *StoragePureRealm) UnsetProtectionGroupSnapshots()`

UnsetProtectionGroupSnapshots ensures that no value is present for ProtectionGroupSnapshots, not even an explicit nil
### GetProtectionGroups

`func (o *StoragePureRealm) GetProtectionGroups() []StoragePureProtectionGroupRelationship`

GetProtectionGroups returns the ProtectionGroups field if non-nil, zero value otherwise.

### GetProtectionGroupsOk

`func (o *StoragePureRealm) GetProtectionGroupsOk() (*[]StoragePureProtectionGroupRelationship, bool)`

GetProtectionGroupsOk returns a tuple with the ProtectionGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroups

`func (o *StoragePureRealm) SetProtectionGroups(v []StoragePureProtectionGroupRelationship)`

SetProtectionGroups sets ProtectionGroups field to given value.

### HasProtectionGroups

`func (o *StoragePureRealm) HasProtectionGroups() bool`

HasProtectionGroups returns a boolean if a field has been set.

### SetProtectionGroupsNil

`func (o *StoragePureRealm) SetProtectionGroupsNil(b bool)`

 SetProtectionGroupsNil sets the value for ProtectionGroups to be an explicit nil

### UnsetProtectionGroups
`func (o *StoragePureRealm) UnsetProtectionGroups()`

UnsetProtectionGroups ensures that no value is present for ProtectionGroups, not even an explicit nil
### GetRegisteredDevice

`func (o *StoragePureRealm) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StoragePureRealm) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StoragePureRealm) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StoragePureRealm) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *StoragePureRealm) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *StoragePureRealm) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
### GetVolumeSnapshots

`func (o *StoragePureRealm) GetVolumeSnapshots() []StoragePureVolumeSnapshotRelationship`

GetVolumeSnapshots returns the VolumeSnapshots field if non-nil, zero value otherwise.

### GetVolumeSnapshotsOk

`func (o *StoragePureRealm) GetVolumeSnapshotsOk() (*[]StoragePureVolumeSnapshotRelationship, bool)`

GetVolumeSnapshotsOk returns a tuple with the VolumeSnapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeSnapshots

`func (o *StoragePureRealm) SetVolumeSnapshots(v []StoragePureVolumeSnapshotRelationship)`

SetVolumeSnapshots sets VolumeSnapshots field to given value.

### HasVolumeSnapshots

`func (o *StoragePureRealm) HasVolumeSnapshots() bool`

HasVolumeSnapshots returns a boolean if a field has been set.

### SetVolumeSnapshotsNil

`func (o *StoragePureRealm) SetVolumeSnapshotsNil(b bool)`

 SetVolumeSnapshotsNil sets the value for VolumeSnapshots to be an explicit nil

### UnsetVolumeSnapshots
`func (o *StoragePureRealm) UnsetVolumeSnapshots()`

UnsetVolumeSnapshots ensures that no value is present for VolumeSnapshots, not even an explicit nil
### GetVolumes

`func (o *StoragePureRealm) GetVolumes() []StoragePureVolumeRelationship`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *StoragePureRealm) GetVolumesOk() (*[]StoragePureVolumeRelationship, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *StoragePureRealm) SetVolumes(v []StoragePureVolumeRelationship)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *StoragePureRealm) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### SetVolumesNil

`func (o *StoragePureRealm) SetVolumesNil(b bool)`

 SetVolumesNil sets the value for Volumes to be an explicit nil

### UnsetVolumes
`func (o *StoragePureRealm) UnsetVolumes()`

UnsetVolumes ensures that no value is present for Volumes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



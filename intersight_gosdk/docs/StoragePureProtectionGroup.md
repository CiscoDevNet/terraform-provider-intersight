# StoragePureProtectionGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.PureProtectionGroup"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.PureProtectionGroup"]
**PodName** | Pointer to **string** | A pod representing a collection of protection groups and volumes is created on one array and stretched to another array, resulting in fully synchronized writes between the two arrays. | [optional] [readonly] 
**RealmName** | Pointer to **string** | A realm is the core multi-tenancy component on a Pure Flash Array, providing a self-contained, virtual storage environment with dedicated policies and quotas for secure data isolation and predictable performance. | [optional] [readonly] 
**Size** | Pointer to **int64** | Overall size of all snapshots in the protection group, represented in bytes. | [optional] [readonly] 
**Source** | Pointer to **string** | Name of PureStorage array name on which the protection group is created. | [optional] [readonly] 
**Targets** | Pointer to **[]string** |  | [optional] 
**Array** | Pointer to [**NullableStoragePureArrayRelationship**](StoragePureArrayRelationship.md) |  | [optional] 
**HostGroups** | Pointer to [**[]StoragePureHostGroupRelationship**](StoragePureHostGroupRelationship.md) | An array of relationships to storagePureHostGroup resources. | [optional] [readonly] 
**Hosts** | Pointer to [**[]StoragePureHostRelationship**](StoragePureHostRelationship.md) | An array of relationships to storagePureHost resources. | [optional] [readonly] 
**Pod** | Pointer to [**NullableStoragePurePodRelationship**](StoragePurePodRelationship.md) |  | [optional] 
**Realm** | Pointer to [**NullableStoragePureRealmRelationship**](StoragePureRealmRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 
**Volumes** | Pointer to [**[]StoragePureVolumeRelationship**](StoragePureVolumeRelationship.md) | An array of relationships to storagePureVolume resources. | [optional] [readonly] 

## Methods

### NewStoragePureProtectionGroup

`func NewStoragePureProtectionGroup(classId string, objectType string, ) *StoragePureProtectionGroup`

NewStoragePureProtectionGroup instantiates a new StoragePureProtectionGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoragePureProtectionGroupWithDefaults

`func NewStoragePureProtectionGroupWithDefaults() *StoragePureProtectionGroup`

NewStoragePureProtectionGroupWithDefaults instantiates a new StoragePureProtectionGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StoragePureProtectionGroup) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StoragePureProtectionGroup) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StoragePureProtectionGroup) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StoragePureProtectionGroup) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StoragePureProtectionGroup) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StoragePureProtectionGroup) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetPodName

`func (o *StoragePureProtectionGroup) GetPodName() string`

GetPodName returns the PodName field if non-nil, zero value otherwise.

### GetPodNameOk

`func (o *StoragePureProtectionGroup) GetPodNameOk() (*string, bool)`

GetPodNameOk returns a tuple with the PodName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodName

`func (o *StoragePureProtectionGroup) SetPodName(v string)`

SetPodName sets PodName field to given value.

### HasPodName

`func (o *StoragePureProtectionGroup) HasPodName() bool`

HasPodName returns a boolean if a field has been set.

### GetRealmName

`func (o *StoragePureProtectionGroup) GetRealmName() string`

GetRealmName returns the RealmName field if non-nil, zero value otherwise.

### GetRealmNameOk

`func (o *StoragePureProtectionGroup) GetRealmNameOk() (*string, bool)`

GetRealmNameOk returns a tuple with the RealmName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmName

`func (o *StoragePureProtectionGroup) SetRealmName(v string)`

SetRealmName sets RealmName field to given value.

### HasRealmName

`func (o *StoragePureProtectionGroup) HasRealmName() bool`

HasRealmName returns a boolean if a field has been set.

### GetSize

`func (o *StoragePureProtectionGroup) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *StoragePureProtectionGroup) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *StoragePureProtectionGroup) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *StoragePureProtectionGroup) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSource

`func (o *StoragePureProtectionGroup) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *StoragePureProtectionGroup) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *StoragePureProtectionGroup) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *StoragePureProtectionGroup) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTargets

`func (o *StoragePureProtectionGroup) GetTargets() []string`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *StoragePureProtectionGroup) GetTargetsOk() (*[]string, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *StoragePureProtectionGroup) SetTargets(v []string)`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *StoragePureProtectionGroup) HasTargets() bool`

HasTargets returns a boolean if a field has been set.

### SetTargetsNil

`func (o *StoragePureProtectionGroup) SetTargetsNil(b bool)`

 SetTargetsNil sets the value for Targets to be an explicit nil

### UnsetTargets
`func (o *StoragePureProtectionGroup) UnsetTargets()`

UnsetTargets ensures that no value is present for Targets, not even an explicit nil
### GetArray

`func (o *StoragePureProtectionGroup) GetArray() StoragePureArrayRelationship`

GetArray returns the Array field if non-nil, zero value otherwise.

### GetArrayOk

`func (o *StoragePureProtectionGroup) GetArrayOk() (*StoragePureArrayRelationship, bool)`

GetArrayOk returns a tuple with the Array field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArray

`func (o *StoragePureProtectionGroup) SetArray(v StoragePureArrayRelationship)`

SetArray sets Array field to given value.

### HasArray

`func (o *StoragePureProtectionGroup) HasArray() bool`

HasArray returns a boolean if a field has been set.

### SetArrayNil

`func (o *StoragePureProtectionGroup) SetArrayNil(b bool)`

 SetArrayNil sets the value for Array to be an explicit nil

### UnsetArray
`func (o *StoragePureProtectionGroup) UnsetArray()`

UnsetArray ensures that no value is present for Array, not even an explicit nil
### GetHostGroups

`func (o *StoragePureProtectionGroup) GetHostGroups() []StoragePureHostGroupRelationship`

GetHostGroups returns the HostGroups field if non-nil, zero value otherwise.

### GetHostGroupsOk

`func (o *StoragePureProtectionGroup) GetHostGroupsOk() (*[]StoragePureHostGroupRelationship, bool)`

GetHostGroupsOk returns a tuple with the HostGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostGroups

`func (o *StoragePureProtectionGroup) SetHostGroups(v []StoragePureHostGroupRelationship)`

SetHostGroups sets HostGroups field to given value.

### HasHostGroups

`func (o *StoragePureProtectionGroup) HasHostGroups() bool`

HasHostGroups returns a boolean if a field has been set.

### SetHostGroupsNil

`func (o *StoragePureProtectionGroup) SetHostGroupsNil(b bool)`

 SetHostGroupsNil sets the value for HostGroups to be an explicit nil

### UnsetHostGroups
`func (o *StoragePureProtectionGroup) UnsetHostGroups()`

UnsetHostGroups ensures that no value is present for HostGroups, not even an explicit nil
### GetHosts

`func (o *StoragePureProtectionGroup) GetHosts() []StoragePureHostRelationship`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *StoragePureProtectionGroup) GetHostsOk() (*[]StoragePureHostRelationship, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *StoragePureProtectionGroup) SetHosts(v []StoragePureHostRelationship)`

SetHosts sets Hosts field to given value.

### HasHosts

`func (o *StoragePureProtectionGroup) HasHosts() bool`

HasHosts returns a boolean if a field has been set.

### SetHostsNil

`func (o *StoragePureProtectionGroup) SetHostsNil(b bool)`

 SetHostsNil sets the value for Hosts to be an explicit nil

### UnsetHosts
`func (o *StoragePureProtectionGroup) UnsetHosts()`

UnsetHosts ensures that no value is present for Hosts, not even an explicit nil
### GetPod

`func (o *StoragePureProtectionGroup) GetPod() StoragePurePodRelationship`

GetPod returns the Pod field if non-nil, zero value otherwise.

### GetPodOk

`func (o *StoragePureProtectionGroup) GetPodOk() (*StoragePurePodRelationship, bool)`

GetPodOk returns a tuple with the Pod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPod

`func (o *StoragePureProtectionGroup) SetPod(v StoragePurePodRelationship)`

SetPod sets Pod field to given value.

### HasPod

`func (o *StoragePureProtectionGroup) HasPod() bool`

HasPod returns a boolean if a field has been set.

### SetPodNil

`func (o *StoragePureProtectionGroup) SetPodNil(b bool)`

 SetPodNil sets the value for Pod to be an explicit nil

### UnsetPod
`func (o *StoragePureProtectionGroup) UnsetPod()`

UnsetPod ensures that no value is present for Pod, not even an explicit nil
### GetRealm

`func (o *StoragePureProtectionGroup) GetRealm() StoragePureRealmRelationship`

GetRealm returns the Realm field if non-nil, zero value otherwise.

### GetRealmOk

`func (o *StoragePureProtectionGroup) GetRealmOk() (*StoragePureRealmRelationship, bool)`

GetRealmOk returns a tuple with the Realm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealm

`func (o *StoragePureProtectionGroup) SetRealm(v StoragePureRealmRelationship)`

SetRealm sets Realm field to given value.

### HasRealm

`func (o *StoragePureProtectionGroup) HasRealm() bool`

HasRealm returns a boolean if a field has been set.

### SetRealmNil

`func (o *StoragePureProtectionGroup) SetRealmNil(b bool)`

 SetRealmNil sets the value for Realm to be an explicit nil

### UnsetRealm
`func (o *StoragePureProtectionGroup) UnsetRealm()`

UnsetRealm ensures that no value is present for Realm, not even an explicit nil
### GetRegisteredDevice

`func (o *StoragePureProtectionGroup) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StoragePureProtectionGroup) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StoragePureProtectionGroup) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StoragePureProtectionGroup) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *StoragePureProtectionGroup) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *StoragePureProtectionGroup) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
### GetVolumes

`func (o *StoragePureProtectionGroup) GetVolumes() []StoragePureVolumeRelationship`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *StoragePureProtectionGroup) GetVolumesOk() (*[]StoragePureVolumeRelationship, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *StoragePureProtectionGroup) SetVolumes(v []StoragePureVolumeRelationship)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *StoragePureProtectionGroup) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### SetVolumesNil

`func (o *StoragePureProtectionGroup) SetVolumesNil(b bool)`

 SetVolumesNil sets the value for Volumes to be an explicit nil

### UnsetVolumes
`func (o *StoragePureProtectionGroup) UnsetVolumes()`

UnsetVolumes ensures that no value is present for Volumes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



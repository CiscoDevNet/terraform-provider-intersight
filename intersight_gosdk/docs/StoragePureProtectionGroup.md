# StoragePureProtectionGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Size** | Pointer to **int64** | Overall size of all snapshots in the protection group, represented in bytes. | [optional] 
**Source** | Pointer to **string** | Name of PureStorage array name on which the protection group is created. | [optional] [readonly] 
**Targets** | Pointer to **[]string** |  | [optional] 
**Array** | Pointer to [**StoragePureArrayRelationship**](storage.PureArray.Relationship.md) |  | [optional] 
**HostGroups** | Pointer to [**[]StoragePureHostGroupRelationship**](storage.PureHostGroup.Relationship.md) | An array of relationships to storagePureHostGroup resources. | [optional] [readonly] 
**Hosts** | Pointer to [**[]StoragePureHostRelationship**](storage.PureHost.Relationship.md) | An array of relationships to storagePureHost resources. | [optional] [readonly] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 
**Volumes** | Pointer to [**[]StoragePureVolumeRelationship**](storage.PureVolume.Relationship.md) | An array of relationships to storagePureVolume resources. | [optional] [readonly] 

## Methods

### NewStoragePureProtectionGroup

`func NewStoragePureProtectionGroup() *StoragePureProtectionGroup`

NewStoragePureProtectionGroup instantiates a new StoragePureProtectionGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoragePureProtectionGroupWithDefaults

`func NewStoragePureProtectionGroupWithDefaults() *StoragePureProtectionGroup`

NewStoragePureProtectionGroupWithDefaults instantiates a new StoragePureProtectionGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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



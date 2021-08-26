# StoragePureProtectionGroupAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.PureProtectionGroup"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.PureProtectionGroup"]
**Size** | Pointer to **int64** | Overall size of all snapshots in the protection group, represented in bytes. | [optional] 
**Source** | Pointer to **string** | Name of PureStorage array name on which the protection group is created. | [optional] [readonly] 
**Targets** | Pointer to **[]string** |  | [optional] 
**Array** | Pointer to [**StoragePureArrayRelationship**](StoragePureArrayRelationship.md) |  | [optional] 
**HostGroups** | Pointer to [**[]StoragePureHostGroupRelationship**](StoragePureHostGroupRelationship.md) | An array of relationships to storagePureHostGroup resources. | [optional] [readonly] 
**Hosts** | Pointer to [**[]StoragePureHostRelationship**](StoragePureHostRelationship.md) | An array of relationships to storagePureHost resources. | [optional] [readonly] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 
**Volumes** | Pointer to [**[]StoragePureVolumeRelationship**](StoragePureVolumeRelationship.md) | An array of relationships to storagePureVolume resources. | [optional] [readonly] 

## Methods

### NewStoragePureProtectionGroupAllOf

`func NewStoragePureProtectionGroupAllOf(classId string, objectType string, ) *StoragePureProtectionGroupAllOf`

NewStoragePureProtectionGroupAllOf instantiates a new StoragePureProtectionGroupAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoragePureProtectionGroupAllOfWithDefaults

`func NewStoragePureProtectionGroupAllOfWithDefaults() *StoragePureProtectionGroupAllOf`

NewStoragePureProtectionGroupAllOfWithDefaults instantiates a new StoragePureProtectionGroupAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StoragePureProtectionGroupAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StoragePureProtectionGroupAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StoragePureProtectionGroupAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StoragePureProtectionGroupAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StoragePureProtectionGroupAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StoragePureProtectionGroupAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetSize

`func (o *StoragePureProtectionGroupAllOf) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *StoragePureProtectionGroupAllOf) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *StoragePureProtectionGroupAllOf) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *StoragePureProtectionGroupAllOf) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSource

`func (o *StoragePureProtectionGroupAllOf) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *StoragePureProtectionGroupAllOf) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *StoragePureProtectionGroupAllOf) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *StoragePureProtectionGroupAllOf) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTargets

`func (o *StoragePureProtectionGroupAllOf) GetTargets() []string`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *StoragePureProtectionGroupAllOf) GetTargetsOk() (*[]string, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *StoragePureProtectionGroupAllOf) SetTargets(v []string)`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *StoragePureProtectionGroupAllOf) HasTargets() bool`

HasTargets returns a boolean if a field has been set.

### SetTargetsNil

`func (o *StoragePureProtectionGroupAllOf) SetTargetsNil(b bool)`

 SetTargetsNil sets the value for Targets to be an explicit nil

### UnsetTargets
`func (o *StoragePureProtectionGroupAllOf) UnsetTargets()`

UnsetTargets ensures that no value is present for Targets, not even an explicit nil
### GetArray

`func (o *StoragePureProtectionGroupAllOf) GetArray() StoragePureArrayRelationship`

GetArray returns the Array field if non-nil, zero value otherwise.

### GetArrayOk

`func (o *StoragePureProtectionGroupAllOf) GetArrayOk() (*StoragePureArrayRelationship, bool)`

GetArrayOk returns a tuple with the Array field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArray

`func (o *StoragePureProtectionGroupAllOf) SetArray(v StoragePureArrayRelationship)`

SetArray sets Array field to given value.

### HasArray

`func (o *StoragePureProtectionGroupAllOf) HasArray() bool`

HasArray returns a boolean if a field has been set.

### GetHostGroups

`func (o *StoragePureProtectionGroupAllOf) GetHostGroups() []StoragePureHostGroupRelationship`

GetHostGroups returns the HostGroups field if non-nil, zero value otherwise.

### GetHostGroupsOk

`func (o *StoragePureProtectionGroupAllOf) GetHostGroupsOk() (*[]StoragePureHostGroupRelationship, bool)`

GetHostGroupsOk returns a tuple with the HostGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostGroups

`func (o *StoragePureProtectionGroupAllOf) SetHostGroups(v []StoragePureHostGroupRelationship)`

SetHostGroups sets HostGroups field to given value.

### HasHostGroups

`func (o *StoragePureProtectionGroupAllOf) HasHostGroups() bool`

HasHostGroups returns a boolean if a field has been set.

### SetHostGroupsNil

`func (o *StoragePureProtectionGroupAllOf) SetHostGroupsNil(b bool)`

 SetHostGroupsNil sets the value for HostGroups to be an explicit nil

### UnsetHostGroups
`func (o *StoragePureProtectionGroupAllOf) UnsetHostGroups()`

UnsetHostGroups ensures that no value is present for HostGroups, not even an explicit nil
### GetHosts

`func (o *StoragePureProtectionGroupAllOf) GetHosts() []StoragePureHostRelationship`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *StoragePureProtectionGroupAllOf) GetHostsOk() (*[]StoragePureHostRelationship, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *StoragePureProtectionGroupAllOf) SetHosts(v []StoragePureHostRelationship)`

SetHosts sets Hosts field to given value.

### HasHosts

`func (o *StoragePureProtectionGroupAllOf) HasHosts() bool`

HasHosts returns a boolean if a field has been set.

### SetHostsNil

`func (o *StoragePureProtectionGroupAllOf) SetHostsNil(b bool)`

 SetHostsNil sets the value for Hosts to be an explicit nil

### UnsetHosts
`func (o *StoragePureProtectionGroupAllOf) UnsetHosts()`

UnsetHosts ensures that no value is present for Hosts, not even an explicit nil
### GetRegisteredDevice

`func (o *StoragePureProtectionGroupAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StoragePureProtectionGroupAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StoragePureProtectionGroupAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StoragePureProtectionGroupAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### GetVolumes

`func (o *StoragePureProtectionGroupAllOf) GetVolumes() []StoragePureVolumeRelationship`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *StoragePureProtectionGroupAllOf) GetVolumesOk() (*[]StoragePureVolumeRelationship, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *StoragePureProtectionGroupAllOf) SetVolumes(v []StoragePureVolumeRelationship)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *StoragePureProtectionGroupAllOf) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### SetVolumesNil

`func (o *StoragePureProtectionGroupAllOf) SetVolumesNil(b bool)`

 SetVolumesNil sets the value for Volumes to be an explicit nil

### UnsetVolumes
`func (o *StoragePureProtectionGroupAllOf) UnsetVolumes()`

UnsetVolumes ensures that no value is present for Volumes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



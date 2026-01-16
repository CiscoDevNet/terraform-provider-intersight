# StoragePureVolume

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.PureVolume"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.PureVolume"]
**Created** | Pointer to **time.Time** | Creation time of the volume. | [optional] [readonly] 
**PodName** | Pointer to **string** | A pod representing a collection of protection groups and volumes is created on one array and stretched to another array, resulting in fully synchronized writes between the two arrays. | [optional] [readonly] 
**RealmName** | Pointer to **string** | A realm is the core multi-tenancy component on a Pure Flash Array, providing a self-contained, virtual storage environment with dedicated policies and quotas for secure data isolation and predictable performance. | [optional] [readonly] 
**Serial** | Pointer to **string** | Serial number of the volume. | [optional] [readonly] 
**Source** | Pointer to **string** | Source from which the volume is created. Applicable only if the volume is cloned from other volume or snapshot. | [optional] [readonly] 
**Array** | Pointer to [**NullableStoragePureArrayRelationship**](StoragePureArrayRelationship.md) |  | [optional] 
**Pod** | Pointer to [**NullableStoragePurePodRelationship**](StoragePurePodRelationship.md) |  | [optional] 
**ProtectionGroup** | Pointer to [**NullableStoragePureProtectionGroupRelationship**](StoragePureProtectionGroupRelationship.md) |  | [optional] 
**Realm** | Pointer to [**NullableStoragePureRealmRelationship**](StoragePureRealmRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewStoragePureVolume

`func NewStoragePureVolume(classId string, objectType string, ) *StoragePureVolume`

NewStoragePureVolume instantiates a new StoragePureVolume object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoragePureVolumeWithDefaults

`func NewStoragePureVolumeWithDefaults() *StoragePureVolume`

NewStoragePureVolumeWithDefaults instantiates a new StoragePureVolume object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StoragePureVolume) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StoragePureVolume) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StoragePureVolume) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StoragePureVolume) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StoragePureVolume) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StoragePureVolume) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCreated

`func (o *StoragePureVolume) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *StoragePureVolume) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *StoragePureVolume) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *StoragePureVolume) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetPodName

`func (o *StoragePureVolume) GetPodName() string`

GetPodName returns the PodName field if non-nil, zero value otherwise.

### GetPodNameOk

`func (o *StoragePureVolume) GetPodNameOk() (*string, bool)`

GetPodNameOk returns a tuple with the PodName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodName

`func (o *StoragePureVolume) SetPodName(v string)`

SetPodName sets PodName field to given value.

### HasPodName

`func (o *StoragePureVolume) HasPodName() bool`

HasPodName returns a boolean if a field has been set.

### GetRealmName

`func (o *StoragePureVolume) GetRealmName() string`

GetRealmName returns the RealmName field if non-nil, zero value otherwise.

### GetRealmNameOk

`func (o *StoragePureVolume) GetRealmNameOk() (*string, bool)`

GetRealmNameOk returns a tuple with the RealmName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmName

`func (o *StoragePureVolume) SetRealmName(v string)`

SetRealmName sets RealmName field to given value.

### HasRealmName

`func (o *StoragePureVolume) HasRealmName() bool`

HasRealmName returns a boolean if a field has been set.

### GetSerial

`func (o *StoragePureVolume) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *StoragePureVolume) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *StoragePureVolume) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *StoragePureVolume) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetSource

`func (o *StoragePureVolume) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *StoragePureVolume) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *StoragePureVolume) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *StoragePureVolume) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetArray

`func (o *StoragePureVolume) GetArray() StoragePureArrayRelationship`

GetArray returns the Array field if non-nil, zero value otherwise.

### GetArrayOk

`func (o *StoragePureVolume) GetArrayOk() (*StoragePureArrayRelationship, bool)`

GetArrayOk returns a tuple with the Array field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArray

`func (o *StoragePureVolume) SetArray(v StoragePureArrayRelationship)`

SetArray sets Array field to given value.

### HasArray

`func (o *StoragePureVolume) HasArray() bool`

HasArray returns a boolean if a field has been set.

### SetArrayNil

`func (o *StoragePureVolume) SetArrayNil(b bool)`

 SetArrayNil sets the value for Array to be an explicit nil

### UnsetArray
`func (o *StoragePureVolume) UnsetArray()`

UnsetArray ensures that no value is present for Array, not even an explicit nil
### GetPod

`func (o *StoragePureVolume) GetPod() StoragePurePodRelationship`

GetPod returns the Pod field if non-nil, zero value otherwise.

### GetPodOk

`func (o *StoragePureVolume) GetPodOk() (*StoragePurePodRelationship, bool)`

GetPodOk returns a tuple with the Pod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPod

`func (o *StoragePureVolume) SetPod(v StoragePurePodRelationship)`

SetPod sets Pod field to given value.

### HasPod

`func (o *StoragePureVolume) HasPod() bool`

HasPod returns a boolean if a field has been set.

### SetPodNil

`func (o *StoragePureVolume) SetPodNil(b bool)`

 SetPodNil sets the value for Pod to be an explicit nil

### UnsetPod
`func (o *StoragePureVolume) UnsetPod()`

UnsetPod ensures that no value is present for Pod, not even an explicit nil
### GetProtectionGroup

`func (o *StoragePureVolume) GetProtectionGroup() StoragePureProtectionGroupRelationship`

GetProtectionGroup returns the ProtectionGroup field if non-nil, zero value otherwise.

### GetProtectionGroupOk

`func (o *StoragePureVolume) GetProtectionGroupOk() (*StoragePureProtectionGroupRelationship, bool)`

GetProtectionGroupOk returns a tuple with the ProtectionGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroup

`func (o *StoragePureVolume) SetProtectionGroup(v StoragePureProtectionGroupRelationship)`

SetProtectionGroup sets ProtectionGroup field to given value.

### HasProtectionGroup

`func (o *StoragePureVolume) HasProtectionGroup() bool`

HasProtectionGroup returns a boolean if a field has been set.

### SetProtectionGroupNil

`func (o *StoragePureVolume) SetProtectionGroupNil(b bool)`

 SetProtectionGroupNil sets the value for ProtectionGroup to be an explicit nil

### UnsetProtectionGroup
`func (o *StoragePureVolume) UnsetProtectionGroup()`

UnsetProtectionGroup ensures that no value is present for ProtectionGroup, not even an explicit nil
### GetRealm

`func (o *StoragePureVolume) GetRealm() StoragePureRealmRelationship`

GetRealm returns the Realm field if non-nil, zero value otherwise.

### GetRealmOk

`func (o *StoragePureVolume) GetRealmOk() (*StoragePureRealmRelationship, bool)`

GetRealmOk returns a tuple with the Realm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealm

`func (o *StoragePureVolume) SetRealm(v StoragePureRealmRelationship)`

SetRealm sets Realm field to given value.

### HasRealm

`func (o *StoragePureVolume) HasRealm() bool`

HasRealm returns a boolean if a field has been set.

### SetRealmNil

`func (o *StoragePureVolume) SetRealmNil(b bool)`

 SetRealmNil sets the value for Realm to be an explicit nil

### UnsetRealm
`func (o *StoragePureVolume) UnsetRealm()`

UnsetRealm ensures that no value is present for Realm, not even an explicit nil
### GetRegisteredDevice

`func (o *StoragePureVolume) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StoragePureVolume) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StoragePureVolume) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StoragePureVolume) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *StoragePureVolume) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *StoragePureVolume) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



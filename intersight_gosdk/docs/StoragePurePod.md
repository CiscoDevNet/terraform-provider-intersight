# StoragePurePod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.PurePod"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.PurePod"]
**ArrayCount** | Pointer to **int64** | Number of arrays in the pod. | [optional] [readonly] 
**Destroyed** | Pointer to **bool** | Indicates whether the pod has been destroyed and is pending eradication. | [optional] [readonly] 
**LinkTargetCount** | Pointer to **int64** | Number of link targets in the pod. | [optional] [readonly] 
**Mediator** | Pointer to **string** | Name of the mediator managing synchronous replication between arrays. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the pod as configured on the Pure array. | [optional] [readonly] 
**PromotionStatus** | Pointer to **string** | Promotion status of the pod. It can be one of the following values Promoting, Promoted, Demoted. | [optional] [readonly] 
**RealmName** | Pointer to **string** | A realm is the core multi-tenancy component on a Pure Flash Array, providing a self-contained, virtual storage environment with dedicated policies and quotas for secure data isolation and predictable performance. | [optional] [readonly] 
**StorageUtilization** | Pointer to [**NullableStoragePurePodUtilization**](StoragePurePodUtilization.md) |  | [optional] 
**Array** | Pointer to [**NullableStoragePureArrayRelationship**](StoragePureArrayRelationship.md) |  | [optional] 
**ProtectionGroups** | Pointer to [**[]StoragePureProtectionGroupRelationship**](StoragePureProtectionGroupRelationship.md) | An array of relationships to storagePureProtectionGroup resources. | [optional] [readonly] 
**Realm** | Pointer to [**NullableStoragePureRealmRelationship**](StoragePureRealmRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 
**Volumes** | Pointer to [**[]StoragePureVolumeRelationship**](StoragePureVolumeRelationship.md) | An array of relationships to storagePureVolume resources. | [optional] [readonly] 

## Methods

### NewStoragePurePod

`func NewStoragePurePod(classId string, objectType string, ) *StoragePurePod`

NewStoragePurePod instantiates a new StoragePurePod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoragePurePodWithDefaults

`func NewStoragePurePodWithDefaults() *StoragePurePod`

NewStoragePurePodWithDefaults instantiates a new StoragePurePod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StoragePurePod) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StoragePurePod) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StoragePurePod) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StoragePurePod) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StoragePurePod) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StoragePurePod) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetArrayCount

`func (o *StoragePurePod) GetArrayCount() int64`

GetArrayCount returns the ArrayCount field if non-nil, zero value otherwise.

### GetArrayCountOk

`func (o *StoragePurePod) GetArrayCountOk() (*int64, bool)`

GetArrayCountOk returns a tuple with the ArrayCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrayCount

`func (o *StoragePurePod) SetArrayCount(v int64)`

SetArrayCount sets ArrayCount field to given value.

### HasArrayCount

`func (o *StoragePurePod) HasArrayCount() bool`

HasArrayCount returns a boolean if a field has been set.

### GetDestroyed

`func (o *StoragePurePod) GetDestroyed() bool`

GetDestroyed returns the Destroyed field if non-nil, zero value otherwise.

### GetDestroyedOk

`func (o *StoragePurePod) GetDestroyedOk() (*bool, bool)`

GetDestroyedOk returns a tuple with the Destroyed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestroyed

`func (o *StoragePurePod) SetDestroyed(v bool)`

SetDestroyed sets Destroyed field to given value.

### HasDestroyed

`func (o *StoragePurePod) HasDestroyed() bool`

HasDestroyed returns a boolean if a field has been set.

### GetLinkTargetCount

`func (o *StoragePurePod) GetLinkTargetCount() int64`

GetLinkTargetCount returns the LinkTargetCount field if non-nil, zero value otherwise.

### GetLinkTargetCountOk

`func (o *StoragePurePod) GetLinkTargetCountOk() (*int64, bool)`

GetLinkTargetCountOk returns a tuple with the LinkTargetCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkTargetCount

`func (o *StoragePurePod) SetLinkTargetCount(v int64)`

SetLinkTargetCount sets LinkTargetCount field to given value.

### HasLinkTargetCount

`func (o *StoragePurePod) HasLinkTargetCount() bool`

HasLinkTargetCount returns a boolean if a field has been set.

### GetMediator

`func (o *StoragePurePod) GetMediator() string`

GetMediator returns the Mediator field if non-nil, zero value otherwise.

### GetMediatorOk

`func (o *StoragePurePod) GetMediatorOk() (*string, bool)`

GetMediatorOk returns a tuple with the Mediator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediator

`func (o *StoragePurePod) SetMediator(v string)`

SetMediator sets Mediator field to given value.

### HasMediator

`func (o *StoragePurePod) HasMediator() bool`

HasMediator returns a boolean if a field has been set.

### GetName

`func (o *StoragePurePod) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StoragePurePod) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StoragePurePod) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StoragePurePod) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPromotionStatus

`func (o *StoragePurePod) GetPromotionStatus() string`

GetPromotionStatus returns the PromotionStatus field if non-nil, zero value otherwise.

### GetPromotionStatusOk

`func (o *StoragePurePod) GetPromotionStatusOk() (*string, bool)`

GetPromotionStatusOk returns a tuple with the PromotionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromotionStatus

`func (o *StoragePurePod) SetPromotionStatus(v string)`

SetPromotionStatus sets PromotionStatus field to given value.

### HasPromotionStatus

`func (o *StoragePurePod) HasPromotionStatus() bool`

HasPromotionStatus returns a boolean if a field has been set.

### GetRealmName

`func (o *StoragePurePod) GetRealmName() string`

GetRealmName returns the RealmName field if non-nil, zero value otherwise.

### GetRealmNameOk

`func (o *StoragePurePod) GetRealmNameOk() (*string, bool)`

GetRealmNameOk returns a tuple with the RealmName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmName

`func (o *StoragePurePod) SetRealmName(v string)`

SetRealmName sets RealmName field to given value.

### HasRealmName

`func (o *StoragePurePod) HasRealmName() bool`

HasRealmName returns a boolean if a field has been set.

### GetStorageUtilization

`func (o *StoragePurePod) GetStorageUtilization() StoragePurePodUtilization`

GetStorageUtilization returns the StorageUtilization field if non-nil, zero value otherwise.

### GetStorageUtilizationOk

`func (o *StoragePurePod) GetStorageUtilizationOk() (*StoragePurePodUtilization, bool)`

GetStorageUtilizationOk returns a tuple with the StorageUtilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageUtilization

`func (o *StoragePurePod) SetStorageUtilization(v StoragePurePodUtilization)`

SetStorageUtilization sets StorageUtilization field to given value.

### HasStorageUtilization

`func (o *StoragePurePod) HasStorageUtilization() bool`

HasStorageUtilization returns a boolean if a field has been set.

### SetStorageUtilizationNil

`func (o *StoragePurePod) SetStorageUtilizationNil(b bool)`

 SetStorageUtilizationNil sets the value for StorageUtilization to be an explicit nil

### UnsetStorageUtilization
`func (o *StoragePurePod) UnsetStorageUtilization()`

UnsetStorageUtilization ensures that no value is present for StorageUtilization, not even an explicit nil
### GetArray

`func (o *StoragePurePod) GetArray() StoragePureArrayRelationship`

GetArray returns the Array field if non-nil, zero value otherwise.

### GetArrayOk

`func (o *StoragePurePod) GetArrayOk() (*StoragePureArrayRelationship, bool)`

GetArrayOk returns a tuple with the Array field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArray

`func (o *StoragePurePod) SetArray(v StoragePureArrayRelationship)`

SetArray sets Array field to given value.

### HasArray

`func (o *StoragePurePod) HasArray() bool`

HasArray returns a boolean if a field has been set.

### SetArrayNil

`func (o *StoragePurePod) SetArrayNil(b bool)`

 SetArrayNil sets the value for Array to be an explicit nil

### UnsetArray
`func (o *StoragePurePod) UnsetArray()`

UnsetArray ensures that no value is present for Array, not even an explicit nil
### GetProtectionGroups

`func (o *StoragePurePod) GetProtectionGroups() []StoragePureProtectionGroupRelationship`

GetProtectionGroups returns the ProtectionGroups field if non-nil, zero value otherwise.

### GetProtectionGroupsOk

`func (o *StoragePurePod) GetProtectionGroupsOk() (*[]StoragePureProtectionGroupRelationship, bool)`

GetProtectionGroupsOk returns a tuple with the ProtectionGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroups

`func (o *StoragePurePod) SetProtectionGroups(v []StoragePureProtectionGroupRelationship)`

SetProtectionGroups sets ProtectionGroups field to given value.

### HasProtectionGroups

`func (o *StoragePurePod) HasProtectionGroups() bool`

HasProtectionGroups returns a boolean if a field has been set.

### SetProtectionGroupsNil

`func (o *StoragePurePod) SetProtectionGroupsNil(b bool)`

 SetProtectionGroupsNil sets the value for ProtectionGroups to be an explicit nil

### UnsetProtectionGroups
`func (o *StoragePurePod) UnsetProtectionGroups()`

UnsetProtectionGroups ensures that no value is present for ProtectionGroups, not even an explicit nil
### GetRealm

`func (o *StoragePurePod) GetRealm() StoragePureRealmRelationship`

GetRealm returns the Realm field if non-nil, zero value otherwise.

### GetRealmOk

`func (o *StoragePurePod) GetRealmOk() (*StoragePureRealmRelationship, bool)`

GetRealmOk returns a tuple with the Realm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealm

`func (o *StoragePurePod) SetRealm(v StoragePureRealmRelationship)`

SetRealm sets Realm field to given value.

### HasRealm

`func (o *StoragePurePod) HasRealm() bool`

HasRealm returns a boolean if a field has been set.

### SetRealmNil

`func (o *StoragePurePod) SetRealmNil(b bool)`

 SetRealmNil sets the value for Realm to be an explicit nil

### UnsetRealm
`func (o *StoragePurePod) UnsetRealm()`

UnsetRealm ensures that no value is present for Realm, not even an explicit nil
### GetRegisteredDevice

`func (o *StoragePurePod) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StoragePurePod) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StoragePurePod) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StoragePurePod) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *StoragePurePod) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *StoragePurePod) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
### GetVolumes

`func (o *StoragePurePod) GetVolumes() []StoragePureVolumeRelationship`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *StoragePurePod) GetVolumesOk() (*[]StoragePureVolumeRelationship, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *StoragePurePod) SetVolumes(v []StoragePureVolumeRelationship)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *StoragePurePod) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### SetVolumesNil

`func (o *StoragePurePod) SetVolumesNil(b bool)`

 SetVolumesNil sets the value for Volumes to be an explicit nil

### UnsetVolumes
`func (o *StoragePurePod) UnsetVolumes()`

UnsetVolumes ensures that no value is present for Volumes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



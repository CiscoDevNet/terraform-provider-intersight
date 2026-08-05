# HciExternalStorage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hci.ExternalStorage"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hci.ExternalStorage"]
**ExternalStorageExtId** | Pointer to **string** | The external identifier of the external storage instance. | [optional] [readonly] 
**FreeCapacityBytes** | Pointer to **int64** | Free capacity of the external storage in bytes. | [optional] [readonly] 
**Health** | Pointer to **string** | Health status of the external storage. Iris API does not populate health, Janus API does. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the external-storage instance. | [optional] [readonly] 
**PodExtId** | Pointer to **string** | Id of the pod this external-storage connects to. | [optional] [readonly] 
**PodName** | Pointer to **string** | Name of the pod this external-storage connects to. | [optional] [readonly] 
**ProviderType** | Pointer to **string** | The type of storage provider. Normalized to EVERPURE_FLASHARRAY for Pure Storage FlashArray (API may return either PURE_STORAGE_FLASHARRAY or EVERPURE_FLASHARRAY). | [optional] [readonly] 
**PureArrayInfos** | Pointer to [**[]HciExternalStorageArray**](HciExternalStorageArray.md) |  | [optional] 
**RealmName** | Pointer to **string** | Realm of the pod this external-storage connects to. | [optional] [readonly] 
**TotalCapacityBytes** | Pointer to **int64** | Total capacity of the external storage in bytes. | [optional] [readonly] 
**PurePod** | Pointer to [**NullableStoragePurePodRelationship**](StoragePurePodRelationship.md) |  | [optional] 
**PureRealm** | Pointer to [**NullableStoragePureRealmRelationship**](StoragePureRealmRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 
**StorageContainers** | Pointer to [**[]HciStorageContainerRelationship**](HciStorageContainerRelationship.md) | An array of relationships to hciStorageContainer resources. | [optional] [readonly] 
**StorageSystems** | Pointer to [**[]StorageBaseArrayRelationship**](StorageBaseArrayRelationship.md) | An array of relationships to storageBaseArray resources. | [optional] [readonly] 

## Methods

### NewHciExternalStorage

`func NewHciExternalStorage(classId string, objectType string, ) *HciExternalStorage`

NewHciExternalStorage instantiates a new HciExternalStorage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHciExternalStorageWithDefaults

`func NewHciExternalStorageWithDefaults() *HciExternalStorage`

NewHciExternalStorageWithDefaults instantiates a new HciExternalStorage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HciExternalStorage) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HciExternalStorage) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HciExternalStorage) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HciExternalStorage) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HciExternalStorage) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HciExternalStorage) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetExternalStorageExtId

`func (o *HciExternalStorage) GetExternalStorageExtId() string`

GetExternalStorageExtId returns the ExternalStorageExtId field if non-nil, zero value otherwise.

### GetExternalStorageExtIdOk

`func (o *HciExternalStorage) GetExternalStorageExtIdOk() (*string, bool)`

GetExternalStorageExtIdOk returns a tuple with the ExternalStorageExtId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalStorageExtId

`func (o *HciExternalStorage) SetExternalStorageExtId(v string)`

SetExternalStorageExtId sets ExternalStorageExtId field to given value.

### HasExternalStorageExtId

`func (o *HciExternalStorage) HasExternalStorageExtId() bool`

HasExternalStorageExtId returns a boolean if a field has been set.

### GetFreeCapacityBytes

`func (o *HciExternalStorage) GetFreeCapacityBytes() int64`

GetFreeCapacityBytes returns the FreeCapacityBytes field if non-nil, zero value otherwise.

### GetFreeCapacityBytesOk

`func (o *HciExternalStorage) GetFreeCapacityBytesOk() (*int64, bool)`

GetFreeCapacityBytesOk returns a tuple with the FreeCapacityBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeCapacityBytes

`func (o *HciExternalStorage) SetFreeCapacityBytes(v int64)`

SetFreeCapacityBytes sets FreeCapacityBytes field to given value.

### HasFreeCapacityBytes

`func (o *HciExternalStorage) HasFreeCapacityBytes() bool`

HasFreeCapacityBytes returns a boolean if a field has been set.

### GetHealth

`func (o *HciExternalStorage) GetHealth() string`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *HciExternalStorage) GetHealthOk() (*string, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *HciExternalStorage) SetHealth(v string)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *HciExternalStorage) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetName

`func (o *HciExternalStorage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HciExternalStorage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HciExternalStorage) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HciExternalStorage) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPodExtId

`func (o *HciExternalStorage) GetPodExtId() string`

GetPodExtId returns the PodExtId field if non-nil, zero value otherwise.

### GetPodExtIdOk

`func (o *HciExternalStorage) GetPodExtIdOk() (*string, bool)`

GetPodExtIdOk returns a tuple with the PodExtId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodExtId

`func (o *HciExternalStorage) SetPodExtId(v string)`

SetPodExtId sets PodExtId field to given value.

### HasPodExtId

`func (o *HciExternalStorage) HasPodExtId() bool`

HasPodExtId returns a boolean if a field has been set.

### GetPodName

`func (o *HciExternalStorage) GetPodName() string`

GetPodName returns the PodName field if non-nil, zero value otherwise.

### GetPodNameOk

`func (o *HciExternalStorage) GetPodNameOk() (*string, bool)`

GetPodNameOk returns a tuple with the PodName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodName

`func (o *HciExternalStorage) SetPodName(v string)`

SetPodName sets PodName field to given value.

### HasPodName

`func (o *HciExternalStorage) HasPodName() bool`

HasPodName returns a boolean if a field has been set.

### GetProviderType

`func (o *HciExternalStorage) GetProviderType() string`

GetProviderType returns the ProviderType field if non-nil, zero value otherwise.

### GetProviderTypeOk

`func (o *HciExternalStorage) GetProviderTypeOk() (*string, bool)`

GetProviderTypeOk returns a tuple with the ProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderType

`func (o *HciExternalStorage) SetProviderType(v string)`

SetProviderType sets ProviderType field to given value.

### HasProviderType

`func (o *HciExternalStorage) HasProviderType() bool`

HasProviderType returns a boolean if a field has been set.

### GetPureArrayInfos

`func (o *HciExternalStorage) GetPureArrayInfos() []HciExternalStorageArray`

GetPureArrayInfos returns the PureArrayInfos field if non-nil, zero value otherwise.

### GetPureArrayInfosOk

`func (o *HciExternalStorage) GetPureArrayInfosOk() (*[]HciExternalStorageArray, bool)`

GetPureArrayInfosOk returns a tuple with the PureArrayInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPureArrayInfos

`func (o *HciExternalStorage) SetPureArrayInfos(v []HciExternalStorageArray)`

SetPureArrayInfos sets PureArrayInfos field to given value.

### HasPureArrayInfos

`func (o *HciExternalStorage) HasPureArrayInfos() bool`

HasPureArrayInfos returns a boolean if a field has been set.

### SetPureArrayInfosNil

`func (o *HciExternalStorage) SetPureArrayInfosNil(b bool)`

 SetPureArrayInfosNil sets the value for PureArrayInfos to be an explicit nil

### UnsetPureArrayInfos
`func (o *HciExternalStorage) UnsetPureArrayInfos()`

UnsetPureArrayInfos ensures that no value is present for PureArrayInfos, not even an explicit nil
### GetRealmName

`func (o *HciExternalStorage) GetRealmName() string`

GetRealmName returns the RealmName field if non-nil, zero value otherwise.

### GetRealmNameOk

`func (o *HciExternalStorage) GetRealmNameOk() (*string, bool)`

GetRealmNameOk returns a tuple with the RealmName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmName

`func (o *HciExternalStorage) SetRealmName(v string)`

SetRealmName sets RealmName field to given value.

### HasRealmName

`func (o *HciExternalStorage) HasRealmName() bool`

HasRealmName returns a boolean if a field has been set.

### GetTotalCapacityBytes

`func (o *HciExternalStorage) GetTotalCapacityBytes() int64`

GetTotalCapacityBytes returns the TotalCapacityBytes field if non-nil, zero value otherwise.

### GetTotalCapacityBytesOk

`func (o *HciExternalStorage) GetTotalCapacityBytesOk() (*int64, bool)`

GetTotalCapacityBytesOk returns a tuple with the TotalCapacityBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCapacityBytes

`func (o *HciExternalStorage) SetTotalCapacityBytes(v int64)`

SetTotalCapacityBytes sets TotalCapacityBytes field to given value.

### HasTotalCapacityBytes

`func (o *HciExternalStorage) HasTotalCapacityBytes() bool`

HasTotalCapacityBytes returns a boolean if a field has been set.

### GetPurePod

`func (o *HciExternalStorage) GetPurePod() StoragePurePodRelationship`

GetPurePod returns the PurePod field if non-nil, zero value otherwise.

### GetPurePodOk

`func (o *HciExternalStorage) GetPurePodOk() (*StoragePurePodRelationship, bool)`

GetPurePodOk returns a tuple with the PurePod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurePod

`func (o *HciExternalStorage) SetPurePod(v StoragePurePodRelationship)`

SetPurePod sets PurePod field to given value.

### HasPurePod

`func (o *HciExternalStorage) HasPurePod() bool`

HasPurePod returns a boolean if a field has been set.

### SetPurePodNil

`func (o *HciExternalStorage) SetPurePodNil(b bool)`

 SetPurePodNil sets the value for PurePod to be an explicit nil

### UnsetPurePod
`func (o *HciExternalStorage) UnsetPurePod()`

UnsetPurePod ensures that no value is present for PurePod, not even an explicit nil
### GetPureRealm

`func (o *HciExternalStorage) GetPureRealm() StoragePureRealmRelationship`

GetPureRealm returns the PureRealm field if non-nil, zero value otherwise.

### GetPureRealmOk

`func (o *HciExternalStorage) GetPureRealmOk() (*StoragePureRealmRelationship, bool)`

GetPureRealmOk returns a tuple with the PureRealm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPureRealm

`func (o *HciExternalStorage) SetPureRealm(v StoragePureRealmRelationship)`

SetPureRealm sets PureRealm field to given value.

### HasPureRealm

`func (o *HciExternalStorage) HasPureRealm() bool`

HasPureRealm returns a boolean if a field has been set.

### SetPureRealmNil

`func (o *HciExternalStorage) SetPureRealmNil(b bool)`

 SetPureRealmNil sets the value for PureRealm to be an explicit nil

### UnsetPureRealm
`func (o *HciExternalStorage) UnsetPureRealm()`

UnsetPureRealm ensures that no value is present for PureRealm, not even an explicit nil
### GetRegisteredDevice

`func (o *HciExternalStorage) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *HciExternalStorage) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *HciExternalStorage) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *HciExternalStorage) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *HciExternalStorage) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *HciExternalStorage) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
### GetStorageContainers

`func (o *HciExternalStorage) GetStorageContainers() []HciStorageContainerRelationship`

GetStorageContainers returns the StorageContainers field if non-nil, zero value otherwise.

### GetStorageContainersOk

`func (o *HciExternalStorage) GetStorageContainersOk() (*[]HciStorageContainerRelationship, bool)`

GetStorageContainersOk returns a tuple with the StorageContainers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageContainers

`func (o *HciExternalStorage) SetStorageContainers(v []HciStorageContainerRelationship)`

SetStorageContainers sets StorageContainers field to given value.

### HasStorageContainers

`func (o *HciExternalStorage) HasStorageContainers() bool`

HasStorageContainers returns a boolean if a field has been set.

### SetStorageContainersNil

`func (o *HciExternalStorage) SetStorageContainersNil(b bool)`

 SetStorageContainersNil sets the value for StorageContainers to be an explicit nil

### UnsetStorageContainers
`func (o *HciExternalStorage) UnsetStorageContainers()`

UnsetStorageContainers ensures that no value is present for StorageContainers, not even an explicit nil
### GetStorageSystems

`func (o *HciExternalStorage) GetStorageSystems() []StorageBaseArrayRelationship`

GetStorageSystems returns the StorageSystems field if non-nil, zero value otherwise.

### GetStorageSystemsOk

`func (o *HciExternalStorage) GetStorageSystemsOk() (*[]StorageBaseArrayRelationship, bool)`

GetStorageSystemsOk returns a tuple with the StorageSystems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSystems

`func (o *HciExternalStorage) SetStorageSystems(v []StorageBaseArrayRelationship)`

SetStorageSystems sets StorageSystems field to given value.

### HasStorageSystems

`func (o *HciExternalStorage) HasStorageSystems() bool`

HasStorageSystems returns a boolean if a field has been set.

### SetStorageSystemsNil

`func (o *HciExternalStorage) SetStorageSystemsNil(b bool)`

 SetStorageSystemsNil sets the value for StorageSystems to be an explicit nil

### UnsetStorageSystems
`func (o *HciExternalStorage) UnsetStorageSystems()`

UnsetStorageSystems ensures that no value is present for StorageSystems, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



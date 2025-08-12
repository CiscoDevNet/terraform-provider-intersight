# StoragePureProtectionGroupSnapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.PureProtectionGroupSnapshot"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.PureProtectionGroupSnapshot"]
**EradicationConfig** | Pointer to **string** | The configuration of eradication feature. | [optional] [readonly] 
**Pod** | Pointer to **string** | A pod representing a collection of protection groups and volumes is created on one array and stretched to another array, resulting in fully synchronized writes between the two arrays. | [optional] [readonly] 
**SnapshotSize** | Pointer to **int64** | The size of the snapshot created. | [optional] [readonly] 
**TotalProvisioned** | Pointer to **int64** | The overall size of the snapshot allocated by the storage array. | [optional] [readonly] 
**Array** | Pointer to [**NullableStoragePureArrayRelationship**](StoragePureArrayRelationship.md) |  | [optional] 
**ProtectionGroup** | Pointer to [**NullableStoragePureProtectionGroupRelationship**](StoragePureProtectionGroupRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewStoragePureProtectionGroupSnapshot

`func NewStoragePureProtectionGroupSnapshot(classId string, objectType string, ) *StoragePureProtectionGroupSnapshot`

NewStoragePureProtectionGroupSnapshot instantiates a new StoragePureProtectionGroupSnapshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoragePureProtectionGroupSnapshotWithDefaults

`func NewStoragePureProtectionGroupSnapshotWithDefaults() *StoragePureProtectionGroupSnapshot`

NewStoragePureProtectionGroupSnapshotWithDefaults instantiates a new StoragePureProtectionGroupSnapshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StoragePureProtectionGroupSnapshot) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StoragePureProtectionGroupSnapshot) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StoragePureProtectionGroupSnapshot) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StoragePureProtectionGroupSnapshot) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StoragePureProtectionGroupSnapshot) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StoragePureProtectionGroupSnapshot) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEradicationConfig

`func (o *StoragePureProtectionGroupSnapshot) GetEradicationConfig() string`

GetEradicationConfig returns the EradicationConfig field if non-nil, zero value otherwise.

### GetEradicationConfigOk

`func (o *StoragePureProtectionGroupSnapshot) GetEradicationConfigOk() (*string, bool)`

GetEradicationConfigOk returns a tuple with the EradicationConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEradicationConfig

`func (o *StoragePureProtectionGroupSnapshot) SetEradicationConfig(v string)`

SetEradicationConfig sets EradicationConfig field to given value.

### HasEradicationConfig

`func (o *StoragePureProtectionGroupSnapshot) HasEradicationConfig() bool`

HasEradicationConfig returns a boolean if a field has been set.

### GetPod

`func (o *StoragePureProtectionGroupSnapshot) GetPod() string`

GetPod returns the Pod field if non-nil, zero value otherwise.

### GetPodOk

`func (o *StoragePureProtectionGroupSnapshot) GetPodOk() (*string, bool)`

GetPodOk returns a tuple with the Pod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPod

`func (o *StoragePureProtectionGroupSnapshot) SetPod(v string)`

SetPod sets Pod field to given value.

### HasPod

`func (o *StoragePureProtectionGroupSnapshot) HasPod() bool`

HasPod returns a boolean if a field has been set.

### GetSnapshotSize

`func (o *StoragePureProtectionGroupSnapshot) GetSnapshotSize() int64`

GetSnapshotSize returns the SnapshotSize field if non-nil, zero value otherwise.

### GetSnapshotSizeOk

`func (o *StoragePureProtectionGroupSnapshot) GetSnapshotSizeOk() (*int64, bool)`

GetSnapshotSizeOk returns a tuple with the SnapshotSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotSize

`func (o *StoragePureProtectionGroupSnapshot) SetSnapshotSize(v int64)`

SetSnapshotSize sets SnapshotSize field to given value.

### HasSnapshotSize

`func (o *StoragePureProtectionGroupSnapshot) HasSnapshotSize() bool`

HasSnapshotSize returns a boolean if a field has been set.

### GetTotalProvisioned

`func (o *StoragePureProtectionGroupSnapshot) GetTotalProvisioned() int64`

GetTotalProvisioned returns the TotalProvisioned field if non-nil, zero value otherwise.

### GetTotalProvisionedOk

`func (o *StoragePureProtectionGroupSnapshot) GetTotalProvisionedOk() (*int64, bool)`

GetTotalProvisionedOk returns a tuple with the TotalProvisioned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalProvisioned

`func (o *StoragePureProtectionGroupSnapshot) SetTotalProvisioned(v int64)`

SetTotalProvisioned sets TotalProvisioned field to given value.

### HasTotalProvisioned

`func (o *StoragePureProtectionGroupSnapshot) HasTotalProvisioned() bool`

HasTotalProvisioned returns a boolean if a field has been set.

### GetArray

`func (o *StoragePureProtectionGroupSnapshot) GetArray() StoragePureArrayRelationship`

GetArray returns the Array field if non-nil, zero value otherwise.

### GetArrayOk

`func (o *StoragePureProtectionGroupSnapshot) GetArrayOk() (*StoragePureArrayRelationship, bool)`

GetArrayOk returns a tuple with the Array field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArray

`func (o *StoragePureProtectionGroupSnapshot) SetArray(v StoragePureArrayRelationship)`

SetArray sets Array field to given value.

### HasArray

`func (o *StoragePureProtectionGroupSnapshot) HasArray() bool`

HasArray returns a boolean if a field has been set.

### SetArrayNil

`func (o *StoragePureProtectionGroupSnapshot) SetArrayNil(b bool)`

 SetArrayNil sets the value for Array to be an explicit nil

### UnsetArray
`func (o *StoragePureProtectionGroupSnapshot) UnsetArray()`

UnsetArray ensures that no value is present for Array, not even an explicit nil
### GetProtectionGroup

`func (o *StoragePureProtectionGroupSnapshot) GetProtectionGroup() StoragePureProtectionGroupRelationship`

GetProtectionGroup returns the ProtectionGroup field if non-nil, zero value otherwise.

### GetProtectionGroupOk

`func (o *StoragePureProtectionGroupSnapshot) GetProtectionGroupOk() (*StoragePureProtectionGroupRelationship, bool)`

GetProtectionGroupOk returns a tuple with the ProtectionGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroup

`func (o *StoragePureProtectionGroupSnapshot) SetProtectionGroup(v StoragePureProtectionGroupRelationship)`

SetProtectionGroup sets ProtectionGroup field to given value.

### HasProtectionGroup

`func (o *StoragePureProtectionGroupSnapshot) HasProtectionGroup() bool`

HasProtectionGroup returns a boolean if a field has been set.

### SetProtectionGroupNil

`func (o *StoragePureProtectionGroupSnapshot) SetProtectionGroupNil(b bool)`

 SetProtectionGroupNil sets the value for ProtectionGroup to be an explicit nil

### UnsetProtectionGroup
`func (o *StoragePureProtectionGroupSnapshot) UnsetProtectionGroup()`

UnsetProtectionGroup ensures that no value is present for ProtectionGroup, not even an explicit nil
### GetRegisteredDevice

`func (o *StoragePureProtectionGroupSnapshot) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StoragePureProtectionGroupSnapshot) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StoragePureProtectionGroupSnapshot) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StoragePureProtectionGroupSnapshot) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *StoragePureProtectionGroupSnapshot) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *StoragePureProtectionGroupSnapshot) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



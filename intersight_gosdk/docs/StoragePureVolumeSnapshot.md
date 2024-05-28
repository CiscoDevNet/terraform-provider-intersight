# StoragePureVolumeSnapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.PureVolumeSnapshot"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.PureVolumeSnapshot"]
**Serial** | Pointer to **string** | Unique serial number of the snapshot allocated by the storage array. | [optional] [readonly] 
**Array** | Pointer to [**NullableStoragePureArrayRelationship**](StoragePureArrayRelationship.md) |  | [optional] 
**ProtectionGroupSnapshot** | Pointer to [**NullableStoragePureProtectionGroupSnapshotRelationship**](StoragePureProtectionGroupSnapshotRelationship.md) |  | [optional] 
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



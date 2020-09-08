# StoragePureVolumeSnapshotAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serial** | Pointer to **string** | Unique serial number of the snapshot allocated by the storage array. | [optional] [readonly] 
**Array** | Pointer to [**StoragePureArrayRelationship**](storage.PureArray.Relationship.md) |  | [optional] 
**ProtectionGroupSnapshot** | Pointer to [**StoragePureProtectionGroupSnapshotRelationship**](storage.PureProtectionGroupSnapshot.Relationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 
**Volume** | Pointer to [**StoragePureVolumeRelationship**](storage.PureVolume.Relationship.md) |  | [optional] 

## Methods

### NewStoragePureVolumeSnapshotAllOf

`func NewStoragePureVolumeSnapshotAllOf() *StoragePureVolumeSnapshotAllOf`

NewStoragePureVolumeSnapshotAllOf instantiates a new StoragePureVolumeSnapshotAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoragePureVolumeSnapshotAllOfWithDefaults

`func NewStoragePureVolumeSnapshotAllOfWithDefaults() *StoragePureVolumeSnapshotAllOf`

NewStoragePureVolumeSnapshotAllOfWithDefaults instantiates a new StoragePureVolumeSnapshotAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerial

`func (o *StoragePureVolumeSnapshotAllOf) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *StoragePureVolumeSnapshotAllOf) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *StoragePureVolumeSnapshotAllOf) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *StoragePureVolumeSnapshotAllOf) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetArray

`func (o *StoragePureVolumeSnapshotAllOf) GetArray() StoragePureArrayRelationship`

GetArray returns the Array field if non-nil, zero value otherwise.

### GetArrayOk

`func (o *StoragePureVolumeSnapshotAllOf) GetArrayOk() (*StoragePureArrayRelationship, bool)`

GetArrayOk returns a tuple with the Array field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArray

`func (o *StoragePureVolumeSnapshotAllOf) SetArray(v StoragePureArrayRelationship)`

SetArray sets Array field to given value.

### HasArray

`func (o *StoragePureVolumeSnapshotAllOf) HasArray() bool`

HasArray returns a boolean if a field has been set.

### GetProtectionGroupSnapshot

`func (o *StoragePureVolumeSnapshotAllOf) GetProtectionGroupSnapshot() StoragePureProtectionGroupSnapshotRelationship`

GetProtectionGroupSnapshot returns the ProtectionGroupSnapshot field if non-nil, zero value otherwise.

### GetProtectionGroupSnapshotOk

`func (o *StoragePureVolumeSnapshotAllOf) GetProtectionGroupSnapshotOk() (*StoragePureProtectionGroupSnapshotRelationship, bool)`

GetProtectionGroupSnapshotOk returns a tuple with the ProtectionGroupSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroupSnapshot

`func (o *StoragePureVolumeSnapshotAllOf) SetProtectionGroupSnapshot(v StoragePureProtectionGroupSnapshotRelationship)`

SetProtectionGroupSnapshot sets ProtectionGroupSnapshot field to given value.

### HasProtectionGroupSnapshot

`func (o *StoragePureVolumeSnapshotAllOf) HasProtectionGroupSnapshot() bool`

HasProtectionGroupSnapshot returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *StoragePureVolumeSnapshotAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StoragePureVolumeSnapshotAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StoragePureVolumeSnapshotAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StoragePureVolumeSnapshotAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### GetVolume

`func (o *StoragePureVolumeSnapshotAllOf) GetVolume() StoragePureVolumeRelationship`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *StoragePureVolumeSnapshotAllOf) GetVolumeOk() (*StoragePureVolumeRelationship, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *StoragePureVolumeSnapshotAllOf) SetVolume(v StoragePureVolumeRelationship)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *StoragePureVolumeSnapshotAllOf) HasVolume() bool`

HasVolume returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# StoragePureVolumeSnapshotAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.PureVolumeSnapshot"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.PureVolumeSnapshot"]
**Serial** | Pointer to **string** | Unique serial number of the snapshot allocated by the storage array. | [optional] [readonly] 
**Array** | Pointer to [**StoragePureArrayRelationship**](StoragePureArrayRelationship.md) |  | [optional] 
**ProtectionGroupSnapshot** | Pointer to [**StoragePureProtectionGroupSnapshotRelationship**](StoragePureProtectionGroupSnapshotRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 
**Volume** | Pointer to [**StoragePureVolumeRelationship**](StoragePureVolumeRelationship.md) |  | [optional] 

## Methods

### NewStoragePureVolumeSnapshotAllOf

`func NewStoragePureVolumeSnapshotAllOf(classId string, objectType string, ) *StoragePureVolumeSnapshotAllOf`

NewStoragePureVolumeSnapshotAllOf instantiates a new StoragePureVolumeSnapshotAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoragePureVolumeSnapshotAllOfWithDefaults

`func NewStoragePureVolumeSnapshotAllOfWithDefaults() *StoragePureVolumeSnapshotAllOf`

NewStoragePureVolumeSnapshotAllOfWithDefaults instantiates a new StoragePureVolumeSnapshotAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StoragePureVolumeSnapshotAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StoragePureVolumeSnapshotAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StoragePureVolumeSnapshotAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StoragePureVolumeSnapshotAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StoragePureVolumeSnapshotAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StoragePureVolumeSnapshotAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


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



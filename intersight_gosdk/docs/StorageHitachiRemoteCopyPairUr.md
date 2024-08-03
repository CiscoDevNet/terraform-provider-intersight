# StorageHitachiRemoteCopyPairUr

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.HitachiRemoteCopyPairUr"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.HitachiRemoteCopyPairUr"]
**MuNumber** | Pointer to **string** | MU (mirror unit) number of the volume. | [optional] [readonly] 
**Name** | Pointer to **string** | Object ID of the remote copy pair. | [optional] [readonly] 
**PvolLdevId** | Pointer to **int64** | LDEV number of primary volume. | [optional] [readonly] 
**PvolStorageSerial** | Pointer to **string** | Serial number of the storage system on the P-VOL. | [optional] [readonly] 
**ReplicationType** | Pointer to **string** | Pair type of the remote copy pair. | [optional] [readonly] 
**Status** | Pointer to **string** | Status of the remote copy pair. | [optional] [readonly] 
**SvolLdevId** | Pointer to **int64** | LDEV number of secondary volume. | [optional] [readonly] 
**SvolStorageSerial** | Pointer to **string** | Serial number of the storage system on the S-VOL. | [optional] [readonly] 
**Array** | Pointer to [**NullableStorageHitachiArrayRelationship**](StorageHitachiArrayRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewStorageHitachiRemoteCopyPairUr

`func NewStorageHitachiRemoteCopyPairUr(classId string, objectType string, ) *StorageHitachiRemoteCopyPairUr`

NewStorageHitachiRemoteCopyPairUr instantiates a new StorageHitachiRemoteCopyPairUr object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageHitachiRemoteCopyPairUrWithDefaults

`func NewStorageHitachiRemoteCopyPairUrWithDefaults() *StorageHitachiRemoteCopyPairUr`

NewStorageHitachiRemoteCopyPairUrWithDefaults instantiates a new StorageHitachiRemoteCopyPairUr object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageHitachiRemoteCopyPairUr) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageHitachiRemoteCopyPairUr) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageHitachiRemoteCopyPairUr) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageHitachiRemoteCopyPairUr) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageHitachiRemoteCopyPairUr) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageHitachiRemoteCopyPairUr) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetMuNumber

`func (o *StorageHitachiRemoteCopyPairUr) GetMuNumber() string`

GetMuNumber returns the MuNumber field if non-nil, zero value otherwise.

### GetMuNumberOk

`func (o *StorageHitachiRemoteCopyPairUr) GetMuNumberOk() (*string, bool)`

GetMuNumberOk returns a tuple with the MuNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuNumber

`func (o *StorageHitachiRemoteCopyPairUr) SetMuNumber(v string)`

SetMuNumber sets MuNumber field to given value.

### HasMuNumber

`func (o *StorageHitachiRemoteCopyPairUr) HasMuNumber() bool`

HasMuNumber returns a boolean if a field has been set.

### GetName

`func (o *StorageHitachiRemoteCopyPairUr) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageHitachiRemoteCopyPairUr) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageHitachiRemoteCopyPairUr) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StorageHitachiRemoteCopyPairUr) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPvolLdevId

`func (o *StorageHitachiRemoteCopyPairUr) GetPvolLdevId() int64`

GetPvolLdevId returns the PvolLdevId field if non-nil, zero value otherwise.

### GetPvolLdevIdOk

`func (o *StorageHitachiRemoteCopyPairUr) GetPvolLdevIdOk() (*int64, bool)`

GetPvolLdevIdOk returns a tuple with the PvolLdevId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPvolLdevId

`func (o *StorageHitachiRemoteCopyPairUr) SetPvolLdevId(v int64)`

SetPvolLdevId sets PvolLdevId field to given value.

### HasPvolLdevId

`func (o *StorageHitachiRemoteCopyPairUr) HasPvolLdevId() bool`

HasPvolLdevId returns a boolean if a field has been set.

### GetPvolStorageSerial

`func (o *StorageHitachiRemoteCopyPairUr) GetPvolStorageSerial() string`

GetPvolStorageSerial returns the PvolStorageSerial field if non-nil, zero value otherwise.

### GetPvolStorageSerialOk

`func (o *StorageHitachiRemoteCopyPairUr) GetPvolStorageSerialOk() (*string, bool)`

GetPvolStorageSerialOk returns a tuple with the PvolStorageSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPvolStorageSerial

`func (o *StorageHitachiRemoteCopyPairUr) SetPvolStorageSerial(v string)`

SetPvolStorageSerial sets PvolStorageSerial field to given value.

### HasPvolStorageSerial

`func (o *StorageHitachiRemoteCopyPairUr) HasPvolStorageSerial() bool`

HasPvolStorageSerial returns a boolean if a field has been set.

### GetReplicationType

`func (o *StorageHitachiRemoteCopyPairUr) GetReplicationType() string`

GetReplicationType returns the ReplicationType field if non-nil, zero value otherwise.

### GetReplicationTypeOk

`func (o *StorageHitachiRemoteCopyPairUr) GetReplicationTypeOk() (*string, bool)`

GetReplicationTypeOk returns a tuple with the ReplicationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationType

`func (o *StorageHitachiRemoteCopyPairUr) SetReplicationType(v string)`

SetReplicationType sets ReplicationType field to given value.

### HasReplicationType

`func (o *StorageHitachiRemoteCopyPairUr) HasReplicationType() bool`

HasReplicationType returns a boolean if a field has been set.

### GetStatus

`func (o *StorageHitachiRemoteCopyPairUr) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StorageHitachiRemoteCopyPairUr) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StorageHitachiRemoteCopyPairUr) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *StorageHitachiRemoteCopyPairUr) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSvolLdevId

`func (o *StorageHitachiRemoteCopyPairUr) GetSvolLdevId() int64`

GetSvolLdevId returns the SvolLdevId field if non-nil, zero value otherwise.

### GetSvolLdevIdOk

`func (o *StorageHitachiRemoteCopyPairUr) GetSvolLdevIdOk() (*int64, bool)`

GetSvolLdevIdOk returns a tuple with the SvolLdevId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSvolLdevId

`func (o *StorageHitachiRemoteCopyPairUr) SetSvolLdevId(v int64)`

SetSvolLdevId sets SvolLdevId field to given value.

### HasSvolLdevId

`func (o *StorageHitachiRemoteCopyPairUr) HasSvolLdevId() bool`

HasSvolLdevId returns a boolean if a field has been set.

### GetSvolStorageSerial

`func (o *StorageHitachiRemoteCopyPairUr) GetSvolStorageSerial() string`

GetSvolStorageSerial returns the SvolStorageSerial field if non-nil, zero value otherwise.

### GetSvolStorageSerialOk

`func (o *StorageHitachiRemoteCopyPairUr) GetSvolStorageSerialOk() (*string, bool)`

GetSvolStorageSerialOk returns a tuple with the SvolStorageSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSvolStorageSerial

`func (o *StorageHitachiRemoteCopyPairUr) SetSvolStorageSerial(v string)`

SetSvolStorageSerial sets SvolStorageSerial field to given value.

### HasSvolStorageSerial

`func (o *StorageHitachiRemoteCopyPairUr) HasSvolStorageSerial() bool`

HasSvolStorageSerial returns a boolean if a field has been set.

### GetArray

`func (o *StorageHitachiRemoteCopyPairUr) GetArray() StorageHitachiArrayRelationship`

GetArray returns the Array field if non-nil, zero value otherwise.

### GetArrayOk

`func (o *StorageHitachiRemoteCopyPairUr) GetArrayOk() (*StorageHitachiArrayRelationship, bool)`

GetArrayOk returns a tuple with the Array field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArray

`func (o *StorageHitachiRemoteCopyPairUr) SetArray(v StorageHitachiArrayRelationship)`

SetArray sets Array field to given value.

### HasArray

`func (o *StorageHitachiRemoteCopyPairUr) HasArray() bool`

HasArray returns a boolean if a field has been set.

### SetArrayNil

`func (o *StorageHitachiRemoteCopyPairUr) SetArrayNil(b bool)`

 SetArrayNil sets the value for Array to be an explicit nil

### UnsetArray
`func (o *StorageHitachiRemoteCopyPairUr) UnsetArray()`

UnsetArray ensures that no value is present for Array, not even an explicit nil
### GetRegisteredDevice

`func (o *StorageHitachiRemoteCopyPairUr) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StorageHitachiRemoteCopyPairUr) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StorageHitachiRemoteCopyPairUr) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StorageHitachiRemoteCopyPairUr) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *StorageHitachiRemoteCopyPairUr) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *StorageHitachiRemoteCopyPairUr) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# StorageHitachiParityGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.HitachiParityGroup"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.HitachiParityGroup"]
**DiskSpeed** | Pointer to **string** | Speed (rpm) of the disk belonging to the parity group. | [optional] [readonly] 
**DiskType** | Pointer to **string** | Type of the disk belonging to the parity group. | [optional] [readonly] 
**IsAcceleratedCompressionEnabled** | Pointer to **bool** | Value of the accelerated compression of the parity group. true, Accelerated compression for the parity group is enabled. false, Accelerated compression for the parity group is disabled. | [optional] [readonly] 
**IsCopyBackModeEnabled** | Pointer to **bool** | Value of the copy back mode setting of the parity group. true, Copy back mode is enabled. false, Copy back mode is disabled. | [optional] [readonly] 
**IsEncryptionEnabled** | Pointer to **bool** | Value of the encryption setting of the parity group. true, Encryption is enabled. false, Encryption is disabled. | [optional] [readonly] 
**Array** | Pointer to [**StorageHitachiArrayRelationship**](StorageHitachiArrayRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewStorageHitachiParityGroup

`func NewStorageHitachiParityGroup(classId string, objectType string, ) *StorageHitachiParityGroup`

NewStorageHitachiParityGroup instantiates a new StorageHitachiParityGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageHitachiParityGroupWithDefaults

`func NewStorageHitachiParityGroupWithDefaults() *StorageHitachiParityGroup`

NewStorageHitachiParityGroupWithDefaults instantiates a new StorageHitachiParityGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageHitachiParityGroup) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageHitachiParityGroup) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageHitachiParityGroup) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageHitachiParityGroup) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageHitachiParityGroup) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageHitachiParityGroup) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDiskSpeed

`func (o *StorageHitachiParityGroup) GetDiskSpeed() string`

GetDiskSpeed returns the DiskSpeed field if non-nil, zero value otherwise.

### GetDiskSpeedOk

`func (o *StorageHitachiParityGroup) GetDiskSpeedOk() (*string, bool)`

GetDiskSpeedOk returns a tuple with the DiskSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSpeed

`func (o *StorageHitachiParityGroup) SetDiskSpeed(v string)`

SetDiskSpeed sets DiskSpeed field to given value.

### HasDiskSpeed

`func (o *StorageHitachiParityGroup) HasDiskSpeed() bool`

HasDiskSpeed returns a boolean if a field has been set.

### GetDiskType

`func (o *StorageHitachiParityGroup) GetDiskType() string`

GetDiskType returns the DiskType field if non-nil, zero value otherwise.

### GetDiskTypeOk

`func (o *StorageHitachiParityGroup) GetDiskTypeOk() (*string, bool)`

GetDiskTypeOk returns a tuple with the DiskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskType

`func (o *StorageHitachiParityGroup) SetDiskType(v string)`

SetDiskType sets DiskType field to given value.

### HasDiskType

`func (o *StorageHitachiParityGroup) HasDiskType() bool`

HasDiskType returns a boolean if a field has been set.

### GetIsAcceleratedCompressionEnabled

`func (o *StorageHitachiParityGroup) GetIsAcceleratedCompressionEnabled() bool`

GetIsAcceleratedCompressionEnabled returns the IsAcceleratedCompressionEnabled field if non-nil, zero value otherwise.

### GetIsAcceleratedCompressionEnabledOk

`func (o *StorageHitachiParityGroup) GetIsAcceleratedCompressionEnabledOk() (*bool, bool)`

GetIsAcceleratedCompressionEnabledOk returns a tuple with the IsAcceleratedCompressionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAcceleratedCompressionEnabled

`func (o *StorageHitachiParityGroup) SetIsAcceleratedCompressionEnabled(v bool)`

SetIsAcceleratedCompressionEnabled sets IsAcceleratedCompressionEnabled field to given value.

### HasIsAcceleratedCompressionEnabled

`func (o *StorageHitachiParityGroup) HasIsAcceleratedCompressionEnabled() bool`

HasIsAcceleratedCompressionEnabled returns a boolean if a field has been set.

### GetIsCopyBackModeEnabled

`func (o *StorageHitachiParityGroup) GetIsCopyBackModeEnabled() bool`

GetIsCopyBackModeEnabled returns the IsCopyBackModeEnabled field if non-nil, zero value otherwise.

### GetIsCopyBackModeEnabledOk

`func (o *StorageHitachiParityGroup) GetIsCopyBackModeEnabledOk() (*bool, bool)`

GetIsCopyBackModeEnabledOk returns a tuple with the IsCopyBackModeEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCopyBackModeEnabled

`func (o *StorageHitachiParityGroup) SetIsCopyBackModeEnabled(v bool)`

SetIsCopyBackModeEnabled sets IsCopyBackModeEnabled field to given value.

### HasIsCopyBackModeEnabled

`func (o *StorageHitachiParityGroup) HasIsCopyBackModeEnabled() bool`

HasIsCopyBackModeEnabled returns a boolean if a field has been set.

### GetIsEncryptionEnabled

`func (o *StorageHitachiParityGroup) GetIsEncryptionEnabled() bool`

GetIsEncryptionEnabled returns the IsEncryptionEnabled field if non-nil, zero value otherwise.

### GetIsEncryptionEnabledOk

`func (o *StorageHitachiParityGroup) GetIsEncryptionEnabledOk() (*bool, bool)`

GetIsEncryptionEnabledOk returns a tuple with the IsEncryptionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEncryptionEnabled

`func (o *StorageHitachiParityGroup) SetIsEncryptionEnabled(v bool)`

SetIsEncryptionEnabled sets IsEncryptionEnabled field to given value.

### HasIsEncryptionEnabled

`func (o *StorageHitachiParityGroup) HasIsEncryptionEnabled() bool`

HasIsEncryptionEnabled returns a boolean if a field has been set.

### GetArray

`func (o *StorageHitachiParityGroup) GetArray() StorageHitachiArrayRelationship`

GetArray returns the Array field if non-nil, zero value otherwise.

### GetArrayOk

`func (o *StorageHitachiParityGroup) GetArrayOk() (*StorageHitachiArrayRelationship, bool)`

GetArrayOk returns a tuple with the Array field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArray

`func (o *StorageHitachiParityGroup) SetArray(v StorageHitachiArrayRelationship)`

SetArray sets Array field to given value.

### HasArray

`func (o *StorageHitachiParityGroup) HasArray() bool`

HasArray returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *StorageHitachiParityGroup) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StorageHitachiParityGroup) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StorageHitachiParityGroup) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StorageHitachiParityGroup) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



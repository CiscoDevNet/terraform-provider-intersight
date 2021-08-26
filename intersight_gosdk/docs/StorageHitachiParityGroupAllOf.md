# StorageHitachiParityGroupAllOf

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

### NewStorageHitachiParityGroupAllOf

`func NewStorageHitachiParityGroupAllOf(classId string, objectType string, ) *StorageHitachiParityGroupAllOf`

NewStorageHitachiParityGroupAllOf instantiates a new StorageHitachiParityGroupAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageHitachiParityGroupAllOfWithDefaults

`func NewStorageHitachiParityGroupAllOfWithDefaults() *StorageHitachiParityGroupAllOf`

NewStorageHitachiParityGroupAllOfWithDefaults instantiates a new StorageHitachiParityGroupAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageHitachiParityGroupAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageHitachiParityGroupAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageHitachiParityGroupAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageHitachiParityGroupAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageHitachiParityGroupAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageHitachiParityGroupAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDiskSpeed

`func (o *StorageHitachiParityGroupAllOf) GetDiskSpeed() string`

GetDiskSpeed returns the DiskSpeed field if non-nil, zero value otherwise.

### GetDiskSpeedOk

`func (o *StorageHitachiParityGroupAllOf) GetDiskSpeedOk() (*string, bool)`

GetDiskSpeedOk returns a tuple with the DiskSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSpeed

`func (o *StorageHitachiParityGroupAllOf) SetDiskSpeed(v string)`

SetDiskSpeed sets DiskSpeed field to given value.

### HasDiskSpeed

`func (o *StorageHitachiParityGroupAllOf) HasDiskSpeed() bool`

HasDiskSpeed returns a boolean if a field has been set.

### GetDiskType

`func (o *StorageHitachiParityGroupAllOf) GetDiskType() string`

GetDiskType returns the DiskType field if non-nil, zero value otherwise.

### GetDiskTypeOk

`func (o *StorageHitachiParityGroupAllOf) GetDiskTypeOk() (*string, bool)`

GetDiskTypeOk returns a tuple with the DiskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskType

`func (o *StorageHitachiParityGroupAllOf) SetDiskType(v string)`

SetDiskType sets DiskType field to given value.

### HasDiskType

`func (o *StorageHitachiParityGroupAllOf) HasDiskType() bool`

HasDiskType returns a boolean if a field has been set.

### GetIsAcceleratedCompressionEnabled

`func (o *StorageHitachiParityGroupAllOf) GetIsAcceleratedCompressionEnabled() bool`

GetIsAcceleratedCompressionEnabled returns the IsAcceleratedCompressionEnabled field if non-nil, zero value otherwise.

### GetIsAcceleratedCompressionEnabledOk

`func (o *StorageHitachiParityGroupAllOf) GetIsAcceleratedCompressionEnabledOk() (*bool, bool)`

GetIsAcceleratedCompressionEnabledOk returns a tuple with the IsAcceleratedCompressionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAcceleratedCompressionEnabled

`func (o *StorageHitachiParityGroupAllOf) SetIsAcceleratedCompressionEnabled(v bool)`

SetIsAcceleratedCompressionEnabled sets IsAcceleratedCompressionEnabled field to given value.

### HasIsAcceleratedCompressionEnabled

`func (o *StorageHitachiParityGroupAllOf) HasIsAcceleratedCompressionEnabled() bool`

HasIsAcceleratedCompressionEnabled returns a boolean if a field has been set.

### GetIsCopyBackModeEnabled

`func (o *StorageHitachiParityGroupAllOf) GetIsCopyBackModeEnabled() bool`

GetIsCopyBackModeEnabled returns the IsCopyBackModeEnabled field if non-nil, zero value otherwise.

### GetIsCopyBackModeEnabledOk

`func (o *StorageHitachiParityGroupAllOf) GetIsCopyBackModeEnabledOk() (*bool, bool)`

GetIsCopyBackModeEnabledOk returns a tuple with the IsCopyBackModeEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCopyBackModeEnabled

`func (o *StorageHitachiParityGroupAllOf) SetIsCopyBackModeEnabled(v bool)`

SetIsCopyBackModeEnabled sets IsCopyBackModeEnabled field to given value.

### HasIsCopyBackModeEnabled

`func (o *StorageHitachiParityGroupAllOf) HasIsCopyBackModeEnabled() bool`

HasIsCopyBackModeEnabled returns a boolean if a field has been set.

### GetIsEncryptionEnabled

`func (o *StorageHitachiParityGroupAllOf) GetIsEncryptionEnabled() bool`

GetIsEncryptionEnabled returns the IsEncryptionEnabled field if non-nil, zero value otherwise.

### GetIsEncryptionEnabledOk

`func (o *StorageHitachiParityGroupAllOf) GetIsEncryptionEnabledOk() (*bool, bool)`

GetIsEncryptionEnabledOk returns a tuple with the IsEncryptionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEncryptionEnabled

`func (o *StorageHitachiParityGroupAllOf) SetIsEncryptionEnabled(v bool)`

SetIsEncryptionEnabled sets IsEncryptionEnabled field to given value.

### HasIsEncryptionEnabled

`func (o *StorageHitachiParityGroupAllOf) HasIsEncryptionEnabled() bool`

HasIsEncryptionEnabled returns a boolean if a field has been set.

### GetArray

`func (o *StorageHitachiParityGroupAllOf) GetArray() StorageHitachiArrayRelationship`

GetArray returns the Array field if non-nil, zero value otherwise.

### GetArrayOk

`func (o *StorageHitachiParityGroupAllOf) GetArrayOk() (*StorageHitachiArrayRelationship, bool)`

GetArrayOk returns a tuple with the Array field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArray

`func (o *StorageHitachiParityGroupAllOf) SetArray(v StorageHitachiArrayRelationship)`

SetArray sets Array field to given value.

### HasArray

`func (o *StorageHitachiParityGroupAllOf) HasArray() bool`

HasArray returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *StorageHitachiParityGroupAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StorageHitachiParityGroupAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StorageHitachiParityGroupAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StorageHitachiParityGroupAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



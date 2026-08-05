# StoragePureFlashBladeFileSystem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.PureFlashBladeFileSystem"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.PureFlashBladeFileSystem"]
**CreatedTime** | Pointer to **time.Time** | Creation timestamp of the file system. | [optional] [readonly] 
**Destroyed** | Pointer to **bool** | Returns a value of true if the file system has been destroyed and is pending eradication. The file system cannot be modified while it is in the destroyed state. | [optional] [readonly] 
**FileSystemId** | Pointer to **string** | A non-modifiable, globally unique ID chosen by the system. | [optional] [readonly] 
**HardLimitEnabled** | Pointer to **bool** | If set to true, the file system&#39;s size, as defined by provisioned, is used as a hard limit quota. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the object (e.g., a file system or snapshot). | [optional] [readonly] 
**PromotionStatus** | Pointer to **string** | The current status of the file system with respect to replication. Possible values are promoted and demoted. | [optional] [readonly] 
**Provisioned** | Pointer to **int64** | The provisioned size of the file system, displayed in bytes. If set to an empty string, the file system is unlimited in size. | [optional] [readonly] 
**StorageUtilization** | Pointer to [**StoragePureFlashBladeUtilization**](StoragePureFlashBladeUtilization.md) |  | [optional] 
**Writable** | Pointer to **bool** | Whether the file system is writable or not. If false, this overrides any protocol or file permission settings and prevents changes. | [optional] [readonly] 
**Array** | Pointer to [**NullableStoragePureFlashBladeSystemRelationship**](StoragePureFlashBladeSystemRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewStoragePureFlashBladeFileSystem

`func NewStoragePureFlashBladeFileSystem(classId string, objectType string, ) *StoragePureFlashBladeFileSystem`

NewStoragePureFlashBladeFileSystem instantiates a new StoragePureFlashBladeFileSystem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoragePureFlashBladeFileSystemWithDefaults

`func NewStoragePureFlashBladeFileSystemWithDefaults() *StoragePureFlashBladeFileSystem`

NewStoragePureFlashBladeFileSystemWithDefaults instantiates a new StoragePureFlashBladeFileSystem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StoragePureFlashBladeFileSystem) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StoragePureFlashBladeFileSystem) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StoragePureFlashBladeFileSystem) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StoragePureFlashBladeFileSystem) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StoragePureFlashBladeFileSystem) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StoragePureFlashBladeFileSystem) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCreatedTime

`func (o *StoragePureFlashBladeFileSystem) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *StoragePureFlashBladeFileSystem) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *StoragePureFlashBladeFileSystem) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *StoragePureFlashBladeFileSystem) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetDestroyed

`func (o *StoragePureFlashBladeFileSystem) GetDestroyed() bool`

GetDestroyed returns the Destroyed field if non-nil, zero value otherwise.

### GetDestroyedOk

`func (o *StoragePureFlashBladeFileSystem) GetDestroyedOk() (*bool, bool)`

GetDestroyedOk returns a tuple with the Destroyed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestroyed

`func (o *StoragePureFlashBladeFileSystem) SetDestroyed(v bool)`

SetDestroyed sets Destroyed field to given value.

### HasDestroyed

`func (o *StoragePureFlashBladeFileSystem) HasDestroyed() bool`

HasDestroyed returns a boolean if a field has been set.

### GetFileSystemId

`func (o *StoragePureFlashBladeFileSystem) GetFileSystemId() string`

GetFileSystemId returns the FileSystemId field if non-nil, zero value otherwise.

### GetFileSystemIdOk

`func (o *StoragePureFlashBladeFileSystem) GetFileSystemIdOk() (*string, bool)`

GetFileSystemIdOk returns a tuple with the FileSystemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSystemId

`func (o *StoragePureFlashBladeFileSystem) SetFileSystemId(v string)`

SetFileSystemId sets FileSystemId field to given value.

### HasFileSystemId

`func (o *StoragePureFlashBladeFileSystem) HasFileSystemId() bool`

HasFileSystemId returns a boolean if a field has been set.

### GetHardLimitEnabled

`func (o *StoragePureFlashBladeFileSystem) GetHardLimitEnabled() bool`

GetHardLimitEnabled returns the HardLimitEnabled field if non-nil, zero value otherwise.

### GetHardLimitEnabledOk

`func (o *StoragePureFlashBladeFileSystem) GetHardLimitEnabledOk() (*bool, bool)`

GetHardLimitEnabledOk returns a tuple with the HardLimitEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardLimitEnabled

`func (o *StoragePureFlashBladeFileSystem) SetHardLimitEnabled(v bool)`

SetHardLimitEnabled sets HardLimitEnabled field to given value.

### HasHardLimitEnabled

`func (o *StoragePureFlashBladeFileSystem) HasHardLimitEnabled() bool`

HasHardLimitEnabled returns a boolean if a field has been set.

### GetName

`func (o *StoragePureFlashBladeFileSystem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StoragePureFlashBladeFileSystem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StoragePureFlashBladeFileSystem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StoragePureFlashBladeFileSystem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPromotionStatus

`func (o *StoragePureFlashBladeFileSystem) GetPromotionStatus() string`

GetPromotionStatus returns the PromotionStatus field if non-nil, zero value otherwise.

### GetPromotionStatusOk

`func (o *StoragePureFlashBladeFileSystem) GetPromotionStatusOk() (*string, bool)`

GetPromotionStatusOk returns a tuple with the PromotionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromotionStatus

`func (o *StoragePureFlashBladeFileSystem) SetPromotionStatus(v string)`

SetPromotionStatus sets PromotionStatus field to given value.

### HasPromotionStatus

`func (o *StoragePureFlashBladeFileSystem) HasPromotionStatus() bool`

HasPromotionStatus returns a boolean if a field has been set.

### GetProvisioned

`func (o *StoragePureFlashBladeFileSystem) GetProvisioned() int64`

GetProvisioned returns the Provisioned field if non-nil, zero value otherwise.

### GetProvisionedOk

`func (o *StoragePureFlashBladeFileSystem) GetProvisionedOk() (*int64, bool)`

GetProvisionedOk returns a tuple with the Provisioned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioned

`func (o *StoragePureFlashBladeFileSystem) SetProvisioned(v int64)`

SetProvisioned sets Provisioned field to given value.

### HasProvisioned

`func (o *StoragePureFlashBladeFileSystem) HasProvisioned() bool`

HasProvisioned returns a boolean if a field has been set.

### GetStorageUtilization

`func (o *StoragePureFlashBladeFileSystem) GetStorageUtilization() StoragePureFlashBladeUtilization`

GetStorageUtilization returns the StorageUtilization field if non-nil, zero value otherwise.

### GetStorageUtilizationOk

`func (o *StoragePureFlashBladeFileSystem) GetStorageUtilizationOk() (*StoragePureFlashBladeUtilization, bool)`

GetStorageUtilizationOk returns a tuple with the StorageUtilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageUtilization

`func (o *StoragePureFlashBladeFileSystem) SetStorageUtilization(v StoragePureFlashBladeUtilization)`

SetStorageUtilization sets StorageUtilization field to given value.

### HasStorageUtilization

`func (o *StoragePureFlashBladeFileSystem) HasStorageUtilization() bool`

HasStorageUtilization returns a boolean if a field has been set.

### GetWritable

`func (o *StoragePureFlashBladeFileSystem) GetWritable() bool`

GetWritable returns the Writable field if non-nil, zero value otherwise.

### GetWritableOk

`func (o *StoragePureFlashBladeFileSystem) GetWritableOk() (*bool, bool)`

GetWritableOk returns a tuple with the Writable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWritable

`func (o *StoragePureFlashBladeFileSystem) SetWritable(v bool)`

SetWritable sets Writable field to given value.

### HasWritable

`func (o *StoragePureFlashBladeFileSystem) HasWritable() bool`

HasWritable returns a boolean if a field has been set.

### GetArray

`func (o *StoragePureFlashBladeFileSystem) GetArray() StoragePureFlashBladeSystemRelationship`

GetArray returns the Array field if non-nil, zero value otherwise.

### GetArrayOk

`func (o *StoragePureFlashBladeFileSystem) GetArrayOk() (*StoragePureFlashBladeSystemRelationship, bool)`

GetArrayOk returns a tuple with the Array field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArray

`func (o *StoragePureFlashBladeFileSystem) SetArray(v StoragePureFlashBladeSystemRelationship)`

SetArray sets Array field to given value.

### HasArray

`func (o *StoragePureFlashBladeFileSystem) HasArray() bool`

HasArray returns a boolean if a field has been set.

### SetArrayNil

`func (o *StoragePureFlashBladeFileSystem) SetArrayNil(b bool)`

 SetArrayNil sets the value for Array to be an explicit nil

### UnsetArray
`func (o *StoragePureFlashBladeFileSystem) UnsetArray()`

UnsetArray ensures that no value is present for Array, not even an explicit nil
### GetRegisteredDevice

`func (o *StoragePureFlashBladeFileSystem) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StoragePureFlashBladeFileSystem) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StoragePureFlashBladeFileSystem) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StoragePureFlashBladeFileSystem) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *StoragePureFlashBladeFileSystem) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *StoragePureFlashBladeFileSystem) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



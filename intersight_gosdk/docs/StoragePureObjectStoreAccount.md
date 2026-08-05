# StoragePureObjectStoreAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.PureObjectStoreAccount"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.PureObjectStoreAccount"]
**Created** | Pointer to **time.Time** | Creation timestamp of the account. | [optional] [readonly] 
**HardLimitEnabled** | Pointer to **bool** | If true, the account size defined by quotaLimit is used as a hard limit quota. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the object store account. | [optional] [readonly] 
**ObjectCount** | Pointer to **int64** | The count of objects within the account. | [optional] [readonly] 
**QuotaLimit** | Pointer to **int64** | The effective quota limit applied against the size of the account, displayed in bytes. | [optional] [readonly] 
**StorageUtilization** | Pointer to [**StoragePureFlashBladeUtilization**](StoragePureFlashBladeUtilization.md) |  | [optional] 
**Uuid** | Pointer to **string** | A non-modifiable, globally unique ID chosen by the system. | [optional] [readonly] 
**Array** | Pointer to [**NullableStoragePureFlashBladeSystemRelationship**](StoragePureFlashBladeSystemRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewStoragePureObjectStoreAccount

`func NewStoragePureObjectStoreAccount(classId string, objectType string, ) *StoragePureObjectStoreAccount`

NewStoragePureObjectStoreAccount instantiates a new StoragePureObjectStoreAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoragePureObjectStoreAccountWithDefaults

`func NewStoragePureObjectStoreAccountWithDefaults() *StoragePureObjectStoreAccount`

NewStoragePureObjectStoreAccountWithDefaults instantiates a new StoragePureObjectStoreAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StoragePureObjectStoreAccount) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StoragePureObjectStoreAccount) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StoragePureObjectStoreAccount) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StoragePureObjectStoreAccount) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StoragePureObjectStoreAccount) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StoragePureObjectStoreAccount) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCreated

`func (o *StoragePureObjectStoreAccount) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *StoragePureObjectStoreAccount) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *StoragePureObjectStoreAccount) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *StoragePureObjectStoreAccount) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetHardLimitEnabled

`func (o *StoragePureObjectStoreAccount) GetHardLimitEnabled() bool`

GetHardLimitEnabled returns the HardLimitEnabled field if non-nil, zero value otherwise.

### GetHardLimitEnabledOk

`func (o *StoragePureObjectStoreAccount) GetHardLimitEnabledOk() (*bool, bool)`

GetHardLimitEnabledOk returns a tuple with the HardLimitEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardLimitEnabled

`func (o *StoragePureObjectStoreAccount) SetHardLimitEnabled(v bool)`

SetHardLimitEnabled sets HardLimitEnabled field to given value.

### HasHardLimitEnabled

`func (o *StoragePureObjectStoreAccount) HasHardLimitEnabled() bool`

HasHardLimitEnabled returns a boolean if a field has been set.

### GetName

`func (o *StoragePureObjectStoreAccount) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StoragePureObjectStoreAccount) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StoragePureObjectStoreAccount) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StoragePureObjectStoreAccount) HasName() bool`

HasName returns a boolean if a field has been set.

### GetObjectCount

`func (o *StoragePureObjectStoreAccount) GetObjectCount() int64`

GetObjectCount returns the ObjectCount field if non-nil, zero value otherwise.

### GetObjectCountOk

`func (o *StoragePureObjectStoreAccount) GetObjectCountOk() (*int64, bool)`

GetObjectCountOk returns a tuple with the ObjectCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectCount

`func (o *StoragePureObjectStoreAccount) SetObjectCount(v int64)`

SetObjectCount sets ObjectCount field to given value.

### HasObjectCount

`func (o *StoragePureObjectStoreAccount) HasObjectCount() bool`

HasObjectCount returns a boolean if a field has been set.

### GetQuotaLimit

`func (o *StoragePureObjectStoreAccount) GetQuotaLimit() int64`

GetQuotaLimit returns the QuotaLimit field if non-nil, zero value otherwise.

### GetQuotaLimitOk

`func (o *StoragePureObjectStoreAccount) GetQuotaLimitOk() (*int64, bool)`

GetQuotaLimitOk returns a tuple with the QuotaLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaLimit

`func (o *StoragePureObjectStoreAccount) SetQuotaLimit(v int64)`

SetQuotaLimit sets QuotaLimit field to given value.

### HasQuotaLimit

`func (o *StoragePureObjectStoreAccount) HasQuotaLimit() bool`

HasQuotaLimit returns a boolean if a field has been set.

### GetStorageUtilization

`func (o *StoragePureObjectStoreAccount) GetStorageUtilization() StoragePureFlashBladeUtilization`

GetStorageUtilization returns the StorageUtilization field if non-nil, zero value otherwise.

### GetStorageUtilizationOk

`func (o *StoragePureObjectStoreAccount) GetStorageUtilizationOk() (*StoragePureFlashBladeUtilization, bool)`

GetStorageUtilizationOk returns a tuple with the StorageUtilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageUtilization

`func (o *StoragePureObjectStoreAccount) SetStorageUtilization(v StoragePureFlashBladeUtilization)`

SetStorageUtilization sets StorageUtilization field to given value.

### HasStorageUtilization

`func (o *StoragePureObjectStoreAccount) HasStorageUtilization() bool`

HasStorageUtilization returns a boolean if a field has been set.

### GetUuid

`func (o *StoragePureObjectStoreAccount) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *StoragePureObjectStoreAccount) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *StoragePureObjectStoreAccount) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *StoragePureObjectStoreAccount) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetArray

`func (o *StoragePureObjectStoreAccount) GetArray() StoragePureFlashBladeSystemRelationship`

GetArray returns the Array field if non-nil, zero value otherwise.

### GetArrayOk

`func (o *StoragePureObjectStoreAccount) GetArrayOk() (*StoragePureFlashBladeSystemRelationship, bool)`

GetArrayOk returns a tuple with the Array field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArray

`func (o *StoragePureObjectStoreAccount) SetArray(v StoragePureFlashBladeSystemRelationship)`

SetArray sets Array field to given value.

### HasArray

`func (o *StoragePureObjectStoreAccount) HasArray() bool`

HasArray returns a boolean if a field has been set.

### SetArrayNil

`func (o *StoragePureObjectStoreAccount) SetArrayNil(b bool)`

 SetArrayNil sets the value for Array to be an explicit nil

### UnsetArray
`func (o *StoragePureObjectStoreAccount) UnsetArray()`

UnsetArray ensures that no value is present for Array, not even an explicit nil
### GetRegisteredDevice

`func (o *StoragePureObjectStoreAccount) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StoragePureObjectStoreAccount) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StoragePureObjectStoreAccount) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StoragePureObjectStoreAccount) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *StoragePureObjectStoreAccount) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *StoragePureObjectStoreAccount) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# StoragePureObjectBucket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.PureObjectBucket"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.PureObjectBucket"]
**BucketName** | Pointer to **string** | Name of the bucket within the object store account. | [optional] [readonly] 
**BucketType** | Pointer to **string** | Type of the bucket (e.g., multi-site-writable). | [optional] [readonly] 
**Created** | Pointer to **time.Time** | Creation time of the bucket. | [optional] [readonly] 
**Destroyed** | Pointer to **bool** | Whether the bucket has been destroyed. | [optional] [readonly] 
**HardLimitEnabled** | Pointer to **bool** | Whether the hard limit is enabled for the bucket. | [optional] [readonly] 
**Name** | Pointer to **string** | Unique name of the bucket (combination of account name and bucket name). | [optional] [readonly] 
**ObjectCount** | Pointer to **int64** | Number of objects in the bucket. | [optional] [readonly] 
**ObjectStoreAccountName** | Pointer to **string** | Name of the object store account that owns this bucket. | [optional] [readonly] 
**QuotaLimit** | Pointer to **int64** | Quota limit for the bucket in bytes. | [optional] [readonly] 
**RetentionLock** | Pointer to **string** | Retention lock status for the bucket. | [optional] [readonly] 
**StorageUtilization** | Pointer to [**StoragePureFlashBladeUtilization**](StoragePureFlashBladeUtilization.md) |  | [optional] 
**Uuid** | Pointer to **string** | A globally unique identifier for the bucket. | [optional] [readonly] 
**Vendor** | Pointer to **string** | Vendor of the storage array. | [optional] [readonly] 
**Version** | Pointer to **string** | Version of the FlashBlade array. | [optional] [readonly] 
**Versioning** | Pointer to **string** | Versioning mode for the bucket. | [optional] [readonly] 
**Array** | Pointer to [**NullableStoragePureFlashBladeSystemRelationship**](StoragePureFlashBladeSystemRelationship.md) |  | [optional] 
**ObjectStoreAccount** | Pointer to [**NullableStoragePureObjectStoreAccountRelationship**](StoragePureObjectStoreAccountRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewStoragePureObjectBucket

`func NewStoragePureObjectBucket(classId string, objectType string, ) *StoragePureObjectBucket`

NewStoragePureObjectBucket instantiates a new StoragePureObjectBucket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoragePureObjectBucketWithDefaults

`func NewStoragePureObjectBucketWithDefaults() *StoragePureObjectBucket`

NewStoragePureObjectBucketWithDefaults instantiates a new StoragePureObjectBucket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StoragePureObjectBucket) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StoragePureObjectBucket) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StoragePureObjectBucket) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StoragePureObjectBucket) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StoragePureObjectBucket) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StoragePureObjectBucket) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBucketName

`func (o *StoragePureObjectBucket) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *StoragePureObjectBucket) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *StoragePureObjectBucket) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.

### HasBucketName

`func (o *StoragePureObjectBucket) HasBucketName() bool`

HasBucketName returns a boolean if a field has been set.

### GetBucketType

`func (o *StoragePureObjectBucket) GetBucketType() string`

GetBucketType returns the BucketType field if non-nil, zero value otherwise.

### GetBucketTypeOk

`func (o *StoragePureObjectBucket) GetBucketTypeOk() (*string, bool)`

GetBucketTypeOk returns a tuple with the BucketType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketType

`func (o *StoragePureObjectBucket) SetBucketType(v string)`

SetBucketType sets BucketType field to given value.

### HasBucketType

`func (o *StoragePureObjectBucket) HasBucketType() bool`

HasBucketType returns a boolean if a field has been set.

### GetCreated

`func (o *StoragePureObjectBucket) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *StoragePureObjectBucket) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *StoragePureObjectBucket) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *StoragePureObjectBucket) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDestroyed

`func (o *StoragePureObjectBucket) GetDestroyed() bool`

GetDestroyed returns the Destroyed field if non-nil, zero value otherwise.

### GetDestroyedOk

`func (o *StoragePureObjectBucket) GetDestroyedOk() (*bool, bool)`

GetDestroyedOk returns a tuple with the Destroyed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestroyed

`func (o *StoragePureObjectBucket) SetDestroyed(v bool)`

SetDestroyed sets Destroyed field to given value.

### HasDestroyed

`func (o *StoragePureObjectBucket) HasDestroyed() bool`

HasDestroyed returns a boolean if a field has been set.

### GetHardLimitEnabled

`func (o *StoragePureObjectBucket) GetHardLimitEnabled() bool`

GetHardLimitEnabled returns the HardLimitEnabled field if non-nil, zero value otherwise.

### GetHardLimitEnabledOk

`func (o *StoragePureObjectBucket) GetHardLimitEnabledOk() (*bool, bool)`

GetHardLimitEnabledOk returns a tuple with the HardLimitEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardLimitEnabled

`func (o *StoragePureObjectBucket) SetHardLimitEnabled(v bool)`

SetHardLimitEnabled sets HardLimitEnabled field to given value.

### HasHardLimitEnabled

`func (o *StoragePureObjectBucket) HasHardLimitEnabled() bool`

HasHardLimitEnabled returns a boolean if a field has been set.

### GetName

`func (o *StoragePureObjectBucket) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StoragePureObjectBucket) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StoragePureObjectBucket) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StoragePureObjectBucket) HasName() bool`

HasName returns a boolean if a field has been set.

### GetObjectCount

`func (o *StoragePureObjectBucket) GetObjectCount() int64`

GetObjectCount returns the ObjectCount field if non-nil, zero value otherwise.

### GetObjectCountOk

`func (o *StoragePureObjectBucket) GetObjectCountOk() (*int64, bool)`

GetObjectCountOk returns a tuple with the ObjectCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectCount

`func (o *StoragePureObjectBucket) SetObjectCount(v int64)`

SetObjectCount sets ObjectCount field to given value.

### HasObjectCount

`func (o *StoragePureObjectBucket) HasObjectCount() bool`

HasObjectCount returns a boolean if a field has been set.

### GetObjectStoreAccountName

`func (o *StoragePureObjectBucket) GetObjectStoreAccountName() string`

GetObjectStoreAccountName returns the ObjectStoreAccountName field if non-nil, zero value otherwise.

### GetObjectStoreAccountNameOk

`func (o *StoragePureObjectBucket) GetObjectStoreAccountNameOk() (*string, bool)`

GetObjectStoreAccountNameOk returns a tuple with the ObjectStoreAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectStoreAccountName

`func (o *StoragePureObjectBucket) SetObjectStoreAccountName(v string)`

SetObjectStoreAccountName sets ObjectStoreAccountName field to given value.

### HasObjectStoreAccountName

`func (o *StoragePureObjectBucket) HasObjectStoreAccountName() bool`

HasObjectStoreAccountName returns a boolean if a field has been set.

### GetQuotaLimit

`func (o *StoragePureObjectBucket) GetQuotaLimit() int64`

GetQuotaLimit returns the QuotaLimit field if non-nil, zero value otherwise.

### GetQuotaLimitOk

`func (o *StoragePureObjectBucket) GetQuotaLimitOk() (*int64, bool)`

GetQuotaLimitOk returns a tuple with the QuotaLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaLimit

`func (o *StoragePureObjectBucket) SetQuotaLimit(v int64)`

SetQuotaLimit sets QuotaLimit field to given value.

### HasQuotaLimit

`func (o *StoragePureObjectBucket) HasQuotaLimit() bool`

HasQuotaLimit returns a boolean if a field has been set.

### GetRetentionLock

`func (o *StoragePureObjectBucket) GetRetentionLock() string`

GetRetentionLock returns the RetentionLock field if non-nil, zero value otherwise.

### GetRetentionLockOk

`func (o *StoragePureObjectBucket) GetRetentionLockOk() (*string, bool)`

GetRetentionLockOk returns a tuple with the RetentionLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionLock

`func (o *StoragePureObjectBucket) SetRetentionLock(v string)`

SetRetentionLock sets RetentionLock field to given value.

### HasRetentionLock

`func (o *StoragePureObjectBucket) HasRetentionLock() bool`

HasRetentionLock returns a boolean if a field has been set.

### GetStorageUtilization

`func (o *StoragePureObjectBucket) GetStorageUtilization() StoragePureFlashBladeUtilization`

GetStorageUtilization returns the StorageUtilization field if non-nil, zero value otherwise.

### GetStorageUtilizationOk

`func (o *StoragePureObjectBucket) GetStorageUtilizationOk() (*StoragePureFlashBladeUtilization, bool)`

GetStorageUtilizationOk returns a tuple with the StorageUtilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageUtilization

`func (o *StoragePureObjectBucket) SetStorageUtilization(v StoragePureFlashBladeUtilization)`

SetStorageUtilization sets StorageUtilization field to given value.

### HasStorageUtilization

`func (o *StoragePureObjectBucket) HasStorageUtilization() bool`

HasStorageUtilization returns a boolean if a field has been set.

### GetUuid

`func (o *StoragePureObjectBucket) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *StoragePureObjectBucket) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *StoragePureObjectBucket) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *StoragePureObjectBucket) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetVendor

`func (o *StoragePureObjectBucket) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *StoragePureObjectBucket) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *StoragePureObjectBucket) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *StoragePureObjectBucket) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetVersion

`func (o *StoragePureObjectBucket) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *StoragePureObjectBucket) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *StoragePureObjectBucket) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *StoragePureObjectBucket) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetVersioning

`func (o *StoragePureObjectBucket) GetVersioning() string`

GetVersioning returns the Versioning field if non-nil, zero value otherwise.

### GetVersioningOk

`func (o *StoragePureObjectBucket) GetVersioningOk() (*string, bool)`

GetVersioningOk returns a tuple with the Versioning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersioning

`func (o *StoragePureObjectBucket) SetVersioning(v string)`

SetVersioning sets Versioning field to given value.

### HasVersioning

`func (o *StoragePureObjectBucket) HasVersioning() bool`

HasVersioning returns a boolean if a field has been set.

### GetArray

`func (o *StoragePureObjectBucket) GetArray() StoragePureFlashBladeSystemRelationship`

GetArray returns the Array field if non-nil, zero value otherwise.

### GetArrayOk

`func (o *StoragePureObjectBucket) GetArrayOk() (*StoragePureFlashBladeSystemRelationship, bool)`

GetArrayOk returns a tuple with the Array field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArray

`func (o *StoragePureObjectBucket) SetArray(v StoragePureFlashBladeSystemRelationship)`

SetArray sets Array field to given value.

### HasArray

`func (o *StoragePureObjectBucket) HasArray() bool`

HasArray returns a boolean if a field has been set.

### SetArrayNil

`func (o *StoragePureObjectBucket) SetArrayNil(b bool)`

 SetArrayNil sets the value for Array to be an explicit nil

### UnsetArray
`func (o *StoragePureObjectBucket) UnsetArray()`

UnsetArray ensures that no value is present for Array, not even an explicit nil
### GetObjectStoreAccount

`func (o *StoragePureObjectBucket) GetObjectStoreAccount() StoragePureObjectStoreAccountRelationship`

GetObjectStoreAccount returns the ObjectStoreAccount field if non-nil, zero value otherwise.

### GetObjectStoreAccountOk

`func (o *StoragePureObjectBucket) GetObjectStoreAccountOk() (*StoragePureObjectStoreAccountRelationship, bool)`

GetObjectStoreAccountOk returns a tuple with the ObjectStoreAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectStoreAccount

`func (o *StoragePureObjectBucket) SetObjectStoreAccount(v StoragePureObjectStoreAccountRelationship)`

SetObjectStoreAccount sets ObjectStoreAccount field to given value.

### HasObjectStoreAccount

`func (o *StoragePureObjectBucket) HasObjectStoreAccount() bool`

HasObjectStoreAccount returns a boolean if a field has been set.

### SetObjectStoreAccountNil

`func (o *StoragePureObjectBucket) SetObjectStoreAccountNil(b bool)`

 SetObjectStoreAccountNil sets the value for ObjectStoreAccount to be an explicit nil

### UnsetObjectStoreAccount
`func (o *StoragePureObjectBucket) UnsetObjectStoreAccount()`

UnsetObjectStoreAccount ensures that no value is present for ObjectStoreAccount, not even an explicit nil
### GetRegisteredDevice

`func (o *StoragePureObjectBucket) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StoragePureObjectBucket) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StoragePureObjectBucket) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StoragePureObjectBucket) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *StoragePureObjectBucket) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *StoragePureObjectBucket) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



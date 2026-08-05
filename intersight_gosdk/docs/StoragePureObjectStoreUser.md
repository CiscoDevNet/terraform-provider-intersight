# StoragePureObjectStoreUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.PureObjectStoreUser"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.PureObjectStoreUser"]
**Created** | Pointer to **time.Time** | Creation time of the user. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the object store user. | [optional] [readonly] 
**ObjectStoreAccountName** | Pointer to **string** | Name of the object store account that owns this user. | [optional] [readonly] 
**Uuid** | Pointer to **string** | A globally unique identifier for the object store user. | [optional] [readonly] 
**Array** | Pointer to [**NullableStoragePureFlashBladeSystemRelationship**](StoragePureFlashBladeSystemRelationship.md) |  | [optional] 
**ObjectStoreAccount** | Pointer to [**NullableStoragePureObjectStoreAccountRelationship**](StoragePureObjectStoreAccountRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewStoragePureObjectStoreUser

`func NewStoragePureObjectStoreUser(classId string, objectType string, ) *StoragePureObjectStoreUser`

NewStoragePureObjectStoreUser instantiates a new StoragePureObjectStoreUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoragePureObjectStoreUserWithDefaults

`func NewStoragePureObjectStoreUserWithDefaults() *StoragePureObjectStoreUser`

NewStoragePureObjectStoreUserWithDefaults instantiates a new StoragePureObjectStoreUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StoragePureObjectStoreUser) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StoragePureObjectStoreUser) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StoragePureObjectStoreUser) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StoragePureObjectStoreUser) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StoragePureObjectStoreUser) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StoragePureObjectStoreUser) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCreated

`func (o *StoragePureObjectStoreUser) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *StoragePureObjectStoreUser) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *StoragePureObjectStoreUser) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *StoragePureObjectStoreUser) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetName

`func (o *StoragePureObjectStoreUser) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StoragePureObjectStoreUser) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StoragePureObjectStoreUser) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StoragePureObjectStoreUser) HasName() bool`

HasName returns a boolean if a field has been set.

### GetObjectStoreAccountName

`func (o *StoragePureObjectStoreUser) GetObjectStoreAccountName() string`

GetObjectStoreAccountName returns the ObjectStoreAccountName field if non-nil, zero value otherwise.

### GetObjectStoreAccountNameOk

`func (o *StoragePureObjectStoreUser) GetObjectStoreAccountNameOk() (*string, bool)`

GetObjectStoreAccountNameOk returns a tuple with the ObjectStoreAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectStoreAccountName

`func (o *StoragePureObjectStoreUser) SetObjectStoreAccountName(v string)`

SetObjectStoreAccountName sets ObjectStoreAccountName field to given value.

### HasObjectStoreAccountName

`func (o *StoragePureObjectStoreUser) HasObjectStoreAccountName() bool`

HasObjectStoreAccountName returns a boolean if a field has been set.

### GetUuid

`func (o *StoragePureObjectStoreUser) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *StoragePureObjectStoreUser) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *StoragePureObjectStoreUser) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *StoragePureObjectStoreUser) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetArray

`func (o *StoragePureObjectStoreUser) GetArray() StoragePureFlashBladeSystemRelationship`

GetArray returns the Array field if non-nil, zero value otherwise.

### GetArrayOk

`func (o *StoragePureObjectStoreUser) GetArrayOk() (*StoragePureFlashBladeSystemRelationship, bool)`

GetArrayOk returns a tuple with the Array field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArray

`func (o *StoragePureObjectStoreUser) SetArray(v StoragePureFlashBladeSystemRelationship)`

SetArray sets Array field to given value.

### HasArray

`func (o *StoragePureObjectStoreUser) HasArray() bool`

HasArray returns a boolean if a field has been set.

### SetArrayNil

`func (o *StoragePureObjectStoreUser) SetArrayNil(b bool)`

 SetArrayNil sets the value for Array to be an explicit nil

### UnsetArray
`func (o *StoragePureObjectStoreUser) UnsetArray()`

UnsetArray ensures that no value is present for Array, not even an explicit nil
### GetObjectStoreAccount

`func (o *StoragePureObjectStoreUser) GetObjectStoreAccount() StoragePureObjectStoreAccountRelationship`

GetObjectStoreAccount returns the ObjectStoreAccount field if non-nil, zero value otherwise.

### GetObjectStoreAccountOk

`func (o *StoragePureObjectStoreUser) GetObjectStoreAccountOk() (*StoragePureObjectStoreAccountRelationship, bool)`

GetObjectStoreAccountOk returns a tuple with the ObjectStoreAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectStoreAccount

`func (o *StoragePureObjectStoreUser) SetObjectStoreAccount(v StoragePureObjectStoreAccountRelationship)`

SetObjectStoreAccount sets ObjectStoreAccount field to given value.

### HasObjectStoreAccount

`func (o *StoragePureObjectStoreUser) HasObjectStoreAccount() bool`

HasObjectStoreAccount returns a boolean if a field has been set.

### SetObjectStoreAccountNil

`func (o *StoragePureObjectStoreUser) SetObjectStoreAccountNil(b bool)`

 SetObjectStoreAccountNil sets the value for ObjectStoreAccount to be an explicit nil

### UnsetObjectStoreAccount
`func (o *StoragePureObjectStoreUser) UnsetObjectStoreAccount()`

UnsetObjectStoreAccount ensures that no value is present for ObjectStoreAccount, not even an explicit nil
### GetRegisteredDevice

`func (o *StoragePureObjectStoreUser) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StoragePureObjectStoreUser) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StoragePureObjectStoreUser) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StoragePureObjectStoreUser) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *StoragePureObjectStoreUser) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *StoragePureObjectStoreUser) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



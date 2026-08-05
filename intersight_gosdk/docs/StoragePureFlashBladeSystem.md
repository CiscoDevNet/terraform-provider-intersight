# StoragePureFlashBladeSystem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.PureFlashBladeSystem"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.PureFlashBladeSystem"]
**FileSystem** | Pointer to [**StoragePureFlashBladeUtilization**](StoragePureFlashBladeUtilization.md) |  | [optional] 
**ObjectStore** | Pointer to [**StoragePureFlashBladeUtilization**](StoragePureFlashBladeUtilization.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewStoragePureFlashBladeSystem

`func NewStoragePureFlashBladeSystem(classId string, objectType string, ) *StoragePureFlashBladeSystem`

NewStoragePureFlashBladeSystem instantiates a new StoragePureFlashBladeSystem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoragePureFlashBladeSystemWithDefaults

`func NewStoragePureFlashBladeSystemWithDefaults() *StoragePureFlashBladeSystem`

NewStoragePureFlashBladeSystemWithDefaults instantiates a new StoragePureFlashBladeSystem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StoragePureFlashBladeSystem) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StoragePureFlashBladeSystem) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StoragePureFlashBladeSystem) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StoragePureFlashBladeSystem) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StoragePureFlashBladeSystem) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StoragePureFlashBladeSystem) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetFileSystem

`func (o *StoragePureFlashBladeSystem) GetFileSystem() StoragePureFlashBladeUtilization`

GetFileSystem returns the FileSystem field if non-nil, zero value otherwise.

### GetFileSystemOk

`func (o *StoragePureFlashBladeSystem) GetFileSystemOk() (*StoragePureFlashBladeUtilization, bool)`

GetFileSystemOk returns a tuple with the FileSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSystem

`func (o *StoragePureFlashBladeSystem) SetFileSystem(v StoragePureFlashBladeUtilization)`

SetFileSystem sets FileSystem field to given value.

### HasFileSystem

`func (o *StoragePureFlashBladeSystem) HasFileSystem() bool`

HasFileSystem returns a boolean if a field has been set.

### GetObjectStore

`func (o *StoragePureFlashBladeSystem) GetObjectStore() StoragePureFlashBladeUtilization`

GetObjectStore returns the ObjectStore field if non-nil, zero value otherwise.

### GetObjectStoreOk

`func (o *StoragePureFlashBladeSystem) GetObjectStoreOk() (*StoragePureFlashBladeUtilization, bool)`

GetObjectStoreOk returns a tuple with the ObjectStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectStore

`func (o *StoragePureFlashBladeSystem) SetObjectStore(v StoragePureFlashBladeUtilization)`

SetObjectStore sets ObjectStore field to given value.

### HasObjectStore

`func (o *StoragePureFlashBladeSystem) HasObjectStore() bool`

HasObjectStore returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *StoragePureFlashBladeSystem) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StoragePureFlashBladeSystem) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StoragePureFlashBladeSystem) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StoragePureFlashBladeSystem) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *StoragePureFlashBladeSystem) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *StoragePureFlashBladeSystem) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



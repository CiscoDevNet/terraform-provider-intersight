# StorageFileItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.FileItem"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.FileItem"]
**Description** | Pointer to **string** | Description of the local Storage File. | [optional] [readonly] 
**FileId** | Pointer to **int64** | The Id of the local Storage File. | [optional] [readonly] 
**HostVisible** | Pointer to **bool** | The mapping visibility of the local Storage File. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the local Storage. | [optional] [readonly] 
**Size** | Pointer to **string** | Total size of the local Storage File. | [optional] [readonly] 
**Type** | Pointer to **string** | File type like CSV, ISO image. | [optional] [readonly] 
**UpdateTime** | Pointer to **string** | Timestamp to indicate the uploaded time for this file. | [optional] [readonly] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 
**StorageItem** | Pointer to [**NullableStorageItemRelationship**](StorageItemRelationship.md) |  | [optional] 

## Methods

### NewStorageFileItem

`func NewStorageFileItem(classId string, objectType string, ) *StorageFileItem`

NewStorageFileItem instantiates a new StorageFileItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageFileItemWithDefaults

`func NewStorageFileItemWithDefaults() *StorageFileItem`

NewStorageFileItemWithDefaults instantiates a new StorageFileItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageFileItem) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageFileItem) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageFileItem) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageFileItem) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageFileItem) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageFileItem) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDescription

`func (o *StorageFileItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StorageFileItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StorageFileItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *StorageFileItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFileId

`func (o *StorageFileItem) GetFileId() int64`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *StorageFileItem) GetFileIdOk() (*int64, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *StorageFileItem) SetFileId(v int64)`

SetFileId sets FileId field to given value.

### HasFileId

`func (o *StorageFileItem) HasFileId() bool`

HasFileId returns a boolean if a field has been set.

### GetHostVisible

`func (o *StorageFileItem) GetHostVisible() bool`

GetHostVisible returns the HostVisible field if non-nil, zero value otherwise.

### GetHostVisibleOk

`func (o *StorageFileItem) GetHostVisibleOk() (*bool, bool)`

GetHostVisibleOk returns a tuple with the HostVisible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostVisible

`func (o *StorageFileItem) SetHostVisible(v bool)`

SetHostVisible sets HostVisible field to given value.

### HasHostVisible

`func (o *StorageFileItem) HasHostVisible() bool`

HasHostVisible returns a boolean if a field has been set.

### GetName

`func (o *StorageFileItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageFileItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageFileItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StorageFileItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSize

`func (o *StorageFileItem) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *StorageFileItem) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *StorageFileItem) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *StorageFileItem) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetType

`func (o *StorageFileItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StorageFileItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StorageFileItem) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *StorageFileItem) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdateTime

`func (o *StorageFileItem) GetUpdateTime() string`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *StorageFileItem) GetUpdateTimeOk() (*string, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *StorageFileItem) SetUpdateTime(v string)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *StorageFileItem) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *StorageFileItem) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StorageFileItem) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StorageFileItem) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StorageFileItem) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *StorageFileItem) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *StorageFileItem) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
### GetStorageItem

`func (o *StorageFileItem) GetStorageItem() StorageItemRelationship`

GetStorageItem returns the StorageItem field if non-nil, zero value otherwise.

### GetStorageItemOk

`func (o *StorageFileItem) GetStorageItemOk() (*StorageItemRelationship, bool)`

GetStorageItemOk returns a tuple with the StorageItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageItem

`func (o *StorageFileItem) SetStorageItem(v StorageItemRelationship)`

SetStorageItem sets StorageItem field to given value.

### HasStorageItem

`func (o *StorageFileItem) HasStorageItem() bool`

HasStorageItem returns a boolean if a field has been set.

### SetStorageItemNil

`func (o *StorageFileItem) SetStorageItemNil(b bool)`

 SetStorageItemNil sets the value for StorageItem to be an explicit nil

### UnsetStorageItem
`func (o *StorageFileItem) UnsetStorageItem()`

UnsetStorageItem ensures that no value is present for StorageItem, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# StorageFileItemAllOf

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
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 
**StorageItem** | Pointer to [**StorageItemRelationship**](StorageItemRelationship.md) |  | [optional] 

## Methods

### NewStorageFileItemAllOf

`func NewStorageFileItemAllOf(classId string, objectType string, ) *StorageFileItemAllOf`

NewStorageFileItemAllOf instantiates a new StorageFileItemAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageFileItemAllOfWithDefaults

`func NewStorageFileItemAllOfWithDefaults() *StorageFileItemAllOf`

NewStorageFileItemAllOfWithDefaults instantiates a new StorageFileItemAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageFileItemAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageFileItemAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageFileItemAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageFileItemAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageFileItemAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageFileItemAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDescription

`func (o *StorageFileItemAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StorageFileItemAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StorageFileItemAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *StorageFileItemAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFileId

`func (o *StorageFileItemAllOf) GetFileId() int64`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *StorageFileItemAllOf) GetFileIdOk() (*int64, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *StorageFileItemAllOf) SetFileId(v int64)`

SetFileId sets FileId field to given value.

### HasFileId

`func (o *StorageFileItemAllOf) HasFileId() bool`

HasFileId returns a boolean if a field has been set.

### GetHostVisible

`func (o *StorageFileItemAllOf) GetHostVisible() bool`

GetHostVisible returns the HostVisible field if non-nil, zero value otherwise.

### GetHostVisibleOk

`func (o *StorageFileItemAllOf) GetHostVisibleOk() (*bool, bool)`

GetHostVisibleOk returns a tuple with the HostVisible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostVisible

`func (o *StorageFileItemAllOf) SetHostVisible(v bool)`

SetHostVisible sets HostVisible field to given value.

### HasHostVisible

`func (o *StorageFileItemAllOf) HasHostVisible() bool`

HasHostVisible returns a boolean if a field has been set.

### GetName

`func (o *StorageFileItemAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageFileItemAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageFileItemAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StorageFileItemAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSize

`func (o *StorageFileItemAllOf) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *StorageFileItemAllOf) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *StorageFileItemAllOf) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *StorageFileItemAllOf) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetType

`func (o *StorageFileItemAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StorageFileItemAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StorageFileItemAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *StorageFileItemAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdateTime

`func (o *StorageFileItemAllOf) GetUpdateTime() string`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *StorageFileItemAllOf) GetUpdateTimeOk() (*string, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *StorageFileItemAllOf) SetUpdateTime(v string)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *StorageFileItemAllOf) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *StorageFileItemAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StorageFileItemAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StorageFileItemAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StorageFileItemAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### GetStorageItem

`func (o *StorageFileItemAllOf) GetStorageItem() StorageItemRelationship`

GetStorageItem returns the StorageItem field if non-nil, zero value otherwise.

### GetStorageItemOk

`func (o *StorageFileItemAllOf) GetStorageItemOk() (*StorageItemRelationship, bool)`

GetStorageItemOk returns a tuple with the StorageItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageItem

`func (o *StorageFileItemAllOf) SetStorageItem(v StorageItemRelationship)`

SetStorageItem sets StorageItem field to given value.

### HasStorageItem

`func (o *StorageFileItemAllOf) HasStorageItem() bool`

HasStorageItem returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# StorageKeySetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.KeySetting"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.KeySetting"]
**RemoteKey** | Pointer to [**NullableStorageRemoteKeySetting**](StorageRemoteKeySetting.md) |  | [optional] 

## Methods

### NewStorageKeySetting

`func NewStorageKeySetting(classId string, objectType string, ) *StorageKeySetting`

NewStorageKeySetting instantiates a new StorageKeySetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageKeySettingWithDefaults

`func NewStorageKeySettingWithDefaults() *StorageKeySetting`

NewStorageKeySettingWithDefaults instantiates a new StorageKeySetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageKeySetting) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageKeySetting) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageKeySetting) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageKeySetting) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageKeySetting) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageKeySetting) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetRemoteKey

`func (o *StorageKeySetting) GetRemoteKey() StorageRemoteKeySetting`

GetRemoteKey returns the RemoteKey field if non-nil, zero value otherwise.

### GetRemoteKeyOk

`func (o *StorageKeySetting) GetRemoteKeyOk() (*StorageRemoteKeySetting, bool)`

GetRemoteKeyOk returns a tuple with the RemoteKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteKey

`func (o *StorageKeySetting) SetRemoteKey(v StorageRemoteKeySetting)`

SetRemoteKey sets RemoteKey field to given value.

### HasRemoteKey

`func (o *StorageKeySetting) HasRemoteKey() bool`

HasRemoteKey returns a boolean if a field has been set.

### SetRemoteKeyNil

`func (o *StorageKeySetting) SetRemoteKeyNil(b bool)`

 SetRemoteKeyNil sets the value for RemoteKey to be an explicit nil

### UnsetRemoteKey
`func (o *StorageKeySetting) UnsetRemoteKey()`

UnsetRemoteKey ensures that no value is present for RemoteKey, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



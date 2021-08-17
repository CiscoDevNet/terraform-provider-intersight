# StorageKeySettingAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.KeySetting"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.KeySetting"]
**KeyType** | Pointer to **string** | Method to be used for fetching the encryption key. * &#x60;None&#x60; - Drive encryption not configured. * &#x60;Manual&#x60; - Drive encryption using manual key. * &#x60;Kmip&#x60; - Remote encryption using KMIP. | [optional] [default to "None"]
**ManualKey** | Pointer to [**NullableStorageLocalKeySetting**](storage.LocalKeySetting.md) |  | [optional] 
**RemoteKey** | Pointer to [**NullableStorageRemoteKeySetting**](storage.RemoteKeySetting.md) |  | [optional] 

## Methods

### NewStorageKeySettingAllOf

`func NewStorageKeySettingAllOf(classId string, objectType string, ) *StorageKeySettingAllOf`

NewStorageKeySettingAllOf instantiates a new StorageKeySettingAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageKeySettingAllOfWithDefaults

`func NewStorageKeySettingAllOfWithDefaults() *StorageKeySettingAllOf`

NewStorageKeySettingAllOfWithDefaults instantiates a new StorageKeySettingAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageKeySettingAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageKeySettingAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageKeySettingAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageKeySettingAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageKeySettingAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageKeySettingAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetKeyType

`func (o *StorageKeySettingAllOf) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *StorageKeySettingAllOf) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *StorageKeySettingAllOf) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.

### HasKeyType

`func (o *StorageKeySettingAllOf) HasKeyType() bool`

HasKeyType returns a boolean if a field has been set.

### GetManualKey

`func (o *StorageKeySettingAllOf) GetManualKey() StorageLocalKeySetting`

GetManualKey returns the ManualKey field if non-nil, zero value otherwise.

### GetManualKeyOk

`func (o *StorageKeySettingAllOf) GetManualKeyOk() (*StorageLocalKeySetting, bool)`

GetManualKeyOk returns a tuple with the ManualKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManualKey

`func (o *StorageKeySettingAllOf) SetManualKey(v StorageLocalKeySetting)`

SetManualKey sets ManualKey field to given value.

### HasManualKey

`func (o *StorageKeySettingAllOf) HasManualKey() bool`

HasManualKey returns a boolean if a field has been set.

### SetManualKeyNil

`func (o *StorageKeySettingAllOf) SetManualKeyNil(b bool)`

 SetManualKeyNil sets the value for ManualKey to be an explicit nil

### UnsetManualKey
`func (o *StorageKeySettingAllOf) UnsetManualKey()`

UnsetManualKey ensures that no value is present for ManualKey, not even an explicit nil
### GetRemoteKey

`func (o *StorageKeySettingAllOf) GetRemoteKey() StorageRemoteKeySetting`

GetRemoteKey returns the RemoteKey field if non-nil, zero value otherwise.

### GetRemoteKeyOk

`func (o *StorageKeySettingAllOf) GetRemoteKeyOk() (*StorageRemoteKeySetting, bool)`

GetRemoteKeyOk returns a tuple with the RemoteKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteKey

`func (o *StorageKeySettingAllOf) SetRemoteKey(v StorageRemoteKeySetting)`

SetRemoteKey sets RemoteKey field to given value.

### HasRemoteKey

`func (o *StorageKeySettingAllOf) HasRemoteKey() bool`

HasRemoteKey returns a boolean if a field has been set.

### SetRemoteKeyNil

`func (o *StorageKeySettingAllOf) SetRemoteKeyNil(b bool)`

 SetRemoteKeyNil sets the value for RemoteKey to be an explicit nil

### UnsetRemoteKey
`func (o *StorageKeySettingAllOf) UnsetRemoteKey()`

UnsetRemoteKey ensures that no value is present for RemoteKey, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



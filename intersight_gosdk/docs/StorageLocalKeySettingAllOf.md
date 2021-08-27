# StorageLocalKeySettingAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.LocalKeySetting"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.LocalKeySetting"]
**ExistingKey** | Pointer to **string** | Existing key which is already configured on the server. | [optional] 
**IsExistingKeySet** | Pointer to **bool** | Indicates whether the value of the &#39;existingKey&#39; property has been set. | [optional] [readonly] [default to false]
**IsNewKeySet** | Pointer to **bool** | Indicates whether the value of the &#39;newKey&#39; property has been set. | [optional] [readonly] [default to false]
**NewKey** | Pointer to **string** | New key to be configured on the controller. | [optional] 

## Methods

### NewStorageLocalKeySettingAllOf

`func NewStorageLocalKeySettingAllOf(classId string, objectType string, ) *StorageLocalKeySettingAllOf`

NewStorageLocalKeySettingAllOf instantiates a new StorageLocalKeySettingAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageLocalKeySettingAllOfWithDefaults

`func NewStorageLocalKeySettingAllOfWithDefaults() *StorageLocalKeySettingAllOf`

NewStorageLocalKeySettingAllOfWithDefaults instantiates a new StorageLocalKeySettingAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageLocalKeySettingAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageLocalKeySettingAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageLocalKeySettingAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageLocalKeySettingAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageLocalKeySettingAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageLocalKeySettingAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetExistingKey

`func (o *StorageLocalKeySettingAllOf) GetExistingKey() string`

GetExistingKey returns the ExistingKey field if non-nil, zero value otherwise.

### GetExistingKeyOk

`func (o *StorageLocalKeySettingAllOf) GetExistingKeyOk() (*string, bool)`

GetExistingKeyOk returns a tuple with the ExistingKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistingKey

`func (o *StorageLocalKeySettingAllOf) SetExistingKey(v string)`

SetExistingKey sets ExistingKey field to given value.

### HasExistingKey

`func (o *StorageLocalKeySettingAllOf) HasExistingKey() bool`

HasExistingKey returns a boolean if a field has been set.

### GetIsExistingKeySet

`func (o *StorageLocalKeySettingAllOf) GetIsExistingKeySet() bool`

GetIsExistingKeySet returns the IsExistingKeySet field if non-nil, zero value otherwise.

### GetIsExistingKeySetOk

`func (o *StorageLocalKeySettingAllOf) GetIsExistingKeySetOk() (*bool, bool)`

GetIsExistingKeySetOk returns a tuple with the IsExistingKeySet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExistingKeySet

`func (o *StorageLocalKeySettingAllOf) SetIsExistingKeySet(v bool)`

SetIsExistingKeySet sets IsExistingKeySet field to given value.

### HasIsExistingKeySet

`func (o *StorageLocalKeySettingAllOf) HasIsExistingKeySet() bool`

HasIsExistingKeySet returns a boolean if a field has been set.

### GetIsNewKeySet

`func (o *StorageLocalKeySettingAllOf) GetIsNewKeySet() bool`

GetIsNewKeySet returns the IsNewKeySet field if non-nil, zero value otherwise.

### GetIsNewKeySetOk

`func (o *StorageLocalKeySettingAllOf) GetIsNewKeySetOk() (*bool, bool)`

GetIsNewKeySetOk returns a tuple with the IsNewKeySet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNewKeySet

`func (o *StorageLocalKeySettingAllOf) SetIsNewKeySet(v bool)`

SetIsNewKeySet sets IsNewKeySet field to given value.

### HasIsNewKeySet

`func (o *StorageLocalKeySettingAllOf) HasIsNewKeySet() bool`

HasIsNewKeySet returns a boolean if a field has been set.

### GetNewKey

`func (o *StorageLocalKeySettingAllOf) GetNewKey() string`

GetNewKey returns the NewKey field if non-nil, zero value otherwise.

### GetNewKeyOk

`func (o *StorageLocalKeySettingAllOf) GetNewKeyOk() (*string, bool)`

GetNewKeyOk returns a tuple with the NewKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewKey

`func (o *StorageLocalKeySettingAllOf) SetNewKey(v string)`

SetNewKey sets NewKey field to given value.

### HasNewKey

`func (o *StorageLocalKeySettingAllOf) HasNewKey() bool`

HasNewKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



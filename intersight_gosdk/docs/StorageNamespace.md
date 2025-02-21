# StorageNamespace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.Namespace"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.Namespace"]
**BlockCapacity** | Pointer to **int64** | Number of blocks of the namespace. | [optional] [readonly] 
**ByteFormatCapacity** | Pointer to **string** | Capacity of the namespace. | [optional] [readonly] 
**LdevId** | Pointer to **int64** | LDEV number of the volume where the namespace is set. | [optional] [readonly] 
**NamespaceId** | Pointer to **int64** | ID of namespace created in the NVM subsystem. | [optional] [readonly] 
**NamespaceNickname** | Pointer to **string** | Nickname of the namespace. | [optional] [readonly] 

## Methods

### NewStorageNamespace

`func NewStorageNamespace(classId string, objectType string, ) *StorageNamespace`

NewStorageNamespace instantiates a new StorageNamespace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNamespaceWithDefaults

`func NewStorageNamespaceWithDefaults() *StorageNamespace`

NewStorageNamespaceWithDefaults instantiates a new StorageNamespace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageNamespace) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageNamespace) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageNamespace) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageNamespace) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageNamespace) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageNamespace) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBlockCapacity

`func (o *StorageNamespace) GetBlockCapacity() int64`

GetBlockCapacity returns the BlockCapacity field if non-nil, zero value otherwise.

### GetBlockCapacityOk

`func (o *StorageNamespace) GetBlockCapacityOk() (*int64, bool)`

GetBlockCapacityOk returns a tuple with the BlockCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockCapacity

`func (o *StorageNamespace) SetBlockCapacity(v int64)`

SetBlockCapacity sets BlockCapacity field to given value.

### HasBlockCapacity

`func (o *StorageNamespace) HasBlockCapacity() bool`

HasBlockCapacity returns a boolean if a field has been set.

### GetByteFormatCapacity

`func (o *StorageNamespace) GetByteFormatCapacity() string`

GetByteFormatCapacity returns the ByteFormatCapacity field if non-nil, zero value otherwise.

### GetByteFormatCapacityOk

`func (o *StorageNamespace) GetByteFormatCapacityOk() (*string, bool)`

GetByteFormatCapacityOk returns a tuple with the ByteFormatCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByteFormatCapacity

`func (o *StorageNamespace) SetByteFormatCapacity(v string)`

SetByteFormatCapacity sets ByteFormatCapacity field to given value.

### HasByteFormatCapacity

`func (o *StorageNamespace) HasByteFormatCapacity() bool`

HasByteFormatCapacity returns a boolean if a field has been set.

### GetLdevId

`func (o *StorageNamespace) GetLdevId() int64`

GetLdevId returns the LdevId field if non-nil, zero value otherwise.

### GetLdevIdOk

`func (o *StorageNamespace) GetLdevIdOk() (*int64, bool)`

GetLdevIdOk returns a tuple with the LdevId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdevId

`func (o *StorageNamespace) SetLdevId(v int64)`

SetLdevId sets LdevId field to given value.

### HasLdevId

`func (o *StorageNamespace) HasLdevId() bool`

HasLdevId returns a boolean if a field has been set.

### GetNamespaceId

`func (o *StorageNamespace) GetNamespaceId() int64`

GetNamespaceId returns the NamespaceId field if non-nil, zero value otherwise.

### GetNamespaceIdOk

`func (o *StorageNamespace) GetNamespaceIdOk() (*int64, bool)`

GetNamespaceIdOk returns a tuple with the NamespaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceId

`func (o *StorageNamespace) SetNamespaceId(v int64)`

SetNamespaceId sets NamespaceId field to given value.

### HasNamespaceId

`func (o *StorageNamespace) HasNamespaceId() bool`

HasNamespaceId returns a boolean if a field has been set.

### GetNamespaceNickname

`func (o *StorageNamespace) GetNamespaceNickname() string`

GetNamespaceNickname returns the NamespaceNickname field if non-nil, zero value otherwise.

### GetNamespaceNicknameOk

`func (o *StorageNamespace) GetNamespaceNicknameOk() (*string, bool)`

GetNamespaceNicknameOk returns a tuple with the NamespaceNickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceNickname

`func (o *StorageNamespace) SetNamespaceNickname(v string)`

SetNamespaceNickname sets NamespaceNickname field to given value.

### HasNamespaceNickname

`func (o *StorageNamespace) HasNamespaceNickname() bool`

HasNamespaceNickname returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



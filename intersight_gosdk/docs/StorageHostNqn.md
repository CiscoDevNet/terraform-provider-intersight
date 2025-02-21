# StorageHostNqn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.HostNqn"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.HostNqn"]
**HostNqn** | Pointer to **string** | Host NQN registered in the NVM subsystem. | [optional] [readonly] 
**HostNqnNickname** | Pointer to **string** | Nickname of the Host NQN. | [optional] [readonly] 

## Methods

### NewStorageHostNqn

`func NewStorageHostNqn(classId string, objectType string, ) *StorageHostNqn`

NewStorageHostNqn instantiates a new StorageHostNqn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageHostNqnWithDefaults

`func NewStorageHostNqnWithDefaults() *StorageHostNqn`

NewStorageHostNqnWithDefaults instantiates a new StorageHostNqn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageHostNqn) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageHostNqn) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageHostNqn) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageHostNqn) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageHostNqn) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageHostNqn) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetHostNqn

`func (o *StorageHostNqn) GetHostNqn() string`

GetHostNqn returns the HostNqn field if non-nil, zero value otherwise.

### GetHostNqnOk

`func (o *StorageHostNqn) GetHostNqnOk() (*string, bool)`

GetHostNqnOk returns a tuple with the HostNqn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostNqn

`func (o *StorageHostNqn) SetHostNqn(v string)`

SetHostNqn sets HostNqn field to given value.

### HasHostNqn

`func (o *StorageHostNqn) HasHostNqn() bool`

HasHostNqn returns a boolean if a field has been set.

### GetHostNqnNickname

`func (o *StorageHostNqn) GetHostNqnNickname() string`

GetHostNqnNickname returns the HostNqnNickname field if non-nil, zero value otherwise.

### GetHostNqnNicknameOk

`func (o *StorageHostNqn) GetHostNqnNicknameOk() (*string, bool)`

GetHostNqnNicknameOk returns a tuple with the HostNqnNickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostNqnNickname

`func (o *StorageHostNqn) SetHostNqnNickname(v string)`

SetHostNqnNickname sets HostNqnNickname field to given value.

### HasHostNqnNickname

`func (o *StorageHostNqn) HasHostNqnNickname() bool`

HasHostNqnNickname returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



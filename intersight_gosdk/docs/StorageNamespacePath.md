# StorageNamespacePath

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.NamespacePath"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.NamespacePath"]
**HostNqn** | Pointer to **string** | Host NQN registered in the NVM subsystem. | [optional] [readonly] 
**LdevId** | Pointer to **int64** | LDEV number of the volume where the namespace is set. | [optional] [readonly] 
**NamespaceId** | Pointer to **int64** | ID of namespace created in the NVM subsystem. | [optional] [readonly] 

## Methods

### NewStorageNamespacePath

`func NewStorageNamespacePath(classId string, objectType string, ) *StorageNamespacePath`

NewStorageNamespacePath instantiates a new StorageNamespacePath object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNamespacePathWithDefaults

`func NewStorageNamespacePathWithDefaults() *StorageNamespacePath`

NewStorageNamespacePathWithDefaults instantiates a new StorageNamespacePath object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageNamespacePath) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageNamespacePath) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageNamespacePath) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageNamespacePath) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageNamespacePath) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageNamespacePath) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetHostNqn

`func (o *StorageNamespacePath) GetHostNqn() string`

GetHostNqn returns the HostNqn field if non-nil, zero value otherwise.

### GetHostNqnOk

`func (o *StorageNamespacePath) GetHostNqnOk() (*string, bool)`

GetHostNqnOk returns a tuple with the HostNqn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostNqn

`func (o *StorageNamespacePath) SetHostNqn(v string)`

SetHostNqn sets HostNqn field to given value.

### HasHostNqn

`func (o *StorageNamespacePath) HasHostNqn() bool`

HasHostNqn returns a boolean if a field has been set.

### GetLdevId

`func (o *StorageNamespacePath) GetLdevId() int64`

GetLdevId returns the LdevId field if non-nil, zero value otherwise.

### GetLdevIdOk

`func (o *StorageNamespacePath) GetLdevIdOk() (*int64, bool)`

GetLdevIdOk returns a tuple with the LdevId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdevId

`func (o *StorageNamespacePath) SetLdevId(v int64)`

SetLdevId sets LdevId field to given value.

### HasLdevId

`func (o *StorageNamespacePath) HasLdevId() bool`

HasLdevId returns a boolean if a field has been set.

### GetNamespaceId

`func (o *StorageNamespacePath) GetNamespaceId() int64`

GetNamespaceId returns the NamespaceId field if non-nil, zero value otherwise.

### GetNamespaceIdOk

`func (o *StorageNamespacePath) GetNamespaceIdOk() (*int64, bool)`

GetNamespaceIdOk returns a tuple with the NamespaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceId

`func (o *StorageNamespacePath) SetNamespaceId(v int64)`

SetNamespaceId sets NamespaceId field to given value.

### HasNamespaceId

`func (o *StorageNamespacePath) HasNamespaceId() bool`

HasNamespaceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



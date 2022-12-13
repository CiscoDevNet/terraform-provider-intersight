# StorageNetAppCifsAclAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.NetAppCifsAcl"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.NetAppCifsAcl"]
**Permission** | Pointer to **string** | Access rights that a user or group has on the defined CIFS share. | [optional] [readonly] 
**Type** | Pointer to **string** | The type of operating system for the CIFS share. | [optional] [readonly] 
**UserOrGroup** | Pointer to **string** | User or group name to add to the access control list of a CIFS share. | [optional] [readonly] 

## Methods

### NewStorageNetAppCifsAclAllOf

`func NewStorageNetAppCifsAclAllOf(classId string, objectType string, ) *StorageNetAppCifsAclAllOf`

NewStorageNetAppCifsAclAllOf instantiates a new StorageNetAppCifsAclAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNetAppCifsAclAllOfWithDefaults

`func NewStorageNetAppCifsAclAllOfWithDefaults() *StorageNetAppCifsAclAllOf`

NewStorageNetAppCifsAclAllOfWithDefaults instantiates a new StorageNetAppCifsAclAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageNetAppCifsAclAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageNetAppCifsAclAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageNetAppCifsAclAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageNetAppCifsAclAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageNetAppCifsAclAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageNetAppCifsAclAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetPermission

`func (o *StorageNetAppCifsAclAllOf) GetPermission() string`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *StorageNetAppCifsAclAllOf) GetPermissionOk() (*string, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *StorageNetAppCifsAclAllOf) SetPermission(v string)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *StorageNetAppCifsAclAllOf) HasPermission() bool`

HasPermission returns a boolean if a field has been set.

### GetType

`func (o *StorageNetAppCifsAclAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StorageNetAppCifsAclAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StorageNetAppCifsAclAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *StorageNetAppCifsAclAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUserOrGroup

`func (o *StorageNetAppCifsAclAllOf) GetUserOrGroup() string`

GetUserOrGroup returns the UserOrGroup field if non-nil, zero value otherwise.

### GetUserOrGroupOk

`func (o *StorageNetAppCifsAclAllOf) GetUserOrGroupOk() (*string, bool)`

GetUserOrGroupOk returns a tuple with the UserOrGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserOrGroup

`func (o *StorageNetAppCifsAclAllOf) SetUserOrGroup(v string)`

SetUserOrGroup sets UserOrGroup field to given value.

### HasUserOrGroup

`func (o *StorageNetAppCifsAclAllOf) HasUserOrGroup() bool`

HasUserOrGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



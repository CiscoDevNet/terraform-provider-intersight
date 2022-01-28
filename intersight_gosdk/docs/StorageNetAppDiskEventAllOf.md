# StorageNetAppDiskEventAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.NetAppDiskEvent"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.NetAppDiskEvent"]
**Disk** | Pointer to [**StorageNetAppBaseDiskRelationship**](StorageNetAppBaseDiskRelationship.md) |  | [optional] 

## Methods

### NewStorageNetAppDiskEventAllOf

`func NewStorageNetAppDiskEventAllOf(classId string, objectType string, ) *StorageNetAppDiskEventAllOf`

NewStorageNetAppDiskEventAllOf instantiates a new StorageNetAppDiskEventAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNetAppDiskEventAllOfWithDefaults

`func NewStorageNetAppDiskEventAllOfWithDefaults() *StorageNetAppDiskEventAllOf`

NewStorageNetAppDiskEventAllOfWithDefaults instantiates a new StorageNetAppDiskEventAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageNetAppDiskEventAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageNetAppDiskEventAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageNetAppDiskEventAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageNetAppDiskEventAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageNetAppDiskEventAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageNetAppDiskEventAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDisk

`func (o *StorageNetAppDiskEventAllOf) GetDisk() StorageNetAppBaseDiskRelationship`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *StorageNetAppDiskEventAllOf) GetDiskOk() (*StorageNetAppBaseDiskRelationship, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *StorageNetAppDiskEventAllOf) SetDisk(v StorageNetAppBaseDiskRelationship)`

SetDisk sets Disk field to given value.

### HasDisk

`func (o *StorageNetAppDiskEventAllOf) HasDisk() bool`

HasDisk returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# StorageNetAppDiskEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.NetAppDiskEvent"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.NetAppDiskEvent"]
**Disk** | Pointer to [**NullableStorageNetAppBaseDiskRelationship**](StorageNetAppBaseDiskRelationship.md) |  | [optional] 

## Methods

### NewStorageNetAppDiskEvent

`func NewStorageNetAppDiskEvent(classId string, objectType string, ) *StorageNetAppDiskEvent`

NewStorageNetAppDiskEvent instantiates a new StorageNetAppDiskEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNetAppDiskEventWithDefaults

`func NewStorageNetAppDiskEventWithDefaults() *StorageNetAppDiskEvent`

NewStorageNetAppDiskEventWithDefaults instantiates a new StorageNetAppDiskEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageNetAppDiskEvent) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageNetAppDiskEvent) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageNetAppDiskEvent) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageNetAppDiskEvent) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageNetAppDiskEvent) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageNetAppDiskEvent) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDisk

`func (o *StorageNetAppDiskEvent) GetDisk() StorageNetAppBaseDiskRelationship`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *StorageNetAppDiskEvent) GetDiskOk() (*StorageNetAppBaseDiskRelationship, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *StorageNetAppDiskEvent) SetDisk(v StorageNetAppBaseDiskRelationship)`

SetDisk sets Disk field to given value.

### HasDisk

`func (o *StorageNetAppDiskEvent) HasDisk() bool`

HasDisk returns a boolean if a field has been set.

### SetDiskNil

`func (o *StorageNetAppDiskEvent) SetDiskNil(b bool)`

 SetDiskNil sets the value for Disk to be an explicit nil

### UnsetDisk
`func (o *StorageNetAppDiskEvent) UnsetDisk()`

UnsetDisk ensures that no value is present for Disk, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



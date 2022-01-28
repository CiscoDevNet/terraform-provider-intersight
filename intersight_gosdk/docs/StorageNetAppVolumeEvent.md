# StorageNetAppVolumeEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.NetAppVolumeEvent"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.NetAppVolumeEvent"]
**Volume** | Pointer to [**StorageNetAppVolumeRelationship**](StorageNetAppVolumeRelationship.md) |  | [optional] 

## Methods

### NewStorageNetAppVolumeEvent

`func NewStorageNetAppVolumeEvent(classId string, objectType string, ) *StorageNetAppVolumeEvent`

NewStorageNetAppVolumeEvent instantiates a new StorageNetAppVolumeEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNetAppVolumeEventWithDefaults

`func NewStorageNetAppVolumeEventWithDefaults() *StorageNetAppVolumeEvent`

NewStorageNetAppVolumeEventWithDefaults instantiates a new StorageNetAppVolumeEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageNetAppVolumeEvent) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageNetAppVolumeEvent) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageNetAppVolumeEvent) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageNetAppVolumeEvent) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageNetAppVolumeEvent) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageNetAppVolumeEvent) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetVolume

`func (o *StorageNetAppVolumeEvent) GetVolume() StorageNetAppVolumeRelationship`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *StorageNetAppVolumeEvent) GetVolumeOk() (*StorageNetAppVolumeRelationship, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *StorageNetAppVolumeEvent) SetVolume(v StorageNetAppVolumeRelationship)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *StorageNetAppVolumeEvent) HasVolume() bool`

HasVolume returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


